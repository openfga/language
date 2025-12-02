namespace OpenFga.Language.Errors;

/// <summary>
/// Contains basic metadata information for parsing errors.
/// Used to store symbol information associated with parsing errors.
/// </summary>
public class Metadata {
    /// <summary>
    /// Gets or sets the symbol that caused the parsing error.
    /// </summary>
    public string Symbol { get; set; } = string.Empty;

    /// <summary>
    /// Initializes a new instance of the Metadata class.
    /// </summary>
    public Metadata() { }

    /// <summary>
    /// Initializes a new instance of the Metadata class with a symbol.
    /// </summary>
    /// <param name="symbol">The symbol that caused the parsing error</param>
    public Metadata(string symbol) {
        Symbol = symbol;
    }
}