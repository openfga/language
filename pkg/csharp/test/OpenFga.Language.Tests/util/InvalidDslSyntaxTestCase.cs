using System.Text.Json.Serialization;
using SharpYaml.Serialization;

namespace OpenFga.Language.Tests.util;

public class InvalidDslSyntaxTestCase
{
    [YamlMember("name")]
    public string Name { get; set; } = string.Empty;

    [YamlMember("dsl")]
    public string Dsl { get; set; } = string.Empty;

    [YamlMember("skip")]
    public bool Skip { get; set; }

    [YamlMember("error_message")]
    public string? ErrorMessage { get; set; }
}
