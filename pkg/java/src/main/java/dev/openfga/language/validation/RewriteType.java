package dev.openfga.language.validation;

public enum RewriteType {
    Direct("direct"),
    ComputedUserset("computed_userset"),
    TupleToUserset("tuple_to_userset");

    private final String value;

    RewriteType(String value) {
        this.value = value;
    }
}
