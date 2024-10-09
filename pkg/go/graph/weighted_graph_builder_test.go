package graph

import (
	"testing"

	language "github.com/openfga/language/pkg/go/transformer"
	"github.com/stretchr/testify/require"
)

func TestWeight1Graph(t *testing.T) {
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
}
