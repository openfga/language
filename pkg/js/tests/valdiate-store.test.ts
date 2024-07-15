import { describe, expect, it } from "@jest/globals";
import { YamlStoreValidator, isStringValue } from "../validator/validate-store";
import { transformDSLToJSONObject } from "../transformer";
import { ValidateFunction } from "ajv";
import * as fs from "fs";
import * as YAML from "yaml";
import { loadStoreTestCases } from "./_testcases";
import path from "path";

describe("validate valid store file", () => {
  const testCases = loadStoreTestCases();

  testCases.forEach((testCase) => {
    const testFn = testCase.skip ? it.skip : it;

    testFn(`should valdiate ${testCase.name} `, () => {
      const schemaValidator: ValidateFunction = YamlStoreValidator();

      const yaml = YAML.parseDocument(fs.readFileSync(testCase.store).toString());

      let jsonModel;
      if (yaml.has("model") && isStringValue(yaml.get("model"))) {
        jsonModel = transformDSLToJSONObject(String(yaml.get("model", false)));
      }

      if (yaml.has("model_file") && isStringValue(yaml.get("model_file"))) {
        jsonModel = transformDSLToJSONObject(
          fs.readFileSync(path.join(testCase.store, "..", String(yaml.get("model_file", false))), "utf-8"),
        );
      }

      if (!schemaValidator.call({ jsonModel }, yaml.toJSON())) {
        if (schemaValidator.errors) {
          expect(schemaValidator.errors).toEqual(testCase.expected_errors);
        }
      }
    });
  });
});
