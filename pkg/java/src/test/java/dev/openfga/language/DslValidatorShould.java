package dev.openfga.language;

import org.junit.jupiter.api.Assumptions;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;
import dev.openfga.language.errors.DslErrorsException;
import dev.openfga.language.errors.ModelValidationSingleError;
import dev.openfga.language.errors.SyntaxError;
import dev.openfga.language.util.TestsData;

import java.util.List;
import java.util.stream.Stream;

import static java.util.stream.Collectors.joining;
import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.catchThrowable;
import static org.junit.jupiter.params.provider.Arguments.arguments;

public class DslValidatorShould {

    @ParameterizedTest(name = "{0}")
    @MethodSource("dslSyntaxTestCases")
    public void verifyDslSyntax(String name, String dsl, List<SyntaxError> expectedErrors, boolean skip) {
        Assumptions.assumeFalse(skip);

        var thrown = catchThrowable(() -> new DslValidator().validate(dsl));

        if (expectedErrors.isEmpty()) {
            assertThat(thrown).isNull();
            return;
        }

        assertThat(thrown).isInstanceOf(DslErrorsException.class);

        var errorsCount = expectedErrors.size();

        var formattedErrors = expectedErrors.stream()
                .map(error -> String.format("syntax error at line=%d, column=%d: %s", error.getLine().getStart(), error.getColumn().getStart(), error.getMessage()))
                .collect(joining("\n\t* "));

        var expectedMessage = String.format("%d error%s occurred:\n\t* %s\n\n",
                errorsCount,
                errorsCount > 1 ? "s" : "",
                formattedErrors);

        assertThat(thrown).hasMessage(expectedMessage);

        var actualErrors = ((DslErrorsException) thrown).getErrors();
        for (int i = 0; i < expectedErrors.size(); i++) {
            var expectedError = expectedErrors.get(i);
            var actualError = actualErrors.get(i);

            assertMatch(expectedError, (SyntaxError) actualError);
        }
    }


    @ParameterizedTest(name = "{0}")
    @MethodSource("dslValidationTestCases")
    public void verifyDslValidation(String name, String dsl, List<ModelValidationSingleError> expectedErrors, boolean skip) {
        Assumptions.assumeFalse(skip);

        var thrown = catchThrowable(() -> new DslValidator().validate(dsl));

        if (expectedErrors.isEmpty()) {
            assertThat(thrown).isNull();
            return;
        }

        assertThat(thrown).isInstanceOf(DslErrorsException.class);

        var errorsCount = expectedErrors.size();

        var formattedErrors = expectedErrors.stream()
                .map(error -> String.format("syntax error at line=%d, column=%d: %s", error.getLine().getStart(), error.getColumn().getStart(), error.getMessage()))
                .collect(joining("\n\t* "));

        var expectedMessage = String.format("%d error%s occurred:\n\t* %s\n\n",
                errorsCount,
                errorsCount > 1 ? "s" : "",
                formattedErrors);

        assertThat(thrown).hasMessage(expectedMessage);

        var actualErrors = ((DslErrorsException) thrown).getErrors();
        for (int i = 0; i < expectedErrors.size(); i++) {
            var expectedError = expectedErrors.get(i);
            var actualError = actualErrors.get(i);

            assertMatch(expectedError, (ModelValidationSingleError) actualError);
        }
    }

    private void assertMatch(SyntaxError expectedError, SyntaxError actualError) {
        assertThat(actualError.getMessage()).isEqualTo(expectedError.getMessage());
        assertThat(actualError.getLine()).isEqualTo(expectedError.getLine());
        assertThat(actualError.getColumn()).isEqualTo(expectedError.getColumn());
        assertThat(actualError.getMetadata()).isEqualTo(expectedError.getMetadata());
    }

    private void assertMatch(ModelValidationSingleError expectedError, ModelValidationSingleError actualError) {
        assertThat(actualError.getMessage()).isEqualTo(expectedError.getMessage());
        assertThat(actualError.getLine()).isEqualTo(expectedError.getLine());
        assertThat(actualError.getColumn()).isEqualTo(expectedError.getColumn());
        assertThat(actualError.getMetadata()).isEqualTo(expectedError.getMetadata());
    }

    private static Stream<Arguments> dslSyntaxTestCases() {
        return TestsData.DSL_SYNTAX_TEST_CASES.stream().map(
                testCase -> arguments(
                        testCase.getName(),
                        testCase.getDsl(),
                        testCase.getExpectedErrors(),
                        testCase.isSkip())
        );
    }

    private static Stream<Arguments> dslValidationTestCases() {
        return TestsData.DSL_VALIDATION_TEST_CASES.stream().map(
                testCase -> arguments(
                        testCase.getName(),
                        testCase.getDsl(),
                        testCase.getExpectedErrors(),
                        testCase.isSkip())
        );
    }
}
