using System.Text.Json.Serialization;

namespace OpenFga.Language.Tests.util;

public sealed class ValidTransformerTestCase
{
    [JsonPropertyName("name")]
    public string Name { get; set; } = string.Empty;

    [JsonPropertyName("dsl")]
    public string Dsl { get; set; } = string.Empty;

    [JsonPropertyName("json")]
    public string Json { get; set; } = string.Empty;

    [JsonPropertyName("skip")]
    public bool Skip { get; set; }

    public ValidTransformerTestCase() { }

    public ValidTransformerTestCase(string name, string dsl, string json, bool skip)
    {
        Name = name;
        Dsl = dsl;
        Json = json;
        Skip = skip;
    }

    public override string ToString()
    {
        return $"TransformerTestCase[name={Name}, dsl={Dsl}, json={Json}, skip={Skip}]";
    }
}
