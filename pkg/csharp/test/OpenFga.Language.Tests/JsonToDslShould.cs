using OpenFga.Language.Tests.util;
using Xunit;
#pragma warning disable xUnit1026
#pragma warning disable IDE0060

namespace OpenFga.Language.Tests;

public class JsonToDslShould {
    [Theory]
    [MemberData(nameof(TransformerTestCases))]
    public void TransformJson(string name, string dsl, string json, bool skip) {
        if (skip) {
            return;
        }

        var generatedDsl = new JsonToDslTransformer().Transform(json);

        Assert.Equal(dsl, generatedDsl);
    }

    [Theory]
    [MemberData(nameof(InvalidJsonSyntaxTestCases))]
    public void ThrowAnExceptionWhenTransformingInvalidJsonToDsl(
        string name, string json, string? errorMessage, bool skip) {
        if (skip) {
            return;
        }

        if (errorMessage == null) {
            // If no error message expected, the transformation should succeed
            var result = new JsonToDslTransformer().Transform(json);
            Assert.NotNull(result);
        }
        else {
            // If error message expected, the transformation should throw an exception
            var exception = Assert.ThrowsAny<Exception>(() => new JsonToDslTransformer().Transform(json));
            Assert.Equal(errorMessage, exception.Message);
        }
    }

    public static IEnumerable<object[]> TransformerTestCases() {
        return TestsData.ValidTransformerTestCases
            .Select(testCase => new object[]
            {
                testCase.Name,
                testCase.Dsl,
                testCase.Json,
                testCase.Skip
            });
    }

    public static IEnumerable<object[]> InvalidJsonSyntaxTestCases() {
        return TestsData.JsonSyntaxTestCases
            .Select(testCase => new object[]
            {
                testCase.Name,
                testCase.Json,
                testCase.ErrorMessage!,
                testCase.Skip
            });
    }
}