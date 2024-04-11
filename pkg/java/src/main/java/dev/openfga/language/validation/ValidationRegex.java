package dev.openfga.language.validation;

import java.util.regex.Pattern;
import java.util.regex.PatternSyntaxException;

class ValidationRegex {

    private final String rule;
    private final Pattern regex;

    private ValidationRegex(String rule, Pattern regex) {
        this.rule = rule;
        this.regex = regex;
    }

    public String getRule() {
        return rule;
    }

    public Pattern getRegex() {
        return regex;
    }

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
