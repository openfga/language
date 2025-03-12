import { describe, expect, it } from "@jest/globals";
import {
  DSLSyntaxSingleError,
  ModelValidationSingleError,
  ModuleTransformationError,
  ModuleTransformationSingleError,
} from "../../errors";
import { transformModuleFilesToModel } from "../../transformer/modules/modules-to-model";
import { loadDSLValidationErrorTestCases, loadModuleTestCases } from "../_testcases";

describe("transformModuleFilesToModel - module test cases", () => {
  const moduleTestCases = loadModuleTestCases();
  moduleTestCases.forEach((testCase) => {
    const testFn = testCase.skip ? it.skip : it;

    if (!testCase.modules) {
      return;
    }

    testFn(`transformModuleFilesToModel ${testCase.name}`, () => {
      if (!testCase.expected_errors) {
        expect(transformModuleFilesToModel(testCase.modules, "1.2")).toEqual(JSON.parse(testCase.json));
      } else {
        try {
          transformModuleFilesToModel(testCase.modules, "1.2");
        } catch (error) {
          expect(error).toBeInstanceOf(ModuleTransformationError);
          const exception = error as ModuleTransformationError;

          const errorsCount = testCase.expected_errors.length;
          expect(exception.message).toEqual(
            `${errorsCount} error${errorsCount === 1 ? "" : "s"} occurred:\n\t* ${testCase.expected_errors
              .map((err: ModelValidationSingleError | ModuleTransformationSingleError | DSLSyntaxSingleError) => {
                let errorType = "transformation-error";
                if ((err as ModelValidationSingleError).metadata?.errorType) {
                  errorType = (err as ModelValidationSingleError).metadata!.errorType;
                } else if ((err as DSLSyntaxSingleError).type) {
                  errorType = err.type;
                }

                let msg = `${errorType} error`;
                if (!err.metadata || err.type) {
                  msg += ` at line=${err.line?.start}, column=${err.column?.start}`;
                }
                msg += `: ${err.msg}`;
                return msg;
              })
              .join("\n\t* ")}\n\n`,
          );

          for (let index = 0; index < errorsCount; index++) {
            // We're asserting an error type against an JSON object here, it works but isn't type correct
            // @ts-expect-error
            expect(exception.errors[index]).toMatchObject(testCase.expected_errors[index]);
          }
        }
      }
    });
  });

  const semanticValidationTestCases = loadDSLValidationErrorTestCases();
  semanticValidationTestCases.forEach((testCase) => {
    let testFn = testCase.skip ? it.skip : it;

    // Skip any test case already using module as it won't be valid
    if (testCase.dsl.trim().startsWith("module")) {
      testFn = it.skip;
    }

    // Skip schema validation errors
    if (testCase.dsl.includes("0.9")) {
      testFn = it.skip;
    }

    testFn(`transformModuleFilesToModel ${testCase.name}`, () => {
      try {
        const dsl = testCase.dsl.replace(/model\n {2}schema 1\.1/, "module test\n");
        transformModuleFilesToModel([{ name: "test.fga", contents: dsl }], "1.2");
      } catch (error) {
        expect(error).toBeInstanceOf(ModuleTransformationError);
        const exception = error as ModuleTransformationError;

        const errorsCount = testCase.expected_errors.length;
        expect(exception.message).toEqual(
          `${errorsCount} error${errorsCount === 1 ? "" : "s"} occurred:\n\t* ${testCase.expected_errors
            .map((err: ModelValidationSingleError | ModuleTransformationSingleError | DSLSyntaxSingleError) => {
              let errorType = "transformation-error";
              if ((err as ModelValidationSingleError).metadata?.errorType) {
                errorType = (err as ModelValidationSingleError).metadata!.errorType;
              } else if ((err as DSLSyntaxSingleError).type) {
                errorType = err.type;
              }

              let msg = `${errorType} error`;
              if (!err.metadata || err.type) {
                msg += ` at line=${err.line?.start}, column=${err.column?.start}`;
              }
              msg += `: ${err.msg}`;
              return msg;
            })
            .join("\n\t* ")}\n\n`,
        );

        for (let index = 0; index < errorsCount; index++) {
          const expected = testCase.expected_errors[index];
          expected.metadata!.module = "test";
          // We're asserting an error type against an JSON object here, it works but isn't type correct
          // @ts-expect-error
          expect(exception.errors[index]).toMatchObject(expected);
        }
      }
    });
  });

  it("should allow passing a custom schema version", () => {
    const model = transformModuleFilesToModel(
      [
        {
          name: "core.fga",
          contents: `module core
  type user`,
        },
      ],
      "1.1",
    );

    expect(model.schema_version).toEqual("1.1");
  });
});
