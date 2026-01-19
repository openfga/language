package transformer

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/hashicorp/go-multierror"
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"google.golang.org/protobuf/encoding/protojson"

	parser "github.com/openfga/language/pkg/go/gen"
)

type RelationDefinitionOperator string

const (
	RELATION_DEFINITION_OPERATOR_NONE    RelationDefinitionOperator = ""        //nolint:stylecheck,revive
	RELATION_DEFINITION_OPERATOR_OR      RelationDefinitionOperator = "or"      //nolint:stylecheck,revive
	RELATION_DEFINITION_OPERATOR_AND     RelationDefinitionOperator = "and"     //nolint:stylecheck,revive
	RELATION_DEFINITION_OPERATOR_BUT_NOT RelationDefinitionOperator = "but not" //nolint:stylecheck,revive
)

// OpenFGA DSL Listener

type relation struct {
	Name     string
	Rewrites []*openfgav1.Userset
	Operator RelationDefinitionOperator
	TypeInfo openfgav1.RelationTypeInfo
}

type stackRelation struct {
	Rewrites []*openfgav1.Userset
	Operator RelationDefinitionOperator
}

type OpenFgaDslListener struct {
	*parser.BaseOpenFGAParserListener

	authorizationModel openfgav1.AuthorizationModel
	currentTypeDef     *openfgav1.TypeDefinition
	currentRelation    *relation
	currentCondition   *openfgav1.Condition
	rewriteStack       []*stackRelation
	isModularModel     bool
	moduleName         string
	typeDefExtensions  map[string]*openfgav1.TypeDefinition

	// Comment tracking
	commentTracker   *CommentTracker
	modelComments    *ModelComments
	typeComments     map[string]*CommentsMetadata
	conditionComments map[string]*CommentInfo
}

func newOpenFgaDslListener() *OpenFgaDslListener {
	return &OpenFgaDslListener{
		typeComments:      make(map[string]*CommentsMetadata),
		conditionComments: make(map[string]*CommentInfo),
	}
}

func ParseExpression(rewrites []*openfgav1.Userset, operator RelationDefinitionOperator) *openfgav1.Userset {
	var relationDef *openfgav1.Userset

	if len(rewrites) == 0 {
		return nil
	}

	if len(rewrites) == 1 {
		relationDef = rewrites[0]
	} else {
		switch operator {
		case RELATION_DEFINITION_OPERATOR_NONE:
			// do nothing
		case RELATION_DEFINITION_OPERATOR_OR:
			relationDef = &openfgav1.Userset{
				Userset: &openfgav1.Userset_Union{
					Union: &openfgav1.Usersets{
						Child: rewrites,
					},
				},
			}
		case RELATION_DEFINITION_OPERATOR_AND:
			relationDef = &openfgav1.Userset{
				Userset: &openfgav1.Userset_Intersection{
					Intersection: &openfgav1.Usersets{
						Child: rewrites,
					},
				},
			}
		case RELATION_DEFINITION_OPERATOR_BUT_NOT:
			relationDef = &openfgav1.Userset{
				Userset: &openfgav1.Userset_Difference{
					Difference: &openfgav1.Difference{
						Base:     rewrites[0],
						Subtract: rewrites[1],
					},
				},
			}
		}
	}

	return relationDef
}

func (l *OpenFgaDslListener) EnterMain(_ *parser.MainContext) {
	l.authorizationModel.Conditions = map[string]*openfgav1.Condition{}
}

func (l *OpenFgaDslListener) ExitModuleHeader(ctx *parser.ModuleHeaderContext) {
	l.isModularModel = true
	if ctx.GetModuleName() != nil {
		l.moduleName = ctx.GetModuleName().GetText()
		l.typeDefExtensions = map[string]*openfgav1.TypeDefinition{}
	}
}

func (l *OpenFgaDslListener) ExitModelHeader(ctx *parser.ModelHeaderContext) {
	if ctx.GetSchemaVersion() != nil {
		l.authorizationModel.SchemaVersion = ctx.GetSchemaVersion().GetText()
	}
}

func (l *OpenFgaDslListener) EnterTypeDef(ctx *parser.TypeDefContext) {
	if ctx.GetTypeName() == nil {
		return
	}

	if ctx.EXTEND() != nil && !l.isModularModel {
		ctx.GetParser().NotifyErrorListeners(
			"extend can only be used in a modular model",
			ctx.GetTypeName().GetStart(),
			nil,
		)
	}

	typeName := ctx.GetTypeName().GetText()
	l.currentTypeDef = &openfgav1.TypeDefinition{
		Type:      typeName,
		Relations: map[string]*openfgav1.Userset{},
		Metadata: &openfgav1.Metadata{
			Relations: map[string]*openfgav1.RelationMetadata{},
		},
	}

	if l.isModularModel {
		l.currentTypeDef.Metadata.Module = l.moduleName
	}

	// Track type comments (line is 1-based from ANTLR, convert to 0-based)
	// Use TYPE token's line number instead of GetStart() which may include preceding newlines
	if l.commentTracker != nil && ctx.TYPE() != nil {
		lineNum := ctx.TYPE().GetSymbol().GetLine() - 1
		commentInfo := l.commentTracker.GetCommentInfoForLine(lineNum)
		if commentInfo != nil {
			if l.typeComments[typeName] == nil {
				l.typeComments[typeName] = &CommentsMetadata{}
			}
			l.typeComments[typeName].Comments = commentInfo
		}
	}
}

func (l *OpenFgaDslListener) EnterConditions(_ *parser.ConditionsContext) {
	l.authorizationModel.Conditions = map[string]*openfgav1.Condition{}
}

func (l *OpenFgaDslListener) EnterCondition(ctx *parser.ConditionContext) {
	if ctx.ConditionName() == nil {
		return
	}

	conditionName := ctx.ConditionName().GetText()
	if l.authorizationModel.GetConditions()[conditionName] != nil {
		ctx.GetParser().NotifyErrorListeners(
			fmt.Sprintf("condition '%s' is already defined in the model", conditionName),
			ctx.ConditionName().GetStart(),
			nil)
	}

	l.currentCondition = &openfgav1.Condition{
		Name:       conditionName,
		Expression: "",
		Parameters: map[string]*openfgav1.ConditionParamTypeRef{},
	}

	if l.isModularModel {
		l.currentCondition.Metadata = &openfgav1.ConditionMetadata{
			Module: l.moduleName,
		}
	}

	// Track condition comments (line is 1-based from ANTLR, convert to 0-based)
	// Use CONDITION token's line number instead of GetStart() which may include preceding newlines
	if l.commentTracker != nil && ctx.CONDITION() != nil {
		lineNum := ctx.CONDITION().GetSymbol().GetLine() - 1
		commentInfo := l.commentTracker.GetCommentInfoForLine(lineNum)
		if commentInfo != nil {
			l.conditionComments[conditionName] = commentInfo
		}
	}
}

func (l *OpenFgaDslListener) ExitConditionParameter(ctx *parser.ConditionParameterContext) {
	if ctx.ParameterName() == nil || ctx.ParameterType() == nil {
		return
	}

	parameterName := ctx.ParameterName().GetText()
	if l.currentCondition.GetParameters()[parameterName] != nil {
		ctx.GetParser().NotifyErrorListeners(
			fmt.Sprintf("parameter '%s' is already defined in the condition '%s'", parameterName, l.currentCondition.GetName()),
			ctx.ParameterName().GetStart(),
			nil)
	}

	paramContainer := ctx.ParameterType().CONDITION_PARAM_CONTAINER()
	typeNameString := ctx.ParameterType().GetText()

	var genericName *openfgav1.ConditionParamTypeRef_TypeName

	if paramContainer != nil {
		typeNameString = paramContainer.GetText()
		genericType := ctx.ParameterType().CONDITION_PARAM_TYPE()

		if genericType != nil {
			genericString := ctx.ParameterType().CONDITION_PARAM_TYPE().GetText()
			genericName = new(openfgav1.ConditionParamTypeRef_TypeName)
			*genericName = openfgav1.ConditionParamTypeRef_TypeName(
				openfgav1.ConditionParamTypeRef_TypeName_value["TYPE_NAME_"+strings.ToUpper(genericString)],
			)
		}
	}

	typeName := new(openfgav1.ConditionParamTypeRef_TypeName)
	*typeName = openfgav1.ConditionParamTypeRef_TypeName(
		openfgav1.ConditionParamTypeRef_TypeName_value["TYPE_NAME_"+strings.ToUpper(typeNameString)],
	)
	conditionParamTypeRef := &openfgav1.ConditionParamTypeRef{
		TypeName:     *typeName,
		GenericTypes: []*openfgav1.ConditionParamTypeRef{},
	}

	if genericName != nil {
		conditionParamTypeRef.GenericTypes = append(conditionParamTypeRef.GenericTypes, &openfgav1.ConditionParamTypeRef{
			TypeName: *genericName,
		})
	}

	l.currentCondition.Parameters[parameterName] = conditionParamTypeRef
}

func (l *OpenFgaDslListener) ExitConditionExpression(ctx *parser.ConditionExpressionContext) {
	l.currentCondition.Expression = strings.TrimRight(ctx.GetText(), "\n")
}

func (l *OpenFgaDslListener) ExitCondition(_ *parser.ConditionContext) {
	if l.currentCondition != nil {
		l.authorizationModel.Conditions[l.currentCondition.GetName()] = l.currentCondition

		l.currentCondition = nil
	}
}

func (l *OpenFgaDslListener) ExitTypeDef(ctx *parser.TypeDefContext) {
	if l.currentTypeDef == nil || l.currentTypeDef.GetType() == "" {
		return
	}

	if !l.isModularModel && len(l.currentTypeDef.GetMetadata().GetRelations()) == 0 {
		l.currentTypeDef.Metadata = nil
	} else if l.isModularModel && len(l.currentTypeDef.GetMetadata().GetRelations()) == 0 {
		l.currentTypeDef.Metadata.Relations = nil
	}

	l.authorizationModel.TypeDefinitions = append(l.authorizationModel.TypeDefinitions, l.currentTypeDef)

	if ctx.EXTEND() != nil && l.isModularModel {
		if typeDefExists := l.typeDefExtensions[l.currentTypeDef.GetType()]; typeDefExists != nil {
			ctx.GetParser().NotifyErrorListeners(
				fmt.Sprintf("'%s' is already extended in file.", l.currentTypeDef.GetType()),
				ctx.GetTypeName().GetStart(),
				nil,
			)
		} else {
			l.typeDefExtensions[l.currentTypeDef.GetType()] = l.currentTypeDef
		}
	}

	l.currentTypeDef = nil
}

func (l *OpenFgaDslListener) EnterRelationDeclaration(_ *parser.RelationDeclarationContext) {
	l.currentRelation = &relation{
		Rewrites: []*openfgav1.Userset{},
		TypeInfo: openfgav1.RelationTypeInfo{DirectlyRelatedUserTypes: []*openfgav1.RelationReference{}},
	}

	l.rewriteStack = []*stackRelation{}
}

func (l *OpenFgaDslListener) ExitRelationDeclaration(ctx *parser.RelationDeclarationContext) {
	if ctx.RelationName() == nil {
		return
	}

	relationName := ctx.RelationName().GetText()
	relationDef := ParseExpression(l.currentRelation.Rewrites, l.currentRelation.Operator)

	if relationDef != nil {
		if l.currentTypeDef.GetRelations()[relationName] != nil {
			ctx.GetParser().NotifyErrorListeners(
				fmt.Sprintf("'%s' is already defined in '%s'", relationName, l.currentTypeDef.GetType()),
				ctx.RelationName().GetStart(),
				nil)
		}

		l.currentTypeDef.Relations[relationName] = relationDef
		directlyRelatedUserTypes := l.currentRelation.TypeInfo.GetDirectlyRelatedUserTypes()
		l.currentTypeDef.Metadata.Relations[relationName] = &openfgav1.RelationMetadata{
			DirectlyRelatedUserTypes: directlyRelatedUserTypes,
		}

		isExtension := false
		if parent, ok := ctx.GetParent().(*parser.TypeDefContext); ok {
			isExtension = parent.EXTEND() != nil
		}

		if l.isModularModel && isExtension {
			l.currentTypeDef.Metadata.Relations[relationName].Module = l.moduleName
		}

		// Track relation comments (line is 1-based from ANTLR, convert to 0-based)
		// Use DEFINE token's line number instead of GetStart() which may include preceding newlines
		if l.commentTracker != nil && l.currentTypeDef != nil && ctx.DEFINE() != nil {
			typeName := l.currentTypeDef.GetType()
			lineNum := ctx.DEFINE().GetSymbol().GetLine() - 1
			commentInfo := l.commentTracker.GetCommentInfoForLine(lineNum)
			if commentInfo != nil {
				if l.typeComments[typeName] == nil {
					l.typeComments[typeName] = &CommentsMetadata{}
				}
				if l.typeComments[typeName].RelationComments == nil {
					l.typeComments[typeName].RelationComments = make(map[string]*CommentInfo)
				}
				l.typeComments[typeName].RelationComments[relationName] = commentInfo
			}
		}
	}

	l.currentRelation = nil
}

func (l *OpenFgaDslListener) EnterRelationDefDirectAssignment(_ *parser.RelationDefDirectAssignmentContext) {
	l.currentRelation.TypeInfo = openfgav1.RelationTypeInfo{DirectlyRelatedUserTypes: []*openfgav1.RelationReference{}}
}

func (l *OpenFgaDslListener) ExitRelationDefDirectAssignment(_ *parser.RelationDefDirectAssignmentContext) {
	partialRewrite := &openfgav1.Userset{Userset: &openfgav1.Userset_This{}}

	l.currentRelation.Rewrites = append(l.currentRelation.Rewrites, partialRewrite)
}

func (l *OpenFgaDslListener) ExitRelationDefTypeRestriction(ctx *parser.RelationDefTypeRestrictionContext) {
	baseRestriction := ctx.RelationDefTypeRestrictionBase()
	if baseRestriction == nil {
		return
	}

	_type := baseRestriction.GetRelationDefTypeRestrictionType()
	usersetRestriction := baseRestriction.GetRelationDefTypeRestrictionRelation()
	wildcardRestriction := baseRestriction.GetRelationDefTypeRestrictionWildcard()
	conditionName := ctx.ConditionName()

	relationRef := &openfgav1.RelationReference{}
	if _type != nil {
		relationRef.Type = _type.GetText()
	}

	if conditionName != nil {
		relationRef.Condition = conditionName.GetText()
	}

	if usersetRestriction != nil {
		relationRef.RelationOrWildcard = &openfgav1.RelationReference_Relation{
			Relation: usersetRestriction.GetText(),
		}
	}

	if wildcardRestriction != nil {
		relationRef.RelationOrWildcard = &openfgav1.RelationReference_Wildcard{Wildcard: &openfgav1.Wildcard{}}
	}

	l.currentRelation.TypeInfo.DirectlyRelatedUserTypes = append(
		l.currentRelation.TypeInfo.DirectlyRelatedUserTypes, relationRef,
	)
}

func (l *OpenFgaDslListener) ExitRelationDefRewrite(ctx *parser.RelationDefRewriteContext) {
	var partialRewrite *openfgav1.Userset

	computedUserset := &openfgav1.ObjectRelation{
		Relation: ctx.GetRewriteComputedusersetName().GetText(),
	}

	if ctx.GetRewriteTuplesetName() == nil {
		partialRewrite = &openfgav1.Userset{Userset: &openfgav1.Userset_ComputedUserset{
			ComputedUserset: computedUserset,
		}}
	} else {
		partialRewrite = &openfgav1.Userset{Userset: &openfgav1.Userset_TupleToUserset{
			TupleToUserset: &openfgav1.TupleToUserset{
				ComputedUserset: computedUserset,
				Tupleset: &openfgav1.ObjectRelation{
					Relation: ctx.GetRewriteTuplesetName().GetText(),
				},
			},
		}}
	}

	l.currentRelation.Rewrites = append(l.currentRelation.Rewrites, partialRewrite)
}

func (l *OpenFgaDslListener) ExitRelationRecurse(_ *parser.RelationRecurseContext) {
	if l.currentRelation == nil {
		return
	}

	relationDef := ParseExpression(l.currentRelation.Rewrites, l.currentRelation.Operator)

	if relationDef != nil {
		l.currentRelation.Rewrites = []*openfgav1.Userset{relationDef}
	}
}

func (l *OpenFgaDslListener) EnterRelationRecurseNoDirect(_ *parser.RelationRecurseNoDirectContext) {
	if l.rewriteStack != nil {
		l.rewriteStack = append(l.rewriteStack, &stackRelation{
			Rewrites: l.currentRelation.Rewrites,
			Operator: l.currentRelation.Operator,
		})
	}

	l.currentRelation.Rewrites = []*openfgav1.Userset{}
}

func (l *OpenFgaDslListener) ExitRelationRecurseNoDirect(_ *parser.RelationRecurseNoDirectContext) {
	if l.currentRelation == nil {
		return
	}

	relationDef := ParseExpression(l.currentRelation.Rewrites, l.currentRelation.Operator)

	popped, stack := l.rewriteStack[len(l.rewriteStack)-1], l.rewriteStack[:len(l.rewriteStack)-1]
	l.rewriteStack = stack

	if relationDef != nil {
		l.currentRelation.Operator = popped.Operator
		l.currentRelation.Rewrites = append(popped.Rewrites, relationDef) //nolint:gocritic
	}
}

func (l *OpenFgaDslListener) EnterRelationDefPartials(ctx *parser.RelationDefPartialsContext) {
	switch {
	case (len(ctx.AllOR()) > 0):
		l.currentRelation.Operator = RELATION_DEFINITION_OPERATOR_OR
	case (len(ctx.AllAND()) > 0):
		l.currentRelation.Operator = RELATION_DEFINITION_OPERATOR_AND
	case (ctx.BUT_NOT() != nil):
		l.currentRelation.Operator = RELATION_DEFINITION_OPERATOR_BUT_NOT
	}
}

//// Error Handling

type OpenFgaDslSyntaxErrorMetadata struct {
	symbol      string
	start, stop int
}

type OpenFgaDslSyntaxError struct {
	line, column int
	msg          string
	metadata     *OpenFgaDslSyntaxErrorMetadata
	e            antlr.RecognitionException //nolint:unused
}

func (err *OpenFgaDslSyntaxError) Error() string {
	return fmt.Sprintf("syntax error at line=%d, column=%d: %s", err.line, err.column, err.msg)
}

type OpenFgaDslSyntaxMultipleError multierror.Error

func (err *OpenFgaDslSyntaxMultipleError) Error() string {
	errors := err.Errors

	pluralS := ""
	if len(errors) > 1 {
		pluralS = "s"
	}

	errorsString := []string{}
	for _, item := range errors {
		errorsString = append(errorsString, item.Error())
	}

	return fmt.Sprintf("%d error%s occurred:\n\t* %s\n\n", len(errors), pluralS, strings.Join(errorsString, "\n\t* "))
}

type OpenFgaDslErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      *multierror.Error
}

func newOpenFgaDslErrorListener() *OpenFgaDslErrorListener {
	return new(OpenFgaDslErrorListener)
}

func (c *OpenFgaDslErrorListener) SyntaxError(
	_ antlr.Recognizer,
	offendingSymbol interface{},
	line, // line is one based, i.e. the first line will be 1
	column int, // column is zero based, i.e. the first column will be 0
	msg string,
	_ antlr.RecognitionException,
) {
	var metadata *OpenFgaDslSyntaxErrorMetadata

	if offendingSymbol != nil {
		symbol, ok := offendingSymbol.(*antlr.CommonToken)
		if ok {
			metadata = &OpenFgaDslSyntaxErrorMetadata{
				symbol: symbol.GetText(),
				start:  symbol.GetStart(),
				stop:   symbol.GetStop(),
			}
		}
	}

	c.Errors = multierror.Append(c.Errors, &OpenFgaDslSyntaxError{
		line:     line - 1,
		column:   column,
		msg:      msg,
		metadata: metadata,
	})
}

///

func ParseDSL(data string) (*OpenFgaDslListener, *OpenFgaDslErrorListener) {
	return ParseDSLWithOptions(data, true)
}

// ParseDSLWithOptions parses DSL with optional comment preservation.
func ParseDSLWithOptions(data string, preserveComments bool) (*OpenFgaDslListener, *OpenFgaDslErrorListener) {
	originalLines := strings.Split(data, "\n")
	cleanedLines := []string{}

	// Build a mapping from cleaned line numbers to original line numbers
	// (both 0-based indices)
	cleanedToOriginal := make(map[int]int)

	for originalIdx, line := range originalLines {
		cleanedLine := ""

		switch {
		case len(strings.TrimLeft(line, " ")) == 0:
			// do nothing, it's an empty line
		case strings.TrimLeft(line, " ")[0:1] == "#":
			cleanedLine = ""
		default:
			cleanedLine = strings.TrimRight(strings.Split(line, " #")[0], " ")
		}

		cleanedIdx := len(cleanedLines)
		cleanedLines = append(cleanedLines, cleanedLine)
		cleanedToOriginal[cleanedIdx] = originalIdx
	}

	// Create comment tracker with original data and line mapping
	var commentTracker *CommentTracker
	if preserveComments {
		commentTracker = NewCommentTrackerWithMapping(data, cleanedToOriginal)
	}

	cleanedData := strings.TrimRight(strings.Join(cleanedLines, "\n"), "\n")
	inputStream := antlr.NewInputStream(cleanedData)
	errorListener := newOpenFgaDslErrorListener()

	// Create the Lexer
	lexer := parser.NewOpenFGALexer(inputStream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	fgaParser := parser.NewOpenFGAParser(stream)
	fgaParser.RemoveErrorListeners()
	fgaParser.AddErrorListener(errorListener)

	listener := newOpenFgaDslListener()
	listener.commentTracker = commentTracker

	antlr.ParseTreeWalkerDefault.Walk(listener, fgaParser.Main())

	// Extract model comments after parsing
	if preserveComments && commentTracker != nil {
		listener.modelComments = commentTracker.GetModelComments()
	}

	return listener, errorListener
}

// TransformDSLToProto - Converts models authored in FGA DSL syntax to the OpenFGA Authorization Model Protobuf format.
func TransformDSLToProto(data string) (*openfgav1.AuthorizationModel, error) {
	listener, errorListener := ParseDSL(data)

	if errorListener.Errors != nil {
		return nil, errorListener.Errors
	}

	return &listener.authorizationModel, nil
}

// MustTransformDSLToProto - Calls TransformDSLToProto - panics if the error fails.
func MustTransformDSLToProto(data string) *openfgav1.AuthorizationModel {
	model, err := TransformDSLToProto(data)
	if err != nil {
		panic(err)
	}

	return model
}

// TransformDSLToJSON - Converts models authored in FGA DSL syntax to the json syntax accepted by the OpenFGA API.
func TransformDSLToJSON(data string) (string, error) {
	model, err := TransformDSLToProto(data)
	if err != nil {
		return "", err
	}

	bytes, err := protojson.Marshal(model)
	if err != nil {
		return "", fmt.Errorf("failed to marshal due to %w", err)
	}

	return string(bytes), nil
}

// MustTransformDSLToJSON - Calls TransformDSLToJSON - panics if the error fails.
func MustTransformDSLToJSON(data string) string {
	jsonString, err := TransformDSLToJSON(data)
	if err != nil {
		panic(err)
	}

	return jsonString
}

// TransformModularDSLToProto - Converts a part of a modular model in DSL syntax to the json syntax accepted by
// OpenFGA API and also returns the type definitions that are extended in the DSL if any are.
func TransformModularDSLToProto(data string) (*openfgav1.AuthorizationModel, map[string]*openfgav1.TypeDefinition, error) {
	listener, errorListener := ParseDSL(data)

	if errorListener.Errors != nil {
		return nil, nil, errorListener.Errors
	}

	return &listener.authorizationModel, listener.typeDefExtensions, nil
}

// TransformResult contains the result of parsing DSL with comments preserved.
type TransformResult struct {
	// AuthorizationModel is the parsed authorization model
	AuthorizationModel *openfgav1.AuthorizationModel
	// ModelComments contains comments at the model level
	ModelComments *ModelComments
	// TypeComments maps type names to their comments metadata
	TypeComments map[string]*CommentsMetadata
	// ConditionComments maps condition names to their comments
	ConditionComments map[string]*CommentInfo
}

// TransformDSLToProtoWithComments - Converts models authored in FGA DSL syntax to the OpenFGA Authorization Model
// Protobuf format while preserving comments metadata.
func TransformDSLToProtoWithComments(data string) (*TransformResult, error) {
	listener, errorListener := ParseDSL(data)

	if errorListener.Errors != nil {
		return nil, errorListener.Errors
	}

	return &TransformResult{
		AuthorizationModel: &listener.authorizationModel,
		ModelComments:      listener.modelComments,
		TypeComments:       listener.typeComments,
		ConditionComments:  listener.conditionComments,
	}, nil
}

// TransformDSLToJSONWithComments - Converts models authored in FGA DSL syntax to the JSON syntax with comments
// preserved in the metadata.
func TransformDSLToJSONWithComments(data string) (string, error) {
	result, err := TransformDSLToProtoWithComments(data)
	if err != nil {
		return "", err
	}

	// First, marshal the proto to JSON
	bytes, err := protojson.Marshal(result.AuthorizationModel)
	if err != nil {
		return "", fmt.Errorf("failed to marshal due to %w", err)
	}

	// Parse into a generic map to inject comments
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(bytes, &jsonMap); err != nil {
		return "", fmt.Errorf("failed to unmarshal for comment injection: %w", err)
	}

	// Inject model comments into top-level metadata
	if result.ModelComments != nil && len(result.ModelComments.PrecedingLines) > 0 {
		if jsonMap["metadata"] == nil {
			jsonMap["metadata"] = make(map[string]interface{})
		}
		metadata, ok := jsonMap["metadata"].(map[string]interface{})
		if ok {
			metadata["model_comments"] = map[string]interface{}{
				"preceding_lines": result.ModelComments.PrecedingLines,
			}
		}
	}

	// Inject type and relation comments
	if typeDefinitions, ok := jsonMap["type_definitions"].([]interface{}); ok {
		for _, td := range typeDefinitions {
			typeDef, ok := td.(map[string]interface{})
			if !ok {
				continue
			}

			typeName, ok := typeDef["type"].(string)
			if !ok {
				continue
			}

			typeComments := result.TypeComments[typeName]
			if typeComments == nil {
				continue
			}

			// Ensure metadata exists
			if typeDef["metadata"] == nil {
				typeDef["metadata"] = make(map[string]interface{})
			}
			metadata, ok := typeDef["metadata"].(map[string]interface{})
			if !ok {
				continue
			}

			// Add type-level comments
			if typeComments.Comments != nil {
				commentObj := make(map[string]interface{})
				if len(typeComments.Comments.PrecedingLines) > 0 {
					commentObj["preceding_lines"] = typeComments.Comments.PrecedingLines
				}
				if typeComments.Comments.Inline != "" {
					commentObj["inline"] = typeComments.Comments.Inline
				}
				if len(commentObj) > 0 {
					metadata["comments"] = commentObj
				}
			}

			// Add relation comments
			if len(typeComments.RelationComments) > 0 {
				if metadata["relations"] == nil {
					metadata["relations"] = make(map[string]interface{})
				}
				relationsMetadata, ok := metadata["relations"].(map[string]interface{})
				if ok {
					for relationName, relationComments := range typeComments.RelationComments {
						// Ensure relation metadata exists
						if relationsMetadata[relationName] == nil {
							relationsMetadata[relationName] = make(map[string]interface{})
						}
						relationMeta, ok := relationsMetadata[relationName].(map[string]interface{})
						if !ok {
							continue
						}

						commentObj := make(map[string]interface{})
						if len(relationComments.PrecedingLines) > 0 {
							commentObj["preceding_lines"] = relationComments.PrecedingLines
						}
						if relationComments.Inline != "" {
							commentObj["inline"] = relationComments.Inline
						}
						if len(commentObj) > 0 {
							relationMeta["comments"] = commentObj
						}
					}
				}
			}
		}
	}

	// Inject condition comments
	if conditions, ok := jsonMap["conditions"].(map[string]interface{}); ok {
		for conditionName, condComments := range result.ConditionComments {
			conditionData, ok := conditions[conditionName].(map[string]interface{})
			if !ok {
				continue
			}

			// Ensure metadata exists
			if conditionData["metadata"] == nil {
				conditionData["metadata"] = make(map[string]interface{})
			}
			metadata, ok := conditionData["metadata"].(map[string]interface{})
			if !ok {
				continue
			}

			commentObj := make(map[string]interface{})
			if len(condComments.PrecedingLines) > 0 {
				commentObj["preceding_lines"] = condComments.PrecedingLines
			}
			if condComments.Inline != "" {
				commentObj["inline"] = condComments.Inline
			}
			if len(commentObj) > 0 {
				metadata["comments"] = commentObj
			}
		}
	}

	// Marshal back to JSON
	finalBytes, err := json.Marshal(jsonMap)
	if err != nil {
		return "", fmt.Errorf("failed to marshal with comments: %w", err)
	}

	return string(finalBytes), nil
}
