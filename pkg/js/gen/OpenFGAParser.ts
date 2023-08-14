// Generated from /app/OpenFGAParser.g4 by ANTLR 4.13.0
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
	public static readonly RULE_main = 0;
	public static readonly RULE_indentation = 1;
	public static readonly RULE_modelHeader = 2;
	public static readonly RULE_typeDefs = 3;
	public static readonly RULE_typeDef = 4;
	public static readonly RULE_relationDeclaration = 5;
	public static readonly RULE_relationDef = 6;
	public static readonly RULE_relationDefPartials = 7;
	public static readonly RULE_relationDefPartialAllOr = 8;
	public static readonly RULE_relationDefPartialAllAnd = 9;
	public static readonly RULE_relationDefPartialAllButNot = 10;
	public static readonly RULE_relationDefDirectAssignment = 11;
	public static readonly RULE_relationDefRewrite = 12;
	public static readonly RULE_relationDefRelationOnSameObject = 13;
	public static readonly RULE_relationDefRelationOnRelatedObject = 14;
	public static readonly RULE_relationDefOperator = 15;
	public static readonly RULE_relationDefOperatorAnd = 16;
	public static readonly RULE_relationDefOperatorOr = 17;
	public static readonly RULE_relationDefOperatorButNot = 18;
	public static readonly RULE_relationDefKeywordFrom = 19;
	public static readonly RULE_relationDefTypeRestriction = 20;
	public static readonly RULE_relationDefTypeRestrictionType = 21;
	public static readonly RULE_relationDefTypeRestrictionRelation = 22;
	public static readonly RULE_relationDefTypeRestrictionWildcard = 23;
	public static readonly RULE_relationDefTypeRestrictionUserset = 24;
	public static readonly RULE_relationDefGrouping = 25;
	public static readonly RULE_rewriteComputedusersetName = 26;
	public static readonly RULE_rewriteTuplesetComputedusersetName = 27;
	public static readonly RULE_rewriteTuplesetName = 28;
	public static readonly RULE_relationName = 29;
	public static readonly RULE_typeName = 30;
	public static readonly RULE_comment = 31;
	public static readonly RULE_multiLineComment = 32;
	public static readonly RULE_spacing = 33;
	public static readonly RULE_newline = 34;
	public static readonly RULE_schemaVersion = 35;
	public static readonly RULE_name = 36;
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
	// tslint:disable:no-trailing-whitespace
	public static readonly ruleNames: string[] = [
		"main", "indentation", "modelHeader", "typeDefs", "typeDef", "relationDeclaration", 
		"relationDef", "relationDefPartials", "relationDefPartialAllOr", "relationDefPartialAllAnd", 
		"relationDefPartialAllButNot", "relationDefDirectAssignment", "relationDefRewrite", 
		"relationDefRelationOnSameObject", "relationDefRelationOnRelatedObject", 
		"relationDefOperator", "relationDefOperatorAnd", "relationDefOperatorOr", 
		"relationDefOperatorButNot", "relationDefKeywordFrom", "relationDefTypeRestriction", 
		"relationDefTypeRestrictionType", "relationDefTypeRestrictionRelation", 
		"relationDefTypeRestrictionWildcard", "relationDefTypeRestrictionUserset", 
		"relationDefGrouping", "rewriteComputedusersetName", "rewriteTuplesetComputedusersetName", 
		"rewriteTuplesetName", "relationName", "typeName", "comment", "multiLineComment", 
		"spacing", "newline", "schemaVersion", "name",
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
			this.state = 74;
			this.modelHeader();
			this.state = 75;
			this.typeDefs();
			this.state = 77;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===19) {
				{
				this.state = 76;
				this.newline();
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
	public indentation(): IndentationContext {
		let localctx: IndentationContext = new IndentationContext(this, this._ctx, this.state);
		this.enterRule(localctx, 2, OpenFGAParser.RULE_indentation);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 79;
			this.match(OpenFGAParser.INDENT);
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
		this.enterRule(localctx, 4, OpenFGAParser.RULE_modelHeader);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 84;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===8 || _la===20) {
				{
				this.state = 81;
				this.multiLineComment();
				this.state = 82;
				this.newline();
				}
			}

			this.state = 86;
			this.match(OpenFGAParser.MODEL);
			this.state = 88;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 87;
				this.spacing();
				}
			}

			this.state = 93;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===19) {
				{
				this.state = 90;
				this.newline();
				this.state = 91;
				this.multiLineComment();
				}
			}

			this.state = 95;
			this.indentation();
			this.state = 96;
			this.match(OpenFGAParser.SCHEMA);
			this.state = 97;
			this.spacing();
			this.state = 98;
			this.schemaVersion();
			this.state = 100;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 99;
				this.spacing();
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
			this.state = 105;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 5, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 102;
					this.typeDef();
					}
					}
				}
				this.state = 107;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 5, this._ctx);
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
			this.state = 111;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 6, this._ctx) ) {
			case 1:
				{
				this.state = 108;
				this.newline();
				this.state = 109;
				this.multiLineComment();
				}
				break;
			}
			this.state = 113;
			this.newline();
			this.state = 114;
			this.match(OpenFGAParser.TYPE);
			this.state = 115;
			this.spacing();
			this.state = 116;
			this.typeName();
			this.state = 118;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 117;
				this.spacing();
				}
			}

			this.state = 130;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===1) {
				{
				this.state = 120;
				this.indentation();
				this.state = 121;
				this.match(OpenFGAParser.RELATIONS);
				this.state = 123;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===20) {
					{
					this.state = 122;
					this.spacing();
					}
				}

				this.state = 126;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 125;
						this.relationDeclaration();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					this.state = 128;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 9, this._ctx);
				} while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER);
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
		this.enterRule(localctx, 10, OpenFGAParser.RULE_relationDeclaration);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 135;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===19) {
				{
				this.state = 132;
				this.newline();
				this.state = 133;
				this.multiLineComment();
				}
			}

			this.state = 137;
			this.indentation();
			this.state = 138;
			this.match(OpenFGAParser.DEFINE);
			this.state = 139;
			this.spacing();
			this.state = 140;
			this.relationName();
			this.state = 142;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 141;
				this.spacing();
				}
			}

			this.state = 144;
			this.match(OpenFGAParser.COLON);
			this.state = 146;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 145;
				this.spacing();
				}
			}

			this.state = 148;
			this.relationDef();
			this.state = 150;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 149;
				this.spacing();
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
	public relationDef(): RelationDefContext {
		let localctx: RelationDefContext = new RelationDefContext(this, this._ctx, this.state);
		this.enterRule(localctx, 12, OpenFGAParser.RULE_relationDef);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 154;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 11:
				{
				this.state = 152;
				this.relationDefDirectAssignment();
				}
				break;
			case 18:
				{
				this.state = 153;
				this.relationDefGrouping();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this.state = 157;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 16, this._ctx) ) {
			case 1:
				{
				this.state = 156;
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
		this.enterRule(localctx, 14, OpenFGAParser.RULE_relationDefPartials);
		try {
			this.state = 162;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 17, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 159;
				this.relationDefPartialAllOr();
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 160;
				this.relationDefPartialAllAnd();
				}
				break;
			case 3:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 161;
				this.relationDefPartialAllButNot();
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
	public relationDefPartialAllOr(): RelationDefPartialAllOrContext {
		let localctx: RelationDefPartialAllOrContext = new RelationDefPartialAllOrContext(this, this._ctx, this.state);
		this.enterRule(localctx, 16, OpenFGAParser.RULE_relationDefPartialAllOr);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 169;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 164;
					this.spacing();
					this.state = 165;
					this.relationDefOperatorOr();
					this.state = 166;
					this.spacing();
					this.state = 167;
					this.relationDefGrouping();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 171;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 18, this._ctx);
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
	public relationDefPartialAllAnd(): RelationDefPartialAllAndContext {
		let localctx: RelationDefPartialAllAndContext = new RelationDefPartialAllAndContext(this, this._ctx, this.state);
		this.enterRule(localctx, 18, OpenFGAParser.RULE_relationDefPartialAllAnd);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 178;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 173;
					this.spacing();
					this.state = 174;
					this.relationDefOperatorAnd();
					this.state = 175;
					this.spacing();
					this.state = 176;
					this.relationDefGrouping();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 180;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 19, this._ctx);
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
	public relationDefPartialAllButNot(): RelationDefPartialAllButNotContext {
		let localctx: RelationDefPartialAllButNotContext = new RelationDefPartialAllButNotContext(this, this._ctx, this.state);
		this.enterRule(localctx, 20, OpenFGAParser.RULE_relationDefPartialAllButNot);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 187;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 182;
					this.spacing();
					this.state = 183;
					this.relationDefOperatorButNot();
					this.state = 184;
					this.spacing();
					this.state = 185;
					this.relationDefGrouping();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 189;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 20, this._ctx);
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
	public relationDefDirectAssignment(): RelationDefDirectAssignmentContext {
		let localctx: RelationDefDirectAssignmentContext = new RelationDefDirectAssignmentContext(this, this._ctx, this.state);
		this.enterRule(localctx, 22, OpenFGAParser.RULE_relationDefDirectAssignment);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 191;
			this.match(OpenFGAParser.L_SQUARE);
			this.state = 193;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 192;
				this.spacing();
				}
			}

			this.state = 195;
			this.relationDefTypeRestriction();
			this.state = 197;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 22, this._ctx) ) {
			case 1:
				{
				this.state = 196;
				this.spacing();
				}
				break;
			}
			this.state = 206;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===13) {
				{
				{
				this.state = 199;
				this.match(OpenFGAParser.COMMA);
				this.state = 201;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===20) {
					{
					this.state = 200;
					this.spacing();
					}
				}

				this.state = 203;
				this.relationDefTypeRestriction();
				}
				}
				this.state = 208;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 210;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 209;
				this.spacing();
				}
			}

			this.state = 212;
			this.match(OpenFGAParser.R_SQUARE);
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
		this.enterRule(localctx, 24, OpenFGAParser.RULE_relationDefRewrite);
		try {
			this.state = 216;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 26, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 214;
				this.relationDefRelationOnSameObject();
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 215;
				this.relationDefRelationOnRelatedObject();
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
	public relationDefRelationOnSameObject(): RelationDefRelationOnSameObjectContext {
		let localctx: RelationDefRelationOnSameObjectContext = new RelationDefRelationOnSameObjectContext(this, this._ctx, this.state);
		this.enterRule(localctx, 26, OpenFGAParser.RULE_relationDefRelationOnSameObject);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 218;
			this.rewriteComputedusersetName();
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
	public relationDefRelationOnRelatedObject(): RelationDefRelationOnRelatedObjectContext {
		let localctx: RelationDefRelationOnRelatedObjectContext = new RelationDefRelationOnRelatedObjectContext(this, this._ctx, this.state);
		this.enterRule(localctx, 28, OpenFGAParser.RULE_relationDefRelationOnRelatedObject);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 220;
			this.rewriteTuplesetComputedusersetName();
			this.state = 221;
			this.spacing();
			this.state = 222;
			this.relationDefKeywordFrom();
			this.state = 223;
			this.spacing();
			this.state = 224;
			this.rewriteTuplesetName();
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
	public relationDefOperator(): RelationDefOperatorContext {
		let localctx: RelationDefOperatorContext = new RelationDefOperatorContext(this, this._ctx, this.state);
		this.enterRule(localctx, 30, OpenFGAParser.RULE_relationDefOperator);
		try {
			this.state = 229;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 15:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 226;
				this.relationDefOperatorOr();
				}
				break;
			case 14:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 227;
				this.relationDefOperatorAnd();
				}
				break;
			case 16:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 228;
				this.relationDefOperatorButNot();
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
	public relationDefOperatorAnd(): RelationDefOperatorAndContext {
		let localctx: RelationDefOperatorAndContext = new RelationDefOperatorAndContext(this, this._ctx, this.state);
		this.enterRule(localctx, 32, OpenFGAParser.RULE_relationDefOperatorAnd);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 231;
			this.match(OpenFGAParser.AND);
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
	public relationDefOperatorOr(): RelationDefOperatorOrContext {
		let localctx: RelationDefOperatorOrContext = new RelationDefOperatorOrContext(this, this._ctx, this.state);
		this.enterRule(localctx, 34, OpenFGAParser.RULE_relationDefOperatorOr);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 233;
			this.match(OpenFGAParser.OR);
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
	public relationDefOperatorButNot(): RelationDefOperatorButNotContext {
		let localctx: RelationDefOperatorButNotContext = new RelationDefOperatorButNotContext(this, this._ctx, this.state);
		this.enterRule(localctx, 36, OpenFGAParser.RULE_relationDefOperatorButNot);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 235;
			this.match(OpenFGAParser.BUT_NOT);
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
	public relationDefKeywordFrom(): RelationDefKeywordFromContext {
		let localctx: RelationDefKeywordFromContext = new RelationDefKeywordFromContext(this, this._ctx, this.state);
		this.enterRule(localctx, 38, OpenFGAParser.RULE_relationDefKeywordFrom);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 237;
			this.match(OpenFGAParser.FROM);
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
		this.enterRule(localctx, 40, OpenFGAParser.RULE_relationDefTypeRestriction);
		try {
			this.state = 242;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 28, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 239;
				this.relationDefTypeRestrictionType();
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 240;
				this.relationDefTypeRestrictionWildcard();
				}
				break;
			case 3:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 241;
				this.relationDefTypeRestrictionUserset();
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
	public relationDefTypeRestrictionType(): RelationDefTypeRestrictionTypeContext {
		let localctx: RelationDefTypeRestrictionTypeContext = new RelationDefTypeRestrictionTypeContext(this, this._ctx, this.state);
		this.enterRule(localctx, 42, OpenFGAParser.RULE_relationDefTypeRestrictionType);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 244;
			this.name();
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
	public relationDefTypeRestrictionRelation(): RelationDefTypeRestrictionRelationContext {
		let localctx: RelationDefTypeRestrictionRelationContext = new RelationDefTypeRestrictionRelationContext(this, this._ctx, this.state);
		this.enterRule(localctx, 44, OpenFGAParser.RULE_relationDefTypeRestrictionRelation);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 246;
			this.name();
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
	public relationDefTypeRestrictionWildcard(): RelationDefTypeRestrictionWildcardContext {
		let localctx: RelationDefTypeRestrictionWildcardContext = new RelationDefTypeRestrictionWildcardContext(this, this._ctx, this.state);
		this.enterRule(localctx, 46, OpenFGAParser.RULE_relationDefTypeRestrictionWildcard);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 248;
			this.relationDefTypeRestrictionType();
			this.state = 249;
			this.match(OpenFGAParser.COLON);
			this.state = 250;
			this.match(OpenFGAParser.WILDCARD);
			this.state = 252;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 29, this._ctx) ) {
			case 1:
				{
				this.state = 251;
				this.spacing();
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
	public relationDefTypeRestrictionUserset(): RelationDefTypeRestrictionUsersetContext {
		let localctx: RelationDefTypeRestrictionUsersetContext = new RelationDefTypeRestrictionUsersetContext(this, this._ctx, this.state);
		this.enterRule(localctx, 48, OpenFGAParser.RULE_relationDefTypeRestrictionUserset);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 254;
			this.relationDefTypeRestrictionType();
			this.state = 255;
			this.match(OpenFGAParser.HASH);
			this.state = 256;
			this.relationDefTypeRestrictionRelation();
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
		this.enterRule(localctx, 50, OpenFGAParser.RULE_relationDefGrouping);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 258;
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
	public rewriteComputedusersetName(): RewriteComputedusersetNameContext {
		let localctx: RewriteComputedusersetNameContext = new RewriteComputedusersetNameContext(this, this._ctx, this.state);
		this.enterRule(localctx, 52, OpenFGAParser.RULE_rewriteComputedusersetName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 260;
			this.name();
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
	public rewriteTuplesetComputedusersetName(): RewriteTuplesetComputedusersetNameContext {
		let localctx: RewriteTuplesetComputedusersetNameContext = new RewriteTuplesetComputedusersetNameContext(this, this._ctx, this.state);
		this.enterRule(localctx, 54, OpenFGAParser.RULE_rewriteTuplesetComputedusersetName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 262;
			this.name();
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
	public rewriteTuplesetName(): RewriteTuplesetNameContext {
		let localctx: RewriteTuplesetNameContext = new RewriteTuplesetNameContext(this, this._ctx, this.state);
		this.enterRule(localctx, 56, OpenFGAParser.RULE_rewriteTuplesetName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 264;
			this.name();
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
		this.enterRule(localctx, 58, OpenFGAParser.RULE_relationName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 266;
			this.name();
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
	public typeName(): TypeNameContext {
		let localctx: TypeNameContext = new TypeNameContext(this, this._ctx, this.state);
		this.enterRule(localctx, 60, OpenFGAParser.RULE_typeName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 268;
			this.name();
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
	public comment(): CommentContext {
		let localctx: CommentContext = new CommentContext(this, this._ctx, this.state);
		this.enterRule(localctx, 62, OpenFGAParser.RULE_comment);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 273;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===20) {
				{
				{
				this.state = 270;
				this.match(OpenFGAParser.WS);
				}
				}
				this.state = 275;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 276;
			this.match(OpenFGAParser.HASH);
			this.state = 280;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 31, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 277;
					_la = this._input.LA(1);
					if(_la<=0 || _la===19) {
					this._errHandler.recoverInline(this);
					}
					else {
						this._errHandler.reportMatch(this);
					    this.consume();
					}
					}
					}
				}
				this.state = 282;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 31, this._ctx);
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
	public multiLineComment(): MultiLineCommentContext {
		let localctx: MultiLineCommentContext = new MultiLineCommentContext(this, this._ctx, this.state);
		this.enterRule(localctx, 64, OpenFGAParser.RULE_multiLineComment);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 283;
			this.comment();
			this.state = 289;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 32, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 284;
					this.newline();
					this.state = 285;
					this.comment();
					}
					}
				}
				this.state = 291;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 32, this._ctx);
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
	public spacing(): SpacingContext {
		let localctx: SpacingContext = new SpacingContext(this, this._ctx, this.state);
		this.enterRule(localctx, 66, OpenFGAParser.RULE_spacing);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 293;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 292;
					this.match(OpenFGAParser.WS);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 295;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 33, this._ctx);
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
	public newline(): NewlineContext {
		let localctx: NewlineContext = new NewlineContext(this, this._ctx, this.state);
		this.enterRule(localctx, 68, OpenFGAParser.RULE_newline);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 298;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			do {
				{
				{
				this.state = 297;
				this.match(OpenFGAParser.NEWLINE);
				}
				}
				this.state = 300;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			} while (_la===19);
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
	public schemaVersion(): SchemaVersionContext {
		let localctx: SchemaVersionContext = new SchemaVersionContext(this, this._ctx, this.state);
		this.enterRule(localctx, 70, OpenFGAParser.RULE_schemaVersion);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 302;
			this.match(OpenFGAParser.SCHEMA_VERSION);
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
	public name(): NameContext {
		let localctx: NameContext = new NameContext(this, this._ctx, this.state);
		this.enterRule(localctx, 72, OpenFGAParser.RULE_name);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 305;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			do {
				{
				{
				this.state = 304;
				this.match(OpenFGAParser.ALPHA_NUMERIC);
				}
				}
				this.state = 307;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			} while (_la===18);
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

	public static readonly _serializedATN: number[] = [4,1,20,310,2,0,7,0,2,
	1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,
	10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,16,2,17,
	7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,2,22,7,22,2,23,7,23,2,24,7,
	24,2,25,7,25,2,26,7,26,2,27,7,27,2,28,7,28,2,29,7,29,2,30,7,30,2,31,7,31,
	2,32,7,32,2,33,7,33,2,34,7,34,2,35,7,35,2,36,7,36,1,0,1,0,1,0,3,0,78,8,
	0,1,1,1,1,1,2,1,2,1,2,3,2,85,8,2,1,2,1,2,3,2,89,8,2,1,2,1,2,1,2,3,2,94,
	8,2,1,2,1,2,1,2,1,2,1,2,3,2,101,8,2,1,3,5,3,104,8,3,10,3,12,3,107,9,3,1,
	4,1,4,1,4,3,4,112,8,4,1,4,1,4,1,4,1,4,1,4,3,4,119,8,4,1,4,1,4,1,4,3,4,124,
	8,4,1,4,4,4,127,8,4,11,4,12,4,128,3,4,131,8,4,1,5,1,5,1,5,3,5,136,8,5,1,
	5,1,5,1,5,1,5,1,5,3,5,143,8,5,1,5,1,5,3,5,147,8,5,1,5,1,5,3,5,151,8,5,1,
	6,1,6,3,6,155,8,6,1,6,3,6,158,8,6,1,7,1,7,1,7,3,7,163,8,7,1,8,1,8,1,8,1,
	8,1,8,4,8,170,8,8,11,8,12,8,171,1,9,1,9,1,9,1,9,1,9,4,9,179,8,9,11,9,12,
	9,180,1,10,1,10,1,10,1,10,1,10,4,10,188,8,10,11,10,12,10,189,1,11,1,11,
	3,11,194,8,11,1,11,1,11,3,11,198,8,11,1,11,1,11,3,11,202,8,11,1,11,5,11,
	205,8,11,10,11,12,11,208,9,11,1,11,3,11,211,8,11,1,11,1,11,1,12,1,12,3,
	12,217,8,12,1,13,1,13,1,14,1,14,1,14,1,14,1,14,1,14,1,15,1,15,1,15,3,15,
	230,8,15,1,16,1,16,1,17,1,17,1,18,1,18,1,19,1,19,1,20,1,20,1,20,3,20,243,
	8,20,1,21,1,21,1,22,1,22,1,23,1,23,1,23,1,23,3,23,253,8,23,1,24,1,24,1,
	24,1,24,1,25,1,25,1,26,1,26,1,27,1,27,1,28,1,28,1,29,1,29,1,30,1,30,1,31,
	5,31,272,8,31,10,31,12,31,275,9,31,1,31,1,31,5,31,279,8,31,10,31,12,31,
	282,9,31,1,32,1,32,1,32,1,32,5,32,288,8,32,10,32,12,32,291,9,32,1,33,4,
	33,294,8,33,11,33,12,33,295,1,34,4,34,299,8,34,11,34,12,34,300,1,35,1,35,
	1,36,4,36,306,8,36,11,36,12,36,307,1,36,0,0,37,0,2,4,6,8,10,12,14,16,18,
	20,22,24,26,28,30,32,34,36,38,40,42,44,46,48,50,52,54,56,58,60,62,64,66,
	68,70,72,0,1,1,0,19,19,311,0,74,1,0,0,0,2,79,1,0,0,0,4,84,1,0,0,0,6,105,
	1,0,0,0,8,111,1,0,0,0,10,135,1,0,0,0,12,154,1,0,0,0,14,162,1,0,0,0,16,169,
	1,0,0,0,18,178,1,0,0,0,20,187,1,0,0,0,22,191,1,0,0,0,24,216,1,0,0,0,26,
	218,1,0,0,0,28,220,1,0,0,0,30,229,1,0,0,0,32,231,1,0,0,0,34,233,1,0,0,0,
	36,235,1,0,0,0,38,237,1,0,0,0,40,242,1,0,0,0,42,244,1,0,0,0,44,246,1,0,
	0,0,46,248,1,0,0,0,48,254,1,0,0,0,50,258,1,0,0,0,52,260,1,0,0,0,54,262,
	1,0,0,0,56,264,1,0,0,0,58,266,1,0,0,0,60,268,1,0,0,0,62,273,1,0,0,0,64,
	283,1,0,0,0,66,293,1,0,0,0,68,298,1,0,0,0,70,302,1,0,0,0,72,305,1,0,0,0,
	74,75,3,4,2,0,75,77,3,6,3,0,76,78,3,68,34,0,77,76,1,0,0,0,77,78,1,0,0,0,
	78,1,1,0,0,0,79,80,5,1,0,0,80,3,1,0,0,0,81,82,3,64,32,0,82,83,3,68,34,0,
	83,85,1,0,0,0,84,81,1,0,0,0,84,85,1,0,0,0,85,86,1,0,0,0,86,88,5,2,0,0,87,
	89,3,66,33,0,88,87,1,0,0,0,88,89,1,0,0,0,89,93,1,0,0,0,90,91,3,68,34,0,
	91,92,3,64,32,0,92,94,1,0,0,0,93,90,1,0,0,0,93,94,1,0,0,0,94,95,1,0,0,0,
	95,96,3,2,1,0,96,97,5,3,0,0,97,98,3,66,33,0,98,100,3,70,35,0,99,101,3,66,
	33,0,100,99,1,0,0,0,100,101,1,0,0,0,101,5,1,0,0,0,102,104,3,8,4,0,103,102,
	1,0,0,0,104,107,1,0,0,0,105,103,1,0,0,0,105,106,1,0,0,0,106,7,1,0,0,0,107,
	105,1,0,0,0,108,109,3,68,34,0,109,110,3,64,32,0,110,112,1,0,0,0,111,108,
	1,0,0,0,111,112,1,0,0,0,112,113,1,0,0,0,113,114,3,68,34,0,114,115,5,5,0,
	0,115,116,3,66,33,0,116,118,3,60,30,0,117,119,3,66,33,0,118,117,1,0,0,0,
	118,119,1,0,0,0,119,130,1,0,0,0,120,121,3,2,1,0,121,123,5,6,0,0,122,124,
	3,66,33,0,123,122,1,0,0,0,123,124,1,0,0,0,124,126,1,0,0,0,125,127,3,10,
	5,0,126,125,1,0,0,0,127,128,1,0,0,0,128,126,1,0,0,0,128,129,1,0,0,0,129,
	131,1,0,0,0,130,120,1,0,0,0,130,131,1,0,0,0,131,9,1,0,0,0,132,133,3,68,
	34,0,133,134,3,64,32,0,134,136,1,0,0,0,135,132,1,0,0,0,135,136,1,0,0,0,
	136,137,1,0,0,0,137,138,3,2,1,0,138,139,5,7,0,0,139,140,3,66,33,0,140,142,
	3,58,29,0,141,143,3,66,33,0,142,141,1,0,0,0,142,143,1,0,0,0,143,144,1,0,
	0,0,144,146,5,9,0,0,145,147,3,66,33,0,146,145,1,0,0,0,146,147,1,0,0,0,147,
	148,1,0,0,0,148,150,3,12,6,0,149,151,3,66,33,0,150,149,1,0,0,0,150,151,
	1,0,0,0,151,11,1,0,0,0,152,155,3,22,11,0,153,155,3,50,25,0,154,152,1,0,
	0,0,154,153,1,0,0,0,155,157,1,0,0,0,156,158,3,14,7,0,157,156,1,0,0,0,157,
	158,1,0,0,0,158,13,1,0,0,0,159,163,3,16,8,0,160,163,3,18,9,0,161,163,3,
	20,10,0,162,159,1,0,0,0,162,160,1,0,0,0,162,161,1,0,0,0,163,15,1,0,0,0,
	164,165,3,66,33,0,165,166,3,34,17,0,166,167,3,66,33,0,167,168,3,50,25,0,
	168,170,1,0,0,0,169,164,1,0,0,0,170,171,1,0,0,0,171,169,1,0,0,0,171,172,
	1,0,0,0,172,17,1,0,0,0,173,174,3,66,33,0,174,175,3,32,16,0,175,176,3,66,
	33,0,176,177,3,50,25,0,177,179,1,0,0,0,178,173,1,0,0,0,179,180,1,0,0,0,
	180,178,1,0,0,0,180,181,1,0,0,0,181,19,1,0,0,0,182,183,3,66,33,0,183,184,
	3,36,18,0,184,185,3,66,33,0,185,186,3,50,25,0,186,188,1,0,0,0,187,182,1,
	0,0,0,188,189,1,0,0,0,189,187,1,0,0,0,189,190,1,0,0,0,190,21,1,0,0,0,191,
	193,5,11,0,0,192,194,3,66,33,0,193,192,1,0,0,0,193,194,1,0,0,0,194,195,
	1,0,0,0,195,197,3,40,20,0,196,198,3,66,33,0,197,196,1,0,0,0,197,198,1,0,
	0,0,198,206,1,0,0,0,199,201,5,13,0,0,200,202,3,66,33,0,201,200,1,0,0,0,
	201,202,1,0,0,0,202,203,1,0,0,0,203,205,3,40,20,0,204,199,1,0,0,0,205,208,
	1,0,0,0,206,204,1,0,0,0,206,207,1,0,0,0,207,210,1,0,0,0,208,206,1,0,0,0,
	209,211,3,66,33,0,210,209,1,0,0,0,210,211,1,0,0,0,211,212,1,0,0,0,212,213,
	5,12,0,0,213,23,1,0,0,0,214,217,3,26,13,0,215,217,3,28,14,0,216,214,1,0,
	0,0,216,215,1,0,0,0,217,25,1,0,0,0,218,219,3,52,26,0,219,27,1,0,0,0,220,
	221,3,54,27,0,221,222,3,66,33,0,222,223,3,38,19,0,223,224,3,66,33,0,224,
	225,3,56,28,0,225,29,1,0,0,0,226,230,3,34,17,0,227,230,3,32,16,0,228,230,
	3,36,18,0,229,226,1,0,0,0,229,227,1,0,0,0,229,228,1,0,0,0,230,31,1,0,0,
	0,231,232,5,14,0,0,232,33,1,0,0,0,233,234,5,15,0,0,234,35,1,0,0,0,235,236,
	5,16,0,0,236,37,1,0,0,0,237,238,5,17,0,0,238,39,1,0,0,0,239,243,3,42,21,
	0,240,243,3,46,23,0,241,243,3,48,24,0,242,239,1,0,0,0,242,240,1,0,0,0,242,
	241,1,0,0,0,243,41,1,0,0,0,244,245,3,72,36,0,245,43,1,0,0,0,246,247,3,72,
	36,0,247,45,1,0,0,0,248,249,3,42,21,0,249,250,5,9,0,0,250,252,5,10,0,0,
	251,253,3,66,33,0,252,251,1,0,0,0,252,253,1,0,0,0,253,47,1,0,0,0,254,255,
	3,42,21,0,255,256,5,8,0,0,256,257,3,44,22,0,257,49,1,0,0,0,258,259,3,24,
	12,0,259,51,1,0,0,0,260,261,3,72,36,0,261,53,1,0,0,0,262,263,3,72,36,0,
	263,55,1,0,0,0,264,265,3,72,36,0,265,57,1,0,0,0,266,267,3,72,36,0,267,59,
	1,0,0,0,268,269,3,72,36,0,269,61,1,0,0,0,270,272,5,20,0,0,271,270,1,0,0,
	0,272,275,1,0,0,0,273,271,1,0,0,0,273,274,1,0,0,0,274,276,1,0,0,0,275,273,
	1,0,0,0,276,280,5,8,0,0,277,279,8,0,0,0,278,277,1,0,0,0,279,282,1,0,0,0,
	280,278,1,0,0,0,280,281,1,0,0,0,281,63,1,0,0,0,282,280,1,0,0,0,283,289,
	3,62,31,0,284,285,3,68,34,0,285,286,3,62,31,0,286,288,1,0,0,0,287,284,1,
	0,0,0,288,291,1,0,0,0,289,287,1,0,0,0,289,290,1,0,0,0,290,65,1,0,0,0,291,
	289,1,0,0,0,292,294,5,20,0,0,293,292,1,0,0,0,294,295,1,0,0,0,295,293,1,
	0,0,0,295,296,1,0,0,0,296,67,1,0,0,0,297,299,5,19,0,0,298,297,1,0,0,0,299,
	300,1,0,0,0,300,298,1,0,0,0,300,301,1,0,0,0,301,69,1,0,0,0,302,303,5,4,
	0,0,303,71,1,0,0,0,304,306,5,18,0,0,305,304,1,0,0,0,306,307,1,0,0,0,307,
	305,1,0,0,0,307,308,1,0,0,0,308,73,1,0,0,0,36,77,84,88,93,100,105,111,118,
	123,128,130,135,142,146,150,154,157,162,171,180,189,193,197,201,206,210,
	216,229,242,252,273,280,289,295,300,307];

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
	public newline(): NewlineContext {
		return this.getTypedRuleContext(NewlineContext, 0) as NewlineContext;
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


export class IndentationContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public INDENT(): TerminalNode {
		return this.getToken(OpenFGAParser.INDENT, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_indentation;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterIndentation) {
	 		listener.enterIndentation(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitIndentation) {
	 		listener.exitIndentation(this);
		}
	}
}


export class ModelHeaderContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public MODEL(): TerminalNode {
		return this.getToken(OpenFGAParser.MODEL, 0);
	}
	public indentation(): IndentationContext {
		return this.getTypedRuleContext(IndentationContext, 0) as IndentationContext;
	}
	public SCHEMA(): TerminalNode {
		return this.getToken(OpenFGAParser.SCHEMA, 0);
	}
	public spacing_list(): SpacingContext[] {
		return this.getTypedRuleContexts(SpacingContext) as SpacingContext[];
	}
	public spacing(i: number): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, i) as SpacingContext;
	}
	public schemaVersion(): SchemaVersionContext {
		return this.getTypedRuleContext(SchemaVersionContext, 0) as SchemaVersionContext;
	}
	public multiLineComment_list(): MultiLineCommentContext[] {
		return this.getTypedRuleContexts(MultiLineCommentContext) as MultiLineCommentContext[];
	}
	public multiLineComment(i: number): MultiLineCommentContext {
		return this.getTypedRuleContext(MultiLineCommentContext, i) as MultiLineCommentContext;
	}
	public newline_list(): NewlineContext[] {
		return this.getTypedRuleContexts(NewlineContext) as NewlineContext[];
	}
	public newline(i: number): NewlineContext {
		return this.getTypedRuleContext(NewlineContext, i) as NewlineContext;
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
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public newline_list(): NewlineContext[] {
		return this.getTypedRuleContexts(NewlineContext) as NewlineContext[];
	}
	public newline(i: number): NewlineContext {
		return this.getTypedRuleContext(NewlineContext, i) as NewlineContext;
	}
	public TYPE(): TerminalNode {
		return this.getToken(OpenFGAParser.TYPE, 0);
	}
	public spacing_list(): SpacingContext[] {
		return this.getTypedRuleContexts(SpacingContext) as SpacingContext[];
	}
	public spacing(i: number): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, i) as SpacingContext;
	}
	public typeName(): TypeNameContext {
		return this.getTypedRuleContext(TypeNameContext, 0) as TypeNameContext;
	}
	public multiLineComment(): MultiLineCommentContext {
		return this.getTypedRuleContext(MultiLineCommentContext, 0) as MultiLineCommentContext;
	}
	public indentation(): IndentationContext {
		return this.getTypedRuleContext(IndentationContext, 0) as IndentationContext;
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
	public indentation(): IndentationContext {
		return this.getTypedRuleContext(IndentationContext, 0) as IndentationContext;
	}
	public DEFINE(): TerminalNode {
		return this.getToken(OpenFGAParser.DEFINE, 0);
	}
	public spacing_list(): SpacingContext[] {
		return this.getTypedRuleContexts(SpacingContext) as SpacingContext[];
	}
	public spacing(i: number): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, i) as SpacingContext;
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
	public newline(): NewlineContext {
		return this.getTypedRuleContext(NewlineContext, 0) as NewlineContext;
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
	public relationDefPartialAllOr(): RelationDefPartialAllOrContext {
		return this.getTypedRuleContext(RelationDefPartialAllOrContext, 0) as RelationDefPartialAllOrContext;
	}
	public relationDefPartialAllAnd(): RelationDefPartialAllAndContext {
		return this.getTypedRuleContext(RelationDefPartialAllAndContext, 0) as RelationDefPartialAllAndContext;
	}
	public relationDefPartialAllButNot(): RelationDefPartialAllButNotContext {
		return this.getTypedRuleContext(RelationDefPartialAllButNotContext, 0) as RelationDefPartialAllButNotContext;
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


export class RelationDefPartialAllOrContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public spacing_list(): SpacingContext[] {
		return this.getTypedRuleContexts(SpacingContext) as SpacingContext[];
	}
	public spacing(i: number): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, i) as SpacingContext;
	}
	public relationDefOperatorOr_list(): RelationDefOperatorOrContext[] {
		return this.getTypedRuleContexts(RelationDefOperatorOrContext) as RelationDefOperatorOrContext[];
	}
	public relationDefOperatorOr(i: number): RelationDefOperatorOrContext {
		return this.getTypedRuleContext(RelationDefOperatorOrContext, i) as RelationDefOperatorOrContext;
	}
	public relationDefGrouping_list(): RelationDefGroupingContext[] {
		return this.getTypedRuleContexts(RelationDefGroupingContext) as RelationDefGroupingContext[];
	}
	public relationDefGrouping(i: number): RelationDefGroupingContext {
		return this.getTypedRuleContext(RelationDefGroupingContext, i) as RelationDefGroupingContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefPartialAllOr;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefPartialAllOr) {
	 		listener.enterRelationDefPartialAllOr(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefPartialAllOr) {
	 		listener.exitRelationDefPartialAllOr(this);
		}
	}
}


export class RelationDefPartialAllAndContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public spacing_list(): SpacingContext[] {
		return this.getTypedRuleContexts(SpacingContext) as SpacingContext[];
	}
	public spacing(i: number): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, i) as SpacingContext;
	}
	public relationDefOperatorAnd_list(): RelationDefOperatorAndContext[] {
		return this.getTypedRuleContexts(RelationDefOperatorAndContext) as RelationDefOperatorAndContext[];
	}
	public relationDefOperatorAnd(i: number): RelationDefOperatorAndContext {
		return this.getTypedRuleContext(RelationDefOperatorAndContext, i) as RelationDefOperatorAndContext;
	}
	public relationDefGrouping_list(): RelationDefGroupingContext[] {
		return this.getTypedRuleContexts(RelationDefGroupingContext) as RelationDefGroupingContext[];
	}
	public relationDefGrouping(i: number): RelationDefGroupingContext {
		return this.getTypedRuleContext(RelationDefGroupingContext, i) as RelationDefGroupingContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefPartialAllAnd;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefPartialAllAnd) {
	 		listener.enterRelationDefPartialAllAnd(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefPartialAllAnd) {
	 		listener.exitRelationDefPartialAllAnd(this);
		}
	}
}


export class RelationDefPartialAllButNotContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public spacing_list(): SpacingContext[] {
		return this.getTypedRuleContexts(SpacingContext) as SpacingContext[];
	}
	public spacing(i: number): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, i) as SpacingContext;
	}
	public relationDefOperatorButNot_list(): RelationDefOperatorButNotContext[] {
		return this.getTypedRuleContexts(RelationDefOperatorButNotContext) as RelationDefOperatorButNotContext[];
	}
	public relationDefOperatorButNot(i: number): RelationDefOperatorButNotContext {
		return this.getTypedRuleContext(RelationDefOperatorButNotContext, i) as RelationDefOperatorButNotContext;
	}
	public relationDefGrouping_list(): RelationDefGroupingContext[] {
		return this.getTypedRuleContexts(RelationDefGroupingContext) as RelationDefGroupingContext[];
	}
	public relationDefGrouping(i: number): RelationDefGroupingContext {
		return this.getTypedRuleContext(RelationDefGroupingContext, i) as RelationDefGroupingContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefPartialAllButNot;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefPartialAllButNot) {
	 		listener.enterRelationDefPartialAllButNot(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefPartialAllButNot) {
	 		listener.exitRelationDefPartialAllButNot(this);
		}
	}
}


export class RelationDefDirectAssignmentContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public L_SQUARE(): TerminalNode {
		return this.getToken(OpenFGAParser.L_SQUARE, 0);
	}
	public relationDefTypeRestriction_list(): RelationDefTypeRestrictionContext[] {
		return this.getTypedRuleContexts(RelationDefTypeRestrictionContext) as RelationDefTypeRestrictionContext[];
	}
	public relationDefTypeRestriction(i: number): RelationDefTypeRestrictionContext {
		return this.getTypedRuleContext(RelationDefTypeRestrictionContext, i) as RelationDefTypeRestrictionContext;
	}
	public R_SQUARE(): TerminalNode {
		return this.getToken(OpenFGAParser.R_SQUARE, 0);
	}
	public spacing_list(): SpacingContext[] {
		return this.getTypedRuleContexts(SpacingContext) as SpacingContext[];
	}
	public spacing(i: number): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, i) as SpacingContext;
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
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public relationDefRelationOnSameObject(): RelationDefRelationOnSameObjectContext {
		return this.getTypedRuleContext(RelationDefRelationOnSameObjectContext, 0) as RelationDefRelationOnSameObjectContext;
	}
	public relationDefRelationOnRelatedObject(): RelationDefRelationOnRelatedObjectContext {
		return this.getTypedRuleContext(RelationDefRelationOnRelatedObjectContext, 0) as RelationDefRelationOnRelatedObjectContext;
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


export class RelationDefRelationOnSameObjectContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public rewriteComputedusersetName(): RewriteComputedusersetNameContext {
		return this.getTypedRuleContext(RewriteComputedusersetNameContext, 0) as RewriteComputedusersetNameContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefRelationOnSameObject;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefRelationOnSameObject) {
	 		listener.enterRelationDefRelationOnSameObject(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefRelationOnSameObject) {
	 		listener.exitRelationDefRelationOnSameObject(this);
		}
	}
}


export class RelationDefRelationOnRelatedObjectContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public rewriteTuplesetComputedusersetName(): RewriteTuplesetComputedusersetNameContext {
		return this.getTypedRuleContext(RewriteTuplesetComputedusersetNameContext, 0) as RewriteTuplesetComputedusersetNameContext;
	}
	public spacing_list(): SpacingContext[] {
		return this.getTypedRuleContexts(SpacingContext) as SpacingContext[];
	}
	public spacing(i: number): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, i) as SpacingContext;
	}
	public relationDefKeywordFrom(): RelationDefKeywordFromContext {
		return this.getTypedRuleContext(RelationDefKeywordFromContext, 0) as RelationDefKeywordFromContext;
	}
	public rewriteTuplesetName(): RewriteTuplesetNameContext {
		return this.getTypedRuleContext(RewriteTuplesetNameContext, 0) as RewriteTuplesetNameContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefRelationOnRelatedObject;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefRelationOnRelatedObject) {
	 		listener.enterRelationDefRelationOnRelatedObject(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefRelationOnRelatedObject) {
	 		listener.exitRelationDefRelationOnRelatedObject(this);
		}
	}
}


export class RelationDefOperatorContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public relationDefOperatorOr(): RelationDefOperatorOrContext {
		return this.getTypedRuleContext(RelationDefOperatorOrContext, 0) as RelationDefOperatorOrContext;
	}
	public relationDefOperatorAnd(): RelationDefOperatorAndContext {
		return this.getTypedRuleContext(RelationDefOperatorAndContext, 0) as RelationDefOperatorAndContext;
	}
	public relationDefOperatorButNot(): RelationDefOperatorButNotContext {
		return this.getTypedRuleContext(RelationDefOperatorButNotContext, 0) as RelationDefOperatorButNotContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefOperator;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefOperator) {
	 		listener.enterRelationDefOperator(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefOperator) {
	 		listener.exitRelationDefOperator(this);
		}
	}
}


export class RelationDefOperatorAndContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public AND(): TerminalNode {
		return this.getToken(OpenFGAParser.AND, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefOperatorAnd;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefOperatorAnd) {
	 		listener.enterRelationDefOperatorAnd(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefOperatorAnd) {
	 		listener.exitRelationDefOperatorAnd(this);
		}
	}
}


export class RelationDefOperatorOrContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public OR(): TerminalNode {
		return this.getToken(OpenFGAParser.OR, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefOperatorOr;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefOperatorOr) {
	 		listener.enterRelationDefOperatorOr(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefOperatorOr) {
	 		listener.exitRelationDefOperatorOr(this);
		}
	}
}


export class RelationDefOperatorButNotContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public BUT_NOT(): TerminalNode {
		return this.getToken(OpenFGAParser.BUT_NOT, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefOperatorButNot;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefOperatorButNot) {
	 		listener.enterRelationDefOperatorButNot(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefOperatorButNot) {
	 		listener.exitRelationDefOperatorButNot(this);
		}
	}
}


export class RelationDefKeywordFromContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public FROM(): TerminalNode {
		return this.getToken(OpenFGAParser.FROM, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefKeywordFrom;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefKeywordFrom) {
	 		listener.enterRelationDefKeywordFrom(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefKeywordFrom) {
	 		listener.exitRelationDefKeywordFrom(this);
		}
	}
}


export class RelationDefTypeRestrictionContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public relationDefTypeRestrictionType(): RelationDefTypeRestrictionTypeContext {
		return this.getTypedRuleContext(RelationDefTypeRestrictionTypeContext, 0) as RelationDefTypeRestrictionTypeContext;
	}
	public relationDefTypeRestrictionWildcard(): RelationDefTypeRestrictionWildcardContext {
		return this.getTypedRuleContext(RelationDefTypeRestrictionWildcardContext, 0) as RelationDefTypeRestrictionWildcardContext;
	}
	public relationDefTypeRestrictionUserset(): RelationDefTypeRestrictionUsersetContext {
		return this.getTypedRuleContext(RelationDefTypeRestrictionUsersetContext, 0) as RelationDefTypeRestrictionUsersetContext;
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


export class RelationDefTypeRestrictionTypeContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public name(): NameContext {
		return this.getTypedRuleContext(NameContext, 0) as NameContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefTypeRestrictionType;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefTypeRestrictionType) {
	 		listener.enterRelationDefTypeRestrictionType(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefTypeRestrictionType) {
	 		listener.exitRelationDefTypeRestrictionType(this);
		}
	}
}


export class RelationDefTypeRestrictionRelationContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public name(): NameContext {
		return this.getTypedRuleContext(NameContext, 0) as NameContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefTypeRestrictionRelation;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefTypeRestrictionRelation) {
	 		listener.enterRelationDefTypeRestrictionRelation(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefTypeRestrictionRelation) {
	 		listener.exitRelationDefTypeRestrictionRelation(this);
		}
	}
}


export class RelationDefTypeRestrictionWildcardContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public relationDefTypeRestrictionType(): RelationDefTypeRestrictionTypeContext {
		return this.getTypedRuleContext(RelationDefTypeRestrictionTypeContext, 0) as RelationDefTypeRestrictionTypeContext;
	}
	public COLON(): TerminalNode {
		return this.getToken(OpenFGAParser.COLON, 0);
	}
	public WILDCARD(): TerminalNode {
		return this.getToken(OpenFGAParser.WILDCARD, 0);
	}
	public spacing(): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, 0) as SpacingContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefTypeRestrictionWildcard;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefTypeRestrictionWildcard) {
	 		listener.enterRelationDefTypeRestrictionWildcard(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefTypeRestrictionWildcard) {
	 		listener.exitRelationDefTypeRestrictionWildcard(this);
		}
	}
}


export class RelationDefTypeRestrictionUsersetContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public relationDefTypeRestrictionType(): RelationDefTypeRestrictionTypeContext {
		return this.getTypedRuleContext(RelationDefTypeRestrictionTypeContext, 0) as RelationDefTypeRestrictionTypeContext;
	}
	public HASH(): TerminalNode {
		return this.getToken(OpenFGAParser.HASH, 0);
	}
	public relationDefTypeRestrictionRelation(): RelationDefTypeRestrictionRelationContext {
		return this.getTypedRuleContext(RelationDefTypeRestrictionRelationContext, 0) as RelationDefTypeRestrictionRelationContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefTypeRestrictionUserset;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefTypeRestrictionUserset) {
	 		listener.enterRelationDefTypeRestrictionUserset(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefTypeRestrictionUserset) {
	 		listener.exitRelationDefTypeRestrictionUserset(this);
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


export class RewriteComputedusersetNameContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public name(): NameContext {
		return this.getTypedRuleContext(NameContext, 0) as NameContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_rewriteComputedusersetName;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRewriteComputedusersetName) {
	 		listener.enterRewriteComputedusersetName(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRewriteComputedusersetName) {
	 		listener.exitRewriteComputedusersetName(this);
		}
	}
}


export class RewriteTuplesetComputedusersetNameContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public name(): NameContext {
		return this.getTypedRuleContext(NameContext, 0) as NameContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_rewriteTuplesetComputedusersetName;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRewriteTuplesetComputedusersetName) {
	 		listener.enterRewriteTuplesetComputedusersetName(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRewriteTuplesetComputedusersetName) {
	 		listener.exitRewriteTuplesetComputedusersetName(this);
		}
	}
}


export class RewriteTuplesetNameContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public name(): NameContext {
		return this.getTypedRuleContext(NameContext, 0) as NameContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_rewriteTuplesetName;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRewriteTuplesetName) {
	 		listener.enterRewriteTuplesetName(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRewriteTuplesetName) {
	 		listener.exitRewriteTuplesetName(this);
		}
	}
}


export class RelationNameContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public name(): NameContext {
		return this.getTypedRuleContext(NameContext, 0) as NameContext;
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


export class TypeNameContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public name(): NameContext {
		return this.getTypedRuleContext(NameContext, 0) as NameContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_typeName;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterTypeName) {
	 		listener.enterTypeName(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitTypeName) {
	 		listener.exitTypeName(this);
		}
	}
}


export class CommentContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public HASH(): TerminalNode {
		return this.getToken(OpenFGAParser.HASH, 0);
	}
	public WS_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WS);
	}
	public WS(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WS, i);
	}
	public NEWLINE_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.NEWLINE);
	}
	public NEWLINE(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.NEWLINE, i);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_comment;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterComment) {
	 		listener.enterComment(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitComment) {
	 		listener.exitComment(this);
		}
	}
}


export class MultiLineCommentContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public comment_list(): CommentContext[] {
		return this.getTypedRuleContexts(CommentContext) as CommentContext[];
	}
	public comment(i: number): CommentContext {
		return this.getTypedRuleContext(CommentContext, i) as CommentContext;
	}
	public newline_list(): NewlineContext[] {
		return this.getTypedRuleContexts(NewlineContext) as NewlineContext[];
	}
	public newline(i: number): NewlineContext {
		return this.getTypedRuleContext(NewlineContext, i) as NewlineContext;
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


export class SpacingContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public WS_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WS);
	}
	public WS(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WS, i);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_spacing;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterSpacing) {
	 		listener.enterSpacing(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitSpacing) {
	 		listener.exitSpacing(this);
		}
	}
}


export class NewlineContext extends ParserRuleContext {
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
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_newline;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterNewline) {
	 		listener.enterNewline(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitNewline) {
	 		listener.exitNewline(this);
		}
	}
}


export class SchemaVersionContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public SCHEMA_VERSION(): TerminalNode {
		return this.getToken(OpenFGAParser.SCHEMA_VERSION, 0);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_schemaVersion;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterSchemaVersion) {
	 		listener.enterSchemaVersion(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitSchemaVersion) {
	 		listener.exitSchemaVersion(this);
		}
	}
}


export class NameContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public ALPHA_NUMERIC_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.ALPHA_NUMERIC);
	}
	public ALPHA_NUMERIC(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.ALPHA_NUMERIC, i);
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_name;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterName) {
	 		listener.enterName(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitName) {
	 		listener.exitName(this);
		}
	}
}
