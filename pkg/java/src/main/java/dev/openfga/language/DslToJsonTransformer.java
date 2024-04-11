package dev.openfga.language;

import static java.util.stream.Collectors.joining;

import dev.openfga.language.antlr.OpenFGALexer;
import dev.openfga.language.antlr.OpenFGAParser;
import dev.openfga.language.errors.DslErrorsException;
import dev.openfga.language.errors.SyntaxError;
import dev.openfga.sdk.api.model.AuthorizationModel;
import java.io.IOException;
import java.util.Arrays;
import java.util.List;
import java.util.regex.Pattern;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.tree.ParseTreeWalker;

public class DslToJsonTransformer {
    public String transform(String dsl) throws IOException, DslErrorsException {
        return JSON.stringify(parseAuthorisationModel(dsl));
    }

    private AuthorizationModel parseAuthorisationModel(String dsl) throws DslErrorsException {
        var result = parseDsl(dsl);
        if (result.IsFailure()) {
            throw new DslErrorsException(result.getErrors());
        }

        return result.getAuthorizationModel();
    }

    private static final Pattern SPACES_LINE_PATTERN = Pattern.compile("^\\s*$");
    private static final Pattern COMMENTED_LINE_PATTERN = Pattern.compile("^\\s*#.*$");

    private String cleanLine(String line) {
        if (SPACES_LINE_PATTERN.matcher(line).matches()
                || COMMENTED_LINE_PATTERN.matcher(line).matches()) {
            return "";
        }
        var cleanedLine = line.split(Pattern.quote(" #"))[0];
        return cleanedLine.stripTrailing();
    }

    public Result parseDsl(String dsl) {
        var cleanedDsl = Arrays.stream(dsl.split("\n")).map(this::cleanLine).collect(joining("\n"));

        var antlrStream = CharStreams.fromString(cleanedDsl);
        var errorListener = new OpenFgaDslErrorListener();

        var lexer = new OpenFGALexer(antlrStream);
        lexer.removeErrorListeners();
        lexer.addErrorListener(errorListener);
        var tokenStream = new CommonTokenStream(lexer);

        var parser = new OpenFGAParser(tokenStream);
        parser.removeErrorListeners();
        parser.addErrorListener(errorListener);

        var listener = new OpenFgaDslListener(parser);
        new ParseTreeWalker().walk(listener, parser.main());

        return new Result(listener.getAuthorizationModel(), errorListener.getErrors());
    }

    public static final class Result {
        private final AuthorizationModel authorizationModel;
        private final List<SyntaxError> errors;

        public Result(AuthorizationModel authorizationModel, List<SyntaxError> errors) {
            this.authorizationModel = authorizationModel;
            this.errors = errors;
        }

        public AuthorizationModel getAuthorizationModel() {
            return authorizationModel;
        }

        public List<SyntaxError> getErrors() {
            return errors;
        }

        public boolean IsSuccess() {
            return errors.isEmpty();
        }

        public boolean IsFailure() {
            return !IsSuccess();
        }
    }
}
