parser grammar OpenFGAParser;
options { tokenVocab=OpenFGALexer; }

main: WHITESPACE? NEWLINE? (modelHeader | moduleHeader) NEWLINE? mixins NEWLINE? NEWLINE? typeDefs NEWLINE? conditions NEWLINE? EOF;

// Model Header
modelHeader: (multiLineComment NEWLINE)? MODEL NEWLINE SCHEMA WHITESPACE schemaVersion=SCHEMA_VERSION WHITESPACE?;
// Module Header
moduleHeader: (multiLineComment NEWLINE)? MODULE WHITESPACE moduleName=identifier WHITESPACE?;

// Type Definitions
typeDefs: typeDef*;
typeDef:  (NEWLINE multiLineComment)? NEWLINE (EXTEND WHITESPACE)? TYPE WHITESPACE typeName=identifier (NEWLINE RELATIONS (mixinDeclaration | relationDeclaration)+)?;

// Relation definitions
mixinDeclaration: (NEWLINE INCLUDE WHITESPACE mixinName (WHITESPACE? COMMA WHITESPACE? mixinName)*);
relationDeclaration: (NEWLINE multiLineComment)? NEWLINE DEFINE WHITESPACE relationName WHITESPACE? COLON WHITESPACE? (relationDef);
relationName: identifier;

relationDef: (relationDefDirectAssignment | relationDefGrouping | relationRecurse) (relationDefPartials)?;
relationDefNoDirect: (relationDefGrouping | relationRecurseNoDirect) (relationDefPartials)?;

relationDefPartials:
    (WHITESPACE OR WHITESPACE (relationDefGrouping | relationRecurseNoDirect))+
    | (WHITESPACE AND WHITESPACE (relationDefGrouping | relationRecurseNoDirect))+
    | (WHITESPACE BUT_NOT WHITESPACE (relationDefGrouping | relationRecurseNoDirect));
    
relationDefGrouping: relationDefRewrite;

relationRecurse:
    LPAREN WHITESPACE* (
    relationDef |
    relationRecurseNoDirect
    ) WHITESPACE* RPAREN;

relationRecurseNoDirect:
    LPAREN WHITESPACE* (
    relationDefNoDirect |
    relationRecurseNoDirect
    ) WHITESPACE* RPAREN;

relationDefDirectAssignment: LBRACKET WHITESPACE? relationDefTypeRestriction WHITESPACE? (COMMA WHITESPACE? relationDefTypeRestriction WHITESPACE?)* RPRACKET;
relationDefRewrite: rewriteComputedusersetName=identifier (WHITESPACE FROM WHITESPACE rewriteTuplesetName=identifier)?;

relationDefTypeRestriction: NEWLINE? (
    relationDefTypeRestrictionBase
    | (relationDefTypeRestrictionBase WHITESPACE KEYWORD_WITH WHITESPACE conditionName)
    ) NEWLINE?;
relationDefTypeRestrictionBase: relationDefTypeRestrictionType=identifier
    ((COLON relationDefTypeRestrictionWildcard=STAR)
     | (HASH relationDefTypeRestrictionRelation=identifier))?;

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

// Mixins
mixins: mixin*;
mixin: (NEWLINE multiLineComment)? NEWLINE
    MIXIN WHITESPACE mixinName WHITESPACE?
    NEWLINE RELATIONS relationDeclaration+;
mixinName: IDENTIFIER;

multiLineComment: HASH (~NEWLINE)* (NEWLINE multiLineComment)?;

identifier: MODEL | SCHEMA | TYPE | RELATION | IDENTIFIER | MODULE | EXTEND;

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