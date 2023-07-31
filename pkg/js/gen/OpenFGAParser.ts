// Generated from /app/OpenFGA.g4 by ANTLR 4.13.0
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
import OpenFGAListener from "./OpenFGAListener.js";
// for running tests with parameters, TODO: discuss strategy for typed parameters in CI
// eslint-disable-next-line no-unused-vars
type int = number;

export default class OpenFGAParser extends Parser {
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
	public static readonly ALPHA_NUMERIC = 22;
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
                                                            "'#'", "'\\r'", 
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
                                                             "ALPHA_NUMERIC" ];
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
	public get grammarFileName(): string { return "OpenFGA.g4"; }
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
			if (_la===18 || _la===19) {
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
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 79;
			_la = this._input.LA(1);
			if(!(_la===1 || _la===2)) {
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
			if (_la===17 || _la===20) {
				{
				this.state = 81;
				this.multiLineComment();
				this.state = 82;
				this.newline();
				}
			}

			this.state = 86;
			this.match(OpenFGAParser.T__2);
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
			switch ( this._interp.adaptivePredict(this._input, 3, this._ctx) ) {
			case 1:
				{
				this.state = 90;
				this.newline();
				this.state = 91;
				this.multiLineComment();
				}
				break;
			}
			this.state = 95;
			this.newline();
			this.state = 96;
			this.indentation();
			this.state = 97;
			this.match(OpenFGAParser.T__3);
			this.state = 98;
			this.spacing();
			this.state = 99;
			this.schemaVersion();
			this.state = 101;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 100;
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
			this.state = 106;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 5, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 103;
					this.typeDef();
					}
					}
				}
				this.state = 108;
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
			this.state = 112;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 6, this._ctx) ) {
			case 1:
				{
				this.state = 109;
				this.newline();
				this.state = 110;
				this.multiLineComment();
				}
				break;
			}
			this.state = 114;
			this.newline();
			this.state = 115;
			this.match(OpenFGAParser.T__4);
			this.state = 116;
			this.spacing();
			this.state = 117;
			this.typeName();
			this.state = 119;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 118;
				this.spacing();
				}
			}

			this.state = 132;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 10, this._ctx) ) {
			case 1:
				{
				this.state = 121;
				this.newline();
				this.state = 122;
				this.indentation();
				this.state = 123;
				this.match(OpenFGAParser.T__5);
				this.state = 125;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===20) {
					{
					this.state = 124;
					this.spacing();
					}
				}

				this.state = 128;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 127;
						this.relationDeclaration();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					this.state = 130;
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
		this.enterRule(localctx, 10, OpenFGAParser.RULE_relationDeclaration);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 137;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 11, this._ctx) ) {
			case 1:
				{
				this.state = 134;
				this.newline();
				this.state = 135;
				this.multiLineComment();
				}
				break;
			}
			this.state = 139;
			this.newline();
			this.state = 140;
			this.indentation();
			this.state = 141;
			this.indentation();
			this.state = 142;
			this.match(OpenFGAParser.T__6);
			this.state = 143;
			this.spacing();
			this.state = 144;
			this.relationName();
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
			this.match(OpenFGAParser.T__7);
			this.state = 150;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 149;
				this.spacing();
				}
			}

			this.state = 152;
			this.relationDef();
			this.state = 154;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 153;
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
			this.state = 158;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 9:
				{
				this.state = 156;
				this.relationDefDirectAssignment();
				}
				break;
			case 22:
				{
				this.state = 157;
				this.relationDefGrouping();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this.state = 161;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 16, this._ctx) ) {
			case 1:
				{
				this.state = 160;
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
			this.state = 166;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 17, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 163;
				this.relationDefPartialAllOr();
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 164;
				this.relationDefPartialAllAnd();
				}
				break;
			case 3:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 165;
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
			this.state = 173;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 168;
					this.spacing();
					this.state = 169;
					this.relationDefOperatorOr();
					this.state = 170;
					this.spacing();
					this.state = 171;
					this.relationDefGrouping();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 175;
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
			this.state = 182;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 177;
					this.spacing();
					this.state = 178;
					this.relationDefOperatorAnd();
					this.state = 179;
					this.spacing();
					this.state = 180;
					this.relationDefGrouping();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 184;
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
			this.state = 191;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 186;
					this.spacing();
					this.state = 187;
					this.relationDefOperatorButNot();
					this.state = 188;
					this.spacing();
					this.state = 189;
					this.relationDefGrouping();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 193;
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
			this.state = 195;
			this.match(OpenFGAParser.T__8);
			this.state = 196;
			this.relationDefTypeRestriction();
			this.state = 198;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 21, this._ctx) ) {
			case 1:
				{
				this.state = 197;
				this.spacing();
				}
				break;
			}
			this.state = 207;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===10) {
				{
				{
				this.state = 200;
				this.match(OpenFGAParser.T__9);
				this.state = 202;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===20) {
					{
					this.state = 201;
					this.spacing();
					}
				}

				this.state = 204;
				this.relationDefTypeRestriction();
				}
				}
				this.state = 209;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 211;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 210;
				this.spacing();
				}
			}

			this.state = 213;
			this.match(OpenFGAParser.T__10);
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
			this.state = 217;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 25, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 215;
				this.relationDefRelationOnSameObject();
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 216;
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
			this.state = 219;
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
			this.state = 221;
			this.rewriteTuplesetComputedusersetName();
			this.state = 222;
			this.spacing();
			this.state = 223;
			this.relationDefKeywordFrom();
			this.state = 224;
			this.spacing();
			this.state = 225;
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
			this.state = 230;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 13:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 227;
				this.relationDefOperatorOr();
				}
				break;
			case 12:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 228;
				this.relationDefOperatorAnd();
				}
				break;
			case 14:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 229;
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
			this.state = 232;
			this.match(OpenFGAParser.T__11);
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
			this.state = 234;
			this.match(OpenFGAParser.T__12);
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
			this.state = 236;
			this.match(OpenFGAParser.T__13);
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
			this.state = 238;
			this.match(OpenFGAParser.T__14);
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
			this.state = 243;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 27, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 240;
				this.relationDefTypeRestrictionType();
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 241;
				this.relationDefTypeRestrictionWildcard();
				}
				break;
			case 3:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 242;
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
			this.state = 245;
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
			this.state = 247;
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
			this.state = 249;
			this.relationDefTypeRestrictionType();
			this.state = 250;
			this.match(OpenFGAParser.T__15);
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
			this.state = 252;
			this.relationDefTypeRestrictionType();
			this.state = 253;
			this.match(OpenFGAParser.T__16);
			this.state = 254;
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
			this.state = 256;
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
			this.state = 258;
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
	public rewriteTuplesetName(): RewriteTuplesetNameContext {
		let localctx: RewriteTuplesetNameContext = new RewriteTuplesetNameContext(this, this._ctx, this.state);
		this.enterRule(localctx, 56, OpenFGAParser.RULE_rewriteTuplesetName);
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
	public relationName(): RelationNameContext {
		let localctx: RelationNameContext = new RelationNameContext(this, this._ctx, this.state);
		this.enterRule(localctx, 58, OpenFGAParser.RULE_relationName);
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
	public typeName(): TypeNameContext {
		let localctx: TypeNameContext = new TypeNameContext(this, this._ctx, this.state);
		this.enterRule(localctx, 60, OpenFGAParser.RULE_typeName);
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
	public comment(): CommentContext {
		let localctx: CommentContext = new CommentContext(this, this._ctx, this.state);
		this.enterRule(localctx, 62, OpenFGAParser.RULE_comment);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 269;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===20) {
				{
				this.state = 268;
				this.spacing();
				}
			}

			this.state = 271;
			this.match(OpenFGAParser.T__16);
			this.state = 275;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while ((((_la) & ~0x1F) === 0 && ((1 << _la) & 7602174) !== 0)) {
				{
				{
				this.state = 272;
				_la = this._input.LA(1);
				if(_la<=0 || _la===18 || _la===19) {
				this._errHandler.recoverInline(this);
				}
				else {
					this._errHandler.reportMatch(this);
				    this.consume();
				}
				}
				}
				this.state = 277;
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
	// @RuleVersion(0)
	public multiLineComment(): MultiLineCommentContext {
		let localctx: MultiLineCommentContext = new MultiLineCommentContext(this, this._ctx, this.state);
		this.enterRule(localctx, 64, OpenFGAParser.RULE_multiLineComment);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 278;
			this.comment();
			this.state = 284;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 30, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 279;
					this.newline();
					this.state = 280;
					this.comment();
					}
					}
				}
				this.state = 286;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 30, this._ctx);
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
			this.state = 288;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 287;
					this.match(OpenFGAParser.T__19);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 290;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 31, this._ctx);
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
			this.state = 293;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			do {
				{
				{
				this.state = 292;
				_la = this._input.LA(1);
				if(!(_la===18 || _la===19)) {
				this._errHandler.recoverInline(this);
				}
				else {
					this._errHandler.reportMatch(this);
				    this.consume();
				}
				}
				}
				this.state = 295;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			} while (_la===18 || _la===19);
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
			this.state = 297;
			this.match(OpenFGAParser.T__20);
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
			this.state = 300;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			do {
				{
				{
				this.state = 299;
				this.match(OpenFGAParser.ALPHA_NUMERIC);
				}
				}
				this.state = 302;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			} while (_la===22);
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

	public static readonly _serializedATN: number[] = [4,1,22,305,2,0,7,0,2,
	1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,
	10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,16,2,17,
	7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,2,22,7,22,2,23,7,23,2,24,7,
	24,2,25,7,25,2,26,7,26,2,27,7,27,2,28,7,28,2,29,7,29,2,30,7,30,2,31,7,31,
	2,32,7,32,2,33,7,33,2,34,7,34,2,35,7,35,2,36,7,36,1,0,1,0,1,0,3,0,78,8,
	0,1,1,1,1,1,2,1,2,1,2,3,2,85,8,2,1,2,1,2,3,2,89,8,2,1,2,1,2,1,2,3,2,94,
	8,2,1,2,1,2,1,2,1,2,1,2,1,2,3,2,102,8,2,1,3,5,3,105,8,3,10,3,12,3,108,9,
	3,1,4,1,4,1,4,3,4,113,8,4,1,4,1,4,1,4,1,4,1,4,3,4,120,8,4,1,4,1,4,1,4,1,
	4,3,4,126,8,4,1,4,4,4,129,8,4,11,4,12,4,130,3,4,133,8,4,1,5,1,5,1,5,3,5,
	138,8,5,1,5,1,5,1,5,1,5,1,5,1,5,1,5,3,5,147,8,5,1,5,1,5,3,5,151,8,5,1,5,
	1,5,3,5,155,8,5,1,6,1,6,3,6,159,8,6,1,6,3,6,162,8,6,1,7,1,7,1,7,3,7,167,
	8,7,1,8,1,8,1,8,1,8,1,8,4,8,174,8,8,11,8,12,8,175,1,9,1,9,1,9,1,9,1,9,4,
	9,183,8,9,11,9,12,9,184,1,10,1,10,1,10,1,10,1,10,4,10,192,8,10,11,10,12,
	10,193,1,11,1,11,1,11,3,11,199,8,11,1,11,1,11,3,11,203,8,11,1,11,5,11,206,
	8,11,10,11,12,11,209,9,11,1,11,3,11,212,8,11,1,11,1,11,1,12,1,12,3,12,218,
	8,12,1,13,1,13,1,14,1,14,1,14,1,14,1,14,1,14,1,15,1,15,1,15,3,15,231,8,
	15,1,16,1,16,1,17,1,17,1,18,1,18,1,19,1,19,1,20,1,20,1,20,3,20,244,8,20,
	1,21,1,21,1,22,1,22,1,23,1,23,1,23,1,24,1,24,1,24,1,24,1,25,1,25,1,26,1,
	26,1,27,1,27,1,28,1,28,1,29,1,29,1,30,1,30,1,31,3,31,270,8,31,1,31,1,31,
	5,31,274,8,31,10,31,12,31,277,9,31,1,32,1,32,1,32,1,32,5,32,283,8,32,10,
	32,12,32,286,9,32,1,33,4,33,289,8,33,11,33,12,33,290,1,34,4,34,294,8,34,
	11,34,12,34,295,1,35,1,35,1,36,4,36,301,8,36,11,36,12,36,302,1,36,0,0,37,
	0,2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,32,34,36,38,40,42,44,46,48,50,
	52,54,56,58,60,62,64,66,68,70,72,0,2,1,0,1,2,1,0,18,19,304,0,74,1,0,0,0,
	2,79,1,0,0,0,4,84,1,0,0,0,6,106,1,0,0,0,8,112,1,0,0,0,10,137,1,0,0,0,12,
	158,1,0,0,0,14,166,1,0,0,0,16,173,1,0,0,0,18,182,1,0,0,0,20,191,1,0,0,0,
	22,195,1,0,0,0,24,217,1,0,0,0,26,219,1,0,0,0,28,221,1,0,0,0,30,230,1,0,
	0,0,32,232,1,0,0,0,34,234,1,0,0,0,36,236,1,0,0,0,38,238,1,0,0,0,40,243,
	1,0,0,0,42,245,1,0,0,0,44,247,1,0,0,0,46,249,1,0,0,0,48,252,1,0,0,0,50,
	256,1,0,0,0,52,258,1,0,0,0,54,260,1,0,0,0,56,262,1,0,0,0,58,264,1,0,0,0,
	60,266,1,0,0,0,62,269,1,0,0,0,64,278,1,0,0,0,66,288,1,0,0,0,68,293,1,0,
	0,0,70,297,1,0,0,0,72,300,1,0,0,0,74,75,3,4,2,0,75,77,3,6,3,0,76,78,3,68,
	34,0,77,76,1,0,0,0,77,78,1,0,0,0,78,1,1,0,0,0,79,80,7,0,0,0,80,3,1,0,0,
	0,81,82,3,64,32,0,82,83,3,68,34,0,83,85,1,0,0,0,84,81,1,0,0,0,84,85,1,0,
	0,0,85,86,1,0,0,0,86,88,5,3,0,0,87,89,3,66,33,0,88,87,1,0,0,0,88,89,1,0,
	0,0,89,93,1,0,0,0,90,91,3,68,34,0,91,92,3,64,32,0,92,94,1,0,0,0,93,90,1,
	0,0,0,93,94,1,0,0,0,94,95,1,0,0,0,95,96,3,68,34,0,96,97,3,2,1,0,97,98,5,
	4,0,0,98,99,3,66,33,0,99,101,3,70,35,0,100,102,3,66,33,0,101,100,1,0,0,
	0,101,102,1,0,0,0,102,5,1,0,0,0,103,105,3,8,4,0,104,103,1,0,0,0,105,108,
	1,0,0,0,106,104,1,0,0,0,106,107,1,0,0,0,107,7,1,0,0,0,108,106,1,0,0,0,109,
	110,3,68,34,0,110,111,3,64,32,0,111,113,1,0,0,0,112,109,1,0,0,0,112,113,
	1,0,0,0,113,114,1,0,0,0,114,115,3,68,34,0,115,116,5,5,0,0,116,117,3,66,
	33,0,117,119,3,60,30,0,118,120,3,66,33,0,119,118,1,0,0,0,119,120,1,0,0,
	0,120,132,1,0,0,0,121,122,3,68,34,0,122,123,3,2,1,0,123,125,5,6,0,0,124,
	126,3,66,33,0,125,124,1,0,0,0,125,126,1,0,0,0,126,128,1,0,0,0,127,129,3,
	10,5,0,128,127,1,0,0,0,129,130,1,0,0,0,130,128,1,0,0,0,130,131,1,0,0,0,
	131,133,1,0,0,0,132,121,1,0,0,0,132,133,1,0,0,0,133,9,1,0,0,0,134,135,3,
	68,34,0,135,136,3,64,32,0,136,138,1,0,0,0,137,134,1,0,0,0,137,138,1,0,0,
	0,138,139,1,0,0,0,139,140,3,68,34,0,140,141,3,2,1,0,141,142,3,2,1,0,142,
	143,5,7,0,0,143,144,3,66,33,0,144,146,3,58,29,0,145,147,3,66,33,0,146,145,
	1,0,0,0,146,147,1,0,0,0,147,148,1,0,0,0,148,150,5,8,0,0,149,151,3,66,33,
	0,150,149,1,0,0,0,150,151,1,0,0,0,151,152,1,0,0,0,152,154,3,12,6,0,153,
	155,3,66,33,0,154,153,1,0,0,0,154,155,1,0,0,0,155,11,1,0,0,0,156,159,3,
	22,11,0,157,159,3,50,25,0,158,156,1,0,0,0,158,157,1,0,0,0,159,161,1,0,0,
	0,160,162,3,14,7,0,161,160,1,0,0,0,161,162,1,0,0,0,162,13,1,0,0,0,163,167,
	3,16,8,0,164,167,3,18,9,0,165,167,3,20,10,0,166,163,1,0,0,0,166,164,1,0,
	0,0,166,165,1,0,0,0,167,15,1,0,0,0,168,169,3,66,33,0,169,170,3,34,17,0,
	170,171,3,66,33,0,171,172,3,50,25,0,172,174,1,0,0,0,173,168,1,0,0,0,174,
	175,1,0,0,0,175,173,1,0,0,0,175,176,1,0,0,0,176,17,1,0,0,0,177,178,3,66,
	33,0,178,179,3,32,16,0,179,180,3,66,33,0,180,181,3,50,25,0,181,183,1,0,
	0,0,182,177,1,0,0,0,183,184,1,0,0,0,184,182,1,0,0,0,184,185,1,0,0,0,185,
	19,1,0,0,0,186,187,3,66,33,0,187,188,3,36,18,0,188,189,3,66,33,0,189,190,
	3,50,25,0,190,192,1,0,0,0,191,186,1,0,0,0,192,193,1,0,0,0,193,191,1,0,0,
	0,193,194,1,0,0,0,194,21,1,0,0,0,195,196,5,9,0,0,196,198,3,40,20,0,197,
	199,3,66,33,0,198,197,1,0,0,0,198,199,1,0,0,0,199,207,1,0,0,0,200,202,5,
	10,0,0,201,203,3,66,33,0,202,201,1,0,0,0,202,203,1,0,0,0,203,204,1,0,0,
	0,204,206,3,40,20,0,205,200,1,0,0,0,206,209,1,0,0,0,207,205,1,0,0,0,207,
	208,1,0,0,0,208,211,1,0,0,0,209,207,1,0,0,0,210,212,3,66,33,0,211,210,1,
	0,0,0,211,212,1,0,0,0,212,213,1,0,0,0,213,214,5,11,0,0,214,23,1,0,0,0,215,
	218,3,26,13,0,216,218,3,28,14,0,217,215,1,0,0,0,217,216,1,0,0,0,218,25,
	1,0,0,0,219,220,3,52,26,0,220,27,1,0,0,0,221,222,3,54,27,0,222,223,3,66,
	33,0,223,224,3,38,19,0,224,225,3,66,33,0,225,226,3,56,28,0,226,29,1,0,0,
	0,227,231,3,34,17,0,228,231,3,32,16,0,229,231,3,36,18,0,230,227,1,0,0,0,
	230,228,1,0,0,0,230,229,1,0,0,0,231,31,1,0,0,0,232,233,5,12,0,0,233,33,
	1,0,0,0,234,235,5,13,0,0,235,35,1,0,0,0,236,237,5,14,0,0,237,37,1,0,0,0,
	238,239,5,15,0,0,239,39,1,0,0,0,240,244,3,42,21,0,241,244,3,46,23,0,242,
	244,3,48,24,0,243,240,1,0,0,0,243,241,1,0,0,0,243,242,1,0,0,0,244,41,1,
	0,0,0,245,246,3,72,36,0,246,43,1,0,0,0,247,248,3,72,36,0,248,45,1,0,0,0,
	249,250,3,42,21,0,250,251,5,16,0,0,251,47,1,0,0,0,252,253,3,42,21,0,253,
	254,5,17,0,0,254,255,3,44,22,0,255,49,1,0,0,0,256,257,3,24,12,0,257,51,
	1,0,0,0,258,259,3,72,36,0,259,53,1,0,0,0,260,261,3,72,36,0,261,55,1,0,0,
	0,262,263,3,72,36,0,263,57,1,0,0,0,264,265,3,72,36,0,265,59,1,0,0,0,266,
	267,3,72,36,0,267,61,1,0,0,0,268,270,3,66,33,0,269,268,1,0,0,0,269,270,
	1,0,0,0,270,271,1,0,0,0,271,275,5,17,0,0,272,274,8,1,0,0,273,272,1,0,0,
	0,274,277,1,0,0,0,275,273,1,0,0,0,275,276,1,0,0,0,276,63,1,0,0,0,277,275,
	1,0,0,0,278,284,3,62,31,0,279,280,3,68,34,0,280,281,3,62,31,0,281,283,1,
	0,0,0,282,279,1,0,0,0,283,286,1,0,0,0,284,282,1,0,0,0,284,285,1,0,0,0,285,
	65,1,0,0,0,286,284,1,0,0,0,287,289,5,20,0,0,288,287,1,0,0,0,289,290,1,0,
	0,0,290,288,1,0,0,0,290,291,1,0,0,0,291,67,1,0,0,0,292,294,7,1,0,0,293,
	292,1,0,0,0,294,295,1,0,0,0,295,293,1,0,0,0,295,296,1,0,0,0,296,69,1,0,
	0,0,297,298,5,21,0,0,298,71,1,0,0,0,299,301,5,22,0,0,300,299,1,0,0,0,301,
	302,1,0,0,0,302,300,1,0,0,0,302,303,1,0,0,0,303,73,1,0,0,0,34,77,84,88,
	93,101,106,112,119,125,130,132,137,146,150,154,158,161,166,175,184,193,
	198,202,207,211,217,230,243,269,275,284,290,295,302];

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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterMain) {
	 		listener.enterMain(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_indentation;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterIndentation) {
	 		listener.enterIndentation(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public newline_list(): NewlineContext[] {
		return this.getTypedRuleContexts(NewlineContext) as NewlineContext[];
	}
	public newline(i: number): NewlineContext {
		return this.getTypedRuleContext(NewlineContext, i) as NewlineContext;
	}
	public indentation(): IndentationContext {
		return this.getTypedRuleContext(IndentationContext, 0) as IndentationContext;
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
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_modelHeader;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterModelHeader) {
	 		listener.enterModelHeader(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterTypeDefs) {
	 		listener.enterTypeDefs(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public relationDeclaration_list(): RelationDeclarationContext[] {
		return this.getTypedRuleContexts(RelationDeclarationContext) as RelationDeclarationContext[];
	}
	public relationDeclaration(i: number): RelationDeclarationContext {
		return this.getTypedRuleContext(RelationDeclarationContext, i) as RelationDeclarationContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_typeDef;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterTypeDef) {
	 		listener.enterTypeDef(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public newline_list(): NewlineContext[] {
		return this.getTypedRuleContexts(NewlineContext) as NewlineContext[];
	}
	public newline(i: number): NewlineContext {
		return this.getTypedRuleContext(NewlineContext, i) as NewlineContext;
	}
	public indentation_list(): IndentationContext[] {
		return this.getTypedRuleContexts(IndentationContext) as IndentationContext[];
	}
	public indentation(i: number): IndentationContext {
		return this.getTypedRuleContext(IndentationContext, i) as IndentationContext;
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
	public relationDef(): RelationDefContext {
		return this.getTypedRuleContext(RelationDefContext, 0) as RelationDefContext;
	}
	public multiLineComment(): MultiLineCommentContext {
		return this.getTypedRuleContext(MultiLineCommentContext, 0) as MultiLineCommentContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDeclaration;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDeclaration) {
	 		listener.enterRelationDeclaration(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDef) {
	 		listener.enterRelationDef(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefPartials) {
	 		listener.enterRelationDefPartials(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefPartialAllOr) {
	 		listener.enterRelationDefPartialAllOr(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefPartialAllAnd) {
	 		listener.enterRelationDefPartialAllAnd(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefPartialAllButNot) {
	 		listener.enterRelationDefPartialAllButNot(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public relationDefTypeRestriction_list(): RelationDefTypeRestrictionContext[] {
		return this.getTypedRuleContexts(RelationDefTypeRestrictionContext) as RelationDefTypeRestrictionContext[];
	}
	public relationDefTypeRestriction(i: number): RelationDefTypeRestrictionContext {
		return this.getTypedRuleContext(RelationDefTypeRestrictionContext, i) as RelationDefTypeRestrictionContext;
	}
	public spacing_list(): SpacingContext[] {
		return this.getTypedRuleContexts(SpacingContext) as SpacingContext[];
	}
	public spacing(i: number): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, i) as SpacingContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefDirectAssignment;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefDirectAssignment) {
	 		listener.enterRelationDefDirectAssignment(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefRewrite) {
	 		listener.enterRelationDefRewrite(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefRelationOnSameObject) {
	 		listener.enterRelationDefRelationOnSameObject(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefRelationOnRelatedObject) {
	 		listener.enterRelationDefRelationOnRelatedObject(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefOperator) {
	 		listener.enterRelationDefOperator(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefOperatorAnd;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefOperatorAnd) {
	 		listener.enterRelationDefOperatorAnd(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefOperatorOr;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefOperatorOr) {
	 		listener.enterRelationDefOperatorOr(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefOperatorButNot;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefOperatorButNot) {
	 		listener.enterRelationDefOperatorButNot(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefKeywordFrom;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefKeywordFrom) {
	 		listener.enterRelationDefKeywordFrom(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefTypeRestriction) {
	 		listener.enterRelationDefTypeRestriction(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefTypeRestrictionType) {
	 		listener.enterRelationDefTypeRestrictionType(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefTypeRestrictionRelation) {
	 		listener.enterRelationDefTypeRestrictionRelation(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefTypeRestrictionWildcard;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefTypeRestrictionWildcard) {
	 		listener.enterRelationDefTypeRestrictionWildcard(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public relationDefTypeRestrictionRelation(): RelationDefTypeRestrictionRelationContext {
		return this.getTypedRuleContext(RelationDefTypeRestrictionRelationContext, 0) as RelationDefTypeRestrictionRelationContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefTypeRestrictionUserset;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefTypeRestrictionUserset) {
	 		listener.enterRelationDefTypeRestrictionUserset(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefGrouping) {
	 		listener.enterRelationDefGrouping(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRewriteComputedusersetName) {
	 		listener.enterRewriteComputedusersetName(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRewriteTuplesetComputedusersetName) {
	 		listener.enterRewriteTuplesetComputedusersetName(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRewriteTuplesetName) {
	 		listener.enterRewriteTuplesetName(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationName) {
	 		listener.enterRelationName(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterTypeName) {
	 		listener.enterTypeName(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public spacing(): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, 0) as SpacingContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_comment;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterComment) {
	 		listener.enterComment(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterMultiLineComment) {
	 		listener.enterMultiLineComment(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_spacing;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterSpacing) {
	 		listener.enterSpacing(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_newline;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterNewline) {
	 		listener.enterNewline(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_schemaVersion;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterSchemaVersion) {
	 		listener.enterSchemaVersion(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
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
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterName) {
	 		listener.enterName(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
	    if(listener.exitName) {
	 		listener.exitName(this);
		}
	}
}
