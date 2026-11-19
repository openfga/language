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

func ConditionParameterMissingGenericTypeError(conditionName string, parameterName string) error {
	return fmt.Errorf( //nolint:goerr113
		"the '%s' parameter of the '%s' condition has no generic type",
		parameterName,
		conditionName,
	)
}
