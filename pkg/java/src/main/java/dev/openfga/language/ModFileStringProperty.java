package dev.openfga.language;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import dev.openfga.language.errors.StartEnd;

public class ModFileStringProperty {
    public static final String JSON_PROPERTY_VALUE = "value";
    private String value;

    public static final String JSON_PROPERTY_LINE = "line";
    private StartEnd line;

    public static final String JSON_PROPERTY_COLUMN = "column";
    private StartEnd column;

    public ModFileStringProperty() {}

    @JsonProperty(JSON_PROPERTY_VALUE)
    @JsonInclude(value = JsonInclude.Include.ALWAYS)
    public String getValue() {
        return value;
    }

    public ModFileStringProperty value(String value) {
        this.value = value;
        return this;
    }

    @JsonProperty(JSON_PROPERTY_LINE)
    @JsonInclude(value = JsonInclude.Include.ALWAYS)
    public StartEnd getLine() {
        return line;
    }

    public ModFileStringProperty line(StartEnd line) {
        this.line = line;
        return this;
    }

    @JsonProperty(JSON_PROPERTY_COLUMN)
    @JsonInclude(value = JsonInclude.Include.ALWAYS)
    public StartEnd getColumn() {
        return column;
    }

    public ModFileStringProperty column(StartEnd column) {
        this.column = column;
        return this;
    }
}
