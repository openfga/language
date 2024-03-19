package dev.openfga.language.util;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import dev.openfga.language.errors.ModelValidationSingleError;

import java.util.ArrayList;
import java.util.List;

public class DslSyntaxTestCase {

    @JsonProperty("name")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private String name;

    @JsonProperty("dsl")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private String dsl;

    @JsonProperty("valid")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private boolean valid;

    @JsonProperty("skip")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private boolean skip;

    @JsonProperty("expected_errors")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private List<ModelValidationSingleError> expectedErrors = new ArrayList<>();

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

    public boolean isValid() {
        return valid;
    }

    public void setValid(boolean valid) {
        this.valid = valid;
    }

    public boolean isSkip() {
        return skip;
    }

    public void setSkip(boolean skip) {
        this.skip = skip;
    }

    public List<ModelValidationSingleError> getExpectedErrors() {
        return expectedErrors;
    }

    public void setExpectedErrors(List<ModelValidationSingleError> expectedErrors) {
        this.expectedErrors = expectedErrors;
    }
}
