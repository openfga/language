import type {
  AuthorizationModel,
  Condition,
  ConditionParamTypeRef,
  ObjectRelation,
  RelationMetadata,
  RelationReference,
  TypeDefinition,
  Userset,
} from "@openfga/sdk";
import * as antlr from "antlr4";
import { ErrorListener, RecognitionException, Recognizer } from "antlr4";
import OpenFGAListener from "../gen/OpenFGAParserListener";
import OpenFGALexer from "../gen/OpenFGALexer";
import OpenFGAParser, {
  ConditionContext,
  ConditionExpressionContext,
  ConditionParameterContext,
  ModelHeaderContext,
  ModuleHeaderContext,
  RelationDeclarationContext,
  RelationDefDirectAssignmentContext,
  RelationDefPartialsContext,
  RelationDefRewriteContext,
  RelationDefTypeRestrictionContext,
  TypeDefContext,
  TypeDefsContext,
} from "../gen/OpenFGAParser";
import { DSLSyntaxError, DSLSyntaxSingleError } from "../errors";
import { TypeName } from "@openfga/sdk";

// Comment tracking types
export interface CommentInfo {
  preceding_lines?: string[];
  inline?: string;
}

export interface CommentsMetadata {
  comments?: CommentInfo;
  relation_comments?: Record<string, CommentInfo>;
}

export interface ModelComments {
  preceding_lines?: string[];
}

// Comment Tracker class for extracting comments from DSL source
class CommentTracker {
  private lines: string[];
  private cleanedToOriginal?: Record<number, number>;

  constructor(source: string, cleanedToOriginal?: Record<number, number>) {
    this.lines = source.split("\n");
    this.cleanedToOriginal = cleanedToOriginal;
  }

  // Get comments that immediately precede the given line number (0-based)
  getPrecedingComments(lineNum: number): string[] {
    if (lineNum <= 0 || lineNum > this.lines.length) {
      return [];
    }

    const comments: string[] = [];
    // Walk backwards from the line before the target
    for (let i = lineNum - 1; i >= 0; i--) {
      const line = this.lines[i].trim();
      if (line.length === 0) {
        // Empty line breaks the contiguous comment block
        break;
      }
      if (line.startsWith("#")) {
        // Prepend to maintain order
        comments.unshift(line);
      } else {
        // Non-comment, non-empty line breaks the block
        break;
      }
    }

    return comments;
  }

  // Get the inline comment for the given line number (0-based)
  getInlineComment(lineNum: number): string {
    if (lineNum < 0 || lineNum >= this.lines.length) {
      return "";
    }

    const line = this.lines[lineNum];
    // Find inline comment
    const inlineIdx = line.indexOf(" #");
    if (inlineIdx !== -1) {
      const beforeComment = line.slice(0, inlineIdx).trim();
      if (beforeComment.length > 0) {
        return line.slice(inlineIdx + 1).trim();
      }
    }
    return "";
  }

  // Get comment info for an element at the given line
  // lineNum is expected to be 0-based and may be from cleaned source (if mapping exists)
  getCommentInfoForLine(lineNum: number): CommentInfo | undefined {
    // If we have a line mapping, convert from cleaned line number to original
    let originalLineNum = lineNum;
    if (this.cleanedToOriginal !== undefined && lineNum in this.cleanedToOriginal) {
      originalLineNum = this.cleanedToOriginal[lineNum];
    }

    const preceding = this.getPrecedingComments(originalLineNum);
    const inline = this.getInlineComment(originalLineNum);

    if (preceding.length === 0 && !inline) {
      return undefined;
    }

    return {
      preceding_lines: preceding.length > 0 ? preceding : undefined,
      inline: inline || undefined,
    };
  }

  // Get model comments (comments before the model declaration)
  getModelComments(): ModelComments | undefined {
    // Find the model declaration line
    let modelLine = -1;
    for (let i = 0; i < this.lines.length; i++) {
      const trimmed = this.lines[i].trim();
      if (trimmed.startsWith("model")) {
        modelLine = i;
        break;
      }
    }

    if (modelLine <= 0) {
      return undefined;
    }

    const precedingComments = this.getPrecedingComments(modelLine);
    if (precedingComments.length === 0) {
      return undefined;
    }

    return {
      preceding_lines: precedingComments,
    };
  }
}

enum RelationDefinitionOperator {
  RELATION_DEFINITION_OPERATOR_NONE = "",
  RELATION_DEFINITION_OPERATOR_OR = "or",
  RELATION_DEFINITION_OPERATOR_AND = "and",
  RELATION_DEFINITION_OPERATOR_BUT_NOT = "but not",
}

type RelationTypeInfo = RelationMetadata;

interface Relation {
  name: string;
  rewrites: Userset[];
  operator: RelationDefinitionOperator;
  typeInfo: RelationTypeInfo;
}

function parseExpression(
  rewrites: Userset[] | undefined,
  operator: RelationDefinitionOperator | undefined,
): Userset | undefined {
  let relationDef: Userset | undefined;

  if (!rewrites?.length) {
    return;
  }
  if (rewrites?.length === 1) {
    relationDef = rewrites[0];
  } else {
    switch (operator) {
      case RelationDefinitionOperator.RELATION_DEFINITION_OPERATOR_OR:
        relationDef = {
          union: {
            child: rewrites,
          },
        };
        break;
      case RelationDefinitionOperator.RELATION_DEFINITION_OPERATOR_AND:
        relationDef = {
          intersection: {
            child: rewrites,
          },
        };
        break;
      case RelationDefinitionOperator.RELATION_DEFINITION_OPERATOR_BUT_NOT:
        relationDef = {
          difference: {
            base: rewrites.shift()!,
            subtract: rewrites.shift()!,
          },
        };
        break;
    }
  }
  return relationDef;
}

interface StackRelation {
  rewrites: Userset[];
  operator: RelationDefinitionOperator;
}

/**
 * This Visitor walks the tree generated by parsers and produces the authorization model
 *
 * @returns {object}
 */
class OpenFgaDslListener extends OpenFGAListener {
  public authorizationModel: Partial<AuthorizationModel> = {};
  public typeDefExtensions: Map<string, TypeDefinition> = new Map();

  // Comment tracking
  public commentTracker?: CommentTracker;
  public modelComments?: ModelComments;
  public typeComments: Record<string, CommentsMetadata> = {};
  public conditionComments: Record<string, CommentInfo> = {};

  private currentTypeDef: Partial<TypeDefinition> | undefined;
  private currentRelation: Partial<Relation> | undefined;
  private currentCondition: Condition | undefined;
  private isModularModel = false;
  private moduleName?: string;

  private rewriteStack: StackRelation[] = [];

  exitModuleHeader = (ctx: ModuleHeaderContext) => {
    if (!ctx._moduleName) {
      return;
    }

    this.isModularModel = true;
    this.moduleName = ctx._moduleName.getText();
  };

  exitModelHeader = (ctx: ModelHeaderContext) => {
    if (ctx.SCHEMA_VERSION()) {
      this.authorizationModel.schema_version = ctx.SCHEMA_VERSION().getText();
    }
  };

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  enterTypeDefs = (_ctx: TypeDefsContext) => {
    this.authorizationModel.type_definitions = [];
  };

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  exitTypeDefs = (_ctx: TypeDefsContext) => {
    if (!this.authorizationModel.type_definitions?.length) {
      delete this.authorizationModel.type_definitions;
    }
  };

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  enterTypeDef = (ctx: TypeDefContext) => {
    if (!ctx._typeName) {
      return;
    }

    if (ctx.EXTEND() && !this.isModularModel) {
      ctx.parser?.notifyErrorListeners("extend can only be used in a modular model", ctx._typeName.start, undefined);
    }

    const typeName = ctx._typeName.getText();
    this.currentTypeDef = {
      type: typeName,
      relations: {},
      metadata: { relations: {} },
    };

    if (this.isModularModel) {
      this.currentTypeDef.metadata!.module = this.moduleName;
    }

    // Track type comments (line is 1-based from ANTLR, convert to 0-based)
    // Use TYPE token's line number instead of ctx.start which may include preceding newlines
    if (this.commentTracker && ctx.TYPE()) {
      const lineNum = ctx.TYPE().symbol.line - 1;
      const commentInfo = this.commentTracker.getCommentInfoForLine(lineNum);
      if (commentInfo) {
        if (!this.typeComments[typeName]) {
          this.typeComments[typeName] = {};
        }
        this.typeComments[typeName].comments = commentInfo;
      }
    }
  };

  exitTypeDef = (ctx: TypeDefContext) => {
    if (!this.currentTypeDef?.type) {
      return;
    }

    if (this.isModularModel && !Object.keys(this.currentTypeDef?.metadata?.relations || {}).length) {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      this.currentTypeDef!.metadata!.relations = undefined as any;
    } else if (!this.isModularModel && !Object.keys(this.currentTypeDef?.metadata?.relations || {}).length) {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      this.currentTypeDef!.metadata = null as any;
    }

    this.authorizationModel.type_definitions?.push(this.currentTypeDef as TypeDefinition);

    if (ctx.EXTEND() && this.isModularModel) {
      if (this.typeDefExtensions.has(this.currentTypeDef.type)) {
        ctx.parser?.notifyErrorListeners(
          `'${this.currentTypeDef.type}' is already extended in file.`,
          ctx._typeName.start,
          undefined,
        );
      } else {
        this.typeDefExtensions.set(this.currentTypeDef.type, this.currentTypeDef as TypeDefinition);
      }
    }

    this.currentTypeDef = undefined;
  };

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  enterRelationDeclaration = (_ctx: RelationDeclarationContext) => {
    this.currentRelation = {
      rewrites: [],
      typeInfo: { directly_related_user_types: [] },
    };

    this.rewriteStack = [];
  };

  exitRelationDeclaration = (ctx: RelationDeclarationContext) => {
    if (!ctx.relationName()) {
      return;
    }

    const relationName = ctx.relationName().getText();
    const rewrites = this.currentRelation?.rewrites;

    const relationDef = parseExpression(rewrites, this.currentRelation?.operator);

    if (relationDef) {
      // Throw error if same named relation occurs more than once in a relationship definition block
      if (this.currentTypeDef!.relations![relationName]) {
        ctx.parser?.notifyErrorListeners(
          `'${relationName}' is already defined in '${this.currentTypeDef?.type}'`,
          ctx.relationName().start,
          undefined,
        );
      }

      this.currentTypeDef!.relations![relationName] = relationDef;
      const directlyRelatedUserTypes = this.currentRelation?.typeInfo?.directly_related_user_types;
      this.currentTypeDef!.metadata!.relations![relationName] = {
        directly_related_user_types: directlyRelatedUserTypes,
      };

      // Only add the module name for a relation when we're parsing an extended type
      if (this.isModularModel && (ctx.parentCtx as TypeDefContext).EXTEND()) {
        this.currentTypeDef!.metadata!.relations![relationName].module = this.moduleName;
      }

      // Track relation comments (line is 1-based from ANTLR, convert to 0-based)
      // Use DEFINE token's line number instead of ctx.start which may include preceding newlines
      if (this.commentTracker && this.currentTypeDef && ctx.DEFINE()) {
        const typeName = this.currentTypeDef.type!;
        const lineNum = ctx.DEFINE().symbol.line - 1;
        const commentInfo = this.commentTracker.getCommentInfoForLine(lineNum);
        if (commentInfo) {
          if (!this.typeComments[typeName]) {
            this.typeComments[typeName] = {};
          }
          if (!this.typeComments[typeName].relation_comments) {
            this.typeComments[typeName].relation_comments = {};
          }
          this.typeComments[typeName].relation_comments![relationName] = commentInfo;
        }
      }
    }

    this.currentRelation = undefined;
  };

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  enterRelationDefDirectAssignment = (_ctx: RelationDefDirectAssignmentContext) => {
    this.currentRelation!.typeInfo = { directly_related_user_types: [] };
  };

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  exitRelationDefDirectAssignment = (_ctx: RelationDefDirectAssignmentContext) => {
    const partialRewrite: Userset = {
      this: {},
    };
    this.currentRelation?.rewrites?.push(partialRewrite);
  };

  exitRelationDefTypeRestriction = (ctx: RelationDefTypeRestrictionContext) => {
    const relationRef: Partial<RelationReference> = {};
    const baseRestriction = ctx.relationDefTypeRestrictionBase();
    if (!baseRestriction) {
      return;
    }

    relationRef.type = baseRestriction._relationDefTypeRestrictionType?.getText();
    const usersetRestriction = baseRestriction._relationDefTypeRestrictionRelation;
    const wildcardRestriction = baseRestriction._relationDefTypeRestrictionWildcard;

    if (ctx.conditionName()) {
      relationRef.condition = ctx.conditionName().getText();
    }

    if (usersetRestriction) {
      relationRef.relation = usersetRestriction.getText();
    }

    if (wildcardRestriction) {
      relationRef.wildcard = {};
    }

    this.currentRelation!.typeInfo!.directly_related_user_types!.push(relationRef as RelationReference);
  };

  exitRelationDefRewrite = (ctx: RelationDefRewriteContext) => {
    let partialRewrite: Userset = {
      computedUserset: {
        relation: ctx._rewriteComputedusersetName.getText(),
      },
    };

    if (ctx._rewriteTuplesetName) {
      partialRewrite = {
        tupleToUserset: {
          ...(partialRewrite as { computedUserset: ObjectRelation }),
          tupleset: {
            relation: ctx._rewriteTuplesetName.getText(),
          },
        },
      };
    }

    this.currentRelation?.rewrites?.push(partialRewrite);
  };

  exitRelationRecurse = () => {
    const rewrites = this.currentRelation?.rewrites;

    const relationDef = parseExpression(rewrites, this.currentRelation?.operator);

    if (relationDef) {
      this.currentRelation!.rewrites = [relationDef];
    }
  };

  enterRelationRecurseNoDirect = () => {
    this.rewriteStack?.push({
      rewrites: this.currentRelation!.rewrites!,
      operator: this.currentRelation!.operator!,
    });

    this.currentRelation!.rewrites = [];
  };

  exitRelationRecurseNoDirect = () => {
    const rewrites = this.currentRelation?.rewrites;

    const relationDef = parseExpression(rewrites, this.currentRelation?.operator);

    const popped = this.rewriteStack.pop();

    if (relationDef) {
      this.currentRelation!.operator = popped?.operator;
      this.currentRelation!.rewrites = [...popped!.rewrites, relationDef];
    }
  };

  enterRelationDefPartials = (ctx: RelationDefPartialsContext) => {
    if (ctx.OR_list().length) {
      this.currentRelation!.operator = RelationDefinitionOperator.RELATION_DEFINITION_OPERATOR_OR;
    } else if (ctx.AND_list().length) {
      this.currentRelation!.operator = RelationDefinitionOperator.RELATION_DEFINITION_OPERATOR_AND;
    } else if (ctx.BUT_NOT()) {
      this.currentRelation!.operator = RelationDefinitionOperator.RELATION_DEFINITION_OPERATOR_BUT_NOT;
    }
  };

  enterCondition = (ctx: ConditionContext) => {
    if (ctx.conditionName() === null) {
      return;
    }
    if (!this.authorizationModel.conditions) {
      this.authorizationModel.conditions = {};
    }

    const conditionName = ctx.conditionName().getText();
    if (this.authorizationModel.conditions![conditionName]) {
      ctx.parser?.notifyErrorListeners(
        `condition '${conditionName}' is already defined in the model`,
        ctx.conditionName().start,
        undefined,
      );
    }

    this.currentCondition = {
      name: conditionName,
      expression: "",
      parameters: {},
    };

    if (this.isModularModel) {
      this.currentCondition.metadata = {
        module: this.moduleName,
      };
    }

    // Track condition comments (line is 1-based from ANTLR, convert to 0-based)
    // Use CONDITION token's line number instead of ctx.start which may include preceding newlines
    if (this.commentTracker && ctx.CONDITION()) {
      const lineNum = ctx.CONDITION().symbol.line - 1;
      const commentInfo = this.commentTracker.getCommentInfoForLine(lineNum);
      if (commentInfo) {
        this.conditionComments[conditionName] = commentInfo;
      }
    }
  };

  exitConditionParameter = (ctx: ConditionParameterContext) => {
    if (!ctx.parameterName() || !ctx.parameterType()) {
      return;
    }

    const parameterName = ctx.parameterName().getText();
    if (this.currentCondition?.parameters?.[parameterName]) {
      ctx.parser?.notifyErrorListeners(
        `parameter '${parameterName}' is already defined in the condition '${this.currentCondition?.name}'`,
        ctx.parameterName().start,
        undefined,
      );
    }

    const paramContainer = ctx.parameterType().CONDITION_PARAM_CONTAINER();
    const conditionParamTypeRef: Partial<ConditionParamTypeRef> = {};
    if (paramContainer) {
      conditionParamTypeRef.type_name = `TYPE_NAME_${paramContainer.getText().toUpperCase()}` as TypeName;
      const genericTypeName =
        ctx.parameterType().CONDITION_PARAM_TYPE() &&
        (`TYPE_NAME_${ctx.parameterType().CONDITION_PARAM_TYPE().getText().toUpperCase()}` as TypeName);
      if (genericTypeName) {
        conditionParamTypeRef.generic_types = [{ type_name: genericTypeName }];
      }
    } else {
      conditionParamTypeRef.type_name = `TYPE_NAME_${ctx.parameterType().getText().toUpperCase()}` as TypeName;
    }

    this.currentCondition!.parameters![parameterName] = conditionParamTypeRef as ConditionParamTypeRef;
  };

  exitConditionExpression = (ctx: ConditionExpressionContext) => {
    this.currentCondition!.expression = ctx.getText().trim();
  };

  exitCondition = () => {
    if (this.currentCondition) {
      this.authorizationModel.conditions![this.currentCondition.name!] = this.currentCondition!;

      this.currentCondition = undefined;
    }
  };
}

class OpenFgaDslErrorListener<T> extends ErrorListener<T> {
  errors: DSLSyntaxSingleError[] = [];

  syntaxError(
    _recognizer: Recognizer<T>,
    offendingSymbol: T,
    line: number, // line is one based, i.e. the first line will be 1
    column: number, // column is zero based, i.e. the first column will be 0
    msg: string,
    e: RecognitionException | undefined,
  ) {
    let metadata = undefined;
    let columnOffset = 0;

    if (offendingSymbol instanceof antlr.Token) {
      metadata = {
        symbol: offendingSymbol.text,
      };
      columnOffset = metadata.symbol.length;
    }

    this.errors.push(
      new DSLSyntaxSingleError(
        {
          line: { start: line - 1, end: line - 1 },
          column: { start: column, end: column + columnOffset },
          msg,
        },
        metadata,
        e,
      ),
    );
  }
}

export function parseDSL(
  data: string,
  preserveComments = true,
): {
  listener: OpenFgaDslListener;
  errorListener: OpenFgaDslErrorListener<unknown>;
} {
  const originalLines = data.split("\n");

  // Build a mapping from cleaned line numbers to original line numbers
  // (both 0-based indices)
  const cleanedToOriginal: Record<number, number> = {};
  const cleanedLines: string[] = [];

  for (let originalIdx = 0; originalIdx < originalLines.length; originalIdx++) {
    const line = originalLines[originalIdx];
    let cleanedLine = "";

    if (line.trimStart().length === 0) {
      // Empty line
    } else if (line.trimStart()[0] === "#") {
      cleanedLine = "";
    } else {
      cleanedLine = line.split(" #")[0].trimEnd();
    }

    const cleanedIdx = cleanedLines.length;
    cleanedLines.push(cleanedLine);
    cleanedToOriginal[cleanedIdx] = originalIdx;
  }

  // Create comment tracker with original data and line mapping
  const commentTracker = preserveComments ? new CommentTracker(data, cleanedToOriginal) : undefined;

  const cleanedData = cleanedLines.join("\n");

  const is = new antlr.InputStream(cleanedData);

  const errorListener = new OpenFgaDslErrorListener();

  // Create the Lexer
  const lexer = new OpenFGALexer(is as antlr.CharStream);
  lexer.removeErrorListeners();
  lexer.addErrorListener(errorListener);
  const stream = new antlr.CommonTokenStream(lexer);

  // Create the Parser
  const parser = new OpenFGAParser(stream);
  parser.removeErrorListeners();
  parser.addErrorListener(errorListener);

  // Finally parse the expression
  const listener = new OpenFgaDslListener();
  listener.commentTracker = commentTracker;

  new antlr.ParseTreeWalker().walk(listener, parser.main());

  // Extract model comments after parsing
  if (preserveComments && commentTracker) {
    listener.modelComments = commentTracker.getModelComments();
  }

  return { listener, errorListener };
}

/**
 * transformDSLToJSONObject - Converts models authored in FGA DSL syntax to the json syntax accepted by the OpenFGA API
 * @param {string} data
 * @returns {AuthorizationModel}
 */
export function transformDSLToJSONObject(data: string): Omit<AuthorizationModel, "id"> {
  const { listener, errorListener } = parseDSL(data);

  if (errorListener.errors.length) {
    throw new DSLSyntaxError(errorListener.errors);
  }

  return listener.authorizationModel as Omit<AuthorizationModel, "id">;
}

/**
 * transformDSLToJSONObject - Converts models authored in FGA DSL syntax to a stringified json representation
 * @param {string} data
 * @returns {string}
 */
export function transformDSLToJSON(data: string): string {
  return JSON.stringify(transformDSLToJSONObject(data));
}

interface ModularDSLTransformResult {
  authorizationModel: Omit<AuthorizationModel, "id">;
  typeDefExtensions: Map<string, TypeDefinition>;
}

/**
 * transformModularDSLToJSONObject - Converts a part of a modular model in DSL syntax to the json syntax accepted by
 * OpenFGA API and also returns the type definitions that are extended in the DSL if any are.
 * @internal
 * @param {string} data
 * @returns {ModularDSLTransformResult}
 */
export function transformModularDSLToJSONObject(data: string): ModularDSLTransformResult {
  const { listener, errorListener } = parseDSL(data);

  if (errorListener.errors.length) {
    throw new DSLSyntaxError(errorListener.errors);
  }

  return {
    authorizationModel: listener.authorizationModel as Omit<AuthorizationModel, "id">,
    typeDefExtensions: listener.typeDefExtensions,
  };
}

// Interface for DSL to JSON transformation result with comments
export interface TransformResultWithComments {
  authorizationModel: Omit<AuthorizationModel, "id">;
  modelComments?: ModelComments;
  typeComments: Record<string, CommentsMetadata>;
  conditionComments: Record<string, CommentInfo>;
}

/**
 * transformDSLToJSONObjectWithComments - Converts models authored in FGA DSL syntax to the json syntax
 * while preserving comments metadata.
 * @param {string} data
 * @returns {TransformResultWithComments}
 */
export function transformDSLToJSONObjectWithComments(data: string): TransformResultWithComments {
  const { listener, errorListener } = parseDSL(data, true);

  if (errorListener.errors.length) {
    throw new DSLSyntaxError(errorListener.errors);
  }

  return {
    authorizationModel: listener.authorizationModel as Omit<AuthorizationModel, "id">,
    modelComments: listener.modelComments,
    typeComments: listener.typeComments,
    conditionComments: listener.conditionComments,
  };
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
type JSONValue = string | number | boolean | null | JSONValue[] | { [key: string]: JSONValue };

/**
 * transformDSLToJSONWithComments - Converts models authored in FGA DSL syntax to a stringified json representation
 * with comments embedded in the metadata.
 * @param {string} data
 * @returns {string}
 */
export function transformDSLToJSONWithComments(data: string): string {
  const result = transformDSLToJSONObjectWithComments(data);

  // Convert to JSON and inject comments
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const jsonObj: Record<string, JSONValue> = result.authorizationModel as any;

  // Add model comments to top-level metadata
  if (result.modelComments && result.modelComments.preceding_lines?.length) {
    if (!jsonObj.metadata) {
      jsonObj.metadata = {};
    }
    (jsonObj.metadata as Record<string, JSONValue>).model_comments = {
      preceding_lines: result.modelComments.preceding_lines,
    };
  }

  // Add type and relation comments
  const typeDefinitions = jsonObj.type_definitions as Record<string, JSONValue>[] | undefined;
  if (typeDefinitions) {
    for (const typeDef of typeDefinitions) {
      const typeName = typeDef.type as string;
      const typeComments = result.typeComments[typeName];
      if (!typeComments) continue;

      // Ensure metadata exists
      if (!typeDef.metadata) {
        typeDef.metadata = {};
      }
      const metadata = typeDef.metadata as Record<string, JSONValue>;

      // Add type-level comments
      if (typeComments.comments) {
        const commentObj: Record<string, JSONValue> = {};
        if (typeComments.comments.preceding_lines?.length) {
          commentObj.preceding_lines = typeComments.comments.preceding_lines;
        }
        if (typeComments.comments.inline) {
          commentObj.inline = typeComments.comments.inline;
        }
        if (Object.keys(commentObj).length > 0) {
          metadata.comments = commentObj;
        }
      }

      // Add relation comments
      if (typeComments.relation_comments) {
        if (!metadata.relations) {
          metadata.relations = {};
        }
        const relationsMetadata = metadata.relations as Record<string, JSONValue>;

        for (const [relationName, relationComments] of Object.entries(typeComments.relation_comments)) {
          if (!relationsMetadata[relationName]) {
            relationsMetadata[relationName] = {};
          }
          const relationMeta = relationsMetadata[relationName] as Record<string, JSONValue>;

          const commentObj: Record<string, JSONValue> = {};
          if (relationComments.preceding_lines?.length) {
            commentObj.preceding_lines = relationComments.preceding_lines;
          }
          if (relationComments.inline) {
            commentObj.inline = relationComments.inline;
          }
          if (Object.keys(commentObj).length > 0) {
            relationMeta.comments = commentObj;
          }
        }
      }
    }
  }

  // Add condition comments
  const conditions = jsonObj.conditions as Record<string, Record<string, JSONValue>> | undefined;
  if (conditions) {
    for (const [conditionName, condComments] of Object.entries(result.conditionComments)) {
      const conditionData = conditions[conditionName];
      if (!conditionData) continue;

      // Ensure metadata exists
      if (!conditionData.metadata) {
        conditionData.metadata = {};
      }
      const metadata = conditionData.metadata as Record<string, JSONValue>;

      const commentObj: Record<string, JSONValue> = {};
      if (condComments.preceding_lines?.length) {
        commentObj.preceding_lines = condComments.preceding_lines;
      }
      if (condComments.inline) {
        commentObj.inline = condComments.inline;
      }
      if (Object.keys(commentObj).length > 0) {
        metadata.comments = commentObj;
      }
    }
  }

  return JSON.stringify(jsonObj);
}
