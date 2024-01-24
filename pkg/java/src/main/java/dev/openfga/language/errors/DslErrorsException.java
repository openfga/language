package dev.openfga.language.errors;

import lombok.Getter;

import java.util.List;

@Getter
public class DslErrorsException extends Exception {

    private final List<? extends ParsingError> errors;

    public DslErrorsException(List<? extends ParsingError> errors) {
        super(Errors.messagesFromErrors(errors));
        this.errors = errors;
    }
}