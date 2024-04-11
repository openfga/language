package dev.openfga.language.util;

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

    public String getName() {
        return name;
    }

    public String getDsl() {
        return dsl;
    }

    public String getJson() {
        return json;
    }

    public boolean isSkip() {
        return skip;
    }

    @Override
    public String toString() {
        return "TransformerTestCase[" + "name="
                + name + ", " + "dsl="
                + dsl + ", " + "json="
                + json + ", " + "skip="
                + skip + ']';
    }
}
