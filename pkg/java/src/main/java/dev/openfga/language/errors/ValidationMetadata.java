package dev.openfga.language.errors;

import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.NoArgsConstructor;

@Data
@EqualsAndHashCode(callSuper = true)
@NoArgsConstructor
public class ValidationMetadata extends Metadata {
    private ValidationError errorType;
    private String relation;
    private String typeName;
    private String conditionName;

    public ValidationMetadata(String symbol, ValidationError errorType) {
        this(symbol, errorType, null, null, null);
    }
    public ValidationMetadata(String symbol, ValidationError errorType, String relation, String typeName, String conditionName) {
        super(symbol);
        this.errorType = errorType;
        this.relation = relation;
        this.typeName = typeName;
        this.conditionName = conditionName;
    }
}