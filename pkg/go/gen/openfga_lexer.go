// Code generated from /app/OpenFGA.g4 by ANTLR 4.13.0. DO NOT EDIT.

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
	modeNames    []string
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
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'  '", "'\\u0009'", "'model'", "'schema'", "'type'", "'relations'",
		"'define'", "':'", "'['", "','", "']'", "'and'", "'or'", "'but not'",
		"'from'", "':*'", "'#'", "'\\r'", "'\\n'", "' '", "'1.1'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "ALPHA_NUMERIC",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "ALPHA_NUMERIC",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 22, 133, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 4, 21, 130, 8, 21,
		11, 21, 12, 21, 131, 0, 0, 22, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 1, 0, 1, 5, 0, 45, 45,
		48, 57, 65, 90, 95, 95, 97, 122, 133, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 1, 45, 1, 0, 0, 0, 3, 48, 1, 0, 0, 0, 5, 50, 1, 0, 0, 0,
		7, 56, 1, 0, 0, 0, 9, 63, 1, 0, 0, 0, 11, 68, 1, 0, 0, 0, 13, 78, 1, 0,
		0, 0, 15, 85, 1, 0, 0, 0, 17, 87, 1, 0, 0, 0, 19, 89, 1, 0, 0, 0, 21, 91,
		1, 0, 0, 0, 23, 93, 1, 0, 0, 0, 25, 97, 1, 0, 0, 0, 27, 100, 1, 0, 0, 0,
		29, 108, 1, 0, 0, 0, 31, 113, 1, 0, 0, 0, 33, 116, 1, 0, 0, 0, 35, 118,
		1, 0, 0, 0, 37, 120, 1, 0, 0, 0, 39, 122, 1, 0, 0, 0, 41, 124, 1, 0, 0,
		0, 43, 129, 1, 0, 0, 0, 45, 46, 5, 32, 0, 0, 46, 47, 5, 32, 0, 0, 47, 2,
		1, 0, 0, 0, 48, 49, 5, 9, 0, 0, 49, 4, 1, 0, 0, 0, 50, 51, 5, 109, 0, 0,
		51, 52, 5, 111, 0, 0, 52, 53, 5, 100, 0, 0, 53, 54, 5, 101, 0, 0, 54, 55,
		5, 108, 0, 0, 55, 6, 1, 0, 0, 0, 56, 57, 5, 115, 0, 0, 57, 58, 5, 99, 0,
		0, 58, 59, 5, 104, 0, 0, 59, 60, 5, 101, 0, 0, 60, 61, 5, 109, 0, 0, 61,
		62, 5, 97, 0, 0, 62, 8, 1, 0, 0, 0, 63, 64, 5, 116, 0, 0, 64, 65, 5, 121,
		0, 0, 65, 66, 5, 112, 0, 0, 66, 67, 5, 101, 0, 0, 67, 10, 1, 0, 0, 0, 68,
		69, 5, 114, 0, 0, 69, 70, 5, 101, 0, 0, 70, 71, 5, 108, 0, 0, 71, 72, 5,
		97, 0, 0, 72, 73, 5, 116, 0, 0, 73, 74, 5, 105, 0, 0, 74, 75, 5, 111, 0,
		0, 75, 76, 5, 110, 0, 0, 76, 77, 5, 115, 0, 0, 77, 12, 1, 0, 0, 0, 78,
		79, 5, 100, 0, 0, 79, 80, 5, 101, 0, 0, 80, 81, 5, 102, 0, 0, 81, 82, 5,
		105, 0, 0, 82, 83, 5, 110, 0, 0, 83, 84, 5, 101, 0, 0, 84, 14, 1, 0, 0,
		0, 85, 86, 5, 58, 0, 0, 86, 16, 1, 0, 0, 0, 87, 88, 5, 91, 0, 0, 88, 18,
		1, 0, 0, 0, 89, 90, 5, 44, 0, 0, 90, 20, 1, 0, 0, 0, 91, 92, 5, 93, 0,
		0, 92, 22, 1, 0, 0, 0, 93, 94, 5, 97, 0, 0, 94, 95, 5, 110, 0, 0, 95, 96,
		5, 100, 0, 0, 96, 24, 1, 0, 0, 0, 97, 98, 5, 111, 0, 0, 98, 99, 5, 114,
		0, 0, 99, 26, 1, 0, 0, 0, 100, 101, 5, 98, 0, 0, 101, 102, 5, 117, 0, 0,
		102, 103, 5, 116, 0, 0, 103, 104, 5, 32, 0, 0, 104, 105, 5, 110, 0, 0,
		105, 106, 5, 111, 0, 0, 106, 107, 5, 116, 0, 0, 107, 28, 1, 0, 0, 0, 108,
		109, 5, 102, 0, 0, 109, 110, 5, 114, 0, 0, 110, 111, 5, 111, 0, 0, 111,
		112, 5, 109, 0, 0, 112, 30, 1, 0, 0, 0, 113, 114, 5, 58, 0, 0, 114, 115,
		5, 42, 0, 0, 115, 32, 1, 0, 0, 0, 116, 117, 5, 35, 0, 0, 117, 34, 1, 0,
		0, 0, 118, 119, 5, 13, 0, 0, 119, 36, 1, 0, 0, 0, 120, 121, 5, 10, 0, 0,
		121, 38, 1, 0, 0, 0, 122, 123, 5, 32, 0, 0, 123, 40, 1, 0, 0, 0, 124, 125,
		5, 49, 0, 0, 125, 126, 5, 46, 0, 0, 126, 127, 5, 49, 0, 0, 127, 42, 1,
		0, 0, 0, 128, 130, 7, 0, 0, 0, 129, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0,
		0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 44, 1, 0, 0, 0, 2,
		0, 131, 0,
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
	l.GrammarFileName = "OpenFGA.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// OpenFGALexer tokens.
const (
	OpenFGALexerT__0          = 1
	OpenFGALexerT__1          = 2
	OpenFGALexerT__2          = 3
	OpenFGALexerT__3          = 4
	OpenFGALexerT__4          = 5
	OpenFGALexerT__5          = 6
	OpenFGALexerT__6          = 7
	OpenFGALexerT__7          = 8
	OpenFGALexerT__8          = 9
	OpenFGALexerT__9          = 10
	OpenFGALexerT__10         = 11
	OpenFGALexerT__11         = 12
	OpenFGALexerT__12         = 13
	OpenFGALexerT__13         = 14
	OpenFGALexerT__14         = 15
	OpenFGALexerT__15         = 16
	OpenFGALexerT__16         = 17
	OpenFGALexerT__17         = 18
	OpenFGALexerT__18         = 19
	OpenFGALexerT__19         = 20
	OpenFGALexerT__20         = 21
	OpenFGALexerALPHA_NUMERIC = 22
)
