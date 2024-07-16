import { describe, expect, it } from "@jest/globals";
import { loadValidTransformerTestCases, loadInvalidJSONSyntaxTestCases, loadModuleTestCases } from "./_testcases";
import { getModulesFromJSON, transformJSONStringToDSL } from "../transformer";

const testCases = loadValidTransformerTestCases();
const invalidTestCases = loadInvalidJSONSyntaxTestCases();
const moduleTestCases = loadModuleTestCases();

describe("jsonToDSL", () => {
  testCases.forEach((testCase) => {
    const testFn = testCase.skip ? it.skip : it;

    testFn(`should transform ${testCase.name} from JSON to DSL`, () => {
      const dslSyntax = transformJSONStringToDSL(testCase.json);
      expect(dslSyntax).toEqual(testCase.dsl);
    });
  });

  invalidTestCases.forEach((testCase) => {
    const testFn = testCase.skip ? it.skip : it;

    testFn(`should throw an error when transforming ${testCase.name} from JSON to DSL`, () => {
      if (testCase.error_message) {
        expect(() => transformJSONStringToDSL(testCase.json)).toThrow(testCase.error_message);
      } else {
        expect(() => transformJSONStringToDSL(testCase.json)).not.toThrow();
      }
    });
  });

  moduleTestCases.forEach((testCase) => {
    if (!testCase.dsl || !testCase.modules) {
      return;
    }
    const testFn = testCase.skip ? it.skip : it;

    testFn(`should transform ${testCase.name} from JSON to DSL without source info`, () => {
      const dslSyntax = transformJSONStringToDSL(testCase.json);
      expect(dslSyntax).toEqual(testCase.dsl);
    });

    testFn(`should transform ${testCase.name} from JSON to DSL with source info`, () => {
      const dslSyntax = transformJSONStringToDSL(testCase.json, { includeSourceInformation: true });
      expect(dslSyntax).toEqual(testCase.dslWithSourceInfo);
    });
  });
});

describe("getModulesFromJSON", () => {
  moduleTestCases.forEach((testCase) => {
    if (!testCase.expected_modules) {
      return;
    }
    const testFn = testCase.skip ? it.skip : it;

    testFn(`should extract modules from ${testCase.name}`, () => {
      const modules = getModulesFromJSON(JSON.parse(testCase.json));
      expect(modules).toEqual(testCase.expected_modules);
    });
  });
});
