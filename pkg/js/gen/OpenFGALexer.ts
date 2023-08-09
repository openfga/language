// Generated from /app/OpenFGA.g4 by ANTLR 4.13.0
// noinspection ES6UnusedImports,JSUnusedGlobalSymbols,JSUnusedLocalSymbols
import {
	ATN,
	ATNDeserializer,
	CharStream,
	DecisionState, DFA,
	Lexer,
	LexerATNSimulator,
	RuleContext,
	PredictionContextCache,
	Token
} from "antlr4";
export default class OpenFGALexer extends Lexer {
	public static readonly INDENT = 1;
	public static readonly MODEL = 2;
	public static readonly TYPE = 3;
	public static readonly SCHEMA = 4;
	public static readonly SCHEMA_VERSION = 5;
	public static readonly RELATIONS = 6;
	public static readonly DEFINE = 7;
	public static readonly AND = 8;
	public static readonly OR = 9;
	public static readonly BUT_NOT = 10;
	public static readonly FROM = 11;
	public static readonly COLON = 12;
	public static readonly HASH = 13;
	public static readonly WILDCARD = 14;
	public static readonly L_SQUARE = 15;
	public static readonly R_SQUARE = 16;
	public static readonly COMMA = 17;
	public static readonly ALPHA_NUMERIC = 18;
	public static readonly NEWLINES = 19;
	public static readonly WS = 20;
	public static readonly EOF = Token.EOF;

	public static readonly channelNames: string[] = [ "DEFAULT_TOKEN_CHANNEL", "HIDDEN" ];
	public static readonly literalNames: (string | null)[] = [ null, null, 
                                                            "'model'", "'type'", 
                                                            "'schema'", 
                                                            null, "'relations'", 
                                                            "'define'", 
                                                            "'and'", "'or'", 
                                                            "'but not'", 
                                                            "'from'", "':'", 
                                                            "'#'", "':*'", 
                                                            "'['", "']'", 
                                                            "','" ];
	public static readonly symbolicNames: (string | null)[] = [ null, "INDENT", 
                                                             "MODEL", "TYPE", 
                                                             "SCHEMA", "SCHEMA_VERSION", 
                                                             "RELATIONS", 
                                                             "DEFINE", "AND", 
                                                             "OR", "BUT_NOT", 
                                                             "FROM", "COLON", 
                                                             "HASH", "WILDCARD", 
                                                             "L_SQUARE", 
                                                             "R_SQUARE", 
                                                             "COMMA", "ALPHA_NUMERIC", 
                                                             "NEWLINES", 
                                                             "WS" ];
	public static readonly modeNames: string[] = [ "DEFAULT_MODE", ];

	public static readonly ruleNames: string[] = [
		"INDENT", "MODEL", "TYPE", "SCHEMA", "SCHEMA_VERSION", "RELATIONS", "DEFINE", 
		"AND", "OR", "BUT_NOT", "FROM", "COLON", "HASH", "WILDCARD", "L_SQUARE", 
		"R_SQUARE", "COMMA", "ALPHA_NUMERIC", "COMMENT_CONTENTS", "NEWLINES", 
		"NEWLINE", "WS",
	];


	constructor(input: CharStream) {
		super(input);
		this._interp = new LexerATNSimulator(this, OpenFGALexer._ATN, OpenFGALexer.DecisionsToDFA, new PredictionContextCache());
	}

	public get grammarFileName(): string { return "OpenFGA.g4"; }

	public get literalNames(): (string | null)[] { return OpenFGALexer.literalNames; }
	public get symbolicNames(): (string | null)[] { return OpenFGALexer.symbolicNames; }
	public get ruleNames(): string[] { return OpenFGALexer.ruleNames; }

	public get serializedATN(): number[] { return OpenFGALexer._serializedATN; }

	public get channelNames(): string[] { return OpenFGALexer.channelNames; }

	public get modeNames(): string[] { return OpenFGALexer.modeNames; }

	public static readonly _serializedATN: number[] = [4,0,20,148,6,-1,2,0,
	7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,
	7,9,2,10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,
	16,2,17,7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,1,0,1,0,1,0,3,0,49,
	8,0,1,1,1,1,1,1,1,1,1,1,1,1,1,2,1,2,1,2,1,2,1,2,1,3,1,3,1,3,1,3,1,3,1,3,
	1,3,1,4,1,4,1,4,1,4,1,4,1,5,1,5,1,5,1,5,1,5,1,5,1,5,1,5,1,5,1,5,1,6,1,6,
	1,6,1,6,1,6,1,6,1,6,1,7,1,7,1,7,1,7,1,8,1,8,1,8,1,9,1,9,1,9,1,9,1,9,1,9,
	1,9,1,9,1,10,1,10,1,10,1,10,1,10,1,11,1,11,1,12,1,12,1,13,1,13,1,13,1,14,
	1,14,1,15,1,15,1,16,1,16,1,17,4,17,125,8,17,11,17,12,17,126,1,18,5,18,130,
	8,18,10,18,12,18,133,9,18,1,19,4,19,136,8,19,11,19,12,19,137,1,20,1,20,
	1,20,3,20,143,8,20,1,21,1,21,1,21,1,21,0,0,22,1,1,3,2,5,3,7,4,9,5,11,6,
	13,7,15,8,17,9,19,10,21,11,23,12,25,13,27,14,29,15,31,16,33,17,35,18,37,
	0,39,19,41,0,43,20,1,0,5,1,0,48,49,5,0,45,45,48,57,65,90,95,95,97,122,3,
	0,10,10,13,13,8232,8233,2,0,10,10,13,13,3,0,9,10,13,13,32,32,150,0,1,1,
	0,0,0,0,3,1,0,0,0,0,5,1,0,0,0,0,7,1,0,0,0,0,9,1,0,0,0,0,11,1,0,0,0,0,13,
	1,0,0,0,0,15,1,0,0,0,0,17,1,0,0,0,0,19,1,0,0,0,0,21,1,0,0,0,0,23,1,0,0,
	0,0,25,1,0,0,0,0,27,1,0,0,0,0,29,1,0,0,0,0,31,1,0,0,0,0,33,1,0,0,0,0,35,
	1,0,0,0,0,39,1,0,0,0,0,43,1,0,0,0,1,48,1,0,0,0,3,50,1,0,0,0,5,56,1,0,0,
	0,7,61,1,0,0,0,9,68,1,0,0,0,11,73,1,0,0,0,13,83,1,0,0,0,15,90,1,0,0,0,17,
	94,1,0,0,0,19,97,1,0,0,0,21,105,1,0,0,0,23,110,1,0,0,0,25,112,1,0,0,0,27,
	114,1,0,0,0,29,117,1,0,0,0,31,119,1,0,0,0,33,121,1,0,0,0,35,124,1,0,0,0,
	37,131,1,0,0,0,39,135,1,0,0,0,41,142,1,0,0,0,43,144,1,0,0,0,45,46,5,32,
	0,0,46,49,5,32,0,0,47,49,5,9,0,0,48,45,1,0,0,0,48,47,1,0,0,0,49,2,1,0,0,
	0,50,51,5,109,0,0,51,52,5,111,0,0,52,53,5,100,0,0,53,54,5,101,0,0,54,55,
	5,108,0,0,55,4,1,0,0,0,56,57,5,116,0,0,57,58,5,121,0,0,58,59,5,112,0,0,
	59,60,5,101,0,0,60,6,1,0,0,0,61,62,5,115,0,0,62,63,5,99,0,0,63,64,5,104,
	0,0,64,65,5,101,0,0,65,66,5,109,0,0,66,67,5,97,0,0,67,8,1,0,0,0,68,69,5,
	49,0,0,69,70,5,46,0,0,70,71,1,0,0,0,71,72,7,0,0,0,72,10,1,0,0,0,73,74,5,
	114,0,0,74,75,5,101,0,0,75,76,5,108,0,0,76,77,5,97,0,0,77,78,5,116,0,0,
	78,79,5,105,0,0,79,80,5,111,0,0,80,81,5,110,0,0,81,82,5,115,0,0,82,12,1,
	0,0,0,83,84,5,100,0,0,84,85,5,101,0,0,85,86,5,102,0,0,86,87,5,105,0,0,87,
	88,5,110,0,0,88,89,5,101,0,0,89,14,1,0,0,0,90,91,5,97,0,0,91,92,5,110,0,
	0,92,93,5,100,0,0,93,16,1,0,0,0,94,95,5,111,0,0,95,96,5,114,0,0,96,18,1,
	0,0,0,97,98,5,98,0,0,98,99,5,117,0,0,99,100,5,116,0,0,100,101,5,32,0,0,
	101,102,5,110,0,0,102,103,5,111,0,0,103,104,5,116,0,0,104,20,1,0,0,0,105,
	106,5,102,0,0,106,107,5,114,0,0,107,108,5,111,0,0,108,109,5,109,0,0,109,
	22,1,0,0,0,110,111,5,58,0,0,111,24,1,0,0,0,112,113,5,35,0,0,113,26,1,0,
	0,0,114,115,5,58,0,0,115,116,5,42,0,0,116,28,1,0,0,0,117,118,5,91,0,0,118,
	30,1,0,0,0,119,120,5,93,0,0,120,32,1,0,0,0,121,122,5,44,0,0,122,34,1,0,
	0,0,123,125,7,1,0,0,124,123,1,0,0,0,125,126,1,0,0,0,126,124,1,0,0,0,126,
	127,1,0,0,0,127,36,1,0,0,0,128,130,8,2,0,0,129,128,1,0,0,0,130,133,1,0,
	0,0,131,129,1,0,0,0,131,132,1,0,0,0,132,38,1,0,0,0,133,131,1,0,0,0,134,
	136,3,41,20,0,135,134,1,0,0,0,136,137,1,0,0,0,137,135,1,0,0,0,137,138,1,
	0,0,0,138,40,1,0,0,0,139,140,5,13,0,0,140,143,5,10,0,0,141,143,7,3,0,0,
	142,139,1,0,0,0,142,141,1,0,0,0,143,42,1,0,0,0,144,145,7,4,0,0,145,146,
	1,0,0,0,146,147,6,21,0,0,147,44,1,0,0,0,6,0,48,126,131,137,142,1,0,1,0];

	private static __ATN: ATN;
	public static get _ATN(): ATN {
		if (!OpenFGALexer.__ATN) {
			OpenFGALexer.__ATN = new ATNDeserializer().deserialize(OpenFGALexer._serializedATN);
		}

		return OpenFGALexer.__ATN;
	}


	static DecisionsToDFA = OpenFGALexer._ATN.decisionToState.map( (ds: DecisionState, index: number) => new DFA(ds, index) );
}