package dev.openfga.language;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.catchThrowable;
import static org.junit.jupiter.params.provider.Arguments.arguments;

import dev.openfga.language.errors.ModuleTransformationError;
import dev.openfga.language.errors.ParsingError;
import dev.openfga.language.util.ModuleExpectedError;
import dev.openfga.language.util.ModuleTransformerTestCase;
import dev.openfga.language.util.TestsData;
import java.util.List;
import java.util.stream.Stream;
import org.junit.jupiter.api.Assumptions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

public class ModuleToModelShould {

    @ParameterizedTest(name = "{0}")
    @MethodSource("moduleTestCases")
    public void transformModulesToModel(String name, ModuleTransformerTestCase testCase) throws Exception {
        Assumptions.assumeFalse(testCase.isSkip());
        Assumptions.assumeTrue(testCase.getModules() != null);

        if (testCase.getExpectedErrors() == null) {
            var expected =
                    JSON.stringify(JSON.parse(testCase.getJson(), dev.openfga.sdk.api.model.AuthorizationModel.class));
            var actual = ModulesToModelTransformer.transform(testCase.getModules(), "1.2");
            assertThat(actual).isEqualTo(expected);
            return;
        }

        var thrown = catchThrowable(() -> ModulesToModelTransformer.transformToModel(testCase.getModules(), "1.2"));

        assertThat(thrown).isInstanceOf(ModuleTransformationError.class);
        var error = (ModuleTransformationError) thrown;

        List<? extends ParsingError> actualErrors = error.getErrors();
        var expectedErrors = testCase.getExpectedErrors();
        assertThat(actualErrors).hasSameSizeAs(expectedErrors);

        var hasSyntaxError = expectedErrors.stream().anyMatch(e -> "syntax".equals(e.getType()));
        if (!hasSyntaxError) {
            // ANTLR emits differing syntax-error text across Java/Go/JS, so only assert the full message when
            // every expected error is a transformation error.
            assertThat(thrown).hasMessage(expectedMessage(expectedErrors));
        }

        for (int i = 0; i < actualErrors.size(); i++) {
            var expected = expectedErrors.get(i);
            var actual = actualErrors.get(i);
            assertThat(actual.getFile()).isEqualTo(expected.getFile());
            if (!"syntax".equals(expected.getType())) {
                assertThat(actual.getMessage()).isEqualTo(expected.getMessage());
                assertThat(actual.getLine()).isEqualTo(expected.getLine());
                assertThat(actual.getColumn()).isEqualTo(expected.getColumn());
            }
        }
    }

    @Test
    public void allowCustomSchemaVersion() throws Exception {
        var model = ModulesToModelTransformer.transformToModel(
                List.of(new ModuleFile("core.fga", "module core\n  type user")), "1.1");
        assertThat(model.getSchemaVersion()).isEqualTo("1.1");
    }

    private static String expectedMessage(List<ModuleExpectedError> errors) {
        var builder = new StringBuilder()
                .append(errors.size())
                .append(" error")
                .append(errors.size() == 1 ? "" : "s")
                .append(" occurred:");
        for (var error : errors) {
            var type = error.getType() != null ? error.getType() : "transformation";
            builder.append("\n\t* ")
                    .append(type)
                    .append(" error at line=")
                    .append(error.getLine().getStart())
                    .append(", column=")
                    .append(error.getColumn().getStart())
                    .append(": ")
                    .append(error.getMessage());
        }
        return builder.append("\n\n").toString();
    }

    private static Stream<Arguments> moduleTestCases() {
        return TestsData.MODULE_TRANSFORMER_TEST_CASES.stream()
                .map(testCase -> arguments(testCase.getName(), testCase));
    }
}
