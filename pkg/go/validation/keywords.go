package validation

// Reserved keywords that cannot be used as type or relation names.
const (
	KeywordSelf   = "self"
	KeywordDefine = "DEFINE"
	KeywordThis   = "this"
)

// ReservedKeywords contains all reserved keywords.
var ReservedKeywords = map[string]bool{
	KeywordSelf: true,
	KeywordThis: true,
}

// IsReservedKeyword checks if a given string is a reserved keyword.
func IsReservedKeyword(keyword string) bool {
	return ReservedKeywords[keyword]
}

// IsReservedTypeName checks if a type name is reserved.
func IsReservedTypeName(typeName string) bool {
	return typeName == KeywordSelf || typeName == KeywordThis
}

// IsReservedRelationName checks if a relation name is reserved.
func IsReservedRelationName(relationName string) bool {
	return relationName == KeywordSelf || relationName == KeywordThis
}
