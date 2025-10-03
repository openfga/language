namespace OpenFga.Language.Errors;

public class Metadata
{
    public string Symbol { get; set; } = string.Empty;

    public Metadata() { }

    public Metadata(string symbol)
    {
        Symbol = symbol;
    }
}
