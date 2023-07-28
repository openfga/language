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
	public static readonly T__2 = 3;
	public static readonly T__3 = 4;
	public static readonly T__4 = 5;
	public static readonly T__5 = 6;
	public static readonly T__6 = 7;
	public static readonly T__7 = 8;
	public static readonly T__8 = 9;
	public static readonly T__9 = 10;
	public static readonly T__10 = 11;
	public static readonly T__11 = 12;
	public static readonly T__12 = 13;
	public static readonly T__13 = 14;
	public static readonly T__14 = 15;
	public static readonly T__15 = 16;
	public static readonly T__16 = 17;
	public static readonly T__17 = 18;
	public static readonly T__18 = 19;
	public static readonly T__19 = 20;
	public static readonly T__20 = 21;
	public static readonly T__21 = 22;
	public static readonly T__22 = 23;
	public static readonly WORD = 24;
	public static readonly EOF = Token.EOF;

	public static readonly channelNames: string[] = [ "DEFAULT_TOKEN_CHANNEL", "HIDDEN" ];
	public static readonly literalNames: (string | null)[] = [ null, "'  '", 
                                                            "'\\u0009'", 
                                                            "'model'", "'schema'", 
                                                            "'type'", "'relations'", 
                                                            "'define'", 
                                                            "':'", "'['", 
                                                            "','", "']'", 
                                                            "'and'", "'or'", 
                                                            "'but not'", 
                                                            "'from'", "':*'", 
                                                            "'#'", "'('", 
                                                            "')'", "'\\r'", 
                                                            "'\\n'", "' '", 
                                                            "'1.1'" ];
	public static readonly symbolicNames: (string | null)[] = [ null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             "WORD" ];
	public static readonly modeNames: string[] = [ "DEFAULT_MODE", ];

	public static readonly ruleNames: string[] = [
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "WORD",
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

	public static readonly _serializedATN: number[] = [4,0,24,141,6,-1,2,0,
	7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,
	7,9,2,10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,
	16,2,17,7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,2,22,7,22,2,23,7,23,
	1,0,1,0,1,0,1,1,1,1,1,2,1,2,1,2,1,2,1,2,1,2,1,3,1,3,1,3,1,3,1,3,1,3,1,3,
	1,4,1,4,1,4,1,4,1,4,1,5,1,5,1,5,1,5,1,5,1,5,1,5,1,5,1,5,1,5,1,6,1,6,1,6,
	1,6,1,6,1,6,1,6,1,7,1,7,1,8,1,8,1,9,1,9,1,10,1,10,1,11,1,11,1,11,1,11,1,
	12,1,12,1,12,1,13,1,13,1,13,1,13,1,13,1,13,1,13,1,13,1,14,1,14,1,14,1,14,
	1,14,1,15,1,15,1,15,1,16,1,16,1,17,1,17,1,18,1,18,1,19,1,19,1,20,1,20,1,
	21,1,21,1,22,1,22,1,22,1,22,1,23,4,23,138,8,23,11,23,12,23,139,0,0,24,1,
	1,3,2,5,3,7,4,9,5,11,6,13,7,15,8,17,9,19,10,21,11,23,12,25,13,27,14,29,
	15,31,16,33,17,35,18,37,19,39,20,41,21,43,22,45,23,47,24,1,0,1,5,0,45,45,
	48,57,65,90,95,95,97,122,141,0,1,1,0,0,0,0,3,1,0,0,0,0,5,1,0,0,0,0,7,1,
	0,0,0,0,9,1,0,0,0,0,11,1,0,0,0,0,13,1,0,0,0,0,15,1,0,0,0,0,17,1,0,0,0,0,
	19,1,0,0,0,0,21,1,0,0,0,0,23,1,0,0,0,0,25,1,0,0,0,0,27,1,0,0,0,0,29,1,0,
	0,0,0,31,1,0,0,0,0,33,1,0,0,0,0,35,1,0,0,0,0,37,1,0,0,0,0,39,1,0,0,0,0,
	41,1,0,0,0,0,43,1,0,0,0,0,45,1,0,0,0,0,47,1,0,0,0,1,49,1,0,0,0,3,52,1,0,
	0,0,5,54,1,0,0,0,7,60,1,0,0,0,9,67,1,0,0,0,11,72,1,0,0,0,13,82,1,0,0,0,
	15,89,1,0,0,0,17,91,1,0,0,0,19,93,1,0,0,0,21,95,1,0,0,0,23,97,1,0,0,0,25,
	101,1,0,0,0,27,104,1,0,0,0,29,112,1,0,0,0,31,117,1,0,0,0,33,120,1,0,0,0,
	35,122,1,0,0,0,37,124,1,0,0,0,39,126,1,0,0,0,41,128,1,0,0,0,43,130,1,0,
	0,0,45,132,1,0,0,0,47,137,1,0,0,0,49,50,5,32,0,0,50,51,5,32,0,0,51,2,1,
	0,0,0,52,53,5,9,0,0,53,4,1,0,0,0,54,55,5,109,0,0,55,56,5,111,0,0,56,57,
	5,100,0,0,57,58,5,101,0,0,58,59,5,108,0,0,59,6,1,0,0,0,60,61,5,115,0,0,
	61,62,5,99,0,0,62,63,5,104,0,0,63,64,5,101,0,0,64,65,5,109,0,0,65,66,5,
	97,0,0,66,8,1,0,0,0,67,68,5,116,0,0,68,69,5,121,0,0,69,70,5,112,0,0,70,
	71,5,101,0,0,71,10,1,0,0,0,72,73,5,114,0,0,73,74,5,101,0,0,74,75,5,108,
	0,0,75,76,5,97,0,0,76,77,5,116,0,0,77,78,5,105,0,0,78,79,5,111,0,0,79,80,
	5,110,0,0,80,81,5,115,0,0,81,12,1,0,0,0,82,83,5,100,0,0,83,84,5,101,0,0,
	84,85,5,102,0,0,85,86,5,105,0,0,86,87,5,110,0,0,87,88,5,101,0,0,88,14,1,
	0,0,0,89,90,5,58,0,0,90,16,1,0,0,0,91,92,5,91,0,0,92,18,1,0,0,0,93,94,5,
	44,0,0,94,20,1,0,0,0,95,96,5,93,0,0,96,22,1,0,0,0,97,98,5,97,0,0,98,99,
	5,110,0,0,99,100,5,100,0,0,100,24,1,0,0,0,101,102,5,111,0,0,102,103,5,114,
	0,0,103,26,1,0,0,0,104,105,5,98,0,0,105,106,5,117,0,0,106,107,5,116,0,0,
	107,108,5,32,0,0,108,109,5,110,0,0,109,110,5,111,0,0,110,111,5,116,0,0,
	111,28,1,0,0,0,112,113,5,102,0,0,113,114,5,114,0,0,114,115,5,111,0,0,115,
	116,5,109,0,0,116,30,1,0,0,0,117,118,5,58,0,0,118,119,5,42,0,0,119,32,1,
	0,0,0,120,121,5,35,0,0,121,34,1,0,0,0,122,123,5,40,0,0,123,36,1,0,0,0,124,
	125,5,41,0,0,125,38,1,0,0,0,126,127,5,13,0,0,127,40,1,0,0,0,128,129,5,10,
	0,0,129,42,1,0,0,0,130,131,5,32,0,0,131,44,1,0,0,0,132,133,5,49,0,0,133,
	134,5,46,0,0,134,135,5,49,0,0,135,46,1,0,0,0,136,138,7,0,0,0,137,136,1,
	0,0,0,138,139,1,0,0,0,139,137,1,0,0,0,139,140,1,0,0,0,140,48,1,0,0,0,2,
	0,139,0];

	private static __ATN: ATN;
	public static get _ATN(): ATN {
		if (!OpenFGALexer.__ATN) {
			OpenFGALexer.__ATN = new ATNDeserializer().deserialize(OpenFGALexer._serializedATN);
		}

		return OpenFGALexer.__ATN;
	}


	static DecisionsToDFA = OpenFGALexer._ATN.decisionToState.map( (ds: DecisionState, index: number) => new DFA(ds, index) );
}