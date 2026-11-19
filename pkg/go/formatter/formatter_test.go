package formatter_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/openfga/language/pkg/go/formatter"
	"github.com/openfga/language/pkg/go/transformer"
)

type formatterTestCase struct {
	Name     string `yaml:"name"`
	DSL      string `yaml:"dsl"`
	Expected string `yaml:"expected"`
}

func loadFormatterTestCases(t *testing.T) []formatterTestCase {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("..", "..", "..", "tests", "data", "formatter-cases.yaml"))
	require.NoError(t, err)

	var testCases []formatterTestCase
	require.NoError(t, yaml.Unmarshal(data, &testCases))
	require.NotEmpty(t, testCases)

	return testCases
}

func TestFormat(t *testing.T) {
	t.Parallel()

	for _, testCase := range loadFormatterTestCases(t) {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			formatted, err := formatter.Format(testCase.DSL)
			require.NoError(t, err)
			require.Equal(t, testCase.Expected, formatted)

			again, err := formatter.Format(formatted)
			require.NoError(t, err)
			require.Equal(t, formatted, again)
		})
	}
}

// TestFormatKeepsCanonicalModels checks that the models produced by the JSON to DSL transformer
// are already formatted.
func TestFormatKeepsCanonicalModels(t *testing.T) {
	t.Parallel()

	files, err := filepath.Glob(filepath.Join("..", "..", "..", "tests", "data", "transformer", "*", "authorization-model.fga"))
	require.NoError(t, err)
	require.NotEmpty(t, files)

	for _, file := range files {
		t.Run(filepath.Base(filepath.Dir(file)), func(t *testing.T) {
			t.Parallel()

			data, err := os.ReadFile(file)
			require.NoError(t, err)

			dsl := strings.TrimRight(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n") + "\n"

			formatted, err := formatter.Format(dsl)
			require.NoError(t, err)
			require.Equal(t, dsl, formatted)
		})
	}
}

// TestFormatModules checks that formatting module files is stable and keeps their meaning.
func TestFormatModules(t *testing.T) {
	t.Parallel()

	files, err := filepath.Glob(filepath.Join("..", "..", "..", "tests", "data", "transformer-module", "*", "module", "*.fga"))
	require.NoError(t, err)
	require.NotEmpty(t, files)

	for _, file := range files {
		t.Run(filepath.Base(filepath.Dir(filepath.Dir(file)))+"/"+filepath.Base(file), func(t *testing.T) {
			t.Parallel()

			data, err := os.ReadFile(file)
			require.NoError(t, err)

			if _, _, err := transformer.TransformModularDSLToProto(string(data)); err != nil {
				t.Skip("fixture is not valid DSL")
			}

			formatted, err := formatter.Format(string(data))
			require.NoError(t, err)

			again, err := formatter.Format(formatted)
			require.NoError(t, err)
			require.Equal(t, formatted, again)
		})
	}
}

func TestFormatWithIndentSize(t *testing.T) {
	t.Parallel()

	formatted, err := formatter.Format("model\nschema 1.1\ntype user\ntype doc\nrelations\ndefine viewer: [user]\n",
		formatter.WithIndentSize(4))
	require.NoError(t, err)
	require.Equal(t, "model\n    schema 1.1\n\ntype user\n\ntype doc\n    relations\n        define viewer: [user]\n", formatted)
}

func TestFormatRejectsInvalidDSL(t *testing.T) {
	t.Parallel()

	_, err := formatter.Format("model\n  schema 1.1\n\ntype user\n  relations\n    define viewer [user]\n")
	require.Error(t, err)
	require.NotErrorIs(t, err, formatter.ErrModelChanged)
}

func TestMustFormatPanicsOnInvalidDSL(t *testing.T) {
	t.Parallel()

	require.Panics(t, func() {
		formatter.MustFormat("not a model")
	})
}

func TestFormatConditionBodyEdgeCases(t *testing.T) {
	t.Parallel()

	const header = "model\n  schema 1.1\n\ntype user\n\n"

	testCases := []struct {
		name     string
		dsl      string
		expected string
	}{
		{
			name:     "comment right after the closing brace",
			dsl:      header + "condition c(x: int) {\n  x > 0\n}// explanation\n",
			expected: header + "condition c(x: int) {\n  x > 0\n}// explanation\n",
		},
		{
			name:     "escaped quote in a string",
			dsl:      header + "condition c(x: string) {\n  x == \"\\\"}\"\n}\n",
			expected: header + "condition c(x: string) {\n  x == \"\\\"}\"\n}\n",
		},
		{
			name:     "brace in a hash comment",
			dsl:      header + "condition c(x: int) {\n  x > 0 # } explanation\n}\n",
			expected: header + "condition c(x: int) {\n  x > 0 # } explanation\n}\n",
		},
		{
			name:     "brace in a slash comment",
			dsl:      header + "condition c(x: int) {\n  x > 0 // } explanation\n}\n",
			expected: header + "condition c(x: int) {\n  x > 0 // } explanation\n}\n",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			formatted, err := formatter.Format(testCase.dsl)
			require.NoError(t, err)
			require.Equal(t, testCase.expected, formatted)
		})
	}
}

func TestFormatCRLFInMultilineLiteral(t *testing.T) {
	t.Parallel()

	const header = "model\r\n  schema 1.1\r\n\r\ntype user\r\n\r\n"

	// Line endings are not part of the model, so a file with CRLF line endings is formatted.
	formatted, err := formatter.Format(header + "condition c(x: int) {\r\n  x > 0\r\n}\r\n")
	require.NoError(t, err)
	require.Equal(t, strings.ReplaceAll(header, "\r\n", "\n")+"condition c(x: int) {\n  x > 0\n}\n", formatted)

	// They are part of a multiline string literal, which must not be changed silently.
	_, err = formatter.Format(header + "condition c(x: string) {\r\n  x == \"\"\"a\r\nb\"\"\"\r\n}\r\n")
	require.ErrorIs(t, err, formatter.ErrModelChanged)
}
