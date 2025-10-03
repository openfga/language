using System.Text.Json;
using SharpYaml.Serialization;

namespace OpenFga.Language.Tests.util;

public static class YamlHelper
{
    public static List<T> ParseList<T>(string yamlContent)
    {
        var serializer = new SharpYaml.Serialization.Serializer();
        var result = serializer.Deserialize<List<T>>(yamlContent);
        return result ?? new List<T>();
    }
}
