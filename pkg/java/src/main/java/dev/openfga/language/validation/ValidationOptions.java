package dev.openfga.language.validation;

import dev.openfga.language.validation.Validator.Rules;

public class ValidationOptions {

    private String typePattern = String.format("^%s$", Rules.TYPE);
    private String relationPattern = String.format("^%s$", Rules.RELATION);

    public String getTypePattern() {
        return typePattern;
    }

    public void setTypePattern(String typePattern) {
        this.typePattern = typePattern;
    }

    public String getRelationPattern() {
        return relationPattern;
    }

    public void setRelationPattern(String relationPattern) {
        this.relationPattern = relationPattern;
    }
}
