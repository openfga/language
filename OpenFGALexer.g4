lexer grammar OpenFGALexer;

fragment SINGLE_INDENT: ('  ' | '	');
fragment DOUBLE_INDENT:  ('    ' | '		');
fragment BOL : [\r\n\f]+;

INDENT: BOL (DOUBLE_INDENT | SINGLE_INDENT);

MODEL: 'model';
SCHEMA: 'schema';
SCHEMA_VERSION: '1.1';
TYPE: 'type';
RELATIONS: 'relations';
DEFINE: 'define';

HASH: '#';
COLON: ':';
WILDCARD: '*';
L_SQUARE: '[';
R_SQUARE: ']';
COMMA: ',';

AND: 'and';
OR: 'or';
BUT_NOT: 'but not';
FROM: 'from';

ALPHA_NUMERIC: [a-zA-Z0-9_-]+;

NEWLINE: [\r\n];
WS: [ \t];
