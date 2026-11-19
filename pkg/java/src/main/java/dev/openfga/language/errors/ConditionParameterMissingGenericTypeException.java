package dev.openfga.language.errors;

public class ConditionParameterMissingGenericTypeException extends RuntimeException {
    public ConditionParameterMissingGenericTypeException(String conditionName, String parameterName) {
        super(String.format(
                "the '%s' parameter of the '%s' condition has no generic type", parameterName, conditionName));
    }
}
