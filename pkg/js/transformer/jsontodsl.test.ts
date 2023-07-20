import { loadTransformerTestCases } from "./_testcases";
import { transformJSONStringToDSL } from "./jsontodsl";

describe("jsonToDSL", () => {
  const testCases = loadTransformerTestCases();

  testCases.forEach((testCase) => {
    const testFn = testCase.skip ? it.skip : it;

    testFn(`should transform ${testCase.name} from JSON to DSL`, () => {
      const dslSyntax = transformJSONStringToDSL(testCase.json);
      expect(dslSyntax).toEqual(testCase.dsl);
    });
  });
});
