parser grammar OpenFGAParser;
options {
	tokenVocab = OpenFGALexer;
}

main:
	WHITESPACE? NEWLINE? modelHeader NEWLINE? typeDefs NEWLINE? conditions NEWLINE? EOF;

// Model Header
modelHeader: (multiLineComment NEWLINE)? MODEL NEWLINE SCHEMA WHITESPACE schemaVersion =
		SCHEMA_VERSION WHITESPACE?;

// Type Definitions
typeDefs: typeDef*;
typeDef: (NEWLINE multiLineComment)? NEWLINE TYPE WHITESPACE typeName = IDENTIFIER (
		NEWLINE RELATIONS relationDeclaration+
	)?;

// Relation definitions
relationDeclaration: (NEWLINE multiLineComment)? NEWLINE DEFINE WHITESPACE relationName WHITESPACE?
		COLON WHITESPACE? (relationDef);
relationName: IDENTIFIER;

relationDef: (
		relationDefDirectAssignment
		| relationDefGrouping
		| relationRecurse
	) (relationDefPartials)?;
relationDefNoDirect: (
		relationDefGrouping
		| relationRecurseNoDirect
	) (relationDefPartials)?;

relationDefPartials:
	(
		WHITESPACE OR WHITESPACE (
			relationDefGrouping
			| relationRecurseNoDirect
		)
	)+
	| (
		WHITESPACE AND WHITESPACE (
			relationDefGrouping
			| relationRecurseNoDirect
		)
	)+
	| (
		WHITESPACE BUT_NOT WHITESPACE (
			relationDefGrouping
			| relationRecurseNoDirect
		)
	);

relationDefGrouping: relationDefRewrite;

relationRecurse:
	LPAREN WHITESPACE* (relationDef | relationRecurseNoDirect) WHITESPACE* RPAREN;

relationRecurseNoDirect:
	LPAREN WHITESPACE* (
		relationDefNoDirect
		| relationRecurseNoDirect
	) WHITESPACE* RPAREN;

relationDefDirectAssignment:
	LBRACKET WHITESPACE? relationDefTypeRestriction WHITESPACE? (
		COMMA WHITESPACE? relationDefTypeRestriction WHITESPACE?
	)* RPRACKET;
relationDefRewrite:
	rewriteComputedusersetName = IDENTIFIER (
		WHITESPACE FROM WHITESPACE rewriteTuplesetName = IDENTIFIER
	)?;

relationDefTypeRestriction:
	NEWLINE? (
		relationDefTypeRestrictionBase
		| (
			relationDefTypeRestrictionBase WHITESPACE KEYWORD_WITH WHITESPACE conditionName
		)
	) NEWLINE?;
relationDefTypeRestrictionBase:
	relationDefTypeRestrictionType = IDENTIFIER (
		(COLON relationDefTypeRestrictionWildcard = STAR)
		| (HASH relationDefTypeRestrictionRelation = IDENTIFIER)
	)?;

// Conditions
conditions: condition*;
condition: (NEWLINE multiLineComment)? NEWLINE CONDITION CONDITION_WHITESPACE conditionDef =
		CONDITION_IDENTIFIER CONDITION_WHITESPACE? CONDITION_LPAREN WHITESPACE? conditionParameter
		WHITESPACE? (
		CONDITION_COMMA CONDITION_WHITESPACE? conditionParameter WHITESPACE?
	)* NEWLINE? CLOSE_CONDITION_DEF WHITESPACE? OPEN_CEL NEWLINE? WHITESPACE? conditionExpression
		NEWLINE? CLOSE_CEL;

conditionName: IDENTIFIER;
conditionParameter:
	NEWLINE? parameterName CONDITION_WHITESPACE? CONDITION_COLON CONDITION_WHITESPACE? parameterType
		;
parameterName: CONDITION_IDENTIFIER;
parameterType:
	CONDITION_PARAM_TYPE
	| (
		CONDITION_PARAM_CONTAINER LESS CONDITION_PARAM_TYPE GREATER
	);

multiLineComment: HASH (~NEWLINE)* (NEWLINE multiLineComment)?;

conditionExpression: (~(CLOSE_CEL))*;