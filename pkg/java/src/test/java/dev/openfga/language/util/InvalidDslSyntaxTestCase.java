package dev.openfga.language.util;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
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
}
