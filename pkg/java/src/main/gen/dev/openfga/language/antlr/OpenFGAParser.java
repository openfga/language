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
		FROM=15, MODULE=16, MODEL=17, SCHEMA=18, SCHEMA_VERSION=19, EXTEND=20, 
		TYPE=21, CONDITION=22, RELATIONS=23, RELATION=24, DEFINE=25, KEYWORD_WITH=26, 
		EQUALS=27, NOT_EQUALS=28, IN=29, LESS_EQUALS=30, GREATER_EQUALS=31, LOGICAL_AND=32, 
		LOGICAL_OR=33, RPRACKET=34, LBRACE=35, RBRACE=36, DOT=37, MINUS=38, EXCLAM=39, 
		QUESTIONMARK=40, PLUS=41, STAR=42, SLASH=43, PERCENT=44, CEL_TRUE=45, 
		CEL_FALSE=46, NUL=47, CEL_COMMENT=48, NUM_FLOAT=49, NUM_INT=50, NUM_UINT=51, 
		STRING=52, BYTES=53, NEWLINE=54, CONDITION_PARAM_CONTAINER=55, CONDITION_PARAM_TYPE=56;
	public static final int
		RULE_main = 0, RULE_modelHeader = 1, RULE_moduleHeader = 2, RULE_typeDefs = 3, 
		RULE_typeDef = 4, RULE_relationDeclaration = 5, RULE_relationName = 6, 
		RULE_relationDef = 7, RULE_relationDefNoDirect = 8, RULE_relationDefPartials = 9, 
		RULE_relationDefGrouping = 10, RULE_relationRecurse = 11, RULE_relationRecurseNoDirect = 12, 
		RULE_relationDefDirectAssignment = 13, RULE_relationDefRewrite = 14, RULE_relationDefTypeRestriction = 15, 
		RULE_relationDefTypeRestrictionBase = 16, RULE_conditions = 17, RULE_condition = 18, 
		RULE_conditionName = 19, RULE_conditionParameter = 20, RULE_parameterName = 21, 
		RULE_parameterType = 22, RULE_multiLineComment = 23, RULE_identifier = 24, 
		RULE_conditionExpression = 25;
	private static String[] makeRuleNames() {
		return new String[] {
			"main", "modelHeader", "moduleHeader", "typeDefs", "typeDef", "relationDeclaration", 
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
			"'#'", "'and'", "'or'", "'but not'", "'from'", "'module'", "'model'", 
			"'schema'", null, "'extend'", "'type'", "'condition'", "'relations'", 
			"'relation'", "'define'", "'with'", "'=='", "'!='", "'in'", "'<='", "'>='", 
			"'&&'", "'||'", "']'", "'{'", "'}'", "'.'", "'-'", "'!'", "'?'", "'+'", 
			"'*'", "'/'", "'%'", "'true'", "'false'", "'null'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "COLON", "COMMA", "LESS", "GREATER", "LBRACKET", "RBRACKET", "LPAREN", 
			"RPAREN", "WHITESPACE", "IDENTIFIER", "HASH", "AND", "OR", "BUT_NOT", 
			"FROM", "MODULE", "MODEL", "SCHEMA", "SCHEMA_VERSION", "EXTEND", "TYPE", 
			"CONDITION", "RELATIONS", "RELATION", "DEFINE", "KEYWORD_WITH", "EQUALS", 
			"NOT_EQUALS", "IN", "LESS_EQUALS", "GREATER_EQUALS", "LOGICAL_AND", "LOGICAL_OR", 
			"RPRACKET", "LBRACE", "RBRACE", "DOT", "MINUS", "EXCLAM", "QUESTIONMARK", 
			"PLUS", "STAR", "SLASH", "PERCENT", "CEL_TRUE", "CEL_FALSE", "NUL", "CEL_COMMENT", 
			"NUM_FLOAT", "NUM_INT", "NUM_UINT", "STRING", "BYTES", "NEWLINE", "CONDITION_PARAM_CONTAINER", 
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
		public TypeDefsContext typeDefs() {
			return getRuleContext(TypeDefsContext.class,0);
		}
		public ConditionsContext conditions() {
			return getRuleContext(ConditionsContext.class,0);
		}
		public TerminalNode EOF() { return getToken(OpenFGAParser.EOF, 0); }
		public ModelHeaderContext modelHeader() {
			return getRuleContext(ModelHeaderContext.class,0);
		}
		public ModuleHeaderContext moduleHeader() {
			return getRuleContext(ModuleHeaderContext.class,0);
		}
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
			setState(53);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(52);
				match(WHITESPACE);
				}
			}

			setState(56);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NEWLINE) {
				{
				setState(55);
				match(NEWLINE);
				}
			}

			setState(60);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				{
				setState(58);
				modelHeader();
				}
				break;
			case 2:
				{
				setState(59);
				moduleHeader();
				}
				break;
			}
			setState(63);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				{
				setState(62);
				match(NEWLINE);
				}
				break;
			}
			setState(65);
			typeDefs();
			setState(67);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				{
				setState(66);
				match(NEWLINE);
				}
				break;
			}
			setState(69);
			conditions();
			setState(71);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NEWLINE) {
				{
				setState(70);
				match(NEWLINE);
				}
			}

			setState(73);
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
			setState(78);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==HASH) {
				{
				setState(75);
				multiLineComment();
				setState(76);
				match(NEWLINE);
				}
			}

			setState(80);
			match(MODEL);
			setState(81);
			match(NEWLINE);
			setState(82);
			match(SCHEMA);
			setState(83);
			match(WHITESPACE);
			setState(84);
			((ModelHeaderContext)_localctx).schemaVersion = match(SCHEMA_VERSION);
			setState(86);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(85);
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
	public static class ModuleHeaderContext extends ParserRuleContext {
		public IdentifierContext moduleName;
		public TerminalNode MODULE() { return getToken(OpenFGAParser.MODULE, 0); }
		public List<TerminalNode> WHITESPACE() { return getTokens(OpenFGAParser.WHITESPACE); }
		public TerminalNode WHITESPACE(int i) {
			return getToken(OpenFGAParser.WHITESPACE, i);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public MultiLineCommentContext multiLineComment() {
			return getRuleContext(MultiLineCommentContext.class,0);
		}
		public TerminalNode NEWLINE() { return getToken(OpenFGAParser.NEWLINE, 0); }
		public ModuleHeaderContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_moduleHeader; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).enterModuleHeader(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof OpenFGAParserListener ) ((OpenFGAParserListener)listener).exitModuleHeader(this);
		}
	}

	public final ModuleHeaderContext moduleHeader() throws RecognitionException {
		ModuleHeaderContext _localctx = new ModuleHeaderContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_moduleHeader);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(91);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==HASH) {
				{
				setState(88);
				multiLineComment();
				setState(89);
				match(NEWLINE);
				}
			}

			setState(93);
			match(MODULE);
			setState(94);
			match(WHITESPACE);
			setState(95);
			((ModuleHeaderContext)_localctx).moduleName = identifier();
			setState(97);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(96);
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
		enterRule(_localctx, 6, RULE_typeDefs);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(102);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(99);
					typeDef();
					}
					} 
				}
				setState(104);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
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
		public List<TerminalNode> WHITESPACE() { return getTokens(OpenFGAParser.WHITESPACE); }
		public TerminalNode WHITESPACE(int i) {
			return getToken(OpenFGAParser.WHITESPACE, i);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public MultiLineCommentContext multiLineComment() {
			return getRuleContext(MultiLineCommentContext.class,0);
		}
		public TerminalNode EXTEND() { return getToken(OpenFGAParser.EXTEND, 0); }
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
		enterRule(_localctx, 8, RULE_typeDef);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(107);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				{
				setState(105);
				match(NEWLINE);
				setState(106);
				multiLineComment();
				}
				break;
			}
			setState(109);
			match(NEWLINE);
			setState(112);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EXTEND) {
				{
				setState(110);
				match(EXTEND);
				setState(111);
				match(WHITESPACE);
				}
			}

			setState(114);
			match(TYPE);
			setState(115);
			match(WHITESPACE);
			setState(116);
			((TypeDefContext)_localctx).typeName = identifier();
			setState(124);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				{
				setState(117);
				match(NEWLINE);
				setState(118);
				match(RELATIONS);
				setState(120); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(119);
						relationDeclaration();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(122); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
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
		enterRule(_localctx, 10, RULE_relationDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				{
				setState(126);
				match(NEWLINE);
				setState(127);
				multiLineComment();
				}
				break;
			}
			setState(130);
			match(NEWLINE);
			setState(131);
			match(DEFINE);
			setState(132);
			match(WHITESPACE);
			setState(133);
			relationName();
			setState(135);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(134);
				match(WHITESPACE);
				}
			}

			setState(137);
			match(COLON);
			setState(139);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(138);
				match(WHITESPACE);
				}
			}

			{
			setState(141);
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
		enterRule(_localctx, 12, RULE_relationName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(143);
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
		enterRule(_localctx, 14, RULE_relationDef);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(148);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LBRACKET:
				{
				setState(145);
				relationDefDirectAssignment();
				}
				break;
			case IDENTIFIER:
			case MODULE:
			case MODEL:
			case SCHEMA:
			case EXTEND:
			case TYPE:
			case RELATION:
				{
				setState(146);
				relationDefGrouping();
				}
				break;
			case LPAREN:
				{
				setState(147);
				relationRecurse();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(151);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				{
				setState(150);
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
		enterRule(_localctx, 16, RULE_relationDefNoDirect);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(155);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
			case MODULE:
			case MODEL:
			case SCHEMA:
			case EXTEND:
			case TYPE:
			case RELATION:
				{
				setState(153);
				relationDefGrouping();
				}
				break;
			case LPAREN:
				{
				setState(154);
				relationRecurseNoDirect();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(158);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				{
				setState(157);
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
		enterRule(_localctx, 18, RULE_relationDefPartials);
		try {
			int _alt;
			setState(189);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(167); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(160);
						match(WHITESPACE);
						setState(161);
						match(OR);
						setState(162);
						match(WHITESPACE);
						setState(165);
						_errHandler.sync(this);
						switch (_input.LA(1)) {
						case IDENTIFIER:
						case MODULE:
						case MODEL:
						case SCHEMA:
						case EXTEND:
						case TYPE:
						case RELATION:
							{
							setState(163);
							relationDefGrouping();
							}
							break;
						case LPAREN:
							{
							setState(164);
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
					setState(169); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(178); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(171);
						match(WHITESPACE);
						setState(172);
						match(AND);
						setState(173);
						match(WHITESPACE);
						setState(176);
						_errHandler.sync(this);
						switch (_input.LA(1)) {
						case IDENTIFIER:
						case MODULE:
						case MODEL:
						case SCHEMA:
						case EXTEND:
						case TYPE:
						case RELATION:
							{
							setState(174);
							relationDefGrouping();
							}
							break;
						case LPAREN:
							{
							setState(175);
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
					setState(180); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,25,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				{
				setState(182);
				match(WHITESPACE);
				setState(183);
				match(BUT_NOT);
				setState(184);
				match(WHITESPACE);
				setState(187);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case IDENTIFIER:
				case MODULE:
				case MODEL:
				case SCHEMA:
				case EXTEND:
				case TYPE:
				case RELATION:
					{
					setState(185);
					relationDefGrouping();
					}
					break;
				case LPAREN:
					{
					setState(186);
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
		enterRule(_localctx, 20, RULE_relationDefGrouping);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(191);
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
		enterRule(_localctx, 22, RULE_relationRecurse);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(193);
			match(LPAREN);
			setState(197);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==WHITESPACE) {
				{
				{
				setState(194);
				match(WHITESPACE);
				}
				}
				setState(199);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(202);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				{
				setState(200);
				relationDef();
				}
				break;
			case 2:
				{
				setState(201);
				relationRecurseNoDirect();
				}
				break;
			}
			setState(207);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==WHITESPACE) {
				{
				{
				setState(204);
				match(WHITESPACE);
				}
				}
				setState(209);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(210);
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
		enterRule(_localctx, 24, RULE_relationRecurseNoDirect);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(212);
			match(LPAREN);
			setState(216);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==WHITESPACE) {
				{
				{
				setState(213);
				match(WHITESPACE);
				}
				}
				setState(218);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(221);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				{
				setState(219);
				relationDefNoDirect();
				}
				break;
			case 2:
				{
				setState(220);
				relationRecurseNoDirect();
				}
				break;
			}
			setState(226);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==WHITESPACE) {
				{
				{
				setState(223);
				match(WHITESPACE);
				}
				}
				setState(228);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(229);
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
		enterRule(_localctx, 26, RULE_relationDefDirectAssignment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(231);
			match(LBRACKET);
			setState(233);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(232);
				match(WHITESPACE);
				}
			}

			setState(235);
			relationDefTypeRestriction();
			setState(237);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(236);
				match(WHITESPACE);
				}
			}

			setState(249);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(239);
				match(COMMA);
				setState(241);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==WHITESPACE) {
					{
					setState(240);
					match(WHITESPACE);
					}
				}

				setState(243);
				relationDefTypeRestriction();
				setState(245);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==WHITESPACE) {
					{
					setState(244);
					match(WHITESPACE);
					}
				}

				}
				}
				setState(251);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(252);
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
		public IdentifierContext rewriteComputedusersetName;
		public IdentifierContext rewriteTuplesetName;
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
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
		enterRule(_localctx, 28, RULE_relationDefRewrite);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(254);
			((RelationDefRewriteContext)_localctx).rewriteComputedusersetName = identifier();
			setState(259);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
			case 1:
				{
				setState(255);
				match(WHITESPACE);
				setState(256);
				match(FROM);
				setState(257);
				match(WHITESPACE);
				setState(258);
				((RelationDefRewriteContext)_localctx).rewriteTuplesetName = identifier();
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
		enterRule(_localctx, 30, RULE_relationDefTypeRestriction);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(262);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NEWLINE) {
				{
				setState(261);
				match(NEWLINE);
				}
			}

			setState(271);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,41,_ctx) ) {
			case 1:
				{
				setState(264);
				relationDefTypeRestrictionBase();
				}
				break;
			case 2:
				{
				{
				setState(265);
				relationDefTypeRestrictionBase();
				setState(266);
				match(WHITESPACE);
				setState(267);
				match(KEYWORD_WITH);
				setState(268);
				match(WHITESPACE);
				setState(269);
				conditionName();
				}
				}
				break;
			}
			setState(274);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NEWLINE) {
				{
				setState(273);
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
		public IdentifierContext relationDefTypeRestrictionType;
		public Token relationDefTypeRestrictionWildcard;
		public IdentifierContext relationDefTypeRestrictionRelation;
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
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
		enterRule(_localctx, 32, RULE_relationDefTypeRestrictionBase);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(276);
			((RelationDefTypeRestrictionBaseContext)_localctx).relationDefTypeRestrictionType = identifier();
			setState(281);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case COLON:
				{
				{
				setState(277);
				match(COLON);
				setState(278);
				((RelationDefTypeRestrictionBaseContext)_localctx).relationDefTypeRestrictionWildcard = match(STAR);
				}
				}
				break;
			case HASH:
				{
				{
				setState(279);
				match(HASH);
				setState(280);
				((RelationDefTypeRestrictionBaseContext)_localctx).relationDefTypeRestrictionRelation = identifier();
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
		enterRule(_localctx, 34, RULE_conditions);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(286);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,44,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(283);
					condition();
					}
					} 
				}
				setState(288);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,44,_ctx);
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
		enterRule(_localctx, 36, RULE_condition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(291);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,45,_ctx) ) {
			case 1:
				{
				setState(289);
				match(NEWLINE);
				setState(290);
				multiLineComment();
				}
				break;
			}
			setState(293);
			match(NEWLINE);
			setState(294);
			match(CONDITION);
			setState(295);
			match(WHITESPACE);
			setState(296);
			conditionName();
			setState(298);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(297);
				match(WHITESPACE);
				}
			}

			setState(300);
			match(LPAREN);
			setState(302);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(301);
				match(WHITESPACE);
				}
			}

			setState(304);
			conditionParameter();
			setState(306);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(305);
				match(WHITESPACE);
				}
			}

			setState(318);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(308);
				match(COMMA);
				setState(310);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==WHITESPACE) {
					{
					setState(309);
					match(WHITESPACE);
					}
				}

				setState(312);
				conditionParameter();
				setState(314);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==WHITESPACE) {
					{
					setState(313);
					match(WHITESPACE);
					}
				}

				}
				}
				setState(320);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(322);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NEWLINE) {
				{
				setState(321);
				match(NEWLINE);
				}
			}

			setState(324);
			match(RPAREN);
			setState(326);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(325);
				match(WHITESPACE);
				}
			}

			setState(328);
			match(LBRACE);
			setState(330);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,54,_ctx) ) {
			case 1:
				{
				setState(329);
				match(NEWLINE);
				}
				break;
			}
			setState(333);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,55,_ctx) ) {
			case 1:
				{
				setState(332);
				match(WHITESPACE);
				}
				break;
			}
			setState(335);
			conditionExpression();
			setState(337);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NEWLINE) {
				{
				setState(336);
				match(NEWLINE);
				}
			}

			setState(339);
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
		enterRule(_localctx, 38, RULE_conditionName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(341);
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
		enterRule(_localctx, 40, RULE_conditionParameter);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(344);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NEWLINE) {
				{
				setState(343);
				match(NEWLINE);
				}
			}

			setState(346);
			parameterName();
			setState(348);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(347);
				match(WHITESPACE);
				}
			}

			setState(350);
			match(COLON);
			setState(352);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHITESPACE) {
				{
				setState(351);
				match(WHITESPACE);
				}
			}

			setState(354);
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
		enterRule(_localctx, 42, RULE_parameterName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(356);
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
		enterRule(_localctx, 44, RULE_parameterType);
		try {
			setState(363);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CONDITION_PARAM_TYPE:
				enterOuterAlt(_localctx, 1);
				{
				setState(358);
				match(CONDITION_PARAM_TYPE);
				}
				break;
			case CONDITION_PARAM_CONTAINER:
				enterOuterAlt(_localctx, 2);
				{
				{
				setState(359);
				match(CONDITION_PARAM_CONTAINER);
				setState(360);
				match(LESS);
				setState(361);
				match(CONDITION_PARAM_TYPE);
				setState(362);
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
		enterRule(_localctx, 46, RULE_multiLineComment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(365);
			match(HASH);
			setState(369);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 126100789566373886L) != 0)) {
				{
				{
				setState(366);
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
				setState(371);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(374);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,62,_ctx) ) {
			case 1:
				{
				setState(372);
				match(NEWLINE);
				setState(373);
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
		public TerminalNode MODULE() { return getToken(OpenFGAParser.MODULE, 0); }
		public TerminalNode EXTEND() { return getToken(OpenFGAParser.EXTEND, 0); }
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
		enterRule(_localctx, 48, RULE_identifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(376);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 20382720L) != 0)) ) {
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
		enterRule(_localctx, 50, RULE_conditionExpression);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(382);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,64,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					setState(380);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,63,_ctx) ) {
					case 1:
						{
						setState(378);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 36028728165271480L) != 0)) ) {
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
						setState(379);
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
				setState(384);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,64,_ctx);
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
		"\u0004\u00018\u0182\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0001\u0000\u0003\u00006\b\u0000\u0001\u0000"+
		"\u0003\u00009\b\u0000\u0001\u0000\u0001\u0000\u0003\u0000=\b\u0000\u0001"+
		"\u0000\u0003\u0000@\b\u0000\u0001\u0000\u0001\u0000\u0003\u0000D\b\u0000"+
		"\u0001\u0000\u0001\u0000\u0003\u0000H\b\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001O\b\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001"+
		"W\b\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002\\\b\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002b\b\u0002\u0001"+
		"\u0003\u0005\u0003e\b\u0003\n\u0003\f\u0003h\t\u0003\u0001\u0004\u0001"+
		"\u0004\u0003\u0004l\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003"+
		"\u0004q\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0004\u0004y\b\u0004\u000b\u0004\f\u0004z\u0003\u0004"+
		"}\b\u0004\u0001\u0005\u0001\u0005\u0003\u0005\u0081\b\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005\u0088\b\u0005"+
		"\u0001\u0005\u0001\u0005\u0003\u0005\u008c\b\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007"+
		"\u0095\b\u0007\u0001\u0007\u0003\u0007\u0098\b\u0007\u0001\b\u0001\b\u0003"+
		"\b\u009c\b\b\u0001\b\u0003\b\u009f\b\b\u0001\t\u0001\t\u0001\t\u0001\t"+
		"\u0001\t\u0003\t\u00a6\b\t\u0004\t\u00a8\b\t\u000b\t\f\t\u00a9\u0001\t"+
		"\u0001\t\u0001\t\u0001\t\u0001\t\u0003\t\u00b1\b\t\u0004\t\u00b3\b\t\u000b"+
		"\t\f\t\u00b4\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0003\t\u00bc\b\t"+
		"\u0003\t\u00be\b\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0005\u000b"+
		"\u00c4\b\u000b\n\u000b\f\u000b\u00c7\t\u000b\u0001\u000b\u0001\u000b\u0003"+
		"\u000b\u00cb\b\u000b\u0001\u000b\u0005\u000b\u00ce\b\u000b\n\u000b\f\u000b"+
		"\u00d1\t\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0005\f\u00d7\b"+
		"\f\n\f\f\f\u00da\t\f\u0001\f\u0001\f\u0003\f\u00de\b\f\u0001\f\u0005\f"+
		"\u00e1\b\f\n\f\f\f\u00e4\t\f\u0001\f\u0001\f\u0001\r\u0001\r\u0003\r\u00ea"+
		"\b\r\u0001\r\u0001\r\u0003\r\u00ee\b\r\u0001\r\u0001\r\u0003\r\u00f2\b"+
		"\r\u0001\r\u0001\r\u0003\r\u00f6\b\r\u0005\r\u00f8\b\r\n\r\f\r\u00fb\t"+
		"\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0003\u000e\u0104\b\u000e\u0001\u000f\u0003\u000f\u0107\b\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0003\u000f\u0110\b\u000f\u0001\u000f\u0003\u000f\u0113\b"+
		"\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0003"+
		"\u0010\u011a\b\u0010\u0001\u0011\u0005\u0011\u011d\b\u0011\n\u0011\f\u0011"+
		"\u0120\t\u0011\u0001\u0012\u0001\u0012\u0003\u0012\u0124\b\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u012b"+
		"\b\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u012f\b\u0012\u0001\u0012"+
		"\u0001\u0012\u0003\u0012\u0133\b\u0012\u0001\u0012\u0001\u0012\u0003\u0012"+
		"\u0137\b\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u013b\b\u0012\u0005"+
		"\u0012\u013d\b\u0012\n\u0012\f\u0012\u0140\t\u0012\u0001\u0012\u0003\u0012"+
		"\u0143\b\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u0147\b\u0012\u0001"+
		"\u0012\u0001\u0012\u0003\u0012\u014b\b\u0012\u0001\u0012\u0003\u0012\u014e"+
		"\b\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u0152\b\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0014\u0003\u0014\u0159\b\u0014"+
		"\u0001\u0014\u0001\u0014\u0003\u0014\u015d\b\u0014\u0001\u0014\u0001\u0014"+
		"\u0003\u0014\u0161\b\u0014\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015"+
		"\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0003\u0016"+
		"\u016c\b\u0016\u0001\u0017\u0001\u0017\u0005\u0017\u0170\b\u0017\n\u0017"+
		"\f\u0017\u0173\t\u0017\u0001\u0017\u0001\u0017\u0003\u0017\u0177\b\u0017"+
		"\u0001\u0018\u0001\u0018\u0001\u0019\u0001\u0019\u0005\u0019\u017d\b\u0019"+
		"\n\u0019\f\u0019\u0180\t\u0019\u0001\u0019\u0000\u0000\u001a\u0000\u0002"+
		"\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e"+
		" \"$&(*,.02\u0000\u0004\u0001\u000066\u0004\u0000\n\n\u0010\u0012\u0014"+
		"\u0015\u0018\u0018\u0004\u0000\u0003\u0005\u0007\n\u001b#%6\u0001\u0000"+
		"$$\u01ab\u00005\u0001\u0000\u0000\u0000\u0002N\u0001\u0000\u0000\u0000"+
		"\u0004[\u0001\u0000\u0000\u0000\u0006f\u0001\u0000\u0000\u0000\bk\u0001"+
		"\u0000\u0000\u0000\n\u0080\u0001\u0000\u0000\u0000\f\u008f\u0001\u0000"+
		"\u0000\u0000\u000e\u0094\u0001\u0000\u0000\u0000\u0010\u009b\u0001\u0000"+
		"\u0000\u0000\u0012\u00bd\u0001\u0000\u0000\u0000\u0014\u00bf\u0001\u0000"+
		"\u0000\u0000\u0016\u00c1\u0001\u0000\u0000\u0000\u0018\u00d4\u0001\u0000"+
		"\u0000\u0000\u001a\u00e7\u0001\u0000\u0000\u0000\u001c\u00fe\u0001\u0000"+
		"\u0000\u0000\u001e\u0106\u0001\u0000\u0000\u0000 \u0114\u0001\u0000\u0000"+
		"\u0000\"\u011e\u0001\u0000\u0000\u0000$\u0123\u0001\u0000\u0000\u0000"+
		"&\u0155\u0001\u0000\u0000\u0000(\u0158\u0001\u0000\u0000\u0000*\u0164"+
		"\u0001\u0000\u0000\u0000,\u016b\u0001\u0000\u0000\u0000.\u016d\u0001\u0000"+
		"\u0000\u00000\u0178\u0001\u0000\u0000\u00002\u017e\u0001\u0000\u0000\u0000"+
		"46\u0005\t\u0000\u000054\u0001\u0000\u0000\u000056\u0001\u0000\u0000\u0000"+
		"68\u0001\u0000\u0000\u000079\u00056\u0000\u000087\u0001\u0000\u0000\u0000"+
		"89\u0001\u0000\u0000\u00009<\u0001\u0000\u0000\u0000:=\u0003\u0002\u0001"+
		"\u0000;=\u0003\u0004\u0002\u0000<:\u0001\u0000\u0000\u0000<;\u0001\u0000"+
		"\u0000\u0000=?\u0001\u0000\u0000\u0000>@\u00056\u0000\u0000?>\u0001\u0000"+
		"\u0000\u0000?@\u0001\u0000\u0000\u0000@A\u0001\u0000\u0000\u0000AC\u0003"+
		"\u0006\u0003\u0000BD\u00056\u0000\u0000CB\u0001\u0000\u0000\u0000CD\u0001"+
		"\u0000\u0000\u0000DE\u0001\u0000\u0000\u0000EG\u0003\"\u0011\u0000FH\u0005"+
		"6\u0000\u0000GF\u0001\u0000\u0000\u0000GH\u0001\u0000\u0000\u0000HI\u0001"+
		"\u0000\u0000\u0000IJ\u0005\u0000\u0000\u0001J\u0001\u0001\u0000\u0000"+
		"\u0000KL\u0003.\u0017\u0000LM\u00056\u0000\u0000MO\u0001\u0000\u0000\u0000"+
		"NK\u0001\u0000\u0000\u0000NO\u0001\u0000\u0000\u0000OP\u0001\u0000\u0000"+
		"\u0000PQ\u0005\u0011\u0000\u0000QR\u00056\u0000\u0000RS\u0005\u0012\u0000"+
		"\u0000ST\u0005\t\u0000\u0000TV\u0005\u0013\u0000\u0000UW\u0005\t\u0000"+
		"\u0000VU\u0001\u0000\u0000\u0000VW\u0001\u0000\u0000\u0000W\u0003\u0001"+
		"\u0000\u0000\u0000XY\u0003.\u0017\u0000YZ\u00056\u0000\u0000Z\\\u0001"+
		"\u0000\u0000\u0000[X\u0001\u0000\u0000\u0000[\\\u0001\u0000\u0000\u0000"+
		"\\]\u0001\u0000\u0000\u0000]^\u0005\u0010\u0000\u0000^_\u0005\t\u0000"+
		"\u0000_a\u00030\u0018\u0000`b\u0005\t\u0000\u0000a`\u0001\u0000\u0000"+
		"\u0000ab\u0001\u0000\u0000\u0000b\u0005\u0001\u0000\u0000\u0000ce\u0003"+
		"\b\u0004\u0000dc\u0001\u0000\u0000\u0000eh\u0001\u0000\u0000\u0000fd\u0001"+
		"\u0000\u0000\u0000fg\u0001\u0000\u0000\u0000g\u0007\u0001\u0000\u0000"+
		"\u0000hf\u0001\u0000\u0000\u0000ij\u00056\u0000\u0000jl\u0003.\u0017\u0000"+
		"ki\u0001\u0000\u0000\u0000kl\u0001\u0000\u0000\u0000lm\u0001\u0000\u0000"+
		"\u0000mp\u00056\u0000\u0000no\u0005\u0014\u0000\u0000oq\u0005\t\u0000"+
		"\u0000pn\u0001\u0000\u0000\u0000pq\u0001\u0000\u0000\u0000qr\u0001\u0000"+
		"\u0000\u0000rs\u0005\u0015\u0000\u0000st\u0005\t\u0000\u0000t|\u00030"+
		"\u0018\u0000uv\u00056\u0000\u0000vx\u0005\u0017\u0000\u0000wy\u0003\n"+
		"\u0005\u0000xw\u0001\u0000\u0000\u0000yz\u0001\u0000\u0000\u0000zx\u0001"+
		"\u0000\u0000\u0000z{\u0001\u0000\u0000\u0000{}\u0001\u0000\u0000\u0000"+
		"|u\u0001\u0000\u0000\u0000|}\u0001\u0000\u0000\u0000}\t\u0001\u0000\u0000"+
		"\u0000~\u007f\u00056\u0000\u0000\u007f\u0081\u0003.\u0017\u0000\u0080"+
		"~\u0001\u0000\u0000\u0000\u0080\u0081\u0001\u0000\u0000\u0000\u0081\u0082"+
		"\u0001\u0000\u0000\u0000\u0082\u0083\u00056\u0000\u0000\u0083\u0084\u0005"+
		"\u0019\u0000\u0000\u0084\u0085\u0005\t\u0000\u0000\u0085\u0087\u0003\f"+
		"\u0006\u0000\u0086\u0088\u0005\t\u0000\u0000\u0087\u0086\u0001\u0000\u0000"+
		"\u0000\u0087\u0088\u0001\u0000\u0000\u0000\u0088\u0089\u0001\u0000\u0000"+
		"\u0000\u0089\u008b\u0005\u0001\u0000\u0000\u008a\u008c\u0005\t\u0000\u0000"+
		"\u008b\u008a\u0001\u0000\u0000\u0000\u008b\u008c\u0001\u0000\u0000\u0000"+
		"\u008c\u008d\u0001\u0000\u0000\u0000\u008d\u008e\u0003\u000e\u0007\u0000"+
		"\u008e\u000b\u0001\u0000\u0000\u0000\u008f\u0090\u00030\u0018\u0000\u0090"+
		"\r\u0001\u0000\u0000\u0000\u0091\u0095\u0003\u001a\r\u0000\u0092\u0095"+
		"\u0003\u0014\n\u0000\u0093\u0095\u0003\u0016\u000b\u0000\u0094\u0091\u0001"+
		"\u0000\u0000\u0000\u0094\u0092\u0001\u0000\u0000\u0000\u0094\u0093\u0001"+
		"\u0000\u0000\u0000\u0095\u0097\u0001\u0000\u0000\u0000\u0096\u0098\u0003"+
		"\u0012\t\u0000\u0097\u0096\u0001\u0000\u0000\u0000\u0097\u0098\u0001\u0000"+
		"\u0000\u0000\u0098\u000f\u0001\u0000\u0000\u0000\u0099\u009c\u0003\u0014"+
		"\n\u0000\u009a\u009c\u0003\u0018\f\u0000\u009b\u0099\u0001\u0000\u0000"+
		"\u0000\u009b\u009a\u0001\u0000\u0000\u0000\u009c\u009e\u0001\u0000\u0000"+
		"\u0000\u009d\u009f\u0003\u0012\t\u0000\u009e\u009d\u0001\u0000\u0000\u0000"+
		"\u009e\u009f\u0001\u0000\u0000\u0000\u009f\u0011\u0001\u0000\u0000\u0000"+
		"\u00a0\u00a1\u0005\t\u0000\u0000\u00a1\u00a2\u0005\r\u0000\u0000\u00a2"+
		"\u00a5\u0005\t\u0000\u0000\u00a3\u00a6\u0003\u0014\n\u0000\u00a4\u00a6"+
		"\u0003\u0018\f\u0000\u00a5\u00a3\u0001\u0000\u0000\u0000\u00a5\u00a4\u0001"+
		"\u0000\u0000\u0000\u00a6\u00a8\u0001\u0000\u0000\u0000\u00a7\u00a0\u0001"+
		"\u0000\u0000\u0000\u00a8\u00a9\u0001\u0000\u0000\u0000\u00a9\u00a7\u0001"+
		"\u0000\u0000\u0000\u00a9\u00aa\u0001\u0000\u0000\u0000\u00aa\u00be\u0001"+
		"\u0000\u0000\u0000\u00ab\u00ac\u0005\t\u0000\u0000\u00ac\u00ad\u0005\f"+
		"\u0000\u0000\u00ad\u00b0\u0005\t\u0000\u0000\u00ae\u00b1\u0003\u0014\n"+
		"\u0000\u00af\u00b1\u0003\u0018\f\u0000\u00b0\u00ae\u0001\u0000\u0000\u0000"+
		"\u00b0\u00af\u0001\u0000\u0000\u0000\u00b1\u00b3\u0001\u0000\u0000\u0000"+
		"\u00b2\u00ab\u0001\u0000\u0000\u0000\u00b3\u00b4\u0001\u0000\u0000\u0000"+
		"\u00b4\u00b2\u0001\u0000\u0000\u0000\u00b4\u00b5\u0001\u0000\u0000\u0000"+
		"\u00b5\u00be\u0001\u0000\u0000\u0000\u00b6\u00b7\u0005\t\u0000\u0000\u00b7"+
		"\u00b8\u0005\u000e\u0000\u0000\u00b8\u00bb\u0005\t\u0000\u0000\u00b9\u00bc"+
		"\u0003\u0014\n\u0000\u00ba\u00bc\u0003\u0018\f\u0000\u00bb\u00b9\u0001"+
		"\u0000\u0000\u0000\u00bb\u00ba\u0001\u0000\u0000\u0000\u00bc\u00be\u0001"+
		"\u0000\u0000\u0000\u00bd\u00a7\u0001\u0000\u0000\u0000\u00bd\u00b2\u0001"+
		"\u0000\u0000\u0000\u00bd\u00b6\u0001\u0000\u0000\u0000\u00be\u0013\u0001"+
		"\u0000\u0000\u0000\u00bf\u00c0\u0003\u001c\u000e\u0000\u00c0\u0015\u0001"+
		"\u0000\u0000\u0000\u00c1\u00c5\u0005\u0007\u0000\u0000\u00c2\u00c4\u0005"+
		"\t\u0000\u0000\u00c3\u00c2\u0001\u0000\u0000\u0000\u00c4\u00c7\u0001\u0000"+
		"\u0000\u0000\u00c5\u00c3\u0001\u0000\u0000\u0000\u00c5\u00c6\u0001\u0000"+
		"\u0000\u0000\u00c6\u00ca\u0001\u0000\u0000\u0000\u00c7\u00c5\u0001\u0000"+
		"\u0000\u0000\u00c8\u00cb\u0003\u000e\u0007\u0000\u00c9\u00cb\u0003\u0018"+
		"\f\u0000\u00ca\u00c8\u0001\u0000\u0000\u0000\u00ca\u00c9\u0001\u0000\u0000"+
		"\u0000\u00cb\u00cf\u0001\u0000\u0000\u0000\u00cc\u00ce\u0005\t\u0000\u0000"+
		"\u00cd\u00cc\u0001\u0000\u0000\u0000\u00ce\u00d1\u0001\u0000\u0000\u0000"+
		"\u00cf\u00cd\u0001\u0000\u0000\u0000\u00cf\u00d0\u0001\u0000\u0000\u0000"+
		"\u00d0\u00d2\u0001\u0000\u0000\u0000\u00d1\u00cf\u0001\u0000\u0000\u0000"+
		"\u00d2\u00d3\u0005\b\u0000\u0000\u00d3\u0017\u0001\u0000\u0000\u0000\u00d4"+
		"\u00d8\u0005\u0007\u0000\u0000\u00d5\u00d7\u0005\t\u0000\u0000\u00d6\u00d5"+
		"\u0001\u0000\u0000\u0000\u00d7\u00da\u0001\u0000\u0000\u0000\u00d8\u00d6"+
		"\u0001\u0000\u0000\u0000\u00d8\u00d9\u0001\u0000\u0000\u0000\u00d9\u00dd"+
		"\u0001\u0000\u0000\u0000\u00da\u00d8\u0001\u0000\u0000\u0000\u00db\u00de"+
		"\u0003\u0010\b\u0000\u00dc\u00de\u0003\u0018\f\u0000\u00dd\u00db\u0001"+
		"\u0000\u0000\u0000\u00dd\u00dc\u0001\u0000\u0000\u0000\u00de\u00e2\u0001"+
		"\u0000\u0000\u0000\u00df\u00e1\u0005\t\u0000\u0000\u00e0\u00df\u0001\u0000"+
		"\u0000\u0000\u00e1\u00e4\u0001\u0000\u0000\u0000\u00e2\u00e0\u0001\u0000"+
		"\u0000\u0000\u00e2\u00e3\u0001\u0000\u0000\u0000\u00e3\u00e5\u0001\u0000"+
		"\u0000\u0000\u00e4\u00e2\u0001\u0000\u0000\u0000\u00e5\u00e6\u0005\b\u0000"+
		"\u0000\u00e6\u0019\u0001\u0000\u0000\u0000\u00e7\u00e9\u0005\u0005\u0000"+
		"\u0000\u00e8\u00ea\u0005\t\u0000\u0000\u00e9\u00e8\u0001\u0000\u0000\u0000"+
		"\u00e9\u00ea\u0001\u0000\u0000\u0000\u00ea\u00eb\u0001\u0000\u0000\u0000"+
		"\u00eb\u00ed\u0003\u001e\u000f\u0000\u00ec\u00ee\u0005\t\u0000\u0000\u00ed"+
		"\u00ec\u0001\u0000\u0000\u0000\u00ed\u00ee\u0001\u0000\u0000\u0000\u00ee"+
		"\u00f9\u0001\u0000\u0000\u0000\u00ef\u00f1\u0005\u0002\u0000\u0000\u00f0"+
		"\u00f2\u0005\t\u0000\u0000\u00f1\u00f0\u0001\u0000\u0000\u0000\u00f1\u00f2"+
		"\u0001\u0000\u0000\u0000\u00f2\u00f3\u0001\u0000\u0000\u0000\u00f3\u00f5"+
		"\u0003\u001e\u000f\u0000\u00f4\u00f6\u0005\t\u0000\u0000\u00f5\u00f4\u0001"+
		"\u0000\u0000\u0000\u00f5\u00f6\u0001\u0000\u0000\u0000\u00f6\u00f8\u0001"+
		"\u0000\u0000\u0000\u00f7\u00ef\u0001\u0000\u0000\u0000\u00f8\u00fb\u0001"+
		"\u0000\u0000\u0000\u00f9\u00f7\u0001\u0000\u0000\u0000\u00f9\u00fa\u0001"+
		"\u0000\u0000\u0000\u00fa\u00fc\u0001\u0000\u0000\u0000\u00fb\u00f9\u0001"+
		"\u0000\u0000\u0000\u00fc\u00fd\u0005\"\u0000\u0000\u00fd\u001b\u0001\u0000"+
		"\u0000\u0000\u00fe\u0103\u00030\u0018\u0000\u00ff\u0100\u0005\t\u0000"+
		"\u0000\u0100\u0101\u0005\u000f\u0000\u0000\u0101\u0102\u0005\t\u0000\u0000"+
		"\u0102\u0104\u00030\u0018\u0000\u0103\u00ff\u0001\u0000\u0000\u0000\u0103"+
		"\u0104\u0001\u0000\u0000\u0000\u0104\u001d\u0001\u0000\u0000\u0000\u0105"+
		"\u0107\u00056\u0000\u0000\u0106\u0105\u0001\u0000\u0000\u0000\u0106\u0107"+
		"\u0001\u0000\u0000\u0000\u0107\u010f\u0001\u0000\u0000\u0000\u0108\u0110"+
		"\u0003 \u0010\u0000\u0109\u010a\u0003 \u0010\u0000\u010a\u010b\u0005\t"+
		"\u0000\u0000\u010b\u010c\u0005\u001a\u0000\u0000\u010c\u010d\u0005\t\u0000"+
		"\u0000\u010d\u010e\u0003&\u0013\u0000\u010e\u0110\u0001\u0000\u0000\u0000"+
		"\u010f\u0108\u0001\u0000\u0000\u0000\u010f\u0109\u0001\u0000\u0000\u0000"+
		"\u0110\u0112\u0001\u0000\u0000\u0000\u0111\u0113\u00056\u0000\u0000\u0112"+
		"\u0111\u0001\u0000\u0000\u0000\u0112\u0113\u0001\u0000\u0000\u0000\u0113"+
		"\u001f\u0001\u0000\u0000\u0000\u0114\u0119\u00030\u0018\u0000\u0115\u0116"+
		"\u0005\u0001\u0000\u0000\u0116\u011a\u0005*\u0000\u0000\u0117\u0118\u0005"+
		"\u000b\u0000\u0000\u0118\u011a\u00030\u0018\u0000\u0119\u0115\u0001\u0000"+
		"\u0000\u0000\u0119\u0117\u0001\u0000\u0000\u0000\u0119\u011a\u0001\u0000"+
		"\u0000\u0000\u011a!\u0001\u0000\u0000\u0000\u011b\u011d\u0003$\u0012\u0000"+
		"\u011c\u011b\u0001\u0000\u0000\u0000\u011d\u0120\u0001\u0000\u0000\u0000"+
		"\u011e\u011c\u0001\u0000\u0000\u0000\u011e\u011f\u0001\u0000\u0000\u0000"+
		"\u011f#\u0001\u0000\u0000\u0000\u0120\u011e\u0001\u0000\u0000\u0000\u0121"+
		"\u0122\u00056\u0000\u0000\u0122\u0124\u0003.\u0017\u0000\u0123\u0121\u0001"+
		"\u0000\u0000\u0000\u0123\u0124\u0001\u0000\u0000\u0000\u0124\u0125\u0001"+
		"\u0000\u0000\u0000\u0125\u0126\u00056\u0000\u0000\u0126\u0127\u0005\u0016"+
		"\u0000\u0000\u0127\u0128\u0005\t\u0000\u0000\u0128\u012a\u0003&\u0013"+
		"\u0000\u0129\u012b\u0005\t\u0000\u0000\u012a\u0129\u0001\u0000\u0000\u0000"+
		"\u012a\u012b\u0001\u0000\u0000\u0000\u012b\u012c\u0001\u0000\u0000\u0000"+
		"\u012c\u012e\u0005\u0007\u0000\u0000\u012d\u012f\u0005\t\u0000\u0000\u012e"+
		"\u012d\u0001\u0000\u0000\u0000\u012e\u012f\u0001\u0000\u0000\u0000\u012f"+
		"\u0130\u0001\u0000\u0000\u0000\u0130\u0132\u0003(\u0014\u0000\u0131\u0133"+
		"\u0005\t\u0000\u0000\u0132\u0131\u0001\u0000\u0000\u0000\u0132\u0133\u0001"+
		"\u0000\u0000\u0000\u0133\u013e\u0001\u0000\u0000\u0000\u0134\u0136\u0005"+
		"\u0002\u0000\u0000\u0135\u0137\u0005\t\u0000\u0000\u0136\u0135\u0001\u0000"+
		"\u0000\u0000\u0136\u0137\u0001\u0000\u0000\u0000\u0137\u0138\u0001\u0000"+
		"\u0000\u0000\u0138\u013a\u0003(\u0014\u0000\u0139\u013b\u0005\t\u0000"+
		"\u0000\u013a\u0139\u0001\u0000\u0000\u0000\u013a\u013b\u0001\u0000\u0000"+
		"\u0000\u013b\u013d\u0001\u0000\u0000\u0000\u013c\u0134\u0001\u0000\u0000"+
		"\u0000\u013d\u0140\u0001\u0000\u0000\u0000\u013e\u013c\u0001\u0000\u0000"+
		"\u0000\u013e\u013f\u0001\u0000\u0000\u0000\u013f\u0142\u0001\u0000\u0000"+
		"\u0000\u0140\u013e\u0001\u0000\u0000\u0000\u0141\u0143\u00056\u0000\u0000"+
		"\u0142\u0141\u0001\u0000\u0000\u0000\u0142\u0143\u0001\u0000\u0000\u0000"+
		"\u0143\u0144\u0001\u0000\u0000\u0000\u0144\u0146\u0005\b\u0000\u0000\u0145"+
		"\u0147\u0005\t\u0000\u0000\u0146\u0145\u0001\u0000\u0000\u0000\u0146\u0147"+
		"\u0001\u0000\u0000\u0000\u0147\u0148\u0001\u0000\u0000\u0000\u0148\u014a"+
		"\u0005#\u0000\u0000\u0149\u014b\u00056\u0000\u0000\u014a\u0149\u0001\u0000"+
		"\u0000\u0000\u014a\u014b\u0001\u0000\u0000\u0000\u014b\u014d\u0001\u0000"+
		"\u0000\u0000\u014c\u014e\u0005\t\u0000\u0000\u014d\u014c\u0001\u0000\u0000"+
		"\u0000\u014d\u014e\u0001\u0000\u0000\u0000\u014e\u014f\u0001\u0000\u0000"+
		"\u0000\u014f\u0151\u00032\u0019\u0000\u0150\u0152\u00056\u0000\u0000\u0151"+
		"\u0150\u0001\u0000\u0000\u0000\u0151\u0152\u0001\u0000\u0000\u0000\u0152"+
		"\u0153\u0001\u0000\u0000\u0000\u0153\u0154\u0005$\u0000\u0000\u0154%\u0001"+
		"\u0000\u0000\u0000\u0155\u0156\u0005\n\u0000\u0000\u0156\'\u0001\u0000"+
		"\u0000\u0000\u0157\u0159\u00056\u0000\u0000\u0158\u0157\u0001\u0000\u0000"+
		"\u0000\u0158\u0159\u0001\u0000\u0000\u0000\u0159\u015a\u0001\u0000\u0000"+
		"\u0000\u015a\u015c\u0003*\u0015\u0000\u015b\u015d\u0005\t\u0000\u0000"+
		"\u015c\u015b\u0001\u0000\u0000\u0000\u015c\u015d\u0001\u0000\u0000\u0000"+
		"\u015d\u015e\u0001\u0000\u0000\u0000\u015e\u0160\u0005\u0001\u0000\u0000"+
		"\u015f\u0161\u0005\t\u0000\u0000\u0160\u015f\u0001\u0000\u0000\u0000\u0160"+
		"\u0161\u0001\u0000\u0000\u0000\u0161\u0162\u0001\u0000\u0000\u0000\u0162"+
		"\u0163\u0003,\u0016\u0000\u0163)\u0001\u0000\u0000\u0000\u0164\u0165\u0005"+
		"\n\u0000\u0000\u0165+\u0001\u0000\u0000\u0000\u0166\u016c\u00058\u0000"+
		"\u0000\u0167\u0168\u00057\u0000\u0000\u0168\u0169\u0005\u0003\u0000\u0000"+
		"\u0169\u016a\u00058\u0000\u0000\u016a\u016c\u0005\u0004\u0000\u0000\u016b"+
		"\u0166\u0001\u0000\u0000\u0000\u016b\u0167\u0001\u0000\u0000\u0000\u016c"+
		"-\u0001\u0000\u0000\u0000\u016d\u0171\u0005\u000b\u0000\u0000\u016e\u0170"+
		"\b\u0000\u0000\u0000\u016f\u016e\u0001\u0000\u0000\u0000\u0170\u0173\u0001"+
		"\u0000\u0000\u0000\u0171\u016f\u0001\u0000\u0000\u0000\u0171\u0172\u0001"+
		"\u0000\u0000\u0000\u0172\u0176\u0001\u0000\u0000\u0000\u0173\u0171\u0001"+
		"\u0000\u0000\u0000\u0174\u0175\u00056\u0000\u0000\u0175\u0177\u0003.\u0017"+
		"\u0000\u0176\u0174\u0001\u0000\u0000\u0000\u0176\u0177\u0001\u0000\u0000"+
		"\u0000\u0177/\u0001\u0000\u0000\u0000\u0178\u0179\u0007\u0001\u0000\u0000"+
		"\u01791\u0001\u0000\u0000\u0000\u017a\u017d\u0007\u0002\u0000\u0000\u017b"+
		"\u017d\b\u0003\u0000\u0000\u017c\u017a\u0001\u0000\u0000\u0000\u017c\u017b"+
		"\u0001\u0000\u0000\u0000\u017d\u0180\u0001\u0000\u0000\u0000\u017e\u017c"+
		"\u0001\u0000\u0000\u0000\u017e\u017f\u0001\u0000\u0000\u0000\u017f3\u0001"+
		"\u0000\u0000\u0000\u0180\u017e\u0001\u0000\u0000\u0000A58<?CGNV[afkpz"+
		"|\u0080\u0087\u008b\u0094\u0097\u009b\u009e\u00a5\u00a9\u00b0\u00b4\u00bb"+
		"\u00bd\u00c5\u00ca\u00cf\u00d8\u00dd\u00e2\u00e9\u00ed\u00f1\u00f5\u00f9"+
		"\u0103\u0106\u010f\u0112\u0119\u011e\u0123\u012a\u012e\u0132\u0136\u013a"+
		"\u013e\u0142\u0146\u014a\u014d\u0151\u0158\u015c\u0160\u016b\u0171\u0176"+
		"\u017c\u017e";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}