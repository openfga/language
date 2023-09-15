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

interface InvalidJSONSyntaxTestCase {
  name: string;
  json: string;
  error_message: string;
  skip?: boolean;
}

export function loadInvalidJsonSyntaxTestCases(): InvalidJSONSyntaxTestCase[] {
  const docs = yaml.parseAllDocuments(
    fs.readFileSync(
      path.join(__dirname, "../../../tests", "data", "json-syntax-transformer-validation-cases.yaml"),
      "utf-8",
    ),
  );

  return docs.map((d) => d.toJSON()) as InvalidJSONSyntaxTestCase[];
}

interface InvalidDslSyntaxTestCase {
  name: string;
  dsl: string;
  valid: boolean;
  error_message: string;
  skip?: boolean;
}

export function loadInvalidDslSyntaxTestCases(): InvalidDslSyntaxTestCase[] {
  const jsonData = fs.readFileSync(path.join(__dirname, "../../../tests", "data", "dsl-syntax-validation-cases.json"));

  return JSON.parse(jsonData.toString("utf8")) as InvalidDslSyntaxTestCase[];
}

interface MultipleInvalidDslSyntaxTestCase extends InvalidDslSyntaxTestCase {
  expected_errors: DSLSyntaxSingleError[];
}

export function loadDslSyntaxErrorTestCases(): MultipleInvalidDslSyntaxTestCase[] {
  return yaml.parse(
    fs.readFileSync(path.join(__dirname, "../../../tests", "data", "dsl-syntax-validation-cases.yaml"), "utf-8"),
  ) as MultipleInvalidDslSyntaxTestCase[];
}

interface MultipleInvalidTestCase extends InvalidDslSyntaxTestCase {
  expected_errors: ModelValidationSingleError[];
}

export function loadDslValidationErrorTestCases(): MultipleInvalidTestCase[] {
  return yaml.parse(
    fs.readFileSync(path.join(__dirname, "../../../tests", "data", "dsl-semantic-validation-cases.yaml"), "utf-8"),
  ) as MultipleInvalidTestCase[];
}
