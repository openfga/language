grammar OpenFGA;

main: modelHeader typeDefs NEWLINES?;

// COMMENT: HASH COMMENT_CONTENTS NEWLINE;

// modelHeader: (multiLineComment NEWLINES)? MODEL  (NEWLINES multiLineComment)? NEWLINES INDENT SCHEMA  schemaVersion ;
// typeDefs: typeDef*;
// typeDef:  (NEWLINES multiLineComment)? NEWLINES TYPE  typeName  (NEWLINES INDENT RELATIONS  relationDeclaration+)?;
// relationDeclaration: (NEWLINES multiLineComment)? NEWLINES INDENT INDENT DEFINE  relationName  COLLON  relationDef ;

modelHeader: MODEL NEWLINES INDENT SCHEMA  schemaVersion ;
typeDefs: typeDef*;
typeDef: NEWLINES TYPE typeName (NEWLINES INDENT RELATIONS relationDeclaration+)?;
relationDeclaration: NEWLINES INDENT INDENT DEFINE  relationName  COLON  relationDef ;

relationDef: (relationDefDirectAssignment | relationDefGrouping) relationDefPartials?;

relationDefPartials: relationDefPartialAllOr | relationDefPartialAllAnd | relationDefPartialAllButNot;
relationDefPartialAllOr: (OR relationDefGrouping)+;
relationDefPartialAllAnd: (AND relationDefGrouping)+;
relationDefPartialAllButNot: (BUT_NOT relationDefGrouping)+;

relationDefDirectAssignment: L_SQUARE relationDefTypeRestriction  (COMMA  relationDefTypeRestriction)*  R_SQUARE;
relationDefRewrite: relationDefRelationOnSameObject | relationDefRelationOnRelatedObject;
relationDefRelationOnSameObject: rewriteComputedusersetName;
relationDefRelationOnRelatedObject: rewriteTuplesetComputedusersetName  FROM  rewriteTuplesetName;

relationDefOperator: OR | AND | BUT_NOT;

relationDefTypeRestriction: relationDefTypeRestrictionType | relationDefTypeRestrictionWildcard | relationDefTypeRestrictionUserset;
relationDefTypeRestrictionType: name;
relationDefTypeRestrictionRelation: name;
relationDefTypeRestrictionWildcard: relationDefTypeRestrictionType WILDCARD;
relationDefTypeRestrictionUserset: relationDefTypeRestrictionType HASH relationDefTypeRestrictionRelation;

relationDefGrouping: relationDefRewrite;

rewriteComputedusersetName: name;
rewriteTuplesetComputedusersetName: name;
rewriteTuplesetName: name;
relationName: name;
typeName: name;

schemaVersion: SCHEMA_VERSION;
name: ALPHA_NUMERIC+;


INDENT: '  ' | '\t';

MODEL: 'model';
TYPE: 'type';
SCHEMA: 'schema';
SCHEMA_VERSION: '1.'[0-1];
RELATIONS: 'relations';
DEFINE: 'define';

AND: 'and';
OR: 'or';
BUT_NOT: 'but not';
FROM: 'from';

COLON: ':';
HASH: '#';
WILDCARD: ':*';
L_SQUARE: '[';
R_SQUARE: ']';
COMMA: ',';

ALPHA_NUMERIC: [a-zA-Z0-9_-]+;
fragment COMMENT_CONTENTS: ~([\n\r\u2028\u2029])*;

NEWLINES: NEWLINE+;
fragment NEWLINE: '\r' '\n' | '\n' | '\r';

WS: [ \t\r\n] -> channel(HIDDEN);