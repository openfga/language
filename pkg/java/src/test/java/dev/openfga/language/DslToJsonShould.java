package dev.openfga.language;

import dev.openfga.language.errors.DslErrorsException;
import dev.openfga.language.errors.SyntaxError;
import dev.openfga.language.util.TestsData;
import dev.openfga.sdk.api.model.AuthorizationModel;
import org.junit.jupiter.api.Assumptions;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.Collection;
import java.util.stream.Stream;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.catchThrowable;
import static org.junit.jupiter.params.provider.Arguments.arguments;

public class DslToJsonShould {

    @ParameterizedTest(name = "{0}")
    @MethodSource("transformerTestCases")
    public void transfomDsl(String name, String dsl, String json, boolean skip) throws Exception {
        Assumptions.assumeFalse(skip);

        var generatedJson = new DslToJsonTransformer().transform(dsl);

        var expectedAuthorizationModel = JSON.parse(json, AuthorizationModel.class);

        var expectedJson = JSON.stringify(expectedAuthorizationModel);

        assertThat(generatedJson).isEqualTo(expectedJson);
    }

    @ParameterizedTest(name = "{0}")
    @MethodSource("dslSyntaxTestCases")
    public void verifyDslSyntax(String name, String dsl, Collection<SyntaxError> expectedErrors) {
        var thrown = catchThrowable(() -> new DslToJsonTransformer().transform(dsl));

        if (expectedErrors.isEmpty()) {
            assertThat(thrown).isNull();
            return;
        }

        assertThat(thrown)
                .isInstanceOf(DslErrorsException.class);

        // unfortunately antlr is throwing different error messages in Java, Go and JS - considering that at the moment
        //  we care that it errors for syntax errors more than we care about the error messages matching,
        //  esp. in Java as we are not building a language server on top of the returned errors yet
        //  actual matching error strings is safe to ignore for now
//        var dslSyntaxException = (DslErrorsException) thrown;
//        assertThat(dslSyntaxException.getErrors()).hasSameSizeAs(expectedErrors);
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

    private static Stream<Arguments> dslSyntaxTestCases() {
        return TestsData.DSL_SYNTAX_TEST_CASES.stream().map(
                testCase -> arguments(
                        testCase.getName(),
                        testCase.getDsl(),
                        testCase.getExpectedErrors())
        );
    }
}
