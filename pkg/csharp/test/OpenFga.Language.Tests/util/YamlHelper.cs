using System.Text.Json;
using SharpYaml.Serialization;
using YamlDotNet.Serialization;

namespace OpenFga.Language.Tests.util;

public static class YamlHelper
{
    public static List<T> ParseList<T>(string yamlContent)
    {
        //var deserializer = new DeserializerBuilder().Build();
        //var result = deserializer.Deserialize<List<T>>(yamlContent);

        var result = new SharpYaml.Serialization.Serializer(new SerializerSettings()
        {
            NamingConvention = new FlatNamingConvention()
        }).Deserialize<List<T>>(yamlContent);
        return result ?? new List<T>();
    }
}
