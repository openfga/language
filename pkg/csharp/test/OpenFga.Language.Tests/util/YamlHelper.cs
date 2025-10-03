using System.Text.Json;
using YamlDotNet.Serialization;

namespace OpenFga.Language.Tests.util;

public static class YamlHelper
{
    public static List<T> ParseList<T>(string yamlContent)
    {
        var deserializer = new DeserializerBuilder().Build();
        var result = deserializer.Deserialize<List<T>>(yamlContent);
        return result ?? new List<T>();
    }
}
