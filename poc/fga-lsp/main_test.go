package main

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

// TestDiagnosticsForTheCycleModels pins what the server publishes when each
// model in pkg/go/graph/testdata/cycles is opened, and when a model of densely
// nested relations around one broken cycle is. It calls diagnose, which is the
// whole of what a didOpen publishes, so it needs no client.
func TestDiagnosticsForTheCycleModels(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		dsl  string
		want string
	}{
		"everything-at-once": {
			dsl:  cycleModel(t, "everything-at-once"),
			want: "cyclic-relation 22, graph-model-unbuildable 1, relation-no-entry-point 2; 25 of 25 placed",
		},
		"file-storage-ten-loops": {
			dsl:  cycleModel(t, "file-storage-ten-loops"),
			want: "cyclic-relation 23; 23 of 23 placed",
		},
		"knowledge-base-seven-loops": {
			dsl:  cycleModel(t, "knowledge-base-seven-loops"),
			want: "cyclic-relation 17, relation-no-entry-point 2; 19 of 19 placed",
		},
		"mixed-missing-from-target": {
			dsl:  cycleModel(t, "mixed-missing-from-target"),
			want: "cyclic-relation 2, relation-no-entry-point 1; 3 of 3 placed",
		},
		"mixed-no-entry-point": {
			dsl:  cycleModel(t, "mixed-no-entry-point"),
			want: "cyclic-relation 2, relation-no-entry-point 2; 4 of 4 placed",
		},
		"mixed-self-rewrite": {
			dsl:  cycleModel(t, "mixed-self-rewrite"),
			want: "cyclic-relation 2, relation-no-entry-point 1; 3 of 3 placed",
		},
		"pattern-matrix-eleven-loops": {
			dsl:  cycleModel(t, "pattern-matrix-eleven-loops"),
			want: "cyclic-relation 22; 22 of 22 placed",
		},
		"8 densely nested relations": {
			dsl:  denselyNested(8),
			want: "cyclic-relation 9; 9 of 9 placed",
		},
		"9 densely nested relations": {
			dsl:  denselyNested(9),
			want: "cyclic-relation 10; 10 of 10 placed",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := summarize(diagnose(test.dsl)); got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}

// cycleModel reads a model from pkg/go/graph/testdata/cycles.
func cycleModel(t *testing.T, name string) string {
	t.Helper()

	dsl, err := os.ReadFile(filepath.Join("..", "..", "pkg", "go", "graph", "testdata", "cycles", name+".fga"))
	if err != nil {
		t.Fatal(err)
	}

	return string(dsl)
}

// denselyNested writes relations that each accept the others' members, and one
// broken cycle through the first of them: bad subtracts blocked from r1's
// members, and r1 accepts bad's.
func denselyNested(relationCount int) string {
	var dsl strings.Builder

	dsl.WriteString("model\n  schema 1.1\ntype user\ntype group\n  relations\n")

	for i := 1; i <= relationCount; i++ {
		references := []string{"user"}

		for j := 1; j <= relationCount; j++ {
			if j != i {
				references = append(references, fmt.Sprintf("group#r%d", j))
			}
		}

		if i == 1 {
			references = append(references, "group#bad")
		}

		fmt.Fprintf(&dsl, "    define r%d: [%s]\n", i, strings.Join(references, ", "))
	}

	dsl.WriteString("    define bad: [user, group#r1] but not blocked\n    define blocked: [user]\n")

	return dsl.String()
}

// summarize counts diagnostics by code, and how many of them have a position.
func summarize(diagnostics []diagnostic) string {
	counts := map[string]int{}
	placed := 0

	for _, d := range diagnostics {
		counts[d.Code]++

		if d.Range != (lspRange{}) {
			placed++
		}
	}

	parts := make([]string, 0, len(counts))
	for _, code := range slices.Sorted(maps.Keys(counts)) {
		parts = append(parts, fmt.Sprintf("%s %d", code, counts[code]))
	}

	return fmt.Sprintf("%s; %d of %d placed", strings.Join(parts, ", "), placed, len(diagnostics))
}
