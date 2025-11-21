#pragma warning disable IDE0060
#pragma warning disable xUnit1026

using OpenFga.Language.Errors;
using OpenFga.Language.Tests.util;
using OpenFga.Language.Validation;
using OpenFga.Sdk.Model;
using System.Text.Json;
using Xunit;

namespace OpenFga.Language.Tests;

public class ModelValidatorShould {
    [Theory]
    [MemberData(nameof(DslSyntaxTestCases))]
    public void VerifyDslSyntax(string name, string dsl, List<ModelValidationSingleError> expectedErrors, bool skip) {
        if (skip) {
            return;
        }

        if (expectedErrors.Count == 0) {
            // If no errors expected, validation should succeed
            ModelValidator.ValidateDsl(dsl);
            return;
        }

        // If errors are expected, validation should throw DslErrorsException
        var exception = Assert.Throws<DslErrorsException>(() => ModelValidator.ValidateDsl(dsl));

        // Unfortunately antlr is throwing different error messages in Java, Go and JS - considering that at the moment
        // we care that it errors for syntax errors more than we care about the error messages matching,
        // esp. in Java as we are not building a language server on top of the returned errors yet
        // actual matching error strings is safe to ignore for now

        // var errorsCount = expectedErrors.Count;
        // var formattedErrors = expectedErrors
        //     .Select(error => $"syntax error at line={error.Line?.Start}, column={error.Column?.Start}: {error.Message}")
        //     .ToList();
        // var expectedMessage = $"{errorsCount} error{(errorsCount > 1 ? "s" : "")} occurred:\n\t* {string.Join("\n\t* ", formattedErrors)}\n\n";
        // Assert.Equal(expectedMessage, exception.Message);
        // var actualErrors = exception.Errors;
        // for (int i = 0; i < expectedErrors.Count; i++)
        // {
        //     var expectedError = expectedErrors[i];
        //     var actualError = actualErrors[i];
        //     AssertMatch(expectedError, actualError);
        // }
    }

    [Theory]
    [MemberData(nameof(DslValidationTestCases))]
    public void VerifyDslValidation(
        string name, string dsl, List<ModelValidationSingleError> expectedErrors, bool skip) {
        if (skip) {
            return;
        }

        if (expectedErrors.Count == 0) {
            // If no errors expected, validation should succeed
            ModelValidator.ValidateDsl(dsl);
            return;
        }

        // If errors are expected, validation should throw DslErrorsException
        var exception = Assert.Throws<DslErrorsException>(() => ModelValidator.ValidateDsl(dsl));

        var errorsCount = expectedErrors.Count;
        var formattedErrors = expectedErrors
            .Select(error => $"validation error at line={error.Line?.Start}, column={error.Column?.Start}: {error.Message}")
            .ToList();
        var expectedMessage = $"{errorsCount} error{(errorsCount > 1 ? "s" : "")} occurred:\n\t* {string.Join("\n\t* ", formattedErrors)}\n\n";

        Assert.Equal(expectedMessage, exception.Message);

        var actualErrors = exception.Errors;
        for (int i = 0; i < expectedErrors.Count; i++) {
            var expectedError = expectedErrors[i];
            var actualError = actualErrors[i];

            AssertMatch(expectedError, (ModelValidationSingleError)actualError);
        }
    }

    [Theory]
    [MemberData(nameof(JsonValidationTestCases))]
    public void VerifyJsonValidation(string name, string json, List<ModelValidationSingleError>? expectedErrors) {
        var model = JsonSerializer.Deserialize<AuthorizationModel>(json)!;

        if (expectedErrors == null || expectedErrors.Count == 0) {
            // If no errors expected, validation should succeed
            ModelValidator.ValidateJson(model);
            return;
        }

        // If errors are expected, validation should throw DslErrorsException
        var exception = Assert.Throws<DslErrorsException>(() => ModelValidator.ValidateJson(model));

        var errorsCount = expectedErrors.Count;
        var formattedErrors = expectedErrors
            .Select(error => $"validation error: {error.Message}")
            .ToList();
        var expectedMessage = $"{errorsCount} error{(errorsCount > 1 ? "s" : "")} occurred:\n\t* {string.Join("\n\t* ", formattedErrors)}\n\n";

        Assert.Equal(expectedMessage, exception.Message);

        var actualErrors = exception.Errors;
        for (int i = 0; i < expectedErrors.Count; i++) {
            var expectedError = expectedErrors[i];
            var actualError = actualErrors[i];

            AssertMatch(expectedError, (ModelValidationSingleError)actualError);
        }
    }

    private void AssertMatch(ParsingError expectedError, ParsingError actualError) {
        Assert.Equal(expectedError.Message, actualError.Message);
        Assert.Equal(expectedError.Line, actualError.Line);
        Assert.Equal(expectedError.Column, actualError.Column);
    }

    private void AssertMatch(ModelValidationSingleError expectedError, ModelValidationSingleError actualError) {
        Assert.Equal(expectedError.Message, actualError.Message);

        if (expectedError.Line != null) {
            Assert.Equal(expectedError.Line, actualError.Line);
        }

        if (expectedError.Column != null) {
            Assert.Equal(expectedError.Column, actualError.Column);
        }

        Assert.Equal(expectedError.Metadata.ErrorType, actualError.Metadata.ErrorType);
        if (expectedError.Metadata.TypeName != null) {
            Assert.Equal(expectedError.Metadata.TypeName, actualError.Metadata.TypeName);
        }
        if (expectedError.Metadata.Relation != null) {
            Assert.Equal(expectedError.Metadata.Relation, actualError.Metadata.Relation);
        }
        if (expectedError.Metadata.Condition != null) {
            Assert.Equal(expectedError.Metadata.Condition, actualError.Metadata.Condition);
        }
    }

    public static IEnumerable<object[]> DslSyntaxTestCases() {
        return TestsData.DslSyntaxTestCases
            .Select(testCase => new object[]
            {
                testCase.Name,
                testCase.Dsl,
                testCase.ExpectedErrors,
                testCase.Skip
            });
    }

    public static IEnumerable<object[]> DslValidationTestCases() {
        return TestsData.DslValidationTestCases
            .Select(testCase => new object[]
            {
                testCase.Name,
                testCase.Dsl,
                testCase.ExpectedErrors,
                testCase.Skip
            });
    }

    public static IEnumerable<object[]> JsonValidationTestCases() {
        return TestsData.JsonValidationTestCases
            .Select(testCase => new object[]
            {
                testCase.Name,
                testCase.Json,
                testCase.ExpectedErrors!
            });
    }
}