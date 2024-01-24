package dev.openfga.language.util;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
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
}
