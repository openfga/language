using System.Text.Json.Serialization;
using SharpYaml.Serialization;
using OpenFga.Language.Errors;

namespace OpenFga.Language.Tests.util;

public class FgaModTestCase
{
    [YamlMember("name")]
    public string Name { get; set; } = string.Empty;

    [YamlMember("modFile")]
    public string ModFile { get; set; } = string.Empty;

    [YamlMember("json")]
    public string? Json { get; set; }

    [YamlMember("skip")]
    public bool Skip { get; set; }

    [YamlMember("expected_errors")]
    public List<ModFileValidationSingleError> ExpectedErrors { get; set; } = new();
}
