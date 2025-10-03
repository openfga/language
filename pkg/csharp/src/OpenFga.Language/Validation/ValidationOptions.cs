namespace OpenFga.Language.Validation;

public class ValidationOptions
{
    private string _typePattern = $"^{Validator.Rules.Type}$";
    private string _relationPattern = $"^{Validator.Rules.Relation}$";

    public string TypePattern
    {
        get => _typePattern;
        set => _typePattern = value;
    }

    public string RelationPattern
    {
        get => _relationPattern;
        set => _relationPattern = value;
    }
}
