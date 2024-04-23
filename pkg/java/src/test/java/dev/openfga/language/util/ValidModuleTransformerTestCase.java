package dev.openfga.language.util;

import java.util.ArrayList;

import dev.openfga.language.ModulesToModelTransformer.ModuleFile;

public class ValidModuleTransformerTestCase {
    private final String name;
    private final ArrayList<ModuleFile>  modules;
    private final String json;
    private final String combinedDsl;
    private final String combinedDslWithSourceInfo;
    private final boolean skip;

    public ValidModuleTransformerTestCase(String name, ArrayList<ModuleFile> modules, String json, String combinedDsl, String combinedDslWithSourceInfo, boolean skip) {
        this.name = name;
        this.modules = modules;
        this.json = json;
        this.skip = skip;
        this.combinedDsl = combinedDsl;
        this.combinedDslWithSourceInfo = combinedDslWithSourceInfo;
    }

    public String getName() {
        return name;
    }

    public ArrayList<ModuleFile>  getModules() {
        return modules;
    }

    public String getJson() {
        return json;
    }

    public String getCombinedDsl() {
        return combinedDsl;
    }

    public String getCombinedDslWithSourceInfo() {
        return combinedDslWithSourceInfo;
    }

    public boolean isSkip() {
        return skip;
    }
}
