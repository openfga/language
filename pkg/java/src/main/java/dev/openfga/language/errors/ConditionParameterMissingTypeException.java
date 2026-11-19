package dev.openfga.language.errors;

public class ConditionParameterMissingTypeException extends RuntimeException {
    public ConditionParameterMissingTypeException(String conditionName, String parameterName) {
        super(String.format("the '%s' parameter of the '%s' condition has no type", parameterName, conditionName));
    }
}
