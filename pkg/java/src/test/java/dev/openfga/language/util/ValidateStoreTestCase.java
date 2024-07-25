package dev.openfga.language.util;

public final class ValidateStoreTestCase {
    private final String name;
    private final String store;
    private final String errors;

    public ValidateStoreTestCase(String name, String store, String errors) {
        this.name = name;
        this.store = store;
        this.errors = errors;
    }

    public String getName() {
        return name;
    }

    public String getStore() {
        return store;
    }

    public String getErrors() {
        return errors;
    }

    @Override
    public String toString() {
        return "TransformerTestCase[" + "name=" + name + ", " + "store=" + store + ", " + "json=" + errors + ", " + ']';
    }
}
