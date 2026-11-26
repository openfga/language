package dev.openfga.language.util;

import dev.openfga.language.ModuleFile;
import java.util.List;

public final class ModuleTransformerTestCase {
    private final String name;
    private final List<ModuleFile> modules;
    private final String json;
    private final String dsl;
    private final String dslWithSourceInfo;
    private final List<String> expectedModules;
    private final List<ModuleExpectedError> expectedErrors;
    private final boolean skip;

    public ModuleTransformerTestCase(
            String name,
            List<ModuleFile> modules,
            String json,
            String dsl,
            String dslWithSourceInfo,
            List<String> expectedModules,
            List<ModuleExpectedError> expectedErrors,
            boolean skip) {
        this.name = name;
        this.modules = modules;
        this.json = json;
        this.dsl = dsl;
        this.dslWithSourceInfo = dslWithSourceInfo;
        this.expectedModules = expectedModules;
        this.expectedErrors = expectedErrors;
        this.skip = skip;
    }

    public String getName() {
        return name;
    }

    public List<ModuleFile> getModules() {
        return modules;
    }

    public String getJson() {
        return json;
    }

    public String getDsl() {
        return dsl;
    }

    public String getDslWithSourceInfo() {
        return dslWithSourceInfo;
    }

    public List<String> getExpectedModules() {
        return expectedModules;
    }

    public List<ModuleExpectedError> getExpectedErrors() {
        return expectedErrors;
    }

    public boolean isSkip() {
        return skip;
    }
}
