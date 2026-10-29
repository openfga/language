package dev.openfga.language.errors;

public class UnsupportedModularModulesException extends RuntimeException {
    public UnsupportedModularModulesException(String schemaVersion) {
        super("model schema version " + schemaVersion + " does not support modules");
    }
}
