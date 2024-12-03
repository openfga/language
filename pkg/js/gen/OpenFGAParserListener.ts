// Generated from /app/OpenFGAParser.g4 by ANTLR 4.13.1

import {ParseTreeListener} from "antlr4";


import { MainContext } from "./OpenFGAParser";
import { ModelHeaderContext } from "./OpenFGAParser";
import { ModuleHeaderContext } from "./OpenFGAParser";
import { TypeDefsContext } from "./OpenFGAParser";
import { TypeDefContext } from "./OpenFGAParser";
import { MixinDeclarationContext } from "./OpenFGAParser";
import { RelationDeclarationContext } from "./OpenFGAParser";
import { RelationNameContext } from "./OpenFGAParser";
import { RelationDefContext } from "./OpenFGAParser";
import { RelationDefNoDirectContext } from "./OpenFGAParser";
import { RelationDefPartialsContext } from "./OpenFGAParser";
import { RelationDefGroupingContext } from "./OpenFGAParser";
import { RelationRecurseContext } from "./OpenFGAParser";
import { RelationRecurseNoDirectContext } from "./OpenFGAParser";
import { RelationDefDirectAssignmentContext } from "./OpenFGAParser";
import { RelationDefRewriteContext } from "./OpenFGAParser";
import { RelationDefTypeRestrictionContext } from "./OpenFGAParser";
import { RelationDefTypeRestrictionBaseContext } from "./OpenFGAParser";
import { ConditionsContext } from "./OpenFGAParser";
import { ConditionContext } from "./OpenFGAParser";
import { ConditionNameContext } from "./OpenFGAParser";
import { ConditionParameterContext } from "./OpenFGAParser";
import { ParameterNameContext } from "./OpenFGAParser";
import { ParameterTypeContext } from "./OpenFGAParser";
import { MixinsContext } from "./OpenFGAParser";
import { MixinContext } from "./OpenFGAParser";
import { MixinNameContext } from "./OpenFGAParser";
import { MultiLineCommentContext } from "./OpenFGAParser";
import { IdentifierContext } from "./OpenFGAParser";
import { ConditionExpressionContext } from "./OpenFGAParser";


/**
 * This interface defines a complete listener for a parse tree produced by
 * `OpenFGAParser`.
 */
export default class OpenFGAParserListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by `OpenFGAParser.main`.
	 * @param ctx the parse tree
	 */
	enterMain?: (ctx: MainContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.main`.
	 * @param ctx the parse tree
	 */
	exitMain?: (ctx: MainContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.modelHeader`.
	 * @param ctx the parse tree
	 */
	enterModelHeader?: (ctx: ModelHeaderContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.modelHeader`.
	 * @param ctx the parse tree
	 */
	exitModelHeader?: (ctx: ModelHeaderContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.moduleHeader`.
	 * @param ctx the parse tree
	 */
	enterModuleHeader?: (ctx: ModuleHeaderContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.moduleHeader`.
	 * @param ctx the parse tree
	 */
	exitModuleHeader?: (ctx: ModuleHeaderContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.typeDefs`.
	 * @param ctx the parse tree
	 */
	enterTypeDefs?: (ctx: TypeDefsContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.typeDefs`.
	 * @param ctx the parse tree
	 */
	exitTypeDefs?: (ctx: TypeDefsContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.typeDef`.
	 * @param ctx the parse tree
	 */
	enterTypeDef?: (ctx: TypeDefContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.typeDef`.
	 * @param ctx the parse tree
	 */
	exitTypeDef?: (ctx: TypeDefContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.mixinDeclaration`.
	 * @param ctx the parse tree
	 */
	enterMixinDeclaration?: (ctx: MixinDeclarationContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.mixinDeclaration`.
	 * @param ctx the parse tree
	 */
	exitMixinDeclaration?: (ctx: MixinDeclarationContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationDeclaration`.
	 * @param ctx the parse tree
	 */
	enterRelationDeclaration?: (ctx: RelationDeclarationContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDeclaration`.
	 * @param ctx the parse tree
	 */
	exitRelationDeclaration?: (ctx: RelationDeclarationContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationName`.
	 * @param ctx the parse tree
	 */
	enterRelationName?: (ctx: RelationNameContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationName`.
	 * @param ctx the parse tree
	 */
	exitRelationName?: (ctx: RelationNameContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationDef`.
	 * @param ctx the parse tree
	 */
	enterRelationDef?: (ctx: RelationDefContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDef`.
	 * @param ctx the parse tree
	 */
	exitRelationDef?: (ctx: RelationDefContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationDefNoDirect`.
	 * @param ctx the parse tree
	 */
	enterRelationDefNoDirect?: (ctx: RelationDefNoDirectContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDefNoDirect`.
	 * @param ctx the parse tree
	 */
	exitRelationDefNoDirect?: (ctx: RelationDefNoDirectContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationDefPartials`.
	 * @param ctx the parse tree
	 */
	enterRelationDefPartials?: (ctx: RelationDefPartialsContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDefPartials`.
	 * @param ctx the parse tree
	 */
	exitRelationDefPartials?: (ctx: RelationDefPartialsContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationDefGrouping`.
	 * @param ctx the parse tree
	 */
	enterRelationDefGrouping?: (ctx: RelationDefGroupingContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDefGrouping`.
	 * @param ctx the parse tree
	 */
	exitRelationDefGrouping?: (ctx: RelationDefGroupingContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationRecurse`.
	 * @param ctx the parse tree
	 */
	enterRelationRecurse?: (ctx: RelationRecurseContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationRecurse`.
	 * @param ctx the parse tree
	 */
	exitRelationRecurse?: (ctx: RelationRecurseContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationRecurseNoDirect`.
	 * @param ctx the parse tree
	 */
	enterRelationRecurseNoDirect?: (ctx: RelationRecurseNoDirectContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationRecurseNoDirect`.
	 * @param ctx the parse tree
	 */
	exitRelationRecurseNoDirect?: (ctx: RelationRecurseNoDirectContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationDefDirectAssignment`.
	 * @param ctx the parse tree
	 */
	enterRelationDefDirectAssignment?: (ctx: RelationDefDirectAssignmentContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDefDirectAssignment`.
	 * @param ctx the parse tree
	 */
	exitRelationDefDirectAssignment?: (ctx: RelationDefDirectAssignmentContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationDefRewrite`.
	 * @param ctx the parse tree
	 */
	enterRelationDefRewrite?: (ctx: RelationDefRewriteContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDefRewrite`.
	 * @param ctx the parse tree
	 */
	exitRelationDefRewrite?: (ctx: RelationDefRewriteContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationDefTypeRestriction`.
	 * @param ctx the parse tree
	 */
	enterRelationDefTypeRestriction?: (ctx: RelationDefTypeRestrictionContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDefTypeRestriction`.
	 * @param ctx the parse tree
	 */
	exitRelationDefTypeRestriction?: (ctx: RelationDefTypeRestrictionContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationDefTypeRestrictionBase`.
	 * @param ctx the parse tree
	 */
	enterRelationDefTypeRestrictionBase?: (ctx: RelationDefTypeRestrictionBaseContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDefTypeRestrictionBase`.
	 * @param ctx the parse tree
	 */
	exitRelationDefTypeRestrictionBase?: (ctx: RelationDefTypeRestrictionBaseContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.conditions`.
	 * @param ctx the parse tree
	 */
	enterConditions?: (ctx: ConditionsContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.conditions`.
	 * @param ctx the parse tree
	 */
	exitConditions?: (ctx: ConditionsContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.condition`.
	 * @param ctx the parse tree
	 */
	enterCondition?: (ctx: ConditionContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.condition`.
	 * @param ctx the parse tree
	 */
	exitCondition?: (ctx: ConditionContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.conditionName`.
	 * @param ctx the parse tree
	 */
	enterConditionName?: (ctx: ConditionNameContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.conditionName`.
	 * @param ctx the parse tree
	 */
	exitConditionName?: (ctx: ConditionNameContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.conditionParameter`.
	 * @param ctx the parse tree
	 */
	enterConditionParameter?: (ctx: ConditionParameterContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.conditionParameter`.
	 * @param ctx the parse tree
	 */
	exitConditionParameter?: (ctx: ConditionParameterContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.parameterName`.
	 * @param ctx the parse tree
	 */
	enterParameterName?: (ctx: ParameterNameContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.parameterName`.
	 * @param ctx the parse tree
	 */
	exitParameterName?: (ctx: ParameterNameContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.parameterType`.
	 * @param ctx the parse tree
	 */
	enterParameterType?: (ctx: ParameterTypeContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.parameterType`.
	 * @param ctx the parse tree
	 */
	exitParameterType?: (ctx: ParameterTypeContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.mixins`.
	 * @param ctx the parse tree
	 */
	enterMixins?: (ctx: MixinsContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.mixins`.
	 * @param ctx the parse tree
	 */
	exitMixins?: (ctx: MixinsContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.mixin`.
	 * @param ctx the parse tree
	 */
	enterMixin?: (ctx: MixinContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.mixin`.
	 * @param ctx the parse tree
	 */
	exitMixin?: (ctx: MixinContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.mixinName`.
	 * @param ctx the parse tree
	 */
	enterMixinName?: (ctx: MixinNameContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.mixinName`.
	 * @param ctx the parse tree
	 */
	exitMixinName?: (ctx: MixinNameContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.multiLineComment`.
	 * @param ctx the parse tree
	 */
	enterMultiLineComment?: (ctx: MultiLineCommentContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.multiLineComment`.
	 * @param ctx the parse tree
	 */
	exitMultiLineComment?: (ctx: MultiLineCommentContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.identifier`.
	 * @param ctx the parse tree
	 */
	enterIdentifier?: (ctx: IdentifierContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.identifier`.
	 * @param ctx the parse tree
	 */
	exitIdentifier?: (ctx: IdentifierContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.conditionExpression`.
	 * @param ctx the parse tree
	 */
	enterConditionExpression?: (ctx: ConditionExpressionContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.conditionExpression`.
	 * @param ctx the parse tree
	 */
	exitConditionExpression?: (ctx: ConditionExpressionContext) => void;
}

