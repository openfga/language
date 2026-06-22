package dev.openfga.language.errors;

import java.util.List;

public class ModuleTransformationError extends Exception {

    private final List<? extends ParsingError> errors;

    public ModuleTransformationError(List<? extends ParsingError> errors) {
        super(Errors.messagesFromErrors(errors));
        this.errors = errors;
    }

    public List<? extends ParsingError> getErrors() {
        return errors;
    }
}
