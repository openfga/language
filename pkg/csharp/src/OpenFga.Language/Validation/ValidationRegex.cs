using System.Text.RegularExpressions;

namespace OpenFga.Language.Validation;

public class ValidationRegex {
    private readonly string _rule;
    private readonly Regex _regex;

    private ValidationRegex(string rule, Regex regex) {
        _rule = rule;
        _regex = regex;
    }

    public string Rule => _rule;
    public Regex Regex => _regex;

    public static ValidationRegex Build(string name, string rule) {
        Regex regex;
        try {
            regex = new Regex(rule);
        }
        catch (ArgumentException ex) {
            var message = $"Incorrect {name} regex specification for {rule}";
            throw new ArgumentException(message, ex);
        }
        return new ValidationRegex(rule, regex);
    }

    public bool Matches(string input) {
        return _regex.IsMatch(input);
    }
}