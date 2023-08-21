parser grammar OpenFGAParser;
options { tokenVocab=OpenFGALexer; }

main: modelHeader typeDefs newline? EOF;

indentation: INDENT;

modelHeader: (multiLineComment newline)? MODEL spacing? (newline multiLineComment)? indentation SCHEMA spacing schemaVersion spacing?;
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

relationDefTypeRestriction: relationDefTypeRestrictionType | relationDefTypeRestrictionWildcard | relationDefTypeRestrictionUserset;
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

comment: WS* HASH ~(NEWLINE)*;
multiLineComment: comment (newline comment)*;
spacing: WS+;
newline: NEWLINE+;
schemaVersion: SCHEMA_VERSION;

name: ALPHA_NUMERIC+;

