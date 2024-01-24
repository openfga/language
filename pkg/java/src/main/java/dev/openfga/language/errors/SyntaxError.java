package dev.openfga.language.errors;

import dev.openfga.sdk.api.model.Metadata;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.antlr.v4.runtime.RecognitionException;

@Getter
@Setter
@EqualsAndHashCode(callSuper = true)
@NoArgsConstructor
public class SyntaxError extends ParsingError {

    private Metadata metadata;
    private RecognitionException cause;

    public SyntaxError(ErrorProperties properties, Metadata metadata, RecognitionException cause) {
        super(ErrorType.SYNTAX.getValue(), properties);
        this.metadata = metadata;
        this.cause = cause;
    }
}
