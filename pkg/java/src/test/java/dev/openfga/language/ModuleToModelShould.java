package dev.openfga.language;

import static org.junit.jupiter.params.provider.Arguments.arguments;

import java.util.ArrayList;
import java.util.stream.Stream;

import org.junit.jupiter.api.Assumptions;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.MethodSource;

import org.junit.jupiter.params.provider.Arguments;

import dev.openfga.language.ModulesToModelTransformer.ModuleFile;
import dev.openfga.language.util.TestsData;

public class ModuleToModelShould {

    @ParameterizedTest(name = "{0}")
    @MethodSource("transformerModuleTestCases")
    public void transformModuleToModel(String name, ArrayList<ModuleFile> modules, String json, boolean skip) throws Exception {
        Assumptions.assumeFalse(skip);
    }

    @ParameterizedTest(name = "{0}")
    @MethodSource("transformerModuleErrorTestCases")
    public void transformModuleToModelErrorCases(String name, ArrayList<ModuleFile> modules, String json, boolean skip) throws Exception {
        Assumptions.assumeFalse(skip);
    }

    private static Stream<Arguments> transformerModuleTestCases() {
        return TestsData.VALID_MODULE_TRANSFORMER_TESTS_CASES.stream()
            .map(testCase -> 
                arguments(testCase.getName(), testCase.getModules(), testCase.getJson(), testCase.isSkip()));
    }

    private static Stream<Arguments> transformerModuleErrorTestCases() {
        return TestsData.INVALID_MODULE_TRANSFORMER_TESTS_CASES.stream()
            .map(testCase -> 
                arguments(testCase.getName(), testCase.getModules(), testCase.getExpectedErrors(), testCase.isSkip()));
    }
}
