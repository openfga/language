package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var schemaPath = filepath.Join("..", "..", "..", "..", "..", "schemas", "authorization-model.schema.json")

// TestSchemaIsUpToDate fails when the checked in schema doesn't match what the
// generator produces from the current API definitions.
func TestSchemaIsUpToDate(t *testing.T) {
	t.Parallel()

	want, err := authorizationModelSchema()
	require.NoError(t, err)

	got, err := os.ReadFile(schemaPath)
	require.NoError(t, err)

	if strings.ReplaceAll(string(got), "\r\n", "\n") != string(want) {
		t.Fatalf("schemas/authorization-model.schema.json is out of date, run `make generate-schemas`. Generated schema:\n%s", want)
	}
}

func generated(t *testing.T) map[string]any {
	t.Helper()

	data, err := authorizationModelSchema()
	require.NoError(t, err)

	var schema map[string]any
	require.NoError(t, json.Unmarshal(data, &schema))

	return schema
}

func definition(t *testing.T, schema map[string]any, name string) map[string]any {
	t.Helper()

	definitions, ok := schema["definitions"].(map[string]any)
	require.True(t, ok)

	def, ok := definitions[name].(map[string]any)
	require.True(t, ok, "missing definition %s", name)

	return def
}

func TestRootRequiresSchemaVersionButNotID(t *testing.T) {
	t.Parallel()

	schema := generated(t)
	require.Equal(t, []any{"schema_version"}, schema["required"])

	properties, ok := schema["properties"].(map[string]any)
	require.True(t, ok)
	require.Contains(t, properties, "type_definitions")
	// The API sets json_name to the snake_case name, so protojson rejects typeDefinitions.
	require.NotContains(t, properties, "typeDefinitions")
}

func TestUsersetSetsExactlyOneRewrite(t *testing.T) {
	t.Parallel()

	userset := definition(t, generated(t), "userset")
	require.InDelta(t, 1, userset["minProperties"], 0)
	require.InDelta(t, 1, userset["maxProperties"], 0)
	require.Equal(t, false, userset["additionalProperties"])
}

func TestTupleToUsersetRequiresTheComputedUsersetUnderEitherName(t *testing.T) {
	t.Parallel()

	ttu := definition(t, generated(t), "tupleToUserset")
	require.Equal(t, []any{"tupleset"}, ttu["required"])
	require.Contains(t, ttu["allOf"], map[string]any{"oneOf": []any{
		map[string]any{"required": []any{"computedUserset"}},
		map[string]any{"required": []any{"computed_userset"}},
	}})
}

func TestRelationReferenceAllowsOneOfRelationAndWildcard(t *testing.T) {
	t.Parallel()

	ref := definition(t, generated(t), "relationReference")
	require.Equal(t, []any{"type"}, ref["required"])
	require.Contains(t, ref["allOf"], map[string]any{"not": map[string]any{"anyOf": []any{
		map[string]any{"required": []any{"relation", "wildcard"}},
	}}})
}

func TestIgnoreEmptyPatternsAcceptAnEmptyString(t *testing.T) {
	t.Parallel()

	metadata := definition(t, generated(t), "metadata")
	properties, ok := metadata["properties"].(map[string]any)
	require.True(t, ok)

	module, ok := properties["module"].(map[string]any)
	require.True(t, ok)
	require.Equal(t, "string", module["type"])
	require.Len(t, module["anyOf"], 2)
}

func TestUsersetsRequireTheirChildren(t *testing.T) {
	t.Parallel()

	usersets := definition(t, generated(t), "usersets")
	require.Equal(t, []any{"child"}, usersets["required"])
}
