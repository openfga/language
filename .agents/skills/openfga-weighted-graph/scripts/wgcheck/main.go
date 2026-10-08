// Command wgcheck reports why an OpenFGA model cannot be loaded by the weighted graph builder.
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strings"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/openfga/language/pkg/go/graph"
	"github.com/openfga/language/pkg/go/transformer"
	"github.com/openfga/openfga/pkg/typesystem"
	"google.golang.org/protobuf/proto"
)

type partialTTU struct {
	object, relation, tupleset, computed string
	missingOn                            []string
}

type relRef struct{ objectType, relation string }

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: wgcheck <model.fga | model.json | fga.mod>")
		os.Exit(2)
	}
	model, err := load(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "load: %v\n", err)
		os.Exit(2)
	}

	fmt.Println("== OpenFGA validation")
	if _, err := typesystem.NewAndValidate(context.Background(), model); err != nil {
		fmt.Printf("INVALID: %v\nFix this first: the weighted graph only matters for models OpenFGA accepts.\n", err)
		os.Exit(1)
	}
	fmt.Println("OK")

	fmt.Println("\n== Weighted graph build")
	err = build(model)
	if err == nil {
		fmt.Println("OK: the model is compatible with the weighted graph")
		return
	}
	fmt.Printf("FAIL: %v\n", err)

	fmt.Println("\n== Incompatibilities")
	working := model
	for _, p := range findPartialTTUs(model) {
		fmt.Printf("[A] %s#%s: '%s from %s' -> %s missing on: %s\n",
			p.object, p.relation, p.computed, p.tupleset, p.computed, strings.Join(p.missingOn, ", "))
		working = stubMissing(working, p)
	}

	// Patch each finding in a copy of the model so the next build surfaces the next one.
	for range 100 {
		err = build(working)
		if err == nil {
			fmt.Println("\nNo other incompatibilities: the model builds once the items above are fixed.")
			break
		}
		switch {
		case errors.Is(err, graph.ErrTupleCycle):
			implicated, ok := findConstrainedCycles(working)
			if !ok {
				fmt.Printf("[C] could not isolate the relations (%v); inspect recursive relations manually\n", err)
				os.Exit(1)
			}
			for _, r := range implicated {
				fmt.Printf("[C] %s#%s: 'and' / 'but not' is part of a recursive cycle\n", r.objectType, r.relation)
			}
			working = neutralizeWhere(working, func(r relRef) bool { return slices.Contains(implicated, r) })
		case strings.Contains(err.Error(), "not all paths return the same type"):
			m := operatorNodeRe.FindStringSubmatch(err.Error())
			if m == nil {
				fmt.Printf("[B] %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("[B] %s#%s: 'and' operands share no user type, so the relation is always false\n", m[1], m[2])
			ref := relRef{m[1], m[2]}
			working = neutralizeWhere(working, func(r relRef) bool { return r == ref })
		default:
			fmt.Printf("[?] %v\n", err)
			os.Exit(1)
		}
	}
	os.Exit(1)
}

var operatorNodeRe = regexp.MustCompile(`node ([^\s#]+)#([^\s:]+):`)

func load(path string) (*openfgav1.AuthorizationModel, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	switch {
	case filepath.Base(path) == "fga.mod":
		mod, err := transformer.TransformModFile(string(data))
		if err != nil {
			return nil, err
		}
		files := make([]transformer.ModuleFile, 0, len(mod.Contents.Value))
		for _, c := range mod.Contents.Value {
			b, err := os.ReadFile(filepath.Join(filepath.Dir(path), c.Value))
			if err != nil {
				return nil, err
			}
			files = append(files, transformer.ModuleFile{Name: c.Value, Contents: string(b)})
		}
		return transformer.TransformModuleFilesToModel(files, mod.Schema.Value)
	case strings.HasSuffix(path, ".json"):
		return transformer.LoadJSONStringToProto(string(data))
	default:
		return transformer.TransformDSLToProto(string(data))
	}
}

func build(model *openfgav1.AuthorizationModel) error {
	_, err := graph.NewWeightedAuthorizationModelGraphBuilder().Build(model)
	return err
}

func sortedTypes(model *openfgav1.AuthorizationModel) []*openfgav1.TypeDefinition {
	tds := append([]*openfgav1.TypeDefinition(nil), model.GetTypeDefinitions()...)
	sort.Slice(tds, func(i, j int) bool { return tds[i].GetType() < tds[j].GetType() })
	return tds
}

func sortedRelations(td *openfgav1.TypeDefinition) []string {
	names := make([]string, 0, len(td.GetRelations()))
	for n := range td.GetRelations() {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

func relationExists(model *openfgav1.AuthorizationModel, objectType, relation string) bool {
	for _, td := range model.GetTypeDefinitions() {
		if td.GetType() == objectType {
			_, ok := td.GetRelations()[relation]
			return ok
		}
	}
	return false
}

func walk(u *openfgav1.Userset, fn func(*openfgav1.Userset)) {
	if u == nil {
		return
	}
	fn(u)
	switch r := u.GetUserset().(type) {
	case *openfgav1.Userset_Union:
		for _, c := range r.Union.GetChild() {
			walk(c, fn)
		}
	case *openfgav1.Userset_Intersection:
		for _, c := range r.Intersection.GetChild() {
			walk(c, fn)
		}
	case *openfgav1.Userset_Difference:
		walk(r.Difference.GetBase(), fn)
		walk(r.Difference.GetSubtract(), fn)
	}
}

func findPartialTTUs(model *openfgav1.AuthorizationModel) []partialTTU {
	var out []partialTTU
	for _, td := range sortedTypes(model) {
		for _, rel := range sortedRelations(td) {
			walk(td.GetRelations()[rel], func(u *openfgav1.Userset) {
				ttu := u.GetTupleToUserset()
				if ttu == nil {
					return
				}
				tupleset := ttu.GetTupleset().GetRelation()
				computed := ttu.GetComputedUserset().GetRelation()
				var missing []string
				seen := map[string]bool{}
				for _, ref := range td.GetMetadata().GetRelations()[tupleset].GetDirectlyRelatedUserTypes() {
					t := ref.GetType()
					if !seen[t] && !relationExists(model, t, computed) {
						missing = append(missing, t)
					}
					seen[t] = true
				}
				if len(missing) > 0 {
					out = append(out, partialTTU{td.GetType(), rel, tupleset, computed, missing})
				}
			})
		}
	}
	return out
}

// stubMissing adds `define <rel>: [<type>]` placeholders so later errors can surface.
func stubMissing(model *openfgav1.AuthorizationModel, p partialTTU) *openfgav1.AuthorizationModel {
	clone := proto.Clone(model).(*openfgav1.AuthorizationModel)
	for _, t := range p.missingOn {
		for _, td := range clone.GetTypeDefinitions() {
			if td.GetType() != t {
				continue
			}
			if td.Relations == nil {
				td.Relations = map[string]*openfgav1.Userset{}
			}
			if td.Metadata == nil {
				td.Metadata = &openfgav1.Metadata{}
			}
			if td.Metadata.Relations == nil {
				td.Metadata.Relations = map[string]*openfgav1.RelationMetadata{}
			}
			td.Relations[p.computed] = &openfgav1.Userset{Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}}
			td.Metadata.Relations[p.computed] = &openfgav1.RelationMetadata{
				DirectlyRelatedUserTypes: []*openfgav1.RelationReference{{Type: t}},
			}
		}
	}
	return clone
}

func hasConstraint(u *openfgav1.Userset) bool {
	found := false
	walk(u, func(n *openfgav1.Userset) {
		if n.GetIntersection() != nil || n.GetDifference() != nil {
			found = true
		}
	})
	return found
}

func union(children ...*openfgav1.Userset) *openfgav1.Userset {
	return &openfgav1.Userset{Userset: &openfgav1.Userset_Union{Union: &openfgav1.Usersets{Child: children}}}
}

// neutralize turns 'and' / 'but not' into 'or', keeping every edge but removing the constraint.
func neutralize(u *openfgav1.Userset) *openfgav1.Userset {
	switch r := u.GetUserset().(type) {
	case *openfgav1.Userset_Union:
		for i, c := range r.Union.GetChild() {
			r.Union.Child[i] = neutralize(c)
		}
	case *openfgav1.Userset_Intersection:
		children := make([]*openfgav1.Userset, 0, len(r.Intersection.GetChild()))
		for _, c := range r.Intersection.GetChild() {
			children = append(children, neutralize(c))
		}
		return union(children...)
	case *openfgav1.Userset_Difference:
		return union(neutralize(r.Difference.GetBase()), neutralize(r.Difference.GetSubtract()))
	}
	return u
}

func neutralizeWhere(model *openfgav1.AuthorizationModel, match func(relRef) bool) *openfgav1.AuthorizationModel {
	clone := proto.Clone(model).(*openfgav1.AuthorizationModel)
	for _, td := range clone.GetTypeDefinitions() {
		for name, rw := range td.GetRelations() {
			if match(relRef{td.GetType(), name}) {
				td.Relations[name] = neutralize(rw)
			}
		}
	}
	return clone
}

// findConstrainedCycles keeps one constrained relation at a time and reports those that reintroduce the cycle error.
func findConstrainedCycles(model *openfgav1.AuthorizationModel) ([]relRef, bool) {
	if errors.Is(build(neutralizeWhere(model, func(relRef) bool { return true })), graph.ErrTupleCycle) {
		return nil, false
	}
	var out []relRef
	for _, td := range sortedTypes(model) {
		for _, rel := range sortedRelations(td) {
			if !hasConstraint(td.GetRelations()[rel]) {
				continue
			}
			ref := relRef{td.GetType(), rel}
			if errors.Is(build(neutralizeWhere(model, func(r relRef) bool { return r != ref })), graph.ErrTupleCycle) {
				out = append(out, ref)
			}
		}
	}
	return out, len(out) > 0
}
