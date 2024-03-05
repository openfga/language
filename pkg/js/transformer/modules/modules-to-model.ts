import { AuthorizationModel, Condition, TypeDefinition } from "@openfga/sdk";
import { validateJSON } from "../../validator";
import { transformModularDSLToJSONObject } from "../dsltojson";
import {
  BaseError,
  DSLSyntaxError,
  ModelValidationError,
  ModuleTransformationError,
  ModuleTransformationSingleError
} from "../../errors";

export interface ModuleFiles {
  name: string;
  contents: string;
}

export const transformModuleFilesToModel = (files: ModuleFiles[]): Omit<AuthorizationModel, "id"> => {
  const model: Omit<AuthorizationModel, "id"> = {
    schema_version: "1.2",
    type_definitions: [],
    conditions: {}
  };

  const typeDefs: TypeDefinition[] = [];
  const types = new Set<string>();
  const extendedTypeDefs: Record<string, TypeDefinition[]> = {};
  const conditions = new Map<string, Condition>();
  const errors: BaseError[] = [];

  for (const { name, contents } of files) {
    try {
      const { authorizationModel, typeDefExtensions } = transformModularDSLToJSONObject(contents);

      for (const typeDef of authorizationModel.type_definitions) {
        // Check if we've already seen this type and it's not marked as an extension
        if (types.has(typeDef.type) && !typeDefExtensions.has(typeDef.type)) {
          errors.push(new ModuleTransformationSingleError({
            msg: `duplicate type definition ${typeDef.type}`
          }));
          continue;
        }

        // If this is an extension mark it to be merged later
        if (typeDefExtensions.has(typeDef.type)) {
          if (!extendedTypeDefs[name]) {
            extendedTypeDefs[name] = [];
          }
          extendedTypeDefs[name].push(typeDef);
          continue;
        }

        types.add(typeDef.type);
        typeDefs.push({
          ...typeDef,
          metadata: {
            ...typeDef.metadata,
            file: name,
          }
        });
      }

      if (authorizationModel.conditions) {
        for (const [conditionName, condition] of Object.entries(authorizationModel.conditions)) {
          // If we have already seen a condition with this name mark it as duplicate
          if (conditions.has(conditionName)) {
            errors.push(new ModuleTransformationSingleError({
              msg: `duplicate condition ${conditionName}`
            }));
            continue;
          }
          conditions.set(conditionName, {
            ...condition as Condition,
            metadata: {
              ...(condition as Condition).metadata,
              file: name,
            }
          });
        }
      }
    } catch (error) {
      if (error instanceof DSLSyntaxError) {
        errors.push(...error.errors);
      } else if (error instanceof Error) {
        errors.push(error as BaseError);
      }
    }
  }

  for (const [filename, extended] of Object.entries(extendedTypeDefs)) {
    for (const typeDef of extended) {
      const originalIndex = typeDefs.findIndex((t) => t.type === typeDef.type);
      const original = typeDefs[originalIndex];

      if (!original) {
        errors.push(new ModuleTransformationSingleError({
          msg: `extended type ${typeDef.type} does not exist`
        }));
        continue;
      }

      if (!original.relations || !Object.keys(original.relations).length) {
        original.relations = typeDef.relations;
        if (!original.metadata) {
          original.metadata = {};
        }
        original.metadata.relations = typeDef.metadata!.relations;

        // Add the file metadata to any relations metadata that exists
        for (const relationName of Object.keys(original.metadata.relations!)) {
          original.metadata.relations![relationName].file = filename;
        }

        typeDefs[originalIndex] = original;
        continue;
      }

      const existingRelationNames = Object.keys(original.relations);

      for (const [name, relation] of Object.entries(typeDef.relations || {})) {
        if (existingRelationNames.includes(name)) {
          errors.push(new ModuleTransformationSingleError({
            msg: `relation ${name} already exists on type ${typeDef.type}`
          }));
          continue;
        }

        const relationsMeta = Object.entries(typeDef.metadata?.relations || {}).find(([n]) => n === name);

        if (!relationsMeta) {
          errors.push(new ModuleTransformationSingleError({
            msg: `unable to find relation metadata for ${name}`
          }));
          continue;
        }

        const [, meta] = relationsMeta;
        meta.file = filename;
        original.relations[name] = relation;
        original.metadata!.relations![name] = meta;
      }

      typeDefs[originalIndex] = original;
    }
  }

  model.type_definitions = Array.from(typeDefs);
  model.conditions = Object.fromEntries(conditions);

  try {
    validateJSON(model as AuthorizationModel);
  } catch (error) {
    if (error instanceof ModelValidationError) {
      errors.push(...error.errors);
    }
  }

  if (errors.length) {
    throw new ModuleTransformationError(errors);
  }

  return model;
};
