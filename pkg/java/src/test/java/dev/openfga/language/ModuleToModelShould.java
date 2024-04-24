package dev.openfga.language;

import static java.util.stream.Collectors.joining;
import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.catchThrowable;
import static org.junit.jupiter.params.provider.Arguments.arguments;

import dev.openfga.language.ModulesToModelTransformer.ModuleFile;
import dev.openfga.language.errors.ModuleTransformationError;
import dev.openfga.language.errors.ModuleTransformationSingleError;
import dev.openfga.language.errors.ParsingError;
import dev.openfga.language.util.TestsData;
import dev.openfga.sdk.api.model.AuthorizationModel;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Stream;
import org.junit.jupiter.api.Assumptions;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

public class ModuleToModelShould {

    @ParameterizedTest(name = "{0}")
    @MethodSource("transformerModuleTestCases")
    public void transformModuleToModel(String name, ArrayList<ModuleFile> modules, String json, boolean skip)
            throws Exception {
        Assumptions.assumeFalse(skip);

        var model = ModulesToModelTransformer.transformToModel(modules, "1.2");

        var expected = JSON.parse(json, AuthorizationModel.class);

        assertThat(model.getSchemaVersion()).isEqualTo(expected.getSchemaVersion());
        assertThat(model.getTypeDefinitions().toArray()).containsExactlyInAnyOrderElementsOf(expected.getTypeDefinitions());
        assertThat(model.getConditions()).containsAllEntriesOf(expected.getConditions());
    }

    @ParameterizedTest(name = "{0}")
    @MethodSource("transformerModuleErrorTestCases")
    public void transformModuleToModelErrorCases(String name, ArrayList<ModuleFile> modules, List<ModuleTransformationSingleError> expectedErrors, boolean skip)
            throws Exception {
        Assumptions.assumeFalse(skip);
        
        var thrown = catchThrowable(() -> ModulesToModelTransformer.transformToModel(modules, "1.2"));

        assertThat(thrown).isInstanceOf(ModuleTransformationError.class);

        var errorsCount = expectedErrors.size();

        var formattedErrors = expectedErrors.stream()
            .map(error -> String.format(
                    "transformation error at line=%d, column=%d: %s",
                    error.getLine().getStart(), error.getColumn().getStart(), error.getMessage()))
            .collect(joining("\n\t* "));

        var expectedMessage = String.format(
            "%d error%s occurred:\n\t* %s\n\n", errorsCount, errorsCount > 1 ? "s" : "", formattedErrors);

        assertThat(thrown).hasMessage(expectedMessage);
    }

    private static Stream<Arguments> transformerModuleTestCases() {
        return TestsData.VALID_MODULE_TRANSFORMER_TESTS_CASES.stream()
                .map(testCase ->
                        arguments(testCase.getName(), testCase.getModules(), testCase.getJson(), testCase.isSkip()));
    }

    private static Stream<Arguments> transformerModuleErrorTestCases() {
        return TestsData.INVALID_MODULE_TRANSFORMER_TESTS_CASES.stream()
                .map(testCase -> arguments(
                        testCase.getName(), testCase.getModules(), testCase.getExpectedErrors(), testCase.isSkip()));
    }
}
