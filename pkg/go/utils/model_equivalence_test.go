package utils_test

import (
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/require"

	"github.com/openfga/language/pkg/go/transformer"
	"github.com/openfga/language/pkg/go/utils"
)

const baseModel = `model
  schema 1.1

type user

type group
  relations
    define member: [user, group#member]

type document
  relations
    define blocked: [user]
    define owner: [user]
    define editor: [user, group#member with in_office] or owner
    define viewer: [user, user:*] or editor or owner but not blocked
    define approver: editor and owner

condition in_office(ip: ipaddress) {
  ip.in_cidr("10.0.0.0/8")
}
`

func TestAreModelsEquivalent(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		dsl        string
		equivalent bool
	}{
		{
			name:       "same model",
			dsl:        baseModel,
			equivalent: true,
		},
		{
			name: "types, relations and conditions in a different order",
			dsl: `model
  schema 1.1

condition in_office(ip: ipaddress) {
  ip.in_cidr("10.0.0.0/8")
}

type document
  relations
    define approver: editor and owner
    define viewer: [user, user:*] or editor or owner but not blocked
    define editor: [user, group#member with in_office] or owner
    define owner: [user]
    define blocked: [user]

type group
  relations
    define member: [user, group#member]

type user
`,
			equivalent: true,
		},
		{
			name: "directly related user types and union and intersection children in a different order",
			dsl: `model
  schema 1.1

type user

type group
  relations
    define member: [group#member, user]

type document
  relations
    define blocked: [user]
    define owner: [user]
    define editor: owner or [group#member with in_office, user]
    define viewer: [user:*, user] or owner or editor but not blocked
    define approver: owner and editor

condition in_office(ip: ipaddress) {
  ip.in_cidr("10.0.0.0/8")
}
`,
			equivalent: true,
		},
		{
			name: "condition expression with different surrounding whitespace",
			dsl: `model
  schema 1.1

type user

type group
  relations
    define member: [user, group#member]

type document
  relations
    define blocked: [user]
    define owner: [user]
    define editor: [user, group#member with in_office] or owner
    define viewer: [user, user:*] or editor or owner but not blocked
    define approver: editor and owner

condition in_office(ip: ipaddress) {

  ip.in_cidr("10.0.0.0/8")

}
`,
			equivalent: true,
		},
		{
			name: "added relation",
			dsl: baseModel[:len(baseModel)-len(`
condition in_office(ip: ipaddress) {
  ip.in_cidr("10.0.0.0/8")
}
`)] + `    define commenter: [user]

condition in_office(ip: ipaddress) {
  ip.in_cidr("10.0.0.0/8")
}
`,
			equivalent: false,
		},
		{
			name: "removed directly related user type",
			dsl: `model
  schema 1.1

type user

type group
  relations
    define member: [user]

type document
  relations
    define blocked: [user]
    define owner: [user]
    define editor: [user, group#member with in_office] or owner
    define viewer: [user, user:*] or editor or owner but not blocked
    define approver: editor and owner

condition in_office(ip: ipaddress) {
  ip.in_cidr("10.0.0.0/8")
}
`,
			equivalent: false,
		},
		{
			name: "union changed to intersection",
			dsl: `model
  schema 1.1

type user

type group
  relations
    define member: [user, group#member]

type document
  relations
    define blocked: [user]
    define owner: [user]
    define editor: [user, group#member with in_office] and owner
    define viewer: [user, user:*] or editor or owner but not blocked
    define approver: editor and owner

condition in_office(ip: ipaddress) {
  ip.in_cidr("10.0.0.0/8")
}
`,
			equivalent: false,
		},
		{
			name: "changed condition expression",
			dsl: `model
  schema 1.1

type user

type group
  relations
    define member: [user, group#member]

type document
  relations
    define blocked: [user]
    define owner: [user]
    define editor: [user, group#member with in_office] or owner
    define viewer: [user, user:*] or editor or owner but not blocked
    define approver: editor and owner

condition in_office(ip: ipaddress) {
  ip.in_cidr("192.168.0.0/16")
}
`,
			equivalent: false,
		},
	}

	base, err := transformer.TransformDSLToProto(baseModel)
	require.NoError(t, err)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			model, err := transformer.TransformDSLToProto(testCase.dsl)
			require.NoError(t, err)

			require.Equal(t, testCase.equivalent, utils.AreModelsEquivalent(base, model))
			require.Equal(t, testCase.equivalent, utils.AreModelsEquivalent(model, base))
		})
	}
}

func TestAreModelsEquivalentIgnoresIDAndSourceInfo(t *testing.T) {
	t.Parallel()

	model1, err := transformer.TransformDSLToProto(baseModel)
	require.NoError(t, err)

	model2, err := transformer.TransformDSLToProto(baseModel)
	require.NoError(t, err)

	model1.Id = "01HVMMBCMGZNT3SED4Z17ECXCA"
	model2.Id = "01HVMMBD123VZ8KV1MPQBNC3EZ"

	for _, typeDef := range model2.GetTypeDefinitions() {
		if typeDef.GetMetadata() == nil {
			typeDef.Metadata = &openfgav1.Metadata{}
		}

		typeDef.Metadata.Module = "documents"
		typeDef.Metadata.SourceInfo = &openfgav1.SourceInfo{File: "documents.fga"}
	}

	require.True(t, utils.AreModelsEquivalent(model1, model2))
	require.NotEqual(t, model1.GetId(), model2.GetId(), "the models passed in must not be changed")
}

func TestAreModelsEquivalentKeepsExclusionOrder(t *testing.T) {
	t.Parallel()

	model1, err := transformer.TransformDSLToProto(`model
  schema 1.1

type user

type document
  relations
    define a: [user]
    define b: [user]
    define viewer: a but not b
`)
	require.NoError(t, err)

	model2, err := transformer.TransformDSLToProto(`model
  schema 1.1

type user

type document
  relations
    define a: [user]
    define b: [user]
    define viewer: b but not a
`)
	require.NoError(t, err)

	require.False(t, utils.AreModelsEquivalent(model1, model2))
}

func TestAreModelsEquivalentFlattensNestedUnions(t *testing.T) {
	t.Parallel()

	model1, err := transformer.TransformDSLToProto(`model
  schema 1.1

type user

type document
  relations
    define a: [user]
    define b: [user]
    define c: [user]
    define viewer: a or (b or c)
`)
	require.NoError(t, err)

	model2, err := transformer.TransformDSLToProto(`model
  schema 1.1

type user

type document
  relations
    define a: [user]
    define b: [user]
    define c: [user]
    define viewer: (a or b) or c
`)
	require.NoError(t, err)

	require.True(t, utils.AreModelsEquivalent(model1, model2))
}

func TestAreModelsEquivalentCollapsesSingleChildRewrites(t *testing.T) {
	t.Parallel()

	model1, err := transformer.TransformDSLToProto(`model
  schema 1.1

type user

type document
  relations
    define a: [user]
    define b: [user]
    define viewer: a or a
    define editor: b and b
`)
	require.NoError(t, err)

	model2, err := transformer.TransformDSLToProto(`model
  schema 1.1

type user

type document
  relations
    define a: [user]
    define b: [user]
    define viewer: a
    define editor: b
`)
	require.NoError(t, err)

	require.True(t, utils.AreModelsEquivalent(model1, model2))
}

func TestAreModelsEquivalentNil(t *testing.T) {
	t.Parallel()

	model, err := transformer.TransformDSLToProto(baseModel)
	require.NoError(t, err)

	require.True(t, utils.AreModelsEquivalent(nil, nil))
	require.False(t, utils.AreModelsEquivalent(model, nil))
	require.False(t, utils.AreModelsEquivalent(nil, model))
}
