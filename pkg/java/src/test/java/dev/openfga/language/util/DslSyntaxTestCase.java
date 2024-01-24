package dev.openfga.language.util;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import dev.openfga.language.errors.ModelValidationSingleError;
import lombok.Getter;
import lombok.Setter;

import java.util.List;

@Getter
@Setter
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
    private List<ModelValidationSingleError> expectedErrors;
}
