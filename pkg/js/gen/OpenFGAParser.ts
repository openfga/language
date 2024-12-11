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
	public static readonly RULE_mixinDeclaration = 5;
	public static readonly RULE_relationDeclaration = 6;
	public static readonly RULE_relationName = 7;
	public static readonly RULE_relationDef = 8;
	public static readonly RULE_relationDefNoDirect = 9;
	public static readonly RULE_relationDefPartials = 10;
	public static readonly RULE_relationDefGrouping = 11;
	public static readonly RULE_relationRecurse = 12;
	public static readonly RULE_relationRecurseNoDirect = 13;
	public static readonly RULE_relationDefDirectAssignment = 14;
	public static readonly RULE_relationDefRewrite = 15;
	public static readonly RULE_relationDefTypeRestriction = 16;
	public static readonly RULE_relationDefTypeRestrictionBase = 17;
	public static readonly RULE_conditions = 18;
	public static readonly RULE_condition = 19;
	public static readonly RULE_conditionName = 20;
	public static readonly RULE_conditionParameter = 21;
	public static readonly RULE_parameterName = 22;
	public static readonly RULE_parameterType = 23;
	public static readonly RULE_mixins = 24;
	public static readonly RULE_mixin = 25;
	public static readonly RULE_mixinName = 26;
	public static readonly RULE_multiLineComment = 27;
	public static readonly RULE_identifier = 28;
	public static readonly RULE_conditionExpression = 29;
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
		"main", "modelHeader", "moduleHeader", "typeDefs", "typeDef", "mixinDeclaration", 
		"relationDeclaration", "relationName", "relationDef", "relationDefNoDirect", 
		"relationDefPartials", "relationDefGrouping", "relationRecurse", "relationRecurseNoDirect", 
		"relationDefDirectAssignment", "relationDefRewrite", "relationDefTypeRestriction", 
		"relationDefTypeRestrictionBase", "conditions", "condition", "conditionName", 
		"conditionParameter", "parameterName", "parameterType", "mixins", "mixin", 
		"mixinName", "multiLineComment", "identifier", "conditionExpression",
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
			this.state = 61;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 60;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 64;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===56) {
				{
				this.state = 63;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 68;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 2, this._ctx) ) {
			case 1:
				{
				this.state = 66;
				this.modelHeader();
				}
				break;
			case 2:
				{
				this.state = 67;
				this.moduleHeader();
				}
				break;
			}
			this.state = 71;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 3, this._ctx) ) {
			case 1:
				{
				this.state = 70;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 73;
			this.mixins();
			this.state = 75;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 4, this._ctx) ) {
			case 1:
				{
				this.state = 74;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 78;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 5, this._ctx) ) {
			case 1:
				{
				this.state = 77;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 80;
			this.typeDefs();
			this.state = 82;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 6, this._ctx) ) {
			case 1:
				{
				this.state = 81;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 84;
			this.conditions();
			this.state = 86;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===56) {
				{
				this.state = 85;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 88;
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
			this.state = 93;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===11) {
				{
				this.state = 90;
				this.multiLineComment();
				this.state = 91;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 95;
			this.match(OpenFGAParser.MODEL);
			this.state = 96;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 97;
			this.match(OpenFGAParser.SCHEMA);
			this.state = 98;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 99;
			localctx._schemaVersion = this.match(OpenFGAParser.SCHEMA_VERSION);
			this.state = 101;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 100;
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
			this.state = 106;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===11) {
				{
				this.state = 103;
				this.multiLineComment();
				this.state = 104;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 108;
			this.match(OpenFGAParser.MODULE);
			this.state = 109;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 110;
			localctx._moduleName = this.identifier();
			this.state = 112;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 111;
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
			this.state = 117;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 12, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 114;
					this.typeDef();
					}
					}
				}
				this.state = 119;
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
			this.state = 122;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 13, this._ctx) ) {
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
			this.state = 127;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 125;
				this.match(OpenFGAParser.EXTEND);
				this.state = 126;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 129;
			this.match(OpenFGAParser.TYPE);
			this.state = 130;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 131;
			localctx._typeName = this.identifier();
			this.state = 140;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 17, this._ctx) ) {
			case 1:
				{
				this.state = 132;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 133;
				this.match(OpenFGAParser.RELATIONS);
				this.state = 136;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						this.state = 136;
						this._errHandler.sync(this);
						switch ( this._interp.adaptivePredict(this._input, 15, this._ctx) ) {
						case 1:
							{
							this.state = 134;
							this.mixinDeclaration();
							}
							break;
						case 2:
							{
							this.state = 135;
							this.relationDeclaration();
							}
							break;
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					this.state = 138;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 16, this._ctx);
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
	public mixinDeclaration(): MixinDeclarationContext {
		let localctx: MixinDeclarationContext = new MixinDeclarationContext(this, this._ctx, this.state);
		this.enterRule(localctx, 10, OpenFGAParser.RULE_mixinDeclaration);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			{
			this.state = 142;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 143;
			this.match(OpenFGAParser.INCLUDE);
			this.state = 144;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 145;
			this.mixinName();
			this.state = 156;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===2 || _la===9) {
				{
				{
				this.state = 147;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 146;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				this.state = 149;
				this.match(OpenFGAParser.COMMA);
				this.state = 151;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 150;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				this.state = 153;
				this.mixinName();
				}
				}
				this.state = 158;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
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
	public relationDeclaration(): RelationDeclarationContext {
		let localctx: RelationDeclarationContext = new RelationDeclarationContext(this, this._ctx, this.state);
		this.enterRule(localctx, 12, OpenFGAParser.RULE_relationDeclaration);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 161;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 21, this._ctx) ) {
			case 1:
				{
				this.state = 159;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 160;
				this.multiLineComment();
				}
				break;
			}
			this.state = 163;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 164;
			this.match(OpenFGAParser.DEFINE);
			this.state = 165;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 166;
			this.relationName();
			this.state = 168;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 167;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 170;
			this.match(OpenFGAParser.COLON);
			this.state = 172;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 171;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			{
			this.state = 174;
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
		this.enterRule(localctx, 14, OpenFGAParser.RULE_relationName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 176;
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
		this.enterRule(localctx, 16, OpenFGAParser.RULE_relationDef);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 181;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 5:
				{
				this.state = 178;
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
				this.state = 179;
				this.relationDefGrouping();
				}
				break;
			case 7:
				{
				this.state = 180;
				this.relationRecurse();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this.state = 184;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 25, this._ctx) ) {
			case 1:
				{
				this.state = 183;
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
		this.enterRule(localctx, 18, OpenFGAParser.RULE_relationDefNoDirect);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 188;
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
				this.state = 186;
				this.relationDefGrouping();
				}
				break;
			case 7:
				{
				this.state = 187;
				this.relationRecurseNoDirect();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this.state = 191;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 27, this._ctx) ) {
			case 1:
				{
				this.state = 190;
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
		this.enterRule(localctx, 20, OpenFGAParser.RULE_relationDefPartials);
		try {
			let _alt: number;
			this.state = 222;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 33, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 200;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 193;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 194;
						this.match(OpenFGAParser.OR);
						this.state = 195;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 198;
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
							this.state = 196;
							this.relationDefGrouping();
							}
							break;
						case 7:
							{
							this.state = 197;
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
					this.state = 202;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 29, this._ctx);
				} while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER);
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 211;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 204;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 205;
						this.match(OpenFGAParser.AND);
						this.state = 206;
						this.match(OpenFGAParser.WHITESPACE);
						this.state = 209;
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
							this.state = 207;
							this.relationDefGrouping();
							}
							break;
						case 7:
							{
							this.state = 208;
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
					this.state = 213;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 31, this._ctx);
				} while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER);
				}
				break;
			case 3:
				this.enterOuterAlt(localctx, 3);
				{
				{
				this.state = 215;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 216;
				this.match(OpenFGAParser.BUT_NOT);
				this.state = 217;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 220;
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
					this.state = 218;
					this.relationDefGrouping();
					}
					break;
				case 7:
					{
					this.state = 219;
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
		this.enterRule(localctx, 22, OpenFGAParser.RULE_relationDefGrouping);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 224;
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
		this.enterRule(localctx, 24, OpenFGAParser.RULE_relationRecurse);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 226;
			this.match(OpenFGAParser.LPAREN);
			this.state = 230;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 227;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 232;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 235;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 35, this._ctx) ) {
			case 1:
				{
				this.state = 233;
				this.relationDef();
				}
				break;
			case 2:
				{
				this.state = 234;
				this.relationRecurseNoDirect();
				}
				break;
			}
			this.state = 240;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 237;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 242;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 243;
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
		this.enterRule(localctx, 26, OpenFGAParser.RULE_relationRecurseNoDirect);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 245;
			this.match(OpenFGAParser.LPAREN);
			this.state = 249;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 246;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 251;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 254;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 38, this._ctx) ) {
			case 1:
				{
				this.state = 252;
				this.relationDefNoDirect();
				}
				break;
			case 2:
				{
				this.state = 253;
				this.relationRecurseNoDirect();
				}
				break;
			}
			this.state = 259;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===9) {
				{
				{
				this.state = 256;
				this.match(OpenFGAParser.WHITESPACE);
				}
				}
				this.state = 261;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 262;
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
		this.enterRule(localctx, 28, OpenFGAParser.RULE_relationDefDirectAssignment);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 264;
			this.match(OpenFGAParser.LBRACKET);
			this.state = 266;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 265;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 268;
			this.relationDefTypeRestriction();
			this.state = 270;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 269;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 282;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===2) {
				{
				{
				this.state = 272;
				this.match(OpenFGAParser.COMMA);
				this.state = 274;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 273;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				this.state = 276;
				this.relationDefTypeRestriction();
				this.state = 278;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 277;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				}
				}
				this.state = 284;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 285;
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
		this.enterRule(localctx, 30, OpenFGAParser.RULE_relationDefRewrite);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 287;
			localctx._rewriteComputedusersetName = this.identifier();
			this.state = 292;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 45, this._ctx) ) {
			case 1:
				{
				this.state = 288;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 289;
				this.match(OpenFGAParser.FROM);
				this.state = 290;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 291;
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
		this.enterRule(localctx, 32, OpenFGAParser.RULE_relationDefTypeRestriction);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 295;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===56) {
				{
				this.state = 294;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 304;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 47, this._ctx) ) {
			case 1:
				{
				this.state = 297;
				this.relationDefTypeRestrictionBase();
				}
				break;
			case 2:
				{
				{
				this.state = 298;
				this.relationDefTypeRestrictionBase();
				this.state = 299;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 300;
				this.match(OpenFGAParser.KEYWORD_WITH);
				this.state = 301;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 302;
				this.conditionName();
				}
				}
				break;
			}
			this.state = 307;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===56) {
				{
				this.state = 306;
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
		this.enterRule(localctx, 34, OpenFGAParser.RULE_relationDefTypeRestrictionBase);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 309;
			localctx._relationDefTypeRestrictionType = this.identifier();
			this.state = 314;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 1:
				{
				{
				this.state = 310;
				this.match(OpenFGAParser.COLON);
				this.state = 311;
				localctx._relationDefTypeRestrictionWildcard = this.match(OpenFGAParser.STAR);
				}
				}
				break;
			case 11:
				{
				{
				this.state = 312;
				this.match(OpenFGAParser.HASH);
				this.state = 313;
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
		this.enterRule(localctx, 36, OpenFGAParser.RULE_conditions);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 319;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 50, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 316;
					this.condition();
					}
					}
				}
				this.state = 321;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 50, this._ctx);
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
		this.enterRule(localctx, 38, OpenFGAParser.RULE_condition);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 324;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 51, this._ctx) ) {
			case 1:
				{
				this.state = 322;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 323;
				this.multiLineComment();
				}
				break;
			}
			this.state = 326;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 327;
			this.match(OpenFGAParser.CONDITION);
			this.state = 328;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 329;
			this.conditionName();
			this.state = 331;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 330;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 333;
			this.match(OpenFGAParser.LPAREN);
			this.state = 335;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 334;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 337;
			this.conditionParameter();
			this.state = 339;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 338;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 351;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===2) {
				{
				{
				this.state = 341;
				this.match(OpenFGAParser.COMMA);
				this.state = 343;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 342;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				this.state = 345;
				this.conditionParameter();
				this.state = 347;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===9) {
					{
					this.state = 346;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				}
				}
				this.state = 353;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 355;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===56) {
				{
				this.state = 354;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 357;
			this.match(OpenFGAParser.RPAREN);
			this.state = 359;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 358;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 361;
			this.match(OpenFGAParser.LBRACE);
			this.state = 363;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 60, this._ctx) ) {
			case 1:
				{
				this.state = 362;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 366;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 61, this._ctx) ) {
			case 1:
				{
				this.state = 365;
				this.match(OpenFGAParser.WHITESPACE);
				}
				break;
			}
			this.state = 368;
			this.conditionExpression();
			this.state = 370;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===56) {
				{
				this.state = 369;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 372;
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
		this.enterRule(localctx, 40, OpenFGAParser.RULE_conditionName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 374;
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
		this.enterRule(localctx, 42, OpenFGAParser.RULE_conditionParameter);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 377;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===56) {
				{
				this.state = 376;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 379;
			this.parameterName();
			this.state = 381;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 380;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 383;
			this.match(OpenFGAParser.COLON);
			this.state = 385;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 384;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 387;
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
		this.enterRule(localctx, 44, OpenFGAParser.RULE_parameterName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 389;
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
		this.enterRule(localctx, 46, OpenFGAParser.RULE_parameterType);
		try {
			this.state = 396;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 58:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 391;
				this.match(OpenFGAParser.CONDITION_PARAM_TYPE);
				}
				break;
			case 57:
				this.enterOuterAlt(localctx, 2);
				{
				{
				this.state = 392;
				this.match(OpenFGAParser.CONDITION_PARAM_CONTAINER);
				this.state = 393;
				this.match(OpenFGAParser.LESS);
				this.state = 394;
				this.match(OpenFGAParser.CONDITION_PARAM_TYPE);
				this.state = 395;
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
		this.enterRule(localctx, 48, OpenFGAParser.RULE_mixins);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 401;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 67, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 398;
					this.mixin();
					}
					}
				}
				this.state = 403;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 67, this._ctx);
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
		this.enterRule(localctx, 50, OpenFGAParser.RULE_mixin);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 406;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 68, this._ctx) ) {
			case 1:
				{
				this.state = 404;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 405;
				this.multiLineComment();
				}
				break;
			}
			this.state = 408;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 409;
			this.match(OpenFGAParser.MIXIN);
			this.state = 410;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 411;
			this.mixinName();
			this.state = 413;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===9) {
				{
				this.state = 412;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 415;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 416;
			this.match(OpenFGAParser.RELATIONS);
			this.state = 418;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 417;
					this.relationDeclaration();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 420;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 70, this._ctx);
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
		this.enterRule(localctx, 52, OpenFGAParser.RULE_mixinName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 422;
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
		this.enterRule(localctx, 54, OpenFGAParser.RULE_multiLineComment);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 424;
			this.match(OpenFGAParser.HASH);
			this.state = 428;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while ((((_la) & ~0x1F) === 0 && ((1 << _la) & 4294967294) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 117440511) !== 0)) {
				{
				{
				this.state = 425;
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
				this.state = 430;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 433;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 72, this._ctx) ) {
			case 1:
				{
				this.state = 431;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 432;
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
		this.enterRule(localctx, 56, OpenFGAParser.RULE_identifier);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 435;
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
		this.enterRule(localctx, 58, OpenFGAParser.RULE_conditionExpression);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 441;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 74, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					this.state = 439;
					this._errHandler.sync(this);
					switch ( this._interp.adaptivePredict(this._input, 73, this._ctx) ) {
					case 1:
						{
						this.state = 437;
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
						this.state = 438;
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
				this.state = 443;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 74, this._ctx);
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

	public static readonly _serializedATN: number[] = [4,1,58,445,2,0,7,0,2,
	1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,
	10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,16,2,17,
	7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,2,22,7,22,2,23,7,23,2,24,7,
	24,2,25,7,25,2,26,7,26,2,27,7,27,2,28,7,28,2,29,7,29,1,0,3,0,62,8,0,1,0,
	3,0,65,8,0,1,0,1,0,3,0,69,8,0,1,0,3,0,72,8,0,1,0,1,0,3,0,76,8,0,1,0,3,0,
	79,8,0,1,0,1,0,3,0,83,8,0,1,0,1,0,3,0,87,8,0,1,0,1,0,1,1,1,1,1,1,3,1,94,
	8,1,1,1,1,1,1,1,1,1,1,1,1,1,3,1,102,8,1,1,2,1,2,1,2,3,2,107,8,2,1,2,1,2,
	1,2,1,2,3,2,113,8,2,1,3,5,3,116,8,3,10,3,12,3,119,9,3,1,4,1,4,3,4,123,8,
	4,1,4,1,4,1,4,3,4,128,8,4,1,4,1,4,1,4,1,4,1,4,1,4,1,4,4,4,137,8,4,11,4,
	12,4,138,3,4,141,8,4,1,5,1,5,1,5,1,5,1,5,3,5,148,8,5,1,5,1,5,3,5,152,8,
	5,1,5,5,5,155,8,5,10,5,12,5,158,9,5,1,6,1,6,3,6,162,8,6,1,6,1,6,1,6,1,6,
	1,6,3,6,169,8,6,1,6,1,6,3,6,173,8,6,1,6,1,6,1,7,1,7,1,8,1,8,1,8,3,8,182,
	8,8,1,8,3,8,185,8,8,1,9,1,9,3,9,189,8,9,1,9,3,9,192,8,9,1,10,1,10,1,10,
	1,10,1,10,3,10,199,8,10,4,10,201,8,10,11,10,12,10,202,1,10,1,10,1,10,1,
	10,1,10,3,10,210,8,10,4,10,212,8,10,11,10,12,10,213,1,10,1,10,1,10,1,10,
	1,10,3,10,221,8,10,3,10,223,8,10,1,11,1,11,1,12,1,12,5,12,229,8,12,10,12,
	12,12,232,9,12,1,12,1,12,3,12,236,8,12,1,12,5,12,239,8,12,10,12,12,12,242,
	9,12,1,12,1,12,1,13,1,13,5,13,248,8,13,10,13,12,13,251,9,13,1,13,1,13,3,
	13,255,8,13,1,13,5,13,258,8,13,10,13,12,13,261,9,13,1,13,1,13,1,14,1,14,
	3,14,267,8,14,1,14,1,14,3,14,271,8,14,1,14,1,14,3,14,275,8,14,1,14,1,14,
	3,14,279,8,14,5,14,281,8,14,10,14,12,14,284,9,14,1,14,1,14,1,15,1,15,1,
	15,1,15,1,15,3,15,293,8,15,1,16,3,16,296,8,16,1,16,1,16,1,16,1,16,1,16,
	1,16,1,16,3,16,305,8,16,1,16,3,16,308,8,16,1,17,1,17,1,17,1,17,1,17,3,17,
	315,8,17,1,18,5,18,318,8,18,10,18,12,18,321,9,18,1,19,1,19,3,19,325,8,19,
	1,19,1,19,1,19,1,19,1,19,3,19,332,8,19,1,19,1,19,3,19,336,8,19,1,19,1,19,
	3,19,340,8,19,1,19,1,19,3,19,344,8,19,1,19,1,19,3,19,348,8,19,5,19,350,
	8,19,10,19,12,19,353,9,19,1,19,3,19,356,8,19,1,19,1,19,3,19,360,8,19,1,
	19,1,19,3,19,364,8,19,1,19,3,19,367,8,19,1,19,1,19,3,19,371,8,19,1,19,1,
	19,1,20,1,20,1,21,3,21,378,8,21,1,21,1,21,3,21,382,8,21,1,21,1,21,3,21,
	386,8,21,1,21,1,21,1,22,1,22,1,23,1,23,1,23,1,23,1,23,3,23,397,8,23,1,24,
	5,24,400,8,24,10,24,12,24,403,9,24,1,25,1,25,3,25,407,8,25,1,25,1,25,1,
	25,1,25,1,25,3,25,414,8,25,1,25,1,25,1,25,4,25,419,8,25,11,25,12,25,420,
	1,26,1,26,1,27,1,27,5,27,427,8,27,10,27,12,27,430,9,27,1,27,1,27,3,27,434,
	8,27,1,28,1,28,1,29,1,29,5,29,440,8,29,10,29,12,29,443,9,29,1,29,0,0,30,
	0,2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,32,34,36,38,40,42,44,46,48,50,
	52,54,56,58,0,4,1,0,56,56,4,0,10,10,16,18,20,21,26,26,4,0,3,5,7,10,29,37,
	39,56,1,0,38,38,492,0,61,1,0,0,0,2,93,1,0,0,0,4,106,1,0,0,0,6,117,1,0,0,
	0,8,122,1,0,0,0,10,142,1,0,0,0,12,161,1,0,0,0,14,176,1,0,0,0,16,181,1,0,
	0,0,18,188,1,0,0,0,20,222,1,0,0,0,22,224,1,0,0,0,24,226,1,0,0,0,26,245,
	1,0,0,0,28,264,1,0,0,0,30,287,1,0,0,0,32,295,1,0,0,0,34,309,1,0,0,0,36,
	319,1,0,0,0,38,324,1,0,0,0,40,374,1,0,0,0,42,377,1,0,0,0,44,389,1,0,0,0,
	46,396,1,0,0,0,48,401,1,0,0,0,50,406,1,0,0,0,52,422,1,0,0,0,54,424,1,0,
	0,0,56,435,1,0,0,0,58,441,1,0,0,0,60,62,5,9,0,0,61,60,1,0,0,0,61,62,1,0,
	0,0,62,64,1,0,0,0,63,65,5,56,0,0,64,63,1,0,0,0,64,65,1,0,0,0,65,68,1,0,
	0,0,66,69,3,2,1,0,67,69,3,4,2,0,68,66,1,0,0,0,68,67,1,0,0,0,69,71,1,0,0,
	0,70,72,5,56,0,0,71,70,1,0,0,0,71,72,1,0,0,0,72,73,1,0,0,0,73,75,3,48,24,
	0,74,76,5,56,0,0,75,74,1,0,0,0,75,76,1,0,0,0,76,78,1,0,0,0,77,79,5,56,0,
	0,78,77,1,0,0,0,78,79,1,0,0,0,79,80,1,0,0,0,80,82,3,6,3,0,81,83,5,56,0,
	0,82,81,1,0,0,0,82,83,1,0,0,0,83,84,1,0,0,0,84,86,3,36,18,0,85,87,5,56,
	0,0,86,85,1,0,0,0,86,87,1,0,0,0,87,88,1,0,0,0,88,89,5,0,0,1,89,1,1,0,0,
	0,90,91,3,54,27,0,91,92,5,56,0,0,92,94,1,0,0,0,93,90,1,0,0,0,93,94,1,0,
	0,0,94,95,1,0,0,0,95,96,5,17,0,0,96,97,5,56,0,0,97,98,5,18,0,0,98,99,5,
	9,0,0,99,101,5,19,0,0,100,102,5,9,0,0,101,100,1,0,0,0,101,102,1,0,0,0,102,
	3,1,0,0,0,103,104,3,54,27,0,104,105,5,56,0,0,105,107,1,0,0,0,106,103,1,
	0,0,0,106,107,1,0,0,0,107,108,1,0,0,0,108,109,5,16,0,0,109,110,5,9,0,0,
	110,112,3,56,28,0,111,113,5,9,0,0,112,111,1,0,0,0,112,113,1,0,0,0,113,5,
	1,0,0,0,114,116,3,8,4,0,115,114,1,0,0,0,116,119,1,0,0,0,117,115,1,0,0,0,
	117,118,1,0,0,0,118,7,1,0,0,0,119,117,1,0,0,0,120,121,5,56,0,0,121,123,
	3,54,27,0,122,120,1,0,0,0,122,123,1,0,0,0,123,124,1,0,0,0,124,127,5,56,
	0,0,125,126,5,20,0,0,126,128,5,9,0,0,127,125,1,0,0,0,127,128,1,0,0,0,128,
	129,1,0,0,0,129,130,5,21,0,0,130,131,5,9,0,0,131,140,3,56,28,0,132,133,
	5,56,0,0,133,136,5,25,0,0,134,137,3,10,5,0,135,137,3,12,6,0,136,134,1,0,
	0,0,136,135,1,0,0,0,137,138,1,0,0,0,138,136,1,0,0,0,138,139,1,0,0,0,139,
	141,1,0,0,0,140,132,1,0,0,0,140,141,1,0,0,0,141,9,1,0,0,0,142,143,5,56,
	0,0,143,144,5,24,0,0,144,145,5,9,0,0,145,156,3,52,26,0,146,148,5,9,0,0,
	147,146,1,0,0,0,147,148,1,0,0,0,148,149,1,0,0,0,149,151,5,2,0,0,150,152,
	5,9,0,0,151,150,1,0,0,0,151,152,1,0,0,0,152,153,1,0,0,0,153,155,3,52,26,
	0,154,147,1,0,0,0,155,158,1,0,0,0,156,154,1,0,0,0,156,157,1,0,0,0,157,11,
	1,0,0,0,158,156,1,0,0,0,159,160,5,56,0,0,160,162,3,54,27,0,161,159,1,0,
	0,0,161,162,1,0,0,0,162,163,1,0,0,0,163,164,5,56,0,0,164,165,5,27,0,0,165,
	166,5,9,0,0,166,168,3,14,7,0,167,169,5,9,0,0,168,167,1,0,0,0,168,169,1,
	0,0,0,169,170,1,0,0,0,170,172,5,1,0,0,171,173,5,9,0,0,172,171,1,0,0,0,172,
	173,1,0,0,0,173,174,1,0,0,0,174,175,3,16,8,0,175,13,1,0,0,0,176,177,3,56,
	28,0,177,15,1,0,0,0,178,182,3,28,14,0,179,182,3,22,11,0,180,182,3,24,12,
	0,181,178,1,0,0,0,181,179,1,0,0,0,181,180,1,0,0,0,182,184,1,0,0,0,183,185,
	3,20,10,0,184,183,1,0,0,0,184,185,1,0,0,0,185,17,1,0,0,0,186,189,3,22,11,
	0,187,189,3,26,13,0,188,186,1,0,0,0,188,187,1,0,0,0,189,191,1,0,0,0,190,
	192,3,20,10,0,191,190,1,0,0,0,191,192,1,0,0,0,192,19,1,0,0,0,193,194,5,
	9,0,0,194,195,5,13,0,0,195,198,5,9,0,0,196,199,3,22,11,0,197,199,3,26,13,
	0,198,196,1,0,0,0,198,197,1,0,0,0,199,201,1,0,0,0,200,193,1,0,0,0,201,202,
	1,0,0,0,202,200,1,0,0,0,202,203,1,0,0,0,203,223,1,0,0,0,204,205,5,9,0,0,
	205,206,5,12,0,0,206,209,5,9,0,0,207,210,3,22,11,0,208,210,3,26,13,0,209,
	207,1,0,0,0,209,208,1,0,0,0,210,212,1,0,0,0,211,204,1,0,0,0,212,213,1,0,
	0,0,213,211,1,0,0,0,213,214,1,0,0,0,214,223,1,0,0,0,215,216,5,9,0,0,216,
	217,5,14,0,0,217,220,5,9,0,0,218,221,3,22,11,0,219,221,3,26,13,0,220,218,
	1,0,0,0,220,219,1,0,0,0,221,223,1,0,0,0,222,200,1,0,0,0,222,211,1,0,0,0,
	222,215,1,0,0,0,223,21,1,0,0,0,224,225,3,30,15,0,225,23,1,0,0,0,226,230,
	5,7,0,0,227,229,5,9,0,0,228,227,1,0,0,0,229,232,1,0,0,0,230,228,1,0,0,0,
	230,231,1,0,0,0,231,235,1,0,0,0,232,230,1,0,0,0,233,236,3,16,8,0,234,236,
	3,26,13,0,235,233,1,0,0,0,235,234,1,0,0,0,236,240,1,0,0,0,237,239,5,9,0,
	0,238,237,1,0,0,0,239,242,1,0,0,0,240,238,1,0,0,0,240,241,1,0,0,0,241,243,
	1,0,0,0,242,240,1,0,0,0,243,244,5,8,0,0,244,25,1,0,0,0,245,249,5,7,0,0,
	246,248,5,9,0,0,247,246,1,0,0,0,248,251,1,0,0,0,249,247,1,0,0,0,249,250,
	1,0,0,0,250,254,1,0,0,0,251,249,1,0,0,0,252,255,3,18,9,0,253,255,3,26,13,
	0,254,252,1,0,0,0,254,253,1,0,0,0,255,259,1,0,0,0,256,258,5,9,0,0,257,256,
	1,0,0,0,258,261,1,0,0,0,259,257,1,0,0,0,259,260,1,0,0,0,260,262,1,0,0,0,
	261,259,1,0,0,0,262,263,5,8,0,0,263,27,1,0,0,0,264,266,5,5,0,0,265,267,
	5,9,0,0,266,265,1,0,0,0,266,267,1,0,0,0,267,268,1,0,0,0,268,270,3,32,16,
	0,269,271,5,9,0,0,270,269,1,0,0,0,270,271,1,0,0,0,271,282,1,0,0,0,272,274,
	5,2,0,0,273,275,5,9,0,0,274,273,1,0,0,0,274,275,1,0,0,0,275,276,1,0,0,0,
	276,278,3,32,16,0,277,279,5,9,0,0,278,277,1,0,0,0,278,279,1,0,0,0,279,281,
	1,0,0,0,280,272,1,0,0,0,281,284,1,0,0,0,282,280,1,0,0,0,282,283,1,0,0,0,
	283,285,1,0,0,0,284,282,1,0,0,0,285,286,5,36,0,0,286,29,1,0,0,0,287,292,
	3,56,28,0,288,289,5,9,0,0,289,290,5,15,0,0,290,291,5,9,0,0,291,293,3,56,
	28,0,292,288,1,0,0,0,292,293,1,0,0,0,293,31,1,0,0,0,294,296,5,56,0,0,295,
	294,1,0,0,0,295,296,1,0,0,0,296,304,1,0,0,0,297,305,3,34,17,0,298,299,3,
	34,17,0,299,300,5,9,0,0,300,301,5,28,0,0,301,302,5,9,0,0,302,303,3,40,20,
	0,303,305,1,0,0,0,304,297,1,0,0,0,304,298,1,0,0,0,305,307,1,0,0,0,306,308,
	5,56,0,0,307,306,1,0,0,0,307,308,1,0,0,0,308,33,1,0,0,0,309,314,3,56,28,
	0,310,311,5,1,0,0,311,315,5,44,0,0,312,313,5,11,0,0,313,315,3,56,28,0,314,
	310,1,0,0,0,314,312,1,0,0,0,314,315,1,0,0,0,315,35,1,0,0,0,316,318,3,38,
	19,0,317,316,1,0,0,0,318,321,1,0,0,0,319,317,1,0,0,0,319,320,1,0,0,0,320,
	37,1,0,0,0,321,319,1,0,0,0,322,323,5,56,0,0,323,325,3,54,27,0,324,322,1,
	0,0,0,324,325,1,0,0,0,325,326,1,0,0,0,326,327,5,56,0,0,327,328,5,22,0,0,
	328,329,5,9,0,0,329,331,3,40,20,0,330,332,5,9,0,0,331,330,1,0,0,0,331,332,
	1,0,0,0,332,333,1,0,0,0,333,335,5,7,0,0,334,336,5,9,0,0,335,334,1,0,0,0,
	335,336,1,0,0,0,336,337,1,0,0,0,337,339,3,42,21,0,338,340,5,9,0,0,339,338,
	1,0,0,0,339,340,1,0,0,0,340,351,1,0,0,0,341,343,5,2,0,0,342,344,5,9,0,0,
	343,342,1,0,0,0,343,344,1,0,0,0,344,345,1,0,0,0,345,347,3,42,21,0,346,348,
	5,9,0,0,347,346,1,0,0,0,347,348,1,0,0,0,348,350,1,0,0,0,349,341,1,0,0,0,
	350,353,1,0,0,0,351,349,1,0,0,0,351,352,1,0,0,0,352,355,1,0,0,0,353,351,
	1,0,0,0,354,356,5,56,0,0,355,354,1,0,0,0,355,356,1,0,0,0,356,357,1,0,0,
	0,357,359,5,8,0,0,358,360,5,9,0,0,359,358,1,0,0,0,359,360,1,0,0,0,360,361,
	1,0,0,0,361,363,5,37,0,0,362,364,5,56,0,0,363,362,1,0,0,0,363,364,1,0,0,
	0,364,366,1,0,0,0,365,367,5,9,0,0,366,365,1,0,0,0,366,367,1,0,0,0,367,368,
	1,0,0,0,368,370,3,58,29,0,369,371,5,56,0,0,370,369,1,0,0,0,370,371,1,0,
	0,0,371,372,1,0,0,0,372,373,5,38,0,0,373,39,1,0,0,0,374,375,5,10,0,0,375,
	41,1,0,0,0,376,378,5,56,0,0,377,376,1,0,0,0,377,378,1,0,0,0,378,379,1,0,
	0,0,379,381,3,44,22,0,380,382,5,9,0,0,381,380,1,0,0,0,381,382,1,0,0,0,382,
	383,1,0,0,0,383,385,5,1,0,0,384,386,5,9,0,0,385,384,1,0,0,0,385,386,1,0,
	0,0,386,387,1,0,0,0,387,388,3,46,23,0,388,43,1,0,0,0,389,390,5,10,0,0,390,
	45,1,0,0,0,391,397,5,58,0,0,392,393,5,57,0,0,393,394,5,3,0,0,394,395,5,
	58,0,0,395,397,5,4,0,0,396,391,1,0,0,0,396,392,1,0,0,0,397,47,1,0,0,0,398,
	400,3,50,25,0,399,398,1,0,0,0,400,403,1,0,0,0,401,399,1,0,0,0,401,402,1,
	0,0,0,402,49,1,0,0,0,403,401,1,0,0,0,404,405,5,56,0,0,405,407,3,54,27,0,
	406,404,1,0,0,0,406,407,1,0,0,0,407,408,1,0,0,0,408,409,5,56,0,0,409,410,
	5,23,0,0,410,411,5,9,0,0,411,413,3,52,26,0,412,414,5,9,0,0,413,412,1,0,
	0,0,413,414,1,0,0,0,414,415,1,0,0,0,415,416,5,56,0,0,416,418,5,25,0,0,417,
	419,3,12,6,0,418,417,1,0,0,0,419,420,1,0,0,0,420,418,1,0,0,0,420,421,1,
	0,0,0,421,51,1,0,0,0,422,423,5,10,0,0,423,53,1,0,0,0,424,428,5,11,0,0,425,
	427,8,0,0,0,426,425,1,0,0,0,427,430,1,0,0,0,428,426,1,0,0,0,428,429,1,0,
	0,0,429,433,1,0,0,0,430,428,1,0,0,0,431,432,5,56,0,0,432,434,3,54,27,0,
	433,431,1,0,0,0,433,434,1,0,0,0,434,55,1,0,0,0,435,436,7,1,0,0,436,57,1,
	0,0,0,437,440,7,2,0,0,438,440,8,3,0,0,439,437,1,0,0,0,439,438,1,0,0,0,440,
	443,1,0,0,0,441,439,1,0,0,0,441,442,1,0,0,0,442,59,1,0,0,0,443,441,1,0,
	0,0,75,61,64,68,71,75,78,82,86,93,101,106,112,117,122,127,136,138,140,147,
	151,156,161,168,172,181,184,188,191,198,202,209,213,220,222,230,235,240,
	249,254,259,266,270,274,278,282,292,295,304,307,314,319,324,331,335,339,
	343,347,351,355,359,363,366,370,377,381,385,396,401,406,413,420,428,433,
	439,441];

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
	public mixinDeclaration_list(): MixinDeclarationContext[] {
		return this.getTypedRuleContexts(MixinDeclarationContext) as MixinDeclarationContext[];
	}
	public mixinDeclaration(i: number): MixinDeclarationContext {
		return this.getTypedRuleContext(MixinDeclarationContext, i) as MixinDeclarationContext;
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


export class MixinDeclarationContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public NEWLINE(): TerminalNode {
		return this.getToken(OpenFGAParser.NEWLINE, 0);
	}
	public INCLUDE(): TerminalNode {
		return this.getToken(OpenFGAParser.INCLUDE, 0);
	}
	public WHITESPACE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WHITESPACE);
	}
	public WHITESPACE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WHITESPACE, i);
	}
	public mixinName_list(): MixinNameContext[] {
		return this.getTypedRuleContexts(MixinNameContext) as MixinNameContext[];
	}
	public mixinName(i: number): MixinNameContext {
		return this.getTypedRuleContext(MixinNameContext, i) as MixinNameContext;
	}
	public COMMA_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.COMMA);
	}
	public COMMA(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.COMMA, i);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_mixinDeclaration;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterMixinDeclaration) {
	 		listener.enterMixinDeclaration(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitMixinDeclaration) {
	 		listener.exitMixinDeclaration(this);
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
