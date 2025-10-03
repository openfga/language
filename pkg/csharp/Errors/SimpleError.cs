using System.Text.Json.Serialization;

namespace OpenFgaLanguage.Errors;

public abstract class SimpleError
{
    [JsonPropertyName("msg")]
    public string Message { get; set; } = string.Empty;

    protected SimpleError() { }

    protected SimpleError(string message)
    {
        Message = message;
    }

    public override string ToString()
    {
        return Message;
    }
}
