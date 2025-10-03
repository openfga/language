using SharpYaml.Serialization;
using System.Text.Json.Serialization;

namespace OpenFga.Language.Tests.util;

public sealed class JsonSyntaxTestCase {
    [YamlMember("name")]
    public string Name { get; set; } = string.Empty;

    [YamlMember("json")]
    public string Json { get; set; } = string.Empty;

    [YamlMember("error_message")]
    public string? ErrorMessage { get; set; }

    [YamlMember("skip")]
    public bool Skip { get; set; }

    [YamlMember("valid")]
    public bool Valid { get; set; }
}