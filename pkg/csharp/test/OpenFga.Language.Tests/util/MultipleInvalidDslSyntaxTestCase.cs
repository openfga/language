using OpenFga.Language.Errors;
using OpenFga.Language.Validation;
using SharpYaml.Serialization;
using System.Text.Json.Serialization;

namespace OpenFga.Language.Tests.util;

public class MultipleInvalidDslSyntaxTestCase : InvalidDslSyntaxTestCase {
    [YamlMember("expected_errors")]
    public List<ModelValidationSingleError> ExpectedErrors { get; set; } = new();
}