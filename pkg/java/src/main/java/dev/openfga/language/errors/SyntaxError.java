package dev.openfga.language.errors;

import org.antlr.v4.runtime.RecognitionException;

public class SyntaxError extends ParsingError {

    private Metadata metadata;
    private RecognitionException cause;

    public SyntaxError(ErrorProperties properties, Metadata metadata, RecognitionException cause) {
        super(ErrorType.SYNTAX.getValue(), properties);
        this.metadata = metadata;
        this.cause = cause;
    }

    public Metadata getMetadata() {
        return metadata;
    }

    public void setMetadata(Metadata metadata) {
        this.metadata = metadata;
    }

    public RecognitionException getCause() {
        return cause;
    }

    public void setCause(RecognitionException cause) {
        this.cause = cause;
    }
}
