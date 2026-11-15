package errors

import (
	"errors"
	"fmt"
)

var ErrUnsupportedJSON = errors.New("json not currently supported by dsl")

func UnsupportedDSLNestingError(typeName string, relationName string) error {
	return fmt.Errorf( //nolint:goerr113
		"the '%s' relation definition under the '%s' type is not supported by the OpenFGA DSL syntax yet",
		relationName,
		typeName,
	)
}

func ConditionNameDoesntMatchError(conditionName string, conditionNestedName string) error {
	return fmt.Errorf( //nolint:goerr113
		"the '%s' condition has a different nested condition name ('%s')",
		conditionName,
		conditionNestedName,
	)
}

func UnsupportedModularModulesError(schemaVersion string) error {
	return fmt.Errorf( //nolint:goerr113
		"model schema version %s does not support modules",
		schemaVersion,
	)
}
