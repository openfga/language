import { loadDslSyntaxErrorTestCases } from "../transformer/_testcases";
import { OpenFgaDslSyntaxMultipleError } from "../transformer/dsltojson";
import validateDsl from "./validate-dsl";

describe("validateDsl", () => {
  const testCases = loadDslSyntaxErrorTestCases();
  testCases.forEach((testCase) => {

    const errorsCount = testCase.expected_errors.length;
    it(`case ${testCase.name} should throw ${errorsCount} errors on validation`, () => {

      try {
        validateDsl(testCase.dsl);
      } catch(thrownError) {

        const exception = thrownError as OpenFgaDslSyntaxMultipleError;

        expect(exception.errors.length).toEqual(errorsCount);

        if (errorsCount) {
          expect(exception.message).toEqual(testCase.error_message);
  
          for (let i = 0; i < errorsCount; i++) {
            const expectedError = testCase.expected_errors[i];
  
            expect(exception.errors[i].msg).toEqual(expectedError.msg);
            expect(exception.errors[i].line).toEqual(expectedError.line);
            expect(exception.errors[i].column).toEqual(expectedError.column);
            
            const resultMetadata = exception.errors[i].metadata;
            const expectedMetadata = expectedError.metadata;

            if (expectedMetadata) {
              expect(resultMetadata?.symbol).toEqual(expectedMetadata.symbol);
              expect(resultMetadata?.start).toEqual(expectedMetadata.start);
              expect(resultMetadata?.stop).toEqual(expectedMetadata.stop);
            } else {
              expect(resultMetadata).toBeUndefined();
            }

          }
        }

      }
      
    });
  });
});