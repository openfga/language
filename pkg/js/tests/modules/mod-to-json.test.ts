import { FGAModFileValidationError, FGAModFileValidationSingleError } from "../../errors";
import { transformModFileToJSON } from "../../transformer/modules/mod-to-json";
import { loadModFileTestCases } from "../_testcases";

describe("modFileToJSON", () => {

    const testCases = loadModFileTestCases();

    testCases.forEach((testCase) => {
        const testFn = testCase.skip ? it.skip : it;
        testFn(`transformModFileToJSON ${testCase.name}`, () => {
            if (!testCase.expected_errors) {
                expect(transformModFileToJSON(testCase.modFile)).toEqual(JSON.parse(testCase.json));
            } else {
                expect(() => transformModFileToJSON(testCase.modFile)).toThrow(FGAModFileValidationError);
                try {
                    transformModFileToJSON(testCase.modFile);
                } catch (error) {
                    const exception = error as FGAModFileValidationError;
                    if (testCase.expected_errors) {
                        const errorsCount = testCase.expected_errors.length;
                        expect(exception.message).toEqual(
                          `${errorsCount} error${errorsCount === 1 ? "" : "s"} occurred:\n\t* ${testCase.expected_errors
                            .map((err: FGAModFileValidationSingleError) => {
                              return `validation error at line=${err.line?.start}, column=${err.column?.start}: ${err.msg}`;
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
            }
        });
    });
});