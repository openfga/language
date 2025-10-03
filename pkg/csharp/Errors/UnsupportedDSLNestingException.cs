namespace OpenFgaLanguage.Errors;

public class UnsupportedDSLNestingException : Exception
{
    public UnsupportedDSLNestingException(string typeName, string relationName) 
        : base($"the '{relationName}' relation definition under the '{typeName}' type is not supported by the OpenFGA DSL syntax yet")
    {
    }
}
