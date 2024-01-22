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
	public static readonly MODEL = 16;
	public static readonly SCHEMA = 17;
	public static readonly SCHEMA_VERSION = 18;
	public static readonly TYPE = 19;
	public static readonly CONDITION = 20;
	public static readonly RELATIONS = 21;
	public static readonly DEFINE = 22;
	public static readonly KEYWORD_WITH = 23;
	public static readonly EQUALS = 24;
	public static readonly NOT_EQUALS = 25;
	public static readonly IN = 26;
	public static readonly LESS_EQUALS = 27;
	public static readonly GREATER_EQUALS = 28;
	public static readonly LOGICAL_AND = 29;
	public static readonly LOGICAL_OR = 30;
	public static readonly RPRACKET = 31;
	public static readonly LBRACE = 32;
	public static readonly RBRACE = 33;
	public static readonly DOT = 34;
	public static readonly MINUS = 35;
	public static readonly EXCLAM = 36;
	public static readonly QUESTIONMARK = 37;
	public static readonly PLUS = 38;
	public static readonly STAR = 39;
	public static readonly SLASH = 40;
	public static readonly PERCENT = 41;
	public static readonly CEL_TRUE = 42;
	public static readonly CEL_FALSE = 43;
	public static readonly NUL = 44;
	public static readonly CEL_COMMENT = 45;
	public static readonly NUM_FLOAT = 46;
	public static readonly NUM_INT = 47;
	public static readonly NUM_UINT = 48;
	public static readonly STRING = 49;
	public static readonly BYTES = 50;
	public static readonly NEWLINE = 51;
	public static readonly CONDITION_PARAM_CONTAINER = 52;
	public static readonly CONDITION_PARAM_TYPE = 53;
	public static readonly EOF = Token.EOF;
	public static readonly RULE_main = 0;
	public static readonly RULE_modelHeader = 1;
	public static readonly RULE_typeDefs = 2;
	public static readonly RULE_typeDef = 3;
	public static readonly RULE_relationDeclaration = 4;
	public static readonly RULE_relationName = 5;
	public static readonly RULE_relationDef = 6;
	public static readonly RULE_relationDefNoDirect = 7;
	public static readonly RULE_relationDefPartials = 8;
	public static readonly RULE_relationDefGrouping = 9;
	public static readonly RULE_relationRecurse = 10;
	public static readonly RULE_relationRecurseNoDirect = 11;
	public static readonly RULE_relationDefDirectAssignment = 12;
	public static readonly RULE_relationDefRewrite = 13;
	public static readonly RULE_relationDefTypeRestriction = 14;
	public static readonly RULE_relationDefTypeRestrictionBase = 15;
	public static readonly RULE_conditions = 16;
	public static readonly RULE_condition = 17;
	public static readonly RULE_conditionName = 18;
	public static readonly RULE_conditionParameter = 19;
	public static readonly RULE_parameterName = 20;
	public static readonly RULE_parameterType = 21;
	public static readonly RULE_multiLineComment = 22;
	public static readonly RULE_conditionExpression = 23;
	public static readonly literalNames: (string | null)[] = [ null, "':'", 
                                                            "','", "'<'", 
                                                            "'>'", "'['", 
                                                            null, "'('", 
                                                            "')'", null, 
                                                            null, "'#'", 
                                                            "'and'", "'or'", 
                                                            "'but not'", 
                                                            "'from'", "'model'", 
                                                            "'schema'", 
                                                            "'1.1'", "'type'", 
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
                                                             "FROM", "MODEL", 
                                                             "SCHEMA", "SCHEMA_VERSION", 
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
		"main", "modelHeader", "typeDefs", "typeDef", "relationDeclaration", "relationName", 
		"relationDef", "relationDefNoDirect", "relationDefPartials", "relationDefGrouping", 
		"relationRecurse", "relationRecurseNoDirect", "relationDefDirectAssignment", 
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
			this.state = 49;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 48;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 52;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===51) {
				{
				this.state = 51;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 54;
			this.modelHeader();
			this.state = 56;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 2, this._ctx) ) {
			case 1:
				{
				this.state = 55;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 58;
			this.typeDefs();
			this.state = 60;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 3, this._ctx) ) {
			case 1:
				{
				this.state = 59;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 62;
			this.conditions();
			this.state = 64;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===51) {
				{
				this.state = 63;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 66;
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
			this.state = 71;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===11) {
				{
				this.state = 68;
				this.multiLineComment();
				this.state = 69;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 73;
			this.match(OpenFGAParser.MODEL);
			this.state = 74;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 75;
			this.match(OpenFGAParser.SCHEMA);
			this.state = 76;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 77;
			localctx._schemaVersion = this.match(OpenFGAParser.SCHEMA_VERSION);
			this.state = 79;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 78;
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
		this.enterRule(localctx, 4, OpenFGAParser.RULE_typeDefs);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 84;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 7, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 81;
					this.typeDef();
					}
					}
				}
				this.state = 86;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 7, this._ctx);
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
		this.enterRule(localctx, 6, OpenFGAParser.RULE_typeDef);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 89;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 8, this._ctx) ) {
			case 1:
				{
				this.state = 87;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 88;
				this.multiLineComment();
				}
				break;
			}
			this.state = 91;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 92;
			this.match(OpenFGAParser.TYPE);
			this.state = 93;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 94;
			localctx._typeName = this.match(OpenFGAParser.IDENTIFIER);
			this.state = 102;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 10, this._ctx) ) {
			case 1:
				{
				this.state = 95;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 96;
				this.match(OpenFGAParser.RELATIONS);
				this.state = 98;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 97;
						this.relationDeclaration();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					this.state = 100;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 9, this._ctx);
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
		this.enterRule(localctx, 8, OpenFGAParser.RULE_relationDeclaration);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 106;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 11, this._ctx) ) {
			case 1:
				{
				this.state = 104;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 105;
				this.multiLineComment();
				}
				break;
			}
			this.state = 108;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 109;
			this.match(OpenFGAParser.DEFINE);
			this.state = 110;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 111;
			this.relationName();
			this.state = 113;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 112;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 115;
			this.match(OpenFGAParser.COLON);
			this.state = 117;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 116;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			{
			this.state = 119;
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
		this.enterRule(localctx, 10, OpenFGAParser.RULE_relationName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 121;
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
		this.enterRule(localctx, 12, OpenFGAParser.RULE_relationDef);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 126;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 5:
				{
				this.state = 123;
				this.relationDefDirectAssignment();
				}
				break;
			case 10:
				{
				this.state = 124;
				this.relationDefGrouping();
				}
				break;
			case 7:
				{
				this.state = 125;
				this.relationRecurse();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this.state = 129;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 15, this._ctx) ) {
			case 1:
				{
				this.state = 128;
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
		this.enterRule(localctx, 14, OpenFGAParser.RULE_relationDefNoDirect);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 133;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 10:
				{
				this.state = 131;
				this.relationDefGrouping();
				}
				break;
			case 7:
				{
				this.state = 132;
				this.relationRecurseNoDirect();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this.state = 136;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 17, this._ctx) ) {
			case 1:
				{
				this.state = 135;
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
		this.enterRule(localctx, 16, OpenFGAParser.RULE_relationDefPartials);
		try {
			let _alt: number;
			this.state = 167;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 23, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 145;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 138;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 139;
						this.match(OpenFGAParser.OR);
						this.state = 140;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 143;
						this._errHandler.sync(this);
						switch (this._input.LA(1)) {
						case 10:
							{
							this.state = 141;
							this.relationDefGrouping();
							}
							break;
						case 7:
							{
							this.state = 142;
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
					this.state = 147;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 19, this._ctx);
				} while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER);
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 156;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 149;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 150;
						this.match(OpenFGAParser.AND);
						this.state = 151;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 154;
						this._errHandler.sync(this);
						switch (this._input.LA(1)) {
						case 10:
							{
							this.state = 152;
							this.relationDefGrouping();
							}
							break;
						case 7:
							{
							this.state = 153;
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
					this.state = 158;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 21, this._ctx);
				} while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER);
				}
				break;
			case 3:
				this.enterOuterAlt(localctx, 3);
				{
				{
				this.state = 160;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 161;
				this.match(OpenFGAParser.BUT_NOT);
				this.state = 162;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 165;
				this._errHandler.sync(this);
				switch (this._input.LA(1)) {
				case 10:
					{
					this.state = 163;
					this.relationDefGrouping();
					}
					break;
				case 7:
					{
					this.state = 164;
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
		this.enterRule(localctx, 18, OpenFGAParser.RULE_relationDefGrouping);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 169;
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
		this.enterRule(localctx, 20, OpenFGAParser.RULE_relationRecurse);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 171;
			this.match(OpenFGAParser.LPAREN);
			this.state = 175;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 172;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 177;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 180;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 25, this._ctx) ) {
			case 1:
				{
				this.state = 178;
				this.relationDef();
				}
				break;
			case 2:
				{
				this.state = 179;
				this.relationRecurseNoDirect();
				}
				break;
			}
			this.state = 185;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 182;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 187;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 188;
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
		this.enterRule(localctx, 22, OpenFGAParser.RULE_relationRecurseNoDirect);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 190;
			this.match(OpenFGAParser.LPAREN);
			this.state = 194;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 191;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 196;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 199;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 28, this._ctx) ) {
			case 1:
				{
				this.state = 197;
				this.relationDefNoDirect();
				}
				break;
			case 2:
				{
				this.state = 198;
				this.relationRecurseNoDirect();
				}
				break;
			}
			this.state = 204;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 201;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 206;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 207;
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
		this.enterRule(localctx, 24, OpenFGAParser.RULE_relationDefDirectAssignment);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 209;
			this.match(OpenFGAParser.LBRACKET);
			this.state = 211;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 210;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 213;
			this.relationDefTypeRestriction();
			this.state = 215;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 214;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 227;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===2) {
				{
				{
				this.state = 217;
				this.match(OpenFGAParser.COMMA);
				this.state = 219;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 218;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				this.state = 221;
				this.relationDefTypeRestriction();
				this.state = 223;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 222;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				}
				}
				this.state = 229;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 230;
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
		this.enterRule(localctx, 26, OpenFGAParser.RULE_relationDefRewrite);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 232;
			localctx._rewriteComputedusersetName = this.match(OpenFGAParser.IDENTIFIER);
			this.state = 237;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 35, this._ctx) ) {
			case 1:
				{
				this.state = 233;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 234;
				this.match(OpenFGAParser.FROM);
				this.state = 235;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 236;
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
		this.enterRule(localctx, 28, OpenFGAParser.RULE_relationDefTypeRestriction);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 240;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===51) {
				{
				this.state = 239;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 249;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 37, this._ctx) ) {
			case 1:
				{
				this.state = 242;
				this.relationDefTypeRestrictionBase();
				}
				break;
			case 2:
				{
				{
				this.state = 243;
				this.relationDefTypeRestrictionBase();
				this.state = 244;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 245;
				this.match(OpenFGAParser.KEYWORD_WITH);
				this.state = 246;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 247;
				this.conditionName();
				}
				}
				break;
			}
			this.state = 252;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===51) {
				{
				this.state = 251;
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
		this.enterRule(localctx, 30, OpenFGAParser.RULE_relationDefTypeRestrictionBase);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 254;
			localctx._relationDefTypeRestrictionType = this.match(OpenFGAParser.IDENTIFIER);
			this.state = 259;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 1:
				{
				{
				this.state = 255;
				this.match(OpenFGAParser.COLON);
				this.state = 256;
				localctx._relationDefTypeRestrictionWildcard = this.match(OpenFGAParser.STAR);
				}
				}
				break;
			case 11:
				{
				{
				this.state = 257;
				this.match(OpenFGAParser.HASH);
				this.state = 258;
				localctx._relationDefTypeRestrictionRelation = this.match(OpenFGAParser.IDENTIFIER);
				}
				}
				break;
			case 2:
			case 9:
			case 31:
			case 51:
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
		this.enterRule(localctx, 32, OpenFGAParser.RULE_conditions);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 264;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 40, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 261;
					this.condition();
					}
					}
				}
				this.state = 266;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 40, this._ctx);
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
		this.enterRule(localctx, 34, OpenFGAParser.RULE_condition);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 269;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 41, this._ctx) ) {
			case 1:
				{
				this.state = 267;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 268;
				this.multiLineComment();
				}
				break;
			}
			this.state = 271;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 272;
			this.match(OpenFGAParser.CONDITION);
			this.state = 273;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 274;
			this.conditionName();
			this.state = 276;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 275;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 278;
			this.match(OpenFGAParser.LPAREN);
			this.state = 280;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 279;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 282;
			this.conditionParameter();
			this.state = 284;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 283;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 296;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===2) {
				{
				{
				this.state = 286;
				this.match(OpenFGAParser.COMMA);
				this.state = 288;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 287;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				this.state = 290;
				this.conditionParameter();
				this.state = 292;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 291;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				}
				}
				this.state = 298;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 300;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===51) {
				{
				this.state = 299;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 302;
			this.match(OpenFGAParser.RPAREN);
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
			this.match(OpenFGAParser.LBRACE);
			this.state = 308;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 50, this._ctx) ) {
			case 1:
				{
				this.state = 307;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 311;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 51, this._ctx) ) {
			case 1:
				{
				this.state = 310;
				this.match(OpenFGAParser.WHITESPACE);
				}
				break;
			}
			this.state = 313;
			this.conditionExpression();
			this.state = 315;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===51) {
				{
				this.state = 314;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 317;
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
		this.enterRule(localctx, 36, OpenFGAParser.RULE_conditionName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 319;
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
		this.enterRule(localctx, 38, OpenFGAParser.RULE_conditionParameter);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 322;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===51) {
				{
				this.state = 321;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 324;
			this.parameterName();
			this.state = 326;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 325;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 328;
			this.match(OpenFGAParser.COLON);
			this.state = 330;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 329;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 332;
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
		this.enterRule(localctx, 40, OpenFGAParser.RULE_parameterName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 334;
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
		this.enterRule(localctx, 42, OpenFGAParser.RULE_parameterType);
		try {
			this.state = 341;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 53:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 336;
				this.match(OpenFGAParser.CONDITION_PARAM_TYPE);
				}
				break;
			case 52:
				this.enterOuterAlt(localctx, 2);
				{
				{
				this.state = 337;
				this.match(OpenFGAParser.CONDITION_PARAM_CONTAINER);
				this.state = 338;
				this.match(OpenFGAParser.LESS);
				this.state = 339;
				this.match(OpenFGAParser.CONDITION_PARAM_TYPE);
				this.state = 340;
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
		this.enterRule(localctx, 44, OpenFGAParser.RULE_multiLineComment);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 343;
			this.match(OpenFGAParser.HASH);
			this.state = 347;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while ((((_la) & ~0x1F) === 0 && ((1 << _la) & 4294967294) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 3670015) !== 0)) {
				{
				{
				this.state = 344;
				_la = this._input.LA(1);
				if(_la<=0 || _la===51) {
				this._errHandler.recoverInline(this);
				}
				else {
					this._errHandler.reportMatch(this);
				    this.consume();
				}
				}
				}
				this.state = 349;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 352;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 58, this._ctx) ) {
			case 1:
				{
				this.state = 350;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 351;
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
		this.enterRule(localctx, 46, OpenFGAParser.RULE_conditionExpression);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 358;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 60, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					this.state = 356;
					this._errHandler.sync(this);
					switch ( this._interp.adaptivePredict(this._input, 59, this._ctx) ) {
					case 1:
						{
						this.state = 354;
						_la = this._input.LA(1);
						if(!((((_la) & ~0x1F) === 0 && ((1 << _la) & 4278192056) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 1048573) !== 0))) {
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
						this.state = 355;
						_la = this._input.LA(1);
						if(_la<=0 || _la===33) {
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
				this.state = 360;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 60, this._ctx);
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

	public static readonly _serializedATN: number[] = [4,1,53,362,2,0,7,0,2,
	1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,
	10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,16,2,17,
	7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,2,22,7,22,2,23,7,23,1,0,3,
	0,50,8,0,1,0,3,0,53,8,0,1,0,1,0,3,0,57,8,0,1,0,1,0,3,0,61,8,0,1,0,1,0,3,
	0,65,8,0,1,0,1,0,1,1,1,1,1,1,3,1,72,8,1,1,1,1,1,1,1,1,1,1,1,1,1,3,1,80,
	8,1,1,2,5,2,83,8,2,10,2,12,2,86,9,2,1,3,1,3,3,3,90,8,3,1,3,1,3,1,3,1,3,
	1,3,1,3,1,3,4,3,99,8,3,11,3,12,3,100,3,3,103,8,3,1,4,1,4,3,4,107,8,4,1,
	4,1,4,1,4,1,4,1,4,3,4,114,8,4,1,4,1,4,3,4,118,8,4,1,4,1,4,1,5,1,5,1,6,1,
	6,1,6,3,6,127,8,6,1,6,3,6,130,8,6,1,7,1,7,3,7,134,8,7,1,7,3,7,137,8,7,1,
	8,1,8,1,8,1,8,1,8,3,8,144,8,8,4,8,146,8,8,11,8,12,8,147,1,8,1,8,1,8,1,8,
	1,8,3,8,155,8,8,4,8,157,8,8,11,8,12,8,158,1,8,1,8,1,8,1,8,1,8,3,8,166,8,
	8,3,8,168,8,8,1,9,1,9,1,10,1,10,5,10,174,8,10,10,10,12,10,177,9,10,1,10,
	1,10,3,10,181,8,10,1,10,5,10,184,8,10,10,10,12,10,187,9,10,1,10,1,10,1,
	11,1,11,5,11,193,8,11,10,11,12,11,196,9,11,1,11,1,11,3,11,200,8,11,1,11,
	5,11,203,8,11,10,11,12,11,206,9,11,1,11,1,11,1,12,1,12,3,12,212,8,12,1,
	12,1,12,3,12,216,8,12,1,12,1,12,3,12,220,8,12,1,12,1,12,3,12,224,8,12,5,
	12,226,8,12,10,12,12,12,229,9,12,1,12,1,12,1,13,1,13,1,13,1,13,1,13,3,13,
	238,8,13,1,14,3,14,241,8,14,1,14,1,14,1,14,1,14,1,14,1,14,1,14,3,14,250,
	8,14,1,14,3,14,253,8,14,1,15,1,15,1,15,1,15,1,15,3,15,260,8,15,1,16,5,16,
	263,8,16,10,16,12,16,266,9,16,1,17,1,17,3,17,270,8,17,1,17,1,17,1,17,1,
	17,1,17,3,17,277,8,17,1,17,1,17,3,17,281,8,17,1,17,1,17,3,17,285,8,17,1,
	17,1,17,3,17,289,8,17,1,17,1,17,3,17,293,8,17,5,17,295,8,17,10,17,12,17,
	298,9,17,1,17,3,17,301,8,17,1,17,1,17,3,17,305,8,17,1,17,1,17,3,17,309,
	8,17,1,17,3,17,312,8,17,1,17,1,17,3,17,316,8,17,1,17,1,17,1,18,1,18,1,19,
	3,19,323,8,19,1,19,1,19,3,19,327,8,19,1,19,1,19,3,19,331,8,19,1,19,1,19,
	1,20,1,20,1,21,1,21,1,21,1,21,1,21,3,21,342,8,21,1,22,1,22,5,22,346,8,22,
	10,22,12,22,349,9,22,1,22,1,22,3,22,353,8,22,1,23,1,23,5,23,357,8,23,10,
	23,12,23,360,9,23,1,23,0,0,24,0,2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,
	32,34,36,38,40,42,44,46,0,3,1,0,51,51,4,0,3,5,7,10,24,32,34,51,1,0,33,33,
	401,0,49,1,0,0,0,2,71,1,0,0,0,4,84,1,0,0,0,6,89,1,0,0,0,8,106,1,0,0,0,10,
	121,1,0,0,0,12,126,1,0,0,0,14,133,1,0,0,0,16,167,1,0,0,0,18,169,1,0,0,0,
	20,171,1,0,0,0,22,190,1,0,0,0,24,209,1,0,0,0,26,232,1,0,0,0,28,240,1,0,
	0,0,30,254,1,0,0,0,32,264,1,0,0,0,34,269,1,0,0,0,36,319,1,0,0,0,38,322,
	1,0,0,0,40,334,1,0,0,0,42,341,1,0,0,0,44,343,1,0,0,0,46,358,1,0,0,0,48,
	50,5,9,0,0,49,48,1,0,0,0,49,50,1,0,0,0,50,52,1,0,0,0,51,53,5,51,0,0,52,
	51,1,0,0,0,52,53,1,0,0,0,53,54,1,0,0,0,54,56,3,2,1,0,55,57,5,51,0,0,56,
	55,1,0,0,0,56,57,1,0,0,0,57,58,1,0,0,0,58,60,3,4,2,0,59,61,5,51,0,0,60,
	59,1,0,0,0,60,61,1,0,0,0,61,62,1,0,0,0,62,64,3,32,16,0,63,65,5,51,0,0,64,
	63,1,0,0,0,64,65,1,0,0,0,65,66,1,0,0,0,66,67,5,0,0,1,67,1,1,0,0,0,68,69,
	3,44,22,0,69,70,5,51,0,0,70,72,1,0,0,0,71,68,1,0,0,0,71,72,1,0,0,0,72,73,
	1,0,0,0,73,74,5,16,0,0,74,75,5,51,0,0,75,76,5,17,0,0,76,77,5,9,0,0,77,79,
	5,18,0,0,78,80,5,9,0,0,79,78,1,0,0,0,79,80,1,0,0,0,80,3,1,0,0,0,81,83,3,
	6,3,0,82,81,1,0,0,0,83,86,1,0,0,0,84,82,1,0,0,0,84,85,1,0,0,0,85,5,1,0,
	0,0,86,84,1,0,0,0,87,88,5,51,0,0,88,90,3,44,22,0,89,87,1,0,0,0,89,90,1,
	0,0,0,90,91,1,0,0,0,91,92,5,51,0,0,92,93,5,19,0,0,93,94,5,9,0,0,94,102,
	5,10,0,0,95,96,5,51,0,0,96,98,5,21,0,0,97,99,3,8,4,0,98,97,1,0,0,0,99,100,
	1,0,0,0,100,98,1,0,0,0,100,101,1,0,0,0,101,103,1,0,0,0,102,95,1,0,0,0,102,
	103,1,0,0,0,103,7,1,0,0,0,104,105,5,51,0,0,105,107,3,44,22,0,106,104,1,
	0,0,0,106,107,1,0,0,0,107,108,1,0,0,0,108,109,5,51,0,0,109,110,5,22,0,0,
	110,111,5,9,0,0,111,113,3,10,5,0,112,114,5,9,0,0,113,112,1,0,0,0,113,114,
	1,0,0,0,114,115,1,0,0,0,115,117,5,1,0,0,116,118,5,9,0,0,117,116,1,0,0,0,
	117,118,1,0,0,0,118,119,1,0,0,0,119,120,3,12,6,0,120,9,1,0,0,0,121,122,
	5,10,0,0,122,11,1,0,0,0,123,127,3,24,12,0,124,127,3,18,9,0,125,127,3,20,
	10,0,126,123,1,0,0,0,126,124,1,0,0,0,126,125,1,0,0,0,127,129,1,0,0,0,128,
	130,3,16,8,0,129,128,1,0,0,0,129,130,1,0,0,0,130,13,1,0,0,0,131,134,3,18,
	9,0,132,134,3,22,11,0,133,131,1,0,0,0,133,132,1,0,0,0,134,136,1,0,0,0,135,
	137,3,16,8,0,136,135,1,0,0,0,136,137,1,0,0,0,137,15,1,0,0,0,138,139,5,9,
	0,0,139,140,5,13,0,0,140,143,5,9,0,0,141,144,3,18,9,0,142,144,3,22,11,0,
	143,141,1,0,0,0,143,142,1,0,0,0,144,146,1,0,0,0,145,138,1,0,0,0,146,147,
	1,0,0,0,147,145,1,0,0,0,147,148,1,0,0,0,148,168,1,0,0,0,149,150,5,9,0,0,
	150,151,5,12,0,0,151,154,5,9,0,0,152,155,3,18,9,0,153,155,3,22,11,0,154,
	152,1,0,0,0,154,153,1,0,0,0,155,157,1,0,0,0,156,149,1,0,0,0,157,158,1,0,
	0,0,158,156,1,0,0,0,158,159,1,0,0,0,159,168,1,0,0,0,160,161,5,9,0,0,161,
	162,5,14,0,0,162,165,5,9,0,0,163,166,3,18,9,0,164,166,3,22,11,0,165,163,
	1,0,0,0,165,164,1,0,0,0,166,168,1,0,0,0,167,145,1,0,0,0,167,156,1,0,0,0,
	167,160,1,0,0,0,168,17,1,0,0,0,169,170,3,26,13,0,170,19,1,0,0,0,171,175,
	5,7,0,0,172,174,5,9,0,0,173,172,1,0,0,0,174,177,1,0,0,0,175,173,1,0,0,0,
	175,176,1,0,0,0,176,180,1,0,0,0,177,175,1,0,0,0,178,181,3,12,6,0,179,181,
	3,22,11,0,180,178,1,0,0,0,180,179,1,0,0,0,181,185,1,0,0,0,182,184,5,9,0,
	0,183,182,1,0,0,0,184,187,1,0,0,0,185,183,1,0,0,0,185,186,1,0,0,0,186,188,
	1,0,0,0,187,185,1,0,0,0,188,189,5,8,0,0,189,21,1,0,0,0,190,194,5,7,0,0,
	191,193,5,9,0,0,192,191,1,0,0,0,193,196,1,0,0,0,194,192,1,0,0,0,194,195,
	1,0,0,0,195,199,1,0,0,0,196,194,1,0,0,0,197,200,3,14,7,0,198,200,3,22,11,
	0,199,197,1,0,0,0,199,198,1,0,0,0,200,204,1,0,0,0,201,203,5,9,0,0,202,201,
	1,0,0,0,203,206,1,0,0,0,204,202,1,0,0,0,204,205,1,0,0,0,205,207,1,0,0,0,
	206,204,1,0,0,0,207,208,5,8,0,0,208,23,1,0,0,0,209,211,5,5,0,0,210,212,
	5,9,0,0,211,210,1,0,0,0,211,212,1,0,0,0,212,213,1,0,0,0,213,215,3,28,14,
	0,214,216,5,9,0,0,215,214,1,0,0,0,215,216,1,0,0,0,216,227,1,0,0,0,217,219,
	5,2,0,0,218,220,5,9,0,0,219,218,1,0,0,0,219,220,1,0,0,0,220,221,1,0,0,0,
	221,223,3,28,14,0,222,224,5,9,0,0,223,222,1,0,0,0,223,224,1,0,0,0,224,226,
	1,0,0,0,225,217,1,0,0,0,226,229,1,0,0,0,227,225,1,0,0,0,227,228,1,0,0,0,
	228,230,1,0,0,0,229,227,1,0,0,0,230,231,5,31,0,0,231,25,1,0,0,0,232,237,
	5,10,0,0,233,234,5,9,0,0,234,235,5,15,0,0,235,236,5,9,0,0,236,238,5,10,
	0,0,237,233,1,0,0,0,237,238,1,0,0,0,238,27,1,0,0,0,239,241,5,51,0,0,240,
	239,1,0,0,0,240,241,1,0,0,0,241,249,1,0,0,0,242,250,3,30,15,0,243,244,3,
	30,15,0,244,245,5,9,0,0,245,246,5,23,0,0,246,247,5,9,0,0,247,248,3,36,18,
	0,248,250,1,0,0,0,249,242,1,0,0,0,249,243,1,0,0,0,250,252,1,0,0,0,251,253,
	5,51,0,0,252,251,1,0,0,0,252,253,1,0,0,0,253,29,1,0,0,0,254,259,5,10,0,
	0,255,256,5,1,0,0,256,260,5,39,0,0,257,258,5,11,0,0,258,260,5,10,0,0,259,
	255,1,0,0,0,259,257,1,0,0,0,259,260,1,0,0,0,260,31,1,0,0,0,261,263,3,34,
	17,0,262,261,1,0,0,0,263,266,1,0,0,0,264,262,1,0,0,0,264,265,1,0,0,0,265,
	33,1,0,0,0,266,264,1,0,0,0,267,268,5,51,0,0,268,270,3,44,22,0,269,267,1,
	0,0,0,269,270,1,0,0,0,270,271,1,0,0,0,271,272,5,51,0,0,272,273,5,20,0,0,
	273,274,5,9,0,0,274,276,3,36,18,0,275,277,5,9,0,0,276,275,1,0,0,0,276,277,
	1,0,0,0,277,278,1,0,0,0,278,280,5,7,0,0,279,281,5,9,0,0,280,279,1,0,0,0,
	280,281,1,0,0,0,281,282,1,0,0,0,282,284,3,38,19,0,283,285,5,9,0,0,284,283,
	1,0,0,0,284,285,1,0,0,0,285,296,1,0,0,0,286,288,5,2,0,0,287,289,5,9,0,0,
	288,287,1,0,0,0,288,289,1,0,0,0,289,290,1,0,0,0,290,292,3,38,19,0,291,293,
	5,9,0,0,292,291,1,0,0,0,292,293,1,0,0,0,293,295,1,0,0,0,294,286,1,0,0,0,
	295,298,1,0,0,0,296,294,1,0,0,0,296,297,1,0,0,0,297,300,1,0,0,0,298,296,
	1,0,0,0,299,301,5,51,0,0,300,299,1,0,0,0,300,301,1,0,0,0,301,302,1,0,0,
	0,302,304,5,8,0,0,303,305,5,9,0,0,304,303,1,0,0,0,304,305,1,0,0,0,305,306,
	1,0,0,0,306,308,5,32,0,0,307,309,5,51,0,0,308,307,1,0,0,0,308,309,1,0,0,
	0,309,311,1,0,0,0,310,312,5,9,0,0,311,310,1,0,0,0,311,312,1,0,0,0,312,313,
	1,0,0,0,313,315,3,46,23,0,314,316,5,51,0,0,315,314,1,0,0,0,315,316,1,0,
	0,0,316,317,1,0,0,0,317,318,5,33,0,0,318,35,1,0,0,0,319,320,5,10,0,0,320,
	37,1,0,0,0,321,323,5,51,0,0,322,321,1,0,0,0,322,323,1,0,0,0,323,324,1,0,
	0,0,324,326,3,40,20,0,325,327,5,9,0,0,326,325,1,0,0,0,326,327,1,0,0,0,327,
	328,1,0,0,0,328,330,5,1,0,0,329,331,5,9,0,0,330,329,1,0,0,0,330,331,1,0,
	0,0,331,332,1,0,0,0,332,333,3,42,21,0,333,39,1,0,0,0,334,335,5,10,0,0,335,
	41,1,0,0,0,336,342,5,53,0,0,337,338,5,52,0,0,338,339,5,3,0,0,339,340,5,
	53,0,0,340,342,5,4,0,0,341,336,1,0,0,0,341,337,1,0,0,0,342,43,1,0,0,0,343,
	347,5,11,0,0,344,346,8,0,0,0,345,344,1,0,0,0,346,349,1,0,0,0,347,345,1,
	0,0,0,347,348,1,0,0,0,348,352,1,0,0,0,349,347,1,0,0,0,350,351,5,51,0,0,
	351,353,3,44,22,0,352,350,1,0,0,0,352,353,1,0,0,0,353,45,1,0,0,0,354,357,
	7,1,0,0,355,357,8,2,0,0,356,354,1,0,0,0,356,355,1,0,0,0,357,360,1,0,0,0,
	358,356,1,0,0,0,358,359,1,0,0,0,359,47,1,0,0,0,360,358,1,0,0,0,61,49,52,
	56,60,64,71,79,84,89,100,102,106,113,117,126,129,133,136,143,147,154,158,
	165,167,175,180,185,194,199,204,211,215,219,223,227,237,240,249,252,259,
	264,269,276,280,284,288,292,296,300,304,308,311,315,322,326,330,341,347,
	352,356,358];

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
	public modelHeader(): ModelHeaderContext {
		return this.getTypedRuleContext(ModelHeaderContext, 0) as ModelHeaderContext;
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
