using Microsoft.VisualStudio.TestPlatform.CommunicationUtilities;
using OpenFga.Language.Errors;
using OpenFga.Language.ModFile;
using OpenFga.Language.Tests.util;
using OpenFga.Language.Transformers;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Serialization;
using Xunit;

namespace OpenFga.Language.Tests;

public class FgaModToJsonShould
{
    private static readonly JsonSerializerOptions JsonOptions = new()
    {
        PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
        WriteIndented = false,
        DefaultIgnoreCondition = JsonIgnoreCondition.Never,
        PropertyNameCaseInsensitive = true
    };

    [Theory]
    [MemberData(nameof(FgaModTestCases))]
    public void TransformFgaMod(
        string name, 
        string fgaMod, 
        string? json, 
        bool skip, 
        List<ModFileValidationSingleError> expectedErrors)
    {
        // Skip test if marked as skip
        if (skip)
        {
            return;
        }

        if (json != null)
        {
            var generatedJson = FgaModTransformer.Transform(fgaMod);
            var expected = JsonSerializer.Serialize(JsonSerializer.Deserialize<FgaModFile>(json, JsonOptions), JsonOptions);

            var different = expected != generatedJson;
            Assert.Equal(expected, generatedJson);
        }
        else
        {
            var exception = Assert.Throws<ModFileValidationError>(() => FgaModTransformer.Transform(fgaMod));

            var errorsCount = expectedErrors.Count;
            var formattedErrors = expectedErrors
                .Select(error => $"validation error at line={error.Line.Start}, column={error.Column.Start}: {error.Message}")
                .ToList();

            var expectedMessage = $"{errorsCount} error{(errorsCount > 1 ? "s" : "")} occurred:\n\t* {string.Join("\n\t* ", formattedErrors)}\n\n";

            var dif = expectedMessage != exception.Message;
            Assert.Equal(expectedMessage, exception.Message);
        }
    }

    public static IEnumerable<object[]> FgaModTestCases()
    {
        return TestsData.FgaModTransformTestCases
            .Select(testCase => new object[]
            {
                testCase.Name,
                testCase.ModFile,
                testCase.Json,
                testCase.Skip,
                testCase.ExpectedErrors
            });
    }
}
