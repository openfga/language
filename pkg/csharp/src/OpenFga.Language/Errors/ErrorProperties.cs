namespace OpenFga.Language.Errors;

public class ErrorProperties {
    public StartEnd? Line { get; set; }
    public StartEnd? Column { get; set; }
    public string Message { get; set; } = string.Empty;

    public ErrorProperties() { }

    public ErrorProperties(StartEnd? line, StartEnd? column, string message) {
        Line = line;
        Column = column;
        Message = message;
    }

    internal string GetFullMessage(string type) {
        if (Line != null && Column != null) {
            return $"{type} error at line={Line.Start}, column={Column.Start}: {Message}";
        }
        else {
            return $"{type} error: {Message}";
        }
    }
}