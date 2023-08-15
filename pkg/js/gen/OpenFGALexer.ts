// Generated from /app/OpenFGALexer.g4 by ANTLR 4.13.0
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
	public static readonly SCHEMA = 3;
	public static readonly SCHEMA_VERSION = 4;
	public static readonly TYPE = 5;
	public static readonly RELATIONS = 6;
	public static readonly DEFINE = 7;
	public static readonly HASH = 8;
	public static readonly COLON = 9;
	public static readonly WILDCARD = 10;
	public static readonly L_SQUARE = 11;
	public static readonly R_SQUARE = 12;
	public static readonly COMMA = 13;
	public static readonly AND = 14;
	public static readonly OR = 15;
	public static readonly BUT_NOT = 16;
	public static readonly FROM = 17;
	public static readonly ALPHA_NUMERIC = 18;
	public static readonly NEWLINE = 19;
	public static readonly WS = 20;
	public static readonly EOF = Token.EOF;

	public static readonly channelNames: string[] = [ "DEFAULT_TOKEN_CHANNEL", "HIDDEN" ];
	public static readonly literalNames: (string | null)[] = [ null, null, 
                                                            "'model'", "'schema'", 
                                                            "'1.1'", "'type'", 
                                                            "'relations'", 
                                                            "'define'", 
                                                            "'#'", "':'", 
                                                            "'*'", "'['", 
                                                            "']'", "','", 
                                                            "'and'", "'or'", 
                                                            "'but not'", 
                                                            "'from'" ];
	public static readonly symbolicNames: (string | null)[] = [ null, "INDENT", 
                                                             "MODEL", "SCHEMA", 
                                                             "SCHEMA_VERSION", 
                                                             "TYPE", "RELATIONS", 
                                                             "DEFINE", "HASH", 
                                                             "COLON", "WILDCARD", 
                                                             "L_SQUARE", 
                                                             "R_SQUARE", 
                                                             "COMMA", "AND", 
                                                             "OR", "BUT_NOT", 
                                                             "FROM", "ALPHA_NUMERIC", 
                                                             "NEWLINE", 
                                                             "WS" ];
	public static readonly modeNames: string[] = [ "DEFAULT_MODE", ];

	public static readonly ruleNames: string[] = [
		"SINGLE_INDENT", "DOUBLE_INDENT", "BOL", "INDENT", "MODEL", "SCHEMA", 
		"SCHEMA_VERSION", "TYPE", "RELATIONS", "DEFINE", "HASH", "COLON", "WILDCARD", 
		"L_SQUARE", "R_SQUARE", "COMMA", "AND", "OR", "BUT_NOT", "FROM", "ALPHA_NUMERIC", 
		"NEWLINE", "WS",
	];


	constructor(input: CharStream) {
		super(input);
		this._interp = new LexerATNSimulator(this, OpenFGALexer._ATN, OpenFGALexer.DecisionsToDFA, new PredictionContextCache());
	}

	public get grammarFileName(): string { return "OpenFGALexer.g4"; }

	public get literalNames(): (string | null)[] { return OpenFGALexer.literalNames; }
	public get symbolicNames(): (string | null)[] { return OpenFGALexer.symbolicNames; }
	public get ruleNames(): string[] { return OpenFGALexer.ruleNames; }

	public get serializedATN(): number[] { return OpenFGALexer._serializedATN; }

	public get channelNames(): string[] { return OpenFGALexer.channelNames; }

	public get modeNames(): string[] { return OpenFGALexer.modeNames; }

	public static readonly _serializedATN: number[] = [4,0,20,150,6,-1,2,0,
	7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,
	7,9,2,10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,
	16,2,17,7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,2,22,7,22,1,0,1,0,
	1,0,3,0,51,8,0,1,1,1,1,1,1,1,1,1,1,1,1,3,1,59,8,1,1,2,4,2,62,8,2,11,2,12,
	2,63,1,3,1,3,1,3,3,3,69,8,3,1,4,1,4,1,4,1,4,1,4,1,4,1,5,1,5,1,5,1,5,1,5,
	1,5,1,5,1,6,1,6,1,6,1,6,1,7,1,7,1,7,1,7,1,7,1,8,1,8,1,8,1,8,1,8,1,8,1,8,
	1,8,1,8,1,8,1,9,1,9,1,9,1,9,1,9,1,9,1,9,1,10,1,10,1,11,1,11,1,12,1,12,1,
	13,1,13,1,14,1,14,1,15,1,15,1,16,1,16,1,16,1,16,1,17,1,17,1,17,1,18,1,18,
	1,18,1,18,1,18,1,18,1,18,1,18,1,19,1,19,1,19,1,19,1,19,1,20,4,20,143,8,
	20,11,20,12,20,144,1,21,1,21,1,22,1,22,0,0,23,1,0,3,0,5,0,7,1,9,2,11,3,
	13,4,15,5,17,6,19,7,21,8,23,9,25,10,27,11,29,12,31,13,33,14,35,15,37,16,
	39,17,41,18,43,19,45,20,1,0,4,2,0,10,10,12,13,5,0,45,45,48,57,65,90,95,
	95,97,122,2,0,10,10,13,13,2,0,9,9,32,32,151,0,7,1,0,0,0,0,9,1,0,0,0,0,11,
	1,0,0,0,0,13,1,0,0,0,0,15,1,0,0,0,0,17,1,0,0,0,0,19,1,0,0,0,0,21,1,0,0,
	0,0,23,1,0,0,0,0,25,1,0,0,0,0,27,1,0,0,0,0,29,1,0,0,0,0,31,1,0,0,0,0,33,
	1,0,0,0,0,35,1,0,0,0,0,37,1,0,0,0,0,39,1,0,0,0,0,41,1,0,0,0,0,43,1,0,0,
	0,0,45,1,0,0,0,1,50,1,0,0,0,3,58,1,0,0,0,5,61,1,0,0,0,7,65,1,0,0,0,9,70,
	1,0,0,0,11,76,1,0,0,0,13,83,1,0,0,0,15,87,1,0,0,0,17,92,1,0,0,0,19,102,
	1,0,0,0,21,109,1,0,0,0,23,111,1,0,0,0,25,113,1,0,0,0,27,115,1,0,0,0,29,
	117,1,0,0,0,31,119,1,0,0,0,33,121,1,0,0,0,35,125,1,0,0,0,37,128,1,0,0,0,
	39,136,1,0,0,0,41,142,1,0,0,0,43,146,1,0,0,0,45,148,1,0,0,0,47,48,5,32,
	0,0,48,51,5,32,0,0,49,51,5,9,0,0,50,47,1,0,0,0,50,49,1,0,0,0,51,2,1,0,0,
	0,52,53,5,32,0,0,53,54,5,32,0,0,54,55,5,32,0,0,55,59,5,32,0,0,56,57,5,9,
	0,0,57,59,5,9,0,0,58,52,1,0,0,0,58,56,1,0,0,0,59,4,1,0,0,0,60,62,7,0,0,
	0,61,60,1,0,0,0,62,63,1,0,0,0,63,61,1,0,0,0,63,64,1,0,0,0,64,6,1,0,0,0,
	65,68,3,5,2,0,66,69,3,3,1,0,67,69,3,1,0,0,68,66,1,0,0,0,68,67,1,0,0,0,69,
	8,1,0,0,0,70,71,5,109,0,0,71,72,5,111,0,0,72,73,5,100,0,0,73,74,5,101,0,
	0,74,75,5,108,0,0,75,10,1,0,0,0,76,77,5,115,0,0,77,78,5,99,0,0,78,79,5,
	104,0,0,79,80,5,101,0,0,80,81,5,109,0,0,81,82,5,97,0,0,82,12,1,0,0,0,83,
	84,5,49,0,0,84,85,5,46,0,0,85,86,5,49,0,0,86,14,1,0,0,0,87,88,5,116,0,0,
	88,89,5,121,0,0,89,90,5,112,0,0,90,91,5,101,0,0,91,16,1,0,0,0,92,93,5,114,
	0,0,93,94,5,101,0,0,94,95,5,108,0,0,95,96,5,97,0,0,96,97,5,116,0,0,97,98,
	5,105,0,0,98,99,5,111,0,0,99,100,5,110,0,0,100,101,5,115,0,0,101,18,1,0,
	0,0,102,103,5,100,0,0,103,104,5,101,0,0,104,105,5,102,0,0,105,106,5,105,
	0,0,106,107,5,110,0,0,107,108,5,101,0,0,108,20,1,0,0,0,109,110,5,35,0,0,
	110,22,1,0,0,0,111,112,5,58,0,0,112,24,1,0,0,0,113,114,5,42,0,0,114,26,
	1,0,0,0,115,116,5,91,0,0,116,28,1,0,0,0,117,118,5,93,0,0,118,30,1,0,0,0,
	119,120,5,44,0,0,120,32,1,0,0,0,121,122,5,97,0,0,122,123,5,110,0,0,123,
	124,5,100,0,0,124,34,1,0,0,0,125,126,5,111,0,0,126,127,5,114,0,0,127,36,
	1,0,0,0,128,129,5,98,0,0,129,130,5,117,0,0,130,131,5,116,0,0,131,132,5,
	32,0,0,132,133,5,110,0,0,133,134,5,111,0,0,134,135,5,116,0,0,135,38,1,0,
	0,0,136,137,5,102,0,0,137,138,5,114,0,0,138,139,5,111,0,0,139,140,5,109,
	0,0,140,40,1,0,0,0,141,143,7,1,0,0,142,141,1,0,0,0,143,144,1,0,0,0,144,
	142,1,0,0,0,144,145,1,0,0,0,145,42,1,0,0,0,146,147,7,2,0,0,147,44,1,0,0,
	0,148,149,7,3,0,0,149,46,1,0,0,0,6,0,50,58,63,68,144,0];

	private static __ATN: ATN;
	public static get _ATN(): ATN {
		if (!OpenFGALexer.__ATN) {
			OpenFGALexer.__ATN = new ATNDeserializer().deserialize(OpenFGALexer._serializedATN);
		}

		return OpenFGALexer.__ATN;
	}


	static DecisionsToDFA = OpenFGALexer._ATN.decisionToState.map( (ds: DecisionState, index: number) => new DFA(ds, index) );
}