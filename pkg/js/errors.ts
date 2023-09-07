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
    InvalidType = "invalid-type",
    RelationNoEntrypoint = "relation-no-entry-point",
    TuplesetNotDirect = "tupleuset-not-direct",
    DuplicatedError = "duplicated-error",
    RequireSchema1_0 = "allowed-type-schema-10",
    AssignableRelationsMustHaveType = "assignable-relation-must-have-type",
    AllowedTypesNotValidOnSchema1_0 = "allowed-type-not-valid-on-schema-1_0",
    InvalidSchema = "invalid-schema",
    InvalidSyntax = "invalid-syntax",
    TypeRestrictionCannotHaveWildcardAndRelation = "type-wildcard-relation",
}

export interface ErrorProperties {
    line: {
        start: number,
        end: number,
    },
    column: {
        start: number,
        end: number,
    },
    msg: string,
}

/**
 * Abstract base class for syntax and validation exceptions
 */ 
export abstract class BaseError extends Error {
    public line: { start: number; end: number; };
    public column: { start: number; end: number; };
    public msg: string;
    
    constructor(
        public properties: ErrorProperties,
        public type: string,
    ) {
        super(`${type} error at line=${properties.line.start}, column=${properties.column.start}: ${properties.msg}`);
        this.line = properties.line;
        this.column = properties.column;
        this.msg = properties.msg;
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
            symbol: string,
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
export class DSLSyntaxError extends Error {
    constructor(public errors: DSLSyntaxSingleError[]) {
        super(`${errors.length} error${errors.length > 1 ? "s" : ""} occurred:\n\t* ${errors.join("\n\t* ")}\n\n`);
        this.errors = errors;
    }

    toString() {
        return this.message;
    }
}

/**
 * Added to reporter as the JSON transformation is being parsed and validated
 */
export class ModelValidationSingleError extends BaseError {
    constructor(
        public properties: ErrorProperties,
        public metadata?: {
            symbol: string,
            errorType: ValidationError,
        },
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
export class ModelValidationError extends Error {
    constructor(public errors: ModelValidationSingleError[]) {
        super(`${errors.length} error${errors.length > 1 ? "s" : ""} occurred:\n\t* ${errors.join("\n\t* ")}\n\n`);
        this.errors = errors;
    }

    toString() {
        return this.message;
    }
}

/**
 * Thrown when improper values are passed.
 */
export class ConfigurationException extends Error {
    constructor(public message: string, public e: Error) {
        super(message);
        if (e?.stack) {
            this.stack = e.stack;
        }
    }
}