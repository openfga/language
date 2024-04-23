package dev.openfga.language.errors;

import java.util.List;

public class ModuleTransformationError extends Exception {
    private final List<ModuleTransformationSingleError> errors;

    public ModuleTransformationError(List<ModuleTransformationSingleError> errors) {
        super(Errors.messagesFromErrors(errors));
        this.errors = errors;
    }

    public List<ModuleTransformationSingleError> getErrors() {
        return errors;
    }
}
