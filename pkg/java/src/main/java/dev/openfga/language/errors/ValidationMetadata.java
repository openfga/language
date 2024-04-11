package dev.openfga.language.errors;

public class ValidationMetadata extends Metadata {
    private ValidationError errorType;
    private String relation;
    private String typeName;
    private String conditionName;

    // Needed for Jackson deserialization
    public ValidationMetadata() {}

    public ValidationMetadata(String symbol, ValidationError errorType) {
        this(symbol, errorType, null, null, null);
    }

    public ValidationMetadata(
            String symbol, ValidationError errorType, String relation, String typeName, String conditionName) {
        super(symbol);
        this.errorType = errorType;
        this.relation = relation;
        this.typeName = typeName;
        this.conditionName = conditionName;
    }

    public ValidationError getErrorType() {
        return errorType;
    }

    public void setErrorType(ValidationError errorType) {
        this.errorType = errorType;
    }

    public String getRelation() {
        return relation;
    }

    public void setRelation(String relation) {
        this.relation = relation;
    }

    public String getTypeName() {
        return typeName;
    }

    public void setTypeName(String typeName) {
        this.typeName = typeName;
    }

    public String getConditionName() {
        return conditionName;
    }

    public void setConditionName(String conditionName) {
        this.conditionName = conditionName;
    }
}
