// Code generated from /app/OpenFGAParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
    "", "':'", "','", "'<'", "'>'", "'['", "", "'('", "')'", "", "", "'#'", 
    "'and'", "'or'", "'but not'", "'from'", "'module'", "'model'", "'schema'", 
    "", "'extend'", "'type'", "'condition'", "'relations'", "'define'", 
    "'with'", "'=='", "'!='", "'in'", "'<='", "'>='", "'&&'", "'||'", "']'", 
    "'{'", "'}'", "'.'", "'-'", "'!'", "'?'", "'+'", "'*'", "'/'", "'%'", 
    "'true'", "'false'", "'null'",
  }
  staticData.SymbolicNames = []string{
    "", "COLON", "COMMA", "LESS", "GREATER", "LBRACKET", "RBRACKET", "LPAREN", 
    "RPAREN", "WHITESPACE", "IDENTIFIER", "HASH", "AND", "OR", "BUT_NOT", 
    "FROM", "MODULE", "MODEL", "SCHEMA", "SCHEMA_VERSION", "EXTEND", "TYPE", 
    "CONDITION", "RELATIONS", "DEFINE", "KEYWORD_WITH", "EQUALS", "NOT_EQUALS", 
    "IN", "LESS_EQUALS", "GREATER_EQUALS", "LOGICAL_AND", "LOGICAL_OR", 
    "RPRACKET", "LBRACE", "RBRACE", "DOT", "MINUS", "EXCLAM", "QUESTIONMARK", 
    "PLUS", "STAR", "SLASH", "PERCENT", "CEL_TRUE", "CEL_FALSE", "NUL", 
    "CEL_COMMENT", "NUM_FLOAT", "NUM_INT", "NUM_UINT", "STRING", "BYTES", 
    "NEWLINE", "CONDITION_PARAM_CONTAINER", "CONDITION_PARAM_TYPE",
  }
  staticData.RuleNames = []string{
    "main", "modelHeader", "moduleHeader", "typeDefs", "typeDef", "relationDeclaration", 
    "relationName", "relationDef", "relationDefNoDirect", "relationDefPartials", 
    "relationDefGrouping", "relationRecurse", "relationRecurseNoDirect", 
    "relationDefDirectAssignment", "relationDefRewrite", "relationDefTypeRestriction", 
    "relationDefTypeRestrictionBase", "conditions", "condition", "conditionName", 
    "conditionParameter", "parameterName", "parameterType", "multiLineComment", 
    "conditionExpression",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 55, 382, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2, 
	21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 3, 0, 52, 8, 
	0, 1, 0, 3, 0, 55, 8, 0, 1, 0, 1, 0, 3, 0, 59, 8, 0, 1, 0, 3, 0, 62, 8, 
	0, 1, 0, 1, 0, 3, 0, 66, 8, 0, 1, 0, 1, 0, 3, 0, 70, 8, 0, 1, 0, 1, 0, 
	1, 1, 1, 1, 1, 1, 3, 1, 77, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 
	1, 85, 8, 1, 1, 2, 1, 2, 1, 2, 3, 2, 90, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 
	3, 2, 96, 8, 2, 1, 3, 5, 3, 99, 8, 3, 10, 3, 12, 3, 102, 9, 3, 1, 4, 1, 
	4, 3, 4, 106, 8, 4, 1, 4, 1, 4, 1, 4, 3, 4, 111, 8, 4, 1, 4, 1, 4, 1, 4, 
	1, 4, 1, 4, 1, 4, 4, 4, 119, 8, 4, 11, 4, 12, 4, 120, 3, 4, 123, 8, 4, 
	1, 5, 1, 5, 3, 5, 127, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 134, 8, 
	5, 1, 5, 1, 5, 3, 5, 138, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 
	7, 3, 7, 147, 8, 7, 1, 7, 3, 7, 150, 8, 7, 1, 8, 1, 8, 3, 8, 154, 8, 8, 
	1, 8, 3, 8, 157, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 164, 8, 9, 4, 
	9, 166, 8, 9, 11, 9, 12, 9, 167, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 175, 
	8, 9, 4, 9, 177, 8, 9, 11, 9, 12, 9, 178, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 
	3, 9, 186, 8, 9, 3, 9, 188, 8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 5, 11, 194, 
	8, 11, 10, 11, 12, 11, 197, 9, 11, 1, 11, 1, 11, 3, 11, 201, 8, 11, 1, 
	11, 5, 11, 204, 8, 11, 10, 11, 12, 11, 207, 9, 11, 1, 11, 1, 11, 1, 12, 
	1, 12, 5, 12, 213, 8, 12, 10, 12, 12, 12, 216, 9, 12, 1, 12, 1, 12, 3, 
	12, 220, 8, 12, 1, 12, 5, 12, 223, 8, 12, 10, 12, 12, 12, 226, 9, 12, 1, 
	12, 1, 12, 1, 13, 1, 13, 3, 13, 232, 8, 13, 1, 13, 1, 13, 3, 13, 236, 8, 
	13, 1, 13, 1, 13, 3, 13, 240, 8, 13, 1, 13, 1, 13, 3, 13, 244, 8, 13, 5, 
	13, 246, 8, 13, 10, 13, 12, 13, 249, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 
	1, 14, 1, 14, 1, 14, 3, 14, 258, 8, 14, 1, 15, 3, 15, 261, 8, 15, 1, 15, 
	1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 270, 8, 15, 1, 15, 3, 
	15, 273, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 280, 8, 16, 1, 
	17, 5, 17, 283, 8, 17, 10, 17, 12, 17, 286, 9, 17, 1, 18, 1, 18, 3, 18, 
	290, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 297, 8, 18, 1, 18, 
	1, 18, 3, 18, 301, 8, 18, 1, 18, 1, 18, 3, 18, 305, 8, 18, 1, 18, 1, 18, 
	3, 18, 309, 8, 18, 1, 18, 1, 18, 3, 18, 313, 8, 18, 5, 18, 315, 8, 18, 
	10, 18, 12, 18, 318, 9, 18, 1, 18, 3, 18, 321, 8, 18, 1, 18, 1, 18, 3, 
	18, 325, 8, 18, 1, 18, 1, 18, 3, 18, 329, 8, 18, 1, 18, 3, 18, 332, 8, 
	18, 1, 18, 1, 18, 3, 18, 336, 8, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 
	3, 20, 343, 8, 20, 1, 20, 1, 20, 3, 20, 347, 8, 20, 1, 20, 1, 20, 3, 20, 
	351, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 
	22, 3, 22, 362, 8, 22, 1, 23, 1, 23, 5, 23, 366, 8, 23, 10, 23, 12, 23, 
	369, 9, 23, 1, 23, 1, 23, 3, 23, 373, 8, 23, 1, 24, 1, 24, 5, 24, 377, 
	8, 24, 10, 24, 12, 24, 380, 9, 24, 1, 24, 0, 0, 25, 0, 2, 4, 6, 8, 10, 
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 
	48, 0, 3, 1, 0, 53, 53, 4, 0, 3, 5, 7, 10, 26, 34, 36, 53, 1, 0, 35, 35, 
	424, 0, 51, 1, 0, 0, 0, 2, 76, 1, 0, 0, 0, 4, 89, 1, 0, 0, 0, 6, 100, 1, 
	0, 0, 0, 8, 105, 1, 0, 0, 0, 10, 126, 1, 0, 0, 0, 12, 141, 1, 0, 0, 0, 
	14, 146, 1, 0, 0, 0, 16, 153, 1, 0, 0, 0, 18, 187, 1, 0, 0, 0, 20, 189, 
	1, 0, 0, 0, 22, 191, 1, 0, 0, 0, 24, 210, 1, 0, 0, 0, 26, 229, 1, 0, 0, 
	0, 28, 252, 1, 0, 0, 0, 30, 260, 1, 0, 0, 0, 32, 274, 1, 0, 0, 0, 34, 284, 
	1, 0, 0, 0, 36, 289, 1, 0, 0, 0, 38, 339, 1, 0, 0, 0, 40, 342, 1, 0, 0, 
	0, 42, 354, 1, 0, 0, 0, 44, 361, 1, 0, 0, 0, 46, 363, 1, 0, 0, 0, 48, 378, 
	1, 0, 0, 0, 50, 52, 5, 9, 0, 0, 51, 50, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 
	52, 54, 1, 0, 0, 0, 53, 55, 5, 53, 0, 0, 54, 53, 1, 0, 0, 0, 54, 55, 1, 
	0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 59, 3, 2, 1, 0, 57, 59, 3, 4, 2, 0, 58, 
	56, 1, 0, 0, 0, 58, 57, 1, 0, 0, 0, 59, 61, 1, 0, 0, 0, 60, 62, 5, 53, 
	0, 0, 61, 60, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 65, 
	3, 6, 3, 0, 64, 66, 5, 53, 0, 0, 65, 64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 
	66, 67, 1, 0, 0, 0, 67, 69, 3, 34, 17, 0, 68, 70, 5, 53, 0, 0, 69, 68, 
	1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 5, 0, 0, 1, 
	72, 1, 1, 0, 0, 0, 73, 74, 3, 46, 23, 0, 74, 75, 5, 53, 0, 0, 75, 77, 1, 
	0, 0, 0, 76, 73, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 
	79, 5, 17, 0, 0, 79, 80, 5, 53, 0, 0, 80, 81, 5, 18, 0, 0, 81, 82, 5, 9, 
	0, 0, 82, 84, 5, 19, 0, 0, 83, 85, 5, 9, 0, 0, 84, 83, 1, 0, 0, 0, 84, 
	85, 1, 0, 0, 0, 85, 3, 1, 0, 0, 0, 86, 87, 3, 46, 23, 0, 87, 88, 5, 53, 
	0, 0, 88, 90, 1, 0, 0, 0, 89, 86, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 91, 
	1, 0, 0, 0, 91, 92, 5, 16, 0, 0, 92, 93, 5, 9, 0, 0, 93, 95, 5, 10, 0, 
	0, 94, 96, 5, 9, 0, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 5, 1, 
	0, 0, 0, 97, 99, 3, 8, 4, 0, 98, 97, 1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100, 
	98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 7, 1, 0, 0, 0, 102, 100, 1, 
	0, 0, 0, 103, 104, 5, 53, 0, 0, 104, 106, 3, 46, 23, 0, 105, 103, 1, 0, 
	0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 110, 5, 53, 0, 0, 
	108, 109, 5, 20, 0, 0, 109, 111, 5, 9, 0, 0, 110, 108, 1, 0, 0, 0, 110, 
	111, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113, 5, 21, 0, 0, 113, 114, 
	5, 9, 0, 0, 114, 122, 5, 10, 0, 0, 115, 116, 5, 53, 0, 0, 116, 118, 5, 
	23, 0, 0, 117, 119, 3, 10, 5, 0, 118, 117, 1, 0, 0, 0, 119, 120, 1, 0, 
	0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 123, 1, 0, 0, 0, 
	122, 115, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 9, 1, 0, 0, 0, 124, 125, 
	5, 53, 0, 0, 125, 127, 3, 46, 23, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 
	0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 129, 5, 53, 0, 0, 129, 130, 5, 24, 
	0, 0, 130, 131, 5, 9, 0, 0, 131, 133, 3, 12, 6, 0, 132, 134, 5, 9, 0, 0, 
	133, 132, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 
	137, 5, 1, 0, 0, 136, 138, 5, 9, 0, 0, 137, 136, 1, 0, 0, 0, 137, 138, 
	1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 3, 14, 7, 0, 140, 11, 1, 0, 
	0, 0, 141, 142, 5, 10, 0, 0, 142, 13, 1, 0, 0, 0, 143, 147, 3, 26, 13, 
	0, 144, 147, 3, 20, 10, 0, 145, 147, 3, 22, 11, 0, 146, 143, 1, 0, 0, 0, 
	146, 144, 1, 0, 0, 0, 146, 145, 1, 0, 0, 0, 147, 149, 1, 0, 0, 0, 148, 
	150, 3, 18, 9, 0, 149, 148, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 15, 
	1, 0, 0, 0, 151, 154, 3, 20, 10, 0, 152, 154, 3, 24, 12, 0, 153, 151, 1, 
	0, 0, 0, 153, 152, 1, 0, 0, 0, 154, 156, 1, 0, 0, 0, 155, 157, 3, 18, 9, 
	0, 156, 155, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 17, 1, 0, 0, 0, 158, 
	159, 5, 9, 0, 0, 159, 160, 5, 13, 0, 0, 160, 163, 5, 9, 0, 0, 161, 164, 
	3, 20, 10, 0, 162, 164, 3, 24, 12, 0, 163, 161, 1, 0, 0, 0, 163, 162, 1, 
	0, 0, 0, 164, 166, 1, 0, 0, 0, 165, 158, 1, 0, 0, 0, 166, 167, 1, 0, 0, 
	0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 188, 1, 0, 0, 0, 169, 
	170, 5, 9, 0, 0, 170, 171, 5, 12, 0, 0, 171, 174, 5, 9, 0, 0, 172, 175, 
	3, 20, 10, 0, 173, 175, 3, 24, 12, 0, 174, 172, 1, 0, 0, 0, 174, 173, 1, 
	0, 0, 0, 175, 177, 1, 0, 0, 0, 176, 169, 1, 0, 0, 0, 177, 178, 1, 0, 0, 
	0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 188, 1, 0, 0, 0, 180, 
	181, 5, 9, 0, 0, 181, 182, 5, 14, 0, 0, 182, 185, 5, 9, 0, 0, 183, 186, 
	3, 20, 10, 0, 184, 186, 3, 24, 12, 0, 185, 183, 1, 0, 0, 0, 185, 184, 1, 
	0, 0, 0, 186, 188, 1, 0, 0, 0, 187, 165, 1, 0, 0, 0, 187, 176, 1, 0, 0, 
	0, 187, 180, 1, 0, 0, 0, 188, 19, 1, 0, 0, 0, 189, 190, 3, 28, 14, 0, 190, 
	21, 1, 0, 0, 0, 191, 195, 5, 7, 0, 0, 192, 194, 5, 9, 0, 0, 193, 192, 1, 
	0, 0, 0, 194, 197, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 
	0, 196, 200, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 198, 201, 3, 14, 7, 0, 199, 
	201, 3, 24, 12, 0, 200, 198, 1, 0, 0, 0, 200, 199, 1, 0, 0, 0, 201, 205, 
	1, 0, 0, 0, 202, 204, 5, 9, 0, 0, 203, 202, 1, 0, 0, 0, 204, 207, 1, 0, 
	0, 0, 205, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 208, 1, 0, 0, 0, 
	207, 205, 1, 0, 0, 0, 208, 209, 5, 8, 0, 0, 209, 23, 1, 0, 0, 0, 210, 214, 
	5, 7, 0, 0, 211, 213, 5, 9, 0, 0, 212, 211, 1, 0, 0, 0, 213, 216, 1, 0, 
	0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 219, 1, 0, 0, 0, 
	216, 214, 1, 0, 0, 0, 217, 220, 3, 16, 8, 0, 218, 220, 3, 24, 12, 0, 219, 
	217, 1, 0, 0, 0, 219, 218, 1, 0, 0, 0, 220, 224, 1, 0, 0, 0, 221, 223, 
	5, 9, 0, 0, 222, 221, 1, 0, 0, 0, 223, 226, 1, 0, 0, 0, 224, 222, 1, 0, 
	0, 0, 224, 225, 1, 0, 0, 0, 225, 227, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 
	227, 228, 5, 8, 0, 0, 228, 25, 1, 0, 0, 0, 229, 231, 5, 5, 0, 0, 230, 232, 
	5, 9, 0, 0, 231, 230, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 233, 1, 0, 
	0, 0, 233, 235, 3, 30, 15, 0, 234, 236, 5, 9, 0, 0, 235, 234, 1, 0, 0, 
	0, 235, 236, 1, 0, 0, 0, 236, 247, 1, 0, 0, 0, 237, 239, 5, 2, 0, 0, 238, 
	240, 5, 9, 0, 0, 239, 238, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 241, 
	1, 0, 0, 0, 241, 243, 3, 30, 15, 0, 242, 244, 5, 9, 0, 0, 243, 242, 1, 
	0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 246, 1, 0, 0, 0, 245, 237, 1, 0, 0, 
	0, 246, 249, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 
	250, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 250, 251, 5, 33, 0, 0, 251, 27, 
	1, 0, 0, 0, 252, 257, 5, 10, 0, 0, 253, 254, 5, 9, 0, 0, 254, 255, 5, 15, 
	0, 0, 255, 256, 5, 9, 0, 0, 256, 258, 5, 10, 0, 0, 257, 253, 1, 0, 0, 0, 
	257, 258, 1, 0, 0, 0, 258, 29, 1, 0, 0, 0, 259, 261, 5, 53, 0, 0, 260, 
	259, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 269, 1, 0, 0, 0, 262, 270, 
	3, 32, 16, 0, 263, 264, 3, 32, 16, 0, 264, 265, 5, 9, 0, 0, 265, 266, 5, 
	25, 0, 0, 266, 267, 5, 9, 0, 0, 267, 268, 3, 38, 19, 0, 268, 270, 1, 0, 
	0, 0, 269, 262, 1, 0, 0, 0, 269, 263, 1, 0, 0, 0, 270, 272, 1, 0, 0, 0, 
	271, 273, 5, 53, 0, 0, 272, 271, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 
	31, 1, 0, 0, 0, 274, 279, 5, 10, 0, 0, 275, 276, 5, 1, 0, 0, 276, 280, 
	5, 41, 0, 0, 277, 278, 5, 11, 0, 0, 278, 280, 5, 10, 0, 0, 279, 275, 1, 
	0, 0, 0, 279, 277, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 33, 1, 0, 0, 
	0, 281, 283, 3, 36, 18, 0, 282, 281, 1, 0, 0, 0, 283, 286, 1, 0, 0, 0, 
	284, 282, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 35, 1, 0, 0, 0, 286, 284, 
	1, 0, 0, 0, 287, 288, 5, 53, 0, 0, 288, 290, 3, 46, 23, 0, 289, 287, 1, 
	0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 292, 5, 53, 0, 
	0, 292, 293, 5, 22, 0, 0, 293, 294, 5, 9, 0, 0, 294, 296, 3, 38, 19, 0, 
	295, 297, 5, 9, 0, 0, 296, 295, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 
	298, 1, 0, 0, 0, 298, 300, 5, 7, 0, 0, 299, 301, 5, 9, 0, 0, 300, 299, 
	1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 304, 3, 40, 
	20, 0, 303, 305, 5, 9, 0, 0, 304, 303, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 
	305, 316, 1, 0, 0, 0, 306, 308, 5, 2, 0, 0, 307, 309, 5, 9, 0, 0, 308, 
	307, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 312, 
	3, 40, 20, 0, 311, 313, 5, 9, 0, 0, 312, 311, 1, 0, 0, 0, 312, 313, 1, 
	0, 0, 0, 313, 315, 1, 0, 0, 0, 314, 306, 1, 0, 0, 0, 315, 318, 1, 0, 0, 
	0, 316, 314, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 320, 1, 0, 0, 0, 318, 
	316, 1, 0, 0, 0, 319, 321, 5, 53, 0, 0, 320, 319, 1, 0, 0, 0, 320, 321, 
	1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 324, 5, 8, 0, 0, 323, 325, 5, 9, 
	0, 0, 324, 323, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 
	326, 328, 5, 34, 0, 0, 327, 329, 5, 53, 0, 0, 328, 327, 1, 0, 0, 0, 328, 
	329, 1, 0, 0, 0, 329, 331, 1, 0, 0, 0, 330, 332, 5, 9, 0, 0, 331, 330, 
	1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 335, 3, 48, 
	24, 0, 334, 336, 5, 53, 0, 0, 335, 334, 1, 0, 0, 0, 335, 336, 1, 0, 0, 
	0, 336, 337, 1, 0, 0, 0, 337, 338, 5, 35, 0, 0, 338, 37, 1, 0, 0, 0, 339, 
	340, 5, 10, 0, 0, 340, 39, 1, 0, 0, 0, 341, 343, 5, 53, 0, 0, 342, 341, 
	1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 346, 3, 42, 
	21, 0, 345, 347, 5, 9, 0, 0, 346, 345, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 
	347, 348, 1, 0, 0, 0, 348, 350, 5, 1, 0, 0, 349, 351, 5, 9, 0, 0, 350, 
	349, 1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 353, 
	3, 44, 22, 0, 353, 41, 1, 0, 0, 0, 354, 355, 5, 10, 0, 0, 355, 43, 1, 0, 
	0, 0, 356, 362, 5, 55, 0, 0, 357, 358, 5, 54, 0, 0, 358, 359, 5, 3, 0, 
	0, 359, 360, 5, 55, 0, 0, 360, 362, 5, 4, 0, 0, 361, 356, 1, 0, 0, 0, 361, 
	357, 1, 0, 0, 0, 362, 45, 1, 0, 0, 0, 363, 367, 5, 11, 0, 0, 364, 366, 
	8, 0, 0, 0, 365, 364, 1, 0, 0, 0, 366, 369, 1, 0, 0, 0, 367, 365, 1, 0, 
	0, 0, 367, 368, 1, 0, 0, 0, 368, 372, 1, 0, 0, 0, 369, 367, 1, 0, 0, 0, 
	370, 371, 5, 53, 0, 0, 371, 373, 3, 46, 23, 0, 372, 370, 1, 0, 0, 0, 372, 
	373, 1, 0, 0, 0, 373, 47, 1, 0, 0, 0, 374, 377, 7, 1, 0, 0, 375, 377, 8, 
	2, 0, 0, 376, 374, 1, 0, 0, 0, 376, 375, 1, 0, 0, 0, 377, 380, 1, 0, 0, 
	0, 378, 376, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 49, 1, 0, 0, 0, 380, 
	378, 1, 0, 0, 0, 65, 51, 54, 58, 61, 65, 69, 76, 84, 89, 95, 100, 105, 
	110, 120, 122, 126, 133, 137, 146, 149, 153, 156, 163, 167, 174, 178, 185, 
	187, 195, 200, 205, 214, 219, 224, 231, 235, 239, 243, 247, 257, 260, 269, 
	272, 279, 284, 289, 296, 300, 304, 308, 312, 316, 320, 324, 328, 331, 335, 
	342, 346, 350, 361, 367, 372, 376, 378,
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
	OpenFGAParserEOF = antlr.TokenEOF
	OpenFGAParserCOLON = 1
	OpenFGAParserCOMMA = 2
	OpenFGAParserLESS = 3
	OpenFGAParserGREATER = 4
	OpenFGAParserLBRACKET = 5
	OpenFGAParserRBRACKET = 6
	OpenFGAParserLPAREN = 7
	OpenFGAParserRPAREN = 8
	OpenFGAParserWHITESPACE = 9
	OpenFGAParserIDENTIFIER = 10
	OpenFGAParserHASH = 11
	OpenFGAParserAND = 12
	OpenFGAParserOR = 13
	OpenFGAParserBUT_NOT = 14
	OpenFGAParserFROM = 15
	OpenFGAParserMODULE = 16
	OpenFGAParserMODEL = 17
	OpenFGAParserSCHEMA = 18
	OpenFGAParserSCHEMA_VERSION = 19
	OpenFGAParserEXTEND = 20
	OpenFGAParserTYPE = 21
	OpenFGAParserCONDITION = 22
	OpenFGAParserRELATIONS = 23
	OpenFGAParserDEFINE = 24
	OpenFGAParserKEYWORD_WITH = 25
	OpenFGAParserEQUALS = 26
	OpenFGAParserNOT_EQUALS = 27
	OpenFGAParserIN = 28
	OpenFGAParserLESS_EQUALS = 29
	OpenFGAParserGREATER_EQUALS = 30
	OpenFGAParserLOGICAL_AND = 31
	OpenFGAParserLOGICAL_OR = 32
	OpenFGAParserRPRACKET = 33
	OpenFGAParserLBRACE = 34
	OpenFGAParserRBRACE = 35
	OpenFGAParserDOT = 36
	OpenFGAParserMINUS = 37
	OpenFGAParserEXCLAM = 38
	OpenFGAParserQUESTIONMARK = 39
	OpenFGAParserPLUS = 40
	OpenFGAParserSTAR = 41
	OpenFGAParserSLASH = 42
	OpenFGAParserPERCENT = 43
	OpenFGAParserCEL_TRUE = 44
	OpenFGAParserCEL_FALSE = 45
	OpenFGAParserNUL = 46
	OpenFGAParserCEL_COMMENT = 47
	OpenFGAParserNUM_FLOAT = 48
	OpenFGAParserNUM_INT = 49
	OpenFGAParserNUM_UINT = 50
	OpenFGAParserSTRING = 51
	OpenFGAParserBYTES = 52
	OpenFGAParserNEWLINE = 53
	OpenFGAParserCONDITION_PARAM_CONTAINER = 54
	OpenFGAParserCONDITION_PARAM_TYPE = 55
)

// OpenFGAParser rules.
const (
	OpenFGAParserRULE_main = 0
	OpenFGAParserRULE_modelHeader = 1
	OpenFGAParserRULE_moduleHeader = 2
	OpenFGAParserRULE_typeDefs = 3
	OpenFGAParserRULE_typeDef = 4
	OpenFGAParserRULE_relationDeclaration = 5
	OpenFGAParserRULE_relationName = 6
	OpenFGAParserRULE_relationDef = 7
	OpenFGAParserRULE_relationDefNoDirect = 8
	OpenFGAParserRULE_relationDefPartials = 9
	OpenFGAParserRULE_relationDefGrouping = 10
	OpenFGAParserRULE_relationRecurse = 11
	OpenFGAParserRULE_relationRecurseNoDirect = 12
	OpenFGAParserRULE_relationDefDirectAssignment = 13
	OpenFGAParserRULE_relationDefRewrite = 14
	OpenFGAParserRULE_relationDefTypeRestriction = 15
	OpenFGAParserRULE_relationDefTypeRestrictionBase = 16
	OpenFGAParserRULE_conditions = 17
	OpenFGAParserRULE_condition = 18
	OpenFGAParserRULE_conditionName = 19
	OpenFGAParserRULE_conditionParameter = 20
	OpenFGAParserRULE_parameterName = 21
	OpenFGAParserRULE_parameterType = 22
	OpenFGAParserRULE_multiLineComment = 23
	OpenFGAParserRULE_conditionExpression = 24
)

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeDefs() ITypeDefsContext
	Conditions() IConditionsContext
	EOF() antlr.TerminalNode
	ModelHeader() IModelHeaderContext
	ModuleHeader() IModuleHeaderContext
	WHITESPACE() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

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

func InitEmptyMainContext(p *MainContext)  {
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

func (s *MainContext) TypeDefs() ITypeDefsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefsContext)
}

func (s *MainContext) Conditions() IConditionsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *MainContext) EOF() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserEOF, 0)
}

func (s *MainContext) ModelHeader() IModelHeaderContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModelHeaderContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModelHeaderContext)
}

func (s *MainContext) ModuleHeader() IModuleHeaderContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModuleHeaderContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModuleHeaderContext)
}

func (s *MainContext) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWHITESPACE, 0)
}

func (s *MainContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNEWLINE)
}

func (s *MainContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINE, i)
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
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(50)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(53)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(56)
			p.ModelHeader()
		}


	case 2:
		{
			p.SetState(57)
			p.ModuleHeader()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(60)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}
	{
		p.SetState(63)
		p.TypeDefs()
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(64)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}
	{
		p.SetState(67)
		p.Conditions()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(68)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(71)
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


// IModelHeaderContext is an interface to support dynamic dispatch.
type IModelHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSchemaVersion returns the schemaVersion token.
	GetSchemaVersion() antlr.Token 


	// SetSchemaVersion sets the schemaVersion token.
	SetSchemaVersion(antlr.Token) 


	// Getter signatures
	MODEL() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	SCHEMA() antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	SCHEMA_VERSION() antlr.TerminalNode
	MultiLineComment() IMultiLineCommentContext

	// IsModelHeaderContext differentiates from other interfaces.
	IsModelHeaderContext()
}

type ModelHeaderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	schemaVersion antlr.Token
}

func NewEmptyModelHeaderContext() *ModelHeaderContext {
	var p = new(ModelHeaderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_modelHeader
	return p
}

func InitEmptyModelHeaderContext(p *ModelHeaderContext)  {
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

func (s *ModelHeaderContext) GetSchemaVersion() antlr.Token { return s.schemaVersion }


func (s *ModelHeaderContext) SetSchemaVersion(v antlr.Token) { s.schemaVersion = v }


func (s *ModelHeaderContext) MODEL() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserMODEL, 0)
}

func (s *ModelHeaderContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNEWLINE)
}

func (s *ModelHeaderContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINE, i)
}

func (s *ModelHeaderContext) SCHEMA() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserSCHEMA, 0)
}

func (s *ModelHeaderContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserWHITESPACE)
}

func (s *ModelHeaderContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWHITESPACE, i)
}

func (s *ModelHeaderContext) SCHEMA_VERSION() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserSCHEMA_VERSION, 0)
}

func (s *ModelHeaderContext) MultiLineComment() IMultiLineCommentContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiLineCommentContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiLineCommentContext)
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
	p.EnterRule(localctx, 2, OpenFGAParserRULE_modelHeader)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserHASH {
		{
			p.SetState(73)
			p.MultiLineComment()
		}
		{
			p.SetState(74)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(78)
		p.Match(OpenFGAParserMODEL)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Match(OpenFGAParserNEWLINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Match(OpenFGAParserSCHEMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(OpenFGAParserWHITESPACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(82)

		var _m = p.Match(OpenFGAParserSCHEMA_VERSION)

		localctx.(*ModelHeaderContext).schemaVersion = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(83)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
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


// IModuleHeaderContext is an interface to support dynamic dispatch.
type IModuleHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetModuleName returns the moduleName token.
	GetModuleName() antlr.Token 


	// SetModuleName sets the moduleName token.
	SetModuleName(antlr.Token) 


	// Getter signatures
	MODULE() antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	MultiLineComment() IMultiLineCommentContext
	NEWLINE() antlr.TerminalNode

	// IsModuleHeaderContext differentiates from other interfaces.
	IsModuleHeaderContext()
}

type ModuleHeaderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	moduleName antlr.Token
}

func NewEmptyModuleHeaderContext() *ModuleHeaderContext {
	var p = new(ModuleHeaderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_moduleHeader
	return p
}

func InitEmptyModuleHeaderContext(p *ModuleHeaderContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_moduleHeader
}

func (*ModuleHeaderContext) IsModuleHeaderContext() {}

func NewModuleHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleHeaderContext {
	var p = new(ModuleHeaderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_moduleHeader

	return p
}

func (s *ModuleHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleHeaderContext) GetModuleName() antlr.Token { return s.moduleName }


func (s *ModuleHeaderContext) SetModuleName(v antlr.Token) { s.moduleName = v }


func (s *ModuleHeaderContext) MODULE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserMODULE, 0)
}

func (s *ModuleHeaderContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserWHITESPACE)
}

func (s *ModuleHeaderContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWHITESPACE, i)
}

func (s *ModuleHeaderContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserIDENTIFIER, 0)
}

func (s *ModuleHeaderContext) MultiLineComment() IMultiLineCommentContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiLineCommentContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiLineCommentContext)
}

func (s *ModuleHeaderContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINE, 0)
}

func (s *ModuleHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ModuleHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterModuleHeader(s)
	}
}

func (s *ModuleHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitModuleHeader(s)
	}
}




func (p *OpenFGAParser) ModuleHeader() (localctx IModuleHeaderContext) {
	localctx = NewModuleHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, OpenFGAParserRULE_moduleHeader)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserHASH {
		{
			p.SetState(86)
			p.MultiLineComment()
		}
		{
			p.SetState(87)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(91)
		p.Match(OpenFGAParserMODULE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Match(OpenFGAParserWHITESPACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(93)

		var _m = p.Match(OpenFGAParserIDENTIFIER)

		localctx.(*ModuleHeaderContext).moduleName = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(94)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
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

func InitEmptyTypeDefsContext(p *TypeDefsContext)  {
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(97)
				p.TypeDef()
			}


		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
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

	// GetTypeName returns the typeName token.
	GetTypeName() antlr.Token 


	// SetTypeName sets the typeName token.
	SetTypeName(antlr.Token) 


	// Getter signatures
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	TYPE() antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	MultiLineComment() IMultiLineCommentContext
	EXTEND() antlr.TerminalNode
	RELATIONS() antlr.TerminalNode
	AllRelationDeclaration() []IRelationDeclarationContext
	RelationDeclaration(i int) IRelationDeclarationContext

	// IsTypeDefContext differentiates from other interfaces.
	IsTypeDefContext()
}

type TypeDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	typeName antlr.Token
}

func NewEmptyTypeDefContext() *TypeDefContext {
	var p = new(TypeDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_typeDef
	return p
}

func InitEmptyTypeDefContext(p *TypeDefContext)  {
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

func (s *TypeDefContext) GetTypeName() antlr.Token { return s.typeName }


func (s *TypeDefContext) SetTypeName(v antlr.Token) { s.typeName = v }


func (s *TypeDefContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNEWLINE)
}

func (s *TypeDefContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINE, i)
}

func (s *TypeDefContext) TYPE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserTYPE, 0)
}

func (s *TypeDefContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserWHITESPACE)
}

func (s *TypeDefContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWHITESPACE, i)
}

func (s *TypeDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserIDENTIFIER, 0)
}

func (s *TypeDefContext) MultiLineComment() IMultiLineCommentContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiLineCommentContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiLineCommentContext)
}

func (s *TypeDefContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserEXTEND, 0)
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	p.SetState(105)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(103)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(104)
			p.MultiLineComment()
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}
	{
		p.SetState(107)
		p.Match(OpenFGAParserNEWLINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserEXTEND {
		{
			p.SetState(108)
			p.Match(OpenFGAParserEXTEND)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(112)
		p.Match(OpenFGAParserTYPE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(OpenFGAParserWHITESPACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(114)

		var _m = p.Match(OpenFGAParserIDENTIFIER)

		localctx.(*TypeDefContext).typeName = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(115)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Match(OpenFGAParserRELATIONS)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
					{
						p.SetState(117)
						p.RelationDeclaration()
					}




			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
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


// IRelationDeclarationContext is an interface to support dynamic dispatch.
type IRelationDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	DEFINE() antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	RelationName() IRelationNameContext
	COLON() antlr.TerminalNode
	RelationDef() IRelationDefContext
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

func InitEmptyRelationDeclarationContext(p *RelationDeclarationContext)  {
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

func (s *RelationDeclarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNEWLINE)
}

func (s *RelationDeclarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINE, i)
}

func (s *RelationDeclarationContext) DEFINE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserDEFINE, 0)
}

func (s *RelationDeclarationContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserWHITESPACE)
}

func (s *RelationDeclarationContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWHITESPACE, i)
}

func (s *RelationDeclarationContext) RelationName() IRelationNameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationNameContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefContext)
}

func (s *RelationDeclarationContext) MultiLineComment() IMultiLineCommentContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiLineCommentContext); ok {
			t = ctx.(antlr.RuleContext);
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
	p.SetState(126)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(124)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(125)
			p.MultiLineComment()
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}
	{
		p.SetState(128)
		p.Match(OpenFGAParserNEWLINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(OpenFGAParserDEFINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Match(OpenFGAParserWHITESPACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(131)
		p.RelationName()
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(132)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(135)
		p.Match(OpenFGAParserCOLON)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(136)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}

	{
		p.SetState(139)
		p.RelationDef()
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
	IDENTIFIER() antlr.TerminalNode

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

func InitEmptyRelationNameContext(p *RelationNameContext)  {
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

func (s *RelationNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserIDENTIFIER, 0)
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
	p.EnterRule(localctx, 12, OpenFGAParserRULE_relationName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(OpenFGAParserIDENTIFIER)
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


// IRelationDefContext is an interface to support dynamic dispatch.
type IRelationDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationDefDirectAssignment() IRelationDefDirectAssignmentContext
	RelationDefGrouping() IRelationDefGroupingContext
	RelationRecurse() IRelationRecurseContext
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

func InitEmptyRelationDefContext(p *RelationDefContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefDirectAssignmentContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefDirectAssignmentContext)
}

func (s *RelationDefContext) RelationDefGrouping() IRelationDefGroupingContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefGroupingContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefGroupingContext)
}

func (s *RelationDefContext) RelationRecurse() IRelationRecurseContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationRecurseContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationRecurseContext)
}

func (s *RelationDefContext) RelationDefPartials() IRelationDefPartialsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefPartialsContext); ok {
			t = ctx.(antlr.RuleContext);
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
	p.EnterRule(localctx, 14, OpenFGAParserRULE_relationDef)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case OpenFGAParserLBRACKET:
		{
			p.SetState(143)
			p.RelationDefDirectAssignment()
		}


	case OpenFGAParserIDENTIFIER:
		{
			p.SetState(144)
			p.RelationDefGrouping()
		}


	case OpenFGAParserLPAREN:
		{
			p.SetState(145)
			p.RelationRecurse()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(148)
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


// IRelationDefNoDirectContext is an interface to support dynamic dispatch.
type IRelationDefNoDirectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationDefGrouping() IRelationDefGroupingContext
	RelationRecurseNoDirect() IRelationRecurseNoDirectContext
	RelationDefPartials() IRelationDefPartialsContext

	// IsRelationDefNoDirectContext differentiates from other interfaces.
	IsRelationDefNoDirectContext()
}

type RelationDefNoDirectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationDefNoDirectContext() *RelationDefNoDirectContext {
	var p = new(RelationDefNoDirectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefNoDirect
	return p
}

func InitEmptyRelationDefNoDirectContext(p *RelationDefNoDirectContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefNoDirect
}

func (*RelationDefNoDirectContext) IsRelationDefNoDirectContext() {}

func NewRelationDefNoDirectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefNoDirectContext {
	var p = new(RelationDefNoDirectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefNoDirect

	return p
}

func (s *RelationDefNoDirectContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefNoDirectContext) RelationDefGrouping() IRelationDefGroupingContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefGroupingContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefGroupingContext)
}

func (s *RelationDefNoDirectContext) RelationRecurseNoDirect() IRelationRecurseNoDirectContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationRecurseNoDirectContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationRecurseNoDirectContext)
}

func (s *RelationDefNoDirectContext) RelationDefPartials() IRelationDefPartialsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefPartialsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefPartialsContext)
}

func (s *RelationDefNoDirectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefNoDirectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RelationDefNoDirectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefNoDirect(s)
	}
}

func (s *RelationDefNoDirectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefNoDirect(s)
	}
}




func (p *OpenFGAParser) RelationDefNoDirect() (localctx IRelationDefNoDirectContext) {
	localctx = NewRelationDefNoDirectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, OpenFGAParserRULE_relationDefNoDirect)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case OpenFGAParserIDENTIFIER:
		{
			p.SetState(151)
			p.RelationDefGrouping()
		}


	case OpenFGAParserLPAREN:
		{
			p.SetState(152)
			p.RelationRecurseNoDirect()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(155)
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
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode
	AllRelationDefGrouping() []IRelationDefGroupingContext
	RelationDefGrouping(i int) IRelationDefGroupingContext
	AllRelationRecurseNoDirect() []IRelationRecurseNoDirectContext
	RelationRecurseNoDirect(i int) IRelationRecurseNoDirectContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode
	BUT_NOT() antlr.TerminalNode

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

func InitEmptyRelationDefPartialsContext(p *RelationDefPartialsContext)  {
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

func (s *RelationDefPartialsContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserWHITESPACE)
}

func (s *RelationDefPartialsContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWHITESPACE, i)
}

func (s *RelationDefPartialsContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserOR)
}

func (s *RelationDefPartialsContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserOR, i)
}

func (s *RelationDefPartialsContext) AllRelationDefGrouping() []IRelationDefGroupingContext {
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

func (s *RelationDefPartialsContext) RelationDefGrouping(i int) IRelationDefGroupingContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefGroupingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *RelationDefPartialsContext) AllRelationRecurseNoDirect() []IRelationRecurseNoDirectContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationRecurseNoDirectContext); ok {
			len++
		}
	}

	tst := make([]IRelationRecurseNoDirectContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationRecurseNoDirectContext); ok {
			tst[i] = t.(IRelationRecurseNoDirectContext)
			i++
		}
	}

	return tst
}

func (s *RelationDefPartialsContext) RelationRecurseNoDirect(i int) IRelationRecurseNoDirectContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationRecurseNoDirectContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationRecurseNoDirectContext)
}

func (s *RelationDefPartialsContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserAND)
}

func (s *RelationDefPartialsContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserAND, i)
}

func (s *RelationDefPartialsContext) BUT_NOT() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserBUT_NOT, 0)
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
	p.EnterRule(localctx, 18, OpenFGAParserRULE_relationDefPartials)
	var _alt int

	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
					{
						p.SetState(158)
						p.Match(OpenFGAParserWHITESPACE)
						if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
						}
					}
					{
						p.SetState(159)
						p.Match(OpenFGAParserOR)
						if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
						}
					}
					{
						p.SetState(160)
						p.Match(OpenFGAParserWHITESPACE)
						if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
						}
					}
					p.SetState(163)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}

					switch p.GetTokenStream().LA(1) {
					case OpenFGAParserIDENTIFIER:
						{
							p.SetState(161)
							p.RelationDefGrouping()
						}


					case OpenFGAParserLPAREN:
						{
							p.SetState(162)
							p.RelationRecurseNoDirect()
						}



					default:
						p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
						goto errorExit
					}




			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
					{
						p.SetState(169)
						p.Match(OpenFGAParserWHITESPACE)
						if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
						}
					}
					{
						p.SetState(170)
						p.Match(OpenFGAParserAND)
						if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
						}
					}
					{
						p.SetState(171)
						p.Match(OpenFGAParserWHITESPACE)
						if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
						}
					}
					p.SetState(174)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}

					switch p.GetTokenStream().LA(1) {
					case OpenFGAParserIDENTIFIER:
						{
							p.SetState(172)
							p.RelationDefGrouping()
						}


					case OpenFGAParserLPAREN:
						{
							p.SetState(173)
							p.RelationRecurseNoDirect()
						}



					default:
						p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
						goto errorExit
					}




			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(178)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(180)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Match(OpenFGAParserBUT_NOT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case OpenFGAParserIDENTIFIER:
			{
				p.SetState(183)
				p.RelationDefGrouping()
			}


		case OpenFGAParserLPAREN:
			{
				p.SetState(184)
				p.RelationRecurseNoDirect()
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
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

func InitEmptyRelationDefGroupingContext(p *RelationDefGroupingContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefRewriteContext); ok {
			t = ctx.(antlr.RuleContext);
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
	p.EnterRule(localctx, 20, OpenFGAParserRULE_relationDefGrouping)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
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


// IRelationRecurseContext is an interface to support dynamic dispatch.
type IRelationRecurseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	RelationDef() IRelationDefContext
	RelationRecurseNoDirect() IRelationRecurseNoDirectContext
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode

	// IsRelationRecurseContext differentiates from other interfaces.
	IsRelationRecurseContext()
}

type RelationRecurseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationRecurseContext() *RelationRecurseContext {
	var p = new(RelationRecurseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationRecurse
	return p
}

func InitEmptyRelationRecurseContext(p *RelationRecurseContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationRecurse
}

func (*RelationRecurseContext) IsRelationRecurseContext() {}

func NewRelationRecurseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationRecurseContext {
	var p = new(RelationRecurseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationRecurse

	return p
}

func (s *RelationRecurseContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationRecurseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserLPAREN, 0)
}

func (s *RelationRecurseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserRPAREN, 0)
}

func (s *RelationRecurseContext) RelationDef() IRelationDefContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefContext)
}

func (s *RelationRecurseContext) RelationRecurseNoDirect() IRelationRecurseNoDirectContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationRecurseNoDirectContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationRecurseNoDirectContext)
}

func (s *RelationRecurseContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserWHITESPACE)
}

func (s *RelationRecurseContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWHITESPACE, i)
}

func (s *RelationRecurseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationRecurseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RelationRecurseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationRecurse(s)
	}
}

func (s *RelationRecurseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationRecurse(s)
	}
}




func (p *OpenFGAParser) RelationRecurse() (localctx IRelationRecurseContext) {
	localctx = NewRelationRecurseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, OpenFGAParserRULE_relationRecurse)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(OpenFGAParserLPAREN)
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


	for _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(192)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(198)
			p.RelationDef()
		}


	case 2:
		{
			p.SetState(199)
			p.RelationRecurseNoDirect()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(202)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(208)
		p.Match(OpenFGAParserRPAREN)
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


// IRelationRecurseNoDirectContext is an interface to support dynamic dispatch.
type IRelationRecurseNoDirectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	RelationDefNoDirect() IRelationDefNoDirectContext
	RelationRecurseNoDirect() IRelationRecurseNoDirectContext
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode

	// IsRelationRecurseNoDirectContext differentiates from other interfaces.
	IsRelationRecurseNoDirectContext()
}

type RelationRecurseNoDirectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationRecurseNoDirectContext() *RelationRecurseNoDirectContext {
	var p = new(RelationRecurseNoDirectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationRecurseNoDirect
	return p
}

func InitEmptyRelationRecurseNoDirectContext(p *RelationRecurseNoDirectContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationRecurseNoDirect
}

func (*RelationRecurseNoDirectContext) IsRelationRecurseNoDirectContext() {}

func NewRelationRecurseNoDirectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationRecurseNoDirectContext {
	var p = new(RelationRecurseNoDirectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationRecurseNoDirect

	return p
}

func (s *RelationRecurseNoDirectContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationRecurseNoDirectContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserLPAREN, 0)
}

func (s *RelationRecurseNoDirectContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserRPAREN, 0)
}

func (s *RelationRecurseNoDirectContext) RelationDefNoDirect() IRelationDefNoDirectContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefNoDirectContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefNoDirectContext)
}

func (s *RelationRecurseNoDirectContext) RelationRecurseNoDirect() IRelationRecurseNoDirectContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationRecurseNoDirectContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationRecurseNoDirectContext)
}

func (s *RelationRecurseNoDirectContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserWHITESPACE)
}

func (s *RelationRecurseNoDirectContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWHITESPACE, i)
}

func (s *RelationRecurseNoDirectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationRecurseNoDirectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RelationRecurseNoDirectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationRecurseNoDirect(s)
	}
}

func (s *RelationRecurseNoDirectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationRecurseNoDirect(s)
	}
}




func (p *OpenFGAParser) RelationRecurseNoDirect() (localctx IRelationRecurseNoDirectContext) {
	localctx = NewRelationRecurseNoDirectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, OpenFGAParserRULE_relationRecurseNoDirect)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(OpenFGAParserLPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(211)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(217)
			p.RelationDefNoDirect()
		}


	case 2:
		{
			p.SetState(218)
			p.RelationRecurseNoDirect()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(221)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(227)
		p.Match(OpenFGAParserRPAREN)
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


// IRelationDefDirectAssignmentContext is an interface to support dynamic dispatch.
type IRelationDefDirectAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	AllRelationDefTypeRestriction() []IRelationDefTypeRestrictionContext
	RelationDefTypeRestriction(i int) IRelationDefTypeRestrictionContext
	RPRACKET() antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
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

func InitEmptyRelationDefDirectAssignmentContext(p *RelationDefDirectAssignmentContext)  {
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

func (s *RelationDefDirectAssignmentContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserLBRACKET, 0)
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefTypeRestrictionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *RelationDefDirectAssignmentContext) RPRACKET() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserRPRACKET, 0)
}

func (s *RelationDefDirectAssignmentContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserWHITESPACE)
}

func (s *RelationDefDirectAssignmentContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWHITESPACE, i)
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
	p.EnterRule(localctx, 26, OpenFGAParserRULE_relationDefDirectAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(OpenFGAParserLBRACKET)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(230)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(233)
		p.RelationDefTypeRestriction()
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(234)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == OpenFGAParserCOMMA {
		{
			p.SetState(237)
			p.Match(OpenFGAParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == OpenFGAParserWHITESPACE {
			{
				p.SetState(238)
				p.Match(OpenFGAParserWHITESPACE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}
		{
			p.SetState(241)
			p.RelationDefTypeRestriction()
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == OpenFGAParserWHITESPACE {
			{
				p.SetState(242)
				p.Match(OpenFGAParserWHITESPACE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}


		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(250)
		p.Match(OpenFGAParserRPRACKET)
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

	// GetRewriteComputedusersetName returns the rewriteComputedusersetName token.
	GetRewriteComputedusersetName() antlr.Token 

	// GetRewriteTuplesetName returns the rewriteTuplesetName token.
	GetRewriteTuplesetName() antlr.Token 


	// SetRewriteComputedusersetName sets the rewriteComputedusersetName token.
	SetRewriteComputedusersetName(antlr.Token) 

	// SetRewriteTuplesetName sets the rewriteTuplesetName token.
	SetRewriteTuplesetName(antlr.Token) 


	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	FROM() antlr.TerminalNode

	// IsRelationDefRewriteContext differentiates from other interfaces.
	IsRelationDefRewriteContext()
}

type RelationDefRewriteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	rewriteComputedusersetName antlr.Token
	rewriteTuplesetName antlr.Token
}

func NewEmptyRelationDefRewriteContext() *RelationDefRewriteContext {
	var p = new(RelationDefRewriteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefRewrite
	return p
}

func InitEmptyRelationDefRewriteContext(p *RelationDefRewriteContext)  {
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

func (s *RelationDefRewriteContext) GetRewriteComputedusersetName() antlr.Token { return s.rewriteComputedusersetName }

func (s *RelationDefRewriteContext) GetRewriteTuplesetName() antlr.Token { return s.rewriteTuplesetName }


func (s *RelationDefRewriteContext) SetRewriteComputedusersetName(v antlr.Token) { s.rewriteComputedusersetName = v }

func (s *RelationDefRewriteContext) SetRewriteTuplesetName(v antlr.Token) { s.rewriteTuplesetName = v }


func (s *RelationDefRewriteContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserIDENTIFIER)
}

func (s *RelationDefRewriteContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserIDENTIFIER, i)
}

func (s *RelationDefRewriteContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserWHITESPACE)
}

func (s *RelationDefRewriteContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWHITESPACE, i)
}

func (s *RelationDefRewriteContext) FROM() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserFROM, 0)
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
	p.EnterRule(localctx, 28, OpenFGAParserRULE_relationDefRewrite)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)

		var _m = p.Match(OpenFGAParserIDENTIFIER)

		localctx.(*RelationDefRewriteContext).rewriteComputedusersetName = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(253)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(254)
			p.Match(OpenFGAParserFROM)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(256)

			var _m = p.Match(OpenFGAParserIDENTIFIER)

			localctx.(*RelationDefRewriteContext).rewriteTuplesetName = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
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


// IRelationDefTypeRestrictionContext is an interface to support dynamic dispatch.
type IRelationDefTypeRestrictionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationDefTypeRestrictionBase() IRelationDefTypeRestrictionBaseContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	KEYWORD_WITH() antlr.TerminalNode
	ConditionName() IConditionNameContext

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

func InitEmptyRelationDefTypeRestrictionContext(p *RelationDefTypeRestrictionContext)  {
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

func (s *RelationDefTypeRestrictionContext) RelationDefTypeRestrictionBase() IRelationDefTypeRestrictionBaseContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefTypeRestrictionBaseContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefTypeRestrictionBaseContext)
}

func (s *RelationDefTypeRestrictionContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNEWLINE)
}

func (s *RelationDefTypeRestrictionContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINE, i)
}

func (s *RelationDefTypeRestrictionContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserWHITESPACE)
}

func (s *RelationDefTypeRestrictionContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWHITESPACE, i)
}

func (s *RelationDefTypeRestrictionContext) KEYWORD_WITH() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserKEYWORD_WITH, 0)
}

func (s *RelationDefTypeRestrictionContext) ConditionName() IConditionNameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionNameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionNameContext)
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
	p.EnterRule(localctx, 30, OpenFGAParserRULE_relationDefTypeRestriction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(259)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(262)
			p.RelationDefTypeRestrictionBase()
		}


	case 2:
		{
			p.SetState(263)
			p.RelationDefTypeRestrictionBase()
		}
		{
			p.SetState(264)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(265)
			p.Match(OpenFGAParserKEYWORD_WITH)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(266)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(267)
			p.ConditionName()
		}


	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(271)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
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


// IRelationDefTypeRestrictionBaseContext is an interface to support dynamic dispatch.
type IRelationDefTypeRestrictionBaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRelationDefTypeRestrictionType returns the relationDefTypeRestrictionType token.
	GetRelationDefTypeRestrictionType() antlr.Token 

	// GetRelationDefTypeRestrictionWildcard returns the relationDefTypeRestrictionWildcard token.
	GetRelationDefTypeRestrictionWildcard() antlr.Token 

	// GetRelationDefTypeRestrictionRelation returns the relationDefTypeRestrictionRelation token.
	GetRelationDefTypeRestrictionRelation() antlr.Token 


	// SetRelationDefTypeRestrictionType sets the relationDefTypeRestrictionType token.
	SetRelationDefTypeRestrictionType(antlr.Token) 

	// SetRelationDefTypeRestrictionWildcard sets the relationDefTypeRestrictionWildcard token.
	SetRelationDefTypeRestrictionWildcard(antlr.Token) 

	// SetRelationDefTypeRestrictionRelation sets the relationDefTypeRestrictionRelation token.
	SetRelationDefTypeRestrictionRelation(antlr.Token) 


	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	COLON() antlr.TerminalNode
	HASH() antlr.TerminalNode
	STAR() antlr.TerminalNode

	// IsRelationDefTypeRestrictionBaseContext differentiates from other interfaces.
	IsRelationDefTypeRestrictionBaseContext()
}

type RelationDefTypeRestrictionBaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	relationDefTypeRestrictionType antlr.Token
	relationDefTypeRestrictionWildcard antlr.Token
	relationDefTypeRestrictionRelation antlr.Token
}

func NewEmptyRelationDefTypeRestrictionBaseContext() *RelationDefTypeRestrictionBaseContext {
	var p = new(RelationDefTypeRestrictionBaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestrictionBase
	return p
}

func InitEmptyRelationDefTypeRestrictionBaseContext(p *RelationDefTypeRestrictionBaseContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestrictionBase
}

func (*RelationDefTypeRestrictionBaseContext) IsRelationDefTypeRestrictionBaseContext() {}

func NewRelationDefTypeRestrictionBaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationDefTypeRestrictionBaseContext {
	var p = new(RelationDefTypeRestrictionBaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationDefTypeRestrictionBase

	return p
}

func (s *RelationDefTypeRestrictionBaseContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationDefTypeRestrictionBaseContext) GetRelationDefTypeRestrictionType() antlr.Token { return s.relationDefTypeRestrictionType }

func (s *RelationDefTypeRestrictionBaseContext) GetRelationDefTypeRestrictionWildcard() antlr.Token { return s.relationDefTypeRestrictionWildcard }

func (s *RelationDefTypeRestrictionBaseContext) GetRelationDefTypeRestrictionRelation() antlr.Token { return s.relationDefTypeRestrictionRelation }


func (s *RelationDefTypeRestrictionBaseContext) SetRelationDefTypeRestrictionType(v antlr.Token) { s.relationDefTypeRestrictionType = v }

func (s *RelationDefTypeRestrictionBaseContext) SetRelationDefTypeRestrictionWildcard(v antlr.Token) { s.relationDefTypeRestrictionWildcard = v }

func (s *RelationDefTypeRestrictionBaseContext) SetRelationDefTypeRestrictionRelation(v antlr.Token) { s.relationDefTypeRestrictionRelation = v }


func (s *RelationDefTypeRestrictionBaseContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserIDENTIFIER)
}

func (s *RelationDefTypeRestrictionBaseContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserIDENTIFIER, i)
}

func (s *RelationDefTypeRestrictionBaseContext) COLON() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserCOLON, 0)
}

func (s *RelationDefTypeRestrictionBaseContext) HASH() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserHASH, 0)
}

func (s *RelationDefTypeRestrictionBaseContext) STAR() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserSTAR, 0)
}

func (s *RelationDefTypeRestrictionBaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefTypeRestrictionBaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RelationDefTypeRestrictionBaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterRelationDefTypeRestrictionBase(s)
	}
}

func (s *RelationDefTypeRestrictionBaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitRelationDefTypeRestrictionBase(s)
	}
}




func (p *OpenFGAParser) RelationDefTypeRestrictionBase() (localctx IRelationDefTypeRestrictionBaseContext) {
	localctx = NewRelationDefTypeRestrictionBaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, OpenFGAParserRULE_relationDefTypeRestrictionBase)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)

		var _m = p.Match(OpenFGAParserIDENTIFIER)

		localctx.(*RelationDefTypeRestrictionBaseContext).relationDefTypeRestrictionType = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case OpenFGAParserCOLON:
		{
			p.SetState(275)
			p.Match(OpenFGAParserCOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(276)

			var _m = p.Match(OpenFGAParserSTAR)

			localctx.(*RelationDefTypeRestrictionBaseContext).relationDefTypeRestrictionWildcard = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	case OpenFGAParserHASH:
		{
			p.SetState(277)
			p.Match(OpenFGAParserHASH)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(278)

			var _m = p.Match(OpenFGAParserIDENTIFIER)

			localctx.(*RelationDefTypeRestrictionBaseContext).relationDefTypeRestrictionRelation = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	case OpenFGAParserCOMMA, OpenFGAParserWHITESPACE, OpenFGAParserRPRACKET, OpenFGAParserNEWLINE:



	default:
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


// IConditionsContext is an interface to support dynamic dispatch.
type IConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCondition() []IConditionContext
	Condition(i int) IConditionContext

	// IsConditionsContext differentiates from other interfaces.
	IsConditionsContext()
}

type ConditionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsContext() *ConditionsContext {
	var p = new(ConditionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_conditions
	return p
}

func InitEmptyConditionsContext(p *ConditionsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_conditions
}

func (*ConditionsContext) IsConditionsContext() {}

func NewConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsContext {
	var p = new(ConditionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_conditions

	return p
}

func (s *ConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *ConditionsContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterConditions(s)
	}
}

func (s *ConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitConditions(s)
	}
}




func (p *OpenFGAParser) Conditions() (localctx IConditionsContext) {
	localctx = NewConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, OpenFGAParserRULE_conditions)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(281)
				p.Condition()
			}


		}
		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
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


// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	CONDITION() antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	ConditionName() IConditionNameContext
	LPAREN() antlr.TerminalNode
	AllConditionParameter() []IConditionParameterContext
	ConditionParameter(i int) IConditionParameterContext
	RPAREN() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	ConditionExpression() IConditionExpressionContext
	RBRACE() antlr.TerminalNode
	MultiLineComment() IMultiLineCommentContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNEWLINE)
}

func (s *ConditionContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINE, i)
}

func (s *ConditionContext) CONDITION() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserCONDITION, 0)
}

func (s *ConditionContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserWHITESPACE)
}

func (s *ConditionContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWHITESPACE, i)
}

func (s *ConditionContext) ConditionName() IConditionNameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionNameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionNameContext)
}

func (s *ConditionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserLPAREN, 0)
}

func (s *ConditionContext) AllConditionParameter() []IConditionParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionParameterContext); ok {
			len++
		}
	}

	tst := make([]IConditionParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionParameterContext); ok {
			tst[i] = t.(IConditionParameterContext)
			i++
		}
	}

	return tst
}

func (s *ConditionContext) ConditionParameter(i int) IConditionParameterContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionParameterContext)
}

func (s *ConditionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserRPAREN, 0)
}

func (s *ConditionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserLBRACE, 0)
}

func (s *ConditionContext) ConditionExpression() IConditionExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionExpressionContext)
}

func (s *ConditionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserRBRACE, 0)
}

func (s *ConditionContext) MultiLineComment() IMultiLineCommentContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiLineCommentContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiLineCommentContext)
}

func (s *ConditionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserCOMMA)
}

func (s *ConditionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserCOMMA, i)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitCondition(s)
	}
}




func (p *OpenFGAParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, OpenFGAParserRULE_condition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(289)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(287)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(288)
			p.MultiLineComment()
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}
	{
		p.SetState(291)
		p.Match(OpenFGAParserNEWLINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(292)
		p.Match(OpenFGAParserCONDITION)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(293)
		p.Match(OpenFGAParserWHITESPACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(294)
		p.ConditionName()
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(295)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(298)
		p.Match(OpenFGAParserLPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(299)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(302)
		p.ConditionParameter()
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(303)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == OpenFGAParserCOMMA {
		{
			p.SetState(306)
			p.Match(OpenFGAParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == OpenFGAParserWHITESPACE {
			{
				p.SetState(307)
				p.Match(OpenFGAParserWHITESPACE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}
		{
			p.SetState(310)
			p.ConditionParameter()
		}
		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == OpenFGAParserWHITESPACE {
			{
				p.SetState(311)
				p.Match(OpenFGAParserWHITESPACE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}


		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(319)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(322)
		p.Match(OpenFGAParserRPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(323)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(326)
		p.Match(OpenFGAParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(328)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(327)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(330)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}
	{
		p.SetState(333)
		p.ConditionExpression()
	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(334)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(337)
		p.Match(OpenFGAParserRBRACE)
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


// IConditionNameContext is an interface to support dynamic dispatch.
type IConditionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsConditionNameContext differentiates from other interfaces.
	IsConditionNameContext()
}

type ConditionNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionNameContext() *ConditionNameContext {
	var p = new(ConditionNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_conditionName
	return p
}

func InitEmptyConditionNameContext(p *ConditionNameContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_conditionName
}

func (*ConditionNameContext) IsConditionNameContext() {}

func NewConditionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionNameContext {
	var p = new(ConditionNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_conditionName

	return p
}

func (s *ConditionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserIDENTIFIER, 0)
}

func (s *ConditionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConditionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterConditionName(s)
	}
}

func (s *ConditionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitConditionName(s)
	}
}




func (p *OpenFGAParser) ConditionName() (localctx IConditionNameContext) {
	localctx = NewConditionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, OpenFGAParserRULE_conditionName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(OpenFGAParserIDENTIFIER)
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


// IConditionParameterContext is an interface to support dynamic dispatch.
type IConditionParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParameterName() IParameterNameContext
	COLON() antlr.TerminalNode
	ParameterType() IParameterTypeContext
	NEWLINE() antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode

	// IsConditionParameterContext differentiates from other interfaces.
	IsConditionParameterContext()
}

type ConditionParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionParameterContext() *ConditionParameterContext {
	var p = new(ConditionParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_conditionParameter
	return p
}

func InitEmptyConditionParameterContext(p *ConditionParameterContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_conditionParameter
}

func (*ConditionParameterContext) IsConditionParameterContext() {}

func NewConditionParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionParameterContext {
	var p = new(ConditionParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_conditionParameter

	return p
}

func (s *ConditionParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionParameterContext) ParameterName() IParameterNameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterNameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterNameContext)
}

func (s *ConditionParameterContext) COLON() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserCOLON, 0)
}

func (s *ConditionParameterContext) ParameterType() IParameterTypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterTypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterTypeContext)
}

func (s *ConditionParameterContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINE, 0)
}

func (s *ConditionParameterContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserWHITESPACE)
}

func (s *ConditionParameterContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWHITESPACE, i)
}

func (s *ConditionParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConditionParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterConditionParameter(s)
	}
}

func (s *ConditionParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitConditionParameter(s)
	}
}




func (p *OpenFGAParser) ConditionParameter() (localctx IConditionParameterContext) {
	localctx = NewConditionParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, OpenFGAParserRULE_conditionParameter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(341)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(344)
		p.ParameterName()
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(345)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(348)
		p.Match(OpenFGAParserCOLON)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(349)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(352)
		p.ParameterType()
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


// IParameterNameContext is an interface to support dynamic dispatch.
type IParameterNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsParameterNameContext differentiates from other interfaces.
	IsParameterNameContext()
}

type ParameterNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterNameContext() *ParameterNameContext {
	var p = new(ParameterNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_parameterName
	return p
}

func InitEmptyParameterNameContext(p *ParameterNameContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_parameterName
}

func (*ParameterNameContext) IsParameterNameContext() {}

func NewParameterNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterNameContext {
	var p = new(ParameterNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_parameterName

	return p
}

func (s *ParameterNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserIDENTIFIER, 0)
}

func (s *ParameterNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParameterNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterParameterName(s)
	}
}

func (s *ParameterNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitParameterName(s)
	}
}




func (p *OpenFGAParser) ParameterName() (localctx IParameterNameContext) {
	localctx = NewParameterNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, OpenFGAParserRULE_parameterName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)
		p.Match(OpenFGAParserIDENTIFIER)
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


// IParameterTypeContext is an interface to support dynamic dispatch.
type IParameterTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONDITION_PARAM_TYPE() antlr.TerminalNode
	CONDITION_PARAM_CONTAINER() antlr.TerminalNode
	LESS() antlr.TerminalNode
	GREATER() antlr.TerminalNode

	// IsParameterTypeContext differentiates from other interfaces.
	IsParameterTypeContext()
}

type ParameterTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterTypeContext() *ParameterTypeContext {
	var p = new(ParameterTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_parameterType
	return p
}

func InitEmptyParameterTypeContext(p *ParameterTypeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_parameterType
}

func (*ParameterTypeContext) IsParameterTypeContext() {}

func NewParameterTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterTypeContext {
	var p = new(ParameterTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_parameterType

	return p
}

func (s *ParameterTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterTypeContext) CONDITION_PARAM_TYPE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserCONDITION_PARAM_TYPE, 0)
}

func (s *ParameterTypeContext) CONDITION_PARAM_CONTAINER() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserCONDITION_PARAM_CONTAINER, 0)
}

func (s *ParameterTypeContext) LESS() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserLESS, 0)
}

func (s *ParameterTypeContext) GREATER() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserGREATER, 0)
}

func (s *ParameterTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParameterTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterParameterType(s)
	}
}

func (s *ParameterTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitParameterType(s)
	}
}




func (p *OpenFGAParser) ParameterType() (localctx IParameterTypeContext) {
	localctx = NewParameterTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, OpenFGAParserRULE_parameterType)
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case OpenFGAParserCONDITION_PARAM_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(356)
			p.Match(OpenFGAParserCONDITION_PARAM_TYPE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case OpenFGAParserCONDITION_PARAM_CONTAINER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(357)
			p.Match(OpenFGAParserCONDITION_PARAM_CONTAINER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(358)
			p.Match(OpenFGAParserLESS)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(359)
			p.Match(OpenFGAParserCONDITION_PARAM_TYPE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(360)
			p.Match(OpenFGAParserGREATER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
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


// IMultiLineCommentContext is an interface to support dynamic dispatch.
type IMultiLineCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HASH() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	MultiLineComment() IMultiLineCommentContext

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

func InitEmptyMultiLineCommentContext(p *MultiLineCommentContext)  {
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

func (s *MultiLineCommentContext) HASH() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserHASH, 0)
}

func (s *MultiLineCommentContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNEWLINE)
}

func (s *MultiLineCommentContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINE, i)
}

func (s *MultiLineCommentContext) MultiLineComment() IMultiLineCommentContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiLineCommentContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiLineCommentContext)
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
	p.EnterRule(localctx, 46, OpenFGAParserRULE_multiLineComment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(OpenFGAParserHASH)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 63050394783186942) != 0) {
		{
			p.SetState(364)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == OpenFGAParserNEWLINE  {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}


		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(370)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(371)
			p.MultiLineComment()
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


// IConditionExpressionContext is an interface to support dynamic dispatch.
type IConditionExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllEQUALS() []antlr.TerminalNode
	EQUALS(i int) antlr.TerminalNode
	AllNOT_EQUALS() []antlr.TerminalNode
	NOT_EQUALS(i int) antlr.TerminalNode
	AllIN() []antlr.TerminalNode
	IN(i int) antlr.TerminalNode
	AllLESS() []antlr.TerminalNode
	LESS(i int) antlr.TerminalNode
	AllLESS_EQUALS() []antlr.TerminalNode
	LESS_EQUALS(i int) antlr.TerminalNode
	AllGREATER_EQUALS() []antlr.TerminalNode
	GREATER_EQUALS(i int) antlr.TerminalNode
	AllGREATER() []antlr.TerminalNode
	GREATER(i int) antlr.TerminalNode
	AllLOGICAL_AND() []antlr.TerminalNode
	LOGICAL_AND(i int) antlr.TerminalNode
	AllLOGICAL_OR() []antlr.TerminalNode
	LOGICAL_OR(i int) antlr.TerminalNode
	AllLBRACKET() []antlr.TerminalNode
	LBRACKET(i int) antlr.TerminalNode
	AllRPRACKET() []antlr.TerminalNode
	RPRACKET(i int) antlr.TerminalNode
	AllLBRACE() []antlr.TerminalNode
	LBRACE(i int) antlr.TerminalNode
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode
	AllEXCLAM() []antlr.TerminalNode
	EXCLAM(i int) antlr.TerminalNode
	AllQUESTIONMARK() []antlr.TerminalNode
	QUESTIONMARK(i int) antlr.TerminalNode
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	AllSTAR() []antlr.TerminalNode
	STAR(i int) antlr.TerminalNode
	AllSLASH() []antlr.TerminalNode
	SLASH(i int) antlr.TerminalNode
	AllPERCENT() []antlr.TerminalNode
	PERCENT(i int) antlr.TerminalNode
	AllCEL_TRUE() []antlr.TerminalNode
	CEL_TRUE(i int) antlr.TerminalNode
	AllCEL_FALSE() []antlr.TerminalNode
	CEL_FALSE(i int) antlr.TerminalNode
	AllNUL() []antlr.TerminalNode
	NUL(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllCEL_COMMENT() []antlr.TerminalNode
	CEL_COMMENT(i int) antlr.TerminalNode
	AllNUM_FLOAT() []antlr.TerminalNode
	NUM_FLOAT(i int) antlr.TerminalNode
	AllNUM_INT() []antlr.TerminalNode
	NUM_INT(i int) antlr.TerminalNode
	AllNUM_UINT() []antlr.TerminalNode
	NUM_UINT(i int) antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllBYTES() []antlr.TerminalNode
	BYTES(i int) antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllRBRACE() []antlr.TerminalNode
	RBRACE(i int) antlr.TerminalNode

	// IsConditionExpressionContext differentiates from other interfaces.
	IsConditionExpressionContext()
}

type ConditionExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionExpressionContext() *ConditionExpressionContext {
	var p = new(ConditionExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_conditionExpression
	return p
}

func InitEmptyConditionExpressionContext(p *ConditionExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_conditionExpression
}

func (*ConditionExpressionContext) IsConditionExpressionContext() {}

func NewConditionExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionExpressionContext {
	var p = new(ConditionExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_conditionExpression

	return p
}

func (s *ConditionExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionExpressionContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserIDENTIFIER)
}

func (s *ConditionExpressionContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserIDENTIFIER, i)
}

func (s *ConditionExpressionContext) AllEQUALS() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserEQUALS)
}

func (s *ConditionExpressionContext) EQUALS(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserEQUALS, i)
}

func (s *ConditionExpressionContext) AllNOT_EQUALS() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNOT_EQUALS)
}

func (s *ConditionExpressionContext) NOT_EQUALS(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNOT_EQUALS, i)
}

func (s *ConditionExpressionContext) AllIN() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserIN)
}

func (s *ConditionExpressionContext) IN(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserIN, i)
}

func (s *ConditionExpressionContext) AllLESS() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserLESS)
}

func (s *ConditionExpressionContext) LESS(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserLESS, i)
}

func (s *ConditionExpressionContext) AllLESS_EQUALS() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserLESS_EQUALS)
}

func (s *ConditionExpressionContext) LESS_EQUALS(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserLESS_EQUALS, i)
}

func (s *ConditionExpressionContext) AllGREATER_EQUALS() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserGREATER_EQUALS)
}

func (s *ConditionExpressionContext) GREATER_EQUALS(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserGREATER_EQUALS, i)
}

func (s *ConditionExpressionContext) AllGREATER() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserGREATER)
}

func (s *ConditionExpressionContext) GREATER(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserGREATER, i)
}

func (s *ConditionExpressionContext) AllLOGICAL_AND() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserLOGICAL_AND)
}

func (s *ConditionExpressionContext) LOGICAL_AND(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserLOGICAL_AND, i)
}

func (s *ConditionExpressionContext) AllLOGICAL_OR() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserLOGICAL_OR)
}

func (s *ConditionExpressionContext) LOGICAL_OR(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserLOGICAL_OR, i)
}

func (s *ConditionExpressionContext) AllLBRACKET() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserLBRACKET)
}

func (s *ConditionExpressionContext) LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserLBRACKET, i)
}

func (s *ConditionExpressionContext) AllRPRACKET() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserRPRACKET)
}

func (s *ConditionExpressionContext) RPRACKET(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserRPRACKET, i)
}

func (s *ConditionExpressionContext) AllLBRACE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserLBRACE)
}

func (s *ConditionExpressionContext) LBRACE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserLBRACE, i)
}

func (s *ConditionExpressionContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserLPAREN)
}

func (s *ConditionExpressionContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserLPAREN, i)
}

func (s *ConditionExpressionContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserRPAREN)
}

func (s *ConditionExpressionContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserRPAREN, i)
}

func (s *ConditionExpressionContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserDOT)
}

func (s *ConditionExpressionContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserDOT, i)
}

func (s *ConditionExpressionContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserMINUS)
}

func (s *ConditionExpressionContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserMINUS, i)
}

func (s *ConditionExpressionContext) AllEXCLAM() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserEXCLAM)
}

func (s *ConditionExpressionContext) EXCLAM(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserEXCLAM, i)
}

func (s *ConditionExpressionContext) AllQUESTIONMARK() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserQUESTIONMARK)
}

func (s *ConditionExpressionContext) QUESTIONMARK(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserQUESTIONMARK, i)
}

func (s *ConditionExpressionContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserPLUS)
}

func (s *ConditionExpressionContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserPLUS, i)
}

func (s *ConditionExpressionContext) AllSTAR() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserSTAR)
}

func (s *ConditionExpressionContext) STAR(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserSTAR, i)
}

func (s *ConditionExpressionContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserSLASH)
}

func (s *ConditionExpressionContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserSLASH, i)
}

func (s *ConditionExpressionContext) AllPERCENT() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserPERCENT)
}

func (s *ConditionExpressionContext) PERCENT(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserPERCENT, i)
}

func (s *ConditionExpressionContext) AllCEL_TRUE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserCEL_TRUE)
}

func (s *ConditionExpressionContext) CEL_TRUE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserCEL_TRUE, i)
}

func (s *ConditionExpressionContext) AllCEL_FALSE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserCEL_FALSE)
}

func (s *ConditionExpressionContext) CEL_FALSE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserCEL_FALSE, i)
}

func (s *ConditionExpressionContext) AllNUL() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNUL)
}

func (s *ConditionExpressionContext) NUL(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNUL, i)
}

func (s *ConditionExpressionContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserWHITESPACE)
}

func (s *ConditionExpressionContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWHITESPACE, i)
}

func (s *ConditionExpressionContext) AllCEL_COMMENT() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserCEL_COMMENT)
}

func (s *ConditionExpressionContext) CEL_COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserCEL_COMMENT, i)
}

func (s *ConditionExpressionContext) AllNUM_FLOAT() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNUM_FLOAT)
}

func (s *ConditionExpressionContext) NUM_FLOAT(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNUM_FLOAT, i)
}

func (s *ConditionExpressionContext) AllNUM_INT() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNUM_INT)
}

func (s *ConditionExpressionContext) NUM_INT(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNUM_INT, i)
}

func (s *ConditionExpressionContext) AllNUM_UINT() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNUM_UINT)
}

func (s *ConditionExpressionContext) NUM_UINT(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNUM_UINT, i)
}

func (s *ConditionExpressionContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserSTRING)
}

func (s *ConditionExpressionContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserSTRING, i)
}

func (s *ConditionExpressionContext) AllBYTES() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserBYTES)
}

func (s *ConditionExpressionContext) BYTES(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserBYTES, i)
}

func (s *ConditionExpressionContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNEWLINE)
}

func (s *ConditionExpressionContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINE, i)
}

func (s *ConditionExpressionContext) AllRBRACE() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserRBRACE)
}

func (s *ConditionExpressionContext) RBRACE(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserRBRACE, i)
}

func (s *ConditionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConditionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.EnterConditionExpression(s)
	}
}

func (s *ConditionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAParserListener); ok {
		listenerT.ExitConditionExpression(s)
	}
}




func (p *OpenFGAParser) ConditionExpression() (localctx IConditionExpressionContext) {
	localctx = NewConditionExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, OpenFGAParserRULE_conditionExpression)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(376)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(374)
					_la = p.GetTokenStream().LA(1)

					if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 18014364082636728) != 0)) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}


			case 2:
				{
					p.SetState(375)
					_la = p.GetTokenStream().LA(1)

					if _la <= 0 || _la == OpenFGAParserRBRACE  {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(380)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext())
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


