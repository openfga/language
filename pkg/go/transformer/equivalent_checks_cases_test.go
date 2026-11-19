package transformer_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
	"unicode/utf16"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/openfga/language/pkg/go/transformer"
)

var updateEquivalentChecksCases = flag.Bool(
	"update",
	false,
	"rewrite tests/data/equivalent-checks-cases.json from tests/data/equivalent-checks-cases.yaml",
)

const (
	equivalentChecksCasesNote = "Generated from equivalent-checks-cases.yaml, with every model as JSON. " +
		"Don't edit this file: edit the YAML file and run `make generate-equivalent-checks-json`."
	jsonPrintWidth = 80
)

var (
	jsonNumberRegex      = regexp.MustCompile(`^-?(0|[1-9][0-9]*)(\.[0-9]+)?([eE][+-]?[0-9]+)?$`)
	yamlBooleanWordRegex = regexp.MustCompile(`^(y|Y|yes|Yes|YES|n|N|no|No|NO|true|True|TRUE|false|False|FALSE|on|On|ON|off|Off|OFF)$`)
)

var errInvalidEquivalentChecksCases = errors.New("invalid equivalent checks cases")

type jsonField struct {
	key   string
	value any
}

// jsonObject keeps the key order of the YAML source.
type jsonObject []jsonField

// TestEquivalentChecksCases makes sure every model in the shared cases converts,
// and that the JSON copy for SDKs without a DSL parser matches the YAML source.
func TestEquivalentChecksCases(t *testing.T) {
	dataPath := filepath.Join("../../../tests", "data")
	yamlPath := filepath.Join(dataPath, "equivalent-checks-cases.yaml")
	jsonPath := filepath.Join(dataPath, "equivalent-checks-cases.json")

	source, err := os.ReadFile(yamlPath)
	require.NoError(t, err)

	generated, err := generateEquivalentChecksCases(source)
	require.NoError(t, err)

	if *updateEquivalentChecksCases {
		require.NoError(t, os.WriteFile(jsonPath, generated, 0o644)) //nolint:gosec

		return
	}

	committed, err := os.ReadFile(jsonPath)
	require.NoError(t, err)

	require.Equal(
		t,
		decodeKeepingNumbers(t, generated),
		decodeKeepingNumbers(t, committed),
		"equivalent-checks-cases.json is out of date: run `make generate-equivalent-checks-json`",
	)
}

// TestEquivalentChecksCasesRejectsInvalidSource makes sure the generator fails instead of
// silently dropping or duplicating data.
func TestEquivalentChecksCasesRejectsInvalidSource(t *testing.T) {
	t.Parallel()

	const validCase = "- name: a\n  json: '{}'\n  checks: []\n  expected_sent: []\n"

	_, err := generateEquivalentChecksCases([]byte(validCase))
	require.NoError(t, err)

	tests := []struct {
		name    string
		source  string
		message string
	}{
		{"a model followed by more JSON", strings.Replace(validCase, "'{}'", "'{} {}'", 1), "expected one JSON value"},
		{"a second YAML document", validCase + "---\n" + validCase, "expected one YAML document"},
		{"a model field", validCase + "  model: {}\n", "use dsl or json, not model"},
		{"a repeated key", validCase + "  checks: []\n", `repeated key "checks"`},
		{"an unquoted boolean word", strings.Replace(validCase, "checks: []", "checks: [yes]", 1), `quote "yes"`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := generateEquivalentChecksCases([]byte(test.source))
			require.ErrorIs(t, err, errInvalidEquivalentChecksCases)
			require.ErrorContains(t, err, test.message)
		})
	}
}

// decodeKeepingNumbers compares numbers by their text, so 1 and 1.0 or 0 and -0.0 stay different.
func decodeKeepingNumbers(t *testing.T, data []byte) any {
	t.Helper()

	value, err := decodeOneJSONValue(data)
	require.NoError(t, err)

	return value
}

// decodeOneJSONValue keeps numbers as written and fails if anything follows the first value.
func decodeOneJSONValue(data []byte) (any, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	var value any
	if err := decoder.Decode(&value); err != nil {
		return nil, err //nolint:wrapcheck
	}

	if err := decoder.Decode(&json.RawMessage{}); !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("%w: expected one JSON value", errInvalidEquivalentChecksCases)
	}

	return value, nil
}

func generateEquivalentChecksCases(source []byte) ([]byte, error) {
	decoder := yaml.NewDecoder(bytes.NewReader(source))

	var document yaml.Node
	if err := decoder.Decode(&document); err != nil {
		return nil, fmt.Errorf("%w: %w", errInvalidEquivalentChecksCases, err)
	}

	if err := decoder.Decode(&yaml.Node{}); !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("%w: expected one YAML document", errInvalidEquivalentChecksCases)
	}

	if len(document.Content) != 1 || document.Content[0].Kind != yaml.SequenceNode {
		return nil, fmt.Errorf("%w: expected a list of cases", errInvalidEquivalentChecksCases)
	}

	names := map[string]bool{}
	cases := []any{}

	for _, node := range document.Content[0].Content {
		testCase, err := equivalentChecksCase(node, names)
		if err != nil {
			return nil, err
		}

		cases = append(cases, testCase)
	}

	var out strings.Builder
	writeJSON(&out, jsonObject{{"note", equivalentChecksCasesNote}, {"cases", cases}}, "", 0, 0)
	out.WriteString("\n")

	return []byte(out.String()), nil
}

func equivalentChecksCase(node *yaml.Node, names map[string]bool) (jsonObject, error) {
	value, err := fromYAML(node)
	if err != nil {
		return nil, err
	}

	fields, ok := value.(jsonObject)
	if !ok {
		return nil, fmt.Errorf("%w: line %d: a case must be a mapping", errInvalidEquivalentChecksCases, node.Line)
	}

	name, _ := fields.get("name").(string)
	if name == "" || names[name] {
		return nil, fmt.Errorf(
			"%w: line %d: a case needs a unique name, got %q",
			errInvalidEquivalentChecksCases,
			node.Line,
			name,
		)
	}

	names[name] = true

	models := 0
	result := jsonObject{}

	for _, field := range fields {
		switch field.key {
		case "dsl", "json":
			models++

			model, err := modelAsJSON(field.key, field.value)
			if err != nil {
				return nil, fmt.Errorf("%w: case %q: %w", errInvalidEquivalentChecksCases, name, err)
			}

			result = append(result, jsonField{"model", model})
		case "model":
			return nil, fmt.Errorf("%w: case %q: use dsl or json, not model", errInvalidEquivalentChecksCases, name)
		default:
			result = append(result, field)
		}
	}

	if models != 1 {
		return nil, fmt.Errorf("%w: case %q needs exactly one of dsl or json", errInvalidEquivalentChecksCases, name)
	}

	return result, nil
}

func modelAsJSON(kind string, value any) (any, error) {
	text, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("%w: %s must be a string", errInvalidEquivalentChecksCases, kind)
	}

	if kind == "dsl" {
		converted, err := transformer.TransformDSLToJSON(text)
		if err != nil {
			return nil, err //nolint:wrapcheck
		}

		text = converted
	}

	model, err := decodeOneJSONValue([]byte(text))
	if err != nil {
		return nil, err
	}

	return sortedKeys(model), nil
}

func (object jsonObject) get(key string) any {
	for _, field := range object {
		if field.key == key {
			return field.value
		}
	}

	return nil
}

// fromYAML keeps every number exactly as written. A plain YAML to JSON conversion
// would turn 1.0 into 1 and -0.0 into -0, and the cases depend on that difference.
func fromYAML(node *yaml.Node) (any, error) {
	switch node.Kind {
	case yaml.MappingNode:
		object := jsonObject{}
		seen := map[string]bool{}

		for i := 0; i < len(node.Content); i += 2 {
			key := node.Content[i]
			if err := checkPortable(key); err != nil {
				return nil, err
			}

			if key.ShortTag() != "!!str" {
				return nil, fmt.Errorf("%w: line %d: keys must be strings", errInvalidEquivalentChecksCases, key.Line)
			}

			// Decoding into yaml.Node skips the parser's own duplicate key check.
			if seen[key.Value] {
				return nil, fmt.Errorf("%w: line %d: repeated key %q", errInvalidEquivalentChecksCases, key.Line, key.Value)
			}

			seen[key.Value] = true

			value, err := fromYAML(node.Content[i+1])
			if err != nil {
				return nil, err
			}

			object = append(object, jsonField{key.Value, value})
		}

		return object, nil
	case yaml.SequenceNode:
		list := []any{}

		for _, item := range node.Content {
			value, err := fromYAML(item)
			if err != nil {
				return nil, err
			}

			list = append(list, value)
		}

		return list, nil
	case yaml.ScalarNode:
		if err := checkPortable(node); err != nil {
			return nil, err
		}

		switch node.ShortTag() {
		case "!!str":
			return node.Value, nil
		case "!!bool":
			// checkPortable allows only true and false.
			return node.Value == "true", nil
		case "!!null":
			return nil, nil
		case "!!int", "!!float":
			if !jsonNumberRegex.MatchString(node.Value) {
				return nil, fmt.Errorf(
					"%w: line %d: %q isn't a JSON number",
					errInvalidEquivalentChecksCases,
					node.Line,
					node.Value,
				)
			}

			return json.Number(node.Value), nil
		}
	}

	return nil, fmt.Errorf("%w: line %d: unsupported YAML value", errInvalidEquivalentChecksCases, node.Line)
}

// checkPortable rejects unquoted words that YAML parsers disagree on, like True, yes or n,
// so every SDK reading the YAML gets the same values as the JSON copy.
func checkPortable(node *yaml.Node) error {
	quoted := node.Style&(yaml.DoubleQuotedStyle|yaml.SingleQuotedStyle|yaml.LiteralStyle|yaml.FoldedStyle) != 0
	if quoted || node.Value == "true" || node.Value == "false" || !yamlBooleanWordRegex.MatchString(node.Value) {
		return nil
	}

	return fmt.Errorf(
		"%w: line %d: write booleans as true or false, and quote %q if it's a string",
		errInvalidEquivalentChecksCases,
		node.Line,
		node.Value,
	)
}

func sortedKeys(value any) any {
	switch typed := value.(type) {
	case map[string]any:
		keys := make([]string, 0, len(typed))
		for key := range typed {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		object := jsonObject{}
		for _, key := range keys {
			object = append(object, jsonField{key, sortedKeys(typed[key])})
		}

		return object
	case []any:
		list := make([]any, 0, len(typed))
		for _, item := range typed {
			list = append(list, sortedKeys(item))
		}

		return list
	default:
		return value
	}
}

// writeJSON writes JSON the way Prettier formats it: objects on several lines, and lists
// of plain values on one line when they fit.
func writeJSON(out *strings.Builder, value any, indent string, column int, trailing int) {
	switch typed := value.(type) {
	case jsonObject:
		if len(typed) == 0 {
			out.WriteString("{}")

			return
		}

		inner := indent + "  "

		out.WriteString("{\n")

		for i, field := range typed {
			key := quoteJSON(field.key) + ": "
			comma := 0

			if i < len(typed)-1 {
				comma = 1
			}

			out.WriteString(inner + key)
			writeJSON(out, field.value, inner, len(inner)+len(key), comma)

			if comma == 1 {
				out.WriteString(",")
			}

			out.WriteString("\n")
		}

		out.WriteString(indent + "}")
	case []any:
		if len(typed) == 0 {
			out.WriteString("[]")

			return
		}

		if line, ok := inlineList(typed); ok && column+len(line)+trailing <= jsonPrintWidth {
			out.WriteString(line)

			return
		}

		inner := indent + "  "

		out.WriteString("[\n")

		for i, item := range typed {
			comma := 0

			if i < len(typed)-1 {
				comma = 1
			}

			out.WriteString(inner)
			writeJSON(out, item, inner, len(inner), comma)

			if comma == 1 {
				out.WriteString(",")
			}

			out.WriteString("\n")
		}

		out.WriteString(indent + "]")
	default:
		out.WriteString(scalarJSON(typed))
	}
}

func inlineList(list []any) (string, bool) {
	items := make([]string, 0, len(list))

	for _, item := range list {
		switch item.(type) {
		case jsonObject, []any:
			return "", false
		}

		items = append(items, scalarJSON(item))
	}

	return "[" + strings.Join(items, ", ") + "]", true
}

func scalarJSON(value any) string {
	switch typed := value.(type) {
	case nil:
		return "null"
	case bool:
		if typed {
			return "true"
		}

		return "false"
	case json.Number:
		return typed.String()
	case string:
		return quoteJSON(typed)
	default:
		panic(fmt.Sprintf("unexpected JSON value %T", value))
	}
}

// quoteJSON escapes every non-ASCII character, so the file stays the same even if an editor
// changes its Unicode form.
func quoteJSON(text string) string {
	var buffer bytes.Buffer

	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode(text)

	var out strings.Builder

	for _, r := range strings.TrimSuffix(buffer.String(), "\n") {
		if r < utf8.RuneSelf {
			out.WriteRune(r)

			continue
		}

		for _, unit := range utf16.Encode([]rune{r}) {
			fmt.Fprintf(&out, `\u%04x`, unit)
		}
	}

	return out.String()
}
