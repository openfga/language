package validation

import (
	"fmt"
	"strings"
	"testing"

	"github.com/openfga/language/pkg/go/transformer"
)

// buildDeepChainDSL builds one type with a long linear chain of computed-userset
// relations r0 -> r1 -> ... -> base -> [user], the worst case for the
// non-branching recursion path in hasEntryPointOrLoop.
func buildDeepChainDSL(depth int) string {
	var b strings.Builder
	b.WriteString("model\n  schema 1.1\ntype user\n")
	b.WriteString("type doc\n  relations\n")
	b.WriteString("    define base: [user]\n")
	prev := "base"
	for i := 0; i < depth; i++ {
		name := fmt.Sprintf("r%d", i)
		fmt.Fprintf(&b, "    define %s: %s\n", name, prev)
		prev = name
	}
	return b.String()
}

// buildWideUnionDSL builds a relation whose rewrite is a wide union of computed
// usersets, each resolving down its own path, stressing the sibling-isolating
// copies in the union branch.
func buildWideUnionDSL(width int) string {
	var b strings.Builder
	b.WriteString("model\n  schema 1.1\ntype user\n")
	b.WriteString("type doc\n  relations\n")
	b.WriteString("    define base: [user]\n")
	members := make([]string, 0, width)
	for i := 0; i < width; i++ {
		name := fmt.Sprintf("m%d", i)
		fmt.Fprintf(&b, "    define %s: base\n", name)
		members = append(members, name)
	}
	fmt.Fprintf(&b, "    define wide: %s\n", strings.Join(members, " or "))
	return b.String()
}

// TestCycleDetection_DeepChainTerminatesWithEntry verifies a long linear chain
// of computed usersets terminating in a direct assignment resolves with an entry
// point and terminates. Guards that the shared visited map still flows down the
// chain after dropping the per-call copy.
func TestCycleDetection_DeepChainTerminatesWithEntry(t *testing.T) {
	dsl := buildDeepChainDSL(1000)
	model, err := transformer.TransformDSLToProto(dsl)
	if err != nil {
		t.Fatalf("failed to transform DSL: %v", err)
	}
	lines := strings.Split(dsl, "\n")
	collector := NewErrorCollector(lines)

	ValidateCyclesAndEntryPoints(collector, model, lines)

	if collector.HasErrors() {
		t.Fatalf("deep computed-userset chain ending in a direct assignment should "+
			"have an entry point, got %d errors: %v", collector.Count(), collector.GetErrors())
	}
}

// TestCycleDetection_WideUnionTerminatesWithEntry verifies that a wide union of
// computed usersets, all reachable down to a concrete type, resolves with an
// entry point and that the sibling-isolating copies don't change the outcome.
func TestCycleDetection_WideUnionTerminatesWithEntry(t *testing.T) {
	dsl := buildWideUnionDSL(1000)
	model, err := transformer.TransformDSLToProto(dsl)
	if err != nil {
		t.Fatalf("failed to transform DSL: %v", err)
	}
	lines := strings.Split(dsl, "\n")
	collector := NewErrorCollector(lines)

	ValidateCyclesAndEntryPoints(collector, model, lines)

	if collector.HasErrors() {
		t.Fatalf("wide union of resolvable members should have an entry point, "+
			"got %d errors: %v", collector.Count(), collector.GetErrors())
	}
}

// TestCycleDetection_SelfReferentialChainIsLoop verifies a relation that
// computes itself through a chain (a->b->c->a) is still detected as a loop with
// no entry point — the shared visited map must accumulate to see the back-edge.
func TestCycleDetection_SelfReferentialChainIsLoop(t *testing.T) {
	dsl := `model
  schema 1.1
type user
type doc
  relations
    define a: b
    define b: c
    define c: a`
	model, err := transformer.TransformDSLToProto(dsl)
	if err != nil {
		t.Fatalf("failed to transform DSL: %v", err)
	}
	lines := strings.Split(dsl, "\n")
	collector := NewErrorCollector(lines)

	ValidateCyclesAndEntryPoints(collector, model, lines)

	if !collector.HasErrors() {
		t.Fatal("self-referential computed chain a->b->c->a should be reported as " +
			"having no entry point")
	}
}

// TestCycleDetection_UnionSiblingIsolation verifies a looping union member does
// not poison a sibling that has an entry point: `mixed: loops or direct` resolves
// with an entry point. This is what the per-branch copyVisited must preserve.
func TestCycleDetection_UnionSiblingIsolation(t *testing.T) {
	dsl := `model
  schema 1.1
type user
type doc
  relations
    define direct: [user]
    define loops: selfref
    define selfref: loops
    define mixed: loops or direct`
	model, err := transformer.TransformDSLToProto(dsl)
	if err != nil {
		t.Fatalf("failed to transform DSL: %v", err)
	}
	lines := strings.Split(dsl, "\n")
	collector := NewErrorCollector(lines)

	ValidateCyclesAndEntryPoints(collector, model, lines)

	// `loops` and `selfref` legitimately have no entry point and are reported.
	// `mixed` and `direct` must NOT be reported.
	for _, e := range collector.GetErrors() {
		if strings.Contains(e.Message, "mixed") || strings.Contains(e.Message, "`direct`") {
			t.Fatalf("relation with a resolvable union branch should have an entry "+
				"point, but got error: %s", e.Message)
		}
	}
}

// TestCycleDetection_DeepChainCountStable checks a chain whose base self-loops
// yields exactly one no-entry-point report per relation — the optimization must
// not suppress or duplicate findings.
func TestCycleDetection_DeepChainCountStable(t *testing.T) {
	var b strings.Builder
	b.WriteString("model\n  schema 1.1\ntype user\ntype doc\n  relations\n")
	b.WriteString("    define base: base\n")
	prev := "base"
	const depth = 50
	for i := 0; i < depth; i++ {
		name := fmt.Sprintf("r%d", i)
		fmt.Fprintf(&b, "    define %s: %s\n", name, prev)
		prev = name
	}
	dsl := b.String()
	model, err := transformer.TransformDSLToProto(dsl)
	if err != nil {
		t.Fatalf("failed to transform DSL: %v", err)
	}
	lines := strings.Split(dsl, "\n")
	collector := NewErrorCollector(lines)

	ValidateCyclesAndEntryPoints(collector, model, lines)

	// base + r0..r49 = depth+1 relations, all with no entry point.
	if collector.Count() != depth+1 {
		t.Fatalf("expected %d no-entry-point errors, got %d: %v",
			depth+1, collector.Count(), collector.GetErrors())
	}
}
