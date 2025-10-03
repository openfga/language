using System.IO;

namespace OpenFga.Language.Tests.util;

public static class TestsData
{
    public const string TransformerCasesFolder = "../../tests/data/transformer";
    public const string DslSyntaxCasesFile = "../../tests/data/dsl-syntax-validation-cases.yaml";
    public const string DslSemanticCasesFile = "../../tests/data/dsl-semantic-validation-cases.yaml";
    public const string JsonSyntaxTransformerCasesFile = "../../tests/data/json-syntax-transformer-validation-cases.yaml";
    public const string FgaModCasesFile = "../../tests/data/fga-mod-transformer-cases.yaml";
    public const string JsonValidationCasesFile = "../../tests/data/json-validation-cases.yaml";
    public const string SkipFile = "test.skip";
    public const string AuthorizationModelJsonFile = "authorization-model.json";
    public const string AuthorizationModelDslFile = "authorization-model.fga";

    public static readonly List<ValidTransformerTestCase> ValidTransformerTestCases = LoadValidTransformerTestCases();
    public static readonly List<DslSyntaxTestCase> DslSyntaxTestCases = LoadDslSyntaxTestCases();
    public static readonly List<MultipleInvalidDslSyntaxTestCase> DslValidationTestCases = LoadDslValidationTestCases();
    public static readonly List<JsonSyntaxTestCase> JsonSyntaxTestCases = LoadJsonSyntaxTestCases();
    public static readonly List<FgaModTestCase> FgaModTransformTestCases = LoadFgaModTransformTestCases();
    public static readonly List<JsonValidationTestCase> JsonValidationTestCases = LoadJsonValidationTestCases();

    private static List<ValidTransformerTestCase> LoadValidTransformerTestCases()
    {
        var transformerCasesFolderPath = PathsGet(TransformerCasesFolder);
        var transformerCasesFolder = new DirectoryInfo(transformerCasesFolderPath);
        var cases = new List<ValidTransformerTestCase>();

        if (!transformerCasesFolder.Exists)
        {
            return cases;
        }

        foreach (var directory in transformerCasesFolder.GetDirectories())
        {
            var name = directory.Name;
            var skipFile = Path.Combine(directory.FullName, SkipFile);
            var jsonFile = Path.Combine(directory.FullName, AuthorizationModelJsonFile);
            var dslFile = Path.Combine(directory.FullName, AuthorizationModelDslFile);

            if (File.Exists(jsonFile) && File.Exists(dslFile))
            {
                var dslContent = File.ReadAllText(dslFile);
                var jsonContent = File.ReadAllText(jsonFile);
                var skip = File.Exists(skipFile);

                cases.Add(new ValidTransformerTestCase(name, dslContent, jsonContent, skip));
            }
        }

        return cases;
    }

    private static List<DslSyntaxTestCase> LoadDslSyntaxTestCases()
    {
        var dslSyntaxCasesFile = PathsGet(DslSyntaxCasesFile);
        var yaml = File.ReadAllText(dslSyntaxCasesFile);


        return YamlHelper.ParseList<DslSyntaxTestCase>(yaml);
    }

    private static List<MultipleInvalidDslSyntaxTestCase> LoadDslValidationTestCases()
    {
        var dslSemanticCasesFile = PathsGet(DslSemanticCasesFile);
        var yaml = File.ReadAllText(dslSemanticCasesFile);
        return YamlHelper.ParseList<MultipleInvalidDslSyntaxTestCase>(yaml);
    }

    private static List<JsonValidationTestCase> LoadJsonValidationTestCases()
    {
        var jsonValidationCasesFile = PathsGet(JsonValidationCasesFile);
        var yaml = File.ReadAllText(jsonValidationCasesFile);
        return YamlHelper.ParseList<JsonValidationTestCase>(yaml);
    }

    private static List<JsonSyntaxTestCase> LoadJsonSyntaxTestCases()
    {
        var jsonSyntaxTransformerCasesFile = PathsGet(JsonSyntaxTransformerCasesFile);
        var yaml = File.ReadAllText(jsonSyntaxTransformerCasesFile);
        return YamlHelper.ParseList<JsonSyntaxTestCase>(yaml);
    }

    private static List<FgaModTestCase> LoadFgaModTransformTestCases()
    {
        var fgaModCasesFile = PathsGet(FgaModCasesFile);
        var yaml = File.ReadAllText(fgaModCasesFile);
        return YamlHelper.ParseList<FgaModTestCase>(yaml);
    }

    private static string PathsGet(string relativePath)
    {
        // Get the directory where the test assembly is located
        var testAssemblyLocation = System.Reflection.Assembly.GetExecutingAssembly().Location;
        var testDirectory = Path.GetDirectoryName(testAssemblyLocation);
        var languageRoot = Path.GetFullPath(Path.Combine(testDirectory, "../../../../.."));
        var fullPath = Path.Combine(languageRoot, relativePath);
        return fullPath;
    }
}
