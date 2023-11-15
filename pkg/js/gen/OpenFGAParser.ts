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
	public static readonly AND = 1;
	public static readonly OR = 2;
	public static readonly BUT_NOT = 3;
	public static readonly FROM = 4;
	public static readonly MODEL = 5;
	public static readonly SCHEMA = 6;
	public static readonly SCHEMA_VERSION = 7;
	public static readonly TYPE = 8;
	public static readonly CONDITION = 9;
	public static readonly CONDITION_PARAM_CONTAINER = 10;
	public static readonly CONDITION_PARAM_TYPE = 11;
	public static readonly RELATIONS = 12;
	public static readonly DEFINE = 13;
	public static readonly KEYWORD_WITH = 14;
	public static readonly IDENTIFIER = 15;
	public static readonly WHITESPACE = 16;
	public static readonly NEWLINE = 17;
	public static readonly DOT = 18;
	public static readonly STAR = 19;
	public static readonly HASH = 20;
	public static readonly COLON = 21;
	public static readonly COMMA = 22;
	public static readonly LPAREN = 23;
	public static readonly RPAREN = 24;
	public static readonly LESS = 25;
	public static readonly GREATER = 26;
	public static readonly LBRACKET = 27;
	public static readonly RPRACKET = 28;
	public static readonly OPEN_CEL = 29;
	public static readonly CEL_HASH = 30;
	public static readonly CEL_COLON = 31;
	public static readonly CEL_COMMA = 32;
	public static readonly EQUALS = 33;
	public static readonly NOT_EQUALS = 34;
	public static readonly IN = 35;
	public static readonly CEL_LESS = 36;
	public static readonly LESS_EQUALS = 37;
	public static readonly GREATER_EQUALS = 38;
	public static readonly CEL_GREATER = 39;
	public static readonly LOGICAL_AND = 40;
	public static readonly LOGICAL_OR = 41;
	public static readonly CEL_LBRACKET = 42;
	public static readonly CEL_RPRACKET = 43;
	public static readonly CEL_LPAREN = 44;
	public static readonly CEL_RPAREN = 45;
	public static readonly CEL_DOT = 46;
	public static readonly MINUS = 47;
	public static readonly EXCLAM = 48;
	public static readonly QUESTIONMARK = 49;
	public static readonly PLUS = 50;
	public static readonly CEL_STAR = 51;
	public static readonly SLASH = 52;
	public static readonly PERCENT = 53;
	public static readonly CEL_TRUE = 54;
	public static readonly CEL_FALSE = 55;
	public static readonly NUL = 56;
	public static readonly CEL_COMMENT = 57;
	public static readonly NUM_FLOAT = 58;
	public static readonly NUM_INT = 59;
	public static readonly NUM_UINT = 60;
	public static readonly STRING = 61;
	public static readonly BYTES = 62;
	public static readonly CEL_IDENTIFIER = 63;
	public static readonly CEL_WHITESPACE = 64;
	public static readonly CEL_NEWLINE = 65;
	public static readonly CLOSE_CEL = 66;
	public static readonly EOF = Token.EOF;
	public static readonly RULE_main = 0;
	public static readonly RULE_modelHeader = 1;
	public static readonly RULE_typeDefs = 2;
	public static readonly RULE_typeDef = 3;
	public static readonly RULE_relationDeclaration = 4;
	public static readonly RULE_relationName = 5;
	public static readonly RULE_relationDef = 6;
	public static readonly RULE_relationDefPartials = 7;
	public static readonly RULE_relationDefGrouping = 8;
	public static readonly RULE_relationDefDirectAssignment = 9;
	public static readonly RULE_relationDefRewrite = 10;
	public static readonly RULE_relationDefTypeRestriction = 11;
	public static readonly RULE_relationDefTypeRestrictionBase = 12;
	public static readonly RULE_conditions = 13;
	public static readonly RULE_condition = 14;
	public static readonly RULE_conditionName = 15;
	public static readonly RULE_conditionParameter = 16;
	public static readonly RULE_parameterName = 17;
	public static readonly RULE_parameterType = 18;
	public static readonly RULE_multiLineComment = 19;
	public static readonly RULE_conditionExpression = 20;
	public static readonly literalNames: (string | null)[] = [ null, "'and'", 
                                                            "'or'", "'but not'", 
                                                            "'from'", "'model'", 
                                                            "'schema'", 
                                                            "'1.1'", "'type'", 
                                                            "'condition'", 
                                                            null, null, 
                                                            "'relations'", 
                                                            "'define'", 
                                                            "'with'", null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, "'{'", 
                                                            null, null, 
                                                            null, "'=='", 
                                                            "'!='", "'in'", 
                                                            null, "'<='", 
                                                            "'>='", null, 
                                                            "'&&'", "'||'", 
                                                            null, null, 
                                                            null, null, 
                                                            null, "'-'", 
                                                            "'!'", "'?'", 
                                                            "'+'", null, 
                                                            "'/'", "'%'", 
                                                            "'true'", "'false'", 
                                                            "'null'", null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, null, 
                                                            "'}'" ];
	public static readonly symbolicNames: (string | null)[] = [ null, "AND", 
                                                             "OR", "BUT_NOT", 
                                                             "FROM", "MODEL", 
                                                             "SCHEMA", "SCHEMA_VERSION", 
                                                             "TYPE", "CONDITION", 
                                                             "CONDITION_PARAM_CONTAINER", 
                                                             "CONDITION_PARAM_TYPE", 
                                                             "RELATIONS", 
                                                             "DEFINE", "KEYWORD_WITH", 
                                                             "IDENTIFIER", 
                                                             "WHITESPACE", 
                                                             "NEWLINE", 
                                                             "DOT", "STAR", 
                                                             "HASH", "COLON", 
                                                             "COMMA", "LPAREN", 
                                                             "RPAREN", "LESS", 
                                                             "GREATER", 
                                                             "LBRACKET", 
                                                             "RPRACKET", 
                                                             "OPEN_CEL", 
                                                             "CEL_HASH", 
                                                             "CEL_COLON", 
                                                             "CEL_COMMA", 
                                                             "EQUALS", "NOT_EQUALS", 
                                                             "IN", "CEL_LESS", 
                                                             "LESS_EQUALS", 
                                                             "GREATER_EQUALS", 
                                                             "CEL_GREATER", 
                                                             "LOGICAL_AND", 
                                                             "LOGICAL_OR", 
                                                             "CEL_LBRACKET", 
                                                             "CEL_RPRACKET", 
                                                             "CEL_LPAREN", 
                                                             "CEL_RPAREN", 
                                                             "CEL_DOT", 
                                                             "MINUS", "EXCLAM", 
                                                             "QUESTIONMARK", 
                                                             "PLUS", "CEL_STAR", 
                                                             "SLASH", "PERCENT", 
                                                             "CEL_TRUE", 
                                                             "CEL_FALSE", 
                                                             "NUL", "CEL_COMMENT", 
                                                             "NUM_FLOAT", 
                                                             "NUM_INT", 
                                                             "NUM_UINT", 
                                                             "STRING", "BYTES", 
                                                             "CEL_IDENTIFIER", 
                                                             "CEL_WHITESPACE", 
                                                             "CEL_NEWLINE", 
                                                             "CLOSE_CEL" ];
	// tslint:disable:no-trailing-whitespace
	public static readonly ruleNames: string[] = [
		"main", "modelHeader", "typeDefs", "typeDef", "relationDeclaration", "relationName", 
		"relationDef", "relationDefPartials", "relationDefGrouping", "relationDefDirectAssignment", 
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
			this.state = 43;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===16) {
				{
				this.state = 42;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 46;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===17) {
				{
				this.state = 45;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 48;
			this.modelHeader();
			this.state = 50;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 2, this._ctx) ) {
			case 1:
				{
				this.state = 49;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 52;
			this.typeDefs();
			this.state = 54;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 3, this._ctx) ) {
			case 1:
				{
				this.state = 53;
				this.match(OpenFGAParser.NEWLINE);
				}
				break;
			}
			this.state = 56;
			this.conditions();
			this.state = 58;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===17) {
				{
				this.state = 57;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 60;
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
			this.state = 65;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 62;
				this.multiLineComment();
				this.state = 63;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 67;
			this.match(OpenFGAParser.MODEL);
			this.state = 68;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 69;
			this.match(OpenFGAParser.SCHEMA);
			this.state = 70;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 71;
			localctx._schemaVersion = this.match(OpenFGAParser.SCHEMA_VERSION);
			this.state = 73;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===16) {
				{
				this.state = 72;
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
			this.state = 78;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 7, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 75;
					this.typeDef();
					}
					}
				}
				this.state = 80;
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
			this.state = 83;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 8, this._ctx) ) {
			case 1:
				{
				this.state = 81;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 82;
				this.multiLineComment();
				}
				break;
			}
			this.state = 85;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 86;
			this.match(OpenFGAParser.TYPE);
			this.state = 87;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 88;
			localctx._typeName = this.match(OpenFGAParser.IDENTIFIER);
			this.state = 96;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 10, this._ctx) ) {
			case 1:
				{
				this.state = 89;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 90;
				this.match(OpenFGAParser.RELATIONS);
				this.state = 92;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 91;
						this.relationDeclaration();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					this.state = 94;
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
			this.state = 98;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 99;
			this.match(OpenFGAParser.DEFINE);
			this.state = 100;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 101;
			this.relationName();
			this.state = 103;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===16) {
				{
				this.state = 102;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 105;
			this.match(OpenFGAParser.COLON);
			this.state = 107;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===16) {
				{
				this.state = 106;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 109;
			this.relationDef();
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
			this.state = 111;
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
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 115;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 27:
				{
				this.state = 113;
				this.relationDefDirectAssignment();
				}
				break;
			case 15:
				{
				this.state = 114;
				this.relationDefGrouping();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this.state = 118;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===16) {
				{
				this.state = 117;
				this.relationDefPartials();
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
	public relationDefPartials(): RelationDefPartialsContext {
		let localctx: RelationDefPartialsContext = new RelationDefPartialsContext(this, this._ctx, this.state);
		this.enterRule(localctx, 14, OpenFGAParser.RULE_relationDefPartials);
		let _la: number;
		try {
			this.state = 144;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 18, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 124;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				do {
					{
					{
					this.state = 120;
					this.match(OpenFGAParser.WHITESPACE);
					this.state = 121;
					this.match(OpenFGAParser.OR);
					this.state = 122;
					this.match(OpenFGAParser.WHITESPACE);
					this.state = 123;
					this.relationDefGrouping();
					}
					}
					this.state = 126;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
				} while (_la===16);
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 132;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				do {
					{
					{
					this.state = 128;
					this.match(OpenFGAParser.WHITESPACE);
					this.state = 129;
					this.match(OpenFGAParser.AND);
					this.state = 130;
					this.match(OpenFGAParser.WHITESPACE);
					this.state = 131;
					this.relationDefGrouping();
					}
					}
					this.state = 134;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
				} while (_la===16);
				}
				break;
			case 3:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 140;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				do {
					{
					{
					this.state = 136;
					this.match(OpenFGAParser.WHITESPACE);
					this.state = 137;
					this.match(OpenFGAParser.BUT_NOT);
					this.state = 138;
					this.match(OpenFGAParser.WHITESPACE);
					this.state = 139;
					this.relationDefGrouping();
					}
					}
					this.state = 142;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
				} while (_la===16);
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
		this.enterRule(localctx, 16, OpenFGAParser.RULE_relationDefGrouping);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 146;
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
	public relationDefDirectAssignment(): RelationDefDirectAssignmentContext {
		let localctx: RelationDefDirectAssignmentContext = new RelationDefDirectAssignmentContext(this, this._ctx, this.state);
		this.enterRule(localctx, 18, OpenFGAParser.RULE_relationDefDirectAssignment);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 148;
			this.match(OpenFGAParser.LBRACKET);
			this.state = 150;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===16) {
				{
				this.state = 149;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 152;
			this.relationDefTypeRestriction();
			this.state = 154;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===16) {
				{
				this.state = 153;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 166;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===22) {
				{
				{
				this.state = 156;
				this.match(OpenFGAParser.COMMA);
				this.state = 158;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===16) {
					{
					this.state = 157;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				this.state = 160;
				this.relationDefTypeRestriction();
				this.state = 162;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===16) {
					{
					this.state = 161;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				}
				}
				this.state = 168;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 169;
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
		this.enterRule(localctx, 20, OpenFGAParser.RULE_relationDefRewrite);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 171;
			localctx._rewriteComputedusersetName = this.match(OpenFGAParser.IDENTIFIER);
			this.state = 176;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 24, this._ctx) ) {
			case 1:
				{
				this.state = 172;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 173;
				this.match(OpenFGAParser.FROM);
				this.state = 174;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 175;
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
		this.enterRule(localctx, 22, OpenFGAParser.RULE_relationDefTypeRestriction);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 179;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===17) {
				{
				this.state = 178;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 188;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 26, this._ctx) ) {
			case 1:
				{
				this.state = 181;
				this.relationDefTypeRestrictionBase();
				}
				break;
			case 2:
				{
				{
				this.state = 182;
				this.relationDefTypeRestrictionBase();
				this.state = 183;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 184;
				this.match(OpenFGAParser.KEYWORD_WITH);
				this.state = 185;
				this.match(OpenFGAParser.WHITESPACE);
				this.state = 186;
				this.conditionName();
				}
				}
				break;
			}
			this.state = 191;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===17) {
				{
				this.state = 190;
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
		this.enterRule(localctx, 24, OpenFGAParser.RULE_relationDefTypeRestrictionBase);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 193;
			localctx._relationDefTypeRestrictionType = this.match(OpenFGAParser.IDENTIFIER);
			this.state = 198;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 21:
				{
				{
				this.state = 194;
				this.match(OpenFGAParser.COLON);
				this.state = 195;
				localctx._relationDefTypeRestrictionWildcard = this.match(OpenFGAParser.STAR);
				}
				}
				break;
			case 20:
				{
				{
				this.state = 196;
				this.match(OpenFGAParser.HASH);
				this.state = 197;
				localctx._relationDefTypeRestrictionRelation = this.match(OpenFGAParser.IDENTIFIER);
				}
				}
				break;
			case 16:
			case 17:
			case 22:
			case 28:
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
		this.enterRule(localctx, 26, OpenFGAParser.RULE_conditions);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 203;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 29, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 200;
					this.condition();
					}
					}
				}
				this.state = 205;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 29, this._ctx);
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
		this.enterRule(localctx, 28, OpenFGAParser.RULE_condition);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 208;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 30, this._ctx) ) {
			case 1:
				{
				this.state = 206;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 207;
				this.multiLineComment();
				}
				break;
			}
			this.state = 210;
			this.match(OpenFGAParser.NEWLINE);
			this.state = 211;
			this.match(OpenFGAParser.CONDITION);
			this.state = 212;
			this.match(OpenFGAParser.WHITESPACE);
			this.state = 213;
			this.conditionName();
			this.state = 215;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===16) {
				{
				this.state = 214;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 217;
			this.match(OpenFGAParser.LPAREN);
			this.state = 219;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===16) {
				{
				this.state = 218;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 221;
			this.conditionParameter();
			this.state = 223;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===16) {
				{
				this.state = 222;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 235;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===22) {
				{
				{
				this.state = 225;
				this.match(OpenFGAParser.COMMA);
				this.state = 227;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===16) {
					{
					this.state = 226;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				this.state = 229;
				this.conditionParameter();
				this.state = 231;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===16) {
					{
					this.state = 230;
					this.match(OpenFGAParser.WHITESPACE);
					}
				}

				}
				}
				this.state = 237;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 239;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===17) {
				{
				this.state = 238;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 241;
			this.match(OpenFGAParser.RPAREN);
			this.state = 243;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===16) {
				{
				this.state = 242;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 245;
			this.match(OpenFGAParser.OPEN_CEL);
			this.state = 246;
			this.conditionExpression();
			this.state = 247;
			this.match(OpenFGAParser.CLOSE_CEL);
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
		this.enterRule(localctx, 30, OpenFGAParser.RULE_conditionName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 249;
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
		this.enterRule(localctx, 32, OpenFGAParser.RULE_conditionParameter);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 252;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===17) {
				{
				this.state = 251;
				this.match(OpenFGAParser.NEWLINE);
				}
			}

			this.state = 254;
			this.parameterName();
			this.state = 256;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===16) {
				{
				this.state = 255;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 258;
			this.match(OpenFGAParser.COLON);
			this.state = 260;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===16) {
				{
				this.state = 259;
				this.match(OpenFGAParser.WHITESPACE);
				}
			}

			this.state = 262;
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
		this.enterRule(localctx, 34, OpenFGAParser.RULE_parameterName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 264;
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
		this.enterRule(localctx, 36, OpenFGAParser.RULE_parameterType);
		try {
			this.state = 271;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 11:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 266;
				this.match(OpenFGAParser.CONDITION_PARAM_TYPE);
				}
				break;
			case 10:
				this.enterOuterAlt(localctx, 2);
				{
				{
				this.state = 267;
				this.match(OpenFGAParser.CONDITION_PARAM_CONTAINER);
				this.state = 268;
				this.match(OpenFGAParser.LESS);
				this.state = 269;
				this.match(OpenFGAParser.CONDITION_PARAM_TYPE);
				this.state = 270;
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
		this.enterRule(localctx, 38, OpenFGAParser.RULE_multiLineComment);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 273;
			this.match(OpenFGAParser.HASH);
			this.state = 277;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while ((((_la) & ~0x1F) === 0 && ((1 << _la) & 4294836222) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 4294967295) !== 0) || ((((_la - 64)) & ~0x1F) === 0 && ((1 << (_la - 64)) & 7) !== 0)) {
				{
				{
				this.state = 274;
				_la = this._input.LA(1);
				if(_la<=0 || _la===17) {
				this._errHandler.recoverInline(this);
				}
				else {
					this._errHandler.reportMatch(this);
				    this.consume();
				}
				}
				}
				this.state = 279;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 282;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 44, this._ctx) ) {
			case 1:
				{
				this.state = 280;
				this.match(OpenFGAParser.NEWLINE);
				this.state = 281;
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
		this.enterRule(localctx, 40, OpenFGAParser.RULE_conditionExpression);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 287;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while ((((_la) & ~0x1F) === 0 && ((1 << _la) & 4294967294) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 4294967295) !== 0) || _la===64 || _la===65) {
				{
				{
				this.state = 284;
				_la = this._input.LA(1);
				if(_la<=0 || _la===66) {
				this._errHandler.recoverInline(this);
				}
				else {
					this._errHandler.reportMatch(this);
				    this.consume();
				}
				}
				}
				this.state = 289;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
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

	public static readonly _serializedATN: number[] = [4,1,66,291,2,0,7,0,2,
	1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,
	10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,16,2,17,
	7,17,2,18,7,18,2,19,7,19,2,20,7,20,1,0,3,0,44,8,0,1,0,3,0,47,8,0,1,0,1,
	0,3,0,51,8,0,1,0,1,0,3,0,55,8,0,1,0,1,0,3,0,59,8,0,1,0,1,0,1,1,1,1,1,1,
	3,1,66,8,1,1,1,1,1,1,1,1,1,1,1,1,1,3,1,74,8,1,1,2,5,2,77,8,2,10,2,12,2,
	80,9,2,1,3,1,3,3,3,84,8,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,4,3,93,8,3,11,3,12,
	3,94,3,3,97,8,3,1,4,1,4,1,4,1,4,1,4,3,4,104,8,4,1,4,1,4,3,4,108,8,4,1,4,
	1,4,1,5,1,5,1,6,1,6,3,6,116,8,6,1,6,3,6,119,8,6,1,7,1,7,1,7,1,7,4,7,125,
	8,7,11,7,12,7,126,1,7,1,7,1,7,1,7,4,7,133,8,7,11,7,12,7,134,1,7,1,7,1,7,
	1,7,4,7,141,8,7,11,7,12,7,142,3,7,145,8,7,1,8,1,8,1,9,1,9,3,9,151,8,9,1,
	9,1,9,3,9,155,8,9,1,9,1,9,3,9,159,8,9,1,9,1,9,3,9,163,8,9,5,9,165,8,9,10,
	9,12,9,168,9,9,1,9,1,9,1,10,1,10,1,10,1,10,1,10,3,10,177,8,10,1,11,3,11,
	180,8,11,1,11,1,11,1,11,1,11,1,11,1,11,1,11,3,11,189,8,11,1,11,3,11,192,
	8,11,1,12,1,12,1,12,1,12,1,12,3,12,199,8,12,1,13,5,13,202,8,13,10,13,12,
	13,205,9,13,1,14,1,14,3,14,209,8,14,1,14,1,14,1,14,1,14,1,14,3,14,216,8,
	14,1,14,1,14,3,14,220,8,14,1,14,1,14,3,14,224,8,14,1,14,1,14,3,14,228,8,
	14,1,14,1,14,3,14,232,8,14,5,14,234,8,14,10,14,12,14,237,9,14,1,14,3,14,
	240,8,14,1,14,1,14,3,14,244,8,14,1,14,1,14,1,14,1,14,1,15,1,15,1,16,3,16,
	253,8,16,1,16,1,16,3,16,257,8,16,1,16,1,16,3,16,261,8,16,1,16,1,16,1,17,
	1,17,1,18,1,18,1,18,1,18,1,18,3,18,272,8,18,1,19,1,19,5,19,276,8,19,10,
	19,12,19,279,9,19,1,19,1,19,3,19,283,8,19,1,20,5,20,286,8,20,10,20,12,20,
	289,9,20,1,20,0,0,21,0,2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,32,34,36,
	38,40,0,2,1,0,17,17,1,0,66,66,317,0,43,1,0,0,0,2,65,1,0,0,0,4,78,1,0,0,
	0,6,83,1,0,0,0,8,98,1,0,0,0,10,111,1,0,0,0,12,115,1,0,0,0,14,144,1,0,0,
	0,16,146,1,0,0,0,18,148,1,0,0,0,20,171,1,0,0,0,22,179,1,0,0,0,24,193,1,
	0,0,0,26,203,1,0,0,0,28,208,1,0,0,0,30,249,1,0,0,0,32,252,1,0,0,0,34,264,
	1,0,0,0,36,271,1,0,0,0,38,273,1,0,0,0,40,287,1,0,0,0,42,44,5,16,0,0,43,
	42,1,0,0,0,43,44,1,0,0,0,44,46,1,0,0,0,45,47,5,17,0,0,46,45,1,0,0,0,46,
	47,1,0,0,0,47,48,1,0,0,0,48,50,3,2,1,0,49,51,5,17,0,0,50,49,1,0,0,0,50,
	51,1,0,0,0,51,52,1,0,0,0,52,54,3,4,2,0,53,55,5,17,0,0,54,53,1,0,0,0,54,
	55,1,0,0,0,55,56,1,0,0,0,56,58,3,26,13,0,57,59,5,17,0,0,58,57,1,0,0,0,58,
	59,1,0,0,0,59,60,1,0,0,0,60,61,5,0,0,1,61,1,1,0,0,0,62,63,3,38,19,0,63,
	64,5,17,0,0,64,66,1,0,0,0,65,62,1,0,0,0,65,66,1,0,0,0,66,67,1,0,0,0,67,
	68,5,5,0,0,68,69,5,17,0,0,69,70,5,6,0,0,70,71,5,16,0,0,71,73,5,7,0,0,72,
	74,5,16,0,0,73,72,1,0,0,0,73,74,1,0,0,0,74,3,1,0,0,0,75,77,3,6,3,0,76,75,
	1,0,0,0,77,80,1,0,0,0,78,76,1,0,0,0,78,79,1,0,0,0,79,5,1,0,0,0,80,78,1,
	0,0,0,81,82,5,17,0,0,82,84,3,38,19,0,83,81,1,0,0,0,83,84,1,0,0,0,84,85,
	1,0,0,0,85,86,5,17,0,0,86,87,5,8,0,0,87,88,5,16,0,0,88,96,5,15,0,0,89,90,
	5,17,0,0,90,92,5,12,0,0,91,93,3,8,4,0,92,91,1,0,0,0,93,94,1,0,0,0,94,92,
	1,0,0,0,94,95,1,0,0,0,95,97,1,0,0,0,96,89,1,0,0,0,96,97,1,0,0,0,97,7,1,
	0,0,0,98,99,5,17,0,0,99,100,5,13,0,0,100,101,5,16,0,0,101,103,3,10,5,0,
	102,104,5,16,0,0,103,102,1,0,0,0,103,104,1,0,0,0,104,105,1,0,0,0,105,107,
	5,21,0,0,106,108,5,16,0,0,107,106,1,0,0,0,107,108,1,0,0,0,108,109,1,0,0,
	0,109,110,3,12,6,0,110,9,1,0,0,0,111,112,5,15,0,0,112,11,1,0,0,0,113,116,
	3,18,9,0,114,116,3,16,8,0,115,113,1,0,0,0,115,114,1,0,0,0,116,118,1,0,0,
	0,117,119,3,14,7,0,118,117,1,0,0,0,118,119,1,0,0,0,119,13,1,0,0,0,120,121,
	5,16,0,0,121,122,5,2,0,0,122,123,5,16,0,0,123,125,3,16,8,0,124,120,1,0,
	0,0,125,126,1,0,0,0,126,124,1,0,0,0,126,127,1,0,0,0,127,145,1,0,0,0,128,
	129,5,16,0,0,129,130,5,1,0,0,130,131,5,16,0,0,131,133,3,16,8,0,132,128,
	1,0,0,0,133,134,1,0,0,0,134,132,1,0,0,0,134,135,1,0,0,0,135,145,1,0,0,0,
	136,137,5,16,0,0,137,138,5,3,0,0,138,139,5,16,0,0,139,141,3,16,8,0,140,
	136,1,0,0,0,141,142,1,0,0,0,142,140,1,0,0,0,142,143,1,0,0,0,143,145,1,0,
	0,0,144,124,1,0,0,0,144,132,1,0,0,0,144,140,1,0,0,0,145,15,1,0,0,0,146,
	147,3,20,10,0,147,17,1,0,0,0,148,150,5,27,0,0,149,151,5,16,0,0,150,149,
	1,0,0,0,150,151,1,0,0,0,151,152,1,0,0,0,152,154,3,22,11,0,153,155,5,16,
	0,0,154,153,1,0,0,0,154,155,1,0,0,0,155,166,1,0,0,0,156,158,5,22,0,0,157,
	159,5,16,0,0,158,157,1,0,0,0,158,159,1,0,0,0,159,160,1,0,0,0,160,162,3,
	22,11,0,161,163,5,16,0,0,162,161,1,0,0,0,162,163,1,0,0,0,163,165,1,0,0,
	0,164,156,1,0,0,0,165,168,1,0,0,0,166,164,1,0,0,0,166,167,1,0,0,0,167,169,
	1,0,0,0,168,166,1,0,0,0,169,170,5,28,0,0,170,19,1,0,0,0,171,176,5,15,0,
	0,172,173,5,16,0,0,173,174,5,4,0,0,174,175,5,16,0,0,175,177,5,15,0,0,176,
	172,1,0,0,0,176,177,1,0,0,0,177,21,1,0,0,0,178,180,5,17,0,0,179,178,1,0,
	0,0,179,180,1,0,0,0,180,188,1,0,0,0,181,189,3,24,12,0,182,183,3,24,12,0,
	183,184,5,16,0,0,184,185,5,14,0,0,185,186,5,16,0,0,186,187,3,30,15,0,187,
	189,1,0,0,0,188,181,1,0,0,0,188,182,1,0,0,0,189,191,1,0,0,0,190,192,5,17,
	0,0,191,190,1,0,0,0,191,192,1,0,0,0,192,23,1,0,0,0,193,198,5,15,0,0,194,
	195,5,21,0,0,195,199,5,19,0,0,196,197,5,20,0,0,197,199,5,15,0,0,198,194,
	1,0,0,0,198,196,1,0,0,0,198,199,1,0,0,0,199,25,1,0,0,0,200,202,3,28,14,
	0,201,200,1,0,0,0,202,205,1,0,0,0,203,201,1,0,0,0,203,204,1,0,0,0,204,27,
	1,0,0,0,205,203,1,0,0,0,206,207,5,17,0,0,207,209,3,38,19,0,208,206,1,0,
	0,0,208,209,1,0,0,0,209,210,1,0,0,0,210,211,5,17,0,0,211,212,5,9,0,0,212,
	213,5,16,0,0,213,215,3,30,15,0,214,216,5,16,0,0,215,214,1,0,0,0,215,216,
	1,0,0,0,216,217,1,0,0,0,217,219,5,23,0,0,218,220,5,16,0,0,219,218,1,0,0,
	0,219,220,1,0,0,0,220,221,1,0,0,0,221,223,3,32,16,0,222,224,5,16,0,0,223,
	222,1,0,0,0,223,224,1,0,0,0,224,235,1,0,0,0,225,227,5,22,0,0,226,228,5,
	16,0,0,227,226,1,0,0,0,227,228,1,0,0,0,228,229,1,0,0,0,229,231,3,32,16,
	0,230,232,5,16,0,0,231,230,1,0,0,0,231,232,1,0,0,0,232,234,1,0,0,0,233,
	225,1,0,0,0,234,237,1,0,0,0,235,233,1,0,0,0,235,236,1,0,0,0,236,239,1,0,
	0,0,237,235,1,0,0,0,238,240,5,17,0,0,239,238,1,0,0,0,239,240,1,0,0,0,240,
	241,1,0,0,0,241,243,5,24,0,0,242,244,5,16,0,0,243,242,1,0,0,0,243,244,1,
	0,0,0,244,245,1,0,0,0,245,246,5,29,0,0,246,247,3,40,20,0,247,248,5,66,0,
	0,248,29,1,0,0,0,249,250,5,15,0,0,250,31,1,0,0,0,251,253,5,17,0,0,252,251,
	1,0,0,0,252,253,1,0,0,0,253,254,1,0,0,0,254,256,3,34,17,0,255,257,5,16,
	0,0,256,255,1,0,0,0,256,257,1,0,0,0,257,258,1,0,0,0,258,260,5,21,0,0,259,
	261,5,16,0,0,260,259,1,0,0,0,260,261,1,0,0,0,261,262,1,0,0,0,262,263,3,
	36,18,0,263,33,1,0,0,0,264,265,5,15,0,0,265,35,1,0,0,0,266,272,5,11,0,0,
	267,268,5,10,0,0,268,269,5,25,0,0,269,270,5,11,0,0,270,272,5,26,0,0,271,
	266,1,0,0,0,271,267,1,0,0,0,272,37,1,0,0,0,273,277,5,20,0,0,274,276,8,0,
	0,0,275,274,1,0,0,0,276,279,1,0,0,0,277,275,1,0,0,0,277,278,1,0,0,0,278,
	282,1,0,0,0,279,277,1,0,0,0,280,281,5,17,0,0,281,283,3,38,19,0,282,280,
	1,0,0,0,282,283,1,0,0,0,283,39,1,0,0,0,284,286,8,1,0,0,285,284,1,0,0,0,
	286,289,1,0,0,0,287,285,1,0,0,0,287,288,1,0,0,0,288,41,1,0,0,0,289,287,
	1,0,0,0,46,43,46,50,54,58,65,73,78,83,94,96,103,107,115,118,126,134,142,
	144,150,154,158,162,166,176,179,188,191,198,203,208,215,219,223,227,231,
	235,239,243,252,256,260,271,277,282,287];

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
	public NEWLINE(): TerminalNode {
		return this.getToken(OpenFGAParser.NEWLINE, 0);
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
	public AND_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.AND);
	}
	public AND(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.AND, i);
	}
	public BUT_NOT_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.BUT_NOT);
	}
	public BUT_NOT(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.BUT_NOT, i);
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
	public OPEN_CEL(): TerminalNode {
		return this.getToken(OpenFGAParser.OPEN_CEL, 0);
	}
	public conditionExpression(): ConditionExpressionContext {
		return this.getTypedRuleContext(ConditionExpressionContext, 0) as ConditionExpressionContext;
	}
	public CLOSE_CEL(): TerminalNode {
		return this.getToken(OpenFGAParser.CLOSE_CEL, 0);
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
	public CLOSE_CEL_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.CLOSE_CEL);
	}
	public CLOSE_CEL(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.CLOSE_CEL, i);
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
