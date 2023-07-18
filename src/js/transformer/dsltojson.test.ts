import { loadTransformerTestCases } from "./_testcases";
import transformDslToJSON from "./dsltojson";

describe("dslToJSON", () => {
  const testCases = loadTransformerTestCases();

  testCases.forEach((testCase) => {
    const testFn = testCase.skip ? it.skip : it;

    testFn(`should transform ${testCase.name} from DSL to JSON`, () => {
      const jsonSyntax = transformDslToJSON(testCase.dsl);
      expect(jsonSyntax).toEqual(JSON.parse(testCase.json));
    });
  });
});
