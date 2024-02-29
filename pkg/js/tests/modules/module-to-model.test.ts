import { BaseError, ModuleTransformationError } from "../../errors";
import { transformModuleFilesToModel } from "../../transformer/modules/modules-to-model";
import { loadModuleTestCases } from "../_testcases";


describe("transformModuleFilesToModel", () => {
    const testCases = loadModuleTestCases();
    testCases.forEach((testCase) => {
        const testFn = testCase.skip ? it.skip : it;

        testFn(`transformModuleFilesToModel ${testCase.name}`, () => {
            if (!testCase.expected_errors) {
                expect(transformModuleFilesToModel(testCase.modules)).toEqual(JSON.parse(testCase.json));
            } else {
                try {
                    transformModuleFilesToModel(testCase.modules);
                } catch (error) {
                    expect(error).toBeInstanceOf(ModuleTransformationError);
                    const exception = error as ModuleTransformationError;

                    const errorsCount = testCase.expected_errors.length;
                    expect(exception.message).toEqual(
                        `${errorsCount} error${errorsCount === 1 ? "" : "s"} occurred:\n\t* ${testCase.expected_errors
                        .map((err: BaseError) => {
                            //@ts-expect-error
                            let msg = `${err.metadata?.errorType || "validation-error"} error: `;
                            if (err?.line) {
                                msg += ` at line ${err.line.start}, column=${err.column?.start} :`;
                            }
                            msg += `${err.msg}`;
                            return msg;
                        })
                        .join("\n\t* ")}\n\n`,
                    );

                    // TODO: need to enable this
                    // for (let index = 0; index < errorsCount; index++) {
                    //     expect(exception.errors[index]).toMatchObject(testCase.expected_errors[index]);
                    // }
                }
            }
        });
    });
});