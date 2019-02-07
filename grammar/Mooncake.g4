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
    : id=DECL_ID
    | fn=function
    | id=DECL_ID fn=function
    ;

errorStmt
    : '[' code=ERROR_CODE ',' info=ERROR_INFO ']' errType=errorType
    ;

linkedStmt
    : LINKED simpleStmt ( linkedStmt )?
    ;

simpleStmt
    : ( inlineStmt )? exprStmt '=>' errorStmt block?
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

EXISTS
    : 'exists'
    ;

EMPTY
    : 'empty'
    ;

FLOAT
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

ATTR_ID
    : ( 'a'..'z' | 'A'..'Z'  | '[' INT ']' )+ ( '.' ATTR_ID )?
    ;

CTX_ID
    : '${' ATTR_ID '}'
    ;

DECL_ID
    : '_' ( 'a'..'z' | 'A'..'Z' )+
    ;

ERROR_CODE
    : ( 'a'..'z' | 'A'..'Z' | [0-9] )+
    ;

ERROR_INFO
    :  '\'' .*? '\''
    ;
