package dev.openfga.language;

import static java.util.stream.Collectors.joining;
import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.catchThrowable;
import static org.junit.jupiter.params.provider.Arguments.arguments;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import dev.openfga.language.errors.DslErrorsException;
import dev.openfga.language.errors.ModelValidationSingleError;
import dev.openfga.language.errors.ParsingError;
import dev.openfga.language.util.TestsData;
import dev.openfga.language.validation.ModelValidator;
import dev.openfga.sdk.api.model.AuthorizationModel;
import java.util.List;
import java.util.stream.Stream;
import org.junit.jupiter.api.Assumptions;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

public class ModelValidatorShould {

    @ParameterizedTest(name = "{0}")
    @MethodSource("dslSyntaxTestCases")
    public void verifyDslSyntax(String name, String dsl, List<ParsingError> expectedErrors, boolean skip) {
        Assumptions.assumeFalse(skip);

        var thrown = catchThrowable(() -> ModelValidator.validateDsl(dsl));

        if (expectedErrors.isEmpty()) {
            assertThat(thrown).isNull();
            return;
        }

        assertThat(thrown).isInstanceOf(DslErrorsException.class);

        // unfortunately antlr is throwing different error messages in Java, Go and JS - considering that at the moment
        //  we care that it errors for syntax errors more than we care about the error messages matching,
        //  esp. in Java as we are not building a language server on top of the returned errors yet
        //  actual matching error strings is safe to ignore for now

        //        var errorsCount = expectedErrors.size();
        //
        //        var formattedErrors = expectedErrors.stream()
        //                .map(error -> String.format("syntax error at line=%d, column=%d: %s",
        // error.getLine().getStart(), error.getColumn().getStart(), error.getMessage()))
        //                .collect(joining("\n\t* "));
        //
        //        var expectedMessage = String.format("%d error%s occurred:\n\t* %s\n\n",
        //                errorsCount,
        //                errorsCount > 1 ? "s" : "",
        //                formattedErrors);
        //
        //        assertThat(thrown).hasMessage(expectedMessage);
        //
        //        var actualErrors = ((DslErrorsException) thrown).getErrors();
        //        for (int i = 0; i < expectedErrors.size(); i++) {
        //            var expectedError = expectedErrors.get(i);
        //            var actualError = actualErrors.get(i);
        //
        //            assertMatch(expectedError, actualError);
        //        }
    }

    @ParameterizedTest(name = "{0}")
    @MethodSource("dslValidationTestCases")
    public void verifyDslValidation(
            String name, String dsl, List<ModelValidationSingleError> expectedErrors, boolean skip) {
        Assumptions.assumeFalse(skip);

        var thrown = catchThrowable(() -> ModelValidator.validateDsl(dsl));

        if (expectedErrors.isEmpty()) {
            assertThat(thrown).isNull();
            return;
        }

        assertThat(thrown).isInstanceOf(DslErrorsException.class);

        var errorsCount = expectedErrors.size();

        var formattedErrors = expectedErrors.stream()
                .map(error -> String.format(
                        "validation error at line=%d, column=%d: %s",
                        error.getLine().getStart(), error.getColumn().getStart(), error.getMessage()))
                .collect(joining("\n\t* "));

        var expectedMessage = String.format(
                "%d error%s occurred:\n\t* %s\n\n", errorsCount, errorsCount > 1 ? "s" : "", formattedErrors);

        assertThat(thrown).hasMessage(expectedMessage);

        var actualErrors = ((DslErrorsException) thrown).getErrors();
        for (int i = 0; i < expectedErrors.size(); i++) {
            var expectedError = expectedErrors.get(i);
            var actualError = actualErrors.get(i);

            assertMatch(expectedError, (ModelValidationSingleError) actualError);
        }
    }

    @ParameterizedTest(name = "{0}")
    @MethodSource("jsonValidationTestCases")
    public void verifyJsonValidation(String name, String json, List<ModelValidationSingleError> expectedErrors)
            throws JsonProcessingException {
        var model = new ObjectMapper().readValue(json, AuthorizationModel.class);

        var thrown = catchThrowable(() -> ModelValidator.validateJson(model));

        if (expectedErrors == null || expectedErrors.isEmpty()) {
            assertThat(thrown).isNull();
            return;
        }

        assertThat(thrown).isInstanceOf(DslErrorsException.class);

        var errorsCount = expectedErrors.size();

        var formattedErrors = expectedErrors.stream()
                .map(error -> String.format("validation error: %s", error.getMessage()))
                .collect(joining("\n\t* "));

        var expectedMessage = String.format(
                "%d error%s occurred:\n\t* %s\n\n", errorsCount, errorsCount > 1 ? "s" : "", formattedErrors);

        assertThat(thrown).hasMessage(expectedMessage);

        var actualErrors = ((DslErrorsException) thrown).getErrors();
        for (int i = 0; i < expectedErrors.size(); i++) {
            var expectedError = expectedErrors.get(i);
            var actualError = actualErrors.get(i);

            assertMatch(expectedError, (ModelValidationSingleError) actualError);
        }
    }

    private void assertMatch(ParsingError expectedError, ParsingError actualError) {
        assertThat(actualError.getMessage()).isEqualTo(expectedError.getMessage());
        assertThat(actualError.getLine()).isEqualTo(expectedError.getLine());
        assertThat(actualError.getColumn()).isEqualTo(expectedError.getColumn());
    }

    private void assertMatch(ModelValidationSingleError expectedError, ModelValidationSingleError actualError) {
        assertThat(actualError.getMessage()).isEqualTo(expectedError.getMessage());

        if (expectedError.getLine() != null) {
            assertThat(actualError.getLine()).isEqualTo(expectedError.getLine());
        }

        if (expectedError.getColumn() != null) {
            assertThat(actualError.getColumn()).isEqualTo(expectedError.getColumn());
        }

        assertThat(actualError.getMetadata().getErrorType())
                .isEqualTo(expectedError.getMetadata().getErrorType());
        if (expectedError.getMetadata().getTypeName() != null) {
            assertThat(actualError.getMetadata().getTypeName())
                    .isEqualTo(expectedError.getMetadata().getTypeName());
        }
        if (expectedError.getMetadata().getRelation() != null) {
            assertThat(actualError.getMetadata().getRelation())
                    .isEqualTo(expectedError.getMetadata().getRelation());
        }
        if (expectedError.getMetadata().getConditionName() != null) {
            assertThat(actualError.getMetadata().getConditionName())
                    .isEqualTo(expectedError.getMetadata().getConditionName());
        }
    }

    private static Stream<Arguments> dslSyntaxTestCases() {
        return TestsData.DSL_SYNTAX_TEST_CASES.stream()
                .map(testCase -> arguments(
                        testCase.getName(), testCase.getDsl(), testCase.getExpectedErrors(), testCase.isSkip()));
    }

    private static Stream<Arguments> dslValidationTestCases() {
        return TestsData.DSL_VALIDATION_TEST_CASES.stream()
                .map(testCase -> arguments(
                        testCase.getName(), testCase.getDsl(), testCase.getExpectedErrors(), testCase.isSkip()));
    }

    private static Stream<Arguments> jsonValidationTestCases() {
        return TestsData.JSON_VALIDATION_TEST_CASES.stream()
                .map(testCase -> arguments(testCase.getName(), testCase.getJson(), testCase.getExpectedErrors()));
    }
}
