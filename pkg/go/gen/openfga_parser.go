// Code generated from /app/OpenFGA.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // OpenFGA

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

var OpenFGAParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func openfgaParserInit() {
  staticData := &OpenFGAParserStaticData
  staticData.LiteralNames = []string{
    "", "'\\r'", "'\\n'", "", "", "'model'", "'type'", "'schema'", "'1.1'", 
    "'relations'", "'define'", "'and'", "'or'", "'but not'", "'from'", "':'", 
    "'#'", "'*'", "'['", "']'", "','",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "MULTILINE_COMMENT", "INDENT", "MODEL", "TYPE", "SCHEMA", 
    "SCHEMA_VERSION", "RELATIONS", "DEFINE", "AND", "OR", "BUT_NOT", "FROM", 
    "COLON", "HASH", "WILDCARD", "L_SQUARE", "R_SQUARE", "COMMA", "SYMBOL", 
    "ALPHA_NUMERIC_CHAR", "ALPHA_NUMERIC", "NEWLINES", "WS",
  }
  staticData.RuleNames = []string{
    "main", "modelHeaderComment", "modelHeader", "typeDefs", "typeDef", 
    "relationDeclaration", "relationDef", "relationDefPartials", "relationDefPartialAllOr", 
    "relationDefPartialAllAnd", "relationDefPartialAllButNot", "relationDefDirectAssignment", 
    "relationDefRewrite", "relationDefRelationOnSameObject", "relationDefRelationOnRelatedObject", 
    "relationDefTypeRestriction", "relationDefTypeRestrictionType", "relationDefTypeRestrictionRelation", 
    "relationDefGrouping", "rewriteComputedusersetName", "rewriteTuplesetComputedusersetName", 
    "rewriteTuplesetName", "relationName", "typeName", "schemaVersion", 
    "name",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 25, 195, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2, 
	21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0, 
	1, 0, 1, 0, 3, 0, 56, 8, 0, 1, 1, 1, 1, 5, 1, 60, 8, 1, 10, 1, 12, 1, 63, 
	9, 1, 3, 1, 65, 8, 1, 1, 1, 3, 1, 68, 8, 1, 1, 2, 1, 2, 1, 2, 3, 2, 73, 
	8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 5, 3, 82, 8, 3, 10, 3, 
	12, 3, 85, 9, 3, 1, 4, 3, 4, 88, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 
	4, 1, 4, 4, 4, 97, 8, 4, 11, 4, 12, 4, 98, 3, 4, 101, 8, 4, 1, 5, 3, 5, 
	104, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 
	3, 6, 116, 8, 6, 1, 6, 3, 6, 119, 8, 6, 1, 7, 1, 7, 1, 7, 3, 7, 124, 8, 
	7, 1, 8, 1, 8, 4, 8, 128, 8, 8, 11, 8, 12, 8, 129, 1, 9, 1, 9, 4, 9, 134, 
	8, 9, 11, 9, 12, 9, 135, 1, 10, 1, 10, 4, 10, 140, 8, 10, 11, 10, 12, 10, 
	141, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 148, 8, 11, 10, 11, 12, 11, 151, 
	9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 3, 12, 157, 8, 12, 1, 13, 1, 13, 1, 
	14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 170, 
	8, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 
	20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 4, 25, 
	191, 8, 25, 11, 25, 12, 25, 192, 1, 25, 0, 0, 26, 0, 2, 4, 6, 8, 10, 12, 
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 
	50, 0, 1, 1, 0, 1, 2, 190, 0, 52, 1, 0, 0, 0, 2, 64, 1, 0, 0, 0, 4, 72, 
	1, 0, 0, 0, 6, 83, 1, 0, 0, 0, 8, 87, 1, 0, 0, 0, 10, 103, 1, 0, 0, 0, 
	12, 115, 1, 0, 0, 0, 14, 123, 1, 0, 0, 0, 16, 127, 1, 0, 0, 0, 18, 133, 
	1, 0, 0, 0, 20, 139, 1, 0, 0, 0, 22, 143, 1, 0, 0, 0, 24, 156, 1, 0, 0, 
	0, 26, 158, 1, 0, 0, 0, 28, 160, 1, 0, 0, 0, 30, 164, 1, 0, 0, 0, 32, 171, 
	1, 0, 0, 0, 34, 173, 1, 0, 0, 0, 36, 175, 1, 0, 0, 0, 38, 177, 1, 0, 0, 
	0, 40, 179, 1, 0, 0, 0, 42, 181, 1, 0, 0, 0, 44, 183, 1, 0, 0, 0, 46, 185, 
	1, 0, 0, 0, 48, 187, 1, 0, 0, 0, 50, 190, 1, 0, 0, 0, 52, 53, 3, 4, 2, 
	0, 53, 55, 3, 6, 3, 0, 54, 56, 5, 24, 0, 0, 55, 54, 1, 0, 0, 0, 55, 56, 
	1, 0, 0, 0, 56, 1, 1, 0, 0, 0, 57, 61, 5, 16, 0, 0, 58, 60, 8, 0, 0, 0, 
	59, 58, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 
	0, 0, 0, 62, 65, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 57, 1, 0, 0, 0, 64, 
	65, 1, 0, 0, 0, 65, 67, 1, 0, 0, 0, 66, 68, 5, 3, 0, 0, 67, 66, 1, 0, 0, 
	0, 67, 68, 1, 0, 0, 0, 68, 3, 1, 0, 0, 0, 69, 70, 3, 2, 1, 0, 70, 71, 5, 
	24, 0, 0, 71, 73, 1, 0, 0, 0, 72, 69, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 
	74, 1, 0, 0, 0, 74, 75, 5, 5, 0, 0, 75, 76, 5, 24, 0, 0, 76, 77, 5, 4, 
	0, 0, 77, 78, 5, 7, 0, 0, 78, 79, 3, 48, 24, 0, 79, 5, 1, 0, 0, 0, 80, 
	82, 3, 8, 4, 0, 81, 80, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 
	0, 83, 84, 1, 0, 0, 0, 84, 7, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 88, 5, 
	3, 0, 0, 87, 86, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 
	90, 5, 24, 0, 0, 90, 91, 5, 6, 0, 0, 91, 100, 3, 46, 23, 0, 92, 93, 5, 
	24, 0, 0, 93, 94, 5, 4, 0, 0, 94, 96, 5, 9, 0, 0, 95, 97, 3, 10, 5, 0, 
	96, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 
	0, 0, 0, 99, 101, 1, 0, 0, 0, 100, 92, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 
	101, 9, 1, 0, 0, 0, 102, 104, 5, 3, 0, 0, 103, 102, 1, 0, 0, 0, 103, 104, 
	1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 106, 5, 24, 0, 0, 106, 107, 5, 4, 
	0, 0, 107, 108, 5, 4, 0, 0, 108, 109, 5, 10, 0, 0, 109, 110, 3, 44, 22, 
	0, 110, 111, 5, 15, 0, 0, 111, 112, 3, 12, 6, 0, 112, 11, 1, 0, 0, 0, 113, 
	116, 3, 22, 11, 0, 114, 116, 3, 36, 18, 0, 115, 113, 1, 0, 0, 0, 115, 114, 
	1, 0, 0, 0, 116, 118, 1, 0, 0, 0, 117, 119, 3, 14, 7, 0, 118, 117, 1, 0, 
	0, 0, 118, 119, 1, 0, 0, 0, 119, 13, 1, 0, 0, 0, 120, 124, 3, 16, 8, 0, 
	121, 124, 3, 18, 9, 0, 122, 124, 3, 20, 10, 0, 123, 120, 1, 0, 0, 0, 123, 
	121, 1, 0, 0, 0, 123, 122, 1, 0, 0, 0, 124, 15, 1, 0, 0, 0, 125, 126, 5, 
	12, 0, 0, 126, 128, 3, 36, 18, 0, 127, 125, 1, 0, 0, 0, 128, 129, 1, 0, 
	0, 0, 129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 17, 1, 0, 0, 0, 
	131, 132, 5, 11, 0, 0, 132, 134, 3, 36, 18, 0, 133, 131, 1, 0, 0, 0, 134, 
	135, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 19, 1, 
	0, 0, 0, 137, 138, 5, 13, 0, 0, 138, 140, 3, 36, 18, 0, 139, 137, 1, 0, 
	0, 0, 140, 141, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 
	142, 21, 1, 0, 0, 0, 143, 144, 5, 18, 0, 0, 144, 149, 3, 30, 15, 0, 145, 
	146, 5, 20, 0, 0, 146, 148, 3, 30, 15, 0, 147, 145, 1, 0, 0, 0, 148, 151, 
	1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 152, 1, 0, 
	0, 0, 151, 149, 1, 0, 0, 0, 152, 153, 5, 19, 0, 0, 153, 23, 1, 0, 0, 0, 
	154, 157, 3, 26, 13, 0, 155, 157, 3, 28, 14, 0, 156, 154, 1, 0, 0, 0, 156, 
	155, 1, 0, 0, 0, 157, 25, 1, 0, 0, 0, 158, 159, 3, 38, 19, 0, 159, 27, 
	1, 0, 0, 0, 160, 161, 3, 40, 20, 0, 161, 162, 5, 14, 0, 0, 162, 163, 3, 
	42, 21, 0, 163, 29, 1, 0, 0, 0, 164, 169, 3, 32, 16, 0, 165, 166, 5, 15, 
	0, 0, 166, 170, 5, 17, 0, 0, 167, 168, 5, 16, 0, 0, 168, 170, 3, 34, 17, 
	0, 169, 165, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 
	31, 1, 0, 0, 0, 171, 172, 3, 50, 25, 0, 172, 33, 1, 0, 0, 0, 173, 174, 
	3, 50, 25, 0, 174, 35, 1, 0, 0, 0, 175, 176, 3, 24, 12, 0, 176, 37, 1, 
	0, 0, 0, 177, 178, 3, 50, 25, 0, 178, 39, 1, 0, 0, 0, 179, 180, 3, 50, 
	25, 0, 180, 41, 1, 0, 0, 0, 181, 182, 3, 50, 25, 0, 182, 43, 1, 0, 0, 0, 
	183, 184, 3, 50, 25, 0, 184, 45, 1, 0, 0, 0, 185, 186, 3, 50, 25, 0, 186, 
	47, 1, 0, 0, 0, 187, 188, 5, 8, 0, 0, 188, 49, 1, 0, 0, 0, 189, 191, 5, 
	23, 0, 0, 190, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 190, 1, 0, 0, 
	0, 192, 193, 1, 0, 0, 0, 193, 51, 1, 0, 0, 0, 20, 55, 61, 64, 67, 72, 83, 
	87, 98, 100, 103, 115, 118, 123, 129, 135, 141, 149, 156, 169, 192,
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
  staticData := &OpenFGAParserStaticData
  staticData.once.Do(openfgaParserInit)
}

// NewOpenFGAParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOpenFGAParser(input antlr.TokenStream) *OpenFGAParser {
	OpenFGAParserInit()
	this := new(OpenFGAParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &OpenFGAParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "OpenFGA.g4"

	return this
}


// OpenFGAParser tokens.
const (
	OpenFGAParserEOF = antlr.TokenEOF
	OpenFGAParserT__0 = 1
	OpenFGAParserT__1 = 2
	OpenFGAParserMULTILINE_COMMENT = 3
	OpenFGAParserINDENT = 4
	OpenFGAParserMODEL = 5
	OpenFGAParserTYPE = 6
	OpenFGAParserSCHEMA = 7
	OpenFGAParserSCHEMA_VERSION = 8
	OpenFGAParserRELATIONS = 9
	OpenFGAParserDEFINE = 10
	OpenFGAParserAND = 11
	OpenFGAParserOR = 12
	OpenFGAParserBUT_NOT = 13
	OpenFGAParserFROM = 14
	OpenFGAParserCOLON = 15
	OpenFGAParserHASH = 16
	OpenFGAParserWILDCARD = 17
	OpenFGAParserL_SQUARE = 18
	OpenFGAParserR_SQUARE = 19
	OpenFGAParserCOMMA = 20
	OpenFGAParserSYMBOL = 21
	OpenFGAParserALPHA_NUMERIC_CHAR = 22
	OpenFGAParserALPHA_NUMERIC = 23
	OpenFGAParserNEWLINES = 24
	OpenFGAParserWS = 25
)

// OpenFGAParser rules.
const (
	OpenFGAParserRULE_main = 0
	OpenFGAParserRULE_modelHeaderComment = 1
	OpenFGAParserRULE_modelHeader = 2
	OpenFGAParserRULE_typeDefs = 3
	OpenFGAParserRULE_typeDef = 4
	OpenFGAParserRULE_relationDeclaration = 5
	OpenFGAParserRULE_relationDef = 6
	OpenFGAParserRULE_relationDefPartials = 7
	OpenFGAParserRULE_relationDefPartialAllOr = 8
	OpenFGAParserRULE_relationDefPartialAllAnd = 9
	OpenFGAParserRULE_relationDefPartialAllButNot = 10
	OpenFGAParserRULE_relationDefDirectAssignment = 11
	OpenFGAParserRULE_relationDefRewrite = 12
	OpenFGAParserRULE_relationDefRelationOnSameObject = 13
	OpenFGAParserRULE_relationDefRelationOnRelatedObject = 14
	OpenFGAParserRULE_relationDefTypeRestriction = 15
	OpenFGAParserRULE_relationDefTypeRestrictionType = 16
	OpenFGAParserRULE_relationDefTypeRestrictionRelation = 17
	OpenFGAParserRULE_relationDefGrouping = 18
	OpenFGAParserRULE_rewriteComputedusersetName = 19
	OpenFGAParserRULE_rewriteTuplesetComputedusersetName = 20
	OpenFGAParserRULE_rewriteTuplesetName = 21
	OpenFGAParserRULE_relationName = 22
	OpenFGAParserRULE_typeName = 23
	OpenFGAParserRULE_schemaVersion = 24
	OpenFGAParserRULE_name = 25
)

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ModelHeader() IModelHeaderContext
	TypeDefs() ITypeDefsContext
	NEWLINES() antlr.TerminalNode

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

func (s *MainContext) NEWLINES() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINES, 0)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitMain(s)
	}
}




func (p *OpenFGAParser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OpenFGAParserRULE_main)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.ModelHeader()
	}
	{
		p.SetState(53)
		p.TypeDefs()
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserNEWLINES {
		{
			p.SetState(54)
			p.Match(OpenFGAParserNEWLINES)
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


// IModelHeaderCommentContext is an interface to support dynamic dispatch.
type IModelHeaderCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HASH() antlr.TerminalNode
	MULTILINE_COMMENT() antlr.TerminalNode

	// IsModelHeaderCommentContext differentiates from other interfaces.
	IsModelHeaderCommentContext()
}

type ModelHeaderCommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelHeaderCommentContext() *ModelHeaderCommentContext {
	var p = new(ModelHeaderCommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_modelHeaderComment
	return p
}

func InitEmptyModelHeaderCommentContext(p *ModelHeaderCommentContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_modelHeaderComment
}

func (*ModelHeaderCommentContext) IsModelHeaderCommentContext() {}

func NewModelHeaderCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelHeaderCommentContext {
	var p = new(ModelHeaderCommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_modelHeaderComment

	return p
}

func (s *ModelHeaderCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelHeaderCommentContext) HASH() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserHASH, 0)
}

func (s *ModelHeaderCommentContext) MULTILINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserMULTILINE_COMMENT, 0)
}

func (s *ModelHeaderCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelHeaderCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ModelHeaderCommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterModelHeaderComment(s)
	}
}

func (s *ModelHeaderCommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitModelHeaderComment(s)
	}
}




func (p *OpenFGAParser) ModelHeaderComment() (localctx IModelHeaderCommentContext) {
	localctx = NewModelHeaderCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OpenFGAParserRULE_modelHeaderComment)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserHASH {
		{
			p.SetState(57)
			p.Match(OpenFGAParserHASH)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(58)
					_la = p.GetTokenStream().LA(1)

					if _la <= 0 || _la == OpenFGAParserT__0 || _la == OpenFGAParserT__1  {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}


			}
			p.SetState(63)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserMULTILINE_COMMENT {
		{
			p.SetState(66)
			p.Match(OpenFGAParserMULTILINE_COMMENT)
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


// IModelHeaderContext is an interface to support dynamic dispatch.
type IModelHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MODEL() antlr.TerminalNode
	AllNEWLINES() []antlr.TerminalNode
	NEWLINES(i int) antlr.TerminalNode
	INDENT() antlr.TerminalNode
	SCHEMA() antlr.TerminalNode
	SchemaVersion() ISchemaVersionContext
	ModelHeaderComment() IModelHeaderCommentContext

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

func (s *ModelHeaderContext) MODEL() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserMODEL, 0)
}

func (s *ModelHeaderContext) AllNEWLINES() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNEWLINES)
}

func (s *ModelHeaderContext) NEWLINES(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINES, i)
}

func (s *ModelHeaderContext) INDENT() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserINDENT, 0)
}

func (s *ModelHeaderContext) SCHEMA() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserSCHEMA, 0)
}

func (s *ModelHeaderContext) SchemaVersion() ISchemaVersionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISchemaVersionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISchemaVersionContext)
}

func (s *ModelHeaderContext) ModelHeaderComment() IModelHeaderCommentContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModelHeaderCommentContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModelHeaderCommentContext)
}

func (s *ModelHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ModelHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterModelHeader(s)
	}
}

func (s *ModelHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitModelHeader(s)
	}
}




func (p *OpenFGAParser) ModelHeader() (localctx IModelHeaderContext) {
	localctx = NewModelHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, OpenFGAParserRULE_modelHeader)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 16842760) != 0) {
		{
			p.SetState(69)
			p.ModelHeaderComment()
		}
		{
			p.SetState(70)
			p.Match(OpenFGAParserNEWLINES)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(74)
		p.Match(OpenFGAParserMODEL)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Match(OpenFGAParserNEWLINES)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(OpenFGAParserINDENT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(77)
		p.Match(OpenFGAParserSCHEMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(78)
		p.SchemaVersion()
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterTypeDefs(s)
	}
}

func (s *TypeDefsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitTypeDefs(s)
	}
}




func (p *OpenFGAParser) TypeDefs() (localctx ITypeDefsContext) {
	localctx = NewTypeDefsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, OpenFGAParserRULE_typeDefs)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(83)
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
				p.SetState(80)
				p.TypeDef()
			}


		}
		p.SetState(85)
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
	AllNEWLINES() []antlr.TerminalNode
	NEWLINES(i int) antlr.TerminalNode
	TYPE() antlr.TerminalNode
	TypeName() ITypeNameContext
	MULTILINE_COMMENT() antlr.TerminalNode
	INDENT() antlr.TerminalNode
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

func (s *TypeDefContext) AllNEWLINES() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserNEWLINES)
}

func (s *TypeDefContext) NEWLINES(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINES, i)
}

func (s *TypeDefContext) TYPE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserTYPE, 0)
}

func (s *TypeDefContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *TypeDefContext) MULTILINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserMULTILINE_COMMENT, 0)
}

func (s *TypeDefContext) INDENT() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserINDENT, 0)
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterTypeDef(s)
	}
}

func (s *TypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitTypeDef(s)
	}
}




func (p *OpenFGAParser) TypeDef() (localctx ITypeDefContext) {
	localctx = NewTypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, OpenFGAParserRULE_typeDef)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserMULTILINE_COMMENT {
		{
			p.SetState(86)
			p.Match(OpenFGAParserMULTILINE_COMMENT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(89)
		p.Match(OpenFGAParserNEWLINES)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(90)
		p.Match(OpenFGAParserTYPE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(91)
		p.TypeName()
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(92)
			p.Match(OpenFGAParserNEWLINES)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Match(OpenFGAParserINDENT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Match(OpenFGAParserRELATIONS)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
					{
						p.SetState(95)
						p.RelationDeclaration()
					}




			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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
	NEWLINES() antlr.TerminalNode
	AllINDENT() []antlr.TerminalNode
	INDENT(i int) antlr.TerminalNode
	DEFINE() antlr.TerminalNode
	RelationName() IRelationNameContext
	COLON() antlr.TerminalNode
	RelationDef() IRelationDefContext
	MULTILINE_COMMENT() antlr.TerminalNode

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

func (s *RelationDeclarationContext) NEWLINES() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserNEWLINES, 0)
}

func (s *RelationDeclarationContext) AllINDENT() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserINDENT)
}

func (s *RelationDeclarationContext) INDENT(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserINDENT, i)
}

func (s *RelationDeclarationContext) DEFINE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserDEFINE, 0)
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

func (s *RelationDeclarationContext) MULTILINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserMULTILINE_COMMENT, 0)
}

func (s *RelationDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RelationDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationDeclaration(s)
	}
}

func (s *RelationDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelationDeclaration(s)
	}
}




func (p *OpenFGAParser) RelationDeclaration() (localctx IRelationDeclarationContext) {
	localctx = NewRelationDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, OpenFGAParserRULE_relationDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == OpenFGAParserMULTILINE_COMMENT {
		{
			p.SetState(102)
			p.Match(OpenFGAParserMULTILINE_COMMENT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(105)
		p.Match(OpenFGAParserNEWLINES)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(OpenFGAParserINDENT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(OpenFGAParserINDENT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(OpenFGAParserDEFINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(109)
		p.RelationName()
	}
	{
		p.SetState(110)
		p.Match(OpenFGAParserCOLON)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(111)
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationDef(s)
	}
}

func (s *RelationDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
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
	case OpenFGAParserL_SQUARE:
		{
			p.SetState(113)
			p.RelationDefDirectAssignment()
		}


	case OpenFGAParserALPHA_NUMERIC:
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


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 14336) != 0) {
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

func (s *RelationDefPartialsContext) RelationDefPartialAllOr() IRelationDefPartialAllOrContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefPartialAllOrContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefPartialAllOrContext)
}

func (s *RelationDefPartialsContext) RelationDefPartialAllAnd() IRelationDefPartialAllAndContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefPartialAllAndContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefPartialAllAndContext)
}

func (s *RelationDefPartialsContext) RelationDefPartialAllButNot() IRelationDefPartialAllButNotContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefPartialAllButNotContext); ok {
			t = ctx.(antlr.RuleContext);
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationDefPartials(s)
	}
}

func (s *RelationDefPartialsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelationDefPartials(s)
	}
}




func (p *OpenFGAParser) RelationDefPartials() (localctx IRelationDefPartialsContext) {
	localctx = NewRelationDefPartialsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, OpenFGAParserRULE_relationDefPartials)
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case OpenFGAParserOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.RelationDefPartialAllOr()
		}


	case OpenFGAParserAND:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.RelationDefPartialAllAnd()
		}


	case OpenFGAParserBUT_NOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(122)
			p.RelationDefPartialAllButNot()
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


// IRelationDefPartialAllOrContext is an interface to support dynamic dispatch.
type IRelationDefPartialAllOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode
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

func InitEmptyRelationDefPartialAllOrContext(p *RelationDefPartialAllOrContext)  {
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

func (s *RelationDefPartialAllOrContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserOR)
}

func (s *RelationDefPartialAllOrContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserOR, i)
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

func (s *RelationDefPartialAllOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefPartialAllOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RelationDefPartialAllOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationDefPartialAllOr(s)
	}
}

func (s *RelationDefPartialAllOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelationDefPartialAllOr(s)
	}
}




func (p *OpenFGAParser) RelationDefPartialAllOr() (localctx IRelationDefPartialAllOrContext) {
	localctx = NewRelationDefPartialAllOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, OpenFGAParserRULE_relationDefPartialAllOr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == OpenFGAParserOR {
		{
			p.SetState(125)
			p.Match(OpenFGAParserOR)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(126)
			p.RelationDefGrouping()
		}


		p.SetState(129)
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


// IRelationDefPartialAllAndContext is an interface to support dynamic dispatch.
type IRelationDefPartialAllAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode
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

func InitEmptyRelationDefPartialAllAndContext(p *RelationDefPartialAllAndContext)  {
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

func (s *RelationDefPartialAllAndContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserAND)
}

func (s *RelationDefPartialAllAndContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserAND, i)
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

func (s *RelationDefPartialAllAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefPartialAllAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RelationDefPartialAllAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationDefPartialAllAnd(s)
	}
}

func (s *RelationDefPartialAllAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelationDefPartialAllAnd(s)
	}
}




func (p *OpenFGAParser) RelationDefPartialAllAnd() (localctx IRelationDefPartialAllAndContext) {
	localctx = NewRelationDefPartialAllAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, OpenFGAParserRULE_relationDefPartialAllAnd)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == OpenFGAParserAND {
		{
			p.SetState(131)
			p.Match(OpenFGAParserAND)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(132)
			p.RelationDefGrouping()
		}


		p.SetState(135)
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


// IRelationDefPartialAllButNotContext is an interface to support dynamic dispatch.
type IRelationDefPartialAllButNotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBUT_NOT() []antlr.TerminalNode
	BUT_NOT(i int) antlr.TerminalNode
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

func InitEmptyRelationDefPartialAllButNotContext(p *RelationDefPartialAllButNotContext)  {
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

func (s *RelationDefPartialAllButNotContext) AllBUT_NOT() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserBUT_NOT)
}

func (s *RelationDefPartialAllButNotContext) BUT_NOT(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserBUT_NOT, i)
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

func (s *RelationDefPartialAllButNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefPartialAllButNotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RelationDefPartialAllButNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationDefPartialAllButNot(s)
	}
}

func (s *RelationDefPartialAllButNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelationDefPartialAllButNot(s)
	}
}




func (p *OpenFGAParser) RelationDefPartialAllButNot() (localctx IRelationDefPartialAllButNotContext) {
	localctx = NewRelationDefPartialAllButNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, OpenFGAParserRULE_relationDefPartialAllButNot)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == OpenFGAParserBUT_NOT {
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
			p.RelationDefGrouping()
		}


		p.SetState(141)
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

func (s *RelationDefDirectAssignmentContext) R_SQUARE() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserR_SQUARE, 0)
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationDefDirectAssignment(s)
	}
}

func (s *RelationDefDirectAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelationDefDirectAssignment(s)
	}
}




func (p *OpenFGAParser) RelationDefDirectAssignment() (localctx IRelationDefDirectAssignmentContext) {
	localctx = NewRelationDefDirectAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, OpenFGAParserRULE_relationDefDirectAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(OpenFGAParserL_SQUARE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(144)
		p.RelationDefTypeRestriction()
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == OpenFGAParserCOMMA {
		{
			p.SetState(145)
			p.Match(OpenFGAParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(146)
			p.RelationDefTypeRestriction()
		}


		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(152)
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

func (s *RelationDefRewriteContext) RelationDefRelationOnSameObject() IRelationDefRelationOnSameObjectContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefRelationOnSameObjectContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefRelationOnSameObjectContext)
}

func (s *RelationDefRewriteContext) RelationDefRelationOnRelatedObject() IRelationDefRelationOnRelatedObjectContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefRelationOnRelatedObjectContext); ok {
			t = ctx.(antlr.RuleContext);
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationDefRewrite(s)
	}
}

func (s *RelationDefRewriteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelationDefRewrite(s)
	}
}




func (p *OpenFGAParser) RelationDefRewrite() (localctx IRelationDefRewriteContext) {
	localctx = NewRelationDefRewriteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, OpenFGAParserRULE_relationDefRewrite)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)
			p.RelationDefRelationOnSameObject()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
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

func InitEmptyRelationDefRelationOnSameObjectContext(p *RelationDefRelationOnSameObjectContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRewriteComputedusersetNameContext); ok {
			t = ctx.(antlr.RuleContext);
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationDefRelationOnSameObject(s)
	}
}

func (s *RelationDefRelationOnSameObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelationDefRelationOnSameObject(s)
	}
}




func (p *OpenFGAParser) RelationDefRelationOnSameObject() (localctx IRelationDefRelationOnSameObjectContext) {
	localctx = NewRelationDefRelationOnSameObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, OpenFGAParserRULE_relationDefRelationOnSameObject)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
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
	FROM() antlr.TerminalNode
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

func InitEmptyRelationDefRelationOnRelatedObjectContext(p *RelationDefRelationOnRelatedObjectContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRewriteTuplesetComputedusersetNameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRewriteTuplesetComputedusersetNameContext)
}

func (s *RelationDefRelationOnRelatedObjectContext) FROM() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserFROM, 0)
}

func (s *RelationDefRelationOnRelatedObjectContext) RewriteTuplesetName() IRewriteTuplesetNameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRewriteTuplesetNameContext); ok {
			t = ctx.(antlr.RuleContext);
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationDefRelationOnRelatedObject(s)
	}
}

func (s *RelationDefRelationOnRelatedObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelationDefRelationOnRelatedObject(s)
	}
}




func (p *OpenFGAParser) RelationDefRelationOnRelatedObject() (localctx IRelationDefRelationOnRelatedObjectContext) {
	localctx = NewRelationDefRelationOnRelatedObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, OpenFGAParserRULE_relationDefRelationOnRelatedObject)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.RewriteTuplesetComputedusersetName()
	}
	{
		p.SetState(161)
		p.Match(OpenFGAParserFROM)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(162)
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


// IRelationDefTypeRestrictionContext is an interface to support dynamic dispatch.
type IRelationDefTypeRestrictionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationDefTypeRestrictionType() IRelationDefTypeRestrictionTypeContext
	COLON() antlr.TerminalNode
	WILDCARD() antlr.TerminalNode
	HASH() antlr.TerminalNode
	RelationDefTypeRestrictionRelation() IRelationDefTypeRestrictionRelationContext

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

func (s *RelationDefTypeRestrictionContext) RelationDefTypeRestrictionType() IRelationDefTypeRestrictionTypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefTypeRestrictionTypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefTypeRestrictionTypeContext)
}

func (s *RelationDefTypeRestrictionContext) COLON() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserCOLON, 0)
}

func (s *RelationDefTypeRestrictionContext) WILDCARD() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserWILDCARD, 0)
}

func (s *RelationDefTypeRestrictionContext) HASH() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserHASH, 0)
}

func (s *RelationDefTypeRestrictionContext) RelationDefTypeRestrictionRelation() IRelationDefTypeRestrictionRelationContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationDefTypeRestrictionRelationContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationDefTypeRestrictionRelationContext)
}

func (s *RelationDefTypeRestrictionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationDefTypeRestrictionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RelationDefTypeRestrictionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationDefTypeRestriction(s)
	}
}

func (s *RelationDefTypeRestrictionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelationDefTypeRestriction(s)
	}
}




func (p *OpenFGAParser) RelationDefTypeRestriction() (localctx IRelationDefTypeRestrictionContext) {
	localctx = NewRelationDefTypeRestrictionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, OpenFGAParserRULE_relationDefTypeRestriction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.RelationDefTypeRestrictionType()
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case OpenFGAParserCOLON:
		{
			p.SetState(165)
			p.Match(OpenFGAParserCOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Match(OpenFGAParserWILDCARD)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	case OpenFGAParserHASH:
		{
			p.SetState(167)
			p.Match(OpenFGAParserHASH)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(168)
			p.RelationDefTypeRestrictionRelation()
		}



	case OpenFGAParserR_SQUARE, OpenFGAParserCOMMA:



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

func InitEmptyRelationDefTypeRestrictionTypeContext(p *RelationDefTypeRestrictionTypeContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationDefTypeRestrictionType(s)
	}
}

func (s *RelationDefTypeRestrictionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelationDefTypeRestrictionType(s)
	}
}




func (p *OpenFGAParser) RelationDefTypeRestrictionType() (localctx IRelationDefTypeRestrictionTypeContext) {
	localctx = NewRelationDefTypeRestrictionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, OpenFGAParserRULE_relationDefTypeRestrictionType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
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

func InitEmptyRelationDefTypeRestrictionRelationContext(p *RelationDefTypeRestrictionRelationContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationDefTypeRestrictionRelation(s)
	}
}

func (s *RelationDefTypeRestrictionRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelationDefTypeRestrictionRelation(s)
	}
}




func (p *OpenFGAParser) RelationDefTypeRestrictionRelation() (localctx IRelationDefTypeRestrictionRelationContext) {
	localctx = NewRelationDefTypeRestrictionRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, OpenFGAParserRULE_relationDefTypeRestrictionRelation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationDefGrouping(s)
	}
}

func (s *RelationDefGroupingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelationDefGrouping(s)
	}
}




func (p *OpenFGAParser) RelationDefGrouping() (localctx IRelationDefGroupingContext) {
	localctx = NewRelationDefGroupingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, OpenFGAParserRULE_relationDefGrouping)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
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

func InitEmptyRewriteComputedusersetNameContext(p *RewriteComputedusersetNameContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRewriteComputedusersetName(s)
	}
}

func (s *RewriteComputedusersetNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRewriteComputedusersetName(s)
	}
}




func (p *OpenFGAParser) RewriteComputedusersetName() (localctx IRewriteComputedusersetNameContext) {
	localctx = NewRewriteComputedusersetNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, OpenFGAParserRULE_rewriteComputedusersetName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
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

func InitEmptyRewriteTuplesetComputedusersetNameContext(p *RewriteTuplesetComputedusersetNameContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRewriteTuplesetComputedusersetName(s)
	}
}

func (s *RewriteTuplesetComputedusersetNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRewriteTuplesetComputedusersetName(s)
	}
}




func (p *OpenFGAParser) RewriteTuplesetComputedusersetName() (localctx IRewriteTuplesetComputedusersetNameContext) {
	localctx = NewRewriteTuplesetComputedusersetNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, OpenFGAParserRULE_rewriteTuplesetComputedusersetName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
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

func InitEmptyRewriteTuplesetNameContext(p *RewriteTuplesetNameContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRewriteTuplesetName(s)
	}
}

func (s *RewriteTuplesetNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRewriteTuplesetName(s)
	}
}




func (p *OpenFGAParser) RewriteTuplesetName() (localctx IRewriteTuplesetNameContext) {
	localctx = NewRewriteTuplesetNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, OpenFGAParserRULE_rewriteTuplesetName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
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

func (s *RelationNameContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationName(s)
	}
}

func (s *RelationNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelationName(s)
	}
}




func (p *OpenFGAParser) RelationName() (localctx IRelationNameContext) {
	localctx = NewRelationNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, OpenFGAParserRULE_relationName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
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

func InitEmptyTypeNameContext(p *TypeNameContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitTypeName(s)
	}
}




func (p *OpenFGAParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, OpenFGAParserRULE_typeName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
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

func InitEmptySchemaVersionContext(p *SchemaVersionContext)  {
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterSchemaVersion(s)
	}
}

func (s *SchemaVersionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitSchemaVersion(s)
	}
}




func (p *OpenFGAParser) SchemaVersion() (localctx ISchemaVersionContext) {
	localctx = NewSchemaVersionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, OpenFGAParserRULE_schemaVersion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
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

func InitEmptyNameContext(p *NameContext)  {
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitName(s)
	}
}




func (p *OpenFGAParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, OpenFGAParserRULE_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == OpenFGAParserALPHA_NUMERIC {
		{
			p.SetState(189)
			p.Match(OpenFGAParserALPHA_NUMERIC)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(192)
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


