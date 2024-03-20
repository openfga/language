import type { AuthorizationModel, RelationReference, RelationMetadata, TypeDefinition, Userset } from "@openfga/sdk";
import { Keyword, ReservedKeywords } from "./keywords";
import { parseDSL } from "../transformer";
import { ConfigurationError, DSLSyntaxError, ModelValidationError, ModelValidationSingleError } from "../errors";
import { ExceptionCollector } from "../util/exceptions";

// eslint-disable-next-line no-useless-escape
const defaultTypeRule = "^[^:#@\\s]{1,254}$";
// eslint-disable-next-line no-useless-escape
const defaultRelationRule = "^[^:#@\\s]{1,50}$";

enum RelationDefOperator {
  Union = "union",
  Intersection = "intersection",
  Difference = "difference",
}

interface ValidationRegex {
  rule: string;
  regex: RegExp;
}

interface ValidationOptions {
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

const deepCopy = <T>(object: T): T => {
  return JSON.parse(JSON.stringify(object));
};

const relationIsSingle = (currentRelation: Userset): boolean => {
  return (
    !Object.prototype.hasOwnProperty.call(currentRelation, RelationDefOperator.Union) &&
    !Object.prototype.hasOwnProperty.call(currentRelation, RelationDefOperator.Intersection) &&
    !Object.prototype.hasOwnProperty.call(currentRelation, RelationDefOperator.Difference)
  );
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

const geConditionLineNumber = (conditionName: string, lines?: string[], skipIndex?: number) => {
  if (!skipIndex) {
    skipIndex = 0;
  }
  if (!lines) {
    return undefined;
  }
  return (
    lines.slice(skipIndex).findIndex((line: string) => line.trim().startsWith(`condition ${conditionName}`)) + skipIndex
  );
};

const getTypeLineNumber = (typeName: string, lines?: string[], skipIndex?: number) => {
  if (!skipIndex) {
    skipIndex = 0;
  }
  if (!lines) {
    return undefined;
  }
  return lines.slice(skipIndex).findIndex((line: string) => line.trim().startsWith(`type ${typeName}`)) + skipIndex;
};

const getRelationLineNumber = (relation: string, lines?: string[], skipIndex?: number) => {
  if (!skipIndex) {
    skipIndex = 0;
  }
  if (!lines) {
    return undefined;
  }
  return (
    lines
      .slice(skipIndex)
      .findIndex((line: string) => line.trim().replace(/ {2,}/g, " ").startsWith(`define ${relation}`)) + skipIndex
  );
};

const getSchemaLineNumber = (schema: string, lines?: string[]) => {
  if (!lines) {
    return undefined;
  }

  const index = lines.findIndex((line: string) => line.trim().replace(/ {2,}/g, " ").startsWith(`schema ${schema}`));

  // As findIndex returns -1 when it doesn't find the line, we want to return 0 instead
  if (index >= 1) {
    return index;
  } else {
    return 0;
  }
};

function checkForDuplicatesTypeNamesInRelation(
  collector: ExceptionCollector,
  relationDef: RelationMetadata,
  relationName: string,
  lines?: string[],
) {
  const typeNameSet = new Set();
  relationDef.directly_related_user_types?.forEach((typeDef) => {
    const typeDefName = getTypeRestrictionString(typeDef);

    if (typeNameSet.has(typeDefName)) {
      const typeIndex = getTypeLineNumber(typeDef.type, lines);
      const lineIndex = getRelationLineNumber(relationName, lines, typeIndex);
      collector.raiseDuplicateTypeRestriction(typeDefName, relationName, lineIndex);
    }
    typeNameSet.add(typeDefName);
  });
}

// ensure all the referenced relations are defined
function checkForDuplicatesInRelation(
  collector: ExceptionCollector,
  typeDef: TypeDefinition,
  relationName: string,
  lines?: string[],
) {
  const relationDef = typeDef.relations![relationName];

  // Union
  const relationUnionNameSet = new Set();
  relationDef.union?.child?.forEach((userset) => {
    const relationDefName = getRelationDefName(userset);
    if (relationDefName && relationUnionNameSet.has(relationDefName)) {
      const typeIndex = getTypeLineNumber(typeDef.type, lines);
      const lineIndex = getRelationLineNumber(relationName, lines, typeIndex);
      collector.raiseDuplicateType(relationDefName, relationName, lineIndex);
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
      collector.raiseDuplicateType(relationDefName, relationName, lineIndex);
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
      collector.raiseDuplicateType(baseName, relationName, lineIndex);
    }
  }
}

// helper function to ensure all childDefs are defined
function childDefDefined(
  collector: ExceptionCollector,
  typeMap: Record<string, TypeDefinition>,
  type: string,
  relation: string,
  childDef: RelationTargetParserResult,
  conditions: AuthorizationModel["conditions"] = {},
  lines?: string[],
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
        collector.raiseAssignableRelationMustHaveTypes(relation, lineIndex);
      }
      for (const item of fromPossibleTypes) {
        const { decodedType, decodedRelation, isWildcard, decodedConditionName } = destructTupleToUserset(item);
        if (!typeMap[decodedType]) {
          // type is not defined
          const typeIndex = getTypeLineNumber(type, lines);
          const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
          collector.raiseInvalidType(`${decodedType}`, decodedType, lineIndex);
        }

        if (decodedConditionName && !conditions[decodedConditionName]) {
          // condition name is not defined
          const typeIndex = getTypeLineNumber(type, lines);
          const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
          collector.raiseInvalidConditionNameInParameter(
            `${decodedConditionName}`,
            type,
            relation,
            decodedConditionName,
            lineIndex,
          );
        }

        if (isWildcard && decodedRelation) {
          // we cannot have both wild carded and relation at the same time
          const typeIndex = getTypeLineNumber(type, lines);
          const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
          collector.raiseAssignableTypeWildcardRelation(item, lineIndex);
        } else if (decodedRelation) {
          if (!typeMap[decodedType] || !typeMap[decodedType].relations![decodedRelation]) {
            // type/relation is not defined
            const typeIndex = getTypeLineNumber(type, lines);
            const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
            collector.raiseInvalidTypeRelation(
              `${decodedType}#${decodedRelation}`,
              decodedType,
              decodedRelation,
              lineIndex,
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
        collector.raiseInvalidRelationError(value, Object.keys(relations), lineIndex);
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
            `${childDef.target} from ${childDef.from}`,
            type,
            childDef.from,
            lineIndex,
          );
        } else {
          const [fromTypes, isValid] = allowableTypes(typeMap, type, childDef.from);
          if (isValid) {
            const childRelationNotValid = [];
            for (const item of fromTypes) {
              const { decodedType, decodedRelation, isWildcard } = destructTupleToUserset(item);
              if (isWildcard || decodedRelation) {
                // we cannot have both wildcard or decoded relation and relation at the same time
                const typeIndex = getTypeLineNumber(type, lines);
                const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
                collector.raiseTupleUsersetRequiresDirect(childDef.from, lineIndex);
              } else {
                // check to see if the relation is defined in any children
                if (!typeMap[decodedType] || !typeMap[decodedType].relations![childDef.target]) {
                  const typeIndex = getTypeLineNumber(type, lines);
                  const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
                  childRelationNotValid.push({
                    symbol: `${childDef.target} from ${childDef.from}`,
                    typeName: decodedType,
                    relationName: childDef.target,
                    lineIndex,
                  });
                }
              }
            }
            // if none of the children have this relation defined, we should raise error.
            // otherwise, the relation is defined in at least 1 child and should be considered valid
            if (childRelationNotValid.length === fromTypes.length) {
              for (const item of childRelationNotValid) {
                const { lineIndex, symbol, typeName, relationName } = item;
                collector.raiseInvalidTypeRelation(symbol, typeName, relationName, lineIndex);
              }
            }
          } else {
            // the from is not allowed.  Only direct assignable types are allowed.
            const typeIndex = getTypeLineNumber(type, lines);
            const lineIndex = getRelationLineNumber(relation, lines, typeIndex);
            collector.raiseTupleUsersetRequiresDirect(childDef.from, lineIndex);
          }
        }
      }
      break;
    }
  }
}

// ensure all the referenced relations are defined
function relationDefined(
  collector: ExceptionCollector,
  typeMap: Record<string, TypeDefinition>,
  type: string,
  relation: string,
  conditions: AuthorizationModel["conditions"],
  lines?: string[],
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
      childDefDefined(collector, typeMap, type, relation, getRelationalParserResult(child), conditions, lines);
    }
  }
}

function modelValidation(
  collector: ExceptionCollector,
  errors: ModelValidationSingleError[],
  authorizationModel: AuthorizationModel,
  //relationsPerType: Record<string, TransformedType>
  lines?: string[],
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
      relationDefined(collector, typeMap, typeName, relationDef, authorizationModel.conditions, lines);
    }
  });

  if (errors.length === 0) {
    const typeSet = new Set();
    authorizationModel.type_definitions?.forEach((typeDef) => {
      const typeName = typeDef.type;
      // check for duplicate types
      if (typeSet.has(typeName)) {
        const typeIndex = getTypeLineNumber(typeName, lines);
        collector.raiseDuplicateTypeName(typeName, typeIndex);
      }
      typeSet.add(typeDef.type);

      for (const relationDefKey in typeDef.metadata?.relations) {
        // check for duplicate type names in the relation
        checkForDuplicatesTypeNamesInRelation(
          collector,
          typeDef.metadata?.relations[relationDefKey],
          relationDefKey,
          lines,
        );
        // check for duplicate relations
        checkForDuplicatesInRelation(collector, typeDef, relationDefKey, lines);
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
            collector.raiseNoEntryPointLoop(relationName, typeName, lineIndex);
          } else {
            collector.raiseNoEntryPoint(relationName, typeName, lineIndex);
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
      collector.raiseUnusedCondition(conditionName, conditionIndex);
    }
  }
}

function populateRelations(
  collector: ExceptionCollector,
  authorizationModel: AuthorizationModel,
  typeRegex: ValidationRegex,
  relationRegex: ValidationRegex,
  lines?: string[],
) {
  // Looking at the types
  authorizationModel.type_definitions?.forEach((typeDef) => {
    const typeName = typeDef.type;

    if (typeName === Keyword.SELF || typeName === ReservedKeywords.THIS) {
      const lineIndex = getTypeLineNumber(typeName, lines);
      collector.raiseReservedTypeName(typeName, lineIndex, {
        file: typeDef.metadata?.source_info?.file,
        module: typeDef.metadata?.module,
      });
    }

    if (!typeRegex.regex.test(typeName)) {
      const lineIndex = getTypeLineNumber(typeName, lines);
      collector.raiseInvalidName(typeName, typeRegex.rule, undefined, lineIndex, {
        file: typeDef.metadata?.source_info?.file,
        module: typeDef.metadata?.module,
      });
    }

    for (const relationKey in typeDef.relations) {
      const relationName = relationKey;
      let relationMeta = typeDef.metadata?.relations?.[relationKey];
      if (!relationMeta?.module) {
        // relation belongs to typedef
        relationMeta = typeDef.metadata;
      }

      if (relationName === Keyword.SELF || relationName === ReservedKeywords.THIS) {
        const typeIndex = getTypeLineNumber(typeName, lines);
        const lineIndex = getRelationLineNumber(relationName, lines, typeIndex);
        collector.raiseReservedRelationName(relationName, lineIndex, {
          file: relationMeta?.source_info?.file,
          module: relationMeta?.module,
        });
      }

      if (!relationRegex.regex.test(relationName)) {
        const typeIndex = getTypeLineNumber(typeName, lines);
        const lineIndex = getRelationLineNumber(relationName, lines, typeIndex);
        collector.raiseInvalidName(relationName, relationRegex.rule, typeName, lineIndex, {
          file: relationMeta?.source_info?.file,
          module: relationMeta?.module,
        });
      }
    }
  });
}

/**
 * validateJSON - Given a JSON string, validates that it is a valid OpenFGA model
 * @param {string} dslString
 * @param {AuthorizationModel} authorizationModel
 * @param {ValidationOptions} options
 */
export function validateJSON(
  authorizationModel: AuthorizationModel,
  options: ValidationOptions = {},
  dslString?: string,
): void {
  const lines = dslString?.split("\n");
  const errors: ModelValidationSingleError[] = [];
  const collector = new ExceptionCollector(errors, lines);
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

  populateRelations(collector, authorizationModel, typeRegex, relationRegex, lines);

  const schemaVersion = authorizationModel.schema_version;

  if (!schemaVersion) {
    collector.raiseSchemaVersionRequired("", 0);
  }

  switch (schemaVersion) {
    case "1.1":
    case "1.2":
      modelValidation(collector, errors, authorizationModel, lines);
      break;
    case undefined:
      break;
    default: {
      const lineIndex = getSchemaLineNumber(schemaVersion, lines);
      collector.raiseInvalidSchemaVersion(schemaVersion, lineIndex);
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

  validateJSON(listener.authorizationModel as AuthorizationModel, options, dsl);
}
