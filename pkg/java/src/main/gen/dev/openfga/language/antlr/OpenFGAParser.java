// Generated from /app/OpenFGAParser.g4 by ANTLR 4.13.1
package dev.openfga.language.antlr;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class OpenFGAParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		COLON=1, COMMA=2, LESS=3, GREATER=4, LBRACKET=5, RBRACKET=6, LPAREN=7, 
		RPAREN=8, WHITESPACE=9, IDENTIFIER=10, HASH=11, AND=12, OR=13, BUT_NOT=14, 
		FROM=15, MODEL=16, SCHEMA=17, SCHEMA_VERSION=18, TYPE=19, CONDITION=20, 
		RELATIONS=21, RELATION=22, DEFINE=23, KEYWORD_WITH=24, EQUALS=25, NOT_EQUALS=26, 
		IN=27, LESS_EQUALS=28, GREATER_EQUALS=29, LOGICAL_AND=30, LOGICAL_OR=31, 
		RPRACKET=32, LBRACE=33, RBRACE=34, DOT=35, MINUS=36, EXCLAM=37, QUESTIONMARK=38, 
		PLUS=39, STAR=40, SLASH=41, PERCENT=42, CEL_TRUE=43, CEL_FALSE=44, NUL=45, 
		CEL_COMMENT=46, NUM_FLOAT=47, NUM_INT=48, NUM_UINT=49, STRING=50, BYTES=51, 
		NEWLINE=52, CONDITION_PARAM_CONTAINER=53, CONDITION_PARAM_TYPE=54;
	public static final int
		RULE_main = 0, RULE_modelHeader = 1, RULE_typeDefs = 2, RULE_typeDef = 3, 
		RULE_relationDeclaration = 4, RULE_relationName = 5, RULE_relationDef = 6, 
		RULE_relationDefNoDirect = 7, RULE_relationDefPartials = 8, RULE_relationDefGrouping = 9, 
		RULE_relationRecurse = 10, RULE_relationRecurseNoDirect = 11, RULE_relationDefDirectAssignment = 12, 
		RULE_relationDefRewrite = 13, RULE_relationDefTypeRestriction = 14, RULE_relationDefTypeRestrictionBase = 15, 
		RULE_conditions = 16, RULE_condition = 17, RULE_conditionName = 18, RULE_conditionParameter = 19, 
		RULE_parameterName = 20, RULE_parameterType = 21, RULE_multiLineComment = 22, 
		RULE_identifier = 23, RULE_conditionExpression = 24;
	private static String[] makeRuleNames() {
		return new String[] {
			"main", "modelHeader", "typeDefs", "typeDef", "relationDeclaration", 
			"relationName", "relationDef", "relationDefNoDirect", "relationDefPartials", 
			"relationDefGrouping", "relationRecurse", "relationRecurseNoDirect", 
			"relationDefDirectAssignment", "relationDefRewrite", "relationDefTypeRestriction", 
			"relationDefTypeRestrictionBase", "conditions", "condition", "conditionName", 
			"conditionParameter", "parameterName", "parameterType", "multiLineComment", 
			"identifier", "conditionExpression"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "':'", "','", "'<'", "'>'", "'['", null, "'('", "')'", null, null, 
			"'#'", "'and'", "'or'", "'but not'", "'from'", "'model'", "'schema'", 
			"'1.1'", "'type'", "'condition'", "'relations'", "'relation'", "'define'", 
			"'with'", "'=='", "'!='", "'in'", "'<='", "'>='", "'&&'", "'||'", "']'", 
			"'{'", "'}'", "'.'", "'-'", "'!'", "'?'", "'+'", "'*'", "'/'", "'%'", 
			"'true'", "'false'", "'null'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "COLON", "COMMA", "LESS", "GREATER", "LBRACKET", "RBRACKET", "LPAREN", 
			"RPAREN", "WHITESPACE", "IDENTIFIER", "HASH", "AND", "OR", "BUT_NOT", 
			"FROM", "MODEL", "SCHEMA", "SCHEMA_VERSION", "TYPE", "CONDITION", "RELATIONS", 
			"RELATION", "DEFINE", "KEYWORD_WITH", "EQUALS", "NOT_EQUALS", "IN", "LESS_EQUALS", 
			"GREATER_EQUALS", "LOGICAL_AND", "LOGICAL_OR", "RPRACKET", "LBRACE", 
			"RBRACE", "DOT", "MINUS", "EXCLAM", "QUESTIONMARK", "PLUS", "STAR", "SLASH", 
			"PERCENT", "CEL_TRUE", "CEL_FALSE", "NUL", "CEL_COMMENT", "NUM_FLOAT", 
			"NUM_INT", "NUM_UINT", "STRING", "BYTES", "NEWLINE", "CONDITION_PARAM_CONTAINER", 
			"CONDITION_PARAM_TYPE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "OpenFGAParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public OpenFGAParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MainContext extends ParserRuleContext {
		public ModelHeaderContext modelHeader() {
			return getRuleContext(ModelHeaderContext.class,0);
		}
		public TypeDefsContext typeDefs() {
			return getRuleContext(TypeDefsContext.class,0);
		}
		public ConditionsContext conditions() {
			return getRuleContext(ConditionsContext.class,0);
		}
		public TerminalNode EOF() { return getToken(OpenFGAParser.EOF, 0); }
		public TerminalNode WHITESPACE() { return getToken(OpenFGAParser.WHITESPACE, 0); }
		public List<TerminalNode> NEWLINE() { return getTokens(OpenFGAParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(OpenFGAParser.NEWLINE, i);
		}
		public MainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_main; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterMain(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitMain(this);
		}
	}

	public final MainContext main() throws RecognitionException {
		MainContext _localctx = new MainContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_main);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(51);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(50);
				match(WHITESPACE);
				}
			}

			setState(54);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NEWLINE) {
				{
				setState(53);
				match(NEWLINE);
				}
			}

			setState(56);
			modelHeader();
			setState(58);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				{
				setState(57);
				match(NEWLINE);
				}
				break;
			}
			setState(60);
			typeDefs();
			setState(62);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				{
				setState(61);
				match(NEWLINE);
				}
				break;
			}
			setState(64);
			conditions();
			setState(66);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NEWLINE) {
				{
				setState(65);
				match(NEWLINE);
				}
			}

			setState(68);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ModelHeaderContext extends ParserRuleContext {
		public Token schemaVersion;
		public TerminalNode MODEL() { return getToken(OpenFGAParser.MODEL, 0); }
		public List<TerminalNode> NEWLINE() { return getTokens(OpenFGAParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(OpenFGAParser.NEWLINE, i);
		}
		public TerminalNode SCHEMA() { return getToken(OpenFGAParser.SCHEMA, 0); }
		public List<TerminalNode> WHITESPACE() { return getTokens(OpenFGAParser.WHITESPACE); }
		public TerminalNode WHITESPACE(int i) {
			return getToken(OpenFGAParser.WHITESPACE, i);
		}
		public TerminalNode SCHEMA_VERSION() { return getToken(OpenFGAParser.SCHEMA_VERSION, 0); }
		public MultiLineCommentContext multiLineComment() {
			return getRuleContext(MultiLineCommentContext.class,0);
		}
		public ModelHeaderContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_modelHeader; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterModelHeader(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitModelHeader(this);
		}
	}

	public final ModelHeaderContext modelHeader() throws RecognitionException {
		ModelHeaderContext _localctx = new ModelHeaderContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_modelHeader);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==HASH) {
				{
				setState(70);
				multiLineComment();
				setState(71);
				match(NEWLINE);
				}
			}

			setState(75);
			match(MODEL);
			setState(76);
			match(NEWLINE);
			setState(77);
			match(SCHEMA);
			setState(78);
			match(WHITESPACE);
			setState(79);
			((ModelHeaderContext)_localctx).schemaVersion = match(SCHEMA_VERSION);
			setState(81);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(80);
				match(WHITESPACE);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeDefsContext extends ParserRuleContext {
		public List<TypeDefContext> typeDef() {
			return getRuleContexts(TypeDefContext.class);
		}
		public TypeDefContext typeDef(int i) {
			return getRuleContext(TypeDefContext.class,i);
		}
		public TypeDefsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeDefs; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterTypeDefs(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitTypeDefs(this);
		}
	}

	public final TypeDefsContext typeDefs() throws RecognitionException {
		TypeDefsContext _localctx = new TypeDefsContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_typeDefs);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(86);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(83);
					typeDef();
					}
					} 
				}
				setState(88);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeDefContext extends ParserRuleContext {
		public IdentifierContext typeName;
		public List<TerminalNode> NEWLINE() { return getTokens(OpenFGAParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(OpenFGAParser.NEWLINE, i);
		}
		public TerminalNode TYPE() { return getToken(OpenFGAParser.TYPE, 0); }
		public TerminalNode WHITESPACE() { return getToken(OpenFGAParser.WHITESPACE, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public MultiLineCommentContext multiLineComment() {
			return getRuleContext(MultiLineCommentContext.class,0);
		}
		public TerminalNode RELATIONS() { return getToken(OpenFGAParser.RELATIONS, 0); }
		public List<RelationDeclarationContext> relationDeclaration() {
			return getRuleContexts(RelationDeclarationContext.class);
		}
		public RelationDeclarationContext relationDeclaration(int i) {
			return getRuleContext(RelationDeclarationContext.class,i);
		}
		public TypeDefContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeDef; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterTypeDef(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitTypeDef(this);
		}
	}

	public final TypeDefContext typeDef() throws RecognitionException {
		TypeDefContext _localctx = new TypeDefContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_typeDef);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(91);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(89);
				match(NEWLINE);
				setState(90);
				multiLineComment();
				}
				break;
			}
			setState(93);
			match(NEWLINE);
			setState(94);
			match(TYPE);
			setState(95);
			match(WHITESPACE);
			setState(96);
			((TypeDefContext)_localctx).typeName = identifier();
			setState(104);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				{
				setState(97);
				match(NEWLINE);
				setState(98);
				match(RELATIONS);
				setState(100); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(99);
						relationDeclaration();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(102); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RelationDeclarationContext extends ParserRuleContext {
		public List<TerminalNode> NEWLINE() { return getTokens(OpenFGAParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(OpenFGAParser.NEWLINE, i);
		}
		public TerminalNode DEFINE() { return getToken(OpenFGAParser.DEFINE, 0); }
		public List<TerminalNode> WHITESPACE() { return getTokens(OpenFGAParser.WHITESPACE); }
		public TerminalNode WHITESPACE(int i) {
			return getToken(OpenFGAParser.WHITESPACE, i);
		}
		public RelationNameContext relationName() {
			return getRuleContext(RelationNameContext.class,0);
		}
		public TerminalNode COLON() { return getToken(OpenFGAParser.COLON, 0); }
		public RelationDefContext relationDef() {
			return getRuleContext(RelationDefContext.class,0);
		}
		public MultiLineCommentContext multiLineComment() {
			return getRuleContext(MultiLineCommentContext.class,0);
		}
		public RelationDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterRelationDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitRelationDeclaration(this);
		}
	}

	public final RelationDeclarationContext relationDeclaration() throws RecognitionException {
		RelationDeclarationContext _localctx = new RelationDeclarationContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_relationDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(108);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				{
				setState(106);
				match(NEWLINE);
				setState(107);
				multiLineComment();
				}
				break;
			}
			setState(110);
			match(NEWLINE);
			setState(111);
			match(DEFINE);
			setState(112);
			match(WHITESPACE);
			setState(113);
			relationName();
			setState(115);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(114);
				match(WHITESPACE);
				}
			}

			setState(117);
			match(COLON);
			setState(119);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(118);
				match(WHITESPACE);
				}
			}

			{
			setState(121);
			relationDef();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RelationNameContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public RelationNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationName; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterRelationName(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitRelationName(this);
		}
	}

	public final RelationNameContext relationName() throws RecognitionException {
		RelationNameContext _localctx = new RelationNameContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_relationName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(123);
			identifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RelationDefContext extends ParserRuleContext {
		public RelationDefDirectAssignmentContext relationDefDirectAssignment() {
			return getRuleContext(RelationDefDirectAssignmentContext.class,0);
		}
		public RelationDefGroupingContext relationDefGrouping() {
			return getRuleContext(RelationDefGroupingContext.class,0);
		}
		public RelationRecurseContext relationRecurse() {
			return getRuleContext(RelationRecurseContext.class,0);
		}
		public RelationDefPartialsContext relationDefPartials() {
			return getRuleContext(RelationDefPartialsContext.class,0);
		}
		public RelationDefContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationDef; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterRelationDef(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitRelationDef(this);
		}
	}

	public final RelationDefContext relationDef() throws RecognitionException {
		RelationDefContext _localctx = new RelationDefContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_relationDef);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LBRACKET:
				{
				setState(125);
				relationDefDirectAssignment();
				}
				break;
			case IDENTIFIER:
				{
				setState(126);
				relationDefGrouping();
				}
				break;
			case LPAREN:
				{
				setState(127);
				relationRecurse();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(131);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				{
				setState(130);
				relationDefPartials();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RelationDefNoDirectContext extends ParserRuleContext {
		public RelationDefGroupingContext relationDefGrouping() {
			return getRuleContext(RelationDefGroupingContext.class,0);
		}
		public RelationRecurseNoDirectContext relationRecurseNoDirect() {
			return getRuleContext(RelationRecurseNoDirectContext.class,0);
		}
		public RelationDefPartialsContext relationDefPartials() {
			return getRuleContext(RelationDefPartialsContext.class,0);
		}
		public RelationDefNoDirectContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationDefNoDirect; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterRelationDefNoDirect(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitRelationDefNoDirect(this);
		}
	}

	public final RelationDefNoDirectContext relationDefNoDirect() throws RecognitionException {
		RelationDefNoDirectContext _localctx = new RelationDefNoDirectContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_relationDefNoDirect);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(135);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				{
				setState(133);
				relationDefGrouping();
				}
				break;
			case LPAREN:
				{
				setState(134);
				relationRecurseNoDirect();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(138);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				{
				setState(137);
				relationDefPartials();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RelationDefPartialsContext extends ParserRuleContext {
		public List<TerminalNode> WHITESPACE() { return getTokens(OpenFGAParser.WHITESPACE); }
		public TerminalNode WHITESPACE(int i) {
			return getToken(OpenFGAParser.WHITESPACE, i);
		}
		public List<TerminalNode> OR() { return getTokens(OpenFGAParser.OR); }
		public TerminalNode OR(int i) {
			return getToken(OpenFGAParser.OR, i);
		}
		public List<RelationDefGroupingContext> relationDefGrouping() {
			return getRuleContexts(RelationDefGroupingContext.class);
		}
		public RelationDefGroupingContext relationDefGrouping(int i) {
			return getRuleContext(RelationDefGroupingContext.class,i);
		}
		public List<RelationRecurseNoDirectContext> relationRecurseNoDirect() {
			return getRuleContexts(RelationRecurseNoDirectContext.class);
		}
		public RelationRecurseNoDirectContext relationRecurseNoDirect(int i) {
			return getRuleContext(RelationRecurseNoDirectContext.class,i);
		}
		public List<TerminalNode> AND() { return getTokens(OpenFGAParser.AND); }
		public TerminalNode AND(int i) {
			return getToken(OpenFGAParser.AND, i);
		}
		public TerminalNode BUT_NOT() { return getToken(OpenFGAParser.BUT_NOT, 0); }
		public RelationDefPartialsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationDefPartials; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterRelationDefPartials(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitRelationDefPartials(this);
		}
	}

	public final RelationDefPartialsContext relationDefPartials() throws RecognitionException {
		RelationDefPartialsContext _localctx = new RelationDefPartialsContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_relationDefPartials);
		try {
			int _alt;
			setState(169);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(147); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(140);
						match(WHITESPACE);
						setState(141);
						match(OR);
						setState(142);
						match(WHITESPACE);
						setState(145);
						_errHandler.sync(this);
						switch (_input.LA(1)) {
						case IDENTIFIER:
							{
							setState(143);
							relationDefGrouping();
							}
							break;
						case LPAREN:
							{
							setState(144);
							relationRecurseNoDirect();
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
					setState(149); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,19,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(158); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(151);
						match(WHITESPACE);
						setState(152);
						match(AND);
						setState(153);
						match(WHITESPACE);
						setState(156);
						_errHandler.sync(this);
						switch (_input.LA(1)) {
						case IDENTIFIER:
							{
							setState(154);
							relationDefGrouping();
							}
							break;
						case LPAREN:
							{
							setState(155);
							relationRecurseNoDirect();
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
					setState(160); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				{
				setState(162);
				match(WHITESPACE);
				setState(163);
				match(BUT_NOT);
				setState(164);
				match(WHITESPACE);
				setState(167);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case IDENTIFIER:
					{
					setState(165);
					relationDefGrouping();
					}
					break;
				case LPAREN:
					{
					setState(166);
					relationRecurseNoDirect();
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
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RelationDefGroupingContext extends ParserRuleContext {
		public RelationDefRewriteContext relationDefRewrite() {
			return getRuleContext(RelationDefRewriteContext.class,0);
		}
		public RelationDefGroupingContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationDefGrouping; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterRelationDefGrouping(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitRelationDefGrouping(this);
		}
	}

	public final RelationDefGroupingContext relationDefGrouping() throws RecognitionException {
		RelationDefGroupingContext _localctx = new RelationDefGroupingContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_relationDefGrouping);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(171);
			relationDefRewrite();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RelationRecurseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(OpenFGAParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(OpenFGAParser.RPAREN, 0); }
		public RelationDefContext relationDef() {
			return getRuleContext(RelationDefContext.class,0);
		}
		public RelationRecurseNoDirectContext relationRecurseNoDirect() {
			return getRuleContext(RelationRecurseNoDirectContext.class,0);
		}
		public List<TerminalNode> WHITESPACE() { return getTokens(OpenFGAParser.WHITESPACE); }
		public TerminalNode WHITESPACE(int i) {
			return getToken(OpenFGAParser.WHITESPACE, i);
		}
		public RelationRecurseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationRecurse; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterRelationRecurse(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitRelationRecurse(this);
		}
	}

	public final RelationRecurseContext relationRecurse() throws RecognitionException {
		RelationRecurseContext _localctx = new RelationRecurseContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_relationRecurse);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(173);
			match(LPAREN);
			setState(177);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==WHITESPACE) {
				{
				{
				setState(174);
				match(WHITESPACE);
				}
				}
				setState(179);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(182);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				{
				setState(180);
				relationDef();
				}
				break;
			case 2:
				{
				setState(181);
				relationRecurseNoDirect();
				}
				break;
			}
			setState(187);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==WHITESPACE) {
				{
				{
				setState(184);
				match(WHITESPACE);
				}
				}
				setState(189);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(190);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RelationRecurseNoDirectContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(OpenFGAParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(OpenFGAParser.RPAREN, 0); }
		public RelationDefNoDirectContext relationDefNoDirect() {
			return getRuleContext(RelationDefNoDirectContext.class,0);
		}
		public RelationRecurseNoDirectContext relationRecurseNoDirect() {
			return getRuleContext(RelationRecurseNoDirectContext.class,0);
		}
		public List<TerminalNode> WHITESPACE() { return getTokens(OpenFGAParser.WHITESPACE); }
		public TerminalNode WHITESPACE(int i) {
			return getToken(OpenFGAParser.WHITESPACE, i);
		}
		public RelationRecurseNoDirectContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationRecurseNoDirect; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterRelationRecurseNoDirect(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitRelationRecurseNoDirect(this);
		}
	}

	public final RelationRecurseNoDirectContext relationRecurseNoDirect() throws RecognitionException {
		RelationRecurseNoDirectContext _localctx = new RelationRecurseNoDirectContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_relationRecurseNoDirect);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(192);
			match(LPAREN);
			setState(196);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==WHITESPACE) {
				{
				{
				setState(193);
				match(WHITESPACE);
				}
				}
				setState(198);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(201);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				{
				setState(199);
				relationDefNoDirect();
				}
				break;
			case 2:
				{
				setState(200);
				relationRecurseNoDirect();
				}
				break;
			}
			setState(206);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==WHITESPACE) {
				{
				{
				setState(203);
				match(WHITESPACE);
				}
				}
				setState(208);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(209);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RelationDefDirectAssignmentContext extends ParserRuleContext {
		public TerminalNode LBRACKET() { return getToken(OpenFGAParser.LBRACKET, 0); }
		public List<RelationDefTypeRestrictionContext> relationDefTypeRestriction() {
			return getRuleContexts(RelationDefTypeRestrictionContext.class);
		}
		public RelationDefTypeRestrictionContext relationDefTypeRestriction(int i) {
			return getRuleContext(RelationDefTypeRestrictionContext.class,i);
		}
		public TerminalNode RPRACKET() { return getToken(OpenFGAParser.RPRACKET, 0); }
		public List<TerminalNode> WHITESPACE() { return getTokens(OpenFGAParser.WHITESPACE); }
		public TerminalNode WHITESPACE(int i) {
			return getToken(OpenFGAParser.WHITESPACE, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(OpenFGAParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(OpenFGAParser.COMMA, i);
		}
		public RelationDefDirectAssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationDefDirectAssignment; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterRelationDefDirectAssignment(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitRelationDefDirectAssignment(this);
		}
	}

	public final RelationDefDirectAssignmentContext relationDefDirectAssignment() throws RecognitionException {
		RelationDefDirectAssignmentContext _localctx = new RelationDefDirectAssignmentContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_relationDefDirectAssignment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(211);
			match(LBRACKET);
			setState(213);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(212);
				match(WHITESPACE);
				}
			}

			setState(215);
			relationDefTypeRestriction();
			setState(217);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(216);
				match(WHITESPACE);
				}
			}

			setState(229);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(219);
				match(COMMA);
				setState(221);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==WHITESPACE) {
					{
					setState(220);
					match(WHITESPACE);
					}
				}

				setState(223);
				relationDefTypeRestriction();
				setState(225);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==WHITESPACE) {
					{
					setState(224);
					match(WHITESPACE);
					}
				}

				}
				}
				setState(231);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(232);
			match(RPRACKET);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RelationDefRewriteContext extends ParserRuleContext {
		public Token rewriteComputedusersetName;
		public Token rewriteTuplesetName;
		public List<TerminalNode> IDENTIFIER() { return getTokens(OpenFGAParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(OpenFGAParser.IDENTIFIER, i);
		}
		public List<TerminalNode> WHITESPACE() { return getTokens(OpenFGAParser.WHITESPACE); }
		public TerminalNode WHITESPACE(int i) {
			return getToken(OpenFGAParser.WHITESPACE, i);
		}
		public TerminalNode FROM() { return getToken(OpenFGAParser.FROM, 0); }
		public RelationDefRewriteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationDefRewrite; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterRelationDefRewrite(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitRelationDefRewrite(this);
		}
	}

	public final RelationDefRewriteContext relationDefRewrite() throws RecognitionException {
		RelationDefRewriteContext _localctx = new RelationDefRewriteContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_relationDefRewrite);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(234);
			((RelationDefRewriteContext)_localctx).rewriteComputedusersetName = match(IDENTIFIER);
			setState(239);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,35,_ctx) ) {
			case 1:
				{
				setState(235);
				match(WHITESPACE);
				setState(236);
				match(FROM);
				setState(237);
				match(WHITESPACE);
				setState(238);
				((RelationDefRewriteContext)_localctx).rewriteTuplesetName = match(IDENTIFIER);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RelationDefTypeRestrictionContext extends ParserRuleContext {
		public RelationDefTypeRestrictionBaseContext relationDefTypeRestrictionBase() {
			return getRuleContext(RelationDefTypeRestrictionBaseContext.class,0);
		}
		public List<TerminalNode> NEWLINE() { return getTokens(OpenFGAParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(OpenFGAParser.NEWLINE, i);
		}
		public List<TerminalNode> WHITESPACE() { return getTokens(OpenFGAParser.WHITESPACE); }
		public TerminalNode WHITESPACE(int i) {
			return getToken(OpenFGAParser.WHITESPACE, i);
		}
		public TerminalNode KEYWORD_WITH() { return getToken(OpenFGAParser.KEYWORD_WITH, 0); }
		public ConditionNameContext conditionName() {
			return getRuleContext(ConditionNameContext.class,0);
		}
		public RelationDefTypeRestrictionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationDefTypeRestriction; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterRelationDefTypeRestriction(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitRelationDefTypeRestriction(this);
		}
	}

	public final RelationDefTypeRestrictionContext relationDefTypeRestriction() throws RecognitionException {
		RelationDefTypeRestrictionContext _localctx = new RelationDefTypeRestrictionContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_relationDefTypeRestriction);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(242);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NEWLINE) {
				{
				setState(241);
				match(NEWLINE);
				}
			}

			setState(251);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
			case 1:
				{
				setState(244);
				relationDefTypeRestrictionBase();
				}
				break;
			case 2:
				{
				{
				setState(245);
				relationDefTypeRestrictionBase();
				setState(246);
				match(WHITESPACE);
				setState(247);
				match(KEYWORD_WITH);
				setState(248);
				match(WHITESPACE);
				setState(249);
				conditionName();
				}
				}
				break;
			}
			setState(254);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NEWLINE) {
				{
				setState(253);
				match(NEWLINE);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RelationDefTypeRestrictionBaseContext extends ParserRuleContext {
		public Token relationDefTypeRestrictionType;
		public Token relationDefTypeRestrictionWildcard;
		public Token relationDefTypeRestrictionRelation;
		public List<TerminalNode> IDENTIFIER() { return getTokens(OpenFGAParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(OpenFGAParser.IDENTIFIER, i);
		}
		public TerminalNode COLON() { return getToken(OpenFGAParser.COLON, 0); }
		public TerminalNode HASH() { return getToken(OpenFGAParser.HASH, 0); }
		public TerminalNode STAR() { return getToken(OpenFGAParser.STAR, 0); }
		public RelationDefTypeRestrictionBaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationDefTypeRestrictionBase; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterRelationDefTypeRestrictionBase(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitRelationDefTypeRestrictionBase(this);
		}
	}

	public final RelationDefTypeRestrictionBaseContext relationDefTypeRestrictionBase() throws RecognitionException {
		RelationDefTypeRestrictionBaseContext _localctx = new RelationDefTypeRestrictionBaseContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_relationDefTypeRestrictionBase);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(256);
			((RelationDefTypeRestrictionBaseContext)_localctx).relationDefTypeRestrictionType = match(IDENTIFIER);
			setState(261);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case COLON:
				{
				{
				setState(257);
				match(COLON);
				setState(258);
				((RelationDefTypeRestrictionBaseContext)_localctx).relationDefTypeRestrictionWildcard = match(STAR);
				}
				}
				break;
			case HASH:
				{
				{
				setState(259);
				match(HASH);
				setState(260);
				((RelationDefTypeRestrictionBaseContext)_localctx).relationDefTypeRestrictionRelation = match(IDENTIFIER);
				}
				}
				break;
			case COMMA:
			case WHITESPACE:
			case RPRACKET:
			case NEWLINE:
				break;
			default:
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionsContext extends ParserRuleContext {
		public List<ConditionContext> condition() {
			return getRuleContexts(ConditionContext.class);
		}
		public ConditionContext condition(int i) {
			return getRuleContext(ConditionContext.class,i);
		}
		public ConditionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditions; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterConditions(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitConditions(this);
		}
	}

	public final ConditionsContext conditions() throws RecognitionException {
		ConditionsContext _localctx = new ConditionsContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_conditions);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(266);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,40,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(263);
					condition();
					}
					} 
				}
				setState(268);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,40,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionContext extends ParserRuleContext {
		public List<TerminalNode> NEWLINE() { return getTokens(OpenFGAParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(OpenFGAParser.NEWLINE, i);
		}
		public TerminalNode CONDITION() { return getToken(OpenFGAParser.CONDITION, 0); }
		public List<TerminalNode> WHITESPACE() { return getTokens(OpenFGAParser.WHITESPACE); }
		public TerminalNode WHITESPACE(int i) {
			return getToken(OpenFGAParser.WHITESPACE, i);
		}
		public ConditionNameContext conditionName() {
			return getRuleContext(ConditionNameContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(OpenFGAParser.LPAREN, 0); }
		public List<ConditionParameterContext> conditionParameter() {
			return getRuleContexts(ConditionParameterContext.class);
		}
		public ConditionParameterContext conditionParameter(int i) {
			return getRuleContext(ConditionParameterContext.class,i);
		}
		public TerminalNode RPAREN() { return getToken(OpenFGAParser.RPAREN, 0); }
		public TerminalNode LBRACE() { return getToken(OpenFGAParser.LBRACE, 0); }
		public ConditionExpressionContext conditionExpression() {
			return getRuleContext(ConditionExpressionContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(OpenFGAParser.RBRACE, 0); }
		public MultiLineCommentContext multiLineComment() {
			return getRuleContext(MultiLineCommentContext.class,0);
		}
		public List<TerminalNode> COMMA() { return getTokens(OpenFGAParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(OpenFGAParser.COMMA, i);
		}
		public ConditionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_condition; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterCondition(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitCondition(this);
		}
	}

	public final ConditionContext condition() throws RecognitionException {
		ConditionContext _localctx = new ConditionContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_condition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(271);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,41,_ctx) ) {
			case 1:
				{
				setState(269);
				match(NEWLINE);
				setState(270);
				multiLineComment();
				}
				break;
			}
			setState(273);
			match(NEWLINE);
			setState(274);
			match(CONDITION);
			setState(275);
			match(WHITESPACE);
			setState(276);
			conditionName();
			setState(278);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(277);
				match(WHITESPACE);
				}
			}

			setState(280);
			match(LPAREN);
			setState(282);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(281);
				match(WHITESPACE);
				}
			}

			setState(284);
			conditionParameter();
			setState(286);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(285);
				match(WHITESPACE);
				}
			}

			setState(298);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(288);
				match(COMMA);
				setState(290);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==WHITESPACE) {
					{
					setState(289);
					match(WHITESPACE);
					}
				}

				setState(292);
				conditionParameter();
				setState(294);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==WHITESPACE) {
					{
					setState(293);
					match(WHITESPACE);
					}
				}

				}
				}
				setState(300);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(302);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NEWLINE) {
				{
				setState(301);
				match(NEWLINE);
				}
			}

			setState(304);
			match(RPAREN);
			setState(306);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(305);
				match(WHITESPACE);
				}
			}

			setState(308);
			match(LBRACE);
			setState(310);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,50,_ctx) ) {
			case 1:
				{
				setState(309);
				match(NEWLINE);
				}
				break;
			}
			setState(313);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,51,_ctx) ) {
			case 1:
				{
				setState(312);
				match(WHITESPACE);
				}
				break;
			}
			setState(315);
			conditionExpression();
			setState(317);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NEWLINE) {
				{
				setState(316);
				match(NEWLINE);
				}
			}

			setState(319);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(OpenFGAParser.IDENTIFIER, 0); }
		public ConditionNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditionName; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterConditionName(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitConditionName(this);
		}
	}

	public final ConditionNameContext conditionName() throws RecognitionException {
		ConditionNameContext _localctx = new ConditionNameContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_conditionName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(321);
			match(IDENTIFIER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionParameterContext extends ParserRuleContext {
		public ParameterNameContext parameterName() {
			return getRuleContext(ParameterNameContext.class,0);
		}
		public TerminalNode COLON() { return getToken(OpenFGAParser.COLON, 0); }
		public ParameterTypeContext parameterType() {
			return getRuleContext(ParameterTypeContext.class,0);
		}
		public TerminalNode NEWLINE() { return getToken(OpenFGAParser.NEWLINE, 0); }
		public List<TerminalNode> WHITESPACE() { return getTokens(OpenFGAParser.WHITESPACE); }
		public TerminalNode WHITESPACE(int i) {
			return getToken(OpenFGAParser.WHITESPACE, i);
		}
		public ConditionParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditionParameter; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterConditionParameter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitConditionParameter(this);
		}
	}

	public final ConditionParameterContext conditionParameter() throws RecognitionException {
		ConditionParameterContext _localctx = new ConditionParameterContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_conditionParameter);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(324);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NEWLINE) {
				{
				setState(323);
				match(NEWLINE);
				}
			}

			setState(326);
			parameterName();
			setState(328);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(327);
				match(WHITESPACE);
				}
			}

			setState(330);
			match(COLON);
			setState(332);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(331);
				match(WHITESPACE);
				}
			}

			setState(334);
			parameterType();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParameterNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(OpenFGAParser.IDENTIFIER, 0); }
		public ParameterNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameterName; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterParameterName(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitParameterName(this);
		}
	}

	public final ParameterNameContext parameterName() throws RecognitionException {
		ParameterNameContext _localctx = new ParameterNameContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_parameterName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(336);
			match(IDENTIFIER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParameterTypeContext extends ParserRuleContext {
		public TerminalNode CONDITION_PARAM_TYPE() { return getToken(OpenFGAParser.CONDITION_PARAM_TYPE, 0); }
		public TerminalNode CONDITION_PARAM_CONTAINER() { return getToken(OpenFGAParser.CONDITION_PARAM_CONTAINER, 0); }
		public TerminalNode LESS() { return getToken(OpenFGAParser.LESS, 0); }
		public TerminalNode GREATER() { return getToken(OpenFGAParser.GREATER, 0); }
		public ParameterTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameterType; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterParameterType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitParameterType(this);
		}
	}

	public final ParameterTypeContext parameterType() throws RecognitionException {
		ParameterTypeContext _localctx = new ParameterTypeContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_parameterType);
		try {
			setState(343);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CONDITION_PARAM_TYPE:
				enterOuterAlt(_localctx, 1);
				{
				setState(338);
				match(CONDITION_PARAM_TYPE);
				}
				break;
			case CONDITION_PARAM_CONTAINER:
				enterOuterAlt(_localctx, 2);
				{
				{
				setState(339);
				match(CONDITION_PARAM_CONTAINER);
				setState(340);
				match(LESS);
				setState(341);
				match(CONDITION_PARAM_TYPE);
				setState(342);
				match(GREATER);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MultiLineCommentContext extends ParserRuleContext {
		public TerminalNode HASH() { return getToken(OpenFGAParser.HASH, 0); }
		public List<TerminalNode> NEWLINE() { return getTokens(OpenFGAParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(OpenFGAParser.NEWLINE, i);
		}
		public MultiLineCommentContext multiLineComment() {
			return getRuleContext(MultiLineCommentContext.class,0);
		}
		public MultiLineCommentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_multiLineComment; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterMultiLineComment(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitMultiLineComment(this);
		}
	}

	public final MultiLineCommentContext multiLineComment() throws RecognitionException {
		MultiLineCommentContext _localctx = new MultiLineCommentContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_multiLineComment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(345);
			match(HASH);
			setState(349);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 31525197391593470L) != 0)) {
				{
				{
				setState(346);
				_la = _input.LA(1);
				if ( _la <= 0 || (_la==NEWLINE) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				}
				setState(351);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(354);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,58,_ctx) ) {
			case 1:
				{
				setState(352);
				match(NEWLINE);
				setState(353);
				multiLineComment();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierContext extends ParserRuleContext {
		public TerminalNode MODEL() { return getToken(OpenFGAParser.MODEL, 0); }
		public TerminalNode SCHEMA() { return getToken(OpenFGAParser.SCHEMA, 0); }
		public TerminalNode TYPE() { return getToken(OpenFGAParser.TYPE, 0); }
		public TerminalNode RELATION() { return getToken(OpenFGAParser.RELATION, 0); }
		public TerminalNode IDENTIFIER() { return getToken(OpenFGAParser.IDENTIFIER, 0); }
		public IdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterIdentifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitIdentifier(this);
		}
	}

	public final IdentifierContext identifier() throws RecognitionException {
		IdentifierContext _localctx = new IdentifierContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_identifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(356);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 4916224L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionExpressionContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(OpenFGAParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(OpenFGAParser.IDENTIFIER, i);
		}
		public List<TerminalNode> EQUALS() { return getTokens(OpenFGAParser.EQUALS); }
		public TerminalNode EQUALS(int i) {
			return getToken(OpenFGAParser.EQUALS, i);
		}
		public List<TerminalNode> NOT_EQUALS() { return getTokens(OpenFGAParser.NOT_EQUALS); }
		public TerminalNode NOT_EQUALS(int i) {
			return getToken(OpenFGAParser.NOT_EQUALS, i);
		}
		public List<TerminalNode> IN() { return getTokens(OpenFGAParser.IN); }
		public TerminalNode IN(int i) {
			return getToken(OpenFGAParser.IN, i);
		}
		public List<TerminalNode> LESS() { return getTokens(OpenFGAParser.LESS); }
		public TerminalNode LESS(int i) {
			return getToken(OpenFGAParser.LESS, i);
		}
		public List<TerminalNode> LESS_EQUALS() { return getTokens(OpenFGAParser.LESS_EQUALS); }
		public TerminalNode LESS_EQUALS(int i) {
			return getToken(OpenFGAParser.LESS_EQUALS, i);
		}
		public List<TerminalNode> GREATER_EQUALS() { return getTokens(OpenFGAParser.GREATER_EQUALS); }
		public TerminalNode GREATER_EQUALS(int i) {
			return getToken(OpenFGAParser.GREATER_EQUALS, i);
		}
		public List<TerminalNode> GREATER() { return getTokens(OpenFGAParser.GREATER); }
		public TerminalNode GREATER(int i) {
			return getToken(OpenFGAParser.GREATER, i);
		}
		public List<TerminalNode> LOGICAL_AND() { return getTokens(OpenFGAParser.LOGICAL_AND); }
		public TerminalNode LOGICAL_AND(int i) {
			return getToken(OpenFGAParser.LOGICAL_AND, i);
		}
		public List<TerminalNode> LOGICAL_OR() { return getTokens(OpenFGAParser.LOGICAL_OR); }
		public TerminalNode LOGICAL_OR(int i) {
			return getToken(OpenFGAParser.LOGICAL_OR, i);
		}
		public List<TerminalNode> LBRACKET() { return getTokens(OpenFGAParser.LBRACKET); }
		public TerminalNode LBRACKET(int i) {
			return getToken(OpenFGAParser.LBRACKET, i);
		}
		public List<TerminalNode> RPRACKET() { return getTokens(OpenFGAParser.RPRACKET); }
		public TerminalNode RPRACKET(int i) {
			return getToken(OpenFGAParser.RPRACKET, i);
		}
		public List<TerminalNode> LBRACE() { return getTokens(OpenFGAParser.LBRACE); }
		public TerminalNode LBRACE(int i) {
			return getToken(OpenFGAParser.LBRACE, i);
		}
		public List<TerminalNode> LPAREN() { return getTokens(OpenFGAParser.LPAREN); }
		public TerminalNode LPAREN(int i) {
			return getToken(OpenFGAParser.LPAREN, i);
		}
		public List<TerminalNode> RPAREN() { return getTokens(OpenFGAParser.RPAREN); }
		public TerminalNode RPAREN(int i) {
			return getToken(OpenFGAParser.RPAREN, i);
		}
		public List<TerminalNode> DOT() { return getTokens(OpenFGAParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(OpenFGAParser.DOT, i);
		}
		public List<TerminalNode> MINUS() { return getTokens(OpenFGAParser.MINUS); }
		public TerminalNode MINUS(int i) {
			return getToken(OpenFGAParser.MINUS, i);
		}
		public List<TerminalNode> EXCLAM() { return getTokens(OpenFGAParser.EXCLAM); }
		public TerminalNode EXCLAM(int i) {
			return getToken(OpenFGAParser.EXCLAM, i);
		}
		public List<TerminalNode> QUESTIONMARK() { return getTokens(OpenFGAParser.QUESTIONMARK); }
		public TerminalNode QUESTIONMARK(int i) {
			return getToken(OpenFGAParser.QUESTIONMARK, i);
		}
		public List<TerminalNode> PLUS() { return getTokens(OpenFGAParser.PLUS); }
		public TerminalNode PLUS(int i) {
			return getToken(OpenFGAParser.PLUS, i);
		}
		public List<TerminalNode> STAR() { return getTokens(OpenFGAParser.STAR); }
		public TerminalNode STAR(int i) {
			return getToken(OpenFGAParser.STAR, i);
		}
		public List<TerminalNode> SLASH() { return getTokens(OpenFGAParser.SLASH); }
		public TerminalNode SLASH(int i) {
			return getToken(OpenFGAParser.SLASH, i);
		}
		public List<TerminalNode> PERCENT() { return getTokens(OpenFGAParser.PERCENT); }
		public TerminalNode PERCENT(int i) {
			return getToken(OpenFGAParser.PERCENT, i);
		}
		public List<TerminalNode> CEL_TRUE() { return getTokens(OpenFGAParser.CEL_TRUE); }
		public TerminalNode CEL_TRUE(int i) {
			return getToken(OpenFGAParser.CEL_TRUE, i);
		}
		public List<TerminalNode> CEL_FALSE() { return getTokens(OpenFGAParser.CEL_FALSE); }
		public TerminalNode CEL_FALSE(int i) {
			return getToken(OpenFGAParser.CEL_FALSE, i);
		}
		public List<TerminalNode> NUL() { return getTokens(OpenFGAParser.NUL); }
		public TerminalNode NUL(int i) {
			return getToken(OpenFGAParser.NUL, i);
		}
		public List<TerminalNode> WHITESPACE() { return getTokens(OpenFGAParser.WHITESPACE); }
		public TerminalNode WHITESPACE(int i) {
			return getToken(OpenFGAParser.WHITESPACE, i);
		}
		public List<TerminalNode> CEL_COMMENT() { return getTokens(OpenFGAParser.CEL_COMMENT); }
		public TerminalNode CEL_COMMENT(int i) {
			return getToken(OpenFGAParser.CEL_COMMENT, i);
		}
		public List<TerminalNode> NUM_FLOAT() { return getTokens(OpenFGAParser.NUM_FLOAT); }
		public TerminalNode NUM_FLOAT(int i) {
			return getToken(OpenFGAParser.NUM_FLOAT, i);
		}
		public List<TerminalNode> NUM_INT() { return getTokens(OpenFGAParser.NUM_INT); }
		public TerminalNode NUM_INT(int i) {
			return getToken(OpenFGAParser.NUM_INT, i);
		}
		public List<TerminalNode> NUM_UINT() { return getTokens(OpenFGAParser.NUM_UINT); }
		public TerminalNode NUM_UINT(int i) {
			return getToken(OpenFGAParser.NUM_UINT, i);
		}
		public List<TerminalNode> STRING() { return getTokens(OpenFGAParser.STRING); }
		public TerminalNode STRING(int i) {
			return getToken(OpenFGAParser.STRING, i);
		}
		public List<TerminalNode> BYTES() { return getTokens(OpenFGAParser.BYTES); }
		public TerminalNode BYTES(int i) {
			return getToken(OpenFGAParser.BYTES, i);
		}
		public List<TerminalNode> NEWLINE() { return getTokens(OpenFGAParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(OpenFGAParser.NEWLINE, i);
		}
		public List<TerminalNode> RBRACE() { return getTokens(OpenFGAParser.RBRACE); }
		public TerminalNode RBRACE(int i) {
			return getToken(OpenFGAParser.RBRACE, i);
		}
		public ConditionExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditionExpression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterConditionExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitConditionExpression(this);
		}
	}

	public final ConditionExpressionContext conditionExpression() throws RecognitionException {
		ConditionExpressionContext _localctx = new ConditionExpressionContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_conditionExpression);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(362);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,60,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					setState(360);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,59,_ctx) ) {
					case 1:
						{
						setState(358);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 9007182041319352L) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						}
						break;
					case 2:
						{
						setState(359);
						_la = _input.LA(1);
						if ( _la <= 0 || (_la==RBRACE) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						}
						break;
					}
					} 
				}
				setState(364);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,60,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u00016\u016e\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0001\u0000\u0003\u00004\b\u0000\u0001\u0000\u0003\u00007\b\u0000\u0001"+
		"\u0000\u0001\u0000\u0003\u0000;\b\u0000\u0001\u0000\u0001\u0000\u0003"+
		"\u0000?\b\u0000\u0001\u0000\u0001\u0000\u0003\u0000C\b\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001J\b\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0003\u0001R\b\u0001\u0001\u0002\u0005\u0002U\b\u0002\n\u0002\f\u0002"+
		"X\t\u0002\u0001\u0003\u0001\u0003\u0003\u0003\\\b\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0004"+
		"\u0003e\b\u0003\u000b\u0003\f\u0003f\u0003\u0003i\b\u0003\u0001\u0004"+
		"\u0001\u0004\u0003\u0004m\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0003\u0004t\b\u0004\u0001\u0004\u0001\u0004"+
		"\u0003\u0004x\b\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u0081\b\u0006\u0001\u0006"+
		"\u0003\u0006\u0084\b\u0006\u0001\u0007\u0001\u0007\u0003\u0007\u0088\b"+
		"\u0007\u0001\u0007\u0003\u0007\u008b\b\u0007\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0003\b\u0092\b\b\u0004\b\u0094\b\b\u000b\b\f\b\u0095\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003\b\u009d\b\b\u0004\b\u009f\b\b"+
		"\u000b\b\f\b\u00a0\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003\b\u00a8"+
		"\b\b\u0003\b\u00aa\b\b\u0001\t\u0001\t\u0001\n\u0001\n\u0005\n\u00b0\b"+
		"\n\n\n\f\n\u00b3\t\n\u0001\n\u0001\n\u0003\n\u00b7\b\n\u0001\n\u0005\n"+
		"\u00ba\b\n\n\n\f\n\u00bd\t\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0005"+
		"\u000b\u00c3\b\u000b\n\u000b\f\u000b\u00c6\t\u000b\u0001\u000b\u0001\u000b"+
		"\u0003\u000b\u00ca\b\u000b\u0001\u000b\u0005\u000b\u00cd\b\u000b\n\u000b"+
		"\f\u000b\u00d0\t\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0003\f"+
		"\u00d6\b\f\u0001\f\u0001\f\u0003\f\u00da\b\f\u0001\f\u0001\f\u0003\f\u00de"+
		"\b\f\u0001\f\u0001\f\u0003\f\u00e2\b\f\u0005\f\u00e4\b\f\n\f\f\f\u00e7"+
		"\t\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0003\r\u00f0"+
		"\b\r\u0001\u000e\u0003\u000e\u00f3\b\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u00fc"+
		"\b\u000e\u0001\u000e\u0003\u000e\u00ff\b\u000e\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u0106\b\u000f\u0001\u0010"+
		"\u0005\u0010\u0109\b\u0010\n\u0010\f\u0010\u010c\t\u0010\u0001\u0011\u0001"+
		"\u0011\u0003\u0011\u0110\b\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0003\u0011\u0117\b\u0011\u0001\u0011\u0001\u0011\u0003"+
		"\u0011\u011b\b\u0011\u0001\u0011\u0001\u0011\u0003\u0011\u011f\b\u0011"+
		"\u0001\u0011\u0001\u0011\u0003\u0011\u0123\b\u0011\u0001\u0011\u0001\u0011"+
		"\u0003\u0011\u0127\b\u0011\u0005\u0011\u0129\b\u0011\n\u0011\f\u0011\u012c"+
		"\t\u0011\u0001\u0011\u0003\u0011\u012f\b\u0011\u0001\u0011\u0001\u0011"+
		"\u0003\u0011\u0133\b\u0011\u0001\u0011\u0001\u0011\u0003\u0011\u0137\b"+
		"\u0011\u0001\u0011\u0003\u0011\u013a\b\u0011\u0001\u0011\u0001\u0011\u0003"+
		"\u0011\u013e\b\u0011\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001"+
		"\u0013\u0003\u0013\u0145\b\u0013\u0001\u0013\u0001\u0013\u0003\u0013\u0149"+
		"\b\u0013\u0001\u0013\u0001\u0013\u0003\u0013\u014d\b\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0003\u0015\u0158\b\u0015\u0001\u0016\u0001\u0016"+
		"\u0005\u0016\u015c\b\u0016\n\u0016\f\u0016\u015f\t\u0016\u0001\u0016\u0001"+
		"\u0016\u0003\u0016\u0163\b\u0016\u0001\u0017\u0001\u0017\u0001\u0018\u0001"+
		"\u0018\u0005\u0018\u0169\b\u0018\n\u0018\f\u0018\u016c\t\u0018\u0001\u0018"+
		"\u0000\u0000\u0019\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014"+
		"\u0016\u0018\u001a\u001c\u001e \"$&(*,.0\u0000\u0004\u0001\u000044\u0004"+
		"\u0000\n\n\u0010\u0011\u0013\u0013\u0016\u0016\u0004\u0000\u0003\u0005"+
		"\u0007\n\u0019!#4\u0001\u0000\"\"\u0194\u00003\u0001\u0000\u0000\u0000"+
		"\u0002I\u0001\u0000\u0000\u0000\u0004V\u0001\u0000\u0000\u0000\u0006["+
		"\u0001\u0000\u0000\u0000\bl\u0001\u0000\u0000\u0000\n{\u0001\u0000\u0000"+
		"\u0000\f\u0080\u0001\u0000\u0000\u0000\u000e\u0087\u0001\u0000\u0000\u0000"+
		"\u0010\u00a9\u0001\u0000\u0000\u0000\u0012\u00ab\u0001\u0000\u0000\u0000"+
		"\u0014\u00ad\u0001\u0000\u0000\u0000\u0016\u00c0\u0001\u0000\u0000\u0000"+
		"\u0018\u00d3\u0001\u0000\u0000\u0000\u001a\u00ea\u0001\u0000\u0000\u0000"+
		"\u001c\u00f2\u0001\u0000\u0000\u0000\u001e\u0100\u0001\u0000\u0000\u0000"+
		" \u010a\u0001\u0000\u0000\u0000\"\u010f\u0001\u0000\u0000\u0000$\u0141"+
		"\u0001\u0000\u0000\u0000&\u0144\u0001\u0000\u0000\u0000(\u0150\u0001\u0000"+
		"\u0000\u0000*\u0157\u0001\u0000\u0000\u0000,\u0159\u0001\u0000\u0000\u0000"+
		".\u0164\u0001\u0000\u0000\u00000\u016a\u0001\u0000\u0000\u000024\u0005"+
		"\t\u0000\u000032\u0001\u0000\u0000\u000034\u0001\u0000\u0000\u000046\u0001"+
		"\u0000\u0000\u000057\u00054\u0000\u000065\u0001\u0000\u0000\u000067\u0001"+
		"\u0000\u0000\u000078\u0001\u0000\u0000\u00008:\u0003\u0002\u0001\u0000"+
		"9;\u00054\u0000\u0000:9\u0001\u0000\u0000\u0000:;\u0001\u0000\u0000\u0000"+
		";<\u0001\u0000\u0000\u0000<>\u0003\u0004\u0002\u0000=?\u00054\u0000\u0000"+
		">=\u0001\u0000\u0000\u0000>?\u0001\u0000\u0000\u0000?@\u0001\u0000\u0000"+
		"\u0000@B\u0003 \u0010\u0000AC\u00054\u0000\u0000BA\u0001\u0000\u0000\u0000"+
		"BC\u0001\u0000\u0000\u0000CD\u0001\u0000\u0000\u0000DE\u0005\u0000\u0000"+
		"\u0001E\u0001\u0001\u0000\u0000\u0000FG\u0003,\u0016\u0000GH\u00054\u0000"+
		"\u0000HJ\u0001\u0000\u0000\u0000IF\u0001\u0000\u0000\u0000IJ\u0001\u0000"+
		"\u0000\u0000JK\u0001\u0000\u0000\u0000KL\u0005\u0010\u0000\u0000LM\u0005"+
		"4\u0000\u0000MN\u0005\u0011\u0000\u0000NO\u0005\t\u0000\u0000OQ\u0005"+
		"\u0012\u0000\u0000PR\u0005\t\u0000\u0000QP\u0001\u0000\u0000\u0000QR\u0001"+
		"\u0000\u0000\u0000R\u0003\u0001\u0000\u0000\u0000SU\u0003\u0006\u0003"+
		"\u0000TS\u0001\u0000\u0000\u0000UX\u0001\u0000\u0000\u0000VT\u0001\u0000"+
		"\u0000\u0000VW\u0001\u0000\u0000\u0000W\u0005\u0001\u0000\u0000\u0000"+
		"XV\u0001\u0000\u0000\u0000YZ\u00054\u0000\u0000Z\\\u0003,\u0016\u0000"+
		"[Y\u0001\u0000\u0000\u0000[\\\u0001\u0000\u0000\u0000\\]\u0001\u0000\u0000"+
		"\u0000]^\u00054\u0000\u0000^_\u0005\u0013\u0000\u0000_`\u0005\t\u0000"+
		"\u0000`h\u0003.\u0017\u0000ab\u00054\u0000\u0000bd\u0005\u0015\u0000\u0000"+
		"ce\u0003\b\u0004\u0000dc\u0001\u0000\u0000\u0000ef\u0001\u0000\u0000\u0000"+
		"fd\u0001\u0000\u0000\u0000fg\u0001\u0000\u0000\u0000gi\u0001\u0000\u0000"+
		"\u0000ha\u0001\u0000\u0000\u0000hi\u0001\u0000\u0000\u0000i\u0007\u0001"+
		"\u0000\u0000\u0000jk\u00054\u0000\u0000km\u0003,\u0016\u0000lj\u0001\u0000"+
		"\u0000\u0000lm\u0001\u0000\u0000\u0000mn\u0001\u0000\u0000\u0000no\u0005"+
		"4\u0000\u0000op\u0005\u0017\u0000\u0000pq\u0005\t\u0000\u0000qs\u0003"+
		"\n\u0005\u0000rt\u0005\t\u0000\u0000sr\u0001\u0000\u0000\u0000st\u0001"+
		"\u0000\u0000\u0000tu\u0001\u0000\u0000\u0000uw\u0005\u0001\u0000\u0000"+
		"vx\u0005\t\u0000\u0000wv\u0001\u0000\u0000\u0000wx\u0001\u0000\u0000\u0000"+
		"xy\u0001\u0000\u0000\u0000yz\u0003\f\u0006\u0000z\t\u0001\u0000\u0000"+
		"\u0000{|\u0003.\u0017\u0000|\u000b\u0001\u0000\u0000\u0000}\u0081\u0003"+
		"\u0018\f\u0000~\u0081\u0003\u0012\t\u0000\u007f\u0081\u0003\u0014\n\u0000"+
		"\u0080}\u0001\u0000\u0000\u0000\u0080~\u0001\u0000\u0000\u0000\u0080\u007f"+
		"\u0001\u0000\u0000\u0000\u0081\u0083\u0001\u0000\u0000\u0000\u0082\u0084"+
		"\u0003\u0010\b\u0000\u0083\u0082\u0001\u0000\u0000\u0000\u0083\u0084\u0001"+
		"\u0000\u0000\u0000\u0084\r\u0001\u0000\u0000\u0000\u0085\u0088\u0003\u0012"+
		"\t\u0000\u0086\u0088\u0003\u0016\u000b\u0000\u0087\u0085\u0001\u0000\u0000"+
		"\u0000\u0087\u0086\u0001\u0000\u0000\u0000\u0088\u008a\u0001\u0000\u0000"+
		"\u0000\u0089\u008b\u0003\u0010\b\u0000\u008a\u0089\u0001\u0000\u0000\u0000"+
		"\u008a\u008b\u0001\u0000\u0000\u0000\u008b\u000f\u0001\u0000\u0000\u0000"+
		"\u008c\u008d\u0005\t\u0000\u0000\u008d\u008e\u0005\r\u0000\u0000\u008e"+
		"\u0091\u0005\t\u0000\u0000\u008f\u0092\u0003\u0012\t\u0000\u0090\u0092"+
		"\u0003\u0016\u000b\u0000\u0091\u008f\u0001\u0000\u0000\u0000\u0091\u0090"+
		"\u0001\u0000\u0000\u0000\u0092\u0094\u0001\u0000\u0000\u0000\u0093\u008c"+
		"\u0001\u0000\u0000\u0000\u0094\u0095\u0001\u0000\u0000\u0000\u0095\u0093"+
		"\u0001\u0000\u0000\u0000\u0095\u0096\u0001\u0000\u0000\u0000\u0096\u00aa"+
		"\u0001\u0000\u0000\u0000\u0097\u0098\u0005\t\u0000\u0000\u0098\u0099\u0005"+
		"\f\u0000\u0000\u0099\u009c\u0005\t\u0000\u0000\u009a\u009d\u0003\u0012"+
		"\t\u0000\u009b\u009d\u0003\u0016\u000b\u0000\u009c\u009a\u0001\u0000\u0000"+
		"\u0000\u009c\u009b\u0001\u0000\u0000\u0000\u009d\u009f\u0001\u0000\u0000"+
		"\u0000\u009e\u0097\u0001\u0000\u0000\u0000\u009f\u00a0\u0001\u0000\u0000"+
		"\u0000\u00a0\u009e\u0001\u0000\u0000\u0000\u00a0\u00a1\u0001\u0000\u0000"+
		"\u0000\u00a1\u00aa\u0001\u0000\u0000\u0000\u00a2\u00a3\u0005\t\u0000\u0000"+
		"\u00a3\u00a4\u0005\u000e\u0000\u0000\u00a4\u00a7\u0005\t\u0000\u0000\u00a5"+
		"\u00a8\u0003\u0012\t\u0000\u00a6\u00a8\u0003\u0016\u000b\u0000\u00a7\u00a5"+
		"\u0001\u0000\u0000\u0000\u00a7\u00a6\u0001\u0000\u0000\u0000\u00a8\u00aa"+
		"\u0001\u0000\u0000\u0000\u00a9\u0093\u0001\u0000\u0000\u0000\u00a9\u009e"+
		"\u0001\u0000\u0000\u0000\u00a9\u00a2\u0001\u0000\u0000\u0000\u00aa\u0011"+
		"\u0001\u0000\u0000\u0000\u00ab\u00ac\u0003\u001a\r\u0000\u00ac\u0013\u0001"+
		"\u0000\u0000\u0000\u00ad\u00b1\u0005\u0007\u0000\u0000\u00ae\u00b0\u0005"+
		"\t\u0000\u0000\u00af\u00ae\u0001\u0000\u0000\u0000\u00b0\u00b3\u0001\u0000"+
		"\u0000\u0000\u00b1\u00af\u0001\u0000\u0000\u0000\u00b1\u00b2\u0001\u0000"+
		"\u0000\u0000\u00b2\u00b6\u0001\u0000\u0000\u0000\u00b3\u00b1\u0001\u0000"+
		"\u0000\u0000\u00b4\u00b7\u0003\f\u0006\u0000\u00b5\u00b7\u0003\u0016\u000b"+
		"\u0000\u00b6\u00b4\u0001\u0000\u0000\u0000\u00b6\u00b5\u0001\u0000\u0000"+
		"\u0000\u00b7\u00bb\u0001\u0000\u0000\u0000\u00b8\u00ba\u0005\t\u0000\u0000"+
		"\u00b9\u00b8\u0001\u0000\u0000\u0000\u00ba\u00bd\u0001\u0000\u0000\u0000"+
		"\u00bb\u00b9\u0001\u0000\u0000\u0000\u00bb\u00bc\u0001\u0000\u0000\u0000"+
		"\u00bc\u00be\u0001\u0000\u0000\u0000\u00bd\u00bb\u0001\u0000\u0000\u0000"+
		"\u00be\u00bf\u0005\b\u0000\u0000\u00bf\u0015\u0001\u0000\u0000\u0000\u00c0"+
		"\u00c4\u0005\u0007\u0000\u0000\u00c1\u00c3\u0005\t\u0000\u0000\u00c2\u00c1"+
		"\u0001\u0000\u0000\u0000\u00c3\u00c6\u0001\u0000\u0000\u0000\u00c4\u00c2"+
		"\u0001\u0000\u0000\u0000\u00c4\u00c5\u0001\u0000\u0000\u0000\u00c5\u00c9"+
		"\u0001\u0000\u0000\u0000\u00c6\u00c4\u0001\u0000\u0000\u0000\u00c7\u00ca"+
		"\u0003\u000e\u0007\u0000\u00c8\u00ca\u0003\u0016\u000b\u0000\u00c9\u00c7"+
		"\u0001\u0000\u0000\u0000\u00c9\u00c8\u0001\u0000\u0000\u0000\u00ca\u00ce"+
		"\u0001\u0000\u0000\u0000\u00cb\u00cd\u0005\t\u0000\u0000\u00cc\u00cb\u0001"+
		"\u0000\u0000\u0000\u00cd\u00d0\u0001\u0000\u0000\u0000\u00ce\u00cc\u0001"+
		"\u0000\u0000\u0000\u00ce\u00cf\u0001\u0000\u0000\u0000\u00cf\u00d1\u0001"+
		"\u0000\u0000\u0000\u00d0\u00ce\u0001\u0000\u0000\u0000\u00d1\u00d2\u0005"+
		"\b\u0000\u0000\u00d2\u0017\u0001\u0000\u0000\u0000\u00d3\u00d5\u0005\u0005"+
		"\u0000\u0000\u00d4\u00d6\u0005\t\u0000\u0000\u00d5\u00d4\u0001\u0000\u0000"+
		"\u0000\u00d5\u00d6\u0001\u0000\u0000\u0000\u00d6\u00d7\u0001\u0000\u0000"+
		"\u0000\u00d7\u00d9\u0003\u001c\u000e\u0000\u00d8\u00da\u0005\t\u0000\u0000"+
		"\u00d9\u00d8\u0001\u0000\u0000\u0000\u00d9\u00da\u0001\u0000\u0000\u0000"+
		"\u00da\u00e5\u0001\u0000\u0000\u0000\u00db\u00dd\u0005\u0002\u0000\u0000"+
		"\u00dc\u00de\u0005\t\u0000\u0000\u00dd\u00dc\u0001\u0000\u0000\u0000\u00dd"+
		"\u00de\u0001\u0000\u0000\u0000\u00de\u00df\u0001\u0000\u0000\u0000\u00df"+
		"\u00e1\u0003\u001c\u000e\u0000\u00e0\u00e2\u0005\t\u0000\u0000\u00e1\u00e0"+
		"\u0001\u0000\u0000\u0000\u00e1\u00e2\u0001\u0000\u0000\u0000\u00e2\u00e4"+
		"\u0001\u0000\u0000\u0000\u00e3\u00db\u0001\u0000\u0000\u0000\u00e4\u00e7"+
		"\u0001\u0000\u0000\u0000\u00e5\u00e3\u0001\u0000\u0000\u0000\u00e5\u00e6"+
		"\u0001\u0000\u0000\u0000\u00e6\u00e8\u0001\u0000\u0000\u0000\u00e7\u00e5"+
		"\u0001\u0000\u0000\u0000\u00e8\u00e9\u0005 \u0000\u0000\u00e9\u0019\u0001"+
		"\u0000\u0000\u0000\u00ea\u00ef\u0005\n\u0000\u0000\u00eb\u00ec\u0005\t"+
		"\u0000\u0000\u00ec\u00ed\u0005\u000f\u0000\u0000\u00ed\u00ee\u0005\t\u0000"+
		"\u0000\u00ee\u00f0\u0005\n\u0000\u0000\u00ef\u00eb\u0001\u0000\u0000\u0000"+
		"\u00ef\u00f0\u0001\u0000\u0000\u0000\u00f0\u001b\u0001\u0000\u0000\u0000"+
		"\u00f1\u00f3\u00054\u0000\u0000\u00f2\u00f1\u0001\u0000\u0000\u0000\u00f2"+
		"\u00f3\u0001\u0000\u0000\u0000\u00f3\u00fb\u0001\u0000\u0000\u0000\u00f4"+
		"\u00fc\u0003\u001e\u000f\u0000\u00f5\u00f6\u0003\u001e\u000f\u0000\u00f6"+
		"\u00f7\u0005\t\u0000\u0000\u00f7\u00f8\u0005\u0018\u0000\u0000\u00f8\u00f9"+
		"\u0005\t\u0000\u0000\u00f9\u00fa\u0003$\u0012\u0000\u00fa\u00fc\u0001"+
		"\u0000\u0000\u0000\u00fb\u00f4\u0001\u0000\u0000\u0000\u00fb\u00f5\u0001"+
		"\u0000\u0000\u0000\u00fc\u00fe\u0001\u0000\u0000\u0000\u00fd\u00ff\u0005"+
		"4\u0000\u0000\u00fe\u00fd\u0001\u0000\u0000\u0000\u00fe\u00ff\u0001\u0000"+
		"\u0000\u0000\u00ff\u001d\u0001\u0000\u0000\u0000\u0100\u0105\u0005\n\u0000"+
		"\u0000\u0101\u0102\u0005\u0001\u0000\u0000\u0102\u0106\u0005(\u0000\u0000"+
		"\u0103\u0104\u0005\u000b\u0000\u0000\u0104\u0106\u0005\n\u0000\u0000\u0105"+
		"\u0101\u0001\u0000\u0000\u0000\u0105\u0103\u0001\u0000\u0000\u0000\u0105"+
		"\u0106\u0001\u0000\u0000\u0000\u0106\u001f\u0001\u0000\u0000\u0000\u0107"+
		"\u0109\u0003\"\u0011\u0000\u0108\u0107\u0001\u0000\u0000\u0000\u0109\u010c"+
		"\u0001\u0000\u0000\u0000\u010a\u0108\u0001\u0000\u0000\u0000\u010a\u010b"+
		"\u0001\u0000\u0000\u0000\u010b!\u0001\u0000\u0000\u0000\u010c\u010a\u0001"+
		"\u0000\u0000\u0000\u010d\u010e\u00054\u0000\u0000\u010e\u0110\u0003,\u0016"+
		"\u0000\u010f\u010d\u0001\u0000\u0000\u0000\u010f\u0110\u0001\u0000\u0000"+
		"\u0000\u0110\u0111\u0001\u0000\u0000\u0000\u0111\u0112\u00054\u0000\u0000"+
		"\u0112\u0113\u0005\u0014\u0000\u0000\u0113\u0114\u0005\t\u0000\u0000\u0114"+
		"\u0116\u0003$\u0012\u0000\u0115\u0117\u0005\t\u0000\u0000\u0116\u0115"+
		"\u0001\u0000\u0000\u0000\u0116\u0117\u0001\u0000\u0000\u0000\u0117\u0118"+
		"\u0001\u0000\u0000\u0000\u0118\u011a\u0005\u0007\u0000\u0000\u0119\u011b"+
		"\u0005\t\u0000\u0000\u011a\u0119\u0001\u0000\u0000\u0000\u011a\u011b\u0001"+
		"\u0000\u0000\u0000\u011b\u011c\u0001\u0000\u0000\u0000\u011c\u011e\u0003"+
		"&\u0013\u0000\u011d\u011f\u0005\t\u0000\u0000\u011e\u011d\u0001\u0000"+
		"\u0000\u0000\u011e\u011f\u0001\u0000\u0000\u0000\u011f\u012a\u0001\u0000"+
		"\u0000\u0000\u0120\u0122\u0005\u0002\u0000\u0000\u0121\u0123\u0005\t\u0000"+
		"\u0000\u0122\u0121\u0001\u0000\u0000\u0000\u0122\u0123\u0001\u0000\u0000"+
		"\u0000\u0123\u0124\u0001\u0000\u0000\u0000\u0124\u0126\u0003&\u0013\u0000"+
		"\u0125\u0127\u0005\t\u0000\u0000\u0126\u0125\u0001\u0000\u0000\u0000\u0126"+
		"\u0127\u0001\u0000\u0000\u0000\u0127\u0129\u0001\u0000\u0000\u0000\u0128"+
		"\u0120\u0001\u0000\u0000\u0000\u0129\u012c\u0001\u0000\u0000\u0000\u012a"+
		"\u0128\u0001\u0000\u0000\u0000\u012a\u012b\u0001\u0000\u0000\u0000\u012b"+
		"\u012e\u0001\u0000\u0000\u0000\u012c\u012a\u0001\u0000\u0000\u0000\u012d"+
		"\u012f\u00054\u0000\u0000\u012e\u012d\u0001\u0000\u0000\u0000\u012e\u012f"+
		"\u0001\u0000\u0000\u0000\u012f\u0130\u0001\u0000\u0000\u0000\u0130\u0132"+
		"\u0005\b\u0000\u0000\u0131\u0133\u0005\t\u0000\u0000\u0132\u0131\u0001"+
		"\u0000\u0000\u0000\u0132\u0133\u0001\u0000\u0000\u0000\u0133\u0134\u0001"+
		"\u0000\u0000\u0000\u0134\u0136\u0005!\u0000\u0000\u0135\u0137\u00054\u0000"+
		"\u0000\u0136\u0135\u0001\u0000\u0000\u0000\u0136\u0137\u0001\u0000\u0000"+
		"\u0000\u0137\u0139\u0001\u0000\u0000\u0000\u0138\u013a\u0005\t\u0000\u0000"+
		"\u0139\u0138\u0001\u0000\u0000\u0000\u0139\u013a\u0001\u0000\u0000\u0000"+
		"\u013a\u013b\u0001\u0000\u0000\u0000\u013b\u013d\u00030\u0018\u0000\u013c"+
		"\u013e\u00054\u0000\u0000\u013d\u013c\u0001\u0000\u0000\u0000\u013d\u013e"+
		"\u0001\u0000\u0000\u0000\u013e\u013f\u0001\u0000\u0000\u0000\u013f\u0140"+
		"\u0005\"\u0000\u0000\u0140#\u0001\u0000\u0000\u0000\u0141\u0142\u0005"+
		"\n\u0000\u0000\u0142%\u0001\u0000\u0000\u0000\u0143\u0145\u00054\u0000"+
		"\u0000\u0144\u0143\u0001\u0000\u0000\u0000\u0144\u0145\u0001\u0000\u0000"+
		"\u0000\u0145\u0146\u0001\u0000\u0000\u0000\u0146\u0148\u0003(\u0014\u0000"+
		"\u0147\u0149\u0005\t\u0000\u0000\u0148\u0147\u0001\u0000\u0000\u0000\u0148"+
		"\u0149\u0001\u0000\u0000\u0000\u0149\u014a\u0001\u0000\u0000\u0000\u014a"+
		"\u014c\u0005\u0001\u0000\u0000\u014b\u014d\u0005\t\u0000\u0000\u014c\u014b"+
		"\u0001\u0000\u0000\u0000\u014c\u014d\u0001\u0000\u0000\u0000\u014d\u014e"+
		"\u0001\u0000\u0000\u0000\u014e\u014f\u0003*\u0015\u0000\u014f\'\u0001"+
		"\u0000\u0000\u0000\u0150\u0151\u0005\n\u0000\u0000\u0151)\u0001\u0000"+
		"\u0000\u0000\u0152\u0158\u00056\u0000\u0000\u0153\u0154\u00055\u0000\u0000"+
		"\u0154\u0155\u0005\u0003\u0000\u0000\u0155\u0156\u00056\u0000\u0000\u0156"+
		"\u0158\u0005\u0004\u0000\u0000\u0157\u0152\u0001\u0000\u0000\u0000\u0157"+
		"\u0153\u0001\u0000\u0000\u0000\u0158+\u0001\u0000\u0000\u0000\u0159\u015d"+
		"\u0005\u000b\u0000\u0000\u015a\u015c\b\u0000\u0000\u0000\u015b\u015a\u0001"+
		"\u0000\u0000\u0000\u015c\u015f\u0001\u0000\u0000\u0000\u015d\u015b\u0001"+
		"\u0000\u0000\u0000\u015d\u015e\u0001\u0000\u0000\u0000\u015e\u0162\u0001"+
		"\u0000\u0000\u0000\u015f\u015d\u0001\u0000\u0000\u0000\u0160\u0161\u0005"+
		"4\u0000\u0000\u0161\u0163\u0003,\u0016\u0000\u0162\u0160\u0001\u0000\u0000"+
		"\u0000\u0162\u0163\u0001\u0000\u0000\u0000\u0163-\u0001\u0000\u0000\u0000"+
		"\u0164\u0165\u0007\u0001\u0000\u0000\u0165/\u0001\u0000\u0000\u0000\u0166"+
		"\u0169\u0007\u0002\u0000\u0000\u0167\u0169\b\u0003\u0000\u0000\u0168\u0166"+
		"\u0001\u0000\u0000\u0000\u0168\u0167\u0001\u0000\u0000\u0000\u0169\u016c"+
		"\u0001\u0000\u0000\u0000\u016a\u0168\u0001\u0000\u0000\u0000\u016a\u016b"+
		"\u0001\u0000\u0000\u0000\u016b1\u0001\u0000\u0000\u0000\u016c\u016a\u0001"+
		"\u0000\u0000\u0000=36:>BIQV[fhlsw\u0080\u0083\u0087\u008a\u0091\u0095"+
		"\u009c\u00a0\u00a7\u00a9\u00b1\u00b6\u00bb\u00c4\u00c9\u00ce\u00d5\u00d9"+
		"\u00dd\u00e1\u00e5\u00ef\u00f2\u00fb\u00fe\u0105\u010a\u010f\u0116\u011a"+
		"\u011e\u0122\u0126\u012a\u012e\u0132\u0136\u0139\u013d\u0144\u0148\u014c"+
		"\u0157\u015d\u0162\u0168\u016a";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}