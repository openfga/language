using OpenFga.Language.Validation;
using System.Text.Json.Serialization;
using OpenFga.Language.Errors;
using YamlDotNet.Serialization;

namespace OpenFga.Language.Tests.util;

public class DslSyntaxTestCase
{
    [YamlMember(Alias = "name")]
    public string Name { get; set; } = string.Empty;

    [YamlMember(Alias = "dsl")]
    public string Dsl { get; set; } = string.Empty;

    [YamlMember(Alias = "valid")]
    public bool Valid { get; set; }

    [YamlMember(Alias = "skip")]
    public bool Skip { get; set; }

    [YamlMember(Alias = "expected_errors")]
    public List<ModelValidationSingleError> ExpectedErrors { get; set; } = new();
}
