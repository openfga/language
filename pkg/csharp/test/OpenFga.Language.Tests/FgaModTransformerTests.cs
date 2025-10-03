using Xunit;
using OpenFga.Language.Errors;
using OpenFga.Language.Transformers;

namespace OpenFga.Language.Tests;

public class FgaModTransformerTests
{
    [Fact]
    public void Parse_ValidModFile_ReturnsCorrectStructure()
    {
        // Arrange
        var modFileContent = """
            schema: "1.2"
            contents:
              - "user.fga"
              - "document.fga"
            """;

        // Act
        var result = FgaModTransformer.Parse(modFileContent);

        // Assert
        Assert.NotNull(result);
        Assert.NotNull(result.Schema);
        Assert.Equal("1.2", result.Schema.Value);
        Assert.NotNull(result.Contents);
        Assert.Equal(2, result.Contents.Value.Count);
        Assert.Equal("user.fga", result.Contents.Value[0].Value);
        Assert.Equal("document.fga", result.Contents.Value[1].Value);
    }

    [Fact]
    public void Parse_MissingSchema_ThrowsValidationError()
    {
        // Arrange
        var modFileContent = """
            contents:
              - "user.fga"
            """;

        // Act & Assert
        var exception = Assert.Throws<ModFileValidationError>(() => FgaModTransformer.Parse(modFileContent));
        Assert.Single(exception.Errors);
        Assert.Contains("missing schema field", exception.Message);
    }

    [Fact]
    public void Parse_MissingContents_ThrowsValidationError()
    {
        // Arrange
        var modFileContent = """
            schema: "1.2"
            """;

        // Act & Assert
        var exception = Assert.Throws<ModFileValidationError>(() => FgaModTransformer.Parse(modFileContent));
        Assert.Single(exception.Errors);
        Assert.Contains("missing contents field", exception.Message);
    }

    [Fact]
    public void Parse_UnsupportedSchemaVersion_ThrowsValidationError()
    {
        // Arrange
        var modFileContent = """
            schema: "1.0"
            contents:
              - "user.fga"
            """;

        // Act & Assert
        var exception = Assert.Throws<ModFileValidationError>(() => FgaModTransformer.Parse(modFileContent));
        Assert.Single(exception.Errors);
        Assert.Contains("unsupported schema version", exception.Message);
    }

    [Fact]
    public void Parse_InvalidFileExtension_ThrowsValidationError()
    {
        // Arrange
        var modFileContent = """
            schema: "1.2"
            contents:
              - "user.txt"
            """;

        // Act & Assert
        var exception = Assert.Throws<ModFileValidationError>(() => FgaModTransformer.Parse(modFileContent));
        Assert.Single(exception.Errors);
        Assert.Contains("contents items should use fga file extension", exception.Message);
    }

    [Fact]
    public void Transform_ValidModFile_ReturnsJsonString()
    {
        // Arrange
        var modFileContent = """
            schema: "1.2"
            contents:
              - "user.fga"
            """;

        // Act
        var result = FgaModTransformer.Transform(modFileContent);

        // Assert
        Assert.NotNull(result);
        Assert.Contains("\"schema\"", result);
        Assert.Contains("\"contents\"", result);
        Assert.Contains("\"user.fga\"", result);
    }
}
