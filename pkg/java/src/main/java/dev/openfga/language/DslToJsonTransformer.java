package dev.openfga.language;

import dev.openfga.language.antlr.OpenFGALexer;
import dev.openfga.language.antlr.OpenFGAParser;
import dev.openfga.language.errors.DslErrorsException;
import dev.openfga.language.errors.SyntaxError;
import dev.openfga.sdk.api.model.AuthorizationModel;
import lombok.Getter;
import org.antlr.v4.runtime.ANTLRInputStream;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.tree.ParseTreeWalker;

import java.io.IOException;
import java.util.Arrays;
import java.util.List;

import static java.util.stream.Collectors.joining;

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

    public Result parseDsl(String dsl) {
        var cleanedDsl = Arrays.stream(dsl.split("\n"))
                .map(String::stripTrailing)
                .collect(joining("\n"));


        var antlrStream = new ANTLRInputStream(cleanedDsl);
        var errorListener = new OpenFgaDslErrorListener();

        var lexer = new OpenFGALexer(antlrStream);
        lexer.removeErrorListeners();
        lexer.addErrorListener(errorListener);
        var tokenStream = new CommonTokenStream(lexer);

        OpenFGAParser parser = new OpenFGAParser(tokenStream);
        parser.removeErrorListeners();
        parser.addErrorListener(errorListener);

        var listener = new OpenFgaDslListener(parser);

        new ParseTreeWalker().walk(listener, parser.main());

        return new Result(listener.getAuthorizationModel(), errorListener.getErrors());
    }

    @Getter
    public static final class Result {
        private final AuthorizationModel authorizationModel;
        private final List<SyntaxError> errors;

        public Result(AuthorizationModel authorizationModel, List<SyntaxError> errors) {
            this.authorizationModel = authorizationModel;
            this.errors = errors;
        }

        public boolean IsSuccess() {
            return errors.isEmpty();
        }

        public boolean IsFailure() {
            return !IsSuccess();
        }
    }
}
