package dev.openfga.language.validation;

import lombok.AllArgsConstructor;
import lombok.Getter;

import java.util.regex.Pattern;
import java.util.regex.PatternSyntaxException;

@Getter
@AllArgsConstructor(access = lombok.AccessLevel.PRIVATE)
class ValidationRegex {

    private final String rule;
    private final Pattern regex;

    public static ValidationRegex build(String name, String rule) {
        Pattern regex = null;
        try {
            regex = Pattern.compile(rule);
        } catch (PatternSyntaxException e) {
            var message = "Incorrect " + name + " regex specification for  + rule";
            throw new IllegalArgumentException(message);
        }
        return new ValidationRegex(rule, regex);
    }

    public boolean matches(String input) {
        return regex.matcher(input).matches();
    }
}


