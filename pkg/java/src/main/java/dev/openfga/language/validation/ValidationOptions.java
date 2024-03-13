package dev.openfga.language.validation;

public class ValidationOptions {

    private static final String DEFAULT_TYPE_PATTERN = "^[^:#@\\s]{1,254}$";
    private static final String DEFAULT_RELATION_PATTERN = "^[^:#@\\s]{1,50}$";

    private String typePattern = DEFAULT_TYPE_PATTERN;
    private String relationPattern = DEFAULT_RELATION_PATTERN;

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
