using OpenFga.Language.Validation;
using System.Text.Json.Serialization;
using OpenFga.Language.Errors;
using YamlDotNet.Serialization;

namespace OpenFga.Language.Tests.util;

public class MultipleInvalidDslSyntaxTestCase : InvalidDslSyntaxTestCase
{
    [YamlMember(Alias = "expected_errors")]
    public List<ModelValidationSingleError> ExpectedErrors { get; set; } = new();
}
