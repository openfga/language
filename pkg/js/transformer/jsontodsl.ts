import type {
  AuthorizationModel,
  Condition,
  ConditionParamTypeRef,
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
    validator
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
  const children = relationDefinition?.union?.child;

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
  const children = relationDefinition?.intersection?.child;

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
  if (relationDefinition.this != null) {
    // Make sure we have no more than 1 reference for direct assignment in a given relation
    validator.occured++;
    return parseThis(typeRestrictions);
  }

  if (relationDefinition.computedUserset != null) {
    return parseComputedUserset(relationDefinition);
  }

  if (relationDefinition.tupleToUserset != null) {
    return parseTupleToUserset(relationDefinition);
  }

  if (relationDefinition.union != null) {
    return `(${parseUnion(typeName, relationName, relationDefinition, typeRestrictions, validator)})`;
  }

  if (relationDefinition.intersection != null) {
    return `(${parseIntersection(typeName, relationName, relationDefinition, typeRestrictions, validator)})`;
  }

  if (relationDefinition.difference != null) {
    return `(${parseDifference(typeName, relationName, relationDefinition, typeRestrictions, validator)})`;
  }

  throw new UnsupportedDSLNestingError(typeName, relationName);
}

function parseRelation(
  typeName: string,
  relationName: string,
  relationDefinition: Userset = {},
  relationMetadata: RelationMetadata = {},
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

  // Check if we have either no direct assignment, or we had exactly 1 direct assignment in the first position
  if (!validator.occured || (validator.occured === 1 && validator.isFirstPosition(relationDefinition))) {
    return parsedRelationString;
  }

  throw new Error(
    `the '${relationName}' relation definition under the '${typeName}' type is not supported by the OpenFGA DSL syntax yet`,
  );
}

const parseType = (typeDef: TypeDefinition): string => {
  const typeName = typeDef.type;
  let parsedTypeString = `\ntype ${typeName}`;

  const relations = typeDef.relations || {};
  const metadata = typeDef.metadata;

  if (Object.keys(relations)?.length) {
    parsedTypeString += "\n  relations";
    for (const relationName in relations) {
      const parsedRelationString = parseRelation(
        typeName,
        relationName,
        relations[relationName],
        metadata?.relations?.[relationName],
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

const parseCondition = (conditionName: string, conditionDef: Condition): string => {
  if (conditionName != conditionDef.name) {
    throw new ConditionNameDoesntMatchError(conditionName, conditionDef.name);
  }

  const paramsString = parseConditionParams(conditionDef.parameters || {});

  return `condition ${conditionName}(${paramsString}) {\n  ${conditionDef.expression}\n}\n`;
};

const parseConditions = (model: Omit<AuthorizationModel, "id">): string => {
  const conditionsMap = model.conditions || {};
  if (!Object.keys(conditionsMap).length) {
    return "";
  }

  let parsedConditionsString = "";
  Object.keys(conditionsMap)
    .sort()
    .forEach((conditionName) => {
      const condition = conditionsMap[conditionName];
      const parsedConditionString = parseCondition(conditionName, condition);

      parsedConditionsString += `\n${parsedConditionString}`;
    });

  return parsedConditionsString;
};

export const transformJSONToDSL = (model: Omit<AuthorizationModel, "id">): string => {
  const schemaVersion = model?.schema_version || "1.1";
  const typeDefinitions = model?.type_definitions?.map((typeDef) => parseType(typeDef));
  const parsedConditionsString = parseConditions(model);

  return `model
  schema ${schemaVersion}
${typeDefinitions ? `${typeDefinitions.join("\n")}\n` : ""}${parsedConditionsString}`;
};

export const transformJSONStringToDSL = (modelString: string): string => {
  const model = JSON.parse(modelString);

  return transformJSONToDSL(model);
};
