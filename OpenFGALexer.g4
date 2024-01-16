lexer grammar OpenFGALexer;

tokens {
	COLON,
	COMMA,
	LESS,
	GREATER,
	LBRACKET,
	RBRACKET,
	LPAREN,
	RPAREN,
	WHITESPACE,
	IDENTIFIER
}

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

// CEL Lexer tokens, slightly modified source:
// https://github.com/google/cel-go/blob/32ac6133c6b8eca8bb76e17e6ad50a1eb757778a/parser/gen/CEL.g4

EQUALS: '==';
NOT_EQUALS: '!=';
IN: 'in';
LESS: '<';
LESS_EQUALS: '<=';
GREATER_EQUALS: '>=';
GREATER: '>';
LOGICAL_AND: '&&';
LOGICAL_OR: '||';

LBRACKET: '[';
RPRACKET: ']';
LBRACE: '{';
RBRACE: '}';
LPAREN: '(';
RPAREN: ')';
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

WHITESPACE: ( '\t' | ' ' | '\u000C')+;
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

IDENTIFIER: (LETTER | '_') (LETTER | DIGIT | '_' | MINUS)*;
	// NOTE: MINUS is not allowed in CEL, but allowed in FGA, CEL will be revalidated after

// END CEL GRAMMAR

NEWLINE:
	WHITESPACE? ('\r'? '\n' | '\r' | '\f') WHITESPACE? NEWLINE?;

mode CONDITION_DEF;

CONDITION_DEF_END: RPAREN -> type(RPAREN), popMode;

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

CONDITION_PARAM_TYPE_LESS: LESS -> type(LESS);
CONDITION_PARAM_TYPE_GREATER: GREATER -> type(GREATER);

CONDITION_OPEN: LPAREN -> type(LPAREN);
CONDITION_COLON: COLON -> type(COLON);
CONDITION_COMMA: COMMA -> type(COMMA);

CONDITION_WS: WHITESPACE -> type(WHITESPACE);
CONDITION_NAME: IDENTIFIER -> type(IDENTIFIER);