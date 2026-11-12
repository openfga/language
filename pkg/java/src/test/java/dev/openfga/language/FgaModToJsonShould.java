package dev.openfga.language;

import static java.util.stream.Collectors.joining;
import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.catchThrowable;
import static org.junit.jupiter.params.provider.Arguments.arguments;

import dev.openfga.language.errors.ModFileValidationError;
import dev.openfga.language.errors.ModFileValidationSingleError;
import dev.openfga.language.util.TestsData;
import java.util.List;
import java.util.stream.Stream;
import org.junit.jupiter.api.Assumptions;
import org.junit.jupiter.api.Timeout;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

public class FgaModToJsonShould {

    @ParameterizedTest(name = "{0}")
    @MethodSource("fgaModTestCases")
    @Timeout(value = 10, threadMode = Timeout.ThreadMode.SEPARATE_THREAD)
    public void transformFgaMod(
            String name, String fgaMod, String json, boolean skip, List<ModFileValidationSingleError> expectedErrors)
            throws Exception, ModFileValidationError {
        Assumptions.assumeFalse(skip);

        if (json != null) {
            var generatedJson = FgaModTransformer.transform(fgaMod);

            var expected = JSON.stringify(JSON.parse(json, FgaModFile.class));

            assertThat(generatedJson).isEqualTo(expected);
        } else {
            var thrown = catchThrowable(() -> FgaModTransformer.transform(fgaMod));

            assertThat(thrown).isInstanceOf(ModFileValidationError.class);

            var errorsCount = expectedErrors.size();

            var formattedErrors = expectedErrors.stream()
                    .map(error -> String.format(
                            "validation error at line=%d, column=%d: %s",
                            error.getLine().getStart(), error.getColumn().getStart(), error.getMessage()))
                    .collect(joining("\n\t* "));

            var expectedMessage = String.format(
                    "%d error%s occurred:\n\t* %s\n\n", errorsCount, errorsCount > 1 ? "s" : "", formattedErrors);

            assertThat(thrown).hasMessage(expectedMessage);
        }
    }

    // Error messages for nested values differ between languages, so these cases are Java-only.
    @ParameterizedTest(name = "{0}")
    @MethodSource("nestedValueTestCases")
    @Timeout(value = 10, threadMode = Timeout.ThreadMode.SEPARATE_THREAD)
    public void skipNestedValuesOfInvalidFields(String name, String fgaMod, List<String> expectedMessages) {
        var thrown = catchThrowable(() -> FgaModTransformer.transform(fgaMod));

        assertThat(thrown).isInstanceOf(ModFileValidationError.class);
        assertThat(((ModFileValidationError) thrown).getErrors())
                .extracting(ModFileValidationSingleError::getMessage)
                .containsExactlyElementsOf(expectedMessages);
    }

    private static Stream<Arguments> nestedValueTestCases() {
        return Stream.of(
                arguments(
                        "schema is an object",
                        "schema:\n  version: \"1.2\"\ncontents:\n  - a.fga\n",
                        List.of("unexpected schema type, expected string got value {")),
                arguments(
                        "contents is an object",
                        "contents:\n  file: a.fga\nschema: \"1.2\"\n",
                        List.of("unexpected contents type, expected list of strings got value {")),
                arguments(
                        "contents items are lists or objects",
                        "schema: \"1.2\"\ncontents:\n  - [b.fga]\n  - {c: d.fga}\n  - e.txt\n",
                        List.of(
                                "unexpected contents item type, expected string got value [",
                                "unexpected contents item type, expected string got value {",
                                "contents items should use fga file extension, got e.txt")));
    }

    private static Stream<Arguments> fgaModTestCases() {
        return TestsData.FGA_MOD_TRANSFORM_TEST_CASES.stream()
                .map(testCase -> arguments(
                        testCase.getName(),
                        testCase.getModFile(),
                        testCase.getJson(),
                        testCase.isSkip(),
                        testCase.getExpectedErrors()));
    }
}
