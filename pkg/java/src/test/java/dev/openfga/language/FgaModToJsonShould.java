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
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

public class FgaModToJsonShould {

    @ParameterizedTest(name = "{0}")
    @MethodSource("fgaModTestCases")
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
