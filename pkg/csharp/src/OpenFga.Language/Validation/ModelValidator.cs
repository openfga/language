using OpenFga.Language.Errors;
using OpenFga.Sdk.Model;
using Metadata = OpenFga.Sdk.Model.Metadata;

namespace OpenFga.Language.Validation;

public class ModelValidator {
    private readonly ValidationOptions _options;
    private readonly AuthorizationModel _authorizationModel;
    private readonly Dsl _dsl;
    private readonly ValidationErrorsBuilder _errors;
    private ValidationRegex? _typeRegex;
    private ValidationRegex? _relationRegex;
    private readonly Dictionary<string, HashSet<string>> _fileToModules = new();

    public ModelValidator(ValidationOptions options, AuthorizationModel authorizationModel, string[]? lines) {
        _options = options;
        _authorizationModel = authorizationModel;
        _dsl = new Dsl(lines);
        _errors = new ValidationErrorsBuilder(lines);
    }

    public static void ValidateJson(AuthorizationModel authorizationModel) {
        ValidateJson(authorizationModel, new ValidationOptions());
    }

    public static void ValidateJson(AuthorizationModel authorizationModel, ValidationOptions options) {
        new ModelValidator(options, authorizationModel, null).Validate();
    }

    public static void ValidateDsl(string dsl) {
        ValidateDsl(dsl, new ValidationOptions());
    }

    public static void ValidateDsl(string dsl, ValidationOptions options) {
        var transformer = new DslToJsonTransformer();
        var result = transformer.ParseDsl(dsl);
        if (result.IsFailure()) {
            throw new DslErrorsException(result.Errors);
        }
        var authorizationModel = result.AuthorizationModel;
        var lines = dsl.Split('\n');

        new ModelValidator(options, authorizationModel, lines).Validate();
    }

    private void Validate() {
        _typeRegex = ValidationRegex.Build("type", _options.TypePattern);
        _relationRegex = ValidationRegex.Build("relation", _options.RelationPattern);

        PopulateRelations();

        var schemaVersion = _authorizationModel.SchemaVersion;
        if (schemaVersion == null) {
            _errors.RaiseSchemaVersionRequired(0, "");
        }

        if (schemaVersion != null && (schemaVersion.Equals("1.1") || schemaVersion.Equals("1.2"))) {
            ModelValidation();
        }
        else if (schemaVersion != null) {
            var lineIndex = _dsl.GetSchemaLineNumber(schemaVersion);
            _errors.RaiseInvalidSchemaVersion(lineIndex, schemaVersion);
        }

        foreach (var entry in _fileToModules) {
            var file = entry.Key;
            var modules = entry.Value;
            if (modules.Count > 1) {
                _errors.RaiseMultipleModulesInSingleFile(file, modules);
            }
        }

        _errors.ThrowIfNotEmpty();
    }

    private void PopulateRelations() {
        foreach (var typeDef in _authorizationModel.TypeDefinitions ?? []) {
            var typeName = typeDef.Type;

            TrackModulesInFile(typeDef.Metadata);

            if (typeName.Equals(Keyword.Self) || typeName.Equals(Keyword.This)) {
                var lineIndex = _dsl.GetTypeLineNumber(typeName);
                _errors.RaiseReservedTypeName(lineIndex, typeName);
            }

            if (!_typeRegex!.Matches(typeName)) {
                var lineIndex = _dsl.GetTypeLineNumber(typeName);
                _errors.RaiseInvalidName(lineIndex, typeName, _typeRegex.Rule);
            }

            var encounteredRelationsInType = new HashSet<string> { Keyword.Self };

            foreach (var kvp in typeDef.Relations ?? []) {
                var relationName = kvp.Key;
                var relation = kvp.Value;

                if (relationName.Equals(Keyword.Self) || relationName.Equals(Keyword.This)) {
                    var typeIndex = _dsl.GetTypeLineNumber(typeName);
                    var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                    _errors.RaiseReservedRelationName(lineIndex, relationName);
                }
                else if (!_relationRegex!.Matches(relationName)) {
                    var typeIndex = _dsl.GetTypeLineNumber(typeName);
                    var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                    _errors.RaiseInvalidName(lineIndex, relationName, _relationRegex.Rule, typeName);
                }
                else if (encounteredRelationsInType.Contains(relationName)) {
                    var typeIndex = _dsl.GetTypeLineNumber(typeName);
                    var initialLineIdx = _dsl.GetRelationLineNumber(relationName, typeIndex);
                    var duplicateLineIdx = _dsl.GetRelationLineNumber(relationName, initialLineIdx + 1);
                    _errors.RaiseDuplicateRelationName(duplicateLineIdx, relationName);
                }
                encounteredRelationsInType.Add(relationName);
            }
        }
    }

    private void ModelValidation() {
        if (!_errors.IsEmpty) {
            return;
        }

        var typeMap = new Dictionary<string, TypeDefinition>();
        var usedConditionNamesSet = new HashSet<string>();

        foreach (var typeDef in _authorizationModel.TypeDefinitions ?? []) {
            var typeName = typeDef.Type;
            typeMap[typeName] = typeDef;

            if (typeDef.Metadata != null) {
                foreach (var kvp in typeDef.Metadata.Relations ?? []) {
                    var relationMetadata = kvp.Value;
                    foreach (var typeRestriction in relationMetadata.DirectlyRelatedUserTypes ?? []) {
                        if (typeRestriction.Condition != null) {
                            usedConditionNamesSet.Add(typeRestriction.Condition);
                        }
                    }
                }
            }
        }

        // first, validate to ensure all the relation are defined
        foreach (var typeDef in _authorizationModel.TypeDefinitions ?? []) {
            var typeName = typeDef.Type;
            foreach (var kvp in typeDef.Relations ?? []) {
                var relationName = kvp.Key;
                RelationDefined(typeMap, typeName, relationName);
            }
        }

        if (_errors.IsEmpty) {
            var typeSet = new HashSet<string>();
            foreach (var typeDef in _authorizationModel.TypeDefinitions ?? []) {
                var typeName = typeDef.Type;
                if (typeSet.Contains(typeName)) {
                    var typeIndex = _dsl.GetTypeLineNumber(typeName);
                    _errors.RaiseDuplicateTypeName(typeIndex, typeName);
                }
                typeSet.Add(typeName);

                if (typeDef.Metadata != null) {
                    foreach (var relationDefKey in (ICollection<string>?)typeDef.Metadata.Relations?.Keys ?? []) {
                        CheckForDuplicatesTypeNamesInRelation(typeDef.Metadata.Relations![relationDefKey], relationDefKey);
                        CheckForDuplicatesInRelation(typeDef, relationDefKey);
                        CheckForInvalidOrderInRelation(typeDef, relationDefKey);
                    }
                }
            }
        }

        // next, ensure all relation have entry point
        // we can skip if there are errors because errors (such as missing relations)
        // will likely lead to no entries
        if (_errors.IsEmpty) {
            foreach (var typeDef in _authorizationModel.TypeDefinitions ?? []) {
                var typeName = typeDef.Type;
                var currentRelations = typeMap[typeName].Relations!;
                var typeDefMetadata = typeDef.Metadata;
                var typeDefRelationsMetadata = typeMap[typeName].Metadata?.Relations;

                foreach (var relationName in typeDef.Relations!.Keys) {
                    var result = EntryPointOrLoop.Compute(
                        typeMap, typeName, relationName, currentRelations[relationName], new Dictionary<string, Dictionary<string, bool>>());

                    TrackModulesInFile(typeDefMetadata, typeDefRelationsMetadata?.GetValueOrDefault(relationName));

                    if (!result.HasEntry()) {
                        var typeIndex = _dsl.GetTypeLineNumber(typeName);
                        var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                        if (result.IsLoop()) {
                            _errors.RaiseNoEntryPointLoop(lineIndex, relationName, typeName);
                        }
                        else {
                            _errors.RaiseNoEntryPoint(lineIndex, relationName, typeName);
                        }
                    }
                }
            }
        }

        foreach (var kvp in _authorizationModel.Conditions ?? []) {
            var conditionName = kvp.Key;
            var condition = kvp.Value;

            TrackModulesInFile(condition.Metadata);

            if (!conditionName.Equals(condition.Name)) {
                _errors.RaiseDifferentNestedConditionName(conditionName, condition.Name);
            }

            if (!usedConditionNamesSet.Contains(conditionName)) {
                var conditionIndex = _dsl.GetConditionLineNumber(conditionName);
                _errors.RaiseUnusedCondition(conditionIndex, conditionName);
            }
        }
    }

    private void CheckForInvalidOrderInRelation(TypeDefinition typeDef, string relationName) {
        var relationDef = typeDef.Relations![relationName];
        var metadataRelation = typeDef.Metadata!.Relations![relationName];

        if (metadataRelation.DirectlyRelatedUserTypes!.Count == 0)
            return;

        // Union
        if (relationDef.Union?.Child.Count > 1 && relationDef.Union.Child[0].This is not { }) {
            var typeIndex = _dsl.GetTypeLineNumber(typeDef.Type);
            var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
            _errors.RaiseThisNotInFirstPlace(lineIndex, relationName);
        }

        // TODO: Should this also be the case for Intersection and Difference?
    }

    private void CheckForDuplicatesInRelation(TypeDefinition typeDef, string relationName) {
        var relationDef = typeDef.Relations![relationName];

        // Union
        var relationUnionNameSet = new HashSet<string>();
        foreach (var userset in relationDef.Union?.Child ?? []) {
            var relationDefName = Dsl.GetRelationDefName(userset);
            if (relationDefName != null && relationUnionNameSet.Contains(relationDefName)) {
                var typeIndex = _dsl.GetTypeLineNumber(typeDef.Type);
                var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                _errors.RaiseDuplicateType(lineIndex, relationDefName, relationName);
            }
            relationUnionNameSet.Add(relationDefName!);
        }

        // Intersection
        var relationIntersectionNameSet = new HashSet<string>();
        foreach (var userset in relationDef.Intersection?.Child ?? []) {
            var relationDefName = Dsl.GetRelationDefName(userset);
            if (relationDefName != null && relationIntersectionNameSet.Contains(relationDefName)) {
                var typeIndex = _dsl.GetTypeLineNumber(typeDef.Type);
                var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                _errors.RaiseDuplicateType(lineIndex, relationDefName, relationName);
            }
            relationIntersectionNameSet.Add(relationDefName!);
        }

        // Difference
        if (relationDef.Difference != null) {
            var baseName = Dsl.GetRelationDefName(relationDef.Difference.Base);
            var subtractName = Dsl.GetRelationDefName(relationDef.Difference.Subtract);
            if (baseName != null && baseName.Equals(subtractName)) {
                var typeIndex = _dsl.GetTypeLineNumber(typeDef.Type);
                var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                _errors.RaiseDuplicateType(lineIndex, baseName, relationName);
            }
        }
    }

    private void CheckForDuplicatesTypeNamesInRelation(RelationMetadata relationDef, string relationName) {
        var typeNameSet = new HashSet<string>();
        foreach (var typeDef in relationDef.DirectlyRelatedUserTypes ?? []) {
            var typeDefName = Dsl.GetTypeRestrictionString(typeDef);
            if (typeNameSet.Contains(typeDefName)) {
                var typeIndex = _dsl.GetTypeLineNumber(typeDef.Type);
                var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                _errors.RaiseDuplicateTypeRestriction(lineIndex, typeDefName, relationName);
            }
            typeNameSet.Add(typeDefName);
        }
    }

    private void RelationDefined(Dictionary<string, TypeDefinition> typeMap, string typeName, string relationName) {
        var relations = typeMap[typeName].Relations;
        if (relations == null || relations.Count == 0 || !relations.ContainsKey(relationName)) {
            return;
        }

        var currentRelation = relations[relationName];
        var children = new List<Userset> { currentRelation };

        while (children.Count > 0) {
            var child = children[0];
            children.RemoveAt(0);

            if (child.Union != null) {
                children.AddRange(child.Union.Child);
            }
            else if (child.Intersection != null) {
                children.AddRange(child.Intersection.Child);
            }
            else if (child.Difference != null && child.Difference.Base != null && child.Difference.Subtract != null) {
                children.Add(child.Difference.Base);
                children.Add(child.Difference.Subtract);
            }
            else {
                ChildDefDefined(typeMap, typeName, relationName, Dsl.GetRelationalParserResult(child));
            }
        }
    }

    private void ChildDefDefined(
        Dictionary<string, TypeDefinition> typeMap,
        string typeName,
        string relationName,
        RelationTargetParserResult childDef) {
        var relations = typeMap[typeName].Relations;
        if (relations == null || relations.Count == 0 || !relations.ContainsKey(relationName)) {
            return;
        }

        RelationMetadata? currentRelationMetadata = null;
        if (typeMap[typeName].Metadata != null) {
            currentRelationMetadata = typeMap[typeName].Metadata!.Relations?.GetValueOrDefault(relationName);
        }

        switch (childDef.Rewrite) {
            case RewriteType.Direct: {
                    var relatedTypes = currentRelationMetadata?.DirectlyRelatedUserTypes ?? new List<RelationReference>();
                    var fromPossibleTypes = Dsl.GetTypeRestrictions(relatedTypes);
                    if (fromPossibleTypes.Count == 0) {
                        var typeIndex = _dsl.GetTypeLineNumber(typeName);
                        var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                        _errors.RaiseAssignableRelationMustHaveTypes(lineIndex, relationName);
                    }
                    foreach (var item in fromPossibleTypes) {
                        var type = DestructuredTupleToUserset.From(item);
                        var decodedType = type.DecodedType;
                        if (!typeMap.ContainsKey(decodedType)) {
                            var typeIndex = _dsl.GetTypeLineNumber(typeName);
                            var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                            _errors.RaiseInvalidType(lineIndex, decodedType, decodedType);
                        }

                        var decodedConditionName = type.DecodedConditionName;
                        if (decodedConditionName != null && !(_authorizationModel.Conditions?.ContainsKey(decodedConditionName) ?? false)) {
                            var typeIndex = _dsl.GetTypeLineNumber(typeName);
                            var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                            _errors.RaiseInvalidConditionNameInParameter(
                                lineIndex, decodedConditionName, typeName, relationName, decodedConditionName);
                        }

                        var decodedRelation = type.DecodedRelation;
                        if (type.IsWildcard && decodedRelation != null) {
                            var typeIndex = _dsl.GetTypeLineNumber(typeName);
                            var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                            _errors.RaiseAssignableTypeWildcardRelation(lineIndex, item);
                        }
                        else if (decodedRelation != null) {
                            if (typeMap[decodedType] == null || !typeMap[decodedType].Relations!.ContainsKey(decodedRelation)) {
                                var typeIndex = _dsl.GetTypeLineNumber(typeName);
                                var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                                _errors.RaiseInvalidTypeRelation(
                                    lineIndex, decodedType + "#" + decodedRelation, decodedType, decodedRelation);
                            }
                        }
                    }
                    break;
                }
            case RewriteType.ComputedUserset: {
                    if (childDef.Target != null && !relations.ContainsKey(childDef.Target)) {
                        var typeIndex = _dsl.GetTypeLineNumber(typeName);
                        var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                        var value = childDef.Target;
                        _errors.RaiseInvalidRelationError(lineIndex, value, relations.Keys);
                    }
                    break;
                }
            case RewriteType.TupleToUserset: {
                    if (childDef.From != null && childDef.Target != null) {
                        if (!relations.ContainsKey(childDef.From)) {
                            var typeIndex = _dsl.GetTypeLineNumber(typeName);
                            var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                            _errors.RaiseInvalidTypeRelation(
                                lineIndex,
                                childDef.Target + " from " + childDef.From,
                                typeName,
                                childDef.From);
                        }
                        else {
                            var allowableTypesResult = AllowableTypes(typeMap, typeName, childDef.From);
                            if (allowableTypesResult.Valid && allowableTypesResult.AllowableTypes.Count > 0) {
                                var childRelationNotValid = new List<InvalidChildRelationMetadata>();
                                var fromTypes = allowableTypesResult.AllowableTypes;
                                foreach (var item in fromTypes) {
                                    var type = DestructuredTupleToUserset.From(item);
                                    var decodedType = type.DecodedType;
                                    var decodedRelation = type.DecodedRelation;
                                    var isWildcard = type.IsWildcard;
                                    if (isWildcard || decodedRelation != null) {
                                        var typeIndex = _dsl.GetTypeLineNumber(typeName);
                                        var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                                        _errors.RaiseTupleUsersetRequiresDirect(lineIndex, childDef.From);
                                    }
                                    else {
                                        if (!typeMap[decodedType].Relations!.ContainsKey(childDef.Target)) {
                                            var typeIndex = _dsl.GetTypeLineNumber(typeName);
                                            var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                                            childRelationNotValid.Add(new InvalidChildRelationMetadata(
                                                lineIndex,
                                                childDef.Target + " from " + childDef.From,
                                                decodedType,
                                                childDef.Target,
                                                childDef.From));
                                        }
                                    }
                                }

                                if (childRelationNotValid.Count == fromTypes.Count) {
                                    foreach (var item in childRelationNotValid) {
                                        _errors.RaiseInvalidRelationOnTupleset(
                                            item.LineIndex,
                                            item.Symbol,
                                            item.TypeName,
                                            typeName,
                                            relationName,
                                            item.RelationName,
                                            item.Parent);
                                    }
                                }
                            }
                            else {
                                var typeIndex = _dsl.GetTypeLineNumber(typeName);
                                var lineIndex = _dsl.GetRelationLineNumber(relationName, typeIndex);
                                _errors.RaiseTupleUsersetRequiresDirect(lineIndex, childDef.From);
                            }
                        }
                    }
                    break;
                }
        }
    }

    private static AllowableTypesResult AllowableTypes(
        Dictionary<string, TypeDefinition> typeMap, string typeName, string relation) {
        var allowedTypes = new List<string>();
        var typeDefinition = typeMap[typeName];
        var currentRelation = typeDefinition.Relations![relation];
        var metadata = typeDefinition.Metadata;
        var relatedTypes = metadata?.Relations?.GetValueOrDefault(relation)?.DirectlyRelatedUserTypes ?? new List<RelationReference>();
        var currentRelationMetadata = Dsl.GetTypeRestrictions(relatedTypes);
        var isValid = IsRelationSingle(currentRelation);
        if (isValid) {
            var childDef = Dsl.GetRelationalParserResult(currentRelation);
            if (childDef.Rewrite == RewriteType.Direct) {
                allowedTypes.AddRange(currentRelationMetadata);
            }
        }
        return new AllowableTypesResult(isValid, allowedTypes);
    }

    private static bool IsRelationSingle(Userset currentRelation) {
        return currentRelation.Union == null
               && currentRelation.Intersection == null
               && currentRelation.Difference == null;
    }

    private void TrackModulesInFile(Metadata? metadata) {
        if (metadata == null) {
            return;
        }

        var sourceInfo = metadata.SourceInfo;
        var module = metadata.Module;
        TrackModulesInFile(module, sourceInfo);
    }

    private void TrackModulesInFile(Metadata? metadata, RelationMetadata? relationMetadata) {
        string? module = null;
        SourceInfo? sourceInfo = null;
        if (relationMetadata != null) {
            module = relationMetadata.Module;
            sourceInfo = relationMetadata.SourceInfo;
        }

        if (module == null) {
            module = metadata?.Module;
            sourceInfo = metadata?.SourceInfo;
        }

        TrackModulesInFile(module, sourceInfo);
    }

    private void TrackModulesInFile(ConditionMetadata? metadata) {
        if (metadata == null) {
            return;
        }

        var sourceInfo = metadata.SourceInfo;
        var module = metadata.Module;
        TrackModulesInFile(module, sourceInfo);
    }

    private void TrackModulesInFile(string? module, SourceInfo? sourceInfo) {
        if (module == null || sourceInfo == null) {
            return;
        }

        if (!_fileToModules.ContainsKey(sourceInfo.File!)) {
            _fileToModules[sourceInfo.File!] = new HashSet<string>();
        }
        _fileToModules[sourceInfo.File!].Add(module);
    }
}