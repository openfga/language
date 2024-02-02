package dev.openfga.language.validation;

import dev.openfga.language.DslToJsonTransformer;
import dev.openfga.language.errors.*;
import dev.openfga.sdk.api.model.*;
import dev.openfga.sdk.api.model.Metadata;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.RequiredArgsConstructor;

import java.io.IOException;
import java.util.*;
import java.util.function.Predicate;
import java.util.regex.Pattern;
import java.util.stream.IntStream;

import static dev.openfga.language.Utils.getNullSafe;
import static dev.openfga.language.Utils.getNullSafeList;
import static java.util.Collections.emptyList;
import static java.util.stream.Collectors.toList;

@RequiredArgsConstructor
public class DslValidator {

    private final ValidationOptions options;
    private final String dsl;
    private String[] lines;

    private AuthorizationModel authorizationModel;
    private final List<ModelValidationSingleError> errors = new ArrayList<>();
    private ValidationRegex typeRegex;
    private ValidationRegex relationRegex;

    public static void validate(String dsl) throws DslErrorsException, IOException {
        validate(dsl, new ValidationOptions());
    }

    public static void validate(String dsl, ValidationOptions options) throws DslErrorsException, IOException {
        new DslValidator(options, dsl).validate();
    }

    public void validate() throws DslErrorsException, IOException {
        var transformer = new DslToJsonTransformer();
        var result = transformer.parseDsl(dsl);
        if (result.IsFailure()) {
            throw new DslErrorsException(result.getErrors());
        }
        authorizationModel = result.getAuthorizationModel();

        lines = dsl.split("\n");

        typeRegex = ValidationRegex.build("type", options.getTypePattern());
        relationRegex = ValidationRegex.build("relation", options.getRelationPattern());

        populateRelations();

        var schemaVersion = authorizationModel.getSchemaVersion();
        if (schemaVersion == null) {
            raiseSchemaVersionRequired(0, "");
        }

        if (schemaVersion.equals("1.1")) {
            modelValidation();
        } else {
            var lineIndex = getSchemaLineNumber(schemaVersion);
            raiseInvalidSchemaVersion(lineIndex, schemaVersion);
        }

        if (!errors.isEmpty()) {
            throw new DslErrorsException(errors);
        }
    }

    private void modelValidation() {
        if (!errors.isEmpty()) {
            return;
        }

        var typeMap = new HashMap<String, TypeDefinition>();
        var usedConditionNamesSet = new HashSet<String>();
        authorizationModel.getTypeDefinitions().forEach(typeDef -> {
            var typeName = typeDef.getType();
            typeMap.put(typeName, typeDef);

            if (typeDef.getMetadata() != null) {
                typeDef.getMetadata().getRelations().forEach((relationName, relationMetadata) -> {
                    relationMetadata.getDirectlyRelatedUserTypes().forEach(typeRestriction -> {
                        if (typeRestriction.getCondition() != null) {
                            usedConditionNamesSet.add(typeRestriction.getCondition());
                        }
                    });
                });
            }
        });

        // first, validate to ensure all the relation are defined
        authorizationModel.getTypeDefinitions().forEach(typeDef -> {
            var typeName = typeDef.getType();
            typeDef.getRelations().forEach((relationName, relationDef) -> {
                relationDefined(typeMap, typeName, relationName);
            });
        });

        if (errors.isEmpty()) {
            var typeSet = new HashSet<String>();
            authorizationModel.getTypeDefinitions().forEach(typeDef -> {
                var typeName = typeDef.getType();
                if (typeSet.contains(typeName)) {
                    var typeIndex = getTypeLineNumber(typeName);
                    raiseDuplicateTypeName(typeIndex, typeName);
                }
                typeSet.add(typeName);

                if (typeDef.getMetadata() != null) {
                    for (var relationDefKey : typeDef.getMetadata().getRelations().keySet()) {
                        checkForDuplicatesTypeNamesInRelation(typeDef.getMetadata().getRelations().get(relationDefKey), relationDefKey);
                        checkForDuplicatesInRelation(typeDef, relationDefKey);
                    }
                }
            });
        }

        // next, ensure all relation have entry point
        // we can skip if there are errors because errors (such as missing relations) will likely lead to no entries
        if (errors.isEmpty()) {
            authorizationModel.getTypeDefinitions().forEach(typeDef -> {
                var typeName = typeDef.getType();
                for (var relationName : typeDef.getRelations().keySet()) {
                    var currentRelations = typeMap.get(typeName).getRelations();
                    var result = hasEntryPointOrLoop(typeMap, typeName, relationName, currentRelations.get(relationName), new HashMap<>());
                    if(!result.isHasEntry()) {
                        var typeIndex = getTypeLineNumber(typeName);
                        var lineIndex = getRelationLineNumber(relationName, typeIndex);
                        if (result.isLoop()) {
                            raiseNoEntryPointLoop(lineIndex, relationName, typeName);
                        } else {
                            raiseNoEntryPoint(lineIndex, relationName, typeName);
                        }
                    }
                }
            });
        }

        authorizationModel.getConditions().forEach((conditionName, condition) -> {
            if (!usedConditionNamesSet.contains(conditionName)) {
                var conditionIndex = getConditionLineNumber(conditionName);
                raiseUnusedCondition(conditionIndex, conditionName);
            }
        });
    }

    private int getConditionLineNumber(String conditionName) {
        return getConditionLineNumber(conditionName, 0);
    }
    private int getConditionLineNumber(String conditionName, int skipIndex) {
        return findLine(
                line -> line.trim().startsWith("condition " + conditionName),
                skipIndex);
    }

    @Getter
    @AllArgsConstructor
    private static class EntryPointOrLoopResult {
        private final boolean hasEntry;
        private final boolean loop;

        public static final EntryPointOrLoopResult BOTH_FALSE = new EntryPointOrLoopResult(false, false);
        public static final EntryPointOrLoopResult HAS_ENTRY_BUT_NO_LOOP = new EntryPointOrLoopResult(true, false);
        public static final EntryPointOrLoopResult NO_ENTRY_WITH_LOOP = new EntryPointOrLoopResult(false, true);
    }

    // for the type/relation, whether there are any unique entry points, and if a loop is found
    // if there are unique entry points (i.e., direct relations) then it will return true
    // otherwise, it will follow its children to see if there are unique entry points
    // if there is a loop during traversal, the function will return a boolean indicating so
    private EntryPointOrLoopResult hasEntryPointOrLoop(Map<String, TypeDefinition> typeMap, String typeName, String relationName, Userset rewrite, Map<String, Map<String, Boolean>> visitedRecords) {
        var visited = deepCopy(visitedRecords);

        if (relationName == null) {
            return EntryPointOrLoopResult.BOTH_FALSE;
        }

        if (!visited.containsKey(typeName)) {
            visited.put(typeName, new HashMap<>());
        }
        visited.get(typeName).put(relationName, true);

        var currentRelations = typeMap.get(typeName).getRelations();
        if (currentRelations == null || !currentRelations.containsKey(relationName)) {
            return EntryPointOrLoopResult.BOTH_FALSE;
        }

        if(typeMap.get(typeName).getRelations() == null || !typeMap.get(typeName).getRelations().containsKey(relationName)) {
            return EntryPointOrLoopResult.BOTH_FALSE;
        }

        var relationsMetada = getNullSafe(typeMap.get(typeName).getMetadata(), Metadata::getRelations);
        if (rewrite.getThis() != null) {
            if (relationsMetada != null) {
                var relationMetadata = relationsMetada.get(relationName);
                var relatedTypes = getNullSafeList(relationMetadata, RelationMetadata::getDirectlyRelatedUserTypes);
                for(var assignableType :  getTypeRestrictions(relatedTypes)) {
                    var destructuredType = destructTupleToUserset(assignableType);
                    var decodedRelation = destructuredType.getDecodedRelation();
                    if (decodedRelation == null || destructuredType.isWildcard()) {
                        return EntryPointOrLoopResult.HAS_ENTRY_BUT_NO_LOOP;
                    }

                    var decodedType = destructuredType.getDecodedType();
                    var assignableRelation = typeMap.get(decodedType).getRelations().get(decodedRelation);
                    if (assignableRelation == null) {
                        return EntryPointOrLoopResult.BOTH_FALSE;
                    }

                    if (getNullSafe(visited.get(decodedType), m -> m.get(decodedRelation)) != null) {
                        continue;
                    }

                    var entryPointOrLoop = hasEntryPointOrLoop(typeMap, decodedType, decodedRelation, assignableRelation, visited);
                    if(entryPointOrLoop.isHasEntry()) {
                        return EntryPointOrLoopResult.HAS_ENTRY_BUT_NO_LOOP;
                    }
                }
            }
            return EntryPointOrLoopResult.BOTH_FALSE;
        } else if (rewrite.getComputedUserset() != null) {
            var computedRelationName = rewrite.getComputedUserset().getRelation();
            if(computedRelationName == null) {
                return EntryPointOrLoopResult.BOTH_FALSE;
            }

            var computedRelation = typeMap.get(typeName).getRelations().get(computedRelationName);
            if(computedRelation == null) {
                return EntryPointOrLoopResult.BOTH_FALSE;
            }

            // Loop detected
            if (visited.get(typeName).containsKey(computedRelationName)) {
                return EntryPointOrLoopResult.NO_ENTRY_WITH_LOOP;
            }

            return hasEntryPointOrLoop(typeMap, typeName, computedRelationName, computedRelation, visited);
        } else if (rewrite.getTupleToUserset() != null) {
            var tuplesetRelationName = rewrite.getTupleToUserset().getTupleset().getRelation();
            var computedRelationName = rewrite.getTupleToUserset().getComputedUserset().getRelation();
            if (tuplesetRelationName == null || computedRelationName == null) {
                return EntryPointOrLoopResult.BOTH_FALSE;
            }

            if (!typeMap.get(typeName).getRelations().containsKey(tuplesetRelationName)) {
                return EntryPointOrLoopResult.BOTH_FALSE;
            }
//            var tuplesetRelation = typeMap.get(typeName).getRelations().get(tuplesetRelationName);


            if (relationsMetada != null) {
                var relationMetadata = relationsMetada.get(tuplesetRelationName);
                var relatedTypes = getNullSafeList(relationMetadata, RelationMetadata::getDirectlyRelatedUserTypes);
                for(var assignableType :  getTypeRestrictions(relatedTypes)) {
                    var assignableRelation = typeMap.get(assignableType).getRelations().get(computedRelationName);
                    if (assignableRelation != null) {
                        if(visited.containsKey(assignableType) && visited.get(assignableType).containsKey(computedRelationName)) {
                            continue;
                        }

                        var entryOrLoop = hasEntryPointOrLoop(typeMap, assignableType, computedRelationName, assignableRelation, visited);
                        if(entryOrLoop.isHasEntry()) {
                            return EntryPointOrLoopResult.HAS_ENTRY_BUT_NO_LOOP;
                        }
                    }
                }
            }
            return EntryPointOrLoopResult.BOTH_FALSE;
        } else if (rewrite.getUnion() != null) {
            var loop = false;

            for (var child : rewrite.getUnion().getChild()) {
                var childEntryOrLoop = hasEntryPointOrLoop(typeMap, typeName, relationName, child, visited);
                if (childEntryOrLoop.isHasEntry()) {
                    return EntryPointOrLoopResult.HAS_ENTRY_BUT_NO_LOOP;
                }
                loop = loop || childEntryOrLoop.isLoop();
            }
            return new EntryPointOrLoopResult(false, loop);
        } else if (rewrite.getIntersection() != null) {
            for (var child : rewrite.getIntersection().getChild()) {
                var childEntryOrLoop = hasEntryPointOrLoop(typeMap, typeName, relationName, child, visited);
                if (!childEntryOrLoop.isHasEntry()) {
                    return childEntryOrLoop;
                }
            }
            return EntryPointOrLoopResult.HAS_ENTRY_BUT_NO_LOOP;
        } else if (rewrite.getDifference() != null) {
            var baseEntryOrLoop = hasEntryPointOrLoop(typeMap, typeName, relationName, rewrite.getDifference().getBase(), visited);
            if(!baseEntryOrLoop.isHasEntry()) {
                return baseEntryOrLoop;
            }
            var substractEntryOrLoop = hasEntryPointOrLoop(typeMap, typeName, relationName, rewrite.getDifference().getSubtract(), visited);
            if(!substractEntryOrLoop.isHasEntry()) {
                return substractEntryOrLoop;
            }
            return EntryPointOrLoopResult.HAS_ENTRY_BUT_NO_LOOP;
        }
        return EntryPointOrLoopResult.BOTH_FALSE;
    }

    private Map<String, Map<String, Boolean>> deepCopy(Map<String, Map<String, Boolean>> records) {
        Map<String, Map<String, Boolean>> copy = new HashMap<>();
        records.forEach((key, value) -> copy.put(key, new HashMap<>(value)));
        return copy;
    }

    private void checkForDuplicatesInRelation(TypeDefinition typeDef, String relationName) {
        var relationDef = typeDef.getRelations().get(relationName);

        // Union
        var relationUnionNameSet = new HashSet<String>();
        getNullSafeList(relationDef.getUnion(), Usersets::getChild).forEach(userset -> {
            var relationDefName = getRelationDefName(userset);
            if(relationDefName != null && relationUnionNameSet.contains(relationDefName)) {
                var typeIndex = getTypeLineNumber(typeDef.getType());
                var lineIndex = getRelationLineNumber(relationName, typeIndex);
                raiseDuplicateType(lineIndex, relationDefName, relationName);
            }
            relationUnionNameSet.add(relationDefName);
        });

        // Intersection
        var relationIntersectionNameSet = new HashSet<String>();
        getNullSafeList(relationDef.getIntersection(), Usersets::getChild).forEach(userset -> {
            var relationDefName = getRelationDefName(userset);
            if(relationDefName != null && relationIntersectionNameSet.contains(relationDefName)) {
                var typeIndex = getTypeLineNumber(typeDef.getType());
                var lineIndex = getRelationLineNumber(relationName, typeIndex);
                raiseDuplicateType(lineIndex, relationDefName, relationName);
            }
            relationIntersectionNameSet.add(relationDefName);
        });

        // Difference
        if (relationDef.getDifference() != null) {
            var baseName = getRelationDefName(relationDef.getDifference().getBase());
            var substractName = getRelationDefName(relationDef.getDifference().getSubtract());
            if (baseName != null && baseName.equals(substractName)) {
                var typeIndex = getTypeLineNumber(typeDef.getType());
                var lineIndex = getRelationLineNumber(relationName, typeIndex);
                raiseDuplicateType(lineIndex, baseName, relationName);
            }
        }
    }

    private String getRelationDefName(Userset userset) {
        var relationDefName = getNullSafe(userset.getComputedUserset(), ObjectRelation::getRelation);
        var parserResult = getRelationalParserResult(userset);
        if (parserResult.getRewrite() == RewriteType.ComputedUserset) {
            relationDefName = parserResult.getTarget();
        } else if (parserResult.getRewrite() == RewriteType.TupleToUserset) {
            relationDefName = parserResult.getTarget() + " from " + parserResult.getFrom();
        }
        return relationDefName;
    }

    private void checkForDuplicatesTypeNamesInRelation(RelationMetadata relationDef, String relationName) {
        var typeNameSet = new HashSet<String>();
        relationDef.getDirectlyRelatedUserTypes().forEach(typeDef -> {
            var typeDefName = getTypeRestrictionString(typeDef);
            if (typeNameSet.contains(typeDefName)) {
                var typeIndex = getTypeLineNumber(typeDef.getType());
                var lineIndex = getRelationLineNumber(relationName, typeIndex);
                raiseDuplicateTypeRestriction(lineIndex, typeDefName, relationName);
            }
            typeNameSet.add(typeDefName);
        });
    }

    private void relationDefined(Map<String, TypeDefinition> typeMap, String typeName, String relationName) {
        var relations = typeMap.get(typeName).getRelations();
        if (relations == null || relations.isEmpty() || !relations.containsKey(relationName)) {
            return;
        }

        var currentRelation = relations.get(relationName);
        var children = new ArrayList<Userset>() {{
            add(currentRelation);
        }};
        while (!children.isEmpty()) {
            var child = children.remove(0);
            if (child.getUnion() != null) {
                children.addAll(child.getUnion().getChild());
            } else if (child.getIntersection() != null) {
                children.addAll(child.getIntersection().getChild());
            } else if (child.getDifference() != null && child.getDifference().getBase() != null && child.getDifference().getSubtract() != null) {
                children.add(child.getDifference().getBase());
                children.add(child.getDifference().getSubtract());
            } else {
                childDefDefined(typeMap, typeName, relationName, getRelationalParserResult(child));
            }
        }
    }

    @AllArgsConstructor
    @Getter
    private static class DestructuredTupleToUserset {
        private final String decodedType;
        private final String decodedRelation;
        private final boolean wildcard;
        private final String decodedConditionName;

    }

    private DestructuredTupleToUserset destructTupleToUserset(String allowableType) {
        var tupleAndCondition = allowableType.split(" with ");
        var tupleString = tupleAndCondition[0];
        var decodedConditionName = tupleAndCondition.length > 1 ? tupleAndCondition[1] : null;
        var isWildcard = tupleString.contains(":*");
        var splittedWords = tupleString.replace(":*", "").split("#");
        return new DestructuredTupleToUserset(
                splittedWords[0],
                splittedWords.length > 1 ? splittedWords[1] : null,
                isWildcard,
                decodedConditionName);
    }

    private void childDefDefined(Map<String, TypeDefinition> typeMap, String typeName, String relationName, RelationTargetParserResult childDef) {
        var relations = typeMap.get(typeName).getRelations();
        if (relations == null || relations.isEmpty() || !relations.containsKey(relationName)) {
            return;
        }

        RelationMetadata currentRelationMetadata = null;
        if (typeMap.get(typeName).getMetadata() != null) {
            currentRelationMetadata = typeMap.get(typeName).getMetadata().getRelations().get(relationName);
        }

        switch (childDef.getRewrite()) {
            case Direct: {
                var relatedTypes = currentRelationMetadata != null
                        ? currentRelationMetadata.getDirectlyRelatedUserTypes()
                        : new ArrayList<RelationReference>();
                var fromPossibleTypes = getTypeRestrictions(relatedTypes);
                if (fromPossibleTypes.isEmpty()) {
                    var typeIndex = getTypeLineNumber(typeName);
                    var lineIndex = getRelationLineNumber(relationName, typeIndex);
                    raiseAssignableRelationMustHaveTypes(lineIndex, relationName);
                }
                for (var item : fromPossibleTypes) {
                    var type = destructTupleToUserset(item);
                    var decodedType = type.getDecodedType();
                    if (!typeMap.containsKey(decodedType)) {
                        var typeIndex = getTypeLineNumber(typeName);
                        var lineIndex = getRelationLineNumber(relationName, typeIndex);
                        raiseInvalidType(lineIndex, decodedType, decodedType);
                    }

                    var decodedConditionName = type.getDecodedConditionName();
                    if (decodedConditionName != null && !authorizationModel.getConditions().containsKey(decodedConditionName)) {
                        var typeIndex = getTypeLineNumber(typeName);
                        var lineIndex = getRelationLineNumber(relationName, typeIndex);
                        raiseInvalidConditionNameInParameter(lineIndex, decodedConditionName, typeName, relationName, decodedConditionName);
                    }

                    var decodedRelation = type.getDecodedRelation();
                    if (type.isWildcard() && decodedRelation != null) {
                        var typeIndex = getTypeLineNumber(typeName);
                        var lineIndex = getRelationLineNumber(relationName, typeIndex);
                        raiseAssignableTypeWildcardRelation(lineIndex, item);
                    } else if (decodedRelation != null) {
                        if (typeMap.get(decodedType) == null || !typeMap.get(decodedType).getRelations().containsKey(decodedRelation)) {
                            var typeIndex = getTypeLineNumber(typeName);
                            var lineIndex = getRelationLineNumber(relationName, typeIndex);
                            raiseInvalidTypeRelation(
                                    lineIndex,
                                    decodedType + "#" + decodedRelation,
                                    decodedType,
                                    decodedRelation);
                        }
                    }
                }
                break;
            }
            case ComputedUserset: {
                if (childDef.getTarget() != null && relations.get(childDef.getTarget()) == null) {
                    var typeIndex = getTypeLineNumber(typeName);
                    var lineIndex = getRelationLineNumber(relationName, typeIndex);
                    var value = childDef.getTarget();
                    raiseInvalidRelationError(lineIndex, value, relations.keySet());
                }
                break;
            }
            case TupleToUserset: {
                if (childDef.getFrom() != null && childDef.getTarget() != null) {
                    if (!relations.containsKey(childDef.getFrom())) {
                        var typeIndex = getTypeLineNumber(typeName);
                        var lineIndex = getRelationLineNumber(relationName, typeIndex);
                        raiseInvalidTypeRelation(
                                lineIndex,
                                childDef.getTarget() + " from " + childDef.getFrom(),
                                typeName,
                                childDef.getFrom()
                        );
                    } else {
                        var allowableTypesResult = allowableTypes(typeMap, typeName, childDef.getFrom());
                        if (allowableTypesResult.isValid()) {
                            var childRelationNotValid = new ArrayList<InvalidChildRelationMetadata>();
                            var fromTypes = allowableTypesResult.getAllowableTypes();
                            for (var item : fromTypes) {
                                var type = destructTupleToUserset(item);
                                var decodedType = type.getDecodedType();
                                var decodedRelation = type.getDecodedRelation();
                                var isWilcard = type.isWildcard();
                                if (isWilcard) {
                                    var typeIndex = getTypeLineNumber(typeName);
                                    var lineIndex = getRelationLineNumber(relationName, typeIndex);
                                    raiseAssignableTypeWildcardRelation(lineIndex, item);
                                } else if (decodedRelation != null) {
                                    var typeIndex = getTypeLineNumber(typeName);
                                    var lineIndex = getRelationLineNumber(relationName, typeIndex);
                                    raiseTupleUsersetRequiresDirect(lineIndex, childDef.getFrom());
                                } else {
                                    if (typeMap.get(decodedType) != null && !typeMap.get(decodedType).getRelations().containsKey(childDef.getTarget())) {
                                        var typeIndex = getTypeLineNumber(typeName);
                                        var lineIndex = getRelationLineNumber(relationName, typeIndex);
                                        childRelationNotValid.add(new InvalidChildRelationMetadata(
                                                lineIndex,
                                                childDef.getTarget() + " from " + childDef.getFrom(),
                                                decodedType,
                                                childDef.getTarget()));
                                    }
                                }
                            }

                            if (childRelationNotValid.size() == fromTypes.size()) {
                                for (var item : childRelationNotValid) {
                                    raiseInvalidTypeRelation(
                                            item.getLineIndex(),
                                            item.getSymbol(),
                                            item.getTypeName(),
                                            item.getRelationName());
                                }
                            }
                        } else {
                            var typeIndex = getTypeLineNumber(typeName);
                            var lineIndex = getRelationLineNumber(relationName, typeIndex);
                            raiseTupleUsersetRequiresDirect(lineIndex, childDef.getFrom());
                        }
                    }
                }
                break;
            }
        }

    }

    @Getter
    @AllArgsConstructor
    private static class InvalidChildRelationMetadata {
        private final int lineIndex;
        private final String symbol;
        private final String typeName;
        private final String relationName;

    }

    @AllArgsConstructor
    @Getter
    private static class AllowableTypesResult {
        private final boolean valid;
        private final List<String> allowableTypes;

    }

    private AllowableTypesResult allowableTypes(Map<String, TypeDefinition> typeMap, String typeName, String relation) {
        var allowedTypes = new ArrayList<String>();
        var typeDefinition = typeMap.get(typeName);
        var currentRelation = typeDefinition.getRelations().get(relation);
        var metadata = typeDefinition.getMetadata();
        Collection<RelationReference> relatedTypes = metadata != null
                ? metadata.getRelations().get(relation).getDirectlyRelatedUserTypes()
                : emptyList();
        var currentRelationMetadata = getTypeRestrictions(relatedTypes);
        var isValid = relationIsSingle(currentRelation);
        if (isValid) {
            var childDef = getRelationalParserResult(currentRelation);
            if (childDef.getRewrite() == RewriteType.Direct) {
                allowedTypes.addAll(currentRelationMetadata);
            }
        }
        return new AllowableTypesResult(isValid, allowedTypes);
    }

    private boolean relationIsSingle(Userset currentRelation) {
        return currentRelation.getUnion() == null
                && currentRelation.getIntersection() == null
                && currentRelation.getDifference() == null;
    }

    private List<String> getTypeRestrictions(Collection<RelationReference> relatedTypes) {
        return relatedTypes.stream()
                .map(this::getTypeRestrictionString)
                .collect(toList());
    }

    private String getTypeRestrictionString(RelationReference typeRestriction) {
        var typeRestrictionString = typeRestriction.getType();
        if (typeRestriction.getWildcard() != null) {
            typeRestrictionString += ":*";
        } else if (typeRestriction.getRelation() != null) {
            typeRestrictionString += "#" + typeRestriction.getRelation();
        }

        if (typeRestriction.getCondition() != null) {
            typeRestrictionString += " with " + typeRestriction.getCondition();
        }

        return typeRestrictionString;
    }

    ;

    private RelationTargetParserResult getRelationalParserResult(Userset userset) {
        String target = null, from = null;

        if (userset.getComputedUserset() != null) {
            target = userset.getComputedUserset().getRelation();
        } else {
            if (userset.getTupleToUserset() != null && userset.getTupleToUserset().getComputedUserset() != null) {
                target = userset.getTupleToUserset().getComputedUserset().getRelation();
            }
            if (userset.getTupleToUserset() != null && userset.getTupleToUserset().getTupleset() != null) {
                from = userset.getTupleToUserset().getTupleset().getRelation();
            }
        }

        var rewrite = RewriteType.Direct;
        if (target != null) {
            rewrite = RewriteType.ComputedUserset;
        }

        if (from != null) {
            rewrite = RewriteType.TupleToUserset;
        }
        return new RelationTargetParserResult(target, from, rewrite);
    }

    private void populateRelations() {
        authorizationModel.getTypeDefinitions().forEach(typeDef -> {
            var typeName = typeDef.getType();

            if (typeName.equals(Keyword.SELF) || typeName.equals(Keyword.THIS)) {
                var lineIndex = getTypeLineNumber(typeName);
                raiseReservedTypeName(lineIndex, typeName);
            }

            if (!typeRegex.matches(typeName)) {
                var lineIndex = getTypeLineNumber(typeName);
                raiseInvalidName(lineIndex, typeName, typeRegex.getRule());
            }

            var encounteredRelationsInType = new HashSet<String>() {{
                add(Keyword.SELF);
            }};
            typeDef.getRelations().forEach((relationName, relation) -> {
                if (relationName.equals(Keyword.SELF) || relationName.equals(Keyword.THIS)) {
                    var typeIndex = getTypeLineNumber(typeName);
                    var lineIndex = getRelationLineNumber(relationName, typeIndex);
                    raiseReservedRelationName(lineIndex, relationName);
                } else if (!relationRegex.matches(relationName)) {
                    var typeIndex = getTypeLineNumber(typeName);
                    var lineIndex = getRelationLineNumber(relationName, typeIndex);
                    raiseInvalidName(lineIndex, relationName, relationRegex.getRule(), typeName);
                } else if (encounteredRelationsInType.contains(relationName)) {
                    var typeIndex = getTypeLineNumber(typeName);
                    var initialLineIdx = getRelationLineNumber(relationName, typeIndex);
                    var duplicateLineIdx = getRelationLineNumber(relationName, initialLineIdx + 1);
                    raiseDuplicateRelationName(duplicateLineIdx, relationName);
                }
                encounteredRelationsInType.add(relationName);
            });
        });

    }

    private int getRelationLineNumber(String relationName) {
        return getRelationLineNumber(relationName, 0);
    }

    private int getRelationLineNumber(String relationName, int skipIndex) {
        return findLine(
                line -> line.trim().replaceAll(" {2,}", " ").startsWith("define " + relationName),
                skipIndex);
    }

    private int getSchemaLineNumber(String schemaVersion) {
        return findLine(
                line -> line.trim().replaceAll(" {2,}", " ").startsWith("schema " + schemaVersion),
                0);
    }

    private int getTypeLineNumber(String typeName) {
        return getTypeLineNumber(typeName, 0);
    }

    private int getTypeLineNumber(String typeName, int skipIndex) {
        return findLine(
                line -> line.trim().startsWith("type " + typeName),
                skipIndex);
    }

    private int findLine(Predicate<String> predicate, int skipIndex) {
        return IntStream.range(skipIndex, lines.length)
                .filter(index -> predicate.test(lines[index]))
                .findFirst().orElse(-1);
    }

    private ErrorProperties buildErrorProperties(String message, int lineIndex, String symbol) {
        return buildErrorProperties(message, lineIndex, symbol, null);
    }

    private ErrorProperties buildErrorProperties(String message, int lineIndex, String symbol, WordResolver wordResolver) {

        var rawLine = lines[lineIndex];
        var regex = Pattern.compile("\\b" + Pattern.quote(symbol) + "\\b");
        var wordIdx = 0;
        var matcher = regex.matcher(rawLine);
        if (matcher.find()) {
            wordIdx = matcher.start() + 1;
        }

        if (wordResolver != null) {
            wordIdx = wordResolver.resolve(wordIdx, rawLine, symbol);
        }

        var line = new StartEnd(lineIndex + 1, lineIndex + 1);
        var column = new StartEnd(wordIdx, wordIdx + symbol.length());
        return new ErrorProperties(line, column, message);
    }

    public void raiseSchemaVersionRequired(int lineIndex, String symbol) {
        var errorProperties = buildErrorProperties("schema version required", lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.SchemaVersionRequired);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseInvalidSchemaVersion(int lineIndex, String symbol) {
        var errorProperties = buildErrorProperties("invalid schema " + symbol, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.InvalidSchema);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    private void raiseReservedTypeName(int lineIndex, String symbol) {
        var errorProperties = buildErrorProperties("a type cannot be named '" + Keyword.SELF + "' or '" + Keyword.THIS + "'.", lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.ReservedTypeKeywords);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    private void raiseInvalidName(int lineIndex, String symbol, String clause) {
        raiseInvalidName(lineIndex, symbol, clause, null);
    }

    private void raiseInvalidName(int lineIndex, String symbol, String clause, String typeName) {
        var messageStart = typeName != null
                ? "relation '" + symbol + "' of type '" + typeName + "'"
                : "type '" + symbol + "'";
        var message = messageStart + " does not match naming rule: '" + clause + "'.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.InvalidName);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    private void raiseReservedRelationName(int lineIndex, String symbol) {
        var errorProperties = buildErrorProperties("a relation cannot be named '" + Keyword.SELF + "' or '" + Keyword.THIS + "'.", lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.ReservedRelationKeywords);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    private void raiseDuplicateRelationName(int lineIndex, String symbol) {
    }

    private void raiseInvalidRelationError(int lineIndex, String symbol, Collection<String> validRelations) {
        var invalid = !validRelations.contains(symbol);
        if (invalid) {
            var message = "the relation `" + symbol + "` does not exist.";
            var errorProperties = buildErrorProperties(message, lineIndex, symbol);
            var metadata = new ValidationMetadata(symbol, ValidationError.ReservedRelationKeywords);
            errors.add(new ModelValidationSingleError(errorProperties, metadata));
        }
    }

    private void raiseAssignableRelationMustHaveTypes(int lineIndex, String symbol) {
        var rawLine = lines[lineIndex];
        var actualValue = rawLine.contains("[")
                ? rawLine.substring(rawLine.indexOf('['), rawLine.lastIndexOf(']') + 1)
                : "self";
        var message = "assignable relation '" + actualValue + "' must have types";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.AssignableRelationsMustHaveType);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    private void raiseInvalidType(int lineIndex, String symbol, String typeName) {
        var message = "`" + typeName + "` is not a valid type.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.InvalidType);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    private void raiseInvalidConditionNameInParameter(int lineIndex, String symbol, String typeName, String relationName, String conditionName) {
        var message = "`" + conditionName + "` is not a defined condition in the model.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.ConditionNotDefined, relationName, typeName, null);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    private void raiseAssignableTypeWildcardRelation(int lineIndex, String symbol) {
        var message = "type restriction `" + symbol + "` cannot contain both wildcard and relation";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.TypeRestrictionCannotHaveWildcardAndRelation);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    private void raiseInvalidTypeRelation(int lineIndex, String symbol, String typeName, String relationName) {
        var message = "`" + relationName + "` is not a valid relation for `" + typeName + "`.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.InvalidRelationType, relationName, typeName, null);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    private void raiseTupleUsersetRequiresDirect(int lineIndex, String symbol) {
        var message = "`" + symbol + "` relation used inside from allows only direct relation.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol, (wordIndex, rawLine, value) -> {
            var clauseStartsAt = rawLine.indexOf("from") + "from".length() + 1;
            wordIndex = clauseStartsAt + rawLine.substring(clauseStartsAt).indexOf(value) + 1;
            return wordIndex;
        });
        var metadata = new ValidationMetadata(symbol, ValidationError.TuplesetNotDirect);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    private void raiseDuplicateTypeName(int lineIndex, String symbol) {
        var message = "the type `" + symbol + "` is a duplicate.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.DuplicatedError);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    private void raiseDuplicateTypeRestriction(int lineIndex, String symbol, String relationName) {
        var message = "the type restriction `" + symbol + "` is a duplicate in the relation `" + relationName + "`.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.DuplicatedError, symbol, null, null);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    private void raiseDuplicateType(int lineIndex, String symbol, String relationName) {
        var message = "the partial relation definition `" + symbol + "` is a duplicate in the relation `" + relationName + "`.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.DuplicatedError, symbol, null, null);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    private void raiseNoEntryPointLoop(int lineIndex, String symbol, String typeName) {
        var message = "`" + symbol + "` is an impossible relation for `" + typeName + "` (potential loop).";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.RelationNoEntrypoint, symbol, null, null);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    private void raiseNoEntryPoint(int lineIndex, String symbol, String typeName) {
        var message = "`" + symbol + "` is an impossible relation for `" + typeName + "` (no entrypoint).";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.RelationNoEntrypoint, symbol, null, null);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    private void raiseUnusedCondition(int lineIndex, String symbol) {
        var message = "`" + symbol + "` condition is not used in the model.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.ConditionNotUsed, null, null, symbol);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

}
