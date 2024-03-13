import {
    DSLSyntaxSingleError,
    ModelValidationSingleError,
    ModuleTransformationError,
    ModuleTransformationSingleError
} from "../../errors";
import { transformModuleFilesToModel } from "../../transformer/modules/modules-to-model";
import { loadModuleTestCases } from "../_testcases";


describe("transformModuleFilesToModel", () => {
    const testCases = loadModuleTestCases();
    testCases.forEach((testCase) => {
        const testFn = testCase.skip ? it.skip : it;

        testFn(`transformModuleFilesToModel ${testCase.name}`, () => {
            if (!testCase.expected_errors) {
                expect(transformModuleFilesToModel(testCase.modules, "1.2")).toEqual(
                    JSON.parse(testCase.json)
                );
            } else {
                try {
                    transformModuleFilesToModel(testCase.modules,  "1.2");
                } catch (error) {
                    expect(error).toBeInstanceOf(ModuleTransformationError);
                    const exception = error as ModuleTransformationError;

                    const errorsCount = testCase.expected_errors.length;
                    expect(exception.message).toEqual(
                        `${errorsCount} error${errorsCount === 1 ? "" : "s"} occurred:\n\t* ${testCase.expected_errors
                        .map((err: ModelValidationSingleError|ModuleTransformationSingleError|DSLSyntaxSingleError) => {
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
                        expect(exception.errors[index]).toMatchObject(testCase.expected_errors[index]);
                    }
                }
            }
        });
    });

    it("should allow passing a custom schema version", () => {
        const model = transformModuleFilesToModel([{
            name: "core.fga",
            contents: `module core
  type user`
        }], "1.1");

        expect(model.schema_version).toEqual("1.1");
    });
});