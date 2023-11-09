lexer grammar OpenFGALexer;
import CELLexer;

HASH: '#';
COLON: ':';
COMMA: ',';
CONDITION_PARAM_CONTAINER: 'map' | 'list';
CONDITION_PARAM_TYPE: 'bool' | 'string' | 'int' | 'uint' |
  'double' | 'duration' | 'timestamp' | 'ipaddress';

AND: 'and';
OR: 'or';
BUT_NOT: 'but not';
FROM: 'from';

MODEL: 'model';
SCHEMA: 'schema';
SCHEMA_VERSION: '1.1';
TYPE: 'type';
CONDITION: 'condition';

RELATIONS: 'relations';
DEFINE: 'define';
KEYWORD_WITH: 'with';

NEWLINE
 : WHITESPACE? ( '\r'? '\n' | '\r' | '\f' ) WHITESPACE? NEWLINE?
 ;
