using System.Text.Json.Serialization;

namespace OpenFga.Language.Tests.util;

public sealed class JsonSyntaxTestCase
{
    [JsonPropertyName("name")]
    public string Name { get; set; } = string.Empty;

    [JsonPropertyName("json")]
    public string Json { get; set; } = string.Empty;

    [JsonPropertyName("error_message")]
    public string? ErrorMessage { get; set; }

    [JsonPropertyName("skip")]
    public bool Skip { get; set; }

    [JsonPropertyName("valid")]
    public bool Valid { get; set; }
}
