namespace OpenFga.Language.Validation;

public class RelationTargetParserResult
{
    public string? Target { get; }
    public string? From { get; }
    public RewriteType Rewrite { get; }

    public RelationTargetParserResult(string? target, string? from, RewriteType rewrite)
    {
        Target = target;
        From = from;
        Rewrite = rewrite;
    }
}
