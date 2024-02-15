import {
  loadDSLSyntaxErrorTestCases,
  loadDSLValidationErrorTestCases,
  loadInvalidJSONSValidationTestCases,
} from "./_testcases";
import { DSLSyntaxError, DSLSyntaxSingleError, ModelValidationError, ModelValidationSingleError } from "../errors";
import { validateDSL, validateJSON } from "../validator";
import { transformDSLToJSON, transformDSLToJSONObject } from "../transformer";
import { AuthorizationModel } from "@openfga/sdk";

describe("validateDSL", () => {
  const syntacticTests = loadDSLSyntaxErrorTestCases();

  syntacticTests.forEach((testCase) => {
    const errorsCount = testCase.expected_errors?.length || 0;

    const testFn = testCase.skip ? it.skip : it;

    testFn(`case ${testCase.name} should throw ${errorsCount} errors on validation`, () => {
      if (!errorsCount) {
        expect(() => validateDSL(testCase.dsl)).not.toThrow();
        return;
      }

      expect(() => validateDSL(testCase.dsl)).toThrow(DSLSyntaxError);
      try {
        validateDSL(testCase.dsl);
      } catch (thrownError) {
        const exception = thrownError as DSLSyntaxError;
        if (errorsCount) {
          expect(exception.message).toEqual(
            `${errorsCount} error${errorsCount === 1 ? "" : "s"} occurred:\n\t* ${testCase.expected_errors
              .map((err: DSLSyntaxSingleError) => {
                return `syntax error at line=${err.line?.start}, column=${err.column?.start}: ${err.msg}`;
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

      expect(() => validateDSL(testCase.dsl)).toThrow(ModelValidationError);
      try {
        validateDSL(testCase.dsl);
      } catch (thrownError) {
        const exception = thrownError as ModelValidationError;
        if (errorsCount) {
          expect(exception.message).toEqual(
            `${errorsCount} error${errorsCount === 1 ? "" : "s"} occurred:\n\t* ${testCase.expected_errors
              .map((err: ModelValidationSingleError) => {
                const errorType = err.metadata?.errorType || "validation";

                return `${errorType} error at line=${err.line?.start}, column=${err.column?.start}: ${err.msg}`;
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

describe("validateJSON", () => {
  const testCases = loadInvalidJSONSValidationTestCases();

  testCases.forEach((testCase) => {
    const errorsCount = testCase.expected_errors?.length || 0;
    const testFn = testCase.skip ? it.skip : it;

    testFn(`case ${testCase.name} should throw ${errorsCount} errors on validation`, () => {
      const json = JSON.parse(testCase.json) as AuthorizationModel;

      if (!errorsCount) {
        expect(() => validateJSON(json)).not.toThrow();
        return;
      }

      expect(() => validateJSON(json)).toThrow(ModelValidationError);
      try {
        validateJSON(json);
      } catch (thrownError) {
        const exception = thrownError as ModelValidationError;
        if (errorsCount) {
          expect(exception.message).toEqual(
            `${errorsCount} error${errorsCount === 1 ? "" : "s"} occurred:\n\t* ${testCase.expected_errors
              .map((err: ModelValidationSingleError) => {
                const errorType = err.metadata?.errorType || "validation";

                return `${errorType} error: ${err.msg}`;
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

  const dslTestCases = loadDSLValidationErrorTestCases();

  dslTestCases.forEach((testCase) => {
    const errorsCount = testCase.expected_errors?.length || 0;
    const testFn = testCase.skip ? it.skip : it;

    testFn(`case ${testCase.name} should throw ${errorsCount} errors on validation`, () => {
      const json = transformDSLToJSONObject(testCase.dsl) as AuthorizationModel;

      if (!errorsCount) {
        expect(() => validateJSON(json)).not.toThrow();
        return;
      }

      expect(() => validateJSON(json)).toThrow();
      try {
        validateJSON(json);
      } catch (thrownError) {
        const exception = thrownError as ModelValidationError;
        if (errorsCount) {
          expect(exception.message).toEqual(
            `${errorsCount} error${errorsCount === 1 ? "" : "s"} occurred:\n\t* ${testCase.expected_errors
              .map((err: ModelValidationSingleError) => {
                const errorType = err.metadata?.errorType || "validation";

                return `${errorType} error: ${err.msg}`;
              })
              .join("\n\t* ")}\n\n`,
          );

          for (let index = 0; index < errorsCount; index++) {
            expect(exception.errors[index].msg === testCase.expected_errors[index].msg);
            expect(exception.errors[index].metadata).toMatchObject(testCase.expected_errors[index].metadata!);
          }
        }
      }
    });
  });
});
