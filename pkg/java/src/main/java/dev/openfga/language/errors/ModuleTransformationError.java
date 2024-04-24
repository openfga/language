package dev.openfga.language.errors;

import java.util.List;

public class ModuleTransformationError extends Exception {
    private final List<Object> errors;

    public ModuleTransformationError(List<Object> errors) {
        super(Errors.messagesFromErrors(errors));
        this.errors = errors;
    }

    public List<Object> getErrors() {
        return errors;
    }
}
