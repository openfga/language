package validation

import (
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/openfga/language/pkg/go/transformer"
)

// modelWithRelations builds the smallest model the phases accept: one type
// whose relations are all direct assignments. It bypasses the parser, which
// would reject the invalid names these tests feed in.
func modelWithRelations(t *testing.T, typeName string, relationNames ...string) *openfgav1.AuthorizationModel {
	t.Helper()

	relations := make(map[string]*openfgav1.Userset, len(relationNames))
	for _, relationName := range relationNames {
		relations[relationName] = &openfgav1.Userset{
			Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}},
		}
	}

	return &openfgav1.AuthorizationModel{
		SchemaVersion: "1.1",
		TypeDefinitions: []*openfgav1.TypeDefinition{
			{Type: typeName, Relations: relations},
		},
	}
}

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

	t.Run("findings are recovered with errors.As", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: editor`

		err := ValidateDSL(mustParse(t, dsl), dsl)
		require.Error(t, err)

		var findings Findings
		require.ErrorAs(t, err, &findings)
		require.Len(t, findings, 1)

		finding := findings[0]
		assert.Equal(t, "the relation `editor` does not exist.", finding.Message)
		assert.Equal(t, MissingDefinition, finding.Metadata.Kind)
		assert.Equal(t, "editor", finding.Metadata.Symbol)
		assert.Equal(t, &Range{Start: 5, End: 5}, finding.Line)
		assert.Equal(t, &Range{Start: 19, End: 25}, finding.Column)
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

		var findings Findings
		require.ErrorAs(t, err, &findings)
		require.Len(t, findings, 1)
		assert.Equal(t, "the relation `editor` does not exist.", findings[0].Message)
		assert.Nil(t, findings[0].Line)
		assert.Nil(t, findings[0].Column)
	})
}

// TestValidateCascadeGate pins the phase gating: once an earlier phase reports,
// the structural phases (duplicates, entry points, operations, wildcards) are
// skipped, so one root cause is not reported as a cascade of derived findings.
func TestValidateCascadeGate(t *testing.T) {
	t.Parallel()

	t.Run("a duplicate suppresses derived entry-point findings", func(t *testing.T) {
		t.Parallel()

		// document is declared twice, and its viewer computes itself — a
		// no-entry-point loop that must NOT be reported alongside the duplicate.
		typeDef := func() *openfgav1.TypeDefinition {
			return &openfgav1.TypeDefinition{
				Type: "document",
				Relations: map[string]*openfgav1.Userset{
					"viewer": {Userset: &openfgav1.Userset_ComputedUserset{
						ComputedUserset: &openfgav1.ObjectRelation{Relation: "viewer"},
					}},
				},
			}
		}
		model := &openfgav1.AuthorizationModel{
			SchemaVersion:   "1.1",
			TypeDefinitions: []*openfgav1.TypeDefinition{typeDef(), typeDef()},
		}

		findings := validate(model, source{})

		require.Len(t, findings, 1)
		assert.Equal(t, DuplicatedError, findings[0].Metadata.Kind)
	})

	t.Run("an entry-point loop is reported when nothing gates it", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: editor
    define editor: viewer`

		var findings Findings
		require.ErrorAs(t, ValidateDSL(mustParse(t, dsl), dsl), &findings)
		require.Len(t, findings, 2)

		for _, finding := range findings {
			assert.Equal(t, RelationNoEntrypoint, finding.Metadata.Kind)
			assert.Contains(t, finding.Message, "(potential loop)")
		}
	})

	t.Run("condition checks run even when the cascade is gated", func(t *testing.T) {
		t.Parallel()

		// The undefined type gates the structural phases, but the unused
		// condition must still be reported, matching the reference.
		dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: [ghost]
condition marked(x: int) {
  x > 0
}`

		var findings Findings
		require.ErrorAs(t, ValidateDSL(mustParse(t, dsl), dsl), &findings)
		require.Len(t, findings, 2)

		assert.Equal(t, InvalidType, findings[0].Metadata.Kind)
		assert.Equal(t, ConditionNotUsed, findings[1].Metadata.Kind)
	})
}

// TestValidateFileAndModule pins where file and module come from: the proto's
// source info, per declaration, not the validated text.
func TestValidateFileAndModule(t *testing.T) {
	t.Parallel()

	model := modelWithRelations(t, "self")
	model.TypeDefinitions[0].Metadata = &openfgav1.Metadata{
		Module:     "core",
		SourceInfo: &openfgav1.SourceInfo{File: "core.fga"},
	}

	findings := validate(model, source{})

	require.Len(t, findings, 1)
	assert.Equal(t, "core.fga", findings[0].File)
	assert.Equal(t, "core", findings[0].Metadata.Module)
}

// TestValidateMultiFile pins the one multi-file rule: a file that would carry
// two modules is reported, with the modules listed in model order.
func TestValidateMultiFile(t *testing.T) {
	t.Parallel()

	model := &openfgav1.AuthorizationModel{
		SchemaVersion: "1.1",
		TypeDefinitions: []*openfgav1.TypeDefinition{
			{Type: "user", Metadata: &openfgav1.Metadata{
				Module:     "core",
				SourceInfo: &openfgav1.SourceInfo{File: "shared.fga"},
			}},
			{Type: "document", Metadata: &openfgav1.Metadata{
				Module:     "docs",
				SourceInfo: &openfgav1.SourceInfo{File: "shared.fga"},
			}},
		},
	}

	findings := validateMultiFile(model)

	require.Len(t, findings, 1)
	assert.Equal(t, MultipleModulesInFile, findings[0].Metadata.Kind)
	assert.Equal(t,
		"file shared.fga would contain multiple module definitions (core, docs) when transforming to DSL. "+
			"Only one module can be defined per file.",
		findings[0].Message)
	assert.Nil(t, findings[0].Line)
}
