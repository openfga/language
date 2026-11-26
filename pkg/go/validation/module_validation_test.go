package validation

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/openfga/language/pkg/go/transformer"
)

// moduleFixtureError is what a model validated from JSON can match of an error
// expected by a tests/data/transformer-module fixture. Its position is within
// a module file, which ValidateJSON has no text for.
type moduleFixtureError struct {
	Message  string `json:"msg"`
	File     string `json:"file"`
	Metadata struct {
		ErrorType Kind `json:"errorType"`
	} `json:"metadata"`
}

// TestValidateModuleFixtures validates the shared module fixtures and checks
// each finding's code, file and message against the fixture's expected errors.
// A fixture without any has to validate clean. One whose errors come from
// combining the modules is left out: pkg/js reports those before it has a
// model to validate.
func TestValidateModuleFixtures(t *testing.T) {
	t.Parallel()

	fixturesDir := filepath.Join(sharedDataDir, "transformer-module")

	fixtures, err := os.ReadDir(fixturesDir)
	require.NoError(t, err)

	for _, fixture := range fixtures {
		if !fixture.IsDir() {
			continue
		}

		dir := filepath.Join(fixturesDir, fixture.Name())

		// A fixture without modules exercises the other direction, from a model
		// to modules.
		if _, err := os.Stat(filepath.Join(dir, "module")); errors.Is(err, fs.ErrNotExist) {
			continue
		}

		var expected []moduleFixtureError

		data, err := os.ReadFile(filepath.Join(dir, "expected_errors.json"))
		if !errors.Is(err, fs.ErrNotExist) {
			require.NoError(t, err)
			require.NoError(t, json.Unmarshal(data, &expected))
		}

		if slices.ContainsFunc(expected, func(e moduleFixtureError) bool { return e.Metadata.ErrorType == "" }) {
			continue
		}

		t.Run(fixture.Name(), func(t *testing.T) {
			t.Parallel()

			if _, err := os.Stat(filepath.Join(dir, "test.skip")); err == nil {
				t.Skip("marked skip in the shared test data")
			}

			model, err := transformer.TransformModuleFilesToModel(readModuleFiles(t, filepath.Join(dir, "module")), "1.2")
			require.NoError(t, err)

			var got []moduleFixtureError
			for _, finding := range ExtractAllAs[*Finding](ValidateJSON(model)) {
				found := moduleFixtureError{Message: finding.Message, File: finding.File}
				found.Metadata.ErrorType = finding.Metadata.Kind
				got = append(got, found)
			}

			assert.ElementsMatch(t, expected, got)
		})
	}
}

func readModuleFiles(t *testing.T, dir string) []transformer.ModuleFile {
	t.Helper()

	entries, err := os.ReadDir(dir)
	require.NoError(t, err)

	var files []transformer.ModuleFile

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".fga") {
			continue
		}

		contents, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		require.NoError(t, err)

		files = append(files, transformer.ModuleFile{Name: entry.Name(), Contents: string(contents)})
	}

	return files
}

// TestValidateExtensionFileAndModule checks that a finding about a relation
// added by `extend type` points at the extension's file and module rather than
// at the file that declares the type, as pkg/js does.
func TestValidateExtensionFileAndModule(t *testing.T) {
	t.Parallel()

	const core = `module core

type user

type document
  relations
    define owner: [user]
`

	tests := []struct {
		name     string
		relation string
		kind     Kind
	}{
		{name: "undefined type", relation: "define viewer: [usr]", kind: InvalidType},
		{name: "missing computed relation", relation: "define viewer: missing", kind: MissingDefinition},
		{name: "missing tupleset relation", relation: "define viewer: owner from nope", kind: InvalidRelationType},
		{name: "reserved relation name", relation: "define this: [user]", kind: ReservedRelationKeywords},
		{name: "duplicate type restriction", relation: "define viewer: [user, user]", kind: DuplicatedError},
		{name: "undefined condition", relation: "define viewer: [user with nope]", kind: ConditionNotDefined},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			model, err := transformer.TransformModuleFilesToModel([]transformer.ModuleFile{
				{Name: "core.fga", Contents: core},
				{Name: "ext.fga", Contents: "module ext\n\nextend type document\n  relations\n    " + test.relation + "\n"},
			}, "1.2")
			require.NoError(t, err)

			findings := ExtractAllAs[*Finding](ValidateJSON(model))

			require.Len(t, findings, 1)
			assert.Equal(t, test.kind, findings[0].Metadata.Kind)
			assert.Equal(t, "ext.fga", findings[0].File)
			assert.Equal(t, "ext", findings[0].Metadata.Module)
		})
	}
}
