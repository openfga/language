package dev.openfga.language.errors;

import com.networknt.schema.ValidationMessage;

public class StoreValidationSingleError extends SimpleError {

    private ValidationMessage message;

    public StoreValidationSingleError(ValidationMessage message) {
        super(message.toString());
        this.message = message;
    }

    public ValidationMessage getValidationMessage() {
        return this.message;
    }
}
