import { describe, expect, it } from "@jest/globals";
import { loadDSLSyntaxErrorTestCases, loadValidTransformerTestCases } from "./_testcases";
import { transformDSLToJSON, transformDSLToJSONObject } from "../transformer";

describe("dslToJSON", () => {
  const validTestCases = loadValidTransformerTestCases();

  validTestCases.forEach((testCase) => {
    const testFn = testCase.skip ? it.skip : it;

    testFn(`should transform ${testCase.name} from DSL to JSON`, () => {
      const jsonSyntax = transformDSLToJSONObject(testCase.dsl);
      expect(jsonSyntax).toEqual(JSON.parse(testCase.json));
    });
  });

  // These just ensure we are calling validate DSL when transforming
  const syntacticTestCases = loadDSLSyntaxErrorTestCases();
  syntacticTestCases.forEach((testCase) => {
    const errorsCount = testCase.expected_errors?.length || 0;
    it(`case ${testCase.name} should pass`, () => {
      if (!errorsCount) {
        expect(() => transformDSLToJSON(testCase.dsl)).not.toThrow();
      } else {
        expect(() => transformDSLToJSON(testCase.dsl)).toThrow(testCase.error_message);
      }
    });
  });
});
