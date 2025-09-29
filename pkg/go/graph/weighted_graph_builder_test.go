package graph

import (
	"testing"

	"github.com/stretchr/testify/require"

	language "github.com/openfga/language/pkg/go/transformer"
)

func TestCompleteWeightedGraph(t *testing.T) {
	t.Parallel()
	model := `
	model
  schema 1.1
type user
type employee

type directs-user
  relations
    define direct: [user]
    define direct_cond: [user with xcond]
    define direct_wild: [user:*]
    define direct_wild_cond: [user:* with xcond]
    define direct_and_direct_cond: [user, user with xcond, employee]
    define direct_and_direct_wild: [user, user:*, employee:*]
    define direct_and_direct_wild_cond: [user, user:* with xcond]
    define direct_cond_and_direct_wild: [user with xcond, user:*]
    define direct_cond_and_direct_wild_cond: [user with xcond, user:* with xcond]
    define direct_wildcard_and_direct_wildcard_cond: [user:*, user:* with xcond]
    define computed: direct
    define computed_cond: direct_cond
    define computed_wild: direct_wild
    define computed_wild_cond: direct_wild_cond
    define computed_computed: computed
    define computed_computed_computed: computed_computed
    define or_computed: computed or computed_cond or direct_wild
    define and_computed: computed_cond and computed_wild
    define butnot_computed: computed_wild_cond but not computed_computed
    define tuple_cycle2: [user, usersets-user#tuple_cycle2, employee]  
    define tuple_cycle3: [user, complexity3#cycle_nested]
    define compute_tuple_cycle3: tuple_cycle3
type directs-employee
  relations
    define direct: [employee]
    define computed: direct
    define direct_cond: [employee with xcond]
    define direct_wild: [employee:*]
    define direct_wild_cond: [employee:* with xcond]
type usersets-user
  relations
    define userset: [directs-user#direct, directs-employee#direct]
    define userset_to_computed: [directs-user#computed, directs-employee#computed]
    define userset_to_computed_cond: [directs-user#computed_cond, directs-employee#direct_cond]
    define userset_to_computed_wild: [directs-user#computed_wild, directs-employee#direct_wild]
    define userset_to_computed_wild_cond: [directs-user#direct_wild_cond, directs-employee#direct_wild_cond]
    define userset_cond: [directs-user#direct with xcond]
    define userset_cond_to_computed: [directs-user#computed with xcond]
    define userset_cond_to_computed_cond: [directs-user#computed_cond with xcond]
    define userset_cond_to_computed_wild: [directs-user#computed_wild with xcond]
    define userset_cond_to_computed_wild_cond: [directs-user#computed_wild_cond with xcond]
    define userset_to_or_computed: [directs-user#or_computed]
    define userset_to_butnot_computed: [directs-user#butnot_computed]
    define userset_to_and_computed:[directs-user#and_computed]
    define userset_recursive: [user, usersets-user#userset_recursive]
    define or_userset: userset or userset_to_computed_cond
    define and_userset: userset_to_computed_cond and userset_to_computed_wild
    define butnot_userset: userset_cond_to_computed_wild but not userset_cond
    define nested_or_userset: userset_to_or_computed or userset_to_butnot_computed
    define nested_and_userset: userset_to_and_computed and userset_to_or_computed
    define ttu_direct_userset: [ttus#direct_pa_direct_ch]
    define ttu_direct_cond_userset: [ttus#direct_cond_pa_direct_ch]
    define ttu_or_direct_userset: [ttus#or_comp_from_direct_parent]
    define ttu_and_direct_userset: [ttus#and_comp_from_direct_parent]
    define tuple_cycle2: [ttus#tuple_cycle2]
    define tuple_cycle3: [directs-user#compute_tuple_cycle3]
type ttus
  relations
    define direct_parent: [directs-user]
    define mult_parent_types: [directs-user, directs-employee]
    define mult_parent_types_cond: [directs-user with xcond, directs-employee with xcond]
    define direct_cond_parent: [directs-user with xcond]
    define userset_parent: [usersets-user]
    define userset_cond_parent: [usersets-user with xcond]
    define tuple_cycle2: tuple_cycle2 from direct_parent
    define tuple_cycle3: tuple_cycle3 from userset_parent
    define direct_pa_direct_ch: direct from mult_parent_types
    define direct_cond_pa_direct_ch: direct from mult_parent_types_cond
    define or_comp_from_direct_parent: or_computed from direct_parent
    define and_comp_from_direct_parent: and_computed from direct_cond_parent
    define butnot_comp_from_direct_parent: butnot_computed from direct_cond_parent
    define userset_pa_userset_ch: userset from userset_parent
    define userset_pa_userset_comp_ch: userset_to_computed from userset_parent
    define userset_pa_userset_comp_cond_ch: userset_to_computed_cond from userset_parent
    define userset_pa_userset_comp_wild_ch: userset_to_computed_wild from userset_parent
    define userset_pa_userset_comp_wild_cond_ch: userset_to_computed_wild_cond from userset_parent
    define userset_cond_userset_ch: userset from userset_cond_parent
    define userset_cond_userset_comp_ch: userset_to_computed from userset_cond_parent
    define userset_cond_userset_comp_cond_ch: userset_to_computed_cond from userset_cond_parent
    define userset_cond_userset_comp_wild_ch: userset_to_computed_wild from userset_cond_parent
    define userset_cond_userset_comp_wild_cond_ch: userset_to_computed_wild_cond from userset_cond_parent
    define or_ttu: direct_pa_direct_ch or direct_cond_pa_direct_ch
    define and_ttu: or_comp_from_direct_parent and direct_pa_direct_ch
    define nested_butnot_ttu: or_comp_from_direct_parent but not userset_pa_userset_comp_wild_ch
type complexity3
  relations
    define ttu_parent: [ttus]
    define userset_parent: [usersets-user]
    define ttu_userset_ttu: ttu_direct_userset from userset_parent
    define ttu_ttu_userset: userset_pa_userset_ch from ttu_parent
    define userset_ttu_userset: [ttus#userset_pa_userset_ch]
    define userset_userset_ttu: [usersets-user#ttu_direct_userset] 
    define compute_ttu_userset_ttu: ttu_userset_ttu
    define compute_userset_ttu_userset: userset_ttu_userset
    define or_compute_complex3: compute_ttu_userset_ttu or compute_userset_ttu_userset
    define and_nested_complex3: [ttus#and_ttu] and compute_ttu_userset_ttu 
    define cycle_nested: [ttus#tuple_cycle3]   
type complexity4
  relations
    define userset_ttu_userset_ttu: [complexity3#ttu_userset_ttu]
    define ttu_ttu_ttu_userset: ttu_ttu_userset from parent
    define userset_or_compute_complex3: [complexity3#or_compute_complex3]
    define ttu_and_nested_complex3: and_nested_complex3 from parent
    define or_complex4: userset_or_compute_complex3 or ttu_and_nested_complex3
    define parent: [complexity3]
condition xcond(x: string) {
  x == '1'
}`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Equal(t, 1, graph.nodes["directs-user#direct"].weights["user"])
	require.Equal(t, 1, graph.nodes["directs-user#direct_cond"].weights["user"])
	require.Equal(t, 1, graph.nodes["directs-user#direct_wild"].weights["user"])
	require.Equal(t, 1, graph.nodes["directs-user#direct_wild_cond"].weights["user"])

	require.Equal(t, 1, graph.nodes["directs-user#direct_and_direct_cond"].weights["user"])
	require.Equal(t, 1, graph.nodes["directs-user#direct_and_direct_cond"].weights["employee"])
	require.Equal(t, 1, graph.nodes["directs-user#direct_and_direct_wild"].weights["user"])
	require.Equal(t, 1, graph.nodes["directs-user#direct_and_direct_wild"].weights["employee"])
	require.Equal(t, 1, graph.nodes["directs-user#direct_and_direct_wild_cond"].weights["user"])
	require.Equal(t, 1, graph.nodes["directs-user#direct_cond_and_direct_wild"].weights["user"])
	require.Equal(t, 1, graph.nodes["directs-user#direct_cond_and_direct_wild_cond"].weights["user"])
	require.Equal(t, 1, graph.nodes["directs-user#direct_wildcard_and_direct_wildcard_cond"].weights["user"])

	require.Equal(t, 1, graph.nodes["directs-user#computed"].weights["user"])
	require.Equal(t, 1, graph.nodes["directs-user#computed_cond"].weights["user"])
	require.Equal(t, 1, graph.nodes["directs-user#computed_wild"].weights["user"])
	require.Equal(t, 1, graph.nodes["directs-user#computed_wild_cond"].weights["user"])
	require.Equal(t, 1, graph.nodes["directs-user#computed_computed"].weights["user"])
	require.Equal(t, 1, graph.nodes["directs-user#computed_computed_computed"].weights["user"])

	require.Equal(t, 1, graph.nodes["directs-user#or_computed"].weights["user"])
	require.Equal(t, 1, graph.nodes["directs-user#and_computed"].weights["user"])
	require.Equal(t, 1, graph.nodes["directs-user#butnot_computed"].weights["user"])

	require.Equal(t, Infinite, graph.nodes["directs-user#tuple_cycle2"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["directs-user#tuple_cycle2"].weights["employee"])
	require.Equal(t, Infinite, graph.nodes["directs-user#tuple_cycle3"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["directs-user#compute_tuple_cycle3"].weights["user"])

	require.Equal(t, 1, graph.nodes["directs-employee#direct"].weights["employee"])
	require.Equal(t, 1, graph.nodes["directs-employee#computed"].weights["employee"])
	require.Equal(t, 1, graph.nodes["directs-employee#direct_cond"].weights["employee"])
	require.Equal(t, 1, graph.nodes["directs-employee#direct_wild"].weights["employee"])
	require.Equal(t, 1, graph.nodes["directs-employee#direct_wild_cond"].weights["employee"])

	require.Equal(t, 2, graph.nodes["usersets-user#userset"].weights["user"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset"].weights["employee"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset_to_computed"].weights["employee"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset_to_computed"].weights["user"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset_to_computed_cond"].weights["employee"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset_to_computed_cond"].weights["user"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset_to_computed_wild"].weights["employee"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset_to_computed_wild"].weights["user"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset_to_computed_wild_cond"].weights["employee"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset_to_computed_wild_cond"].weights["user"])

	require.Equal(t, 2, graph.nodes["usersets-user#userset_cond"].weights["user"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset_cond_to_computed"].weights["user"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset_cond_to_computed_cond"].weights["user"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset_cond_to_computed_wild"].weights["user"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset_cond_to_computed_wild_cond"].weights["user"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset_to_or_computed"].weights["user"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset_to_butnot_computed"].weights["user"])
	require.Equal(t, 2, graph.nodes["usersets-user#userset_to_and_computed"].weights["user"])

	require.Equal(t, Infinite, graph.nodes["usersets-user#userset_recursive"].weights["user"])

	require.Equal(t, 2, graph.nodes["usersets-user#or_userset"].weights["user"])
	require.Equal(t, 2, graph.nodes["usersets-user#or_userset"].weights["employee"])
	require.Equal(t, 2, graph.nodes["usersets-user#and_userset"].weights["user"])
	require.Equal(t, 2, graph.nodes["usersets-user#and_userset"].weights["employee"])
	require.Equal(t, 2, graph.nodes["usersets-user#butnot_userset"].weights["user"])
	require.Equal(t, 2, graph.nodes["usersets-user#nested_or_userset"].weights["user"])
	require.Equal(t, 2, graph.nodes["usersets-user#nested_and_userset"].weights["user"])

	require.Equal(t, 3, graph.nodes["usersets-user#ttu_direct_userset"].weights["user"])
	require.Equal(t, 3, graph.nodes["usersets-user#ttu_direct_userset"].weights["employee"])
	require.Equal(t, 3, graph.nodes["usersets-user#ttu_direct_cond_userset"].weights["user"])
	require.Equal(t, 3, graph.nodes["usersets-user#ttu_direct_cond_userset"].weights["employee"])
	require.Equal(t, 3, graph.nodes["usersets-user#ttu_or_direct_userset"].weights["user"])
	require.Equal(t, 3, graph.nodes["usersets-user#ttu_and_direct_userset"].weights["user"])

	require.Equal(t, Infinite, graph.nodes["usersets-user#tuple_cycle2"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["usersets-user#tuple_cycle2"].weights["employee"])
	require.Equal(t, Infinite, graph.nodes["usersets-user#tuple_cycle3"].weights["user"])

	require.Equal(t, 1, graph.nodes["ttus#direct_parent"].weights["directs-user"])
	require.Equal(t, 1, graph.nodes["ttus#mult_parent_types"].weights["directs-user"])
	require.Equal(t, 1, graph.nodes["ttus#mult_parent_types"].weights["directs-employee"])
	require.Equal(t, 1, graph.nodes["ttus#mult_parent_types_cond"].weights["directs-user"])
	require.Equal(t, 1, graph.nodes["ttus#mult_parent_types_cond"].weights["directs-employee"])
	require.Equal(t, 1, graph.nodes["ttus#direct_cond_parent"].weights["directs-user"])

	require.Equal(t, 1, graph.nodes["ttus#userset_parent"].weights["usersets-user"])
	require.Equal(t, 1, graph.nodes["ttus#userset_cond_parent"].weights["usersets-user"])

	require.Equal(t, Infinite, graph.nodes["ttus#tuple_cycle2"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["ttus#tuple_cycle2"].weights["employee"])
	require.Equal(t, Infinite, graph.nodes["ttus#tuple_cycle3"].weights["user"])

	require.Equal(t, 2, graph.nodes["ttus#direct_pa_direct_ch"].weights["user"])
	require.Equal(t, 2, graph.nodes["ttus#direct_pa_direct_ch"].weights["employee"])
	require.Equal(t, 2, graph.nodes["ttus#direct_cond_pa_direct_ch"].weights["user"])
	require.Equal(t, 2, graph.nodes["ttus#direct_cond_pa_direct_ch"].weights["employee"])
	require.Equal(t, 2, graph.nodes["ttus#or_comp_from_direct_parent"].weights["user"])
	require.Equal(t, 2, graph.nodes["ttus#and_comp_from_direct_parent"].weights["user"])
	require.Equal(t, 2, graph.nodes["ttus#butnot_comp_from_direct_parent"].weights["user"])

	require.Equal(t, 3, graph.nodes["ttus#userset_pa_userset_ch"].weights["user"])
	require.Equal(t, 3, graph.nodes["ttus#userset_pa_userset_ch"].weights["employee"])
	require.Equal(t, 3, graph.nodes["ttus#userset_pa_userset_comp_ch"].weights["user"])
	require.Equal(t, 3, graph.nodes["ttus#userset_pa_userset_comp_ch"].weights["employee"])
	require.Equal(t, 3, graph.nodes["ttus#userset_pa_userset_comp_cond_ch"].weights["user"])
	require.Equal(t, 3, graph.nodes["ttus#userset_pa_userset_comp_cond_ch"].weights["employee"])
	require.Equal(t, 3, graph.nodes["ttus#userset_pa_userset_comp_wild_ch"].weights["user"])
	require.Equal(t, 3, graph.nodes["ttus#userset_pa_userset_comp_wild_ch"].weights["employee"])
	require.Equal(t, 3, graph.nodes["ttus#userset_pa_userset_comp_wild_cond_ch"].weights["employee"])
	require.Equal(t, 3, graph.nodes["ttus#userset_pa_userset_comp_wild_cond_ch"].weights["user"])

	require.Equal(t, 3, graph.nodes["ttus#userset_cond_userset_ch"].weights["user"])
	require.Equal(t, 3, graph.nodes["ttus#userset_cond_userset_ch"].weights["employee"])
	require.Equal(t, 3, graph.nodes["ttus#userset_cond_userset_comp_ch"].weights["user"])
	require.Equal(t, 3, graph.nodes["ttus#userset_cond_userset_comp_ch"].weights["employee"])
	require.Equal(t, 3, graph.nodes["ttus#userset_cond_userset_comp_cond_ch"].weights["user"])
	require.Equal(t, 3, graph.nodes["ttus#userset_cond_userset_comp_cond_ch"].weights["employee"])
	require.Equal(t, 3, graph.nodes["ttus#userset_cond_userset_comp_wild_ch"].weights["user"])
	require.Equal(t, 3, graph.nodes["ttus#userset_cond_userset_comp_wild_ch"].weights["employee"])
	require.Equal(t, 3, graph.nodes["ttus#userset_cond_userset_comp_wild_cond_ch"].weights["employee"])
	require.Equal(t, 3, graph.nodes["ttus#userset_cond_userset_comp_wild_cond_ch"].weights["user"])

	require.Equal(t, 2, graph.nodes["ttus#or_ttu"].weights["user"])
	require.Equal(t, 2, graph.nodes["ttus#or_ttu"].weights["employee"])
	require.Equal(t, 2, graph.nodes["ttus#and_ttu"].weights["user"])
	require.Equal(t, 3, graph.nodes["ttus#nested_butnot_ttu"].weights["user"])

	require.Equal(t, 1, graph.nodes["complexity3#ttu_parent"].weights["ttus"])
	require.Equal(t, 1, graph.nodes["complexity3#userset_parent"].weights["usersets-user"])
	require.Equal(t, 4, graph.nodes["complexity3#ttu_userset_ttu"].weights["user"])
	require.Equal(t, 4, graph.nodes["complexity3#ttu_userset_ttu"].weights["employee"])
	require.Equal(t, 4, graph.nodes["complexity3#ttu_ttu_userset"].weights["user"])
	require.Equal(t, 4, graph.nodes["complexity3#ttu_ttu_userset"].weights["employee"])
	require.Equal(t, 4, graph.nodes["complexity3#userset_ttu_userset"].weights["user"])
	require.Equal(t, 4, graph.nodes["complexity3#userset_ttu_userset"].weights["employee"])
	require.Equal(t, 4, graph.nodes["complexity3#userset_userset_ttu"].weights["user"])
	require.Equal(t, 4, graph.nodes["complexity3#userset_userset_ttu"].weights["employee"])
	require.Equal(t, 4, graph.nodes["complexity3#compute_ttu_userset_ttu"].weights["user"])
	require.Equal(t, 4, graph.nodes["complexity3#compute_ttu_userset_ttu"].weights["employee"])
	require.Equal(t, 4, graph.nodes["complexity3#compute_userset_ttu_userset"].weights["user"])
	require.Equal(t, 4, graph.nodes["complexity3#compute_userset_ttu_userset"].weights["employee"])
	require.Equal(t, 4, graph.nodes["complexity3#or_compute_complex3"].weights["user"])
	require.Equal(t, 4, graph.nodes["complexity3#or_compute_complex3"].weights["employee"])
	require.Equal(t, 4, graph.nodes["complexity3#and_nested_complex3"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["complexity3#cycle_nested"].weights["user"])

	require.Equal(t, 5, graph.nodes["complexity4#userset_ttu_userset_ttu"].weights["user"])
	require.Equal(t, 5, graph.nodes["complexity4#userset_ttu_userset_ttu"].weights["employee"])
	require.Equal(t, 1, graph.nodes["complexity4#parent"].weights["complexity3"])
	require.Equal(t, 5, graph.nodes["complexity4#ttu_ttu_ttu_userset"].weights["user"])
	require.Equal(t, 5, graph.nodes["complexity4#ttu_ttu_ttu_userset"].weights["employee"])
	require.Equal(t, 5, graph.nodes["complexity4#userset_or_compute_complex3"].weights["user"])
	require.Equal(t, 5, graph.nodes["complexity4#userset_or_compute_complex3"].weights["employee"])
	require.Equal(t, 5, graph.nodes["complexity4#ttu_and_nested_complex3"].weights["user"])
	require.Equal(t, 5, graph.nodes["complexity4#or_complex4"].weights["user"])

	require.Len(t, graph.nodes["directs-user#direct_wild"].wildcards, 1)
	require.Len(t, graph.nodes["directs-user#direct_wild_cond"].wildcards, 1)
	require.Len(t, graph.nodes["directs-user#direct_and_direct_wild"].wildcards, 2)
	require.Len(t, graph.nodes["directs-user#direct_and_direct_wild_cond"].wildcards, 1)
	require.Len(t, graph.nodes["directs-user#direct_cond_and_direct_wild"].wildcards, 1)
	require.Len(t, graph.nodes["directs-user#direct_cond_and_direct_wild_cond"].wildcards, 1)
	require.Len(t, graph.nodes["directs-user#direct_wildcard_and_direct_wildcard_cond"].wildcards, 1)
	require.Len(t, graph.nodes["directs-user#computed_wild"].wildcards, 1)
	require.Len(t, graph.nodes["directs-user#or_computed"].wildcards, 1)
	require.Len(t, graph.nodes["directs-user#and_computed"].wildcards, 1)
	require.Len(t, graph.nodes["directs-user#butnot_computed"].wildcards, 1)
	require.Empty(t, graph.nodes["directs-user#direct"].wildcards)
	require.Empty(t, graph.nodes["directs-user#direct_cond"].wildcards)
	require.Empty(t, graph.nodes["directs-user#direct_and_direct_cond"].wildcards)
	require.Empty(t, graph.nodes["directs-user#computed"].wildcards)
	require.Empty(t, graph.nodes["directs-user#computed_cond"].wildcards)
	require.Empty(t, graph.nodes["directs-user#computed_computed"].wildcards)
	require.Empty(t, graph.nodes["directs-user#computed_computed_computed"].wildcards)
	require.Empty(t, graph.nodes["directs-user#tuple_cycle2"].wildcards)
	require.Empty(t, graph.nodes["directs-user#tuple_cycle3"].wildcards)
	require.Empty(t, graph.nodes["directs-user#compute_tuple_cycle3"].wildcards)

	require.Empty(t, graph.nodes["directs-employee#direct"].wildcards)
	require.Empty(t, graph.nodes["directs-employee#computed"].wildcards)
	require.Empty(t, graph.nodes["directs-employee#direct_cond"].wildcards)
	require.Len(t, graph.nodes["directs-employee#direct_wild"].wildcards, 1)
	require.Len(t, graph.nodes["directs-employee#direct_wild_cond"].wildcards, 1)

	require.Empty(t, graph.nodes["usersets-user#userset"].wildcards)
	require.Empty(t, graph.nodes["usersets-user#userset_to_computed"].wildcards)
	require.Empty(t, graph.nodes["usersets-user#userset_to_computed_cond"].wildcards)
	require.Empty(t, graph.nodes["usersets-user#userset_cond"].wildcards)
	require.Empty(t, graph.nodes["usersets-user#userset_cond_to_computed"].wildcards)
	require.Empty(t, graph.nodes["usersets-user#userset_cond_to_computed_cond"].wildcards)
	require.Empty(t, graph.nodes["usersets-user#userset_recursive"].wildcards)
	require.Empty(t, graph.nodes["usersets-user#or_userset"].wildcards)
	require.Empty(t, graph.nodes["usersets-user#ttu_direct_userset"].wildcards)
	require.Empty(t, graph.nodes["usersets-user#ttu_direct_cond_userset"].wildcards)
	require.Empty(t, graph.nodes["usersets-user#tuple_cycle2"].wildcards)
	require.Empty(t, graph.nodes["usersets-user#tuple_cycle3"].wildcards)

	require.Len(t, graph.nodes["usersets-user#userset_to_computed_wild"].wildcards, 2)
	require.Len(t, graph.nodes["usersets-user#userset_to_computed_wild_cond"].wildcards, 2)
	require.Len(t, graph.nodes["usersets-user#userset_cond_to_computed_wild"].wildcards, 1)
	require.Len(t, graph.nodes["usersets-user#userset_cond_to_computed_wild_cond"].wildcards, 1)
	require.Len(t, graph.nodes["usersets-user#userset_to_or_computed"].wildcards, 1)
	require.Len(t, graph.nodes["usersets-user#userset_to_butnot_computed"].wildcards, 1)
	require.Len(t, graph.nodes["usersets-user#userset_to_and_computed"].wildcards, 1)
	require.Len(t, graph.nodes["usersets-user#and_userset"].wildcards, 2)
	require.Len(t, graph.nodes["usersets-user#butnot_userset"].wildcards, 1)
	require.Len(t, graph.nodes["usersets-user#nested_or_userset"].wildcards, 1)
	require.Len(t, graph.nodes["usersets-user#nested_and_userset"].wildcards, 1)
	require.Len(t, graph.nodes["usersets-user#ttu_or_direct_userset"].wildcards, 1)
	require.Len(t, graph.nodes["usersets-user#ttu_and_direct_userset"].wildcards, 1)

	require.Empty(t, graph.nodes["ttus#direct_parent"].wildcards)
	require.Empty(t, graph.nodes["ttus#mult_parent_types"].wildcards)
	require.Empty(t, graph.nodes["ttus#mult_parent_types_cond"].wildcards)
	require.Empty(t, graph.nodes["ttus#direct_cond_parent"].wildcards)
	require.Empty(t, graph.nodes["ttus#userset_parent"].wildcards)
	require.Empty(t, graph.nodes["ttus#userset_cond_parent"].wildcards)
	require.Empty(t, graph.nodes["ttus#tuple_cycle2"].wildcards)
	require.Empty(t, graph.nodes["ttus#tuple_cycle3"].wildcards)
	require.Empty(t, graph.nodes["ttus#direct_pa_direct_ch"].wildcards)
	require.Empty(t, graph.nodes["ttus#direct_cond_pa_direct_ch"].wildcards)
	require.Empty(t, graph.nodes["ttus#userset_pa_userset_ch"].wildcards)
	require.Empty(t, graph.nodes["ttus#userset_pa_userset_comp_ch"].wildcards)
	require.Empty(t, graph.nodes["ttus#userset_pa_userset_comp_cond_ch"].wildcards)
	require.Empty(t, graph.nodes["ttus#userset_cond_userset_ch"].wildcards)
	require.Empty(t, graph.nodes["ttus#userset_cond_userset_comp_ch"].wildcards)
	require.Empty(t, graph.nodes["ttus#userset_cond_userset_comp_cond_ch"].wildcards)
	require.Empty(t, graph.nodes["ttus#or_ttu"].wildcards)

	require.Len(t, graph.nodes["ttus#or_comp_from_direct_parent"].wildcards, 1)
	require.Len(t, graph.nodes["ttus#and_comp_from_direct_parent"].wildcards, 1)
	require.Len(t, graph.nodes["ttus#butnot_comp_from_direct_parent"].wildcards, 1)
	require.Len(t, graph.nodes["ttus#userset_pa_userset_comp_wild_ch"].wildcards, 2)
	require.Len(t, graph.nodes["ttus#userset_pa_userset_comp_wild_cond_ch"].wildcards, 2)
	require.Len(t, graph.nodes["ttus#userset_cond_userset_comp_wild_ch"].wildcards, 2)
	require.Len(t, graph.nodes["ttus#userset_cond_userset_comp_wild_cond_ch"].wildcards, 2)
	require.Len(t, graph.nodes["ttus#and_ttu"].wildcards, 1)
	require.Len(t, graph.nodes["ttus#nested_butnot_ttu"].wildcards, 2)

	require.Empty(t, graph.nodes["complexity3#ttu_parent"].wildcards)
	require.Empty(t, graph.nodes["complexity3#userset_parent"].wildcards)
	require.Empty(t, graph.nodes["complexity3#ttu_userset_ttu"].wildcards)
	require.Empty(t, graph.nodes["complexity3#userset_ttu_userset"].wildcards)
	require.Empty(t, graph.nodes["complexity3#userset_userset_ttu"].wildcards)
	require.Empty(t, graph.nodes["complexity3#compute_ttu_userset_ttu"].wildcards)
	require.Empty(t, graph.nodes["complexity3#compute_userset_ttu_userset"].wildcards)
	require.Empty(t, graph.nodes["complexity3#or_compute_complex3"].wildcards)
	require.Empty(t, graph.nodes["complexity3#cycle_nested"].wildcards)
	require.Len(t, graph.nodes["complexity3#and_nested_complex3"].wildcards, 1)

	require.Empty(t, graph.nodes["complexity4#userset_ttu_userset_ttu"].wildcards)
	require.Empty(t, graph.nodes["complexity4#parent"].wildcards)
	require.Empty(t, graph.nodes["complexity4#ttu_ttu_ttu_userset"].wildcards)
	require.Empty(t, graph.nodes["complexity4#userset_or_compute_complex3"].wildcards)
	require.Len(t, graph.nodes["complexity4#ttu_and_nested_complex3"].wildcards, 1)
	require.Len(t, graph.nodes["complexity4#or_complex4"].wildcards, 1)
}

func TestInvalidGraphNoRelationDefined(t *testing.T) {
	t.Parallel()
	model := `
	model
  		schema 1.1
		type user
		type group
			relations
				define member: [user]
		type folder
			relations
				define viewer: [group#computed_member]
`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	_, err := wgb.Build(authorizationModel)
	require.ErrorIs(t, err, ErrInvalidModel)
}

func TestInvalidGraphTupleCycleWithExclusion(t *testing.T) {
	t.Parallel()
	model := `
	model
  		schema 1.1
		type user
		type document
			relations
				define viewer: [user] but not restricted
				define restricted: [user, document#viewer]
`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	_, err := wgb.Build(authorizationModel)
	require.ErrorIs(t, err, ErrContrainstTupleCycle)
}

func TestInvalidGraphTupleCycleWithExclusionCase2(t *testing.T) {
	t.Parallel()
	model := `
	model
  		schema 1.1
		type user
		type document
			relations
				define viewer: [user] but not restricteda
				define restricteda: restrictedb
				define restrictedb: restrictedc
				define restrictedc: [user, document#viewer]
`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	_, err := wgb.Build(authorizationModel)
	require.ErrorIs(t, err, ErrContrainstTupleCycle)
}

func TestInvalidGraphModelCycle(t *testing.T) {
	t.Parallel()
	model := `
	model
  		schema 1.1
		type user
		type document
			relations
				define x: y
				define y: x
`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	_, err := wgb.Build(authorizationModel)
	require.ErrorIs(t, err, ErrModelCycle)
}

func TestValidGraphModel(t *testing.T) {
	t.Parallel()
	model := `
	model
  		schema 1.1
		type user
        type role
            relations
                define assignee: [user]
        type permission
            relations
                define assignee: [role#assignee]
        type job
            relations
                define can_read: assignee from permission
                define permission: [permission]
                define cannot_read: [user] but not can_read
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)
	require.Equal(t, 3, graph.nodes["job#can_read"].weights["user"])
	butnotNode := graph.edges["job#cannot_read"][0].to
	edges := graph.edges[butnotNode.uniqueLabel]
	require.Equal(t, "job#cannot_read", edges[0].GetRelationDefinition())
}

func TestCompleteWeightedGraphWithExclusion(t *testing.T) {
	t.Parallel()
	t.Run("B_appears_in_A_infinite", func(t *testing.T) {
		model := `
	model
		schema 1.1
		type user
		type other
		type employee
		type group
			relations
				define parent: [group]
				define admin: [user, employee] or admin from parent
				define banned: [user] or banned from parent
				define allowed: ([user, employee, other] or admin) but not banned
		type role
			relations
				define owner: [group]
				define allowed: allowed from owner
`
		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		require.NoError(t, err)
		require.Len(t, graph.nodes["role#allowed"].weights, 3)
		require.Equal(t, Infinite, graph.nodes["role#allowed"].weights["user"])
		require.Equal(t, 2, graph.nodes["role#allowed"].weights["other"])
		require.Equal(t, Infinite, graph.nodes["role#allowed"].weights["employee"])

		butnotNode := graph.edges["group#allowed"][0].to
		orNode := graph.edges[butnotNode.uniqueLabel][0].to
		orEdges := graph.edges[orNode.uniqueLabel]
		require.Equal(t, "group#allowed", orEdges[0].GetRelationDefinition())
		require.Equal(t, DirectLogicalEdge, orEdges[0].GetEdgeType())
		logicalUsersetNode := orEdges[0].to
		require.Equal(t, LogicalDirectGrouping, logicalUsersetNode.GetNodeType())
		directEdges := graph.edges[logicalUsersetNode.uniqueLabel]
		require.Equal(t, DirectEdge, directEdges[0].GetEdgeType())
		require.Equal(t, "group#allowed", directEdges[0].GetRelationDefinition())
		require.Equal(t, DirectEdge, directEdges[1].GetEdgeType())
		require.Equal(t, "group#allowed", directEdges[1].GetRelationDefinition())
		require.Equal(t, DirectEdge, directEdges[2].GetEdgeType())
		require.Equal(t, "group#allowed", directEdges[2].GetRelationDefinition())
	})
	t.Run("B_appears_in_A_finite", func(t *testing.T) {
		model := `
	model
		schema 1.1
		type user
		type other
		type employee
		type group
			relations
				define parent: [group]
				define admin: [user, employee] or admin from parent
				define banned: [other]
				define allowed: ([user, employee, other] or admin) but not banned
		type role
			relations
				define owner: [group]
				define allowed: allowed from owner
`
		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		require.NoError(t, err)
		require.Len(t, graph.nodes["role#allowed"].weights, 3)
		require.Equal(t, Infinite, graph.nodes["role#allowed"].weights["user"])
		require.Equal(t, 2, graph.nodes["role#allowed"].weights["other"])
		require.Equal(t, Infinite, graph.nodes["role#allowed"].weights["employee"])

		butnotNode := graph.edges["group#allowed"][0].to
		orNode := graph.edges[butnotNode.uniqueLabel][0].to
		orEdges := graph.edges[orNode.uniqueLabel]
		require.Equal(t, "group#allowed", orEdges[0].GetRelationDefinition())
		require.Equal(t, DirectLogicalEdge, orEdges[0].GetEdgeType())
		logicalUsersetNode := orEdges[0].to
		require.Equal(t, LogicalDirectGrouping, logicalUsersetNode.GetNodeType())
		directEdges := graph.edges[logicalUsersetNode.uniqueLabel]
		require.Equal(t, DirectEdge, directEdges[0].GetEdgeType())
		require.Equal(t, "group#allowed", directEdges[0].GetRelationDefinition())
		require.Equal(t, DirectEdge, directEdges[1].GetEdgeType())
		require.Equal(t, "group#allowed", directEdges[1].GetRelationDefinition())
		require.Equal(t, DirectEdge, directEdges[2].GetEdgeType())
		require.Equal(t, "group#allowed", directEdges[2].GetRelationDefinition())
	})
	t.Run("B_not_appear_in_A", func(t *testing.T) {
		model := `
	model
		schema 1.1
		type user
		type other
		type employee
		type group
			relations
				define parent: [group]
				define admin: [user, employee] or admin from parent
				define banned: [other] or banned from parent
				define allowed: ([user, employee] or admin) but not banned
		type role
			relations
				define owner: [group]
				define allowed: allowed from owner
`
		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		require.NoError(t, err)
		require.Len(t, graph.nodes["role#allowed"].weights, 2)
		require.Equal(t, Infinite, graph.nodes["role#allowed"].weights["user"])
		require.Equal(t, Infinite, graph.nodes["role#allowed"].weights["employee"])
		_, found := graph.nodes["role#allowed"].weights["other"]
		require.False(t, found)

		butnotNode := graph.edges["group#allowed"][0].to
		orNode := graph.edges[butnotNode.uniqueLabel][0].to
		orEdges := graph.edges[orNode.uniqueLabel]
		require.Equal(t, "group#allowed", orEdges[0].GetRelationDefinition())
		require.Equal(t, DirectLogicalEdge, orEdges[0].GetEdgeType())
		logicalUsersetNode := orEdges[0].to
		require.Equal(t, LogicalDirectGrouping, logicalUsersetNode.GetNodeType())
		directEdges := graph.edges[logicalUsersetNode.uniqueLabel]
		require.Equal(t, DirectEdge, directEdges[0].GetEdgeType())
		require.Equal(t, "group#allowed", directEdges[0].GetRelationDefinition())
		require.Equal(t, DirectEdge, directEdges[1].GetEdgeType())
		require.Equal(t, "group#allowed", directEdges[1].GetRelationDefinition())
	})
}

func TestValidConditionalGraphModel(t *testing.T) {
	t.Parallel()
	model := `
	model
  		schema 1.1
		type user
        type role
            relations
                define assignee: [user]
        type permission
            relations
                define assignee: [role#assignee, role#assignee with condX]
				define member: [user, permission#member, permission#member with condX]
        type job
            relations
                define can_read: assignee from permission
				define can_view: [user] or can_view from owner
				define owner: [job, job with condX]
                define permission: [permission, permission with condX]
		condition condX (x:int) {
					x > 0
				}
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)
	require.Len(t, graph.nodes, 13)
	require.Len(t, graph.edges, 9)
	edges, _ := graph.GetEdgesFromNode(graph.nodes["permission#assignee"])
	require.Len(t, edges, 1)
	conditions := edges[0].conditions
	require.Empty(t, edges[0].tuplesetRelation)
	require.Len(t, conditions, 2)
	require.Equal(t, "none", conditions[0])
	require.Equal(t, "condX", conditions[1])

	edges, _ = graph.GetEdgesFromNode(graph.nodes["job#can_read"])
	require.Len(t, edges, 1)
	conditions = edges[0].conditions
	require.Len(t, conditions, 1)
	require.Equal(t, "none", conditions[0])
	require.Equal(t, "job#permission", edges[0].tuplesetRelation)
	edges, _ = graph.GetEdgesFromNode(graph.nodes["job#permission"])
	require.Len(t, edges, 1)
	conditions = edges[0].conditions
	require.Len(t, conditions, 2)
	require.Equal(t, "none", conditions[0])
	require.Equal(t, "condX", conditions[1])
	require.Equal(t, "", edges[0].tuplesetRelation)
	edges, _ = graph.GetEdgesFromNode(graph.nodes["role#assignee"])
	require.Len(t, edges, 1)
	conditions = edges[0].conditions
	require.Len(t, conditions, 1)
	require.Equal(t, "none", conditions[0])
	require.Equal(t, "", edges[0].tuplesetRelation)
	edges, _ = graph.GetEdgesFromNode(graph.nodes["permission#member"])
	require.Len(t, edges, 2)
	var recursiveEdge *WeightedAuthorizationModelEdge
	var userEdge *WeightedAuthorizationModelEdge
	if edges[0].weights["user"] == Infinite {
		recursiveEdge = edges[0]
		userEdge = edges[1]
	} else {
		recursiveEdge = edges[1]
		userEdge = edges[0]
	}
	conditions = recursiveEdge.conditions
	require.Len(t, conditions, 2)
	require.Equal(t, "none", conditions[0])
	require.Equal(t, "condX", conditions[1])
	require.Equal(t, "", recursiveEdge.tuplesetRelation)
	conditions = userEdge.conditions
	require.Len(t, conditions, 1)
	require.Equal(t, "none", conditions[0])
	require.Equal(t, "", userEdge.tuplesetRelation)
	edges, _ = graph.GetEdgesFromNode(graph.nodes["job#owner"])
	require.Len(t, edges, 1)
	conditions = edges[0].conditions
	require.Len(t, conditions, 2)
	require.Equal(t, "none", conditions[0])
	require.Equal(t, "condX", conditions[1])
	require.Equal(t, "", edges[0].tuplesetRelation)
	edges, _ = graph.GetEdgesFromNode(graph.nodes["job#can_view"])
	require.Len(t, edges, 1)
	conditions = edges[0].conditions
	require.Len(t, conditions, 1)
	require.Equal(t, "none", conditions[0])
	require.Equal(t, "", edges[0].tuplesetRelation)
	edges, _ = graph.GetEdgesFromNode(edges[0].to) // OR node
	require.Len(t, edges, 2)
	if edges[0].weights["user"] == Infinite {
		recursiveEdge = edges[0]
		userEdge = edges[1]
	} else {
		recursiveEdge = edges[1]
		userEdge = edges[0]
	}
	conditions = recursiveEdge.conditions
	require.Len(t, conditions, 1)
	require.Equal(t, "none", conditions[0])
	require.Equal(t, "job#owner", recursiveEdge.tuplesetRelation)
	conditions = userEdge.conditions
	require.Len(t, conditions, 1)
	require.Equal(t, "", userEdge.tuplesetRelation)
	require.Equal(t, "none", conditions[0])

	require.Equal(t, 2, graph.nodes["permission#assignee"].weights["user"])
	require.Equal(t, 3, graph.nodes["job#can_read"].weights["user"])
	require.Equal(t, 1, graph.nodes["role#assignee"].weights["user"])
	require.Equal(t, 1, graph.nodes["job#permission"].weights["permission"])
	require.Equal(t, Infinite, graph.nodes["permission#member"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["job#can_view"].weights["user"])
	require.Equal(t, 1, graph.nodes["job#owner"].weights["job"])

	assigneeEdges := graph.edges["permission#assignee"]
	require.Len(t, assigneeEdges, 1)
	require.Equal(t, "permission#assignee", assigneeEdges[0].GetRelationDefinition())
	memberEdges := graph.edges["permission#member"]
	require.Len(t, memberEdges, 2)
	require.Equal(t, "permission#member", memberEdges[0].GetRelationDefinition())
	require.Equal(t, "permission#member", memberEdges[1].GetRelationDefinition())
}

func TestGraphConstructionOrderedExclusion(t *testing.T) {
	t.Parallel()
	model := `
	model
  		schema 1.1
		type user
		type employee
        type role
            relations
                define assignee: [user]
				define cannot_read: [user]
        type permission
            relations
                define assignee: [employee]
				define cannot_read: cannot_read from role
				define role: [role]
        type job
            relations
                define can_read: [user, employee, role#assignee, permission#assignee] but not cannot_read
                define cannot_read: cannot_read from permission
				define permission: [permission]
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 15)
	require.Len(t, graph.edges, 10)
	exclusionNodeID := graph.edges["job#can_read"][0].to.uniqueLabel
	exclusionEdges := graph.edges[exclusionNodeID]
	require.Len(t, exclusionEdges, 2)
	cannotreadID := exclusionEdges[1].to.uniqueLabel
	require.Equal(t, "job#cannot_read", cannotreadID)
	exclusionNodeWeights := graph.nodes[exclusionNodeID].weights
	jobcannotReadWeight := graph.nodes[cannotreadID].weights
	require.Len(t, exclusionNodeWeights, 2)
	require.Equal(t, 3, exclusionNodeWeights["user"])
	require.Equal(t, 2, exclusionNodeWeights["employee"])
	require.Len(t, jobcannotReadWeight, 1)
	require.Equal(t, 3, jobcannotReadWeight["user"])

	logicalDirectGroupingNode := exclusionEdges[0].to
	require.Equal(t, LogicalDirectGrouping, logicalDirectGroupingNode.nodeType)
	directEdges := graph.edges[logicalDirectGroupingNode.uniqueLabel]
	require.Len(t, directEdges, 4)
	require.Len(t, logicalDirectGroupingNode.weights, 2)
	require.Equal(t, 2, logicalDirectGroupingNode.weights["user"])
	require.Equal(t, 2, logicalDirectGroupingNode.weights["employee"])

	require.Equal(t, "job#can_read", directEdges[0].GetRelationDefinition())
	require.Equal(t, DirectEdge, directEdges[0].GetEdgeType())
	require.Len(t, directEdges[0].weights, 1)
	require.Equal(t, 1, directEdges[0].weights["user"])

	require.Equal(t, "job#can_read", directEdges[1].GetRelationDefinition())
	require.Equal(t, DirectEdge, directEdges[1].GetEdgeType())
	require.Len(t, directEdges[1].weights, 1)
	require.Equal(t, 1, directEdges[1].weights["employee"])

	require.Equal(t, "job#can_read", directEdges[2].GetRelationDefinition())
	require.Equal(t, DirectEdge, directEdges[2].GetEdgeType())
	require.Len(t, directEdges[2].weights, 1)
	require.Equal(t, 2, directEdges[2].weights["user"])

	require.Equal(t, "job#can_read", directEdges[3].GetRelationDefinition())
	require.Equal(t, DirectEdge, directEdges[3].GetEdgeType())
	require.Len(t, directEdges[3].weights, 1)
	require.Equal(t, 2, directEdges[3].weights["employee"])
}

func TestGraphConstructionOrderedExclusionWithLogicalTTU(t *testing.T) {
	t.Parallel()
	model := `
	model
  		schema 1.1
		type user
        type role
            relations
				define cannot_read: [user]
		type role_assignment
            relations
				define cannot_read: [user]
        type permission
            relations
				define cannot_read: cannot_read from role or assignee
				define role: [role, role_assignment]
				define assignee: [user]
        type job
            relations
                define can_read: [user] but not cannot_read
                define cannot_read: cannot_read from permission
				define permission: [permission]
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 16)
	require.Len(t, graph.edges, 11)
	exclusionNodeID := graph.edges["job#can_read"][0].to.uniqueLabel
	exclusionEdges := graph.edges[exclusionNodeID]
	require.Len(t, exclusionEdges, 2)
	require.Equal(t, DirectEdge, exclusionEdges[0].GetEdgeType())
	cannotreadID := exclusionEdges[1].to.uniqueLabel
	require.Equal(t, "job#cannot_read", cannotreadID)
	exclusionNodeWeights := graph.nodes[exclusionNodeID].weights
	jobcannotReadWeights := graph.nodes[cannotreadID].weights
	require.Len(t, exclusionNodeWeights, 1)
	require.Equal(t, 3, exclusionNodeWeights["user"])
	require.Len(t, jobcannotReadWeights, 1)
	require.Equal(t, 3, jobcannotReadWeights["user"])
	require.Len(t, graph.edges[cannotreadID], 1)
	require.Equal(t, TTUEdge, graph.edges[cannotreadID][0].GetEdgeType())
	permissionCannotReadNode := graph.edges[cannotreadID][0].to.uniqueLabel
	require.Len(t, graph.edges[permissionCannotReadNode], 1)
	require.Equal(t, RewriteEdge, graph.edges[permissionCannotReadNode][0].GetEdgeType())
	orNode := graph.edges[permissionCannotReadNode][0].to.uniqueLabel
	require.Len(t, graph.edges[orNode], 2)
	require.Equal(t, LogicalTTUGrouping, graph.edges[orNode][0].to.GetNodeType())
	require.Equal(t, TTULogicalEdge, graph.edges[orNode][0].GetEdgeType())
	require.Equal(t, RewriteEdge, graph.edges[orNode][1].GetEdgeType())
	logicalTTUNode := graph.edges[orNode][0].to.uniqueLabel
	require.Len(t, graph.edges[logicalTTUNode], 2)
	require.Equal(t, TTUEdge, graph.edges[logicalTTUNode][0].GetEdgeType())
	require.Equal(t, TTUEdge, graph.edges[logicalTTUNode][1].GetEdgeType())
}

func TestGraphConstructionDirectAssignation(t *testing.T) {
	t.Parallel()
	model := `
	    model
                    schema 1.1
                type folder
                    relations
                        define viewer: [user]
                type user
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 3)
	require.Len(t, graph.edges, 1)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#viewer"].nodeType)
	require.Equal(t, "user", graph.edges["folder#viewer"][0].to.uniqueLabel)
}

func TestGraphConstructionMultipleUsersetWithoutOrder(t *testing.T) {
	t.Parallel()
	model := `
	    model
        schema 1.1
      type user
      type directs
          relations
              define direct: [user]
			  define mixed: [user]
			  define da: [user]
      type usersets-user
          relations
          define da: [user]
          define mixed: [directs#direct, user] and da
      type wrapper
          relations
          define parent: [usersets-user, directs]
          define assigned: mixed from parent but not da from parent
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 16)
	require.Len(t, graph.edges, 12)
	require.Len(t, graph.edges["wrapper#assigned"], 1)
	weight, _ := graph.nodes["wrapper#assigned"].GetWeight("user")
	require.Equal(t, 3, weight)

	butnot_node := graph.edges["wrapper#assigned"][0].to
	require.Len(t, graph.edges[butnot_node.uniqueLabel], 2)
	mixedLogicalTTU_node := graph.edges[butnot_node.uniqueLabel][0].to
	daLogicalTTU_node := graph.edges[butnot_node.uniqueLabel][1].to

	weight, _ = mixedLogicalTTU_node.GetWeight("user")
	require.Equal(t, 3, weight)
	weight, _ = daLogicalTTU_node.GetWeight("user")
	require.Equal(t, 2, weight)

	weight, _ = graph.edges[mixedLogicalTTU_node.uniqueLabel][0].GetWeight("user")
	require.Equal(t, 3, weight)
	weight, _ = graph.edges[mixedLogicalTTU_node.uniqueLabel][1].GetWeight("user")
	require.Equal(t, 2, weight)

	weight, _ = graph.edges[daLogicalTTU_node.uniqueLabel][0].GetWeight("user")
	require.Equal(t, 2, weight)
	weight, _ = graph.edges[daLogicalTTU_node.uniqueLabel][1].GetWeight("user")
	require.Equal(t, 2, weight)

	require.Len(t, graph.edges["usersets-user#mixed"], 1)
	weight, _ = graph.nodes["usersets-user#mixed"].GetWeight("user")
	require.Equal(t, 2, weight)
	and_node := graph.edges["usersets-user#mixed"][0].to
	weight, _ = graph.nodes[and_node.uniqueLabel].GetWeight("user")
	require.Equal(t, 2, weight)
	edges := graph.edges[and_node.uniqueLabel]
	ta_node := edges[1].to
	logical_node := edges[0].to

	require.Len(t, graph.edges[ta_node.uniqueLabel], 1)
	weight, _ = ta_node.GetWeight("user")
	require.Equal(t, 1, weight)

	require.Len(t, graph.edges[logical_node.uniqueLabel], 2)
	weight, _ = logical_node.GetWeight("user")
	require.Equal(t, 2, weight)

	weight, _ = graph.edges[logical_node.uniqueLabel][0].GetWeight("user")
	require.Equal(t, 2, weight)
	weight, _ = graph.edges[logical_node.uniqueLabel][1].GetWeight("user")
	require.Equal(t, 1, weight)

}

func TestGraphConstructionSuperNestedCycles(t *testing.T) {
	t.Parallel()
	model := `
	    model
	schema 1.1

	type user

	type document
		relations
			define viewer: [team#member, org#employee]

	type team
		relations
			define member: [user, document#viewer, org#employee]
			
	type org
		relations
			define employee: [user, document#viewer, team#member]
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 7)
	require.Len(t, graph.edges, 3)
	require.Len(t, graph.edges["document#viewer"], 2)
	weight, _ := graph.edges["document#viewer"][0].GetWeight("user")
	require.True(t, graph.edges["document#viewer"][0].IsPartOfTupleCycle())
	require.Equal(t, Infinite, weight)
	weight, _ = graph.edges["document#viewer"][1].GetWeight("user")
	require.Equal(t, Infinite, weight)
	require.True(t, graph.edges["document#viewer"][1].IsPartOfTupleCycle())

	require.Len(t, graph.edges["team#member"], 3)
	weight, _ = graph.edges["team#member"][0].GetWeight("user")
	require.Equal(t, 1, weight)
	require.False(t, graph.edges["team#member"][0].IsPartOfTupleCycle())
	weight, _ = graph.edges["team#member"][1].GetWeight("user")
	require.Equal(t, Infinite, weight)
	require.True(t, graph.edges["team#member"][1].IsPartOfTupleCycle())
	weight, _ = graph.edges["team#member"][2].GetWeight("user")
	require.Equal(t, Infinite, weight)
	require.True(t, graph.edges["team#member"][2].IsPartOfTupleCycle())

	require.Len(t, graph.edges["org#employee"], 3)
	weight, _ = graph.edges["org#employee"][0].GetWeight("user")
	require.False(t, graph.edges["org#employee"][0].IsPartOfTupleCycle())
	require.Equal(t, 1, weight)
	weight, _ = graph.edges["org#employee"][1].GetWeight("user")
	require.Equal(t, Infinite, weight)
	require.True(t, graph.edges["org#employee"][1].IsPartOfTupleCycle())
	weight, _ = graph.edges["org#employee"][2].GetWeight("user")
	require.Equal(t, Infinite, weight)
	require.True(t, graph.edges["org#employee"][2].IsPartOfTupleCycle())

}

func TestGraphConstructionAlgebraicWithNestedCycles(t *testing.T) {
	t.Parallel()
	model := `
	    model
	schema 1.1

	type user

	type document
		relations
			define viewer: [team#member] or (employee from org or member)
			define member: [user]
			define org: [org]

	type team
		relations
			define member: [user, org#employee]
			
	type org
		relations
			define employee: [user, document#viewer]
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 11)
	require.Len(t, graph.edges, 7)
	require.Len(t, graph.edges["document#viewer"], 1)
	weight, _ := graph.edges["document#viewer"][0].GetWeight("user")
	require.True(t, graph.edges["document#viewer"][0].IsPartOfTupleCycle())
	require.Equal(t, Infinite, weight)

	outer_or := graph.edges["document#viewer"][0].GetTo()
	outer_or_edges := graph.edges[outer_or.GetUniqueLabel()]
	require.Len(t, outer_or_edges, 2)
	require.True(t, outer_or.IsPartOfTupleCycle())
	require.Equal(t, Infinite, outer_or.weights["user"])
	weight, _ = outer_or_edges[0].GetWeight("user")
	require.Equal(t, Infinite, weight)
	require.True(t, outer_or_edges[0].IsPartOfTupleCycle())
	weight, _ = outer_or_edges[1].GetWeight("user")
	require.Equal(t, Infinite, weight)
	require.True(t, outer_or_edges[1].IsPartOfTupleCycle())

	inner_or := outer_or_edges[1].GetTo()
	inner_or_edges := graph.edges[inner_or.GetUniqueLabel()]
	require.True(t, inner_or.IsPartOfTupleCycle())
	require.Equal(t, Infinite, inner_or.weights["user"])
	require.Len(t, inner_or_edges, 2)
	weight, _ = inner_or_edges[0].GetWeight("user")
	require.Equal(t, Infinite, weight)
	require.True(t, inner_or_edges[0].IsPartOfTupleCycle())
	weight, _ = inner_or_edges[1].GetWeight("user")
	require.Equal(t, 1, weight)
	require.False(t, inner_or_edges[1].IsPartOfTupleCycle())

	require.Len(t, graph.edges["team#member"], 2)
	weight, _ = graph.edges["team#member"][0].GetWeight("user")
	require.Equal(t, 1, weight)
	require.False(t, graph.edges["team#member"][0].IsPartOfTupleCycle())
	weight, _ = graph.edges["team#member"][1].GetWeight("user")
	require.Equal(t, Infinite, weight)
	require.True(t, graph.edges["team#member"][1].IsPartOfTupleCycle())

	require.Len(t, graph.edges["org#employee"], 2)
	weight, _ = graph.edges["org#employee"][0].GetWeight("user")
	require.False(t, graph.edges["org#employee"][0].IsPartOfTupleCycle())
	require.Equal(t, 1, weight)
	weight, _ = graph.edges["org#employee"][1].GetWeight("user")
	require.Equal(t, Infinite, weight)
	require.True(t, graph.edges["org#employee"][1].IsPartOfTupleCycle())

}

func TestGraphConstructionRecursioncWithNestedCycles(t *testing.T) {
	t.Parallel()
	model := `
	    model
	schema 1.1

	type user

	type document
		relations
			define viewer: [document#viewer] or (employee from org or member)
			define member: [user]
			define org: [org]

	type team
		relations
			define member: [user, document#viewer]
			
	type org
		relations
			define employee: [user, team#member]
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 11)
	require.Len(t, graph.edges, 7)
	require.Len(t, graph.edges["document#viewer"], 1)
	weight, _ := graph.nodes["document#viewer"].GetWeight("user")
	require.True(t, graph.nodes["document#viewer"].IsPartOfTupleCycle())
	require.Equal(t, Infinite, weight)
	require.Equal(t, "document#viewer", graph.nodes["document#viewer"].recursiveRelation)

	outer_or := graph.edges["document#viewer"][0].GetTo()
	outer_or_edges := graph.edges[outer_or.GetUniqueLabel()]
	require.Equal(t, "document#viewer", outer_or.recursiveRelation)
	require.Len(t, outer_or_edges, 2)
	require.True(t, outer_or.IsPartOfTupleCycle())
	require.Equal(t, Infinite, outer_or.weights["user"])
	weight, _ = outer_or_edges[0].GetWeight("user")
	require.Equal(t, Infinite, weight)
	require.True(t, outer_or_edges[0].IsPartOfTupleCycle())
	require.Equal(t, "document#viewer", outer_or_edges[0].recursiveRelation)
	weight, _ = outer_or_edges[1].GetWeight("user")
	require.Equal(t, Infinite, weight)
	require.True(t, outer_or_edges[1].IsPartOfTupleCycle())
	require.Empty(t, outer_or_edges[1].recursiveRelation)

	inner_or := outer_or_edges[1].GetTo()
	inner_or_edges := graph.edges[inner_or.GetUniqueLabel()]
	require.True(t, inner_or.IsPartOfTupleCycle())
	require.Empty(t, inner_or.recursiveRelation)

	require.Equal(t, Infinite, inner_or.weights["user"])
	require.Len(t, inner_or_edges, 2)
	weight, _ = inner_or_edges[0].GetWeight("user")
	require.Equal(t, Infinite, weight)
	require.True(t, inner_or_edges[0].IsPartOfTupleCycle())
	weight, _ = inner_or_edges[1].GetWeight("user")
	require.Equal(t, 1, weight)
	require.False(t, inner_or_edges[1].IsPartOfTupleCycle())

	require.Len(t, graph.edges["team#member"], 2)
	weight, _ = graph.edges["team#member"][0].GetWeight("user")
	require.Equal(t, 1, weight)
	require.False(t, graph.edges["team#member"][0].IsPartOfTupleCycle())
	weight, _ = graph.edges["team#member"][1].GetWeight("user")
	require.Equal(t, Infinite, weight)
	require.True(t, graph.edges["team#member"][1].IsPartOfTupleCycle())

	require.Len(t, graph.edges["org#employee"], 2)
	weight, _ = graph.edges["org#employee"][0].GetWeight("user")
	require.False(t, graph.edges["org#employee"][0].IsPartOfTupleCycle())
	require.Equal(t, 1, weight)
	weight, _ = graph.edges["org#employee"][1].GetWeight("user")
	require.Equal(t, Infinite, weight)
	require.True(t, graph.edges["org#employee"][1].IsPartOfTupleCycle())

}

func TestGraphConstructionWildcardAssignation(t *testing.T) {
	t.Parallel()
	model := `
	    model
                    schema 1.1
                type folder
                    relations
                        define viewer: [user:*]
                type user
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 4)
	require.Len(t, graph.edges, 1)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificTypeWildcard, graph.nodes["user:*"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#viewer"].nodeType)
	require.Equal(t, "user:*", graph.edges["folder#viewer"][0].to.uniqueLabel)
}

func TestGraphConstructionDirectAssignmentWildardAndType(t *testing.T) {
	t.Parallel()
	model := `
	    model
                    schema 1.1
                type folder
                    relations
                        define viewer: [user:*, user]
                type user
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 4)
	require.Len(t, graph.edges, 1)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificTypeWildcard, graph.nodes["user:*"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#viewer"].nodeType)
	require.Len(t, graph.edges["folder#viewer"], 2)
	require.Equal(t, "user:*", graph.edges["folder#viewer"][0].to.uniqueLabel)
	require.Equal(t, "user", graph.edges["folder#viewer"][1].to.uniqueLabel)
}

func TestGraphConstructionDirectAssignmentWithUsersets(t *testing.T) {
	t.Parallel()
	model := `
	      model
                    schema 1.1
                type folder
                    relations
                        define viewer: [group#member]
                type group
                    relations
                        define member: [user]
                type user
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 5)
	require.Len(t, graph.edges, 2)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["group"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#viewer"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["group#member"].nodeType)
	require.Len(t, graph.edges["folder#viewer"], 1)
	require.Len(t, graph.edges["group#member"], 1)
}

func TestGraphConstructionDirectAssignmentWithUsersetRecursive(t *testing.T) {
	t.Parallel()
	model := `
	      model
                    schema 1.1
                type folder
                    relations
                        define viewer: [user, folder#viewer]
                type user
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 3)
	require.Len(t, graph.edges, 1)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#viewer"].nodeType)
	require.Len(t, graph.edges["folder#viewer"], 2)
	require.Equal(t, "user", graph.edges["folder#viewer"][0].to.uniqueLabel)
	require.Equal(t, "folder#viewer", graph.edges["folder#viewer"][1].to.uniqueLabel)
}

func TestGraphConstructionDirectAssignmentWithConditions(t *testing.T) {
	t.Parallel()
	model := `
	      model
                    schema 1.1
                type folder
                    relations
                        define viewer: [user with condX, user]
                type user
				condition condX (x:int) {
                    x > 0
                }
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 3)
	require.Len(t, graph.edges, 1)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#viewer"].nodeType)
	require.Len(t, graph.edges["folder#viewer"], 1)
	require.Len(t, graph.edges["folder#viewer"][0].conditions, 2)
	require.Equal(t, "user", graph.edges["folder#viewer"][0].to.uniqueLabel)
	require.Equal(t, "condX", graph.edges["folder#viewer"][0].conditions[0])
	require.Equal(t, "none", graph.edges["folder#viewer"][0].conditions[1])
}

func TestGraphConstructioComputedRelation(t *testing.T) {
	t.Parallel()
	model := `
	     model
                    schema 1.1
                type folder
                    relations
                        define x: y
                        define y: [user]
                type user
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 4)
	require.Len(t, graph.edges, 2)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#x"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#y"].nodeType)
	require.Len(t, graph.edges["folder#x"], 1)
	require.Len(t, graph.edges["folder#y"], 1)
	require.Equal(t, ComputedEdge, graph.edges["folder#x"][0].edgeType)
	require.Len(t, graph.edges["folder#y"][0].conditions, 1)
	require.Equal(t, "none", graph.edges["folder#y"][0].conditions[0])
	require.Equal(t, DirectEdge, graph.edges["folder#y"][0].edgeType)
}

func TestGraphConstructioComputedWithCycle(t *testing.T) {
	t.Parallel()
	model := `
	     model
                schema 1.1
                type folder
                    relations
                        define x: y
                        define y: z
                        define z: x
                type user
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	_, err := wgb.Build(authorizationModel)
	require.ErrorIs(t, err, ErrModelCycle)
}

func TestGraphConstructionTTU(t *testing.T) {
	t.Parallel()
	model := `
	       model
				schema 1.1
			type user
			type document
				relations
					define parent: [folder]
					define viewer: admin from parent
			type folder
				relations
					define admin: [user]
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 6)
	require.Len(t, graph.edges, 3)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["document"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#admin"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["document#parent"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["document#viewer"].nodeType)
	require.Equal(t, "folder#admin", graph.edges["document#viewer"][0].to.uniqueLabel)
	require.Equal(t, "document#parent", graph.edges["document#viewer"][0].tuplesetRelation)
	require.Equal(t, TTUEdge, graph.edges["document#viewer"][0].edgeType)

}

func TestGraphConstructionTTUConditional(t *testing.T) {
	t.Parallel()
	model := `
	       model
                    schema 1.1
                type user
                type document
                    relations
                        define parent: [folder, folder with condX]
                        define viewer: admin from parent
                type folder
                    relations
                        define admin: [user]
                condition condX (x:int) {
                    x > 0
                }
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 6)
	require.Len(t, graph.edges, 3)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["document"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#admin"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["document#parent"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["document#viewer"].nodeType)
	require.Len(t, graph.edges["document#viewer"], 1)
	require.Equal(t, "folder#admin", graph.edges["document#viewer"][0].to.uniqueLabel)
	require.Equal(t, "document#parent", graph.edges["document#viewer"][0].tuplesetRelation)
	require.Equal(t, TTUEdge, graph.edges["document#viewer"][0].edgeType)
	require.Len(t, graph.edges["document#viewer"][0].conditions, 1)
	require.Len(t, graph.edges["document#parent"], 1)
	require.Len(t, graph.edges["document#parent"][0].conditions, 2)
	require.Equal(t, "none", graph.edges["document#parent"][0].conditions[0])
	require.Equal(t, "condX", graph.edges["document#parent"][0].conditions[1])
}

func TestGraphConstructionUsersetConditional(t *testing.T) {
	t.Parallel()
	model := `
	      model
                    schema 1.1
                type folder
                    relations
                        define viewer: [group#member, group#member with condX]
                type group
                    relations
                        define member: [user]
                type user
                condition condX (x:int) {
                    x > 0
                }
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 5)
	require.Len(t, graph.edges, 2)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["group"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#viewer"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["group#member"].nodeType)
	require.Len(t, graph.edges["folder#viewer"], 1)
	require.Equal(t, "group#member", graph.edges["folder#viewer"][0].to.uniqueLabel)
	require.Equal(t, DirectEdge, graph.edges["folder#viewer"][0].edgeType)
	require.Len(t, graph.edges["folder#viewer"][0].conditions, 2)
	require.Equal(t, "none", graph.edges["folder#viewer"][0].conditions[0])
	require.Equal(t, "condX", graph.edges["folder#viewer"][0].conditions[1])
}

func TestGraphConstructionTTURecursive(t *testing.T) {
	t.Parallel()
	model := `
	     model
                    schema 1.1
                type user
                type folder
                    relations
                        define parent: [folder]
                        define viewer: [user] or viewer from parent
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 5)
	require.Len(t, graph.edges, 3)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#viewer"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#parent"].nodeType)

	require.Len(t, graph.edges["folder#parent"], 1)
	require.Len(t, graph.edges["folder#viewer"], 1)
	require.Equal(t, RewriteEdge, graph.edges["folder#viewer"][0].edgeType)
	require.Equal(t, OperatorNode, graph.edges["folder#viewer"][0].to.nodeType)
	unionNode := graph.edges["folder#viewer"][0].to
	require.Equal(t, "union", unionNode.label)
	require.Len(t, graph.edges[unionNode.uniqueLabel], 2)
	require.Equal(t, "user", graph.edges[unionNode.uniqueLabel][0].to.uniqueLabel)
	require.Equal(t, TTUEdge, graph.edges[unionNode.uniqueLabel][1].edgeType)
	require.Equal(t, "folder#parent", graph.edges[unionNode.uniqueLabel][1].tuplesetRelation)
	require.Equal(t, "folder#viewer", graph.edges[unionNode.uniqueLabel][1].to.uniqueLabel)
}

func TestGraphConstructionTTUWithTwoParents(t *testing.T) {
	t.Parallel()
	model := `
	     model
                    schema 1.1
                type user
                type document
                    relations
                        define parent: [folder, folder2]
                        define viewer: admin from parent
                type folder
                    relations
                        define admin: [user]
                type folder2
                    relations
                        define admin: [user]
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 8)
	require.Len(t, graph.edges, 4)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["document"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder2"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["document#parent"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["document#viewer"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#admin"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder2#admin"].nodeType)
	require.Len(t, graph.edges["document#parent"], 2)
	require.Len(t, graph.edges["document#viewer"], 2)
	require.Equal(t, TTUEdge, graph.edges["document#viewer"][0].edgeType)
	require.Equal(t, "folder#admin", graph.edges["document#viewer"][0].to.uniqueLabel)

	require.Equal(t, TTUEdge, graph.edges["document#viewer"][1].edgeType)
	require.Equal(t, "folder2#admin", graph.edges["document#viewer"][1].to.uniqueLabel)
}

func TestGraphConstructionInvalidTTU(t *testing.T) {
	t.Parallel()
	model := `
	    model
                    schema 1.1
                type user
                type document
                    relations
                        define parent: [folder, folder2]
                        define viewer: admin from parent
                type folder
                    relations
                        define admin: [user]
                type folder2
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	_, err := wgb.Build(authorizationModel)
	require.ErrorIs(t, err, ErrInvalidModel)
}

func TestGraphConstructionIntersection(t *testing.T) {
	t.Run("two_operands", func(t *testing.T) {
		t.Run("balanced", func(t *testing.T) {
			t.Parallel()
			model := `
	      model
		  	schema 1.1
		  type user
		  type folder
		  	  relations
				define a: [user]
				define b: [user]
				define c: a and b
		  `
			authorizationModel := language.MustTransformDSLToProto(model)
			wgb := NewWeightedAuthorizationModelGraphBuilder()
			graph, err := wgb.Build(authorizationModel)
			require.NoError(t, err)

			require.Len(t, graph.nodes, 6)
			require.Len(t, graph.edges, 4)
			require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
			require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
			require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#a"].nodeType)
			require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#b"].nodeType)
			require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#c"].nodeType)
			andNode := graph.edges["folder#c"][0].to
			require.Equal(t, OperatorNode, andNode.nodeType)
			require.Equal(t, "intersection", andNode.label)
			require.Len(t, graph.edges[andNode.uniqueLabel], 2)
			require.Equal(t, "folder#a", graph.edges[andNode.uniqueLabel][0].to.uniqueLabel)
			require.Equal(t, RewriteEdge, graph.edges[andNode.uniqueLabel][0].edgeType)
			require.Equal(t, "folder#b", graph.edges[andNode.uniqueLabel][1].to.uniqueLabel)
			require.Equal(t, RewriteEdge, graph.edges[andNode.uniqueLabel][1].edgeType)
			require.Len(t, graph.edges["folder#a"], 1)
			require.Equal(t, DirectEdge, graph.edges["folder#a"][0].edgeType)
			require.Len(t, graph.edges["folder#b"], 1)
			require.Equal(t, DirectEdge, graph.edges["folder#b"][0].edgeType)
			require.Len(t, graph.edges["folder#c"], 1)
			require.Equal(t, RewriteEdge, graph.edges["folder#c"][0].edgeType)
		})

		t.Run("unbalanced", func(t *testing.T) {
			t.Parallel()
			model := `
	      model
		  	schema 1.1
		  type user1
		  type user2
		  type folder
		  	  relations
				define a: [user1,user2]
				define b: [user1]
				define c: a and b
		  `
			authorizationModel := language.MustTransformDSLToProto(model)
			wgb := NewWeightedAuthorizationModelGraphBuilder()
			graph, err := wgb.Build(authorizationModel)
			require.NoError(t, err)

			require.Len(t, graph.nodes, 7)
			require.Len(t, graph.edges["folder#c"], 1)
			require.Equal(t, map[string]int{"user1": 1}, graph.edges["folder#c"][0].GetWeights()) // Because only [user1] on both sides of AND
		})
	})

	t.Run("three_operands_balanced_direct", func(t *testing.T) {
		t.Parallel()
		model := `
			model
				schema 1.1
			type user1
			type user2
			type user3

			type document
				relations
					define viewer: [user1,user2,user3] and can_view
					define can_view: [user1,user2,user3]`

		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		require.NoError(t, err)

		require.Len(t, graph.nodes, 8)
		require.Len(t, graph.edges["document#viewer"], 1)
		require.Equal(t, map[string]int{"user1": 1, "user2": 1, "user3": 1}, graph.edges["document#viewer"][0].GetWeights()) // Because all [user1, user2, user3] on both sides of AND
	})

	t.Run("three_operands_unbalanced_direct", func(t *testing.T) {
		t.Run("left_unbalanced", func(t *testing.T) {
			t.Parallel()
			model := `
				model
					schema 1.1
				type user1
				type user2
				type user3

				type document
					relations
						define viewer: [user1,user3] and can_view
						define can_view: [user1,user2,user3]`

			authorizationModel := language.MustTransformDSLToProto(model)
			wgb := NewWeightedAuthorizationModelGraphBuilder()
			graph, err := wgb.Build(authorizationModel)
			require.NoError(t, err)

			require.Len(t, graph.nodes, 8)
			require.Len(t, graph.edges["document#viewer"], 1)
			require.Equal(t, map[string]int{"user1": 1, "user3": 1}, graph.edges["document#viewer"][0].GetWeights()) // Because only [user1, user3] on left of AND
		})
		t.Run("right_unbalanced", func(t *testing.T) {
			t.Parallel()
			model := `
				model
					schema 1.1
				type user1
				type user2
				type user3

				type document
					relations
						define viewer: [user1,user2,user3] and can_view
						define can_view: [user1,user3]`

			authorizationModel := language.MustTransformDSLToProto(model)
			wgb := NewWeightedAuthorizationModelGraphBuilder()
			graph, err := wgb.Build(authorizationModel)
			require.NoError(t, err)

			require.Len(t, graph.nodes, 8)
			require.Len(t, graph.edges["document#viewer"], 1)
			require.Equal(t, map[string]int{"user1": 1, "user3": 1}, graph.edges["document#viewer"][0].GetWeights()) // Because only [user1, user3] on right of AND
		})
	})

	t.Run("three_operands_unbalanced_rewrite", func(t *testing.T) {
		t.Parallel()
		t.Run("left_unbalanced", func(t *testing.T) {
			model := `
			model
				schema 1.1
			type user1
			type user2
			type user3

			type document
				relations
					define viewer: viewer1 and viewer2
					define viewer1: [user1,user3]
					define viewer2: [user1,user2,user3]`

			authorizationModel := language.MustTransformDSLToProto(model)
			wgb := NewWeightedAuthorizationModelGraphBuilder()
			graph, err := wgb.Build(authorizationModel)
			require.NoError(t, err)

			require.Len(t, graph.nodes, 8)
			require.Len(t, graph.edges["document#viewer"], 1)
			require.Equal(t, map[string]int{"user1": 1, "user3": 1}, graph.edges["document#viewer"][0].GetWeights()) // Because only [user1, user3] on left of AND
		})
		t.Run("right_unbalanced", func(t *testing.T) {
			model := `
				model
					schema 1.1
				type user1
				type user2
				type user3

				type document
					relations
						define viewer: viewer1 and viewer2
						define viewer1: [user1,user3]
						define viewer2: [user1,user3]`

			authorizationModel := language.MustTransformDSLToProto(model)
			wgb := NewWeightedAuthorizationModelGraphBuilder()
			graph, err := wgb.Build(authorizationModel)
			require.NoError(t, err)

			require.Len(t, graph.nodes, 8)
			require.Len(t, graph.edges["document#viewer"], 1)
			require.Equal(t, map[string]int{"user1": 1, "user3": 1}, graph.edges["document#viewer"][0].GetWeights()) // Because only [user1, user3] on right of AND
		})
	})

	t.Run("three_operands_balanced_rewrite", func(t *testing.T) {
		t.Parallel()
		model := `
			model
				schema 1.1
			type user1
			type user2
			type user3

			type document
				relations
					define viewer: viewer1 and viewer2
					define viewer1: [user1,user2,user3]
					define viewer2: [user1,user2,user3]`

		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		require.NoError(t, err)

		require.Len(t, graph.nodes, 8)
		require.Len(t, graph.edges["document#viewer"], 1)
		require.Equal(t, map[string]int{"user1": 1, "user2": 1, "user3": 1}, graph.edges["document#viewer"][0].GetWeights()) // Because all [user1, user2, user3] on both sides of AND
	})

	t.Run("three_operands_balanced_rewrite_userset_directly_assigned", func(t *testing.T) {
		t.Parallel()
		model := `
			model
				schema 1.1
			type user1
			type user2
			type user3

			type document
				relations
					define viewer: [document#viewer1] and viewer2
					define viewer1: [user1,user2,user3]
					define viewer2: [user1,user2,user3]`

		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		require.NoError(t, err)

		require.Len(t, graph.nodes, 8)
		require.Len(t, graph.edges["document#viewer"], 1)
		require.Equal(t, map[string]int{"user1": 2, "user2": 2, "user3": 2}, graph.edges["document#viewer"][0].GetWeights()) // Because all [user1, user2, user3, document#viewer1] on both sides of AND
	})

	t.Run("three_operands_balanced_rewrite_userset_directly_assigned_mixed", func(t *testing.T) {
		t.Parallel()
		model := `
			model
				schema 1.1
			type user1
			type user2
			type user3

			type document
				relations
					define viewer: [user1,user2, user3, document#viewer1] and viewer2
					define viewer1: [user1,user2,user3]
					define viewer2: [user1,user2,user3]`

		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		require.NoError(t, err)

		require.Len(t, graph.nodes, 9)
		require.Len(t, graph.edges["document#viewer"], 1)
		require.Equal(t, map[string]int{"user1": 2, "user2": 2, "user3": 2}, graph.edges["document#viewer"][0].GetWeights()) // Because all [user1, user2, user3, document#viewer1] on both sides of AND
	})

	t.Run("three_operands_balanced_rewrite_mixed_wildcards", func(t *testing.T) {
		t.Parallel()
		model := `
			model
				schema 1.1
			type user1
			type user2
			type user3

			type document
				relations
					define viewer: viewer1 and viewer2
					define viewer1: [user1,user2:*,user3]
					define viewer2: [user1:*,user2,user3:*]`

		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		require.NoError(t, err)

		require.Len(t, graph.nodes, 11)
		require.Len(t, graph.edges["document#viewer"], 1)
		require.Equal(t, map[string]int{"user1": 1, "user2": 1, "user3": 1}, graph.edges["document#viewer"][0].GetWeights()) // Because all [user1, user2, user3] on both sides of AND
	})

	t.Run("invalid_model", func(t *testing.T) {
		t.Run("with_direct_types", func(t *testing.T) {

			t.Parallel()
			model := `
				model
					schema 1.1
				type user1
				type user2
				type user3

				type document
					relations
						define viewer: [user1,user2] and owner
						define owner: [user3]`

			authorizationModel := language.MustTransformDSLToProto(model)
			wgb := NewWeightedAuthorizationModelGraphBuilder()
			_, err := wgb.Build(authorizationModel)
			require.ErrorContains(t, err, "invalid model: not all paths return the same type for the node intersection:")
		})
		t.Run("without_direct_types", func(t *testing.T) {
			t.Parallel()
			model := `
				model
					schema 1.1
				type user1
				type user2
				type user3

				type document
					relations
						define viewer: viewer1 and viewer2
						define viewer1: [user1,user2]
						define viewer2: [user3]`

			authorizationModel := language.MustTransformDSLToProto(model)
			wgb := NewWeightedAuthorizationModelGraphBuilder()
			_, err := wgb.Build(authorizationModel)
			require.ErrorContains(t, err, "invalid model: not all paths return the same type for the node intersection:")
		})
	})

	t.Run("multiple_intersections", func(t *testing.T) {
		t.Run("balanced", func(t *testing.T) {
			t.Parallel()
			model := `
		      model
			  	schema 1.1
			  type user1
			  type user2
			  type user3
			  type document
			  	  relations
					define a: [user1,user2,user3] and b and c
					define b: [user1,user2,user3]
					define c: [user1,user2,user3]
			  `
			authorizationModel := language.MustTransformDSLToProto(model)
			wgb := NewWeightedAuthorizationModelGraphBuilder()
			graph, err := wgb.Build(authorizationModel)
			require.NoError(t, err)

			require.Len(t, graph.nodes, 9)
			require.Len(t, graph.edges, 5)
			require.Len(t, graph.edges["document#a"], 1)
			require.Equal(t, map[string]int{"user1": 1, "user2": 1, "user3": 1}, graph.edges["document#a"][0].GetWeights())
		})

		t.Run("unbalanced", func(t *testing.T) {
			t.Parallel()
			model := `
		      model
			  	schema 1.1
			  type user1
			  type user2
			  type user3
			  type document
			  	  relations
					define a: [user1,user2] and b and c
					define b: [user1,user3]
					define c: [user1,user2]
			  `
			authorizationModel := language.MustTransformDSLToProto(model)
			wgb := NewWeightedAuthorizationModelGraphBuilder()
			graph, err := wgb.Build(authorizationModel)
			require.NoError(t, err)

			require.Len(t, graph.nodes, 9)
			require.Len(t, graph.edges, 5)
			require.Len(t, graph.edges["document#a"], 1)
			require.Equal(t, map[string]int{"user1": 1}, graph.edges["document#a"][0].GetWeights())
		})
	})

}

func TestGraphConstructionIntersectionWithType(t *testing.T) {
	t.Parallel()
	model := `
	      model
                    schema 1.1
                type user
                type folder
                   relations
                     define a: [user]
                     define b: [user] and a
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 5)
	require.Len(t, graph.edges, 3)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#a"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#b"].nodeType)
	andNode := graph.edges["folder#b"][0].to
	require.Equal(t, OperatorNode, andNode.nodeType)
	require.Equal(t, "intersection", andNode.label)
	require.Len(t, graph.edges[andNode.uniqueLabel], 2)
	require.Equal(t, "user", graph.edges[andNode.uniqueLabel][0].to.uniqueLabel)
	require.Equal(t, DirectEdge, graph.edges[andNode.uniqueLabel][0].edgeType)
	require.Equal(t, "folder#a", graph.edges[andNode.uniqueLabel][1].to.uniqueLabel)
	require.Equal(t, RewriteEdge, graph.edges[andNode.uniqueLabel][1].edgeType)
	require.Len(t, graph.edges["folder#a"], 1)
	require.Equal(t, DirectEdge, graph.edges["folder#a"][0].edgeType)
}

func TestGraphConstructionNestedIntersection(t *testing.T) {
	t.Parallel()
	model := `
	      model
                    schema 1.1
                type user
                type folder
                   relations
                     define a: [user]
                     define b: [user]
                     define c: [user]
                     define d: [user]
                     define e: (a and b and c) and d
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 9)
	require.Len(t, graph.edges, 7)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#a"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#b"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#c"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#d"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#e"].nodeType)

	outerAndNode := graph.edges["folder#e"][0].to
	require.Equal(t, OperatorNode, outerAndNode.nodeType)
	require.Equal(t, "intersection", outerAndNode.label)
	require.Len(t, graph.edges[outerAndNode.uniqueLabel], 2)
	require.Equal(t, "folder#d", graph.edges[outerAndNode.uniqueLabel][1].to.uniqueLabel)
	require.Equal(t, RewriteEdge, graph.edges[outerAndNode.uniqueLabel][1].edgeType)
	require.Equal(t, RewriteEdge, graph.edges[outerAndNode.uniqueLabel][0].edgeType)
	innerAndNode := graph.edges[outerAndNode.uniqueLabel][0].to
	require.Equal(t, OperatorNode, innerAndNode.nodeType)
	require.Equal(t, "intersection", innerAndNode.label)
	require.Len(t, graph.edges[innerAndNode.uniqueLabel], 3)

	require.Equal(t, RewriteEdge, graph.edges[innerAndNode.uniqueLabel][0].edgeType)
	require.Equal(t, "folder#a", graph.edges[innerAndNode.uniqueLabel][0].to.uniqueLabel)

	require.Equal(t, RewriteEdge, graph.edges[innerAndNode.uniqueLabel][1].edgeType)
	require.Equal(t, "folder#b", graph.edges[innerAndNode.uniqueLabel][1].to.uniqueLabel)

	require.Equal(t, RewriteEdge, graph.edges[innerAndNode.uniqueLabel][2].edgeType)
	require.Equal(t, "folder#c", graph.edges[innerAndNode.uniqueLabel][2].to.uniqueLabel)
}

func TestGraphConstructionNestedUnion(t *testing.T) {
	t.Parallel()
	model := `
	      model
                    schema 1.1
                type user
                type folder
                   relations
                     define a: [user]
                     define b: [user]
                     define c: [user]
                     define d: [user]
                     define e: (a or b or c) or d
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 9)
	require.Len(t, graph.edges, 7)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#a"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#b"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#c"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#d"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#e"].nodeType)

	outerOrNode := graph.edges["folder#e"][0].to
	require.Equal(t, OperatorNode, outerOrNode.nodeType)
	require.Equal(t, "union", outerOrNode.label)
	require.Len(t, graph.edges[outerOrNode.uniqueLabel], 2)
	require.Equal(t, "folder#d", graph.edges[outerOrNode.uniqueLabel][1].to.uniqueLabel)
	require.Equal(t, RewriteEdge, graph.edges[outerOrNode.uniqueLabel][1].edgeType)
	require.Equal(t, RewriteEdge, graph.edges[outerOrNode.uniqueLabel][0].edgeType)
	innerOrNode := graph.edges[outerOrNode.uniqueLabel][0].to
	require.Equal(t, OperatorNode, innerOrNode.nodeType)
	require.Equal(t, "union", innerOrNode.label)
	require.Len(t, graph.edges[innerOrNode.uniqueLabel], 3)

	require.Equal(t, RewriteEdge, graph.edges[innerOrNode.uniqueLabel][0].edgeType)
	require.Equal(t, "folder#a", graph.edges[innerOrNode.uniqueLabel][0].to.uniqueLabel)

	require.Equal(t, RewriteEdge, graph.edges[innerOrNode.uniqueLabel][1].edgeType)
	require.Equal(t, "folder#b", graph.edges[innerOrNode.uniqueLabel][1].to.uniqueLabel)

	require.Equal(t, RewriteEdge, graph.edges[innerOrNode.uniqueLabel][2].edgeType)
	require.Equal(t, "folder#c", graph.edges[innerOrNode.uniqueLabel][2].to.uniqueLabel)
}

func TestGraphConstructionInvalidRecursiveUnion(t *testing.T) {
	t.Parallel()

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("TestUserFail should have panicked!")
			}
		}()
		// This function should cause a panic
		model := `
	      model
                    schema 1.1
                type user
                types folder
                   relations
                     define a: [user] or b
                     define b: [user] or c
                     define c: [user] or a
	`
		language.MustTransformDSLToProto(model)
	}()
}

func TestGraphConstructionMultigraph(t *testing.T) {
	t.Parallel()
	model := `
	        model
                  schema 1.1
                
                type user
                
                type state
                  relations
                    define can_view: [user]
                
                type transition
                  relations
                    define start: [state]
                    define end: [state]
                    define can_apply: [user] and can_view from start and can_view from end
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 8)
	require.Len(t, graph.edges, 5)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["state"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["transition"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["state#can_view"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["transition#start"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["transition#end"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["transition#can_apply"].nodeType)

	andNode := graph.edges["transition#can_apply"][0].to
	require.Equal(t, OperatorNode, andNode.nodeType)
	require.Equal(t, "intersection", andNode.label)
	require.Len(t, graph.edges[andNode.uniqueLabel], 3)
	require.Equal(t, "user", graph.edges[andNode.uniqueLabel][0].to.uniqueLabel)
	require.Equal(t, DirectEdge, graph.edges[andNode.uniqueLabel][0].edgeType)
	require.Equal(t, TTUEdge, graph.edges[andNode.uniqueLabel][1].edgeType)
	require.Equal(t, "state#can_view", graph.edges[andNode.uniqueLabel][1].to.uniqueLabel)
	require.Equal(t, "transition#start", graph.edges[andNode.uniqueLabel][1].tuplesetRelation)

	require.Equal(t, TTUEdge, graph.edges[andNode.uniqueLabel][2].edgeType)
	require.Equal(t, "state#can_view", graph.edges[andNode.uniqueLabel][2].to.uniqueLabel)
	require.Equal(t, "transition#end", graph.edges[andNode.uniqueLabel][2].tuplesetRelation)
}

func TestGraphConstructionExclusion(t *testing.T) {
	t.Parallel()
	model := `
	      model
			schema 1.1
		type user
		type folder
			relations
				define a: [user]
				define b: [user]
				define c: a but not b
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 6)
	require.Len(t, graph.edges, 4)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#a"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#b"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#c"].nodeType)
	exclusionNode := graph.edges["folder#c"][0].to
	require.Equal(t, OperatorNode, exclusionNode.nodeType)
	require.Equal(t, "exclusion", exclusionNode.label)
	require.Len(t, graph.edges[exclusionNode.uniqueLabel], 2)
	require.Equal(t, "folder#a", graph.edges[exclusionNode.uniqueLabel][0].to.uniqueLabel)
	require.Equal(t, RewriteEdge, graph.edges[exclusionNode.uniqueLabel][0].edgeType)
	require.Equal(t, "folder#b", graph.edges[exclusionNode.uniqueLabel][1].to.uniqueLabel)
	require.Equal(t, RewriteEdge, graph.edges[exclusionNode.uniqueLabel][1].edgeType)
	require.Len(t, graph.edges["folder#a"], 1)
	require.Equal(t, DirectEdge, graph.edges["folder#a"][0].edgeType)
	require.Len(t, graph.edges["folder#b"], 1)
	require.Equal(t, DirectEdge, graph.edges["folder#b"][0].edgeType)
	require.Len(t, graph.edges["folder#c"], 1)
	require.Equal(t, RewriteEdge, graph.edges["folder#c"][0].edgeType)
}

func TestGraphConstructionExclusionWithType(t *testing.T) {
	t.Parallel()
	model := `
	      model
                    schema 1.1
                type user
                type folder
                   relations
                     define a: [user]
                     define b: [user] but not a
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 5)
	require.Len(t, graph.edges, 3)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#a"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#b"].nodeType)
	exclusionNode := graph.edges["folder#b"][0].to
	require.Equal(t, OperatorNode, exclusionNode.nodeType)
	require.Equal(t, "exclusion", exclusionNode.label)
	require.Len(t, graph.edges[exclusionNode.uniqueLabel], 2)
	require.Equal(t, "user", graph.edges[exclusionNode.uniqueLabel][0].to.uniqueLabel)
	require.Equal(t, DirectEdge, graph.edges[exclusionNode.uniqueLabel][0].edgeType)
	require.Equal(t, "folder#a", graph.edges[exclusionNode.uniqueLabel][1].to.uniqueLabel)
	require.Equal(t, RewriteEdge, graph.edges[exclusionNode.uniqueLabel][1].edgeType)
	require.Len(t, graph.edges["folder#a"], 1)
	require.Equal(t, DirectEdge, graph.edges["folder#a"][0].edgeType)
	require.Len(t, graph.edges["folder#b"], 1)
	require.Equal(t, RewriteEdge, graph.edges["folder#b"][0].edgeType)
}

func TestGraphConstructionMixedAlg(t *testing.T) {
	t.Parallel()
	model := `
		                model
                    schema 1.1
                type user
                type folder
                    relations
                        define a: [user]
                        define b: [user]
                        define c: [user]
                        define d: [user]
                        define e: (a or b or c) but not d  
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)

	require.Len(t, graph.nodes, 9)
	require.Len(t, graph.edges, 7)
	require.Equal(t, SpecificType, graph.nodes["user"].nodeType)
	require.Equal(t, SpecificType, graph.nodes["folder"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#a"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#b"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#c"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#d"].nodeType)
	require.Equal(t, SpecificTypeAndRelation, graph.nodes["folder#e"].nodeType)

	exclusionNode := graph.edges["folder#e"][0].to
	require.Equal(t, OperatorNode, exclusionNode.nodeType)
	require.Equal(t, "exclusion", exclusionNode.label)
	require.Len(t, graph.edges[exclusionNode.uniqueLabel], 2)
	orNode := graph.edges[exclusionNode.uniqueLabel][0].to
	require.Equal(t, OperatorNode, orNode.nodeType)
	require.Equal(t, "union", orNode.label)
	require.Len(t, graph.edges[orNode.uniqueLabel], 3)
	require.Equal(t, "folder#a", graph.edges[orNode.uniqueLabel][0].to.uniqueLabel)
	require.Equal(t, RewriteEdge, graph.edges[orNode.uniqueLabel][0].edgeType)
	require.Equal(t, "folder#b", graph.edges[orNode.uniqueLabel][1].to.uniqueLabel)
	require.Equal(t, RewriteEdge, graph.edges[orNode.uniqueLabel][1].edgeType)
	require.Equal(t, "folder#c", graph.edges[orNode.uniqueLabel][2].to.uniqueLabel)
	require.Equal(t, RewriteEdge, graph.edges[orNode.uniqueLabel][2].edgeType)
	require.Equal(t, "folder#d", graph.edges[exclusionNode.uniqueLabel][1].to.uniqueLabel)
	require.Equal(t, RewriteEdge, graph.edges[exclusionNode.uniqueLabel][1].edgeType)
	require.Len(t, graph.edges["folder#a"], 1)
	require.Equal(t, DirectEdge, graph.edges["folder#a"][0].edgeType)
	require.Len(t, graph.edges["folder#b"], 1)
	require.Equal(t, DirectEdge, graph.edges["folder#b"][0].edgeType)
	require.Len(t, graph.edges["folder#c"], 1)
	require.Equal(t, DirectEdge, graph.edges["folder#c"][0].edgeType)
	require.Len(t, graph.edges["folder#d"], 1)
	require.Equal(t, DirectEdge, graph.edges["folder#d"][0].edgeType)
}

func TestGraphConstructionInvalidDSL(t *testing.T) {
	t.Parallel()
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("TestUserFail should have panicked!")
			}
		}()
		// This function should cause a panic
		model := `
	      model
                    schema 1.1
                type user
                types folder
                   relations
                      define x: y
                        define y: x
                        define a: [user] or x or b
                        define b: [user] or c
                        define c: [user] or a
	`
		language.MustTransformDSLToProto(model)
	}()
}

func TestGraphConstructionInvalidModelCycle(t *testing.T) {
	t.Parallel()
	model := `
	                model
      schema 1.1
    type user
    type group
      relations
        define member: [user] or memberA or memberB or memberC
        define memberA: [user] or member or memberB or memberC
        define memberB: [user] or member or memberA or memberC
        define memberC: [user] or member or memberA or memberB
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	_, err := wgb.Build(authorizationModel)
	require.ErrorIs(t, err, ErrModelCycle)

}

func TestGraphConstructionInvalidModelCycle2(t *testing.T) {
	t.Parallel()
	model := `
	   model
      schema 1.1
    type user
    type account
      relations
        define admin: [user] or member or super_admin or owner
        define member: [user] or owner or admin or super_admin
        define super_admin: [user] or admin or member or owner
        define owner: [user]
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	_, err := wgb.Build(authorizationModel)
	require.ErrorIs(t, err, ErrModelCycle)

}

func TestGraphConstructionInvalidModelCycle3(t *testing.T) {
	t.Parallel()
	model := `
	   model
      schema 1.1
      type user

    type document
      relations
        define admin: [user]
        define action1: admin and action2 and action3
        define action2: admin and action1 and action3
        define action3: admin and action1 and action2
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	_, err := wgb.Build(authorizationModel)
	require.ErrorIs(t, err, ErrModelCycle)

}

func TestGraphConstructionTupleCycles(t *testing.T) {
	t.Run("no_cycles", func(t *testing.T) {
		t.Parallel()
		model := `
		model
			schema 1.1
		type user
		type document
			relations
				define viewer: [user]
				define can_view: [document#viewer, user]
		`
		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		require.NoError(t, err)

		for _, node := range graph.nodes {
			require.Empty(t, node.GetRecursiveRelation())
			require.Empty(t, node.IsPartOfTupleCycle())
		}

		for _, edges := range graph.edges {
			for _, edge := range edges {
				require.Empty(t, edge.GetRecursiveRelation())
				require.Empty(t, edge.IsPartOfTupleCycle())
			}
		}
	})

	t.Run("tuple_two_usersets", func(t *testing.T) {
		t.Parallel()
		model := `
		model
			schema 1.1
		type user
		type folder
			relations
				define viewer: [user, folder#can_view]
				define can_view: [user, folder#viewer]`

		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		require.NoError(t, err)

		require.Len(t, graph.nodes, 4)
		require.Len(t, graph.edges, 2)
		require.True(t, graph.nodes["user"].nodeType == SpecificType)
		require.True(t, graph.nodes["folder"].nodeType == SpecificType)
		require.True(t, graph.nodes["folder#viewer"].nodeType == SpecificTypeAndRelation)
		require.True(t, graph.nodes["folder#can_view"].nodeType == SpecificTypeAndRelation)

		for _, node := range graph.nodes {
			require.Empty(t, node.GetRecursiveRelation())
		}
		require.True(t, graph.nodes["folder#can_view"].IsPartOfTupleCycle())
		require.True(t, graph.nodes["folder#viewer"].IsPartOfTupleCycle())

		require.Len(t, graph.edges["folder#can_view"], 2)
		require.Len(t, graph.edges["folder#viewer"], 2)
		require.True(t, graph.edges["folder#viewer"][0].edgeType == DirectEdge)
		require.True(t, graph.edges["folder#viewer"][0].to.nodeType == SpecificType)
		require.True(t, graph.edges["folder#viewer"][1].edgeType == DirectEdge)
		require.True(t, graph.edges["folder#viewer"][1].to.nodeType == SpecificTypeAndRelation)
		require.True(t, graph.edges["folder#viewer"][1].to.uniqueLabel == "folder#can_view")

		require.True(t, graph.edges["folder#can_view"][0].edgeType == DirectEdge)
		require.True(t, graph.edges["folder#can_view"][0].to.nodeType == SpecificType)
		require.True(t, graph.edges["folder#can_view"][1].edgeType == DirectEdge)
		require.True(t, graph.edges["folder#can_view"][1].to.nodeType == SpecificTypeAndRelation)
		require.True(t, graph.edges["folder#can_view"][1].to.uniqueLabel == "folder#viewer")
	})

	t.Run("recursive_cycles_with_intermediate_relation", func(t *testing.T) {
		model := `
			model
				schema 1.1
			type user
			type group
				relations
					define inherited_member: member from parent
					define member: [user] or inherited_member
					define parent: [group]
		`
		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		require.NoError(t, err)

		require.Len(t, graph.nodes, 6)
		require.Len(t, graph.edges, 4)

		for _, node := range graph.nodes {
			require.Empty(t, node.GetRecursiveRelation())
		}

		require.True(t, graph.nodes["group#inherited_member"].IsPartOfTupleCycle())
		require.True(t, graph.nodes["group#member"].IsPartOfTupleCycle())
		require.False(t, graph.nodes["group#parent"].IsPartOfTupleCycle())

		require.Empty(t, graph.edges["group#parent"][0].GetRecursiveRelation())
		require.False(t, graph.edges["group#parent"][0].IsPartOfTupleCycle())
		require.Empty(t, graph.edges["group#inherited_member"][0].GetRecursiveRelation())
		require.True(t, graph.edges["group#inherited_member"][0].IsPartOfTupleCycle())
		require.Empty(t, graph.edges["group#member"][0].GetRecursiveRelation())
		require.True(t, graph.edges["group#member"][0].IsPartOfTupleCycle())
	})

	t.Run("recursive_cycle_userset", func(t *testing.T) {
		model := `
			model
				schema 1.1
			type user
			type group
				relations
					define member: [user, group#member]
		`
		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		require.NoError(t, err)

		require.Len(t, graph.nodes, 3)
		require.Len(t, graph.edges, 1)

		require.Equal(t, "group#member", graph.nodes["group#member"].GetRecursiveRelation())
		require.False(t, graph.nodes["group#member"].IsPartOfTupleCycle())
		require.Empty(t, graph.edges["group#member"][0].GetRecursiveRelation())
		require.Equal(t, "group#member", graph.edges["group#member"][1].GetRecursiveRelation())
		require.False(t, graph.edges["group#member"][0].IsPartOfTupleCycle())
	})

	t.Run("recursive_cycle_ttu", func(t *testing.T) {
		model := `
			model
				schema 1.1
			type user
			type group
				relations
					define member: [user] or member from parent
					define parent: [group]
		`
		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		require.NoError(t, err)

		require.Len(t, graph.nodes, 5)
		require.Len(t, graph.edges, 3)

		var unionNodeID string
		for _, node := range graph.nodes {
			if node.GetNodeType() == OperatorNode {
				unionNodeID = node.GetUniqueLabel()
			}
		}

		require.Equal(t, "group#member", graph.nodes[unionNodeID].GetRecursiveRelation())
		require.False(t, graph.nodes[unionNodeID].IsPartOfTupleCycle())
		require.Empty(t, graph.nodes["group#parent"].GetRecursiveRelation())
		require.Empty(t, graph.edges["group#parent"][0].GetRecursiveRelation())
		require.False(t, graph.nodes["group#member"].IsPartOfTupleCycle())
		require.Equal(t, "group#member", graph.nodes["group#member"].GetRecursiveRelation())
		for _, node := range graph.nodes {
			if node.GetLabel() == "union" {
				require.Equal(t, "group#member", node.GetRecursiveRelation())
			}
		}
		require.Equal(t, "group#member", graph.edges["group#member"][0].GetRecursiveRelation())
	})

	t.Run("both_recursion_and_tuple_cycles", func(t *testing.T) {
		model := `
			model
				schema 1.1
			type user
			type group
				relations
					define inherited_member: member from parent
					define member: [user, group#member] or inherited_member
					define parent: [group]`

		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		require.NoError(t, err)

		var unionNodeID string
		for _, node := range graph.nodes {
			if node.GetNodeType() == OperatorNode {
				unionNodeID = node.GetUniqueLabel()
			}
		}

		require.True(t, graph.nodes[unionNodeID].IsPartOfTupleCycle())
		require.Equal(t, "group#member", graph.nodes[unionNodeID].GetRecursiveRelation())
		require.Equal(t, "group#member", graph.nodes["group#member"].GetRecursiveRelation())
		require.Empty(t, graph.nodes["group#parent"].GetRecursiveRelation())
		require.Empty(t, graph.nodes["group#inherited_member"].GetRecursiveRelation())
		require.True(t, graph.nodes["group#member"].IsPartOfTupleCycle())
		require.True(t, graph.nodes["group#inherited_member"].IsPartOfTupleCycle())
		for _, edge := range graph.edges["group#member"] {
			require.Equal(t, "group#member", edge.GetRecursiveRelation())
		}
	})

	t.Run("both_recursion_and_tuple_cycles_wildcard", func(t *testing.T) {
		model := `
			model
				schema 1.1
			type user
			type group
				relations
					define inherited_member: member from parent
					define member: [user:*, group#member] or inherited_member
					define parent: [group]`

		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		require.NoError(t, err)

		var unionNodeID string
		for _, node := range graph.nodes {
			if node.GetNodeType() == OperatorNode {
				unionNodeID = node.GetUniqueLabel()
			}
		}

		require.True(t, graph.nodes[unionNodeID].IsPartOfTupleCycle())
		require.Equal(t, "group#member", graph.nodes[unionNodeID].GetRecursiveRelation())
		require.Equal(t, "group#member", graph.nodes["group#member"].GetRecursiveRelation())
		require.Empty(t, graph.nodes["group#parent"].GetRecursiveRelation())
		require.Empty(t, graph.nodes["group#inherited_member"].GetRecursiveRelation())
		require.True(t, graph.nodes["group#member"].IsPartOfTupleCycle())
		require.True(t, graph.nodes["group#inherited_member"].IsPartOfTupleCycle())
		for _, edge := range graph.edges["group#member"] {
			require.Equal(t, "group#member", edge.GetRecursiveRelation())
		}
	})
}
