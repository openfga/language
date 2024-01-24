package dev.openfga.language.errors;

import lombok.Getter;

import java.util.Collection;
import java.util.List;

import static java.util.stream.Collectors.joining;

@Getter
public abstract class Errors<T> extends SimpleError {

    private final List<T> errors;

    public Errors(List<T> errors) {
        super(messagesFromErrors(errors));
        this.errors = errors;
    }


    static String messagesFromErrors(Collection<?> errors) {
        var delimiter = "\n\t* ";
        var errorsPlural = errors.size() > 1 ? "s" : "";
        var prefix = String.format("%d error%s occurred:%s", errors.size(), errorsPlural, delimiter);
        var suffix = "\n\n";
        return errors.stream()
                .map(Object::toString)
                .collect(joining("\n\t* ", prefix, suffix));
    }
}