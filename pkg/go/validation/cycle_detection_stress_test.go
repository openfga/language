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

// entryPointsFor transforms the DSL and runs the entry-point phase alone.
func entryPointsFor(t *testing.T, dsl string) []*Finding {
	t.Helper()

	model, err := transformer.TransformDSLToProto(dsl)
	if err != nil {
		t.Fatalf("failed to transform DSL: %v", err)
	}

	return ExtractAllAs[*Finding](validateEntryPoints(newIndex(model), newSource(dsl)))
}

// TestCycleDetection_DeepChainTerminatesWithEntry verifies a long linear chain
// of computed usersets terminating in a direct assignment resolves with an entry
// point and terminates: the visited map must flow down the whole chain rather
// than being copied per hop, or deep chains go quadratic.
func TestCycleDetection_DeepChainTerminatesWithEntry(t *testing.T) {
	findings := entryPointsFor(t, buildDeepChainDSL(1000))

	if len(findings) > 0 {
		t.Fatalf("deep computed-userset chain ending in a direct assignment should "+
			"have an entry point, got %d findings: %v", len(findings), findings)
	}
}

// TestCycleDetection_WideUnionTerminatesWithEntry verifies that a wide union of
// computed usersets, all reachable down to a concrete type, resolves with an
// entry point and that the sibling-isolating copies don't change the outcome.
func TestCycleDetection_WideUnionTerminatesWithEntry(t *testing.T) {
	findings := entryPointsFor(t, buildWideUnionDSL(1000))

	if len(findings) > 0 {
		t.Fatalf("wide union of resolvable members should have an entry point, "+
			"got %d findings: %v", len(findings), findings)
	}
}

// TestCycleDetection_SelfReferentialChainIsLoop verifies a relation that
// computes itself through a chain (a->b->c->a) is still detected as a loop with
// no entry point — the shared visited map must accumulate to see the back-edge.
func TestCycleDetection_SelfReferentialChainIsLoop(t *testing.T) {
	findings := entryPointsFor(t, `model
  schema 1.1
type user
type doc
  relations
    define a: b
    define b: c
    define c: a`)

	if len(findings) == 0 {
		t.Fatal("self-referential computed chain a->b->c->a should be reported as " +
			"having no entry point")
	}
}

// TestCycleDetection_UnionSiblingIsolation verifies a looping union member does
// not poison a sibling that has an entry point: `mixed: loops or direct` resolves
// with an entry point. This is what the per-branch copyVisited must preserve.
func TestCycleDetection_UnionSiblingIsolation(t *testing.T) {
	findings := entryPointsFor(t, `model
  schema 1.1
type user
type doc
  relations
    define direct: [user]
    define loops: selfref
    define selfref: loops
    define mixed: loops or direct`)

	// `loops` and `selfref` legitimately have no entry point and are reported.
	// `mixed` and `direct` must NOT be reported.
	for _, finding := range findings {
		if strings.Contains(finding.Message, "mixed") || strings.Contains(finding.Message, "`direct`") {
			t.Fatalf("relation with a resolvable union branch should have an entry "+
				"point, but got finding: %s", finding.Message)
		}
	}
}

// TestCycleDetection_DeepChainCountStable checks a chain whose base self-loops
// yields exactly one no-entry-point report per relation — neither suppressed
// nor duplicated by the visited-map sharing on the linear path.
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

	findings := entryPointsFor(t, b.String())

	// base + r0..r49 = depth+1 relations, all with no entry point.
	if len(findings) != depth+1 {
		t.Fatalf("expected %d no-entry-point findings, got %d: %v",
			depth+1, len(findings), findings)
	}
}
