import { loadDslSyntaxErrorTestCases } from "../transformer/_testcases";
import validateDsl from "./validate-dsl";

const testCases = loadDslSyntaxErrorTestCases();
  testCases.forEach((testCase) => {

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