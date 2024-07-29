package dev.openfga.language.util;

import static java.util.Collections.unmodifiableList;

import com.fasterxml.jackson.core.type.TypeReference;
import java.io.IOException;
import java.nio.file.DirectoryStream;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;

public class TestsData {

    public static final String TRANSFORMER_CASES_FOLDER = "../../tests/data/transformer";
    public static final String DSL_SYNTAX_CASES_FILE = "../../tests/data/dsl-syntax-validation-cases.yaml";
    public static final String DSL_SEMANTIC_CASES_FILE = "../../tests/data/dsl-semantic-validation-cases.yaml";
    public static final String JSON_SYNTAX_TRANSFORMER_CASES_FILE =
            "../../tests/data/json-syntax-transformer-validation-cases.yaml";
    public static final String FGA_MOD_CASES_FILE = "../../tests/data/fga-mod-transformer-cases.yaml";
    public static final String JSON_VALIDATION_CASES_FILE = "../../tests/data/json-validation-cases.yaml";
    public static final String SKIP_FILE = "test.skip";
    public static final String AUTHORIZATION_MODEL_JSON_FILE = "authorization-model.json";
    public static final String AUTHORIZATION_MODEL_DSL_FILE = "authorization-model.fga";

    public static final List<ValidTransformerTestCase> VALID_TRANSFORMER_TEST_CASES = loadValidTransformerTestCases();
    public static final List<DslSyntaxTestCase> DSL_SYNTAX_TEST_CASES = loadDslSyntaxTestCases();
    public static final List<MultipleInvalidDslSyntaxTestCase> DSL_VALIDATION_TEST_CASES = loadDslValidationTestCases();
    public static final List<JsonSyntaxTestCase> JSON_SYNTAX_TEST_CASES = loadJsonSyntaxTestCases();
    public static final List<FgaModTestCase> FGA_MOD_TRANSFORM_TEST_CASES = loadFgaModTransformTestCases();
    public static final List<JsonValidationTestCase> JSON_VALIDATION_TEST_CASES = loadJsonValidationTestCases();

    public static final String STORE_VALIDATION_YAML_FILE = "store.fga.yaml";
    public static final String STORE_VALIDATION_ERRORS_JSON_FILE = "java_expected_errors.json";
    public static final String STORE_VALIDATION_CASES_FOLDER = "../../tests/data/stores";
    public static final List<ValidateStoreTestCase> STORE_VALIDATION_TEST_CASES = loadStoreValidationTestCases();

    private static List<ValidateStoreTestCase> loadStoreValidationTestCases() {
        var storeValidationCasesFolder = Paths.get(STORE_VALIDATION_CASES_FOLDER);

        List<ValidateStoreTestCase> cases = new ArrayList<>();
        try (DirectoryStream<Path> stream = Files.newDirectoryStream(storeValidationCasesFolder)) {
            for (Path path : stream) {
                if (!Files.isDirectory(path)) {
                    continue;
                }

                var name = path.getFileName().toString();
                var storeFile = path.resolve(STORE_VALIDATION_YAML_FILE);
                var errorFile = path.resolve(STORE_VALIDATION_ERRORS_JSON_FILE);

                if (errorFile.toFile().exists()) {
                    cases.add(
                            new ValidateStoreTestCase(name, Files.readString(storeFile), Files.readString(errorFile)));
                } else {
                    cases.add(new ValidateStoreTestCase(name, Files.readString(storeFile), null));
                }
            }
        } catch (IOException e) {
            throw new RuntimeException(e);
        }

        return unmodifiableList(cases);
    }

    private static List<ValidTransformerTestCase> loadValidTransformerTestCases() {
        var transformerCasesFolder = Paths.get(TRANSFORMER_CASES_FOLDER);

        List<ValidTransformerTestCase> cases = new ArrayList<>();
        try (DirectoryStream<Path> stream = Files.newDirectoryStream(transformerCasesFolder)) {
            for (Path path : stream) {
                if (!Files.isDirectory(path)) {
                    continue;
                }

                var name = path.getFileName().toString();
                var skipFile = path.resolve(SKIP_FILE);
                var jsonFile = path.resolve(AUTHORIZATION_MODEL_JSON_FILE);
                var dslFile = path.resolve(AUTHORIZATION_MODEL_DSL_FILE);

                cases.add(new ValidTransformerTestCase(
                        name, Files.readString(dslFile), Files.readString(jsonFile), Files.exists(skipFile)));
            }
        } catch (IOException e) {
            throw new RuntimeException(e);
        }

        return unmodifiableList(cases);
    }

    private static List<DslSyntaxTestCase> loadDslSyntaxTestCases() {
        var dslSyntaxCasesFile = Paths.get(DSL_SYNTAX_CASES_FILE);
        try {
            var json = Files.readString(dslSyntaxCasesFile);
            return YAML.parseList(json, new TypeReference<>() {});
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    private static List<MultipleInvalidDslSyntaxTestCase> loadDslValidationTestCases() {
        var dslSyntaxCasesFile = Paths.get(DSL_SEMANTIC_CASES_FILE);
        try {
            var json = Files.readString(dslSyntaxCasesFile);
            return YAML.parseList(json, new TypeReference<>() {});
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    private static List<JsonValidationTestCase> loadJsonValidationTestCases() {
        var jsonValidationCasesFile = Paths.get(JSON_VALIDATION_CASES_FILE);
        try {
            var yaml = Files.readString(jsonValidationCasesFile);
            return YAML.parseList(yaml, new TypeReference<>() {});
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    private static List<JsonSyntaxTestCase> loadJsonSyntaxTestCases() {
        var dslSyntaxCasesFile = Paths.get(JSON_SYNTAX_TRANSFORMER_CASES_FILE);
        try {
            var json = Files.readString(dslSyntaxCasesFile);
            return YAML.parseList(json, new TypeReference<>() {});
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    private static List<FgaModTestCase> loadFgaModTransformTestCases() {
        var fgaModCasesFile = Paths.get(FGA_MOD_CASES_FILE);
        try {
            var yaml = Files.readString(fgaModCasesFile);
            return YAML.parseList(yaml, new TypeReference<>() {});
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
