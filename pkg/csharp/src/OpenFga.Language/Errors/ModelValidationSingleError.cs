using System.Text.Json.Serialization;
using YamlDotNet.Serialization;

namespace OpenFga.Language.Errors;

public class ModelValidationSingleError : ParsingError
{
    [JsonPropertyName("metadata")]
    [YamlMember(Alias = "metadata")]
    public ValidationMetadata Metadata { get; set; } = new();

    public ModelValidationSingleError() { }

    public ModelValidationSingleError(ErrorProperties properties, ValidationMetadata metadata) 
        : base("validation", properties)
    {
        Metadata = metadata;
    }
}
