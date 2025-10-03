using System.Text;
using System.Web;
using YamlDotNet.Serialization;
using OpenFgaLanguage.Errors;
using OpenFgaLanguage.ModFile;
using OpenFgaLanguage.Utils;

namespace OpenFgaLanguage.Transformers;

public class FgaModTransformer
{
    private readonly List<ModFileValidationSingleError> _errors = new();
    private readonly string _modFileContents;

    public FgaModTransformer(string modFileContents)
    {
        _modFileContents = modFileContents;
    }

    public static string Transform(string modFileContents)
    {
        return Json.Stringify(Parse(modFileContents));
    }

    public static FgaModFile Parse(string modFileContents)
    {
        return new FgaModTransformer(modFileContents).Parse();
    }

    public FgaModFile Parse()
    {
        var deserializer = new DeserializerBuilder().Build();
        
        try
        {
            var yamlData = deserializer.Deserialize<Dictionary<object, object>>(_modFileContents);
            var modFile = new FgaModFile();
            var seenFields = new List<string>();

            foreach (var kvp in yamlData)
            {
                var key = kvp.Key.ToString();
                if (key == "schema")
                {
                    seenFields.Add("schema");
                    HandleSchema(kvp.Value, modFile);
                }
                else if (key == "contents")
                {
                    seenFields.Add("contents");
                    HandleContents(kvp.Value, modFile);
                }
            }

            if (!seenFields.Contains("schema"))
            {
                AddError("missing schema field", new StartEnd(0, 0), new StartEnd(0, 0));
            }

            if (!seenFields.Contains("contents"))
            {
                AddError("missing contents field", new StartEnd(0, 0), new StartEnd(0, 0));
            }

            if (_errors.Count > 0)
            {
                throw new ModFileValidationError(_errors);
            }

            return modFile;
        }
        catch (ModFileValidationError)
        {
            // Re-throw validation errors as-is
            throw;
        }
        catch (Exception ex)
        {
            throw new InvalidOperationException($"Failed to parse YAML: {ex.Message}", ex);
        }
    }

    private void HandleSchema(object? value, FgaModFile modFile)
    {
        if (value is string schemaValue)
        {
            if (!schemaValue.Equals("1.2"))
            {
                AddError("unsupported schema version, fga.mod only supported in version `1.2`", new StartEnd(0, 0), new StartEnd(0, 0));
            }
            else
            {
                var property = new ModFileStringProperty()
                    .SetValue(schemaValue)
                    .SetLine(new StartEnd(0, 0))
                    .SetColumn(new StartEnd(0, 0));
                modFile.SetSchema(property);
            }
        }
        else
        {
            AddError("unexpected schema type, expected string", new StartEnd(0, 0), new StartEnd(0, 0));
        }
    }

    private void HandleContents(object? value, FgaModFile modFile)
    {
        if (value is IEnumerable<object> contentsList)
        {
            var contents = new List<ModFileStringProperty>();
            
            foreach (var item in contentsList)
            {
                if (item is string rawValue)
                {
                    var decodedValue = HttpUtility.UrlDecode(rawValue, Encoding.UTF8);
                    decodedValue = decodedValue.Replace("\\", "/");

                    if (decodedValue.Contains("../") || decodedValue.StartsWith("/"))
                    {
                        AddError($"invalid contents item {rawValue}", new StartEnd(0, 0), new StartEnd(0, 0));
                        continue;
                    }

                    if (!decodedValue.EndsWith(".fga"))
                    {
                        AddError($"contents items should use fga file extension, got {decodedValue}", new StartEnd(0, 0), new StartEnd(0, 0));
                        continue;
                    }

                    contents.Add(new ModFileStringProperty()
                        .SetValue(decodedValue)
                        .SetLine(new StartEnd(0, 0))
                        .SetColumn(new StartEnd(0, 0)));
                }
                else
                {
                    AddError("unexpected contents item type, expected string", new StartEnd(0, 0), new StartEnd(0, 0));
                }
            }

            var property = new ModFileArrayProperty()
                .SetValue(contents)
                .SetLine(new StartEnd(0, 0))
                .SetColumn(new StartEnd(0, 0));

            modFile.SetContents(property);
        }
        else
        {
            AddError("unexpected contents type, expected list of strings", new StartEnd(0, 0), new StartEnd(0, 0));
        }
    }

    private void AddError(string message, StartEnd line, StartEnd column)
    {
        _errors.Add(new ModFileValidationSingleError(new ErrorProperties(line, column, message)));
    }
}