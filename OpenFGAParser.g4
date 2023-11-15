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
    OPEN_CEL
    conditionExpression
    CLOSE_CEL;
conditionName: IDENTIFIER;
conditionParameter: NEWLINE? parameterName WHITESPACE? COLON WHITESPACE? parameterType;
parameterName: IDENTIFIER;
parameterType: CONDITION_PARAM_TYPE | (CONDITION_PARAM_CONTAINER LESS CONDITION_PARAM_TYPE GREATER);

multiLineComment: HASH (~NEWLINE)* (NEWLINE multiLineComment)?;

conditionExpression: (~(CLOSE_CEL))*;
