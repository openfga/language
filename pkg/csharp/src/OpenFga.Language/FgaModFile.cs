using System.Text.Json.Serialization;

namespace OpenFga.Language.ModFile;

public class FgaModFile
{
    public const string JsonPropertySchema = "schema";
    public const string JsonPropertyContents = "contents";

    [JsonPropertyName(JsonPropertySchema)]
    public ModFileStringProperty? Schema { get; set; }

    [JsonPropertyName(JsonPropertyContents)]
    public ModFileArrayProperty? Contents { get; set; }

    public FgaModFile() { }

    public FgaModFile SetSchema(ModFileStringProperty? schema)
    {
        Schema = schema;
        return this;
    }

    public FgaModFile SetContents(ModFileArrayProperty? contents)
    {
        Contents = contents;
        return this;
    }
}
