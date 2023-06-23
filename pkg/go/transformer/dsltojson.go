package transformer

import (
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/openfga/language/pkg/go/gen"
	pb "go.buf.build/openfga/go/openfga/api/openfga/v1"
)

type RelationDefinitionOperator string

const (
	RELATION_DEFINITION_OPERATOR_NONE    RelationDefinitionOperator = ""
	RELATION_DEFINITION_OPERATOR_OR      RelationDefinitionOperator = "or"
	RELATION_DEFINITION_OPERATOR_AND     RelationDefinitionOperator = "and"
	RELATION_DEFINITION_OPERATOR_BUT_NOT RelationDefinitionOperator = "but not"
)

type Relation struct {
	Name     string
	Rewrites []*pb.Userset
	Operator RelationDefinitionOperator
	TypeInfo pb.RelationTypeInfo
}

type OpenFgaDslListener struct {
	*parser.BaseOpenFGAListener

	authorizationModel pb.AuthorizationModel
	currentTypeDef     *pb.TypeDefinition
	currentRelation    *Relation
}

func NewOpenFgaDslListener() *OpenFgaDslListener {
	return new(OpenFgaDslListener)
}

func (l *OpenFgaDslListener) ExitSchemaVersion(ctx *parser.SchemaVersionContext) {
	l.authorizationModel.SchemaVersion = ctx.GetText()
}

func (l *OpenFgaDslListener) EnterTypeDef(_ctx *parser.TypeDefContext) {
	l.currentTypeDef = &pb.TypeDefinition{
		Relations: map[string]*pb.Userset{},
		Metadata:  nil,
	}
}

func (l *OpenFgaDslListener) ExitTypeDef(ctx *parser.TypeDefContext) {
	typeName := ctx.TypeName().GetText()
	typeDef := l.currentTypeDef
	typeDef.Type = typeName

	l.authorizationModel.TypeDefinitions = append(l.authorizationModel.TypeDefinitions, l.currentTypeDef)

	l.currentTypeDef = nil
}

func (l *OpenFgaDslListener) EnterRelationDeclaration(_ctx *parser.RelationDeclarationContext) {
	l.currentRelation = &Relation{
		Rewrites: []*pb.Userset{},
		TypeInfo: pb.RelationTypeInfo{DirectlyRelatedUserTypes: []*pb.RelationReference{}},
	}
}

func (l *OpenFgaDslListener) ExitRelationDeclaration(ctx *parser.RelationDeclarationContext) {
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
	if len(directlyRelatedUserTypes) > 0 {
		if l.currentTypeDef.Metadata == nil {
			l.currentTypeDef.Metadata = &pb.Metadata{
				Relations: map[string]*pb.RelationMetadata{},
			}
		}

		l.currentTypeDef.Metadata.Relations[relationName] = &pb.RelationMetadata{DirectlyRelatedUserTypes: directlyRelatedUserTypes}
	} else if l.currentTypeDef.Metadata != nil {
		l.currentTypeDef.Metadata.Relations[relationName] = &pb.RelationMetadata{DirectlyRelatedUserTypes: []*pb.RelationReference{}}
	}

	l.currentRelation = nil
}

func (l *OpenFgaDslListener) EnterRelationDefDirectAssignment(_ctx *parser.RelationDefDirectAssignmentContext) {
	l.currentRelation.TypeInfo = pb.RelationTypeInfo{DirectlyRelatedUserTypes: []*pb.RelationReference{}}
}

func (l *OpenFgaDslListener) ExitRelationDefDirectAssignment(_ctx *parser.RelationDefDirectAssignmentContext) {
	partialRewrite := &pb.Userset{Userset: &pb.Userset_This{}}

	l.currentRelation.Rewrites = append(l.currentRelation.Rewrites, partialRewrite)
}

func (l *OpenFgaDslListener) ExitRelationDefTypeRestriction(ctx *parser.RelationDefTypeRestrictionContext) {
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

func (l *OpenFgaDslListener) ExitRelationDefRelationOnSameObject(ctx *parser.RelationDefRelationOnSameObjectContext) {
	partialRewrite := &pb.Userset{Userset: &pb.Userset_ComputedUserset{
		ComputedUserset: &pb.ObjectRelation{
			Object:   "",
			Relation: ctx.RewriteComputedusersetName().GetText(),
		},
	}}
	l.currentRelation.Rewrites = append(l.currentRelation.Rewrites, partialRewrite)
}

func (l *OpenFgaDslListener) ExitRelationDefRelationOnRelatedObject(
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

func (l *OpenFgaDslListener) EnterRelationDefPartialAllOr(_ctx *parser.RelationDefPartialAllOrContext) {
	l.currentRelation.Operator = RELATION_DEFINITION_OPERATOR_OR
}

func (l *OpenFgaDslListener) EnterRelationDefPartialAllAnd(_ctx *parser.RelationDefPartialAllAndContext) {
	l.currentRelation.Operator = RELATION_DEFINITION_OPERATOR_AND
}

func (l *OpenFgaDslListener) EnterRelationDefPartialAllButNot(_ctx *parser.RelationDefPartialAllButNotContext) {
	l.currentRelation.Operator = RELATION_DEFINITION_OPERATOR_BUT_NOT
}

func TransformDslToJSON(data string) *pb.AuthorizationModel {
	is := antlr.NewInputStream(data)

	// Create the Lexer
	lexer := parser.NewOpenFGALexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewOpenFGAParser(stream)

	// Finally parse the expression
	l := NewOpenFgaDslListener()
	antlr.ParseTreeWalkerDefault.Walk(l, p.Main())

	return &l.authorizationModel
}
