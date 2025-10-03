using System.Text.Json.Serialization;
using OpenFga.Language.Validation;

namespace OpenFga.Language.Tests.util;

public class MultipleInvalidDslSyntaxTestCase : InvalidDslSyntaxTestCase
{
    [JsonPropertyName("expected_errors")]
    public List<ModelValidationSingleError> ExpectedErrors { get; set; } = new();
}
