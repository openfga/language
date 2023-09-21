import * as path from "path";
import * as fs from "fs";
import * as yaml from "yaml";
import { DSLSyntaxSingleError, ModelValidationSingleError } from "../errors";

interface ValidTestCase {
  name: string;
  dsl: string;
  json: string;
  skip?: boolean;
}

interface InvalidJSONSyntaxTestCase {
  name: string;
  json: string;
  error_message: string;
  skip?: boolean;
}

interface InvalidDSLSyntaxTestCase {
  name: string;
  dsl: string;
  error_message: string;
  skip?: boolean;
}

interface MultipleInvalidDSLSyntaxTestCase extends InvalidDSLSyntaxTestCase {
  expected_errors: DSLSyntaxSingleError[];
}

interface MultipleInvalidTestCase extends InvalidDSLSyntaxTestCase {
  expected_errors: ModelValidationSingleError[];
}

export function loadValidTransformerTestCases(): ValidTestCase[] {
  const testDataPath = path.join(__dirname, "../../../tests", "data", "transformer");
  const entries = fs.readdirSync(testDataPath, { withFileTypes: true });

  const testCases: ValidTestCase[] = [];

  for (const entry of entries) {
    if (!entry.isDirectory()) {
      continue;
    }

    const testCase: Partial<ValidTestCase> = {
      name: entry.name,
    };

    try {
      const skipFile = fs.readFileSync(path.join(testDataPath, testCase.name!, "test.skip"));
      if (skipFile) {
        testCase.skip = true;
      }
    } catch (e) {
      // do nothing
    }

    const jsonData = fs.readFileSync(path.join(testDataPath, testCase.name!, "authorization-model.json"));
    testCase.json = jsonData.toString("utf8");

    const dslData = fs.readFileSync(path.join(testDataPath, testCase.name!, "authorization-model.fga"));
    testCase.dsl = dslData.toString("utf8");

    testCases.push(testCase as ValidTestCase);
  }

  return testCases;
}

export function loadDSLSyntaxErrorTestCases(): MultipleInvalidDSLSyntaxTestCase[] {
  return yaml.parse(
    fs.readFileSync(path.join(__dirname, "../../../tests", "data", "dsl-syntax-validation-cases.yaml"), "utf-8"),
  ) as MultipleInvalidDSLSyntaxTestCase[];
}

export function loadDSLValidationErrorTestCases(): MultipleInvalidTestCase[] {
  return yaml.parse(
    fs.readFileSync(path.join(__dirname, "../../../tests", "data", "dsl-semantic-validation-cases.yaml"), "utf-8"),
  ) as MultipleInvalidTestCase[];
}

export function loadInvalidJSONSyntaxTestCases(): InvalidJSONSyntaxTestCase[] {
  const docs = yaml.parseAllDocuments(
    fs.readFileSync(
      path.join(__dirname, "../../../tests", "data", "json-syntax-transformer-validation-cases.yaml"),
      "utf-8",
    ),
  );

  return docs.map((d) => d.toJSON()) as InvalidJSONSyntaxTestCase[];
}
