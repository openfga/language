namespace OpenFga.Language.Errors;

public class UnsupportedDslNestingException(string typeName, string relationName) 
    : Exception($"the '{relationName}' relation definition under the '{typeName}' type is not supported by the OpenFGA DSL syntax yet");
