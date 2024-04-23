package dev.openfga.language.util;

import java.util.ArrayList;

import dev.openfga.language.ModulesToModelTransformer.ModuleFile;


public class InvalidModuleTransformerTestCases {
    private final String name;
    private final ArrayList<ModuleFile> modules;
    private final String expectedErrors;
    private final boolean skip;

    public InvalidModuleTransformerTestCases(String name, ArrayList<ModuleFile> modules, String expectedErrors, boolean skip) {
        this.name = name;
        this.modules = modules;
        this.expectedErrors = expectedErrors;
        this.skip = skip;
    }

    public String getName() {
        return name;
    }

    public ArrayList<ModuleFile> getModules() {
        return modules;
    }

    public String getExpectedErrors() {
        return expectedErrors;
    }

    public boolean isSkip() {
        return skip;
    }
}
