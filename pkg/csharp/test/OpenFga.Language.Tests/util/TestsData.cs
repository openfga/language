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
        var transformerCasesFolder = new DirectoryInfo(TransformerCasesFolder);
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
        var dslSyntaxCasesFile = new FileInfo(DslSyntaxCasesFile);
        if (!dslSyntaxCasesFile.Exists)
        {
            return new List<DslSyntaxTestCase>();
        }

        var yaml = File.ReadAllText(dslSyntaxCasesFile.FullName);
        return YamlHelper.ParseList<DslSyntaxTestCase>(yaml);
    }

    private static List<MultipleInvalidDslSyntaxTestCase> LoadDslValidationTestCases()
    {
        var dslSyntaxCasesFile = new FileInfo(DslSemanticCasesFile);
        if (!dslSyntaxCasesFile.Exists)
        {
            return new List<MultipleInvalidDslSyntaxTestCase>();
        }

        var yaml = File.ReadAllText(dslSyntaxCasesFile.FullName);
        return YamlHelper.ParseList<MultipleInvalidDslSyntaxTestCase>(yaml);
    }

    private static List<JsonValidationTestCase> LoadJsonValidationTestCases()
    {
        var jsonValidationCasesFile = new FileInfo(JsonValidationCasesFile);
        if (!jsonValidationCasesFile.Exists)
        {
            return new List<JsonValidationTestCase>();
        }

        var yaml = File.ReadAllText(jsonValidationCasesFile.FullName);
        return YamlHelper.ParseList<JsonValidationTestCase>(yaml);
    }

    private static List<JsonSyntaxTestCase> LoadJsonSyntaxTestCases()
    {
        var dslSyntaxCasesFile = new FileInfo(JsonSyntaxTransformerCasesFile);
        if (!dslSyntaxCasesFile.Exists)
        {
            return new List<JsonSyntaxTestCase>();
        }

        var yaml = File.ReadAllText(dslSyntaxCasesFile.FullName);
        return YamlHelper.ParseList<JsonSyntaxTestCase>(yaml);
    }

    private static List<FgaModTestCase> LoadFgaModTransformTestCases()
    {
        var fgaModCasesFile = new FileInfo(FgaModCasesFile);
        if (!fgaModCasesFile.Exists)
        {
            return new List<FgaModTestCase>();
        }

        var yaml = File.ReadAllText(fgaModCasesFile.FullName);
        return YamlHelper.ParseList<FgaModTestCase>(yaml);
    }
}
