import { AuthorizationModel, Difference, TypeDefinition, Userset, Usersets } from "@openfga/sdk";
import { getSchemaVersionFromString, checkSchemaVersionSupportsModules } from "./schema_version";

/**
 * getModuleForObjectTypeRelation - returns the module for the given object type and relation in that type.
 * @param typeDef - A TypeDefinition object which contains metadata about the type.
 * @param relation - A string representing the relation whose module is to be retrieved.
 * @return string - A string representing the module for the given object type and relation.
 * @error error - An error if the relation does not exist.
 */
export function getModuleForObjectTypeRelation(typeDef: TypeDefinition, relation: string): string | undefined {
  if (!typeDef.relations?.[relation]) {
    throw new Error(`relation ${relation} does not exist in type ${typeDef.type}`);
  }

  const relationsMetadata = typeDef?.metadata?.relations || {};
  const relationMetadata = relationsMetadata[relation];

  if (!relationMetadata?.module) {
    return typeDef?.metadata?.module || undefined;
  }

  return relationMetadata.module;
}

/**
 * isRelationAssignable - returns true if the relation is assignable, as in the relation definition has a key "this" or any of its children have a key "this".
 * @param relDef - A Userset object representing a relation definition.
 * @return boolean - A boolean representing whether the relation definition has a key "this".
 */
export function isRelationAssignable(relDef: Userset): boolean {
  for (const key of Object.keys(relDef)) {
    const val = relDef[key as keyof Userset];
    if (key === "this") {
      return true;
    }

    if (key === "union" || key === "intersection") {
      for (const item of (val as Usersets).child) {
        if (isRelationAssignable(item)) {
          return true;
        }
      }
    }

    if (key === "difference") {
      if (isRelationAssignable((val as Difference).base) || isRelationAssignable((val as Difference).subtract)) {
        return true;
      }
    }
  }

  // ComputedUserset and TupleToUserset are not assignable
  return false;
}

/**
 * isModelModular - returns true if the model is modular.
 * A model is modular if it has schema version 1.2 and has at least one relation or object that has a module defined in its metadata.
 * @param model - An AuthorizationModel object.
 * @return boolean - A boolean representing whether the model is modular.
 * @throws error - An error if the model's schema version is not recognized.
 */
export function isModelModular(model: AuthorizationModel): boolean {
  const schemaVersion = getSchemaVersionFromString(model.schema_version);
  if (!checkSchemaVersionSupportsModules(schemaVersion)) {
    return false;
  }

  // Check if any type definition has a module defined
  for (const typeDef of model.type_definitions || []) {
    // Check if the type itself has a module
    if (typeDef.metadata?.module) {
      return true;
    }

    // Check if any relation has a module defined
    const relationsMetadata = typeDef.metadata?.relations || {};
    for (const relationMetadata of Object.values(relationsMetadata)) {
      if (relationMetadata.module) {
        return true;
      }
    }
  }

  return false;
}
