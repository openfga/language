using System.Text.Json.Serialization;
using OpenFga.Language.Validation;

namespace OpenFga.Language.Tests.util;

public class JsonValidationTestCase
{
    [JsonPropertyName("name")]
    public string Name { get; set; } = string.Empty;

    [JsonPropertyName("json")]
    public string Json { get; set; } = string.Empty;

    [JsonPropertyName("expected_errors")]
    public List<ModelValidationSingleError> ExpectedErrors { get; set; } = new();
}
