package dev.openfga.language.validation;

import dev.openfga.language.errors.*;
import java.util.ArrayList;
import java.util.Collection;
import java.util.List;
import java.util.Set;
import java.util.regex.Pattern;

class ValidationErrorsBuilder {

    private final String[] lines;
    private final List<ModelValidationSingleError> errors = new ArrayList<>();

    ValidationErrorsBuilder(String[] lines) {
        this.lines = lines;
    }

    private ErrorProperties buildErrorProperties(String message, int lineIndex, String symbol) {
        return buildErrorProperties(message, lineIndex, symbol, null);
    }

    private ErrorProperties buildErrorProperties(
            String message, int lineIndex, String symbol, WordResolver wordResolver) {

        var properties = new ErrorProperties(null, null, message);

        if (lines != null) {
            var rawLine = lines[lineIndex];
            var regex = Pattern.compile("\\b" + Pattern.quote(symbol) + "((?=[\\b\\s,\\]])|$)");
            var wordIdx = 0;
            var matcher = regex.matcher(rawLine);
            if (matcher.find()) {
                wordIdx = matcher.start();
            }

            if (wordResolver != null) {
                wordIdx = wordResolver.resolve(wordIdx, rawLine, symbol);
            }

            properties.setLine(new StartEnd(lineIndex, lineIndex));
            properties.setColumn(new StartEnd(wordIdx, wordIdx + symbol.length()));
        }

        return properties;
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

    public void raiseReservedTypeName(int lineIndex, String symbol) {
        var errorProperties = buildErrorProperties(
                "a type cannot be named '" + Keyword.SELF + "' or '" + Keyword.THIS + "'.", lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.ReservedTypeKeywords);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseInvalidName(int lineIndex, String symbol, String clause) {
        raiseInvalidName(lineIndex, symbol, clause, null);
    }

    public void raiseInvalidName(int lineIndex, String symbol, String clause, String typeName) {
        var messageStart =
                typeName != null ? "relation '" + symbol + "' of type '" + typeName + "'" : "type '" + symbol + "'";
        var message = messageStart + " does not match naming rule: '" + clause + "'.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.InvalidName);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseReservedRelationName(int lineIndex, String symbol) {
        var errorProperties = buildErrorProperties(
                "a relation cannot be named '" + Keyword.SELF + "' or '" + Keyword.THIS + "'.", lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.ReservedRelationKeywords);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseDuplicateRelationName(int lineIndex, String symbol) {}

    public void raiseInvalidRelationError(int lineIndex, String symbol, Collection<String> validRelations) {
        var invalid = !validRelations.contains(symbol);
        if (invalid) {
            var message = "the relation `" + symbol + "` does not exist.";
            var errorProperties = buildErrorProperties(message, lineIndex, symbol);
            var metadata = new ValidationMetadata(symbol, ValidationError.MissingDefinition);
            errors.add(new ModelValidationSingleError(errorProperties, metadata));
        }
    }

    public void raiseAssignableRelationMustHaveTypes(int lineIndex, String symbol) {
        var rawLine = lines[lineIndex];
        var actualValue =
                rawLine.contains("[") ? rawLine.substring(rawLine.indexOf('['), rawLine.lastIndexOf(']') + 1) : "self";
        var message = "assignable relation '" + actualValue + "' must have types";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.AssignableRelationsMustHaveType);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseInvalidType(int lineIndex, String symbol, String typeName) {
        var message = "`" + typeName + "` is not a valid type.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol, (wordIdx, rawLine, type) -> {
            // Split line at definition as InvalidType should mark the value, not the key
            var splitLine = rawLine.split(":");
            return splitLine[0].length() + splitLine[1].indexOf(typeName) + 1;
        });
        var metadata = new ValidationMetadata(symbol, ValidationError.InvalidType);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseInvalidConditionNameInParameter(
            int lineIndex, String symbol, String typeName, String relationName, String conditionName) {
        var message = "`" + conditionName + "` is not a defined condition in the model.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata =
                new ValidationMetadata(symbol, ValidationError.ConditionNotDefined, relationName, typeName, null);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseAssignableTypeWildcardRelation(int lineIndex, String symbol) {
        var message = "type restriction `" + symbol + "` cannot contain both wildcard and relation";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.TypeRestrictionCannotHaveWildcardAndRelation);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseInvalidTypeRelation(int lineIndex, String symbol, String typeName, String relationName) {
        var message = "`" + relationName + "` is not a valid relation for `" + typeName + "`.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata =
                new ValidationMetadata(symbol, ValidationError.InvalidRelationType, relationName, typeName, null);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseInvalidRelationOnTupleset(
            int lineIndex,
            String symbol,
            String typeName,
            String typeDef,
            String relationName,
            String offendingRelation,
            String parent) {
        var message = "the `" + offendingRelation + "` relation definition on type `" + typeDef + "` is not valid: `"
                + offendingRelation + "` does not exist on `" + parent + "`, which is of type `" + typeName + "`.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata =
                new ValidationMetadata(symbol, ValidationError.InvalidRelationOnTupleset, relationName, typeName, null);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseTupleUsersetRequiresDirect(int lineIndex, String symbol) {
        var message = "`" + symbol + "` relation used inside from allows only direct relation.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol, (wordIndex, rawLine, value) -> {
            var clauseStartsAt = rawLine.indexOf("from") + "from".length();
            wordIndex = clauseStartsAt + rawLine.substring(clauseStartsAt).indexOf(value);
            return wordIndex;
        });
        var metadata = new ValidationMetadata(symbol, ValidationError.TuplesetNotDirect);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseDuplicateTypeName(int lineIndex, String symbol) {
        var message = "the type `" + symbol + "` is a duplicate.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.DuplicatedError);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseDuplicateTypeRestriction(int lineIndex, String symbol, String relationName) {
        var message = "the type restriction `" + symbol + "` is a duplicate in the relation `" + relationName + "`.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.DuplicatedError, symbol, null, null);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseDuplicateType(int lineIndex, String symbol, String relationName) {
        var message = "the partial relation definition `" + symbol + "` is a duplicate in the relation `" + relationName
                + "`.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.DuplicatedError, symbol, null, null);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseNoEntryPointLoop(int lineIndex, String symbol, String typeName) {
        var message = "`" + symbol + "` is an impossible relation for `" + typeName + "` (potential loop).";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.RelationNoEntrypoint, symbol, null, null);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseNoEntryPoint(int lineIndex, String symbol, String typeName) {
        var message = "`" + symbol + "` is an impossible relation for `" + typeName + "` (no entrypoint).";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.RelationNoEntrypoint, symbol, null, null);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseUnusedCondition(int lineIndex, String symbol) {
        var message = "`" + symbol + "` condition is not used in the model.";
        var errorProperties = buildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.ConditionNotUsed, null, null, symbol);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseDifferentNestedConditionName(String conditionKey, String nestedConditionName) {
        var message = "condition key is `" + conditionKey + "` but nested name property is " + nestedConditionName;
        var errorProperties = buildErrorProperties(message, 0, nestedConditionName);
        var metadata = new ValidationMetadata(
                nestedConditionName, ValidationError.DifferentNestedConditionName, null, null, null);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void raiseMultipleModulesInSingleFile(String file, Set<String> modules) {
        var modulesString = String.join(", ", modules);
        var message = "file " + file + " would contain multiple module definitions (" + modulesString
                + ") when transforming to DSL. Only one module can be defined per file.";
        var errorProperties = buildErrorProperties(message, 0, file);
        var metadata = new ValidationMetadata(file, ValidationError.MultipleModulesInFile, null, null, null);
        errors.add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public boolean isEmpty() {
        return errors.isEmpty();
    }

    public void throwIfNotEmpty() throws DslErrorsException {
        if (!errors.isEmpty()) {
            throw new DslErrorsException(errors);
        }
    }
}
