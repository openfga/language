import { loadDSLSyntaxErrorTestCases, loadDSLValidationErrorTestCases } from "./_testcases";
import { DSLSyntaxError, DSLSyntaxSingleError, ModelValidationError, ModelValidationSingleError } from "../errors";
import { validateDSL } from "../validator";

describe("validateDSL", () => {
  const syntacticTests = loadDSLSyntaxErrorTestCases();
  [syntacticTests[2]].forEach((testCase) => {
    const errorsCount = testCase.expected_errors?.length || 0;

    const testFn = testCase.skip ? it.skip : it;

    testFn(`case ${testCase.name} should throw ${errorsCount} errors on validation`, () => {
      if (!errorsCount) {
        expect(() => validateDSL(testCase.dsl)).not.toThrow();
        return;
      }

      expect(() => validateDSL(testCase.dsl)).toThrowError(DSLSyntaxError);
      try {
        validateDSL(testCase.dsl);
      } catch (thrownError) {
        const exception = thrownError as DSLSyntaxError;
        if (errorsCount) {
          expect(exception.message).toEqual(
            `${errorsCount} error${errorsCount === 1 ? "" : "s"} occurred:\n\t* ${testCase.expected_errors
              .map((err: DSLSyntaxSingleError) => {
                return `syntax error at line=${err.line.start}, column=${err.column.start}: ${err.msg}`;
              })
              .join("\n\t* ")}\n\n`,
          );

          for (let index = 0; index < errorsCount; index++) {
            expect(exception.errors[index]).toMatchObject(testCase.expected_errors[index]);
          }
        }
      }
    });
  });

  const semanticTests = loadDSLValidationErrorTestCases();
  semanticTests.forEach((testCase) => {
    const errorsCount = testCase.expected_errors?.length || 0;
    const testFn = testCase.skip ? it.skip : it;

    testFn(`case ${testCase.name} should throw ${errorsCount} errors on validation`, () => {
      if (!errorsCount) {
        expect(() => validateDSL(testCase.dsl)).not.toThrow();
        return;
      }

      expect(() => validateDSL(testCase.dsl)).toThrowError(ModelValidationError);
      try {
        validateDSL(testCase.dsl);
      } catch (thrownError) {
        const exception = thrownError as ModelValidationError;
        if (errorsCount) {
          expect(exception.message).toEqual(
            `${errorsCount} error${errorsCount === 1 ? "" : "s"} occurred:\n\t* ${testCase.expected_errors
              .map((err: ModelValidationSingleError) => {
                const errorType = err.metadata?.errorType || "validation";
                return `${errorType} error at line=${err.line.start}, column=${err.column.start}: ${err.msg}`;
              })
              .join("\n\t* ")}\n\n`,
          );

          for (let index = 0; index < errorsCount; index++) {
            expect(exception.errors[index]).toMatchObject(testCase.expected_errors[index]);
          }
        }
      }
    });
  });
});
