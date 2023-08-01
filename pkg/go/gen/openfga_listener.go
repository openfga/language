// Code generated from /app/OpenFGA.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // OpenFGA

import "github.com/antlr4-go/antlr/v4"

// OpenFGAListener is a complete listener for a parse tree produced by OpenFGAParser.
type OpenFGAListener interface {
	antlr.ParseTreeListener

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// EnterIndentation is called when entering the indentation production.
	EnterIndentation(c *IndentationContext)

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

	// EnterRelationDefOperator is called when entering the relationDefOperator production.
	EnterRelationDefOperator(c *RelationDefOperatorContext)

	// EnterRelationDefOperatorAnd is called when entering the relationDefOperatorAnd production.
	EnterRelationDefOperatorAnd(c *RelationDefOperatorAndContext)

	// EnterRelationDefOperatorOr is called when entering the relationDefOperatorOr production.
	EnterRelationDefOperatorOr(c *RelationDefOperatorOrContext)

	// EnterRelationDefOperatorButNot is called when entering the relationDefOperatorButNot production.
	EnterRelationDefOperatorButNot(c *RelationDefOperatorButNotContext)

	// EnterRelationDefKeywordFrom is called when entering the relationDefKeywordFrom production.
	EnterRelationDefKeywordFrom(c *RelationDefKeywordFromContext)

	// EnterRelationDefTypeRestriction is called when entering the relationDefTypeRestriction production.
	EnterRelationDefTypeRestriction(c *RelationDefTypeRestrictionContext)

	// EnterRelationDefTypeRestrictionType is called when entering the relationDefTypeRestrictionType production.
	EnterRelationDefTypeRestrictionType(c *RelationDefTypeRestrictionTypeContext)

	// EnterRelationDefTypeRestrictionRelation is called when entering the relationDefTypeRestrictionRelation production.
	EnterRelationDefTypeRestrictionRelation(c *RelationDefTypeRestrictionRelationContext)

	// EnterRelationDefTypeRestrictionWildcard is called when entering the relationDefTypeRestrictionWildcard production.
	EnterRelationDefTypeRestrictionWildcard(c *RelationDefTypeRestrictionWildcardContext)

	// EnterRelationDefTypeRestrictionUserset is called when entering the relationDefTypeRestrictionUserset production.
	EnterRelationDefTypeRestrictionUserset(c *RelationDefTypeRestrictionUsersetContext)

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

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterMultiLineComment is called when entering the multiLineComment production.
	EnterMultiLineComment(c *MultiLineCommentContext)

	// EnterSpacing is called when entering the spacing production.
	EnterSpacing(c *SpacingContext)

	// EnterNewline is called when entering the newline production.
	EnterNewline(c *NewlineContext)

	// EnterSchemaVersion is called when entering the schemaVersion production.
	EnterSchemaVersion(c *SchemaVersionContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitIndentation is called when exiting the indentation production.
	ExitIndentation(c *IndentationContext)

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

	// ExitRelationDefOperator is called when exiting the relationDefOperator production.
	ExitRelationDefOperator(c *RelationDefOperatorContext)

	// ExitRelationDefOperatorAnd is called when exiting the relationDefOperatorAnd production.
	ExitRelationDefOperatorAnd(c *RelationDefOperatorAndContext)

	// ExitRelationDefOperatorOr is called when exiting the relationDefOperatorOr production.
	ExitRelationDefOperatorOr(c *RelationDefOperatorOrContext)

	// ExitRelationDefOperatorButNot is called when exiting the relationDefOperatorButNot production.
	ExitRelationDefOperatorButNot(c *RelationDefOperatorButNotContext)

	// ExitRelationDefKeywordFrom is called when exiting the relationDefKeywordFrom production.
	ExitRelationDefKeywordFrom(c *RelationDefKeywordFromContext)

	// ExitRelationDefTypeRestriction is called when exiting the relationDefTypeRestriction production.
	ExitRelationDefTypeRestriction(c *RelationDefTypeRestrictionContext)

	// ExitRelationDefTypeRestrictionType is called when exiting the relationDefTypeRestrictionType production.
	ExitRelationDefTypeRestrictionType(c *RelationDefTypeRestrictionTypeContext)

	// ExitRelationDefTypeRestrictionRelation is called when exiting the relationDefTypeRestrictionRelation production.
	ExitRelationDefTypeRestrictionRelation(c *RelationDefTypeRestrictionRelationContext)

	// ExitRelationDefTypeRestrictionWildcard is called when exiting the relationDefTypeRestrictionWildcard production.
	ExitRelationDefTypeRestrictionWildcard(c *RelationDefTypeRestrictionWildcardContext)

	// ExitRelationDefTypeRestrictionUserset is called when exiting the relationDefTypeRestrictionUserset production.
	ExitRelationDefTypeRestrictionUserset(c *RelationDefTypeRestrictionUsersetContext)

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

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitMultiLineComment is called when exiting the multiLineComment production.
	ExitMultiLineComment(c *MultiLineCommentContext)

	// ExitSpacing is called when exiting the spacing production.
	ExitSpacing(c *SpacingContext)

	// ExitNewline is called when exiting the newline production.
	ExitNewline(c *NewlineContext)

	// ExitSchemaVersion is called when exiting the schemaVersion production.
	ExitSchemaVersion(c *SchemaVersionContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)
}
