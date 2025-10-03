using SharpYaml.Serialization;
using System.Text.Json.Serialization;

namespace OpenFga.Language.Errors;

/// <summary>
/// Abstract base class for simple errors that contain only a message.
/// Provides the foundation for all error types in the OpenFGA language library.
/// </summary>
public abstract class SimpleError {
    /// <summary>
    /// Gets or sets the error message.
    /// </summary>
    [JsonPropertyName("msg")]
    [YamlMember("msg")]
    public string Message { get; set; } = string.Empty;

    /// <summary>
    /// Initializes a new instance of the SimpleError class.
    /// </summary>
    protected SimpleError() { }

    /// <summary>
    /// Initializes a new instance of the SimpleError class with a message.
    /// </summary>
    /// <param name="message">The error message</param>
    protected SimpleError(string message) {
        Message = message;
    }

    /// <summary>
    /// Returns a string representation of the error.
    /// </summary>
    /// <returns>The error message</returns>
    public override string ToString() {
        return Message;
    }
}