package dev.openfga.language;

import static java.util.Collections.unmodifiableList;

import dev.openfga.language.errors.ErrorProperties;
import dev.openfga.language.errors.Metadata;
import dev.openfga.language.errors.StartEnd;
import dev.openfga.language.errors.SyntaxError;
import java.util.ArrayList;
import java.util.BitSet;
import java.util.List;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.ATNConfigSet;
import org.antlr.v4.runtime.dfa.DFA;

public class OpenFgaDslErrorListener implements ANTLRErrorListener {

    private final List<SyntaxError> errors = new ArrayList<>();

    public List<SyntaxError> getErrors() {
        return unmodifiableList(errors);
    }

    @Override
    // line is one based, i.e. the first line will be 1
    // column is zero based, i.e. the first column will 0
    public void syntaxError(
            Recognizer<?, ?> recognizer,
            Object offendingSymbol,
            int line,
            int column,
            String message,
            RecognitionException e) {
        Metadata metadata = null;
        var columnOffset = 0;

        if (offendingSymbol instanceof Token) {
            metadata = new Metadata(((Token) offendingSymbol).getText());
            columnOffset = metadata.getSymbol().length();
        }

        var properties =
                new ErrorProperties(new StartEnd(line, line), new StartEnd(column, column + columnOffset), message);
        this.errors.add(new SyntaxError(properties, metadata, e));
    }

    @Override
    public void reportAmbiguity(
            Parser parser, DFA dfa, int i, int i1, boolean b, BitSet bitSet, ATNConfigSet atnConfigSet) {}

    @Override
    public void reportAttemptingFullContext(
            Parser parser, DFA dfa, int i, int i1, BitSet bitSet, ATNConfigSet atnConfigSet) {}

    @Override
    public void reportContextSensitivity(Parser parser, DFA dfa, int i, int i1, int i2, ATNConfigSet atnConfigSet) {}
}
