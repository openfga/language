// Generated from /app/OpenFGAParser.g4 by ANTLR 4.13.0

import {ParseTreeVisitor} from 'antlr4';


import { MainContext } from "./OpenFGAParser";
import { IndentationContext } from "./OpenFGAParser";
import { ModelHeaderContext } from "./OpenFGAParser";
import { TypeDefsContext } from "./OpenFGAParser";
import { TypeDefContext } from "./OpenFGAParser";
import { RelationDeclarationContext } from "./OpenFGAParser";
import { RelationDefContext } from "./OpenFGAParser";
import { RelationDefPartialsContext } from "./OpenFGAParser";
import { RelationDefPartialAllOrContext } from "./OpenFGAParser";
import { RelationDefPartialAllAndContext } from "./OpenFGAParser";
import { RelationDefPartialAllButNotContext } from "./OpenFGAParser";
import { RelationDefDirectAssignmentContext } from "./OpenFGAParser";
import { RelationDefRewriteContext } from "./OpenFGAParser";
import { RelationDefRelationOnSameObjectContext } from "./OpenFGAParser";
import { RelationDefRelationOnRelatedObjectContext } from "./OpenFGAParser";
import { RelationDefOperatorContext } from "./OpenFGAParser";
import { RelationDefOperatorAndContext } from "./OpenFGAParser";
import { RelationDefOperatorOrContext } from "./OpenFGAParser";
import { RelationDefOperatorButNotContext } from "./OpenFGAParser";
import { RelationDefKeywordFromContext } from "./OpenFGAParser";
import { RelationDefTypeRestrictionContext } from "./OpenFGAParser";
import { RelationDefTypeRestrictionTypeContext } from "./OpenFGAParser";
import { RelationDefTypeRestrictionRelationContext } from "./OpenFGAParser";
import { RelationDefTypeRestrictionWildcardContext } from "./OpenFGAParser";
import { RelationDefTypeRestrictionUsersetContext } from "./OpenFGAParser";
import { RelationDefGroupingContext } from "./OpenFGAParser";
import { RewriteComputedusersetNameContext } from "./OpenFGAParser";
import { RewriteTuplesetComputedusersetNameContext } from "./OpenFGAParser";
import { RewriteTuplesetNameContext } from "./OpenFGAParser";
import { RelationNameContext } from "./OpenFGAParser";
import { TypeNameContext } from "./OpenFGAParser";
import { CommentContext } from "./OpenFGAParser";
import { MultiLineCommentContext } from "./OpenFGAParser";
import { SpacingContext } from "./OpenFGAParser";
import { NewlineContext } from "./OpenFGAParser";
import { SchemaVersionContext } from "./OpenFGAParser";
import { NameContext } from "./OpenFGAParser";


/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by `OpenFGAParser`.
 *
 * @param <Result> The return type of the visit operation. Use `void` for
 * operations with no return type.
 */
export default class OpenFGAParserVisitor<Result> extends ParseTreeVisitor<Result> {
	/**
	 * Visit a parse tree produced by `OpenFGAParser.main`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitMain?: (ctx: MainContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.indentation`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitIndentation?: (ctx: IndentationContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.modelHeader`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitModelHeader?: (ctx: ModelHeaderContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.typeDefs`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitTypeDefs?: (ctx: TypeDefsContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.typeDef`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitTypeDef?: (ctx: TypeDefContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDeclaration`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDeclaration?: (ctx: RelationDeclarationContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDef`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDef?: (ctx: RelationDefContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefPartials`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefPartials?: (ctx: RelationDefPartialsContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefPartialAllOr`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefPartialAllOr?: (ctx: RelationDefPartialAllOrContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefPartialAllAnd`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefPartialAllAnd?: (ctx: RelationDefPartialAllAndContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefPartialAllButNot`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefPartialAllButNot?: (ctx: RelationDefPartialAllButNotContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefDirectAssignment`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefDirectAssignment?: (ctx: RelationDefDirectAssignmentContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefRewrite`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefRewrite?: (ctx: RelationDefRewriteContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefRelationOnSameObject`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefRelationOnSameObject?: (ctx: RelationDefRelationOnSameObjectContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefRelationOnRelatedObject`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefRelationOnRelatedObject?: (ctx: RelationDefRelationOnRelatedObjectContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefOperator`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefOperator?: (ctx: RelationDefOperatorContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefOperatorAnd`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefOperatorAnd?: (ctx: RelationDefOperatorAndContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefOperatorOr`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefOperatorOr?: (ctx: RelationDefOperatorOrContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefOperatorButNot`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefOperatorButNot?: (ctx: RelationDefOperatorButNotContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefKeywordFrom`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefKeywordFrom?: (ctx: RelationDefKeywordFromContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefTypeRestriction`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefTypeRestriction?: (ctx: RelationDefTypeRestrictionContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefTypeRestrictionType`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefTypeRestrictionType?: (ctx: RelationDefTypeRestrictionTypeContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefTypeRestrictionRelation`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefTypeRestrictionRelation?: (ctx: RelationDefTypeRestrictionRelationContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefTypeRestrictionWildcard`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefTypeRestrictionWildcard?: (ctx: RelationDefTypeRestrictionWildcardContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefTypeRestrictionUserset`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefTypeRestrictionUserset?: (ctx: RelationDefTypeRestrictionUsersetContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationDefGrouping`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationDefGrouping?: (ctx: RelationDefGroupingContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.rewriteComputedusersetName`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRewriteComputedusersetName?: (ctx: RewriteComputedusersetNameContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.rewriteTuplesetComputedusersetName`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRewriteTuplesetComputedusersetName?: (ctx: RewriteTuplesetComputedusersetNameContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.rewriteTuplesetName`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRewriteTuplesetName?: (ctx: RewriteTuplesetNameContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.relationName`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitRelationName?: (ctx: RelationNameContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.typeName`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitTypeName?: (ctx: TypeNameContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.comment`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitComment?: (ctx: CommentContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.multiLineComment`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitMultiLineComment?: (ctx: MultiLineCommentContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.spacing`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitSpacing?: (ctx: SpacingContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.newline`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitNewline?: (ctx: NewlineContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.schemaVersion`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitSchemaVersion?: (ctx: SchemaVersionContext) => Result;
	/**
	 * Visit a parse tree produced by `OpenFGAParser.name`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitName?: (ctx: NameContext) => Result;
}

