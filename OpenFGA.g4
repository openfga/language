grammar OpenFGA;

main: modelHeader typeDefs newline*;

indentation: '  ' | '	';

modelHeader: 'model' (newline+ multiLineComment)? newline indentation 'schema' spacing schemaVersion;
typeDefs: typeDef*;
typeDef:  (newline+ multiLineComment)? newline+ 'type' spacing typeName (newline indentation 'relations' relationDeclaration+)?;
relationDeclaration: (newline+ multiLineComment)? newline+ indentation indentation 'define' spacing relationName ':' spacing? relationDef;

relationDef: (relationDefDirectAssignment | relationDefRewrite) relationDefPartials?;

relationDefPartials: relationDefPartialAllOr | relationDefPartialAllAnd | relationDefPartialAllButNot;
relationDefPartialAllOr: (spacing relationDefOperatorOr spacing relationDefRewrite)+;
relationDefPartialAllAnd: (spacing relationDefOperatorAnd spacing relationDefRewrite)+;
relationDefPartialAllButNot: (spacing relationDefOperatorButNot spacing relationDefRewrite)+;

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

//relationDefComplexGrouping: relationDefGroup (spacing relationDefOperator spacing relationDefGroup)*;
//relationDefGroup: '(' spacing? (relationDefOrGrouping | relationDefAndGrouping | relationDefButNotGrouping) spacing? ')';
//relationDefAndGrouping: relationDefRewrite spacing relationDefOperatorAnd (spacing relationDefAndGrouping)+;
//relationDefOrGrouping: relationDefRewrite spacing relationDefOperatorOr (spacing relationDefOrGrouping)+;
//relationDefButNotGrouping: relationDefRewrite spacing relationDefOperatorButNot (spacing relationDefButNotGrouping)+;

rewriteComputedusersetName: name;
rewriteTuplesetComputedusersetName: name;
rewriteTuplesetName: name;
relationName: name;
typeName: name;

comment
  : spacing*  '#' ~( '\r' | '\n' )*
  ;
multiLineComment: comment (newline+ comment)*;
spacing: ' '+;
newline: '\n';
schemaVersion: '1.1';

name: WORD+;
WORD: [a-zA-Z0-9_-]+;
