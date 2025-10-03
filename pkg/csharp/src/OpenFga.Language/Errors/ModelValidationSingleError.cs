using SharpYaml.Serialization;
using System.Text.Json.Serialization;

namespace OpenFga.Language.Errors;

public class ModelValidationSingleError : ParsingError {
    [JsonPropertyName("metadata")]
    [YamlMember("metadata")]
    public ValidationMetadata Metadata { get; set; } = new();

    public ModelValidationSingleError() { }

    public ModelValidationSingleError(ErrorProperties properties, ValidationMetadata metadata)
        : base("validation", properties) {
        Metadata = metadata;
    }
}