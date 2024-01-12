lexer grammar OpenFGALexer;

HASH: '#';
COLON: ':';
COMMA: ',';

AND: 'and';
OR: 'or';
BUT_NOT: 'but not';
FROM: 'from';

MODEL: 'model';
SCHEMA: 'schema';
SCHEMA_VERSION: '1.1';
TYPE: 'type';
CONDITION: 'condition' -> pushMode(CONDITION_DEF);

RELATIONS: 'relations';
DEFINE: 'define';
KEYWORD_WITH: 'with';

LBRACKET: '[';
RPRACKET: ']';
LPAREN: '(';
RPAREN: ')';

IDENTIFIER: (LETTER | '_') (LETTER | DIGIT | '_' | MINUS)*;

WHITESPACE: ( '\t' | ' ' | '\u000C')+;

NEWLINE:
	WHITESPACE? ('\r'? '\n' | '\r' | '\f') WHITESPACE? NEWLINE?;

// CEL Lexer tokens, slightly modified source:
// https://github.com/google/cel-go/blob/32ac6133c6b8eca8bb76e17e6ad50a1eb757778a/parser/gen/CEL.g4

OPEN_CEL: '{' -> pushMode(CEL);

mode CEL;

CLOSE_CEL: '}' -> popMode;

EQUALS: '==';
NOT_EQUALS: '!=';
IN: 'in';
LESS: '<';
LESS_EQUALS: '<=';
GREATER_EQUALS: '>=';
GREATER: '>';
LOGICAL_AND: '&&';
LOGICAL_OR: '||';

CEL_LBRACKET: '[';
CEL_RPRACKET: ']';

DOT: '.';
MINUS: '-';
EXCLAM: '!';
QUESTIONMARK: '?';
PLUS: '+';
STAR: '*';
SLASH: '/';
PERCENT: '%';
CEL_TRUE: 'true';
CEL_FALSE: 'false';
NUL: 'null';

fragment BACKSLASH: '\\';
fragment LETTER: 'A' ..'Z' | 'a' ..'z';
fragment DIGIT: '0' ..'9';
fragment EXPONENT: ('e' | 'E') ( '+' | '-')? DIGIT+;
fragment HEXDIGIT: ('0' ..'9' | 'a' ..'f' | 'A' ..'F');
fragment RAW: 'r' | 'R';

fragment ESC_SEQ:
	ESC_CHAR_SEQ
	| ESC_BYTE_SEQ
	| ESC_UNI_SEQ
	| ESC_OCT_SEQ;

fragment ESC_CHAR_SEQ:
	BACKSLASH (
		'a'
		| 'b'
		| 'f'
		| 'n'
		| 'r'
		| 't'
		| 'v'
		| '"'
		| '\''
		| '\\'
		| '?'
		| '`'
	);

fragment ESC_OCT_SEQ:
	BACKSLASH ('0' ..'3') ('0' ..'7') ('0' ..'7');

fragment ESC_BYTE_SEQ: BACKSLASH ( 'x' | 'X') HEXDIGIT HEXDIGIT;

fragment ESC_UNI_SEQ:
	BACKSLASH 'u' HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT
	| BACKSLASH 'U' HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT;

CEL_COMMENT: '//' (~'\n')* -> channel(HIDDEN);

NUM_FLOAT: (
		DIGIT+ ('.' DIGIT+) EXPONENT?
		| DIGIT+ EXPONENT
		| '.' DIGIT+ EXPONENT?
	);

NUM_INT: ( DIGIT+ | '0x' HEXDIGIT+);

NUM_UINT: DIGIT+ ( 'u' | 'U') | '0x' HEXDIGIT+ ( 'u' | 'U');

STRING:
	'"' (ESC_SEQ | ~('\\' | '"' | '\n' | '\r'))* '"'
	| '\'' (ESC_SEQ | ~('\\' | '\'' | '\n' | '\r'))* '\''
	| '"""' (ESC_SEQ | ~('\\'))*? '"""'
	| '\'\'\'' (ESC_SEQ | ~('\\'))*? '\'\'\''
	| RAW '"' ~('"' | '\n' | '\r')* '"'
	| RAW '\'' ~('\'' | '\n' | '\r')* '\''
	| RAW '"""' .*? '"""'
	| RAW '\'\'\'' .*? '\'\'\'';

BYTES: ('b' | 'B') STRING;

CEL_IDENTIFIER: (LETTER | '_') (LETTER | DIGIT | '_' | MINUS)*;
// NOTE: MINUS is not allowed in CEL, but allowed in FGA, CEL will be revalidated after

CEL_WHITESPACE: ( '\t' | ' ' | '\u000C')+;

CEL_NEWLINE:
	CEL_WHITESPACE? ('\r'? '\n' | '\r' | '\f') CEL_WHITESPACE? CEL_NEWLINE?;

// END CEL GRAMMAR

mode CONDITION_DEF;

CLOSE_CONDITION_DEF: ')' -> popMode;

CONDITION_COLON: ':';
CONDITION_COMMA: ',';
CONDITION_LPAREN: '(';

CONDITION_PARAM_CONTAINER: 'map' | 'list';
CONDITION_PARAM_TYPE:
	'bool'
	| 'string'
	| 'int'
	| 'uint'
	| 'double'
	| 'duration'
	| 'timestamp'
	| 'ipaddress';

CONDITION_IDENTIFIER: (LETTER | '_') (
		LETTER
		| DIGIT
		| '_'
		| MINUS
	)*;

CONDITION_WHITESPACE: ( '\t' | ' ' | '\u000C')+;

CONDITION_NEWLINE:
	CONDITION_WHITESPACE? ('\r'? '\n' | '\r' | '\f') CEL_WHITESPACE? CEL_NEWLINE?;