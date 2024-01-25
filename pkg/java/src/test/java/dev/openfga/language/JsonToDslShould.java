package dev.openfga.language;

import org.junit.jupiter.api.Assumptions;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;
import dev.openfga.language.util.TestsData;

import java.util.stream.Stream;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.catchThrowable;
import static org.junit.jupiter.params.provider.Arguments.arguments;

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
    public void throwAnExceptionWhenTransformingInvalidJsonToDsl(String name, String json, String errorMessage, boolean skip) {
        Assumptions.assumeFalse(skip);

        var thrown = catchThrowable(() -> new JsonToDslTransformer().transform(json));

        if(errorMessage == null) {
            assertThat(thrown).isNull();
        } else {
            assertThat(thrown).hasMessage(errorMessage);
        }
    }

    private static Stream<Arguments> transformerTestCases() {
        return TestsData.VALID_TRANSFORMER_TEST_CASES.stream().map(
                testCase -> arguments(
                        testCase.getName(),
                        testCase.getDsl(),
                        testCase.getJson(),
                        testCase.isSkip())
        );
    }

    private static Stream<Arguments> invalidJsonSyntaxTestCases() {
        return TestsData.JSON_SYNTAX_TEST_CASES.stream()
                .map(testCase -> arguments(
                        testCase.getName(),
                        testCase.getJson(),
                        testCase.getErrorMessage(),
                        testCase.isSkip())
                );
    }
}