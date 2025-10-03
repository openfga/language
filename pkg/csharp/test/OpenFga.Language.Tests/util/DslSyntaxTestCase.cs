using OpenFga.Language.Errors;
using OpenFga.Language.Validation;
using SharpYaml.Serialization;
using System.Text.Json.Serialization;

namespace OpenFga.Language.Tests.util;

public class DslSyntaxTestCase {
    [YamlMember("name")]
    public string Name { get; set; } = string.Empty;

    [YamlMember("dsl")]
    public string Dsl { get; set; } = string.Empty;

    [YamlMember("valid")]
    public bool Valid { get; set; }

    [YamlMember("skip")]
    public bool Skip { get; set; }

    [YamlMember("expected_errors")]
    public List<ModelValidationSingleError> ExpectedErrors { get; set; } = new();
}