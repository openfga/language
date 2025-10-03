using OpenFga.Sdk.Model;

namespace OpenFga.Language;

/// <summary>
/// Helper class to store relation information on the stack during parsing.
/// Used internally during DSL parsing to manage relation definitions and operators.
/// </summary>
/// <param name="rewrites">The list of usersets for this relation</param>
/// <param name="operator">The operator used in the relation definition</param>
public class StackRelation(List<Userset> rewrites, string? @operator) {
    /// <summary>
    /// Gets or sets the list of usersets for this relation.
    /// </summary>
    public List<Userset> Rewrites { get; set; } = rewrites;

    /// <summary>
    /// Gets or sets the operator used in the relation definition.
    /// </summary>
    public string? Operator { get; set; } = @operator;
}