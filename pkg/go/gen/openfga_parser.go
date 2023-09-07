// Code generated from /app/OpenFGAParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // OpenFGAParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type OpenFGAParser struct {
	*antlr.BaseParser
}

var OpenFGAParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func openfgaparserParserInit() {
	staticData := &OpenFGAParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'model'", "'schema'", "'1.1'", "'type'", "'relations'", "'define'",
		"'#'", "':'", "'*'", "'['", "']'", "','", "'and'", "'or'", "'but not'",
		"'from'",
	}
	staticData.SymbolicNames = []string{
		"", "INDENT", "MODEL", "SCHEMA", "SCHEMA_VERSION", "TYPE", "RELATIONS",
		"DEFINE", "HASH", "COLON", "WILDCARD", "L_SQUARE", "R_SQUARE", "COMMA",
		"AND", "OR", "BUT_NOT", "FROM", "ALPHA_NUMERIC", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"main", "indentation", "modelHeader", "typeDefs", "typeDef", "relationDeclaration",
		"relationDef", "relationDefPartials", "relationDefPartialAllOr", "relationDefPartialAllAnd",
		"relationDefPartialAllButNot", "relationDefDirectAssignment", "relationDefRewrite",
		"relationDefRelationOnSameObject", "relationDefRelationOnRelatedObject",
		"relationDefOperator", "relationDefOperatorAnd", "relationDefOperatorOr",
		"relationDefOperatorButNot", "relationDefKeywordFrom", "relationDefTypeRestriction",
		"relationDefTypeRestrictionType", "relationDefTypeRestrictionRelation",
		"relationDefTypeRestrictionWildcard", "relationDefTypeRestrictionUserset",
		"relationDefGrouping", "rewriteComputedusersetName", "rewriteTuplesetComputedusersetName",
		"rewriteTuplesetName", "relationName", "typeName", "comment", "multiLineComment",
		"spacing", "newline", "schemaVersion", "name",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 20, 312, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		1, 0, 1, 0, 1, 0, 3, 0, 78, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 3, 2, 87, 8, 2, 1, 2, 1, 2, 3, 2, 91, 8, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		96, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 103, 8, 2, 1, 3, 5, 3, 106,
		8, 3, 10, 3, 12, 3, 109, 9, 3, 1, 4, 1, 4, 1, 4, 3, 4, 114, 8, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 121, 8, 4, 1, 4, 1, 4, 1, 4, 3, 4, 126, 8,
		4, 1, 4, 4, 4, 129, 8, 4, 11, 4, 12, 4, 130, 3, 4, 133, 8, 4, 1, 5, 1,
		5, 1, 5, 3, 5, 138, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 145, 8, 5,
		1, 5, 1, 5, 3, 5, 149, 8, 5, 1, 5, 1, 5, 3, 5, 153, 8, 5, 1, 6, 1, 6, 3,
		6, 157, 8, 6, 1, 6, 3, 6, 160, 8, 6, 1, 7, 1, 7, 1, 7, 3, 7, 165, 8, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 4, 8, 172, 8, 8, 11, 8, 12, 8, 173, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 4, 9, 181, 8, 9, 11, 9, 12, 9, 182, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 4, 10, 190, 8, 10, 11, 10, 12, 10, 191, 1, 11, 1,
		11, 3, 11, 196, 8, 11, 1, 11, 1, 11, 3, 11, 200, 8, 11, 1, 11, 1, 11, 3,
		11, 204, 8, 11, 1, 11, 5, 11, 207, 8, 11, 10, 11, 12, 11, 210, 9, 11, 1,
		11, 3, 11, 213, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 3, 12, 219, 8, 12, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		3, 15, 232, 8, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 20, 3, 20, 245, 8, 20, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 255, 8, 23, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 31, 5, 31, 274, 8, 31, 10, 31, 12, 31, 277, 9, 31, 1,
		31, 1, 31, 5, 31, 281, 8, 31, 10, 31, 12, 31, 284, 9, 31, 1, 32, 1, 32,
		1, 32, 1, 32, 5, 32, 290, 8, 32, 10, 32, 12, 32, 293, 9, 32, 1, 33, 4,
		33, 296, 8, 33, 11, 33, 12, 33, 297, 1, 34, 4, 34, 301, 8, 34, 11, 34,
		12, 34, 302, 1, 35, 1, 35, 1, 36, 4, 36, 308, 8, 36, 11, 36, 12, 36, 309,
		1, 36, 0, 0, 37, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
		66, 68, 70, 72, 0, 1, 1, 0, 19, 19, 313, 0, 74, 1, 0, 0, 0, 2, 81, 1, 0,
		0, 0, 4, 86, 1, 0, 0, 0, 6, 107, 1, 0, 0, 0, 8, 113, 1, 0, 0, 0, 10, 137,
		1, 0, 0, 0, 12, 156, 1, 0, 0, 0, 14, 164, 1, 0, 0, 0, 16, 171, 1, 0, 0,
		0, 18, 180, 1, 0, 0, 0, 20, 189, 1, 0, 0, 0, 22, 193, 1, 0, 0, 0, 24, 218,
		1, 0, 0, 0, 26, 220, 1, 0, 0, 0, 28, 222, 1, 0, 0, 0, 30, 231, 1, 0, 0,
		0, 32, 233, 1, 0, 0, 0, 34, 235, 1, 0, 0, 0, 36, 237, 1, 0, 0, 0, 38, 239,
		1, 0, 0, 0, 40, 244, 1, 0, 0, 0, 42, 246, 1, 0, 0, 0, 44, 248, 1, 0, 0,
		0, 46, 250, 1, 0, 0, 0, 48, 256, 1, 0, 0, 0, 50, 260, 1, 0, 0, 0, 52, 262,
		1, 0, 0, 0, 54, 264, 1, 0, 0, 0, 56, 266, 1, 0, 0, 0, 58, 268, 1, 0, 0,
		0, 60, 270, 1, 0, 0, 0, 62, 275, 1, 0, 0, 0, 64, 285, 1, 0, 0, 0, 66, 295,
		1, 0, 0, 0, 68, 300, 1, 0, 0, 0, 70, 304, 1, 0, 0, 0, 72, 307, 1, 0, 0,
		0, 74, 75, 3, 4, 2, 0, 75, 77, 3, 6, 3, 0, 76, 78, 3, 68, 34, 0, 77, 76,
		1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 5, 0, 0, 1,
		80, 1, 1, 0, 0, 0, 81, 82, 5, 1, 0, 0, 82, 3, 1, 0, 0, 0, 83, 84, 3, 64,
		32, 0, 84, 85, 3, 68, 34, 0, 85, 87, 1, 0, 0, 0, 86, 83, 1, 0, 0, 0, 86,
		87, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 90, 5, 2, 0, 0, 89, 91, 3, 66,
		33, 0, 90, 89, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 95, 1, 0, 0, 0, 92,
		93, 3, 68, 34, 0, 93, 94, 3, 64, 32, 0, 94, 96, 1, 0, 0, 0, 95, 92, 1,
		0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 98, 3, 2, 1, 0, 98,
		99, 5, 3, 0, 0, 99, 100, 3, 66, 33, 0, 100, 102, 3, 70, 35, 0, 101, 103,
		3, 66, 33, 0, 102, 101, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 5, 1, 0,
		0, 0, 104, 106, 3, 8, 4, 0, 105, 104, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0,
		107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 7, 1, 0, 0, 0, 109, 107,
		1, 0, 0, 0, 110, 111, 3, 68, 34, 0, 111, 112, 3, 64, 32, 0, 112, 114, 1,
		0, 0, 0, 113, 110, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 1, 0, 0,
		0, 115, 116, 3, 68, 34, 0, 116, 117, 5, 5, 0, 0, 117, 118, 3, 66, 33, 0,
		118, 120, 3, 60, 30, 0, 119, 121, 3, 66, 33, 0, 120, 119, 1, 0, 0, 0, 120,
		121, 1, 0, 0, 0, 121, 132, 1, 0, 0, 0, 122, 123, 3, 2, 1, 0, 123, 125,
		5, 6, 0, 0, 124, 126, 3, 66, 33, 0, 125, 124, 1, 0, 0, 0, 125, 126, 1,
		0, 0, 0, 126, 128, 1, 0, 0, 0, 127, 129, 3, 10, 5, 0, 128, 127, 1, 0, 0,
		0, 129, 130, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131,
		133, 1, 0, 0, 0, 132, 122, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 9, 1,
		0, 0, 0, 134, 135, 3, 68, 34, 0, 135, 136, 3, 64, 32, 0, 136, 138, 1, 0,
		0, 0, 137, 134, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0,
		139, 140, 3, 2, 1, 0, 140, 141, 5, 7, 0, 0, 141, 142, 3, 66, 33, 0, 142,
		144, 3, 58, 29, 0, 143, 145, 3, 66, 33, 0, 144, 143, 1, 0, 0, 0, 144, 145,
		1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 148, 5, 9, 0, 0, 147, 149, 3, 66,
		33, 0, 148, 147, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0,
		150, 152, 3, 12, 6, 0, 151, 153, 3, 66, 33, 0, 152, 151, 1, 0, 0, 0, 152,
		153, 1, 0, 0, 0, 153, 11, 1, 0, 0, 0, 154, 157, 3, 22, 11, 0, 155, 157,
		3, 50, 25, 0, 156, 154, 1, 0, 0, 0, 156, 155, 1, 0, 0, 0, 157, 159, 1,
		0, 0, 0, 158, 160, 3, 14, 7, 0, 159, 158, 1, 0, 0, 0, 159, 160, 1, 0, 0,
		0, 160, 13, 1, 0, 0, 0, 161, 165, 3, 16, 8, 0, 162, 165, 3, 18, 9, 0, 163,
		165, 3, 20, 10, 0, 164, 161, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 163,
		1, 0, 0, 0, 165, 15, 1, 0, 0, 0, 166, 167, 3, 66, 33, 0, 167, 168, 3, 34,
		17, 0, 168, 169, 3, 66, 33, 0, 169, 170, 3, 50, 25, 0, 170, 172, 1, 0,
		0, 0, 171, 166, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0,
		173, 174, 1, 0, 0, 0, 174, 17, 1, 0, 0, 0, 175, 176, 3, 66, 33, 0, 176,
		177, 3, 32, 16, 0, 177, 178, 3, 66, 33, 0, 178, 179, 3, 50, 25, 0, 179,
		181, 1, 0, 0, 0, 180, 175, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 180,
		1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 19, 1, 0, 0, 0, 184, 185, 3, 66,
		33, 0, 185, 186, 3, 36, 18, 0, 186, 187, 3, 66, 33, 0, 187, 188, 3, 50,
		25, 0, 188, 190, 1, 0, 0, 0, 189, 184, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0,
		191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 21, 1, 0, 0, 0, 193, 195,
		5, 11, 0, 0, 194, 196, 3, 66, 33, 0, 195, 194, 1, 0, 0, 0, 195, 196, 1,
		0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 199, 3, 40, 20, 0, 198, 200, 3, 66,
		33, 0, 199, 198, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 208, 1, 0, 0, 0,
		201, 203, 5, 13, 0, 0, 202, 204, 3, 66, 33, 0, 203, 202, 1, 0, 0, 0, 203,
		204, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 207, 3, 40, 20, 0, 206, 201,
		1, 0, 0, 0, 207, 210, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0,
		0, 0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 211, 213, 3, 66, 33,
		0, 212, 211, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214,
		215, 5, 12, 0, 0, 215, 23, 1, 0, 0, 0, 216, 219, 3, 26, 13, 0, 217, 219,
		3, 28, 14, 0, 218, 216, 1, 0, 0, 0, 218, 217, 1, 0, 0, 0, 219, 25, 1, 0,
		0, 0, 220, 221, 3, 52, 26, 0, 221, 27, 1, 0, 0, 0, 222, 223, 3, 54, 27,
		0, 223, 224, 3, 66, 33, 0, 224, 225, 3, 38, 19, 0, 225, 226, 3, 66, 33,
		0, 226, 227, 3, 56, 28, 0, 227, 29, 1, 0, 0, 0, 228, 232, 3, 34, 17, 0,
		229, 232, 3, 32, 16, 0, 230, 232, 3, 36, 18, 0, 231, 228, 1, 0, 0, 0, 231,
		229, 1, 0, 0, 0, 231, 230, 1, 0, 0, 0, 232, 31, 1, 0, 0, 0, 233, 234, 5,
		14, 0, 0, 234, 33, 1, 0, 0, 0, 235, 236, 5, 15, 0, 0, 236, 35, 1, 0, 0,
		0, 237, 238, 5, 16, 0, 0, 238, 37, 1, 0, 0, 0, 239, 240, 5, 17, 0, 0, 240,
		39, 1, 0, 0, 0, 241, 245, 3, 42, 21, 0, 242, 245, 3, 46, 23, 0, 243, 245,
		3, 48, 24, 0, 244, 241, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 244, 243, 1,
		0, 0, 0, 245, 41, 1, 0, 0, 0, 246, 247, 3, 72, 36, 0, 247, 43, 1, 0, 0,
		0, 248, 249, 3, 72, 36, 0, 249, 45, 1, 0, 0, 0, 250, 251, 3, 42, 21, 0,
		251, 252, 5, 9, 0, 0, 252, 254, 5, 10, 0, 0, 253, 255, 3, 66, 33, 0, 254,
		253, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 47, 1, 0, 0, 0, 256, 257, 3,
		42, 21, 0, 257, 258, 5, 8, 0, 0, 258, 259, 3, 44, 22, 0, 259, 49, 1, 0,
		0, 0, 260, 261, 3, 24, 12, 0, 261, 51, 1, 0, 0, 0, 262, 263, 3, 72, 36,
		0, 263, 53, 1, 0, 0, 0, 264, 265, 3, 72, 36, 0, 265, 55, 1, 0, 0, 0, 266,
		267, 3, 72, 36, 0, 267, 57, 1, 0, 0, 0, 268, 269, 3, 72, 36, 0, 269, 59,
		1, 0, 0, 0, 270, 271, 3, 72, 36, 0, 271, 61, 1, 0, 0, 0, 272, 274, 5, 20,
		0, 0, 273, 272, 1, 0, 0, 0, 274, 277, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0,
		275, 276, 1, 0, 0, 0, 276, 278, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 278,
		282, 5, 8, 0, 0, 279, 281, 8, 0, 0, 0, 280, 279, 1, 0, 0, 0, 281, 284,
		1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 63, 1, 0,
		0, 0, 284, 282, 1, 0, 0, 0, 285, 291, 3, 62, 31, 0, 286, 287, 3, 68, 34,
		0, 287, 288, 3, 62, 31, 0, 288, 290, 1, 0, 0, 0, 289, 286, 1, 0, 0, 0,
		290, 293, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292,
		65, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 294, 296, 5, 20, 0, 0, 295, 294,
		1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 297, 298, 1, 0,
		0, 0, 298, 67, 1, 0, 0, 0, 299, 301, 5, 19, 0, 0, 300, 299, 1, 0, 0, 0,
		301, 302, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303,
		69, 1, 0, 0, 0, 304, 305, 5, 4, 0, 0, 305, 71, 1, 0, 0, 0, 306, 308, 5,
		18, 0, 0, 307, 306, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 307, 1, 0, 0,
		0, 309, 310, 1, 0, 0, 0, 310, 73, 1, 0, 0, 0, 36, 77, 86, 90, 95, 102,
		107, 113, 120, 125, 130, 132, 137, 144, 148, 152, 156, 159, 164, 173, 182,
		191, 195, 199, 203, 208, 212, 218, 231, 244, 254, 275, 282, 291, 297, 302,
		309,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// OpenFGAParserInit initializes any static state used to implement OpenFGAParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOpenFGAParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OpenFGAParserInit() {
	staticData := &OpenFGAParserParserStaticData
	staticData.once.Do(openfgaparserParserInit)
}

// NewOpenFGAParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOpenFGAParser(input antlr.TokenStream) *OpenFGAParser {
	OpenFGAParserInit()
	this := new(OpenFGAParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &OpenFGAParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "OpenFGAParser.g4"

	return this
}

// OpenFGAParser tokens.
const (
	OpenFGAParserEOF            = antlr.TokenEOF
	OpenFGAParserINDENT         = 1
	OpenFGAParserMODEL          = 2
	OpenFGAParserSCHEMA         = 3
	OpenFGAParserSCHEMA_VERSION = 4
	OpenFGAParserTYPE           = 5
	OpenFGAParserRELATIONS      = 6
	OpenFGAParserDEFINE         = 7
	OpenFGAParserHASH           = 8
	OpenFGAParserCOLON          = 9
	OpenFGAParserWILDCARD       = 10
	OpenFGAParserL_SQUARE       = 11
	OpenFGAParserR_SQUARE       = 12
	OpenFGAParserCOMMA          = 13
	OpenFGAParserAND            = 14
	OpenFGAParserOR             = 15
	OpenFGAParserBUT_NOT        = 16
	OpenFGAParserFROM           = 17
	OpenFGAParserALPHA_NUMERIC  = 18
	OpenFGAParserNEWLINE        = 19
	OpenFGAParserWS             = 20
)

// OpenFGAParser rules.
const (
	OpenFGAParserRULE_main                               = 0
	OpenFGAParserRULE_indentation                        = 1
	OpenFGAParserRULE_modelHeader                        = 2
	OpenFGAParserRULE_typeDefs                           = 3
	OpenFGAParserRULE_typeDef                            = 4
	OpenFGAParserRULE_relationDeclaration                = 5
	OpenFGAParserRULE_relationDef                        = 6
	OpenFGAParserRULE_relationDefPartials                = 7
	OpenFGAParserRULE_relationDefPartialAllOr            = 8
	OpenFGAParserRULE_relationDefPartialAllAnd           = 9
	OpenFGAParserRULE_relationDefPartialAllButNot        = 10
	OpenFGAParserRULE_relationDefDirectAssignment        = 11
	OpenFGAParserRULE_relationDefRewrite                 = 12
	OpenFGAParserRULE_relationDefRelationOnSameObject    = 13
	OpenFGAParserRULE_relationDefRelationOnRelatedObject = 14
	OpenFGAParserRULE_relationDefOperator                = 15
	OpenFGAParserRULE_relationDefOperatorAnd             = 16
	OpenFGAParserRULE_relationDefOperatorOr              = 17
	OpenFGAParserRULE_relationDefOperatorButNot          = 18
	OpenFGAParserRULE_relationDefKeywordFrom             = 19
	OpenFGAParserRULE_relationDefTypeRestriction         = 20
	OpenFGAParserRULE_relationDefTypeRestrictionType     = 21
	OpenFGAParserRULE_relationDefTypeRestrictionRelation = 22
	OpenFGAParserRULE_relationDefTypeRestrictionWildcard = 23
	OpenFGAParserRULE_relationDefTypeRestrictionUserset  = 24
	OpenFGAParserRULE_relationDefGrouping                = 25
	OpenFGAParserRULE_rewriteComputedusersetName         = 26
	OpenFGAParserRULE_rewriteTuplesetComputedusersetName = 27
	OpenFGAParserRULE_rewriteTuplesetName                = 28
	OpenFGAParserRULE_relationName                       = 29
	OpenFGAParserRULE_typeName                           = 30
	OpenFGAParserRULE_comment                            = 31
	OpenFGAParserRULE_multiLineComment                   = 32
	OpenFGAParserRULE_spacing                            = 33
	OpenFGAParserRULE_newline                            = 34
	OpenFGAParserRULE_schemaVersion                      = 35
	OpenFGAParserRULE_name                               = 36
)

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ModelHeader() IModelHeaderContext
	TypeDefs() ITypeDefsContext
	EOF() antlr.TerminalNode
	Newline() INewlineContext

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_main
	return p
}

func InitEmptyMainContext(p *MainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_main
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) ModelHeader() IModelHeaderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModelHeaderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModelHeaderContext)
}

func (s *MainContext) TypeDefs() ITypeDefsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefsContext)
}

func (s *MainContext) EOF() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserEOF, 0)
}

func (s *MainContext) Newline() INewlineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INewlineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INewlineContext)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitMain(s)
	}
}

func (p *OpenFGAParser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OpenFGAParserRULE_main)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.ModelHeader()
	}
	{
		p.SetState(75)
		p.TypeDefs()
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(76)
			p.Newline()
		}

	}
	{
		p.SetState(79)
		p.Match(OpenFGAParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndentationContext is an interface to support dynamic dispatch.
type IIndentationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INDENT() antlr.TerminalNode

	// IsIndentationContext differentiates from other interfaces.
	IsIndentationContext()
}

type IndentationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndentationContext() *IndentationContext {
	var p = new(IndentationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_indentation
	return p
}

func InitEmptyIndentationContext(p *IndentationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_indentation
}

func (*IndentationContext) IsIndentationContext() {}

func NewIndentationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndentationContext {
	var p = new(IndentationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_indentation

	return p
}

func (s *IndentationContext) GetParser() antlr.Parser { return s.parser }

func (s *IndentationContext) INDENT() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserINDENT, 0)
}

func (s *IndentationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndentationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndentationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterIndentation(s)
	}
}

func (s *IndentationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitIndentation(s)
	}
}

func (p *OpenFGAParser) Indentation() (localctx IIndentationContext) {
	localctx = NewIndentationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OpenFGAParserRULE_indentation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(OpenFGAParserINDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IModelHeaderContext is an interface to support dynamic dispatch.
type IModelHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MODEL() antlr.TerminalNode
	Indentation() IIndentationContext
	SCHEMA() antlr.TerminalNode
	AllSpacing() []ISpacingContext
	Spacing(i int) ISpacingContext
	SchemaVersion() ISchemaVersionContext
	AllMultiLineComment() []IMultiLineCommentContext
	MultiLineComment(i int) IMultiLineCommentContext
	AllNewline() []INewlineContext
	Newline(i int) INewlineContext

	// IsModelHeaderContext differentiates from other interfaces.
	IsModelHeaderContext()
}

type ModelHeaderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelHeaderContext() *ModelHeaderContext {
	var p = new(ModelHeaderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_modelHeader
	return p
}

func InitEmptyModelHeaderContext(p *ModelHeaderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_modelHeader
}

func (*ModelHeaderContext) IsModelHeaderContext() {}

func NewModelHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelHeaderContext {
	var p = new(ModelHeaderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_modelHeader

	return p
}

func (s *ModelHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelHeaderContext) MODEL() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserMODEL, 0)
}

func (s *ModelHeaderContext) Indentation() IIndentationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndentationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndentationContext)
}

func (s *ModelHeaderContext) SCHEMA() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserSCHEMA, 0)
}

func (s *ModelHeaderContext) AllSpacing() []ISpacingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISpacingContext); ok {
			len++
		}
	}

	tst := make([]ISpacingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISpacingContext); ok {
			tst[i] = t.(ISpacingContext)
			i++
		}
	}

	return tst
}

func (s *ModelHeaderContext) Spacing(i int) ISpacingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpacingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpacingContext)
}

func (s *ModelHeaderContext) SchemaVersion() ISchemaVersionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISchemaVersionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISchemaVersionContext)
}

func (s *ModelHeaderContext) AllMultiLineComment() []IMultiLineCommentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiLineCommentContext); ok {
			len++
		}
	}

	tst := make([]IMultiLineCommentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiLineCommentContext); ok {
			tst[i] = t.(IMultiLineCommentContext)
			i++
		}
	}

	return tst
}

func (s *ModelHeaderContext) MultiLineComment(i int) IMultiLineCommentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiLineCommentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiLineCommentContext)
}

func (s *ModelHeaderContext) AllNewline() []INewlineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INewlineContext); ok {
			len++
		}
	}

	tst := make([]INewlineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INewlineContext); ok {
			tst[i] = t.(INewlineContext)
			i++
		}
	}

	return tst
}

func (s *ModelHeaderContext) Newline(i int) INewlineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INewlineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INewlineContext)
}

func (s *ModelHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterModelHeader(s)
	}
}

func (s *ModelHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitModelHeader(s)
	}
}

func (p *OpenFGAParser) ModelHeader() (localctx IModelHeaderContext) {
	localctx = NewModelHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, OpenFGAParserRULE_modelHeader)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserHASH || _la == OpenFGAParserWS {
		{
			p.SetState(83)
			p.MultiLineComment()
		}
		{
			p.SetState(84)
			p.Newline()
		}

	}
	{
		p.SetState(88)
		p.Match(OpenFGAParserMODEL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserWS {
		{
			p.SetState(89)
			p.Spacing()
		}

	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(92)
			p.Newline()
		}
		{
			p.SetState(93)
			p.MultiLineComment()
		}

	}
	{
		p.SetState(97)
		p.Indentation()
	}
	{
		p.SetState(98)
		p.Match(OpenFGAParserSCHEMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.Spacing()
	}
	{
		p.SetState(100)
		p.SchemaVersion()
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserWS {
		{
			p.SetState(101)
			p.Spacing()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeDefsContext is an interface to support dynamic dispatch.
type ITypeDefsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTypeDef() []ITypeDefContext
	TypeDef(i int) ITypeDefContext

	// IsTypeDefsContext differentiates from other interfaces.
	IsTypeDefsContext()
}

type TypeDefsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefsContext() *TypeDefsContext {
	var p = new(TypeDefsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_typeDefs
	return p
}

func InitEmptyTypeDefsContext(p *TypeDefsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_typeDefs
}

func (*TypeDefsContext) IsTypeDefsContext() {}

func NewTypeDefsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefsContext {
	var p = new(TypeDefsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_typeDefs

	return p
}

func (s *TypeDefsContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefsContext) AllTypeDef() []ITypeDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDefContext); ok {
			len++
		}
	}

	tst := make([]ITypeDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDefContext); ok {
			tst[i] = t.(ITypeDefContext)
			i++
		}
	}

	return tst
}

func (s *TypeDefsContext) TypeDef(i int) ITypeDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *TypeDefsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterTypeDefs(s)
	}
}

func (s *TypeDefsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitTypeDefs(s)
	}
}

func (p *OpenFGAParser) TypeDefs() (localctx ITypeDefsContext) {
	localctx = NewTypeDefsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, OpenFGAParserRULE_typeDefs)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(104)
				p.TypeDef()
			}

		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeDefContext is an interface to support dynamic dispatch.
type ITypeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNewline() []INewlineContext
	Newline(i int) INewlineContext
	TYPE() antlr.TerminalNode
	AllSpacing() []ISpacingContext
	Spacing(i int) ISpacingContext
	TypeName() ITypeNameContext
	MultiLineComment() IMultiLineCommentContext
	Indentation() IIndentationContext
	RELATIONS() antlr.TerminalNode
	AllRelationDeclaration() []IRelationDeclarationContext
	RelationDeclaration(i int) IRelationDeclarationContext

	// IsTypeDefContext differentiates from other interfaces.
	IsTypeDefContext()
}

type TypeDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefContext() *TypeDefContext {
	var p = new(TypeDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_typeDef
	return p
}

func InitEmptyTypeDefContext(p *TypeDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_typeDef
}

func (*TypeDefContext) IsTypeDefContext() {}

func NewTypeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefContext {
	var p = new(TypeDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_typeDef

	return p
}

func (s *TypeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefContext) AllNewline() []INewlineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INewlineContext); ok {
			len++
		}
	}

	tst := make([]INewlineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INewlineContext); ok {
			tst[i] = t.(INewlineContext)
			i++
		}
	}

	return tst
}

func (s *TypeDefContext) Newline(i int) INewlineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INewlineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INewlineContext)
}

func (s *TypeDefContext) TYPE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserTYPE, 0)
}

func (s *TypeDefContext) AllSpacing() []ISpacingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISpacingContext); ok {
			len++
		}
	}

	tst := make([]ISpacingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISpacingContext); ok {
			tst[i] = t.(ISpacingContext)
			i++
		}
	}

	return tst
}

func (s *TypeDefContext) Spacing(i int) ISpacingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpacingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpacingContext)
}

func (s *TypeDefContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *TypeDefContext) MultiLineComment() IMultiLineCommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiLineCommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiLineCommentContext)
}

func (s *TypeDefContext) Indentation() IIndentationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndentationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndentationContext)
}

func (s *TypeDefContext) RELATIONS() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserRELATIONS, 0)
}

func (s *TypeDefContext) AllRelationDeclaration() []IRelationDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IRelationDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationDeclarationContext); ok {
			tst[i] = t.(IRelationDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *TypeDefContext) RelationDeclaration(i int) IRelationDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDeclarationContext)
}

func (s *TypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterTypeDef(s)
	}
}

func (s *TypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitTypeDef(s)
	}
}

func (p *OpenFGAParser) TypeDef() (localctx ITypeDefContext) {
	localctx = NewTypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, OpenFGAParserRULE_typeDef)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(110)
			p.Newline()
		}
		{
			p.SetState(111)
			p.MultiLineComment()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(115)
		p.Newline()
	}
	{
		p.SetState(116)
		p.Match(OpenFGAParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Spacing()
	}
	{
		p.SetState(118)
		p.TypeName()
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserWS {
		{
			p.SetState(119)
			p.Spacing()
		}

	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserINDENT {
		{
			p.SetState(122)
			p.Indentation()
		}
		{
			p.SetState(123)
			p.Match(OpenFGAParserRELATIONS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == OpenFGAParserWS {
			{
				p.SetState(124)
				p.Spacing()
			}

		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(127)
					p.RelationDeclaration()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(130)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDeclarationContext is an interface to support dynamic dispatch.
type IRelationDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Indentation() IIndentationContext
	DEFINE() antlr.TerminalNode
	AllSpacing() []ISpacingContext
	Spacing(i int) ISpacingContext
	RelationName() IRelationNameContext
	COLON() antlr.TerminalNode
	RelationDef() IRelationDefContext
	Newline() INewlineContext
	MultiLineComment() IMultiLineCommentContext

	// IsRelationDeclarationContext differentiates from other interfaces.
	IsRelationDeclarationContext()
}

type RelationDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDeclarationContext() *RelationDeclarationContext {
	var p = new(RelationDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDeclaration
	return p
}

func InitEmptyRelationDeclarationContext(p *RelationDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDeclaration
}

func (*RelationDeclarationContext) IsRelationDeclarationContext() {}

func NewRelationDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDeclarationContext {
	var p = new(RelationDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDeclaration

	return p
}

func (s *RelationDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDeclarationContext) Indentation() IIndentationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndentationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndentationContext)
}

func (s *RelationDeclarationContext) DEFINE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserDEFINE, 0)
}

func (s *RelationDeclarationContext) AllSpacing() []ISpacingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISpacingContext); ok {
			len++
		}
	}

	tst := make([]ISpacingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISpacingContext); ok {
			tst[i] = t.(ISpacingContext)
			i++
		}
	}

	return tst
}

func (s *RelationDeclarationContext) Spacing(i int) ISpacingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpacingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpacingContext)
}

func (s *RelationDeclarationContext) RelationName() IRelationNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationNameContext)
}

func (s *RelationDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserCOLON, 0)
}

func (s *RelationDeclarationContext) RelationDef() IRelationDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefContext)
}

func (s *RelationDeclarationContext) Newline() INewlineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INewlineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INewlineContext)
}

func (s *RelationDeclarationContext) MultiLineComment() IMultiLineCommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiLineCommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiLineCommentContext)
}

func (s *RelationDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDeclaration(s)
	}
}

func (s *RelationDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDeclaration(s)
	}
}

func (p *OpenFGAParser) RelationDeclaration() (localctx IRelationDeclarationContext) {
	localctx = NewRelationDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, OpenFGAParserRULE_relationDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(134)
			p.Newline()
		}
		{
			p.SetState(135)
			p.MultiLineComment()
		}

	}
	{
		p.SetState(139)
		p.Indentation()
	}
	{
		p.SetState(140)
		p.Match(OpenFGAParserDEFINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Spacing()
	}
	{
		p.SetState(142)
		p.RelationName()
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserWS {
		{
			p.SetState(143)
			p.Spacing()
		}

	}
	{
		p.SetState(146)
		p.Match(OpenFGAParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserWS {
		{
			p.SetState(147)
			p.Spacing()
		}

	}
	{
		p.SetState(150)
		p.RelationDef()
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserWS {
		{
			p.SetState(151)
			p.Spacing()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefContext is an interface to support dynamic dispatch.
type IRelationDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationDefDirectAssignment() IRelationDefDirectAssignmentContext
	RelationDefGrouping() IRelationDefGroupingContext
	RelationDefPartials() IRelationDefPartialsContext

	// IsRelationDefContext differentiates from other interfaces.
	IsRelationDefContext()
}

type RelationDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefContext() *RelationDefContext {
	var p = new(RelationDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDef
	return p
}

func InitEmptyRelationDefContext(p *RelationDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDef
}

func (*RelationDefContext) IsRelationDefContext() {}

func NewRelationDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefContext {
	var p = new(RelationDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDef

	return p
}

func (s *RelationDefContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefContext) RelationDefDirectAssignment() IRelationDefDirectAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefDirectAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefDirectAssignmentContext)
}

func (s *RelationDefContext) RelationDefGrouping() IRelationDefGroupingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefGroupingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefGroupingContext)
}

func (s *RelationDefContext) RelationDefPartials() IRelationDefPartialsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefPartialsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefPartialsContext)
}

func (s *RelationDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDef(s)
	}
}

func (s *RelationDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDef(s)
	}
}

func (p *OpenFGAParser) RelationDef() (localctx IRelationDefContext) {
	localctx = NewRelationDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, OpenFGAParserRULE_relationDef)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case OpenFGAParserL_SQUARE:
		{
			p.SetState(154)
			p.RelationDefDirectAssignment()
		}

	case OpenFGAParserALPHA_NUMERIC:
		{
			p.SetState(155)
			p.RelationDefGrouping()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(158)
			p.RelationDefPartials()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefPartialsContext is an interface to support dynamic dispatch.
type IRelationDefPartialsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationDefPartialAllOr() IRelationDefPartialAllOrContext
	RelationDefPartialAllAnd() IRelationDefPartialAllAndContext
	RelationDefPartialAllButNot() IRelationDefPartialAllButNotContext

	// IsRelationDefPartialsContext differentiates from other interfaces.
	IsRelationDefPartialsContext()
}

type RelationDefPartialsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefPartialsContext() *RelationDefPartialsContext {
	var p = new(RelationDefPartialsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefPartials
	return p
}

func InitEmptyRelationDefPartialsContext(p *RelationDefPartialsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefPartials
}

func (*RelationDefPartialsContext) IsRelationDefPartialsContext() {}

func NewRelationDefPartialsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefPartialsContext {
	var p = new(RelationDefPartialsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefPartials

	return p
}

func (s *RelationDefPartialsContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefPartialsContext) RelationDefPartialAllOr() IRelationDefPartialAllOrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefPartialAllOrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefPartialAllOrContext)
}

func (s *RelationDefPartialsContext) RelationDefPartialAllAnd() IRelationDefPartialAllAndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefPartialAllAndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefPartialAllAndContext)
}

func (s *RelationDefPartialsContext) RelationDefPartialAllButNot() IRelationDefPartialAllButNotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefPartialAllButNotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefPartialAllButNotContext)
}

func (s *RelationDefPartialsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefPartialsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefPartialsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefPartials(s)
	}
}

func (s *RelationDefPartialsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefPartials(s)
	}
}

func (p *OpenFGAParser) RelationDefPartials() (localctx IRelationDefPartialsContext) {
	localctx = NewRelationDefPartialsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, OpenFGAParserRULE_relationDefPartials)
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.RelationDefPartialAllOr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.RelationDefPartialAllAnd()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(163)
			p.RelationDefPartialAllButNot()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefPartialAllOrContext is an interface to support dynamic dispatch.
type IRelationDefPartialAllOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSpacing() []ISpacingContext
	Spacing(i int) ISpacingContext
	AllRelationDefOperatorOr() []IRelationDefOperatorOrContext
	RelationDefOperatorOr(i int) IRelationDefOperatorOrContext
	AllRelationDefGrouping() []IRelationDefGroupingContext
	RelationDefGrouping(i int) IRelationDefGroupingContext

	// IsRelationDefPartialAllOrContext differentiates from other interfaces.
	IsRelationDefPartialAllOrContext()
}

type RelationDefPartialAllOrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefPartialAllOrContext() *RelationDefPartialAllOrContext {
	var p = new(RelationDefPartialAllOrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefPartialAllOr
	return p
}

func InitEmptyRelationDefPartialAllOrContext(p *RelationDefPartialAllOrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefPartialAllOr
}

func (*RelationDefPartialAllOrContext) IsRelationDefPartialAllOrContext() {}

func NewRelationDefPartialAllOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefPartialAllOrContext {
	var p = new(RelationDefPartialAllOrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefPartialAllOr

	return p
}

func (s *RelationDefPartialAllOrContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefPartialAllOrContext) AllSpacing() []ISpacingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISpacingContext); ok {
			len++
		}
	}

	tst := make([]ISpacingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISpacingContext); ok {
			tst[i] = t.(ISpacingContext)
			i++
		}
	}

	return tst
}

func (s *RelationDefPartialAllOrContext) Spacing(i int) ISpacingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpacingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpacingContext)
}

func (s *RelationDefPartialAllOrContext) AllRelationDefOperatorOr() []IRelationDefOperatorOrContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationDefOperatorOrContext); ok {
			len++
		}
	}

	tst := make([]IRelationDefOperatorOrContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationDefOperatorOrContext); ok {
			tst[i] = t.(IRelationDefOperatorOrContext)
			i++
		}
	}

	return tst
}

func (s *RelationDefPartialAllOrContext) RelationDefOperatorOr(i int) IRelationDefOperatorOrContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefOperatorOrContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefOperatorOrContext)
}

func (s *RelationDefPartialAllOrContext) AllRelationDefGrouping() []IRelationDefGroupingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationDefGroupingContext); ok {
			len++
		}
	}

	tst := make([]IRelationDefGroupingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationDefGroupingContext); ok {
			tst[i] = t.(IRelationDefGroupingContext)
			i++
		}
	}

	return tst
}

func (s *RelationDefPartialAllOrContext) RelationDefGrouping(i int) IRelationDefGroupingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefGroupingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefGroupingContext)
}

func (s *RelationDefPartialAllOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefPartialAllOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefPartialAllOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefPartialAllOr(s)
	}
}

func (s *RelationDefPartialAllOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefPartialAllOr(s)
	}
}

func (p *OpenFGAParser) RelationDefPartialAllOr() (localctx IRelationDefPartialAllOrContext) {
	localctx = NewRelationDefPartialAllOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, OpenFGAParserRULE_relationDefPartialAllOr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(166)
				p.Spacing()
			}
			{
				p.SetState(167)
				p.RelationDefOperatorOr()
			}
			{
				p.SetState(168)
				p.Spacing()
			}
			{
				p.SetState(169)
				p.RelationDefGrouping()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefPartialAllAndContext is an interface to support dynamic dispatch.
type IRelationDefPartialAllAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSpacing() []ISpacingContext
	Spacing(i int) ISpacingContext
	AllRelationDefOperatorAnd() []IRelationDefOperatorAndContext
	RelationDefOperatorAnd(i int) IRelationDefOperatorAndContext
	AllRelationDefGrouping() []IRelationDefGroupingContext
	RelationDefGrouping(i int) IRelationDefGroupingContext

	// IsRelationDefPartialAllAndContext differentiates from other interfaces.
	IsRelationDefPartialAllAndContext()
}

type RelationDefPartialAllAndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefPartialAllAndContext() *RelationDefPartialAllAndContext {
	var p = new(RelationDefPartialAllAndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefPartialAllAnd
	return p
}

func InitEmptyRelationDefPartialAllAndContext(p *RelationDefPartialAllAndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefPartialAllAnd
}

func (*RelationDefPartialAllAndContext) IsRelationDefPartialAllAndContext() {}

func NewRelationDefPartialAllAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefPartialAllAndContext {
	var p = new(RelationDefPartialAllAndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefPartialAllAnd

	return p
}

func (s *RelationDefPartialAllAndContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefPartialAllAndContext) AllSpacing() []ISpacingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISpacingContext); ok {
			len++
		}
	}

	tst := make([]ISpacingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISpacingContext); ok {
			tst[i] = t.(ISpacingContext)
			i++
		}
	}

	return tst
}

func (s *RelationDefPartialAllAndContext) Spacing(i int) ISpacingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpacingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpacingContext)
}

func (s *RelationDefPartialAllAndContext) AllRelationDefOperatorAnd() []IRelationDefOperatorAndContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationDefOperatorAndContext); ok {
			len++
		}
	}

	tst := make([]IRelationDefOperatorAndContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationDefOperatorAndContext); ok {
			tst[i] = t.(IRelationDefOperatorAndContext)
			i++
		}
	}

	return tst
}

func (s *RelationDefPartialAllAndContext) RelationDefOperatorAnd(i int) IRelationDefOperatorAndContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefOperatorAndContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefOperatorAndContext)
}

func (s *RelationDefPartialAllAndContext) AllRelationDefGrouping() []IRelationDefGroupingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationDefGroupingContext); ok {
			len++
		}
	}

	tst := make([]IRelationDefGroupingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationDefGroupingContext); ok {
			tst[i] = t.(IRelationDefGroupingContext)
			i++
		}
	}

	return tst
}

func (s *RelationDefPartialAllAndContext) RelationDefGrouping(i int) IRelationDefGroupingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefGroupingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefGroupingContext)
}

func (s *RelationDefPartialAllAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefPartialAllAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefPartialAllAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefPartialAllAnd(s)
	}
}

func (s *RelationDefPartialAllAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefPartialAllAnd(s)
	}
}

func (p *OpenFGAParser) RelationDefPartialAllAnd() (localctx IRelationDefPartialAllAndContext) {
	localctx = NewRelationDefPartialAllAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, OpenFGAParserRULE_relationDefPartialAllAnd)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(175)
				p.Spacing()
			}
			{
				p.SetState(176)
				p.RelationDefOperatorAnd()
			}
			{
				p.SetState(177)
				p.Spacing()
			}
			{
				p.SetState(178)
				p.RelationDefGrouping()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefPartialAllButNotContext is an interface to support dynamic dispatch.
type IRelationDefPartialAllButNotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSpacing() []ISpacingContext
	Spacing(i int) ISpacingContext
	AllRelationDefOperatorButNot() []IRelationDefOperatorButNotContext
	RelationDefOperatorButNot(i int) IRelationDefOperatorButNotContext
	AllRelationDefGrouping() []IRelationDefGroupingContext
	RelationDefGrouping(i int) IRelationDefGroupingContext

	// IsRelationDefPartialAllButNotContext differentiates from other interfaces.
	IsRelationDefPartialAllButNotContext()
}

type RelationDefPartialAllButNotContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefPartialAllButNotContext() *RelationDefPartialAllButNotContext {
	var p = new(RelationDefPartialAllButNotContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefPartialAllButNot
	return p
}

func InitEmptyRelationDefPartialAllButNotContext(p *RelationDefPartialAllButNotContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefPartialAllButNot
}

func (*RelationDefPartialAllButNotContext) IsRelationDefPartialAllButNotContext() {}

func NewRelationDefPartialAllButNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefPartialAllButNotContext {
	var p = new(RelationDefPartialAllButNotContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefPartialAllButNot

	return p
}

func (s *RelationDefPartialAllButNotContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefPartialAllButNotContext) AllSpacing() []ISpacingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISpacingContext); ok {
			len++
		}
	}

	tst := make([]ISpacingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISpacingContext); ok {
			tst[i] = t.(ISpacingContext)
			i++
		}
	}

	return tst
}

func (s *RelationDefPartialAllButNotContext) Spacing(i int) ISpacingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpacingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpacingContext)
}

func (s *RelationDefPartialAllButNotContext) AllRelationDefOperatorButNot() []IRelationDefOperatorButNotContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationDefOperatorButNotContext); ok {
			len++
		}
	}

	tst := make([]IRelationDefOperatorButNotContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationDefOperatorButNotContext); ok {
			tst[i] = t.(IRelationDefOperatorButNotContext)
			i++
		}
	}

	return tst
}

func (s *RelationDefPartialAllButNotContext) RelationDefOperatorButNot(i int) IRelationDefOperatorButNotContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefOperatorButNotContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefOperatorButNotContext)
}

func (s *RelationDefPartialAllButNotContext) AllRelationDefGrouping() []IRelationDefGroupingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationDefGroupingContext); ok {
			len++
		}
	}

	tst := make([]IRelationDefGroupingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationDefGroupingContext); ok {
			tst[i] = t.(IRelationDefGroupingContext)
			i++
		}
	}

	return tst
}

func (s *RelationDefPartialAllButNotContext) RelationDefGrouping(i int) IRelationDefGroupingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefGroupingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefGroupingContext)
}

func (s *RelationDefPartialAllButNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefPartialAllButNotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefPartialAllButNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefPartialAllButNot(s)
	}
}

func (s *RelationDefPartialAllButNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefPartialAllButNot(s)
	}
}

func (p *OpenFGAParser) RelationDefPartialAllButNot() (localctx IRelationDefPartialAllButNotContext) {
	localctx = NewRelationDefPartialAllButNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, OpenFGAParserRULE_relationDefPartialAllButNot)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(184)
				p.Spacing()
			}
			{
				p.SetState(185)
				p.RelationDefOperatorButNot()
			}
			{
				p.SetState(186)
				p.Spacing()
			}
			{
				p.SetState(187)
				p.RelationDefGrouping()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefDirectAssignmentContext is an interface to support dynamic dispatch.
type IRelationDefDirectAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_SQUARE() antlr.TerminalNode
	AllRelationDefTypeRestriction() []IRelationDefTypeRestrictionContext
	RelationDefTypeRestriction(i int) IRelationDefTypeRestrictionContext
	R_SQUARE() antlr.TerminalNode
	AllSpacing() []ISpacingContext
	Spacing(i int) ISpacingContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsRelationDefDirectAssignmentContext differentiates from other interfaces.
	IsRelationDefDirectAssignmentContext()
}

type RelationDefDirectAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefDirectAssignmentContext() *RelationDefDirectAssignmentContext {
	var p = new(RelationDefDirectAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefDirectAssignment
	return p
}

func InitEmptyRelationDefDirectAssignmentContext(p *RelationDefDirectAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefDirectAssignment
}

func (*RelationDefDirectAssignmentContext) IsRelationDefDirectAssignmentContext() {}

func NewRelationDefDirectAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefDirectAssignmentContext {
	var p = new(RelationDefDirectAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefDirectAssignment

	return p
}

func (s *RelationDefDirectAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefDirectAssignmentContext) L_SQUARE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserL_SQUARE, 0)
}

func (s *RelationDefDirectAssignmentContext) AllRelationDefTypeRestriction() []IRelationDefTypeRestrictionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationDefTypeRestrictionContext); ok {
			len++
		}
	}

	tst := make([]IRelationDefTypeRestrictionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationDefTypeRestrictionContext); ok {
			tst[i] = t.(IRelationDefTypeRestrictionContext)
			i++
		}
	}

	return tst
}

func (s *RelationDefDirectAssignmentContext) RelationDefTypeRestriction(i int) IRelationDefTypeRestrictionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefTypeRestrictionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefTypeRestrictionContext)
}

func (s *RelationDefDirectAssignmentContext) R_SQUARE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserR_SQUARE, 0)
}

func (s *RelationDefDirectAssignmentContext) AllSpacing() []ISpacingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISpacingContext); ok {
			len++
		}
	}

	tst := make([]ISpacingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISpacingContext); ok {
			tst[i] = t.(ISpacingContext)
			i++
		}
	}

	return tst
}

func (s *RelationDefDirectAssignmentContext) Spacing(i int) ISpacingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpacingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpacingContext)
}

func (s *RelationDefDirectAssignmentContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserCOMMA)
}

func (s *RelationDefDirectAssignmentContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserCOMMA, i)
}

func (s *RelationDefDirectAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefDirectAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefDirectAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefDirectAssignment(s)
	}
}

func (s *RelationDefDirectAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefDirectAssignment(s)
	}
}

func (p *OpenFGAParser) RelationDefDirectAssignment() (localctx IRelationDefDirectAssignmentContext) {
	localctx = NewRelationDefDirectAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, OpenFGAParserRULE_relationDefDirectAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(OpenFGAParserL_SQUARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserWS {
		{
			p.SetState(194)
			p.Spacing()
		}

	}
	{
		p.SetState(197)
		p.RelationDefTypeRestriction()
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(198)
			p.Spacing()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == OpenFGAParserCOMMA {
		{
			p.SetState(201)
			p.Match(OpenFGAParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == OpenFGAParserWS {
			{
				p.SetState(202)
				p.Spacing()
			}

		}
		{
			p.SetState(205)
			p.RelationDefTypeRestriction()
		}

		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserWS {
		{
			p.SetState(211)
			p.Spacing()
		}

	}
	{
		p.SetState(214)
		p.Match(OpenFGAParserR_SQUARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefRewriteContext is an interface to support dynamic dispatch.
type IRelationDefRewriteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationDefRelationOnSameObject() IRelationDefRelationOnSameObjectContext
	RelationDefRelationOnRelatedObject() IRelationDefRelationOnRelatedObjectContext

	// IsRelationDefRewriteContext differentiates from other interfaces.
	IsRelationDefRewriteContext()
}

type RelationDefRewriteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefRewriteContext() *RelationDefRewriteContext {
	var p = new(RelationDefRewriteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefRewrite
	return p
}

func InitEmptyRelationDefRewriteContext(p *RelationDefRewriteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefRewrite
}

func (*RelationDefRewriteContext) IsRelationDefRewriteContext() {}

func NewRelationDefRewriteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefRewriteContext {
	var p = new(RelationDefRewriteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefRewrite

	return p
}

func (s *RelationDefRewriteContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefRewriteContext) RelationDefRelationOnSameObject() IRelationDefRelationOnSameObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefRelationOnSameObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefRelationOnSameObjectContext)
}

func (s *RelationDefRewriteContext) RelationDefRelationOnRelatedObject() IRelationDefRelationOnRelatedObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefRelationOnRelatedObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefRelationOnRelatedObjectContext)
}

func (s *RelationDefRewriteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefRewriteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefRewriteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefRewrite(s)
	}
}

func (s *RelationDefRewriteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefRewrite(s)
	}
}

func (p *OpenFGAParser) RelationDefRewrite() (localctx IRelationDefRewriteContext) {
	localctx = NewRelationDefRewriteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, OpenFGAParserRULE_relationDefRewrite)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(216)
			p.RelationDefRelationOnSameObject()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.RelationDefRelationOnRelatedObject()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefRelationOnSameObjectContext is an interface to support dynamic dispatch.
type IRelationDefRelationOnSameObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RewriteComputedusersetName() IRewriteComputedusersetNameContext

	// IsRelationDefRelationOnSameObjectContext differentiates from other interfaces.
	IsRelationDefRelationOnSameObjectContext()
}

type RelationDefRelationOnSameObjectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefRelationOnSameObjectContext() *RelationDefRelationOnSameObjectContext {
	var p = new(RelationDefRelationOnSameObjectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefRelationOnSameObject
	return p
}

func InitEmptyRelationDefRelationOnSameObjectContext(p *RelationDefRelationOnSameObjectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefRelationOnSameObject
}

func (*RelationDefRelationOnSameObjectContext) IsRelationDefRelationOnSameObjectContext() {}

func NewRelationDefRelationOnSameObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefRelationOnSameObjectContext {
	var p = new(RelationDefRelationOnSameObjectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefRelationOnSameObject

	return p
}

func (s *RelationDefRelationOnSameObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefRelationOnSameObjectContext) RewriteComputedusersetName() IRewriteComputedusersetNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRewriteComputedusersetNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRewriteComputedusersetNameContext)
}

func (s *RelationDefRelationOnSameObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefRelationOnSameObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefRelationOnSameObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefRelationOnSameObject(s)
	}
}

func (s *RelationDefRelationOnSameObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefRelationOnSameObject(s)
	}
}

func (p *OpenFGAParser) RelationDefRelationOnSameObject() (localctx IRelationDefRelationOnSameObjectContext) {
	localctx = NewRelationDefRelationOnSameObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, OpenFGAParserRULE_relationDefRelationOnSameObject)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.RewriteComputedusersetName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefRelationOnRelatedObjectContext is an interface to support dynamic dispatch.
type IRelationDefRelationOnRelatedObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RewriteTuplesetComputedusersetName() IRewriteTuplesetComputedusersetNameContext
	AllSpacing() []ISpacingContext
	Spacing(i int) ISpacingContext
	RelationDefKeywordFrom() IRelationDefKeywordFromContext
	RewriteTuplesetName() IRewriteTuplesetNameContext

	// IsRelationDefRelationOnRelatedObjectContext differentiates from other interfaces.
	IsRelationDefRelationOnRelatedObjectContext()
}

type RelationDefRelationOnRelatedObjectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefRelationOnRelatedObjectContext() *RelationDefRelationOnRelatedObjectContext {
	var p = new(RelationDefRelationOnRelatedObjectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefRelationOnRelatedObject
	return p
}

func InitEmptyRelationDefRelationOnRelatedObjectContext(p *RelationDefRelationOnRelatedObjectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefRelationOnRelatedObject
}

func (*RelationDefRelationOnRelatedObjectContext) IsRelationDefRelationOnRelatedObjectContext() {}

func NewRelationDefRelationOnRelatedObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefRelationOnRelatedObjectContext {
	var p = new(RelationDefRelationOnRelatedObjectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefRelationOnRelatedObject

	return p
}

func (s *RelationDefRelationOnRelatedObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefRelationOnRelatedObjectContext) RewriteTuplesetComputedusersetName() IRewriteTuplesetComputedusersetNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRewriteTuplesetComputedusersetNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRewriteTuplesetComputedusersetNameContext)
}

func (s *RelationDefRelationOnRelatedObjectContext) AllSpacing() []ISpacingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISpacingContext); ok {
			len++
		}
	}

	tst := make([]ISpacingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISpacingContext); ok {
			tst[i] = t.(ISpacingContext)
			i++
		}
	}

	return tst
}

func (s *RelationDefRelationOnRelatedObjectContext) Spacing(i int) ISpacingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpacingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpacingContext)
}

func (s *RelationDefRelationOnRelatedObjectContext) RelationDefKeywordFrom() IRelationDefKeywordFromContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefKeywordFromContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefKeywordFromContext)
}

func (s *RelationDefRelationOnRelatedObjectContext) RewriteTuplesetName() IRewriteTuplesetNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRewriteTuplesetNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRewriteTuplesetNameContext)
}

func (s *RelationDefRelationOnRelatedObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefRelationOnRelatedObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefRelationOnRelatedObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefRelationOnRelatedObject(s)
	}
}

func (s *RelationDefRelationOnRelatedObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefRelationOnRelatedObject(s)
	}
}

func (p *OpenFGAParser) RelationDefRelationOnRelatedObject() (localctx IRelationDefRelationOnRelatedObjectContext) {
	localctx = NewRelationDefRelationOnRelatedObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, OpenFGAParserRULE_relationDefRelationOnRelatedObject)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.RewriteTuplesetComputedusersetName()
	}
	{
		p.SetState(223)
		p.Spacing()
	}
	{
		p.SetState(224)
		p.RelationDefKeywordFrom()
	}
	{
		p.SetState(225)
		p.Spacing()
	}
	{
		p.SetState(226)
		p.RewriteTuplesetName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefOperatorContext is an interface to support dynamic dispatch.
type IRelationDefOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationDefOperatorOr() IRelationDefOperatorOrContext
	RelationDefOperatorAnd() IRelationDefOperatorAndContext
	RelationDefOperatorButNot() IRelationDefOperatorButNotContext

	// IsRelationDefOperatorContext differentiates from other interfaces.
	IsRelationDefOperatorContext()
}

type RelationDefOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefOperatorContext() *RelationDefOperatorContext {
	var p = new(RelationDefOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefOperator
	return p
}

func InitEmptyRelationDefOperatorContext(p *RelationDefOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefOperator
}

func (*RelationDefOperatorContext) IsRelationDefOperatorContext() {}

func NewRelationDefOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefOperatorContext {
	var p = new(RelationDefOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefOperator

	return p
}

func (s *RelationDefOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefOperatorContext) RelationDefOperatorOr() IRelationDefOperatorOrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefOperatorOrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefOperatorOrContext)
}

func (s *RelationDefOperatorContext) RelationDefOperatorAnd() IRelationDefOperatorAndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefOperatorAndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefOperatorAndContext)
}

func (s *RelationDefOperatorContext) RelationDefOperatorButNot() IRelationDefOperatorButNotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefOperatorButNotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefOperatorButNotContext)
}

func (s *RelationDefOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefOperator(s)
	}
}

func (s *RelationDefOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefOperator(s)
	}
}

func (p *OpenFGAParser) RelationDefOperator() (localctx IRelationDefOperatorContext) {
	localctx = NewRelationDefOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, OpenFGAParserRULE_relationDefOperator)
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case OpenFGAParserOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.RelationDefOperatorOr()
		}

	case OpenFGAParserAND:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.RelationDefOperatorAnd()
		}

	case OpenFGAParserBUT_NOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(230)
			p.RelationDefOperatorButNot()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefOperatorAndContext is an interface to support dynamic dispatch.
type IRelationDefOperatorAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND() antlr.TerminalNode

	// IsRelationDefOperatorAndContext differentiates from other interfaces.
	IsRelationDefOperatorAndContext()
}

type RelationDefOperatorAndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefOperatorAndContext() *RelationDefOperatorAndContext {
	var p = new(RelationDefOperatorAndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefOperatorAnd
	return p
}

func InitEmptyRelationDefOperatorAndContext(p *RelationDefOperatorAndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefOperatorAnd
}

func (*RelationDefOperatorAndContext) IsRelationDefOperatorAndContext() {}

func NewRelationDefOperatorAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefOperatorAndContext {
	var p = new(RelationDefOperatorAndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefOperatorAnd

	return p
}

func (s *RelationDefOperatorAndContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefOperatorAndContext) AND() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserAND, 0)
}

func (s *RelationDefOperatorAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefOperatorAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefOperatorAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefOperatorAnd(s)
	}
}

func (s *RelationDefOperatorAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefOperatorAnd(s)
	}
}

func (p *OpenFGAParser) RelationDefOperatorAnd() (localctx IRelationDefOperatorAndContext) {
	localctx = NewRelationDefOperatorAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, OpenFGAParserRULE_relationDefOperatorAnd)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.Match(OpenFGAParserAND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefOperatorOrContext is an interface to support dynamic dispatch.
type IRelationDefOperatorOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OR() antlr.TerminalNode

	// IsRelationDefOperatorOrContext differentiates from other interfaces.
	IsRelationDefOperatorOrContext()
}

type RelationDefOperatorOrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefOperatorOrContext() *RelationDefOperatorOrContext {
	var p = new(RelationDefOperatorOrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefOperatorOr
	return p
}

func InitEmptyRelationDefOperatorOrContext(p *RelationDefOperatorOrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefOperatorOr
}

func (*RelationDefOperatorOrContext) IsRelationDefOperatorOrContext() {}

func NewRelationDefOperatorOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefOperatorOrContext {
	var p = new(RelationDefOperatorOrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefOperatorOr

	return p
}

func (s *RelationDefOperatorOrContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefOperatorOrContext) OR() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserOR, 0)
}

func (s *RelationDefOperatorOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefOperatorOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefOperatorOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefOperatorOr(s)
	}
}

func (s *RelationDefOperatorOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefOperatorOr(s)
	}
}

func (p *OpenFGAParser) RelationDefOperatorOr() (localctx IRelationDefOperatorOrContext) {
	localctx = NewRelationDefOperatorOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, OpenFGAParserRULE_relationDefOperatorOr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(OpenFGAParserOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefOperatorButNotContext is an interface to support dynamic dispatch.
type IRelationDefOperatorButNotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BUT_NOT() antlr.TerminalNode

	// IsRelationDefOperatorButNotContext differentiates from other interfaces.
	IsRelationDefOperatorButNotContext()
}

type RelationDefOperatorButNotContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefOperatorButNotContext() *RelationDefOperatorButNotContext {
	var p = new(RelationDefOperatorButNotContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefOperatorButNot
	return p
}

func InitEmptyRelationDefOperatorButNotContext(p *RelationDefOperatorButNotContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefOperatorButNot
}

func (*RelationDefOperatorButNotContext) IsRelationDefOperatorButNotContext() {}

func NewRelationDefOperatorButNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefOperatorButNotContext {
	var p = new(RelationDefOperatorButNotContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefOperatorButNot

	return p
}

func (s *RelationDefOperatorButNotContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefOperatorButNotContext) BUT_NOT() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserBUT_NOT, 0)
}

func (s *RelationDefOperatorButNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefOperatorButNotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefOperatorButNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefOperatorButNot(s)
	}
}

func (s *RelationDefOperatorButNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefOperatorButNot(s)
	}
}

func (p *OpenFGAParser) RelationDefOperatorButNot() (localctx IRelationDefOperatorButNotContext) {
	localctx = NewRelationDefOperatorButNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, OpenFGAParserRULE_relationDefOperatorButNot)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(OpenFGAParserBUT_NOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefKeywordFromContext is an interface to support dynamic dispatch.
type IRelationDefKeywordFromContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FROM() antlr.TerminalNode

	// IsRelationDefKeywordFromContext differentiates from other interfaces.
	IsRelationDefKeywordFromContext()
}

type RelationDefKeywordFromContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefKeywordFromContext() *RelationDefKeywordFromContext {
	var p = new(RelationDefKeywordFromContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefKeywordFrom
	return p
}

func InitEmptyRelationDefKeywordFromContext(p *RelationDefKeywordFromContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefKeywordFrom
}

func (*RelationDefKeywordFromContext) IsRelationDefKeywordFromContext() {}

func NewRelationDefKeywordFromContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefKeywordFromContext {
	var p = new(RelationDefKeywordFromContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefKeywordFrom

	return p
}

func (s *RelationDefKeywordFromContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefKeywordFromContext) FROM() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserFROM, 0)
}

func (s *RelationDefKeywordFromContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefKeywordFromContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefKeywordFromContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefKeywordFrom(s)
	}
}

func (s *RelationDefKeywordFromContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefKeywordFrom(s)
	}
}

func (p *OpenFGAParser) RelationDefKeywordFrom() (localctx IRelationDefKeywordFromContext) {
	localctx = NewRelationDefKeywordFromContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, OpenFGAParserRULE_relationDefKeywordFrom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(OpenFGAParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefTypeRestrictionContext is an interface to support dynamic dispatch.
type IRelationDefTypeRestrictionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationDefTypeRestrictionType() IRelationDefTypeRestrictionTypeContext
	RelationDefTypeRestrictionWildcard() IRelationDefTypeRestrictionWildcardContext
	RelationDefTypeRestrictionUserset() IRelationDefTypeRestrictionUsersetContext

	// IsRelationDefTypeRestrictionContext differentiates from other interfaces.
	IsRelationDefTypeRestrictionContext()
}

type RelationDefTypeRestrictionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefTypeRestrictionContext() *RelationDefTypeRestrictionContext {
	var p = new(RelationDefTypeRestrictionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestriction
	return p
}

func InitEmptyRelationDefTypeRestrictionContext(p *RelationDefTypeRestrictionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestriction
}

func (*RelationDefTypeRestrictionContext) IsRelationDefTypeRestrictionContext() {}

func NewRelationDefTypeRestrictionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefTypeRestrictionContext {
	var p = new(RelationDefTypeRestrictionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestriction

	return p
}

func (s *RelationDefTypeRestrictionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefTypeRestrictionContext) RelationDefTypeRestrictionType() IRelationDefTypeRestrictionTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefTypeRestrictionTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefTypeRestrictionTypeContext)
}

func (s *RelationDefTypeRestrictionContext) RelationDefTypeRestrictionWildcard() IRelationDefTypeRestrictionWildcardContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefTypeRestrictionWildcardContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefTypeRestrictionWildcardContext)
}

func (s *RelationDefTypeRestrictionContext) RelationDefTypeRestrictionUserset() IRelationDefTypeRestrictionUsersetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefTypeRestrictionUsersetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefTypeRestrictionUsersetContext)
}

func (s *RelationDefTypeRestrictionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefTypeRestrictionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefTypeRestrictionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefTypeRestriction(s)
	}
}

func (s *RelationDefTypeRestrictionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefTypeRestriction(s)
	}
}

func (p *OpenFGAParser) RelationDefTypeRestriction() (localctx IRelationDefTypeRestrictionContext) {
	localctx = NewRelationDefTypeRestrictionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, OpenFGAParserRULE_relationDefTypeRestriction)
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.RelationDefTypeRestrictionType()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.RelationDefTypeRestrictionWildcard()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(243)
			p.RelationDefTypeRestrictionUserset()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefTypeRestrictionTypeContext is an interface to support dynamic dispatch.
type IRelationDefTypeRestrictionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsRelationDefTypeRestrictionTypeContext differentiates from other interfaces.
	IsRelationDefTypeRestrictionTypeContext()
}

type RelationDefTypeRestrictionTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefTypeRestrictionTypeContext() *RelationDefTypeRestrictionTypeContext {
	var p = new(RelationDefTypeRestrictionTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestrictionType
	return p
}

func InitEmptyRelationDefTypeRestrictionTypeContext(p *RelationDefTypeRestrictionTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestrictionType
}

func (*RelationDefTypeRestrictionTypeContext) IsRelationDefTypeRestrictionTypeContext() {}

func NewRelationDefTypeRestrictionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefTypeRestrictionTypeContext {
	var p = new(RelationDefTypeRestrictionTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestrictionType

	return p
}

func (s *RelationDefTypeRestrictionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefTypeRestrictionTypeContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RelationDefTypeRestrictionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefTypeRestrictionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefTypeRestrictionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefTypeRestrictionType(s)
	}
}

func (s *RelationDefTypeRestrictionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefTypeRestrictionType(s)
	}
}

func (p *OpenFGAParser) RelationDefTypeRestrictionType() (localctx IRelationDefTypeRestrictionTypeContext) {
	localctx = NewRelationDefTypeRestrictionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, OpenFGAParserRULE_relationDefTypeRestrictionType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefTypeRestrictionRelationContext is an interface to support dynamic dispatch.
type IRelationDefTypeRestrictionRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsRelationDefTypeRestrictionRelationContext differentiates from other interfaces.
	IsRelationDefTypeRestrictionRelationContext()
}

type RelationDefTypeRestrictionRelationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefTypeRestrictionRelationContext() *RelationDefTypeRestrictionRelationContext {
	var p = new(RelationDefTypeRestrictionRelationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestrictionRelation
	return p
}

func InitEmptyRelationDefTypeRestrictionRelationContext(p *RelationDefTypeRestrictionRelationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestrictionRelation
}

func (*RelationDefTypeRestrictionRelationContext) IsRelationDefTypeRestrictionRelationContext() {}

func NewRelationDefTypeRestrictionRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefTypeRestrictionRelationContext {
	var p = new(RelationDefTypeRestrictionRelationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestrictionRelation

	return p
}

func (s *RelationDefTypeRestrictionRelationContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefTypeRestrictionRelationContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RelationDefTypeRestrictionRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefTypeRestrictionRelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefTypeRestrictionRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefTypeRestrictionRelation(s)
	}
}

func (s *RelationDefTypeRestrictionRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefTypeRestrictionRelation(s)
	}
}

func (p *OpenFGAParser) RelationDefTypeRestrictionRelation() (localctx IRelationDefTypeRestrictionRelationContext) {
	localctx = NewRelationDefTypeRestrictionRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, OpenFGAParserRULE_relationDefTypeRestrictionRelation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefTypeRestrictionWildcardContext is an interface to support dynamic dispatch.
type IRelationDefTypeRestrictionWildcardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationDefTypeRestrictionType() IRelationDefTypeRestrictionTypeContext
	COLON() antlr.TerminalNode
	WILDCARD() antlr.TerminalNode
	Spacing() ISpacingContext

	// IsRelationDefTypeRestrictionWildcardContext differentiates from other interfaces.
	IsRelationDefTypeRestrictionWildcardContext()
}

type RelationDefTypeRestrictionWildcardContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefTypeRestrictionWildcardContext() *RelationDefTypeRestrictionWildcardContext {
	var p = new(RelationDefTypeRestrictionWildcardContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestrictionWildcard
	return p
}

func InitEmptyRelationDefTypeRestrictionWildcardContext(p *RelationDefTypeRestrictionWildcardContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestrictionWildcard
}

func (*RelationDefTypeRestrictionWildcardContext) IsRelationDefTypeRestrictionWildcardContext() {}

func NewRelationDefTypeRestrictionWildcardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefTypeRestrictionWildcardContext {
	var p = new(RelationDefTypeRestrictionWildcardContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestrictionWildcard

	return p
}

func (s *RelationDefTypeRestrictionWildcardContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefTypeRestrictionWildcardContext) RelationDefTypeRestrictionType() IRelationDefTypeRestrictionTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefTypeRestrictionTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefTypeRestrictionTypeContext)
}

func (s *RelationDefTypeRestrictionWildcardContext) COLON() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserCOLON, 0)
}

func (s *RelationDefTypeRestrictionWildcardContext) WILDCARD() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWILDCARD, 0)
}

func (s *RelationDefTypeRestrictionWildcardContext) Spacing() ISpacingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpacingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpacingContext)
}

func (s *RelationDefTypeRestrictionWildcardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefTypeRestrictionWildcardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefTypeRestrictionWildcardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefTypeRestrictionWildcard(s)
	}
}

func (s *RelationDefTypeRestrictionWildcardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefTypeRestrictionWildcard(s)
	}
}

func (p *OpenFGAParser) RelationDefTypeRestrictionWildcard() (localctx IRelationDefTypeRestrictionWildcardContext) {
	localctx = NewRelationDefTypeRestrictionWildcardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, OpenFGAParserRULE_relationDefTypeRestrictionWildcard)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.RelationDefTypeRestrictionType()
	}
	{
		p.SetState(251)
		p.Match(OpenFGAParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(252)
		p.Match(OpenFGAParserWILDCARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(253)
			p.Spacing()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefTypeRestrictionUsersetContext is an interface to support dynamic dispatch.
type IRelationDefTypeRestrictionUsersetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationDefTypeRestrictionType() IRelationDefTypeRestrictionTypeContext
	HASH() antlr.TerminalNode
	RelationDefTypeRestrictionRelation() IRelationDefTypeRestrictionRelationContext

	// IsRelationDefTypeRestrictionUsersetContext differentiates from other interfaces.
	IsRelationDefTypeRestrictionUsersetContext()
}

type RelationDefTypeRestrictionUsersetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefTypeRestrictionUsersetContext() *RelationDefTypeRestrictionUsersetContext {
	var p = new(RelationDefTypeRestrictionUsersetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestrictionUserset
	return p
}

func InitEmptyRelationDefTypeRestrictionUsersetContext(p *RelationDefTypeRestrictionUsersetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestrictionUserset
}

func (*RelationDefTypeRestrictionUsersetContext) IsRelationDefTypeRestrictionUsersetContext() {}

func NewRelationDefTypeRestrictionUsersetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefTypeRestrictionUsersetContext {
	var p = new(RelationDefTypeRestrictionUsersetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestrictionUserset

	return p
}

func (s *RelationDefTypeRestrictionUsersetContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefTypeRestrictionUsersetContext) RelationDefTypeRestrictionType() IRelationDefTypeRestrictionTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefTypeRestrictionTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefTypeRestrictionTypeContext)
}

func (s *RelationDefTypeRestrictionUsersetContext) HASH() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserHASH, 0)
}

func (s *RelationDefTypeRestrictionUsersetContext) RelationDefTypeRestrictionRelation() IRelationDefTypeRestrictionRelationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefTypeRestrictionRelationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefTypeRestrictionRelationContext)
}

func (s *RelationDefTypeRestrictionUsersetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefTypeRestrictionUsersetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefTypeRestrictionUsersetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefTypeRestrictionUserset(s)
	}
}

func (s *RelationDefTypeRestrictionUsersetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefTypeRestrictionUserset(s)
	}
}

func (p *OpenFGAParser) RelationDefTypeRestrictionUserset() (localctx IRelationDefTypeRestrictionUsersetContext) {
	localctx = NewRelationDefTypeRestrictionUsersetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, OpenFGAParserRULE_relationDefTypeRestrictionUserset)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.RelationDefTypeRestrictionType()
	}
	{
		p.SetState(257)
		p.Match(OpenFGAParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.RelationDefTypeRestrictionRelation()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationDefGroupingContext is an interface to support dynamic dispatch.
type IRelationDefGroupingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationDefRewrite() IRelationDefRewriteContext

	// IsRelationDefGroupingContext differentiates from other interfaces.
	IsRelationDefGroupingContext()
}

type RelationDefGroupingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefGroupingContext() *RelationDefGroupingContext {
	var p = new(RelationDefGroupingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefGrouping
	return p
}

func InitEmptyRelationDefGroupingContext(p *RelationDefGroupingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefGrouping
}

func (*RelationDefGroupingContext) IsRelationDefGroupingContext() {}

func NewRelationDefGroupingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefGroupingContext {
	var p = new(RelationDefGroupingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefGrouping

	return p
}

func (s *RelationDefGroupingContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefGroupingContext) RelationDefRewrite() IRelationDefRewriteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefRewriteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefRewriteContext)
}

func (s *RelationDefGroupingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefGroupingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationDefGroupingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefGrouping(s)
	}
}

func (s *RelationDefGroupingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefGrouping(s)
	}
}

func (p *OpenFGAParser) RelationDefGrouping() (localctx IRelationDefGroupingContext) {
	localctx = NewRelationDefGroupingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, OpenFGAParserRULE_relationDefGrouping)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.RelationDefRewrite()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRewriteComputedusersetNameContext is an interface to support dynamic dispatch.
type IRewriteComputedusersetNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsRewriteComputedusersetNameContext differentiates from other interfaces.
	IsRewriteComputedusersetNameContext()
}

type RewriteComputedusersetNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRewriteComputedusersetNameContext() *RewriteComputedusersetNameContext {
	var p = new(RewriteComputedusersetNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_rewriteComputedusersetName
	return p
}

func InitEmptyRewriteComputedusersetNameContext(p *RewriteComputedusersetNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_rewriteComputedusersetName
}

func (*RewriteComputedusersetNameContext) IsRewriteComputedusersetNameContext() {}

func NewRewriteComputedusersetNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RewriteComputedusersetNameContext {
	var p = new(RewriteComputedusersetNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_rewriteComputedusersetName

	return p
}

func (s *RewriteComputedusersetNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RewriteComputedusersetNameContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RewriteComputedusersetNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RewriteComputedusersetNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RewriteComputedusersetNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRewriteComputedusersetName(s)
	}
}

func (s *RewriteComputedusersetNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRewriteComputedusersetName(s)
	}
}

func (p *OpenFGAParser) RewriteComputedusersetName() (localctx IRewriteComputedusersetNameContext) {
	localctx = NewRewriteComputedusersetNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, OpenFGAParserRULE_rewriteComputedusersetName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.Name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRewriteTuplesetComputedusersetNameContext is an interface to support dynamic dispatch.
type IRewriteTuplesetComputedusersetNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsRewriteTuplesetComputedusersetNameContext differentiates from other interfaces.
	IsRewriteTuplesetComputedusersetNameContext()
}

type RewriteTuplesetComputedusersetNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRewriteTuplesetComputedusersetNameContext() *RewriteTuplesetComputedusersetNameContext {
	var p = new(RewriteTuplesetComputedusersetNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_rewriteTuplesetComputedusersetName
	return p
}

func InitEmptyRewriteTuplesetComputedusersetNameContext(p *RewriteTuplesetComputedusersetNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_rewriteTuplesetComputedusersetName
}

func (*RewriteTuplesetComputedusersetNameContext) IsRewriteTuplesetComputedusersetNameContext() {}

func NewRewriteTuplesetComputedusersetNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RewriteTuplesetComputedusersetNameContext {
	var p = new(RewriteTuplesetComputedusersetNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_rewriteTuplesetComputedusersetName

	return p
}

func (s *RewriteTuplesetComputedusersetNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RewriteTuplesetComputedusersetNameContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RewriteTuplesetComputedusersetNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RewriteTuplesetComputedusersetNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RewriteTuplesetComputedusersetNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRewriteTuplesetComputedusersetName(s)
	}
}

func (s *RewriteTuplesetComputedusersetNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRewriteTuplesetComputedusersetName(s)
	}
}

func (p *OpenFGAParser) RewriteTuplesetComputedusersetName() (localctx IRewriteTuplesetComputedusersetNameContext) {
	localctx = NewRewriteTuplesetComputedusersetNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, OpenFGAParserRULE_rewriteTuplesetComputedusersetName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRewriteTuplesetNameContext is an interface to support dynamic dispatch.
type IRewriteTuplesetNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsRewriteTuplesetNameContext differentiates from other interfaces.
	IsRewriteTuplesetNameContext()
}

type RewriteTuplesetNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRewriteTuplesetNameContext() *RewriteTuplesetNameContext {
	var p = new(RewriteTuplesetNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_rewriteTuplesetName
	return p
}

func InitEmptyRewriteTuplesetNameContext(p *RewriteTuplesetNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_rewriteTuplesetName
}

func (*RewriteTuplesetNameContext) IsRewriteTuplesetNameContext() {}

func NewRewriteTuplesetNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RewriteTuplesetNameContext {
	var p = new(RewriteTuplesetNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_rewriteTuplesetName

	return p
}

func (s *RewriteTuplesetNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RewriteTuplesetNameContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RewriteTuplesetNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RewriteTuplesetNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RewriteTuplesetNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRewriteTuplesetName(s)
	}
}

func (s *RewriteTuplesetNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRewriteTuplesetName(s)
	}
}

func (p *OpenFGAParser) RewriteTuplesetName() (localctx IRewriteTuplesetNameContext) {
	localctx = NewRewriteTuplesetNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, OpenFGAParserRULE_rewriteTuplesetName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationNameContext is an interface to support dynamic dispatch.
type IRelationNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsRelationNameContext differentiates from other interfaces.
	IsRelationNameContext()
}

type RelationNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationNameContext() *RelationNameContext {
	var p = new(RelationNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationName
	return p
}

func InitEmptyRelationNameContext(p *RelationNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationName
}

func (*RelationNameContext) IsRelationNameContext() {}

func NewRelationNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationNameContext {
	var p = new(RelationNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationName

	return p
}

func (s *RelationNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationNameContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RelationNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationName(s)
	}
}

func (s *RelationNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationName(s)
	}
}

func (p *OpenFGAParser) RelationName() (localctx IRelationNameContext) {
	localctx = NewRelationNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, OpenFGAParserRULE_relationName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_typeName
	return p
}

func InitEmptyTypeNameContext(p *TypeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_typeName
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (p *OpenFGAParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, OpenFGAParserRULE_typeName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HASH() antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_comment
	return p
}

func InitEmptyCommentContext(p *CommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_comment
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) HASH() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserHASH, 0)
}

func (s *CommentContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserWS)
}

func (s *CommentContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWS, i)
}

func (s *CommentContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNEWLINE)
}

func (s *CommentContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINE, i)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *OpenFGAParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, OpenFGAParserRULE_comment)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == OpenFGAParserWS {
		{
			p.SetState(272)
			p.Match(OpenFGAParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(278)
		p.Match(OpenFGAParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(279)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || _la == OpenFGAParserNEWLINE {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiLineCommentContext is an interface to support dynamic dispatch.
type IMultiLineCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllComment() []ICommentContext
	Comment(i int) ICommentContext
	AllNewline() []INewlineContext
	Newline(i int) INewlineContext

	// IsMultiLineCommentContext differentiates from other interfaces.
	IsMultiLineCommentContext()
}

type MultiLineCommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiLineCommentContext() *MultiLineCommentContext {
	var p = new(MultiLineCommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_multiLineComment
	return p
}

func InitEmptyMultiLineCommentContext(p *MultiLineCommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_multiLineComment
}

func (*MultiLineCommentContext) IsMultiLineCommentContext() {}

func NewMultiLineCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiLineCommentContext {
	var p = new(MultiLineCommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_multiLineComment

	return p
}

func (s *MultiLineCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiLineCommentContext) AllComment() []ICommentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommentContext); ok {
			len++
		}
	}

	tst := make([]ICommentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommentContext); ok {
			tst[i] = t.(ICommentContext)
			i++
		}
	}

	return tst
}

func (s *MultiLineCommentContext) Comment(i int) ICommentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *MultiLineCommentContext) AllNewline() []INewlineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INewlineContext); ok {
			len++
		}
	}

	tst := make([]INewlineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INewlineContext); ok {
			tst[i] = t.(INewlineContext)
			i++
		}
	}

	return tst
}

func (s *MultiLineCommentContext) Newline(i int) INewlineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INewlineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INewlineContext)
}

func (s *MultiLineCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiLineCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiLineCommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterMultiLineComment(s)
	}
}

func (s *MultiLineCommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitMultiLineComment(s)
	}
}

func (p *OpenFGAParser) MultiLineComment() (localctx IMultiLineCommentContext) {
	localctx = NewMultiLineCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, OpenFGAParserRULE_multiLineComment)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Comment()
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(286)
				p.Newline()
			}
			{
				p.SetState(287)
				p.Comment()
			}

		}
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISpacingContext is an interface to support dynamic dispatch.
type ISpacingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode

	// IsSpacingContext differentiates from other interfaces.
	IsSpacingContext()
}

type SpacingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpacingContext() *SpacingContext {
	var p = new(SpacingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_spacing
	return p
}

func InitEmptySpacingContext(p *SpacingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_spacing
}

func (*SpacingContext) IsSpacingContext() {}

func NewSpacingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpacingContext {
	var p = new(SpacingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_spacing

	return p
}

func (s *SpacingContext) GetParser() antlr.Parser { return s.parser }

func (s *SpacingContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserWS)
}

func (s *SpacingContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWS, i)
}

func (s *SpacingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpacingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpacingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterSpacing(s)
	}
}

func (s *SpacingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitSpacing(s)
	}
}

func (p *OpenFGAParser) Spacing() (localctx ISpacingContext) {
	localctx = NewSpacingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, OpenFGAParserRULE_spacing)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(294)
				p.Match(OpenFGAParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INewlineContext is an interface to support dynamic dispatch.
type INewlineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsNewlineContext differentiates from other interfaces.
	IsNewlineContext()
}

type NewlineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNewlineContext() *NewlineContext {
	var p = new(NewlineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_newline
	return p
}

func InitEmptyNewlineContext(p *NewlineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_newline
}

func (*NewlineContext) IsNewlineContext() {}

func NewNewlineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NewlineContext {
	var p = new(NewlineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_newline

	return p
}

func (s *NewlineContext) GetParser() antlr.Parser { return s.parser }

func (s *NewlineContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNEWLINE)
}

func (s *NewlineContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINE, i)
}

func (s *NewlineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewlineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NewlineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterNewline(s)
	}
}

func (s *NewlineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitNewline(s)
	}
}

func (p *OpenFGAParser) Newline() (localctx INewlineContext) {
	localctx = NewNewlineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, OpenFGAParserRULE_newline)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == OpenFGAParserNEWLINE {
		{
			p.SetState(299)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISchemaVersionContext is an interface to support dynamic dispatch.
type ISchemaVersionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SCHEMA_VERSION() antlr.TerminalNode

	// IsSchemaVersionContext differentiates from other interfaces.
	IsSchemaVersionContext()
}

type SchemaVersionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaVersionContext() *SchemaVersionContext {
	var p = new(SchemaVersionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_schemaVersion
	return p
}

func InitEmptySchemaVersionContext(p *SchemaVersionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_schemaVersion
}

func (*SchemaVersionContext) IsSchemaVersionContext() {}

func NewSchemaVersionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaVersionContext {
	var p = new(SchemaVersionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_schemaVersion

	return p
}

func (s *SchemaVersionContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaVersionContext) SCHEMA_VERSION() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserSCHEMA_VERSION, 0)
}

func (s *SchemaVersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaVersionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaVersionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterSchemaVersion(s)
	}
}

func (s *SchemaVersionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitSchemaVersion(s)
	}
}

func (p *OpenFGAParser) SchemaVersion() (localctx ISchemaVersionContext) {
	localctx = NewSchemaVersionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, OpenFGAParserRULE_schemaVersion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Match(OpenFGAParserSCHEMA_VERSION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllALPHA_NUMERIC() []antlr.TerminalNode
	ALPHA_NUMERIC(i int) antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) AllALPHA_NUMERIC() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserALPHA_NUMERIC)
}

func (s *NameContext) ALPHA_NUMERIC(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserALPHA_NUMERIC, i)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *OpenFGAParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, OpenFGAParserRULE_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == OpenFGAParserALPHA_NUMERIC {
		{
			p.SetState(306)
			p.Match(OpenFGAParserALPHA_NUMERIC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
