package transformer

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/hashicorp/go-multierror"
	pb "github.com/openfga/api/proto/openfga/v1"
	parser "github.com/openfga/language/pkg/go/gen"
)

type RelationDefinitionOperator string

const (
	RELATION_DEFINITION_OPERATOR_NONE    RelationDefinitionOperator = ""
	RELATION_DEFINITION_OPERATOR_OR      RelationDefinitionOperator = "or"
	RELATION_DEFINITION_OPERATOR_AND     RelationDefinitionOperator = "and"
	RELATION_DEFINITION_OPERATOR_BUT_NOT RelationDefinitionOperator = "but not"
)

// OpenFGA DSL Listener

type relation struct {
	Name     string
	Rewrites []*pb.Userset
	Operator RelationDefinitionOperator
	TypeInfo pb.RelationTypeInfo
}

type openFgaDslListener struct {
	*parser.BaseOpenFGAParserListener

	authorizationModel pb.AuthorizationModel
	currentTypeDef     *pb.TypeDefinition
	currentRelation    *relation
}

func newOpenFgaDslListener() *openFgaDslListener {
	return new(openFgaDslListener)
}

func (l *openFgaDslListener) ExitSchemaVersion(ctx *parser.SchemaVersionContext) {
	l.authorizationModel.SchemaVersion = ctx.GetText()
}

func (l *openFgaDslListener) EnterTypeDef(_ctx *parser.TypeDefContext) {
	l.currentTypeDef = &pb.TypeDefinition{
		Relations: map[string]*pb.Userset{},
		Metadata: &pb.Metadata{
			Relations: map[string]*pb.RelationMetadata{},
		},
	}
}

func (l *openFgaDslListener) ExitTypeDef(ctx *parser.TypeDefContext) {
	typeName := ctx.TypeName().GetText()
	typeDef := l.currentTypeDef
	typeDef.Type = typeName

	if len(l.currentTypeDef.Metadata.Relations) == 0 {
		l.currentTypeDef.Metadata = nil
	}

	l.authorizationModel.TypeDefinitions = append(l.authorizationModel.TypeDefinitions, l.currentTypeDef)

	l.currentTypeDef = nil
}

func (l *openFgaDslListener) EnterRelationDeclaration(_ctx *parser.RelationDeclarationContext) {
	l.currentRelation = &relation{
		Rewrites: []*pb.Userset{},
		TypeInfo: pb.RelationTypeInfo{DirectlyRelatedUserTypes: []*pb.RelationReference{}},
	}
}

func (l *openFgaDslListener) ExitRelationDeclaration(ctx *parser.RelationDeclarationContext) {
	relationName := ctx.RelationName().GetText()

	if len(l.currentRelation.Rewrites) == 1 {
		l.currentTypeDef.Relations[relationName] = l.currentRelation.Rewrites[0]
	} else {
		if l.currentRelation.Operator == RELATION_DEFINITION_OPERATOR_OR {
			l.currentTypeDef.Relations[relationName] = &pb.Userset{
				Userset: &pb.Userset_Union{
					Union: &pb.Usersets{
						Child: l.currentRelation.Rewrites,
					},
				},
			}
		} else if l.currentRelation.Operator == RELATION_DEFINITION_OPERATOR_AND {
			l.currentTypeDef.Relations[relationName] = &pb.Userset{
				Userset: &pb.Userset_Intersection{
					Intersection: &pb.Usersets{
						Child: l.currentRelation.Rewrites,
					},
				},
			}
		} else if l.currentRelation.Operator == RELATION_DEFINITION_OPERATOR_BUT_NOT {
			l.currentTypeDef.Relations[relationName] = &pb.Userset{
				Userset: &pb.Userset_Difference{
					Difference: &pb.Difference{
						Base:     l.currentRelation.Rewrites[0],
						Subtract: l.currentRelation.Rewrites[1],
					},
				},
			}
		}
	}

	directlyRelatedUserTypes := l.currentRelation.TypeInfo.GetDirectlyRelatedUserTypes()

	l.currentTypeDef.Metadata.Relations[relationName] = &pb.RelationMetadata{DirectlyRelatedUserTypes: directlyRelatedUserTypes}

	l.currentRelation = nil
}

func (l *openFgaDslListener) EnterRelationDefDirectAssignment(_ctx *parser.RelationDefDirectAssignmentContext) {
	l.currentRelation.TypeInfo = pb.RelationTypeInfo{DirectlyRelatedUserTypes: []*pb.RelationReference{}}
}

func (l *openFgaDslListener) ExitRelationDefDirectAssignment(_ctx *parser.RelationDefDirectAssignmentContext) {
	partialRewrite := &pb.Userset{Userset: &pb.Userset_This{}}

	l.currentRelation.Rewrites = append(l.currentRelation.Rewrites, partialRewrite)
}

func (l *openFgaDslListener) ExitRelationDefTypeRestriction(ctx *parser.RelationDefTypeRestrictionContext) {
	relationRef := &pb.RelationReference{}
	_type := ctx.RelationDefTypeRestrictionType()
	usersetRestriction := ctx.RelationDefTypeRestrictionUserset()
	wildcardRestriction := ctx.RelationDefTypeRestrictionWildcard()

	if _type != nil {
		relationRef.Type = _type.GetText()
	}

	if usersetRestriction != nil {
		relationRef.Type = usersetRestriction.RelationDefTypeRestrictionType().GetText()
		relationRef.RelationOrWildcard = &pb.RelationReference_Relation{
			Relation: usersetRestriction.RelationDefTypeRestrictionRelation().GetText(),
		}
	}

	if wildcardRestriction != nil {
		relationRef.Type = wildcardRestriction.RelationDefTypeRestrictionType().GetText()
		relationRef.RelationOrWildcard = &pb.RelationReference_Wildcard{Wildcard: &pb.Wildcard{}}
	}

	l.currentRelation.TypeInfo.DirectlyRelatedUserTypes = append(l.currentRelation.TypeInfo.DirectlyRelatedUserTypes, relationRef)
}

func (l *openFgaDslListener) ExitRelationDefRelationOnSameObject(ctx *parser.RelationDefRelationOnSameObjectContext) {
	partialRewrite := &pb.Userset{Userset: &pb.Userset_ComputedUserset{
		ComputedUserset: &pb.ObjectRelation{
			Object:   "",
			Relation: ctx.RewriteComputedusersetName().GetText(),
		},
	}}
	l.currentRelation.Rewrites = append(l.currentRelation.Rewrites, partialRewrite)
}

func (l *openFgaDslListener) ExitRelationDefRelationOnRelatedObject(
	ctx *parser.RelationDefRelationOnRelatedObjectContext,
) {
	partialRewrite := &pb.Userset{Userset: &pb.Userset_TupleToUserset{
		TupleToUserset: &pb.TupleToUserset{
			ComputedUserset: &pb.ObjectRelation{
				Object:   "",
				Relation: ctx.RewriteTuplesetComputedusersetName().GetText(),
			},
			Tupleset: &pb.ObjectRelation{
				Relation: ctx.RewriteTuplesetName().GetText(),
			},
		},
	}}
	l.currentRelation.Rewrites = append(l.currentRelation.Rewrites, partialRewrite)
}

func (l *openFgaDslListener) EnterRelationDefPartialAllOr(_ctx *parser.RelationDefPartialAllOrContext) {
	l.currentRelation.Operator = RELATION_DEFINITION_OPERATOR_OR
}

func (l *openFgaDslListener) EnterRelationDefPartialAllAnd(_ctx *parser.RelationDefPartialAllAndContext) {
	l.currentRelation.Operator = RELATION_DEFINITION_OPERATOR_AND
}

func (l *openFgaDslListener) EnterRelationDefPartialAllButNot(_ctx *parser.RelationDefPartialAllButNotContext) {
	l.currentRelation.Operator = RELATION_DEFINITION_OPERATOR_BUT_NOT
}

//// Error Handling

type OpenFgaDslSyntaxError struct {
	line, column int
	msg          string
}

func (err *OpenFgaDslSyntaxError) Error() string {
	return fmt.Sprintf("syntax error at line=%d, column=%d: %s", err.line, err.column, err.msg)
}

type openFgaDslErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      *multierror.Error
}

func (c *openFgaDslErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = multierror.Append(c.Errors, &OpenFgaDslSyntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

///

// TransformDslToJSON - Converts models authored in FGA DSL syntax to the json syntax accepted by the OpenFGA API
func TransformDslToJSON(data string) (*pb.AuthorizationModel, error) {
	is := antlr.NewInputStream(data)

	errorListener := openFgaDslErrorListener{}

	// Create the Lexer
	lexer := parser.NewOpenFGALexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&errorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewOpenFGAParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&errorListener)

	// Finally parse the expression
	l := newOpenFgaDslListener()
	antlr.ParseTreeWalkerDefault.Walk(l, p.Main())

	if errorListener.Errors != nil {
		return nil, errorListener.Errors
	}

	return &l.authorizationModel, nil
}
