package transformer

import (
	"fmt"
	"sort"
	"strings"

	pb "github.com/openfga/api/proto/openfga/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

func parseTypeRestriction(restriction *pb.RelationReference) string {
	typeName := restriction.GetType()
	relation := restriction.GetRelation()
	wildcard := restriction.GetWildcard()

	if wildcard != nil {
		return fmt.Sprintf("%v:*", typeName)
	}

	if relation != "" {
		return fmt.Sprintf("%v#%v", typeName, relation)
	}

	return fmt.Sprintf("%v", typeName)
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

func parseDifference(relationDefinition *pb.Userset, typeRestrictions []*pb.RelationReference) string {
	return fmt.Sprintf(
		"%v but not %v",
		parseSubRelation(relationDefinition.GetDifference().GetBase(), typeRestrictions),
		parseSubRelation(relationDefinition.GetDifference().GetSubtract(), typeRestrictions),
	)
}

func parseUnion(relationDefinition *pb.Userset, typeRestrictions []*pb.RelationReference) string {
	parsedString := []string{}
	children := relationDefinition.GetUnion().GetChild()

	for index := 0; index < len(children); index++ {
		parsedString = append(parsedString, parseSubRelation(children[index], typeRestrictions))
	}

	return fmt.Sprintf(
		"%v",
		strings.Join(parsedString, " or "),
	)
}

func parseIntersection(relationDefinition *pb.Userset, typeRestrictions []*pb.RelationReference) string {
	parsedString := []string{}
	children := relationDefinition.GetIntersection().GetChild()

	for index := 0; index < len(children); index++ {
		parsedString = append(parsedString, parseSubRelation(children[index], typeRestrictions))
	}

	return fmt.Sprintf(
		"%v",
		strings.Join(parsedString, " and "),
	)
}

func parseSubRelation(relationDefinition *pb.Userset, typeRestrictions []*pb.RelationReference) string {
	if relationDefinition.GetThis() != nil {
		return parseThis(typeRestrictions)
	}

	if relationDefinition.GetComputedUserset() != nil {
		return parseComputedUserset(relationDefinition)
	}

	if relationDefinition.GetTupleToUserset() != nil {
		return parseTupleToUserset(relationDefinition)
	}

	return ""
}

func parseRelation(
	relationName string,
	relationDefinition *pb.Userset,
	relationMetadata *pb.RelationMetadata,
) string {
	parsedRelationString := fmt.Sprintf(`    define %v: `, relationName)

	typeRestrictions := relationMetadata.GetDirectlyRelatedUserTypes()

	if relationDefinition.GetDifference() != nil {
		parsedRelationString += parseDifference(relationDefinition, typeRestrictions)
	} else if relationDefinition.GetUnion() != nil {
		parsedRelationString += parseUnion(relationDefinition, typeRestrictions)
	} else if relationDefinition.GetIntersection() != nil {
		parsedRelationString += parseIntersection(relationDefinition, typeRestrictions)
	} else {
		parsedRelationString += parseSubRelation(relationDefinition, typeRestrictions)
	}

	return parsedRelationString
}

func parseType(typeDefinition *pb.TypeDefinition) string {
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
			relation := relationsList[index]
			userset := relations[relation]
			meta := metadata.GetRelations()[relation]
			parsedRelationString := parseRelation(relation, userset, meta)
			parsedTypeString += fmt.Sprintf("\n%v", parsedRelationString)
		}
	}

	return parsedTypeString
}

// TransformJSONProtoToDSL - Converts models from the protobuf representation of the JSON syntax to the OpenFGA DSL
func TransformJSONProtoToDSL(model *pb.AuthorizationModel) string {
	schemaVersion := model.SchemaVersion

	typeDefinitions := []string{}
	typeDefs := model.GetTypeDefinitions()

	for index := 0; index < len(typeDefs); index++ {
		typeDef := typeDefs[index]
		typeDefinitions = append(typeDefinitions, fmt.Sprintf("\n%v", parseType(typeDef)))
	}

	typeDefsString := strings.Join(typeDefinitions, "\n")
	if len(typeDefinitions) > 0 {
		typeDefsString += "\n"
	}

	return fmt.Sprintf(`model
  schema %v
%v`, schemaVersion, typeDefsString)
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

	dsl := TransformJSONProtoToDSL(model)

	return &dsl, nil
}
