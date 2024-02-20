// Generated from /app/OpenFGAParser.g4 by ANTLR 4.13.1
// noinspection ES6UnusedImports,JSUnusedGlobalSymbols,JSUnusedLocalSymbols

import {
	ATN,
	ATNDeserializer, DecisionState, DFA, FailedPredicateException,
	RecognitionException, NoViableAltException, BailErrorStrategy,
	Parser, ParserATNSimulator,
	RuleContext, ParserRuleContext, PredictionMode, PredictionContextCache,
	TerminalNode, RuleNode,
	Token, TokenStream,
	Interval, IntervalSet
} from 'antlr4';
import OpenFGAParserListener from "./OpenFGAParserListener.js";
// for running tests with parameters, TODO: discuss strategy for typed parameters in CI
// eslint-disable-next-line no-unused-vars
type int = number;

export default class OpenFGAParser extends Parser {
	public static readonly COLON = 1;
	public static readonly COMMA = 2;
	public static readonly LESS = 3;
	public static readonly GREATER = 4;
	public static readonly LBRACKET = 5;
	public static readonly RBRACKET = 6;
	public static readonly LPAREN = 7;
	public static readonly RPAREN = 8;
	public static readonly WHITESPACE = 9;
	public static readonly IDENTIFIER = 10;
	public static readonly HASH = 11;
	public static readonly AND = 12;
	public static readonly OR = 13;
	public static readonly BUT_NOT = 14;
	public static readonly FROM = 15;
	public static readonly MODULE = 16;
	public static readonly MODEL = 17;
	public static readonly SCHEMA = 18;
	public static readonly SCHEMA_VERSION = 19;
	public static readonly TYPE = 20;
	public static readonly CONDITION = 21;
	public static readonly RELATIONS = 22;
	public static readonly DEFINE = 23;
	public static readonly KEYWORD_WITH = 24;
	public static readonly EQUALS = 25;
	public static readonly NOT_EQUALS = 26;
	public static readonly IN = 27;
	public static readonly LESS_EQUALS = 28;
	public static readonly GREATER_EQUALS = 29;
	public static readonly LOGICAL_AND = 30;
	public static readonly LOGICAL_OR = 31;
	public static readonly RPRACKET = 32;
	public static readonly LBRACE = 33;
	public static readonly RBRACE = 34;
	public static readonly DOT = 35;
	public static readonly MINUS = 36;
	public static readonly EXCLAM = 37;
	public static readonly QUESTIONMARK = 38;
	public static readonly PLUS = 39;
	public static readonly STAR = 40;
	public static readonly SLASH = 41;
	public static readonly PERCENT = 42;
	public static readonly CEL_TRUE = 43;
	public static readonly CEL_FALSE = 44;
	public static readonly NUL = 45;
	public static readonly CEL_COMMENT = 46;
	public static readonly NUM_FLOAT = 47;
	public static readonly NUM_INT = 48;
	public static readonly NUM_UINT = 49;
	public static readonly STRING = 50;
	public static readonly BYTES = 51;
	public static readonly NEWLINE = 52;
	public static readonly CONDITION_PARAM_CONTAINER = 53;
	public static readonly CONDITION_PARAM_TYPE = 54;
	public static readonly EOF = Token.EOF;
	public static readonly RULE_main = 0;
	public static readonly RULE_modelHeader = 1;
	public static readonly RULE_moduleHeader = 2;
	public static readonly RULE_typeDefs = 3;
	public static readonly RULE_typeDef = 4;
	public static readonly RULE_relationDeclaration = 5;
	public static readonly RULE_relationName = 6;
	public static readonly RULE_relationDef = 7;
	public static readonly RULE_relationDefNoDirect = 8;
	public static readonly RULE_relationDefPartials = 9;
	public static readonly RULE_relationDefGrouping = 10;
	public static readonly RULE_relationRecurse = 11;
	public static readonly RULE_relationRecurseNoDirect = 12;
	public static readonly RULE_relationDefDirectAssignment = 13;
	public static readonly RULE_relationDefRewrite = 14;
	public static readonly RULE_relationDefTypeRestriction = 15;
	public static readonly RULE_relationDefTypeRestrictionBase = 16;
	public static readonly RULE_conditions = 17;
	public static readonly RULE_condition = 18;
	public static readonly RULE_conditionName = 19;
	public static readonly RULE_conditionParameter = 20;
	public static readonly RULE_parameterName = 21;
	public static readonly RULE_parameterType = 22;
	public static readonly RULE_multiLineComment = 23;
	public static readonly RULE_conditionExpression = 24;
	public static readonly literalNames: (string | null)[] = [ null, "':'", 
                                                            "','", "'<'", 
                                                            "'>'", "'['", 
                                                            null, "'('", 
                                                            "')'", null, 
                                                            null, "'#'", 
                                                            "'and'", "'or'", 
                                                            "'but not'", 
                                                            "'from'", "'module'", 
                                                            "'model'", "'schema'", 
                                                            null, "'type'", 
                                                            "'condition'", 
                                                            "'relations'", 
                                                            "'define'", 
                                                            "'with'", "'=='", 
                                                            "'!='", "'in'", 
                                                            "'<='", "'>='", 
                                                            "'&&'", "'||'", 
                                                            "']'", "'{'", 
                                                            "'}'", "'.'", 
                                                            "'-'", "'!'", 
                                                            "'?'", "'+'", 
                                                            "'*'", "'/'", 
                                                            "'%'", "'true'", 
                                                            "'false'", "'null'" ];
	public static readonly symbolicNames: (string | null)[] = [ null, "COLON", 
                                                             "COMMA", "LESS", 
                                                             "GREATER", 
                                                             "LBRACKET", 
                                                             "RBRACKET", 
                                                             "LPAREN", "RPAREN", 
                                                             "WHITESPACE", 
                                                             "IDENTIFIER", 
                                                             "HASH", "AND", 
                                                             "OR", "BUT_NOT", 
                                                             "FROM", "MODULE", 
                                                             "MODEL", "SCHEMA", 
                                                             "SCHEMA_VERSION", 
                                                             "TYPE", "CONDITION", 
                                                             "RELATIONS", 
                                                             "DEFINE", "KEYWORD_WITH", 
                                                             "EQUALS", "NOT_EQUALS", 
                                                             "IN", "LESS_EQUALS", 
                                                             "GREATER_EQUALS", 
                                                             "LOGICAL_AND", 
                                                             "LOGICAL_OR", 
                                                             "RPRACKET", 
                                                             "LBRACE", "RBRACE", 
                                                             "DOT", "MINUS", 
                                                             "EXCLAM", "QUESTIONMARK", 
                                                             "PLUS", "STAR", 
                                                             "SLASH", "PERCENT", 
                                                             "CEL_TRUE", 
                                                             "CEL_FALSE", 
                                                             "NUL", "CEL_COMMENT", 
                                                             "NUM_FLOAT", 
                                                             "NUM_INT", 
                                                             "NUM_UINT", 
                                                             "STRING", "BYTES", 
                                                             "NEWLINE", 
                                                             "CONDITION_PARAM_CONTAINER", 
                                                             "CONDITION_PARAM_TYPE" ];
	// tslint:disable:no-trailing-whitespace
	public static readonly ruleNames: string[] = [
		"main", "modelHeader", "moduleHeader", "typeDefs", "typeDef", "relationDeclaration", 
		"relationName", "relationDef", "relationDefNoDirect", "relationDefPartials", 
		"relationDefGrouping", "relationRecurse", "relationRecurseNoDirect", "relationDefDirectAssignment", 
		"relationDefRewrite", "relationDefTypeRestriction", "relationDefTypeRestrictionBase", 
		"conditions", "condition", "conditionName", "conditionParameter", "parameterName", 
		"parameterType", "multiLineComment", "conditionExpression",
	];
	public get grammarFileName(): string { return "OpenFGAParser.g4"; }
	public get literalNames(): (string | null)[] { return OpenFGAParser.literalNames; }
	public get symbolicNames(): (string | null)[] { return OpenFGAParser.symbolicNames; }
	public get ruleNames(): string[] { return OpenFGAParser.ruleNames; }
	public get serializedATN(): number[] { return OpenFGAParser._serializedATN; }

	protected createFailedPredicateException(predicate?: string, message?: string): FailedPredicateException {
		return new FailedPredicateException(this, predicate, message);
	}

	constructor(input: TokenStream) {
		super(input);
		this._interp = new ParserATNSimulator(this, OpenFGAParser._ATN, OpenFGAParser.DecisionsToDFA, new PredictionContextCache());
	}
	// @RuleVersion(0)
	public main(): MainContext {
		let localctx: MainContext = new MainContext(this, this._ctx, this.state);
		this.enterRule(localctx, 0, OpenFGAParser.RULE_main);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 51;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 50;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 54;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===52) {
				{
				this.state = 53;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 58;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 2, this._ctx) ) {
			case 1:
				{
				this.state = 56;
				this.modelHeader();
				}
				break;
			case 2:
				{
				this.state = 57;
				this.moduleHeader();
				}
				break;
			}
			this.state = 61;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 3, this._ctx) ) {
			case 1:
				{
				this.state = 60;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 63;
			this.typeDefs();
			this.state = 65;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 4, this._ctx) ) {
			case 1:
				{
				this.state = 64;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 67;
			this.conditions();
			this.state = 69;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===52) {
				{
				this.state = 68;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 71;
			this.match(OpenFGAParser.EOF);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public modelHeader(): ModelHeaderContext {
		let localctx: ModelHeaderContext = new ModelHeaderContext(this, this._ctx, this.state);
		this.enterRule(localctx, 2, OpenFGAParser.RULE_modelHeader);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 76;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===11) {
				{
				this.state = 73;
				this.multiLineComment();
				this.state = 74;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 78;
			this.match(OpenFGAParser.MODEL);
			this.state = 79;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 80;
			this.match(OpenFGAParser.SCHEMA);
			this.state = 81;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 82;
			localctx._schemaVersion = this.match(OpenFGAParser.SCHEMA_VERSION);
			this.state = 84;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 83;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public moduleHeader(): ModuleHeaderContext {
		let localctx: ModuleHeaderContext = new ModuleHeaderContext(this, this._ctx, this.state);
		this.enterRule(localctx, 4, OpenFGAParser.RULE_moduleHeader);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 89;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===11) {
				{
				this.state = 86;
				this.multiLineComment();
				this.state = 87;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 91;
			this.match(OpenFGAParser.MODULE);
			this.state = 92;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 93;
			localctx._moduleName = this.match(OpenFGAParser.IDENTIFIER);
			this.state = 95;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 94;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public typeDefs(): TypeDefsContext {
		let localctx: TypeDefsContext = new TypeDefsContext(this, this._ctx, this.state);
		this.enterRule(localctx, 6, OpenFGAParser.RULE_typeDefs);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 100;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 10, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 97;
					this.typeDef();
					}
					}
				}
				this.state = 102;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 10, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public typeDef(): TypeDefContext {
		let localctx: TypeDefContext = new TypeDefContext(this, this._ctx, this.state);
		this.enterRule(localctx, 8, OpenFGAParser.RULE_typeDef);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 105;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 11, this._ctx) ) {
			case 1:
				{
				this.state = 103;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 104;
				this.multiLineComment();
				}
				break;
			}
			this.state = 107;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 108;
			this.match(OpenFGAParser.TYPE);
			this.state = 109;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 110;
			localctx._typeName = this.match(OpenFGAParser.IDENTIFIER);
			this.state = 118;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 13, this._ctx) ) {
			case 1:
				{
				this.state = 111;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 112;
				this.match(OpenFGAParser.RELATIONS);
				this.state = 114;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 113;
						this.relationDeclaration();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					this.state = 116;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 12, this._ctx);
				} while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER);
				}
				break;
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public relationDeclaration(): RelationDeclarationContext {
		let localctx: RelationDeclarationContext = new RelationDeclarationContext(this, this._ctx, this.state);
		this.enterRule(localctx, 10, OpenFGAParser.RULE_relationDeclaration);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 122;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 14, this._ctx) ) {
			case 1:
				{
				this.state = 120;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 121;
				this.multiLineComment();
				}
				break;
			}
			this.state = 124;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 125;
			this.match(OpenFGAParser.DEFINE);
			this.state = 126;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 127;
			this.relationName();
			this.state = 129;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 128;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 131;
			this.match(OpenFGAParser.COLON);
			this.state = 133;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 132;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			{
			this.state = 135;
			this.relationDef();
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public relationName(): RelationNameContext {
		let localctx: RelationNameContext = new RelationNameContext(this, this._ctx, this.state);
		this.enterRule(localctx, 12, OpenFGAParser.RULE_relationName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 137;
			this.match(OpenFGAParser.IDENTIFIER);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public relationDef(): RelationDefContext {
		let localctx: RelationDefContext = new RelationDefContext(this, this._ctx, this.state);
		this.enterRule(localctx, 14, OpenFGAParser.RULE_relationDef);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 142;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 5:
				{
				this.state = 139;
				this.relationDefDirectAssignment();
				}
				break;
			case 10:
				{
				this.state = 140;
				this.relationDefGrouping();
				}
				break;
			case 7:
				{
				this.state = 141;
				this.relationRecurse();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this.state = 145;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 18, this._ctx) ) {
			case 1:
				{
				this.state = 144;
				this.relationDefPartials();
				}
				break;
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public relationDefNoDirect(): RelationDefNoDirectContext {
		let localctx: RelationDefNoDirectContext = new RelationDefNoDirectContext(this, this._ctx, this.state);
		this.enterRule(localctx, 16, OpenFGAParser.RULE_relationDefNoDirect);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 149;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 10:
				{
				this.state = 147;
				this.relationDefGrouping();
				}
				break;
			case 7:
				{
				this.state = 148;
				this.relationRecurseNoDirect();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this.state = 152;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 20, this._ctx) ) {
			case 1:
				{
				this.state = 151;
				this.relationDefPartials();
				}
				break;
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public relationDefPartials(): RelationDefPartialsContext {
		let localctx: RelationDefPartialsContext = new RelationDefPartialsContext(this, this._ctx, this.state);
		this.enterRule(localctx, 18, OpenFGAParser.RULE_relationDefPartials);
		try {
			let _alt: number;
			this.state = 183;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 26, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 161;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 154;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 155;
						this.match(OpenFGAParser.OR);
						this.state = 156;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 159;
						this._errHandler.sync(this);
						switch (this._input.LA(1)) {
						case 10:
							{
							this.state = 157;
							this.relationDefGrouping();
							}
							break;
						case 7:
							{
							this.state = 158;
							this.relationRecurseNoDirect();
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					this.state = 163;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 22, this._ctx);
				} while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER);
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 172;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 165;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 166;
						this.match(OpenFGAParser.AND);
						this.state = 167;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 170;
						this._errHandler.sync(this);
						switch (this._input.LA(1)) {
						case 10:
							{
							this.state = 168;
							this.relationDefGrouping();
							}
							break;
						case 7:
							{
							this.state = 169;
							this.relationRecurseNoDirect();
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					this.state = 174;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 24, this._ctx);
				} while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER);
				}
				break;
			case 3:
				this.enterOuterAlt(localctx, 3);
				{
				{
				this.state = 176;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 177;
				this.match(OpenFGAParser.BUT_NOT);
				this.state = 178;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 181;
				this._errHandler.sync(this);
				switch (this._input.LA(1)) {
				case 10:
					{
					this.state = 179;
					this.relationDefGrouping();
					}
					break;
				case 7:
					{
					this.state = 180;
					this.relationRecurseNoDirect();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				}
				break;
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public relationDefGrouping(): RelationDefGroupingContext {
		let localctx: RelationDefGroupingContext = new RelationDefGroupingContext(this, this._ctx, this.state);
		this.enterRule(localctx, 20, OpenFGAParser.RULE_relationDefGrouping);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 185;
			this.relationDefRewrite();
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public relationRecurse(): RelationRecurseContext {
		let localctx: RelationRecurseContext = new RelationRecurseContext(this, this._ctx, this.state);
		this.enterRule(localctx, 22, OpenFGAParser.RULE_relationRecurse);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 187;
			this.match(OpenFGAParser.LPAREN);
			this.state = 191;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 188;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 193;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 196;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 28, this._ctx) ) {
			case 1:
				{
				this.state = 194;
				this.relationDef();
				}
				break;
			case 2:
				{
				this.state = 195;
				this.relationRecurseNoDirect();
				}
				break;
			}
			this.state = 201;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 198;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 203;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 204;
			this.match(OpenFGAParser.RPAREN);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public relationRecurseNoDirect(): RelationRecurseNoDirectContext {
		let localctx: RelationRecurseNoDirectContext = new RelationRecurseNoDirectContext(this, this._ctx, this.state);
		this.enterRule(localctx, 24, OpenFGAParser.RULE_relationRecurseNoDirect);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 206;
			this.match(OpenFGAParser.LPAREN);
			this.state = 210;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 207;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 212;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 215;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 31, this._ctx) ) {
			case 1:
				{
				this.state = 213;
				this.relationDefNoDirect();
				}
				break;
			case 2:
				{
				this.state = 214;
				this.relationRecurseNoDirect();
				}
				break;
			}
			this.state = 220;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 217;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 222;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 223;
			this.match(OpenFGAParser.RPAREN);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public relationDefDirectAssignment(): RelationDefDirectAssignmentContext {
		let localctx: RelationDefDirectAssignmentContext = new RelationDefDirectAssignmentContext(this, this._ctx, this.state);
		this.enterRule(localctx, 26, OpenFGAParser.RULE_relationDefDirectAssignment);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 225;
			this.match(OpenFGAParser.LBRACKET);
			this.state = 227;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 226;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 229;
			this.relationDefTypeRestriction();
			this.state = 231;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 230;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 243;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===2) {
				{
				{
				this.state = 233;
				this.match(OpenFGAParser.COMMA);
				this.state = 235;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 234;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				this.state = 237;
				this.relationDefTypeRestriction();
				this.state = 239;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 238;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				}
				}
				this.state = 245;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 246;
			this.match(OpenFGAParser.RPRACKET);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public relationDefRewrite(): RelationDefRewriteContext {
		let localctx: RelationDefRewriteContext = new RelationDefRewriteContext(this, this._ctx, this.state);
		this.enterRule(localctx, 28, OpenFGAParser.RULE_relationDefRewrite);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 248;
			localctx._rewriteComputedusersetName = this.match(OpenFGAParser.IDENTIFIER);
			this.state = 253;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 38, this._ctx) ) {
			case 1:
				{
				this.state = 249;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 250;
				this.match(OpenFGAParser.FROM);
				this.state = 251;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 252;
				localctx._rewriteTuplesetName = this.match(OpenFGAParser.IDENTIFIER);
				}
				break;
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public relationDefTypeRestriction(): RelationDefTypeRestrictionContext {
		let localctx: RelationDefTypeRestrictionContext = new RelationDefTypeRestrictionContext(this, this._ctx, this.state);
		this.enterRule(localctx, 30, OpenFGAParser.RULE_relationDefTypeRestriction);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 256;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===52) {
				{
				this.state = 255;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 265;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 40, this._ctx) ) {
			case 1:
				{
				this.state = 258;
				this.relationDefTypeRestrictionBase();
				}
				break;
			case 2:
				{
				{
				this.state = 259;
				this.relationDefTypeRestrictionBase();
				this.state = 260;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 261;
				this.match(OpenFGAParser.KEYWORD_WITH);
				this.state = 262;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 263;
				this.conditionName();
				}
				}
				break;
			}
			this.state = 268;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===52) {
				{
				this.state = 267;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public relationDefTypeRestrictionBase(): RelationDefTypeRestrictionBaseContext {
		let localctx: RelationDefTypeRestrictionBaseContext = new RelationDefTypeRestrictionBaseContext(this, this._ctx, this.state);
		this.enterRule(localctx, 32, OpenFGAParser.RULE_relationDefTypeRestrictionBase);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 270;
			localctx._relationDefTypeRestrictionType = this.match(OpenFGAParser.IDENTIFIER);
			this.state = 275;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 1:
				{
				{
				this.state = 271;
				this.match(OpenFGAParser.COLON);
				this.state = 272;
				localctx._relationDefTypeRestrictionWildcard = this.match(OpenFGAParser.STAR);
				}
				}
				break;
			case 11:
				{
				{
				this.state = 273;
				this.match(OpenFGAParser.HASH);
				this.state = 274;
				localctx._relationDefTypeRestrictionRelation = this.match(OpenFGAParser.IDENTIFIER);
				}
				}
				break;
			case 2:
			case 9:
			case 32:
			case 52:
				break;
			default:
				break;
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public conditions(): ConditionsContext {
		let localctx: ConditionsContext = new ConditionsContext(this, this._ctx, this.state);
		this.enterRule(localctx, 34, OpenFGAParser.RULE_conditions);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 280;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 43, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 277;
					this.condition();
					}
					}
				}
				this.state = 282;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 43, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public condition(): ConditionContext {
		let localctx: ConditionContext = new ConditionContext(this, this._ctx, this.state);
		this.enterRule(localctx, 36, OpenFGAParser.RULE_condition);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 285;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 44, this._ctx) ) {
			case 1:
				{
				this.state = 283;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 284;
				this.multiLineComment();
				}
				break;
			}
			this.state = 287;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 288;
			this.match(OpenFGAParser.CONDITION);
			this.state = 289;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 290;
			this.conditionName();
			this.state = 292;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 291;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 294;
			this.match(OpenFGAParser.LPAREN);
			this.state = 296;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 295;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 298;
			this.conditionParameter();
			this.state = 300;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 299;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 312;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===2) {
				{
				{
				this.state = 302;
				this.match(OpenFGAParser.COMMA);
				this.state = 304;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 303;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				this.state = 306;
				this.conditionParameter();
				this.state = 308;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 307;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				}
				}
				this.state = 314;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 316;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===52) {
				{
				this.state = 315;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 318;
			this.match(OpenFGAParser.RPAREN);
			this.state = 320;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 319;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 322;
			this.match(OpenFGAParser.LBRACE);
			this.state = 324;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 53, this._ctx) ) {
			case 1:
				{
				this.state = 323;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 327;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 54, this._ctx) ) {
			case 1:
				{
				this.state = 326;
				this.match(OpenFGAParser.WHITESPACE);
				}
				break;
			}
			this.state = 329;
			this.conditionExpression();
			this.state = 331;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===52) {
				{
				this.state = 330;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 333;
			this.match(OpenFGAParser.RBRACE);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public conditionName(): ConditionNameContext {
		let localctx: ConditionNameContext = new ConditionNameContext(this, this._ctx, this.state);
		this.enterRule(localctx, 38, OpenFGAParser.RULE_conditionName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 335;
			this.match(OpenFGAParser.IDENTIFIER);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public conditionParameter(): ConditionParameterContext {
		let localctx: ConditionParameterContext = new ConditionParameterContext(this, this._ctx, this.state);
		this.enterRule(localctx, 40, OpenFGAParser.RULE_conditionParameter);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 338;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===52) {
				{
				this.state = 337;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 340;
			this.parameterName();
			this.state = 342;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 341;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 344;
			this.match(OpenFGAParser.COLON);
			this.state = 346;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 345;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 348;
			this.parameterType();
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public parameterName(): ParameterNameContext {
		let localctx: ParameterNameContext = new ParameterNameContext(this, this._ctx, this.state);
		this.enterRule(localctx, 42, OpenFGAParser.RULE_parameterName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 350;
			this.match(OpenFGAParser.IDENTIFIER);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public parameterType(): ParameterTypeContext {
		let localctx: ParameterTypeContext = new ParameterTypeContext(this, this._ctx, this.state);
		this.enterRule(localctx, 44, OpenFGAParser.RULE_parameterType);
		try {
			this.state = 357;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 54:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 352;
				this.match(OpenFGAParser.CONDITION_PARAM_TYPE);
				}
				break;
			case 53:
				this.enterOuterAlt(localctx, 2);
				{
				{
				this.state = 353;
				this.match(OpenFGAParser.CONDITION_PARAM_CONTAINER);
				this.state = 354;
				this.match(OpenFGAParser.LESS);
				this.state = 355;
				this.match(OpenFGAParser.CONDITION_PARAM_TYPE);
				this.state = 356;
				this.match(OpenFGAParser.GREATER);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public multiLineComment(): MultiLineCommentContext {
		let localctx: MultiLineCommentContext = new MultiLineCommentContext(this, this._ctx, this.state);
		this.enterRule(localctx, 46, OpenFGAParser.RULE_multiLineComment);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 359;
			this.match(OpenFGAParser.HASH);
			this.state = 363;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while ((((_la) & ~0x1F) === 0 && ((1 << _la) & 4294967294) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 7340031) !== 0)) {
				{
				{
				this.state = 360;
				_la = this._input.LA(1);
				if(_la<=0 || _la===52) {
				this._errHandler.recoverInline(this);
				}
				else {
					this._errHandler.reportMatch(this);
				    this.consume();
				}
				}
				}
				this.state = 365;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 368;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 61, this._ctx) ) {
			case 1:
				{
				this.state = 366;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 367;
				this.multiLineComment();
				}
				break;
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public conditionExpression(): ConditionExpressionContext {
		let localctx: ConditionExpressionContext = new ConditionExpressionContext(this, this._ctx, this.state);
		this.enterRule(localctx, 48, OpenFGAParser.RULE_conditionExpression);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 374;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 63, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					this.state = 372;
					this._errHandler.sync(this);
					switch ( this._interp.adaptivePredict(this._input, 62, this._ctx) ) {
					case 1:
						{
						this.state = 370;
						_la = this._input.LA(1);
						if(!((((_la) & ~0x1F) === 0 && ((1 << _la) & 4261414840) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 2097147) !== 0))) {
						this._errHandler.recoverInline(this);
						}
						else {
							this._errHandler.reportMatch(this);
						    this.consume();
						}
						}
						break;
					case 2:
						{
						this.state = 371;
						_la = this._input.LA(1);
						if(_la<=0 || _la===34) {
						this._errHandler.recoverInline(this);
						}
						else {
							this._errHandler.reportMatch(this);
						    this.consume();
						}
						}
						break;
					}
					}
				}
				this.state = 376;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 63, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}

	public static readonly _serializedATN: number[] = [4,1,54,378,2,0,7,0,2,
	1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,
	10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,16,2,17,
	7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,2,22,7,22,2,23,7,23,2,24,7,
	24,1,0,3,0,52,8,0,1,0,3,0,55,8,0,1,0,1,0,3,0,59,8,0,1,0,3,0,62,8,0,1,0,
	1,0,3,0,66,8,0,1,0,1,0,3,0,70,8,0,1,0,1,0,1,1,1,1,1,1,3,1,77,8,1,1,1,1,
	1,1,1,1,1,1,1,1,1,3,1,85,8,1,1,2,1,2,1,2,3,2,90,8,2,1,2,1,2,1,2,1,2,3,2,
	96,8,2,1,3,5,3,99,8,3,10,3,12,3,102,9,3,1,4,1,4,3,4,106,8,4,1,4,1,4,1,4,
	1,4,1,4,1,4,1,4,4,4,115,8,4,11,4,12,4,116,3,4,119,8,4,1,5,1,5,3,5,123,8,
	5,1,5,1,5,1,5,1,5,1,5,3,5,130,8,5,1,5,1,5,3,5,134,8,5,1,5,1,5,1,6,1,6,1,
	7,1,7,1,7,3,7,143,8,7,1,7,3,7,146,8,7,1,8,1,8,3,8,150,8,8,1,8,3,8,153,8,
	8,1,9,1,9,1,9,1,9,1,9,3,9,160,8,9,4,9,162,8,9,11,9,12,9,163,1,9,1,9,1,9,
	1,9,1,9,3,9,171,8,9,4,9,173,8,9,11,9,12,9,174,1,9,1,9,1,9,1,9,1,9,3,9,182,
	8,9,3,9,184,8,9,1,10,1,10,1,11,1,11,5,11,190,8,11,10,11,12,11,193,9,11,
	1,11,1,11,3,11,197,8,11,1,11,5,11,200,8,11,10,11,12,11,203,9,11,1,11,1,
	11,1,12,1,12,5,12,209,8,12,10,12,12,12,212,9,12,1,12,1,12,3,12,216,8,12,
	1,12,5,12,219,8,12,10,12,12,12,222,9,12,1,12,1,12,1,13,1,13,3,13,228,8,
	13,1,13,1,13,3,13,232,8,13,1,13,1,13,3,13,236,8,13,1,13,1,13,3,13,240,8,
	13,5,13,242,8,13,10,13,12,13,245,9,13,1,13,1,13,1,14,1,14,1,14,1,14,1,14,
	3,14,254,8,14,1,15,3,15,257,8,15,1,15,1,15,1,15,1,15,1,15,1,15,1,15,3,15,
	266,8,15,1,15,3,15,269,8,15,1,16,1,16,1,16,1,16,1,16,3,16,276,8,16,1,17,
	5,17,279,8,17,10,17,12,17,282,9,17,1,18,1,18,3,18,286,8,18,1,18,1,18,1,
	18,1,18,1,18,3,18,293,8,18,1,18,1,18,3,18,297,8,18,1,18,1,18,3,18,301,8,
	18,1,18,1,18,3,18,305,8,18,1,18,1,18,3,18,309,8,18,5,18,311,8,18,10,18,
	12,18,314,9,18,1,18,3,18,317,8,18,1,18,1,18,3,18,321,8,18,1,18,1,18,3,18,
	325,8,18,1,18,3,18,328,8,18,1,18,1,18,3,18,332,8,18,1,18,1,18,1,19,1,19,
	1,20,3,20,339,8,20,1,20,1,20,3,20,343,8,20,1,20,1,20,3,20,347,8,20,1,20,
	1,20,1,21,1,21,1,22,1,22,1,22,1,22,1,22,3,22,358,8,22,1,23,1,23,5,23,362,
	8,23,10,23,12,23,365,9,23,1,23,1,23,3,23,369,8,23,1,24,1,24,5,24,373,8,
	24,10,24,12,24,376,9,24,1,24,0,0,25,0,2,4,6,8,10,12,14,16,18,20,22,24,26,
	28,30,32,34,36,38,40,42,44,46,48,0,3,1,0,52,52,4,0,3,5,7,10,25,33,35,52,
	1,0,34,34,419,0,51,1,0,0,0,2,76,1,0,0,0,4,89,1,0,0,0,6,100,1,0,0,0,8,105,
	1,0,0,0,10,122,1,0,0,0,12,137,1,0,0,0,14,142,1,0,0,0,16,149,1,0,0,0,18,
	183,1,0,0,0,20,185,1,0,0,0,22,187,1,0,0,0,24,206,1,0,0,0,26,225,1,0,0,0,
	28,248,1,0,0,0,30,256,1,0,0,0,32,270,1,0,0,0,34,280,1,0,0,0,36,285,1,0,
	0,0,38,335,1,0,0,0,40,338,1,0,0,0,42,350,1,0,0,0,44,357,1,0,0,0,46,359,
	1,0,0,0,48,374,1,0,0,0,50,52,5,9,0,0,51,50,1,0,0,0,51,52,1,0,0,0,52,54,
	1,0,0,0,53,55,5,52,0,0,54,53,1,0,0,0,54,55,1,0,0,0,55,58,1,0,0,0,56,59,
	3,2,1,0,57,59,3,4,2,0,58,56,1,0,0,0,58,57,1,0,0,0,59,61,1,0,0,0,60,62,5,
	52,0,0,61,60,1,0,0,0,61,62,1,0,0,0,62,63,1,0,0,0,63,65,3,6,3,0,64,66,5,
	52,0,0,65,64,1,0,0,0,65,66,1,0,0,0,66,67,1,0,0,0,67,69,3,34,17,0,68,70,
	5,52,0,0,69,68,1,0,0,0,69,70,1,0,0,0,70,71,1,0,0,0,71,72,5,0,0,1,72,1,1,
	0,0,0,73,74,3,46,23,0,74,75,5,52,0,0,75,77,1,0,0,0,76,73,1,0,0,0,76,77,
	1,0,0,0,77,78,1,0,0,0,78,79,5,17,0,0,79,80,5,52,0,0,80,81,5,18,0,0,81,82,
	5,9,0,0,82,84,5,19,0,0,83,85,5,9,0,0,84,83,1,0,0,0,84,85,1,0,0,0,85,3,1,
	0,0,0,86,87,3,46,23,0,87,88,5,52,0,0,88,90,1,0,0,0,89,86,1,0,0,0,89,90,
	1,0,0,0,90,91,1,0,0,0,91,92,5,16,0,0,92,93,5,9,0,0,93,95,5,10,0,0,94,96,
	5,9,0,0,95,94,1,0,0,0,95,96,1,0,0,0,96,5,1,0,0,0,97,99,3,8,4,0,98,97,1,
	0,0,0,99,102,1,0,0,0,100,98,1,0,0,0,100,101,1,0,0,0,101,7,1,0,0,0,102,100,
	1,0,0,0,103,104,5,52,0,0,104,106,3,46,23,0,105,103,1,0,0,0,105,106,1,0,
	0,0,106,107,1,0,0,0,107,108,5,52,0,0,108,109,5,20,0,0,109,110,5,9,0,0,110,
	118,5,10,0,0,111,112,5,52,0,0,112,114,5,22,0,0,113,115,3,10,5,0,114,113,
	1,0,0,0,115,116,1,0,0,0,116,114,1,0,0,0,116,117,1,0,0,0,117,119,1,0,0,0,
	118,111,1,0,0,0,118,119,1,0,0,0,119,9,1,0,0,0,120,121,5,52,0,0,121,123,
	3,46,23,0,122,120,1,0,0,0,122,123,1,0,0,0,123,124,1,0,0,0,124,125,5,52,
	0,0,125,126,5,23,0,0,126,127,5,9,0,0,127,129,3,12,6,0,128,130,5,9,0,0,129,
	128,1,0,0,0,129,130,1,0,0,0,130,131,1,0,0,0,131,133,5,1,0,0,132,134,5,9,
	0,0,133,132,1,0,0,0,133,134,1,0,0,0,134,135,1,0,0,0,135,136,3,14,7,0,136,
	11,1,0,0,0,137,138,5,10,0,0,138,13,1,0,0,0,139,143,3,26,13,0,140,143,3,
	20,10,0,141,143,3,22,11,0,142,139,1,0,0,0,142,140,1,0,0,0,142,141,1,0,0,
	0,143,145,1,0,0,0,144,146,3,18,9,0,145,144,1,0,0,0,145,146,1,0,0,0,146,
	15,1,0,0,0,147,150,3,20,10,0,148,150,3,24,12,0,149,147,1,0,0,0,149,148,
	1,0,0,0,150,152,1,0,0,0,151,153,3,18,9,0,152,151,1,0,0,0,152,153,1,0,0,
	0,153,17,1,0,0,0,154,155,5,9,0,0,155,156,5,13,0,0,156,159,5,9,0,0,157,160,
	3,20,10,0,158,160,3,24,12,0,159,157,1,0,0,0,159,158,1,0,0,0,160,162,1,0,
	0,0,161,154,1,0,0,0,162,163,1,0,0,0,163,161,1,0,0,0,163,164,1,0,0,0,164,
	184,1,0,0,0,165,166,5,9,0,0,166,167,5,12,0,0,167,170,5,9,0,0,168,171,3,
	20,10,0,169,171,3,24,12,0,170,168,1,0,0,0,170,169,1,0,0,0,171,173,1,0,0,
	0,172,165,1,0,0,0,173,174,1,0,0,0,174,172,1,0,0,0,174,175,1,0,0,0,175,184,
	1,0,0,0,176,177,5,9,0,0,177,178,5,14,0,0,178,181,5,9,0,0,179,182,3,20,10,
	0,180,182,3,24,12,0,181,179,1,0,0,0,181,180,1,0,0,0,182,184,1,0,0,0,183,
	161,1,0,0,0,183,172,1,0,0,0,183,176,1,0,0,0,184,19,1,0,0,0,185,186,3,28,
	14,0,186,21,1,0,0,0,187,191,5,7,0,0,188,190,5,9,0,0,189,188,1,0,0,0,190,
	193,1,0,0,0,191,189,1,0,0,0,191,192,1,0,0,0,192,196,1,0,0,0,193,191,1,0,
	0,0,194,197,3,14,7,0,195,197,3,24,12,0,196,194,1,0,0,0,196,195,1,0,0,0,
	197,201,1,0,0,0,198,200,5,9,0,0,199,198,1,0,0,0,200,203,1,0,0,0,201,199,
	1,0,0,0,201,202,1,0,0,0,202,204,1,0,0,0,203,201,1,0,0,0,204,205,5,8,0,0,
	205,23,1,0,0,0,206,210,5,7,0,0,207,209,5,9,0,0,208,207,1,0,0,0,209,212,
	1,0,0,0,210,208,1,0,0,0,210,211,1,0,0,0,211,215,1,0,0,0,212,210,1,0,0,0,
	213,216,3,16,8,0,214,216,3,24,12,0,215,213,1,0,0,0,215,214,1,0,0,0,216,
	220,1,0,0,0,217,219,5,9,0,0,218,217,1,0,0,0,219,222,1,0,0,0,220,218,1,0,
	0,0,220,221,1,0,0,0,221,223,1,0,0,0,222,220,1,0,0,0,223,224,5,8,0,0,224,
	25,1,0,0,0,225,227,5,5,0,0,226,228,5,9,0,0,227,226,1,0,0,0,227,228,1,0,
	0,0,228,229,1,0,0,0,229,231,3,30,15,0,230,232,5,9,0,0,231,230,1,0,0,0,231,
	232,1,0,0,0,232,243,1,0,0,0,233,235,5,2,0,0,234,236,5,9,0,0,235,234,1,0,
	0,0,235,236,1,0,0,0,236,237,1,0,0,0,237,239,3,30,15,0,238,240,5,9,0,0,239,
	238,1,0,0,0,239,240,1,0,0,0,240,242,1,0,0,0,241,233,1,0,0,0,242,245,1,0,
	0,0,243,241,1,0,0,0,243,244,1,0,0,0,244,246,1,0,0,0,245,243,1,0,0,0,246,
	247,5,32,0,0,247,27,1,0,0,0,248,253,5,10,0,0,249,250,5,9,0,0,250,251,5,
	15,0,0,251,252,5,9,0,0,252,254,5,10,0,0,253,249,1,0,0,0,253,254,1,0,0,0,
	254,29,1,0,0,0,255,257,5,52,0,0,256,255,1,0,0,0,256,257,1,0,0,0,257,265,
	1,0,0,0,258,266,3,32,16,0,259,260,3,32,16,0,260,261,5,9,0,0,261,262,5,24,
	0,0,262,263,5,9,0,0,263,264,3,38,19,0,264,266,1,0,0,0,265,258,1,0,0,0,265,
	259,1,0,0,0,266,268,1,0,0,0,267,269,5,52,0,0,268,267,1,0,0,0,268,269,1,
	0,0,0,269,31,1,0,0,0,270,275,5,10,0,0,271,272,5,1,0,0,272,276,5,40,0,0,
	273,274,5,11,0,0,274,276,5,10,0,0,275,271,1,0,0,0,275,273,1,0,0,0,275,276,
	1,0,0,0,276,33,1,0,0,0,277,279,3,36,18,0,278,277,1,0,0,0,279,282,1,0,0,
	0,280,278,1,0,0,0,280,281,1,0,0,0,281,35,1,0,0,0,282,280,1,0,0,0,283,284,
	5,52,0,0,284,286,3,46,23,0,285,283,1,0,0,0,285,286,1,0,0,0,286,287,1,0,
	0,0,287,288,5,52,0,0,288,289,5,21,0,0,289,290,5,9,0,0,290,292,3,38,19,0,
	291,293,5,9,0,0,292,291,1,0,0,0,292,293,1,0,0,0,293,294,1,0,0,0,294,296,
	5,7,0,0,295,297,5,9,0,0,296,295,1,0,0,0,296,297,1,0,0,0,297,298,1,0,0,0,
	298,300,3,40,20,0,299,301,5,9,0,0,300,299,1,0,0,0,300,301,1,0,0,0,301,312,
	1,0,0,0,302,304,5,2,0,0,303,305,5,9,0,0,304,303,1,0,0,0,304,305,1,0,0,0,
	305,306,1,0,0,0,306,308,3,40,20,0,307,309,5,9,0,0,308,307,1,0,0,0,308,309,
	1,0,0,0,309,311,1,0,0,0,310,302,1,0,0,0,311,314,1,0,0,0,312,310,1,0,0,0,
	312,313,1,0,0,0,313,316,1,0,0,0,314,312,1,0,0,0,315,317,5,52,0,0,316,315,
	1,0,0,0,316,317,1,0,0,0,317,318,1,0,0,0,318,320,5,8,0,0,319,321,5,9,0,0,
	320,319,1,0,0,0,320,321,1,0,0,0,321,322,1,0,0,0,322,324,5,33,0,0,323,325,
	5,52,0,0,324,323,1,0,0,0,324,325,1,0,0,0,325,327,1,0,0,0,326,328,5,9,0,
	0,327,326,1,0,0,0,327,328,1,0,0,0,328,329,1,0,0,0,329,331,3,48,24,0,330,
	332,5,52,0,0,331,330,1,0,0,0,331,332,1,0,0,0,332,333,1,0,0,0,333,334,5,
	34,0,0,334,37,1,0,0,0,335,336,5,10,0,0,336,39,1,0,0,0,337,339,5,52,0,0,
	338,337,1,0,0,0,338,339,1,0,0,0,339,340,1,0,0,0,340,342,3,42,21,0,341,343,
	5,9,0,0,342,341,1,0,0,0,342,343,1,0,0,0,343,344,1,0,0,0,344,346,5,1,0,0,
	345,347,5,9,0,0,346,345,1,0,0,0,346,347,1,0,0,0,347,348,1,0,0,0,348,349,
	3,44,22,0,349,41,1,0,0,0,350,351,5,10,0,0,351,43,1,0,0,0,352,358,5,54,0,
	0,353,354,5,53,0,0,354,355,5,3,0,0,355,356,5,54,0,0,356,358,5,4,0,0,357,
	352,1,0,0,0,357,353,1,0,0,0,358,45,1,0,0,0,359,363,5,11,0,0,360,362,8,0,
	0,0,361,360,1,0,0,0,362,365,1,0,0,0,363,361,1,0,0,0,363,364,1,0,0,0,364,
	368,1,0,0,0,365,363,1,0,0,0,366,367,5,52,0,0,367,369,3,46,23,0,368,366,
	1,0,0,0,368,369,1,0,0,0,369,47,1,0,0,0,370,373,7,1,0,0,371,373,8,2,0,0,
	372,370,1,0,0,0,372,371,1,0,0,0,373,376,1,0,0,0,374,372,1,0,0,0,374,375,
	1,0,0,0,375,49,1,0,0,0,376,374,1,0,0,0,64,51,54,58,61,65,69,76,84,89,95,
	100,105,116,118,122,129,133,142,145,149,152,159,163,170,174,181,183,191,
	196,201,210,215,220,227,231,235,239,243,253,256,265,268,275,280,285,292,
	296,300,304,308,312,316,320,324,327,331,338,342,346,357,363,368,372,374];

	private static __ATN: ATN;
	public static get _ATN(): ATN {
		if (!OpenFGAParser.__ATN) {
			OpenFGAParser.__ATN = new ATNDeserializer().deserialize(OpenFGAParser._serializedATN);
		}

		return OpenFGAParser.__ATN;
	}


	static DecisionsToDFA = OpenFGAParser._ATN.decisionToState.map( (ds: DecisionState, index: number) => new DFA(ds, index) );

}

export class MainContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public typeDefs(): TypeDefsContext {
		return this.getTypedRuleContext(TypeDefsContext, 0) as TypeDefsContext;
	}
	public conditions(): ConditionsContext {
		return this.getTypedRuleContext(ConditionsContext, 0) as ConditionsContext;
	}
	public EOF(): TerminalNode {
		return this.getToken(OpenFGAParser.EOF, 0);
	}
	public modelHeader(): ModelHeaderContext {
		return this.getTypedRuleContext(ModelHeaderContext, 0) as ModelHeaderContext;
	}
	public moduleHeader(): ModuleHeaderContext {
		return this.getTypedRuleContext(ModuleHeaderContext, 0) as ModuleHeaderContext;
	}
	public WHITESPACE(): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, 0);
	}
	public NEWLINE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.NEWLINE);
	}
	public NEWLINE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.NEWLINE, i);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_main;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterMain) {
	 		listener.enterMain(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitMain) {
	 		listener.exitMain(this);
		}
	}
}


export class ModelHeaderContext extends ParserRuleContext {
	public _schemaVersion!: Token;
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public MODEL(): TerminalNode {
		return this.getToken(OpenFGAParser.MODEL, 0);
	}
	public NEWLINE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.NEWLINE);
	}
	public NEWLINE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.NEWLINE, i);
	}
	public SCHEMA(): TerminalNode {
		return this.getToken(OpenFGAParser.SCHEMA, 0);
	}
	public WHITESPACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WHITESPACE);
	}
	public WHITESPACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, i);
	}
	public SCHEMA_VERSION(): TerminalNode {
		return this.getToken(OpenFGAParser.SCHEMA_VERSION, 0);
	}
	public multiLineComment(): MultiLineCommentContext {
		return this.getTypedRuleContext(MultiLineCommentContext, 0) as MultiLineCommentContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_modelHeader;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterModelHeader) {
	 		listener.enterModelHeader(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitModelHeader) {
	 		listener.exitModelHeader(this);
		}
	}
}


export class ModuleHeaderContext extends ParserRuleContext {
	public _moduleName!: Token;
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public MODULE(): TerminalNode {
		return this.getToken(OpenFGAParser.MODULE, 0);
	}
	public WHITESPACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WHITESPACE);
	}
	public WHITESPACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, i);
	}
	public IDENTIFIER(): TerminalNode {
		return this.getToken(OpenFGAParser.IDENTIFIER, 0);
	}
	public multiLineComment(): MultiLineCommentContext {
		return this.getTypedRuleContext(MultiLineCommentContext, 0) as MultiLineCommentContext;
	}
	public NEWLINE(): TerminalNode {
		return this.getToken(OpenFGAParser.NEWLINE, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_moduleHeader;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterModuleHeader) {
	 		listener.enterModuleHeader(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitModuleHeader) {
	 		listener.exitModuleHeader(this);
		}
	}
}


export class TypeDefsContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public typeDef_list(): TypeDefContext[] {
		return this.getTypedRuleContexts(TypeDefContext) as TypeDefContext[];
	}
	public typeDef(i: number): TypeDefContext {
		return this.getTypedRuleContext(TypeDefContext, i) as TypeDefContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_typeDefs;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterTypeDefs) {
	 		listener.enterTypeDefs(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitTypeDefs) {
	 		listener.exitTypeDefs(this);
		}
	}
}


export class TypeDefContext extends ParserRuleContext {
	public _typeName!: Token;
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public NEWLINE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.NEWLINE);
	}
	public NEWLINE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.NEWLINE, i);
	}
	public TYPE(): TerminalNode {
		return this.getToken(OpenFGAParser.TYPE, 0);
	}
	public WHITESPACE(): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, 0);
	}
	public IDENTIFIER(): TerminalNode {
		return this.getToken(OpenFGAParser.IDENTIFIER, 0);
	}
	public multiLineComment(): MultiLineCommentContext {
		return this.getTypedRuleContext(MultiLineCommentContext, 0) as MultiLineCommentContext;
	}
	public RELATIONS(): TerminalNode {
		return this.getToken(OpenFGAParser.RELATIONS, 0);
	}
	public relationDeclaration_list(): RelationDeclarationContext[] {
		return this.getTypedRuleContexts(RelationDeclarationContext) as RelationDeclarationContext[];
	}
	public relationDeclaration(i: number): RelationDeclarationContext {
		return this.getTypedRuleContext(RelationDeclarationContext, i) as RelationDeclarationContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_typeDef;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterTypeDef) {
	 		listener.enterTypeDef(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitTypeDef) {
	 		listener.exitTypeDef(this);
		}
	}
}


export class RelationDeclarationContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public NEWLINE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.NEWLINE);
	}
	public NEWLINE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.NEWLINE, i);
	}
	public DEFINE(): TerminalNode {
		return this.getToken(OpenFGAParser.DEFINE, 0);
	}
	public WHITESPACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WHITESPACE);
	}
	public WHITESPACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, i);
	}
	public relationName(): RelationNameContext {
		return this.getTypedRuleContext(RelationNameContext, 0) as RelationNameContext;
	}
	public COLON(): TerminalNode {
		return this.getToken(OpenFGAParser.COLON, 0);
	}
	public relationDef(): RelationDefContext {
		return this.getTypedRuleContext(RelationDefContext, 0) as RelationDefContext;
	}
	public multiLineComment(): MultiLineCommentContext {
		return this.getTypedRuleContext(MultiLineCommentContext, 0) as MultiLineCommentContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDeclaration;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDeclaration) {
	 		listener.enterRelationDeclaration(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDeclaration) {
	 		listener.exitRelationDeclaration(this);
		}
	}
}


export class RelationNameContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public IDENTIFIER(): TerminalNode {
		return this.getToken(OpenFGAParser.IDENTIFIER, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationName;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationName) {
	 		listener.enterRelationName(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationName) {
	 		listener.exitRelationName(this);
		}
	}
}


export class RelationDefContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public relationDefDirectAssignment(): RelationDefDirectAssignmentContext {
		return this.getTypedRuleContext(RelationDefDirectAssignmentContext, 0) as RelationDefDirectAssignmentContext;
	}
	public relationDefGrouping(): RelationDefGroupingContext {
		return this.getTypedRuleContext(RelationDefGroupingContext, 0) as RelationDefGroupingContext;
	}
	public relationRecurse(): RelationRecurseContext {
		return this.getTypedRuleContext(RelationRecurseContext, 0) as RelationRecurseContext;
	}
	public relationDefPartials(): RelationDefPartialsContext {
		return this.getTypedRuleContext(RelationDefPartialsContext, 0) as RelationDefPartialsContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDef;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDef) {
	 		listener.enterRelationDef(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDef) {
	 		listener.exitRelationDef(this);
		}
	}
}


export class RelationDefNoDirectContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public relationDefGrouping(): RelationDefGroupingContext {
		return this.getTypedRuleContext(RelationDefGroupingContext, 0) as RelationDefGroupingContext;
	}
	public relationRecurseNoDirect(): RelationRecurseNoDirectContext {
		return this.getTypedRuleContext(RelationRecurseNoDirectContext, 0) as RelationRecurseNoDirectContext;
	}
	public relationDefPartials(): RelationDefPartialsContext {
		return this.getTypedRuleContext(RelationDefPartialsContext, 0) as RelationDefPartialsContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefNoDirect;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefNoDirect) {
	 		listener.enterRelationDefNoDirect(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefNoDirect) {
	 		listener.exitRelationDefNoDirect(this);
		}
	}
}


export class RelationDefPartialsContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public WHITESPACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WHITESPACE);
	}
	public WHITESPACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, i);
	}
	public OR_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.OR);
	}
	public OR(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.OR, i);
	}
	public relationDefGrouping_list(): RelationDefGroupingContext[] {
		return this.getTypedRuleContexts(RelationDefGroupingContext) as RelationDefGroupingContext[];
	}
	public relationDefGrouping(i: number): RelationDefGroupingContext {
		return this.getTypedRuleContext(RelationDefGroupingContext, i) as RelationDefGroupingContext;
	}
	public relationRecurseNoDirect_list(): RelationRecurseNoDirectContext[] {
		return this.getTypedRuleContexts(RelationRecurseNoDirectContext) as RelationRecurseNoDirectContext[];
	}
	public relationRecurseNoDirect(i: number): RelationRecurseNoDirectContext {
		return this.getTypedRuleContext(RelationRecurseNoDirectContext, i) as RelationRecurseNoDirectContext;
	}
	public AND_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.AND);
	}
	public AND(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.AND, i);
	}
	public BUT_NOT(): TerminalNode {
		return this.getToken(OpenFGAParser.BUT_NOT, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefPartials;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefPartials) {
	 		listener.enterRelationDefPartials(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefPartials) {
	 		listener.exitRelationDefPartials(this);
		}
	}
}


export class RelationDefGroupingContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public relationDefRewrite(): RelationDefRewriteContext {
		return this.getTypedRuleContext(RelationDefRewriteContext, 0) as RelationDefRewriteContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefGrouping;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefGrouping) {
	 		listener.enterRelationDefGrouping(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefGrouping) {
	 		listener.exitRelationDefGrouping(this);
		}
	}
}


export class RelationRecurseContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public LPAREN(): TerminalNode {
		return this.getToken(OpenFGAParser.LPAREN, 0);
	}
	public RPAREN(): TerminalNode {
		return this.getToken(OpenFGAParser.RPAREN, 0);
	}
	public relationDef(): RelationDefContext {
		return this.getTypedRuleContext(RelationDefContext, 0) as RelationDefContext;
	}
	public relationRecurseNoDirect(): RelationRecurseNoDirectContext {
		return this.getTypedRuleContext(RelationRecurseNoDirectContext, 0) as RelationRecurseNoDirectContext;
	}
	public WHITESPACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WHITESPACE);
	}
	public WHITESPACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, i);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationRecurse;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationRecurse) {
	 		listener.enterRelationRecurse(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationRecurse) {
	 		listener.exitRelationRecurse(this);
		}
	}
}


export class RelationRecurseNoDirectContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public LPAREN(): TerminalNode {
		return this.getToken(OpenFGAParser.LPAREN, 0);
	}
	public RPAREN(): TerminalNode {
		return this.getToken(OpenFGAParser.RPAREN, 0);
	}
	public relationDefNoDirect(): RelationDefNoDirectContext {
		return this.getTypedRuleContext(RelationDefNoDirectContext, 0) as RelationDefNoDirectContext;
	}
	public relationRecurseNoDirect(): RelationRecurseNoDirectContext {
		return this.getTypedRuleContext(RelationRecurseNoDirectContext, 0) as RelationRecurseNoDirectContext;
	}
	public WHITESPACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WHITESPACE);
	}
	public WHITESPACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, i);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationRecurseNoDirect;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationRecurseNoDirect) {
	 		listener.enterRelationRecurseNoDirect(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationRecurseNoDirect) {
	 		listener.exitRelationRecurseNoDirect(this);
		}
	}
}


export class RelationDefDirectAssignmentContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public LBRACKET(): TerminalNode {
		return this.getToken(OpenFGAParser.LBRACKET, 0);
	}
	public relationDefTypeRestriction_list(): RelationDefTypeRestrictionContext[] {
		return this.getTypedRuleContexts(RelationDefTypeRestrictionContext) as RelationDefTypeRestrictionContext[];
	}
	public relationDefTypeRestriction(i: number): RelationDefTypeRestrictionContext {
		return this.getTypedRuleContext(RelationDefTypeRestrictionContext, i) as RelationDefTypeRestrictionContext;
	}
	public RPRACKET(): TerminalNode {
		return this.getToken(OpenFGAParser.RPRACKET, 0);
	}
	public WHITESPACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WHITESPACE);
	}
	public WHITESPACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, i);
	}
	public COMMA_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.COMMA);
	}
	public COMMA(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.COMMA, i);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefDirectAssignment;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefDirectAssignment) {
	 		listener.enterRelationDefDirectAssignment(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefDirectAssignment) {
	 		listener.exitRelationDefDirectAssignment(this);
		}
	}
}


export class RelationDefRewriteContext extends ParserRuleContext {
	public _rewriteComputedusersetName!: Token;
	public _rewriteTuplesetName!: Token;
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public IDENTIFIER_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.IDENTIFIER);
	}
	public IDENTIFIER(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.IDENTIFIER, i);
	}
	public WHITESPACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WHITESPACE);
	}
	public WHITESPACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, i);
	}
	public FROM(): TerminalNode {
		return this.getToken(OpenFGAParser.FROM, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefRewrite;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefRewrite) {
	 		listener.enterRelationDefRewrite(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefRewrite) {
	 		listener.exitRelationDefRewrite(this);
		}
	}
}


export class RelationDefTypeRestrictionContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public relationDefTypeRestrictionBase(): RelationDefTypeRestrictionBaseContext {
		return this.getTypedRuleContext(RelationDefTypeRestrictionBaseContext, 0) as RelationDefTypeRestrictionBaseContext;
	}
	public NEWLINE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.NEWLINE);
	}
	public NEWLINE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.NEWLINE, i);
	}
	public WHITESPACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WHITESPACE);
	}
	public WHITESPACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, i);
	}
	public KEYWORD_WITH(): TerminalNode {
		return this.getToken(OpenFGAParser.KEYWORD_WITH, 0);
	}
	public conditionName(): ConditionNameContext {
		return this.getTypedRuleContext(ConditionNameContext, 0) as ConditionNameContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefTypeRestriction;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefTypeRestriction) {
	 		listener.enterRelationDefTypeRestriction(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefTypeRestriction) {
	 		listener.exitRelationDefTypeRestriction(this);
		}
	}
}


export class RelationDefTypeRestrictionBaseContext extends ParserRuleContext {
	public _relationDefTypeRestrictionType!: Token;
	public _relationDefTypeRestrictionWildcard!: Token;
	public _relationDefTypeRestrictionRelation!: Token;
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public IDENTIFIER_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.IDENTIFIER);
	}
	public IDENTIFIER(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.IDENTIFIER, i);
	}
	public COLON(): TerminalNode {
		return this.getToken(OpenFGAParser.COLON, 0);
	}
	public HASH(): TerminalNode {
		return this.getToken(OpenFGAParser.HASH, 0);
	}
	public STAR(): TerminalNode {
		return this.getToken(OpenFGAParser.STAR, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefTypeRestrictionBase;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefTypeRestrictionBase) {
	 		listener.enterRelationDefTypeRestrictionBase(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefTypeRestrictionBase) {
	 		listener.exitRelationDefTypeRestrictionBase(this);
		}
	}
}


export class ConditionsContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public condition_list(): ConditionContext[] {
		return this.getTypedRuleContexts(ConditionContext) as ConditionContext[];
	}
	public condition(i: number): ConditionContext {
		return this.getTypedRuleContext(ConditionContext, i) as ConditionContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_conditions;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterConditions) {
	 		listener.enterConditions(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitConditions) {
	 		listener.exitConditions(this);
		}
	}
}


export class ConditionContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public NEWLINE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.NEWLINE);
	}
	public NEWLINE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.NEWLINE, i);
	}
	public CONDITION(): TerminalNode {
		return this.getToken(OpenFGAParser.CONDITION, 0);
	}
	public WHITESPACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WHITESPACE);
	}
	public WHITESPACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, i);
	}
	public conditionName(): ConditionNameContext {
		return this.getTypedRuleContext(ConditionNameContext, 0) as ConditionNameContext;
	}
	public LPAREN(): TerminalNode {
		return this.getToken(OpenFGAParser.LPAREN, 0);
	}
	public conditionParameter_list(): ConditionParameterContext[] {
		return this.getTypedRuleContexts(ConditionParameterContext) as ConditionParameterContext[];
	}
	public conditionParameter(i: number): ConditionParameterContext {
		return this.getTypedRuleContext(ConditionParameterContext, i) as ConditionParameterContext;
	}
	public RPAREN(): TerminalNode {
		return this.getToken(OpenFGAParser.RPAREN, 0);
	}
	public LBRACE(): TerminalNode {
		return this.getToken(OpenFGAParser.LBRACE, 0);
	}
	public conditionExpression(): ConditionExpressionContext {
		return this.getTypedRuleContext(ConditionExpressionContext, 0) as ConditionExpressionContext;
	}
	public RBRACE(): TerminalNode {
		return this.getToken(OpenFGAParser.RBRACE, 0);
	}
	public multiLineComment(): MultiLineCommentContext {
		return this.getTypedRuleContext(MultiLineCommentContext, 0) as MultiLineCommentContext;
	}
	public COMMA_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.COMMA);
	}
	public COMMA(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.COMMA, i);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_condition;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterCondition) {
	 		listener.enterCondition(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitCondition) {
	 		listener.exitCondition(this);
		}
	}
}


export class ConditionNameContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public IDENTIFIER(): TerminalNode {
		return this.getToken(OpenFGAParser.IDENTIFIER, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_conditionName;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterConditionName) {
	 		listener.enterConditionName(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitConditionName) {
	 		listener.exitConditionName(this);
		}
	}
}


export class ConditionParameterContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public parameterName(): ParameterNameContext {
		return this.getTypedRuleContext(ParameterNameContext, 0) as ParameterNameContext;
	}
	public COLON(): TerminalNode {
		return this.getToken(OpenFGAParser.COLON, 0);
	}
	public parameterType(): ParameterTypeContext {
		return this.getTypedRuleContext(ParameterTypeContext, 0) as ParameterTypeContext;
	}
	public NEWLINE(): TerminalNode {
		return this.getToken(OpenFGAParser.NEWLINE, 0);
	}
	public WHITESPACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WHITESPACE);
	}
	public WHITESPACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, i);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_conditionParameter;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterConditionParameter) {
	 		listener.enterConditionParameter(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitConditionParameter) {
	 		listener.exitConditionParameter(this);
		}
	}
}


export class ParameterNameContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public IDENTIFIER(): TerminalNode {
		return this.getToken(OpenFGAParser.IDENTIFIER, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_parameterName;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterParameterName) {
	 		listener.enterParameterName(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitParameterName) {
	 		listener.exitParameterName(this);
		}
	}
}


export class ParameterTypeContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public CONDITION_PARAM_TYPE(): TerminalNode {
		return this.getToken(OpenFGAParser.CONDITION_PARAM_TYPE, 0);
	}
	public CONDITION_PARAM_CONTAINER(): TerminalNode {
		return this.getToken(OpenFGAParser.CONDITION_PARAM_CONTAINER, 0);
	}
	public LESS(): TerminalNode {
		return this.getToken(OpenFGAParser.LESS, 0);
	}
	public GREATER(): TerminalNode {
		return this.getToken(OpenFGAParser.GREATER, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_parameterType;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterParameterType) {
	 		listener.enterParameterType(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitParameterType) {
	 		listener.exitParameterType(this);
		}
	}
}


export class MultiLineCommentContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public HASH(): TerminalNode {
		return this.getToken(OpenFGAParser.HASH, 0);
	}
	public NEWLINE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.NEWLINE);
	}
	public NEWLINE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.NEWLINE, i);
	}
	public multiLineComment(): MultiLineCommentContext {
		return this.getTypedRuleContext(MultiLineCommentContext, 0) as MultiLineCommentContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_multiLineComment;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterMultiLineComment) {
	 		listener.enterMultiLineComment(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitMultiLineComment) {
	 		listener.exitMultiLineComment(this);
		}
	}
}


export class ConditionExpressionContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public IDENTIFIER_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.IDENTIFIER);
	}
	public IDENTIFIER(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.IDENTIFIER, i);
	}
	public EQUALS_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.EQUALS);
	}
	public EQUALS(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.EQUALS, i);
	}
	public NOT_EQUALS_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.NOT_EQUALS);
	}
	public NOT_EQUALS(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.NOT_EQUALS, i);
	}
	public IN_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.IN);
	}
	public IN(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.IN, i);
	}
	public LESS_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.LESS);
	}
	public LESS(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.LESS, i);
	}
	public LESS_EQUALS_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.LESS_EQUALS);
	}
	public LESS_EQUALS(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.LESS_EQUALS, i);
	}
	public GREATER_EQUALS_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.GREATER_EQUALS);
	}
	public GREATER_EQUALS(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.GREATER_EQUALS, i);
	}
	public GREATER_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.GREATER);
	}
	public GREATER(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.GREATER, i);
	}
	public LOGICAL_AND_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.LOGICAL_AND);
	}
	public LOGICAL_AND(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.LOGICAL_AND, i);
	}
	public LOGICAL_OR_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.LOGICAL_OR);
	}
	public LOGICAL_OR(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.LOGICAL_OR, i);
	}
	public LBRACKET_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.LBRACKET);
	}
	public LBRACKET(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.LBRACKET, i);
	}
	public RPRACKET_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.RPRACKET);
	}
	public RPRACKET(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.RPRACKET, i);
	}
	public LBRACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.LBRACE);
	}
	public LBRACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.LBRACE, i);
	}
	public LPAREN_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.LPAREN);
	}
	public LPAREN(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.LPAREN, i);
	}
	public RPAREN_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.RPAREN);
	}
	public RPAREN(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.RPAREN, i);
	}
	public DOT_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.DOT);
	}
	public DOT(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.DOT, i);
	}
	public MINUS_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.MINUS);
	}
	public MINUS(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.MINUS, i);
	}
	public EXCLAM_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.EXCLAM);
	}
	public EXCLAM(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.EXCLAM, i);
	}
	public QUESTIONMARK_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.QUESTIONMARK);
	}
	public QUESTIONMARK(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.QUESTIONMARK, i);
	}
	public PLUS_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.PLUS);
	}
	public PLUS(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.PLUS, i);
	}
	public STAR_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.STAR);
	}
	public STAR(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.STAR, i);
	}
	public SLASH_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.SLASH);
	}
	public SLASH(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.SLASH, i);
	}
	public PERCENT_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.PERCENT);
	}
	public PERCENT(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.PERCENT, i);
	}
	public CEL_TRUE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.CEL_TRUE);
	}
	public CEL_TRUE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.CEL_TRUE, i);
	}
	public CEL_FALSE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.CEL_FALSE);
	}
	public CEL_FALSE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.CEL_FALSE, i);
	}
	public NUL_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.NUL);
	}
	public NUL(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.NUL, i);
	}
	public WHITESPACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WHITESPACE);
	}
	public WHITESPACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, i);
	}
	public CEL_COMMENT_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.CEL_COMMENT);
	}
	public CEL_COMMENT(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.CEL_COMMENT, i);
	}
	public NUM_FLOAT_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.NUM_FLOAT);
	}
	public NUM_FLOAT(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.NUM_FLOAT, i);
	}
	public NUM_INT_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.NUM_INT);
	}
	public NUM_INT(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.NUM_INT, i);
	}
	public NUM_UINT_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.NUM_UINT);
	}
	public NUM_UINT(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.NUM_UINT, i);
	}
	public STRING_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.STRING);
	}
	public STRING(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.STRING, i);
	}
	public BYTES_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.BYTES);
	}
	public BYTES(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.BYTES, i);
	}
	public NEWLINE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.NEWLINE);
	}
	public NEWLINE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.NEWLINE, i);
	}
	public RBRACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.RBRACE);
	}
	public RBRACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.RBRACE, i);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_conditionExpression;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterConditionExpression) {
	 		listener.enterConditionExpression(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitConditionExpression) {
	 		listener.exitConditionExpression(this);
		}
	}
}
