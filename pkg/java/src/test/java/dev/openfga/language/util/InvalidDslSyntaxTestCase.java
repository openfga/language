package dev.openfga.language.util;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class InvalidDslSyntaxTestCase {

    @JsonProperty("name")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private String name;

    @JsonProperty("dsl")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private String dsl;

    @JsonProperty("skip")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private boolean skip;

    @JsonProperty("error_message")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private String errorMessage;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDsl() {
        return dsl;
    }

    public void setDsl(String dsl) {
        this.dsl = dsl;
    }

    public boolean isSkip() {
        return skip;
    }

    public void setSkip(boolean skip) {
        this.skip = skip;
    }

    public String getErrorMessage() {
        return errorMessage;
    }

    public void setErrorMessage(String errorMessage) {
        this.errorMessage = errorMessage;
    }
}
