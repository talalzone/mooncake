grammar Mooncake;

mcrule
    : statementList
    ;

statementList
    : ( statement COMMENT? )*
    ;

statement
    : block
    | linkedStmt
    | simpleStmt
	;

block
    : '{' statementList '}'
    ;

linkedStmt
    : LINKED simpleStmt ( linkedStmt )?
    ;

simpleStmt
    : ( inlineStmt )? exprStmt '=>' errorStmt block?
    ;

inlineStmt
    : id=DECL_ID
    | fn=function
    | id=DECL_ID fn=function
    ;

errorStmt
    : '[' code=STRING ',' info=STRING ']' errType=errorType
    ;

exprStmt
    : id=identifier  op=operator  val=literal
    ;

identifier
    : attributeIdentifier
    | declarationIdentifier
    ;

literal
    : intLiteral
    | floatLiteral
    | boolLiteral
    | nullLiteral
    | stringLiteral
    | ctxLiteral
    ;

function
    : lengthFunction
    | dateTimeLongFunction
    | afterCurrentTimeFunction
    ;

operator
    : equalOperator
    | notEqualOperator
    | greaterThanOperator
    | lessThanOperator
    | greaterThanOrEqualOperator
    | lessThanOrEqualOperator
    ;

errorType
    : fatalError
    | severeError
    | warningError
    ;

fatalError
    : FATAL
    ;

severeError
    : SEVERE
    ;

warningError
    : WARNING
    ;

equalOperator
    : EQ
    ;

notEqualOperator
    : NE
    ;

greaterThanOperator
    : GT
    ;

lessThanOperator
    : LT
    ;

greaterThanOrEqualOperator
    : GTE
    ;

lessThanOrEqualOperator
    : LTE
    ;

intLiteral
    : INT
    ;

floatLiteral
    : FLOAT
    ;

boolLiteral
    : BOOL
    ;

nullLiteral
    : NULL
    ;

stringLiteral
    : STRING
    ;

ctxLiteral
    : CTX_ID
    ;

attributeIdentifier
    : ATTR_ID
    ;

declarationIdentifier
    : DECL_ID
    ;

lengthFunction
    : LEN_FUNC
    ;

dateTimeLongFunction
    : DATETIME_LONG
    ;

afterCurrentTimeFunction
    : AFTER_CURR_TIME
    ;

FATAL
    : '!!!'
    ;

SEVERE
    : '!!'
    ;

WARNING
    : '!'
    ;

LINKED
    : '~'
    ;

EQ  : 'eq'
    | '=='
    ;

NE  : 'ne'
    | '!='
    ;

AND : 'and'
    | '&&'
    ;

OR  : 'or'
    | '||'
    ;

GT  : 'gt'
    | '>'
    ;

LT  : 'lt'
    | '<'
    ;

GTE : 'gte'
    | '>='
    ;

LTE : 'lte'
    | '<='
    ;

LEN_FUNC
    : '@len'
    ;

FLOAT_FUNC
    : '@float'
    ;

DATETIME_LONG
    : '@dateTimeLong'
    ;

AFTER_CURR_TIME
    : '@afterCurrentTime'
    ;

FLOAT
    :  INT '.' Numeric+
    ;

INT
    : '-'? Numeric+
    ;

BOOL
    : TRUE
    | FALSE
    ;

NULL
    : 'nil'
    | 'null'
    ;

STRING
    : SingleQuoteString
    ;

TRUE
    : 'true'
    ;

FALSE
    : 'false'
    ;

COMMENT
    : ('#' ~[\r\n]*) -> skip
    ;

WS  :  [ \t]+ -> channel(HIDDEN)
    ;

TERMINATOR
	: [\r\n]+ -> channel(HIDDEN)
	;

DECL_ID
    : Underscore AlphaNumeric
    ;

ATTR_ID
    : ( AlphaNumeric | '[' Numeric ']' | Underscore )+ ( '.' ATTR_ID )?
    ;

CTX_ID
    : '${' ATTR_ID '}'
    ;

fragment Character
    : ('a'..'z' | 'A'..'Z')
    ;

fragment Numeric
    : [0-9]
    ;

fragment AlphaNumeric
    : (Character | Numeric)
    ;

fragment Underscore
    : '_'
    ;

fragment SingleQuote
    : '\''
    ;

fragment DoubleQuote
    : '"'
    ;

fragment SingleQuoteString
    : SingleQuote ( EscSeq | ~['\r\n\\] )* SingleQuote
    ;

fragment EscSeq
	:	Esc ( [btnfr"'\\] | . | EOF	)
	;

fragment EscAny
	:	Esc .
	;

fragment Esc
    : '\\'
    ;

