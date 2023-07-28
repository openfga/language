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
	public static readonly T__21 = 22;
	public static readonly T__22 = 23;
	public static readonly WORD = 24;
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
	public static readonly RULE_relationDefGroup = 26;
	public static readonly RULE_rewriteComputedusersetName = 27;
	public static readonly RULE_rewriteTuplesetComputedusersetName = 28;
	public static readonly RULE_rewriteTuplesetName = 29;
	public static readonly RULE_relationName = 30;
	public static readonly RULE_typeName = 31;
	public static readonly RULE_comment = 32;
	public static readonly RULE_multiLineComment = 33;
	public static readonly RULE_spacing = 34;
	public static readonly RULE_newline = 35;
	public static readonly RULE_schemaVersion = 36;
	public static readonly RULE_name = 37;
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
		"relationDefGrouping", "relationDefGroup", "rewriteComputedusersetName", 
		"rewriteTuplesetComputedusersetName", "rewriteTuplesetName", "relationName", 
		"typeName", "comment", "multiLineComment", "spacing", "newline", "schemaVersion", 
		"name",
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
			this.state = 76;
			this.modelHeader();
			this.state = 77;
			this.typeDefs();
			this.state = 79;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===21) {
				{
				this.state = 78;
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
			this.state = 81;
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
			this.state = 86;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===17 || _la===22) {
				{
				this.state = 83;
				this.multiLineComment();
				this.state = 84;
				this.newline();
				}
			}

			this.state = 88;
			this.match(OpenFGAParser.T__2);
			this.state = 96;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 3, this._ctx) ) {
			case 1:
				{
				this.state = 90;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				do {
					{
					{
					this.state = 89;
					this.newline();
					}
					}
					this.state = 92;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
				} while (_la===21);
				this.state = 94;
				this.multiLineComment();
				}
				break;
			}
			this.state = 98;
			this.newline();
			this.state = 99;
			this.indentation();
			this.state = 100;
			this.match(OpenFGAParser.T__3);
			this.state = 101;
			this.spacing();
			this.state = 102;
			this.schemaVersion();
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
			this.state = 107;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 4, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 104;
					this.typeDef();
					}
					}
				}
				this.state = 109;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 4, this._ctx);
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
			this.state = 113;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 5, this._ctx) ) {
			case 1:
				{
				this.state = 110;
				this.newline();
				this.state = 111;
				this.multiLineComment();
				}
				break;
			}
			this.state = 116;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			do {
				{
				{
				this.state = 115;
				this.newline();
				}
				}
				this.state = 118;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			} while (_la===21);
			this.state = 120;
			this.match(OpenFGAParser.T__4);
			this.state = 121;
			this.spacing();
			this.state = 122;
			this.typeName();
			this.state = 131;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 8, this._ctx) ) {
			case 1:
				{
				this.state = 123;
				this.newline();
				this.state = 124;
				this.indentation();
				this.state = 125;
				this.match(OpenFGAParser.T__5);
				this.state = 127;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 126;
						this.relationDeclaration();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					this.state = 129;
					this._errHandler.sync(this);
					_alt = this._interp.adaptivePredict(this._input, 7, this._ctx);
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
			this.state = 136;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 9, this._ctx) ) {
			case 1:
				{
				this.state = 133;
				this.newline();
				this.state = 134;
				this.multiLineComment();
				}
				break;
			}
			this.state = 138;
			this.newline();
			this.state = 139;
			this.indentation();
			this.state = 140;
			this.indentation();
			this.state = 141;
			this.match(OpenFGAParser.T__6);
			this.state = 142;
			this.spacing();
			this.state = 143;
			this.relationName();
			this.state = 145;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===22) {
				{
				this.state = 144;
				this.spacing();
				}
			}

			this.state = 147;
			this.match(OpenFGAParser.T__7);
			this.state = 149;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===22) {
				{
				this.state = 148;
				this.spacing();
				}
			}

			this.state = 151;
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
	public relationDef(): RelationDefContext {
		let localctx: RelationDefContext = new RelationDefContext(this, this._ctx, this.state);
		this.enterRule(localctx, 12, OpenFGAParser.RULE_relationDef);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 155;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 9:
				{
				this.state = 153;
				this.relationDefDirectAssignment();
				}
				break;
			case 24:
				{
				this.state = 154;
				this.relationDefGrouping();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this.state = 158;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===22) {
				{
				this.state = 157;
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
		try {
			this.state = 163;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 14, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 160;
				this.relationDefPartialAllOr();
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 161;
				this.relationDefPartialAllAnd();
				}
				break;
			case 3:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 162;
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
			this.state = 170;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 165;
					this.spacing();
					this.state = 166;
					this.relationDefOperatorOr();
					this.state = 167;
					this.spacing();
					this.state = 168;
					this.relationDefGrouping();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 172;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 15, this._ctx);
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
			this.state = 179;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 174;
					this.spacing();
					this.state = 175;
					this.relationDefOperatorAnd();
					this.state = 176;
					this.spacing();
					this.state = 177;
					this.relationDefGrouping();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 181;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 16, this._ctx);
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
			this.state = 188;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 183;
					this.spacing();
					this.state = 184;
					this.relationDefOperatorButNot();
					this.state = 185;
					this.spacing();
					this.state = 186;
					this.relationDefGrouping();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 190;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 17, this._ctx);
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
			this.state = 192;
			this.match(OpenFGAParser.T__8);
			this.state = 193;
			this.relationDefTypeRestriction();
			this.state = 195;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 18, this._ctx) ) {
			case 1:
				{
				this.state = 194;
				this.spacing();
				}
				break;
			}
			this.state = 204;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===10) {
				{
				{
				this.state = 197;
				this.match(OpenFGAParser.T__9);
				this.state = 199;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===22) {
					{
					this.state = 198;
					this.spacing();
					}
				}

				this.state = 201;
				this.relationDefTypeRestriction();
				}
				}
				this.state = 206;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 208;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===22) {
				{
				this.state = 207;
				this.spacing();
				}
			}

			this.state = 210;
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
			this.state = 214;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 22, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 212;
				this.relationDefRelationOnSameObject();
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 213;
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
			this.state = 216;
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
			this.state = 218;
			this.rewriteTuplesetComputedusersetName();
			this.state = 219;
			this.spacing();
			this.state = 220;
			this.relationDefKeywordFrom();
			this.state = 221;
			this.spacing();
			this.state = 222;
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
			this.state = 227;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 13:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 224;
				this.relationDefOperatorOr();
				}
				break;
			case 12:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 225;
				this.relationDefOperatorAnd();
				}
				break;
			case 14:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 226;
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
			this.state = 229;
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
			this.state = 231;
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
			this.state = 233;
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
			this.state = 235;
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
			this.state = 240;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 24, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 237;
				this.relationDefTypeRestrictionType();
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 238;
				this.relationDefTypeRestrictionWildcard();
				}
				break;
			case 3:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 239;
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
			this.state = 242;
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
	public relationDefTypeRestrictionWildcard(): RelationDefTypeRestrictionWildcardContext {
		let localctx: RelationDefTypeRestrictionWildcardContext = new RelationDefTypeRestrictionWildcardContext(this, this._ctx, this.state);
		this.enterRule(localctx, 46, OpenFGAParser.RULE_relationDefTypeRestrictionWildcard);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 246;
			this.relationDefTypeRestrictionType();
			this.state = 247;
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
			this.state = 249;
			this.relationDefTypeRestrictionType();
			this.state = 250;
			this.match(OpenFGAParser.T__16);
			this.state = 251;
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
			this.state = 253;
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
	public relationDefGroup(): RelationDefGroupContext {
		let localctx: RelationDefGroupContext = new RelationDefGroupContext(this, this._ctx, this.state);
		this.enterRule(localctx, 52, OpenFGAParser.RULE_relationDefGroup);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 255;
			this.match(OpenFGAParser.T__17);
			this.state = 256;
			this.relationDefGrouping();
			this.state = 260;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===22) {
				{
				{
				this.state = 257;
				this.relationDefPartials();
				}
				}
				this.state = 262;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 263;
			this.match(OpenFGAParser.T__18);
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
		this.enterRule(localctx, 54, OpenFGAParser.RULE_rewriteComputedusersetName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 265;
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
		this.enterRule(localctx, 56, OpenFGAParser.RULE_rewriteTuplesetComputedusersetName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 267;
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
		this.enterRule(localctx, 58, OpenFGAParser.RULE_rewriteTuplesetName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 269;
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
		this.enterRule(localctx, 60, OpenFGAParser.RULE_relationName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 271;
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
		this.enterRule(localctx, 62, OpenFGAParser.RULE_typeName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 273;
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
		this.enterRule(localctx, 64, OpenFGAParser.RULE_comment);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 278;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===22) {
				{
				{
				this.state = 275;
				this.spacing();
				}
				}
				this.state = 280;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 281;
			this.match(OpenFGAParser.T__16);
			this.state = 285;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while ((((_la) & ~0x1F) === 0 && ((1 << _la) & 30408702) !== 0)) {
				{
				{
				this.state = 282;
				_la = this._input.LA(1);
				if(_la<=0 || _la===20 || _la===21) {
				this._errHandler.recoverInline(this);
				}
				else {
					this._errHandler.reportMatch(this);
				    this.consume();
				}
				}
				}
				this.state = 287;
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
		this.enterRule(localctx, 66, OpenFGAParser.RULE_multiLineComment);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 288;
			this.comment();
			this.state = 294;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 28, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 289;
					this.newline();
					this.state = 290;
					this.comment();
					}
					}
				}
				this.state = 296;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 28, this._ctx);
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
		this.enterRule(localctx, 68, OpenFGAParser.RULE_spacing);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 298;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 297;
					this.match(OpenFGAParser.T__21);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 300;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 29, this._ctx);
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
		this.enterRule(localctx, 70, OpenFGAParser.RULE_newline);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 303;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 302;
					this.match(OpenFGAParser.T__20);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 305;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 30, this._ctx);
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
	public schemaVersion(): SchemaVersionContext {
		let localctx: SchemaVersionContext = new SchemaVersionContext(this, this._ctx, this.state);
		this.enterRule(localctx, 72, OpenFGAParser.RULE_schemaVersion);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 307;
			this.match(OpenFGAParser.T__22);
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
		this.enterRule(localctx, 74, OpenFGAParser.RULE_name);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 310;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			do {
				{
				{
				this.state = 309;
				this.match(OpenFGAParser.WORD);
				}
				}
				this.state = 312;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			} while (_la===24);
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

	public static readonly _serializedATN: number[] = [4,1,24,315,2,0,7,0,2,
	1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,
	10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,16,2,17,
	7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,2,22,7,22,2,23,7,23,2,24,7,
	24,2,25,7,25,2,26,7,26,2,27,7,27,2,28,7,28,2,29,7,29,2,30,7,30,2,31,7,31,
	2,32,7,32,2,33,7,33,2,34,7,34,2,35,7,35,2,36,7,36,2,37,7,37,1,0,1,0,1,0,
	3,0,80,8,0,1,1,1,1,1,2,1,2,1,2,3,2,87,8,2,1,2,1,2,4,2,91,8,2,11,2,12,2,
	92,1,2,1,2,3,2,97,8,2,1,2,1,2,1,2,1,2,1,2,1,2,1,3,5,3,106,8,3,10,3,12,3,
	109,9,3,1,4,1,4,1,4,3,4,114,8,4,1,4,4,4,117,8,4,11,4,12,4,118,1,4,1,4,1,
	4,1,4,1,4,1,4,1,4,4,4,128,8,4,11,4,12,4,129,3,4,132,8,4,1,5,1,5,1,5,3,5,
	137,8,5,1,5,1,5,1,5,1,5,1,5,1,5,1,5,3,5,146,8,5,1,5,1,5,3,5,150,8,5,1,5,
	1,5,1,6,1,6,3,6,156,8,6,1,6,3,6,159,8,6,1,7,1,7,1,7,3,7,164,8,7,1,8,1,8,
	1,8,1,8,1,8,4,8,171,8,8,11,8,12,8,172,1,9,1,9,1,9,1,9,1,9,4,9,180,8,9,11,
	9,12,9,181,1,10,1,10,1,10,1,10,1,10,4,10,189,8,10,11,10,12,10,190,1,11,
	1,11,1,11,3,11,196,8,11,1,11,1,11,3,11,200,8,11,1,11,5,11,203,8,11,10,11,
	12,11,206,9,11,1,11,3,11,209,8,11,1,11,1,11,1,12,1,12,3,12,215,8,12,1,13,
	1,13,1,14,1,14,1,14,1,14,1,14,1,14,1,15,1,15,1,15,3,15,228,8,15,1,16,1,
	16,1,17,1,17,1,18,1,18,1,19,1,19,1,20,1,20,1,20,3,20,241,8,20,1,21,1,21,
	1,22,1,22,1,23,1,23,1,23,1,24,1,24,1,24,1,24,1,25,1,25,1,26,1,26,1,26,5,
	26,259,8,26,10,26,12,26,262,9,26,1,26,1,26,1,27,1,27,1,28,1,28,1,29,1,29,
	1,30,1,30,1,31,1,31,1,32,5,32,277,8,32,10,32,12,32,280,9,32,1,32,1,32,5,
	32,284,8,32,10,32,12,32,287,9,32,1,33,1,33,1,33,1,33,5,33,293,8,33,10,33,
	12,33,296,9,33,1,34,4,34,299,8,34,11,34,12,34,300,1,35,4,35,304,8,35,11,
	35,12,35,305,1,36,1,36,1,37,4,37,311,8,37,11,37,12,37,312,1,37,0,0,38,0,
	2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,32,34,36,38,40,42,44,46,48,50,
	52,54,56,58,60,62,64,66,68,70,72,74,0,2,1,0,1,2,1,0,20,21,311,0,76,1,0,
	0,0,2,81,1,0,0,0,4,86,1,0,0,0,6,107,1,0,0,0,8,113,1,0,0,0,10,136,1,0,0,
	0,12,155,1,0,0,0,14,163,1,0,0,0,16,170,1,0,0,0,18,179,1,0,0,0,20,188,1,
	0,0,0,22,192,1,0,0,0,24,214,1,0,0,0,26,216,1,0,0,0,28,218,1,0,0,0,30,227,
	1,0,0,0,32,229,1,0,0,0,34,231,1,0,0,0,36,233,1,0,0,0,38,235,1,0,0,0,40,
	240,1,0,0,0,42,242,1,0,0,0,44,244,1,0,0,0,46,246,1,0,0,0,48,249,1,0,0,0,
	50,253,1,0,0,0,52,255,1,0,0,0,54,265,1,0,0,0,56,267,1,0,0,0,58,269,1,0,
	0,0,60,271,1,0,0,0,62,273,1,0,0,0,64,278,1,0,0,0,66,288,1,0,0,0,68,298,
	1,0,0,0,70,303,1,0,0,0,72,307,1,0,0,0,74,310,1,0,0,0,76,77,3,4,2,0,77,79,
	3,6,3,0,78,80,3,70,35,0,79,78,1,0,0,0,79,80,1,0,0,0,80,1,1,0,0,0,81,82,
	7,0,0,0,82,3,1,0,0,0,83,84,3,66,33,0,84,85,3,70,35,0,85,87,1,0,0,0,86,83,
	1,0,0,0,86,87,1,0,0,0,87,88,1,0,0,0,88,96,5,3,0,0,89,91,3,70,35,0,90,89,
	1,0,0,0,91,92,1,0,0,0,92,90,1,0,0,0,92,93,1,0,0,0,93,94,1,0,0,0,94,95,3,
	66,33,0,95,97,1,0,0,0,96,90,1,0,0,0,96,97,1,0,0,0,97,98,1,0,0,0,98,99,3,
	70,35,0,99,100,3,2,1,0,100,101,5,4,0,0,101,102,3,68,34,0,102,103,3,72,36,
	0,103,5,1,0,0,0,104,106,3,8,4,0,105,104,1,0,0,0,106,109,1,0,0,0,107,105,
	1,0,0,0,107,108,1,0,0,0,108,7,1,0,0,0,109,107,1,0,0,0,110,111,3,70,35,0,
	111,112,3,66,33,0,112,114,1,0,0,0,113,110,1,0,0,0,113,114,1,0,0,0,114,116,
	1,0,0,0,115,117,3,70,35,0,116,115,1,0,0,0,117,118,1,0,0,0,118,116,1,0,0,
	0,118,119,1,0,0,0,119,120,1,0,0,0,120,121,5,5,0,0,121,122,3,68,34,0,122,
	131,3,62,31,0,123,124,3,70,35,0,124,125,3,2,1,0,125,127,5,6,0,0,126,128,
	3,10,5,0,127,126,1,0,0,0,128,129,1,0,0,0,129,127,1,0,0,0,129,130,1,0,0,
	0,130,132,1,0,0,0,131,123,1,0,0,0,131,132,1,0,0,0,132,9,1,0,0,0,133,134,
	3,70,35,0,134,135,3,66,33,0,135,137,1,0,0,0,136,133,1,0,0,0,136,137,1,0,
	0,0,137,138,1,0,0,0,138,139,3,70,35,0,139,140,3,2,1,0,140,141,3,2,1,0,141,
	142,5,7,0,0,142,143,3,68,34,0,143,145,3,60,30,0,144,146,3,68,34,0,145,144,
	1,0,0,0,145,146,1,0,0,0,146,147,1,0,0,0,147,149,5,8,0,0,148,150,3,68,34,
	0,149,148,1,0,0,0,149,150,1,0,0,0,150,151,1,0,0,0,151,152,3,12,6,0,152,
	11,1,0,0,0,153,156,3,22,11,0,154,156,3,50,25,0,155,153,1,0,0,0,155,154,
	1,0,0,0,156,158,1,0,0,0,157,159,3,14,7,0,158,157,1,0,0,0,158,159,1,0,0,
	0,159,13,1,0,0,0,160,164,3,16,8,0,161,164,3,18,9,0,162,164,3,20,10,0,163,
	160,1,0,0,0,163,161,1,0,0,0,163,162,1,0,0,0,164,15,1,0,0,0,165,166,3,68,
	34,0,166,167,3,34,17,0,167,168,3,68,34,0,168,169,3,50,25,0,169,171,1,0,
	0,0,170,165,1,0,0,0,171,172,1,0,0,0,172,170,1,0,0,0,172,173,1,0,0,0,173,
	17,1,0,0,0,174,175,3,68,34,0,175,176,3,32,16,0,176,177,3,68,34,0,177,178,
	3,50,25,0,178,180,1,0,0,0,179,174,1,0,0,0,180,181,1,0,0,0,181,179,1,0,0,
	0,181,182,1,0,0,0,182,19,1,0,0,0,183,184,3,68,34,0,184,185,3,36,18,0,185,
	186,3,68,34,0,186,187,3,50,25,0,187,189,1,0,0,0,188,183,1,0,0,0,189,190,
	1,0,0,0,190,188,1,0,0,0,190,191,1,0,0,0,191,21,1,0,0,0,192,193,5,9,0,0,
	193,195,3,40,20,0,194,196,3,68,34,0,195,194,1,0,0,0,195,196,1,0,0,0,196,
	204,1,0,0,0,197,199,5,10,0,0,198,200,3,68,34,0,199,198,1,0,0,0,199,200,
	1,0,0,0,200,201,1,0,0,0,201,203,3,40,20,0,202,197,1,0,0,0,203,206,1,0,0,
	0,204,202,1,0,0,0,204,205,1,0,0,0,205,208,1,0,0,0,206,204,1,0,0,0,207,209,
	3,68,34,0,208,207,1,0,0,0,208,209,1,0,0,0,209,210,1,0,0,0,210,211,5,11,
	0,0,211,23,1,0,0,0,212,215,3,26,13,0,213,215,3,28,14,0,214,212,1,0,0,0,
	214,213,1,0,0,0,215,25,1,0,0,0,216,217,3,54,27,0,217,27,1,0,0,0,218,219,
	3,56,28,0,219,220,3,68,34,0,220,221,3,38,19,0,221,222,3,68,34,0,222,223,
	3,58,29,0,223,29,1,0,0,0,224,228,3,34,17,0,225,228,3,32,16,0,226,228,3,
	36,18,0,227,224,1,0,0,0,227,225,1,0,0,0,227,226,1,0,0,0,228,31,1,0,0,0,
	229,230,5,12,0,0,230,33,1,0,0,0,231,232,5,13,0,0,232,35,1,0,0,0,233,234,
	5,14,0,0,234,37,1,0,0,0,235,236,5,15,0,0,236,39,1,0,0,0,237,241,3,42,21,
	0,238,241,3,46,23,0,239,241,3,48,24,0,240,237,1,0,0,0,240,238,1,0,0,0,240,
	239,1,0,0,0,241,41,1,0,0,0,242,243,3,74,37,0,243,43,1,0,0,0,244,245,3,74,
	37,0,245,45,1,0,0,0,246,247,3,42,21,0,247,248,5,16,0,0,248,47,1,0,0,0,249,
	250,3,42,21,0,250,251,5,17,0,0,251,252,3,44,22,0,252,49,1,0,0,0,253,254,
	3,24,12,0,254,51,1,0,0,0,255,256,5,18,0,0,256,260,3,50,25,0,257,259,3,14,
	7,0,258,257,1,0,0,0,259,262,1,0,0,0,260,258,1,0,0,0,260,261,1,0,0,0,261,
	263,1,0,0,0,262,260,1,0,0,0,263,264,5,19,0,0,264,53,1,0,0,0,265,266,3,74,
	37,0,266,55,1,0,0,0,267,268,3,74,37,0,268,57,1,0,0,0,269,270,3,74,37,0,
	270,59,1,0,0,0,271,272,3,74,37,0,272,61,1,0,0,0,273,274,3,74,37,0,274,63,
	1,0,0,0,275,277,3,68,34,0,276,275,1,0,0,0,277,280,1,0,0,0,278,276,1,0,0,
	0,278,279,1,0,0,0,279,281,1,0,0,0,280,278,1,0,0,0,281,285,5,17,0,0,282,
	284,8,1,0,0,283,282,1,0,0,0,284,287,1,0,0,0,285,283,1,0,0,0,285,286,1,0,
	0,0,286,65,1,0,0,0,287,285,1,0,0,0,288,294,3,64,32,0,289,290,3,70,35,0,
	290,291,3,64,32,0,291,293,1,0,0,0,292,289,1,0,0,0,293,296,1,0,0,0,294,292,
	1,0,0,0,294,295,1,0,0,0,295,67,1,0,0,0,296,294,1,0,0,0,297,299,5,22,0,0,
	298,297,1,0,0,0,299,300,1,0,0,0,300,298,1,0,0,0,300,301,1,0,0,0,301,69,
	1,0,0,0,302,304,5,21,0,0,303,302,1,0,0,0,304,305,1,0,0,0,305,303,1,0,0,
	0,305,306,1,0,0,0,306,71,1,0,0,0,307,308,5,23,0,0,308,73,1,0,0,0,309,311,
	5,24,0,0,310,309,1,0,0,0,311,312,1,0,0,0,312,310,1,0,0,0,312,313,1,0,0,
	0,313,75,1,0,0,0,32,79,86,92,96,107,113,118,129,131,136,145,149,155,158,
	163,172,181,190,195,199,204,208,214,227,240,260,278,285,294,300,305,312];

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
	public spacing(): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, 0) as SpacingContext;
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
	public spacing(): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, 0) as SpacingContext;
	}
	public typeName(): TypeNameContext {
		return this.getTypedRuleContext(TypeNameContext, 0) as TypeNameContext;
	}
	public newline_list(): NewlineContext[] {
		return this.getTypedRuleContexts(NewlineContext) as NewlineContext[];
	}
	public newline(i: number): NewlineContext {
		return this.getTypedRuleContext(NewlineContext, i) as NewlineContext;
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


export class RelationDefGroupContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public relationDefGrouping(): RelationDefGroupingContext {
		return this.getTypedRuleContext(RelationDefGroupingContext, 0) as RelationDefGroupingContext;
	}
	public relationDefPartials_list(): RelationDefPartialsContext[] {
		return this.getTypedRuleContexts(RelationDefPartialsContext) as RelationDefPartialsContext[];
	}
	public relationDefPartials(i: number): RelationDefPartialsContext {
		return this.getTypedRuleContext(RelationDefPartialsContext, i) as RelationDefPartialsContext;
	}
    public get ruleIndex(): number {
    	return OpenFGAParser.RULE_relationDefGroup;
	}
	public enterRule(listener: OpenFGAListener): void {
	    if(listener.enterRelationDefGroup) {
	 		listener.enterRelationDefGroup(this);
		}
	}
	public exitRule(listener: OpenFGAListener): void {
	    if(listener.exitRelationDefGroup) {
	 		listener.exitRelationDefGroup(this);
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
	public spacing_list(): SpacingContext[] {
		return this.getTypedRuleContexts(SpacingContext) as SpacingContext[];
	}
	public spacing(i: number): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, i) as SpacingContext;
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
	public WORD_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.WORD);
	}
	public WORD(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.WORD, i);
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
