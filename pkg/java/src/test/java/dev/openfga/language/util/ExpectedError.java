package dev.openfga.language.util;

import com.fasterxml.jackson.annotation.JsonProperty;
import dev.openfga.language.errors.StartEnd;

public final class ExpectedError {
    @JsonProperty("msg")
    private String message;

    private StartEnd line;
    private StartEnd column;

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public StartEnd getLine() {
        return line;
    }

    public void setLine(StartEnd line) {
        this.line = line;
    }

    public StartEnd getColumn() {
        return column;
    }

    public void setColumn(StartEnd column) {
        this.column = column;
    }
}
