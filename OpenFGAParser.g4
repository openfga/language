parser grammar OpenFGAParser;
options { tokenVocab=OpenFGALexer; }

main: WHITESPACE? NEWLINE? modelHeader NEWLINE? typeDefs NEWLINE? conditions NEWLINE? EOF;

// Model Header
modelHeader: (multiLineComment NEWLINE)? MODEL NEWLINE SCHEMA WHITESPACE schemaVersion=SCHEMA_VERSION WHITESPACE?;

// Type Definitions
typeDefs: typeDef*;
typeDef:  (NEWLINE multiLineComment)? NEWLINE TYPE WHITESPACE typeName=IDENTIFIER (NEWLINE RELATIONS relationDeclaration+)?;

relationDeclaration: NEWLINE DEFINE WHITESPACE relationName WHITESPACE? COLON WHITESPACE? relationDef;
relationName: IDENTIFIER;
relationDef: (relationDefDirectAssignment | relationDefGrouping) (relationDefPartials)?;
relationDefPartials: (WHITESPACE OR WHITESPACE relationDefGrouping)+ | (WHITESPACE AND WHITESPACE relationDefGrouping)+ | (WHITESPACE BUT_NOT WHITESPACE relationDefGrouping)+;
relationDefGrouping: relationDefRewrite;

relationDefDirectAssignment: LBRACKET WHITESPACE? relationDefTypeRestriction WHITESPACE? (COMMA WHITESPACE? relationDefTypeRestriction WHITESPACE?)* RPRACKET;
relationDefRewrite: rewriteComputedusersetName=IDENTIFIER (WHITESPACE FROM WHITESPACE rewriteTuplesetName=IDENTIFIER)?;

relationDefTypeRestriction: NEWLINE? (
    relationDefTypeRestrictionBase
    | (relationDefTypeRestrictionBase WHITESPACE KEYWORD_WITH WHITESPACE conditionName)
    ) NEWLINE?;
relationDefTypeRestrictionBase: relationDefTypeRestrictionType=IDENTIFIER
    ((COLON relationDefTypeRestrictionWildcard=STAR)
     | (HASH relationDefTypeRestrictionRelation=IDENTIFIER))?;

// Conditions
conditions: condition*;
condition: (NEWLINE multiLineComment)? NEWLINE
    CONDITION WHITESPACE conditionName WHITESPACE?
    LPAREN WHITESPACE? conditionParameter WHITESPACE? (COMMA WHITESPACE? conditionParameter WHITESPACE?)* NEWLINE? RPAREN WHITESPACE?
    LBRACE NEWLINE? WHITESPACE?
    conditionExpression
    NEWLINE? RBRACE;
conditionName: IDENTIFIER;
conditionParameter: NEWLINE? parameterName WHITESPACE? COLON WHITESPACE? parameterType;
parameterName: IDENTIFIER;
parameterType: CONDITION_PARAM_TYPE | (CONDITION_PARAM_CONTAINER LESS CONDITION_PARAM_TYPE GREATER);

multiLineComment: HASH (~NEWLINE)* (NEWLINE multiLineComment)?;

conditionExpression: ((
IDENTIFIER |
EQUALS |
NOT_EQUALS |
IN |
LESS |
LESS_EQUALS |
GREATER_EQUALS |
GREATER |
LOGICAL_AND |
LOGICAL_OR |
LBRACKET |
RPRACKET |
LBRACE |
LPAREN |
RPAREN |
DOT |
MINUS |
EXCLAM |
QUESTIONMARK |
PLUS |
STAR |
SLASH |
PERCENT |
CEL_TRUE |
CEL_FALSE |
NUL |
WHITESPACE |
CEL_COMMENT |
NUM_FLOAT |
NUM_INT |
NUM_UINT |
STRING |
BYTES |
NEWLINE |
WHITESPACE
)|~(RBRACE))*;
//
//// CEL Grammar, source: https://github.com/google/cel-go/blob/32ac6133c6b8eca8bb76e17e6ad50a1eb757778a/parser/gen/CEL.g4
//// License (Apache-2.0): https://github.com/google/cel-go/blob/32ac6133c6b8eca8bb76e17e6ad50a1eb757778a/LICENSE
//conditionExpression
//    : e=expr
//    ;
//
//expr
//    : e=conditionalOr (op='?' e1=conditionalOr ':' e2=expr)?
//    ;
//
//conditionalOr
//    : e=conditionalAnd (ops+='||' e1+=conditionalAnd)*
//    ;
//
//conditionalAnd
//    : e=relation (ops+='&&' e1+=relation)*
//    ;
//
//relation
//    : calc
//    | relation op=('<'|'<='|'>='|'>'|'=='|'!='|'in') relation
//    ;
//
//calc
//    : unary
//    | calc op=('*'|'/'|'%') calc
//    | calc op=('+'|'-') calc
//    ;
//
//unary
//    : member                                                        # MemberExpr
//    | (ops+='!')+ member                                            # LogicalNot
//    | (ops+='-')+ member                                            # Negate
//    ;
//
//member
//    : primary                                                       # PrimaryExpr
//    | member op='.' (opt='?')? id=IDENTIFIER                        # Select
//    | member op='.' id=IDENTIFIER open='(' args=exprList? ')'       # MemberCall
//    | member op='[' (opt='?')? index=expr ']'                       # Index
//    ;
//
//primary
//    : leadingDot='.'? id=IDENTIFIER (op='(' args=exprList? ')')?    # IdentOrGlobalCall
//    | '(' e=expr ')'                                                # Nested
//    | op='[' elems=listInit? ','? ']'                               # CreateList
//    | op='{' entries=mapInitializerList? ','? '}'                   # CreateStruct
//    | leadingDot='.'? ids+=IDENTIFIER (ops+='.' ids+=IDENTIFIER)*
//        op='{' entries=fieldInitializerList? ','? '}'               # CreateMessage
//    | literal                                                       # ConstantLiteral
//    ;
//
//exprList
//    : e+=expr (',' e+=expr)*
//    ;
//
//listInit
//    : elems+=optExpr (',' elems+=optExpr)*
//    ;
//
//fieldInitializerList
//    : fields+=optField cols+=':' values+=expr (',' fields+=optField cols+=':' values+=expr)*
//    ;
//
//optField
//    : (opt='?')? IDENTIFIER
//    ;
//
//mapInitializerList
//    : keys+=optExpr cols+=':' values+=expr (',' keys+=optExpr cols+=':' values+=expr)*
//    ;
//
//optExpr
//    : (opt='?')? e=expr
//    ;
//
//literal
//    : sign=MINUS? tok=NUM_INT   # Int
//    | tok=NUM_UINT              # Uint
//    | sign=MINUS? tok=NUM_FLOAT # Double
//    | tok=STRING                # String
//    | tok=BYTES                 # Bytes
//    | tok=CEL_TRUE              # BoolTrue
//    | tok=CEL_FALSE             # BoolFalse
//    | tok=NUL                   # Null
//    ;
//// END CEL Grammer
