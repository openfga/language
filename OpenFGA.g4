grammar OpenFGA;

main: modelHeader typeDefs newline?;

indentation: '  ' | '	';

modelHeader: (multiLineComment newline)? 'model' spacing? (newline multiLineComment)? newline indentation 'schema' spacing schemaVersion spacing?;
typeDefs: typeDef*;
typeDef:  (newline multiLineComment)? newline 'type' spacing typeName spacing? (newline indentation 'relations' spacing? relationDeclaration+)?;
relationDeclaration: (newline multiLineComment)? newline indentation indentation 'define' spacing relationName spacing? ':' spacing? relationDef spacing?;

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

rewriteComputedusersetName: name;
rewriteTuplesetComputedusersetName: name;
rewriteTuplesetName: name;
relationName: name;
typeName: name;

comment
  : spacing?  '#' ~( '\r' | '\n' )*
  ;
multiLineComment: comment (newline comment)*;
spacing: ' '+;
newline: ('\r' | '\n')+;
schemaVersion: '1.1';

name: ALPHA_NUMERIC+;
ALPHA_NUMERIC: [a-zA-Z0-9_-]+;
