using System.Text.RegularExpressions;
using Antlr4.Runtime;
using Antlr4.Runtime.Tree;
using OpenFga.Sdk.Model;
using OpenFga.Language;
using OpenFgaLanguage.Errors;
using OpenFgaLanguage.Listeners;
using OpenFgaLanguage.Utils;

namespace OpenFgaLanguage.Transformers;

public class DslToJsonTransformer
{
    private static readonly Regex SpacesLinePattern = new(@"^\s*$");
    private static readonly Regex CommentedLinePattern = new(@"^\s*#.*$");

    public string Transform(string dsl)
    {
        return Json.Stringify(ParseAuthorisationModel(dsl));
    }

    private AuthorizationModel ParseAuthorisationModel(string dsl)
    {
        var result = ParseDsl(dsl);
        if (result.IsFailure())
        {
            throw new DslErrorsException(result.Errors);
        }

        return result.AuthorizationModel;
    }

    private string CleanLine(string line)
    {
        if (SpacesLinePattern.IsMatch(line) || CommentedLinePattern.IsMatch(line))
        {
            return string.Empty;
        }
        
        var cleanedLine = line.Split(" #")[0];
        return cleanedLine.TrimEnd();
    }

    public Result ParseDsl(string dsl)
    {
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

    public class Result
    {
        public AuthorizationModel AuthorizationModel { get; }
        public List<SyntaxError> Errors { get; }

        public Result(AuthorizationModel authorizationModel, List<SyntaxError> errors)
        {
            AuthorizationModel = authorizationModel;
            Errors = errors;
        }

        public bool IsSuccess()
        {
            return Errors.Count == 0;
        }

        public bool IsFailure()
        {
            return !IsSuccess();
        }
    }
}
