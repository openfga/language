package dev.openfga.language.errors;

import java.util.List;

public class ModFileValidationError extends Exception {

    private final List<ModFileValidationSingleError> errors;

    public ModFileValidationError(List<ModFileValidationSingleError> errors) {
        super(Errors.messagesFromErrors(errors));
        this.errors = errors;
    }

    public List<ModFileValidationSingleError> getErrors() {
        return errors;
    }
}
