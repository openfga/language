package utils

import "testing"

func TestIsNameByte(t *testing.T) {
	t.Parallel()

	// IDENTIFIER: (LETTER | '_') (LETTER | DIGIT | '_' | MINUS)*, and
	// EXTENDED_IDENTIFIER adds SLASH and DOT (OpenFGALexer.g4), so the name bytes
	// are exactly the ASCII letters, digits, `_`, `-`, `/` and `.`. Checking every
	// byte pins the table to those ranges, so a mistyped entry cannot survive.
	isName := func(b byte) bool {
		switch {
		case b >= 'a' && b <= 'z', b >= 'A' && b <= 'Z', b >= '0' && b <= '9':
			return true
		case b == '_', b == '-', b == '/', b == '.':
			return true
		}

		return false
	}

	for i := range 256 {
		b := byte(i)
		if got, want := IsNameByte(b), isName(b); got != want {
			t.Errorf("IsNameByte(0x%02X) = %v, want %v", b, got, want)
		}
	}
}

func TestNormalizeWhitespace(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "already normal returns unchanged", in: "define owner: [user]", want: "define owner: [user]"},
		{name: "collapses repeated spaces", in: "define  owner:", want: "define owner:"},
		{name: "folds a tab", in: "define\towner:", want: "define owner:"},
		{name: "folds a form feed", in: "define\fowner:", want: "define owner:"},
		{name: "folds a mixed run", in: "extend \t type\f\forg", want: "extend type org"},
		{name: "folds leading and trailing runs", in: "\t define owner \f", want: " define owner "},
		{name: "leaves other bytes alone", in: "a\vb", want: "a\vb"},
		{name: "empty line", in: "", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NormalizeWhitespace(tt.in); got != tt.want {
				t.Errorf("NormalizeWhitespace(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestGetTypeLineNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		typeName string
		lines    []string
		want     int
	}{
		{
			name:     "finds the declaration",
			typeName: "user",
			lines:    []string{"model", "  schema 1.2", "type user", "  relations"},
			want:     2,
		},
		{
			name:     "does not match a type whose name extends the one asked for",
			typeName: "document",
			lines:    []string{"type documentation", "  relations", "type document", "  relations"},
			want:     2,
		},
		{
			name:     "allows a trailing module comment",
			typeName: "other",
			lines:    []string{"type user", "type other # module: core, file: core.fga"},
			want:     1,
		},
		{
			// `#`, `/`, `.` and `-` are all name bytes or name terminators per the
			// grammar's identifier rules; `#` terminates, so this resolves to the
			// first line. `type doc#x` is not valid DSL, so it cannot reach here
			// from a parsed model — the case is pinned only to document the choice.
			name:     "treats a glued hash as ending the name",
			typeName: "doc",
			lines:    []string{"type doc#x", "type doc"},
			want:     0,
		},
		{
			// EXTENDED_IDENTIFIER admits `.`, `/` and `-` inside a name, so a
			// module-qualified name must not match a longer one that shares its
			// prefix.
			name:     "does not match a dotted name that extends the one asked for",
			typeName: "core.doc",
			lines:    []string{"type core.doc2", "type core.doc"},
			want:     1,
		},
		{
			name:     "does not match a hyphenated name that extends the one asked for",
			typeName: "my-doc",
			lines:    []string{"type my-doc-archive", "type my-doc"},
			want:     1,
		},
		{
			name:     "does not match a slashed name that extends the one asked for",
			typeName: "internal",
			lines:    []string{"type internal/doc", "type internal"},
			want:     1,
		},
		{
			name:     "does not match an extend declaration",
			typeName: "user",
			lines:    []string{"extend type user", "type user"},
			want:     1,
		},
		{
			// WHITESPACE admits \f, so it must separate a name from its trailing
			// comment as a space would; trimming only " \t" collapsed this to 0:0.
			name:     "allows a form feed before a trailing comment",
			typeName: "other",
			lines:    []string{"type user", "type other\f# module: core, file: core.fga"},
			want:     1,
		},
		{
			// WHITESPACE is ('\t' | ' ' | '\u000C')+, so any of the three can
			// separate `type` from the name.
			name:     "allows a tab between type and the name",
			typeName: "user",
			lines:    []string{"model", "type\tuser"},
			want:     1,
		},
		{
			name:     "allows a form feed between type and the name",
			typeName: "user",
			lines:    []string{"model", "type\fuser"},
			want:     1,
		},
		{
			name:     "collapses a mixed run of whitespace between type and the name",
			typeName: "user",
			lines:    []string{"model", "type \t user"},
			want:     1,
		},
		{
			name:     "returns -1 when absent",
			typeName: "missing",
			lines:    []string{"type user", "type org"},
			want:     -1,
		},
		{
			name:     "returns -1 for no lines",
			typeName: "user",
			lines:    nil,
			want:     -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := GetTypeLineNumber(tt.typeName, tt.lines); got != tt.want {
				t.Errorf("GetTypeLineNumber(%q) = %d, want %d", tt.typeName, got, tt.want)
			}
		})
	}
}

func TestGetExtendedTypeLineNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		typeName string
		lines    []string
		want     int
	}{
		{
			name:     "finds the extend declaration",
			typeName: "user",
			lines:    []string{"type user", "extend type user", "  relations"},
			want:     1,
		},
		{
			name:     "does not match a type whose name extends the one asked for",
			typeName: "document",
			lines:    []string{"extend type documentation", "extend type document"},
			want:     1,
		},
		{
			name:     "allows a trailing module comment",
			typeName: "org",
			lines:    []string{"extend type org # module: org, file: org.fga"},
			want:     0,
		},
		{
			name:     "allows a form feed before a trailing comment",
			typeName: "org",
			lines:    []string{"extend type org\f# module: org, file: org.fga"},
			want:     0,
		},
		{
			name:     "allows tabs between extend, type and the name",
			typeName: "org",
			lines:    []string{"type org", "extend\ttype\torg"},
			want:     1,
		},
		{
			name:     "returns -1 when absent",
			typeName: "user",
			lines:    []string{"type user"},
			want:     -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := GetExtendedTypeLineNumber(tt.typeName, tt.lines); got != tt.want {
				t.Errorf("GetExtendedTypeLineNumber(%q) = %d, want %d", tt.typeName, got, tt.want)
			}
		})
	}
}

func TestGetRelationLineNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		relation string
		lines    []string
		want     int
	}{
		{
			name:     "finds the definition",
			relation: "owner",
			lines:    []string{"type doc", "  relations", "    define owner: [user]"},
			want:     2,
		},
		{
			name:     "does not match a relation whose name extends the one asked for",
			relation: "owner",
			lines:    []string{"    define owner_group: [group]", "    define owner: [user]"},
			want:     1,
		},
		{
			name:     "tolerates repeated spaces after define",
			relation: "owner",
			lines:    []string{"    define  owner: [user]"},
			want:     0,
		},
		{
			name:     "allows whitespace before the colon",
			relation: "owner",
			lines:    []string{"    define owner : [user]"},
			want:     0,
		},
		{
			// `define owner\f: [user]` parses, so it must be findable.
			name:     "allows a form feed before the colon",
			relation: "owner",
			lines:    []string{"    define owner\f: [user]"},
			want:     0,
		},
		{
			// `define\towner: [user]` parses — WHITESPACE admits tabs — so the
			// definition must be findable.
			name:     "allows a tab between define and the name",
			relation: "owner",
			lines:    []string{"    define\towner: [user]"},
			want:     0,
		},
		{
			name:     "allows a form feed between define and the name",
			relation: "owner",
			lines:    []string{"    define\fowner: [user]"},
			want:     0,
		},
		{
			name:     "collapses a mixed run of whitespace between define and the name",
			relation: "owner",
			lines:    []string{"    define \t owner: [user]"},
			want:     0,
		},
		{
			name:     "returns -1 when absent",
			relation: "missing",
			lines:    []string{"    define owner: [user]"},
			want:     -1,
		},
		{
			name:     "returns -1 for no lines",
			relation: "owner",
			lines:    nil,
			want:     -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := GetRelationLineNumber(tt.relation, tt.lines); got != tt.want {
				t.Errorf("GetRelationLineNumber(%q) = %d, want %d", tt.relation, got, tt.want)
			}
		})
	}
}

func TestGetConditionLineNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		conditionName string
		lines         []string
		want          int
	}{
		{
			name:          "finds the declaration",
			conditionName: "in_range",
			lines:         []string{"type user", "condition in_range(x: int) {"},
			want:          1,
		},
		{
			name:          "does not match a condition whose name extends the one asked for",
			conditionName: "less",
			lines:         []string{"condition less_than(x: int) {", "condition less(x: int) {"},
			want:          1,
		},
		{
			name:          "allows whitespace before the parameter list",
			conditionName: "less",
			lines:         []string{"condition less (x: int) {"},
			want:          0,
		},
		{
			// `condition less\f(x: int) {` parses, so it must be findable.
			name:          "allows a form feed before the parameter list",
			conditionName: "less",
			lines:         []string{"condition less\f(x: int) {"},
			want:          0,
		},
		{
			name:          "allows a tab between condition and the name",
			conditionName: "less",
			lines:         []string{"condition\tless(x: int) {"},
			want:          0,
		},
		{
			name:          "collapses a mixed run of whitespace between condition and the name",
			conditionName: "less",
			lines:         []string{"condition \t less(x: int) {"},
			want:          0,
		},
		{
			name:          "returns -1 when absent",
			conditionName: "missing",
			lines:         []string{"condition less(x: int) {"},
			want:          -1,
		},
		{
			name:          "returns -1 for no lines",
			conditionName: "less",
			lines:         nil,
			want:          -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := GetConditionLineNumber(tt.conditionName, tt.lines); got != tt.want {
				t.Errorf("GetConditionLineNumber(%q) = %d, want %d", tt.conditionName, got, tt.want)
			}
		})
	}
}
