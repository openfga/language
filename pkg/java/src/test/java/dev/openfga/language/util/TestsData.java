package dev.openfga.language.util;

import static java.util.Collections.unmodifiableList;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import dev.openfga.language.ModuleFile;
import java.io.IOException;
import java.nio.file.DirectoryStream;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class TestsData {

    public static final String TRANSFORMER_CASES_FOLDER = "../../tests/data/transformer";
    public static final String DSL_SYNTAX_CASES_FILE = "../../tests/data/dsl-syntax-validation-cases.yaml";
    public static final String DSL_SEMANTIC_CASES_FILE = "../../tests/data/dsl-semantic-validation-cases.yaml";
    public static final String JSON_SYNTAX_TRANSFORMER_CASES_FILE =
            "../../tests/data/json-syntax-transformer-validation-cases.yaml";
    public static final String FGA_MOD_CASES_FILE = "../../tests/data/fga-mod-transformer-cases.yaml";
    public static final String JSON_VALIDATION_CASES_FILE = "../../tests/data/json-validation-cases.yaml";
    public static final String TRANSFORMER_MODULE_CASES_FOLDER = "../../tests/data/transformer-module";
    public static final String SKIP_FILE = "test.skip";
    public static final String AUTHORIZATION_MODEL_JSON_FILE = "authorization-model.json";
    public static final String AUTHORIZATION_MODEL_DSL_FILE = "authorization-model.fga";

    public static final List<ValidTransformerTestCase> VALID_TRANSFORMER_TEST_CASES = loadValidTransformerTestCases();
    public static final List<DslSyntaxTestCase> DSL_SYNTAX_TEST_CASES = loadDslSyntaxTestCases();
    public static final List<MultipleInvalidDslSyntaxTestCase> DSL_VALIDATION_TEST_CASES = loadDslValidationTestCases();
    public static final List<JsonSyntaxTestCase> JSON_SYNTAX_TEST_CASES = loadJsonSyntaxTestCases();
    public static final List<FgaModTestCase> FGA_MOD_TRANSFORM_TEST_CASES = loadFgaModTransformTestCases();
    public static final List<JsonValidationTestCase> JSON_VALIDATION_TEST_CASES = loadJsonValidationTestCases();
    public static final List<ModuleTransformerTestCase> MODULE_TRANSFORMER_TEST_CASES =
            loadModuleTransformerTestCases();

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

    private static List<ModuleTransformerTestCase> loadModuleTransformerTestCases() {
        var mapper = new ObjectMapper();
        var casesFolder = Paths.get(TRANSFORMER_MODULE_CASES_FOLDER);

        List<ModuleTransformerTestCase> cases = new ArrayList<>();
        try (DirectoryStream<Path> stream = Files.newDirectoryStream(casesFolder)) {
            for (Path path : stream) {
                if (!Files.isDirectory(path)) {
                    continue;
                }

                var name = path.getFileName().toString();
                var skip = Files.exists(path.resolve(SKIP_FILE));

                var json = readIfExists(path.resolve(AUTHORIZATION_MODEL_JSON_FILE));
                var dsl = readIfExists(path.resolve("combined.fga"));
                var dslWithSourceInfo = readIfExists(path.resolve("combined-sourceinfo.fga"));

                List<String> expectedModules = null;
                var modulesJson = readIfExists(path.resolve("expected_modules.json"));
                if (modulesJson != null) {
                    expectedModules = mapper.readValue(modulesJson, new TypeReference<>() {});
                }

                List<ModuleExpectedError> expectedErrors = null;
                var errorsJson = readIfExists(path.resolve("expected_errors.json"));
                if (errorsJson != null) {
                    List<ModuleExpectedError> allErrors = mapper.readValue(errorsJson, new TypeReference<>() {});
                    // The Go-equivalent transformer surfaces transformation and syntax errors only; model-validation
                    // errors are filtered out. A case left with no errors is skipped.
                    expectedErrors = allErrors.stream()
                            .filter(error -> !error.isValidationError())
                            .collect(Collectors.toList());
                    if (expectedErrors.isEmpty()) {
                        skip = true;
                    }
                }

                cases.add(new ModuleTransformerTestCase(
                        name,
                        loadModuleFiles(path.resolve("module")),
                        json,
                        dsl,
                        dslWithSourceInfo,
                        expectedModules,
                        expectedErrors,
                        skip));
            }
        } catch (IOException e) {
            throw new RuntimeException(e);
        }

        return unmodifiableList(cases);
    }

    private static List<ModuleFile> loadModuleFiles(Path moduleDir) throws IOException {
        if (!Files.isDirectory(moduleDir)) {
            return null;
        }

        try (Stream<Path> files = Files.walk(moduleDir)) {
            return files.filter(Files::isRegularFile)
                    .filter(file -> file.getFileName().toString().endsWith(".fga"))
                    .sorted(Comparator.comparing(file -> file.getFileName().toString()))
                    .map(file -> {
                        try {
                            return new ModuleFile(file.getFileName().toString(), Files.readString(file));
                        } catch (IOException e) {
                            throw new RuntimeException(e);
                        }
                    })
                    .collect(Collectors.toList());
        }
    }

    private static String readIfExists(Path path) throws IOException {
        return Files.exists(path) ? Files.readString(path) : null;
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
