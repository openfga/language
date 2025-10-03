namespace OpenFga.Language.Validation;

internal class InvalidChildRelationMetadata
{
    public int LineIndex { get; }
    public string Symbol { get; }
    public string TypeName { get; }
    public string RelationName { get; }
    public string Parent { get; }

    public InvalidChildRelationMetadata(int lineIndex, string symbol, string typeName, string relationName, string parent)
    {
        LineIndex = lineIndex;
        Symbol = symbol;
        TypeName = typeName;
        RelationName = relationName;
        Parent = parent;
    }
}
