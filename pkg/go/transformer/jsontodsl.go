package transformer

import (
	"cmp"
	"encoding/json"
	"fmt"
	"slices"
	"sort"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"

	"github.com/openfga/language/pkg/go/errors"
)

// JSONModelWithComments represents the JSON structure with comments embedded.
type JSONModelWithComments struct {
	SchemaVersion   string                              `json:"schema_version"`
	Metadata        *JSONModelMetadata                  `json:"metadata,omitempty"`
	TypeDefinitions []JSONTypeDefinitionWithComments    `json:"type_definitions,omitempty"`
	Conditions      map[string]JSONConditionWithComments `json:"conditions,omitempty"`
}

// JSONModelMetadata represents top-level model metadata including model comments.
type JSONModelMetadata struct {
	ModelComments *JSONCommentBlock `json:"model_comments,omitempty"`
}

// JSONCommentBlock represents a block of comments.
type JSONCommentBlock struct {
	PrecedingLines []string `json:"preceding_lines,omitempty"`
	Inline         string   `json:"inline,omitempty"`
}

// JSONTypeDefinitionWithComments represents a type definition with comments.
type JSONTypeDefinitionWithComments struct {
	Type      string                                  `json:"type"`
	Relations map[string]json.RawMessage              `json:"relations,omitempty"`
	Metadata  *JSONTypeMetadataWithComments           `json:"metadata,omitempty"`
}

// JSONTypeMetadataWithComments represents type metadata with comments.
type JSONTypeMetadataWithComments struct {
	Module    string                                    `json:"module,omitempty"`
	Comments  *JSONCommentBlock                         `json:"comments,omitempty"`
	Relations map[string]*JSONRelationMetadataWithComments `json:"relations,omitempty"`
}

// JSONRelationMetadataWithComments represents relation metadata with comments.
type JSONRelationMetadataWithComments struct {
	DirectlyRelatedUserTypes []json.RawMessage `json:"directly_related_user_types,omitempty"`
	Module                   string            `json:"module,omitempty"`
	Comments                 *JSONCommentBlock `json:"comments,omitempty"`
}

// JSONConditionWithComments represents a condition with comments.
type JSONConditionWithComments struct {
	Name       string                           `json:"name"`
	Expression string                           `json:"expression"`
	Parameters map[string]json.RawMessage       `json:"parameters,omitempty"`
	Metadata   *JSONConditionMetadataWithComments `json:"metadata,omitempty"`
}

// JSONConditionMetadataWithComments represents condition metadata with comments.
type JSONConditionMetadataWithComments struct {
	Module   string            `json:"module,omitempty"`
	Comments *JSONCommentBlock `json:"comments,omitempty"`
}

type DirectAssignmentValidator struct {
	occurred int
}

func (v *DirectAssignmentValidator) incr() {
	v.occurred++
}

func (v *DirectAssignmentValidator) occurrences() int {
	return v.occurred
}

func (v *DirectAssignmentValidator) isFirstPosition(userset *openfgav1.Userset) bool { //nolint:cyclop
	if userset.GetThis() != nil {
		return true
	}

	switch {
	case userset.GetDifference() != nil && userset.GetDifference().GetBase() != nil:
		if userset.GetDifference().GetBase().GetThis() != nil {
			return true
		}

		return v.isFirstPosition(userset.GetDifference().GetBase())
	case userset.GetIntersection() != nil &&
		userset.GetIntersection().GetChild() != nil &&
		len(userset.GetIntersection().GetChild()) > 0:
		// For union and intersection, we are moving `this` to first position in the parse,
		// so even if it is not in the first position here, we're fine
		children := userset.GetIntersection().GetChild()
		for _, child := range children {
			if child.GetThis() != nil {
				return true
			}
		}

		return v.isFirstPosition(children[0])

	case userset.GetUnion() != nil && len(userset.GetUnion().GetChild()) > 0:
		children := userset.GetUnion().GetChild()
		if len(children) > 0 {
			for _, child := range children {
				if child.GetThis() != nil {
					return true
				}
			}

			return v.isFirstPosition(children[0])
		}
	}

	return false
}

func parseTypeRestriction(restriction *openfgav1.RelationReference) string {
	typeName := restriction.GetType()
	relation := restriction.GetRelation()
	wildcard := restriction.GetWildcard()
	condition := restriction.GetCondition()

	typeRestriction := typeName
	if wildcard != nil {
		typeRestriction += ":*"
	}

	if relation != "" {
		typeRestriction += fmt.Sprintf("#%v", relation)
	}

	if condition != "" {
		typeRestriction += fmt.Sprintf(" with %v", condition)
	}

	return typeRestriction
}

func parseTypeRestrictions(restrictions []*openfgav1.RelationReference) []string {
	parsedTypeRestrictions := []string{}
	for index := 0; index < len(restrictions); index++ {
		parsedTypeRestrictions = append(parsedTypeRestrictions, parseTypeRestriction(restrictions[index]))
	}

	return parsedTypeRestrictions
}

func parseThis(typeRestrictions []*openfgav1.RelationReference) string {
	parsedTypeRestrictions := parseTypeRestrictions(typeRestrictions)

	return fmt.Sprintf("[%v]", strings.Join(parsedTypeRestrictions, ", "))
}

func parseTupleToUserset(relationDefinition *openfgav1.Userset) string {
	return fmt.Sprintf(
		"%v from %v",
		relationDefinition.GetTupleToUserset().GetComputedUserset().GetRelation(),
		relationDefinition.GetTupleToUserset().GetTupleset().GetRelation(),
	)
}

func parseComputedUserset(relationDefinition *openfgav1.Userset) string {
	return relationDefinition.GetComputedUserset().GetRelation()
}

func parseDifference(
	typeName string,
	relationName string,
	relationDefinition *openfgav1.Userset,
	typeRestrictions []*openfgav1.RelationReference,
	validator *DirectAssignmentValidator,
) (string, error) {
	parsedSubStringBase, err := parseSubRelation(
		typeName,
		relationName,
		relationDefinition.GetDifference().GetBase(),
		typeRestrictions,
		validator,
	)
	if err != nil {
		return "", err
	}

	parsedSubStringSubtract, err := parseSubRelation(
		typeName,
		relationName,
		relationDefinition.GetDifference().GetSubtract(),
		typeRestrictions,
		validator,
	)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%v but not %v",
		parsedSubStringBase,
		parsedSubStringSubtract,
	), nil
}

func parseUnion(
	typeName string,
	relationName string,
	relationDefinition *openfgav1.Userset,
	typeRestrictions []*openfgav1.RelationReference,
	validator *DirectAssignmentValidator,
) (string, error) {
	parsedString := []string{}
	children := prioritizeDirectAssignment(relationDefinition.GetUnion().GetChild())

	for index := 0; index < len(children); index++ {
		parsedSubString, err := parseSubRelation(typeName, relationName, children[index], typeRestrictions, validator)
		if err != nil {
			return "", err
		}

		parsedString = append(parsedString, parsedSubString)
	}

	return strings.Join(parsedString, " or "), nil
}

func parseIntersection(
	typeName string,
	relationName string,
	relationDefinition *openfgav1.Userset,
	typeRestrictions []*openfgav1.RelationReference,
	validator *DirectAssignmentValidator,
) (string, error) {
	parsedString := []string{}
	children := prioritizeDirectAssignment(relationDefinition.GetIntersection().GetChild())

	for index := 0; index < len(children); index++ {
		parsedSubString, err := parseSubRelation(typeName, relationName, children[index], typeRestrictions, validator)
		if err != nil {
			return "", err
		}

		parsedString = append(parsedString, parsedSubString)
	}

	return strings.Join(parsedString, " and "), nil
}

func parseSubRelation(
	typeName string,
	relationName string,
	relationDefinition *openfgav1.Userset,
	typeRestrictions []*openfgav1.RelationReference,
	validator *DirectAssignmentValidator,
) (string, error) {
	if relationDefinition.GetThis() != nil {
		// Make sure we have no more than 1 reference for direct assignment in a given relation
		validator.incr()

		return parseThis(typeRestrictions), nil
	}

	if relationDefinition.GetComputedUserset() != nil {
		return parseComputedUserset(relationDefinition), nil
	}

	if relationDefinition.GetTupleToUserset() != nil {
		return parseTupleToUserset(relationDefinition), nil
	}

	if relationDefinition.GetUnion() != nil {
		parsedUnion, err := parseUnion(typeName, relationName, relationDefinition, typeRestrictions, validator)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("(%s)", parsedUnion), nil
	}

	if relationDefinition.GetIntersection() != nil {
		parsedIntersection, err := parseIntersection(typeName, relationName, relationDefinition, typeRestrictions, validator)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("(%s)", parsedIntersection), nil
	}

	if relationDefinition.GetDifference() != nil {
		parsedDiff, err := parseDifference(typeName, relationName, relationDefinition, typeRestrictions, validator)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("(%s)", parsedDiff), nil
	}

	return "", errors.UnsupportedDSLNestingError(typeName, relationName)
}

func parseRelation(
	typeName string,
	relationName string,
	relationDefinition *openfgav1.Userset,
	relationMetadata *openfgav1.RelationMetadata,
	includeSourceInformation bool,
) (string, error) {
	validator := DirectAssignmentValidator{
		occurred: 0,
	}

	sourceString := constructSourceComment(
		relationMetadata.GetModule(),
		relationMetadata.GetSourceInfo().GetFile(),
		" extended by:",
		includeSourceInformation,
	)

	typeRestrictions := relationMetadata.GetDirectlyRelatedUserTypes()

	parseFn := parseSubRelation

	switch {
	case relationDefinition.GetDifference() != nil:
		parseFn = parseDifference
	case relationDefinition.GetUnion() != nil:
		parseFn = parseUnion
	case relationDefinition.GetIntersection() != nil:
		parseFn = parseIntersection
	}

	parsedRelationString, err := parseFn(typeName, relationName, relationDefinition, typeRestrictions, &validator)
	if err != nil {
		return "", err
	}

	// Check if we have either no direct assignment, or we had exactly 1 direct assignment in the first position
	if validator.occurrences() == 0 || (validator.occurrences() == 1 && validator.isFirstPosition(relationDefinition)) {
		return fmt.Sprintf(`    define %v: %v%s`, relationName, parsedRelationString, sourceString), nil
	}

	return "", errors.UnsupportedDSLNestingError(typeName, relationName)
}

func prioritizeDirectAssignment(usersets []*openfgav1.Userset) []*openfgav1.Userset {
	if len(usersets) > 0 {
		thisPosition := -1

		for index, userset := range usersets {
			if userset.GetThis() != nil {
				thisPosition = index

				break
			}
		}

		if thisPosition > 0 {
			newUsersets := []*openfgav1.Userset{usersets[thisPosition]}
			newUsersets = append(newUsersets, usersets[:thisPosition]...)
			newUsersets = append(newUsersets, usersets[thisPosition+1:]...)

			return newUsersets
		}
	}

	return usersets
}

func parseType(typeDefinition *openfgav1.TypeDefinition, isModularModel, includeSourceInformation bool) (string, error) {
	typeName := typeDefinition.GetType()
	sourceString := constructSourceComment(
		typeDefinition.GetMetadata().GetModule(),
		typeDefinition.GetMetadata().GetSourceInfo().GetFile(),
		"",
		includeSourceInformation,
	)

	parsedTypeString := fmt.Sprintf(`type %v%s`, typeName, sourceString)
	relations := typeDefinition.GetRelations()
	metadata := typeDefinition.GetMetadata()

	if len(relations) > 0 {
		parsedTypeString += "\n  relations"

		relationsList := []string{}
		for relation := range relations {
			relationsList = append(relationsList, relation)
		}

		// We are doing this in two loops (and sorting in between)
		// to make sure we have a deterministic behaviour that matches the API
		if isModularModel {
			slices.SortStableFunc(relationsList, func(aName, bName string) int {
				aMeta := metadata.GetRelations()[aName]
				bMeta := metadata.GetRelations()[bName]

				return sortByModule(
					aName, bName,
					aMeta.GetModule(), bMeta.GetModule(),
					aMeta.GetSourceInfo().GetFile(), bMeta.GetSourceInfo().GetFile(),
				)
			})
		} else {
			sort.Strings(relationsList)
		}

		for index := 0; index < len(relationsList); index++ {
			relationName := relationsList[index]
			userset := relations[relationName]
			meta := metadata.GetRelations()[relationName]

			parsedRelationString, err := parseRelation(typeName, relationName, userset, meta, includeSourceInformation)
			if err != nil {
				return "", err
			}

			parsedTypeString += fmt.Sprintf("\n%v", parsedRelationString)
		}
	}

	return parsedTypeString, nil
}

func parseConditionParams(parameterMap map[string]*openfgav1.ConditionParamTypeRef) string {
	parametersStringArray := []string{}

	parameterNames := []string{}
	for parameterType := range parameterMap {
		parameterNames = append(parameterNames, parameterType)
	}

	// We are doing this in two loops (and sorting in between)
	// to make sure we have a deterministic behaviour that matches the API
	sort.Strings(parameterNames)

	for _, parameterName := range parameterNames {
		parameterType := parameterMap[parameterName]
		parameterTypeString := strings.ToLower(strings.ReplaceAll(parameterType.GetTypeName().String(), "TYPE_NAME_", ""))

		if parameterTypeString == "list" || parameterTypeString == "map" {
			genericTypeString := strings.ToLower(
				strings.ReplaceAll(
					parameterType.GetGenericTypes()[0].GetTypeName().String(), "TYPE_NAME_", ""),
			)
			parameterTypeString = fmt.Sprintf("%s<%s>", parameterTypeString, genericTypeString)
		}

		parametersStringArray = append(parametersStringArray, fmt.Sprintf("%s: %s", parameterName, parameterTypeString))
	}

	return strings.Join(parametersStringArray, ", ")
}

func parseCondition(conditionName string, conditionDef *openfgav1.Condition, includeSourceInformation bool) (string, error) {
	if conditionName != conditionDef.GetName() {
		return "", errors.ConditionNameDoesntMatchError(conditionName, conditionDef.GetName())
	}

	paramsString := parseConditionParams(conditionDef.GetParameters())
	sourceString := constructSourceComment(
		conditionDef.GetMetadata().GetModule(),
		conditionDef.GetMetadata().GetSourceInfo().GetFile(),
		"", includeSourceInformation,
	)

	return fmt.Sprintf(
		"condition %s(%s) {\n  %s\n}%s\n",
		conditionDef.GetName(),
		paramsString,
		conditionDef.GetExpression(),
		sourceString,
	), nil
}

func parseConditions(model *openfgav1.AuthorizationModel, includeSourceInformation bool) (string, error) {
	conditionsMap := model.GetConditions()
	if len(conditionsMap) == 0 {
		return "", nil
	}

	parsedConditionsString := ""

	conditionNames := []string{}
	for conditionName := range conditionsMap {
		conditionNames = append(conditionNames, conditionName)
	}

	slices.SortStableFunc(conditionNames, func(aName, bName string) int {
		aMeta := conditionsMap[aName].GetMetadata()
		bMeta := conditionsMap[bName].GetMetadata()

		return sortByModule(
			aName, bName,
			aMeta.GetModule(), bMeta.GetModule(),
			aMeta.GetSourceInfo().GetFile(), bMeta.GetSourceInfo().GetFile(),
		)
	})

	for index := 0; index < len(conditionNames); index++ {
		conditionName := conditionNames[index]
		condition := conditionsMap[conditionName]

		parsedConditionString, err := parseCondition(conditionName, condition, includeSourceInformation)
		if err != nil {
			return "", err
		}

		parsedConditionsString += fmt.Sprintf("\n%v", parsedConditionString)
	}

	return parsedConditionsString, nil
}

func constructSourceComment(module, file, leadingString string, includeSourceInformation bool) string {
	if (module == "" && file == "") || !includeSourceInformation {
		return ""
	}

	return fmt.Sprintf(" #%s module: %s, file: %s", leadingString, module, file)
}

type transformOptions struct {
	includeSourceInformation bool
}

type TransformOption func(t *transformOptions)

// WithIncludeSourceInformation - Configures whether to append file and module information to types,
// relations, and conditions.
func WithIncludeSourceInformation(includeSourceInformation bool) TransformOption {
	return func(t *transformOptions) {
		t.includeSourceInformation = includeSourceInformation
	}
}

// TransformJSONProtoToDSL - Converts models from the protobuf representation of the JSON syntax to the OpenFGA DSL.
func TransformJSONProtoToDSL(model *openfgav1.AuthorizationModel, opts ...TransformOption) (string, error) {
	schemaVersion := model.GetSchemaVersion()

	transformOpts := &transformOptions{
		includeSourceInformation: false,
	}

	for _, opt := range opts {
		opt(transformOpts)
	}

	typeDefinitions := []string{}
	typeDefs := model.GetTypeDefinitions()
	isModularModel := false

	for index := 0; index < len(typeDefs); index++ {
		typeDef := typeDefs[index]

		if typeDef.GetMetadata().GetModule() != "" {
			isModularModel = true

			break
		}
	}

	if isModularModel {
		slices.SortStableFunc(typeDefs, func(a, b *openfgav1.TypeDefinition) int {
			return sortByModule(
				a.GetType(), b.GetType(),
				a.GetMetadata().GetModule(), b.GetMetadata().GetModule(),
				a.GetMetadata().GetSourceInfo().GetFile(), b.GetMetadata().GetSourceInfo().GetFile(),
			)
		})
	}

	for index := 0; index < len(typeDefs); index++ {
		typeDef := typeDefs[index]

		parsedType, err := parseType(typeDef, isModularModel, transformOpts.includeSourceInformation)
		if err != nil {
			return "", err
		}

		typeDefinitions = append(typeDefinitions, fmt.Sprintf("\n%v", parsedType))
	}

	typeDefsString := strings.Join(typeDefinitions, "\n")
	if len(typeDefinitions) > 0 {
		typeDefsString += "\n"
	}

	parsedConditionsString, err := parseConditions(model, transformOpts.includeSourceInformation)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`model
  schema %v
%v%v`, schemaVersion, typeDefsString, parsedConditionsString), nil
}

// LoadJSONStringToProto - Converts models authored in OpenFGA JSON syntax to the protobuf representation.
func LoadJSONStringToProto(modelString string) (*openfgav1.AuthorizationModel, error) {
	model := &openfgav1.AuthorizationModel{}
	unmarshaller := protojson.UnmarshalOptions{
		AllowPartial:   false,
		DiscardUnknown: true,
	}

	if err := unmarshaller.Unmarshal([]byte(modelString), model); err != nil {
		return nil, err //nolint:wrapcheck
	}

	return model, nil
}

// TransformJSONStringToDSL - Converts models authored in OpenFGA JSON syntax to the DSL syntax.
func TransformJSONStringToDSL(modelString string, opts ...TransformOption) (*string, error) {
	model, err := LoadJSONStringToProto(modelString)
	if err != nil {
		return nil, err
	}

	dsl, err := TransformJSONProtoToDSL(model, opts...)
	if err != nil {
		return nil, err
	}

	return &dsl, nil
}

func sortByModule(aName, bName, aModule, bModule, aFile, bFile string) int {
	if aModule == "" && bModule == "" {
		return cmp.Compare(aName, bName)
	}

	if aModule == "" {
		return -1
	}

	if bModule == "" {
		return 1
	}

	if aModule != bModule {
		return cmp.Compare(aModule, bModule)
	} else if aFile != bFile {
		return cmp.Compare(aFile, bFile)
	}

	return cmp.Compare(aName, bName)
}

// formatCommentLines formats comment lines for DSL output.
func formatCommentLines(comments []string) string {
	if len(comments) == 0 {
		return ""
	}
	result := ""
	for _, comment := range comments {
		result += comment + "\n"
	}
	return result
}

// formatInlineComment formats an inline comment for DSL output.
func formatInlineComment(comment string) string {
	if comment == "" {
		return ""
	}
	// Ensure the comment starts with # if it doesn't already
	if !strings.HasPrefix(comment, "#") {
		return " #" + comment
	}
	return " " + comment
}

// TransformJSONStringToDSLWithComments - Converts models authored in OpenFGA JSON syntax to the DSL syntax,
// preserving any comments stored in the JSON metadata.
func TransformJSONStringToDSLWithComments(modelString string) (*string, error) {
	// First try to parse as JSON with comments to extract comment metadata
	var modelWithComments JSONModelWithComments
	if err := json.Unmarshal([]byte(modelString), &modelWithComments); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Also load via proto for the actual model transformation
	model, err := LoadJSONStringToProto(modelString)
	if err != nil {
		return nil, err
	}

	dsl, err := transformJSONProtoToDSLWithComments(model, &modelWithComments)
	if err != nil {
		return nil, err
	}

	return &dsl, nil
}

// transformJSONProtoToDSLWithComments converts the proto model to DSL while injecting comments.
func transformJSONProtoToDSLWithComments(model *openfgav1.AuthorizationModel, commentsModel *JSONModelWithComments) (string, error) {
	schemaVersion := model.GetSchemaVersion()

	// Build model comments prefix
	modelCommentsStr := ""
	if commentsModel.Metadata != nil && commentsModel.Metadata.ModelComments != nil {
		modelCommentsStr = formatCommentLines(commentsModel.Metadata.ModelComments.PrecedingLines)
	}

	// Build type comments map
	typeCommentsMap := make(map[string]*JSONTypeMetadataWithComments)
	for _, typeDef := range commentsModel.TypeDefinitions {
		if typeDef.Metadata != nil {
			typeCommentsMap[typeDef.Type] = typeDef.Metadata
		}
	}

	typeDefinitions := []string{}
	typeDefs := model.GetTypeDefinitions()
	isModularModel := false

	for index := 0; index < len(typeDefs); index++ {
		typeDef := typeDefs[index]

		if typeDef.GetMetadata().GetModule() != "" {
			isModularModel = true

			break
		}
	}

	if isModularModel {
		slices.SortStableFunc(typeDefs, func(a, b *openfgav1.TypeDefinition) int {
			return sortByModule(
				a.GetType(), b.GetType(),
				a.GetMetadata().GetModule(), b.GetMetadata().GetModule(),
				a.GetMetadata().GetSourceInfo().GetFile(), b.GetMetadata().GetSourceInfo().GetFile(),
			)
		})
	}

	for index := 0; index < len(typeDefs); index++ {
		typeDef := typeDefs[index]

		typeComments := typeCommentsMap[typeDef.GetType()]
		parsedType, err := parseTypeWithComments(typeDef, isModularModel, false, typeComments)
		if err != nil {
			return "", err
		}

		typeDefinitions = append(typeDefinitions, fmt.Sprintf("\n%v", parsedType))
	}

	typeDefsString := strings.Join(typeDefinitions, "\n")
	if len(typeDefinitions) > 0 {
		typeDefsString += "\n"
	}

	parsedConditionsString, err := parseConditionsWithComments(model, false, commentsModel.Conditions)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`%vmodel
  schema %v
%v%v`, modelCommentsStr, schemaVersion, typeDefsString, parsedConditionsString), nil
}

// parseTypeWithComments parses a type definition including comments.
func parseTypeWithComments(typeDefinition *openfgav1.TypeDefinition, isModularModel, includeSourceInformation bool, typeComments *JSONTypeMetadataWithComments) (string, error) {
	typeName := typeDefinition.GetType()

	// Build type comment prefix
	typeCommentsStr := ""
	typeInlineComment := ""
	if typeComments != nil && typeComments.Comments != nil {
		typeCommentsStr = formatCommentLines(typeComments.Comments.PrecedingLines)
		typeInlineComment = formatInlineComment(typeComments.Comments.Inline)
	}

	sourceString := constructSourceComment(
		typeDefinition.GetMetadata().GetModule(),
		typeDefinition.GetMetadata().GetSourceInfo().GetFile(),
		"",
		includeSourceInformation,
	)

	parsedTypeString := fmt.Sprintf(`%vtype %v%s%s`, typeCommentsStr, typeName, typeInlineComment, sourceString)
	relations := typeDefinition.GetRelations()
	metadata := typeDefinition.GetMetadata()

	if len(relations) > 0 {
		parsedTypeString += "\n  relations"

		relationsList := []string{}
		for relation := range relations {
			relationsList = append(relationsList, relation)
		}

		// We are doing this in two loops (and sorting in between)
		// to make sure we have a deterministic behaviour that matches the API
		if isModularModel {
			slices.SortStableFunc(relationsList, func(aName, bName string) int {
				aMeta := metadata.GetRelations()[aName]
				bMeta := metadata.GetRelations()[bName]

				return sortByModule(
					aName, bName,
					aMeta.GetModule(), bMeta.GetModule(),
					aMeta.GetSourceInfo().GetFile(), bMeta.GetSourceInfo().GetFile(),
				)
			})
		} else {
			sort.Strings(relationsList)
		}

		for index := 0; index < len(relationsList); index++ {
			relationName := relationsList[index]
			userset := relations[relationName]
			meta := metadata.GetRelations()[relationName]

			// Get relation comments
			var relationComments *JSONCommentBlock
			if typeComments != nil && typeComments.Relations != nil {
				if relMeta := typeComments.Relations[relationName]; relMeta != nil {
					relationComments = relMeta.Comments
				}
			}

			parsedRelationString, err := parseRelationWithComments(typeName, relationName, userset, meta, includeSourceInformation, relationComments)
			if err != nil {
				return "", err
			}

			parsedTypeString += fmt.Sprintf("\n%v", parsedRelationString)
		}
	}

	return parsedTypeString, nil
}

// parseRelationWithComments parses a relation including comments.
func parseRelationWithComments(
	typeName string,
	relationName string,
	relationDefinition *openfgav1.Userset,
	relationMetadata *openfgav1.RelationMetadata,
	includeSourceInformation bool,
	comments *JSONCommentBlock,
) (string, error) {
	validator := DirectAssignmentValidator{
		occurred: 0,
	}

	// Build relation comment prefix
	relationCommentsStr := ""
	relationInlineComment := ""
	if comments != nil {
		for _, line := range comments.PrecedingLines {
			relationCommentsStr += "    " + line + "\n"
		}
		relationInlineComment = formatInlineComment(comments.Inline)
	}

	sourceString := constructSourceComment(
		relationMetadata.GetModule(),
		relationMetadata.GetSourceInfo().GetFile(),
		" extended by:",
		includeSourceInformation,
	)

	typeRestrictions := relationMetadata.GetDirectlyRelatedUserTypes()

	parseFn := parseSubRelation

	switch {
	case relationDefinition.GetDifference() != nil:
		parseFn = parseDifference
	case relationDefinition.GetUnion() != nil:
		parseFn = parseUnion
	case relationDefinition.GetIntersection() != nil:
		parseFn = parseIntersection
	}

	parsedRelationString, err := parseFn(typeName, relationName, relationDefinition, typeRestrictions, &validator)
	if err != nil {
		return "", err
	}

	// Check if we have either no direct assignment, or we had exactly 1 direct assignment in the first position
	if validator.occurrences() == 0 || (validator.occurrences() == 1 && validator.isFirstPosition(relationDefinition)) {
		return fmt.Sprintf(`%s    define %v: %v%s%s`, relationCommentsStr, relationName, parsedRelationString, relationInlineComment, sourceString), nil
	}

	return "", errors.UnsupportedDSLNestingError(typeName, relationName)
}

// parseConditionsWithComments parses conditions with comments.
func parseConditionsWithComments(model *openfgav1.AuthorizationModel, includeSourceInformation bool, conditionsWithComments map[string]JSONConditionWithComments) (string, error) {
	conditionsMap := model.GetConditions()
	if len(conditionsMap) == 0 {
		return "", nil
	}

	parsedConditionsString := ""

	conditionNames := []string{}
	for conditionName := range conditionsMap {
		conditionNames = append(conditionNames, conditionName)
	}

	slices.SortStableFunc(conditionNames, func(aName, bName string) int {
		aMeta := conditionsMap[aName].GetMetadata()
		bMeta := conditionsMap[bName].GetMetadata()

		return sortByModule(
			aName, bName,
			aMeta.GetModule(), bMeta.GetModule(),
			aMeta.GetSourceInfo().GetFile(), bMeta.GetSourceInfo().GetFile(),
		)
	})

	for index := 0; index < len(conditionNames); index++ {
		conditionName := conditionNames[index]
		condition := conditionsMap[conditionName]

		// Get condition comments
		var comments *JSONCommentBlock
		if condWithComments, ok := conditionsWithComments[conditionName]; ok && condWithComments.Metadata != nil {
			comments = condWithComments.Metadata.Comments
		}

		parsedConditionString, err := parseConditionWithComments(conditionName, condition, includeSourceInformation, comments)
		if err != nil {
			return "", err
		}

		parsedConditionsString += fmt.Sprintf("\n%v", parsedConditionString)
	}

	return parsedConditionsString, nil
}

// parseConditionWithComments parses a condition with comments.
func parseConditionWithComments(conditionName string, conditionDef *openfgav1.Condition, includeSourceInformation bool, comments *JSONCommentBlock) (string, error) {
	if conditionName != conditionDef.GetName() {
		return "", errors.ConditionNameDoesntMatchError(conditionName, conditionDef.GetName())
	}

	// Build condition comment prefix
	conditionCommentsStr := ""
	if comments != nil {
		conditionCommentsStr = formatCommentLines(comments.PrecedingLines)
	}

	paramsString := parseConditionParams(conditionDef.GetParameters())
	sourceString := constructSourceComment(
		conditionDef.GetMetadata().GetModule(),
		conditionDef.GetMetadata().GetSourceInfo().GetFile(),
		"", includeSourceInformation,
	)

	return fmt.Sprintf(
		"%scondition %s(%s) {\n  %s\n}%s\n",
		conditionCommentsStr,
		conditionDef.GetName(),
		paramsString,
		conditionDef.GetExpression(),
		sourceString,
	), nil
}
