package validation

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	fgaerrors "github.com/openfga/language/pkg/go/errors"
)

// severityFixtureFile holds the Go-only expectations for severity, category and
// criticality. See the header of that file for why it is not in tests/data.
const severityFixtureFile = "testdata/severity-category-cases.yaml"

type severityFixtureScope struct {
	ObjectType string `yaml:"object_type"`
	Relation   string `yaml:"relation"`
	Condition  string `yaml:"condition"`
}

type severityFixtureExpectation struct {
	ErrorType string               `yaml:"error_type"`
	Severity  string               `yaml:"severity"`
	Category  string               `yaml:"category"`
	Critical  bool                 `yaml:"critical"`
	Sentinel  string               `yaml:"sentinel"`
	Scope     severityFixtureScope `yaml:"scope"`
}

type severityFixtureCase struct {
	Name     string                       `yaml:"name"`
	DSL      string                       `yaml:"dsl"`
	Expected []severityFixtureExpectation `yaml:"expected"`
}

// sentinelsByName resolves the sentinel a fixture names. YAML cannot reference a Go
// value, so fixtures name the sentinel as a string; an unknown name fails rather
// than being skipped.
var sentinelsByName = map[string]error{
	"ErrInvalidType":                fgaerrors.ErrInvalidType,
	"ErrDuplicateDefinition":        fgaerrors.ErrDuplicateDefinition,
	"ErrNoEntrypoints":              fgaerrors.ErrNoEntrypoints,
	"ErrConditionUnReferenced":      fgaerrors.ErrConditionUnReferenced,
	"ErrReservedKeywords":           fgaerrors.ErrReservedKeywords,
	"ErrObjectTypeUndefined":        fgaerrors.ErrObjectTypeUndefined,
	"ErrRelationUndefined":          fgaerrors.ErrRelationUndefined,
	"ErrInvalidRelationType":        fgaerrors.ErrInvalidRelationType,
	"ErrInvalidSchemaVersion":       fgaerrors.ErrInvalidSchemaVersion,
	"ErrMultipleModulesInFile":      fgaerrors.ErrMultipleModulesInFile,
	"ErrInvalidWildcard":            fgaerrors.ErrInvalidWildcard,
	"ErrConditionUndefined":         fgaerrors.ErrConditionUndefined,
	"ErrConditionNameMismatch":      fgaerrors.ErrConditionNameMismatch,
	"ErrInvalidName":                fgaerrors.ErrInvalidName,
	"ErrDirectlyAssignableRelation": fgaerrors.ErrDirectlyAssignableRelation,
}

func loadSeverityFixtures(t *testing.T) []severityFixtureCase {
	t.Helper()

	contents, err := os.ReadFile(severityFixtureFile)
	require.NoError(t, err, "reading %s", severityFixtureFile)

	var cases []severityFixtureCase
	require.NoError(t, yaml.Unmarshal(contents, &cases))
	require.NotEmpty(t, cases, "fixture file parsed to no cases")

	return cases
}

// TestSeverityFixtures runs the Go-only fixtures. Each expectation names a finding
// the validator must produce, with its severity, category, criticality, the sentinel
// errors.Is matches and the scope errors.As exposes.
func TestSeverityFixtures(t *testing.T) {
	t.Parallel()

	for _, fixture := range loadSeverityFixtures(t) {
		t.Run(fixture.Name, func(t *testing.T) {
			t.Parallel()

			validationErrors := validateDSL(t, fixture.DSL)
			require.NotNil(t, validationErrors)

			findings := validationErrors.AllFindings()

			for _, want := range fixture.Expected {
				sentinel, ok := sentinelsByName[want.Sentinel]
				require.Truef(t, ok,
					"fixture names sentinel %q, which is not in sentinelsByName", want.Sentinel)

				matched := findSeverityFixtureMatch(findings, want)
				require.NotNilf(t, matched,
					"no finding matched error_type=%q scope=%+v; got %s",
					want.ErrorType, want.Scope, validationErrors.Error())

				assert.Equal(t, want.Severity, matched.Severity.String())
				assert.Equal(t, want.Category, matched.Category.String())
				assert.Equal(t, want.Critical, isCriticalErrorType(matched.Metadata.ErrorType))
				assert.Equal(t, want.Severity == "error", matched.Blocks())

				require.ErrorIsf(t, error(matched), sentinel,
					"finding %q does not match %s via errors.Is", want.ErrorType, want.Sentinel)

				// The finding was selected on the scope errors.As reports, so
				// asserting the metadata here checks the two agree: the metadata
				// is derived from the cause and must not drift from it.
				require.NotNil(t, matched.Metadata)
				assert.Equal(t, want.Scope.ObjectType, matched.Metadata.Type)
				assert.Equal(t, want.Scope.Relation, matched.Metadata.Relation)
				assert.Equal(t, want.Scope.Condition, matched.Metadata.Condition)
			}
		})
	}
}

// findSeverityFixtureMatch locates the finding an expectation refers to. Matching
// on error type alone is not enough: the no-entrypoint case produces one finding
// per relation, so the scope is part of the identity.
func findSeverityFixtureMatch(
	findings []*ValidationError, want severityFixtureExpectation,
) *ValidationError {
	for _, finding := range findings {
		if finding.Metadata == nil || string(finding.Metadata.ErrorType) != want.ErrorType {
			continue
		}

		if finding.Unwrap() == nil {
			continue
		}

		objectType, relation, condition := causeScope(error(finding))
		if objectType == want.Scope.ObjectType &&
			relation == want.Scope.Relation &&
			condition == want.Scope.Condition {
			return finding
		}
	}

	return nil
}

// TestSeverityFixturesAreNotInTheSharedCorpus keeps these keys out of tests/data. An
// unknown key there breaks the Java suite on deserialisation and the JS suite on its
// toMatchObject assertions; see the fixture file's header.
func TestSeverityFixturesAreNotInTheSharedCorpus(t *testing.T) {
	t.Parallel()

	sharedCorpus := filepath.Join("..", "..", "..", "tests", "data", "dsl-semantic-validation-cases.yaml")

	contents, err := os.ReadFile(sharedCorpus)
	require.NoError(t, err, "reading the shared corpus")

	var cases []map[string]any
	require.NoError(t, yaml.Unmarshal(contents, &cases))

	// Keys pkg/js and pkg/java have no field for.
	goOnlyKeys := []string{"severity", "category", "critical", "sentinel", "scope"}

	for index, testCase := range cases {
		for _, key := range goOnlyKeys {
			_, present := testCase[key]
			assert.Falsef(t, present,
				"shared corpus case %d has Go-only key %q; it belongs in %s until "+
					"pkg/js and pkg/java can read it", index, key, severityFixtureFile)
		}

		expectedErrors, ok := testCase["expected_errors"].([]any)
		if !ok {
			continue
		}

		for _, raw := range expectedErrors {
			expectedError, ok := raw.(map[string]any)
			if !ok {
				continue
			}

			for _, key := range goOnlyKeys {
				_, present := expectedError[key]
				assert.Falsef(t, present,
					"shared corpus case %d has Go-only key %q inside expected_errors", index, key)
			}
		}
	}
}

// TestEveryFixtureSentinelIsReal stops sentinelsByName from drifting into a map
// of names that no longer exist, which would make the fixtures silently skip.
func TestEveryFixtureSentinelIsReal(t *testing.T) {
	t.Parallel()

	for name, sentinel := range sentinelsByName {
		require.Errorf(t, sentinel, "%s resolves to a nil error", name)
	}

	for _, fixture := range loadSeverityFixtures(t) {
		for _, want := range fixture.Expected {
			_, ok := sentinelsByName[want.Sentinel]
			assert.Truef(t, ok, "fixture %q names unknown sentinel %q", fixture.Name, want.Sentinel)
		}
	}
}
