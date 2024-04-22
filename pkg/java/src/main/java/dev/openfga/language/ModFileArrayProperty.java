package dev.openfga.language;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import dev.openfga.language.errors.StartEnd;
import java.util.List;

public class ModFileArrayProperty {
    public static final String JSON_PROPERTY_VALUE = "value";
    private List<ModFileStringProperty> value;

    public static final String JSON_PROPERTY_LINE = "line";
    private StartEnd line;

    public static final String JSON_PROPERTY_COLUMN = "column";
    private StartEnd column;

    public ModFileArrayProperty() {}

    @JsonProperty(JSON_PROPERTY_VALUE)
    @JsonInclude(value = JsonInclude.Include.ALWAYS)
    public List<ModFileStringProperty> getValue() {
        return value;
    }

    public ModFileArrayProperty value(List<ModFileStringProperty> value) {
        this.value = value;
        return this;
    }

    @JsonProperty(JSON_PROPERTY_LINE)
    @JsonInclude(value = JsonInclude.Include.ALWAYS)
    public StartEnd getLine() {
        return line;
    }

    public ModFileArrayProperty line(StartEnd line) {
        this.line = line;
        return this;
    }

    @JsonProperty(JSON_PROPERTY_COLUMN)
    @JsonInclude(value = JsonInclude.Include.ALWAYS)
    public StartEnd getColumn() {
        return column;
    }

    public ModFileArrayProperty column(StartEnd column) {
        this.column = column;
        return this;
    }
}
