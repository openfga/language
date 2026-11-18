import { describe, expect, it } from "@jest/globals";
import { YamlStoreValidator, YamlTuplesValidator, isStringValue } from "../validator/validate-store";
import type { AuthorizationModel } from "@openfga/sdk";
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

    testFn(`should validate ${testCase.name} `, () => {
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

describe("validate store file with a user that is missing its type prefix", () => {
  const model = `model
  schema 1.1
type user
type document
  relations
    define viewer: [user]
`;

  it("should report errors for a tuple user with no colon rather than failing to run", () => {
    const schemaValidator: ValidateFunction = YamlStoreValidator();
    const jsonModel = transformDSLToJSONObject(model);

    const isValid = schemaValidator.call(
      { jsonModel },
      {
        model,
        tuples: [{ user: "anne", relation: "viewer", object: "document:1" }],
        tests: [],
      },
    );

    expect(isValid).toBe(false);
    expect(schemaValidator.errors?.length).toBeGreaterThan(0);
  });

  it("should report errors for a check user with no colon rather than failing to run", () => {
    const schemaValidator: ValidateFunction = YamlStoreValidator();
    const jsonModel = transformDSLToJSONObject(model);

    const isValid = schemaValidator.call(
      { jsonModel },
      {
        model,
        tests: [
          {
            name: "no type prefix on the check user",
            check: [{ user: "anne", object: "document:1", assertions: { viewer: true } }],
          },
        ],
      },
    );

    expect(isValid).toBe(false);
    expect(schemaValidator.errors?.length).toBeGreaterThan(0);
  });
});

describe("validate tuples against a model whose metadata omits a relation", () => {
  it("should report the missing type restriction rather than failing to run", () => {
    const schemaValidator: ValidateFunction = YamlTuplesValidator();
    const jsonModel = {
      schema_version: "1.1",
      type_definitions: [
        { type: "user" },
        {
          type: "document",
          relations: { viewer: { this: {} }, editor: { this: {} } },
          metadata: {
            relations: { editor: { directly_related_user_types: [{ type: "user" }] } },
          },
        },
      ],
    } as unknown as AuthorizationModel;

    const isValid = schemaValidator.call({ jsonModel }, [
      { user: "user:anne", relation: "viewer", object: "document:1" },
    ]);

    expect(isValid).toBe(false);
    expect(schemaValidator.errors?.length).toBeGreaterThan(0);
  });
});
