package validation

import (
	"fmt"
	"strings"
	"testing"

	"github.com/openfga/language/pkg/go/transformer"
)

// buildDeepChainDSL builds one type with a long linear chain of computed-userset
// relations r0 -> r1 -> ... -> base -> [user], each relation depending on the
// one before it.
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

// buildWideUnionDSL builds a relation `wide` that is a union of width computed
// usersets. Every member but the last computes itself and has no entry point;
// the last resolves through base, so the union reaches it only after walking
// every other member.
func buildWideUnionDSL(width int) string {
	var b strings.Builder
	b.WriteString("model\n  schema 1.1\ntype user\n")
	b.WriteString("type doc\n  relations\n")
	b.WriteString("    define base: [user]\n")
	members := make([]string, 0, width)
	for i := 0; i < width; i++ {
		name := fmt.Sprintf("m%d", i)
		target := name
		if i == width-1 {
			target = "base"
		}
		fmt.Fprintf(&b, "    define %s: %s\n", name, target)
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

// TestCycleDetection_DeepChainTerminatesWithEntry checks that a chain of a
// thousand computed usersets ending in a direct assignment resolves with an
// entry point.
func TestCycleDetection_DeepChainTerminatesWithEntry(t *testing.T) {
	findings := entryPointsFor(t, buildDeepChainDSL(1000))

	if len(findings) > 0 {
		t.Fatalf("deep computed-userset chain ending in a direct assignment should "+
			"have an entry point, got %d findings: %v", len(findings), findings)
	}
}

// TestCycleDetection_WideUnionTerminatesWithEntry checks that a union keeps
// looking past members with no entry point: `wide` resolves through its last
// member, after 999 that loop.
func TestCycleDetection_WideUnionTerminatesWithEntry(t *testing.T) {
	const width = 1000
	findings := entryPointsFor(t, buildWideUnionDSL(width))

	for _, finding := range findings {
		if finding.Metadata.Relation == "wide" {
			t.Fatalf("union with a resolvable last member should have an entry "+
				"point, but got finding: %s", finding.Message)
		}
	}

	// Each looping member is reported on its own.
	if len(findings) != width-1 {
		t.Fatalf("expected %d no-entry-point findings, got %d", width-1, len(findings))
	}
}

// TestCycleDetection_SelfReferentialChainIsLoop verifies a relation that
// computes itself through a chain (a->b->c->a) is still detected as a loop with
// no entry point.
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

// TestCycleDetection_UnionSiblingIsolation checks that a looping union member
// does not stop the union from finding a sibling's entry point: `mixed: loops or
// direct` resolves through direct.
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

// TestCycleDetection_IntersectionSiblingIsolation checks that one intersection
// operand does not hide a relation from the next: `a` and `b` both reach `x`
// through a type restriction, and `b` must still find the entry point there
// after `a` has.
func TestCycleDetection_IntersectionSiblingIsolation(t *testing.T) {
	findings := entryPointsFor(t, `model
  schema 1.1
type user
type doc
  relations
    define x: [user]
    define a: [doc#x]
    define b: [doc#x]
    define r: a and b`)

	for _, finding := range findings {
		t.Errorf("unexpected finding: %s", finding.Message)
	}
}

// ladderDepth is the number of levels buildLadderDSL stacks on a0 and b0, and
// so the number of times walking every path through a ladder doubles.
const ladderDepth = 40

// buildLadderDSL builds one type with relations a0 and b0, then for each level
// i up to ladderDepth a relation ai and a relation bi that both combine a(i-1)
// and b(i-1) with operator. Every relation above the first level shares both of
// the relations below it.
func buildLadderDSL(operator, a0, b0 string) string {
	var b strings.Builder
	b.WriteString("model\n  schema 1.1\ntype user\n")
	b.WriteString("type doc\n  relations\n")
	fmt.Fprintf(&b, "    define a0: %s\n    define b0: %s\n", a0, b0)
	for i := 1; i <= ladderDepth; i++ {
		fmt.Fprintf(&b, "    define a%d: a%d %s b%d\n", i, i-1, operator, i-1)
		fmt.Fprintf(&b, "    define b%d: a%d %s b%d\n", i, i-1, operator, i-1)
	}
	return b.String()
}

// TestCycleDetection_SharedDependencies checks models whose relations share
// dependencies, which walking every path through them would take exponential
// time to resolve: each level of these ladders doubles the number of paths.
func TestCycleDetection_SharedDependencies(t *testing.T) {
	tests := []struct {
		name     string
		dsl      string
		findings int
	}{
		{
			name: "intersections",
			dsl:  buildLadderDSL("and", "[user]", "[user]"),
		},
		{
			name: "differences",
			dsl:  buildLadderDSL("but not", "[user]", "[user]"),
		},
		{
			name:     "unions over a loop",
			dsl:      buildLadderDSL("or", "b0", "a0"),
			findings: 2*ladderDepth + 2,
		},
		{
			name:     "unions around a cycle",
			dsl:      buildLadderDSL("or", fmt.Sprintf("a%d", ladderDepth), "a0"),
			findings: 2*ladderDepth + 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			findings := entryPointsFor(t, test.dsl)

			if len(findings) != test.findings {
				t.Fatalf("expected %d findings, got %d: %v", test.findings, len(findings), findings)
			}

			for _, finding := range findings {
				if !strings.HasSuffix(finding.Message, "(potential loop).") {
					t.Errorf("expected a potential loop, got: %s", finding.Message)
				}
			}
		})
	}
}
