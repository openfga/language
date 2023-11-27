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
    "", "'and'", "'or'", "'but not'", "'from'", "'model'", "'schema'", "'1.1'", 
    "'type'", "'condition'", "", "", "'relations'", "'define'", "'with'", 
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'{'", "", "", 
    "", "'=='", "'!='", "'in'", "", "'<='", "'>='", "", "'&&'", "'||'", 
    "", "", "", "", "", "'-'", "'!'", "'?'", "'+'", "", "'/'", "'%'", "'true'", 
    "'false'", "'null'", "", "", "", "", "", "", "", "", "", "'}'",
  }
  staticData.SymbolicNames = []string{
    "", "AND", "OR", "BUT_NOT", "FROM", "MODEL", "SCHEMA", "SCHEMA_VERSION", 
    "TYPE", "CONDITION", "CONDITION_PARAM_CONTAINER", "CONDITION_PARAM_TYPE", 
    "RELATIONS", "DEFINE", "KEYWORD_WITH", "IDENTIFIER", "WHITESPACE", "NEWLINE", 
    "DOT", "STAR", "HASH", "COLON", "COMMA", "LPAREN", "RPAREN", "LESS", 
    "GREATER", "LBRACKET", "RPRACKET", "OPEN_CEL", "CEL_HASH", "CEL_COLON", 
    "CEL_COMMA", "EQUALS", "NOT_EQUALS", "IN", "CEL_LESS", "LESS_EQUALS", 
    "GREATER_EQUALS", "CEL_GREATER", "LOGICAL_AND", "LOGICAL_OR", "CEL_LBRACKET", 
    "CEL_RPRACKET", "CEL_LPAREN", "CEL_RPAREN", "CEL_DOT", "MINUS", "EXCLAM", 
    "QUESTIONMARK", "PLUS", "CEL_STAR", "SLASH", "PERCENT", "CEL_TRUE", 
    "CEL_FALSE", "NUL", "CEL_COMMENT", "NUM_FLOAT", "NUM_INT", "NUM_UINT", 
    "STRING", "BYTES", "CEL_IDENTIFIER", "CEL_WHITESPACE", "CEL_NEWLINE", 
    "CLOSE_CEL",
  }
  staticData.RuleNames = []string{
    "main", "modelHeader", "typeDefs", "typeDef", "relationDeclaration", 
    "relationName", "relationDef", "relationDefPartials", "relationDefGrouping", 
    "relationDefDirectAssignment", "relationDefRewrite", "relationDefTypeRestriction", 
    "relationDefTypeRestrictionBase", "conditions", "condition", "conditionName", 
    "conditionParameter", "parameterName", "parameterType", "multiLineComment", 
    "conditionExpression",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 66, 291, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1, 
	0, 3, 0, 44, 8, 0, 1, 0, 3, 0, 47, 8, 0, 1, 0, 1, 0, 3, 0, 51, 8, 0, 1, 
	0, 1, 0, 3, 0, 55, 8, 0, 1, 0, 1, 0, 3, 0, 59, 8, 0, 1, 0, 1, 0, 1, 1, 
	1, 1, 1, 1, 3, 1, 66, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 74, 
	8, 1, 1, 2, 5, 2, 77, 8, 2, 10, 2, 12, 2, 80, 9, 2, 1, 3, 1, 3, 3, 3, 84, 
	8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 4, 3, 93, 8, 3, 11, 3, 
	12, 3, 94, 3, 3, 97, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 104, 8, 
	4, 1, 4, 1, 4, 3, 4, 108, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 3, 
	6, 116, 8, 6, 1, 6, 3, 6, 119, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 4, 7, 125, 
	8, 7, 11, 7, 12, 7, 126, 1, 7, 1, 7, 1, 7, 1, 7, 4, 7, 133, 8, 7, 11, 7, 
	12, 7, 134, 1, 7, 1, 7, 1, 7, 1, 7, 4, 7, 141, 8, 7, 11, 7, 12, 7, 142, 
	3, 7, 145, 8, 7, 1, 8, 1, 8, 1, 9, 1, 9, 3, 9, 151, 8, 9, 1, 9, 1, 9, 3, 
	9, 155, 8, 9, 1, 9, 1, 9, 3, 9, 159, 8, 9, 1, 9, 1, 9, 3, 9, 163, 8, 9, 
	5, 9, 165, 8, 9, 10, 9, 12, 9, 168, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 
	10, 1, 10, 1, 10, 3, 10, 177, 8, 10, 1, 11, 3, 11, 180, 8, 11, 1, 11, 1, 
	11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 189, 8, 11, 1, 11, 3, 11, 
	192, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 199, 8, 12, 1, 13, 
	5, 13, 202, 8, 13, 10, 13, 12, 13, 205, 9, 13, 1, 14, 1, 14, 3, 14, 209, 
	8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 216, 8, 14, 1, 14, 1, 
	14, 3, 14, 220, 8, 14, 1, 14, 1, 14, 3, 14, 224, 8, 14, 1, 14, 1, 14, 3, 
	14, 228, 8, 14, 1, 14, 1, 14, 3, 14, 232, 8, 14, 5, 14, 234, 8, 14, 10, 
	14, 12, 14, 237, 9, 14, 1, 14, 3, 14, 240, 8, 14, 1, 14, 1, 14, 3, 14, 
	244, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 3, 16, 253, 
	8, 16, 1, 16, 1, 16, 3, 16, 257, 8, 16, 1, 16, 1, 16, 3, 16, 261, 8, 16, 
	1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 272, 
	8, 18, 1, 19, 1, 19, 5, 19, 276, 8, 19, 10, 19, 12, 19, 279, 9, 19, 1, 
	19, 1, 19, 3, 19, 283, 8, 19, 1, 20, 5, 20, 286, 8, 20, 10, 20, 12, 20, 
	289, 9, 20, 1, 20, 0, 0, 21, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 
	24, 26, 28, 30, 32, 34, 36, 38, 40, 0, 2, 1, 0, 17, 17, 1, 0, 66, 66, 317, 
	0, 43, 1, 0, 0, 0, 2, 65, 1, 0, 0, 0, 4, 78, 1, 0, 0, 0, 6, 83, 1, 0, 0, 
	0, 8, 98, 1, 0, 0, 0, 10, 111, 1, 0, 0, 0, 12, 115, 1, 0, 0, 0, 14, 144, 
	1, 0, 0, 0, 16, 146, 1, 0, 0, 0, 18, 148, 1, 0, 0, 0, 20, 171, 1, 0, 0, 
	0, 22, 179, 1, 0, 0, 0, 24, 193, 1, 0, 0, 0, 26, 203, 1, 0, 0, 0, 28, 208, 
	1, 0, 0, 0, 30, 249, 1, 0, 0, 0, 32, 252, 1, 0, 0, 0, 34, 264, 1, 0, 0, 
	0, 36, 271, 1, 0, 0, 0, 38, 273, 1, 0, 0, 0, 40, 287, 1, 0, 0, 0, 42, 44, 
	5, 16, 0, 0, 43, 42, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 46, 1, 0, 0, 0, 
	45, 47, 5, 17, 0, 0, 46, 45, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 48, 1, 
	0, 0, 0, 48, 50, 3, 2, 1, 0, 49, 51, 5, 17, 0, 0, 50, 49, 1, 0, 0, 0, 50, 
	51, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 54, 3, 4, 2, 0, 53, 55, 5, 17, 
	0, 0, 54, 53, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 58, 
	3, 26, 13, 0, 57, 59, 5, 17, 0, 0, 58, 57, 1, 0, 0, 0, 58, 59, 1, 0, 0, 
	0, 59, 60, 1, 0, 0, 0, 60, 61, 5, 0, 0, 1, 61, 1, 1, 0, 0, 0, 62, 63, 3, 
	38, 19, 0, 63, 64, 5, 17, 0, 0, 64, 66, 1, 0, 0, 0, 65, 62, 1, 0, 0, 0, 
	65, 66, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 5, 5, 0, 0, 68, 69, 5, 
	17, 0, 0, 69, 70, 5, 6, 0, 0, 70, 71, 5, 16, 0, 0, 71, 73, 5, 7, 0, 0, 
	72, 74, 5, 16, 0, 0, 73, 72, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 3, 1, 
	0, 0, 0, 75, 77, 3, 6, 3, 0, 76, 75, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 
	76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 5, 1, 0, 0, 0, 80, 78, 1, 0, 0, 
	0, 81, 82, 5, 17, 0, 0, 82, 84, 3, 38, 19, 0, 83, 81, 1, 0, 0, 0, 83, 84, 
	1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 5, 17, 0, 0, 86, 87, 5, 8, 0, 0, 
	87, 88, 5, 16, 0, 0, 88, 96, 5, 15, 0, 0, 89, 90, 5, 17, 0, 0, 90, 92, 
	5, 12, 0, 0, 91, 93, 3, 8, 4, 0, 92, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 
	94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 97, 1, 0, 0, 0, 96, 89, 1, 
	0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 7, 1, 0, 0, 0, 98, 99, 5, 17, 0, 0, 99, 
	100, 5, 13, 0, 0, 100, 101, 5, 16, 0, 0, 101, 103, 3, 10, 5, 0, 102, 104, 
	5, 16, 0, 0, 103, 102, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 1, 0, 
	0, 0, 105, 107, 5, 21, 0, 0, 106, 108, 5, 16, 0, 0, 107, 106, 1, 0, 0, 
	0, 107, 108, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 110, 3, 12, 6, 0, 110, 
	9, 1, 0, 0, 0, 111, 112, 5, 15, 0, 0, 112, 11, 1, 0, 0, 0, 113, 116, 3, 
	18, 9, 0, 114, 116, 3, 16, 8, 0, 115, 113, 1, 0, 0, 0, 115, 114, 1, 0, 
	0, 0, 116, 118, 1, 0, 0, 0, 117, 119, 3, 14, 7, 0, 118, 117, 1, 0, 0, 0, 
	118, 119, 1, 0, 0, 0, 119, 13, 1, 0, 0, 0, 120, 121, 5, 16, 0, 0, 121, 
	122, 5, 2, 0, 0, 122, 123, 5, 16, 0, 0, 123, 125, 3, 16, 8, 0, 124, 120, 
	1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0, 
	0, 0, 127, 145, 1, 0, 0, 0, 128, 129, 5, 16, 0, 0, 129, 130, 5, 1, 0, 0, 
	130, 131, 5, 16, 0, 0, 131, 133, 3, 16, 8, 0, 132, 128, 1, 0, 0, 0, 133, 
	134, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 145, 
	1, 0, 0, 0, 136, 137, 5, 16, 0, 0, 137, 138, 5, 3, 0, 0, 138, 139, 5, 16, 
	0, 0, 139, 141, 3, 16, 8, 0, 140, 136, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 
	142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 145, 1, 0, 0, 0, 144, 
	124, 1, 0, 0, 0, 144, 132, 1, 0, 0, 0, 144, 140, 1, 0, 0, 0, 145, 15, 1, 
	0, 0, 0, 146, 147, 3, 20, 10, 0, 147, 17, 1, 0, 0, 0, 148, 150, 5, 27, 
	0, 0, 149, 151, 5, 16, 0, 0, 150, 149, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 
	151, 152, 1, 0, 0, 0, 152, 154, 3, 22, 11, 0, 153, 155, 5, 16, 0, 0, 154, 
	153, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 166, 1, 0, 0, 0, 156, 158, 
	5, 22, 0, 0, 157, 159, 5, 16, 0, 0, 158, 157, 1, 0, 0, 0, 158, 159, 1, 
	0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 162, 3, 22, 11, 0, 161, 163, 5, 16, 
	0, 0, 162, 161, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 165, 1, 0, 0, 0, 
	164, 156, 1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 
	167, 1, 0, 0, 0, 167, 169, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 170, 
	5, 28, 0, 0, 170, 19, 1, 0, 0, 0, 171, 176, 5, 15, 0, 0, 172, 173, 5, 16, 
	0, 0, 173, 174, 5, 4, 0, 0, 174, 175, 5, 16, 0, 0, 175, 177, 5, 15, 0, 
	0, 176, 172, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 21, 1, 0, 0, 0, 178, 
	180, 5, 17, 0, 0, 179, 178, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 188, 
	1, 0, 0, 0, 181, 189, 3, 24, 12, 0, 182, 183, 3, 24, 12, 0, 183, 184, 5, 
	16, 0, 0, 184, 185, 5, 14, 0, 0, 185, 186, 5, 16, 0, 0, 186, 187, 3, 30, 
	15, 0, 187, 189, 1, 0, 0, 0, 188, 181, 1, 0, 0, 0, 188, 182, 1, 0, 0, 0, 
	189, 191, 1, 0, 0, 0, 190, 192, 5, 17, 0, 0, 191, 190, 1, 0, 0, 0, 191, 
	192, 1, 0, 0, 0, 192, 23, 1, 0, 0, 0, 193, 198, 5, 15, 0, 0, 194, 195, 
	5, 21, 0, 0, 195, 199, 5, 19, 0, 0, 196, 197, 5, 20, 0, 0, 197, 199, 5, 
	15, 0, 0, 198, 194, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 
	0, 199, 25, 1, 0, 0, 0, 200, 202, 3, 28, 14, 0, 201, 200, 1, 0, 0, 0, 202, 
	205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 27, 1, 
	0, 0, 0, 205, 203, 1, 0, 0, 0, 206, 207, 5, 17, 0, 0, 207, 209, 3, 38, 
	19, 0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 
	210, 211, 5, 17, 0, 0, 211, 212, 5, 9, 0, 0, 212, 213, 5, 16, 0, 0, 213, 
	215, 3, 30, 15, 0, 214, 216, 5, 16, 0, 0, 215, 214, 1, 0, 0, 0, 215, 216, 
	1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 219, 5, 23, 0, 0, 218, 220, 5, 16, 
	0, 0, 219, 218, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 
	221, 223, 3, 32, 16, 0, 222, 224, 5, 16, 0, 0, 223, 222, 1, 0, 0, 0, 223, 
	224, 1, 0, 0, 0, 224, 235, 1, 0, 0, 0, 225, 227, 5, 22, 0, 0, 226, 228, 
	5, 16, 0, 0, 227, 226, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 229, 1, 0, 
	0, 0, 229, 231, 3, 32, 16, 0, 230, 232, 5, 16, 0, 0, 231, 230, 1, 0, 0, 
	0, 231, 232, 1, 0, 0, 0, 232, 234, 1, 0, 0, 0, 233, 225, 1, 0, 0, 0, 234, 
	237, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 239, 
	1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 238, 240, 5, 17, 0, 0, 239, 238, 1, 0, 
	0, 0, 239, 240, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 243, 5, 24, 0, 0, 
	242, 244, 5, 16, 0, 0, 243, 242, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 
	245, 1, 0, 0, 0, 245, 246, 5, 29, 0, 0, 246, 247, 3, 40, 20, 0, 247, 248, 
	5, 66, 0, 0, 248, 29, 1, 0, 0, 0, 249, 250, 5, 15, 0, 0, 250, 31, 1, 0, 
	0, 0, 251, 253, 5, 17, 0, 0, 252, 251, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 
	253, 254, 1, 0, 0, 0, 254, 256, 3, 34, 17, 0, 255, 257, 5, 16, 0, 0, 256, 
	255, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 260, 
	5, 21, 0, 0, 259, 261, 5, 16, 0, 0, 260, 259, 1, 0, 0, 0, 260, 261, 1, 
	0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 263, 3, 36, 18, 0, 263, 33, 1, 0, 0, 
	0, 264, 265, 5, 15, 0, 0, 265, 35, 1, 0, 0, 0, 266, 272, 5, 11, 0, 0, 267, 
	268, 5, 10, 0, 0, 268, 269, 5, 25, 0, 0, 269, 270, 5, 11, 0, 0, 270, 272, 
	5, 26, 0, 0, 271, 266, 1, 0, 0, 0, 271, 267, 1, 0, 0, 0, 272, 37, 1, 0, 
	0, 0, 273, 277, 5, 20, 0, 0, 274, 276, 8, 0, 0, 0, 275, 274, 1, 0, 0, 0, 
	276, 279, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 
	282, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 280, 281, 5, 17, 0, 0, 281, 283, 
	3, 38, 19, 0, 282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 39, 1, 0, 
	0, 0, 284, 286, 8, 1, 0, 0, 285, 284, 1, 0, 0, 0, 286, 289, 1, 0, 0, 0, 
	287, 285, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 41, 1, 0, 0, 0, 289, 287, 
	1, 0, 0, 0, 46, 43, 46, 50, 54, 58, 65, 73, 78, 83, 94, 96, 103, 107, 115, 
	118, 126, 134, 142, 144, 150, 154, 158, 162, 166, 176, 179, 188, 191, 198, 
	203, 208, 215, 219, 223, 227, 231, 235, 239, 243, 252, 256, 260, 271, 277, 
	282, 287,
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
	OpenFGAParserAND = 1
	OpenFGAParserOR = 2
	OpenFGAParserBUT_NOT = 3
	OpenFGAParserFROM = 4
	OpenFGAParserMODEL = 5
	OpenFGAParserSCHEMA = 6
	OpenFGAParserSCHEMA_VERSION = 7
	OpenFGAParserTYPE = 8
	OpenFGAParserCONDITION = 9
	OpenFGAParserCONDITION_PARAM_CONTAINER = 10
	OpenFGAParserCONDITION_PARAM_TYPE = 11
	OpenFGAParserRELATIONS = 12
	OpenFGAParserDEFINE = 13
	OpenFGAParserKEYWORD_WITH = 14
	OpenFGAParserIDENTIFIER = 15
	OpenFGAParserWHITESPACE = 16
	OpenFGAParserNEWLINE = 17
	OpenFGAParserDOT = 18
	OpenFGAParserSTAR = 19
	OpenFGAParserHASH = 20
	OpenFGAParserCOLON = 21
	OpenFGAParserCOMMA = 22
	OpenFGAParserLPAREN = 23
	OpenFGAParserRPAREN = 24
	OpenFGAParserLESS = 25
	OpenFGAParserGREATER = 26
	OpenFGAParserLBRACKET = 27
	OpenFGAParserRPRACKET = 28
	OpenFGAParserOPEN_CEL = 29
	OpenFGAParserCEL_HASH = 30
	OpenFGAParserCEL_COLON = 31
	OpenFGAParserCEL_COMMA = 32
	OpenFGAParserEQUALS = 33
	OpenFGAParserNOT_EQUALS = 34
	OpenFGAParserIN = 35
	OpenFGAParserCEL_LESS = 36
	OpenFGAParserLESS_EQUALS = 37
	OpenFGAParserGREATER_EQUALS = 38
	OpenFGAParserCEL_GREATER = 39
	OpenFGAParserLOGICAL_AND = 40
	OpenFGAParserLOGICAL_OR = 41
	OpenFGAParserCEL_LBRACKET = 42
	OpenFGAParserCEL_RPRACKET = 43
	OpenFGAParserCEL_LPAREN = 44
	OpenFGAParserCEL_RPAREN = 45
	OpenFGAParserCEL_DOT = 46
	OpenFGAParserMINUS = 47
	OpenFGAParserEXCLAM = 48
	OpenFGAParserQUESTIONMARK = 49
	OpenFGAParserPLUS = 50
	OpenFGAParserCEL_STAR = 51
	OpenFGAParserSLASH = 52
	OpenFGAParserPERCENT = 53
	OpenFGAParserCEL_TRUE = 54
	OpenFGAParserCEL_FALSE = 55
	OpenFGAParserNUL = 56
	OpenFGAParserCEL_COMMENT = 57
	OpenFGAParserNUM_FLOAT = 58
	OpenFGAParserNUM_INT = 59
	OpenFGAParserNUM_UINT = 60
	OpenFGAParserSTRING = 61
	OpenFGAParserBYTES = 62
	OpenFGAParserCEL_IDENTIFIER = 63
	OpenFGAParserCEL_WHITESPACE = 64
	OpenFGAParserCEL_NEWLINE = 65
	OpenFGAParserCLOSE_CEL = 66
)

// OpenFGAParser rules.
const (
	OpenFGAParserRULE_main = 0
	OpenFGAParserRULE_modelHeader = 1
	OpenFGAParserRULE_typeDefs = 2
	OpenFGAParserRULE_typeDef = 3
	OpenFGAParserRULE_relationDeclaration = 4
	OpenFGAParserRULE_relationName = 5
	OpenFGAParserRULE_relationDef = 6
	OpenFGAParserRULE_relationDefPartials = 7
	OpenFGAParserRULE_relationDefGrouping = 8
	OpenFGAParserRULE_relationDefDirectAssignment = 9
	OpenFGAParserRULE_relationDefRewrite = 10
	OpenFGAParserRULE_relationDefTypeRestriction = 11
	OpenFGAParserRULE_relationDefTypeRestrictionBase = 12
	OpenFGAParserRULE_conditions = 13
	OpenFGAParserRULE_condition = 14
	OpenFGAParserRULE_conditionName = 15
	OpenFGAParserRULE_conditionParameter = 16
	OpenFGAParserRULE_parameterName = 17
	OpenFGAParserRULE_parameterType = 18
	OpenFGAParserRULE_multiLineComment = 19
	OpenFGAParserRULE_conditionExpression = 20
)

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ModelHeader() IModelHeaderContext
	TypeDefs() ITypeDefsContext
	Conditions() IConditionsContext
	EOF() antlr.TerminalNode
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
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(42)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(45)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(48)
		p.ModelHeader()
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(49)
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
		p.SetState(52)
		p.TypeDefs()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(53)
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
		p.SetState(56)
		p.Conditions()
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(57)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(60)
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserHASH {
		{
			p.SetState(62)
			p.MultiLineComment()
		}
		{
			p.SetState(63)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(67)
		p.Match(OpenFGAParserMODEL)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Match(OpenFGAParserNEWLINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Match(OpenFGAParserSCHEMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Match(OpenFGAParserWHITESPACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(71)

		var _m = p.Match(OpenFGAParserSCHEMA_VERSION)

		localctx.(*ModelHeaderContext).schemaVersion = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(72)
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
	p.EnterRule(localctx, 4, OpenFGAParserRULE_typeDefs)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(75)
				p.TypeDef()
			}


		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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
	WHITESPACE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	MultiLineComment() IMultiLineCommentContext
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

func (s *TypeDefContext) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWHITESPACE, 0)
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
	p.EnterRule(localctx, 6, OpenFGAParserRULE_typeDef)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(81)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(82)
			p.MultiLineComment()
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}
	{
		p.SetState(85)
		p.Match(OpenFGAParserNEWLINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Match(OpenFGAParserTYPE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Match(OpenFGAParserWHITESPACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(88)

		var _m = p.Match(OpenFGAParserIDENTIFIER)

		localctx.(*TypeDefContext).typeName = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(89)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Match(OpenFGAParserRELATIONS)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
					{
						p.SetState(91)
						p.RelationDeclaration()
					}




			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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
	NEWLINE() antlr.TerminalNode
	DEFINE() antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	RelationName() IRelationNameContext
	COLON() antlr.TerminalNode
	RelationDef() IRelationDefContext

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

func (s *RelationDeclarationContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINE, 0)
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
	p.EnterRule(localctx, 8, OpenFGAParserRULE_relationDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(OpenFGAParserNEWLINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(99)
		p.Match(OpenFGAParserDEFINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(OpenFGAParserWHITESPACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(101)
		p.RelationName()
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(102)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(105)
		p.Match(OpenFGAParserCOLON)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(106)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(109)
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
	p.EnterRule(localctx, 10, OpenFGAParserRULE_relationName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
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
	p.EnterRule(localctx, 12, OpenFGAParserRULE_relationDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case OpenFGAParserLBRACKET:
		{
			p.SetState(113)
			p.RelationDefDirectAssignment()
		}


	case OpenFGAParserIDENTIFIER:
		{
			p.SetState(114)
			p.RelationDefGrouping()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(117)
			p.RelationDefPartials()
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
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode
	AllBUT_NOT() []antlr.TerminalNode
	BUT_NOT(i int) antlr.TerminalNode

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

func (s *RelationDefPartialsContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserAND)
}

func (s *RelationDefPartialsContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserAND, i)
}

func (s *RelationDefPartialsContext) AllBUT_NOT() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserBUT_NOT)
}

func (s *RelationDefPartialsContext) BUT_NOT(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserBUT_NOT, i)
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
	var _la int

	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == OpenFGAParserWHITESPACE {
			{
				p.SetState(120)
				p.Match(OpenFGAParserWHITESPACE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(121)
				p.Match(OpenFGAParserOR)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(122)
				p.Match(OpenFGAParserWHITESPACE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(123)
				p.RelationDefGrouping()
			}


			p.SetState(126)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == OpenFGAParserWHITESPACE {
			{
				p.SetState(128)
				p.Match(OpenFGAParserWHITESPACE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(129)
				p.Match(OpenFGAParserAND)
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
				p.RelationDefGrouping()
			}


			p.SetState(134)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == OpenFGAParserWHITESPACE {
			{
				p.SetState(136)
				p.Match(OpenFGAParserWHITESPACE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(137)
				p.Match(OpenFGAParserBUT_NOT)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(138)
				p.Match(OpenFGAParserWHITESPACE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(139)
				p.RelationDefGrouping()
			}


			p.SetState(142)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 16, OpenFGAParserRULE_relationDefGrouping)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
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
	p.EnterRule(localctx, 18, OpenFGAParserRULE_relationDefDirectAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(OpenFGAParserLBRACKET)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(149)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(152)
		p.RelationDefTypeRestriction()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(153)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == OpenFGAParserCOMMA {
		{
			p.SetState(156)
			p.Match(OpenFGAParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == OpenFGAParserWHITESPACE {
			{
				p.SetState(157)
				p.Match(OpenFGAParserWHITESPACE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}
		{
			p.SetState(160)
			p.RelationDefTypeRestriction()
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == OpenFGAParserWHITESPACE {
			{
				p.SetState(161)
				p.Match(OpenFGAParserWHITESPACE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}


		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(169)
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
	p.EnterRule(localctx, 20, OpenFGAParserRULE_relationDefRewrite)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)

		var _m = p.Match(OpenFGAParserIDENTIFIER)

		localctx.(*RelationDefRewriteContext).rewriteComputedusersetName = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(172)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Match(OpenFGAParserFROM)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(175)

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
	p.EnterRule(localctx, 22, OpenFGAParserRULE_relationDefTypeRestriction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(178)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(181)
			p.RelationDefTypeRestrictionBase()
		}


	case 2:
		{
			p.SetState(182)
			p.RelationDefTypeRestrictionBase()
		}
		{
			p.SetState(183)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(184)
			p.Match(OpenFGAParserKEYWORD_WITH)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(186)
			p.ConditionName()
		}


	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(190)
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
	p.EnterRule(localctx, 24, OpenFGAParserRULE_relationDefTypeRestrictionBase)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)

		var _m = p.Match(OpenFGAParserIDENTIFIER)

		localctx.(*RelationDefTypeRestrictionBaseContext).relationDefTypeRestrictionType = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case OpenFGAParserCOLON:
		{
			p.SetState(194)
			p.Match(OpenFGAParserCOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(195)

			var _m = p.Match(OpenFGAParserSTAR)

			localctx.(*RelationDefTypeRestrictionBaseContext).relationDefTypeRestrictionWildcard = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	case OpenFGAParserHASH:
		{
			p.SetState(196)
			p.Match(OpenFGAParserHASH)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(197)

			var _m = p.Match(OpenFGAParserIDENTIFIER)

			localctx.(*RelationDefTypeRestrictionBaseContext).relationDefTypeRestrictionRelation = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	case OpenFGAParserWHITESPACE, OpenFGAParserNEWLINE, OpenFGAParserCOMMA, OpenFGAParserRPRACKET:



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
	p.EnterRule(localctx, 26, OpenFGAParserRULE_conditions)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(200)
				p.Condition()
			}


		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
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
	OPEN_CEL() antlr.TerminalNode
	ConditionExpression() IConditionExpressionContext
	CLOSE_CEL() antlr.TerminalNode
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

func (s *ConditionContext) OPEN_CEL() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserOPEN_CEL, 0)
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

func (s *ConditionContext) CLOSE_CEL() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserCLOSE_CEL, 0)
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
	p.EnterRule(localctx, 28, OpenFGAParserRULE_condition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(208)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(206)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(207)
			p.MultiLineComment()
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}
	{
		p.SetState(210)
		p.Match(OpenFGAParserNEWLINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Match(OpenFGAParserCONDITION)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(212)
		p.Match(OpenFGAParserWHITESPACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(213)
		p.ConditionName()
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(214)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(217)
		p.Match(OpenFGAParserLPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(218)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(221)
		p.ConditionParameter()
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(222)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == OpenFGAParserCOMMA {
		{
			p.SetState(225)
			p.Match(OpenFGAParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == OpenFGAParserWHITESPACE {
			{
				p.SetState(226)
				p.Match(OpenFGAParserWHITESPACE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}
		{
			p.SetState(229)
			p.ConditionParameter()
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


		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(238)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(241)
		p.Match(OpenFGAParserRPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
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
	{
		p.SetState(245)
		p.Match(OpenFGAParserOPEN_CEL)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(246)
		p.ConditionExpression()
	}
	{
		p.SetState(247)
		p.Match(OpenFGAParserCLOSE_CEL)
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
	p.EnterRule(localctx, 30, OpenFGAParserRULE_conditionName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
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
	p.EnterRule(localctx, 32, OpenFGAParserRULE_conditionParameter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserNEWLINE {
		{
			p.SetState(251)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(254)
		p.ParameterName()
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(255)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(258)
		p.Match(OpenFGAParserCOLON)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserWHITESPACE {
		{
			p.SetState(259)
			p.Match(OpenFGAParserWHITESPACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(262)
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
	p.EnterRule(localctx, 34, OpenFGAParserRULE_parameterName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
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
	p.EnterRule(localctx, 36, OpenFGAParserRULE_parameterType)
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case OpenFGAParserCONDITION_PARAM_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.Match(OpenFGAParserCONDITION_PARAM_TYPE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case OpenFGAParserCONDITION_PARAM_CONTAINER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
			p.Match(OpenFGAParserCONDITION_PARAM_CONTAINER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(268)
			p.Match(OpenFGAParserLESS)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(269)
			p.Match(OpenFGAParserCONDITION_PARAM_TYPE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(270)
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
	p.EnterRule(localctx, 38, OpenFGAParserRULE_multiLineComment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Match(OpenFGAParserHASH)
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


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & -131074) != 0) || ((int64((_la - 64)) & ^0x3f) == 0 && ((int64(1) << (_la - 64)) & 7) != 0) {
		{
			p.SetState(274)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == OpenFGAParserNEWLINE  {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}


		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(280)
			p.Match(OpenFGAParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(281)
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
	AllCLOSE_CEL() []antlr.TerminalNode
	CLOSE_CEL(i int) antlr.TerminalNode

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

func (s *ConditionExpressionContext) AllCLOSE_CEL() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserCLOSE_CEL)
}

func (s *ConditionExpressionContext) CLOSE_CEL(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserCLOSE_CEL, i)
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
	p.EnterRule(localctx, 40, OpenFGAParserRULE_conditionExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & -2) != 0) || _la == OpenFGAParserCEL_WHITESPACE || _la == OpenFGAParserCEL_NEWLINE {
		{
			p.SetState(284)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == OpenFGAParserCLOSE_CEL  {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}


		p.SetState(289)
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


