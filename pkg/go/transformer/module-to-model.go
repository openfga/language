package transformer

import (
	"fmt"
	"slices"
	"strings"

	"github.com/hashicorp/go-multierror"
	pb "github.com/openfga/api/proto/openfga/v1"
)

type ModuleFile struct {
	Name     string
	Contents string
}

type ModuleTransformationSingleMetadata struct{}

// ModuleTransformationSingleError is an error occurred during transformation of a module. Line and
// column number provided are one based.
type ModuleTransformationSingleError struct {
	Line, Column int
	Msg          string
}

func (e *ModuleTransformationSingleError) Error() string {
	return fmt.Sprintf("transformation error at line=%d, column=%d: %s", e.Line, e.Column, e.Msg)
}

type ModuleValidationMultipleError multierror.Error

func (e *ModuleValidationMultipleError) Error() string {
	errors := e.Errors

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

// TransformModuleFilesToModel transforms the provided modules into a singular authorization model.
func TransformModuleFilesToModel(modules []ModuleFile) (*pb.AuthorizationModel, error) {
	model := &pb.AuthorizationModel{
		SchemaVersion:   "1.2",
		TypeDefinitions: []*pb.TypeDefinition{},
		Conditions:      map[string]*pb.Condition{},
	}

	rawTypeDefs := []*pb.TypeDefinition{}
	types := []string{}
	extendedTypeDefs := map[string][]*pb.TypeDefinition{}
	conditions := map[string]*pb.Condition{}

	errors := &multierror.Error{}

	for _, module := range modules {
		mdl, typeDefExtensions, err := TransformModularDSLToProto(module.Contents)
		if err != nil {
			// add
			continue
		}

		for _, typeDef := range mdl.GetTypeDefinitions() {
			_, extension := typeDefExtensions[typeDef.GetType()]
			if slices.Contains(types, typeDef.GetType()) && !extension {
				errors = multierror.Append(errors, &ModuleTransformationSingleError{
					Msg: "duplicate type definition " + typeDef.GetType(),
				})
			}

			if extension {
				if extendedTypeDefs[module.Name] == nil {
					extendedTypeDefs[module.Name] = []*pb.TypeDefinition{}
				}

				extendedTypeDefs[module.Name] = append(extendedTypeDefs[module.Name], typeDef)

				continue
			}

			types = append(types, typeDef.GetType())
			typeDef.Metadata.File = module.Name
			rawTypeDefs = append(rawTypeDefs, typeDef)
		}

		for name, condition := range mdl.GetConditions() {
			if _, ok := conditions[name]; ok {
				errors = multierror.Append(errors, &ModuleTransformationSingleError{
					Msg: "duplicate condition " + name,
				})

				continue
			}

			condition.Metadata.File = module.Name
			conditions[name] = condition
		}
	}

	for filename, typeDefs := range extendedTypeDefs {
		for _, typeDef := range typeDefs {
			originalIndex := slices.IndexFunc(rawTypeDefs, func(t *pb.TypeDefinition) bool {
				return t.GetType() == typeDef.GetType()
			})
			original := rawTypeDefs[originalIndex]

			if original == nil {
				errors = multierror.Append(errors, &ModuleTransformationSingleError{
					Msg: fmt.Sprintf("extended type %s does not exist", typeDef.GetType()),
				})

				continue
			}

			if original.Relations == nil || len(original.GetRelations()) == 0 {
				original.Relations = typeDef.GetRelations()

				if original.GetMetadata() == nil {
					original.Metadata = &pb.Metadata{}
				}

				original.Metadata.Relations = typeDef.GetMetadata().GetRelations()

				if original.Metadata.Relations != nil {
					for name := range original.GetMetadata().GetRelations() {
						original.Metadata.Relations[name].File = filename
					}
				}

				rawTypeDefs[originalIndex] = original

				continue
			}

			existingRelationNames := []string{}
			for name := range original.GetRelations() {
				existingRelationNames = append(existingRelationNames, name)
			}

			for name, relation := range typeDef.GetRelations() {
				if slices.Contains(existingRelationNames, name) {
					errors = multierror.Append(&ModuleTransformationSingleError{
						Msg: fmt.Sprintf("relation %s already exists on type %s", name, typeDef.GetType()),
					})

					continue
				}

				var relationsMeta *pb.RelationMetadata

				for relationMetName, relationM := range typeDef.GetMetadata().GetRelations() {
					if relationMetName == name {
						relationsMeta = relationM

						break
					}
				}

				relationsMeta.File = filename
				original.Relations[name] = relation
				original.Metadata.Relations[name] = relationsMeta
			}
		}
	}

	model.TypeDefinitions = rawTypeDefs
	model.Conditions = conditions

	if len(errors.Errors) != 0 {
		return nil, &ModuleValidationMultipleError{
			Errors: errors.Errors,
		}
	}

	return model, nil
}
