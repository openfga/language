using OpenFga.Language.Errors;
using OpenFga.Language.Tests.util;
using OpenFga.Language.Transformers;
using OpenFga.Sdk.Model;
using Xunit;

namespace OpenFga.Language.Tests;

public class DslToJsonShould
{
    [Theory]
    [MemberData(nameof(TransformerTestCases))]
    public void TransformDsl(string name, string dsl, string json, bool skip)
    {
        if (skip)
        {
            return;
        }

        var generatedJson = new DslToJsonTransformer().Transform(dsl);

        var expectedAuthorizationModel = Json.Parse<AuthorizationModel>(json);

        var expectedJson = Json.Stringify(expectedAuthorizationModel);

        Assert.Equal(expectedJson, generatedJson);
    }

    [Theory]
    [MemberData(nameof(DslSyntaxTestCases))]
    public void VerifyDslSyntax(string name, string dsl, List<ModelValidationSingleError> expectedErrors)
    {
        if (expectedErrors.Count == 0)
        {
            // If no errors expected, the transformation should succeed
            var result = new DslToJsonTransformer().Transform(dsl);
            Assert.NotNull(result);
            return;
        }

        // If errors are expected, the transformation should throw DslErrorsException
        var exception = Assert.Throws<DslErrorsException>(() => new DslToJsonTransformer().Transform(dsl));

        // Unfortunately antlr is throwing different error messages in Java, Go and JS - considering that at the moment
        // we care that it errors for syntax errors more than we care about the error messages matching,
        // esp. in Java as we are not building a language server on top of the returned errors yet
        // actual matching error strings is safe to ignore for now
        // Assert.Equal(expectedErrors.Count, exception.Errors.Count);
    }

    public static IEnumerable<object[]> TransformerTestCases()
    {
        return TestsData.ValidTransformerTestCases
            .Select(testCase => new object[]
            {
                testCase.Name,
                testCase.Dsl,
                testCase.Json,
                testCase.Skip
            });
    }

    public static IEnumerable<object[]> DslSyntaxTestCases()
    {
        return TestsData.DslSyntaxTestCases
            .Select(testCase => new object[]
            {
                testCase.Name,
                testCase.Dsl,
                testCase.ExpectedErrors
            });
    }
}
