import { ModelValidationSingleError, ModuleTransformationSingleError, ValidationError } from "../errors";
import { Keyword, ReservedKeywords } from "../validator/keywords";

interface ValidationErrorProps {
  message: string;
  metadata: {
    symbol: string;
    errorType: ValidationError;
    relation?: string;
    typeName?: string;
    conditionName?: string;
    file?: string;
    module?: string;
  };
  lines?: string[];
  lineIndex?: number;
  customResolver?: (wordIndex: number, rawLine: string, symbol: string) => number;
}

interface BaseProps {
  errors: ModelValidationSingleError[];
  symbol: string;
  lines?: string[];
  lineIndex?: number;
  file?: string;
  module?: string;
}

const createInvalidName = (props: BaseProps, clause: string, typeName?: string) => {
  const { errors, lines, lineIndex, symbol, file, module } = props;
  const errorMessage =
    (typeName ? `relation '${symbol}' of type '${typeName}' ` : `type '${symbol}' `) +
    `does not match naming rule: '${clause}'.`;
  errors.push(
    constructValidationError({
      message: errorMessage,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.InvalidName, file, module },
    }),
  );
};

const createReservedTypeNameError = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol, file, module } = props;
  errors.push(
    constructValidationError({
      message: `a type cannot be named '${Keyword.SELF}' or '${ReservedKeywords.THIS}'.`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.ReservedTypeKeywords, file, module },
    }),
  );
};

const createReservedRelationNameError = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol, file, module } = props;
  errors.push(
    constructValidationError({
      message: `a relation cannot be named '${Keyword.SELF}' or '${ReservedKeywords.THIS}'.`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.ReservedRelationKeywords, file, module },
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

  if (!lines?.length || lineIndex === undefined) {
    const actualValue = "";
    errors.push(
      constructValidationError({
        message: `assignable relation '${actualValue}' must have types`,
        metadata: { symbol: actualValue, errorType: ValidationError.AssignableRelationsMustHaveType },
      }),
    );
    return;
  }

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

  if (!lines?.length || lineIndex === undefined) {
    errors.push(
      new ModelValidationSingleError(
        {
          msg: `duplicate relationship definition \`${symbol}\`.`,
        },
        { symbol, errorType: ValidationError.DuplicatedError },
      ),
    );
    return;
  }

  const rawLine = lines[lineIndex];

  errors.push(
    new ModelValidationSingleError(
      {
        msg: `duplicate relationship definition \`${symbol}\`.`,
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
      message: "schema version required",
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

  if (!lines?.length || lineIndex === undefined) {
    return new ModelValidationSingleError(
      {
        msg: message,
        file: metadata.file,
      },
      { symbol: metadata?.symbol, errorType: metadata.errorType, module: metadata.module },
    );
  }

  const rawLine = lines[lineIndex];

  const re = new RegExp("\\b" + metadata.symbol + "\\b");
  let wordIdx = rawLine?.search(re) + 1;

  if (isNaN(wordIdx) || wordIdx === 0) {
    wordIdx = 1;
  }

  if (typeof customResolver === "function") {
    wordIdx = customResolver(wordIdx, rawLine, metadata.symbol);
  }

  return new ModelValidationSingleError(
    {
      line: { start: lineIndex + 1, end: lineIndex + 1 },
      column: { start: wordIdx, end: wordIdx + (metadata.symbol?.length || 0) },
      msg: message,
      file: metadata.file,
    },
    { symbol: metadata?.symbol, errorType: metadata.errorType, module: metadata.module },
  );
}

interface Meta {
  file?: string;
  module?: string;
}

export class ExceptionCollector {
  constructor(
    private errors: ModelValidationSingleError[],
    private lines?: string[],
  ) {}

  raiseInvalidName(symbol: string, clause: string, typeName?: string, lineIndex?: number, metadata?: Meta) {
    createInvalidName(
      { errors: this.errors, lines: this.lines, lineIndex, symbol, file: metadata?.file, module: metadata?.module },
      clause,
      typeName,
    );
  }

  raiseReservedTypeName(symbol: string, lineIndex?: number, metadata?: Meta) {
    createReservedTypeNameError({
      errors: this.errors,
      lines: this.lines,
      lineIndex,
      symbol,
      file: metadata?.file,
      module: metadata?.module,
    });
  }

  raiseReservedRelationName(symbol: string, lineIndex?: number, metadata?: Meta) {
    createReservedRelationNameError({
      errors: this.errors,
      lines: this.lines,
      lineIndex,
      symbol,
      file: metadata?.file,
      module: metadata?.module,
    });
  }

  raiseTupleUsersetRequiresDirect(symbol: string, lineIndex?: number) {
    createTupleUsersetRequireDirectError({ errors: this.errors, lines: this.lines, lineIndex, symbol });
  }

  raiseDuplicateTypeName(symbol: string, lineIndex?: number) {
    createDuplicateTypeNameError({ errors: this.errors, lines: this.lines, lineIndex, symbol });
  }

  raiseDuplicateTypeRestriction(symbol: string, relationName: string, lineIndex?: number) {
    createDuplicateTypeRestrictionError({ errors: this.errors, lines: this.lines, lineIndex, symbol }, relationName);
  }

  raiseDuplicateType(symbol: string, relationName: string, lineIndex?: number) {
    createDuplicateRelationError({ errors: this.errors, lines: this.lines, lineIndex, symbol }, relationName);
  }

  raiseDuplicateRelationshipDefinition(symbol: string, lineIndex?: number) {
    createDuplicateRelationshipDefinitionError({ errors: this.errors, lines: this.lines, lineIndex, symbol });
  }

  raiseNoEntryPointLoop(symbol: string, typeName: string, lineIndex?: number) {
    createNoEntryPointLoopError({ errors: this.errors, lines: this.lines, lineIndex, symbol }, typeName);
  }

  raiseNoEntryPoint(symbol: string, typeName: string, lineIndex?: number) {
    createNoEntryPointError({ errors: this.errors, lines: this.lines, lineIndex, symbol }, typeName);
  }

  raiseInvalidTypeRelation(symbol: string, typeName: string, relationName: string, lineIndex?: number) {
    createInvalidTypeRelationError(
      {
        errors: this.errors,
        lines: this.lines,
        lineIndex,
        symbol,
      },
      typeName,
      relationName,
    );
  }

  raiseInvalidType(symbol: string, typeName: string, lineIndex?: number) {
    createInvalidTypeError({ errors: this.errors, lines: this.lines, lineIndex, symbol }, typeName);
  }

  raiseAssignableRelationMustHaveTypes(symbol: string, lineIndex?: number) {
    createAssignableRelationMustHaveTypesError({ errors: this.errors, lines: this.lines, lineIndex, symbol });
  }

  raiseAssignableTypeWildcardRelation(symbol: string, lineIndex?: number) {
    createAssignableTypeWildcardRelationError({ errors: this.errors, lines: this.lines, lineIndex, symbol });
  }

  raiseInvalidRelationError(symbol: string, validRelations: string[], lineIndex?: number) {
    createInvalidRelationError({ errors: this.errors, lines: this.lines, lineIndex, symbol }, validRelations);
  }

  raiseInvalidSchemaVersion(symbol: string, lineIndex?: number) {
    createInvalidSchemaVersionError({ errors: this.errors, lines: this.lines, lineIndex, symbol });
  }

  raiseSchemaVersionRequired(symbol: string, lineIndex?: number) {
    createSchemaVersionRequiredError({ errors: this.errors, lines: this.lines, lineIndex, symbol });
  }

  raiseMaximumOneDirectRelationship(symbol: string, lineIndex?: number) {
    createMaximumOneDirectRelationship({ errors: this.errors, lines: this.lines, lineIndex, symbol });
  }

  raiseInvalidConditionNameInParameter(
    symbol: string,
    typeName: string,
    relationName: string,
    conditionName: string,
    lineIndex?: number,
  ) {
    createInvalidConditionNameInParameterError(
      { errors: this.errors, lines: this.lines, lineIndex, symbol },
      typeName,
      relationName,
      conditionName,
    );
  }

  raiseUnusedCondition(symbol: string, lineIndex?: number) {
    createUnusedConditionError({ errors: this.errors, lines: this.lines, lineIndex, symbol });
  }
}

interface TransformationErrorProps {
  message: string;
  metadata: {
    symbol: string;
    relation?: string;
    typeName?: string;
    conditionName?: string;
    file?: string;
  };
  lines?: string[];
  lineIndex?: number;
  customResolver?: (wordIndex: number, rawLine: string, symbol: string) => number;
}

export function constructTransformationError(props: TransformationErrorProps) {
  const { message, lines, lineIndex, metadata } = props;

  if (!lines?.length || lineIndex === undefined) {
    return new ModuleTransformationSingleError(
      {
        msg: message,
      },
      { symbol: metadata?.symbol },
    );
  }

  const rawLine = lines[lineIndex];

  const re = new RegExp("\\b" + metadata.symbol + "\\b");
  let wordIdx = rawLine?.search(re) + 1;

  if (isNaN(wordIdx) || wordIdx === 0) {
    wordIdx = 1;
  }

  return new ModuleTransformationSingleError(
    {
      line: { start: lineIndex + 1, end: lineIndex + 1 },
      column: { start: wordIdx, end: wordIdx + (metadata.symbol?.length || 0) },
      msg: message,
      file: metadata.file,
    },
    { symbol: metadata?.symbol },
  );
}
