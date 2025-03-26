// Generated from /app/OpenFGAParser.g4 by ANTLR 4.13.1
package dev.openfga.language.antlr;
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link OpenFGAParser}.
 */
public interface OpenFGAParserListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#main}.
	 * @param ctx the parse tree
	 */
	void enterMain(OpenFGAParser.MainContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#main}.
	 * @param ctx the parse tree
	 */
	void exitMain(OpenFGAParser.MainContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#modelHeader}.
	 * @param ctx the parse tree
	 */
	void enterModelHeader(OpenFGAParser.ModelHeaderContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#modelHeader}.
	 * @param ctx the parse tree
	 */
	void exitModelHeader(OpenFGAParser.ModelHeaderContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#moduleHeader}.
	 * @param ctx the parse tree
	 */
	void enterModuleHeader(OpenFGAParser.ModuleHeaderContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#moduleHeader}.
	 * @param ctx the parse tree
	 */
	void exitModuleHeader(OpenFGAParser.ModuleHeaderContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#typeDefs}.
	 * @param ctx the parse tree
	 */
	void enterTypeDefs(OpenFGAParser.TypeDefsContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#typeDefs}.
	 * @param ctx the parse tree
	 */
	void exitTypeDefs(OpenFGAParser.TypeDefsContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#typeDef}.
	 * @param ctx the parse tree
	 */
	void enterTypeDef(OpenFGAParser.TypeDefContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#typeDef}.
	 * @param ctx the parse tree
	 */
	void exitTypeDef(OpenFGAParser.TypeDefContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#relationDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterRelationDeclaration(OpenFGAParser.RelationDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#relationDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitRelationDeclaration(OpenFGAParser.RelationDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#relationName}.
	 * @param ctx the parse tree
	 */
	void enterRelationName(OpenFGAParser.RelationNameContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#relationName}.
	 * @param ctx the parse tree
	 */
	void exitRelationName(OpenFGAParser.RelationNameContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#relationDef}.
	 * @param ctx the parse tree
	 */
	void enterRelationDef(OpenFGAParser.RelationDefContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#relationDef}.
	 * @param ctx the parse tree
	 */
	void exitRelationDef(OpenFGAParser.RelationDefContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#relationDefNoDirect}.
	 * @param ctx the parse tree
	 */
	void enterRelationDefNoDirect(OpenFGAParser.RelationDefNoDirectContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#relationDefNoDirect}.
	 * @param ctx the parse tree
	 */
	void exitRelationDefNoDirect(OpenFGAParser.RelationDefNoDirectContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#relationDefPartials}.
	 * @param ctx the parse tree
	 */
	void enterRelationDefPartials(OpenFGAParser.RelationDefPartialsContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#relationDefPartials}.
	 * @param ctx the parse tree
	 */
	void exitRelationDefPartials(OpenFGAParser.RelationDefPartialsContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#relationDefGrouping}.
	 * @param ctx the parse tree
	 */
	void enterRelationDefGrouping(OpenFGAParser.RelationDefGroupingContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#relationDefGrouping}.
	 * @param ctx the parse tree
	 */
	void exitRelationDefGrouping(OpenFGAParser.RelationDefGroupingContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#relationRecurse}.
	 * @param ctx the parse tree
	 */
	void enterRelationRecurse(OpenFGAParser.RelationRecurseContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#relationRecurse}.
	 * @param ctx the parse tree
	 */
	void exitRelationRecurse(OpenFGAParser.RelationRecurseContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#relationRecurseNoDirect}.
	 * @param ctx the parse tree
	 */
	void enterRelationRecurseNoDirect(OpenFGAParser.RelationRecurseNoDirectContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#relationRecurseNoDirect}.
	 * @param ctx the parse tree
	 */
	void exitRelationRecurseNoDirect(OpenFGAParser.RelationRecurseNoDirectContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#relationDefDirectAssignment}.
	 * @param ctx the parse tree
	 */
	void enterRelationDefDirectAssignment(OpenFGAParser.RelationDefDirectAssignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#relationDefDirectAssignment}.
	 * @param ctx the parse tree
	 */
	void exitRelationDefDirectAssignment(OpenFGAParser.RelationDefDirectAssignmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#relationDefRewrite}.
	 * @param ctx the parse tree
	 */
	void enterRelationDefRewrite(OpenFGAParser.RelationDefRewriteContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#relationDefRewrite}.
	 * @param ctx the parse tree
	 */
	void exitRelationDefRewrite(OpenFGAParser.RelationDefRewriteContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#relationDefTypeRestriction}.
	 * @param ctx the parse tree
	 */
	void enterRelationDefTypeRestriction(OpenFGAParser.RelationDefTypeRestrictionContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#relationDefTypeRestriction}.
	 * @param ctx the parse tree
	 */
	void exitRelationDefTypeRestriction(OpenFGAParser.RelationDefTypeRestrictionContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#relationDefTypeRestrictionBase}.
	 * @param ctx the parse tree
	 */
	void enterRelationDefTypeRestrictionBase(OpenFGAParser.RelationDefTypeRestrictionBaseContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#relationDefTypeRestrictionBase}.
	 * @param ctx the parse tree
	 */
	void exitRelationDefTypeRestrictionBase(OpenFGAParser.RelationDefTypeRestrictionBaseContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#conditions}.
	 * @param ctx the parse tree
	 */
	void enterConditions(OpenFGAParser.ConditionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#conditions}.
	 * @param ctx the parse tree
	 */
	void exitConditions(OpenFGAParser.ConditionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#condition}.
	 * @param ctx the parse tree
	 */
	void enterCondition(OpenFGAParser.ConditionContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#condition}.
	 * @param ctx the parse tree
	 */
	void exitCondition(OpenFGAParser.ConditionContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#conditionName}.
	 * @param ctx the parse tree
	 */
	void enterConditionName(OpenFGAParser.ConditionNameContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#conditionName}.
	 * @param ctx the parse tree
	 */
	void exitConditionName(OpenFGAParser.ConditionNameContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#conditionParameter}.
	 * @param ctx the parse tree
	 */
	void enterConditionParameter(OpenFGAParser.ConditionParameterContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#conditionParameter}.
	 * @param ctx the parse tree
	 */
	void exitConditionParameter(OpenFGAParser.ConditionParameterContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#parameterName}.
	 * @param ctx the parse tree
	 */
	void enterParameterName(OpenFGAParser.ParameterNameContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#parameterName}.
	 * @param ctx the parse tree
	 */
	void exitParameterName(OpenFGAParser.ParameterNameContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#parameterType}.
	 * @param ctx the parse tree
	 */
	void enterParameterType(OpenFGAParser.ParameterTypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#parameterType}.
	 * @param ctx the parse tree
	 */
	void exitParameterType(OpenFGAParser.ParameterTypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#multiLineComment}.
	 * @param ctx the parse tree
	 */
	void enterMultiLineComment(OpenFGAParser.MultiLineCommentContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#multiLineComment}.
	 * @param ctx the parse tree
	 */
	void exitMultiLineComment(OpenFGAParser.MultiLineCommentContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#identifier}.
	 * @param ctx the parse tree
	 */
	void enterIdentifier(OpenFGAParser.IdentifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#identifier}.
	 * @param ctx the parse tree
	 */
	void exitIdentifier(OpenFGAParser.IdentifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#extended_identifier}.
	 * @param ctx the parse tree
	 */
	void enterExtended_identifier(OpenFGAParser.Extended_identifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#extended_identifier}.
	 * @param ctx the parse tree
	 */
	void exitExtended_identifier(OpenFGAParser.Extended_identifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link OpenFGAParser#conditionExpression}.
	 * @param ctx the parse tree
	 */
	void enterConditionExpression(OpenFGAParser.ConditionExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link OpenFGAParser#conditionExpression}.
	 * @param ctx the parse tree
	 */
	void exitConditionExpression(OpenFGAParser.ConditionExpressionContext ctx);
}