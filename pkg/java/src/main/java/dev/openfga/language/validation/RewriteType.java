package dev.openfga.language.validation;

import lombok.RequiredArgsConstructor;

@RequiredArgsConstructor
public enum RewriteType {

    Direct("direct"),
    ComputedUserset("computed_userset"),
    TupleToUserset("tuple_to_userset");

    private final String value;
}
