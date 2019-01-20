grammar Mooncake;

mcrule
    : statementList
    ;

block
    : '{' statementList '}'
    ;

statementList
    : ( statement COMMENT? )*
    ;

statement
    : block
    | linkedStmt
    | simpleStmt
	;

inlineStmt
    : id=INLINE_ID
    | fn=function
    | id=INLINE_ID fn=function
    ;

errorStmt
    : '[' code=ERROR_CODE ',' info=ERROR_INFO ']' errType=errorType
    ;

linkedStmt
    : LINKED simpleStmt ( linkedStmt )?
    ;

simpleStmt
    : ( inlineStmt )? expression '=>' errorStmt block?
    ;

expression
    : id=identifier  op=operator  val=literal
    ;

identifier
    : IDENTIFIER
    | INLINE_ID
    ;

literal
    : INT
    | DOUBLE
    | BOOL
    | NULL
    | CTX_ID
    ;

function
    : LEN_FUNC
    | FLOAT_FUNC
    | DATETIME_LONG
    | AFTER_CURR_TIME
    ;

operator
    : EQ
    | NE
    | GT
    | LT
    | GTE
    | LTE
    ;

errorType
    : FATAL
    | SEVERE
    | WARNING
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

EXISTS
    : 'exists'
    ;

EMPTY
    : 'empty'
    ;

DOUBLE
    : '-'? INT '.' [0-9]+
    ;

INT
    : [0-9] [0-9]*
    ;

BOOL
    : TRUE
    | FALSE
    ;

TRUE
    : 'true'
    ;
FALSE
    : 'false'
    ;

NULL
    : 'nil'
    ;

COMMENT
    : '#' ~[\r\n]* -> skip
    ;

WS  :  [ \t]+ -> channel(HIDDEN)
    ;

TERMINATOR
	: [\r\n]+ -> channel(HIDDEN)
	;

IDENTIFIER
    : ( 'a'..'z' | 'A'..'Z' | '[' INT ']' )+ ( '.' IDENTIFIER )?
    ;

CTX_ID
    : '${' IDENTIFIER '}'
    ;

INLINE_ID
    : '_' ( 'a'..'z' | 'A'..'Z' )+
    ;

ERROR_CODE
    : ( 'a'..'z' | 'A'..'Z' | [0-9] )+
    ;

ERROR_INFO
    :  '\'' .*? '\''
    ;
