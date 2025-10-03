using OpenFga.Language.Validation;
using System.Text.Json.Serialization;
using OpenFga.Language.Errors;
using YamlDotNet.Serialization;

namespace OpenFga.Language.Tests.util;

public class JsonValidationTestCase
{
    [YamlMember(Alias = "name")]
    public string Name { get; set; } = string.Empty;

    [YamlMember(Alias ="json")]
    public string Json { get; set; } = string.Empty;

    [YamlMember(Alias ="expected_errors")]
    public List<ModelValidationSingleError>? ExpectedErrors { get; set; } = new();
}
