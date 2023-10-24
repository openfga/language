lexer grammar OpenFGALexer;

fragment SINGLE_INDENT: ('  ' | '	');
fragment DOUBLE_INDENT:  ('    ' | '		');
fragment BOL : [\r\n\f]+;

INDENT: BOL (DOUBLE_INDENT | SINGLE_INDENT);

MODEL: 'model';
SCHEMA: 'schema';
SCHEMA_VERSION: '1.1';
TYPE: 'type';
CONDITION: 'condition';

RELATIONS: 'relations';
DEFINE: 'define';
WTH: 'with';

HASH: '#';
COLON: ':';
WILDCARD: '*';
L_SQUARE: '[';
R_SQUARE: ']';
L_PARANTHESES: '(';
R_PARANTHESES: ')';
L_BRACES: '{';
R_BRACES: '}';
L_ANGLE_BRACKET: '<';
R_ANGLE_BRACKET: '>';
COMMA: ',';
CONDITION_PARAM_CONTAINER: 'map' | 'list';
CONDITION_PARAM_TYPE: 'bool' | 'string' | 'int' | 'uint' |
  'double' | 'duration' | 'timestamp' | 'ipaddress';
CONDITION_SYMBOL:
 '==' | '!=' | 'in' | L_ANGLE_BRACKET | '<=' | '>=' | R_ANGLE_BRACKET | '&&' | '||'
 | L_SQUARE | R_SQUARE | L_BRACES | L_PARANTHESES | R_PARANTHESES | '.' | COMMA | '-' | '!' | '?' | COLON | '+' | WILDCARD | '/' | '%' | 'true' | 'false' | 'null' | '"';



AND: 'and';
OR: 'or';
BUT_NOT: 'but not';
FROM: 'from';

ALPHA_NUMERIC: [a-zA-Z0-9_-]+;

NEWLINE: [\r\n];
WS: [ \t];
