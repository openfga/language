import { ModelValidationSingleError, ValidationError } from "../errors";
import { Keyword, ReservedKeywords } from "../validator/keywords";

interface ValidationErrorProps {
  message: string;
  lines: string[];
  lineIndex: number;
  metadata: {
    symbol: string;
    errorType: ValidationError;
    relation?: string;
    typeName?: string;
    conditionName?: string;
  };
  customResolver?: (wordIndex: number, rawLine: string, symbol: string) => number;
}

interface BaseProps {
  errors: ModelValidationSingleError[];
  lines: string[];
  lineIndex: number;
  symbol: string;
}

const createInvalidName = (props: BaseProps, clause: string, typeName?: string) => {
  const { errors, lines, lineIndex, symbol } = props;
  const errorMessage =
    (typeName ? `relation '${symbol}' of type '${typeName}' ` : `type '${symbol}' `) +
    `does not match naming rule: '${clause}'.`;
  errors.push(
    constructValidationError({
      message: errorMessage,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.InvalidName },
    }),
  );
};

const createReservedTypeNameError = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol } = props;
  errors.push(
    constructValidationError({
      message: `a type cannot be named '${Keyword.SELF}' or '${ReservedKeywords.THIS}'.`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.ReservedTypeKeywords },
    }),
  );
};

const createReservedRelationNameError = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol } = props;
  errors.push(
    constructValidationError({
      message: `a relation cannot be named '${Keyword.SELF}' or '${ReservedKeywords.THIS}'.`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.ReservedRelationKeywords },
    }),
  );
};

const createTupleUsersetRequireDirectError = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol } = props;

  errors.push(
    constructValidationError({
      message: `\`${symbol}\` relation used inside from allows only direct relation.`,
      lines,
      lineIndex,
      customResolver: (wordIdx, rawLine, value) => {
        const clauseStartsAt = rawLine.indexOf("from") + "from".length + 1;
        wordIdx = clauseStartsAt + rawLine.slice(clauseStartsAt).indexOf(value) + 1;
        return wordIdx;
      },
      metadata: { symbol, errorType: ValidationError.TuplesetNotDirect },
    }),
  );
};

const createNoEntryPointLoopError = (props: BaseProps, typeName: string) => {
  const { errors, lines, lineIndex, symbol } = props;
  errors.push(
    constructValidationError({
      message: `\`${symbol}\` is an impossible relation for \`${typeName}\` (potential loop).`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.RelationNoEntrypoint, relation: symbol },
    }),
  );
};

const createNoEntryPointError = (props: BaseProps, typeName: string) => {
  const { errors, lines, lineIndex, symbol } = props;
  errors.push(
    constructValidationError({
      message: `\`${symbol}\` is an impossible relation for \`${typeName}\` (no entrypoint).`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.RelationNoEntrypoint, relation: symbol },
    }),
  );
};

const createInvalidTypeRelationError = (props: BaseProps, typeName: string, relationName: string) => {
  const { errors, lines, lineIndex, symbol } = props;
  errors.push(
    constructValidationError({
      message: `\`${relationName}\` is not a valid relation for \`${typeName}\`.`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.InvalidRelationType, relation: relationName, typeName },
    }),
  );
};

const createInvalidConditionNameInParameterError = (
  props: BaseProps,
  typeName: string,
  relationName: string,
  conditionName: string,
) => {
  const { errors, lines, lineIndex, symbol } = props;
  // const rawLine = lines[lineIndex];
  // const actualValue = rawLine.includes("[")
  //   ? rawLine.slice(rawLine.indexOf("["), rawLine.lastIndexOf("]") + 1)
  //   : "self";
  errors.push(
    constructValidationError({
      message: `\`${conditionName}\` is not a defined condition in the model.`,
      lines,
      lineIndex,
      customResolver: (wordIdx, rawLine, symbol) => {
        wordIdx = rawLine.indexOf(symbol.substring(1));
        return wordIdx;
      },
      metadata: {
        symbol,
        errorType: ValidationError.ConditionNotDefined,
        relation: relationName,
        typeName,
      },
    }),
  );
};

const createUnusedConditionError = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol } = props;
  errors.push(
    constructValidationError({
      message: `\`${symbol}\` condition is not used in the model.`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.ConditionNotUsed, conditionName: symbol },
    }),
  );
};

const createInvalidTypeError = (props: BaseProps, typeName: string) => {
  const { errors, lines, lineIndex, symbol } = props;
  errors.push(
    constructValidationError({
      message: `\`${typeName}\` is not a valid type.`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.InvalidType, typeName: typeName },
    }),
  );
};

const createAssignableRelationMustHaveTypesError = (props: BaseProps) => {
  const { errors, lines, lineIndex } = props;
  const rawLine = lines[lineIndex];
  const actualValue = rawLine.includes("[")
    ? rawLine.slice(rawLine.indexOf("["), rawLine.lastIndexOf("]") + 1)
    : "self";

  errors.push(
    constructValidationError({
      message: `assignable relation '${actualValue}' must have types`,
      lines,
      lineIndex,
      customResolver: (wordIdx, rawLine, symbol) => {
        wordIdx = rawLine.indexOf(symbol.substring(1));
        return wordIdx;
      },
      metadata: { symbol: actualValue, errorType: ValidationError.AssignableRelationsMustHaveType },
    }),
  );
};

const createDuplicateTypeNameError = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol } = props;
  errors.push(
    constructValidationError({
      message: `the type \`${symbol}\` is a duplicate.`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.DuplicatedError },
    }),
  );
};

const createDuplicateTypeRestrictionError = (props: BaseProps, relationName: string) => {
  const { errors, lines, lineIndex, symbol } = props;
  errors.push(
    constructValidationError({
      message: `the type restriction \`${symbol}\` is a duplicate in the relation \`${relationName}\`.`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.DuplicatedError, relation: symbol },
    }),
  );
};

const createDuplicateRelationError = (props: BaseProps, relationName: string) => {
  const { errors, lines, lineIndex, symbol } = props;
  errors.push(
    constructValidationError({
      message: `the partial relation definition \`${symbol}\` is a duplicate in the relation \`${relationName}\`.`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.DuplicatedError, relation: symbol },
    }),
  );
};

const createDuplicateRelationshipDefinitionError = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol } = props;
  const rawLine = lines[lineIndex];

  errors.push(
    new ModelValidationSingleError(
      {
        msg: `duplicate definition \`${symbol}\`.`,
        line: {
          start: lineIndex + 1,
          end: lineIndex + 1,
        },
        column: {
          start: rawLine.indexOf(Keyword.DEFINE) + 1,
          end: rawLine.length + 1,
        },
      },
      { symbol, errorType: ValidationError.DuplicatedError },
    ),
  );
};

const createAssignableTypeWildcardRelationError = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol } = props;
  errors.push(
    constructValidationError({
      message: `type restriction \`${symbol}\` cannot contain both wildcard and relation`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.TypeRestrictionCannotHaveWildcardAndRelation },
    }),
  );
};

const createInvalidRelationError = (props: BaseProps, validRelations: string[]) => {
  const { errors, lines, lineIndex, symbol } = props;
  const isInValid = !validRelations?.includes(symbol);
  if (isInValid) {
    errors.push(
      constructValidationError({
        message: `the relation \`${symbol}\` does not exist.`,
        lines,
        lineIndex,
        metadata: { symbol, errorType: ValidationError.MissingDefinition, relation: symbol },
      }),
    );
  }
};

export const createInvalidSchemaVersionError = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol } = props;
  errors.push(
    constructValidationError({
      message: `invalid schema ${symbol}`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.InvalidSchema },
    }),
  );
};

export const createSchemaVersionRequiredError = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol } = props;
  errors.push(
    constructValidationError({
      message: `schema version required`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.SchemaVersionRequired },
    }),
  );
};

export const createMaximumOneDirectRelationship = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol } = props;
  errors.push(
    constructValidationError({
      message: "each relationship must have at most 1 set of direct relations defined.",
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.AssignableRelationsMustHaveType },
      customResolver: (wordIdx, rawLine, symbol) => {
        wordIdx = rawLine.indexOf(symbol.substring(1));
        return wordIdx;
      },
    }),
  );
};

function constructValidationError(props: ValidationErrorProps): ModelValidationSingleError {
  const { message, lines, lineIndex, customResolver, metadata } = props;

  const rawLine = lines[lineIndex];
  const re = new RegExp("\\b" + metadata.symbol + "\\b");
  let wordIdx = rawLine?.search(re) + 1;

  if (typeof customResolver === "function") {
    wordIdx = customResolver(wordIdx, rawLine, metadata.symbol);
  }

  return new ModelValidationSingleError(
    {
      line: { start: lineIndex + 1, end: lineIndex + 1 },
      column: { start: wordIdx, end: wordIdx + metadata.symbol.length },
      msg: message,
    },
    { symbol: metadata?.symbol, errorType: metadata.errorType },
  );
}

export const exceptionCollector = (errors: ModelValidationSingleError[], lines: string[]) => {
  return {
    raiseInvalidName(lineIndex: number, symbol: string, clause: string, typeName?: string) {
      createInvalidName({ errors, lines, lineIndex, symbol }, clause, typeName);
    },
    raiseReservedTypeName(lineIndex: number, symbol: string) {
      createReservedTypeNameError({ errors, lines, lineIndex, symbol });
    },
    raiseReservedRelationName(lineIndex: number, symbol: string) {
      createReservedRelationNameError({ errors, lines, lineIndex, symbol });
    },
    raiseTupleUsersetRequiresDirect(lineIndex: number, symbol: string) {
      createTupleUsersetRequireDirectError({ errors, lines, lineIndex, symbol });
    },
    raiseDuplicateTypeName(lineIndex: number, symbol: string) {
      createDuplicateTypeNameError({ errors, lines, lineIndex, symbol });
    },
    raiseDuplicateTypeRestriction(lineIndex: number, symbol: string, relationName: string) {
      createDuplicateTypeRestrictionError({ errors, lines, lineIndex, symbol }, relationName);
    },
    raiseDuplicateType(lineIndex: number, symbol: string, relationName: string) {
      createDuplicateRelationError({ errors, lines, lineIndex, symbol }, relationName);
    },
    raiseDuplicateRelationshipDefinition(lineIndex: number, symbol: string) {
      createDuplicateRelationshipDefinitionError({ errors, lines, lineIndex, symbol });
    },
    raiseNoEntryPointLoop(lineIndex: number, symbol: string, typeName: string) {
      createNoEntryPointLoopError({ errors, lines, lineIndex, symbol }, typeName);
    },
    raiseNoEntryPoint(lineIndex: number, symbol: string, typeName: string) {
      createNoEntryPointError({ errors, lines, lineIndex, symbol }, typeName);
    },
    raiseInvalidTypeRelation(lineIndex: number, symbol: string, typeName: string, relationName: string) {
      createInvalidTypeRelationError({ errors, lines, lineIndex, symbol }, typeName, relationName);
    },
    raiseInvalidType(lineIndex: number, symbol: string, typeName: string) {
      createInvalidTypeError({ errors, lines, lineIndex, symbol }, typeName);
    },
    raiseAssignableRelationMustHaveTypes(lineIndex: number, symbol: string) {
      createAssignableRelationMustHaveTypesError({ errors, lines, lineIndex, symbol });
    },
    raiseAssignableTypeWildcardRelation(lineIndex: number, symbol: string) {
      createAssignableTypeWildcardRelationError({ errors, lines, lineIndex, symbol });
    },
    raiseInvalidRelationError(lineIndex: number, symbol: string, validRelations: string[]) {
      createInvalidRelationError({ errors, lines, lineIndex, symbol }, validRelations);
    },
    raiseInvalidSchemaVersion(lineIndex: number, symbol: string) {
      createInvalidSchemaVersionError({ errors, lines, lineIndex, symbol });
    },
    raiseSchemaVersionRequired(lineIndex: number, symbol: string) {
      createSchemaVersionRequiredError({ errors, lines, lineIndex, symbol });
    },
    raiseMaximumOneDirectRelationship(lineIndex: number, symbol: string) {
      createMaximumOneDirectRelationship({ errors, lines, lineIndex, symbol });
    },
    raiseInvalidConditionNameInParameter(
      lineIndex: number,
      symbol: string,
      typeName: string,
      relationName: string,
      conditionName: string,
    ) {
      createInvalidConditionNameInParameterError(
        { errors, lines, lineIndex, symbol },
        typeName,
        relationName,
        conditionName,
      );
    },
    raiseUnusedCondition(lineIndex: number, symbol: string) {
      createUnusedConditionError({ errors, lines, lineIndex, symbol });
    },
  };
};
