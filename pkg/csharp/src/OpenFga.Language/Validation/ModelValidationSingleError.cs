using System.Text.Json.Serialization;
using OpenFga.Language.Errors;

namespace OpenFga.Language.Validation;

public class ModelValidationSingleError : ParsingError
{
    [JsonPropertyName("metadata")]
    public ValidationMetadata Metadata { get; set; } = new();

    public ModelValidationSingleError() { }

    public ModelValidationSingleError(ErrorProperties properties, ValidationMetadata metadata) 
        : base("validation", properties)
    {
        Metadata = metadata;
    }
}
