package dev.openfga.language.util;

import java.util.List;
import dev.openfga.language.ModulesToModelTransformer.ModuleFile;
import dev.openfga.language.errors.ModuleTransformationSingleError;

import java.util.ArrayList;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.databind.ObjectMapper;

public class InvalidModuleTransformerTestCases {
    private final String name;
    private final ArrayList<ModuleFile> modules;
    private final List<ModuleTransformationSingleError> expectedErrors;
    private final boolean skip;

    public InvalidModuleTransformerTestCases(
            String name, ArrayList<ModuleFile> modules, String expectedErrors, boolean skip) throws JsonMappingException, JsonProcessingException {
        this.name = name;
        this.modules = modules;
        // this.expectedErrors = expectedErrors;
        ObjectMapper objectMapper = new ObjectMapper();
        this.expectedErrors = objectMapper.readValue(expectedErrors, new TypeReference<List<ModuleTransformationSingleError>>(){});
        this.skip = skip;
    }

    public String getName() {
        return name;
    }

    public ArrayList<ModuleFile> getModules() {
        return modules;
    }

    public List<ModuleTransformationSingleError> getExpectedErrors() {
        return expectedErrors;
    }

    public boolean isSkip() {
        return skip;
    }
}
