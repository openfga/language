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
	public static readonly T__0 = 1;
	public static readonly T__1 = 2;
	public static readonly MULTILINE_COMMENT = 3;
	public static readonly INDENT = 4;
	public static readonly MODEL = 5;
	public static readonly TYPE = 6;
	public static readonly SCHEMA = 7;
	public static readonly SCHEMA_VERSION = 8;
	public static readonly RELATIONS = 9;
	public static readonly DEFINE = 10;
	public static readonly AND = 11;
	public static readonly OR = 12;
	public static readonly BUT_NOT = 13;
	public static readonly FROM = 14;
	public static readonly COLON = 15;
	public static readonly HASH = 16;
	public static readonly WILDCARD = 17;
	public static readonly L_SQUARE = 18;
	public static readonly R_SQUARE = 19;
	public static readonly COMMA = 20;
	public static readonly SYMBOL = 21;
	public static readonly ALPHA_NUMERIC_CHAR = 22;
	public static readonly ALPHA_NUMERIC = 23;
	public static readonly NEWLINES = 24;
	public static readonly WS = 25;
	public static readonly EOF = Token.EOF;

	public static readonly channelNames: string[] = [ "DEFAULT_TOKEN_CHANNEL", "HIDDEN" ];
	public static readonly literalNames: (string | null)[] = [ null, "'\\r'", 
                                                            "'\\n'", null, 
                                                            null, "'model'", 
                                                            "'type'", "'schema'", 
                                                            "'1.1'", "'relations'", 
                                                            "'define'", 
                                                            "'and'", "'or'", 
                                                            "'but not'", 
                                                            "'from'", "':'", 
                                                            "'#'", "'*'", 
                                                            "'['", "']'", 
                                                            "','" ];
	public static readonly symbolicNames: (string | null)[] = [ null, null, 
                                                             null, "MULTILINE_COMMENT", 
                                                             "INDENT", "MODEL", 
                                                             "TYPE", "SCHEMA", 
                                                             "SCHEMA_VERSION", 
                                                             "RELATIONS", 
                                                             "DEFINE", "AND", 
                                                             "OR", "BUT_NOT", 
                                                             "FROM", "COLON", 
                                                             "HASH", "WILDCARD", 
                                                             "L_SQUARE", 
                                                             "R_SQUARE", 
                                                             "COMMA", "SYMBOL", 
                                                             "ALPHA_NUMERIC_CHAR", 
                                                             "ALPHA_NUMERIC", 
                                                             "NEWLINES", 
                                                             "WS" ];
	public static readonly modeNames: string[] = [ "DEFAULT_MODE", ];

	public static readonly ruleNames: string[] = [
		"T__0", "T__1", "MULTILINE_COMMENT", "INDENT", "MODEL", "TYPE", "SCHEMA", 
		"SCHEMA_VERSION", "RELATIONS", "DEFINE", "AND", "OR", "BUT_NOT", "FROM", 
		"COLON", "HASH", "WILDCARD", "L_SQUARE", "R_SQUARE", "COMMA", "SYMBOL", 
		"ALPHA_NUMERIC_CHAR", "ALPHA_NUMERIC", "COMMENT_CONTENTS", "NEWLINES", 
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

	public static readonly _serializedATN: number[] = [4,0,25,179,6,-1,2,0,
	7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,
	7,9,2,10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,
	16,2,17,7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,2,22,7,22,2,23,7,23,
	2,24,7,24,2,25,7,25,2,26,7,26,1,0,1,0,1,1,1,1,1,2,1,2,5,2,62,8,2,10,2,12,
	2,65,9,2,1,2,1,2,1,2,4,2,70,8,2,11,2,12,2,71,1,3,1,3,1,3,3,3,77,8,3,1,4,
	1,4,1,4,1,4,1,4,1,4,1,5,1,5,1,5,1,5,1,5,1,6,1,6,1,6,1,6,1,6,1,6,1,6,1,7,
	1,7,1,7,1,7,1,8,1,8,1,8,1,8,1,8,1,8,1,8,1,8,1,8,1,8,1,9,1,9,1,9,1,9,1,9,
	1,9,1,9,1,10,1,10,1,10,1,10,1,11,1,11,1,11,1,12,1,12,1,12,1,12,1,12,1,12,
	1,12,1,12,1,13,1,13,1,13,1,13,1,13,1,14,1,14,1,15,1,15,1,16,1,16,1,17,1,
	17,1,18,1,18,1,19,1,19,1,20,3,20,151,8,20,1,21,1,21,1,22,4,22,156,8,22,
	11,22,12,22,157,1,23,5,23,161,8,23,10,23,12,23,164,9,23,1,24,4,24,167,8,
	24,11,24,12,24,168,1,25,1,25,1,25,3,25,174,8,25,1,26,1,26,1,26,1,26,0,0,
	27,1,1,3,2,5,3,7,4,9,5,11,6,13,7,15,8,17,9,19,10,21,11,23,12,25,13,27,14,
	29,15,31,16,33,17,35,18,37,19,39,20,41,21,43,22,45,23,47,0,49,24,51,0,53,
	25,1,0,5,6,0,33,47,58,62,64,64,91,96,123,123,125,126,5,0,45,45,48,57,65,
	90,95,95,97,122,3,0,10,10,13,13,8232,8233,2,0,10,10,13,13,2,0,9,9,32,32,
	183,0,1,1,0,0,0,0,3,1,0,0,0,0,5,1,0,0,0,0,7,1,0,0,0,0,9,1,0,0,0,0,11,1,
	0,0,0,0,13,1,0,0,0,0,15,1,0,0,0,0,17,1,0,0,0,0,19,1,0,0,0,0,21,1,0,0,0,
	0,23,1,0,0,0,0,25,1,0,0,0,0,27,1,0,0,0,0,29,1,0,0,0,0,31,1,0,0,0,0,33,1,
	0,0,0,0,35,1,0,0,0,0,37,1,0,0,0,0,39,1,0,0,0,0,41,1,0,0,0,0,43,1,0,0,0,
	0,45,1,0,0,0,0,49,1,0,0,0,0,53,1,0,0,0,1,55,1,0,0,0,3,57,1,0,0,0,5,69,1,
	0,0,0,7,76,1,0,0,0,9,78,1,0,0,0,11,84,1,0,0,0,13,89,1,0,0,0,15,96,1,0,0,
	0,17,100,1,0,0,0,19,110,1,0,0,0,21,117,1,0,0,0,23,121,1,0,0,0,25,124,1,
	0,0,0,27,132,1,0,0,0,29,137,1,0,0,0,31,139,1,0,0,0,33,141,1,0,0,0,35,143,
	1,0,0,0,37,145,1,0,0,0,39,147,1,0,0,0,41,150,1,0,0,0,43,152,1,0,0,0,45,
	155,1,0,0,0,47,162,1,0,0,0,49,166,1,0,0,0,51,173,1,0,0,0,53,175,1,0,0,0,
	55,56,5,13,0,0,56,2,1,0,0,0,57,58,5,10,0,0,58,4,1,0,0,0,59,63,3,49,24,0,
	60,62,3,53,26,0,61,60,1,0,0,0,62,65,1,0,0,0,63,61,1,0,0,0,63,64,1,0,0,0,
	64,66,1,0,0,0,65,63,1,0,0,0,66,67,3,31,15,0,67,68,3,47,23,0,68,70,1,0,0,
	0,69,59,1,0,0,0,70,71,1,0,0,0,71,69,1,0,0,0,71,72,1,0,0,0,72,6,1,0,0,0,
	73,74,5,32,0,0,74,77,5,32,0,0,75,77,5,9,0,0,76,73,1,0,0,0,76,75,1,0,0,0,
	77,8,1,0,0,0,78,79,5,109,0,0,79,80,5,111,0,0,80,81,5,100,0,0,81,82,5,101,
	0,0,82,83,5,108,0,0,83,10,1,0,0,0,84,85,5,116,0,0,85,86,5,121,0,0,86,87,
	5,112,0,0,87,88,5,101,0,0,88,12,1,0,0,0,89,90,5,115,0,0,90,91,5,99,0,0,
	91,92,5,104,0,0,92,93,5,101,0,0,93,94,5,109,0,0,94,95,5,97,0,0,95,14,1,
	0,0,0,96,97,5,49,0,0,97,98,5,46,0,0,98,99,5,49,0,0,99,16,1,0,0,0,100,101,
	5,114,0,0,101,102,5,101,0,0,102,103,5,108,0,0,103,104,5,97,0,0,104,105,
	5,116,0,0,105,106,5,105,0,0,106,107,5,111,0,0,107,108,5,110,0,0,108,109,
	5,115,0,0,109,18,1,0,0,0,110,111,5,100,0,0,111,112,5,101,0,0,112,113,5,
	102,0,0,113,114,5,105,0,0,114,115,5,110,0,0,115,116,5,101,0,0,116,20,1,
	0,0,0,117,118,5,97,0,0,118,119,5,110,0,0,119,120,5,100,0,0,120,22,1,0,0,
	0,121,122,5,111,0,0,122,123,5,114,0,0,123,24,1,0,0,0,124,125,5,98,0,0,125,
	126,5,117,0,0,126,127,5,116,0,0,127,128,5,32,0,0,128,129,5,110,0,0,129,
	130,5,111,0,0,130,131,5,116,0,0,131,26,1,0,0,0,132,133,5,102,0,0,133,134,
	5,114,0,0,134,135,5,111,0,0,135,136,5,109,0,0,136,28,1,0,0,0,137,138,5,
	58,0,0,138,30,1,0,0,0,139,140,5,35,0,0,140,32,1,0,0,0,141,142,5,42,0,0,
	142,34,1,0,0,0,143,144,5,91,0,0,144,36,1,0,0,0,145,146,5,93,0,0,146,38,
	1,0,0,0,147,148,5,44,0,0,148,40,1,0,0,0,149,151,7,0,0,0,150,149,1,0,0,0,
	151,42,1,0,0,0,152,153,7,1,0,0,153,44,1,0,0,0,154,156,3,43,21,0,155,154,
	1,0,0,0,156,157,1,0,0,0,157,155,1,0,0,0,157,158,1,0,0,0,158,46,1,0,0,0,
	159,161,8,2,0,0,160,159,1,0,0,0,161,164,1,0,0,0,162,160,1,0,0,0,162,163,
	1,0,0,0,163,48,1,0,0,0,164,162,1,0,0,0,165,167,3,51,25,0,166,165,1,0,0,
	0,167,168,1,0,0,0,168,166,1,0,0,0,168,169,1,0,0,0,169,50,1,0,0,0,170,171,
	5,13,0,0,171,174,5,10,0,0,172,174,7,3,0,0,173,170,1,0,0,0,173,172,1,0,0,
	0,174,52,1,0,0,0,175,176,7,4,0,0,176,177,1,0,0,0,177,178,6,26,0,0,178,54,
	1,0,0,0,9,0,63,71,76,150,157,162,168,173,1,0,1,0];

	private static __ATN: ATN;
	public static get _ATN(): ATN {
		if (!OpenFGALexer.__ATN) {
			OpenFGALexer.__ATN = new ATNDeserializer().deserialize(OpenFGALexer._serializedATN);
		}

		return OpenFGALexer.__ATN;
	}


	static DecisionsToDFA = OpenFGALexer._ATN.decisionToState.map( (ds: DecisionState, index: number) => new DFA(ds, index) );
}