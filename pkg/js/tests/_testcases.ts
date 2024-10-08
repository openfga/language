import * as path from "path";
import * as fs from "fs";
import * as yaml from "yaml";
import {
  BaseError,
  DSLSyntaxSingleError,
  FGAModFileValidationSingleError,
  ModelValidationSingleError,
} from "../errors";
import { ModuleFile } from "../transformer";
import { ErrorObject } from "ajv";

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

interface InvalidJSONValidationTestCase {
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

interface MultipleInvalidJSONTestCase extends InvalidJSONValidationTestCase {
  expected_errors: ModelValidationSingleError[];
}

interface MultipleInvalidDSLSyntaxTestCase extends InvalidDSLSyntaxTestCase {
  expected_errors: DSLSyntaxSingleError[];
}

interface MultipleInvalidTestCase extends InvalidDSLSyntaxTestCase {
  expected_errors: ModelValidationSingleError[];
}

interface FGAModFileTestCase extends Omit<ValidTestCase, "dsl"> {
  modFile: string;
  expected_errors?: FGAModFileValidationSingleError[];
}

interface ModuleTestCase extends ValidTestCase {
  modules: ModuleFile[];
  dslWithSourceInfo?: string;
  expected_modules?: string[];
  expected_errors?: BaseError[];
}

interface MultipleInvalidStoreTestCase {
  name: string;
  store: string;
  skip?: boolean;
  expected_errors: ErrorObject[];
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
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
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
  return yaml.parse(
    fs.readFileSync(
      path.join(__dirname, "../../../tests", "data", "json-syntax-transformer-validation-cases.yaml"),
      "utf-8",
    ),
  ) as InvalidJSONSyntaxTestCase[];
}

export function loadInvalidJSONSValidationTestCases(): MultipleInvalidJSONTestCase[] {
  return yaml.parse(
    fs.readFileSync(path.join(__dirname, "../../../tests", "data", "json-validation-cases.yaml"), "utf-8"),
  ) as MultipleInvalidJSONTestCase[];
}

export function loadModFileTestCases(): FGAModFileTestCase[] {
  return yaml.parse(
    fs.readFileSync(path.join(__dirname, "../../../tests", "data", "fga-mod-transformer-cases.yaml"), "utf-8"),
  ) as FGAModFileTestCase[];
}

export function loadModuleTestCases(): ModuleTestCase[] {
  const testDataPath = path.join(__dirname, "../../../tests", "data", "transformer-module");
  const entries = fs.readdirSync(testDataPath, { withFileTypes: true });

  const testCases: ModuleTestCase[] = [];

  for (const entry of entries) {
    if (!entry.isDirectory()) {
      continue;
    }

    const testCase: Partial<ModuleTestCase> = {
      name: entry.name,
    };

    try {
      const skipFile = fs.readFileSync(path.join(testDataPath, testCase.name!, "test.skip"));
      if (skipFile) {
        testCase.skip = true;
      }
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
    } catch (e) {
      // do nothing
    }

    const modelPath = path.join(testDataPath, testCase.name!, "authorization-model.json");
    if (fs.existsSync(modelPath)) {
      const jsonData = fs.readFileSync(modelPath);
      testCase.json = jsonData.toString("utf8");
    }

    const expectedModulesPath = path.join(testDataPath, testCase.name!, "expected_modules.json");
    if (fs.existsSync(expectedModulesPath)) {
      const expectedModules = fs.readFileSync(expectedModulesPath);
      testCase.expected_modules = JSON.parse(expectedModules.toString("utf8"));
    }

    const errorsPath = path.join(testDataPath, testCase.name!, "expected_errors.json");
    if (fs.existsSync(errorsPath)) {
      const expectedErrors = fs.readFileSync(errorsPath);
      testCase.expected_errors = JSON.parse(expectedErrors.toString("utf8"));
    }

    const dslPath = path.join(testDataPath, testCase.name!, "combined.fga");
    if (fs.existsSync(dslPath)) {
      const dsl = fs.readFileSync(dslPath);
      testCase.dsl = dsl.toString("utf8");
    }

    const dslWithSourceInfoPath = path.join(testDataPath, testCase.name!, "combined-sourceinfo.fga");
    if (fs.existsSync(dslWithSourceInfoPath)) {
      const dsl = fs.readFileSync(dslWithSourceInfoPath);
      testCase.dslWithSourceInfo = dsl.toString("utf8");
    }

    const modules: ModuleFile[] = [];
    const modulesPath = path.join(testDataPath, testCase.name!, "module");
    if (fs.existsSync(modulesPath)) {
      const files = fs.readdirSync(modulesPath, {
        withFileTypes: true,
        recursive: true,
      });

      for (const file of files) {
        if (!file.isFile() || !file.name || !file.name.endsWith(".fga")) {
          continue;
        }

        modules.push({
          name: file.name,
          contents: fs.readFileSync(path.join(testDataPath, testCase.name!, "module", file.name), "utf8"),
        });
      }

      testCase.modules = modules;
    }

    testCases.push(testCase as ModuleTestCase);
  }

  return testCases;
}

export function loadStoreTestCases(): MultipleInvalidStoreTestCase[] {
  const testDataPath = path.join(__dirname, "../../../tests", "data", "stores");
  const entries = fs.readdirSync(testDataPath, { withFileTypes: true });

  const errors: MultipleInvalidStoreTestCase[] = [];

  for (const entry of entries) {
    if (!entry.isDirectory()) {
      continue;
    }

    const testCase: Partial<MultipleInvalidStoreTestCase> = {
      name: entry.name,
    };

    testCase.store = path.join(testDataPath, testCase.name!, "store.fga.yaml");

    const errorsPath = path.join(testDataPath, testCase.name!, "expected_errors.json");
    if (fs.existsSync(errorsPath)) {
      const expectedErrors = fs.readFileSync(errorsPath);
      testCase.expected_errors = JSON.parse(expectedErrors.toString("utf8"));
    }

    errors.push(testCase as MultipleInvalidStoreTestCase);
  }

  return errors;
}
