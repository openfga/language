using System.Text.Json.Serialization;

namespace OpenFga.Language.Tests.util;

public class InvalidDslSyntaxTestCase
{
    [JsonPropertyName("name")]
    public string Name { get; set; } = string.Empty;

    [JsonPropertyName("dsl")]
    public string Dsl { get; set; } = string.Empty;

    [JsonPropertyName("skip")]
    public bool Skip { get; set; }

    [JsonPropertyName("error_message")]
    public string? ErrorMessage { get; set; }
}
