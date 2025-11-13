package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeywordConstants(t *testing.T) {
	// Test that keyword constants are defined correctly
	assert.Equal(t, "self", KeywordSelf)
	assert.Equal(t, "DEFINE", KeywordDefine)
	assert.Equal(t, "this", KeywordThis)
}

func TestReservedKeywords(t *testing.T) {
	// Test that reserved keywords map is properly initialized
	assert.NotNil(t, ReservedKeywords)

	// Test that all expected keywords are in the map
	assert.True(t, ReservedKeywords[KeywordSelf])
	assert.True(t, ReservedKeywords[KeywordThis])

	// Test that non-reserved words are not in the map
	assert.False(t, ReservedKeywords["document"])
	assert.False(t, ReservedKeywords["user"])
	assert.False(t, ReservedKeywords["viewer"])
}

func TestIsReservedKeyword(t *testing.T) {
	tests := []struct {
		name     string
		keyword  string
		expected bool
	}{
		{
			name:     "self is reserved",
			keyword:  "self",
			expected: true,
		},
		{
			name:     "this is reserved",
			keyword:  "this",
			expected: true,
		},
		{
			name:     "document is not reserved",
			keyword:  "document",
			expected: false,
		},
		{
			name:     "user is not reserved",
			keyword:  "user",
			expected: false,
		},
		{
			name:     "viewer is not reserved",
			keyword:  "viewer",
			expected: false,
		},
		{
			name:     "admin is not reserved",
			keyword:  "admin",
			expected: false,
		},
		{
			name:     "empty string is not reserved",
			keyword:  "",
			expected: false,
		},
		{
			name:     "SELF (uppercase) is not reserved",
			keyword:  "SELF",
			expected: false,
		},
		{
			name:     "THIS (uppercase) is not reserved",
			keyword:  "THIS",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsReservedKeyword(tt.keyword)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsReservedTypeName(t *testing.T) {
	tests := []struct {
		name     string
		typeName string
		expected bool
	}{
		{
			name:     "self is reserved type name",
			typeName: "self",
			expected: true,
		},
		{
			name:     "this is reserved type name",
			typeName: "this",
			expected: true,
		},
		{
			name:     "document is not reserved type name",
			typeName: "document",
			expected: false,
		},
		{
			name:     "user is not reserved type name",
			typeName: "user",
			expected: false,
		},
		{
			name:     "folder is not reserved type name",
			typeName: "folder",
			expected: false,
		},
		{
			name:     "group is not reserved type name",
			typeName: "group",
			expected: false,
		},
		{
			name:     "empty string is not reserved type name",
			typeName: "",
			expected: false,
		},
		{
			name:     "SELF (uppercase) is not reserved type name",
			typeName: "SELF",
			expected: false,
		},
		{
			name:     "THIS (uppercase) is not reserved type name",
			typeName: "THIS",
			expected: false,
		},
		{
			name:     "Self (mixed case) is not reserved type name",
			typeName: "Self",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsReservedTypeName(tt.typeName)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsReservedRelationName(t *testing.T) {
	tests := []struct {
		name         string
		relationName string
		expected     bool
	}{
		{
			name:         "self is reserved relation name",
			relationName: "self",
			expected:     true,
		},
		{
			name:         "this is reserved relation name",
			relationName: "this",
			expected:     true,
		},
		{
			name:         "viewer is not reserved relation name",
			relationName: "viewer",
			expected:     false,
		},
		{
			name:         "admin is not reserved relation name",
			relationName: "admin",
			expected:     false,
		},
		{
			name:         "owner is not reserved relation name",
			relationName: "owner",
			expected:     false,
		},
		{
			name:         "member is not reserved relation name",
			relationName: "member",
			expected:     false,
		},
		{
			name:         "can_view is not reserved relation name",
			relationName: "can_view",
			expected:     false,
		},
		{
			name:         "empty string is not reserved relation name",
			relationName: "",
			expected:     false,
		},
		{
			name:         "SELF (uppercase) is not reserved relation name",
			relationName: "SELF",
			expected:     false,
		},
		{
			name:         "THIS (uppercase) is not reserved relation name",
			relationName: "THIS",
			expected:     false,
		},
		{
			name:         "This (mixed case) is not reserved relation name",
			relationName: "This",
			expected:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsReservedRelationName(tt.relationName)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestReservedKeywordsConsistency(t *testing.T) {
	// Test that IsReservedKeyword, IsReservedTypeName, and IsReservedRelationName
	// are consistent with each other for reserved keywords

	reservedKeywords := []string{"self", "this"}

	for _, keyword := range reservedKeywords {
		assert.True(t, IsReservedKeyword(keyword), "IsReservedKeyword should return true for %s", keyword)
		assert.True(t, IsReservedTypeName(keyword), "IsReservedTypeName should return true for %s", keyword)
		assert.True(t, IsReservedRelationName(keyword), "IsReservedRelationName should return true for %s", keyword)
	}

	// Test that they're consistent for non-reserved words
	nonReservedWords := []string{"document", "user", "viewer", "admin", "owner"}

	for _, word := range nonReservedWords {
		assert.False(t, IsReservedKeyword(word), "IsReservedKeyword should return false for %s", word)
		assert.False(t, IsReservedTypeName(word), "IsReservedTypeName should return false for %s", word)
		assert.False(t, IsReservedRelationName(word), "IsReservedRelationName should return false for %s", word)
	}
}

func TestReservedKeywordsMatchJSImplementation(t *testing.T) {
	// Test that our reserved keywords match exactly with the JS implementation
	// Based on the JS keywords.ts file:
	// - Keyword.SELF = "self"
	// - ReservedKeywords.THIS = "this"

	// Test exact keyword values
	assert.Equal(t, "self", KeywordSelf)
	assert.Equal(t, "this", KeywordThis)

	// Test that these are the only reserved keywords for types and relations
	assert.True(t, IsReservedTypeName("self"))
	assert.True(t, IsReservedTypeName("this"))
	assert.True(t, IsReservedRelationName("self"))
	assert.True(t, IsReservedRelationName("this"))

	// Test case sensitivity (JS implementation is case-sensitive)
	assert.False(t, IsReservedTypeName("SELF"))
	assert.False(t, IsReservedTypeName("Self"))
	assert.False(t, IsReservedTypeName("THIS"))
	assert.False(t, IsReservedTypeName("This"))
}

func TestReservedKeywordsValidation(t *testing.T) {
	collector := NewErrorCollector(nil)

	// Test type name validation - should pass for valid names
	isValid := ValidateTypeName("document", collector, nil, nil)
	assert.True(t, isValid)
	assert.Empty(t, collector.GetErrors())

	// Test type name validation - should fail for reserved keywords
	collector = NewErrorCollector(nil)
	isValid = ValidateTypeName("this", collector, nil, nil)
	assert.False(t, isValid)
	assert.NotEmpty(t, collector.GetErrors())

	// Test relation name validation - should pass for valid names
	collector = NewErrorCollector(nil)
	isValid = ValidateRelationName("viewer", "document", collector, nil, nil)
	assert.True(t, isValid)
	assert.Empty(t, collector.GetErrors())

	// Test relation name validation - should fail for reserved keywords
	collector = NewErrorCollector(nil)
	isValid = ValidateRelationName("self", "document", collector, nil, nil)
	assert.False(t, isValid)
	assert.NotEmpty(t, collector.GetErrors())
}
