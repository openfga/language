package dev.openfga.language.errors;

import lombok.Getter;

@Getter
public enum ErrorType {
    SYNTAX("syntax"),
    VALIDATION("validation");

    private final String value;

    ErrorType(String value) {
        this.value = value;
    }

    @Override
    public String toString() {
        return value;
    }
}