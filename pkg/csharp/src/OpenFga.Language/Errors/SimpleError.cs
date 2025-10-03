using SharpYaml.Serialization;
using System.Text.Json.Serialization;

namespace OpenFga.Language.Errors;

public abstract class SimpleError {
    [JsonPropertyName("msg")]
    [YamlMember("msg")]
    public string Message { get; set; } = string.Empty;

    protected SimpleError() { }

    protected SimpleError(string message) {
        Message = message;
    }

    public override string ToString() {
        return Message;
    }
}