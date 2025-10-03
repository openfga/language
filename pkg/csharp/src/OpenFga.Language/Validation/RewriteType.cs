using System.Text.Json.Serialization;

namespace OpenFga.Language.Validation;

public enum RewriteType {
    [JsonPropertyName("direct")]
    Direct,

    [JsonPropertyName("computed_userset")]
    ComputedUserset,

    [JsonPropertyName("tuple_to_userset")]
    TupleToUserset
}