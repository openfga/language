import { 
  loadInvalidDslSyntaxTestCases,
  loadValidTransformerTestCases, 
  loadDslSyntaxErrorTestCases
} from "./_testcases";
import transformDslToJSON, { validateDsl } from "./dsltojson";

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

    const errorsCount = testCase.expected_errors.length;
    it(`case ${testCase.name} should throw ${errorsCount} errors`, () => {

      const result = validateDsl(testCase.dsl);

      expect(result.errors.length).toEqual(errorsCount);

      if (result.errors.length) {
        expect(result.message).toEqual(testCase.error_message);

        for(let i = 0; i < result.errors.length; i++) {
          const expectedError = testCase.expected_errors[i];

          expect(result.errors[i].msg).toEqual(expectedError.msg);
          expect(result.errors[i].line).toEqual(expectedError.line);
          expect(result.errors[i].column).toEqual(expectedError.column);
        }
      }
    });
  });
});