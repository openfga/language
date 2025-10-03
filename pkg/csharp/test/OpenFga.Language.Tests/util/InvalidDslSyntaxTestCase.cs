using System.Text.Json.Serialization;
using YamlDotNet.Serialization;

namespace OpenFga.Language.Tests.util;

public class InvalidDslSyntaxTestCase
{
    [YamlMember(Alias = "name")]
    public string Name { get; set; } = string.Empty;

    [YamlMember(Alias = "dsl")]
    public string Dsl { get; set; } = string.Empty;

    [YamlMember(Alias = "skip")]
    public bool Skip { get; set; }

    [YamlMember(Alias = "error_message")]
    public string? ErrorMessage { get; set; }
}
