import { loadDslSyntaxErrorTestCases } from "../transformer/_testcases";
import validateDsl from "./validate-dsl";

describe("validateDsl", () => {
  const testCases = loadDslSyntaxErrorTestCases();
  testCases.forEach((testCase) => {

    const errorsCount = testCase.expected_errors.length;
    it(`case ${testCase.name} should return ${errorsCount} errors on validation`, () => {

      const result = validateDsl(testCase.dsl);

      expect(result.errors.length).toEqual(errorsCount);

      if (errorsCount) {
        expect(result.message).toEqual(testCase.error_message);

        for (let i = 0; i < errorsCount; i++) {
          const expectedError = testCase.expected_errors[i];

          expect(result.errors[i].msg).toEqual(expectedError.msg);
          expect(result.errors[i].line).toEqual(expectedError.line);
          expect(result.errors[i].column).toEqual(expectedError.column);
          
          if (expectedError.metadata) {
            const resultMetadata = result.errors[i].metadata;
            const expectedMetadata = expectedError.metadata;

            expect(resultMetadata?.symbol).toEqual(expectedMetadata.symbol);
            expect(resultMetadata?.start).toEqual(expectedMetadata.start);
            expect(resultMetadata?.stop).toEqual(expectedMetadata.stop);
          }
        }
      }
    });
  });
});