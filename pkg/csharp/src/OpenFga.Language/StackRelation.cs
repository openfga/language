using OpenFga.Sdk.Model;

namespace OpenFga.Language;

/// <summary>
/// Helper class to store relation information on the stack during parsing
/// </summary>
public class StackRelation(List<Userset> rewrites, string? @operator)
{
    public List<Userset> Rewrites { get; set; } = rewrites;
    public string? Operator { get; set; } = @operator;
}