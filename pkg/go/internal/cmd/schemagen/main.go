// Command schemagen writes schemas/authorization-model.schema.json from the
// OpenFGA API protobuf definitions, so the schema follows the field names and
// validation rules of the API instead of being maintained by hand.
//
// Run it from pkg/go with `go run ./internal/cmd/schemagen`.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/envoyproxy/protoc-gen-validate/validate"
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	schemaID          = "https://raw.githubusercontent.com/openfga/language/main/schemas/authorization-model.schema.json"
	schemaTitle       = "OpenFGA authorization model"
	schemaDescription = "An authorization model in JSON, as returned by the API, written by `fga model write` and " +
		"produced by the DSL transformers. It checks the shape of the model; use the language validators to " +
		"check that its types and relations are consistent. Generated from the OpenFGA API protobuf " +
		"definitions by pkg/go/internal/cmd/schemagen, do not edit by hand."
)

// optionalFields lists fields whose API validation rules reject an empty value,
// but which a model file doesn't need: the id is assigned when the model is written.
var optionalFields = map[protoreflect.FullName]bool{
	"openfga.v1.AuthorizationModel.id": true,
}

// requiredFields lists fields the API doesn't require but a model can't work
// without: a union or intersection with no child list is malformed.
var requiredFields = map[protoreflect.FullName]bool{
	"openfga.v1.Usersets.child": true,
}

func main() {
	out := flag.String("out", "../../schemas/authorization-model.schema.json", "file to write the schema to")
	flag.Parse()

	schema, err := authorizationModelSchema()
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(*out, schema, 0o600); err != nil {
		log.Fatal(err)
	}
}

// authorizationModelSchema returns the JSON schema of openfga.v1.AuthorizationModel.
func authorizationModelSchema() ([]byte, error) {
	g := &generator{definitions: map[string]any{}}

	root := g.message((&openfgav1.AuthorizationModel{}).ProtoReflect().Descriptor())
	root["$schema"] = "http://json-schema.org/draft-07/schema#"
	root["$id"] = schemaID
	root["title"] = schemaTitle
	root["description"] = schemaDescription
	root["definitions"] = g.definitions

	var buf bytes.Buffer

	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(root); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// generator builds JSON schemas from protobuf message descriptors.
type generator struct {
	definitions map[string]any
}

// message returns the schema of a message, adding the messages it uses to the definitions.
func (g *generator) message(md protoreflect.MessageDescriptor) map[string]any {
	properties := map[string]any{}
	required := []string{}
	allOf := []any{}

	fields := md.Fields()
	for i := range fields.Len() {
		fd := fields.Get(i)
		names := fieldNames(fd)

		schema, isRequired := g.field(fd)
		for _, name := range names {
			properties[name] = schema
		}

		if !isRequired {
			continue
		}

		if len(names) == 1 {
			required = append(required, names[0])

			continue
		}

		// protojson takes either name, so require one of them.
		allOf = append(allOf, map[string]any{"oneOf": requireEach(names)})
	}

	schema := map[string]any{
		"type":                 "object",
		"additionalProperties": false,
		"properties":           properties,
	}

	allOf = append(allOf, g.oneofs(md, schema)...)

	if len(required) > 0 {
		schema["required"] = required
	}

	if len(allOf) > 0 {
		schema["allOf"] = allOf
	}

	return schema
}

// oneofs applies the oneof groups of a message. A message whose fields all
// belong to one oneof (like Userset) means nothing without one of them, so it
// must set exactly one; otherwise at most one field of each group may be set.
func (g *generator) oneofs(md protoreflect.MessageDescriptor, schema map[string]any) []any {
	allOf := []any{}

	oneofs := md.Oneofs()
	for i := range oneofs.Len() {
		od := oneofs.Get(i)
		if od.IsSynthetic() {
			continue
		}

		groupFields := od.Fields()
		if groupFields.Len() == md.Fields().Len() {
			schema["minProperties"] = 1
			schema["maxProperties"] = 1

			continue
		}

		names := []string{}
		for j := range groupFields.Len() {
			names = append(names, fieldNames(groupFields.Get(j))...)
		}

		pairs := []any{}

		for a := range names {
			for b := a + 1; b < len(names); b++ {
				pairs = append(pairs, map[string]any{"required": []string{names[a], names[b]}})
			}
		}

		if len(pairs) > 0 {
			allOf = append(allOf, map[string]any{"not": map[string]any{"anyOf": pairs}})
		}

		if required, _ := proto.GetExtension(od.Options(), validate.E_Required).(bool); required {
			allOf = append(allOf, map[string]any{"anyOf": requireEach(names)})
		}
	}

	return allOf
}

// field returns the schema of a field and whether the field must be present.
func (g *generator) field(fd protoreflect.FieldDescriptor) (map[string]any, bool) {
	rules, _ := proto.GetExtension(fd.Options(), validate.E_Rules).(*validate.FieldRules)

	var schema map[string]any

	switch {
	case fd.IsMap():
		schema = g.mapField(fd, rules.GetMap())
	case fd.IsList():
		schema = g.listField(fd, rules.GetRepeated())
	default:
		schema = g.value(fd, rules)
	}

	required := isRequired(fd, rules)
	if !required && (fd.IsMap() || fd.IsList() || fd.Kind() == protoreflect.MessageKind) {
		// protojson reads null as an unset field.
		schema = nullable(schema)
	}

	return schema, required
}

// mapField returns the schema of a map field.
func (g *generator) mapField(fd protoreflect.FieldDescriptor, rules *validate.MapRules) map[string]any {
	schema := map[string]any{
		"type":                 "object",
		"additionalProperties": g.value(fd.MapValue(), rules.GetValues()),
	}

	if pattern := rules.GetKeys().GetString_().GetPattern(); pattern != "" {
		schema["propertyNames"] = map[string]any{"pattern": pattern}
	}

	if minPairs := rules.GetMinPairs(); minPairs > 0 {
		schema["minProperties"] = minPairs
	}

	if maxPairs := rules.GetMaxPairs(); maxPairs > 0 {
		schema["maxProperties"] = maxPairs
	}

	return schema
}

// listField returns the schema of a repeated field.
func (g *generator) listField(fd protoreflect.FieldDescriptor, rules *validate.RepeatedRules) map[string]any {
	schema := map[string]any{
		"type":  "array",
		"items": g.value(fd, rules.GetItems()),
	}

	if minItems := rules.GetMinItems(); minItems > 0 {
		schema["minItems"] = minItems
	}

	if maxItems := rules.GetMaxItems(); maxItems > 0 {
		schema["maxItems"] = maxItems
	}

	return schema
}

// value returns the schema of a single value of the field's kind.
func (g *generator) value(fd protoreflect.FieldDescriptor, rules *validate.FieldRules) map[string]any {
	switch fd.Kind() {
	case protoreflect.StringKind:
		return stringSchema(rules.GetString_())
	case protoreflect.BytesKind:
		return map[string]any{"type": "string"}
	case protoreflect.BoolKind:
		return map[string]any{"type": "boolean"}
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return map[string]any{"type": "number"}
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		// protojson writes 64-bit integers as strings and reads both.
		return map[string]any{"type": []string{"integer", "string"}}
	case protoreflect.EnumKind:
		return enumSchema(fd.Enum())
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return g.ref(fd.Message())
	default:
		return map[string]any{"type": "integer"}
	}
}

// ref returns a reference to the definition of a message, creating it if needed.
func (g *generator) ref(md protoreflect.MessageDescriptor) map[string]any {
	if schema, ok := wellKnown(md); ok {
		return schema
	}

	name := definitionName(md)
	if _, ok := g.definitions[name]; !ok {
		// Reserve the name first, messages like Userset refer to themselves.
		g.definitions[name] = map[string]any{}
		g.definitions[name] = g.message(md)
	}

	return map[string]any{"$ref": "#/definitions/" + name}
}

// stringSchema returns the schema of a string with its validation rules.
func stringSchema(rules *validate.StringRules) map[string]any {
	constraints := map[string]any{}

	if pattern := rules.GetPattern(); pattern != "" {
		constraints["pattern"] = pattern
	}

	// JSON schema counts characters; byte limits are kept as character limits,
	// which can only be more lenient for multi-byte strings.
	if maxLength := minPositive(rules.GetMaxLen(), rules.GetMaxBytes()); maxLength > 0 {
		constraints["maxLength"] = maxLength
	}

	if minLength := max(rules.GetMinLen(), rules.GetMinBytes()); minLength > 0 {
		constraints["minLength"] = minLength
	}

	if in := rules.GetIn(); len(in) > 0 {
		constraints["enum"] = in
	}

	schema := map[string]any{"type": "string"}

	if rules.GetIgnoreEmpty() && len(constraints) > 0 {
		schema["anyOf"] = []any{map[string]any{"maxLength": 0}, constraints}

		return schema
	}

	for key, value := range constraints {
		schema[key] = value
	}

	return schema
}

// enumSchema accepts the names and the numbers of an enum, as protojson does.
func enumSchema(ed protoreflect.EnumDescriptor) map[string]any {
	names := []any{}
	numbers := []any{}

	values := ed.Values()
	for i := range values.Len() {
		value := values.Get(i)
		names = append(names, string(value.Name()))
		numbers = append(numbers, int32(value.Number()))
	}

	return map[string]any{"enum": append(names, numbers...)}
}

// isRequired reports whether a model must set the field for it to pass the API's validation.
func isRequired(fd protoreflect.FieldDescriptor, rules *validate.FieldRules) bool {
	if optionalFields[fd.FullName()] {
		return false
	}

	if requiredFields[fd.FullName()] {
		return true
	}

	// Members of a oneof are covered by the oneof rules instead.
	if od := fd.ContainingOneof(); od != nil && !od.IsSynthetic() {
		return false
	}

	switch {
	case fd.IsList():
		return rules.GetRepeated().GetMinItems() > 0
	case fd.IsMap():
		return rules.GetMap().GetMinPairs() > 0
	case fd.Kind() == protoreflect.MessageKind:
		return rules.GetMessage().GetRequired()
	case fd.Kind() == protoreflect.StringKind:
		stringRules := rules.GetString_()
		if stringRules == nil || stringRules.GetIgnoreEmpty() {
			return false
		}

		if stringRules.GetMinLen() > 0 || stringRules.GetMinBytes() > 0 {
			return true
		}

		// A missing string reads as "", so a pattern that rejects "" makes it required.
		pattern := stringRules.GetPattern()

		return pattern != "" && !regexp.MustCompile(pattern).MatchString("")
	default:
		return false
	}
}

// wellKnown returns the JSON form of the protobuf well-known types.
func wellKnown(md protoreflect.MessageDescriptor) (map[string]any, bool) {
	switch md.FullName() {
	case "google.protobuf.Struct":
		return map[string]any{"type": "object"}, true
	case "google.protobuf.Value":
		return map[string]any{}, true
	case "google.protobuf.ListValue":
		return map[string]any{"type": "array"}, true
	case "google.protobuf.Timestamp", "google.protobuf.Duration", "google.protobuf.StringValue",
		"google.protobuf.BytesValue", "google.protobuf.FieldMask":
		return map[string]any{"type": "string"}, true
	case "google.protobuf.BoolValue":
		return map[string]any{"type": "boolean"}, true
	case "google.protobuf.Int32Value", "google.protobuf.UInt32Value":
		return map[string]any{"type": "integer"}, true
	case "google.protobuf.Int64Value", "google.protobuf.UInt64Value":
		return map[string]any{"type": []string{"integer", "string"}}, true
	case "google.protobuf.FloatValue", "google.protobuf.DoubleValue":
		return map[string]any{"type": "number"}, true
	case "google.protobuf.Empty":
		return map[string]any{"type": "object", "additionalProperties": false}, true
	default:
		return nil, false
	}
}

// fieldNames returns the names protojson accepts for a field: the JSON name and,
// when it differs, the original proto name.
func fieldNames(fd protoreflect.FieldDescriptor) []string {
	jsonName := fd.JSONName()
	protoName := string(fd.Name())

	if jsonName == protoName {
		return []string{protoName}
	}

	names := []string{jsonName, protoName}
	sort.Strings(names)

	return names
}

// definitionName turns a message name like TypeDefinition into typeDefinition.
func definitionName(md protoreflect.MessageDescriptor) string {
	name := string(md.Name())

	return strings.ToLower(name[:1]) + name[1:]
}

// nullable lets a schema also accept null.
func nullable(schema map[string]any) map[string]any {
	if _, ok := schema["$ref"]; ok {
		return map[string]any{"anyOf": []any{map[string]any{"type": "null"}, schema}}
	}

	if typ, ok := schema["type"].(string); ok {
		schema["type"] = []string{typ, "null"}
	}

	return schema
}

// requireEach returns one "required" clause per name.
func requireEach(names []string) []any {
	clauses := make([]any, 0, len(names))
	for _, name := range names {
		clauses = append(clauses, map[string]any{"required": []string{name}})
	}

	return clauses
}

// minPositive returns the smaller of two limits, ignoring unset (zero) ones.
func minPositive(a, b uint64) uint64 {
	switch {
	case a == 0:
		return b
	case b == 0:
		return a
	default:
		return min(a, b)
	}
}
