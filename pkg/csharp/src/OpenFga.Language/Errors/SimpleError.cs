using System.Text.Json.Serialization;
using YamlDotNet.Serialization;

namespace OpenFga.Language.Errors;

public abstract class SimpleError
{
    [JsonPropertyName("msg")]
    [YamlMember(Alias = "msg")]
    [SharpYaml.Serialization.YamlMember("msg")]
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
