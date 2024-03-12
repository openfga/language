import { AuthorizationModel, Condition , TypeDefinition } from "@openfga/sdk";
import { validateJSON } from "../../validator";
import { transformModularDSLToJSONObject } from "../dsltojson";
import {
  BaseError,
  DSLSyntaxError,
  ModelValidationError,
  ModuleTransformationError,
  ModuleTransformationSingleError
} from "../../errors";
import { getTypeLineNumber, getConditionLineNumber, getRelationLineNumber } from "../../util/line-numbers";
import { constructTransformationError } from "../../util/exceptions";

export interface ModuleFile {
  name: string;
  contents: string;
}

export const transformModuleFilesToModel = (
  files: ModuleFile[],
  schemaVersion: string
): Omit<AuthorizationModel, "id"> => {
  const model: Omit<AuthorizationModel, "id"> = {
    schema_version: schemaVersion,
    type_definitions: [],
    conditions: {}
  };

  const typeDefs: TypeDefinition[] = [];
  const types = new Set<string>();
  const extendedTypeDefs: Record<string, TypeDefinition[]> = {};
  const conditions = new Map<string, Condition>();
  const errors: BaseError[] = [];
  const moduleFiles = new Map(files.map(file => [file.name, file.contents]));

  for (const { name, contents } of files) {
    try {
      const lines = contents.split("\n");
      const { authorizationModel, typeDefExtensions } = transformModularDSLToJSONObject(contents);

      for (const typeDef of authorizationModel.type_definitions) {
        // Check if we've already seen this type and it's not marked as an extension
        if (types.has(typeDef.type) && !typeDefExtensions.has(typeDef.type)) {
          const lineIndex = getTypeLineNumber(typeDef.type, lines);
          errors.push(constructTransformationError({
            message: `duplicate type definition ${typeDef.type}`,
            lines, lineIndex,
            metadata: {
              symbol: typeDef.type,
              file: name,
            }
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

        typeDef.metadata!.file = name;
        types.add(typeDef.type);
        typeDefs.push(typeDef);
      }

      if (authorizationModel.conditions) {
        for (const [conditionName, condition] of Object.entries(authorizationModel.conditions)) {
          // If we have already seen a condition with this name mark it as duplicate
          if (conditions.has(conditionName)) {
            const lineIndex = getConditionLineNumber(conditionName, lines);
            errors.push(constructTransformationError({
              message: `duplicate condition ${conditionName}`,
              lines, lineIndex,
              metadata: {
                symbol: conditionName,
                file: name,
              }
            }));
            continue;
          }
          condition.metadata!.file = name;
          conditions.set(conditionName, condition);
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
    const lines = moduleFiles.get(filename)?.split("\n");
    for (const typeDef of extended) {

      if (!typeDef.relations) {
        // TODO: Maybe should be an error case or at least a warning?
        continue;
      }

      const originalIndex = typeDefs.findIndex((t) => t.type === typeDef.type);
      const original = typeDefs[originalIndex];

      if (!original) {
        const lineIndex = getTypeLineNumber(typeDef.type, lines, 0, true);
        errors.push(constructTransformationError({
          message: `extended type ${typeDef.type} does not exist`,
          lines, lineIndex,
          metadata: {
            symbol: typeDef.type,
            file: filename
          }
        }));
        continue;
      }

      const existingRelationNames = Object.keys(original.relations || {});

      if (!existingRelationNames || !existingRelationNames.length) {
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

      for (const [name, relation] of Object.entries(typeDef.relations)) {
        if (existingRelationNames.includes(name)) {
          const lineIndex = getRelationLineNumber(name, lines);
          errors.push(constructTransformationError({
            message: `relation ${name} already exists on type ${typeDef.type}`,
            lines, lineIndex,
            metadata: {
              symbol: name,
              file: filename,
            }
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
        original.relations![name] = relation;
        original.metadata!.relations![name] = meta;
      }

      typeDefs[originalIndex] = original;
    }
  }

  model.type_definitions = typeDefs;
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
