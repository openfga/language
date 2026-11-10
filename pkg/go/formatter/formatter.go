// Package formatter rewrites OpenFGA DSL files into a consistent layout.
package formatter

import (
	"errors"
	"strings"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"google.golang.org/protobuf/proto"

	"github.com/openfga/language/pkg/go/transformer"
)

// ErrModelChanged is returned when the formatted output would not parse back to the same model.
// It means the input uses a layout the formatter does not handle yet, and the input is left as it is.
var ErrModelChanged = errors.New("formatting would change the meaning of the model")

const defaultIndentSize = 2

type options struct {
	indent string
}

// Option configures Format.
type Option func(*options)

// WithIndentSize sets the number of spaces used for each indentation level. The default is 2.
// Tabs are not offered because the DSL parser only treats space indented lines that start with
// `#` as comments.
func WithIndentSize(size int) Option {
	return func(o *options) {
		if size > 0 {
			o.indent = strings.Repeat(" ", size)
		}
	}
}

// Format takes the contents of a DSL file (a model or a module) and returns it in the canonical
// layout: one indentation level for `schema` and `relations`, two for `define`, a single blank line
// between types and conditions, single spaces around operators and after commas, and a trailing
// newline. Comments are kept and move with the line they belong to. The body of a condition is
// kept as written, apart from its indentation when it fits on one line.
//
// The input must be valid DSL. Before returning, Format parses the output and checks that it
// produces the same model as the input, so formatting never changes what the model means.
func Format(dsl string, opts ...Option) (string, error) {
	o := options{indent: strings.Repeat(" ", defaultIndentSize)}
	for _, opt := range opts {
		opt(&o)
	}

	dsl = strings.ReplaceAll(dsl, "\r\n", "\n")

	before, beforeExtensions, err := transformer.TransformModularDSLToProto(dsl)
	if err != nil {
		return "", err
	}

	p := &printer{indent: o.indent}
	formatted := p.format(strings.Split(dsl, "\n"))

	after, afterExtensions, err := transformer.TransformModularDSLToProto(formatted)
	if err != nil {
		return "", ErrModelChanged
	}

	if !sameModel(before, after) || !sameExtensions(beforeExtensions, afterExtensions) {
		return "", ErrModelChanged
	}

	return formatted, nil
}

// MustFormat calls Format and panics if it returns an error.
func MustFormat(dsl string, opts ...Option) string {
	formatted, err := Format(dsl, opts...)
	if err != nil {
		panic(err)
	}

	return formatted
}

func sameModel(a, b *openfgav1.AuthorizationModel) bool {
	return proto.Equal(normalizeModel(a), normalizeModel(b))
}

// normalizeModel trims the condition expressions, since moving a closing brace to its own line
// only changes the whitespace around the expression.
func normalizeModel(model *openfgav1.AuthorizationModel) *openfgav1.AuthorizationModel {
	clone, _ := proto.Clone(model).(*openfgav1.AuthorizationModel)
	for _, condition := range clone.GetConditions() {
		condition.Expression = strings.TrimSpace(condition.GetExpression())
	}

	return clone
}

func sameExtensions(a, b map[string]*openfgav1.TypeDefinition) bool {
	if len(a) != len(b) {
		return false
	}

	for name, typeDef := range a {
		if !proto.Equal(typeDef, b[name]) {
			return false
		}
	}

	return true
}

type statementKind int

const (
	kindNone statementKind = iota
	kindHeader
	kindType
	kindRelations
	kindDefine
	kindCondition
)

type printer struct {
	indent string
	out    []string

	// pending holds full line comments (and "" for blank lines) seen since the last statement.
	pending []string
	last    statementKind
}

func (p *printer) format(lines []string) string {
	for i := 0; i < len(lines); i++ {
		trimmed := strings.TrimSpace(lines[i])

		switch {
		case trimmed == "":
			p.pending = append(p.pending, "")
		case strings.HasPrefix(trimmed, "#"):
			p.pending = append(p.pending, trimmed)
		default:
			i = p.statement(lines, i)
		}
	}

	p.flushPending()

	return strings.Join(p.out, "\n") + "\n"
}

// statement prints the statement that starts at lines[start] and returns the index of its last line.
func (p *printer) statement(lines []string, start int) int {
	code, comment := splitComment(lines[start])
	fields := strings.Fields(code)
	if len(fields) == 0 {
		p.pending = append(p.pending, comment)

		return start
	}

	switch fields[0] {
	case "model":
		p.emit(kindHeader, 0, "model", comment)
	case "schema":
		p.emit(kindHeader, 1, strings.Join(fields, " "), comment)
	case "module":
		p.emit(kindHeader, 0, strings.Join(fields, " "), comment)
	case "type", "extend":
		p.emit(kindType, 0, strings.Join(fields, " "), comment)
	case "relations":
		p.emit(kindRelations, 1, "relations", comment)
	case "define":
		return p.define(lines, start)
	case "condition":
		return p.condition(lines, start)
	default:
		p.emit(kindNone, 0, strings.TrimRight(lines[start], " "), "")
	}

	return start
}

func (p *printer) define(lines []string, start int) int {
	code, comment := splitComment(lines[start])
	end := start

	// A list of type restrictions can span several lines, so keep reading until the brackets
	// and parentheses are closed.
	for depth(code) > 0 && end+1 < len(lines) {
		end++

		nextCode, nextComment := splitComment(lines[end])
		if nextComment != "" || strings.HasPrefix(strings.TrimSpace(lines[end]), "#") {
			// Comments inside a multi-line definition have nowhere to go once it is joined,
			// so leave the definition as it is.
			p.emit(kindDefine, 2, strings.TrimSpace(lines[start]), "")
			for _, line := range lines[start+1 : end+1] {
				p.out = append(p.out, strings.TrimRight(line, " "))
			}

			for end+1 < len(lines) && depthOf(lines[start:end+1]) > 0 {
				end++
				p.out = append(p.out, strings.TrimRight(lines[end], " "))
			}

			return end
		}

		code += " " + nextCode
	}

	rest := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(code), "define"))
	name, expression, _ := strings.Cut(rest, ":")

	p.emit(kindDefine, 2, "define "+strings.TrimSpace(name)+": "+normalizeExpression(expression), comment)

	return end
}

func (p *printer) condition(lines []string, start int) int {
	// Read the header up to the opening brace.
	header := ""
	headerComment := ""
	end := start
	body := ""

	for ; end < len(lines); end++ {
		code, comment := splitComment(lines[end])
		if comment != "" {
			headerComment = comment
		}

		if before, after, found := strings.Cut(code, "{"); found {
			header += " " + before
			body = after

			break
		}

		header += " " + code
	}

	if end >= len(lines) {
		// No opening brace: not something we can format, print it back unchanged.
		p.emit(kindCondition, 0, strings.TrimSpace(lines[start]), "")
		for _, line := range lines[start+1:] {
			p.out = append(p.out, strings.TrimRight(line, " "))
		}

		return len(lines) - 1
	}

	// Collect the body up to the matching closing brace.
	bodyLines := []string{}
	closingComment := ""
	braces := 1
	current := body
	line := end

	for {
		if idx := closingBrace(current, &braces); idx >= 0 {
			bodyLines = append(bodyLines, current[:idx])
			_, closingComment = splitComment(current[idx+1:])

			break
		}

		bodyLines = append(bodyLines, current)

		line++
		if line >= len(lines) {
			break
		}

		current = lines[line]
	}

	p.emit(kindCondition, 0, normalizeConditionHeader(header)+" {", headerComment)

	bodyLines = trimBlankLines(bodyLines)
	if len(bodyLines) == 1 {
		p.out = append(p.out, p.indent+strings.TrimSpace(bodyLines[0]))
	} else {
		for _, bodyLine := range bodyLines {
			p.out = append(p.out, strings.TrimRight(bodyLine, " "))
		}
	}

	closing := "}"
	if closingComment != "" {
		closing += " " + closingComment
	}

	p.out = append(p.out, closing)

	if line >= len(lines) {
		return len(lines) - 1
	}

	return line
}

// emit prints a statement at the given indentation level, after any comments that came before it.
func (p *printer) emit(kind statementKind, level int, text, comment string) {
	// Types and conditions always get one blank line before them (and before their comments),
	// definitions keep a single blank line between them when the input had one, and everything
	// else drops blank lines.
	forceBlank := kind == kindType || kind == kindCondition
	keepBlank := kind == kindDefine && p.last == kindDefine

	p.flushPendingFor(level, forceBlank, keepBlank)

	line := strings.Repeat(p.indent, level) + text
	if comment != "" {
		line += " " + comment
	}

	p.out = append(p.out, line)
	p.last = kind
}

func (p *printer) flushPendingFor(level int, forceBlank, keepBlank bool) {
	hadBlank := len(p.pending) > 0 && p.pending[0] == ""

	if len(p.out) > 0 && (forceBlank || (keepBlank && hadBlank)) {
		p.out = append(p.out, "")
	}

	p.writePending(level)
}

// flushPending prints the comments left at the end of the file.
func (p *printer) flushPending() {
	if len(trimBlankLines(p.pending)) == 0 {
		p.pending = nil

		return
	}

	if len(p.out) > 0 && p.pending[0] == "" {
		p.out = append(p.out, "")
	}

	p.writePending(0)
}

func (p *printer) writePending(level int) {
	previousBlank := true

	for _, entry := range trimBlankLines(p.pending) {
		if entry == "" {
			if !previousBlank {
				p.out = append(p.out, "")
			}

			previousBlank = true

			continue
		}

		p.out = append(p.out, strings.Repeat(p.indent, level)+entry)
		previousBlank = false
	}

	p.pending = nil
}

// splitComment splits a line into its code and its trailing comment the same way the parser
// does: a comment starts at the first " #".
func splitComment(line string) (string, string) {
	code, comment, found := strings.Cut(line, " #")
	if !found {
		return line, ""
	}

	return code, "#" + strings.TrimRight(comment, " ")
}

func depth(code string) int {
	return strings.Count(code, "[") + strings.Count(code, "(") - strings.Count(code, "]") - strings.Count(code, ")")
}

func depthOf(lines []string) int {
	total := 0

	for _, line := range lines {
		code, _ := splitComment(line)
		if strings.HasPrefix(strings.TrimSpace(line), "#") {
			continue
		}

		total += depth(code)
	}

	return total
}

// normalizeExpression puts single spaces between the words of a relation definition, no spaces
// inside brackets and parentheses, and a space after each comma.
func normalizeExpression(expression string) string {
	var builder strings.Builder

	previous := ""

	for _, token := range tokenize(expression) {
		if previous != "" && previous != "[" && previous != "(" &&
			token != "]" && token != ")" && token != "," {
			builder.WriteByte(' ')
		}

		builder.WriteString(token)
		previous = token
	}

	return builder.String()
}

func tokenize(expression string) []string {
	tokens := []string{}

	for _, field := range strings.Fields(expression) {
		start := 0

		for i, char := range field {
			switch char {
			case '[', ']', '(', ')', ',':
				if i > start {
					tokens = append(tokens, field[start:i])
				}

				tokens = append(tokens, string(char))
				start = i + 1
			}
		}

		if start < len(field) {
			tokens = append(tokens, field[start:])
		}
	}

	return tokens
}

// normalizeConditionHeader turns `condition  name ( a:int,b : string )` into
// `condition name(a: int, b: string)`.
func normalizeConditionHeader(header string) string {
	header = strings.TrimSpace(header)
	header = strings.TrimSpace(strings.TrimPrefix(header, "condition"))

	name, params, found := strings.Cut(header, "(")
	if !found {
		return "condition " + strings.Join(strings.Fields(header), " ")
	}

	params = strings.TrimSpace(params)
	params = strings.TrimSuffix(params, ")")

	normalized := []string{}

	for _, param := range strings.Split(params, ",") {
		paramName, paramType, _ := strings.Cut(param, ":")
		normalized = append(normalized,
			strings.TrimSpace(paramName)+": "+strings.Join(strings.Fields(paramType), ""))
	}

	return "condition " + strings.TrimSpace(name) + "(" + strings.Join(normalized, ", ") + ")"
}

// closingBrace returns the index of the brace that closes the condition body in line, or -1.
// Braces inside string literals are ignored.
func closingBrace(line string, braces *int) int {
	var quote rune

	for i, char := range line {
		switch {
		case quote != 0:
			if char == quote {
				quote = 0
			}
		case char == '"' || char == '\'':
			quote = char
		case char == '{':
			*braces++
		case char == '}':
			*braces--
			if *braces == 0 {
				return i
			}
		}
	}

	return -1
}

func trimBlankLines(lines []string) []string {
	start := 0
	for start < len(lines) && strings.TrimSpace(lines[start]) == "" {
		start++
	}

	end := len(lines)
	for end > start && strings.TrimSpace(lines[end-1]) == "" {
		end--
	}

	return lines[start:end]
}
