package dev.openfga.language.validation;

import java.util.List;

class AllowableTypesResult {
    private final boolean valid;
    private final List<String> allowableTypes;

    public AllowableTypesResult(boolean valid, List<String> allowableTypes) {
        this.valid = valid;
        this.allowableTypes = allowableTypes;
    }

    public boolean isValid() {
        return valid;
    }

    public List<String> getAllowableTypes() {
        return allowableTypes;
    }
}
