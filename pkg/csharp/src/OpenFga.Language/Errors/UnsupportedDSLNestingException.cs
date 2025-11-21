namespace OpenFga.Language.Errors;

/// <summary>
/// Exception thrown when attempting to transform a JSON authorization model to DSL format
/// that contains unsupported nesting patterns in relation definitions.
/// </summary>
/// <param name="typeName">The name of the type containing the unsupported relation</param>
/// <param name="relationName">The name of the relation with unsupported nesting</param>
public class UnsupportedDslNestingException(string typeName, string relationName)
    : Exception($"the '{relationName}' relation definition under the '{typeName}' type is not supported by the OpenFGA DSL syntax yet");