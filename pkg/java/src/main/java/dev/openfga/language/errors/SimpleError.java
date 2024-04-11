package dev.openfga.language.errors;

import com.fasterxml.jackson.annotation.JsonProperty;

public abstract class SimpleError {

    @JsonProperty("msg")
    private String message;

    public SimpleError() {}

    public SimpleError(String message) {
        this.message = message;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public String toString() {
        return message;
    }
}
