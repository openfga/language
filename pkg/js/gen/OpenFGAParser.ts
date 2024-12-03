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
	public static readonly MIXIN = 23;
	public static readonly INCLUDE = 24;
	public static readonly RELATIONS = 25;
	public static readonly RELATION = 26;
	public static readonly DEFINE = 27;
	public static readonly KEYWORD_WITH = 28;
	public static readonly EQUALS = 29;
	public static readonly NOT_EQUALS = 30;
	public static readonly IN = 31;
	public static readonly LESS_EQUALS = 32;
	public static readonly GREATER_EQUALS = 33;
	public static readonly LOGICAL_AND = 34;
	public static readonly LOGICAL_OR = 35;
	public static readonly RPRACKET = 36;
	public static readonly LBRACE = 37;
	public static readonly RBRACE = 38;
	public static readonly DOT = 39;
	public static readonly MINUS = 40;
	public static readonly EXCLAM = 41;
	public static readonly QUESTIONMARK = 42;
	public static readonly PLUS = 43;
	public static readonly STAR = 44;
	public static readonly SLASH = 45;
	public static readonly PERCENT = 46;
	public static readonly CEL_TRUE = 47;
	public static readonly CEL_FALSE = 48;
	public static readonly NUL = 49;
	public static readonly CEL_COMMENT = 50;
	public static readonly NUM_FLOAT = 51;
	public static readonly NUM_INT = 52;
	public static readonly NUM_UINT = 53;
	public static readonly STRING = 54;
	public static readonly BYTES = 55;
	public static readonly NEWLINE = 56;
	public static readonly CONDITION_PARAM_CONTAINER = 57;
	public static readonly CONDITION_PARAM_TYPE = 58;
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
	public static readonly RULE_mixins = 23;
	public static readonly RULE_mixin = 24;
	public static readonly RULE_mixinName = 25;
	public static readonly RULE_multiLineComment = 26;
	public static readonly RULE_identifier = 27;
	public static readonly RULE_conditionExpression = 28;
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
                                                            "'mixin'", "'include'", 
                                                            "'relations'", 
                                                            "'relation'", 
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
                                                             "MIXIN", "INCLUDE", 
                                                             "RELATIONS", 
                                                             "RELATION", 
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
		"parameterType", "mixins", "mixin", "mixinName", "multiLineComment", "identifier", 
		"conditionExpression",
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
			this.state = 59;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 58;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 62;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===56) {
				{
				this.state = 61;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 66;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 2, this._ctx) ) {
			case 1:
				{
				this.state = 64;
				this.modelHeader();
				}
				break;
			case 2:
				{
				this.state = 65;
				this.moduleHeader();
				}
				break;
			}
			this.state = 69;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 3, this._ctx) ) {
			case 1:
				{
				this.state = 68;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 71;
			this.mixins();
			this.state = 73;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 4, this._ctx) ) {
			case 1:
				{
				this.state = 72;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 76;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 5, this._ctx) ) {
			case 1:
				{
				this.state = 75;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 78;
			this.typeDefs();
			this.state = 80;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 6, this._ctx) ) {
			case 1:
				{
				this.state = 79;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 82;
			this.conditions();
			this.state = 84;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===56) {
				{
				this.state = 83;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 86;
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
			this.state = 91;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===11) {
				{
				this.state = 88;
				this.multiLineComment();
				this.state = 89;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 93;
			this.match(OpenFGAParser.MODEL);
			this.state = 94;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 95;
			this.match(OpenFGAParser.SCHEMA);
			this.state = 96;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 97;
			localctx._schemaVersion = this.match(OpenFGAParser.SCHEMA_VERSION);
			this.state = 99;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 98;
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
			this.state = 104;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===11) {
				{
				this.state = 101;
				this.multiLineComment();
				this.state = 102;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 106;
			this.match(OpenFGAParser.MODULE);
			this.state = 107;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 108;
			localctx._moduleName = this.identifier();
			this.state = 110;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 109;
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
			this.state = 115;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 12, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 112;
					this.typeDef();
					}
					}
				}
				this.state = 117;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 12, this._ctx);
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
			this.state = 120;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 13, this._ctx) ) {
			case 1:
				{
				this.state = 118;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 119;
				this.multiLineComment();
				}
				break;
			}
			this.state = 122;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 125;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 123;
				this.match(OpenFGAParser.EXTEND);
				this.state = 124;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 127;
			this.match(OpenFGAParser.TYPE);
			this.state = 128;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 129;
			localctx._typeName = this.identifier();
			this.state = 137;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 16, this._ctx) ) {
			case 1:
				{
				this.state = 130;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 131;
				this.match(OpenFGAParser.RELATIONS);
				this.state = 133;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 132;
						this.relationDeclaration();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					this.state = 135;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 15, this._ctx);
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
			this.state = 141;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 17, this._ctx) ) {
			case 1:
				{
				this.state = 139;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 140;
				this.multiLineComment();
				}
				break;
			}
			this.state = 143;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 144;
			this.match(OpenFGAParser.DEFINE);
			this.state = 145;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 146;
			this.relationName();
			this.state = 148;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 147;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 150;
			this.match(OpenFGAParser.COLON);
			this.state = 152;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 151;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			{
			this.state = 154;
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
			this.state = 156;
			this.identifier();
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
			this.state = 161;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 5:
				{
				this.state = 158;
				this.relationDefDirectAssignment();
				}
				break;
			case 10:
			case 16:
			case 17:
			case 18:
			case 20:
			case 21:
			case 26:
				{
				this.state = 159;
				this.relationDefGrouping();
				}
				break;
			case 7:
				{
				this.state = 160;
				this.relationRecurse();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this.state = 164;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 21, this._ctx) ) {
			case 1:
				{
				this.state = 163;
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
			this.state = 168;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 10:
			case 16:
			case 17:
			case 18:
			case 20:
			case 21:
			case 26:
				{
				this.state = 166;
				this.relationDefGrouping();
				}
				break;
			case 7:
				{
				this.state = 167;
				this.relationRecurseNoDirect();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this.state = 171;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 23, this._ctx) ) {
			case 1:
				{
				this.state = 170;
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
			this.state = 202;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 29, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 180;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 173;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 174;
						this.match(OpenFGAParser.OR);
						this.state = 175;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 178;
						this._errHandler.sync(this);
						switch (this._input.LA(1)) {
						case 10:
						case 16:
						case 17:
						case 18:
						case 20:
						case 21:
						case 26:
							{
							this.state = 176;
							this.relationDefGrouping();
							}
							break;
						case 7:
							{
							this.state = 177;
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
					this.state = 182;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 25, this._ctx);
				} while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER);
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 191;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 184;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 185;
						this.match(OpenFGAParser.AND);
						this.state = 186;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 189;
						this._errHandler.sync(this);
						switch (this._input.LA(1)) {
						case 10:
						case 16:
						case 17:
						case 18:
						case 20:
						case 21:
						case 26:
							{
							this.state = 187;
							this.relationDefGrouping();
							}
							break;
						case 7:
							{
							this.state = 188;
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
					this.state = 193;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 27, this._ctx);
				} while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER);
				}
				break;
			case 3:
				this.enterOuterAlt(localctx, 3);
				{
				{
				this.state = 195;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 196;
				this.match(OpenFGAParser.BUT_NOT);
				this.state = 197;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 200;
				this._errHandler.sync(this);
				switch (this._input.LA(1)) {
				case 10:
				case 16:
				case 17:
				case 18:
				case 20:
				case 21:
				case 26:
					{
					this.state = 198;
					this.relationDefGrouping();
					}
					break;
				case 7:
					{
					this.state = 199;
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
			this.state = 204;
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
				this.relationDef();
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
	public relationRecurseNoDirect(): RelationRecurseNoDirectContext {
		let localctx: RelationRecurseNoDirectContext = new RelationRecurseNoDirectContext(this, this._ctx, this.state);
		this.enterRule(localctx, 24, OpenFGAParser.RULE_relationRecurseNoDirect);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 225;
			this.match(OpenFGAParser.LPAREN);
			this.state = 229;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 226;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 231;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 234;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 34, this._ctx) ) {
			case 1:
				{
				this.state = 232;
				this.relationDefNoDirect();
				}
				break;
			case 2:
				{
				this.state = 233;
				this.relationRecurseNoDirect();
				}
				break;
			}
			this.state = 239;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 236;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 241;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 242;
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
			this.state = 244;
			this.match(OpenFGAParser.LBRACKET);
			this.state = 246;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 245;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 248;
			this.relationDefTypeRestriction();
			this.state = 250;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 249;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 262;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===2) {
				{
				{
				this.state = 252;
				this.match(OpenFGAParser.COMMA);
				this.state = 254;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 253;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				this.state = 256;
				this.relationDefTypeRestriction();
				this.state = 258;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 257;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				}
				}
				this.state = 264;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 265;
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
			this.state = 267;
			localctx._rewriteComputedusersetName = this.identifier();
			this.state = 272;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 41, this._ctx) ) {
			case 1:
				{
				this.state = 268;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 269;
				this.match(OpenFGAParser.FROM);
				this.state = 270;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 271;
				localctx._rewriteTuplesetName = this.identifier();
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
			this.state = 275;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===56) {
				{
				this.state = 274;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 284;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 43, this._ctx) ) {
			case 1:
				{
				this.state = 277;
				this.relationDefTypeRestrictionBase();
				}
				break;
			case 2:
				{
				{
				this.state = 278;
				this.relationDefTypeRestrictionBase();
				this.state = 279;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 280;
				this.match(OpenFGAParser.KEYWORD_WITH);
				this.state = 281;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 282;
				this.conditionName();
				}
				}
				break;
			}
			this.state = 287;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===56) {
				{
				this.state = 286;
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
			this.state = 289;
			localctx._relationDefTypeRestrictionType = this.identifier();
			this.state = 294;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 1:
				{
				{
				this.state = 290;
				this.match(OpenFGAParser.COLON);
				this.state = 291;
				localctx._relationDefTypeRestrictionWildcard = this.match(OpenFGAParser.STAR);
				}
				}
				break;
			case 11:
				{
				{
				this.state = 292;
				this.match(OpenFGAParser.HASH);
				this.state = 293;
				localctx._relationDefTypeRestrictionRelation = this.identifier();
				}
				}
				break;
			case 2:
			case 9:
			case 36:
			case 56:
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
			this.state = 299;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 46, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 296;
					this.condition();
					}
					}
				}
				this.state = 301;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 46, this._ctx);
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
			this.state = 304;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 47, this._ctx) ) {
			case 1:
				{
				this.state = 302;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 303;
				this.multiLineComment();
				}
				break;
			}
			this.state = 306;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 307;
			this.match(OpenFGAParser.CONDITION);
			this.state = 308;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 309;
			this.conditionName();
			this.state = 311;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 310;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 313;
			this.match(OpenFGAParser.LPAREN);
			this.state = 315;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 314;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 317;
			this.conditionParameter();
			this.state = 319;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 318;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 331;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===2) {
				{
				{
				this.state = 321;
				this.match(OpenFGAParser.COMMA);
				this.state = 323;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 322;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				this.state = 325;
				this.conditionParameter();
				this.state = 327;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 326;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				}
				}
				this.state = 333;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 335;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===56) {
				{
				this.state = 334;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 337;
			this.match(OpenFGAParser.RPAREN);
			this.state = 339;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 338;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 341;
			this.match(OpenFGAParser.LBRACE);
			this.state = 343;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 56, this._ctx) ) {
			case 1:
				{
				this.state = 342;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 346;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 57, this._ctx) ) {
			case 1:
				{
				this.state = 345;
				this.match(OpenFGAParser.WHITESPACE);
				}
				break;
			}
			this.state = 348;
			this.conditionExpression();
			this.state = 350;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===56) {
				{
				this.state = 349;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 352;
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
	public conditionParameter(): ConditionParameterContext {
		let localctx: ConditionParameterContext = new ConditionParameterContext(this, this._ctx, this.state);
		this.enterRule(localctx, 40, OpenFGAParser.RULE_conditionParameter);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 357;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===56) {
				{
				this.state = 356;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 359;
			this.parameterName();
			this.state = 361;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 360;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 363;
			this.match(OpenFGAParser.COLON);
			this.state = 365;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 364;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 367;
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
			this.state = 369;
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
			this.state = 376;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 58:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 371;
				this.match(OpenFGAParser.CONDITION_PARAM_TYPE);
				}
				break;
			case 57:
				this.enterOuterAlt(localctx, 2);
				{
				{
				this.state = 372;
				this.match(OpenFGAParser.CONDITION_PARAM_CONTAINER);
				this.state = 373;
				this.match(OpenFGAParser.LESS);
				this.state = 374;
				this.match(OpenFGAParser.CONDITION_PARAM_TYPE);
				this.state = 375;
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
	public mixins(): MixinsContext {
		let localctx: MixinsContext = new MixinsContext(this, this._ctx, this.state);
		this.enterRule(localctx, 46, OpenFGAParser.RULE_mixins);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 381;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 63, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 378;
					this.mixin();
					}
					}
				}
				this.state = 383;
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
	// @RuleVersion(0)
	public mixin(): MixinContext {
		let localctx: MixinContext = new MixinContext(this, this._ctx, this.state);
		this.enterRule(localctx, 48, OpenFGAParser.RULE_mixin);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 386;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 64, this._ctx) ) {
			case 1:
				{
				this.state = 384;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 385;
				this.multiLineComment();
				}
				break;
			}
			this.state = 388;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 389;
			this.match(OpenFGAParser.MIXIN);
			this.state = 390;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 391;
			this.mixinName();
			this.state = 393;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 392;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 395;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 396;
			this.match(OpenFGAParser.RELATIONS);
			this.state = 398;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 397;
					this.relationDeclaration();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 400;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 66, this._ctx);
			} while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER);
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
	public mixinName(): MixinNameContext {
		let localctx: MixinNameContext = new MixinNameContext(this, this._ctx, this.state);
		this.enterRule(localctx, 50, OpenFGAParser.RULE_mixinName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 402;
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
	public multiLineComment(): MultiLineCommentContext {
		let localctx: MultiLineCommentContext = new MultiLineCommentContext(this, this._ctx, this.state);
		this.enterRule(localctx, 52, OpenFGAParser.RULE_multiLineComment);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 404;
			this.match(OpenFGAParser.HASH);
			this.state = 408;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while ((((_la) & ~0x1F) === 0 && ((1 << _la) & 4294967294) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 117440511) !== 0)) {
				{
				{
				this.state = 405;
				_la = this._input.LA(1);
				if(_la<=0 || _la===56) {
				this._errHandler.recoverInline(this);
				}
				else {
					this._errHandler.reportMatch(this);
				    this.consume();
				}
				}
				}
				this.state = 410;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 413;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 68, this._ctx) ) {
			case 1:
				{
				this.state = 411;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 412;
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
	public identifier(): IdentifierContext {
		let localctx: IdentifierContext = new IdentifierContext(this, this._ctx, this.state);
		this.enterRule(localctx, 54, OpenFGAParser.RULE_identifier);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 415;
			_la = this._input.LA(1);
			if(!((((_la) & ~0x1F) === 0 && ((1 << _la) & 70714368) !== 0))) {
			this._errHandler.recoverInline(this);
			}
			else {
				this._errHandler.reportMatch(this);
			    this.consume();
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
		this.enterRule(localctx, 56, OpenFGAParser.RULE_conditionExpression);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 421;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 70, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					this.state = 419;
					this._errHandler.sync(this);
					switch ( this._interp.adaptivePredict(this._input, 69, this._ctx) ) {
					case 1:
						{
						this.state = 417;
						_la = this._input.LA(1);
						if(!((((_la) & ~0x1F) === 0 && ((1 << _la) & 3758098360) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 33554367) !== 0))) {
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
						this.state = 418;
						_la = this._input.LA(1);
						if(_la<=0 || _la===38) {
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
				this.state = 423;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 70, this._ctx);
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

	public static readonly _serializedATN: number[] = [4,1,58,425,2,0,7,0,2,
	1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,
	10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,16,2,17,
	7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,2,22,7,22,2,23,7,23,2,24,7,
	24,2,25,7,25,2,26,7,26,2,27,7,27,2,28,7,28,1,0,3,0,60,8,0,1,0,3,0,63,8,
	0,1,0,1,0,3,0,67,8,0,1,0,3,0,70,8,0,1,0,1,0,3,0,74,8,0,1,0,3,0,77,8,0,1,
	0,1,0,3,0,81,8,0,1,0,1,0,3,0,85,8,0,1,0,1,0,1,1,1,1,1,1,3,1,92,8,1,1,1,
	1,1,1,1,1,1,1,1,1,1,3,1,100,8,1,1,2,1,2,1,2,3,2,105,8,2,1,2,1,2,1,2,1,2,
	3,2,111,8,2,1,3,5,3,114,8,3,10,3,12,3,117,9,3,1,4,1,4,3,4,121,8,4,1,4,1,
	4,1,4,3,4,126,8,4,1,4,1,4,1,4,1,4,1,4,1,4,4,4,134,8,4,11,4,12,4,135,3,4,
	138,8,4,1,5,1,5,3,5,142,8,5,1,5,1,5,1,5,1,5,1,5,3,5,149,8,5,1,5,1,5,3,5,
	153,8,5,1,5,1,5,1,6,1,6,1,7,1,7,1,7,3,7,162,8,7,1,7,3,7,165,8,7,1,8,1,8,
	3,8,169,8,8,1,8,3,8,172,8,8,1,9,1,9,1,9,1,9,1,9,3,9,179,8,9,4,9,181,8,9,
	11,9,12,9,182,1,9,1,9,1,9,1,9,1,9,3,9,190,8,9,4,9,192,8,9,11,9,12,9,193,
	1,9,1,9,1,9,1,9,1,9,3,9,201,8,9,3,9,203,8,9,1,10,1,10,1,11,1,11,5,11,209,
	8,11,10,11,12,11,212,9,11,1,11,1,11,3,11,216,8,11,1,11,5,11,219,8,11,10,
	11,12,11,222,9,11,1,11,1,11,1,12,1,12,5,12,228,8,12,10,12,12,12,231,9,12,
	1,12,1,12,3,12,235,8,12,1,12,5,12,238,8,12,10,12,12,12,241,9,12,1,12,1,
	12,1,13,1,13,3,13,247,8,13,1,13,1,13,3,13,251,8,13,1,13,1,13,3,13,255,8,
	13,1,13,1,13,3,13,259,8,13,5,13,261,8,13,10,13,12,13,264,9,13,1,13,1,13,
	1,14,1,14,1,14,1,14,1,14,3,14,273,8,14,1,15,3,15,276,8,15,1,15,1,15,1,15,
	1,15,1,15,1,15,1,15,3,15,285,8,15,1,15,3,15,288,8,15,1,16,1,16,1,16,1,16,
	1,16,3,16,295,8,16,1,17,5,17,298,8,17,10,17,12,17,301,9,17,1,18,1,18,3,
	18,305,8,18,1,18,1,18,1,18,1,18,1,18,3,18,312,8,18,1,18,1,18,3,18,316,8,
	18,1,18,1,18,3,18,320,8,18,1,18,1,18,3,18,324,8,18,1,18,1,18,3,18,328,8,
	18,5,18,330,8,18,10,18,12,18,333,9,18,1,18,3,18,336,8,18,1,18,1,18,3,18,
	340,8,18,1,18,1,18,3,18,344,8,18,1,18,3,18,347,8,18,1,18,1,18,3,18,351,
	8,18,1,18,1,18,1,19,1,19,1,20,3,20,358,8,20,1,20,1,20,3,20,362,8,20,1,20,
	1,20,3,20,366,8,20,1,20,1,20,1,21,1,21,1,22,1,22,1,22,1,22,1,22,3,22,377,
	8,22,1,23,5,23,380,8,23,10,23,12,23,383,9,23,1,24,1,24,3,24,387,8,24,1,
	24,1,24,1,24,1,24,1,24,3,24,394,8,24,1,24,1,24,1,24,4,24,399,8,24,11,24,
	12,24,400,1,25,1,25,1,26,1,26,5,26,407,8,26,10,26,12,26,410,9,26,1,26,1,
	26,3,26,414,8,26,1,27,1,27,1,28,1,28,5,28,420,8,28,10,28,12,28,423,9,28,
	1,28,0,0,29,0,2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,32,34,36,38,40,42,
	44,46,48,50,52,54,56,0,4,1,0,56,56,4,0,10,10,16,18,20,21,26,26,4,0,3,5,
	7,10,29,37,39,56,1,0,38,38,469,0,59,1,0,0,0,2,91,1,0,0,0,4,104,1,0,0,0,
	6,115,1,0,0,0,8,120,1,0,0,0,10,141,1,0,0,0,12,156,1,0,0,0,14,161,1,0,0,
	0,16,168,1,0,0,0,18,202,1,0,0,0,20,204,1,0,0,0,22,206,1,0,0,0,24,225,1,
	0,0,0,26,244,1,0,0,0,28,267,1,0,0,0,30,275,1,0,0,0,32,289,1,0,0,0,34,299,
	1,0,0,0,36,304,1,0,0,0,38,354,1,0,0,0,40,357,1,0,0,0,42,369,1,0,0,0,44,
	376,1,0,0,0,46,381,1,0,0,0,48,386,1,0,0,0,50,402,1,0,0,0,52,404,1,0,0,0,
	54,415,1,0,0,0,56,421,1,0,0,0,58,60,5,9,0,0,59,58,1,0,0,0,59,60,1,0,0,0,
	60,62,1,0,0,0,61,63,5,56,0,0,62,61,1,0,0,0,62,63,1,0,0,0,63,66,1,0,0,0,
	64,67,3,2,1,0,65,67,3,4,2,0,66,64,1,0,0,0,66,65,1,0,0,0,67,69,1,0,0,0,68,
	70,5,56,0,0,69,68,1,0,0,0,69,70,1,0,0,0,70,71,1,0,0,0,71,73,3,46,23,0,72,
	74,5,56,0,0,73,72,1,0,0,0,73,74,1,0,0,0,74,76,1,0,0,0,75,77,5,56,0,0,76,
	75,1,0,0,0,76,77,1,0,0,0,77,78,1,0,0,0,78,80,3,6,3,0,79,81,5,56,0,0,80,
	79,1,0,0,0,80,81,1,0,0,0,81,82,1,0,0,0,82,84,3,34,17,0,83,85,5,56,0,0,84,
	83,1,0,0,0,84,85,1,0,0,0,85,86,1,0,0,0,86,87,5,0,0,1,87,1,1,0,0,0,88,89,
	3,52,26,0,89,90,5,56,0,0,90,92,1,0,0,0,91,88,1,0,0,0,91,92,1,0,0,0,92,93,
	1,0,0,0,93,94,5,17,0,0,94,95,5,56,0,0,95,96,5,18,0,0,96,97,5,9,0,0,97,99,
	5,19,0,0,98,100,5,9,0,0,99,98,1,0,0,0,99,100,1,0,0,0,100,3,1,0,0,0,101,
	102,3,52,26,0,102,103,5,56,0,0,103,105,1,0,0,0,104,101,1,0,0,0,104,105,
	1,0,0,0,105,106,1,0,0,0,106,107,5,16,0,0,107,108,5,9,0,0,108,110,3,54,27,
	0,109,111,5,9,0,0,110,109,1,0,0,0,110,111,1,0,0,0,111,5,1,0,0,0,112,114,
	3,8,4,0,113,112,1,0,0,0,114,117,1,0,0,0,115,113,1,0,0,0,115,116,1,0,0,0,
	116,7,1,0,0,0,117,115,1,0,0,0,118,119,5,56,0,0,119,121,3,52,26,0,120,118,
	1,0,0,0,120,121,1,0,0,0,121,122,1,0,0,0,122,125,5,56,0,0,123,124,5,20,0,
	0,124,126,5,9,0,0,125,123,1,0,0,0,125,126,1,0,0,0,126,127,1,0,0,0,127,128,
	5,21,0,0,128,129,5,9,0,0,129,137,3,54,27,0,130,131,5,56,0,0,131,133,5,25,
	0,0,132,134,3,10,5,0,133,132,1,0,0,0,134,135,1,0,0,0,135,133,1,0,0,0,135,
	136,1,0,0,0,136,138,1,0,0,0,137,130,1,0,0,0,137,138,1,0,0,0,138,9,1,0,0,
	0,139,140,5,56,0,0,140,142,3,52,26,0,141,139,1,0,0,0,141,142,1,0,0,0,142,
	143,1,0,0,0,143,144,5,56,0,0,144,145,5,27,0,0,145,146,5,9,0,0,146,148,3,
	12,6,0,147,149,5,9,0,0,148,147,1,0,0,0,148,149,1,0,0,0,149,150,1,0,0,0,
	150,152,5,1,0,0,151,153,5,9,0,0,152,151,1,0,0,0,152,153,1,0,0,0,153,154,
	1,0,0,0,154,155,3,14,7,0,155,11,1,0,0,0,156,157,3,54,27,0,157,13,1,0,0,
	0,158,162,3,26,13,0,159,162,3,20,10,0,160,162,3,22,11,0,161,158,1,0,0,0,
	161,159,1,0,0,0,161,160,1,0,0,0,162,164,1,0,0,0,163,165,3,18,9,0,164,163,
	1,0,0,0,164,165,1,0,0,0,165,15,1,0,0,0,166,169,3,20,10,0,167,169,3,24,12,
	0,168,166,1,0,0,0,168,167,1,0,0,0,169,171,1,0,0,0,170,172,3,18,9,0,171,
	170,1,0,0,0,171,172,1,0,0,0,172,17,1,0,0,0,173,174,5,9,0,0,174,175,5,13,
	0,0,175,178,5,9,0,0,176,179,3,20,10,0,177,179,3,24,12,0,178,176,1,0,0,0,
	178,177,1,0,0,0,179,181,1,0,0,0,180,173,1,0,0,0,181,182,1,0,0,0,182,180,
	1,0,0,0,182,183,1,0,0,0,183,203,1,0,0,0,184,185,5,9,0,0,185,186,5,12,0,
	0,186,189,5,9,0,0,187,190,3,20,10,0,188,190,3,24,12,0,189,187,1,0,0,0,189,
	188,1,0,0,0,190,192,1,0,0,0,191,184,1,0,0,0,192,193,1,0,0,0,193,191,1,0,
	0,0,193,194,1,0,0,0,194,203,1,0,0,0,195,196,5,9,0,0,196,197,5,14,0,0,197,
	200,5,9,0,0,198,201,3,20,10,0,199,201,3,24,12,0,200,198,1,0,0,0,200,199,
	1,0,0,0,201,203,1,0,0,0,202,180,1,0,0,0,202,191,1,0,0,0,202,195,1,0,0,0,
	203,19,1,0,0,0,204,205,3,28,14,0,205,21,1,0,0,0,206,210,5,7,0,0,207,209,
	5,9,0,0,208,207,1,0,0,0,209,212,1,0,0,0,210,208,1,0,0,0,210,211,1,0,0,0,
	211,215,1,0,0,0,212,210,1,0,0,0,213,216,3,14,7,0,214,216,3,24,12,0,215,
	213,1,0,0,0,215,214,1,0,0,0,216,220,1,0,0,0,217,219,5,9,0,0,218,217,1,0,
	0,0,219,222,1,0,0,0,220,218,1,0,0,0,220,221,1,0,0,0,221,223,1,0,0,0,222,
	220,1,0,0,0,223,224,5,8,0,0,224,23,1,0,0,0,225,229,5,7,0,0,226,228,5,9,
	0,0,227,226,1,0,0,0,228,231,1,0,0,0,229,227,1,0,0,0,229,230,1,0,0,0,230,
	234,1,0,0,0,231,229,1,0,0,0,232,235,3,16,8,0,233,235,3,24,12,0,234,232,
	1,0,0,0,234,233,1,0,0,0,235,239,1,0,0,0,236,238,5,9,0,0,237,236,1,0,0,0,
	238,241,1,0,0,0,239,237,1,0,0,0,239,240,1,0,0,0,240,242,1,0,0,0,241,239,
	1,0,0,0,242,243,5,8,0,0,243,25,1,0,0,0,244,246,5,5,0,0,245,247,5,9,0,0,
	246,245,1,0,0,0,246,247,1,0,0,0,247,248,1,0,0,0,248,250,3,30,15,0,249,251,
	5,9,0,0,250,249,1,0,0,0,250,251,1,0,0,0,251,262,1,0,0,0,252,254,5,2,0,0,
	253,255,5,9,0,0,254,253,1,0,0,0,254,255,1,0,0,0,255,256,1,0,0,0,256,258,
	3,30,15,0,257,259,5,9,0,0,258,257,1,0,0,0,258,259,1,0,0,0,259,261,1,0,0,
	0,260,252,1,0,0,0,261,264,1,0,0,0,262,260,1,0,0,0,262,263,1,0,0,0,263,265,
	1,0,0,0,264,262,1,0,0,0,265,266,5,36,0,0,266,27,1,0,0,0,267,272,3,54,27,
	0,268,269,5,9,0,0,269,270,5,15,0,0,270,271,5,9,0,0,271,273,3,54,27,0,272,
	268,1,0,0,0,272,273,1,0,0,0,273,29,1,0,0,0,274,276,5,56,0,0,275,274,1,0,
	0,0,275,276,1,0,0,0,276,284,1,0,0,0,277,285,3,32,16,0,278,279,3,32,16,0,
	279,280,5,9,0,0,280,281,5,28,0,0,281,282,5,9,0,0,282,283,3,38,19,0,283,
	285,1,0,0,0,284,277,1,0,0,0,284,278,1,0,0,0,285,287,1,0,0,0,286,288,5,56,
	0,0,287,286,1,0,0,0,287,288,1,0,0,0,288,31,1,0,0,0,289,294,3,54,27,0,290,
	291,5,1,0,0,291,295,5,44,0,0,292,293,5,11,0,0,293,295,3,54,27,0,294,290,
	1,0,0,0,294,292,1,0,0,0,294,295,1,0,0,0,295,33,1,0,0,0,296,298,3,36,18,
	0,297,296,1,0,0,0,298,301,1,0,0,0,299,297,1,0,0,0,299,300,1,0,0,0,300,35,
	1,0,0,0,301,299,1,0,0,0,302,303,5,56,0,0,303,305,3,52,26,0,304,302,1,0,
	0,0,304,305,1,0,0,0,305,306,1,0,0,0,306,307,5,56,0,0,307,308,5,22,0,0,308,
	309,5,9,0,0,309,311,3,38,19,0,310,312,5,9,0,0,311,310,1,0,0,0,311,312,1,
	0,0,0,312,313,1,0,0,0,313,315,5,7,0,0,314,316,5,9,0,0,315,314,1,0,0,0,315,
	316,1,0,0,0,316,317,1,0,0,0,317,319,3,40,20,0,318,320,5,9,0,0,319,318,1,
	0,0,0,319,320,1,0,0,0,320,331,1,0,0,0,321,323,5,2,0,0,322,324,5,9,0,0,323,
	322,1,0,0,0,323,324,1,0,0,0,324,325,1,0,0,0,325,327,3,40,20,0,326,328,5,
	9,0,0,327,326,1,0,0,0,327,328,1,0,0,0,328,330,1,0,0,0,329,321,1,0,0,0,330,
	333,1,0,0,0,331,329,1,0,0,0,331,332,1,0,0,0,332,335,1,0,0,0,333,331,1,0,
	0,0,334,336,5,56,0,0,335,334,1,0,0,0,335,336,1,0,0,0,336,337,1,0,0,0,337,
	339,5,8,0,0,338,340,5,9,0,0,339,338,1,0,0,0,339,340,1,0,0,0,340,341,1,0,
	0,0,341,343,5,37,0,0,342,344,5,56,0,0,343,342,1,0,0,0,343,344,1,0,0,0,344,
	346,1,0,0,0,345,347,5,9,0,0,346,345,1,0,0,0,346,347,1,0,0,0,347,348,1,0,
	0,0,348,350,3,56,28,0,349,351,5,56,0,0,350,349,1,0,0,0,350,351,1,0,0,0,
	351,352,1,0,0,0,352,353,5,38,0,0,353,37,1,0,0,0,354,355,5,10,0,0,355,39,
	1,0,0,0,356,358,5,56,0,0,357,356,1,0,0,0,357,358,1,0,0,0,358,359,1,0,0,
	0,359,361,3,42,21,0,360,362,5,9,0,0,361,360,1,0,0,0,361,362,1,0,0,0,362,
	363,1,0,0,0,363,365,5,1,0,0,364,366,5,9,0,0,365,364,1,0,0,0,365,366,1,0,
	0,0,366,367,1,0,0,0,367,368,3,44,22,0,368,41,1,0,0,0,369,370,5,10,0,0,370,
	43,1,0,0,0,371,377,5,58,0,0,372,373,5,57,0,0,373,374,5,3,0,0,374,375,5,
	58,0,0,375,377,5,4,0,0,376,371,1,0,0,0,376,372,1,0,0,0,377,45,1,0,0,0,378,
	380,3,48,24,0,379,378,1,0,0,0,380,383,1,0,0,0,381,379,1,0,0,0,381,382,1,
	0,0,0,382,47,1,0,0,0,383,381,1,0,0,0,384,385,5,56,0,0,385,387,3,52,26,0,
	386,384,1,0,0,0,386,387,1,0,0,0,387,388,1,0,0,0,388,389,5,56,0,0,389,390,
	5,23,0,0,390,391,5,9,0,0,391,393,3,50,25,0,392,394,5,9,0,0,393,392,1,0,
	0,0,393,394,1,0,0,0,394,395,1,0,0,0,395,396,5,56,0,0,396,398,5,25,0,0,397,
	399,3,10,5,0,398,397,1,0,0,0,399,400,1,0,0,0,400,398,1,0,0,0,400,401,1,
	0,0,0,401,49,1,0,0,0,402,403,5,10,0,0,403,51,1,0,0,0,404,408,5,11,0,0,405,
	407,8,0,0,0,406,405,1,0,0,0,407,410,1,0,0,0,408,406,1,0,0,0,408,409,1,0,
	0,0,409,413,1,0,0,0,410,408,1,0,0,0,411,412,5,56,0,0,412,414,3,52,26,0,
	413,411,1,0,0,0,413,414,1,0,0,0,414,53,1,0,0,0,415,416,7,1,0,0,416,55,1,
	0,0,0,417,420,7,2,0,0,418,420,8,3,0,0,419,417,1,0,0,0,419,418,1,0,0,0,420,
	423,1,0,0,0,421,419,1,0,0,0,421,422,1,0,0,0,422,57,1,0,0,0,423,421,1,0,
	0,0,71,59,62,66,69,73,76,80,84,91,99,104,110,115,120,125,135,137,141,148,
	152,161,164,168,171,178,182,189,193,200,202,210,215,220,229,234,239,246,
	250,254,258,262,272,275,284,287,294,299,304,311,315,319,323,327,331,335,
	339,343,346,350,357,361,365,376,381,386,393,400,408,413,419,421];

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
	public mixins(): MixinsContext {
		return this.getTypedRuleContext(MixinsContext, 0) as MixinsContext;
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
	public _moduleName!: IdentifierContext;
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
	public identifier(): IdentifierContext {
		return this.getTypedRuleContext(IdentifierContext, 0) as IdentifierContext;
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
	public _typeName!: IdentifierContext;
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
	public identifier(): IdentifierContext {
		return this.getTypedRuleContext(IdentifierContext, 0) as IdentifierContext;
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
	public identifier(): IdentifierContext {
		return this.getTypedRuleContext(IdentifierContext, 0) as IdentifierContext;
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
	public _rewriteComputedusersetName!: IdentifierContext;
	public _rewriteTuplesetName!: IdentifierContext;
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public identifier_list(): IdentifierContext[] {
		return this.getTypedRuleContexts(IdentifierContext) as IdentifierContext[];
	}
	public identifier(i: number): IdentifierContext {
		return this.getTypedRuleContext(IdentifierContext, i) as IdentifierContext;
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
	public _relationDefTypeRestrictionType!: IdentifierContext;
	public _relationDefTypeRestrictionWildcard!: Token;
	public _relationDefTypeRestrictionRelation!: IdentifierContext;
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public identifier_list(): IdentifierContext[] {
		return this.getTypedRuleContexts(IdentifierContext) as IdentifierContext[];
	}
	public identifier(i: number): IdentifierContext {
		return this.getTypedRuleContext(IdentifierContext, i) as IdentifierContext;
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


export class MixinsContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public mixin_list(): MixinContext[] {
		return this.getTypedRuleContexts(MixinContext) as MixinContext[];
	}
	public mixin(i: number): MixinContext {
		return this.getTypedRuleContext(MixinContext, i) as MixinContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_mixins;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterMixins) {
	 		listener.enterMixins(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitMixins) {
	 		listener.exitMixins(this);
		}
	}
}


export class MixinContext extends ParserRuleContext {
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
	public MIXIN(): TerminalNode {
		return this.getToken(OpenFGAParser.MIXIN, 0);
	}
	public WHITESPACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WHITESPACE);
	}
	public WHITESPACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, i);
	}
	public mixinName(): MixinNameContext {
		return this.getTypedRuleContext(MixinNameContext, 0) as MixinNameContext;
	}
	public RELATIONS(): TerminalNode {
		return this.getToken(OpenFGAParser.RELATIONS, 0);
	}
	public multiLineComment(): MultiLineCommentContext {
		return this.getTypedRuleContext(MultiLineCommentContext, 0) as MultiLineCommentContext;
	}
	public relationDeclaration_list(): RelationDeclarationContext[] {
		return this.getTypedRuleContexts(RelationDeclarationContext) as RelationDeclarationContext[];
	}
	public relationDeclaration(i: number): RelationDeclarationContext {
		return this.getTypedRuleContext(RelationDeclarationContext, i) as RelationDeclarationContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_mixin;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterMixin) {
	 		listener.enterMixin(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitMixin) {
	 		listener.exitMixin(this);
		}
	}
}


export class MixinNameContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public IDENTIFIER(): TerminalNode {
		return this.getToken(OpenFGAParser.IDENTIFIER, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_mixinName;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterMixinName) {
	 		listener.enterMixinName(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitMixinName) {
	 		listener.exitMixinName(this);
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


export class IdentifierContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public MODEL(): TerminalNode {
		return this.getToken(OpenFGAParser.MODEL, 0);
	}
	public SCHEMA(): TerminalNode {
		return this.getToken(OpenFGAParser.SCHEMA, 0);
	}
	public TYPE(): TerminalNode {
		return this.getToken(OpenFGAParser.TYPE, 0);
	}
	public RELATION(): TerminalNode {
		return this.getToken(OpenFGAParser.RELATION, 0);
	}
	public IDENTIFIER(): TerminalNode {
		return this.getToken(OpenFGAParser.IDENTIFIER, 0);
	}
	public MODULE(): TerminalNode {
		return this.getToken(OpenFGAParser.MODULE, 0);
	}
	public EXTEND(): TerminalNode {
		return this.getToken(OpenFGAParser.EXTEND, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_identifier;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterIdentifier) {
	 		listener.enterIdentifier(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitIdentifier) {
	 		listener.exitIdentifier(this);
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
