// Code generated from /app/OpenFGA.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // OpenFGA

import "github.com/antlr4-go/antlr/v4"


// OpenFGAListener is a complete listener for a parse tree produced by OpenFGAParser.
type OpenFGAListener interface {
	antlr.ParseTreeListener

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// EnterModelHeaderComment is called when entering the modelHeaderComment production.
	EnterModelHeaderComment(c *ModelHeaderCommentContext)

	// EnterModelHeader is called when entering the modelHeader production.
	EnterModelHeader(c *ModelHeaderContext)

	// EnterTypeDefs is called when entering the typeDefs production.
	EnterTypeDefs(c *TypeDefsContext)

	// EnterTypeDef is called when entering the typeDef production.
	EnterTypeDef(c *TypeDefContext)

	// EnterRelationDeclaration is called when entering the relationDeclaration production.
	EnterRelationDeclaration(c *RelationDeclarationContext)

	// EnterRelationDef is called when entering the relationDef production.
	EnterRelationDef(c *RelationDefContext)

	// EnterRelationDefPartials is called when entering the relationDefPartials production.
	EnterRelationDefPartials(c *RelationDefPartialsContext)

	// EnterRelationDefPartialAllOr is called when entering the relationDefPartialAllOr production.
	EnterRelationDefPartialAllOr(c *RelationDefPartialAllOrContext)

	// EnterRelationDefPartialAllAnd is called when entering the relationDefPartialAllAnd production.
	EnterRelationDefPartialAllAnd(c *RelationDefPartialAllAndContext)

	// EnterRelationDefPartialAllButNot is called when entering the relationDefPartialAllButNot production.
	EnterRelationDefPartialAllButNot(c *RelationDefPartialAllButNotContext)

	// EnterRelationDefDirectAssignment is called when entering the relationDefDirectAssignment production.
	EnterRelationDefDirectAssignment(c *RelationDefDirectAssignmentContext)

	// EnterRelationDefRewrite is called when entering the relationDefRewrite production.
	EnterRelationDefRewrite(c *RelationDefRewriteContext)

	// EnterRelationDefRelationOnSameObject is called when entering the relationDefRelationOnSameObject production.
	EnterRelationDefRelationOnSameObject(c *RelationDefRelationOnSameObjectContext)

	// EnterRelationDefRelationOnRelatedObject is called when entering the relationDefRelationOnRelatedObject production.
	EnterRelationDefRelationOnRelatedObject(c *RelationDefRelationOnRelatedObjectContext)

	// EnterRelationDefTypeRestriction is called when entering the relationDefTypeRestriction production.
	EnterRelationDefTypeRestriction(c *RelationDefTypeRestrictionContext)

	// EnterRelationDefTypeRestrictionType is called when entering the relationDefTypeRestrictionType production.
	EnterRelationDefTypeRestrictionType(c *RelationDefTypeRestrictionTypeContext)

	// EnterRelationDefTypeRestrictionRelation is called when entering the relationDefTypeRestrictionRelation production.
	EnterRelationDefTypeRestrictionRelation(c *RelationDefTypeRestrictionRelationContext)

	// EnterRelationDefGrouping is called when entering the relationDefGrouping production.
	EnterRelationDefGrouping(c *RelationDefGroupingContext)

	// EnterRewriteComputedusersetName is called when entering the rewriteComputedusersetName production.
	EnterRewriteComputedusersetName(c *RewriteComputedusersetNameContext)

	// EnterRewriteTuplesetComputedusersetName is called when entering the rewriteTuplesetComputedusersetName production.
	EnterRewriteTuplesetComputedusersetName(c *RewriteTuplesetComputedusersetNameContext)

	// EnterRewriteTuplesetName is called when entering the rewriteTuplesetName production.
	EnterRewriteTuplesetName(c *RewriteTuplesetNameContext)

	// EnterRelationName is called when entering the relationName production.
	EnterRelationName(c *RelationNameContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterSchemaVersion is called when entering the schemaVersion production.
	EnterSchemaVersion(c *SchemaVersionContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitModelHeaderComment is called when exiting the modelHeaderComment production.
	ExitModelHeaderComment(c *ModelHeaderCommentContext)

	// ExitModelHeader is called when exiting the modelHeader production.
	ExitModelHeader(c *ModelHeaderContext)

	// ExitTypeDefs is called when exiting the typeDefs production.
	ExitTypeDefs(c *TypeDefsContext)

	// ExitTypeDef is called when exiting the typeDef production.
	ExitTypeDef(c *TypeDefContext)

	// ExitRelationDeclaration is called when exiting the relationDeclaration production.
	ExitRelationDeclaration(c *RelationDeclarationContext)

	// ExitRelationDef is called when exiting the relationDef production.
	ExitRelationDef(c *RelationDefContext)

	// ExitRelationDefPartials is called when exiting the relationDefPartials production.
	ExitRelationDefPartials(c *RelationDefPartialsContext)

	// ExitRelationDefPartialAllOr is called when exiting the relationDefPartialAllOr production.
	ExitRelationDefPartialAllOr(c *RelationDefPartialAllOrContext)

	// ExitRelationDefPartialAllAnd is called when exiting the relationDefPartialAllAnd production.
	ExitRelationDefPartialAllAnd(c *RelationDefPartialAllAndContext)

	// ExitRelationDefPartialAllButNot is called when exiting the relationDefPartialAllButNot production.
	ExitRelationDefPartialAllButNot(c *RelationDefPartialAllButNotContext)

	// ExitRelationDefDirectAssignment is called when exiting the relationDefDirectAssignment production.
	ExitRelationDefDirectAssignment(c *RelationDefDirectAssignmentContext)

	// ExitRelationDefRewrite is called when exiting the relationDefRewrite production.
	ExitRelationDefRewrite(c *RelationDefRewriteContext)

	// ExitRelationDefRelationOnSameObject is called when exiting the relationDefRelationOnSameObject production.
	ExitRelationDefRelationOnSameObject(c *RelationDefRelationOnSameObjectContext)

	// ExitRelationDefRelationOnRelatedObject is called when exiting the relationDefRelationOnRelatedObject production.
	ExitRelationDefRelationOnRelatedObject(c *RelationDefRelationOnRelatedObjectContext)

	// ExitRelationDefTypeRestriction is called when exiting the relationDefTypeRestriction production.
	ExitRelationDefTypeRestriction(c *RelationDefTypeRestrictionContext)

	// ExitRelationDefTypeRestrictionType is called when exiting the relationDefTypeRestrictionType production.
	ExitRelationDefTypeRestrictionType(c *RelationDefTypeRestrictionTypeContext)

	// ExitRelationDefTypeRestrictionRelation is called when exiting the relationDefTypeRestrictionRelation production.
	ExitRelationDefTypeRestrictionRelation(c *RelationDefTypeRestrictionRelationContext)

	// ExitRelationDefGrouping is called when exiting the relationDefGrouping production.
	ExitRelationDefGrouping(c *RelationDefGroupingContext)

	// ExitRewriteComputedusersetName is called when exiting the rewriteComputedusersetName production.
	ExitRewriteComputedusersetName(c *RewriteComputedusersetNameContext)

	// ExitRewriteTuplesetComputedusersetName is called when exiting the rewriteTuplesetComputedusersetName production.
	ExitRewriteTuplesetComputedusersetName(c *RewriteTuplesetComputedusersetNameContext)

	// ExitRewriteTuplesetName is called when exiting the rewriteTuplesetName production.
	ExitRewriteTuplesetName(c *RewriteTuplesetNameContext)

	// ExitRelationName is called when exiting the relationName production.
	ExitRelationName(c *RelationNameContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitSchemaVersion is called when exiting the schemaVersion production.
	ExitSchemaVersion(c *SchemaVersionContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)
}
