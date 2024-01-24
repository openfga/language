package dev.openfga.language.errors;

public class UnsupportedDSLNestingException extends RuntimeException {
    public UnsupportedDSLNestingException(String typeName, String relationName) {
        super(String.format("the '%s' relation under the '%s' type has mixed operators which are not supported by the OpenFGA DSL syntax yet", relationName, typeName));
    }
}