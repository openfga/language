parser grammar OpenFGAParser;
options { tokenVocab=OpenFGALexer; }

main: modelHeader typeDefs conditions newline? EOF;

indentation: INDENT;

// Model Header
modelHeader: (multiLineComment newline)? MODEL spacing? (newline multiLineComment)? indentation SCHEMA spacing schemaVersion spacing?;

// Type Definitions
typeDefs: typeDef*;
typeDef:  (newline multiLineComment)? newline TYPE spacing typeName spacing? (indentation RELATIONS spacing? relationDeclaration+)?;

relationDeclaration: (newline multiLineComment)? indentation DEFINE spacing relationName spacing? COLON spacing? relationDef spacing?;
relationDef: (relationDefDirectAssignment | relationDefGrouping) relationDefPartials?;

relationDefPartials: relationDefPartialAllOr | relationDefPartialAllAnd | relationDefPartialAllButNot;
relationDefPartialAllOr: (spacing relationDefOperatorOr spacing relationDefGrouping)+;
relationDefPartialAllAnd: (spacing relationDefOperatorAnd spacing relationDefGrouping)+;
relationDefPartialAllButNot: (spacing relationDefOperatorButNot spacing relationDefGrouping)+;

relationDefDirectAssignment: L_SQUARE spacing? relationDefTypeRestriction spacing? (COMMA spacing? relationDefTypeRestriction)* spacing? R_SQUARE;
relationDefRewrite: relationDefRelationOnSameObject | relationDefRelationOnRelatedObject;
relationDefRelationOnSameObject: rewriteComputedusersetName;
relationDefRelationOnRelatedObject: rewriteTuplesetComputedusersetName spacing relationDefKeywordFrom spacing rewriteTuplesetName;

relationDefOperator: relationDefOperatorOr | relationDefOperatorAnd | relationDefOperatorButNot;
relationDefOperatorAnd: AND;
relationDefOperatorOr: OR;
relationDefOperatorButNot: BUT_NOT;
relationDefKeywordFrom: FROM;

relationDefTypeRestriction: relationDefTypeRestrictionType | relationDefTypeRestrictionWildcard | relationDefTypeRestrictionUserset | relationDefTypeRestrictionWithCondition;
relationDefTypeRestrictionWithCondition: (relationDefTypeRestrictionType | relationDefTypeRestrictionWildcard | relationDefTypeRestrictionUserset) spacing WTH spacing conditionName;
relationDefTypeRestrictionType: name;
relationDefTypeRestrictionRelation: name;
relationDefTypeRestrictionWildcard: relationDefTypeRestrictionType COLON WILDCARD spacing?;
relationDefTypeRestrictionUserset: relationDefTypeRestrictionType HASH relationDefTypeRestrictionRelation;

relationDefGrouping: relationDefRewrite;

rewriteComputedusersetName: name;
rewriteTuplesetComputedusersetName: name;
rewriteTuplesetName: name;
relationName: name;
typeName: name;

// Conditions
conditions: condition*;
condition: (newline multiLineComment)? newline
    CONDITION spacing conditionName spacing?
    L_PARANTHESES conditionParameter spacing? (COMMA spacing? conditionParameter spacing?)* R_PARANTHESES spacing?
    L_BRACES
    conditionExpression
    R_BRACES;

conditionParameter: parameterName spacing? COLON spacing? parameterType;
parameterName: name;
conditionName: name;
parameterType: CONDITION_PARAM_TYPE;
conditionExpression: (CONDITION_SYMBOL|~(R_BRACES))*;

// Base
comment: WS* HASH ~(NEWLINE)*;
multiLineComment: comment (newline comment)*;
spacing: WS+;
newline: NEWLINE+;
schemaVersion: SCHEMA_VERSION;

name: ALPHA_NUMERIC+;