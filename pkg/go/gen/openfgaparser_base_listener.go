// Code generated from /app/OpenFGAParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // OpenFGAParser

import "github.com/antlr4-go/antlr/v4"

// BaseOpenFGAParserListener is a complete listener for a parse tree produced by OpenFGAParser.
type BaseOpenFGAParserListener struct{}

var _ OpenFGAParserListener = &BaseOpenFGAParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOpenFGAParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOpenFGAParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOpenFGAParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOpenFGAParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMain is called when production main is entered.
func (s *BaseOpenFGAParserListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BaseOpenFGAParserListener) ExitMain(ctx *MainContext) {}

// EnterIndentation is called when production indentation is entered.
func (s *BaseOpenFGAParserListener) EnterIndentation(ctx *IndentationContext) {}

// ExitIndentation is called when production indentation is exited.
func (s *BaseOpenFGAParserListener) ExitIndentation(ctx *IndentationContext) {}

// EnterModelHeader is called when production modelHeader is entered.
func (s *BaseOpenFGAParserListener) EnterModelHeader(ctx *ModelHeaderContext) {}

// ExitModelHeader is called when production modelHeader is exited.
func (s *BaseOpenFGAParserListener) ExitModelHeader(ctx *ModelHeaderContext) {}

// EnterTypeDefs is called when production typeDefs is entered.
func (s *BaseOpenFGAParserListener) EnterTypeDefs(ctx *TypeDefsContext) {}

// ExitTypeDefs is called when production typeDefs is exited.
func (s *BaseOpenFGAParserListener) ExitTypeDefs(ctx *TypeDefsContext) {}

// EnterTypeDef is called when production typeDef is entered.
func (s *BaseOpenFGAParserListener) EnterTypeDef(ctx *TypeDefContext) {}

// ExitTypeDef is called when production typeDef is exited.
func (s *BaseOpenFGAParserListener) ExitTypeDef(ctx *TypeDefContext) {}

// EnterRelationDeclaration is called when production relationDeclaration is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDeclaration(ctx *RelationDeclarationContext) {}

// ExitRelationDeclaration is called when production relationDeclaration is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDeclaration(ctx *RelationDeclarationContext) {}

// EnterRelationDef is called when production relationDef is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDef(ctx *RelationDefContext) {}

// ExitRelationDef is called when production relationDef is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDef(ctx *RelationDefContext) {}

// EnterRelationDefPartials is called when production relationDefPartials is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefPartials(ctx *RelationDefPartialsContext) {}

// ExitRelationDefPartials is called when production relationDefPartials is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefPartials(ctx *RelationDefPartialsContext) {}

// EnterRelationDefPartialAllOr is called when production relationDefPartialAllOr is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefPartialAllOr(ctx *RelationDefPartialAllOrContext) {
}

// ExitRelationDefPartialAllOr is called when production relationDefPartialAllOr is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefPartialAllOr(ctx *RelationDefPartialAllOrContext) {
}

// EnterRelationDefPartialAllAnd is called when production relationDefPartialAllAnd is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefPartialAllAnd(ctx *RelationDefPartialAllAndContext) {
}

// ExitRelationDefPartialAllAnd is called when production relationDefPartialAllAnd is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefPartialAllAnd(ctx *RelationDefPartialAllAndContext) {
}

// EnterRelationDefPartialAllButNot is called when production relationDefPartialAllButNot is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefPartialAllButNot(ctx *RelationDefPartialAllButNotContext) {
}

// ExitRelationDefPartialAllButNot is called when production relationDefPartialAllButNot is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefPartialAllButNot(ctx *RelationDefPartialAllButNotContext) {
}

// EnterRelationDefDirectAssignment is called when production relationDefDirectAssignment is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefDirectAssignment(ctx *RelationDefDirectAssignmentContext) {
}

// ExitRelationDefDirectAssignment is called when production relationDefDirectAssignment is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefDirectAssignment(ctx *RelationDefDirectAssignmentContext) {
}

// EnterRelationDefRewrite is called when production relationDefRewrite is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefRewrite(ctx *RelationDefRewriteContext) {}

// ExitRelationDefRewrite is called when production relationDefRewrite is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefRewrite(ctx *RelationDefRewriteContext) {}

// EnterRelationDefRelationOnSameObject is called when production relationDefRelationOnSameObject is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefRelationOnSameObject(ctx *RelationDefRelationOnSameObjectContext) {
}

// ExitRelationDefRelationOnSameObject is called when production relationDefRelationOnSameObject is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefRelationOnSameObject(ctx *RelationDefRelationOnSameObjectContext) {
}

// EnterRelationDefRelationOnRelatedObject is called when production relationDefRelationOnRelatedObject is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefRelationOnRelatedObject(ctx *RelationDefRelationOnRelatedObjectContext) {
}

// ExitRelationDefRelationOnRelatedObject is called when production relationDefRelationOnRelatedObject is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefRelationOnRelatedObject(ctx *RelationDefRelationOnRelatedObjectContext) {
}

// EnterRelationDefOperator is called when production relationDefOperator is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefOperator(ctx *RelationDefOperatorContext) {}

// ExitRelationDefOperator is called when production relationDefOperator is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefOperator(ctx *RelationDefOperatorContext) {}

// EnterRelationDefOperatorAnd is called when production relationDefOperatorAnd is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefOperatorAnd(ctx *RelationDefOperatorAndContext) {}

// ExitRelationDefOperatorAnd is called when production relationDefOperatorAnd is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefOperatorAnd(ctx *RelationDefOperatorAndContext) {}

// EnterRelationDefOperatorOr is called when production relationDefOperatorOr is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefOperatorOr(ctx *RelationDefOperatorOrContext) {}

// ExitRelationDefOperatorOr is called when production relationDefOperatorOr is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefOperatorOr(ctx *RelationDefOperatorOrContext) {}

// EnterRelationDefOperatorButNot is called when production relationDefOperatorButNot is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefOperatorButNot(ctx *RelationDefOperatorButNotContext) {
}

// ExitRelationDefOperatorButNot is called when production relationDefOperatorButNot is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefOperatorButNot(ctx *RelationDefOperatorButNotContext) {
}

// EnterRelationDefKeywordFrom is called when production relationDefKeywordFrom is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefKeywordFrom(ctx *RelationDefKeywordFromContext) {}

// ExitRelationDefKeywordFrom is called when production relationDefKeywordFrom is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefKeywordFrom(ctx *RelationDefKeywordFromContext) {}

// EnterRelationDefTypeRestriction is called when production relationDefTypeRestriction is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefTypeRestriction(ctx *RelationDefTypeRestrictionContext) {
}

// ExitRelationDefTypeRestriction is called when production relationDefTypeRestriction is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefTypeRestriction(ctx *RelationDefTypeRestrictionContext) {
}

// EnterRelationDefTypeRestrictionType is called when production relationDefTypeRestrictionType is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefTypeRestrictionType(ctx *RelationDefTypeRestrictionTypeContext) {
}

// ExitRelationDefTypeRestrictionType is called when production relationDefTypeRestrictionType is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefTypeRestrictionType(ctx *RelationDefTypeRestrictionTypeContext) {
}

// EnterRelationDefTypeRestrictionRelation is called when production relationDefTypeRestrictionRelation is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefTypeRestrictionRelation(ctx *RelationDefTypeRestrictionRelationContext) {
}

// ExitRelationDefTypeRestrictionRelation is called when production relationDefTypeRestrictionRelation is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefTypeRestrictionRelation(ctx *RelationDefTypeRestrictionRelationContext) {
}

// EnterRelationDefTypeRestrictionWildcard is called when production relationDefTypeRestrictionWildcard is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefTypeRestrictionWildcard(ctx *RelationDefTypeRestrictionWildcardContext) {
}

// ExitRelationDefTypeRestrictionWildcard is called when production relationDefTypeRestrictionWildcard is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefTypeRestrictionWildcard(ctx *RelationDefTypeRestrictionWildcardContext) {
}

// EnterRelationDefTypeRestrictionUserset is called when production relationDefTypeRestrictionUserset is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefTypeRestrictionUserset(ctx *RelationDefTypeRestrictionUsersetContext) {
}

// ExitRelationDefTypeRestrictionUserset is called when production relationDefTypeRestrictionUserset is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefTypeRestrictionUserset(ctx *RelationDefTypeRestrictionUsersetContext) {
}

// EnterRelationDefGrouping is called when production relationDefGrouping is entered.
func (s *BaseOpenFGAParserListener) EnterRelationDefGrouping(ctx *RelationDefGroupingContext) {}

// ExitRelationDefGrouping is called when production relationDefGrouping is exited.
func (s *BaseOpenFGAParserListener) ExitRelationDefGrouping(ctx *RelationDefGroupingContext) {}

// EnterRewriteComputedusersetName is called when production rewriteComputedusersetName is entered.
func (s *BaseOpenFGAParserListener) EnterRewriteComputedusersetName(ctx *RewriteComputedusersetNameContext) {
}

// ExitRewriteComputedusersetName is called when production rewriteComputedusersetName is exited.
func (s *BaseOpenFGAParserListener) ExitRewriteComputedusersetName(ctx *RewriteComputedusersetNameContext) {
}

// EnterRewriteTuplesetComputedusersetName is called when production rewriteTuplesetComputedusersetName is entered.
func (s *BaseOpenFGAParserListener) EnterRewriteTuplesetComputedusersetName(ctx *RewriteTuplesetComputedusersetNameContext) {
}

// ExitRewriteTuplesetComputedusersetName is called when production rewriteTuplesetComputedusersetName is exited.
func (s *BaseOpenFGAParserListener) ExitRewriteTuplesetComputedusersetName(ctx *RewriteTuplesetComputedusersetNameContext) {
}

// EnterRewriteTuplesetName is called when production rewriteTuplesetName is entered.
func (s *BaseOpenFGAParserListener) EnterRewriteTuplesetName(ctx *RewriteTuplesetNameContext) {}

// ExitRewriteTuplesetName is called when production rewriteTuplesetName is exited.
func (s *BaseOpenFGAParserListener) ExitRewriteTuplesetName(ctx *RewriteTuplesetNameContext) {}

// EnterRelationName is called when production relationName is entered.
func (s *BaseOpenFGAParserListener) EnterRelationName(ctx *RelationNameContext) {}

// ExitRelationName is called when production relationName is exited.
func (s *BaseOpenFGAParserListener) ExitRelationName(ctx *RelationNameContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseOpenFGAParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseOpenFGAParserListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseOpenFGAParserListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseOpenFGAParserListener) ExitComment(ctx *CommentContext) {}

// EnterMultiLineComment is called when production multiLineComment is entered.
func (s *BaseOpenFGAParserListener) EnterMultiLineComment(ctx *MultiLineCommentContext) {}

// ExitMultiLineComment is called when production multiLineComment is exited.
func (s *BaseOpenFGAParserListener) ExitMultiLineComment(ctx *MultiLineCommentContext) {}

// EnterSpacing is called when production spacing is entered.
func (s *BaseOpenFGAParserListener) EnterSpacing(ctx *SpacingContext) {}

// ExitSpacing is called when production spacing is exited.
func (s *BaseOpenFGAParserListener) ExitSpacing(ctx *SpacingContext) {}

// EnterNewline is called when production newline is entered.
func (s *BaseOpenFGAParserListener) EnterNewline(ctx *NewlineContext) {}

// ExitNewline is called when production newline is exited.
func (s *BaseOpenFGAParserListener) ExitNewline(ctx *NewlineContext) {}

// EnterSchemaVersion is called when production schemaVersion is entered.
func (s *BaseOpenFGAParserListener) EnterSchemaVersion(ctx *SchemaVersionContext) {}

// ExitSchemaVersion is called when production schemaVersion is exited.
func (s *BaseOpenFGAParserListener) ExitSchemaVersion(ctx *SchemaVersionContext) {}

// EnterName is called when production name is entered.
func (s *BaseOpenFGAParserListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseOpenFGAParserListener) ExitName(ctx *NameContext) {}
