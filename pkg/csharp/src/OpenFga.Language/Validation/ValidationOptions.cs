namespace OpenFga.Language.Validation;

/// <summary>
/// Configuration options for validating OpenFGA authorization models.
/// Allows customization of validation patterns for type names and relation names.
/// </summary>
public class ValidationOptions {
    private string _typePattern = $"^{Validator.Rules.Type}$";
    private string _relationPattern = $"^{Validator.Rules.Relation}$";

    /// <summary>
    /// Gets or sets the regular expression pattern used to validate type names.
    /// Defaults to the standard FGA type name pattern.
    /// </summary>
    public string TypePattern {
        get => _typePattern;
        set => _typePattern = value;
    }

    /// <summary>
    /// Gets or sets the regular expression pattern used to validate relation names.
    /// Defaults to the standard FGA relation name pattern.
    /// </summary>
    public string RelationPattern {
        get => _relationPattern;
        set => _relationPattern = value;
    }
}