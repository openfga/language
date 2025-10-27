using Antlr4.Runtime;
using Antlr4.Runtime.Tree;
using OpenFga.Language.Errors;
using OpenFga.Sdk.Model;
using System.Text.RegularExpressions;

namespace OpenFga.Language;

/// <summary>
/// Transforms OpenFGA DSL (Domain Specific Language) syntax into JSON format.
/// This class provides functionality to parse and convert FGA authorization model definitions
/// from their textual DSL representation to structured JSON.
/// </summary>
public class DslToJsonTransformer {
    private static readonly Regex SpacesLinePattern = new(@"^\s*$");
    private static readonly Regex CommentedLinePattern = new(@"^\s*#.*$");

    /// <summary>
    /// Transforms the provided DSL string into JSON format.
    /// </summary>
    /// <param name="dsl">The DSL string to transform</param>
    /// <returns>A JSON string representation of the authorization model</returns>
    /// <exception cref="DslErrorsException">Thrown when the DSL contains syntax errors</exception>
    public string Transform(string dsl) {
        return Json.Stringify(ParseAuthorisationModel(dsl));
    }

    private AuthorizationModel ParseAuthorisationModel(string dsl) {
        var result = ParseDsl(dsl);
        if (result.IsFailure()) {
            throw new DslErrorsException(result.Errors);
        }

        return result.AuthorizationModel;
    }

    private string CleanLine(string line) {
        if (SpacesLinePattern.IsMatch(line) || CommentedLinePattern.IsMatch(line)) {
            return string.Empty;
        }

        var cleanedLine = line.Split(" #")[0];
        return cleanedLine.TrimEnd();
    }

    /// <summary>
    /// Parses a DSL string and returns a result containing the authorization model and any parsing errors.
    /// </summary>
    /// <param name="dsl">The DSL string to parse</param>
    /// <returns>A <see cref="Result"/> containing the parsed authorization model and any errors</returns>
    public Result ParseDsl(string dsl) {
        var cleanedDsl = string.Join("\n", dsl.Split('\n').Select(CleanLine));

        var inputStream = new AntlrInputStream(cleanedDsl);
        var errorListener = new OpenFgaDslErrorListener();

        var lexer = new OpenFGALexer(inputStream);
        lexer.RemoveErrorListeners();
        lexer.AddErrorListener(errorListener);
        var tokenStream = new CommonTokenStream(lexer);

        var parser = new OpenFGAParser(tokenStream);
        parser.RemoveErrorListeners();
        parser.AddErrorListener(errorListener);

        var listener = new OpenFgaDslListener(parser);
        ParseTreeWalker.Default.Walk(listener, parser.main());

        return new Result(listener.GetAuthorizationModel(), errorListener.Errors.ToList());
    }

    /// <summary>
    /// Represents the result of parsing a DSL string, containing either the parsed authorization model
    /// or a list of syntax errors encountered during parsing.
    /// </summary>
    /// <param name="authorizationModel">The parsed authorization model</param>
    /// <param name="errors">List of syntax errors encountered during parsing</param>
    public class Result(AuthorizationModel authorizationModel, List<SyntaxError> errors) {

        /// <summary>
        /// Gets the parsed authorization model.
        /// </summary>
        public AuthorizationModel AuthorizationModel { get; } = authorizationModel;

        /// <summary>
        /// Gets the list of syntax errors encountered during parsing.
        /// </summary>
        public List<SyntaxError> Errors { get; } = errors;

        /// <summary>
        /// Determines whether the parsing was successful (no errors).
        /// </summary>
        /// <returns>True if parsing was successful, false otherwise</returns>
        public bool IsSuccess() {
            return Errors.Count == 0;
        }

        /// <summary>
        /// Determines whether the parsing failed (contains errors).
        /// </summary>
        /// <returns>True if parsing failed, false otherwise</returns>
        public bool IsFailure() {
            return !IsSuccess();
        }
    }
}