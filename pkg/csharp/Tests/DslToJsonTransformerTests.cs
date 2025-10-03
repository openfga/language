using Xunit;
using OpenFgaLanguage.Errors;
using OpenFgaLanguage.Transformers;

namespace OpenFgaLanguage.Tests;

public class DslToJsonTransformerTests
{
    [Fact]
    public void Transform_ValidDsl_ReturnsJsonString()
    {
        // Arrange
        var dsl = """
            model
              schema 1.1
            type user
            """;

        var transformer = new DslToJsonTransformer();

        // Act
        var result = transformer.Transform(dsl);

        // Assert
        Assert.NotNull(result);
        Assert.Contains("\"schema_version\"", result);
        Assert.Contains("\"type_definitions\"", result);
    }

    [Fact]
    public void ParseDsl_ValidDsl_ReturnsSuccessResult()
    {
        // Arrange
        var dsl = """
            model
              schema 1.1
            type user
            """;

        var transformer = new DslToJsonTransformer();

        // Act
        var result = transformer.ParseDsl(dsl);

        // Assert
        Assert.True(result.IsSuccess());
        Assert.False(result.IsFailure());
        Assert.Empty(result.Errors);
        Assert.NotNull(result.AuthorizationModel);
    }

    [Fact]
    public void ParseDsl_InvalidDsl_ReturnsFailureResult()
    {
        // Arrange
        var dsl = """
            invalid syntax here
            """;

        var transformer = new DslToJsonTransformer();

        // Act
        var result = transformer.ParseDsl(dsl);

        // Assert
        Assert.False(result.IsSuccess());
        Assert.True(result.IsFailure());
        Assert.NotEmpty(result.Errors);
    }

    [Fact]
    public void Transform_InvalidDsl_ThrowsDslErrorsException()
    {
        // Arrange
        var dsl = """
            invalid syntax here
            """;

        var transformer = new DslToJsonTransformer();

        // Act & Assert
        var exception = Assert.Throws<DslErrorsException>(() => transformer.Transform(dsl));
        Assert.NotEmpty(exception.Errors);
    }

    [Fact]
    public void ParseDsl_WithComments_IgnoresComments()
    {
        // Arrange
        var dsl = """
            model
              schema 1.1
            type user # This is a comment
            """;

        var transformer = new DslToJsonTransformer();

        // Act
        var result = transformer.ParseDsl(dsl);

        // Assert
        Assert.True(result.IsSuccess());
        Assert.Empty(result.Errors);
    }
}
