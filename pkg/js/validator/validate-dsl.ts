import type { AuthorizationModel, RelationReference, RelationMetadata, TypeDefinition, Userset } from "@openfga/sdk";
import { Keyword, ReservedKeywords } from "./keywords";
import { parseDSL } from "../transformer";
import { ConfigurationError, DSLSyntaxError, ModelValidationError, ModelValidationSingleError } from "../errors";
import { exceptionCollector } from "../util/exceptions";

// eslint-disable-next-line no-useless-escape
export const defaultTypeRule = "^[^:#@\\s]{1,254}$";
// eslint-disable-next-line no-useless-escape
export const defaultRelationRule = "^[^:#@\\s]{1,50}$";

enum RelationDefOperator {
  Union = "union",
  Intersection = "intersection",
  Difference = "difference",
}

export interface ValidationRegex {
  rule: string;
  regex: RegExp;
}

export interface ValidationOptions {
  typeValidation?: string;
  relationValidation?: string;
}

enum RewriteType {
  Direct = "direct",
  ComputedUserset = "computed_userset",
  TupleToUserset = "tuple_to_userset",
}

interface RelationTargetParserResult {
  target?: string;
  from?: string;
  rewrite: RewriteType;
}

const deepCopy = <T>(object: T): T => {
  return JSON.parse(JSON.stringify(object));
};

const geConditionLineNumber = (conditionName: string, lines: string[], skipIndex?: number) => {
  if (!skipIndex) {
    skipIndex = 0;
  }
  return (
    lines.slice(skipIndex).findIndex((line: string) => line.trim().startsWith(`condition ${conditionName}`)) + skipIndex
  );
};

const getTypeLineNumber = (typeName: string, lines: string[], skipIndex?: number) => {
  if (!skipIndex) {
    skipIndex = 0;
  }
  return lines.slice(skipIndex).findIndex((line: string) => line.trim().startsWith(`type ${typeName}`)) + skipIndex;
};

const getRelationLineNumber = (relation: string, lines: string[], skipIndex?: number) => {
  if (!skipIndex) {
    skipIndex = 0;
  }
  return (
    lines
      .slice(skipIndex)
      .findIndex((line: string) => line.trim().replace(/ {2,}/g, " ").startsWith(`define ${relation}`)) + skipIndex
  );
};

const getSchemaLineNumber = (schema: string, lines: string[]) => {
  return lines.findIndex((line: string) => line.trim().replace(/ {2,}/g, " ").startsWith(`schema ${schema}`));
};

const getTypeRestrictionString = (typeRestriction: RelationReference): string => {
  let typeRestrictionString = typeRestriction.type;
  if (typeRestriction.wildcard) {
    typeRestrictionString += ":*";
  } else if (typeRestriction.relation) {
    typeRestrictionString += `#${typeRestriction.relation}`;
  }

  if ((typeRestriction as RelationReference).condition) {
    typeRestrictionString += ` with ${(typeRestriction as RelationReference).condition}`;
  }

  return typeRestrictionString;
};

const getTypeRestrictions = (relatedTypes: Array<RelationReference>): string[] => {
  return relatedTypes.map((u) => getTypeRestrictionString(u));
};

const getRelationalParserResult = (userset: Userset): RelationTargetParserResult => {
  let target,
    from = undefined;

  if (userset.computedUserset) {
    target = userset.computedUserset.relation || undefined;
  } else {
    target = userset.tupleToUserset?.computedUserset?.relation || undefined;
    from = userset.tupleToUserset?.tupleset?.relation || undefined;
  }

  let rewrite = RewriteType.Direct;
  if (target) {
    rewrite = RewriteType.ComputedUserset;
  }
  if (from) {
    rewrite = RewriteType.TupleToUserset;
  }
  return { target, from, rewrite };
};

interface DestructedAssignableType {
  decodedType: string;
  decodedRelation?: string;
  isWildcard: boolean;
  decodedConditionName?: string;
}

// helper function to figure out whether the specified allowable types
// are tuple to user set.  If so, return the type and relationship.
// Otherwise, return null as relationship
const destructTupleToUserset = (allowableType: string): DestructedAssignableType => {
  const [tupleString, decodedConditionName] = allowableType.split(" with ");
  const isWildcard = tupleString.includes(":*");
  const splittedWords = tupleString.replace(":*", "").split("#");
  return { decodedType: splittedWords[0], decodedRelation: splittedWords[1], isWildcard, decodedConditionName };
};

const relationIsSingle = (currentRelation: Userset): boolean => {
  return (
    !Object.prototype.hasOwnProperty.call(currentRelation, RelationDefOperator.Union) &&
    !Object.prototype.hasOwnProperty.call(currentRelation, RelationDefOperator.Intersection) &&
    !Object.prototype.hasOwnProperty.call(currentRelation, RelationDefOperator.Difference)
  );
};

const getRelationDefName = (userset: Userset): string | undefined => {
  let relationDefName = userset.computedUserset?.relation;

  const parserResult = getRelationalParserResult(userset);
  if (parserResult.rewrite === RewriteType.ComputedUserset) {
    relationDefName = parserResult.target;
  } else if (parserResult.rewrite === RewriteType.TupleToUserset) {
    relationDefName = `${parserResult.target} from ${parserResult.from}`;
  }
  return relationDefName;
};

// Return all the allowable types for the specified type/relation
function allowableTypes(typeName: Record<string, TypeDefinition>, type: string, relation: string): [string[], boolean] {
  const allowedTypes: string[] = [];
  const currentRelations = typeName[type].relations![relation];
  const currentRelationMetadata = getTypeRestrictions(
    typeName[type].metadata?.relations![relation].directly_related_user_types || [],
  );

  const isValid = relationIsSingle(currentRelations);
  // for now, we assume that the type/relation must be single and rewrite is direct
  if (isValid) {
    const childDef = getRelationalParserResult(currentRelations);

    switch (childDef.rewrite) {
      case RewriteType.Direct: {
        allowedTypes.push(...currentRelationMetadata);
      }
    }
  }
  return [allowedTypes, isValid];
}

// for the type/relation, whether there are any unique entry points, and if a loop is found
// if there are unique entry points (i.e., direct relations) then it will return true
// otherwise, it will follow its children to see if there are unique entry points
// if there is a loop during traversal, the function will return a boolean indicating so
function hasEntryPointOrLoop(
  typeMap: Record<string, TypeDefinition>,
  typeName: string,
  relationName: string | undefined,
  rewrite: Userset,
  visitedRecords: Record<string, Record<string, boolean>>,
): [boolean, boolean] {
  // Deep copy
  const visited = deepCopy(visitedRecords);

  if (!relationName) {
    // nothing to do if relation is undefined
    return [false, false];
  }

  if (!visited[typeName]) {
    visited[typeName] = {};
  }
  visited[typeName][relationName] = true;

  const currentRelation = typeMap[typeName].relations;
  if (!currentRelation || !currentRelation[relationName]) {
    return [false, false];
  }

  const relationMetadata = typeMap[typeName].metadata?.relations;

  if (!typeMap[typeName].relations || !typeMap[typeName].relations![relationName]) {
    return [false, false];
  }

  if (rewrite.this) {
    for (const assignableType of getTypeRestrictions(
      relationMetadata?.[relationName]?.directly_related_user_types || [],
    )) {
      const { decodedType, decodedRelation, isWildcard } = destructTupleToUserset(assignableType);
      if (!decodedRelation || isWildcard) {
        return [true, false];
      }

      const assignableRelation = typeMap[decodedType].relations![decodedRelation];
      if (!assignableRelation) {
        return [false, false];
      }

      if (visited[decodedType]?.[decodedRelation]) {
        continue;
      }

      const [hasEntry] = hasEntryPointOrLoop(typeMap, decodedType, decodedRelation, assignableRelation, visited);
      if (hasEntry) {
        return [true, false];
      }
    }

    return [false, false];
  } else if (rewrite.computedUserset) {
    const computedRelationName = rewrite.computedUserset.relation;
    if (!computedRelationName) {
      return [false, false];
    }

    if (!typeMap[typeName].relations![computedRelationName]) {
      return [false, false];
    }

    const computedRelation = typeMap[typeName].relations![computedRelationName];
    if (!computedRelation) {
      return [false, false];
    }

    // Loop detected
    if (visited[typeName][computedRelationName]) {
      return [false, true];
    }

    const [hasEntry, loop] = hasEntryPointOrLoop(typeMap, typeName, computedRelationName, computedRelation, visited);
    return [hasEntry, loop];
  } else if (rewrite.tupleToUserset) {
    const tuplesetRelationName = rewrite.tupleToUserset.tupleset.relation;
    const computedRelationName = rewrite.tupleToUserset.computedUserset.relation;

    if (!tuplesetRelationName || !computedRelationName) {
      return [false, false];
    }

    const tuplesetRelation = typeMap[typeName].relations![tuplesetRelationName];
    if (!tuplesetRelation) {
      return [false, false];
    }

    for (const assignableType of getTypeRestrictions(
      relationMetadata?.[tuplesetRelationName]?.directly_related_user_types || [],
    )) {
      const assignableRelation = typeMap[assignableType].relations![computedRelationName];
      if (assignableRelation) {
        if (visited[assignableType] && visited[assignableType][computedRelationName]) {
          continue;
        }

        const [hasEntry] = hasEntryPointOrLoop(
          typeMap,
          assignableType,
          computedRelationName,
          assignableRelation,
          visited,
        );
        if (hasEntry) {
          return [true, false];
        }
      }
    }
    return [false, false];
  } else if (rewrite.union) {
    let loop = false;

    for (const child of rewrite.union.child) {
      const [entryPoint, childLoop] = hasEntryPointOrLoop(typeMap, typeName, relationName, child, deepCopy(visited));
      if (entryPoint) {
        return [true, false];
      }
      loop = loop || childLoop;
    }
    return [false, loop];
  } else if (rewrite.intersection) {
    for (const child of rewrite.intersection.child) {
      const [hasEntry, childLoop] = hasEntryPointOrLoop(typeMap, typeName, relationName, child, deepCopy(visited));
      if (!hasEntry) {
        return [false, childLoop];
      }
    }

    return [true, false];
  } else if (rewrite.difference) {
    const visited = deepCopy(visitedRecords);

    const [hasEntryBase, loopBase] = hasEntryPointOrLoop(
      typeMap,
      typeName,
      relationName,
      rewrite.difference.base,
      visited,
    );
    if (!hasEntryBase) {
      return [false, loopBase];
    }

    const [hasEntrySubtract, loopSubtract] = hasEntryPointOrLoop(
      typeMap,
      typeName,
      relationName,
      rewrite.difference.subtract,
      visited,
    );
    if (!hasEntrySubtract) {
      return [false, loopSubtract];
    }

    return [true, false];
  }

  return [false, false];
}

function checkForDuplicatesTypeNamesInRelation(
  lines: string[],
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  collector: any,
  relationDef: RelationMetadata,
  relationName: string,
) {
  const typeNameSet = new Set();
  relationDef.directly_related_user_types?.forEach((typeDef) => {
    const typeDefName = getTypeRestrictionString(typeDef);

    if (typeNameSet.has(typeDefName)) {
      const typeIndex = getTypeLineNumber(typeDef.type, lines);
      const lineIndex = getRelationLineNumber(relationName, lines, typeIndex);
      collector.raiseDuplicateTypeRestriction(lineIndex, typeDefName, relationName);
    }
    typeNameSet.add(typeDefName);
  });
}

// ensure all the referenced relations are defined
function checkForDuplicatesInRelation(
  lines: string[],
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  collector: any,
  typeDef: TypeDefinition,
  relationName: string,
) {
  const relationDef = typeDef.relations![relationName];

  // Union
  const relationUnionNameSet = new Set();
  relationDef.union?.child?.forEach((userset) => {
    const relationDefName = getRelationDefName(userset);
    if (relationDefName && relationUnionNameSet.has(relationDefName)) {
      const typeIndex = getTypeLineNumber(typeDef.type, lines);
      const lineIndex = getRelationLineNumber(relationName, lines, typeIndex);
      collector.raiseDuplicateType(lineIndex, relationDefName, relationName);
    }
    relationUnionNameSet.add(relationDefName);
  });

  // Intersection
  const relationIntersectionNameSet = new Set();
  relationDef.intersection?.child?.forEach((userset) => {
    const relationDefName = getRelationDefName(userset);
    if (relationDefName && relationIntersectionNameSet.has(relationDefName)) {
      const typeIndex = getTypeLineNumber(typeDef.type, lines);
      const lineIndex = getRelationLineNumber(relationName, lines, typeIndex);
      collector.raiseDuplicateType(lineIndex, relationDefName, relationName);
    }
    relationIntersectionNameSet.add(relationDefName);
  });

  // Difference
  if (Object.prototype.hasOwnProperty.call(relationDef, RelationDefOperator.Difference)) {
    const baseName = getRelationDefName(relationDef.difference!.base);
    const subtractName = getRelationDefName(relationDef.difference!.subtract);
    if (baseName && baseName === subtractName) {
      const typeIndex = getTypeLineNumber(typeDef.type, lines);
      const lineIndex = getRelationLineNumber(relationName, lines, typeIndex);
      collector.raiseDuplicateType(lineIndex, baseName, relationName);
    }
  }
}

// helper function to ensure all childDefs are defined
function childDefDefined(
  lines: string[],
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  collector: any,
  typeMap: Record<string, TypeDefinition>,
  type: string,
  relation: string,
  childDef: RelationTargetParserResult,
  conditions: AuthorizationModel["conditions"] = {},
) {
  const relations = typeMap[type].relations;
  if (!relations || !relations[relation]) {
    return;
  }

  const currentRelationMetadata = typeMap[type].metadata?.relations![relation];

  switch (childDef.rewrite) {
    case RewriteType.Direct: {
      // for this case, as long as the type / type+relation defined, we should be fine
      const fromPossibleTypes = getTypeRestrictions(currentRelationMetadata?.directly_related_user_types || []);
      if (!fromPossibleTypes.length) {
        const typeIndex = getTypeLineNumber(type, lines);
        const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
        collector.raiseAssignableRelationMustHaveTypes(lineIndex, relation);
      }
      for (const item of fromPossibleTypes) {
        const { decodedType, decodedRelation, isWildcard, decodedConditionName } = destructTupleToUserset(item);
        if (!typeMap[decodedType]) {
          // type is not defined
          const typeIndex = getTypeLineNumber(type, lines);
          const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
          collector.raiseInvalidType(lineIndex, `${decodedType}`, decodedType);
        }

        if (decodedConditionName && !conditions[decodedConditionName]) {
          // condition name is not defined
          const typeIndex = getTypeLineNumber(type, lines);
          const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
          collector.raiseInvalidConditionNameInParameter(
            lineIndex,
            `${decodedConditionName}`,
            type,
            relation,
            decodedConditionName,
          );
        }

        if (isWildcard && decodedRelation) {
          // we cannot have both wild carded and relation at the same time
          const typeIndex = getTypeLineNumber(type, lines);
          const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
          collector.raiseAssignableTypeWildcardRelation(lineIndex, item);
        } else if (decodedRelation) {
          if (!typeMap[decodedType] || !typeMap[decodedType].relations![decodedRelation]) {
            // type/relation is not defined
            const typeIndex = getTypeLineNumber(type, lines);
            const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
            collector.raiseInvalidTypeRelation(
              lineIndex,
              `${decodedType}#${decodedRelation}`,
              decodedType,
              decodedRelation,
            );
          }
        }
      }
      break;
    }
    case RewriteType.ComputedUserset: {
      if (childDef.target && !relations![childDef.target]) {
        const typeIndex = getTypeLineNumber(type, lines);
        const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
        const value = childDef.target;
        collector.raiseInvalidRelationError(lineIndex, value, Object.keys(relations));
      }
      break;
    }
    case RewriteType.TupleToUserset: {
      // for this case, we need to consider both the "from" and "relation"
      if (childDef.from && childDef.target) {
        // First, check to see if the childDef.from exists
        if (!relations[childDef.from]) {
          const typeIndex = getTypeLineNumber(type, lines);
          const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
          collector.raiseInvalidTypeRelation(
            lineIndex,
            `${childDef.target} from ${childDef.from}`,
            type,
            childDef.from,
          );
        } else {
          const [fromTypes, isValid] = allowableTypes(typeMap, type, childDef.from);
          if (isValid) {
            const childRelationNotValid = [];
            for (const item of fromTypes) {
              const { decodedType, decodedRelation, isWildcard } = destructTupleToUserset(item);
              if (isWildcard) {
                // we cannot have both wild carded and relation at the same time
                const typeIndex = getTypeLineNumber(type, lines);
                const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
                collector.raiseAssignableTypeWildcardRelation(lineIndex, item);
              } else if (decodedRelation) {
                const typeIndex = getTypeLineNumber(type, lines);
                const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
                collector.raiseTupleUsersetRequiresDirect(lineIndex, childDef.from);
              } else {
                // check to see if the relation is defined in any children
                if (!typeMap[decodedType] || !typeMap[decodedType].relations![childDef.target]) {
                  const typeIndex = getTypeLineNumber(type, lines);
                  const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
                  childRelationNotValid.push({
                    lineIndex,
                    symbol: `${childDef.target} from ${childDef.from}`,
                    typeName: decodedType,
                    relationName: childDef.target,
                  });
                }
              }
            }
            // if none of the children have this relation defined, we should raise error.
            // otherwise, the relation is defined in at least 1 child and should be considered valid
            if (childRelationNotValid.length === fromTypes.length) {
              for (const item of childRelationNotValid) {
                const { lineIndex, symbol, typeName, relationName } = item;
                collector.raiseInvalidTypeRelation(lineIndex, symbol, typeName, relationName);
              }
            }
          } else {
            // the from is not allowed.  Only direct assignable types are allowed.
            const typeIndex = getTypeLineNumber(type, lines);
            const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
            collector.raiseTupleUsersetRequiresDirect(lineIndex, childDef.from);
          }
        }
      }
      break;
    }
  }
}

// ensure all the referenced relations are defined
function relationDefined(
  lines: string[],
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  collector: any,
  typeMap: Record<string, TypeDefinition>,
  type: string,
  relation: string,
  conditions: AuthorizationModel["conditions"],
) {
  const relations = typeMap[type].relations;
  if (!relations || !relations[relation]) {
    return;
  }

  const currentRelation = { ...relations[relation] };
  const children: Userset[] = [currentRelation];

  while (children.length) {
    const child = children.shift();

    if (child?.union?.child.length) {
      children.push(...child.union.child);
    } else if (child?.intersection?.child.length) {
      children.push(...child.intersection.child);
    } else if (child?.difference?.base && child.difference.subtract) {
      children.push(child?.difference?.base, child.difference.subtract);
    } else if (child) {
      childDefDefined(lines, collector, typeMap, type, relation, getRelationalParserResult(child), conditions);
    }
  }
}

function modelValidation(
  lines: string[],
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  collector: any,
  errors: ModelValidationSingleError[],
  authorizationModel: AuthorizationModel,
  //relationsPerType: Record<string, TransformedType>
) {
  if (errors.length) {
    // no point in looking at directly assignable types if the model itself already
    // has other problems
    return;
  }

  const typeMap: Record<string, TypeDefinition> = {};
  const usedConditionNamesSet = new Set();
  authorizationModel.type_definitions?.forEach((typeDef) => {
    const typeName = typeDef.type;
    typeMap[typeName] = typeDef;
    for (const relationName in typeDef.metadata?.relations) {
      (typeDef.metadata?.relations[relationName].directly_related_user_types || []).forEach((typeRestriction) => {
        if (typeRestriction.condition) {
          usedConditionNamesSet.add(typeRestriction.condition);
        }
      });
    }
  });

  // first, validate to ensure all the relation are defined
  authorizationModel.type_definitions?.forEach((typeDef) => {
    const typeName = typeDef.type;

    // parse through each of the relations to do validation
    for (const relationDef in typeDef.relations) {
      relationDefined(lines, collector, typeMap, typeName, relationDef, authorizationModel.conditions);
    }
  });

  if (errors.length === 0) {
    const typeSet = new Set();
    authorizationModel.type_definitions?.forEach((typeDef) => {
      const typeName = typeDef.type;
      // check for duplicate types
      if (typeSet.has(typeName)) {
        const typeIndex = getTypeLineNumber(typeName, lines);
        collector.raiseDuplicateTypeName(typeIndex, typeName);
      }
      typeSet.add(typeDef.type);

      for (const relationDefKey in typeDef.metadata?.relations) {
        // check for duplicate type names in the relation
        checkForDuplicatesTypeNamesInRelation(
          lines,
          collector,
          typeDef.metadata?.relations[relationDefKey],
          relationDefKey,
        );
        // check for duplicate relations
        checkForDuplicatesInRelation(lines, collector, typeDef, relationDefKey);
      }
    });
  }

  // next, ensure all relation have entry point
  // we can skip if there are errors because errors (such as missing relations) will likely lead to no entries
  if (errors.length === 0) {
    authorizationModel.type_definitions?.forEach((typeDef) => {
      const typeName = typeDef.type;
      // parse through each of the relations to do validation
      for (const relationName in typeDef.relations) {
        const currentRelation = typeMap[typeName].relations;
        const [hasEntry, loop] = hasEntryPointOrLoop(
          typeMap,
          typeName,
          relationName,
          currentRelation![relationName],
          {},
        );
        if (!hasEntry) {
          const typeIndex = getTypeLineNumber(typeName, lines);
          const lineIndex = getRelationLineNumber(relationName, lines, typeIndex);
          if (loop) {
            collector.raiseNoEntryPointLoop(lineIndex, relationName, typeName);
          } else {
            collector.raiseNoEntryPoint(lineIndex, relationName, typeName);
          }
        }
      }
    });
  }

  for (const conditionName in authorizationModel.conditions) {
    // const condition = authorizationModel.conditions[conditionName];
    // Ensure that the nested condition name matches
    // TODO: This does not make sense for the DSL, and is a JSON only error
    // if (conditionName != condition.name) {
    //   const conditionIndex = geConditionLineNumber(conditionName, lines);
    //   collector.raiseDifferentNestedConditionName(conditionIndex, conditionName);
    // }

    // Ensure that the condition has been used
    if (!usedConditionNamesSet.has(conditionName)) {
      const conditionIndex = geConditionLineNumber(conditionName, lines);
      collector.raiseUnusedCondition(conditionIndex, conditionName);
    }
  }
}

function populateRelations(
  lines: string[],
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  collector: any,
  authorizationModel: AuthorizationModel,
  typeRegex: ValidationRegex,
  relationRegex: ValidationRegex,
) {
  // Looking at the types
  authorizationModel.type_definitions?.forEach((typeDef) => {
    const typeName = typeDef.type;

    if (typeName === Keyword.SELF || typeName === ReservedKeywords.THIS) {
      const lineIndex = getTypeLineNumber(typeName, lines);
      collector.raiseReservedTypeName(lineIndex, typeName);
    }

    if (!typeRegex.regex.test(typeName)) {
      const lineIndex = getTypeLineNumber(typeName, lines);
      collector.raiseInvalidName(lineIndex, typeName, typeRegex.rule);
    }

    // Include keyword
    const encounteredRelationsInType: Record<string, boolean> = { [Keyword.SELF]: true };

    for (const relationKey in typeDef.relations) {
      const relationName = relationKey;

      if (relationName === Keyword.SELF || relationName === ReservedKeywords.THIS) {
        const typeIndex = getTypeLineNumber(typeName, lines);
        const lineIndex = getRelationLineNumber(relationName, lines, typeIndex);
        collector.raiseReservedRelationName(lineIndex, relationName);
      } else if (!relationRegex.regex.test(relationName)) {
        const typeIndex = getTypeLineNumber(typeName, lines);
        const lineIndex = getRelationLineNumber(relationName, lines, typeIndex);
        collector.raiseInvalidName(lineIndex, relationName, relationRegex.rule, typeName);
      } else if (encounteredRelationsInType[relationName]) {
        // Check if we have any duplicate relations
        // figure out what is the lineIdx in question
        const typeIndex = getTypeLineNumber(typeName, lines);
        const initialLineIdx = getRelationLineNumber(relationName, lines, typeIndex);
        const duplicateLineIdx = getRelationLineNumber(relationName, lines, initialLineIdx + 1);
        collector.raiseDuplicateDefinition(duplicateLineIdx, relationName);
      }
      encounteredRelationsInType[relationName] = true;
    }
  });
}

/**
 * validateJSON - Given a JSON string, validates that it is a valid OpenFGA model
 * @param {string} jsonString
 * @param {AuthorizationModel} authorizationModel
 * @param {ValidationOptions} options
 */
export function validateJSON(
  jsonString: string,
  authorizationModel: AuthorizationModel,
  options: ValidationOptions = {},
): void {
  const lines = jsonString.split("\n");
  const errors: ModelValidationSingleError[] = [];
  const collector = exceptionCollector(errors, lines);
  const typeValidation = options.typeValidation || defaultTypeRule;
  const relationValidation = options.relationValidation || defaultRelationRule;
  const defaultRegex = new RegExp("[a-zA-Z]*");

  let typeRegex: ValidationRegex = {
    regex: defaultRegex,
    rule: typeValidation,
  };
  try {
    typeRegex = {
      regex: new RegExp(typeValidation),
      rule: typeValidation,
    };
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
  } catch (e: any) {
    throw new ConfigurationError(`Incorrect type regex specification for ${typeValidation}`, e);
  }

  let relationRegex: ValidationRegex = {
    regex: defaultRegex,
    rule: relationValidation,
  };
  try {
    relationRegex = {
      regex: new RegExp(relationValidation),
      rule: relationValidation,
    };
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
  } catch (e: any) {
    throw new ConfigurationError(`Incorrect relation regex specification for ${relationValidation}`, e);
  }

  populateRelations(lines, collector, authorizationModel, typeRegex, relationRegex);

  const schemaVersion = authorizationModel.schema_version;

  if (!schemaVersion) {
    collector.raiseSchemaVersionRequired(0, "");
    throw new ModelValidationError(errors);
  }

  switch (schemaVersion) {
    case "1.1":
      modelValidation(lines, collector, errors, authorizationModel);
      break;
    default: {
      const lineIndex = getSchemaLineNumber(schemaVersion, lines);
      collector.raiseInvalidSchemaVersion(lineIndex, schemaVersion);
      break;
    }
  }

  if (errors.length) {
    throw new ModelValidationError(errors);
  }
}

/**
 * validateDSL - Given a string, validates that it is in valid FGA DSL syntax
 * @param {string} dsl
 * @param {ValidationOptions} options
 * @throws {DSLSyntaxError}
 */
export function validateDSL(dsl: string, options: ValidationOptions = {}): void {
  const { listener, errorListener } = parseDSL(dsl);

  if (errorListener.errors.length) {
    throw new DSLSyntaxError(errorListener.errors);
  }

  validateJSON(dsl, listener.authorizationModel as AuthorizationModel, options);
}
