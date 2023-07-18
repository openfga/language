import { AuthorizationModel, RelationMetadata, RelationReference, TypeDefinition, Userset } from "@openfga/sdk";

function parseTypeRestriction(restriction: RelationReference): string {
  const typeName = restriction.type;
  const relation = restriction.relation;
  const wildcard = restriction.wildcard;

  if (wildcard != null) {
    return `${typeName}:*`;
  }

  if (relation != null) {
    return `${typeName}#${relation}`;
  }

  return typeName;
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
  return relationDefinition?.computedUserset?.relation!;
}

function parseDifference(relationDefinition: Userset, typeRestrictions: RelationReference[]): string {
  const base = parseSubRelation(relationDefinition?.difference?.base!, typeRestrictions);
  const difference = parseSubRelation(relationDefinition?.difference?.subtract!, typeRestrictions);
  return `${base} but not ${difference}`;
}

function parseUnion(relationDefinition: Userset, typeRestrictions: RelationReference[]): string {
  let parsedString: string[] = [];
  const children = relationDefinition?.union?.child;

  for (const child of children || []) {
    parsedString.push(parseSubRelation(child, typeRestrictions));
  }

  return parsedString.join(" or ");
}

function parseIntersection(relationDefinition: Userset, typeRestrictions: RelationReference[]): string {
  let parsedString: string[] = [];
  const children = relationDefinition?.intersection?.child;

  for (const child of children || []) {
    parsedString.push(parseSubRelation(child, typeRestrictions));
  }

  return parsedString.join(" and ");
}

function parseSubRelation(relationDefinition: Userset, typeRestrictions: RelationReference[]): string {
  if (relationDefinition.this != null) {
    return parseThis(typeRestrictions);
  }

  if (relationDefinition.computedUserset != null) {
    return parseComputedUserset(relationDefinition);
  }

  if (relationDefinition.tupleToUserset != null) {
    return parseTupleToUserset(relationDefinition);
  }

  return "";
}

function parseRelation(
  relationName: string,
  relationDefinition: Userset = {},
  relationMetadata: RelationMetadata = {},
) {
  let parsedRelationString = `    define ${relationName}: `;
  const typeRestrictions = relationMetadata.directly_related_user_types || [];

  if (relationDefinition.difference != null) {
    parsedRelationString += parseDifference(relationDefinition, typeRestrictions);
  } else if (relationDefinition.union != null) {
    parsedRelationString += parseUnion(relationDefinition, typeRestrictions);
  } else if (relationDefinition.intersection != null) {
    parsedRelationString += parseIntersection(relationDefinition, typeRestrictions);
  } else {
    parsedRelationString += parseSubRelation(relationDefinition, typeRestrictions);
  }

  return parsedRelationString;
}

const parseType = (typeDef: TypeDefinition): string => {
  const typeName = typeDef.type;
  let parsedTypeString = `\ntype ${typeName}`;

  const relations = typeDef.relations || {};
  const metadata = typeDef.metadata;

  if (Object.keys(relations)?.length) {
    parsedTypeString += "\n  relations";
    for (const relation in relations) {
      const parsedRelationString = parseRelation(relation, relations[relation], metadata?.relations?.[relation]);
      parsedTypeString += `\n${parsedRelationString}`;
    }
  }

  return parsedTypeString;
};

export const transformJSONToDSL = (model: AuthorizationModel): string => {
  const schemaVersion = model?.schema_version || "1.1";
  const typeDefinitions = model?.type_definitions?.map((typeDef) => parseType(typeDef));

  return `model
  schema ${schemaVersion}
${typeDefinitions ? `${typeDefinitions.join("\n")}\n` : ""}`;
};

export const transformJSONStringToDSL = (modelString: string): string => {
  const model = JSON.parse(modelString);

  return transformJSONToDSL(model);
};
