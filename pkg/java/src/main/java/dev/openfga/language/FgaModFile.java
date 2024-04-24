package dev.openfga.language;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class FgaModFile {

    public static final String JSON_PROPERTY_SCHEMA = "schema";
    private ModFileStringProperty schema;

    public static final String JSON_PROPERTY_CONTENTS = "contents";
    private ModFileArrayProperty contents;

    public FgaModFile() {}

    public FgaModFile schema(ModFileStringProperty schema) {
        this.schema = schema;
        return this;
    }

    @JsonProperty(JSON_PROPERTY_SCHEMA)
    @JsonInclude(value = JsonInclude.Include.ALWAYS)
    public ModFileStringProperty getSchema() {
        return this.schema;
    }

    public FgaModFile contents(ModFileArrayProperty contents) {
        this.contents = contents;
        return this;
    }

    @JsonProperty(JSON_PROPERTY_CONTENTS)
    @JsonInclude(value = JsonInclude.Include.ALWAYS)
    public ModFileArrayProperty getContents() {
        return this.contents;
    }
}
