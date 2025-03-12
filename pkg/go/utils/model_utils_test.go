package utils

import (
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

func TestGetModuleForObjectTypeRelation(t *testing.T) {
	t.Parallel()

	type args struct {
		typeDef  *openfgav1.TypeDefinition
		relation string
	}

	typeDefWithModule := &openfgav1.TypeDefinition{
		Type: "type1",
		Relations: map[string]*openfgav1.Userset{
			"relation1": {},
			"relation2": {},
			"relation3": {},
			"relation4": {},
		},
		Metadata: &openfgav1.Metadata{
			Module: "type_module1",
			Relations: map[string]*openfgav1.RelationMetadata{
				"relation1": {Module: "module1"},
				"relation2": {Module: ""},
				"relation3": {},
				"relation5": {},
			},
		},
	}

	typeDefWithoutModule := &openfgav1.TypeDefinition{
		Type: "type2",
		Relations: map[string]*openfgav1.Userset{
			"relation7": {},
		},
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Relation exists and has a module",
			args: args{
				typeDef:  typeDefWithModule,
				relation: "relation1",
			},
			want:    "module1",
			wantErr: false,
		},
		{
			name: "Relation exists but has an empty string as a module, type has a module",
			args: args{
				typeDef:  typeDefWithModule,
				relation: "relation2",
			},
			want:    "type_module1",
			wantErr: false,
		},
		{
			name: "Relation exists but does not have a module, type has a module",
			args: args{
				typeDef:  typeDefWithModule,
				relation: "relation3",
			},
			want:    "type_module1",
			wantErr: false,
		},
		{
			name: "Relation exists but does not have metadata, type has a module",
			args: args{
				typeDef:  typeDefWithModule,
				relation: "relation4",
			},
			want:    "type_module1",
			wantErr: false,
		},
		{
			name: "Relation does not exist",
			args: args{
				typeDef:  typeDefWithModule,
				relation: "relation5",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Relation does not exist 2",
			args: args{
				typeDef:  typeDefWithModule,
				relation: "relation6",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Relation exists but does not have a module, type does not have a module",
			args: args{
				typeDef:  typeDefWithoutModule,
				relation: "relation7",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := GetModuleForObjectTypeRelation(tt.args.typeDef, tt.args.relation)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetModuleForObjectTypeRelation() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if got != tt.want {
				t.Errorf("GetModuleForObjectTypeRelation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRelationAssignable(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		relDef   *openfgav1.Userset
		expected bool
	}{
		{
			name:     "Relation definition has a key 'this'",
			relDef:   &openfgav1.Userset{Userset: &openfgav1.Userset_This{}},
			expected: true,
		},
		{
			name: "Relation definition has a key 'union' with a child that has a key 'this'",
			relDef: &openfgav1.Userset{Userset: &openfgav1.Userset_Union{
				Union: &openfgav1.Usersets{
					Child: []*openfgav1.Userset{
						{Userset: &openfgav1.Userset_This{}},
					},
				},
			}},
			expected: true,
		},
		{
			name: "Relation definition has a key 'intersection' with a child that has a key 'this'",
			relDef: &openfgav1.Userset{Userset: &openfgav1.Userset_Intersection{
				Intersection: &openfgav1.Usersets{
					Child: []*openfgav1.Userset{
						{Userset: &openfgav1.Userset_This{}},
					},
				},
			}},
			expected: true,
		},
		{
			name: "Relation definition has a key 'difference' with base having a key 'this'",
			relDef: &openfgav1.Userset{Userset: &openfgav1.Userset_Difference{
				Difference: &openfgav1.Difference{
					Base:     &openfgav1.Userset{Userset: &openfgav1.Userset_This{}},
					Subtract: &openfgav1.Userset{},
				},
			}},
			expected: true,
		},
		{
			name: "Relation definition has a key 'difference' with subtract having a key 'this'",
			relDef: &openfgav1.Userset{Userset: &openfgav1.Userset_Difference{
				Difference: &openfgav1.Difference{
					Base:     &openfgav1.Userset{},
					Subtract: &openfgav1.Userset{Userset: &openfgav1.Userset_This{}},
				},
			}},
			expected: true,
		},
		{
			name: "Relation definition does not have any assignable keys",
			relDef: &openfgav1.Userset{Userset: &openfgav1.Userset_Union{
				Union: &openfgav1.Usersets{
					Child: []*openfgav1.Userset{
						{Userset: &openfgav1.Userset_Intersection{
							Intersection: &openfgav1.Usersets{
								Child: []*openfgav1.Userset{
									{},
								},
							},
						}},
					},
				},
			}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := IsRelationAssignable(tt.relDef)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
