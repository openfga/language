package dev.openfga.language.util;

import java.util.ArrayList;
import java.util.List;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

import dev.openfga.language.errors.ModelValidationSingleError;

public class JsonValidationTestCase {
    @JsonProperty("name")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private String name;

    @JsonProperty("json")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private String json;

    @JsonProperty("expected_errors")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private List<ModelValidationSingleError> expectedErrors = new ArrayList<>();



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

    public List<ModelValidationSingleError> getExpectedErrors() {
        return expectedErrors;
    }

    public void setExpectedErrors(List<ModelValidationSingleError> expectedErrors) {
        this.expectedErrors = expectedErrors;
    }
}
