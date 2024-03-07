package dev.openfga.language.util;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public final class JsonSyntaxTestCase {

    @JsonProperty("name")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private String name;

    @JsonProperty("json")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private String json;

    @JsonProperty("error_message")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private String errorMessage;

    @JsonProperty("skip")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private boolean skip;

    @JsonProperty("valid")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private boolean valid;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getJson() {
        return json;
    }

    public void setJson(String json) {
        this.json = json;
    }

    public String getErrorMessage() {
        return errorMessage;
    }

    public void setErrorMessage(String errorMessage) {
        this.errorMessage = errorMessage;
    }

    public boolean isSkip() {
        return skip;
    }

    public void setSkip(boolean skip) {
        this.skip = skip;
    }

    public boolean isValid() {
        return valid;
    }

    public void setValid(boolean valid) {
        this.valid = valid;
    }
}
