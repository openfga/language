package dev.openfga.language.util;

import lombok.EqualsAndHashCode;
import lombok.Getter;

@Getter
@EqualsAndHashCode
public final class ValidTransformerTestCase {
    private final String name;
    private final String dsl;
    private final String json;
    private final boolean skip;

    public ValidTransformerTestCase(String name, String dsl, String json, boolean skip) {
        this.name = name;
        this.dsl = dsl;
        this.json = json;
        this.skip = skip;
    }

    @Override
    public String toString() {
        return "TransformerTestCase[" +
                "name=" + name + ", " +
                "dsl=" + dsl + ", " +
                "json=" + json + ", " +
                "skip=" + skip + ']';
    }

}

