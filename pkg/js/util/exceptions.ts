import {
  ErrorMetadata,
  ErrorProperties,
  ModelValidationSingleError,
  ModuleTransformationSingleError,
  ValidationError,
} from "../errors";
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
    offendingType?: string;
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
  type?: string;
  relation?: string;
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
  const { errors, lines, lineIndex, symbol, file, module, type, relation } = props;

  errors.push(
    constructValidationError({
      message: `\`${symbol}\` relation used inside from allows only direct relation.`,
      lines,
      lineIndex,
      customResolver: (wordIdx, rawLine, value) => {
        const clauseStartsAt = rawLine.indexOf("from") + "from".length;
        wordIdx = clauseStartsAt + rawLine.slice(clauseStartsAt).indexOf(value);
        return wordIdx;
      },
      metadata: { symbol, errorType: ValidationError.TuplesetNotDirect, file, module, typeName: type, relation },
    }),
  );
};

const createNoEntryPointLoopError = (props: BaseProps, typeName: string) => {
  const { errors, lines, lineIndex, symbol, file, module } = props;
  errors.push(
    constructValidationError({
      message: `\`${symbol}\` is an impossible relation for \`${typeName}\` (potential loop).`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.RelationNoEntrypoint, relation: symbol, typeName, file, module },
    }),
  );
};

const createNoEntryPointError = (props: BaseProps, typeName: string) => {
  const { errors, lines, lineIndex, symbol, module, file } = props;
  errors.push(
    constructValidationError({
      message: `\`${symbol}\` is an impossible relation for \`${typeName}\` (no entrypoint).`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.RelationNoEntrypoint, relation: symbol, typeName, module, file },
    }),
  );
};
// targetType and targetRelation, typeRestriction
const createInvalidTypeRelationError = (
  props: BaseProps,
  typeName: string,
  relationName: string,
  offendingRelation: string,
  offendingType: string,
) => {
  const { errors, lines, lineIndex, symbol, file, module } = props;
  errors.push(
    constructValidationError({
      message: `\`${offendingRelation}\` is not a valid relation for \`${typeName}\`.`,
      lines,
      lineIndex,
      metadata: {
        symbol,
        errorType: ValidationError.InvalidRelationType,
        relation: relationName,
        typeName,
        file,
        module,
        offendingType,
      },
    }),
  );
};

const createInvalidRelationOnTuplesetError = (
  props: BaseProps,
  typeName: string,
  typeDef: string,
  relationName: string,
  offendingRelation: string,
  parent: string,
) => {
  const { errors, lines, lineIndex, symbol, file, module } = props;
  errors.push(
    constructValidationError({
      message: `the \`${offendingRelation}\` relation definition on type \`${typeDef}\` is not valid: \`${offendingRelation}\` does not exist on \`${parent}\`, which is of type \`${typeName}\`.`,
      lines,
      lineIndex,
      metadata: {
        symbol,
        errorType: ValidationError.InvalidRelationOnTupleset,
        relation: relationName,
        typeName: typeDef,
        file,
        module,
      },
    }),
  );
};

const createInvalidConditionNameInParameterError = (
  props: BaseProps,
  typeName: string,
  relationName: string,
  conditionName: string,
) => {
  const { errors, lines, lineIndex, symbol, module, file } = props;
  errors.push(
    constructValidationError({
      message: `\`${conditionName}\` is not a defined condition in the model.`,
      lines,
      lineIndex,
      metadata: {
        symbol,
        errorType: ValidationError.ConditionNotDefined,
        relation: relationName,
        typeName,
        file,
        module,
      },
    }),
  );
};

const createUnusedConditionError = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol, module, file } = props;
  errors.push(
    constructValidationError({
      message: `\`${symbol}\` condition is not used in the model.`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.ConditionNotUsed, conditionName: symbol, module, file },
    }),
  );
};

const createInvalidTypeError = (props: BaseProps, typeName: string) => {
  const { errors, lines, lineIndex, symbol, file, module, relation } = props;
  errors.push(
    constructValidationError({
      message: `\`${symbol}\` is not a valid type.`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.InvalidType, typeName: typeName, file, module, relation },
      customResolver: (wordIndex: number, rawLine: string, symbol: string): number => {
        // Split line at definition as InvalidType should mark the value, not the key
        const re = new RegExp("\\b" + symbol + "\\b");
        const splitLine = rawLine.split(":");
        return splitLine[0].length + splitLine[1].search(re) + 1;
      },
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
  const actualValue = rawLine.includes("[") ? rawLine.slice(rawLine.indexOf("["), rawLine.lastIndexOf("]")) : "self";

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
  const { errors, lines, lineIndex, symbol, file, module } = props;
  errors.push(
    constructValidationError({
      message: `the type \`${symbol}\` is a duplicate.`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.DuplicatedError, typeName: symbol, file, module },
    }),
  );
};

const createDuplicateTypeRestrictionError = (props: BaseProps, relationName: string, typeName: string) => {
  const { errors, lines, lineIndex, symbol, file, module } = props;
  errors.push(
    constructValidationError({
      message: `the type restriction \`${symbol}\` is a duplicate in the relation \`${relationName}\`.`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.DuplicatedError, relation: relationName, typeName, file, module },
    }),
  );
};

const createDuplicateRelationError = (props: BaseProps, relationName: string, typeName: string) => {
  const { errors, lines, lineIndex, symbol, file, module } = props;
  errors.push(
    constructValidationError({
      message: `the partial relation definition \`${symbol}\` is a duplicate in the relation \`${relationName}\`.`,
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.DuplicatedError, relation: relationName, typeName, file, module },
    }),
  );
};

const createDuplicateRelationshipDefinitionError = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol, file, module } = props;

  if (!lines?.length || lineIndex === undefined) {
    errors.push(
      new ModelValidationSingleError(
        {
          msg: `duplicate relationship definition \`${symbol}\`.`,
          file,
        },
        { symbol, errorType: ValidationError.DuplicatedError, module },
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
          start: lineIndex,
          end: lineIndex,
        },
        column: {
          start: rawLine.indexOf(Keyword.DEFINE),
          end: rawLine.length,
        },
        file,
      },
      { symbol, errorType: ValidationError.DuplicatedError, module },
    ),
  );
};

const createAssignableTypeWildcardRelationError = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol, file, module, type, relation } = props;
  errors.push(
    constructValidationError({
      message: `type restriction \`${symbol}\` cannot contain both wildcard and relation`,
      lines,
      lineIndex,
      metadata: {
        symbol,
        errorType: ValidationError.TypeRestrictionCannotHaveWildcardAndRelation,
        relation,
        typeName: type,
        module,
        file,
      },
    }),
  );
};

const createInvalidRelationError = (props: BaseProps, validRelations: string[]) => {
  const { errors, lines, lineIndex, symbol, file, module, type, relation } = props;
  const isInValid = !validRelations?.includes(symbol);
  if (isInValid) {
    errors.push(
      constructValidationError({
        message: `the relation \`${symbol}\` does not exist.`,
        lines,
        lineIndex,
        metadata: { symbol, errorType: ValidationError.MissingDefinition, relation, file, module, typeName: type },
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

export const createSchemaVersionUnsupportedError = (props: BaseProps) => {
  const { errors, lines, lineIndex, symbol } = props;
  errors.push(
    constructValidationError({
      message: "schema version no longer supported",
      lines,
      lineIndex,
      metadata: { symbol, errorType: ValidationError.SchemaVersionUnsupported },
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

const createDifferentNestedConditionNameError = (props: BaseProps, condition: string, nestedConditionName: string) => {
  const { errors } = props;
  errors.push(
    constructValidationError({
      message: `condition key is \`${condition}\` but nested name property is ${nestedConditionName}`,
      metadata: { symbol: nestedConditionName, errorType: ValidationError.DifferentNestedConditionName },
    }),
  );
};

function createMultipleModuleInSingleFileError(props: BaseProps, file: string, modules: string[]) {
  const { errors } = props;
  errors.push(
    constructValidationError({
      message: `file ${file} would contain multiple module definitions (${modules.join(", ")}) when transforming to DSL. Only one module can be defined per file.`,
      metadata: { symbol: file, errorType: ValidationError.MultipleModulesInFile },
    }),
  );
}

function constructValidationError(props: ValidationErrorProps): ModelValidationSingleError {
  const { message, lines, lineIndex, customResolver, metadata } = props;

  const errorProps: ErrorProperties = {
    msg: message,
    file: metadata.file,
  };

  const errorMetadata: ErrorMetadata = {
    symbol: metadata?.symbol,
    errorType: metadata.errorType,
    module: metadata.module,
    relation: metadata.relation,
    type: metadata.typeName,
    offendingType: metadata.offendingType,
  };

  if (lines?.length && lineIndex != undefined) {
    const rawLine = lines[lineIndex];

    const re = new RegExp("\\b" + metadata.symbol + "\\b");
    let wordIdx = rawLine?.search(re);

    if (isNaN(wordIdx) || wordIdx === -1) {
      wordIdx = 0;
    }

    if (typeof customResolver === "function") {
      wordIdx = customResolver(wordIdx, rawLine, metadata.symbol);
    }

    errorProps.line = { start: lineIndex, end: lineIndex };
    errorProps.column = { start: wordIdx, end: wordIdx + (metadata.symbol?.length || 0) };
  }

  return new ModelValidationSingleError(errorProps, errorMetadata);
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

  raiseTupleUsersetRequiresDirect(symbol: string, type: string, relation: string, meta: Meta, lineIndex?: number) {
    createTupleUsersetRequireDirectError({
      errors: this.errors,
      lines: this.lines,
      lineIndex,
      symbol,
      file: meta.file,
      module: meta.module,
      relation,
      type,
    });
  }

  raiseDuplicateTypeName(symbol: string, meta: Meta, lineIndex?: number) {
    createDuplicateTypeNameError({
      errors: this.errors,
      lines: this.lines,
      lineIndex,
      symbol,
      module: meta.module,
      file: meta.file,
    });
  }

  raiseDuplicateTypeRestriction(
    symbol: string,
    relationName: string,
    typeName: string,
    meta: Meta,
    lineIndex?: number,
  ) {
    createDuplicateTypeRestrictionError(
      { errors: this.errors, lines: this.lines, lineIndex, symbol, module: meta.module, file: meta.file },
      relationName,
      typeName,
    );
  }

  raiseDuplicateType(symbol: string, relationName: string, typeName: string, meta: Meta, lineIndex?: number) {
    createDuplicateRelationError(
      { errors: this.errors, lines: this.lines, lineIndex, symbol, module: meta.module, file: meta.file },
      relationName,
      typeName,
    );
  }

  raiseDuplicateRelationshipDefinition(symbol: string, meta: Meta, lineIndex?: number) {
    createDuplicateRelationshipDefinitionError({
      errors: this.errors,
      lines: this.lines,
      lineIndex,
      symbol,
      module: meta.module,
      file: meta.file,
    });
  }

  raiseNoEntryPointLoop(symbol: string, typeName: string, meta: Meta, lineIndex?: number) {
    createNoEntryPointLoopError(
      { errors: this.errors, lines: this.lines, lineIndex, symbol, module: meta.module, file: meta.file },
      typeName,
    );
  }

  raiseNoEntryPoint(symbol: string, typeName: string, meta: Meta, lineIndex?: number) {
    createNoEntryPointError(
      { errors: this.errors, lines: this.lines, lineIndex, symbol, file: meta.file, module: meta.module },
      typeName,
    );
  }

  raiseInvalidRelationOnTupleset(
    symbol: string,
    typeName: string,
    typeDef: string,
    relationName: string,
    offendingRelation: string,
    parent: string,
    lineIndex?: number,
    meta?: Meta,
  ) {
    createInvalidRelationOnTuplesetError(
      {
        errors: this.errors,
        lines: this.lines,
        lineIndex,
        symbol,
        file: meta?.file,
        module: meta?.module,
      },
      typeName,
      typeDef,
      relationName,
      offendingRelation,
      parent,
    );
  }

  raiseInvalidTypeRelation(
    symbol: string,
    typeName: string,
    relationName: string,
    offendingRelation: string,
    offendingType: string,
    lineIndex?: number,
    meta?: Meta,
  ) {
    createInvalidTypeRelationError(
      {
        errors: this.errors,
        lines: this.lines,
        lineIndex,
        symbol,
        file: meta?.file,
        module: meta?.module,
      },
      typeName,
      relationName,
      offendingRelation,
      offendingType,
    );
  }

  raiseInvalidType(symbol: string, typeName: string, relation: string, meta: Meta, lineIndex?: number) {
    createInvalidTypeError(
      { errors: this.errors, lines: this.lines, lineIndex, symbol, module: meta.module, file: meta.file, relation },
      typeName,
    );
  }

  raiseAssignableRelationMustHaveTypes(symbol: string, lineIndex?: number) {
    createAssignableRelationMustHaveTypesError({ errors: this.errors, lines: this.lines, lineIndex, symbol });
  }

  raiseAssignableTypeWildcardRelation(symbol: string, type: string, relation: string, meta: Meta, lineIndex?: number) {
    createAssignableTypeWildcardRelationError({
      errors: this.errors,
      lines: this.lines,
      lineIndex,
      symbol,
      relation,
      type,
      module: meta.module,
      file: meta.file,
    });
  }

  raiseInvalidRelationError(
    symbol: string,
    type: string,
    relation: string,
    validRelations: string[],
    lineIndex?: number,
    metadata?: Meta,
  ) {
    createInvalidRelationError(
      {
        errors: this.errors,
        lines: this.lines,
        lineIndex,
        symbol,
        file: metadata?.file,
        module: metadata?.module,
        type,
        relation,
      },
      validRelations,
    );
  }

  raiseInvalidSchemaVersion(symbol: string, lineIndex?: number) {
    createInvalidSchemaVersionError({ errors: this.errors, lines: this.lines, lineIndex, symbol });
  }

  raiseSchemaVersionRequired(symbol: string, lineIndex?: number) {
    createSchemaVersionRequiredError({ errors: this.errors, lines: this.lines, lineIndex, symbol });
  }

  raiseSchemaVersionUnsupported(symbol: string, lineIndex?: number) {
    createSchemaVersionUnsupportedError({ errors: this.errors, lines: this.lines, lineIndex, symbol });
  }

  raiseMaximumOneDirectRelationship(symbol: string, lineIndex?: number) {
    createMaximumOneDirectRelationship({ errors: this.errors, lines: this.lines, lineIndex, symbol });
  }

  raiseInvalidConditionNameInParameter(
    symbol: string,
    typeName: string,
    relationName: string,
    conditionName: string,
    meta: Meta,
    lineIndex?: number,
  ) {
    createInvalidConditionNameInParameterError(
      { errors: this.errors, lines: this.lines, lineIndex, symbol, module: meta.module, file: meta.file },
      typeName,
      relationName,
      conditionName,
    );
  }

  raiseUnusedCondition(symbol: string, meta: Meta, lineIndex?: number) {
    createUnusedConditionError({
      errors: this.errors,
      lines: this.lines,
      lineIndex,
      symbol,
      module: meta.module,
      file: meta.file,
    });
  }

  raiseDifferentNestedConditionName(condition: string, nestedConditionName: string) {
    createDifferentNestedConditionNameError(
      {
        errors: this.errors,
        symbol: condition,
      },
      condition,
      nestedConditionName,
    );
  }

  raiseMultipleModulesInSingleFile(file: string, modules: Set<string>) {
    createMultipleModuleInSingleFileError(
      {
        errors: this.errors,
        lines: this.lines,
        symbol: file,
      },
      file,
      Array.from(modules),
    );
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
  let wordIdx = rawLine?.search(re);

  if (isNaN(wordIdx) || wordIdx === -1) {
    wordIdx = 0;
  }

  return new ModuleTransformationSingleError(
    {
      line: { start: lineIndex, end: lineIndex },
      column: { start: wordIdx, end: wordIdx + (metadata.symbol?.length || 0) },
      msg: message,
      file: metadata.file,
    },
    { symbol: metadata?.symbol },
  );
}
