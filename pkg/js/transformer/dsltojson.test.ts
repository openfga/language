import { loadDslSyntaxErrorTestCases, loadValidTransformerTestCases } from "./_testcases";
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

  // These just ensure we are calling validate DSL when transforming
  const syntacticTestCases = loadDslSyntaxErrorTestCases();
  syntacticTestCases.forEach((testCase) => {
    const errorsCount = testCase.expected_errors?.length || 0;
    it(`case ${testCase.name} should pass`, () => {
      if (!errorsCount) {
        expect(() => transformDslToJSON(testCase.dsl)).not.toThrow();
      } else {
        expect(() => transformDslToJSON(testCase.dsl)).toThrow(testCase.error_message);
      }
    });
  });
});
