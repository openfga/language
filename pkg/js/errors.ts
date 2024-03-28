import { RecognitionException } from "antlr4";

export enum ValidationError {
  SchemaVersionRequired = "schema-version-required",
  SchemaVersionUnsupported = "schema-version-unsupported",
  ReservedTypeKeywords = "reserved-type-keywords",
  ReservedRelationKeywords = "reserved-relation-keywords",
  SelfError = "self-error",
  InvalidName = "invalid-name",
  MissingDefinition = "missing-definition",
  InvalidRelationType = "invalid-relation-type",
  InvalidRelationOnTupleset = "invalid-relation-on-tupleset",
  InvalidType = "invalid-type",
  RelationNoEntrypoint = "relation-no-entry-point",
  TuplesetNotDirect = "tupleuserset-not-direct",
  DuplicatedError = "duplicated-error",
  RequireSchema1_0 = "allowed-type-schema-10",
  AssignableRelationsMustHaveType = "assignable-relation-must-have-type",
  AllowedTypesNotValidOnSchema1_0 = "allowed-type-not-valid-on-schema-1_0",
  InvalidSchema = "invalid-schema",
  InvalidSyntax = "invalid-syntax",
  TypeRestrictionCannotHaveWildcardAndRelation = "type-wildcard-relation",
  ConditionNotDefined = "condition-not-defined",
  ConditionNotUsed = "condition-not-used",
  DifferentNestedConditionName = "different-nested-condition-name",
}

export interface ErrorProperties {
  line?: {
    start: number;
    end: number;
  };
  column?: {
    start: number;
    end: number;
  };
  msg: string;
  file?: string;
}

export interface ErrorMetadata {
  symbol: string;
  errorType: ValidationError;
  module?: string;
  type?: string;
  relation?: string;
  condition?: string;
}

/**
 * Abstract base class for syntax and validation exceptions
 */
export abstract class BaseError extends Error {
  public line: { start: number; end: number } | undefined;
  public column: { start: number; end: number } | undefined;
  public file: string | undefined;
  public msg: string;

  constructor(
    public properties: ErrorProperties,
    public type: string,
  ) {
    super(
      `${type} error${
        properties.line !== undefined && properties.column !== undefined
          ? ` at line=${properties.line.start}, column=${properties.column.start}`
          : ""
      }: ${properties.msg}`,
    );
    this.line = properties.line;
    this.column = properties.column;
    this.msg = properties.msg;
    this.file = properties.file;
  }

  toString() {
    return this.message;
  }
}

/**
 * Abstract base class for errors that collate other errors.
 */
export abstract class BaseMultiError<T = BaseError> extends Error {
  constructor(public errors: Array<T>) {
    super(`${errors.length} error${errors.length > 1 ? "s" : ""} occurred:\n\t* ${errors.join("\n\t* ")}\n\n`);
    this.errors = errors;
  }

  toString() {
    return this.message;
  }
}

/**
 * Added to listener during syntax parsing, when syntax errors are encountered
 */
export class DSLSyntaxSingleError extends BaseError {
  constructor(
    public properties: ErrorProperties,
    public metadata?: {
      symbol: string;
    },
    e?: RecognitionException,
  ) {
    super(properties, "syntax");
    if (e?.stack) {
      this.stack = e.stack;
    }
    this.metadata = metadata;
  }

  toString() {
    return this.message;
  }
}

/**
 * Thrown at the end of syntax parsing, collecting all Syntax errors encountered during parsing
 */
export class DSLSyntaxError extends BaseMultiError<DSLSyntaxSingleError> {}

/**
 * Added to reporter as the JSON transformation is being parsed and validated
 */
export class ModelValidationSingleError extends BaseError {
  constructor(
    public properties: ErrorProperties,
    public metadata?: ErrorMetadata,
  ) {
    super(properties, metadata?.errorType || "validation");
    this.metadata = metadata;
  }

  toString() {
    return this.message;
  }
}

/**
 * Thrown at end of checkDSL validation, collecting all encountered validation errors
 */
export class ModelValidationError extends BaseMultiError<ModelValidationSingleError> {}

/**
 * Thrown when improper values are passed.
 */
export class ConfigurationError extends Error {
  constructor(
    public message: string,
    public e: Error,
  ) {
    super(message);
    if (e?.stack) {
      this.stack = e.stack;
    }
  }
}

export class UnsupportedDSLNestingError extends Error {
  constructor(
    public typeName: string,
    public relationName: string,
  ) {
    super(
      `the '${relationName}' relation definition under the '${typeName}' type is not supported by the OpenFGA DSL syntax yet`,
    );
  }
}

export class ConditionNameDoesntMatchError extends Error {
  constructor(
    public conditionName: string,
    public conditionNestedName: string,
  ) {
    super(`the '${conditionName}' condition has a different nested condition name ('${conditionNestedName}')`);
  }
}

/**
 * Represents an individual error returned during validation of `fga.mod`.
 * Line and column numbers returned as part of this are one based.
 */
export class FGAModFileValidationSingleError extends BaseError {
  constructor(public properties: ErrorProperties) {
    super(properties, "validation");
  }

  toString() {
    return this.message;
  }
}

/**
 * Thrown when an `fga.mod` file is invalid.
 */
export class FGAModFileValidationError extends BaseMultiError {}

/*
 * Represents an individual error returned during transformation of a module.
 * Line and column numbers returned as part of this are one based.
 */
export class ModuleTransformationSingleError extends BaseError {
  constructor(
    public properties: ErrorProperties,
    public metadata?: {
      symbol: string;
    },
  ) {
    super(properties, "transformation-error");
    this.metadata = metadata;
  }

  toString() {
    return this.message;
  }
}

/**
 * Thrown when a module is invalid.
 */
export class ModuleTransformationError extends BaseMultiError {}
