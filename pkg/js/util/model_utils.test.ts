import { getModuleForObjectTypeRelation, isRelationAssignable, isModelModular } from "./model_utils";
import { AuthorizationModel, TypeDefinition, Userset, Usersets, Difference } from "@openfga/sdk";
import { describe } from "@jest/globals";

describe("model_utils", () => {
  describe("getModuleForObjectTypeRelation", () => {
    const typeDefWithModule: TypeDefinition = {
      type: "type1",
      relations: {
        relation1: {},
        relation2: {},
        relation3: {},
        relation4: {},
      },
      metadata: {
        module: "type_module1",
        relations: {
          relation1: { module: "module1" },
          relation2: { module: "" },
          relation3: {},
          relation5: {},
        },
      },
    };

    const typeDefWithoutModule: TypeDefinition = {
      type: "type2",
      relations: {
        relation7: {},
      },
    };

    test("Relation exists and has a module", () => {
      const result = getModuleForObjectTypeRelation(typeDefWithModule, "relation1");
      expect(result).toBe("module1");
    });

    test("Relation exists but has an empty string as a module, type has a module", () => {
      const result = getModuleForObjectTypeRelation(typeDefWithModule, "relation2");
      expect(result).toBe("type_module1");
    });

    test("Relation exists but does not have a module, type has a module", () => {
      const result = getModuleForObjectTypeRelation(typeDefWithModule, "relation3");
      expect(result).toBe("type_module1");
    });

    test("Relation exists but does not have metadata, type has a module", () => {
      const result = getModuleForObjectTypeRelation(typeDefWithModule, "relation4");
      expect(result).toBe("type_module1");
    });

    test("Relation does not exist", () => {
      expect(() => {
        getModuleForObjectTypeRelation(typeDefWithModule, "relation5");
      }).toThrow("relation relation5 does not exist in type type1");
    });

    test("Relation does not exist 2", () => {
      expect(() => {
        getModuleForObjectTypeRelation(typeDefWithModule, "relation6");
      }).toThrow("relation relation6 does not exist in type type1");
    });

    test("Relation exists but does not have a module, type does not have a module", () => {
      const result = getModuleForObjectTypeRelation(typeDefWithoutModule, "relation7");
      expect(result).toBeUndefined();
    });
  });

  describe("isRelationAssignable", () => {
    test("Relation definition has a key 'this'", () => {
      const relDef: Userset = { this: {} };
      expect(isRelationAssignable(relDef)).toBe(true);
    });

    test("Relation definition has a key 'union' with a child that has a key 'this'", () => {
      const relDef: Userset = { union: { child: [{ this: {} }] } as Usersets };
      expect(isRelationAssignable(relDef)).toBe(true);
    });

    test("Relation definition has a key 'intersection' with a child that has a key 'this'", () => {
      const relDef: Userset = { intersection: { child: [{ this: {} }] } as Usersets };
      expect(isRelationAssignable(relDef)).toBe(true);
    });

    test("Relation definition has a key 'difference' with base having a key 'this'", () => {
      const relDef: Userset = { difference: { base: { this: {} }, subtract: {} } as Difference };
      expect(isRelationAssignable(relDef)).toBe(true);
    });

    test("Relation definition has a key 'difference' with subtract having a key 'this'", () => {
      const relDef: Userset = { difference: { base: {}, subtract: { this: {} } } as Difference };
      expect(isRelationAssignable(relDef)).toBe(true);
    });

    test("Relation definition does not have any assignable keys", () => {
      const relDef: Userset = { union: { child: [{ intersection: { child: [{}] } }] } as Usersets };
      expect(isRelationAssignable(relDef)).toBe(false);
    });
  });

  describe("isModelModular", () => {
    test("Model with schema 1.2 and type with module is modular", () => {
      const model = {
        id: "test",
        schema_version: "1.2",
        type_definitions: [
          {
            type: "user",
            relations: {
              viewer: {},
            },
            metadata: {
              module: "user_module",
            },
          },
        ],
      } as AuthorizationModel;
      expect(isModelModular(model)).toBe(true);
    });

    test("Model with schema 1.2 and relation with module is modular", () => {
      const model = {
        id: "test",
        schema_version: "1.2",
        type_definitions: [
          {
            type: "document",
            relations: {
              viewer: {},
            },
            metadata: {
              relations: {
                viewer: { module: "viewer_module" },
              },
            },
          },
        ],
      } as AuthorizationModel;
      expect(isModelModular(model)).toBe(true);
    });

    test("Model with schema 1.1 and type with module is not modular", () => {
      const model = {
        id: "test",
        schema_version: "1.1",
        type_definitions: [
          {
            type: "user",
            relations: {
              viewer: {},
            },
            metadata: {
              module: "user_module",
            },
          },
        ],
      } as AuthorizationModel;
      expect(isModelModular(model)).toBe(false);
    });

    test("Model with schema 1.2 but no modules is not modular", () => {
      const model = {
        id: "test",
        schema_version: "1.2",
        type_definitions: [
          {
            type: "user",
            relations: {
              viewer: {},
            },
          },
        ],
      } as AuthorizationModel;
      expect(isModelModular(model)).toBe(false);
    });

    test("Model with schema 1.2 and empty string module is not modular", () => {
      const model = {
        id: "test",
        schema_version: "1.2",
        type_definitions: [
          {
            type: "user",
            relations: {
              viewer: {},
            },
            metadata: {
              module: "",
            },
          },
        ],
      } as AuthorizationModel;
      expect(isModelModular(model)).toBe(false);
    });

    test("Model with schema 1.2, no type module, but relation with module is modular", () => {
      const model = {
        id: "test",
        schema_version: "1.2",
        type_definitions: [
          {
            type: "document",
            relations: {
              viewer: {},
              editor: {},
            },
            metadata: {
              relations: {
                editor: { module: "editor_module" },
              },
            },
          },
        ],
      } as AuthorizationModel;
      expect(isModelModular(model)).toBe(true);
    });

    test("Model with no schema version throws error", () => {
      const model = {
        id: "test",
        schema_version: undefined as unknown as string,
        type_definitions: [
          {
            type: "user",
            relations: {
              viewer: {},
            },
            metadata: {
              module: "user_module",
            },
          },
        ],
      } as AuthorizationModel;
      expect(() => isModelModular(model)).toThrow("Unsupported schema version: undefined");
    });

    test("Model with unsupported schema version throws error", () => {
      const model = {
        id: "test",
        schema_version: "2.0",
        type_definitions: [
          {
            type: "user",
            relations: {
              viewer: {},
            },
            metadata: {
              module: "user_module",
            },
          },
        ],
      } as AuthorizationModel;
      expect(() => isModelModular(model)).toThrow("Unsupported schema version: 2.0");
    });

    test("Model with empty type definitions is not modular", () => {
      const model = {
        id: "test",
        schema_version: "1.2",
        type_definitions: [],
      } as AuthorizationModel;
      expect(isModelModular(model)).toBe(false);
    });
  });
});
