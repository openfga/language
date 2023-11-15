// Code generated from /app/OpenFGALexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
  	"sync"
	"unicode"
	"github.com/antlr4-go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type OpenFGALexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var OpenFGALexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  ChannelNames           []string
  ModeNames              []string
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func openfgalexerLexerInit() {
  staticData := &OpenFGALexerLexerStaticData
  staticData.ChannelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.ModeNames = []string{
    "DEFAULT_MODE", "CEL",
  }
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
    "AND", "OR", "BUT_NOT", "FROM", "MODEL", "SCHEMA", "SCHEMA_VERSION", 
    "TYPE", "CONDITION", "CONDITION_PARAM_CONTAINER", "CONDITION_PARAM_TYPE", 
    "RELATIONS", "DEFINE", "KEYWORD_WITH", "IDENTIFIER", "WHITESPACE", "NEWLINE", 
    "DOT", "STAR", "HASH", "COLON", "COMMA", "LPAREN", "RPAREN", "LESS", 
    "GREATER", "LBRACKET", "RPRACKET", "OPEN_CEL", "CEL_HASH", "CEL_COLON", 
    "CEL_COMMA", "EQUALS", "NOT_EQUALS", "IN", "CEL_LESS", "LESS_EQUALS", 
    "GREATER_EQUALS", "CEL_GREATER", "LOGICAL_AND", "LOGICAL_OR", "CEL_LBRACKET", 
    "CEL_RPRACKET", "CEL_LPAREN", "CEL_RPAREN", "CEL_DOT", "MINUS", "EXCLAM", 
    "QUESTIONMARK", "PLUS", "CEL_STAR", "SLASH", "PERCENT", "CEL_TRUE", 
    "CEL_FALSE", "NUL", "BACKSLASH", "LETTER", "DIGIT", "EXPONENT", "HEXDIGIT", 
    "RAW", "ESC_SEQ", "ESC_CHAR_SEQ", "ESC_OCT_SEQ", "ESC_BYTE_SEQ", "ESC_UNI_SEQ", 
    "CEL_COMMENT", "NUM_FLOAT", "NUM_INT", "NUM_UINT", "STRING", "BYTES", 
    "CEL_IDENTIFIER", "CEL_WHITESPACE", "CEL_NEWLINE", "CLOSE_CEL",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 66, 695, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 
	7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 
	7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 
	14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 
	2, 20, 7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 
	25, 7, 25, 2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 
	7, 30, 2, 31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 
	35, 2, 36, 7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 
	2, 41, 7, 41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 
	46, 7, 46, 2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 
	7, 51, 2, 52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 
	56, 2, 57, 7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 
	2, 62, 7, 62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 
	67, 7, 67, 2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 
	7, 72, 2, 73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 1, 0, 1, 
	0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 
	2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 
	4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 
	7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 
	8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 216, 8, 
	9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 
	1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 
	10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 
	1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 
	10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 267, 8, 10, 
	1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 
	12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 
	1, 13, 1, 14, 1, 14, 3, 14, 293, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 
	14, 299, 8, 14, 10, 14, 12, 14, 302, 9, 14, 1, 15, 4, 15, 305, 8, 15, 11, 
	15, 12, 15, 306, 1, 16, 3, 16, 310, 8, 16, 1, 16, 3, 16, 313, 8, 16, 1, 
	16, 1, 16, 3, 16, 317, 8, 16, 1, 16, 3, 16, 320, 8, 16, 1, 16, 3, 16, 323, 
	8, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 
	21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 
	1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 
	31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 
	1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 
	39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 
	1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 
	48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 
	1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 55, 1, 
	55, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 
	1, 59, 3, 59, 430, 8, 59, 1, 59, 4, 59, 433, 8, 59, 11, 59, 12, 59, 434, 
	1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 62, 3, 62, 445, 8, 
	62, 1, 63, 1, 63, 1, 63, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 65, 1, 65, 
	1, 65, 1, 65, 1, 65, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 
	66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 
	3, 66, 478, 8, 66, 1, 67, 1, 67, 1, 67, 1, 67, 5, 67, 484, 8, 67, 10, 67, 
	12, 67, 487, 9, 67, 1, 67, 1, 67, 1, 68, 4, 68, 492, 8, 68, 11, 68, 12, 
	68, 493, 1, 68, 1, 68, 4, 68, 498, 8, 68, 11, 68, 12, 68, 499, 1, 68, 3, 
	68, 503, 8, 68, 1, 68, 4, 68, 506, 8, 68, 11, 68, 12, 68, 507, 1, 68, 1, 
	68, 1, 68, 1, 68, 4, 68, 514, 8, 68, 11, 68, 12, 68, 515, 1, 68, 3, 68, 
	519, 8, 68, 3, 68, 521, 8, 68, 1, 69, 4, 69, 524, 8, 69, 11, 69, 12, 69, 
	525, 1, 69, 1, 69, 1, 69, 1, 69, 4, 69, 532, 8, 69, 11, 69, 12, 69, 533, 
	3, 69, 536, 8, 69, 1, 70, 4, 70, 539, 8, 70, 11, 70, 12, 70, 540, 1, 70, 
	1, 70, 1, 70, 1, 70, 1, 70, 1, 70, 4, 70, 549, 8, 70, 11, 70, 12, 70, 550, 
	1, 70, 1, 70, 3, 70, 555, 8, 70, 1, 71, 1, 71, 1, 71, 5, 71, 560, 8, 71, 
	10, 71, 12, 71, 563, 9, 71, 1, 71, 1, 71, 1, 71, 1, 71, 5, 71, 569, 8, 
	71, 10, 71, 12, 71, 572, 9, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 
	1, 71, 5, 71, 581, 8, 71, 10, 71, 12, 71, 584, 9, 71, 1, 71, 1, 71, 1, 
	71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 5, 71, 595, 8, 71, 10, 71, 
	12, 71, 598, 9, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 5, 71, 606, 
	8, 71, 10, 71, 12, 71, 609, 9, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 5, 
	71, 616, 8, 71, 10, 71, 12, 71, 619, 9, 71, 1, 71, 1, 71, 1, 71, 1, 71, 
	1, 71, 1, 71, 1, 71, 1, 71, 5, 71, 629, 8, 71, 10, 71, 12, 71, 632, 9, 
	71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 
	5, 71, 644, 8, 71, 10, 71, 12, 71, 647, 9, 71, 1, 71, 1, 71, 1, 71, 1, 
	71, 3, 71, 653, 8, 71, 1, 72, 1, 72, 1, 72, 1, 73, 1, 73, 3, 73, 660, 8, 
	73, 1, 73, 1, 73, 1, 73, 1, 73, 5, 73, 666, 8, 73, 10, 73, 12, 73, 669, 
	9, 73, 1, 74, 4, 74, 672, 8, 74, 11, 74, 12, 74, 673, 1, 75, 3, 75, 677, 
	8, 75, 1, 75, 3, 75, 680, 8, 75, 1, 75, 1, 75, 3, 75, 684, 8, 75, 1, 75, 
	3, 75, 687, 8, 75, 1, 75, 3, 75, 690, 8, 75, 1, 76, 1, 76, 1, 76, 1, 76, 
	4, 582, 596, 630, 645, 0, 77, 2, 1, 4, 2, 6, 3, 8, 4, 10, 5, 12, 6, 14, 
	7, 16, 8, 18, 9, 20, 10, 22, 11, 24, 12, 26, 13, 28, 14, 30, 15, 32, 16, 
	34, 17, 36, 18, 38, 19, 40, 20, 42, 21, 44, 22, 46, 23, 48, 24, 50, 25, 
	52, 26, 54, 27, 56, 28, 58, 29, 60, 30, 62, 31, 64, 32, 66, 33, 68, 34, 
	70, 35, 72, 36, 74, 37, 76, 38, 78, 39, 80, 40, 82, 41, 84, 42, 86, 43, 
	88, 44, 90, 45, 92, 46, 94, 47, 96, 48, 98, 49, 100, 50, 102, 51, 104, 
	52, 106, 53, 108, 54, 110, 55, 112, 56, 114, 0, 116, 0, 118, 0, 120, 0, 
	122, 0, 124, 0, 126, 0, 128, 0, 130, 0, 132, 0, 134, 0, 136, 57, 138, 58, 
	140, 59, 142, 60, 144, 61, 146, 62, 148, 63, 150, 64, 152, 65, 154, 66, 
	2, 0, 1, 16, 3, 0, 9, 9, 12, 12, 32, 32, 2, 0, 65, 90, 97, 122, 2, 0, 69, 
	69, 101, 101, 2, 0, 43, 43, 45, 45, 3, 0, 48, 57, 65, 70, 97, 102, 2, 0, 
	82, 82, 114, 114, 10, 0, 34, 34, 39, 39, 63, 63, 92, 92, 96, 98, 102, 102, 
	110, 110, 114, 114, 116, 116, 118, 118, 2, 0, 88, 88, 120, 120, 1, 0, 10, 
	10, 2, 0, 85, 85, 117, 117, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 4, 0, 
	10, 10, 13, 13, 39, 39, 92, 92, 1, 0, 92, 92, 3, 0, 10, 10, 13, 13, 34, 
	34, 3, 0, 10, 10, 13, 13, 39, 39, 2, 0, 66, 66, 98, 98, 752, 0, 2, 1, 0, 
	0, 0, 0, 4, 1, 0, 0, 0, 0, 6, 1, 0, 0, 0, 0, 8, 1, 0, 0, 0, 0, 10, 1, 0, 
	0, 0, 0, 12, 1, 0, 0, 0, 0, 14, 1, 0, 0, 0, 0, 16, 1, 0, 0, 0, 0, 18, 1, 
	0, 0, 0, 0, 20, 1, 0, 0, 0, 0, 22, 1, 0, 0, 0, 0, 24, 1, 0, 0, 0, 0, 26, 
	1, 0, 0, 0, 0, 28, 1, 0, 0, 0, 0, 30, 1, 0, 0, 0, 0, 32, 1, 0, 0, 0, 0, 
	34, 1, 0, 0, 0, 0, 36, 1, 0, 0, 0, 0, 38, 1, 0, 0, 0, 0, 40, 1, 0, 0, 0, 
	0, 42, 1, 0, 0, 0, 0, 44, 1, 0, 0, 0, 0, 46, 1, 0, 0, 0, 0, 48, 1, 0, 0, 
	0, 0, 50, 1, 0, 0, 0, 0, 52, 1, 0, 0, 0, 0, 54, 1, 0, 0, 0, 0, 56, 1, 0, 
	0, 0, 0, 58, 1, 0, 0, 0, 1, 60, 1, 0, 0, 0, 1, 62, 1, 0, 0, 0, 1, 64, 1, 
	0, 0, 0, 1, 66, 1, 0, 0, 0, 1, 68, 1, 0, 0, 0, 1, 70, 1, 0, 0, 0, 1, 72, 
	1, 0, 0, 0, 1, 74, 1, 0, 0, 0, 1, 76, 1, 0, 0, 0, 1, 78, 1, 0, 0, 0, 1, 
	80, 1, 0, 0, 0, 1, 82, 1, 0, 0, 0, 1, 84, 1, 0, 0, 0, 1, 86, 1, 0, 0, 0, 
	1, 88, 1, 0, 0, 0, 1, 90, 1, 0, 0, 0, 1, 92, 1, 0, 0, 0, 1, 94, 1, 0, 0, 
	0, 1, 96, 1, 0, 0, 0, 1, 98, 1, 0, 0, 0, 1, 100, 1, 0, 0, 0, 1, 102, 1, 
	0, 0, 0, 1, 104, 1, 0, 0, 0, 1, 106, 1, 0, 0, 0, 1, 108, 1, 0, 0, 0, 1, 
	110, 1, 0, 0, 0, 1, 112, 1, 0, 0, 0, 1, 136, 1, 0, 0, 0, 1, 138, 1, 0, 
	0, 0, 1, 140, 1, 0, 0, 0, 1, 142, 1, 0, 0, 0, 1, 144, 1, 0, 0, 0, 1, 146, 
	1, 0, 0, 0, 1, 148, 1, 0, 0, 0, 1, 150, 1, 0, 0, 0, 1, 152, 1, 0, 0, 0, 
	1, 154, 1, 0, 0, 0, 2, 156, 1, 0, 0, 0, 4, 160, 1, 0, 0, 0, 6, 163, 1, 
	0, 0, 0, 8, 171, 1, 0, 0, 0, 10, 176, 1, 0, 0, 0, 12, 182, 1, 0, 0, 0, 
	14, 189, 1, 0, 0, 0, 16, 193, 1, 0, 0, 0, 18, 198, 1, 0, 0, 0, 20, 215, 
	1, 0, 0, 0, 22, 266, 1, 0, 0, 0, 24, 268, 1, 0, 0, 0, 26, 278, 1, 0, 0, 
	0, 28, 285, 1, 0, 0, 0, 30, 292, 1, 0, 0, 0, 32, 304, 1, 0, 0, 0, 34, 309, 
	1, 0, 0, 0, 36, 324, 1, 0, 0, 0, 38, 326, 1, 0, 0, 0, 40, 328, 1, 0, 0, 
	0, 42, 330, 1, 0, 0, 0, 44, 332, 1, 0, 0, 0, 46, 334, 1, 0, 0, 0, 48, 336, 
	1, 0, 0, 0, 50, 338, 1, 0, 0, 0, 52, 340, 1, 0, 0, 0, 54, 342, 1, 0, 0, 
	0, 56, 344, 1, 0, 0, 0, 58, 346, 1, 0, 0, 0, 60, 350, 1, 0, 0, 0, 62, 352, 
	1, 0, 0, 0, 64, 354, 1, 0, 0, 0, 66, 356, 1, 0, 0, 0, 68, 359, 1, 0, 0, 
	0, 70, 362, 1, 0, 0, 0, 72, 365, 1, 0, 0, 0, 74, 367, 1, 0, 0, 0, 76, 370, 
	1, 0, 0, 0, 78, 373, 1, 0, 0, 0, 80, 375, 1, 0, 0, 0, 82, 378, 1, 0, 0, 
	0, 84, 381, 1, 0, 0, 0, 86, 383, 1, 0, 0, 0, 88, 385, 1, 0, 0, 0, 90, 387, 
	1, 0, 0, 0, 92, 389, 1, 0, 0, 0, 94, 391, 1, 0, 0, 0, 96, 393, 1, 0, 0, 
	0, 98, 395, 1, 0, 0, 0, 100, 397, 1, 0, 0, 0, 102, 399, 1, 0, 0, 0, 104, 
	401, 1, 0, 0, 0, 106, 403, 1, 0, 0, 0, 108, 405, 1, 0, 0, 0, 110, 410, 
	1, 0, 0, 0, 112, 416, 1, 0, 0, 0, 114, 421, 1, 0, 0, 0, 116, 423, 1, 0, 
	0, 0, 118, 425, 1, 0, 0, 0, 120, 427, 1, 0, 0, 0, 122, 436, 1, 0, 0, 0, 
	124, 438, 1, 0, 0, 0, 126, 444, 1, 0, 0, 0, 128, 446, 1, 0, 0, 0, 130, 
	449, 1, 0, 0, 0, 132, 454, 1, 0, 0, 0, 134, 477, 1, 0, 0, 0, 136, 479, 
	1, 0, 0, 0, 138, 520, 1, 0, 0, 0, 140, 535, 1, 0, 0, 0, 142, 554, 1, 0, 
	0, 0, 144, 652, 1, 0, 0, 0, 146, 654, 1, 0, 0, 0, 148, 659, 1, 0, 0, 0, 
	150, 671, 1, 0, 0, 0, 152, 676, 1, 0, 0, 0, 154, 691, 1, 0, 0, 0, 156, 
	157, 5, 97, 0, 0, 157, 158, 5, 110, 0, 0, 158, 159, 5, 100, 0, 0, 159, 
	3, 1, 0, 0, 0, 160, 161, 5, 111, 0, 0, 161, 162, 5, 114, 0, 0, 162, 5, 
	1, 0, 0, 0, 163, 164, 5, 98, 0, 0, 164, 165, 5, 117, 0, 0, 165, 166, 5, 
	116, 0, 0, 166, 167, 5, 32, 0, 0, 167, 168, 5, 110, 0, 0, 168, 169, 5, 
	111, 0, 0, 169, 170, 5, 116, 0, 0, 170, 7, 1, 0, 0, 0, 171, 172, 5, 102, 
	0, 0, 172, 173, 5, 114, 0, 0, 173, 174, 5, 111, 0, 0, 174, 175, 5, 109, 
	0, 0, 175, 9, 1, 0, 0, 0, 176, 177, 5, 109, 0, 0, 177, 178, 5, 111, 0, 
	0, 178, 179, 5, 100, 0, 0, 179, 180, 5, 101, 0, 0, 180, 181, 5, 108, 0, 
	0, 181, 11, 1, 0, 0, 0, 182, 183, 5, 115, 0, 0, 183, 184, 5, 99, 0, 0, 
	184, 185, 5, 104, 0, 0, 185, 186, 5, 101, 0, 0, 186, 187, 5, 109, 0, 0, 
	187, 188, 5, 97, 0, 0, 188, 13, 1, 0, 0, 0, 189, 190, 5, 49, 0, 0, 190, 
	191, 5, 46, 0, 0, 191, 192, 5, 49, 0, 0, 192, 15, 1, 0, 0, 0, 193, 194, 
	5, 116, 0, 0, 194, 195, 5, 121, 0, 0, 195, 196, 5, 112, 0, 0, 196, 197, 
	5, 101, 0, 0, 197, 17, 1, 0, 0, 0, 198, 199, 5, 99, 0, 0, 199, 200, 5, 
	111, 0, 0, 200, 201, 5, 110, 0, 0, 201, 202, 5, 100, 0, 0, 202, 203, 5, 
	105, 0, 0, 203, 204, 5, 116, 0, 0, 204, 205, 5, 105, 0, 0, 205, 206, 5, 
	111, 0, 0, 206, 207, 5, 110, 0, 0, 207, 19, 1, 0, 0, 0, 208, 209, 5, 109, 
	0, 0, 209, 210, 5, 97, 0, 0, 210, 216, 5, 112, 0, 0, 211, 212, 5, 108, 
	0, 0, 212, 213, 5, 105, 0, 0, 213, 214, 5, 115, 0, 0, 214, 216, 5, 116, 
	0, 0, 215, 208, 1, 0, 0, 0, 215, 211, 1, 0, 0, 0, 216, 21, 1, 0, 0, 0, 
	217, 218, 5, 98, 0, 0, 218, 219, 5, 111, 0, 0, 219, 220, 5, 111, 0, 0, 
	220, 267, 5, 108, 0, 0, 221, 222, 5, 115, 0, 0, 222, 223, 5, 116, 0, 0, 
	223, 224, 5, 114, 0, 0, 224, 225, 5, 105, 0, 0, 225, 226, 5, 110, 0, 0, 
	226, 267, 5, 103, 0, 0, 227, 228, 5, 105, 0, 0, 228, 229, 5, 110, 0, 0, 
	229, 267, 5, 116, 0, 0, 230, 231, 5, 117, 0, 0, 231, 232, 5, 105, 0, 0, 
	232, 233, 5, 110, 0, 0, 233, 267, 5, 116, 0, 0, 234, 235, 5, 100, 0, 0, 
	235, 236, 5, 111, 0, 0, 236, 237, 5, 117, 0, 0, 237, 238, 5, 98, 0, 0, 
	238, 239, 5, 108, 0, 0, 239, 267, 5, 101, 0, 0, 240, 241, 5, 100, 0, 0, 
	241, 242, 5, 117, 0, 0, 242, 243, 5, 114, 0, 0, 243, 244, 5, 97, 0, 0, 
	244, 245, 5, 116, 0, 0, 245, 246, 5, 105, 0, 0, 246, 247, 5, 111, 0, 0, 
	247, 267, 5, 110, 0, 0, 248, 249, 5, 116, 0, 0, 249, 250, 5, 105, 0, 0, 
	250, 251, 5, 109, 0, 0, 251, 252, 5, 101, 0, 0, 252, 253, 5, 115, 0, 0, 
	253, 254, 5, 116, 0, 0, 254, 255, 5, 97, 0, 0, 255, 256, 5, 109, 0, 0, 
	256, 267, 5, 112, 0, 0, 257, 258, 5, 105, 0, 0, 258, 259, 5, 112, 0, 0, 
	259, 260, 5, 97, 0, 0, 260, 261, 5, 100, 0, 0, 261, 262, 5, 100, 0, 0, 
	262, 263, 5, 114, 0, 0, 263, 264, 5, 101, 0, 0, 264, 265, 5, 115, 0, 0, 
	265, 267, 5, 115, 0, 0, 266, 217, 1, 0, 0, 0, 266, 221, 1, 0, 0, 0, 266, 
	227, 1, 0, 0, 0, 266, 230, 1, 0, 0, 0, 266, 234, 1, 0, 0, 0, 266, 240, 
	1, 0, 0, 0, 266, 248, 1, 0, 0, 0, 266, 257, 1, 0, 0, 0, 267, 23, 1, 0, 
	0, 0, 268, 269, 5, 114, 0, 0, 269, 270, 5, 101, 0, 0, 270, 271, 5, 108, 
	0, 0, 271, 272, 5, 97, 0, 0, 272, 273, 5, 116, 0, 0, 273, 274, 5, 105, 
	0, 0, 274, 275, 5, 111, 0, 0, 275, 276, 5, 110, 0, 0, 276, 277, 5, 115, 
	0, 0, 277, 25, 1, 0, 0, 0, 278, 279, 5, 100, 0, 0, 279, 280, 5, 101, 0, 
	0, 280, 281, 5, 102, 0, 0, 281, 282, 5, 105, 0, 0, 282, 283, 5, 110, 0, 
	0, 283, 284, 5, 101, 0, 0, 284, 27, 1, 0, 0, 0, 285, 286, 5, 119, 0, 0, 
	286, 287, 5, 105, 0, 0, 287, 288, 5, 116, 0, 0, 288, 289, 5, 104, 0, 0, 
	289, 29, 1, 0, 0, 0, 290, 293, 3, 116, 57, 0, 291, 293, 5, 95, 0, 0, 292, 
	290, 1, 0, 0, 0, 292, 291, 1, 0, 0, 0, 293, 300, 1, 0, 0, 0, 294, 299, 
	3, 116, 57, 0, 295, 299, 3, 118, 58, 0, 296, 299, 5, 95, 0, 0, 297, 299, 
	3, 94, 46, 0, 298, 294, 1, 0, 0, 0, 298, 295, 1, 0, 0, 0, 298, 296, 1, 
	0, 0, 0, 298, 297, 1, 0, 0, 0, 299, 302, 1, 0, 0, 0, 300, 298, 1, 0, 0, 
	0, 300, 301, 1, 0, 0, 0, 301, 31, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 303, 
	305, 7, 0, 0, 0, 304, 303, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 304, 
	1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 33, 1, 0, 0, 0, 308, 310, 3, 32, 
	15, 0, 309, 308, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 316, 1, 0, 0, 0, 
	311, 313, 5, 13, 0, 0, 312, 311, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 
	314, 1, 0, 0, 0, 314, 317, 5, 10, 0, 0, 315, 317, 2, 12, 13, 0, 316, 312, 
	1, 0, 0, 0, 316, 315, 1, 0, 0, 0, 317, 319, 1, 0, 0, 0, 318, 320, 3, 32, 
	15, 0, 319, 318, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 322, 1, 0, 0, 0, 
	321, 323, 3, 34, 16, 0, 322, 321, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 
	35, 1, 0, 0, 0, 324, 325, 5, 46, 0, 0, 325, 37, 1, 0, 0, 0, 326, 327, 5, 
	42, 0, 0, 327, 39, 1, 0, 0, 0, 328, 329, 5, 35, 0, 0, 329, 41, 1, 0, 0, 
	0, 330, 331, 5, 58, 0, 0, 331, 43, 1, 0, 0, 0, 332, 333, 5, 44, 0, 0, 333, 
	45, 1, 0, 0, 0, 334, 335, 5, 40, 0, 0, 335, 47, 1, 0, 0, 0, 336, 337, 5, 
	41, 0, 0, 337, 49, 1, 0, 0, 0, 338, 339, 5, 60, 0, 0, 339, 51, 1, 0, 0, 
	0, 340, 341, 5, 62, 0, 0, 341, 53, 1, 0, 0, 0, 342, 343, 5, 91, 0, 0, 343, 
	55, 1, 0, 0, 0, 344, 345, 5, 93, 0, 0, 345, 57, 1, 0, 0, 0, 346, 347, 5, 
	123, 0, 0, 347, 348, 1, 0, 0, 0, 348, 349, 6, 28, 0, 0, 349, 59, 1, 0, 
	0, 0, 350, 351, 5, 35, 0, 0, 351, 61, 1, 0, 0, 0, 352, 353, 5, 58, 0, 0, 
	353, 63, 1, 0, 0, 0, 354, 355, 5, 44, 0, 0, 355, 65, 1, 0, 0, 0, 356, 357, 
	5, 61, 0, 0, 357, 358, 5, 61, 0, 0, 358, 67, 1, 0, 0, 0, 359, 360, 5, 33, 
	0, 0, 360, 361, 5, 61, 0, 0, 361, 69, 1, 0, 0, 0, 362, 363, 5, 105, 0, 
	0, 363, 364, 5, 110, 0, 0, 364, 71, 1, 0, 0, 0, 365, 366, 5, 60, 0, 0, 
	366, 73, 1, 0, 0, 0, 367, 368, 5, 60, 0, 0, 368, 369, 5, 61, 0, 0, 369, 
	75, 1, 0, 0, 0, 370, 371, 5, 62, 0, 0, 371, 372, 5, 61, 0, 0, 372, 77, 
	1, 0, 0, 0, 373, 374, 5, 62, 0, 0, 374, 79, 1, 0, 0, 0, 375, 376, 5, 38, 
	0, 0, 376, 377, 5, 38, 0, 0, 377, 81, 1, 0, 0, 0, 378, 379, 5, 124, 0, 
	0, 379, 380, 5, 124, 0, 0, 380, 83, 1, 0, 0, 0, 381, 382, 5, 91, 0, 0, 
	382, 85, 1, 0, 0, 0, 383, 384, 5, 93, 0, 0, 384, 87, 1, 0, 0, 0, 385, 386, 
	5, 40, 0, 0, 386, 89, 1, 0, 0, 0, 387, 388, 5, 41, 0, 0, 388, 91, 1, 0, 
	0, 0, 389, 390, 5, 46, 0, 0, 390, 93, 1, 0, 0, 0, 391, 392, 5, 45, 0, 0, 
	392, 95, 1, 0, 0, 0, 393, 394, 5, 33, 0, 0, 394, 97, 1, 0, 0, 0, 395, 396, 
	5, 63, 0, 0, 396, 99, 1, 0, 0, 0, 397, 398, 5, 43, 0, 0, 398, 101, 1, 0, 
	0, 0, 399, 400, 5, 42, 0, 0, 400, 103, 1, 0, 0, 0, 401, 402, 5, 47, 0, 
	0, 402, 105, 1, 0, 0, 0, 403, 404, 5, 37, 0, 0, 404, 107, 1, 0, 0, 0, 405, 
	406, 5, 116, 0, 0, 406, 407, 5, 114, 0, 0, 407, 408, 5, 117, 0, 0, 408, 
	409, 5, 101, 0, 0, 409, 109, 1, 0, 0, 0, 410, 411, 5, 102, 0, 0, 411, 412, 
	5, 97, 0, 0, 412, 413, 5, 108, 0, 0, 413, 414, 5, 115, 0, 0, 414, 415, 
	5, 101, 0, 0, 415, 111, 1, 0, 0, 0, 416, 417, 5, 110, 0, 0, 417, 418, 5, 
	117, 0, 0, 418, 419, 5, 108, 0, 0, 419, 420, 5, 108, 0, 0, 420, 113, 1, 
	0, 0, 0, 421, 422, 5, 92, 0, 0, 422, 115, 1, 0, 0, 0, 423, 424, 7, 1, 0, 
	0, 424, 117, 1, 0, 0, 0, 425, 426, 2, 48, 57, 0, 426, 119, 1, 0, 0, 0, 
	427, 429, 7, 2, 0, 0, 428, 430, 7, 3, 0, 0, 429, 428, 1, 0, 0, 0, 429, 
	430, 1, 0, 0, 0, 430, 432, 1, 0, 0, 0, 431, 433, 3, 118, 58, 0, 432, 431, 
	1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 432, 1, 0, 0, 0, 434, 435, 1, 0, 
	0, 0, 435, 121, 1, 0, 0, 0, 436, 437, 7, 4, 0, 0, 437, 123, 1, 0, 0, 0, 
	438, 439, 7, 5, 0, 0, 439, 125, 1, 0, 0, 0, 440, 445, 3, 128, 63, 0, 441, 
	445, 3, 132, 65, 0, 442, 445, 3, 134, 66, 0, 443, 445, 3, 130, 64, 0, 444, 
	440, 1, 0, 0, 0, 444, 441, 1, 0, 0, 0, 444, 442, 1, 0, 0, 0, 444, 443, 
	1, 0, 0, 0, 445, 127, 1, 0, 0, 0, 446, 447, 3, 114, 56, 0, 447, 448, 7, 
	6, 0, 0, 448, 129, 1, 0, 0, 0, 449, 450, 3, 114, 56, 0, 450, 451, 2, 48, 
	51, 0, 451, 452, 2, 48, 55, 0, 452, 453, 2, 48, 55, 0, 453, 131, 1, 0, 
	0, 0, 454, 455, 3, 114, 56, 0, 455, 456, 7, 7, 0, 0, 456, 457, 3, 122, 
	60, 0, 457, 458, 3, 122, 60, 0, 458, 133, 1, 0, 0, 0, 459, 460, 3, 114, 
	56, 0, 460, 461, 5, 117, 0, 0, 461, 462, 3, 122, 60, 0, 462, 463, 3, 122, 
	60, 0, 463, 464, 3, 122, 60, 0, 464, 465, 3, 122, 60, 0, 465, 478, 1, 0, 
	0, 0, 466, 467, 3, 114, 56, 0, 467, 468, 5, 85, 0, 0, 468, 469, 3, 122, 
	60, 0, 469, 470, 3, 122, 60, 0, 470, 471, 3, 122, 60, 0, 471, 472, 3, 122, 
	60, 0, 472, 473, 3, 122, 60, 0, 473, 474, 3, 122, 60, 0, 474, 475, 3, 122, 
	60, 0, 475, 476, 3, 122, 60, 0, 476, 478, 1, 0, 0, 0, 477, 459, 1, 0, 0, 
	0, 477, 466, 1, 0, 0, 0, 478, 135, 1, 0, 0, 0, 479, 480, 5, 47, 0, 0, 480, 
	481, 5, 47, 0, 0, 481, 485, 1, 0, 0, 0, 482, 484, 8, 8, 0, 0, 483, 482, 
	1, 0, 0, 0, 484, 487, 1, 0, 0, 0, 485, 483, 1, 0, 0, 0, 485, 486, 1, 0, 
	0, 0, 486, 488, 1, 0, 0, 0, 487, 485, 1, 0, 0, 0, 488, 489, 6, 67, 1, 0, 
	489, 137, 1, 0, 0, 0, 490, 492, 3, 118, 58, 0, 491, 490, 1, 0, 0, 0, 492, 
	493, 1, 0, 0, 0, 493, 491, 1, 0, 0, 0, 493, 494, 1, 0, 0, 0, 494, 495, 
	1, 0, 0, 0, 495, 497, 5, 46, 0, 0, 496, 498, 3, 118, 58, 0, 497, 496, 1, 
	0, 0, 0, 498, 499, 1, 0, 0, 0, 499, 497, 1, 0, 0, 0, 499, 500, 1, 0, 0, 
	0, 500, 502, 1, 0, 0, 0, 501, 503, 3, 120, 59, 0, 502, 501, 1, 0, 0, 0, 
	502, 503, 1, 0, 0, 0, 503, 521, 1, 0, 0, 0, 504, 506, 3, 118, 58, 0, 505, 
	504, 1, 0, 0, 0, 506, 507, 1, 0, 0, 0, 507, 505, 1, 0, 0, 0, 507, 508, 
	1, 0, 0, 0, 508, 509, 1, 0, 0, 0, 509, 510, 3, 120, 59, 0, 510, 521, 1, 
	0, 0, 0, 511, 513, 5, 46, 0, 0, 512, 514, 3, 118, 58, 0, 513, 512, 1, 0, 
	0, 0, 514, 515, 1, 0, 0, 0, 515, 513, 1, 0, 0, 0, 515, 516, 1, 0, 0, 0, 
	516, 518, 1, 0, 0, 0, 517, 519, 3, 120, 59, 0, 518, 517, 1, 0, 0, 0, 518, 
	519, 1, 0, 0, 0, 519, 521, 1, 0, 0, 0, 520, 491, 1, 0, 0, 0, 520, 505, 
	1, 0, 0, 0, 520, 511, 1, 0, 0, 0, 521, 139, 1, 0, 0, 0, 522, 524, 3, 118, 
	58, 0, 523, 522, 1, 0, 0, 0, 524, 525, 1, 0, 0, 0, 525, 523, 1, 0, 0, 0, 
	525, 526, 1, 0, 0, 0, 526, 536, 1, 0, 0, 0, 527, 528, 5, 48, 0, 0, 528, 
	529, 5, 120, 0, 0, 529, 531, 1, 0, 0, 0, 530, 532, 3, 122, 60, 0, 531, 
	530, 1, 0, 0, 0, 532, 533, 1, 0, 0, 0, 533, 531, 1, 0, 0, 0, 533, 534, 
	1, 0, 0, 0, 534, 536, 1, 0, 0, 0, 535, 523, 1, 0, 0, 0, 535, 527, 1, 0, 
	0, 0, 536, 141, 1, 0, 0, 0, 537, 539, 3, 118, 58, 0, 538, 537, 1, 0, 0, 
	0, 539, 540, 1, 0, 0, 0, 540, 538, 1, 0, 0, 0, 540, 541, 1, 0, 0, 0, 541, 
	542, 1, 0, 0, 0, 542, 543, 7, 9, 0, 0, 543, 555, 1, 0, 0, 0, 544, 545, 
	5, 48, 0, 0, 545, 546, 5, 120, 0, 0, 546, 548, 1, 0, 0, 0, 547, 549, 3, 
	122, 60, 0, 548, 547, 1, 0, 0, 0, 549, 550, 1, 0, 0, 0, 550, 548, 1, 0, 
	0, 0, 550, 551, 1, 0, 0, 0, 551, 552, 1, 0, 0, 0, 552, 553, 7, 9, 0, 0, 
	553, 555, 1, 0, 0, 0, 554, 538, 1, 0, 0, 0, 554, 544, 1, 0, 0, 0, 555, 
	143, 1, 0, 0, 0, 556, 561, 5, 34, 0, 0, 557, 560, 3, 126, 62, 0, 558, 560, 
	8, 10, 0, 0, 559, 557, 1, 0, 0, 0, 559, 558, 1, 0, 0, 0, 560, 563, 1, 0, 
	0, 0, 561, 559, 1, 0, 0, 0, 561, 562, 1, 0, 0, 0, 562, 564, 1, 0, 0, 0, 
	563, 561, 1, 0, 0, 0, 564, 653, 5, 34, 0, 0, 565, 570, 5, 39, 0, 0, 566, 
	569, 3, 126, 62, 0, 567, 569, 8, 11, 0, 0, 568, 566, 1, 0, 0, 0, 568, 567, 
	1, 0, 0, 0, 569, 572, 1, 0, 0, 0, 570, 568, 1, 0, 0, 0, 570, 571, 1, 0, 
	0, 0, 571, 573, 1, 0, 0, 0, 572, 570, 1, 0, 0, 0, 573, 653, 5, 39, 0, 0, 
	574, 575, 5, 34, 0, 0, 575, 576, 5, 34, 0, 0, 576, 577, 5, 34, 0, 0, 577, 
	582, 1, 0, 0, 0, 578, 581, 3, 126, 62, 0, 579, 581, 8, 12, 0, 0, 580, 578, 
	1, 0, 0, 0, 580, 579, 1, 0, 0, 0, 581, 584, 1, 0, 0, 0, 582, 583, 1, 0, 
	0, 0, 582, 580, 1, 0, 0, 0, 583, 585, 1, 0, 0, 0, 584, 582, 1, 0, 0, 0, 
	585, 586, 5, 34, 0, 0, 586, 587, 5, 34, 0, 0, 587, 653, 5, 34, 0, 0, 588, 
	589, 5, 39, 0, 0, 589, 590, 5, 39, 0, 0, 590, 591, 5, 39, 0, 0, 591, 596, 
	1, 0, 0, 0, 592, 595, 3, 126, 62, 0, 593, 595, 8, 12, 0, 0, 594, 592, 1, 
	0, 0, 0, 594, 593, 1, 0, 0, 0, 595, 598, 1, 0, 0, 0, 596, 597, 1, 0, 0, 
	0, 596, 594, 1, 0, 0, 0, 597, 599, 1, 0, 0, 0, 598, 596, 1, 0, 0, 0, 599, 
	600, 5, 39, 0, 0, 600, 601, 5, 39, 0, 0, 601, 653, 5, 39, 0, 0, 602, 603, 
	3, 124, 61, 0, 603, 607, 5, 34, 0, 0, 604, 606, 8, 13, 0, 0, 605, 604, 
	1, 0, 0, 0, 606, 609, 1, 0, 0, 0, 607, 605, 1, 0, 0, 0, 607, 608, 1, 0, 
	0, 0, 608, 610, 1, 0, 0, 0, 609, 607, 1, 0, 0, 0, 610, 611, 5, 34, 0, 0, 
	611, 653, 1, 0, 0, 0, 612, 613, 3, 124, 61, 0, 613, 617, 5, 39, 0, 0, 614, 
	616, 8, 14, 0, 0, 615, 614, 1, 0, 0, 0, 616, 619, 1, 0, 0, 0, 617, 615, 
	1, 0, 0, 0, 617, 618, 1, 0, 0, 0, 618, 620, 1, 0, 0, 0, 619, 617, 1, 0, 
	0, 0, 620, 621, 5, 39, 0, 0, 621, 653, 1, 0, 0, 0, 622, 623, 3, 124, 61, 
	0, 623, 624, 5, 34, 0, 0, 624, 625, 5, 34, 0, 0, 625, 626, 5, 34, 0, 0, 
	626, 630, 1, 0, 0, 0, 627, 629, 9, 0, 0, 0, 628, 627, 1, 0, 0, 0, 629, 
	632, 1, 0, 0, 0, 630, 631, 1, 0, 0, 0, 630, 628, 1, 0, 0, 0, 631, 633, 
	1, 0, 0, 0, 632, 630, 1, 0, 0, 0, 633, 634, 5, 34, 0, 0, 634, 635, 5, 34, 
	0, 0, 635, 636, 5, 34, 0, 0, 636, 653, 1, 0, 0, 0, 637, 638, 3, 124, 61, 
	0, 638, 639, 5, 39, 0, 0, 639, 640, 5, 39, 0, 0, 640, 641, 5, 39, 0, 0, 
	641, 645, 1, 0, 0, 0, 642, 644, 9, 0, 0, 0, 643, 642, 1, 0, 0, 0, 644, 
	647, 1, 0, 0, 0, 645, 646, 1, 0, 0, 0, 645, 643, 1, 0, 0, 0, 646, 648, 
	1, 0, 0, 0, 647, 645, 1, 0, 0, 0, 648, 649, 5, 39, 0, 0, 649, 650, 5, 39, 
	0, 0, 650, 651, 5, 39, 0, 0, 651, 653, 1, 0, 0, 0, 652, 556, 1, 0, 0, 0, 
	652, 565, 1, 0, 0, 0, 652, 574, 1, 0, 0, 0, 652, 588, 1, 0, 0, 0, 652, 
	602, 1, 0, 0, 0, 652, 612, 1, 0, 0, 0, 652, 622, 1, 0, 0, 0, 652, 637, 
	1, 0, 0, 0, 653, 145, 1, 0, 0, 0, 654, 655, 7, 15, 0, 0, 655, 656, 3, 144, 
	71, 0, 656, 147, 1, 0, 0, 0, 657, 660, 3, 116, 57, 0, 658, 660, 5, 95, 
	0, 0, 659, 657, 1, 0, 0, 0, 659, 658, 1, 0, 0, 0, 660, 667, 1, 0, 0, 0, 
	661, 666, 3, 116, 57, 0, 662, 666, 3, 118, 58, 0, 663, 666, 5, 95, 0, 0, 
	664, 666, 3, 94, 46, 0, 665, 661, 1, 0, 0, 0, 665, 662, 1, 0, 0, 0, 665, 
	663, 1, 0, 0, 0, 665, 664, 1, 0, 0, 0, 666, 669, 1, 0, 0, 0, 667, 665, 
	1, 0, 0, 0, 667, 668, 1, 0, 0, 0, 668, 149, 1, 0, 0, 0, 669, 667, 1, 0, 
	0, 0, 670, 672, 7, 0, 0, 0, 671, 670, 1, 0, 0, 0, 672, 673, 1, 0, 0, 0, 
	673, 671, 1, 0, 0, 0, 673, 674, 1, 0, 0, 0, 674, 151, 1, 0, 0, 0, 675, 
	677, 3, 150, 74, 0, 676, 675, 1, 0, 0, 0, 676, 677, 1, 0, 0, 0, 677, 683, 
	1, 0, 0, 0, 678, 680, 5, 13, 0, 0, 679, 678, 1, 0, 0, 0, 679, 680, 1, 0, 
	0, 0, 680, 681, 1, 0, 0, 0, 681, 684, 5, 10, 0, 0, 682, 684, 2, 12, 13, 
	0, 683, 679, 1, 0, 0, 0, 683, 682, 1, 0, 0, 0, 684, 686, 1, 0, 0, 0, 685, 
	687, 3, 150, 74, 0, 686, 685, 1, 0, 0, 0, 686, 687, 1, 0, 0, 0, 687, 689, 
	1, 0, 0, 0, 688, 690, 3, 152, 75, 0, 689, 688, 1, 0, 0, 0, 689, 690, 1, 
	0, 0, 0, 690, 153, 1, 0, 0, 0, 691, 692, 5, 125, 0, 0, 692, 693, 1, 0, 
	0, 0, 693, 694, 6, 76, 2, 0, 694, 155, 1, 0, 0, 0, 53, 0, 1, 215, 266, 
	292, 298, 300, 306, 309, 312, 316, 319, 322, 429, 434, 444, 477, 485, 493, 
	499, 502, 507, 515, 518, 520, 525, 533, 535, 540, 550, 554, 559, 561, 568, 
	570, 580, 582, 594, 596, 607, 617, 630, 645, 652, 659, 665, 667, 673, 676, 
	679, 683, 686, 689, 3, 5, 1, 0, 0, 1, 0, 4, 0, 0,
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

// OpenFGALexerInit initializes any static state used to implement OpenFGALexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewOpenFGALexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func OpenFGALexerInit() {
  staticData := &OpenFGALexerLexerStaticData
  staticData.once.Do(openfgalexerLexerInit)
}

// NewOpenFGALexer produces a new lexer instance for the optional input antlr.CharStream.
func NewOpenFGALexer(input antlr.CharStream) *OpenFGALexer {
  OpenFGALexerInit()
	l := new(OpenFGALexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &OpenFGALexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "OpenFGALexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// OpenFGALexer tokens.
const (
	OpenFGALexerAND = 1
	OpenFGALexerOR = 2
	OpenFGALexerBUT_NOT = 3
	OpenFGALexerFROM = 4
	OpenFGALexerMODEL = 5
	OpenFGALexerSCHEMA = 6
	OpenFGALexerSCHEMA_VERSION = 7
	OpenFGALexerTYPE = 8
	OpenFGALexerCONDITION = 9
	OpenFGALexerCONDITION_PARAM_CONTAINER = 10
	OpenFGALexerCONDITION_PARAM_TYPE = 11
	OpenFGALexerRELATIONS = 12
	OpenFGALexerDEFINE = 13
	OpenFGALexerKEYWORD_WITH = 14
	OpenFGALexerIDENTIFIER = 15
	OpenFGALexerWHITESPACE = 16
	OpenFGALexerNEWLINE = 17
	OpenFGALexerDOT = 18
	OpenFGALexerSTAR = 19
	OpenFGALexerHASH = 20
	OpenFGALexerCOLON = 21
	OpenFGALexerCOMMA = 22
	OpenFGALexerLPAREN = 23
	OpenFGALexerRPAREN = 24
	OpenFGALexerLESS = 25
	OpenFGALexerGREATER = 26
	OpenFGALexerLBRACKET = 27
	OpenFGALexerRPRACKET = 28
	OpenFGALexerOPEN_CEL = 29
	OpenFGALexerCEL_HASH = 30
	OpenFGALexerCEL_COLON = 31
	OpenFGALexerCEL_COMMA = 32
	OpenFGALexerEQUALS = 33
	OpenFGALexerNOT_EQUALS = 34
	OpenFGALexerIN = 35
	OpenFGALexerCEL_LESS = 36
	OpenFGALexerLESS_EQUALS = 37
	OpenFGALexerGREATER_EQUALS = 38
	OpenFGALexerCEL_GREATER = 39
	OpenFGALexerLOGICAL_AND = 40
	OpenFGALexerLOGICAL_OR = 41
	OpenFGALexerCEL_LBRACKET = 42
	OpenFGALexerCEL_RPRACKET = 43
	OpenFGALexerCEL_LPAREN = 44
	OpenFGALexerCEL_RPAREN = 45
	OpenFGALexerCEL_DOT = 46
	OpenFGALexerMINUS = 47
	OpenFGALexerEXCLAM = 48
	OpenFGALexerQUESTIONMARK = 49
	OpenFGALexerPLUS = 50
	OpenFGALexerCEL_STAR = 51
	OpenFGALexerSLASH = 52
	OpenFGALexerPERCENT = 53
	OpenFGALexerCEL_TRUE = 54
	OpenFGALexerCEL_FALSE = 55
	OpenFGALexerNUL = 56
	OpenFGALexerCEL_COMMENT = 57
	OpenFGALexerNUM_FLOAT = 58
	OpenFGALexerNUM_INT = 59
	OpenFGALexerNUM_UINT = 60
	OpenFGALexerSTRING = 61
	OpenFGALexerBYTES = 62
	OpenFGALexerCEL_IDENTIFIER = 63
	OpenFGALexerCEL_WHITESPACE = 64
	OpenFGALexerCEL_NEWLINE = 65
	OpenFGALexerCLOSE_CEL = 66
)

// OpenFGALexerCEL is the OpenFGALexer mode.
const OpenFGALexerCEL = 1

