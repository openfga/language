import { describe, expect, it } from "@jest/globals";
import Ajv from "ajv";
import * as fs from "fs";
import * as path from "path";
import * as YAML from "yaml";
import { Rules, Validator } from "../validator/validate-rules";

const schemasDir = path.join(__dirname, "../../../schemas");
const storesDir = path.join(__dirname, "../../../tests/data/stores");

const readSchema = (name: string) => JSON.parse(fs.readFileSync(path.join(schemasDir, name), "utf-8"));

const tuplesSchema = readSchema("tuples.schema.json");
const storeSchema = readSchema("store.schema.json");

const newAjv = () => new Ajv({ allErrors: true }).addSchema(tuplesSchema);

const tupleProps = tuplesSchema.definitions.tuple.properties;

describe("published JSON schemas", () => {
  it("are valid draft-07 schemas", () => {
    const ajv = newAjv();
    expect(ajv.validateSchema(tuplesSchema)).toBe(true);
    expect(ajv.validateSchema(storeSchema)).toBe(true);
    expect(() => ajv.compile(storeSchema)).not.toThrow();
  });

  it("build their patterns from the shared validation rules", () => {
    const sized = "(?=\\S{2,256}$)";
    expect(tupleProps.relation.pattern).toBe(`^${Rules.relation}$`);
    expect(tupleProps.condition.properties.name.pattern).toBe(`^${Rules.condition}$`);
    expect(tupleProps.object.pattern).toBe(`^${sized}${Rules.type}:${Rules.id}$`);
    expect(tupleProps.user.pattern).toBe(
      `^(?:${Rules.type}:${Rules.id}#${Rules.relation}|${sized}${Rules.type}:${Rules.id}|${Rules.type}:\\*)$`,
    );
  });

  describe("user and object patterns agree with the validator", () => {
    const values = [
      "user:anne",
      "user:*",
      "group:eng#member",
      "document:2021-budget",
      "folder:a/b.c|d+e@f",
      "user",
      "user:",
      ":anne",
      "user:anne#",
      "user:#member",
      "user anne:1",
      "group:*#member",
      "a:b",
      `document:${"x".repeat(260)}`,
    ];

    values.forEach((value) => {
      it(`for ${value.length > 40 ? value.slice(0, 40) + "..." : value}`, () => {
        expect(new RegExp(tupleProps.user.pattern).test(value)).toBe(Validator.user(value));
        expect(new RegExp(tupleProps.object.pattern).test(value)).toBe(Validator.object(value));
      });
    });
  });

  it("accept a tuples file", () => {
    const validate = new Ajv({ allErrors: true }).compile(tuplesSchema);
    const tuples = YAML.parse(`
- user: user:anne
  relation: viewer
  object: document:roadmap
- user: group:eng#member
  relation: viewer
  object: document:roadmap
  condition:
    name: in_office_hours
    context:
      hour: 9
`);
    expect(validate(tuples)).toBe(true);
  });

  it("reject invalid tuples", () => {
    const validate = new Ajv({ allErrors: true }).compile(tuplesSchema);
    expect(validate([{ user: "anne", relation: "viewer", object: "document:roadmap" }])).toBe(false);
    expect(validate([{ user: "user:anne", object: "document:roadmap" }])).toBe(false);
    expect(validate([{ user: "user:anne", relation: "viewer", object: "document:roadmap", extra: 1 }])).toBe(false);
  });

  describe("store schema on the store test data", () => {
    const validate = newAjv().compile(storeSchema);

    fs.readdirSync(storesDir).forEach((dir) => {
      const expectsErrors = fs.existsSync(path.join(storesDir, dir, "expected_errors.json"));
      const store = YAML.parse(fs.readFileSync(path.join(storesDir, dir, "store.fga.yaml"), "utf-8"));
      const structural = dir.includes("additional-properties");

      // missing-model is a semantic error the schema cannot see, so only check structural cases
      if (!expectsErrors || structural) {
        it(`${expectsErrors ? "rejects" : "accepts"} ${dir}`, () => {
          expect(validate(store)).toBe(!expectsErrors);
        });
      }
    });

    it("accepts the CLI's users, objects and tuple_files fields", () => {
      expect(
        validate({
          tuple_files: ["a.yaml", "b.yaml"],
          tests: [
            {
              check: [
                {
                  users: ["user:anne", "user:beth"],
                  objects: ["document:roadmap"],
                  assertions: { viewer: true },
                },
              ],
            },
          ],
        }),
      ).toBe(true);
    });

    it("rejects a check with both user and users, or both object and objects", () => {
      const assertions = { viewer: true };
      const bothUsers = { user: "user:anne", users: ["user:beth"], object: "document:a", assertions };
      const bothObjects = { user: "user:anne", object: "document:a", objects: ["document:b"], assertions };
      expect(validate({ tests: [{ check: [bothUsers] }] })).toBe(false);
      expect(validate({ tests: [{ check: [bothObjects] }] })).toBe(false);
    });

    it("requires a user and an object in a check", () => {
      expect(validate({ tests: [{ check: [{ object: "document:roadmap", assertions: { viewer: true } }] }] })).toBe(
        false,
      );
    });
  });
});
