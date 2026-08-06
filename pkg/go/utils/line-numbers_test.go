package utils

import "testing"

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
