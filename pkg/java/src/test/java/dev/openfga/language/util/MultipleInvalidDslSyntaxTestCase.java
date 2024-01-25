package dev.openfga.language.util;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import dev.openfga.language.errors.ModelValidationSingleError;
import lombok.Getter;
import lombok.Setter;

import java.util.ArrayList;
import java.util.List;

@Getter
@Setter
public class MultipleInvalidDslSyntaxTestCase extends InvalidDslSyntaxTestCase {

    @JsonProperty("expected_errors")
    @JsonInclude(JsonInclude.Include.USE_DEFAULTS)
    private List<ModelValidationSingleError> expectedErrors = new ArrayList<>();
}
