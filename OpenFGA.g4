grammar OpenFGA;

main: modelHeader typeDefs newline?;

indentation: '  ' | '	';

modelHeader: (multiLineComment newline)? 'model' (newline+ multiLineComment)? newline indentation 'schema' spacing schemaVersion;
typeDefs: typeDef*;
typeDef:  (newline multiLineComment)? newline+ 'type' spacing typeName (newline indentation 'relations' relationDeclaration+)?;
relationDeclaration: (newline multiLineComment)? newline indentation indentation 'define' spacing relationName spacing? ':' spacing? relationDef;

relationDef: (relationDefDirectAssignment | relationDefGrouping) relationDefPartials?;

relationDefPartials: relationDefPartialAllOr | relationDefPartialAllAnd | relationDefPartialAllButNot;
relationDefPartialAllOr: (spacing relationDefOperatorOr spacing relationDefGrouping)+;
relationDefPartialAllAnd: (spacing relationDefOperatorAnd spacing relationDefGrouping)+;
relationDefPartialAllButNot: (spacing relationDefOperatorButNot spacing relationDefGrouping)+;

relationDefDirectAssignment: '[' relationDefTypeRestriction spacing? (',' spacing? relationDefTypeRestriction)* spacing? ']';
relationDefRewrite: relationDefRelationOnSameObject | relationDefRelationOnRelatedObject;
relationDefRelationOnSameObject: rewriteComputedusersetName;
relationDefRelationOnRelatedObject: rewriteTuplesetComputedusersetName spacing relationDefKeywordFrom spacing rewriteTuplesetName;

relationDefOperator: relationDefOperatorOr | relationDefOperatorAnd | relationDefOperatorButNot;
relationDefOperatorAnd: 'and';
relationDefOperatorOr: 'or';
relationDefOperatorButNot: 'but not';
relationDefKeywordFrom: 'from';

relationDefTypeRestriction: relationDefTypeRestrictionType | relationDefTypeRestrictionWildcard | relationDefTypeRestrictionUserset;
relationDefTypeRestrictionType: name;
relationDefTypeRestrictionRelation: name;
relationDefTypeRestrictionWildcard: relationDefTypeRestrictionType ':*';
relationDefTypeRestrictionUserset: relationDefTypeRestrictionType '#' relationDefTypeRestrictionRelation;

relationDefGrouping: relationDefRewrite;
relationDefGroup: '('  relationDefGrouping relationDefPartials* ')';

rewriteComputedusersetName: name;
rewriteTuplesetComputedusersetName: name;
rewriteTuplesetName: name;
relationName: name;
typeName: name;

comment
  : spacing*  '#' ~( '\r' | '\n' )*
  ;
multiLineComment: comment (newline comment)*;
spacing: ' '+;
newline: '\n'+;
schemaVersion: '1.1';

name: WORD+;
WORD: [a-zA-Z0-9_-]+;
