import { describe, expect, it } from "@jest/globals";
import { transformDSLToJSONWithComments, transformDSLToJSONObjectWithComments } from "../transformer/dsltojson";
import { transformJSONStringToDSLWithComments, transformJSONToDSLWithComments } from "../transformer/jsontodsl";

describe("Comment Preservation", () => {
  describe("DSL to JSON with Comments", () => {
    it("should preserve model comments", () => {
      const dsl = `# OpenFGA Model
# Version 1.0
model
  schema 1.1

type user`;

      const result = transformDSLToJSONObjectWithComments(dsl);

      expect(result.modelComments).toBeDefined();
      expect(result.modelComments?.preceding_lines).toEqual(["# OpenFGA Model", "# Version 1.0"]);
    });

    it("should preserve type comments", () => {
      const dsl = `model
  schema 1.1

# User type comment
type user`;

      const result = transformDSLToJSONObjectWithComments(dsl);

      expect(result.typeComments["user"]).toBeDefined();
      expect(result.typeComments["user"].comments?.preceding_lines).toEqual(["# User type comment"]);
    });

    it("should preserve relation comments", () => {
      const dsl = `model
  schema 1.1

type document
  relations
    # Owner comment
    define owner: [user]`;

      const result = transformDSLToJSONObjectWithComments(dsl);

      expect(result.typeComments["document"]).toBeDefined();
      expect(result.typeComments["document"].relation_comments?.["owner"]).toBeDefined();
      expect(result.typeComments["document"].relation_comments?.["owner"].preceding_lines).toEqual(["# Owner comment"]);
    });

    it("should preserve condition comments", () => {
      const dsl = `model
  schema 1.1

type user

# IP-based access control
condition ip_check(ip: string) {
  ip == "127.0.0.1"
}`;

      const result = transformDSLToJSONObjectWithComments(dsl);

      expect(result.conditionComments["ip_check"]).toBeDefined();
      expect(result.conditionComments["ip_check"].preceding_lines).toEqual(["# IP-based access control"]);
    });

    it("should embed comments in JSON output", () => {
      const dsl = `# Model comment
model
  schema 1.1

# User type
type user`;

      const jsonStr = transformDSLToJSONWithComments(dsl);
      const json = JSON.parse(jsonStr);

      expect(json.metadata?.model_comments?.preceding_lines).toEqual(["# Model comment"]);
      expect(json.type_definitions?.[0].metadata?.comments?.preceding_lines).toEqual(["# User type"]);
    });
  });

  describe("JSON to DSL with Comments", () => {
    it("should emit model comments", () => {
      const json = {
        schema_version: "1.1",
        metadata: {
          model_comments: {
            preceding_lines: ["# Model comment"],
          },
        },
        type_definitions: [{ type: "user" }],
      };

      const dsl = transformJSONToDSLWithComments(json);

      expect(dsl).toContain("# Model comment\nmodel");
    });

    it("should emit type comments", () => {
      const json = {
        schema_version: "1.1",
        type_definitions: [
          {
            type: "user",
            metadata: {
              comments: {
                preceding_lines: ["# User type"],
              },
            },
          },
        ],
      };

      const dsl = transformJSONToDSLWithComments(json);

      expect(dsl).toContain("# User type\ntype user");
    });

    it("should emit relation comments", () => {
      const json = {
        schema_version: "1.1",
        type_definitions: [
          {
            type: "document",
            relations: {
              owner: { this: {} },
            },
            metadata: {
              relations: {
                owner: {
                  directly_related_user_types: [{ type: "user" }],
                  comments: {
                    preceding_lines: ["# Owner relation"],
                  },
                },
              },
            },
          },
        ],
      };

      const dsl = transformJSONToDSLWithComments(json);

      expect(dsl).toContain("# Owner relation\n    define owner:");
    });

    it("should emit condition comments", () => {
      const json = {
        schema_version: "1.1",
        type_definitions: [{ type: "user" }],
        conditions: {
          ip_check: {
            name: "ip_check",
            // eslint-disable-next-line quotes
            expression: 'ip == "127.0.0.1"',
            parameters: {
              ip: { type_name: "TYPE_NAME_STRING" as const },
            },
            metadata: {
              comments: {
                preceding_lines: ["# IP condition"],
              },
            },
          },
        },
      };

      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      const dsl = transformJSONToDSLWithComments(json as any);

      expect(dsl).toContain("# IP condition\ncondition ip_check");
    });

    it("should parse JSON string with comments", () => {
      const jsonStr = `{
        "schema_version": "1.1",
        "metadata": {
          "model_comments": {
            "preceding_lines": ["# Model comment"]
          }
        },
        "type_definitions": [{"type": "user"}]
      }`;

      const dsl = transformJSONStringToDSLWithComments(jsonStr);

      expect(dsl).toContain("# Model comment\nmodel");
    });
  });

  describe("Round Trip Comment Preservation", () => {
    it("should preserve comments through DSL -> JSON -> DSL", () => {
      const originalDSL = `# Model header comment
model
  schema 1.1

# User type
type user

# Document type
type document
  relations
    # Owner of document
    define owner: [user]
`;

      // DSL -> JSON
      const jsonStr = transformDSLToJSONWithComments(originalDSL);

      // JSON -> DSL
      const dsl = transformJSONStringToDSLWithComments(jsonStr);

      // Verify comments are preserved
      expect(dsl).toContain("# Model header comment");
      expect(dsl).toContain("# User type");
      expect(dsl).toContain("# Document type");
      expect(dsl).toContain("# Owner of document");
    });

    it("should handle multiple preceding comments", () => {
      const originalDSL = `# Comment line 1
# Comment line 2
# Comment line 3
model
  schema 1.1

type user
`;

      const jsonStr = transformDSLToJSONWithComments(originalDSL);
      const dsl = transformJSONStringToDSLWithComments(jsonStr);

      expect(dsl).toContain("# Comment line 1");
      expect(dsl).toContain("# Comment line 2");
      expect(dsl).toContain("# Comment line 3");
    });
  });

  describe("Backward Compatibility", () => {
    it("should handle JSON without comments", () => {
      const json = {
        schema_version: "1.1",
        type_definitions: [{ type: "user" }],
      };

      const dsl = transformJSONToDSLWithComments(json);

      expect(dsl).toContain("type user");
      // Should not throw and should produce valid DSL
    });

    it("should handle DSL without comments", () => {
      const dsl = `model
  schema 1.1

type user`;

      const result = transformDSLToJSONObjectWithComments(dsl);

      // Should not have model comments
      expect(result.modelComments).toBeUndefined();

      // Type comments should be empty
      expect(result.typeComments["user"]).toBeUndefined();
    });
  });

  describe("Edge Cases", () => {
    it("should handle empty lines between comments and elements", () => {
      // Empty lines should break comment association
      const dsl = `# This comment should not be associated

model
  schema 1.1

type user`;

      const result = transformDSLToJSONObjectWithComments(dsl);

      // The empty line breaks the association
      expect(result.modelComments).toBeUndefined();
    });

    it("should preserve inline comments", () => {
      const dsl = `model
  schema 1.1

type user # inline comment`;

      const result = transformDSLToJSONObjectWithComments(dsl);

      expect(result.typeComments["user"]).toBeDefined();
      expect(result.typeComments["user"].comments?.inline).toBe("# inline comment");
    });

    it("should handle special characters in comments", () => {
      const dsl = `# Comment with special chars: @#$%^&*()
model
  schema 1.1

type user`;

      const jsonStr = transformDSLToJSONWithComments(dsl);
      const dslResult = transformJSONStringToDSLWithComments(jsonStr);

      expect(dslResult).toContain("# Comment with special chars: @#$%^&*()");
    });
  });
});
