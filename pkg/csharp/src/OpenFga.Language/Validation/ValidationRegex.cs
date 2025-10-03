using System.Text.RegularExpressions;

namespace OpenFga.Language.Validation;

/// <summary>
/// Represents a compiled regular expression with associated validation rules.
/// Provides a wrapper around .NET Regex with additional metadata for validation purposes.
/// </summary>
public class ValidationRegex {
    private readonly string _rule;
    private readonly Regex _regex;

    private ValidationRegex(string rule, Regex regex) {
        _rule = rule;
        _regex = regex;
    }

    /// <summary>
    /// Gets the original regex rule string.
    /// </summary>
    public string Rule => _rule;

    /// <summary>
    /// Gets the compiled Regex instance.
    /// </summary>
    public Regex Regex => _regex;

    /// <summary>
    /// Creates a new ValidationRegex instance from a rule string.
    /// </summary>
    /// <param name="name">The name of the regex pattern for error reporting</param>
    /// <param name="rule">The regex pattern string</param>
    /// <returns>A new ValidationRegex instance</returns>
    /// <exception cref="ArgumentException">Thrown when the regex pattern is invalid</exception>
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

    /// <summary>
    /// Tests whether the input string matches the regex pattern.
    /// </summary>
    /// <param name="input">The string to test</param>
    /// <returns>True if the input matches the pattern, false otherwise</returns>
    public bool Matches(string input) {
        return _regex.IsMatch(input);
    }
}