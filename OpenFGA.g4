grammar OpenFGA;

main: modelHeader typeDefs NEWLINES?;

modelHeaderComment: (HASH ~( '\r' | '\n' )*)? MULTILINE_COMMENT?;
modelHeader: (modelHeaderComment NEWLINES)? MODEL NEWLINES INDENT SCHEMA schemaVersion;
typeDefs: typeDef*;
typeDef: MULTILINE_COMMENT? NEWLINES TYPE typeName (NEWLINES INDENT RELATIONS relationDeclaration+)?;
relationDeclaration: MULTILINE_COMMENT? NEWLINES INDENT INDENT DEFINE relationName COLON relationDef ;

relationDef: (relationDefDirectAssignment | relationDefGrouping) relationDefPartials?;

relationDefPartials: relationDefPartialAllOr | relationDefPartialAllAnd | relationDefPartialAllButNot;
relationDefPartialAllOr: (OR relationDefGrouping)+;
relationDefPartialAllAnd: (AND relationDefGrouping)+;
relationDefPartialAllButNot: (BUT_NOT relationDefGrouping)+;

relationDefDirectAssignment: L_SQUARE relationDefTypeRestriction (COMMA relationDefTypeRestriction)* R_SQUARE;
relationDefRewrite: relationDefRelationOnSameObject | relationDefRelationOnRelatedObject;
relationDefRelationOnSameObject: rewriteComputedusersetName;
relationDefRelationOnRelatedObject: rewriteTuplesetComputedusersetName  FROM  rewriteTuplesetName;

relationDefTypeRestriction: relationDefTypeRestrictionType ((COLON WILDCARD) | (HASH relationDefTypeRestrictionRelation))?;
relationDefTypeRestrictionType: name;
relationDefTypeRestrictionRelation: name;

relationDefGrouping: relationDefRewrite;

rewriteComputedusersetName: name;
rewriteTuplesetComputedusersetName: name;
rewriteTuplesetName: name;
relationName: name;
typeName: name;

schemaVersion: SCHEMA_VERSION;
name: ALPHA_NUMERIC+;

MULTILINE_COMMENT: (NEWLINES WS* HASH COMMENT_CONTENTS)+;
INDENT: '  ' | '\t';

MODEL: 'model';
TYPE: 'type';
SCHEMA: 'schema';
SCHEMA_VERSION: '1.1';
RELATIONS: 'relations';
DEFINE: 'define';

AND: 'and';
OR: 'or';
BUT_NOT: 'but not';
FROM: 'from';

COLON: ':';
HASH: '#';
WILDCARD: '*';
L_SQUARE: '[';
R_SQUARE: ']';
COMMA: ',';

SYMBOL: [~!@#$%^&*()_[\]{}:";',.\\/<>`] | '+' | '=' | '-';
ALPHA_NUMERIC_CHAR: [a-zA-Z0-9_-];
ALPHA_NUMERIC: ALPHA_NUMERIC_CHAR+;
fragment COMMENT_CONTENTS: ~([\n\r\u2028\u2029])*;

NEWLINES: NEWLINE+;
fragment NEWLINE: '\r' '\n' | '\n' | '\r';


WS: [ \t] -> channel(HIDDEN);