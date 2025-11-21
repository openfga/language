namespace OpenFga.Language.Validation;

internal class AllowableTypesResult {
    public bool Valid { get; }
    public List<string> AllowableTypes { get; }

    public AllowableTypesResult(bool valid, List<string> allowableTypes) {
        Valid = valid;
        AllowableTypes = allowableTypes;
    }
}