using OpenFga.Sdk.Model;

namespace OpenFga.Language;

/// <summary>
/// Helper class to hold relation information during parsing
/// This bridges the gap between the ANTLR parsing and the OpenFGA SDK model classes
/// </summary>
public class Relation(string? name, List<Userset> rewrites, string? @operator, RelationMetadata typeInfo) {
    public string? Name { get; set; } = name;
    public List<Userset> Rewrites { get; set; } = rewrites;
    public string? Operator { get; set; } = @operator;
    public RelationMetadata TypeInfo { get; set; } = typeInfo;
}