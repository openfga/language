import { loadValidTransformerTestCases, loadInvalidJSONSyntaxTestCases, loadModuleTestCases } from "./_testcases";
import { transformJSONStringToDSL } from "../transformer";

describe("jsonToDSL", () => {
  const testCases = loadValidTransformerTestCases();
  const invalidTestCases = loadInvalidJSONSyntaxTestCases();
  const moduleTestCases = loadModuleTestCases();

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
    if (!testCase.dsl) {
      return;
    }
    const testFn = testCase.skip ? it.skip : it;

    testFn(`should transform ${testCase.name} from JSON to DSL with comments`, () => {
      const dslSyntax = transformJSONStringToDSL(testCase.json);
      expect(dslSyntax).toEqual(testCase.dsl);
    });
  });
});
