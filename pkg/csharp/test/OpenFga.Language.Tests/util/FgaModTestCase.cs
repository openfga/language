using System.Text.Json.Serialization;
using YamlDotNet.Serialization;
using OpenFga.Language.Errors;

namespace OpenFga.Language.Tests.util;

public class FgaModTestCase
{
    [YamlMember(Alias = "name")]
    public string Name { get; set; } = string.Empty;

    [YamlMember(Alias = "modFile")]
    [SharpYaml.Serialization.YamlMember("modFile")]
    public string ModFile { get; set; } = string.Empty;

    [YamlMember(Alias = "json")]
    public string? Json { get; set; }

    [YamlMember(Alias = "skip")]
    public bool Skip { get; set; }

    [YamlMember(Alias = "expected_errors")]
    public List<ModFileValidationSingleError> ExpectedErrors { get; set; } = new();
}
