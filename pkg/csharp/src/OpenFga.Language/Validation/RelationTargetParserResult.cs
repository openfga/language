namespace OpenFga.Language.Validation;

/// <summary>
/// Represents the result of parsing a relation target definition.
/// Contains information about the target relation, source relation, and rewrite type.
/// </summary>
public class RelationTargetParserResult {
    /// <summary>
    /// Gets the target relation name.
    /// </summary>
    public string? Target { get; }

    /// <summary>
    /// Gets the source relation name (for tuple-to-userset rewrites).
    /// </summary>
    public string? From { get; }

    /// <summary>
    /// Gets the rewrite type for this relation target.
    /// </summary>
    public RewriteType Rewrite { get; }

    /// <summary>
    /// Initializes a new instance of the RelationTargetParserResult class.
    /// </summary>
    /// <param name="target">The target relation name</param>
    /// <param name="from">The source relation name</param>
    /// <param name="rewrite">The rewrite type</param>
    public RelationTargetParserResult(string? target, string? from, RewriteType rewrite) {
        Target = target;
        From = from;
        Rewrite = rewrite;
    }
}