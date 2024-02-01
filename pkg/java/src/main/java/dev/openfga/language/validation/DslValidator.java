package dev.openfga.language.validation;

import dev.openfga.language.DslToJsonTransformer;
import dev.openfga.language.errors.*;
import dev.openfga.sdk.api.model.*;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.RequiredArgsConstructor;

import java.io.IOException;
import java.util.*;
import java.util.function.Predicate;
import java.util.regex.Pattern;
import java.util.stream.IntStream;

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

        authorizationModel.getTypeDefinitions().forEach(typeDef -> {
            var typeName = typeDef.getType();
            typeDef.getRelations().forEach((relationName, relationDef) -> {
                relationDefined(typeMap, typeName, relationName);
            });
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
                break;
            }
        }

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
                0);
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

}
