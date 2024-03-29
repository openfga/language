package transformer

import (
	"fmt"
	"sort"
	"strings"

	pb "github.com/openfga/api/proto/openfga/v1"
	"github.com/openfga/language/pkg/go/errors"
	"google.golang.org/protobuf/encoding/protojson"
)

type DirectAssignmentValidator struct {
	occurred int
}

func (v *DirectAssignmentValidator) incr() {
	v.occurred++
}

func (v *DirectAssignmentValidator) occurrences() int {
	return v.occurred
}

func (v *DirectAssignmentValidator) isFirstPosition(userset *pb.Userset) bool {
	if userset.GetThis() != nil {
		return true
	}

	if userset.GetDifference() != nil && userset.GetDifference().GetBase() != nil {
		if userset.GetDifference().GetBase().GetThis() != nil {
			return true
		} else {
			return v.isFirstPosition(userset.GetDifference().GetBase())
		}
	} else if userset.GetIntersection() != nil && userset.GetIntersection().GetChild() != nil {
		// For union and intersection, we are moving `this` to first position in the parse,
		// so even if it is not in the first position here, we're fine
		children := userset.GetIntersection().GetChild()
		if len(children) > 0 {
			for _, child := range children {
				if child.GetThis() != nil {
					return true
				}
			}

			return v.isFirstPosition(children[0])
		}
	} else if userset.GetUnion() != nil && userset.GetUnion().GetChild() != nil {
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

func parseTypeRestriction(restriction *pb.RelationReference) string {
	typeName := restriction.GetType()
	relation := restriction.GetRelation()
	wildcard := restriction.GetWildcard()
	condition := restriction.GetCondition()

	typeRestriction := fmt.Sprintf("%v", typeName)
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

func parseTypeRestrictions(restrictions []*pb.RelationReference) []string {
	parsedTypeRestrictions := []string{}
	for index := 0; index < len(restrictions); index++ {
		parsedTypeRestrictions = append(parsedTypeRestrictions, parseTypeRestriction(restrictions[index]))
	}

	return parsedTypeRestrictions
}

func parseThis(typeRestrictions []*pb.RelationReference) string {
	parsedTypeRestrictions := parseTypeRestrictions(typeRestrictions)

	return fmt.Sprintf("[%v]", strings.Join(parsedTypeRestrictions, ", "))
}

func parseTupleToUserset(relationDefinition *pb.Userset) string {
	return fmt.Sprintf(
		"%v from %v",
		relationDefinition.GetTupleToUserset().GetComputedUserset().GetRelation(),
		relationDefinition.GetTupleToUserset().GetTupleset().GetRelation(),
	)
}

func parseComputedUserset(relationDefinition *pb.Userset) string {
	return fmt.Sprintf("%v", relationDefinition.GetComputedUserset().GetRelation())
}

func parseDifference(typeName string, relationName string, relationDefinition *pb.Userset, typeRestrictions []*pb.RelationReference, validator *DirectAssignmentValidator) (string, error) {
	parsedSubStringBase, err := parseSubRelation(typeName, relationName, relationDefinition.GetDifference().GetBase(), typeRestrictions, validator)
	if err != nil {
		return "", err
	}

	parsedSubStringSubtract, err := parseSubRelation(typeName, relationName, relationDefinition.GetDifference().GetSubtract(), typeRestrictions, validator)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%v but not %v",
		parsedSubStringBase,
		parsedSubStringSubtract,
	), nil
}

func parseUnion(typeName string, relationName string, relationDefinition *pb.Userset, typeRestrictions []*pb.RelationReference, validator *DirectAssignmentValidator) (string, error) {
	parsedString := []string{}
	children := prioritizeDirectAssignment(relationDefinition.GetUnion().GetChild())

	for index := 0; index < len(children); index++ {
		parsedSubString, err := parseSubRelation(typeName, relationName, children[index], typeRestrictions, validator)
		if err != nil {
			return "", err
		}

		parsedString = append(parsedString, parsedSubString)
	}

	return fmt.Sprintf(
		"%v",
		strings.Join(parsedString, " or "),
	), nil
}

func parseIntersection(typeName string, relationName string, relationDefinition *pb.Userset, typeRestrictions []*pb.RelationReference, validator *DirectAssignmentValidator) (string, error) {
	parsedString := []string{}
	children := prioritizeDirectAssignment(relationDefinition.GetIntersection().GetChild())

	for index := 0; index < len(children); index++ {
		parsedSubString, err := parseSubRelation(typeName, relationName, children[index], typeRestrictions, validator)
		if err != nil {
			return "", err
		}

		parsedString = append(parsedString, parsedSubString)
	}

	return fmt.Sprintf(
		"%v",
		strings.Join(parsedString, " and "),
	), nil
}

func parseSubRelation(typeName string, relationName string, relationDefinition *pb.Userset, typeRestrictions []*pb.RelationReference, validator *DirectAssignmentValidator) (string, error) {
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
	relationDefinition *pb.Userset,
	relationMetadata *pb.RelationMetadata,
) (string, error) {
	validator := DirectAssignmentValidator{
		occurred: 0,
	}

	typeRestrictions := relationMetadata.GetDirectlyRelatedUserTypes()

	parseFn := parseSubRelation

	if relationDefinition.GetDifference() != nil {
		parseFn = parseDifference
	} else if relationDefinition.GetUnion() != nil {
		parseFn = parseUnion
	} else if relationDefinition.GetIntersection() != nil {
		parseFn = parseIntersection
	}

	parsedRelationString, err := parseFn(typeName, relationName, relationDefinition, typeRestrictions, &validator)
	if err != nil {
		return "", err
	}

	// Check if we have either no direct assignment, or we had exactly 1 direct assignment in the first position
	if validator.occurrences() == 0 || (validator.occurrences() == 1 && validator.isFirstPosition(relationDefinition)) {
		return fmt.Sprintf(`    define %v: %v`, relationName, parsedRelationString), nil
	}

	return "", errors.UnsupportedDSLNestingError(typeName, relationName)
}

func prioritizeDirectAssignment(usersets []*pb.Userset) []*pb.Userset {
	if len(usersets) > 0 {
		thisPosition := -1
		for index, userset := range usersets {
			if userset.GetThis() != nil {
				thisPosition = index
				break
			}
		}

		if thisPosition > 0 {
			newUsersets := []*pb.Userset{usersets[thisPosition]}
			newUsersets = append(newUsersets, usersets[:thisPosition]...)
			newUsersets = append(newUsersets, usersets[thisPosition+1:]...)

			return newUsersets
		}
	}

	return usersets
}

func parseType(typeDefinition *pb.TypeDefinition) (string, error) {
	typeName := typeDefinition.GetType()
	parsedTypeString := fmt.Sprintf(`type %v`, typeName)
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
		sort.Strings(relationsList)

		for index := 0; index < len(relationsList); index++ {
			relationName := relationsList[index]
			userset := relations[relationName]
			meta := metadata.GetRelations()[relationName]

			parsedRelationString, err := parseRelation(typeName, relationName, userset, meta)
			if err != nil {
				return "", err
			}

			parsedTypeString += fmt.Sprintf("\n%v", parsedRelationString)
		}
	}

	return parsedTypeString, nil
}

func parseConditionParams(parameterMap map[string]*pb.ConditionParamTypeRef) (string, error) {
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
		parameterTypeString := strings.ToLower(strings.ReplaceAll(parameterType.TypeName.String(), "TYPE_NAME_", ""))
		if parameterTypeString == "list" || parameterTypeString == "map" {
			genericTypeString := strings.ToLower(strings.ReplaceAll(parameterType.GetGenericTypes()[0].TypeName.String(), "TYPE_NAME_", ""))
			parameterTypeString = fmt.Sprintf("%s<%s>", parameterTypeString, genericTypeString)
		}

		parametersStringArray = append(parametersStringArray, fmt.Sprintf("%s: %s", parameterName, parameterTypeString))
	}

	return strings.Join(parametersStringArray, ", "), nil
}

func parseCondition(conditionName string, conditionDef *pb.Condition) (string, error) {
	if conditionName != conditionDef.GetName() {
		return "", errors.ConditionNameDoesntMatchError(conditionName, conditionDef.GetName())
	}

	paramsString, err := parseConditionParams(conditionDef.GetParameters())
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("condition %s(%s) {\n  %s\n}\n", conditionDef.Name, paramsString, conditionDef.GetExpression()), nil
}

func parseConditions(model *pb.AuthorizationModel) (string, error) {
	conditionsMap := model.GetConditions()
	if len(conditionsMap) == 0 {
		return "", nil
	}

	parsedConditionsString := ""

	conditionNames := []string{}
	for conditionName := range conditionsMap {
		conditionNames = append(conditionNames, conditionName)
	}

	// We are doing this in two loops (and sorting in between)
	// to make sure we have a deterministic behaviour that matches the API
	sort.Strings(conditionNames)

	for index := 0; index < len(conditionNames); index++ {
		conditionName := conditionNames[index]
		condition := conditionsMap[conditionName]

		parsedConditionString, err := parseCondition(conditionName, condition)
		if err != nil {
			return "", err
		}

		parsedConditionsString += fmt.Sprintf("\n%v", parsedConditionString)
	}

	return parsedConditionsString, nil
}

// TransformJSONProtoToDSL - Converts models from the protobuf representation of the JSON syntax to the OpenFGA DSL
func TransformJSONProtoToDSL(model *pb.AuthorizationModel) (string, error) {
	schemaVersion := model.SchemaVersion

	typeDefinitions := []string{}
	typeDefs := model.GetTypeDefinitions()

	for index := 0; index < len(typeDefs); index++ {
		typeDef := typeDefs[index]

		parsedType, err := parseType(typeDef)
		if err != nil {
			return "", err
		}

		typeDefinitions = append(typeDefinitions, fmt.Sprintf("\n%v", parsedType))
	}

	typeDefsString := strings.Join(typeDefinitions, "\n")
	if len(typeDefinitions) > 0 {
		typeDefsString += "\n"
	}

	parsedConditionsString, err := parseConditions(model)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`model
  schema %v
%v%v`, schemaVersion, typeDefsString, parsedConditionsString), nil
}

// LoadJSONStringToProto - Converts models authored in OpenFGA JSON syntax to the protobuf representation
func LoadJSONStringToProto(modelString string) (*pb.AuthorizationModel, error) {
	model := &pb.AuthorizationModel{}
	unmarshaller := protojson.UnmarshalOptions{
		AllowPartial:   false,
		DiscardUnknown: true,
	}

	if err := unmarshaller.Unmarshal([]byte(modelString), model); err != nil {
		return nil, err
	}

	return model, nil
}

// TransformJSONStringToDSL - Converts models authored in OpenFGA JSON syntax to the DSL syntax
func TransformJSONStringToDSL(modelString string) (*string, error) {
	model, err := LoadJSONStringToProto(modelString)
	if err != nil {
		return nil, err
	}

	dsl, err := TransformJSONProtoToDSL(model)
	if err != nil {
		return nil, err
	}

	return &dsl, nil
}
