import * as path from "path";
import * as fs from "fs";

interface TestCase {
  name: string;
  dsl: string;
  json: string;
  skip: boolean;
}

export function loadTransformerTestCases(): TestCase[] {
  const testDataPath = path.join(__dirname, "../../../tests", "data", "transformer");
  const entries = fs.readdirSync(testDataPath, { withFileTypes: true });

  const testCases: TestCase[] = [];

  for (const entry of entries) {
    if (!entry.isDirectory()) {
      continue;
    }

    const testCase: Partial<TestCase> = {
      name: entry.name,
    };

    try {
      const skipFile = fs.readFileSync(path.join(testDataPath, testCase.name!, "test.skip"));
      if (skipFile) {
        testCase.skip = true;
      }
    } catch (e) {
      // do nothing
    }

    const jsonData = fs.readFileSync(path.join(testDataPath, testCase.name!, "authorization-model.json"));
    testCase.json = jsonData.toString("utf8");

    const dslData = fs.readFileSync(path.join(testDataPath, testCase.name!, "authorization-model.fga"));
    testCase.dsl = dslData.toString("utf8");

    testCases.push(testCase as TestCase);
  }

  return testCases;
}
