// Code generated from /app/OpenFGA.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // OpenFGA

import "github.com/antlr4-go/antlr/v4"

// BaseOpenFGAListener is a complete listener for a parse tree produced by OpenFGAParser.
type BaseOpenFGAListener struct{}

var _ OpenFGAListener = &BaseOpenFGAListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOpenFGAListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOpenFGAListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOpenFGAListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOpenFGAListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMain is called when production main is entered.
func (s *BaseOpenFGAListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BaseOpenFGAListener) ExitMain(ctx *MainContext) {}

// EnterIndentation is called when production indentation is entered.
func (s *BaseOpenFGAListener) EnterIndentation(ctx *IndentationContext) {}

// ExitIndentation is called when production indentation is exited.
func (s *BaseOpenFGAListener) ExitIndentation(ctx *IndentationContext) {}

// EnterModelHeader is called when production modelHeader is entered.
func (s *BaseOpenFGAListener) EnterModelHeader(ctx *ModelHeaderContext) {}

// ExitModelHeader is called when production modelHeader is exited.
func (s *BaseOpenFGAListener) ExitModelHeader(ctx *ModelHeaderContext) {}

// EnterTypeDefs is called when production typeDefs is entered.
func (s *BaseOpenFGAListener) EnterTypeDefs(ctx *TypeDefsContext) {}

// ExitTypeDefs is called when production typeDefs is exited.
func (s *BaseOpenFGAListener) ExitTypeDefs(ctx *TypeDefsContext) {}

// EnterTypeDef is called when production typeDef is entered.
func (s *BaseOpenFGAListener) EnterTypeDef(ctx *TypeDefContext) {}

// ExitTypeDef is called when production typeDef is exited.
func (s *BaseOpenFGAListener) ExitTypeDef(ctx *TypeDefContext) {}

// EnterRelationDeclaration is called when production relationDeclaration is entered.
func (s *BaseOpenFGAListener) EnterRelationDeclaration(ctx *RelationDeclarationContext) {}

// ExitRelationDeclaration is called when production relationDeclaration is exited.
func (s *BaseOpenFGAListener) ExitRelationDeclaration(ctx *RelationDeclarationContext) {}

// EnterRelationDef is called when production relationDef is entered.
func (s *BaseOpenFGAListener) EnterRelationDef(ctx *RelationDefContext) {}

// ExitRelationDef is called when production relationDef is exited.
func (s *BaseOpenFGAListener) ExitRelationDef(ctx *RelationDefContext) {}

// EnterRelationDefPartials is called when production relationDefPartials is entered.
func (s *BaseOpenFGAListener) EnterRelationDefPartials(ctx *RelationDefPartialsContext) {}

// ExitRelationDefPartials is called when production relationDefPartials is exited.
func (s *BaseOpenFGAListener) ExitRelationDefPartials(ctx *RelationDefPartialsContext) {}

// EnterRelationDefPartialAllOr is called when production relationDefPartialAllOr is entered.
func (s *BaseOpenFGAListener) EnterRelationDefPartialAllOr(ctx *RelationDefPartialAllOrContext) {}

// ExitRelationDefPartialAllOr is called when production relationDefPartialAllOr is exited.
func (s *BaseOpenFGAListener) ExitRelationDefPartialAllOr(ctx *RelationDefPartialAllOrContext) {}

// EnterRelationDefPartialAllAnd is called when production relationDefPartialAllAnd is entered.
func (s *BaseOpenFGAListener) EnterRelationDefPartialAllAnd(ctx *RelationDefPartialAllAndContext) {}

// ExitRelationDefPartialAllAnd is called when production relationDefPartialAllAnd is exited.
func (s *BaseOpenFGAListener) ExitRelationDefPartialAllAnd(ctx *RelationDefPartialAllAndContext) {}

// EnterRelationDefPartialAllButNot is called when production relationDefPartialAllButNot is entered.
func (s *BaseOpenFGAListener) EnterRelationDefPartialAllButNot(ctx *RelationDefPartialAllButNotContext) {
}

// ExitRelationDefPartialAllButNot is called when production relationDefPartialAllButNot is exited.
func (s *BaseOpenFGAListener) ExitRelationDefPartialAllButNot(ctx *RelationDefPartialAllButNotContext) {
}

// EnterRelationDefDirectAssignment is called when production relationDefDirectAssignment is entered.
func (s *BaseOpenFGAListener) EnterRelationDefDirectAssignment(ctx *RelationDefDirectAssignmentContext) {
}

// ExitRelationDefDirectAssignment is called when production relationDefDirectAssignment is exited.
func (s *BaseOpenFGAListener) ExitRelationDefDirectAssignment(ctx *RelationDefDirectAssignmentContext) {
}

// EnterRelationDefRewrite is called when production relationDefRewrite is entered.
func (s *BaseOpenFGAListener) EnterRelationDefRewrite(ctx *RelationDefRewriteContext) {}

// ExitRelationDefRewrite is called when production relationDefRewrite is exited.
func (s *BaseOpenFGAListener) ExitRelationDefRewrite(ctx *RelationDefRewriteContext) {}

// EnterRelationDefRelationOnSameObject is called when production relationDefRelationOnSameObject is entered.
func (s *BaseOpenFGAListener) EnterRelationDefRelationOnSameObject(ctx *RelationDefRelationOnSameObjectContext) {
}

// ExitRelationDefRelationOnSameObject is called when production relationDefRelationOnSameObject is exited.
func (s *BaseOpenFGAListener) ExitRelationDefRelationOnSameObject(ctx *RelationDefRelationOnSameObjectContext) {
}

// EnterRelationDefRelationOnRelatedObject is called when production relationDefRelationOnRelatedObject is entered.
func (s *BaseOpenFGAListener) EnterRelationDefRelationOnRelatedObject(ctx *RelationDefRelationOnRelatedObjectContext) {
}

// ExitRelationDefRelationOnRelatedObject is called when production relationDefRelationOnRelatedObject is exited.
func (s *BaseOpenFGAListener) ExitRelationDefRelationOnRelatedObject(ctx *RelationDefRelationOnRelatedObjectContext) {
}

// EnterRelationDefOperator is called when production relationDefOperator is entered.
func (s *BaseOpenFGAListener) EnterRelationDefOperator(ctx *RelationDefOperatorContext) {}

// ExitRelationDefOperator is called when production relationDefOperator is exited.
func (s *BaseOpenFGAListener) ExitRelationDefOperator(ctx *RelationDefOperatorContext) {}

// EnterRelationDefOperatorAnd is called when production relationDefOperatorAnd is entered.
func (s *BaseOpenFGAListener) EnterRelationDefOperatorAnd(ctx *RelationDefOperatorAndContext) {}

// ExitRelationDefOperatorAnd is called when production relationDefOperatorAnd is exited.
func (s *BaseOpenFGAListener) ExitRelationDefOperatorAnd(ctx *RelationDefOperatorAndContext) {}

// EnterRelationDefOperatorOr is called when production relationDefOperatorOr is entered.
func (s *BaseOpenFGAListener) EnterRelationDefOperatorOr(ctx *RelationDefOperatorOrContext) {}

// ExitRelationDefOperatorOr is called when production relationDefOperatorOr is exited.
func (s *BaseOpenFGAListener) ExitRelationDefOperatorOr(ctx *RelationDefOperatorOrContext) {}

// EnterRelationDefOperatorButNot is called when production relationDefOperatorButNot is entered.
func (s *BaseOpenFGAListener) EnterRelationDefOperatorButNot(ctx *RelationDefOperatorButNotContext) {}

// ExitRelationDefOperatorButNot is called when production relationDefOperatorButNot is exited.
func (s *BaseOpenFGAListener) ExitRelationDefOperatorButNot(ctx *RelationDefOperatorButNotContext) {}

// EnterRelationDefKeywordFrom is called when production relationDefKeywordFrom is entered.
func (s *BaseOpenFGAListener) EnterRelationDefKeywordFrom(ctx *RelationDefKeywordFromContext) {}

// ExitRelationDefKeywordFrom is called when production relationDefKeywordFrom is exited.
func (s *BaseOpenFGAListener) ExitRelationDefKeywordFrom(ctx *RelationDefKeywordFromContext) {}

// EnterRelationDefTypeRestriction is called when production relationDefTypeRestriction is entered.
func (s *BaseOpenFGAListener) EnterRelationDefTypeRestriction(ctx *RelationDefTypeRestrictionContext) {
}

// ExitRelationDefTypeRestriction is called when production relationDefTypeRestriction is exited.
func (s *BaseOpenFGAListener) ExitRelationDefTypeRestriction(ctx *RelationDefTypeRestrictionContext) {
}

// EnterRelationDefTypeRestrictionType is called when production relationDefTypeRestrictionType is entered.
func (s *BaseOpenFGAListener) EnterRelationDefTypeRestrictionType(ctx *RelationDefTypeRestrictionTypeContext) {
}

// ExitRelationDefTypeRestrictionType is called when production relationDefTypeRestrictionType is exited.
func (s *BaseOpenFGAListener) ExitRelationDefTypeRestrictionType(ctx *RelationDefTypeRestrictionTypeContext) {
}

// EnterRelationDefTypeRestrictionRelation is called when production relationDefTypeRestrictionRelation is entered.
func (s *BaseOpenFGAListener) EnterRelationDefTypeRestrictionRelation(ctx *RelationDefTypeRestrictionRelationContext) {
}

// ExitRelationDefTypeRestrictionRelation is called when production relationDefTypeRestrictionRelation is exited.
func (s *BaseOpenFGAListener) ExitRelationDefTypeRestrictionRelation(ctx *RelationDefTypeRestrictionRelationContext) {
}

// EnterRelationDefTypeRestrictionWildcard is called when production relationDefTypeRestrictionWildcard is entered.
func (s *BaseOpenFGAListener) EnterRelationDefTypeRestrictionWildcard(ctx *RelationDefTypeRestrictionWildcardContext) {
}

// ExitRelationDefTypeRestrictionWildcard is called when production relationDefTypeRestrictionWildcard is exited.
func (s *BaseOpenFGAListener) ExitRelationDefTypeRestrictionWildcard(ctx *RelationDefTypeRestrictionWildcardContext) {
}

// EnterRelationDefTypeRestrictionUserset is called when production relationDefTypeRestrictionUserset is entered.
func (s *BaseOpenFGAListener) EnterRelationDefTypeRestrictionUserset(ctx *RelationDefTypeRestrictionUsersetContext) {
}

// ExitRelationDefTypeRestrictionUserset is called when production relationDefTypeRestrictionUserset is exited.
func (s *BaseOpenFGAListener) ExitRelationDefTypeRestrictionUserset(ctx *RelationDefTypeRestrictionUsersetContext) {
}

// EnterRewriteComputedusersetName is called when production rewriteComputedusersetName is entered.
func (s *BaseOpenFGAListener) EnterRewriteComputedusersetName(ctx *RewriteComputedusersetNameContext) {
}

// ExitRewriteComputedusersetName is called when production rewriteComputedusersetName is exited.
func (s *BaseOpenFGAListener) ExitRewriteComputedusersetName(ctx *RewriteComputedusersetNameContext) {
}

// EnterRewriteTuplesetComputedusersetName is called when production rewriteTuplesetComputedusersetName is entered.
func (s *BaseOpenFGAListener) EnterRewriteTuplesetComputedusersetName(ctx *RewriteTuplesetComputedusersetNameContext) {
}

// ExitRewriteTuplesetComputedusersetName is called when production rewriteTuplesetComputedusersetName is exited.
func (s *BaseOpenFGAListener) ExitRewriteTuplesetComputedusersetName(ctx *RewriteTuplesetComputedusersetNameContext) {
}

// EnterRewriteTuplesetName is called when production rewriteTuplesetName is entered.
func (s *BaseOpenFGAListener) EnterRewriteTuplesetName(ctx *RewriteTuplesetNameContext) {}

// ExitRewriteTuplesetName is called when production rewriteTuplesetName is exited.
func (s *BaseOpenFGAListener) ExitRewriteTuplesetName(ctx *RewriteTuplesetNameContext) {}

// EnterRelationName is called when production relationName is entered.
func (s *BaseOpenFGAListener) EnterRelationName(ctx *RelationNameContext) {}

// ExitRelationName is called when production relationName is exited.
func (s *BaseOpenFGAListener) ExitRelationName(ctx *RelationNameContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseOpenFGAListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseOpenFGAListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseOpenFGAListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseOpenFGAListener) ExitComment(ctx *CommentContext) {}

// EnterMultiLineComment is called when production multiLineComment is entered.
func (s *BaseOpenFGAListener) EnterMultiLineComment(ctx *MultiLineCommentContext) {}

// ExitMultiLineComment is called when production multiLineComment is exited.
func (s *BaseOpenFGAListener) ExitMultiLineComment(ctx *MultiLineCommentContext) {}

// EnterSpacing is called when production spacing is entered.
func (s *BaseOpenFGAListener) EnterSpacing(ctx *SpacingContext) {}

// ExitSpacing is called when production spacing is exited.
func (s *BaseOpenFGAListener) ExitSpacing(ctx *SpacingContext) {}

// EnterNewline is called when production newline is entered.
func (s *BaseOpenFGAListener) EnterNewline(ctx *NewlineContext) {}

// ExitNewline is called when production newline is exited.
func (s *BaseOpenFGAListener) ExitNewline(ctx *NewlineContext) {}

// EnterSchemaVersion is called when production schemaVersion is entered.
func (s *BaseOpenFGAListener) EnterSchemaVersion(ctx *SchemaVersionContext) {}

// ExitSchemaVersion is called when production schemaVersion is exited.
func (s *BaseOpenFGAListener) ExitSchemaVersion(ctx *SchemaVersionContext) {}

// EnterName is called when production name is entered.
func (s *BaseOpenFGAListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseOpenFGAListener) ExitName(ctx *NameContext) {}
