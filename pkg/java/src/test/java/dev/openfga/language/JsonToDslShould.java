package dev.openfga.language;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.catchThrowable;
import static org.junit.jupiter.params.provider.Arguments.arguments;

import dev.openfga.language.util.ModuleTransformerTestCase;
import dev.openfga.language.util.TestsData;
import dev.openfga.sdk.api.model.AuthorizationModel;
import java.util.stream.Stream;
import org.junit.jupiter.api.Assumptions;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

public class JsonToDslShould {

    @ParameterizedTest(name = "{0}")
    @MethodSource("transformerTestCases")
    public void transfomJson(String name, String dsl, String json, boolean skip) throws Exception {
        Assumptions.assumeFalse(skip);

        var generatedDsl = new JsonToDslTransformer().transform(json);

        assertThat(generatedDsl).isEqualTo(dsl);
    }

    @ParameterizedTest(name = "{0}")
    @MethodSource("invalidJsonSyntaxTestCases")
    public void throwAnExceptionWhenTransformingInvalidJsonToDsl(
            String name, String json, String errorMessage, boolean skip) {
        Assumptions.assumeFalse(skip);

        var thrown = catchThrowable(() -> new JsonToDslTransformer().transform(json));

        if (errorMessage == null) {
            assertThat(thrown).isNull();
        } else {
            assertThat(thrown).hasMessage(errorMessage);
        }
    }

    @ParameterizedTest(name = "{0}")
    @MethodSource("moduleTestCases")
    public void transformModularJsonToDsl(String name, ModuleTransformerTestCase testCase) throws Exception {
        Assumptions.assumeFalse(testCase.isSkip());
        Assumptions.assumeTrue(testCase.getDsl() != null && testCase.getModules() != null);

        var transformer = new JsonToDslTransformer();
        assertThat(transformer.transform(testCase.getJson())).isEqualTo(testCase.getDsl());
        assertThat(transformer.transform(testCase.getJson(), true)).isEqualTo(testCase.getDslWithSourceInfo());
    }

    @ParameterizedTest(name = "{0}")
    @MethodSource("moduleTestCases")
    public void extractModulesFromJson(String name, ModuleTransformerTestCase testCase) throws Exception {
        Assumptions.assumeFalse(testCase.isSkip());
        Assumptions.assumeTrue(testCase.getExpectedModules() != null);

        var model = JSON.parse(testCase.getJson(), AuthorizationModel.class);
        assertThat(JsonToDslTransformer.getModulesFromJSON(model)).isEqualTo(testCase.getExpectedModules());
    }

    private static Stream<Arguments> transformerTestCases() {
        return TestsData.VALID_TRANSFORMER_TEST_CASES.stream()
                .map(testCase ->
                        arguments(testCase.getName(), testCase.getDsl(), testCase.getJson(), testCase.isSkip()));
    }

    private static Stream<Arguments> invalidJsonSyntaxTestCases() {
        return TestsData.JSON_SYNTAX_TEST_CASES.stream()
                .map(testCase -> arguments(
                        testCase.getName(), testCase.getJson(), testCase.getErrorMessage(), testCase.isSkip()));
    }

    private static Stream<Arguments> moduleTestCases() {
        return TestsData.MODULE_TRANSFORMER_TEST_CASES.stream()
                .map(testCase -> arguments(testCase.getName(), testCase));
    }
}
