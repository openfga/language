using OpenFga.Language.Errors;
using OpenFga.Language.Validation;
using SharpYaml.Serialization;
using System.Text.Json.Serialization;

namespace OpenFga.Language.Tests.util;

public class JsonValidationTestCase {
    [YamlMember("name")]
    public string Name { get; set; } = string.Empty;

    [YamlMember("json")]
    public string Json { get; set; } = string.Empty;

    [YamlMember("expected_errors")]
    public List<ModelValidationSingleError>? ExpectedErrors { get; set; } = new();
}