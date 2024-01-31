package dev.openfga.language.errors;

import lombok.*;

@Getter
@Setter
@NoArgsConstructor
public class ModelValidationSingleError extends ParsingError {

    private ValidationMetadata metadata;

    public ModelValidationSingleError(ErrorProperties properties, ValidationMetadata metadata) {
        super("syntax", properties);
        this.metadata = metadata;
    }
}