using System.Text.Json.Serialization;

namespace OpenFga.Language.Validation;

/// <summary>
/// Enumeration of rewrite types used in OpenFGA relation definitions.
/// Defines the different ways relations can be rewritten or computed.
/// </summary>
public enum RewriteType {
    /// <summary>Direct assignment rewrite type.</summary>
    [JsonPropertyName("direct")]
    Direct,

    /// <summary>Computed userset rewrite type.</summary>
    [JsonPropertyName("computed_userset")]
    ComputedUserset,

    /// <summary>Tuple to userset rewrite type.</summary>
    [JsonPropertyName("tuple_to_userset")]
    TupleToUserset
}