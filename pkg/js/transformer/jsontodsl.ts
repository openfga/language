import type { AuthorizationModel, RelationMetadata, RelationReference, TypeDefinition, Userset } from "@openfga/sdk";
import { UnsupportedDSLNestingError } from "../errors";

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
  return relationDefinition!.computedUserset!.relation!;
}

function parseDifference(
  typeName: string,
  relationName: string,
  relationDefinition: Userset,
  typeRestrictions: RelationReference[],
): string {
  const base = parseSubRelation(typeName, relationName, relationDefinition!.difference!.base!, typeRestrictions);
  const difference = parseSubRelation(
    typeName,
    relationName,
    relationDefinition!.difference!.subtract!,
    typeRestrictions,
  );
  return `${base} but not ${difference}`;
}

function parseUnion(
  typeName: string,
  relationName: string,
  relationDefinition: Userset,
  typeRestrictions: RelationReference[],
): string {
  const parsedString: string[] = [];
  const children = relationDefinition?.union?.child;

  for (const child of children || []) {
    parsedString.push(parseSubRelation(typeName, relationName, child, typeRestrictions));
  }

  return parsedString.join(" or ");
}

function parseIntersection(
  typeName: string,
  relationName: string,
  relationDefinition: Userset,
  typeRestrictions: RelationReference[],
): string {
  const parsedString: string[] = [];
  const children = relationDefinition?.intersection?.child;

  for (const child of children || []) {
    parsedString.push(parseSubRelation(typeName, relationName, child, typeRestrictions));
  }

  return parsedString.join(" and ");
}

function parseSubRelation(
  typeName: string,
  relationName: string,
  relationDefinition: Userset,
  typeRestrictions: RelationReference[],
): string {
  if (relationDefinition.this != null) {
    return parseThis(typeRestrictions);
  }

  if (relationDefinition.computedUserset != null) {
    return parseComputedUserset(relationDefinition);
  }

  if (relationDefinition.tupleToUserset != null) {
    return parseTupleToUserset(relationDefinition);
  }

  throw new UnsupportedDSLNestingError(typeName, relationName);
}

function parseRelation(
  typeName: string,
  relationName: string,
  relationDefinition: Userset = {},
  relationMetadata: RelationMetadata = {},
) {
  let parsedRelationString = `    define ${relationName}: `;
  const typeRestrictions = relationMetadata.directly_related_user_types || [];

  if (relationDefinition.difference != null) {
    parsedRelationString += parseDifference(typeName, relationName, relationDefinition, typeRestrictions);
  } else if (relationDefinition.union != null) {
    parsedRelationString += parseUnion(typeName, relationName, relationDefinition, typeRestrictions);
  } else if (relationDefinition.intersection != null) {
    parsedRelationString += parseIntersection(typeName, relationName, relationDefinition, typeRestrictions);
  } else {
    parsedRelationString += parseSubRelation(typeName, relationName, relationDefinition, typeRestrictions);
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
