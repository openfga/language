package dev.openfga.language.errors;

import java.util.List;

public class DslErrorsException extends Exception {

    private final List<? extends ParsingError> errors;

    public DslErrorsException(List<? extends ParsingError> errors) {
        super(Errors.messagesFromErrors(errors));
        this.errors = errors;
    }

    public List<? extends ParsingError> getErrors() {
        return errors;
    }
}