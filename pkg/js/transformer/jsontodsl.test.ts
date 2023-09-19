import { loadValidTransformerTestCases, loadInvalidJSONSyntaxTestCases } from "./_testcases";
import { transformJSONStringToDSL } from "./jsontodsl";

describe("jsonToDSL", () => {
  const testCases = loadValidTransformerTestCases();
  const invalidTestCases = loadInvalidJSONSyntaxTestCases();

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
      expect(() => transformJSONStringToDSL(testCase.json)).toThrow(testCase.error_message);
    });
  });
});
