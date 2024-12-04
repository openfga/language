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
  MixinContext,
  MixinDeclarationContext,
  MixinsContext,
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

interface Mixin {
  name: string;
  relations: Map<string, Partial<Relation>>;
}

/**
 * This Visitor walks the tree generated by parsers and produces Python code
 *
 * @returns {object}
 */
class OpenFgaDslListener extends OpenFGAListener {
  public authorizationModel: Partial<AuthorizationModel> = {};
  public typeDefExtensions: Map<string, TypeDefinition> = new Map();

  private mixins : Map<string, Mixin> = new Map();
  private currentTypeDef: Partial<TypeDefinition> | undefined;
  private currentRelation: Partial<Relation> | undefined;
  private currentCondition: Condition | undefined;
  private currentMixin: Mixin | undefined;
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

    this.currentTypeDef = {
      type: ctx._typeName.getText(),
      relations: {},
      metadata: { relations: {} },
    };

    if (this.isModularModel) {
      this.currentTypeDef.metadata!.module = this.moduleName;
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

  enterMixins = (ctx: MixinsContext) => {
    this.mixins = new Map();
  }

  exitMixins = () => {
  }

  /*
    mixin foo
      relations
        define relationship1 : [user]
  */
  enterMixin = (ctx: MixinContext) => {
    const mixinName = ctx.mixinName().getText();
    
    // Cannot define multiple mixinx with the same name
    if (this.mixins.has(mixinName)) {
      ctx.parser?.notifyErrorListeners(`mixin '${mixinName}' is already defined`, ctx.mixinName().start, undefined);
    }

    const mixin: Mixin = {
      name: mixinName,
      relations: new Map()
    }

    this.mixins.set(mixin.name, mixin);
    this.currentMixin = mixin;
  }

  exitMixin = () => {
    this.currentMixin = undefined;
  }

  // include foo
  enterMixinDeclaration = (ctx: MixinDeclarationContext) => {
    const mixinName = ctx.mixinName().getText();
    const mixin = this.mixins.get(mixinName);

    if (!mixin) {
      ctx.parser?.notifyErrorListeners(`mixin '${mixinName}' is not defined`, ctx.mixinName().start, undefined);
    }

    if (mixin?.relations) {
      mixin?.relations.forEach((relation, relationName) => {
        const relationDef = parseExpression(relation.rewrites, relation.operator);

        if (relationDef) {
          this.currentTypeDef!.relations![relationName] = relationDef;
          const directlyRelatedUserTypes = relation.typeInfo?.directly_related_user_types;

          this.currentTypeDef!.metadata!.relations![relationName] = {
            directly_related_user_types: directlyRelatedUserTypes,
            mixin: mixinName,
          };

          // Only add the module name for a relation when we're parsing an extended type
          if (this.isModularModel && (ctx.parentCtx as TypeDefContext).EXTEND()) {
            this.currentTypeDef!.metadata!.relations![relationName].module = this.moduleName;
          }
        }
      });
    }
  }

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
      if (this.currentMixin) {
        if (this.currentMixin.relations.has(relationName)) {
          // Throw error if same named relation occurs more than once in a mixin block
          ctx.parser?.notifyErrorListeners(
            `'${relationName}' is already defined in mixin '${this.currentMixin.name}'`,
            ctx.relationName().start,
            undefined,
          );
        }

        if (this.currentRelation) {
          this.currentMixin.relations.set(relationName, this.currentRelation);
        }
      } else if (this.currentTypeDef) {
        if (this.currentTypeDef!.relations![relationName]) {
          // Throw error if same named relation occurs more than once in a relationship definition block
          if (this.currentTypeDef!.metadata!.relations![relationName]?.mixin) {
            // .. but if it's a duplicate because it was defined in a mixin, tweak the message
            ctx.parser?.notifyErrorListeners(
              `'${relationName}' is already defined in mixin '${this.currentTypeDef!.metadata!.relations![relationName].mixin}'`,
              ctx.relationName().start,
              undefined,
            );
          } else  {
            ctx.parser?.notifyErrorListeners(
              `'${relationName}' is already defined in '${this.currentTypeDef?.type}'`,
              ctx.relationName().start,
              undefined,
            );
          }
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

export function parseDSL(data: string): {
  listener: OpenFgaDslListener;
  errorListener: OpenFgaDslErrorListener<unknown>;
} {
  const cleanedData = data
    .split("\n")
    .map((line) => {
      if (line.trimStart()[0] === "#") {
        return "";
      }

      return line.split(" #")[0].trimEnd();
    })
    .join("\n");

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
  new antlr.ParseTreeWalker().walk(listener, parser.main());

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