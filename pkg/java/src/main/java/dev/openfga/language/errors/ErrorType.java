package dev.openfga.language.errors;

public enum ErrorType {
    SYNTAX("syntax"),
    VALIDATION("validation");

    private final String value;

    ErrorType(String value) {
        this.value = value;
    }

    public String getValue() {
        return value;
    }

    @Override
    public String toString() {
        return value;
    }
}