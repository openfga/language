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
	public static readonly EXTEND = 20;
	public static readonly TYPE = 21;
	public static readonly CONDITION = 22;
	public static readonly RELATIONS = 23;
	public static readonly DEFINE = 24;
	public static readonly KEYWORD_WITH = 25;
	public static readonly EQUALS = 26;
	public static readonly NOT_EQUALS = 27;
	public static readonly IN = 28;
	public static readonly LESS_EQUALS = 29;
	public static readonly GREATER_EQUALS = 30;
	public static readonly LOGICAL_AND = 31;
	public static readonly LOGICAL_OR = 32;
	public static readonly RPRACKET = 33;
	public static readonly LBRACE = 34;
	public static readonly RBRACE = 35;
	public static readonly DOT = 36;
	public static readonly MINUS = 37;
	public static readonly EXCLAM = 38;
	public static readonly QUESTIONMARK = 39;
	public static readonly PLUS = 40;
	public static readonly STAR = 41;
	public static readonly SLASH = 42;
	public static readonly PERCENT = 43;
	public static readonly CEL_TRUE = 44;
	public static readonly CEL_FALSE = 45;
	public static readonly NUL = 46;
	public static readonly CEL_COMMENT = 47;
	public static readonly NUM_FLOAT = 48;
	public static readonly NUM_INT = 49;
	public static readonly NUM_UINT = 50;
	public static readonly STRING = 51;
	public static readonly BYTES = 52;
	public static readonly NEWLINE = 53;
	public static readonly CONDITION_PARAM_CONTAINER = 54;
	public static readonly CONDITION_PARAM_TYPE = 55;
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
                                                            null, "'extend'", 
                                                            "'type'", "'condition'", 
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
                                                             "EXTEND", "TYPE", 
                                                             "CONDITION", 
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
			if (_la===53) {
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
			if (_la===53) {
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
		let _la: number;
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
			this.state = 110;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 108;
				this.match(OpenFGAParser.EXTEND);
				this.state = 109;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 112;
			this.match(OpenFGAParser.TYPE);
			this.state = 113;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 114;
			localctx._typeName = this.match(OpenFGAParser.IDENTIFIER);
			this.state = 122;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 14, this._ctx) ) {
			case 1:
				{
				this.state = 115;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 116;
				this.match(OpenFGAParser.RELATIONS);
				this.state = 118;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 117;
						this.relationDeclaration();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					this.state = 120;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 13, this._ctx);
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
			this.state = 126;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 15, this._ctx) ) {
			case 1:
				{
				this.state = 124;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 125;
				this.multiLineComment();
				}
				break;
			}
			this.state = 128;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 129;
			this.match(OpenFGAParser.DEFINE);
			this.state = 130;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 131;
			this.relationName();
			this.state = 133;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 132;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 135;
			this.match(OpenFGAParser.COLON);
			this.state = 137;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 136;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			{
			this.state = 139;
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
			this.state = 141;
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
			this.state = 146;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 5:
				{
				this.state = 143;
				this.relationDefDirectAssignment();
				}
				break;
			case 10:
				{
				this.state = 144;
				this.relationDefGrouping();
				}
				break;
			case 7:
				{
				this.state = 145;
				this.relationRecurse();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this.state = 149;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 19, this._ctx) ) {
			case 1:
				{
				this.state = 148;
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
			this.state = 153;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 10:
				{
				this.state = 151;
				this.relationDefGrouping();
				}
				break;
			case 7:
				{
				this.state = 152;
				this.relationRecurseNoDirect();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this.state = 156;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 21, this._ctx) ) {
			case 1:
				{
				this.state = 155;
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
			this.state = 187;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 27, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 165;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 158;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 159;
						this.match(OpenFGAParser.OR);
						this.state = 160;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 163;
						this._errHandler.sync(this);
						switch (this._input.LA(1)) {
						case 10:
							{
							this.state = 161;
							this.relationDefGrouping();
							}
							break;
						case 7:
							{
							this.state = 162;
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
					this.state = 167;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 23, this._ctx);
				} while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER);
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 176;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 169;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 170;
						this.match(OpenFGAParser.AND);
						this.state = 171;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 174;
						this._errHandler.sync(this);
						switch (this._input.LA(1)) {
						case 10:
							{
							this.state = 172;
							this.relationDefGrouping();
							}
							break;
						case 7:
							{
							this.state = 173;
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
					this.state = 178;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 25, this._ctx);
				} while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER);
				}
				break;
			case 3:
				this.enterOuterAlt(localctx, 3);
				{
				{
				this.state = 180;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 181;
				this.match(OpenFGAParser.BUT_NOT);
				this.state = 182;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 185;
				this._errHandler.sync(this);
				switch (this._input.LA(1)) {
				case 10:
					{
					this.state = 183;
					this.relationDefGrouping();
					}
					break;
				case 7:
					{
					this.state = 184;
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
			this.state = 189;
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
			this.state = 191;
			this.match(OpenFGAParser.LPAREN);
			this.state = 195;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 192;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 197;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 200;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 29, this._ctx) ) {
			case 1:
				{
				this.state = 198;
				this.relationDef();
				}
				break;
			case 2:
				{
				this.state = 199;
				this.relationRecurseNoDirect();
				}
				break;
			}
			this.state = 205;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 202;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 207;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 208;
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
			this.state = 210;
			this.match(OpenFGAParser.LPAREN);
			this.state = 214;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 211;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 216;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 219;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 32, this._ctx) ) {
			case 1:
				{
				this.state = 217;
				this.relationDefNoDirect();
				}
				break;
			case 2:
				{
				this.state = 218;
				this.relationRecurseNoDirect();
				}
				break;
			}
			this.state = 224;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 221;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 226;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 227;
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
			this.state = 229;
			this.match(OpenFGAParser.LBRACKET);
			this.state = 231;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 230;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 233;
			this.relationDefTypeRestriction();
			this.state = 235;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 234;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 247;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===2) {
				{
				{
				this.state = 237;
				this.match(OpenFGAParser.COMMA);
				this.state = 239;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 238;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				this.state = 241;
				this.relationDefTypeRestriction();
				this.state = 243;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 242;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				}
				}
				this.state = 249;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 250;
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
			this.state = 252;
			localctx._rewriteComputedusersetName = this.match(OpenFGAParser.IDENTIFIER);
			this.state = 257;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 39, this._ctx) ) {
			case 1:
				{
				this.state = 253;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 254;
				this.match(OpenFGAParser.FROM);
				this.state = 255;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 256;
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
			this.state = 260;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===53) {
				{
				this.state = 259;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 269;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 41, this._ctx) ) {
			case 1:
				{
				this.state = 262;
				this.relationDefTypeRestrictionBase();
				}
				break;
			case 2:
				{
				{
				this.state = 263;
				this.relationDefTypeRestrictionBase();
				this.state = 264;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 265;
				this.match(OpenFGAParser.KEYWORD_WITH);
				this.state = 266;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 267;
				this.conditionName();
				}
				}
				break;
			}
			this.state = 272;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===53) {
				{
				this.state = 271;
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
			this.state = 274;
			localctx._relationDefTypeRestrictionType = this.match(OpenFGAParser.IDENTIFIER);
			this.state = 279;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 1:
				{
				{
				this.state = 275;
				this.match(OpenFGAParser.COLON);
				this.state = 276;
				localctx._relationDefTypeRestrictionWildcard = this.match(OpenFGAParser.STAR);
				}
				}
				break;
			case 11:
				{
				{
				this.state = 277;
				this.match(OpenFGAParser.HASH);
				this.state = 278;
				localctx._relationDefTypeRestrictionRelation = this.match(OpenFGAParser.IDENTIFIER);
				}
				}
				break;
			case 2:
			case 9:
			case 33:
			case 53:
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
			this.state = 284;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 44, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 281;
					this.condition();
					}
					}
				}
				this.state = 286;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 44, this._ctx);
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
			this.state = 289;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 45, this._ctx) ) {
			case 1:
				{
				this.state = 287;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 288;
				this.multiLineComment();
				}
				break;
			}
			this.state = 291;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 292;
			this.match(OpenFGAParser.CONDITION);
			this.state = 293;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 294;
			this.conditionName();
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
			this.match(OpenFGAParser.LPAREN);
			this.state = 300;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 299;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 302;
			this.conditionParameter();
			this.state = 304;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 303;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 316;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===2) {
				{
				{
				this.state = 306;
				this.match(OpenFGAParser.COMMA);
				this.state = 308;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 307;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				this.state = 310;
				this.conditionParameter();
				this.state = 312;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 311;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				}
				}
				this.state = 318;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 320;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===53) {
				{
				this.state = 319;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 322;
			this.match(OpenFGAParser.RPAREN);
			this.state = 324;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 323;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 326;
			this.match(OpenFGAParser.LBRACE);
			this.state = 328;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 54, this._ctx) ) {
			case 1:
				{
				this.state = 327;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 331;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 55, this._ctx) ) {
			case 1:
				{
				this.state = 330;
				this.match(OpenFGAParser.WHITESPACE);
				}
				break;
			}
			this.state = 333;
			this.conditionExpression();
			this.state = 335;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===53) {
				{
				this.state = 334;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 337;
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
			this.state = 339;
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
			this.state = 342;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===53) {
				{
				this.state = 341;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 344;
			this.parameterName();
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
			this.match(OpenFGAParser.COLON);
			this.state = 350;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 349;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 352;
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
			this.state = 354;
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
			this.state = 361;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 55:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 356;
				this.match(OpenFGAParser.CONDITION_PARAM_TYPE);
				}
				break;
			case 54:
				this.enterOuterAlt(localctx, 2);
				{
				{
				this.state = 357;
				this.match(OpenFGAParser.CONDITION_PARAM_CONTAINER);
				this.state = 358;
				this.match(OpenFGAParser.LESS);
				this.state = 359;
				this.match(OpenFGAParser.CONDITION_PARAM_TYPE);
				this.state = 360;
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
			this.state = 363;
			this.match(OpenFGAParser.HASH);
			this.state = 367;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while ((((_la) & ~0x1F) === 0 && ((1 << _la) & 4294967294) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 14680063) !== 0)) {
				{
				{
				this.state = 364;
				_la = this._input.LA(1);
				if(_la<=0 || _la===53) {
				this._errHandler.recoverInline(this);
				}
				else {
					this._errHandler.reportMatch(this);
				    this.consume();
				}
				}
				}
				this.state = 369;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 372;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 62, this._ctx) ) {
			case 1:
				{
				this.state = 370;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 371;
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
			this.state = 378;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 64, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					this.state = 376;
					this._errHandler.sync(this);
					switch ( this._interp.adaptivePredict(this._input, 63, this._ctx) ) {
					case 1:
						{
						this.state = 374;
						_la = this._input.LA(1);
						if(!((((_la) & ~0x1F) === 0 && ((1 << _la) & 4227860408) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 4194295) !== 0))) {
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
						this.state = 375;
						_la = this._input.LA(1);
						if(_la<=0 || _la===35) {
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
				this.state = 380;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 64, this._ctx);
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

	public static readonly _serializedATN: number[] = [4,1,55,382,2,0,7,0,2,
	1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,
	10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,16,2,17,
	7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,2,22,7,22,2,23,7,23,2,24,7,
	24,1,0,3,0,52,8,0,1,0,3,0,55,8,0,1,0,1,0,3,0,59,8,0,1,0,3,0,62,8,0,1,0,
	1,0,3,0,66,8,0,1,0,1,0,3,0,70,8,0,1,0,1,0,1,1,1,1,1,1,3,1,77,8,1,1,1,1,
	1,1,1,1,1,1,1,1,1,3,1,85,8,1,1,2,1,2,1,2,3,2,90,8,2,1,2,1,2,1,2,1,2,3,2,
	96,8,2,1,3,5,3,99,8,3,10,3,12,3,102,9,3,1,4,1,4,3,4,106,8,4,1,4,1,4,1,4,
	3,4,111,8,4,1,4,1,4,1,4,1,4,1,4,1,4,4,4,119,8,4,11,4,12,4,120,3,4,123,8,
	4,1,5,1,5,3,5,127,8,5,1,5,1,5,1,5,1,5,1,5,3,5,134,8,5,1,5,1,5,3,5,138,8,
	5,1,5,1,5,1,6,1,6,1,7,1,7,1,7,3,7,147,8,7,1,7,3,7,150,8,7,1,8,1,8,3,8,154,
	8,8,1,8,3,8,157,8,8,1,9,1,9,1,9,1,9,1,9,3,9,164,8,9,4,9,166,8,9,11,9,12,
	9,167,1,9,1,9,1,9,1,9,1,9,3,9,175,8,9,4,9,177,8,9,11,9,12,9,178,1,9,1,9,
	1,9,1,9,1,9,3,9,186,8,9,3,9,188,8,9,1,10,1,10,1,11,1,11,5,11,194,8,11,10,
	11,12,11,197,9,11,1,11,1,11,3,11,201,8,11,1,11,5,11,204,8,11,10,11,12,11,
	207,9,11,1,11,1,11,1,12,1,12,5,12,213,8,12,10,12,12,12,216,9,12,1,12,1,
	12,3,12,220,8,12,1,12,5,12,223,8,12,10,12,12,12,226,9,12,1,12,1,12,1,13,
	1,13,3,13,232,8,13,1,13,1,13,3,13,236,8,13,1,13,1,13,3,13,240,8,13,1,13,
	1,13,3,13,244,8,13,5,13,246,8,13,10,13,12,13,249,9,13,1,13,1,13,1,14,1,
	14,1,14,1,14,1,14,3,14,258,8,14,1,15,3,15,261,8,15,1,15,1,15,1,15,1,15,
	1,15,1,15,1,15,3,15,270,8,15,1,15,3,15,273,8,15,1,16,1,16,1,16,1,16,1,16,
	3,16,280,8,16,1,17,5,17,283,8,17,10,17,12,17,286,9,17,1,18,1,18,3,18,290,
	8,18,1,18,1,18,1,18,1,18,1,18,3,18,297,8,18,1,18,1,18,3,18,301,8,18,1,18,
	1,18,3,18,305,8,18,1,18,1,18,3,18,309,8,18,1,18,1,18,3,18,313,8,18,5,18,
	315,8,18,10,18,12,18,318,9,18,1,18,3,18,321,8,18,1,18,1,18,3,18,325,8,18,
	1,18,1,18,3,18,329,8,18,1,18,3,18,332,8,18,1,18,1,18,3,18,336,8,18,1,18,
	1,18,1,19,1,19,1,20,3,20,343,8,20,1,20,1,20,3,20,347,8,20,1,20,1,20,3,20,
	351,8,20,1,20,1,20,1,21,1,21,1,22,1,22,1,22,1,22,1,22,3,22,362,8,22,1,23,
	1,23,5,23,366,8,23,10,23,12,23,369,9,23,1,23,1,23,3,23,373,8,23,1,24,1,
	24,5,24,377,8,24,10,24,12,24,380,9,24,1,24,0,0,25,0,2,4,6,8,10,12,14,16,
	18,20,22,24,26,28,30,32,34,36,38,40,42,44,46,48,0,3,1,0,53,53,4,0,3,5,7,
	10,26,34,36,53,1,0,35,35,424,0,51,1,0,0,0,2,76,1,0,0,0,4,89,1,0,0,0,6,100,
	1,0,0,0,8,105,1,0,0,0,10,126,1,0,0,0,12,141,1,0,0,0,14,146,1,0,0,0,16,153,
	1,0,0,0,18,187,1,0,0,0,20,189,1,0,0,0,22,191,1,0,0,0,24,210,1,0,0,0,26,
	229,1,0,0,0,28,252,1,0,0,0,30,260,1,0,0,0,32,274,1,0,0,0,34,284,1,0,0,0,
	36,289,1,0,0,0,38,339,1,0,0,0,40,342,1,0,0,0,42,354,1,0,0,0,44,361,1,0,
	0,0,46,363,1,0,0,0,48,378,1,0,0,0,50,52,5,9,0,0,51,50,1,0,0,0,51,52,1,0,
	0,0,52,54,1,0,0,0,53,55,5,53,0,0,54,53,1,0,0,0,54,55,1,0,0,0,55,58,1,0,
	0,0,56,59,3,2,1,0,57,59,3,4,2,0,58,56,1,0,0,0,58,57,1,0,0,0,59,61,1,0,0,
	0,60,62,5,53,0,0,61,60,1,0,0,0,61,62,1,0,0,0,62,63,1,0,0,0,63,65,3,6,3,
	0,64,66,5,53,0,0,65,64,1,0,0,0,65,66,1,0,0,0,66,67,1,0,0,0,67,69,3,34,17,
	0,68,70,5,53,0,0,69,68,1,0,0,0,69,70,1,0,0,0,70,71,1,0,0,0,71,72,5,0,0,
	1,72,1,1,0,0,0,73,74,3,46,23,0,74,75,5,53,0,0,75,77,1,0,0,0,76,73,1,0,0,
	0,76,77,1,0,0,0,77,78,1,0,0,0,78,79,5,17,0,0,79,80,5,53,0,0,80,81,5,18,
	0,0,81,82,5,9,0,0,82,84,5,19,0,0,83,85,5,9,0,0,84,83,1,0,0,0,84,85,1,0,
	0,0,85,3,1,0,0,0,86,87,3,46,23,0,87,88,5,53,0,0,88,90,1,0,0,0,89,86,1,0,
	0,0,89,90,1,0,0,0,90,91,1,0,0,0,91,92,5,16,0,0,92,93,5,9,0,0,93,95,5,10,
	0,0,94,96,5,9,0,0,95,94,1,0,0,0,95,96,1,0,0,0,96,5,1,0,0,0,97,99,3,8,4,
	0,98,97,1,0,0,0,99,102,1,0,0,0,100,98,1,0,0,0,100,101,1,0,0,0,101,7,1,0,
	0,0,102,100,1,0,0,0,103,104,5,53,0,0,104,106,3,46,23,0,105,103,1,0,0,0,
	105,106,1,0,0,0,106,107,1,0,0,0,107,110,5,53,0,0,108,109,5,20,0,0,109,111,
	5,9,0,0,110,108,1,0,0,0,110,111,1,0,0,0,111,112,1,0,0,0,112,113,5,21,0,
	0,113,114,5,9,0,0,114,122,5,10,0,0,115,116,5,53,0,0,116,118,5,23,0,0,117,
	119,3,10,5,0,118,117,1,0,0,0,119,120,1,0,0,0,120,118,1,0,0,0,120,121,1,
	0,0,0,121,123,1,0,0,0,122,115,1,0,0,0,122,123,1,0,0,0,123,9,1,0,0,0,124,
	125,5,53,0,0,125,127,3,46,23,0,126,124,1,0,0,0,126,127,1,0,0,0,127,128,
	1,0,0,0,128,129,5,53,0,0,129,130,5,24,0,0,130,131,5,9,0,0,131,133,3,12,
	6,0,132,134,5,9,0,0,133,132,1,0,0,0,133,134,1,0,0,0,134,135,1,0,0,0,135,
	137,5,1,0,0,136,138,5,9,0,0,137,136,1,0,0,0,137,138,1,0,0,0,138,139,1,0,
	0,0,139,140,3,14,7,0,140,11,1,0,0,0,141,142,5,10,0,0,142,13,1,0,0,0,143,
	147,3,26,13,0,144,147,3,20,10,0,145,147,3,22,11,0,146,143,1,0,0,0,146,144,
	1,0,0,0,146,145,1,0,0,0,147,149,1,0,0,0,148,150,3,18,9,0,149,148,1,0,0,
	0,149,150,1,0,0,0,150,15,1,0,0,0,151,154,3,20,10,0,152,154,3,24,12,0,153,
	151,1,0,0,0,153,152,1,0,0,0,154,156,1,0,0,0,155,157,3,18,9,0,156,155,1,
	0,0,0,156,157,1,0,0,0,157,17,1,0,0,0,158,159,5,9,0,0,159,160,5,13,0,0,160,
	163,5,9,0,0,161,164,3,20,10,0,162,164,3,24,12,0,163,161,1,0,0,0,163,162,
	1,0,0,0,164,166,1,0,0,0,165,158,1,0,0,0,166,167,1,0,0,0,167,165,1,0,0,0,
	167,168,1,0,0,0,168,188,1,0,0,0,169,170,5,9,0,0,170,171,5,12,0,0,171,174,
	5,9,0,0,172,175,3,20,10,0,173,175,3,24,12,0,174,172,1,0,0,0,174,173,1,0,
	0,0,175,177,1,0,0,0,176,169,1,0,0,0,177,178,1,0,0,0,178,176,1,0,0,0,178,
	179,1,0,0,0,179,188,1,0,0,0,180,181,5,9,0,0,181,182,5,14,0,0,182,185,5,
	9,0,0,183,186,3,20,10,0,184,186,3,24,12,0,185,183,1,0,0,0,185,184,1,0,0,
	0,186,188,1,0,0,0,187,165,1,0,0,0,187,176,1,0,0,0,187,180,1,0,0,0,188,19,
	1,0,0,0,189,190,3,28,14,0,190,21,1,0,0,0,191,195,5,7,0,0,192,194,5,9,0,
	0,193,192,1,0,0,0,194,197,1,0,0,0,195,193,1,0,0,0,195,196,1,0,0,0,196,200,
	1,0,0,0,197,195,1,0,0,0,198,201,3,14,7,0,199,201,3,24,12,0,200,198,1,0,
	0,0,200,199,1,0,0,0,201,205,1,0,0,0,202,204,5,9,0,0,203,202,1,0,0,0,204,
	207,1,0,0,0,205,203,1,0,0,0,205,206,1,0,0,0,206,208,1,0,0,0,207,205,1,0,
	0,0,208,209,5,8,0,0,209,23,1,0,0,0,210,214,5,7,0,0,211,213,5,9,0,0,212,
	211,1,0,0,0,213,216,1,0,0,0,214,212,1,0,0,0,214,215,1,0,0,0,215,219,1,0,
	0,0,216,214,1,0,0,0,217,220,3,16,8,0,218,220,3,24,12,0,219,217,1,0,0,0,
	219,218,1,0,0,0,220,224,1,0,0,0,221,223,5,9,0,0,222,221,1,0,0,0,223,226,
	1,0,0,0,224,222,1,0,0,0,224,225,1,0,0,0,225,227,1,0,0,0,226,224,1,0,0,0,
	227,228,5,8,0,0,228,25,1,0,0,0,229,231,5,5,0,0,230,232,5,9,0,0,231,230,
	1,0,0,0,231,232,1,0,0,0,232,233,1,0,0,0,233,235,3,30,15,0,234,236,5,9,0,
	0,235,234,1,0,0,0,235,236,1,0,0,0,236,247,1,0,0,0,237,239,5,2,0,0,238,240,
	5,9,0,0,239,238,1,0,0,0,239,240,1,0,0,0,240,241,1,0,0,0,241,243,3,30,15,
	0,242,244,5,9,0,0,243,242,1,0,0,0,243,244,1,0,0,0,244,246,1,0,0,0,245,237,
	1,0,0,0,246,249,1,0,0,0,247,245,1,0,0,0,247,248,1,0,0,0,248,250,1,0,0,0,
	249,247,1,0,0,0,250,251,5,33,0,0,251,27,1,0,0,0,252,257,5,10,0,0,253,254,
	5,9,0,0,254,255,5,15,0,0,255,256,5,9,0,0,256,258,5,10,0,0,257,253,1,0,0,
	0,257,258,1,0,0,0,258,29,1,0,0,0,259,261,5,53,0,0,260,259,1,0,0,0,260,261,
	1,0,0,0,261,269,1,0,0,0,262,270,3,32,16,0,263,264,3,32,16,0,264,265,5,9,
	0,0,265,266,5,25,0,0,266,267,5,9,0,0,267,268,3,38,19,0,268,270,1,0,0,0,
	269,262,1,0,0,0,269,263,1,0,0,0,270,272,1,0,0,0,271,273,5,53,0,0,272,271,
	1,0,0,0,272,273,1,0,0,0,273,31,1,0,0,0,274,279,5,10,0,0,275,276,5,1,0,0,
	276,280,5,41,0,0,277,278,5,11,0,0,278,280,5,10,0,0,279,275,1,0,0,0,279,
	277,1,0,0,0,279,280,1,0,0,0,280,33,1,0,0,0,281,283,3,36,18,0,282,281,1,
	0,0,0,283,286,1,0,0,0,284,282,1,0,0,0,284,285,1,0,0,0,285,35,1,0,0,0,286,
	284,1,0,0,0,287,288,5,53,0,0,288,290,3,46,23,0,289,287,1,0,0,0,289,290,
	1,0,0,0,290,291,1,0,0,0,291,292,5,53,0,0,292,293,5,22,0,0,293,294,5,9,0,
	0,294,296,3,38,19,0,295,297,5,9,0,0,296,295,1,0,0,0,296,297,1,0,0,0,297,
	298,1,0,0,0,298,300,5,7,0,0,299,301,5,9,0,0,300,299,1,0,0,0,300,301,1,0,
	0,0,301,302,1,0,0,0,302,304,3,40,20,0,303,305,5,9,0,0,304,303,1,0,0,0,304,
	305,1,0,0,0,305,316,1,0,0,0,306,308,5,2,0,0,307,309,5,9,0,0,308,307,1,0,
	0,0,308,309,1,0,0,0,309,310,1,0,0,0,310,312,3,40,20,0,311,313,5,9,0,0,312,
	311,1,0,0,0,312,313,1,0,0,0,313,315,1,0,0,0,314,306,1,0,0,0,315,318,1,0,
	0,0,316,314,1,0,0,0,316,317,1,0,0,0,317,320,1,0,0,0,318,316,1,0,0,0,319,
	321,5,53,0,0,320,319,1,0,0,0,320,321,1,0,0,0,321,322,1,0,0,0,322,324,5,
	8,0,0,323,325,5,9,0,0,324,323,1,0,0,0,324,325,1,0,0,0,325,326,1,0,0,0,326,
	328,5,34,0,0,327,329,5,53,0,0,328,327,1,0,0,0,328,329,1,0,0,0,329,331,1,
	0,0,0,330,332,5,9,0,0,331,330,1,0,0,0,331,332,1,0,0,0,332,333,1,0,0,0,333,
	335,3,48,24,0,334,336,5,53,0,0,335,334,1,0,0,0,335,336,1,0,0,0,336,337,
	1,0,0,0,337,338,5,35,0,0,338,37,1,0,0,0,339,340,5,10,0,0,340,39,1,0,0,0,
	341,343,5,53,0,0,342,341,1,0,0,0,342,343,1,0,0,0,343,344,1,0,0,0,344,346,
	3,42,21,0,345,347,5,9,0,0,346,345,1,0,0,0,346,347,1,0,0,0,347,348,1,0,0,
	0,348,350,5,1,0,0,349,351,5,9,0,0,350,349,1,0,0,0,350,351,1,0,0,0,351,352,
	1,0,0,0,352,353,3,44,22,0,353,41,1,0,0,0,354,355,5,10,0,0,355,43,1,0,0,
	0,356,362,5,55,0,0,357,358,5,54,0,0,358,359,5,3,0,0,359,360,5,55,0,0,360,
	362,5,4,0,0,361,356,1,0,0,0,361,357,1,0,0,0,362,45,1,0,0,0,363,367,5,11,
	0,0,364,366,8,0,0,0,365,364,1,0,0,0,366,369,1,0,0,0,367,365,1,0,0,0,367,
	368,1,0,0,0,368,372,1,0,0,0,369,367,1,0,0,0,370,371,5,53,0,0,371,373,3,
	46,23,0,372,370,1,0,0,0,372,373,1,0,0,0,373,47,1,0,0,0,374,377,7,1,0,0,
	375,377,8,2,0,0,376,374,1,0,0,0,376,375,1,0,0,0,377,380,1,0,0,0,378,376,
	1,0,0,0,378,379,1,0,0,0,379,49,1,0,0,0,380,378,1,0,0,0,65,51,54,58,61,65,
	69,76,84,89,95,100,105,110,120,122,126,133,137,146,149,153,156,163,167,
	174,178,185,187,195,200,205,214,219,224,231,235,239,243,247,257,260,269,
	272,279,284,289,296,300,304,308,312,316,320,324,328,331,335,342,346,350,
	361,367,372,376,378];

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
	public EXTEND(): TerminalNode {
		return this.getToken(OpenFGAParser.EXTEND, 0);
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
