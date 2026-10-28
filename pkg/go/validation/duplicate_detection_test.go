package validation

import (
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
)

// TestOperandName covers the operand shapes the DSL cannot produce, such as a
// tuple-to-userset with one side missing. The shared cases cover the rest.
func TestOperandName(t *testing.T) {
	t.Parallel()

	computed := func(relation string) *openfgav1.Userset {
		return &openfgav1.Userset{Userset: &openfgav1.Userset_ComputedUserset{
			ComputedUserset: &openfgav1.ObjectRelation{Relation: relation},
		}}
	}
	tupleToUserset := func(target, from string) *openfgav1.Userset {
		return &openfgav1.Userset{Userset: &openfgav1.Userset_TupleToUserset{
			TupleToUserset: &openfgav1.TupleToUserset{
				ComputedUserset: &openfgav1.ObjectRelation{Relation: target},
				Tupleset:        &openfgav1.ObjectRelation{Relation: from},
			},
		}}
	}

	tests := []struct {
		name    string
		userset *openfgav1.Userset
		want    string
	}{
		{"computed userset", computed("viewer"), "viewer"},
		{"computed userset with no relation", computed(""), ""},
		{"tuple to userset", tupleToUserset("viewer", "parent"), "viewer from parent"},
		{"tuple to userset with no tupleset relation", tupleToUserset("viewer", ""), "viewer"},
		{"tuple to userset with no computed relation", tupleToUserset("", "parent"), ""},
		{"direct assignment", &openfgav1.Userset{Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}}, ""},
		{"empty userset", &openfgav1.Userset{}, ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.want, operandName(test.userset))
		})
	}
}
