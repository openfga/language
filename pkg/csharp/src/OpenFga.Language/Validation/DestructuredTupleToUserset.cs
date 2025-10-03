namespace OpenFga.Language.Validation;

internal class DestructuredTupleToUserset
{
    public string DecodedType { get; }
    public string? DecodedRelation { get; }
    public bool IsWildcard { get; }
    public string? DecodedConditionName { get; }

    public DestructuredTupleToUserset(string decodedType, string? decodedRelation, bool isWildcard, string? decodedConditionName)
    {
        DecodedType = decodedType;
        DecodedRelation = decodedRelation;
        IsWildcard = isWildcard;
        DecodedConditionName = decodedConditionName;
    }

    public static DestructuredTupleToUserset From(string allowableType)
    {
        var tupleAndCondition = allowableType.Split(" with ");
        var tupleString = tupleAndCondition[0];
        var decodedConditionName = tupleAndCondition.Length > 1 ? tupleAndCondition[1] : null;
        var isWildcard = tupleString.Contains(":*");
        var splittedWords = tupleString.Replace(":*", "").Split('#');
        return new DestructuredTupleToUserset(
            splittedWords[0], 
            splittedWords.Length > 1 ? splittedWords[1] : null, 
            isWildcard, 
            decodedConditionName);
    }
}
