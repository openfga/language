using System.Text.Json.Serialization;
using OpenFga.Language.Errors;

namespace OpenFga.Language.Tests.util;

public class FgaModTestCase
{
    [JsonPropertyName("name")]
    public string Name { get; set; } = string.Empty;

    [JsonPropertyName("modFile")]
    public string ModFile { get; set; } = string.Empty;

    [JsonPropertyName("json")]
    public string Json { get; set; } = string.Empty;

    [JsonPropertyName("skip")]
    public bool Skip { get; set; }

    [JsonPropertyName("expected_errors")]
    public List<ModFileValidationSingleError> ExpectedErrors { get; set; } = new();
}
