import { loadDslSyntaxErrorTestCases, loadDslValidationErrorTestCases } from "../transformer/_testcases";
import { DSLSyntaxError, DSLSyntaxSingleError, ModelValidationError, ModelValidationSingleError } from "../errors";
import validateDsl from "./validate-dsl";

describe("validateDsl", () => {
  const syntacticTests = loadDslSyntaxErrorTestCases();
  syntacticTests.forEach((testCase) => {

    const errorsCount = testCase.expected_errors.length;

    const testFn = testCase.skip ? it.skip : it;

    testFn(`case ${testCase.name} should throw ${errorsCount} errors on validation`, () => {

      if (!testCase.expected_errors?.length) {
        expect(() => validateDsl(testCase.dsl)).not.toThrow();
        return;
      }

      expect(() => validateDsl(testCase.dsl)).toThrowError(DSLSyntaxError);
      try {
        validateDsl(testCase.dsl);
      } catch (thrownError) {

        const exception = thrownError as DSLSyntaxError;
        if (errorsCount) {
          expect(exception.message).toEqual(
            `${errorsCount} error${errorsCount === 1 ? "" : "s"} occurred:\n\t* ${testCase.expected_errors.map(
              (err: DSLSyntaxSingleError) => {
                return `syntax error at line=${err.line.start}, column=${err.column.start}: ${err.msg}`;
              }).join("\n\t* ")}\n\n`);

          for (let i = 0; i < errorsCount; i++) {
            expect(exception.errors[i]).toMatchObject(testCase.expected_errors[i]);
          }
        }
      }
    });
  });

  const semanticTests = loadDslValidationErrorTestCases();
  semanticTests.forEach((testCase) => {

    const errorsCount = testCase.expected_errors.length;
    const testFn = testCase.skip ? it.skip : it;

    testFn(`case ${testCase.name} should throw ${errorsCount} errors on validation`, () => {

      if (!testCase.expected_errors?.length) {
        expect(() => validateDsl(testCase.dsl)).not.toThrow();
        return;
      }

      expect(() => validateDsl(testCase.dsl)).toThrowError(ModelValidationError);
      try {
        validateDsl(testCase.dsl);
      } catch (thrownError) {

        const exception = thrownError as ModelValidationError;
        if (errorsCount) {
          expect(exception.message).toEqual(
            `${errorsCount} error${errorsCount === 1 ? "" : "s"} occurred:\n\t* ${testCase.expected_errors.map(
              (err: ModelValidationSingleError) => {
                const errorType = err.metadata?.errorType || "validation";
                return `${errorType} error at line=${err.line.start}, column=${err.column.start}: ${err.msg}`;
            }).join("\n\t* ")}\n\n`);

          for (let i = 0; i < errorsCount; i++) {
            expect(exception.errors[i]).toMatchObject(testCase.expected_errors[i]);
          }
        }
      }
    });
  });

});
