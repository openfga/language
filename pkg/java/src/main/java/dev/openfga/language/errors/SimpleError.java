package dev.openfga.language.errors;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.*;

@Getter
@Setter
@EqualsAndHashCode
@NoArgsConstructor
@AllArgsConstructor
public abstract class SimpleError {

    @JsonProperty("msg")
    private String message;

    public String toString() {
        return message;
    }
}