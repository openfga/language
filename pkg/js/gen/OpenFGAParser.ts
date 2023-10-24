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
	public static readonly CONDITION = 6;
	public static readonly RELATIONS = 7;
	public static readonly DEFINE = 8;
	public static readonly WTH = 9;
	public static readonly HASH = 10;
	public static readonly COLON = 11;
	public static readonly WILDCARD = 12;
	public static readonly L_SQUARE = 13;
	public static readonly R_SQUARE = 14;
	public static readonly L_PARANTHESES = 15;
	public static readonly R_PARANTHESES = 16;
	public static readonly L_BRACES = 17;
	public static readonly R_BRACES = 18;
	public static readonly L_ANGLE_BRACKET = 19;
	public static readonly R_ANGLE_BRACKET = 20;
	public static readonly COMMA = 21;
	public static readonly CONDITION_PARAM_CONTAINER = 22;
	public static readonly CONDITION_PARAM_TYPE = 23;
	public static readonly CONDITION_SYMBOL = 24;
	public static readonly AND = 25;
	public static readonly OR = 26;
	public static readonly BUT_NOT = 27;
	public static readonly FROM = 28;
	public static readonly ALPHA_NUMERIC = 29;
	public static readonly NEWLINE = 30;
	public static readonly WS = 31;
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
	public static readonly RULE_relationDefTypeRestrictionWithCondition = 21;
	public static readonly RULE_relationDefTypeRestrictionType = 22;
	public static readonly RULE_relationDefTypeRestrictionRelation = 23;
	public static readonly RULE_relationDefTypeRestrictionWildcard = 24;
	public static readonly RULE_relationDefTypeRestrictionUserset = 25;
	public static readonly RULE_relationDefGrouping = 26;
	public static readonly RULE_rewriteComputedusersetName = 27;
	public static readonly RULE_rewriteTuplesetComputedusersetName = 28;
	public static readonly RULE_rewriteTuplesetName = 29;
	public static readonly RULE_relationName = 30;
	public static readonly RULE_typeName = 31;
	public static readonly RULE_conditions = 32;
	public static readonly RULE_condition = 33;
	public static readonly RULE_conditionParameter = 34;
	public static readonly RULE_parameterName = 35;
	public static readonly RULE_conditionName = 36;
	public static readonly RULE_parameterType = 37;
	public static readonly RULE_conditionExpression = 38;
	public static readonly RULE_comment = 39;
	public static readonly RULE_multiLineComment = 40;
	public static readonly RULE_spacing = 41;
	public static readonly RULE_newline = 42;
	public static readonly RULE_schemaVersion = 43;
	public static readonly RULE_name = 44;
	public static readonly literalNames: (string | null)[] = [ null, null, 
                                                            "'model'", "'schema'", 
                                                            "'1.1'", "'type'", 
                                                            "'condition'", 
                                                            "'relations'", 
                                                            "'define'", 
                                                            "'with'", "'#'", 
                                                            "':'", "'*'", 
                                                            "'['", "']'", 
                                                            "'('", "')'", 
                                                            "'{'", "'}'", 
                                                            "'<'", "'>'", 
                                                            "','", null, 
                                                            null, null, 
                                                            "'and'", "'or'", 
                                                            "'but not'", 
                                                            "'from'" ];
	public static readonly symbolicNames: (string | null)[] = [ null, "INDENT", 
                                                             "MODEL", "SCHEMA", 
                                                             "SCHEMA_VERSION", 
                                                             "TYPE", "CONDITION", 
                                                             "RELATIONS", 
                                                             "DEFINE", "WTH", 
                                                             "HASH", "COLON", 
                                                             "WILDCARD", 
                                                             "L_SQUARE", 
                                                             "R_SQUARE", 
                                                             "L_PARANTHESES", 
                                                             "R_PARANTHESES", 
                                                             "L_BRACES", 
                                                             "R_BRACES", 
                                                             "L_ANGLE_BRACKET", 
                                                             "R_ANGLE_BRACKET", 
                                                             "COMMA", "CONDITION_PARAM_CONTAINER", 
                                                             "CONDITION_PARAM_TYPE", 
                                                             "CONDITION_SYMBOL", 
                                                             "AND", "OR", 
                                                             "BUT_NOT", 
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
		"relationDefTypeRestrictionWithCondition", "relationDefTypeRestrictionType", 
		"relationDefTypeRestrictionRelation", "relationDefTypeRestrictionWildcard", 
		"relationDefTypeRestrictionUserset", "relationDefGrouping", "rewriteComputedusersetName", 
		"rewriteTuplesetComputedusersetName", "rewriteTuplesetName", "relationName", 
		"typeName", "conditions", "condition", "conditionParameter", "parameterName", 
		"conditionName", "parameterType", "conditionExpression", "comment", "multiLineComment", 
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
			this.state = 90;
			this.modelHeader();
			this.state = 91;
			this.typeDefs();
			this.state = 92;
			this.conditions();
			this.state = 94;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===30) {
				{
				this.state = 93;
				this.newline();
				}
			}

			this.state = 96;
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
	public indentation(): IndentationContext {
		let localctx: IndentationContext = new IndentationContext(this, this._ctx, this.state);
		this.enterRule(localctx, 2, OpenFGAParser.RULE_indentation);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 98;
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
			this.state = 103;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===10 || _la===31) {
				{
				this.state = 100;
				this.multiLineComment();
				this.state = 101;
				this.newline();
				}
			}

			this.state = 105;
			this.match(OpenFGAParser.MODEL);
			this.state = 107;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===31) {
				{
				this.state = 106;
				this.spacing();
				}
			}

			this.state = 112;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===30) {
				{
				this.state = 109;
				this.newline();
				this.state = 110;
				this.multiLineComment();
				}
			}

			this.state = 114;
			this.indentation();
			this.state = 115;
			this.match(OpenFGAParser.SCHEMA);
			this.state = 116;
			this.spacing();
			this.state = 117;
			this.schemaVersion();
			this.state = 119;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===31) {
				{
				this.state = 118;
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
			this.state = 124;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 5, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 121;
					this.typeDef();
					}
					}
				}
				this.state = 126;
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
			this.state = 130;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 6, this._ctx) ) {
			case 1:
				{
				this.state = 127;
				this.newline();
				this.state = 128;
				this.multiLineComment();
				}
				break;
			}
			this.state = 132;
			this.newline();
			this.state = 133;
			this.match(OpenFGAParser.TYPE);
			this.state = 134;
			this.spacing();
			this.state = 135;
			this.typeName();
			this.state = 137;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===31) {
				{
				this.state = 136;
				this.spacing();
				}
			}

			this.state = 149;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===1) {
				{
				this.state = 139;
				this.indentation();
				this.state = 140;
				this.match(OpenFGAParser.RELATIONS);
				this.state = 142;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===31) {
					{
					this.state = 141;
					this.spacing();
					}
				}

				this.state = 145;
				this._errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						this.state = 144;
						this.relationDeclaration();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					this.state = 147;
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
			this.state = 154;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===30) {
				{
				this.state = 151;
				this.newline();
				this.state = 152;
				this.multiLineComment();
				}
			}

			this.state = 156;
			this.indentation();
			this.state = 157;
			this.match(OpenFGAParser.DEFINE);
			this.state = 158;
			this.spacing();
			this.state = 159;
			this.relationName();
			this.state = 161;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===31) {
				{
				this.state = 160;
				this.spacing();
				}
			}

			this.state = 163;
			this.match(OpenFGAParser.COLON);
			this.state = 165;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===31) {
				{
				this.state = 164;
				this.spacing();
				}
			}

			this.state = 167;
			this.relationDef();
			this.state = 169;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===31) {
				{
				this.state = 168;
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
			this.state = 173;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 13:
				{
				this.state = 171;
				this.relationDefDirectAssignment();
				}
				break;
			case 29:
				{
				this.state = 172;
				this.relationDefGrouping();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this.state = 176;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 16, this._ctx) ) {
			case 1:
				{
				this.state = 175;
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
			this.state = 181;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 17, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 178;
				this.relationDefPartialAllOr();
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 179;
				this.relationDefPartialAllAnd();
				}
				break;
			case 3:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 180;
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
					this.relationDefOperatorOr();
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
			this.state = 197;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 192;
					this.spacing();
					this.state = 193;
					this.relationDefOperatorAnd();
					this.state = 194;
					this.spacing();
					this.state = 195;
					this.relationDefGrouping();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 199;
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
			this.state = 206;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 201;
					this.spacing();
					this.state = 202;
					this.relationDefOperatorButNot();
					this.state = 203;
					this.spacing();
					this.state = 204;
					this.relationDefGrouping();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 208;
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
			this.state = 210;
			this.match(OpenFGAParser.L_SQUARE);
			this.state = 212;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===31) {
				{
				this.state = 211;
				this.spacing();
				}
			}

			this.state = 214;
			this.relationDefTypeRestriction();
			this.state = 216;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 22, this._ctx) ) {
			case 1:
				{
				this.state = 215;
				this.spacing();
				}
				break;
			}
			this.state = 225;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===21) {
				{
				{
				this.state = 218;
				this.match(OpenFGAParser.COMMA);
				this.state = 220;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===31) {
					{
					this.state = 219;
					this.spacing();
					}
				}

				this.state = 222;
				this.relationDefTypeRestriction();
				}
				}
				this.state = 227;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 229;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===31) {
				{
				this.state = 228;
				this.spacing();
				}
			}

			this.state = 231;
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
			this.state = 235;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 26, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 233;
				this.relationDefRelationOnSameObject();
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 234;
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
			this.state = 237;
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
			this.state = 239;
			this.rewriteTuplesetComputedusersetName();
			this.state = 240;
			this.spacing();
			this.state = 241;
			this.relationDefKeywordFrom();
			this.state = 242;
			this.spacing();
			this.state = 243;
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
			this.state = 248;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 26:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 245;
				this.relationDefOperatorOr();
				}
				break;
			case 25:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 246;
				this.relationDefOperatorAnd();
				}
				break;
			case 27:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 247;
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
			this.state = 250;
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
			this.state = 252;
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
			this.state = 254;
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
			this.state = 256;
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
			this.state = 262;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 28, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 258;
				this.relationDefTypeRestrictionType();
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 259;
				this.relationDefTypeRestrictionWildcard();
				}
				break;
			case 3:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 260;
				this.relationDefTypeRestrictionUserset();
				}
				break;
			case 4:
				this.enterOuterAlt(localctx, 4);
				{
				this.state = 261;
				this.relationDefTypeRestrictionWithCondition();
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
	public relationDefTypeRestrictionWithCondition(): RelationDefTypeRestrictionWithConditionContext {
		let localctx: RelationDefTypeRestrictionWithConditionContext = new RelationDefTypeRestrictionWithConditionContext(this, this._ctx, this.state);
		this.enterRule(localctx, 42, OpenFGAParser.RULE_relationDefTypeRestrictionWithCondition);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 267;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 29, this._ctx) ) {
			case 1:
				{
				this.state = 264;
				this.relationDefTypeRestrictionType();
				}
				break;
			case 2:
				{
				this.state = 265;
				this.relationDefTypeRestrictionWildcard();
				}
				break;
			case 3:
				{
				this.state = 266;
				this.relationDefTypeRestrictionUserset();
				}
				break;
			}
			this.state = 269;
			this.spacing();
			this.state = 270;
			this.match(OpenFGAParser.WTH);
			this.state = 271;
			this.spacing();
			this.state = 272;
			this.conditionName();
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
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
		this.enterRule(localctx, 44, OpenFGAParser.RULE_relationDefTypeRestrictionType);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 274;
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
		this.enterRule(localctx, 46, OpenFGAParser.RULE_relationDefTypeRestrictionRelation);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 276;
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
		this.enterRule(localctx, 48, OpenFGAParser.RULE_relationDefTypeRestrictionWildcard);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 278;
			this.relationDefTypeRestrictionType();
			this.state = 279;
			this.match(OpenFGAParser.COLON);
			this.state = 280;
			this.match(OpenFGAParser.WILDCARD);
			this.state = 282;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 30, this._ctx) ) {
			case 1:
				{
				this.state = 281;
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
		this.enterRule(localctx, 50, OpenFGAParser.RULE_relationDefTypeRestrictionUserset);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 284;
			this.relationDefTypeRestrictionType();
			this.state = 285;
			this.match(OpenFGAParser.HASH);
			this.state = 286;
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
		this.enterRule(localctx, 52, OpenFGAParser.RULE_relationDefGrouping);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 288;
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
		this.enterRule(localctx, 54, OpenFGAParser.RULE_rewriteComputedusersetName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 290;
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
			this.state = 292;
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
			this.state = 294;
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
			this.state = 296;
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
			this.state = 298;
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
	public conditions(): ConditionsContext {
		let localctx: ConditionsContext = new ConditionsContext(this, this._ctx, this.state);
		this.enterRule(localctx, 64, OpenFGAParser.RULE_conditions);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 303;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 31, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 300;
					this.condition();
					}
					}
				}
				this.state = 305;
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
	public condition(): ConditionContext {
		let localctx: ConditionContext = new ConditionContext(this, this._ctx, this.state);
		this.enterRule(localctx, 66, OpenFGAParser.RULE_condition);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 309;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 32, this._ctx) ) {
			case 1:
				{
				this.state = 306;
				this.newline();
				this.state = 307;
				this.multiLineComment();
				}
				break;
			}
			this.state = 311;
			this.newline();
			this.state = 312;
			this.match(OpenFGAParser.CONDITION);
			this.state = 313;
			this.spacing();
			this.state = 314;
			this.conditionName();
			this.state = 316;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===31) {
				{
				this.state = 315;
				this.spacing();
				}
			}

			this.state = 318;
			this.match(OpenFGAParser.L_PARANTHESES);
			this.state = 319;
			this.conditionParameter();
			this.state = 321;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===31) {
				{
				this.state = 320;
				this.spacing();
				}
			}

			this.state = 333;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===21) {
				{
				{
				this.state = 323;
				this.match(OpenFGAParser.COMMA);
				this.state = 325;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===31) {
					{
					this.state = 324;
					this.spacing();
					}
				}

				this.state = 327;
				this.conditionParameter();
				this.state = 329;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===31) {
					{
					this.state = 328;
					this.spacing();
					}
				}

				}
				}
				this.state = 335;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 336;
			this.match(OpenFGAParser.R_PARANTHESES);
			this.state = 338;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===31) {
				{
				this.state = 337;
				this.spacing();
				}
			}

			this.state = 340;
			this.match(OpenFGAParser.L_BRACES);
			this.state = 341;
			this.conditionExpression();
			this.state = 342;
			this.match(OpenFGAParser.R_BRACES);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
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
		this.enterRule(localctx, 68, OpenFGAParser.RULE_conditionParameter);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 344;
			this.parameterName();
			this.state = 346;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===31) {
				{
				this.state = 345;
				this.spacing();
				}
			}

			this.state = 348;
			this.match(OpenFGAParser.COLON);
			this.state = 350;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===31) {
				{
				this.state = 349;
				this.spacing();
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
		this.enterRule(localctx, 70, OpenFGAParser.RULE_parameterName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 354;
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
	public conditionName(): ConditionNameContext {
		let localctx: ConditionNameContext = new ConditionNameContext(this, this._ctx, this.state);
		this.enterRule(localctx, 72, OpenFGAParser.RULE_conditionName);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 356;
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
	public parameterType(): ParameterTypeContext {
		let localctx: ParameterTypeContext = new ParameterTypeContext(this, this._ctx, this.state);
		this.enterRule(localctx, 74, OpenFGAParser.RULE_parameterType);
		try {
			this.state = 363;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 23:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 358;
				this.match(OpenFGAParser.CONDITION_PARAM_TYPE);
				}
				break;
			case 22:
				this.enterOuterAlt(localctx, 2);
				{
				{
				this.state = 359;
				this.match(OpenFGAParser.CONDITION_PARAM_CONTAINER);
				this.state = 360;
				this.match(OpenFGAParser.L_ANGLE_BRACKET);
				this.state = 361;
				this.match(OpenFGAParser.CONDITION_PARAM_TYPE);
				this.state = 362;
				this.match(OpenFGAParser.R_ANGLE_BRACKET);
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
	public conditionExpression(): ConditionExpressionContext {
		let localctx: ConditionExpressionContext = new ConditionExpressionContext(this, this._ctx, this.state);
		this.enterRule(localctx, 76, OpenFGAParser.RULE_conditionExpression);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 369;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while ((((_la) & ~0x1F) === 0 && ((1 << _la) & 4294705150) !== 0)) {
				{
				this.state = 367;
				this._errHandler.sync(this);
				switch ( this._interp.adaptivePredict(this._input, 42, this._ctx) ) {
				case 1:
					{
					this.state = 365;
					this.match(OpenFGAParser.CONDITION_SYMBOL);
					}
					break;
				case 2:
					{
					this.state = 366;
					_la = this._input.LA(1);
					if(_la<=0 || _la===18) {
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
				this.state = 371;
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
	public comment(): CommentContext {
		let localctx: CommentContext = new CommentContext(this, this._ctx, this.state);
		this.enterRule(localctx, 78, OpenFGAParser.RULE_comment);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 375;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===31) {
				{
				{
				this.state = 372;
				this.match(OpenFGAParser.WS);
				}
				}
				this.state = 377;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 378;
			this.match(OpenFGAParser.HASH);
			this.state = 382;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 45, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 379;
					_la = this._input.LA(1);
					if(_la<=0 || _la===30) {
					this._errHandler.recoverInline(this);
					}
					else {
						this._errHandler.reportMatch(this);
					    this.consume();
					}
					}
					}
				}
				this.state = 384;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 45, this._ctx);
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
		this.enterRule(localctx, 80, OpenFGAParser.RULE_multiLineComment);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 385;
			this.comment();
			this.state = 391;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 46, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 386;
					this.newline();
					this.state = 387;
					this.comment();
					}
					}
				}
				this.state = 393;
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
	public spacing(): SpacingContext {
		let localctx: SpacingContext = new SpacingContext(this, this._ctx, this.state);
		this.enterRule(localctx, 82, OpenFGAParser.RULE_spacing);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 395;
			this._errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					this.state = 394;
					this.match(OpenFGAParser.WS);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				this.state = 397;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 47, this._ctx);
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
		this.enterRule(localctx, 84, OpenFGAParser.RULE_newline);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 400;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			do {
				{
				{
				this.state = 399;
				this.match(OpenFGAParser.NEWLINE);
				}
				}
				this.state = 402;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			} while (_la===30);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
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
		this.enterRule(localctx, 86, OpenFGAParser.RULE_schemaVersion);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 404;
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
		this.enterRule(localctx, 88, OpenFGAParser.RULE_name);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 407;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			do {
				{
				{
				this.state = 406;
				this.match(OpenFGAParser.ALPHA_NUMERIC);
				}
				}
				this.state = 409;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			} while (_la===29);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}

	public static readonly _serializedATN: number[] = [4,1,31,412,2,0,7,0,2,
	1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,
	10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,16,2,17,
	7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,2,22,7,22,2,23,7,23,2,24,7,
	24,2,25,7,25,2,26,7,26,2,27,7,27,2,28,7,28,2,29,7,29,2,30,7,30,2,31,7,31,
	2,32,7,32,2,33,7,33,2,34,7,34,2,35,7,35,2,36,7,36,2,37,7,37,2,38,7,38,2,
	39,7,39,2,40,7,40,2,41,7,41,2,42,7,42,2,43,7,43,2,44,7,44,1,0,1,0,1,0,1,
	0,3,0,95,8,0,1,0,1,0,1,1,1,1,1,2,1,2,1,2,3,2,104,8,2,1,2,1,2,3,2,108,8,
	2,1,2,1,2,1,2,3,2,113,8,2,1,2,1,2,1,2,1,2,1,2,3,2,120,8,2,1,3,5,3,123,8,
	3,10,3,12,3,126,9,3,1,4,1,4,1,4,3,4,131,8,4,1,4,1,4,1,4,1,4,1,4,3,4,138,
	8,4,1,4,1,4,1,4,3,4,143,8,4,1,4,4,4,146,8,4,11,4,12,4,147,3,4,150,8,4,1,
	5,1,5,1,5,3,5,155,8,5,1,5,1,5,1,5,1,5,1,5,3,5,162,8,5,1,5,1,5,3,5,166,8,
	5,1,5,1,5,3,5,170,8,5,1,6,1,6,3,6,174,8,6,1,6,3,6,177,8,6,1,7,1,7,1,7,3,
	7,182,8,7,1,8,1,8,1,8,1,8,1,8,4,8,189,8,8,11,8,12,8,190,1,9,1,9,1,9,1,9,
	1,9,4,9,198,8,9,11,9,12,9,199,1,10,1,10,1,10,1,10,1,10,4,10,207,8,10,11,
	10,12,10,208,1,11,1,11,3,11,213,8,11,1,11,1,11,3,11,217,8,11,1,11,1,11,
	3,11,221,8,11,1,11,5,11,224,8,11,10,11,12,11,227,9,11,1,11,3,11,230,8,11,
	1,11,1,11,1,12,1,12,3,12,236,8,12,1,13,1,13,1,14,1,14,1,14,1,14,1,14,1,
	14,1,15,1,15,1,15,3,15,249,8,15,1,16,1,16,1,17,1,17,1,18,1,18,1,19,1,19,
	1,20,1,20,1,20,1,20,3,20,263,8,20,1,21,1,21,1,21,3,21,268,8,21,1,21,1,21,
	1,21,1,21,1,21,1,22,1,22,1,23,1,23,1,24,1,24,1,24,1,24,3,24,283,8,24,1,
	25,1,25,1,25,1,25,1,26,1,26,1,27,1,27,1,28,1,28,1,29,1,29,1,30,1,30,1,31,
	1,31,1,32,5,32,302,8,32,10,32,12,32,305,9,32,1,33,1,33,1,33,3,33,310,8,
	33,1,33,1,33,1,33,1,33,1,33,3,33,317,8,33,1,33,1,33,1,33,3,33,322,8,33,
	1,33,1,33,3,33,326,8,33,1,33,1,33,3,33,330,8,33,5,33,332,8,33,10,33,12,
	33,335,9,33,1,33,1,33,3,33,339,8,33,1,33,1,33,1,33,1,33,1,34,1,34,3,34,
	347,8,34,1,34,1,34,3,34,351,8,34,1,34,1,34,1,35,1,35,1,36,1,36,1,37,1,37,
	1,37,1,37,1,37,3,37,364,8,37,1,38,1,38,5,38,368,8,38,10,38,12,38,371,9,
	38,1,39,5,39,374,8,39,10,39,12,39,377,9,39,1,39,1,39,5,39,381,8,39,10,39,
	12,39,384,9,39,1,40,1,40,1,40,1,40,5,40,390,8,40,10,40,12,40,393,9,40,1,
	41,4,41,396,8,41,11,41,12,41,397,1,42,4,42,401,8,42,11,42,12,42,402,1,43,
	1,43,1,44,4,44,408,8,44,11,44,12,44,409,1,44,0,0,45,0,2,4,6,8,10,12,14,
	16,18,20,22,24,26,28,30,32,34,36,38,40,42,44,46,48,50,52,54,56,58,60,62,
	64,66,68,70,72,74,76,78,80,82,84,86,88,0,2,1,0,18,18,1,0,30,30,421,0,90,
	1,0,0,0,2,98,1,0,0,0,4,103,1,0,0,0,6,124,1,0,0,0,8,130,1,0,0,0,10,154,1,
	0,0,0,12,173,1,0,0,0,14,181,1,0,0,0,16,188,1,0,0,0,18,197,1,0,0,0,20,206,
	1,0,0,0,22,210,1,0,0,0,24,235,1,0,0,0,26,237,1,0,0,0,28,239,1,0,0,0,30,
	248,1,0,0,0,32,250,1,0,0,0,34,252,1,0,0,0,36,254,1,0,0,0,38,256,1,0,0,0,
	40,262,1,0,0,0,42,267,1,0,0,0,44,274,1,0,0,0,46,276,1,0,0,0,48,278,1,0,
	0,0,50,284,1,0,0,0,52,288,1,0,0,0,54,290,1,0,0,0,56,292,1,0,0,0,58,294,
	1,0,0,0,60,296,1,0,0,0,62,298,1,0,0,0,64,303,1,0,0,0,66,309,1,0,0,0,68,
	344,1,0,0,0,70,354,1,0,0,0,72,356,1,0,0,0,74,363,1,0,0,0,76,369,1,0,0,0,
	78,375,1,0,0,0,80,385,1,0,0,0,82,395,1,0,0,0,84,400,1,0,0,0,86,404,1,0,
	0,0,88,407,1,0,0,0,90,91,3,4,2,0,91,92,3,6,3,0,92,94,3,64,32,0,93,95,3,
	84,42,0,94,93,1,0,0,0,94,95,1,0,0,0,95,96,1,0,0,0,96,97,5,0,0,1,97,1,1,
	0,0,0,98,99,5,1,0,0,99,3,1,0,0,0,100,101,3,80,40,0,101,102,3,84,42,0,102,
	104,1,0,0,0,103,100,1,0,0,0,103,104,1,0,0,0,104,105,1,0,0,0,105,107,5,2,
	0,0,106,108,3,82,41,0,107,106,1,0,0,0,107,108,1,0,0,0,108,112,1,0,0,0,109,
	110,3,84,42,0,110,111,3,80,40,0,111,113,1,0,0,0,112,109,1,0,0,0,112,113,
	1,0,0,0,113,114,1,0,0,0,114,115,3,2,1,0,115,116,5,3,0,0,116,117,3,82,41,
	0,117,119,3,86,43,0,118,120,3,82,41,0,119,118,1,0,0,0,119,120,1,0,0,0,120,
	5,1,0,0,0,121,123,3,8,4,0,122,121,1,0,0,0,123,126,1,0,0,0,124,122,1,0,0,
	0,124,125,1,0,0,0,125,7,1,0,0,0,126,124,1,0,0,0,127,128,3,84,42,0,128,129,
	3,80,40,0,129,131,1,0,0,0,130,127,1,0,0,0,130,131,1,0,0,0,131,132,1,0,0,
	0,132,133,3,84,42,0,133,134,5,5,0,0,134,135,3,82,41,0,135,137,3,62,31,0,
	136,138,3,82,41,0,137,136,1,0,0,0,137,138,1,0,0,0,138,149,1,0,0,0,139,140,
	3,2,1,0,140,142,5,7,0,0,141,143,3,82,41,0,142,141,1,0,0,0,142,143,1,0,0,
	0,143,145,1,0,0,0,144,146,3,10,5,0,145,144,1,0,0,0,146,147,1,0,0,0,147,
	145,1,0,0,0,147,148,1,0,0,0,148,150,1,0,0,0,149,139,1,0,0,0,149,150,1,0,
	0,0,150,9,1,0,0,0,151,152,3,84,42,0,152,153,3,80,40,0,153,155,1,0,0,0,154,
	151,1,0,0,0,154,155,1,0,0,0,155,156,1,0,0,0,156,157,3,2,1,0,157,158,5,8,
	0,0,158,159,3,82,41,0,159,161,3,60,30,0,160,162,3,82,41,0,161,160,1,0,0,
	0,161,162,1,0,0,0,162,163,1,0,0,0,163,165,5,11,0,0,164,166,3,82,41,0,165,
	164,1,0,0,0,165,166,1,0,0,0,166,167,1,0,0,0,167,169,3,12,6,0,168,170,3,
	82,41,0,169,168,1,0,0,0,169,170,1,0,0,0,170,11,1,0,0,0,171,174,3,22,11,
	0,172,174,3,52,26,0,173,171,1,0,0,0,173,172,1,0,0,0,174,176,1,0,0,0,175,
	177,3,14,7,0,176,175,1,0,0,0,176,177,1,0,0,0,177,13,1,0,0,0,178,182,3,16,
	8,0,179,182,3,18,9,0,180,182,3,20,10,0,181,178,1,0,0,0,181,179,1,0,0,0,
	181,180,1,0,0,0,182,15,1,0,0,0,183,184,3,82,41,0,184,185,3,34,17,0,185,
	186,3,82,41,0,186,187,3,52,26,0,187,189,1,0,0,0,188,183,1,0,0,0,189,190,
	1,0,0,0,190,188,1,0,0,0,190,191,1,0,0,0,191,17,1,0,0,0,192,193,3,82,41,
	0,193,194,3,32,16,0,194,195,3,82,41,0,195,196,3,52,26,0,196,198,1,0,0,0,
	197,192,1,0,0,0,198,199,1,0,0,0,199,197,1,0,0,0,199,200,1,0,0,0,200,19,
	1,0,0,0,201,202,3,82,41,0,202,203,3,36,18,0,203,204,3,82,41,0,204,205,3,
	52,26,0,205,207,1,0,0,0,206,201,1,0,0,0,207,208,1,0,0,0,208,206,1,0,0,0,
	208,209,1,0,0,0,209,21,1,0,0,0,210,212,5,13,0,0,211,213,3,82,41,0,212,211,
	1,0,0,0,212,213,1,0,0,0,213,214,1,0,0,0,214,216,3,40,20,0,215,217,3,82,
	41,0,216,215,1,0,0,0,216,217,1,0,0,0,217,225,1,0,0,0,218,220,5,21,0,0,219,
	221,3,82,41,0,220,219,1,0,0,0,220,221,1,0,0,0,221,222,1,0,0,0,222,224,3,
	40,20,0,223,218,1,0,0,0,224,227,1,0,0,0,225,223,1,0,0,0,225,226,1,0,0,0,
	226,229,1,0,0,0,227,225,1,0,0,0,228,230,3,82,41,0,229,228,1,0,0,0,229,230,
	1,0,0,0,230,231,1,0,0,0,231,232,5,14,0,0,232,23,1,0,0,0,233,236,3,26,13,
	0,234,236,3,28,14,0,235,233,1,0,0,0,235,234,1,0,0,0,236,25,1,0,0,0,237,
	238,3,54,27,0,238,27,1,0,0,0,239,240,3,56,28,0,240,241,3,82,41,0,241,242,
	3,38,19,0,242,243,3,82,41,0,243,244,3,58,29,0,244,29,1,0,0,0,245,249,3,
	34,17,0,246,249,3,32,16,0,247,249,3,36,18,0,248,245,1,0,0,0,248,246,1,0,
	0,0,248,247,1,0,0,0,249,31,1,0,0,0,250,251,5,25,0,0,251,33,1,0,0,0,252,
	253,5,26,0,0,253,35,1,0,0,0,254,255,5,27,0,0,255,37,1,0,0,0,256,257,5,28,
	0,0,257,39,1,0,0,0,258,263,3,44,22,0,259,263,3,48,24,0,260,263,3,50,25,
	0,261,263,3,42,21,0,262,258,1,0,0,0,262,259,1,0,0,0,262,260,1,0,0,0,262,
	261,1,0,0,0,263,41,1,0,0,0,264,268,3,44,22,0,265,268,3,48,24,0,266,268,
	3,50,25,0,267,264,1,0,0,0,267,265,1,0,0,0,267,266,1,0,0,0,268,269,1,0,0,
	0,269,270,3,82,41,0,270,271,5,9,0,0,271,272,3,82,41,0,272,273,3,72,36,0,
	273,43,1,0,0,0,274,275,3,88,44,0,275,45,1,0,0,0,276,277,3,88,44,0,277,47,
	1,0,0,0,278,279,3,44,22,0,279,280,5,11,0,0,280,282,5,12,0,0,281,283,3,82,
	41,0,282,281,1,0,0,0,282,283,1,0,0,0,283,49,1,0,0,0,284,285,3,44,22,0,285,
	286,5,10,0,0,286,287,3,46,23,0,287,51,1,0,0,0,288,289,3,24,12,0,289,53,
	1,0,0,0,290,291,3,88,44,0,291,55,1,0,0,0,292,293,3,88,44,0,293,57,1,0,0,
	0,294,295,3,88,44,0,295,59,1,0,0,0,296,297,3,88,44,0,297,61,1,0,0,0,298,
	299,3,88,44,0,299,63,1,0,0,0,300,302,3,66,33,0,301,300,1,0,0,0,302,305,
	1,0,0,0,303,301,1,0,0,0,303,304,1,0,0,0,304,65,1,0,0,0,305,303,1,0,0,0,
	306,307,3,84,42,0,307,308,3,80,40,0,308,310,1,0,0,0,309,306,1,0,0,0,309,
	310,1,0,0,0,310,311,1,0,0,0,311,312,3,84,42,0,312,313,5,6,0,0,313,314,3,
	82,41,0,314,316,3,72,36,0,315,317,3,82,41,0,316,315,1,0,0,0,316,317,1,0,
	0,0,317,318,1,0,0,0,318,319,5,15,0,0,319,321,3,68,34,0,320,322,3,82,41,
	0,321,320,1,0,0,0,321,322,1,0,0,0,322,333,1,0,0,0,323,325,5,21,0,0,324,
	326,3,82,41,0,325,324,1,0,0,0,325,326,1,0,0,0,326,327,1,0,0,0,327,329,3,
	68,34,0,328,330,3,82,41,0,329,328,1,0,0,0,329,330,1,0,0,0,330,332,1,0,0,
	0,331,323,1,0,0,0,332,335,1,0,0,0,333,331,1,0,0,0,333,334,1,0,0,0,334,336,
	1,0,0,0,335,333,1,0,0,0,336,338,5,16,0,0,337,339,3,82,41,0,338,337,1,0,
	0,0,338,339,1,0,0,0,339,340,1,0,0,0,340,341,5,17,0,0,341,342,3,76,38,0,
	342,343,5,18,0,0,343,67,1,0,0,0,344,346,3,70,35,0,345,347,3,82,41,0,346,
	345,1,0,0,0,346,347,1,0,0,0,347,348,1,0,0,0,348,350,5,11,0,0,349,351,3,
	82,41,0,350,349,1,0,0,0,350,351,1,0,0,0,351,352,1,0,0,0,352,353,3,74,37,
	0,353,69,1,0,0,0,354,355,3,88,44,0,355,71,1,0,0,0,356,357,3,88,44,0,357,
	73,1,0,0,0,358,364,5,23,0,0,359,360,5,22,0,0,360,361,5,19,0,0,361,362,5,
	23,0,0,362,364,5,20,0,0,363,358,1,0,0,0,363,359,1,0,0,0,364,75,1,0,0,0,
	365,368,5,24,0,0,366,368,8,0,0,0,367,365,1,0,0,0,367,366,1,0,0,0,368,371,
	1,0,0,0,369,367,1,0,0,0,369,370,1,0,0,0,370,77,1,0,0,0,371,369,1,0,0,0,
	372,374,5,31,0,0,373,372,1,0,0,0,374,377,1,0,0,0,375,373,1,0,0,0,375,376,
	1,0,0,0,376,378,1,0,0,0,377,375,1,0,0,0,378,382,5,10,0,0,379,381,8,1,0,
	0,380,379,1,0,0,0,381,384,1,0,0,0,382,380,1,0,0,0,382,383,1,0,0,0,383,79,
	1,0,0,0,384,382,1,0,0,0,385,391,3,78,39,0,386,387,3,84,42,0,387,388,3,78,
	39,0,388,390,1,0,0,0,389,386,1,0,0,0,390,393,1,0,0,0,391,389,1,0,0,0,391,
	392,1,0,0,0,392,81,1,0,0,0,393,391,1,0,0,0,394,396,5,31,0,0,395,394,1,0,
	0,0,396,397,1,0,0,0,397,395,1,0,0,0,397,398,1,0,0,0,398,83,1,0,0,0,399,
	401,5,30,0,0,400,399,1,0,0,0,401,402,1,0,0,0,402,400,1,0,0,0,402,403,1,
	0,0,0,403,85,1,0,0,0,404,405,5,4,0,0,405,87,1,0,0,0,406,408,5,29,0,0,407,
	406,1,0,0,0,408,409,1,0,0,0,409,407,1,0,0,0,409,410,1,0,0,0,410,89,1,0,
	0,0,50,94,103,107,112,119,124,130,137,142,147,149,154,161,165,169,173,176,
	181,190,199,208,212,216,220,225,229,235,248,262,267,282,303,309,316,321,
	325,329,333,338,346,350,363,367,369,375,382,391,397,402,409];

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
	public relationDefTypeRestrictionWithCondition(): RelationDefTypeRestrictionWithConditionContext {
		return this.getTypedRuleContext(RelationDefTypeRestrictionWithConditionContext, 0) as RelationDefTypeRestrictionWithConditionContext;
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


export class RelationDefTypeRestrictionWithConditionContext extends ParserRuleContext {
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
	public WTH(): TerminalNode {
		return this.getToken(OpenFGAParser.WTH, 0);
	}
	public conditionName(): ConditionNameContext {
		return this.getTypedRuleContext(ConditionNameContext, 0) as ConditionNameContext;
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
    	return OpenFGAParser.RULE_relationDefTypeRestrictionWithCondition;
	}
	public enterRule(listener: OpenFGAParserListener): void {
	    if(listener.enterRelationDefTypeRestrictionWithCondition) {
	 		listener.enterRelationDefTypeRestrictionWithCondition(this);
		}
	}
	public exitRule(listener: OpenFGAParserListener): void {
	    if(listener.exitRelationDefTypeRestrictionWithCondition) {
	 		listener.exitRelationDefTypeRestrictionWithCondition(this);
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
	public newline_list(): NewlineContext[] {
		return this.getTypedRuleContexts(NewlineContext) as NewlineContext[];
	}
	public newline(i: number): NewlineContext {
		return this.getTypedRuleContext(NewlineContext, i) as NewlineContext;
	}
	public CONDITION(): TerminalNode {
		return this.getToken(OpenFGAParser.CONDITION, 0);
	}
	public spacing_list(): SpacingContext[] {
		return this.getTypedRuleContexts(SpacingContext) as SpacingContext[];
	}
	public spacing(i: number): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, i) as SpacingContext;
	}
	public conditionName(): ConditionNameContext {
		return this.getTypedRuleContext(ConditionNameContext, 0) as ConditionNameContext;
	}
	public L_PARANTHESES(): TerminalNode {
		return this.getToken(OpenFGAParser.L_PARANTHESES, 0);
	}
	public conditionParameter_list(): ConditionParameterContext[] {
		return this.getTypedRuleContexts(ConditionParameterContext) as ConditionParameterContext[];
	}
	public conditionParameter(i: number): ConditionParameterContext {
		return this.getTypedRuleContext(ConditionParameterContext, i) as ConditionParameterContext;
	}
	public R_PARANTHESES(): TerminalNode {
		return this.getToken(OpenFGAParser.R_PARANTHESES, 0);
	}
	public L_BRACES(): TerminalNode {
		return this.getToken(OpenFGAParser.L_BRACES, 0);
	}
	public conditionExpression(): ConditionExpressionContext {
		return this.getTypedRuleContext(ConditionExpressionContext, 0) as ConditionExpressionContext;
	}
	public R_BRACES(): TerminalNode {
		return this.getToken(OpenFGAParser.R_BRACES, 0);
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
	public spacing_list(): SpacingContext[] {
		return this.getTypedRuleContexts(SpacingContext) as SpacingContext[];
	}
	public spacing(i: number): SpacingContext {
		return this.getTypedRuleContext(SpacingContext, i) as SpacingContext;
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
	public name(): NameContext {
		return this.getTypedRuleContext(NameContext, 0) as NameContext;
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


export class ConditionNameContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public name(): NameContext {
		return this.getTypedRuleContext(NameContext, 0) as NameContext;
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
	public L_ANGLE_BRACKET(): TerminalNode {
		return this.getToken(OpenFGAParser.L_ANGLE_BRACKET, 0);
	}
	public R_ANGLE_BRACKET(): TerminalNode {
		return this.getToken(OpenFGAParser.R_ANGLE_BRACKET, 0);
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


export class ConditionExpressionContext extends ParserRuleContext {
	constructor(parser?: OpenFGAParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public CONDITION_SYMBOL_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.CONDITION_SYMBOL);
	}
	public CONDITION_SYMBOL(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.CONDITION_SYMBOL, i);
	}
	public R_BRACES_list(): TerminalNode[] {
	    	return this.getTokens(OpenFGAParser.R_BRACES);
	}
	public R_BRACES(i: number): TerminalNode {
		return this.getToken(OpenFGAParser.R_BRACES, i);
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
