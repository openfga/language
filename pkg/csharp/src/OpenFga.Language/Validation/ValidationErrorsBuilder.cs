using System.Text.RegularExpressions;
using OpenFga.Language.Errors;

namespace OpenFga.Language.Validation;

internal class ValidationErrorsBuilder
{
    private readonly string[]? _lines;
    private readonly List<ModelValidationSingleError> _errors = new();

    public ValidationErrorsBuilder(string[]? lines)
    {
        _lines = lines;
    }

    private ErrorProperties BuildErrorProperties(string message, int lineIndex, string symbol)
    {
        return BuildErrorProperties(message, lineIndex, symbol, null);
    }

    private ErrorProperties BuildErrorProperties(string message, int lineIndex, string symbol, Func<int, string, string, int>? wordResolver)
    {
        var properties = new ErrorProperties(null, null, message);

        if (_lines != null && lineIndex < _lines.Length)
        {
            var rawLine = _lines[lineIndex];
            var regex = new Regex($@"\b{Regex.Escape(symbol)}((?=\W)|$)");
            var wordIdx = 0;
            var match = regex.Match(rawLine);
            if (!string.IsNullOrEmpty(symbol) && match.Success)
            {
                wordIdx = match.Index;
            }

            if (wordResolver != null)
            {
                wordIdx = wordResolver(wordIdx, rawLine, symbol);
            }

            properties.Line = new StartEnd(lineIndex, lineIndex);
            properties.Column = new StartEnd(wordIdx, wordIdx + symbol.Length);
        }

        return properties;
    }

    public void RaiseSchemaVersionRequired(int lineIndex, string symbol)
    {
        var errorProperties = BuildErrorProperties("schema version required", lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.SchemaVersionRequired);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseInvalidSchemaVersion(int lineIndex, string symbol)
    {
        var errorProperties = BuildErrorProperties("invalid schema " + symbol, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.InvalidSchema);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseReservedTypeName(int lineIndex, string symbol)
    {
        var errorProperties = BuildErrorProperties(
            $"a type cannot be named '{Keyword.Self}' or '{Keyword.This}'.", lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.ReservedTypeKeywords);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseInvalidName(int lineIndex, string symbol, string clause)
    {
        RaiseInvalidName(lineIndex, symbol, clause, null);
    }

    public void RaiseInvalidName(int lineIndex, string symbol, string clause, string? typeName)
    {
        var messageStart = typeName != null ? $"relation '{symbol}' of type '{typeName}'" : $"type '{symbol}'";
        var message = $"{messageStart} does not match naming rule: '{clause}'.";
        var errorProperties = BuildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.InvalidName);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseReservedRelationName(int lineIndex, string symbol)
    {
        var errorProperties = BuildErrorProperties(
            $"a relation cannot be named '{Keyword.Self}' or '{Keyword.This}'.", lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.ReservedRelationKeywords);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseDuplicateRelationName(int lineIndex, string symbol)
    {
        // Empty implementation as per Java version
    }

    public void RaiseInvalidRelationError(int lineIndex, string symbol, IEnumerable<string> validRelations)
    {
        var invalid = !validRelations.Contains(symbol);
        if (invalid)
        {
            var message = $"the relation `{symbol}` does not exist.";
            var errorProperties = BuildErrorProperties(message, lineIndex, symbol);
            var metadata = new ValidationMetadata(symbol, ValidationError.MissingDefinition);
            _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
        }
    }

    public void RaiseAssignableRelationMustHaveTypes(int lineIndex, string symbol)
    {
        var rawLine = _lines?[lineIndex] ?? "";
        var actualValue = rawLine.Contains('[') 
            ? rawLine.Substring(rawLine.IndexOf('['), rawLine.LastIndexOf(']') - rawLine.IndexOf('[') + 1) 
            : "self";
        var message = $"assignable relation '{actualValue}' must have types";
        var errorProperties = BuildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.AssignableRelationsMustHaveType);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseInvalidType(int lineIndex, string symbol, string typeName)
    {
        var message = $"`{typeName}` is not a valid type.";
        var errorProperties = BuildErrorProperties(message, lineIndex, symbol, (wordIdx, rawLine, type) =>
        {
            // Split line at definition as InvalidType should mark the value, not the key
            var splitLine = rawLine.Split(':');
            return splitLine[0].Length + splitLine[1].IndexOf(typeName) + 1;
        });
        var metadata = new ValidationMetadata(symbol, ValidationError.InvalidType);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseInvalidConditionNameInParameter(int lineIndex, string symbol, string typeName, string relationName, string conditionName)
    {
        var message = $"`{conditionName}` is not a defined condition in the model.";
        var errorProperties = BuildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.ConditionNotDefined, relationName, typeName, null);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseAssignableTypeWildcardRelation(int lineIndex, string symbol)
    {
        var message = $"type restriction `{symbol}` cannot contain both wildcard and relation";
        var errorProperties = BuildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.TypeRestrictionCannotHaveWildcardAndRelation);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseInvalidTypeRelation(int lineIndex, string symbol, string typeName, string relationName)
    {
        var message = $"`{relationName}` is not a valid relation for `{typeName}`.";
        var errorProperties = BuildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.InvalidRelationType, relationName, typeName, null);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseInvalidRelationOnTupleset(int lineIndex, string symbol, string typeName, string typeDef, string relationName, string offendingRelation, string parent)
    {
        var message = $"the `{offendingRelation}` relation definition on type `{typeDef}` is not valid: `{offendingRelation}` does not exist on `{parent}`, which is of type `{typeName}`.";
        var errorProperties = BuildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.InvalidRelationOnTupleset, relationName, typeName, null);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseTupleUsersetRequiresDirect(int lineIndex, string symbol)
    {
        var message = $"`{symbol}` relation used inside from allows only direct relation.";
        var errorProperties = BuildErrorProperties(message, lineIndex, symbol, (wordIndex, rawLine, value) =>
        {
            var clauseStartsAt = rawLine.IndexOf("from") + "from".Length;
            wordIndex = clauseStartsAt + rawLine.Substring(clauseStartsAt).IndexOf(value);
            return wordIndex;
        });
        var metadata = new ValidationMetadata(symbol, ValidationError.TuplesetNotDirect);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseDuplicateTypeName(int lineIndex, string symbol)
    {
        var message = $"the type `{symbol}` is a duplicate.";
        var errorProperties = BuildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.DuplicatedError);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseDuplicateTypeRestriction(int lineIndex, string symbol, string relationName)
    {
        var message = $"the type restriction `{symbol}` is a duplicate in the relation `{relationName}`.";
        var errorProperties = BuildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.DuplicatedError, symbol, null, null);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseDuplicateType(int lineIndex, string symbol, string relationName)
    {
        var message = $"the partial relation definition `{symbol}` is a duplicate in the relation `{relationName}`.";
        var errorProperties = BuildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.DuplicatedError, symbol, null, null);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseThisNotInFirstPlace(int lineIndex, string relationName)
    {
        var message = $"this must be the first element in relation definition `{relationName}`.";
        var errorProperties = BuildErrorProperties(message, lineIndex, relationName);
        var metadata = new ValidationMetadata(relationName, ValidationError.ThisNotInFirstPlace, null, null, null);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseNoEntryPointLoop(int lineIndex, string symbol, string typeName)
    {
        var message = $"`{symbol}` is an impossible relation for `{typeName}` (potential loop).";
        var errorProperties = BuildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.RelationNoEntrypoint, symbol, null, null);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseNoEntryPoint(int lineIndex, string symbol, string typeName)
    {
        var message = $"`{symbol}` is an impossible relation for `{typeName}` (no entrypoint).";
        var errorProperties = BuildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.RelationNoEntrypoint, symbol, null, null);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseUnusedCondition(int lineIndex, string symbol)
    {
        var message = $"`{symbol}` condition is not used in the model.";
        var errorProperties = BuildErrorProperties(message, lineIndex, symbol);
        var metadata = new ValidationMetadata(symbol, ValidationError.ConditionNotUsed, null, null, symbol);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseDifferentNestedConditionName(string conditionKey, string nestedConditionName)
    {
        var message = $"condition key is `{conditionKey}` but nested name property is {nestedConditionName}";
        var errorProperties = BuildErrorProperties(message, 0, nestedConditionName);
        var metadata = new ValidationMetadata(nestedConditionName, ValidationError.DifferentNestedConditionName, null, null, null);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public void RaiseMultipleModulesInSingleFile(string file, IEnumerable<string> modules)
    {
        var modulesString = string.Join(", ", modules);
        var message = $"file {file} would contain multiple module definitions ({modulesString}) when transforming to DSL. Only one module can be defined per file.";
        var errorProperties = BuildErrorProperties(message, 0, file);
        var metadata = new ValidationMetadata(file, ValidationError.MultipleModulesInFile, null, null, null);
        _errors.Add(new ModelValidationSingleError(errorProperties, metadata));
    }

    public bool IsEmpty => _errors.Count == 0;

    public void ThrowIfNotEmpty()
    {
        if (!IsEmpty)
        {
            throw new DslErrorsException(_errors);
        }
    }
}
