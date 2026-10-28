package validation

import (
	"strings"
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/openfga/language/pkg/go/transformer"
)

// mustParse transforms a DSL that is expected to be grammatical.
func mustParse(t *testing.T, dsl string) *openfgav1.AuthorizationModel {
	t.Helper()

	model, err := transformer.TransformDSLToProto(dsl)
	require.NoError(t, err)

	return model
}

func TestValidateDSL(t *testing.T) {
	t.Parallel()

	t.Run("valid model returns nil", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: [user]`

		require.NoError(t, ValidateDSL(mustParse(t, dsl), dsl))
	})

	t.Run("nil model returns nil", func(t *testing.T) {
		t.Parallel()

		require.NoError(t, ValidateDSL(nil, ""))
		require.NoError(t, ValidateJSON(nil))
	})

	t.Run("findings are recovered with ExtractAllAs", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: editor`

		err := ValidateDSL(mustParse(t, dsl), dsl)
		require.Error(t, err)

		findings := ExtractAllAs[*Finding](err)
		require.Len(t, findings, 1)

		finding := findings[0]
		assert.Equal(t, MissingDefinition, finding.Metadata.Kind)
		assert.Equal(t, "editor", finding.Metadata.Symbol)
		assert.Equal(t, &Range{Start: 5, End: 5}, finding.Line)
		assert.Equal(t, &Range{Start: 19, End: 25}, finding.Column)
	})

	t.Run("an invalid condition name is reported at its declaration", func(t *testing.T) {
		t.Parallel()

		// The naming rule caps a condition name at 50 characters; the grammar
		// does not.
		name := strings.Repeat("c", 51)
		dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: [user with ` + name + `]
condition ` + name + `(x: int) {
  x > 0
}`

		err := ValidateDSL(mustParse(t, dsl), dsl)
		require.Error(t, err)

		findings := ExtractAllAs[*Finding](err)
		require.Len(t, findings, 1)

		finding := findings[0]
		assert.Equal(t, InvalidName, finding.Metadata.Kind)
		assert.Equal(t, &Range{Start: 6, End: 6}, finding.Line)
		assert.Equal(t, &Range{Start: 10, End: 61}, finding.Column)
	})
}

func TestValidateJSON(t *testing.T) {
	t.Parallel()

	t.Run("valid model returns nil", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: [user]`

		require.NoError(t, ValidateJSON(mustParse(t, dsl)))
	})

	t.Run("findings carry no position", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: editor`

		err := ValidateJSON(mustParse(t, dsl))
		require.Error(t, err)

		findings := ExtractAllAs[*Finding](err)
		require.Len(t, findings, 1)
		assert.Nil(t, findings[0].Line)
		assert.Nil(t, findings[0].Column)
	})
}

// TestValidateFileAndModule checks where file and module come from: the proto's
// source info, per declaration, not the validated text.
func TestValidateFileAndModule(t *testing.T) {
	t.Parallel()

	model := mustParse(t, "model\n  schema 1.1\ntype self")
	model.TypeDefinitions[0].Metadata = &openfgav1.Metadata{
		Module:     "core",
		SourceInfo: &openfgav1.SourceInfo{File: "core.fga"},
	}

	findings := ExtractAllAs[*Finding](validate(model, source{}))

	require.Len(t, findings, 1)
	assert.Equal(t, "core.fga", findings[0].File)
	assert.Equal(t, "core", findings[0].Metadata.Module)
}
