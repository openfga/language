import { loadInvalidDslSyntaxTestCases, loadValidTransformerTestCases } from "./_testcases";
import transformDslToJSON from "./dsltojson";

describe("dslToJSON", () => {
  const validTestCases = loadValidTransformerTestCases();

  validTestCases.forEach((testCase) => {
    const testFn = testCase.skip ? it.skip : it;

    testFn(`should transform ${testCase.name} from DSL to JSON`, () => {
      const jsonSyntax = transformDslToJSON(testCase.dsl);
      expect(jsonSyntax).toEqual(JSON.parse(testCase.json));
    });
  });

  const invalidTestCases = loadInvalidDslSyntaxTestCases();
  invalidTestCases.forEach((testCase) => {
    it(`case ${testCase.name} should pass`, () => {
      if (testCase.valid) {
        expect(() => transformDslToJSON(testCase.dsl)).not.toThrow();
      } else {
        expect(() => transformDslToJSON(testCase.dsl)).toThrow(testCase.error_message);
      }
    });
  });
});
