using System.Text.Json.Serialization;
using OpenFga.Language.Validation;

namespace OpenFga.Language.Tests.util;

public class DslSyntaxTestCase
{
    [JsonPropertyName("name")]
    public string Name { get; set; } = string.Empty;

    [JsonPropertyName("dsl")]
    public string Dsl { get; set; } = string.Empty;

    [JsonPropertyName("valid")]
    public bool Valid { get; set; }

    [JsonPropertyName("skip")]
    public bool Skip { get; set; }

    [JsonPropertyName("expected_errors")]
    public List<ModelValidationSingleError> ExpectedErrors { get; set; } = new();
}
