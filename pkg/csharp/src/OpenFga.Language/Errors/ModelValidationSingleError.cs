using SharpYaml.Serialization;
using System.Text.Json.Serialization;

namespace OpenFga.Language.Errors;

/// <summary>
/// Represents a single validation error that occurred during authorization model validation.
/// Contains detailed metadata about the validation error including symbol information and error context.
/// </summary>
public class ModelValidationSingleError : ParsingError {
    /// <summary>
    /// Gets or sets the validation metadata containing detailed error information.
    /// </summary>
    [JsonPropertyName("metadata")]
    [YamlMember("metadata")]
    public ValidationMetadata Metadata { get; set; } = new();

    /// <summary>
    /// Initializes a new instance of the ModelValidationSingleError class.
    /// </summary>
    public ModelValidationSingleError() { }

    /// <summary>
    /// Initializes a new instance of the ModelValidationSingleError class with error properties and metadata.
    /// </summary>
    /// <param name="properties">The error properties including location and message</param>
    /// <param name="metadata">The validation metadata containing detailed error information</param>
    public ModelValidationSingleError(ErrorProperties properties, ValidationMetadata metadata)
        : base("validation", properties) {
        Metadata = metadata;
    }
}