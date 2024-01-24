package dev.openfga.language.util;

import com.fasterxml.jackson.annotation.JsonProperty;
import dev.openfga.language.errors.StartEnd;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public final class ExpectedError {
    @JsonProperty("msg")
    private String message;
    private StartEnd line;
    private StartEnd column;
}