// Generated from /app/OpenFGA.g4 by ANTLR 4.13.0

import {ParseTreeListener} from "antlr4";


import { MainContext } from "./OpenFGAParser";
import { ModelHeaderCommentContext } from "./OpenFGAParser";
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
import { RelationDefTypeRestrictionContext } from "./OpenFGAParser";
import { RelationDefTypeRestrictionTypeContext } from "./OpenFGAParser";
import { RelationDefTypeRestrictionRelationContext } from "./OpenFGAParser";
import { RelationDefGroupingContext } from "./OpenFGAParser";
import { RewriteComputedusersetNameContext } from "./OpenFGAParser";
import { RewriteTuplesetComputedusersetNameContext } from "./OpenFGAParser";
import { RewriteTuplesetNameContext } from "./OpenFGAParser";
import { RelationNameContext } from "./OpenFGAParser";
import { TypeNameContext } from "./OpenFGAParser";
import { SchemaVersionContext } from "./OpenFGAParser";
import { NameContext } from "./OpenFGAParser";


/**
 * This interface defines a complete listener for a parse tree produced by
 * `OpenFGAParser`.
 */
export default class OpenFGAListener extends ParseTreeListener {
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
	 * Enter a parse tree produced by `OpenFGAParser.modelHeaderComment`.
	 * @param ctx the parse tree
	 */
	enterModelHeaderComment?: (ctx: ModelHeaderCommentContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.modelHeaderComment`.
	 * @param ctx the parse tree
	 */
	exitModelHeaderComment?: (ctx: ModelHeaderCommentContext) => void;
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
	 * Enter a parse tree produced by `OpenFGAParser.relationDefPartialAllOr`.
	 * @param ctx the parse tree
	 */
	enterRelationDefPartialAllOr?: (ctx: RelationDefPartialAllOrContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDefPartialAllOr`.
	 * @param ctx the parse tree
	 */
	exitRelationDefPartialAllOr?: (ctx: RelationDefPartialAllOrContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationDefPartialAllAnd`.
	 * @param ctx the parse tree
	 */
	enterRelationDefPartialAllAnd?: (ctx: RelationDefPartialAllAndContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDefPartialAllAnd`.
	 * @param ctx the parse tree
	 */
	exitRelationDefPartialAllAnd?: (ctx: RelationDefPartialAllAndContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationDefPartialAllButNot`.
	 * @param ctx the parse tree
	 */
	enterRelationDefPartialAllButNot?: (ctx: RelationDefPartialAllButNotContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDefPartialAllButNot`.
	 * @param ctx the parse tree
	 */
	exitRelationDefPartialAllButNot?: (ctx: RelationDefPartialAllButNotContext) => void;
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
	 * Enter a parse tree produced by `OpenFGAParser.relationDefRelationOnSameObject`.
	 * @param ctx the parse tree
	 */
	enterRelationDefRelationOnSameObject?: (ctx: RelationDefRelationOnSameObjectContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDefRelationOnSameObject`.
	 * @param ctx the parse tree
	 */
	exitRelationDefRelationOnSameObject?: (ctx: RelationDefRelationOnSameObjectContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationDefRelationOnRelatedObject`.
	 * @param ctx the parse tree
	 */
	enterRelationDefRelationOnRelatedObject?: (ctx: RelationDefRelationOnRelatedObjectContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDefRelationOnRelatedObject`.
	 * @param ctx the parse tree
	 */
	exitRelationDefRelationOnRelatedObject?: (ctx: RelationDefRelationOnRelatedObjectContext) => void;
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
	 * Enter a parse tree produced by `OpenFGAParser.relationDefTypeRestrictionType`.
	 * @param ctx the parse tree
	 */
	enterRelationDefTypeRestrictionType?: (ctx: RelationDefTypeRestrictionTypeContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDefTypeRestrictionType`.
	 * @param ctx the parse tree
	 */
	exitRelationDefTypeRestrictionType?: (ctx: RelationDefTypeRestrictionTypeContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.relationDefTypeRestrictionRelation`.
	 * @param ctx the parse tree
	 */
	enterRelationDefTypeRestrictionRelation?: (ctx: RelationDefTypeRestrictionRelationContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.relationDefTypeRestrictionRelation`.
	 * @param ctx the parse tree
	 */
	exitRelationDefTypeRestrictionRelation?: (ctx: RelationDefTypeRestrictionRelationContext) => void;
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
	 * Enter a parse tree produced by `OpenFGAParser.rewriteComputedusersetName`.
	 * @param ctx the parse tree
	 */
	enterRewriteComputedusersetName?: (ctx: RewriteComputedusersetNameContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.rewriteComputedusersetName`.
	 * @param ctx the parse tree
	 */
	exitRewriteComputedusersetName?: (ctx: RewriteComputedusersetNameContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.rewriteTuplesetComputedusersetName`.
	 * @param ctx the parse tree
	 */
	enterRewriteTuplesetComputedusersetName?: (ctx: RewriteTuplesetComputedusersetNameContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.rewriteTuplesetComputedusersetName`.
	 * @param ctx the parse tree
	 */
	exitRewriteTuplesetComputedusersetName?: (ctx: RewriteTuplesetComputedusersetNameContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.rewriteTuplesetName`.
	 * @param ctx the parse tree
	 */
	enterRewriteTuplesetName?: (ctx: RewriteTuplesetNameContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.rewriteTuplesetName`.
	 * @param ctx the parse tree
	 */
	exitRewriteTuplesetName?: (ctx: RewriteTuplesetNameContext) => void;
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
	 * Enter a parse tree produced by `OpenFGAParser.typeName`.
	 * @param ctx the parse tree
	 */
	enterTypeName?: (ctx: TypeNameContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.typeName`.
	 * @param ctx the parse tree
	 */
	exitTypeName?: (ctx: TypeNameContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.schemaVersion`.
	 * @param ctx the parse tree
	 */
	enterSchemaVersion?: (ctx: SchemaVersionContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.schemaVersion`.
	 * @param ctx the parse tree
	 */
	exitSchemaVersion?: (ctx: SchemaVersionContext) => void;
	/**
	 * Enter a parse tree produced by `OpenFGAParser.name`.
	 * @param ctx the parse tree
	 */
	enterName?: (ctx: NameContext) => void;
	/**
	 * Exit a parse tree produced by `OpenFGAParser.name`.
	 * @param ctx the parse tree
	 */
	exitName?: (ctx: NameContext) => void;
}

