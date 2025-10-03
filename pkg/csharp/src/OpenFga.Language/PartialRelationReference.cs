using OpenFga.Sdk.Model;

namespace OpenFga.Language;

class PartialRelationReference
{
    public string? Type { get; set; }
    public string? Relation { get; set; }
    public object? Wildcard { get; set; }
    public string? Condition { get; set; }

    public RelationReference AsRelationReference()
    {
        return new RelationReference()
        {
            Type = Type!,
            Relation = Relation,
            Wildcard = Wildcard,
            Condition = Condition,
        };
    }
}
