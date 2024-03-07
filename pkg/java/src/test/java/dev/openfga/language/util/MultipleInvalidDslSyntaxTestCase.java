package dev.openfga.language.util;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import dev.openfga.language.errors.ModelValidationSingleError;

import java.util.ArrayList;
import java.util.List;

public class MultipleInvalidDslSyntaxTestCase extends InvalidDslSyntaxTestCase {

    @JsonProperty("expected_errors")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private List<ModelValidationSingleError> expectedErrors = new ArrayList<>();

    public List<ModelValidationSingleError> getExpectedErrors() {
        return expectedErrors;
    }

    public void setExpectedErrors(List<ModelValidationSingleError> expectedErrors) {
        this.expectedErrors = expectedErrors;
    }
}
