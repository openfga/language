lexer grammar OpenFGALexer;

AND: 'and';
OR: 'or';
BUT_NOT: 'but not';
FROM: 'from';

MODEL: 'model';
SCHEMA: 'schema';
SCHEMA_VERSION: '1.1';
TYPE: 'type';
CONDITION: 'condition';

CONDITION_PARAM_CONTAINER: 'map' | 'list';
CONDITION_PARAM_TYPE: 'bool' | 'string' | 'int' | 'uint' |
  'double' | 'duration' | 'timestamp' | 'ipaddress';

RELATIONS: 'relations';
DEFINE: 'define';
KEYWORD_WITH: 'with';

IDENTIFIER : (LETTER | '_') ( LETTER | DIGIT | '_' | MINUS)*; // NOTE: MINUS is not allowed in CEL, but allowed in FGA, CEL will be revalidated after

WHITESPACE : ( '\t' | ' ' | '\u000C' )+;

NEWLINE
 : WHITESPACE? ( '\r'? '\n' | '\r' | '\f' ) WHITESPACE? NEWLINE?
 ;

DOT : '.';
STAR : '*';
HASH: '#';
COLON: ':';
COMMA: ',';
LPAREN : '(';
RPAREN : ')';
LESS : '<';
GREATER : '>';
LBRACKET : '[';
RPRACKET : ']';

// CEL Lexer tokens, slightly modified
// source: https://github.com/google/cel-go/blob/32ac6133c6b8eca8bb76e17e6ad50a1eb757778a/parser/gen/CEL.g4


OPEN_CEL: '{' -> pushMode(CEL);

mode CEL;
CEL_HASH: '#';
CEL_COLON: ':';
CEL_COMMA: ',';

EQUALS : '==';
NOT_EQUALS : '!=';
IN: 'in';
CEL_LESS : '<';
LESS_EQUALS : '<=';
GREATER_EQUALS : '>=';
CEL_GREATER : '>';
LOGICAL_AND : '&&';
LOGICAL_OR : '||';

CEL_LBRACKET : '[';
CEL_RPRACKET : ']';


CEL_LPAREN : '(';
CEL_RPAREN : ')';
CEL_DOT : '.';
MINUS : '-';
EXCLAM : '!';
QUESTIONMARK : '?';
PLUS : '+';
CEL_STAR : '*';
SLASH : '/';
PERCENT : '%';
CEL_TRUE : 'true';
CEL_FALSE : 'false';
NUL : 'null';

fragment BACKSLASH : '\\';
fragment LETTER : 'A'..'Z' | 'a'..'z' ;
fragment DIGIT  : '0'..'9' ;
fragment EXPONENT : ('e' | 'E') ( '+' | '-' )? DIGIT+ ;
fragment HEXDIGIT : ('0'..'9'|'a'..'f'|'A'..'F') ;
fragment RAW : 'r' | 'R';

fragment ESC_SEQ
    : ESC_CHAR_SEQ
    | ESC_BYTE_SEQ
    | ESC_UNI_SEQ
    | ESC_OCT_SEQ
    ;

fragment ESC_CHAR_SEQ
    : BACKSLASH ('a'|'b'|'f'|'n'|'r'|'t'|'v'|'"'|'\''|'\\'|'?'|'`')
    ;

fragment ESC_OCT_SEQ
    : BACKSLASH ('0'..'3') ('0'..'7') ('0'..'7')
    ;

fragment ESC_BYTE_SEQ
    : BACKSLASH ( 'x' | 'X' ) HEXDIGIT HEXDIGIT
    ;

fragment ESC_UNI_SEQ
    : BACKSLASH 'u' HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT
    | BACKSLASH 'U' HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT
    ;


CEL_COMMENT : '//' (~'\n')* -> channel(HIDDEN) ;

NUM_FLOAT
  : ( DIGIT+ ('.' DIGIT+) EXPONENT?
    | DIGIT+ EXPONENT
    | '.' DIGIT+ EXPONENT?
    )
  ;

NUM_INT
  : ( DIGIT+ | '0x' HEXDIGIT+ );

NUM_UINT
   : DIGIT+ ( 'u' | 'U' )
   | '0x' HEXDIGIT+ ( 'u' | 'U' )
   ;

STRING
  : '"' (ESC_SEQ | ~('\\'|'"'|'\n'|'\r'))* '"'
  | '\'' (ESC_SEQ | ~('\\'|'\''|'\n'|'\r'))* '\''
  | '"""' (ESC_SEQ | ~('\\'))*? '"""'
  | '\'\'\'' (ESC_SEQ | ~('\\'))*? '\'\'\''
  | RAW '"' ~('"'|'\n'|'\r')* '"'
  | RAW '\'' ~('\''|'\n'|'\r')* '\''
  | RAW '"""' .*? '"""'
  | RAW '\'\'\'' .*? '\'\'\''
  ;

BYTES : ('b' | 'B') STRING;

CEL_IDENTIFIER : (LETTER | '_') ( LETTER | DIGIT | '_' | MINUS)*; // NOTE: MINUS is not allowed in CEL, but allowed in FGA, CEL will be revalidated after

CEL_WHITESPACE : ( '\t' | ' ' | '\u000C' )+;

CEL_NEWLINE
 : CEL_WHITESPACE? ( '\r'? '\n' | '\r' | '\f' ) CEL_WHITESPACE? CEL_NEWLINE?
 ;

CLOSE_CEL: '}' -> popMode;

// END CEL GRAMMAR

