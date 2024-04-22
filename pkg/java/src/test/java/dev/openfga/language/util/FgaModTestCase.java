package dev.openfga.language.util;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import dev.openfga.language.errors.ModFileValidationSingleError;
import java.util.ArrayList;
import java.util.List;

public class FgaModTestCase {
    @JsonProperty("name")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private String name;

    @JsonProperty("modFile")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private String modFile;

    @JsonProperty("json")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private String json;

    @JsonProperty("skip")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private boolean skip;

    @JsonProperty("expected_errors")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private List<ModFileValidationSingleError> expectedErrors = new ArrayList<>();

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getModFile() {
        return modFile;
    }

    public void setDsl(String modFile) {
        this.modFile = modFile;
    }

    public String getJson() {
        return json;
    }

    public void setJson(String json) {
        this.json = json;
    }

    public boolean isSkip() {
        return skip;
    }

    public void setSkip(boolean skip) {
        this.skip = skip;
    }

    public List<ModFileValidationSingleError> getExpectedErrors() {
        return expectedErrors;
    }

    public void setExpectedErrors(List<ModFileValidationSingleError> expectedErrors) {
        this.expectedErrors = expectedErrors;
    }
}
