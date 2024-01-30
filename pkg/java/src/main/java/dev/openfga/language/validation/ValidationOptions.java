package dev.openfga.language.validation;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class ValidationOptions {

    private static final String DEFAULT_TYPE_PATTERN = "^[^:#@\\s]{1,254}$";
    private static final String DEFAULT_RELATION_PATTERN = "^[^:#@\\s]{1,50}$";

    private String typePattern = DEFAULT_TYPE_PATTERN;
    private String relationPattern = DEFAULT_RELATION_PATTERN;
}
