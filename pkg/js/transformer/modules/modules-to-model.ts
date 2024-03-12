import { AuthorizationModel, Condition , TypeDefinition } from "@openfga/sdk";
import { validateJSON } from "../../validator";
import { transformModularDSLToJSONObject } from "../dsltojson";
import {
  BaseError,
  DSLSyntaxError,
  ModelValidationError,
  ModelValidationSingleError,
  ModuleTransformationError,
  ModuleTransformationSingleError,
  ValidationError
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

        typeDef.metadata!.source_info = {
          file: name
        };
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
          condition.metadata!.source_info = {
            file: name
          };
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
          original.metadata.relations![relationName].source_info = {
            file: filename
          };
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
        meta.source_info = {
          file: filename
        };
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
      for (const e of error.errors) {
        if (!e.file || !e.metadata?.module || !e.metadata.symbol) {
          errors.push(e);
          continue;
        }

        const lines = moduleFiles.get(e.file)?.split("\n");
        if (!lines) {
          errors.push(e);
          continue;
        }

        const lineIndex = resolveLineIndex(e, lines);

        if (lineIndex === -1) {
          errors.push(e);
          continue;
        }

        const line = lines[lineIndex];
        const wordIndex = resolveWordIndex(e, line);

        e.line = { start: lineIndex + 1, end: lineIndex + 1 };
        e.column = { start: wordIndex, end: wordIndex + (e.metadata.symbol?.length || 0) };
        errors.push(e);
      }

    }
  }

  if (errors.length) {
    throw new ModuleTransformationError(errors);
  }

  return model;
};

function resolveLineIndex (e: ModelValidationSingleError, lines?: string[]): number {
  const { metadata } = e;
  let lineIndex;

  if (!metadata || !lines) {
    return -1;
  }

  switch(metadata?.errorType) {
    case ValidationError.ReservedTypeKeywords:
      lineIndex = getTypeLineNumber(metadata.symbol, lines);
     break;
    case ValidationError.InvalidName:
      // handle type vs relation, this isn't ideal but no other way to check
      if (e.message.startsWith("invalid-name error: relation")) {
        lineIndex =  getRelationLineNumber(metadata.symbol, lines);
      } else {
        lineIndex = getTypeLineNumber(metadata.symbol, lines);
      }
      break;
    case ValidationError.ReservedRelationKeywords:
      lineIndex =  getRelationLineNumber(metadata.symbol, lines);
      break;
  }

  if (lineIndex === undefined) {
    return -1;
  }

  return lineIndex;
}

function resolveWordIndex (e: ModelValidationSingleError, line: string): number {
  const { metadata } = e;

  if (!metadata) {
    return -1;
  }

  const re = new RegExp("\\b" + metadata.symbol + "\\b");
  let wordIdx = line?.search(re) + 1;
  if (wordIdx == undefined || isNaN(wordIdx) || wordIdx === 0) {
    wordIdx = 1;
  }

  return wordIdx;
}