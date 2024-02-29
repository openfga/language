import { AuthorizationModel as OGAuthorizationModel, Condition as OGCondition, TypeDefinition as OGTypeDefinition, Metadata as OGMetadata } from "@openfga/sdk";
import { validateJSON } from "../../validator";
import { transformDSLToJSONObject } from "../dsltojson";
import { BaseError, DSLSyntaxError, ModelValidationError, ModuleTransformationError, ModuleTransformationSingleError } from "../../errors";

export interface ModuleFiles {
  name: string;
  contents: string;
}

// TODO: need to move these out once we have them added to the API
interface Condition extends OGCondition {
  metadata: {
    file?: string;
    module?: string
  }
}

interface Metadata extends OGMetadata {
  file?: string;
  module?: string
}

interface TypeDefinition extends OGTypeDefinition {
  metadata: Metadata
}

interface AuthorizationModel extends OGAuthorizationModel {
  name?: string
}

export const transformModuleFilesToModel = (files: ModuleFiles[]): Omit<AuthorizationModel, "id"> => {
  // Build all the individual models
  // Validate them (no conflicting names that aren't extensions?)
  // Copy over any extensions?
  // Stitch them into the main model
  const model: Omit<AuthorizationModel, "id"> = {
    schema_version: "1.2",
    type_definitions: [],
    conditions: {}
  };

  const types: TypeDefinition[] = [];
  const conditions = new Map<string, Condition>();
  const errors: BaseError[] = [];

  for (const { name, contents } of files) {
    try {
      const model = transformDSLToJSONObject(contents) as AuthorizationModel;

      types.push(...model.type_definitions.map((type) => ({
        ...type,
        metadata: {
          ...type.metadata,
          file: name,
        }
      })));

      if (model.conditions) {
        for (const [conditionName, condition] of Object.entries(model.conditions)) {
          if (conditions.has(conditionName)) {
            errors.push(new ModuleTransformationSingleError({
              msg: `duplicate condition ${conditionName}`
            }));
          }
          conditions.set(conditionName, {
            ...condition as Condition,
            metadata: {
              ...(condition as Condition).metadata,
              file: conditionName,
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

  model.type_definitions = types;
  model.conditions = Object.fromEntries(conditions);

  try {
    validateJSON(model as AuthorizationModel);
  } catch (error) {
    if (error instanceof ModelValidationError) {
      errors.push(...error.errors);
    }
  }

  if (errors.length !== 0) {
    throw new ModuleTransformationError(errors);
  }


  return model;
};
