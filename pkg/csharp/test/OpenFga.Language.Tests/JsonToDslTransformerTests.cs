using Xunit;
using OpenFga.Language.Errors;
using OpenFga.Language.Transformers;

namespace OpenFga.Language.Tests;

public class JsonToDslTransformerTests
{
    [Fact]
    public void Transform_ValidJson_ReturnsDslString()
    {
        // Arrange
        var json = """
            {
              "schema_version": "1.1",
              "type_definitions": [
                {
                  "type": "user"
                }
              ]
            }
            """;

        var transformer = new JsonToDslTransformer();

        // Act
        var result = transformer.Transform(json);

        // Assert
        Assert.NotNull(result);
        Assert.Contains("model", result);
        Assert.Contains("schema 1.1", result);
        Assert.Contains("type user", result);
    }

    [Fact]
    public void Transform_JsonWithRelations_ReturnsDslWithRelations()
    {
        // Arrange
        var json = """
            {
              "schema_version": "1.1",
              "type_definitions": [
                {
                  "type": "document",
                  "relations": {
                    "reader": {
                      "this": {}
                    }
                  }
                }
              ]
            }
            """;

        var transformer = new JsonToDslTransformer();

        // Act
        var result = transformer.Transform(json);

        // Assert
        Assert.NotNull(result);
        Assert.Contains("type document", result);
        Assert.Contains("relations", result);
        Assert.Contains("define reader:", result);
    }

    [Fact]
    public void Transform_JsonWithConditions_ReturnsDslWithConditions()
    {
        // Arrange
        var json = """
            {
              "schema_version": "1.1",
              "type_definitions": [
                {
                  "type": "user"
                }
              ],
              "conditions": {
                "is_public": {
                  "name": "is_public",
                  "expression": "user == \"public\"",
                  "parameters": {}
                }
              }
            }
            """;

        var transformer = new JsonToDslTransformer();

        // Act
        var result = transformer.Transform(json);

        // Assert
        Assert.NotNull(result);
        Assert.Contains("condition is_public()", result);
        Assert.Contains("user == \"public\"", result);
    }

    [Fact]
    public void Transform_EmptyJson_ReturnsBasicDsl()
    {
        // Arrange
        var json = "{}";

        var transformer = new JsonToDslTransformer();

        // Act
        var result = transformer.Transform(json);

        // Assert
        Assert.NotNull(result);
        Assert.Contains("model", result);
        Assert.Contains("schema 1.1", result);
    }

    [Fact]
    public void Transform_JsonWithNullModel_ReturnsBasicDsl()
    {
        // Arrange
        var json = "null";

        var transformer = new JsonToDslTransformer();

        // Act
        var result = transformer.Transform(json);

        // Assert
        Assert.NotNull(result);
        Assert.Contains("model", result);
        Assert.Contains("schema 1.1", result);
    }
}
