import { 
  loadInvalidDslSyntaxTestCases,
  loadValidTransformerTestCases, 
  loadDslSyntaxErrorTestCases
} from "./_testcases";
import transformDslToJSON from "./dsltojson";

describe("dslToJSON", () => {
  const testCases = loadValidTransformerTestCases();

  testCases.forEach((testCase) => {
    const testFn = testCase.skip ? it.skip : it;

    testFn(`should transform ${testCase.name} from DSL to JSON`, () => {
      const jsonSyntax = transformDslToJSON(testCase.dsl);
      expect(jsonSyntax).toEqual(JSON.parse(testCase.json));
    });
  });

  const testCases2 = loadInvalidDslSyntaxTestCases();
  testCases2.forEach((testCase) => {
    it(`case ${testCase.name} should pass`, () => {
      if (testCase.valid) {
        expect(() => transformDslToJSON(testCase.dsl)).not.toThrow();
      } else {
        expect(() => transformDslToJSON(testCase.dsl)).toThrow(testCase.error_message);
      }
    });
  });

  const testCase3 = loadDslSyntaxErrorTestCases();
  testCase3.forEach((testCase) => {
    const errorsCount = testCase.expectedError.length;
    it(`case ${testCase.name} should throw ${errorsCount} errors`, () => {

      if (errorsCount === 0) {
        expect(() => transformDslToJSON(testCase.dsl)).not.toThrow();
      } else {
        expect(() => transformDslToJSON(testCase.dsl)).toThrow(testCase.error_message);
      }
    });
  });
});