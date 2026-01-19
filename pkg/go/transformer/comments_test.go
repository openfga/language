package transformer_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/openfga/language/pkg/go/transformer"
)

func TestCommentTracker(t *testing.T) {
	t.Parallel()

	t.Run("extracts preceding comments", func(t *testing.T) {
		t.Parallel()

		source := `# Comment 1
# Comment 2
model
  schema 1.1`

		tracker := transformer.NewCommentTracker(source)
		comments := tracker.GetPrecedingComments(2)

		require.Equal(t, []string{"# Comment 1", "# Comment 2"}, comments)
	})

	t.Run("extracts inline comments", func(t *testing.T) {
		t.Parallel()

		source := `model
  schema 1.1

type user # inline comment`

		tracker := transformer.NewCommentTracker(source)
		inline := tracker.GetInlineComment(3)

		require.Equal(t, "# inline comment", inline)
	})

	t.Run("extracts model comments", func(t *testing.T) {
		t.Parallel()

		source := `# Model comment
model
  schema 1.1`

		tracker := transformer.NewCommentTracker(source)
		modelComments := tracker.GetModelComments()

		require.NotNil(t, modelComments)
		require.Equal(t, []string{"# Model comment"}, modelComments.PrecedingLines)
	})

	t.Run("handles no comments", func(t *testing.T) {
		t.Parallel()

		source := `model
  schema 1.1

type user`

		tracker := transformer.NewCommentTracker(source)
		comments := tracker.GetPrecedingComments(3)
		inline := tracker.GetInlineComment(3)

		require.Empty(t, comments)
		require.Empty(t, inline)
	})
}

func TestDSLToJSONWithComments(t *testing.T) {
	t.Parallel()

	t.Run("preserves model comments", func(t *testing.T) {
		t.Parallel()

		dsl := `# OpenFGA Model
# Version 1.0
model
  schema 1.1

type user`

		jsonStr, err := transformer.TransformDSLToJSONWithComments(dsl)
		require.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &result)
		require.NoError(t, err)

		metadata, ok := result["metadata"].(map[string]interface{})
		require.True(t, ok, "metadata should exist")

		modelComments, ok := metadata["model_comments"].(map[string]interface{})
		require.True(t, ok, "model_comments should exist")

		precedingLines, ok := modelComments["preceding_lines"].([]interface{})
		require.True(t, ok, "preceding_lines should exist")
		require.Len(t, precedingLines, 2)
		require.Equal(t, "# OpenFGA Model", precedingLines[0])
		require.Equal(t, "# Version 1.0", precedingLines[1])
	})

	t.Run("preserves type comments", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1

# User type comment
type user`

		jsonStr, err := transformer.TransformDSLToJSONWithComments(dsl)
		require.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &result)
		require.NoError(t, err)

		typeDefinitions, ok := result["type_definitions"].([]interface{})
		require.True(t, ok)
		require.Len(t, typeDefinitions, 1)

		typeDef := typeDefinitions[0].(map[string]interface{})
		metadata, ok := typeDef["metadata"].(map[string]interface{})
		require.True(t, ok, "type metadata should exist")

		comments, ok := metadata["comments"].(map[string]interface{})
		require.True(t, ok, "comments should exist")

		precedingLines, ok := comments["preceding_lines"].([]interface{})
		require.True(t, ok)
		require.Len(t, precedingLines, 1)
		require.Equal(t, "# User type comment", precedingLines[0])
	})

	t.Run("preserves relation comments", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1

type document
  relations
    # Owner comment
    define owner: [user]`

		jsonStr, err := transformer.TransformDSLToJSONWithComments(dsl)
		require.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &result)
		require.NoError(t, err)

		typeDefinitions := result["type_definitions"].([]interface{})
		typeDef := typeDefinitions[0].(map[string]interface{})
		metadata := typeDef["metadata"].(map[string]interface{})
		relations := metadata["relations"].(map[string]interface{})
		ownerMeta := relations["owner"].(map[string]interface{})
		comments := ownerMeta["comments"].(map[string]interface{})
		precedingLines := comments["preceding_lines"].([]interface{})

		require.Len(t, precedingLines, 1)
		require.Equal(t, "# Owner comment", precedingLines[0])
	})

	t.Run("preserves condition comments", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1

type user

# IP-based access control
condition ip_check(ip: string) {
  ip == "127.0.0.1"
}`

		jsonStr, err := transformer.TransformDSLToJSONWithComments(dsl)
		require.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &result)
		require.NoError(t, err)

		conditions := result["conditions"].(map[string]interface{})
		ipCheck := conditions["ip_check"].(map[string]interface{})
		metadata := ipCheck["metadata"].(map[string]interface{})
		comments := metadata["comments"].(map[string]interface{})
		precedingLines := comments["preceding_lines"].([]interface{})

		require.Len(t, precedingLines, 1)
		require.Equal(t, "# IP-based access control", precedingLines[0])
	})
}

func TestJSONToDSLWithComments(t *testing.T) {
	t.Parallel()

	t.Run("emits model comments", func(t *testing.T) {
		t.Parallel()

		jsonStr := `{
			"schema_version": "1.1",
			"metadata": {
				"model_comments": {
					"preceding_lines": ["# Model comment"]
				}
			},
			"type_definitions": [{"type": "user"}]
		}`

		dsl, err := transformer.TransformJSONStringToDSLWithComments(jsonStr)
		require.NoError(t, err)
		require.NotNil(t, dsl)
		require.Contains(t, *dsl, "# Model comment\nmodel")
	})

	t.Run("emits type comments", func(t *testing.T) {
		t.Parallel()

		jsonStr := `{
			"schema_version": "1.1",
			"type_definitions": [{
				"type": "user",
				"metadata": {
					"comments": {
						"preceding_lines": ["# User type"]
					}
				}
			}]
		}`

		dsl, err := transformer.TransformJSONStringToDSLWithComments(jsonStr)
		require.NoError(t, err)
		require.NotNil(t, dsl)
		require.Contains(t, *dsl, "# User type\ntype user")
	})

	t.Run("emits relation comments", func(t *testing.T) {
		t.Parallel()

		jsonStr := `{
			"schema_version": "1.1",
			"type_definitions": [{
				"type": "document",
				"relations": {
					"owner": {"this": {}}
				},
				"metadata": {
					"relations": {
						"owner": {
							"directly_related_user_types": [{"type": "user"}],
							"comments": {
								"preceding_lines": ["# Owner relation"]
							}
						}
					}
				}
			}]
		}`

		dsl, err := transformer.TransformJSONStringToDSLWithComments(jsonStr)
		require.NoError(t, err)
		require.NotNil(t, dsl)
		require.Contains(t, *dsl, "# Owner relation\n    define owner:")
	})

	t.Run("emits condition comments", func(t *testing.T) {
		t.Parallel()

		jsonStr := `{
			"schema_version": "1.1",
			"type_definitions": [{"type": "user"}],
			"conditions": {
				"ip_check": {
					"name": "ip_check",
					"expression": "ip == \"127.0.0.1\"",
					"parameters": {
						"ip": {"type_name": "TYPE_NAME_STRING"}
					},
					"metadata": {
						"comments": {
							"preceding_lines": ["# IP condition"]
						}
					}
				}
			}
		}`

		dsl, err := transformer.TransformJSONStringToDSLWithComments(jsonStr)
		require.NoError(t, err)
		require.NotNil(t, dsl)
		require.Contains(t, *dsl, "# IP condition\ncondition ip_check")
	})
}

func TestRoundTripCommentPreservation(t *testing.T) {
	t.Parallel()

	t.Run("round trip preserves comments", func(t *testing.T) {
		t.Parallel()

		originalDSL := `# Model header comment
model
  schema 1.1

# User type
type user

# Document type
type document
  relations
    # Owner of document
    define owner: [user]
`

		// DSL -> JSON
		jsonStr, err := transformer.TransformDSLToJSONWithComments(originalDSL)
		require.NoError(t, err)

		// JSON -> DSL
		dsl, err := transformer.TransformJSONStringToDSLWithComments(jsonStr)
		require.NoError(t, err)
		require.NotNil(t, dsl)

		// Verify comments are preserved
		require.Contains(t, *dsl, "# Model header comment")
		require.Contains(t, *dsl, "# User type")
		require.Contains(t, *dsl, "# Document type")
		require.Contains(t, *dsl, "# Owner of document")
	})
}

func TestBackwardCompatibility(t *testing.T) {
	t.Parallel()

	t.Run("JSON without comments transforms correctly", func(t *testing.T) {
		t.Parallel()

		jsonStr := `{
			"schema_version": "1.1",
			"type_definitions": [{
				"type": "user"
			}]
		}`

		dsl, err := transformer.TransformJSONStringToDSLWithComments(jsonStr)
		require.NoError(t, err)
		require.NotNil(t, dsl)
		require.Contains(t, *dsl, "type user")
	})

	t.Run("DSL without comments transforms correctly", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1

type user`

		jsonStr, err := transformer.TransformDSLToJSONWithComments(dsl)
		require.NoError(t, err)
		require.NotEmpty(t, jsonStr)

		// Should not have model_comments if there were no comments
		var result map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &result)
		require.NoError(t, err)

		// Check metadata doesn't have model_comments if no comments
		if metadata, ok := result["metadata"].(map[string]interface{}); ok {
			_, hasModelComments := metadata["model_comments"]
			require.False(t, hasModelComments)
		}
	})
}
