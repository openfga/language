import { getModuleForObjectTypeRelation, isRelationAssignable } from "./model_utils";
import { TypeDefinition, Userset, Usersets, Difference } from "@openfga/sdk";
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
});
