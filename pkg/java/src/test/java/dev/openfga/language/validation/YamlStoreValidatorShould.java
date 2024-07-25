package dev.openfga.language.validation;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.params.provider.Arguments.arguments;

import com.fasterxml.jackson.databind.ObjectMapper;
import dev.openfga.language.util.TestsData;
import java.util.Objects;
import java.util.stream.Stream;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;
import org.snakeyaml.engine.v2.api.Load;
import org.snakeyaml.engine.v2.api.LoadSettings;

public class YamlStoreValidatorShould {

    @ParameterizedTest(name = "{0}")
    @MethodSource("validationTestCases")
    public void testStoreValidation(String name, String store, String errors) throws Exception {
        var validator = new YamlStoreValidator();

        var settings = LoadSettings.builder().setUseMarks(true).build();

        var load = new Load(settings);
        var loader = load.loadFromString(store);
        var mapper = new ObjectMapper();
        var mapped = mapper.writeValueAsString(loader);
        var results = validator.validate(mapped);

        assertNotNull(results);

        if (!Objects.isNull(errors)) {
            assertEquals(mapper.readTree(errors), mapper.readTree(mapper.writeValueAsString(results)));
        }
    }

    private static Stream<Arguments> validationTestCases() {
        return TestsData.STORE_VALIDATION_TEST_CASES.stream()
                .map(testCase -> arguments(testCase.getName(), testCase.getStore(), testCase.getErrors()));
    }
}
