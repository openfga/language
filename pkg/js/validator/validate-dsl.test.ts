import { loadDslSyntaxErrorTestCases } from "../transformer/_testcases";
import { OpenFgaDslSyntaxMultipleError } from "../transformer/dsltojson";
import validateDsl from "./validate-dsl";

describe("validateDsl", () => {
  const testCases = loadDslSyntaxErrorTestCases();
  testCases.forEach((testCase) => {

    const errorsCount = testCase.expected_errors.length;
    it(`case ${testCase.name} should throw ${errorsCount} errors on validation`, () => {

      if (!testCase.expected_errors?.length) {
        expect(() => validateDsl(testCase.dsl)).not.toThrow();
        return;
      }

      expect(() => validateDsl(testCase.dsl)).toThrowError(OpenFgaDslSyntaxMultipleError);
      try {
        validateDsl(testCase.dsl);
      } catch(thrownError) {

        const exception = thrownError as OpenFgaDslSyntaxMultipleError;
        if (errorsCount) {
          expect(exception.message).toEqual(`${testCase.expected_errors.length} error${testCase.expected_errors.length === 1 ? "" : "s"} occurred:\n\t* ${testCase.expected_errors.map(err => `syntax error at line=${err.line}, column=${err.column}: ${err.msg}`).join("\n\t* ")}\n\n`);

          for (let i = 0; i < errorsCount; i++) {
            expect(exception.errors[i]).toMatchObject(testCase.expected_errors[i]);
          }
        }
      }
    });
  });
});