using System.Text.Json;
using System.Text.Json.Serialization;

namespace OpenFga.Language;

internal static class Json
{
    private static readonly JsonSerializerOptions Options = new()
    {
        PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
        WriteIndented = false,
        DefaultIgnoreCondition = JsonIgnoreCondition.Never,
        PropertyNameCaseInsensitive = true
    };

    public static T Parse<T>(string json) where T : class
    {
        if (string.IsNullOrEmpty(json) || json == "null")
        {
            return Activator.CreateInstance<T>();
        }
        return JsonSerializer.Deserialize<T>(json, Options) ?? throw new InvalidOperationException("Failed to deserialize JSON");
    }

    public static string Stringify(object obj)
    {
        return JsonSerializer.Serialize(obj, Options);
    }
}
