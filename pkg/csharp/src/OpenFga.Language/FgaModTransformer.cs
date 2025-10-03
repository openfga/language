using System.Text;
using System.Text.Encodings.Web;
using System.Text.RegularExpressions;
using System.Web;
using YamlDotNet.Serialization;
using OpenFga.Language.Errors;
using OpenFga.Language.ModFile;
using YamlDotNet.Core;
using YamlDotNet.Core.Events;
using YamlDotNet.Core.Tokens;
using Scalar = YamlDotNet.Core.Events.Scalar;

namespace OpenFga.Language.Transformers;

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
        var s = new Scanner(new StringReader(_modFileContents));
        List<Token> l = [];
        while (s.MoveNext())
        {
            l.Add(s.Current);
        }

        var parser = new Parser(new Scanner(new StringReader(_modFileContents)));
        var modFile = new FgaModFile();
        var seenFields = new HashSet<string>();

        while (parser.MoveNext())
        {
            var currentToken = parser.Current;
            if (currentToken is not Scalar key)
            {
                continue;
            }

            var name = key.Value;

            // move to value
            parser.MoveNext();
            currentToken = parser.Current;
            var location = currentToken.End;

            if (name == "schema")
            {
                seenFields.Add("schema");
                HandleSchema(parser, modFile, currentToken, location);
            }
            else if (name == "contents")
            {
                seenFields.Add("contents");
                HandleContents(parser, modFile, currentToken, location);
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

    private void HandleSchema(Parser parser, FgaModFile modFile, ParsingEvent currentToken, Mark location)
    {
        if (currentToken is not Scalar scalar)
        {
            // Without a scalar there is no value so we can't match the expected error
            return;
        }

        var value = scalar.Value;

        if (!Regex.IsMatch(scalar.Value, @"^\d+\.\d+$"))
        {
            var line = GetLine(location);
            var column = GetColumn(location, value);
            AddError("unexpected schema type, expected string got value " + value, line, column);
        }
        else if (scalar.Value != "1.2")
        {
            var line = GetLine(location);
            var column = GetColumn(location, "\"" + value + "\"");
            AddError("unsupported schema version, fga.mod only supported in version `1.2`", line, column);
        }
        else
        {
            var line = GetLine(location);
            var column = GetColumn(location, "\"" + value + "\"");
            var property = new ModFileStringProperty()
            {
                Value = value,
                Line = line,
                Column = column
            };
            modFile.Schema = property;
        }
    }

    private void HandleContents(Parser parser, FgaModFile modFile, ParsingEvent currentToken, Mark startLocation)
    {
        if (currentToken is not SequenceStart)
        {
            if (currentToken is not Scalar scalar)
            {
                // Without a scalar there is no value so we can't match the expected error
                return;
            }

            var value = scalar.Value;
            StartEnd line = GetLine(startLocation);
            StartEnd column = GetColumn(startLocation, value);
            AddError("unexpected contents type, expected list of strings got value " + value, line, column);
            return;
        }

        List<ModFileStringProperty> contents = new List<ModFileStringProperty>();
        while (parser.MoveNext())
        {
            currentToken = parser.Current;
            if (currentToken is SequenceEnd)
            {
                break;
            }

            if (currentToken is not Scalar scalar)
            {
                // Without a scalar there is no value so we can't match the expected error
                continue;
            }

            var currentLoc = currentToken.End;
            var rawValue = scalar.Value;

            StartEnd line = GetLine(currentLoc);
            StartEnd column = GetColumn(currentLoc, rawValue);

            var value = rawValue.Replace(@"\\", @"\"); // The double backslash is not handled by YamlDotNet

            // YamlDotNet parses true and 1 as Scalar. The tests want this to be boolean and number which YamlDotNet does not support.
            // We could also check for ends with .fga here as that is the real determinant of a valid contents item.
            if (!rawValue.Contains("."))
            {
                AddError("unexpected contents item type, expected string got value " + rawValue, line, column);
                continue;
            }

            value = HttpUtility.UrlDecode(scalar.Value, Encoding.UTF8);

            value = value.Replace(@"\", "/");

            if (value.Contains("../") || value.StartsWith("/"))
            {
                AddError("invalid contents item " + rawValue, line, column);
                continue;
            }

            if (!value.EndsWith(".fga"))
            {
                AddError("contents items should use fga file extension, got " + value, line, column);
                continue;
            }

            contents.Add(new ModFileStringProperty()
            {
                Value = value,
                Line = line,
                Column = column
            });
        }
        var endLoc = parser.Current.Start;

        modFile.Contents = new ModFileArrayProperty
        {
            Line = new StartEnd(startLocation.Line - 1, endLoc.Line - 1),
            Column = new StartEnd(startLocation.Column - 1, endLoc.Column - 1),
            Value = contents
        };
    }

    private StartEnd GetLine(Mark loc)
    {
        var line = loc.Line - 1;
        return new StartEnd(line, line);
    }

    private StartEnd GetColumn(Mark loc, string text)
    {
        // As the yaml parser does not expose an easy way of checking if the value is wrapped in
        // quotes, scan the modFile string to see if it is and add an offset to the string length.
        var quotesOffset = 0;
        if (_modFileContents.Contains("'" + text + "'") || _modFileContents.Contains("\"" + text + "\""))
        {
            quotesOffset = 2;
        }

        int columnEnd = loc.Column - 1;
        int columnStart = columnEnd - (text.Length + quotesOffset);
        return new StartEnd(columnStart, columnEnd);
    }
    private void AddError(string message, StartEnd line, StartEnd column)
    {
        _errors.Add(new ModFileValidationSingleError(new ErrorProperties(line, column, message)));
    }
}