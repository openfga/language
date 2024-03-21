import type {
  AuthorizationModel,
  Condition,
  ConditionMetadata,
  ConditionParamTypeRef,
  Metadata,
  RelationMetadata,
  RelationReference,
  TypeDefinition,
  Userset,
} from "@openfga/sdk";
import { ConditionNameDoesntMatchError, UnsupportedDSLNestingError } from "../errors";

class DirectAssignmentValidator {
  occured: number = 0;

  isFirstPosition = (userset: Userset): boolean => {
    // Throw error if direct assignment is present, and not the first element.
    if (userset.this) {
      return true;
    }

    if (userset.difference?.base) {
      if (userset.difference.base.this) {
        return true;
      } else {
        return this.isFirstPosition(userset.difference.base);
      }
    } else if (userset.intersection?.child.length) {
      if (userset.intersection.child[0].this) {
        return true;
      } else {
        return this.isFirstPosition(userset.intersection.child[0]);
      }
    } else if (userset.union?.child.length) {
      if (userset.union.child[0].this) {
        return true;
      } else {
        return this.isFirstPosition(userset.union.child[0]);
      }
    }

    return false;
  };
}

function parseTypeRestriction(restriction: RelationReference): string {
  const typeName = restriction.type;
  const relation = restriction.relation;
  const wildcard = restriction.wildcard;
  const condition = restriction.condition;

  let typeRestriction = typeName;
  if (wildcard) {
    typeRestriction = `${typeRestriction}:*`;
  }

  if (relation) {
    typeRestriction = `${typeRestriction}#${relation}`;
  }

  if (condition) {
    typeRestriction = `${typeRestriction} with ${condition}`;
  }

  return typeRestriction;
}

function parseTypeRestrictions(restrictions: RelationReference[]): string[] {
  const parsedTypeRestrictions: string[] = [];
  for (let index = 0; index < restrictions?.length; index++) {
    parsedTypeRestrictions.push(parseTypeRestriction(restrictions[index]));
  }

  return parsedTypeRestrictions;
}

function parseThis(typeRestrictions: RelationReference[]): string {
  const parsedTypeRestrictions = parseTypeRestrictions(typeRestrictions);
  return `[${parsedTypeRestrictions.join(", ")}]`;
}

function parseTupleToUserset(relationDefinition: Userset): string {
  const computedUserset = relationDefinition?.tupleToUserset?.computedUserset?.relation;
  const tupleset = relationDefinition?.tupleToUserset?.tupleset?.relation;
  return `${computedUserset} from ${tupleset}`;
}

function parseComputedUserset(relationDefinition: Userset): string {
  return relationDefinition!.computedUserset!.relation!;
}

function parseDifference(
  typeName: string,
  relationName: string,
  relationDefinition: Userset,
  typeRestrictions: RelationReference[],
  validator: DirectAssignmentValidator,
): string {
  const base = parseSubRelation(
    typeName,
    relationName,
    relationDefinition!.difference!.base!,
    typeRestrictions,
    validator,
  );
  const difference = parseSubRelation(
    typeName,
    relationName,
    relationDefinition!.difference!.subtract!,
    typeRestrictions,
    validator,
  );
  return `${base} but not ${difference}`;
}

function parseUnion(
  typeName: string,
  relationName: string,
  relationDefinition: Userset,
  typeRestrictions: RelationReference[],
  validator: DirectAssignmentValidator,
): string {
  const parsedString: string[] = [];
  const children = prioritizeDirectAssignment(relationDefinition?.union?.child);

  for (const child of children || []) {
    parsedString.push(parseSubRelation(typeName, relationName, child, typeRestrictions, validator));
  }

  return parsedString.join(" or ");
}

function parseIntersection(
  typeName: string,
  relationName: string,
  relationDefinition: Userset,
  typeRestrictions: RelationReference[],
  validator: DirectAssignmentValidator,
): string {
  const parsedString: string[] = [];
  const children = prioritizeDirectAssignment(relationDefinition?.intersection?.child);

  for (const child of children || []) {
    parsedString.push(parseSubRelation(typeName, relationName, child, typeRestrictions, validator));
  }

  return parsedString.join(" and ");
}

function parseSubRelation(
  typeName: string,
  relationName: string,
  relationDefinition: Userset,
  typeRestrictions: RelationReference[],
  validator: DirectAssignmentValidator,
): string {
  if (relationDefinition.this) {
    // Make sure we have no more than 1 reference for direct assignment in a given relation
    validator.occured++;
    return parseThis(typeRestrictions);
  }

  if (relationDefinition.computedUserset) {
    return parseComputedUserset(relationDefinition);
  }

  if (relationDefinition.tupleToUserset) {
    return parseTupleToUserset(relationDefinition);
  }

  if (relationDefinition.union) {
    return `(${parseUnion(typeName, relationName, relationDefinition, typeRestrictions, validator)})`;
  }

  if (relationDefinition.intersection) {
    return `(${parseIntersection(typeName, relationName, relationDefinition, typeRestrictions, validator)})`;
  }

  if (relationDefinition.difference) {
    return `(${parseDifference(typeName, relationName, relationDefinition, typeRestrictions, validator)})`;
  }

  throw new UnsupportedDSLNestingError(typeName, relationName);
}

function parseRelation(
  typeName: string,
  relationName: string,
  relationDefinition: Userset = {},
  relationMetadata: RelationMetadata = {},
  includeSourceInformation = false,
) {
  const validator = new DirectAssignmentValidator();

  let parsedRelationString = `    define ${relationName}: `;
  const typeRestrictions: RelationReference[] = relationMetadata.directly_related_user_types || [];

  if (relationDefinition.difference != null) {
    parsedRelationString += parseDifference(typeName, relationName, relationDefinition, typeRestrictions, validator);
  } else if (relationDefinition.union != null) {
    parsedRelationString += parseUnion(typeName, relationName, relationDefinition, typeRestrictions, validator);
  } else if (relationDefinition.intersection != null) {
    parsedRelationString += parseIntersection(typeName, relationName, relationDefinition, typeRestrictions, validator);
  } else {
    parsedRelationString += parseSubRelation(typeName, relationName, relationDefinition, typeRestrictions, validator);
  }

  parsedRelationString += constructSourceComment(relationMetadata, " extended by:", includeSourceInformation);

  // Check if we have either no direct assignment, or we had exactly 1 direct assignment in the first position
  if (!validator.occured || (validator.occured === 1 && validator.isFirstPosition(relationDefinition))) {
    return parsedRelationString;
  }

  throw new Error(
    `the '${relationName}' relation definition under the '${typeName}' type is not supported by the OpenFGA DSL syntax yet`,
  );
}

const prioritizeDirectAssignment = (usersets: Userset[] | undefined): Userset[] | undefined => {
  if (usersets?.length) {
    const thisPosition = usersets.findIndex((userset) => userset.this);
    if (thisPosition > 0) {
      usersets.unshift(...usersets.splice(thisPosition, 1));
    }
  }

  return usersets;
};

const parseType = (typeDef: TypeDefinition, isModularModel: boolean, includeSourceInformation = false): string => {
  const typeName = typeDef.type;

  const sourceString = constructSourceComment(typeDef.metadata, "", includeSourceInformation);
  let parsedTypeString = `\ntype ${typeName}${sourceString}`;

  const relations = typeDef.relations || {};
  const metadata = typeDef.metadata;

  if (Object.keys(relations)?.length) {
    parsedTypeString += "\n  relations";
    const sortedRelations = Object.entries(relations).sort(([aName], [bName]) => {
      if (!isModularModel) {
        return 0;
      }
      const aMetadata = metadata?.relations?.[aName] || {};
      const bMetadata = metadata?.relations?.[bName] || {};

      return sortByModule(aName, bName, aMetadata, bMetadata);
    });
    for (const [name, definition] of sortedRelations) {
      const parsedRelationString = parseRelation(
        typeName,
        name,
        definition,
        metadata?.relations?.[name],
        includeSourceInformation,
      );
      parsedTypeString += `\n${parsedRelationString}`;
    }
  }

  return parsedTypeString;
};

const parseConditionParams = (parameterMap: Record<string, ConditionParamTypeRef>): string => {
  const parametersStringArray: string[] = [];

  Object.keys(parameterMap)
    .sort()
    .forEach((parameterName) => {
      const parameterType = parameterMap[parameterName];
      let parameterTypeString = parameterType.type_name.replace("TYPE_NAME_", "").toLowerCase();
      if (parameterTypeString === "list" || parameterTypeString === "map") {
        const genericTypeString = parameterType.generic_types?.[0].type_name.replace("TYPE_NAME_", "").toLowerCase();
        parameterTypeString = `${parameterTypeString}<${genericTypeString}>`;
      }
      parametersStringArray.push(`${parameterName}: ${parameterTypeString}`);
    });

  return parametersStringArray.join(", ");
};

const parseCondition = (conditionName: string, conditionDef: Condition, includeSourceInformation = false): string => {
  if (conditionName != conditionDef.name) {
    throw new ConditionNameDoesntMatchError(conditionName, conditionDef.name);
  }

  const paramsString = parseConditionParams(conditionDef.parameters || {});
  const sourceString = constructSourceComment(conditionDef.metadata, "", includeSourceInformation);

  return `condition ${conditionName}(${paramsString}) {\n  ${conditionDef.expression}\n}${sourceString}\n`;
};

const parseConditions = (
  model: Omit<AuthorizationModel, "id">,
  isModularModel: boolean,
  includeSourceInformation = false,
): string => {
  const conditionsMap = model.conditions || {};
  if (!Object.keys(conditionsMap).length) {
    return "";
  }

  let parsedConditionsString = "";

  Object.entries(conditionsMap)
    .sort(([aName, aCondition], [bName, bCondition]) => {
      if (!isModularModel) {
        return aName.localeCompare(bName);
      }
      return sortByModule(aName, bName, aCondition.metadata, bCondition.metadata);
    })
    .forEach(([conditionName, condition]) => {
      const parsedConditionString = parseCondition(conditionName, condition, includeSourceInformation);

      parsedConditionsString += `\n${parsedConditionString}`;
    });

  return parsedConditionsString;
};

const constructSourceComment = (
  metadata?: ConditionMetadata | Metadata,
  leadingString = "",
  includeSourceInformation = false,
): string => {
  return metadata?.module && includeSourceInformation
    ? ` #${leadingString} module: ${metadata.module}, file: ${metadata.source_info?.file}`
    : "";
};

/**
 * Configuration options for printing the DSL.
 */
export interface TransformOptions {
  /**
   * If true, comments are appended to types, relations, and conditions with file and module information.
   */
  includeSourceInformation?: boolean;
}

export const transformJSONToDSL = (model: Omit<AuthorizationModel, "id">, options?: TransformOptions): string => {
  const schemaVersion = model?.schema_version || "1.1";
  const isModularModel = model.type_definitions?.some((typeDef) => typeDef.metadata?.module);

  const typeDefinitions = (
    isModularModel
      ? model?.type_definitions.sort((a, b) => sortByModule(a.type, b.type, a.metadata, b.metadata))
      : model?.type_definitions
  )?.map((typeDef) => parseType(typeDef, isModularModel, options?.includeSourceInformation));
  const parsedConditionsString = parseConditions(model, isModularModel, options?.includeSourceInformation);

  return `model
  schema ${schemaVersion}
${typeDefinitions ? `${typeDefinitions.join("\n")}\n` : ""}${parsedConditionsString}`;
};

export const transformJSONStringToDSL = (modelString: string, options?: TransformOptions): string => {
  const model = JSON.parse(modelString);

  return transformJSONToDSL(model, options);
};

function sortByModule(aName: string, bName: string, aMeta?: Metadata, bMeta?: Metadata) {
  // If we have no module information for both, sort by name
  if (!aMeta?.module && !bMeta?.module) {
    return aName.localeCompare(bName);
  }
  // If there is no module then it belongs to the same file as the type so sort it at the top
  if (aMeta?.module == undefined) {
    return -1;
  }

  if (bMeta?.module === undefined) {
    return 1;
  }

  // First we sort by module name
  if (aMeta.module !== bMeta.module) {
    return aMeta.module!.localeCompare(bMeta.module!);
  }

  // If the module name is the same then sort by file name
  if (aMeta.source_info?.file !== bMeta.source_info?.file) {
    return aMeta.source_info!.file!.localeCompare(bMeta.source_info!.file!);
  }

  // If the module name and file name are the same then sort based on name
  return aName.localeCompare(bName);
}
