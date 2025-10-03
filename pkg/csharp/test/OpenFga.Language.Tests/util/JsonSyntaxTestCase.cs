using System.Text.Json.Serialization;
using YamlDotNet.Serialization;

namespace OpenFga.Language.Tests.util;

public sealed class JsonSyntaxTestCase
{
    [YamlMember(Alias = "name")]
    public string Name { get; set; } = string.Empty;

    [YamlMember(Alias = "json")]
    public string Json { get; set; } = string.Empty;

    [YamlMember(Alias = "error_message")]
    public string? ErrorMessage { get; set; }

    [YamlMember(Alias = "skip")]
    public bool Skip { get; set; }

    [YamlMember(Alias = "valid")]
    public bool Valid { get; set; }
}
