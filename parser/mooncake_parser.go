// Code generated from /Users/talal/go/src/mooncake/grammar/Mooncake.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Mooncake

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 37, 192, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34, 
	9, 34, 4, 35, 9, 35, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 75, 10, 3, 7, 3, 77, 
	10, 3, 12, 3, 14, 3, 80, 11, 3, 3, 4, 3, 4, 3, 4, 5, 4, 85, 10, 4, 3, 5, 
	3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 5, 6, 94, 10, 6, 3, 7, 5, 7, 97, 10, 
	7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 103, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 5, 
	8, 109, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 
	3, 10, 3, 10, 3, 11, 3, 11, 5, 11, 124, 10, 11, 3, 12, 3, 12, 3, 12, 3, 
	12, 3, 12, 3, 12, 5, 12, 132, 10, 12, 3, 13, 3, 13, 3, 13, 5, 13, 137, 
	10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 145, 10, 14, 3, 
	15, 3, 15, 3, 15, 5, 15, 150, 10, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 
	3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 
	23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 
	3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 
	34, 3, 34, 3, 35, 3, 35, 3, 35, 2, 2, 36, 2, 4, 6, 8, 10, 12, 14, 16, 18, 
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 
	56, 58, 60, 62, 64, 66, 68, 2, 2, 2, 181, 2, 70, 3, 2, 2, 2, 4, 78, 3, 
	2, 2, 2, 6, 84, 3, 2, 2, 2, 8, 86, 3, 2, 2, 2, 10, 90, 3, 2, 2, 2, 12, 
	96, 3, 2, 2, 2, 14, 108, 3, 2, 2, 2, 16, 110, 3, 2, 2, 2, 18, 117, 3, 2, 
	2, 2, 20, 123, 3, 2, 2, 2, 22, 131, 3, 2, 2, 2, 24, 136, 3, 2, 2, 2, 26, 
	144, 3, 2, 2, 2, 28, 149, 3, 2, 2, 2, 30, 151, 3, 2, 2, 2, 32, 153, 3, 
	2, 2, 2, 34, 155, 3, 2, 2, 2, 36, 157, 3, 2, 2, 2, 38, 159, 3, 2, 2, 2, 
	40, 161, 3, 2, 2, 2, 42, 163, 3, 2, 2, 2, 44, 165, 3, 2, 2, 2, 46, 167, 
	3, 2, 2, 2, 48, 169, 3, 2, 2, 2, 50, 171, 3, 2, 2, 2, 52, 173, 3, 2, 2, 
	2, 54, 175, 3, 2, 2, 2, 56, 177, 3, 2, 2, 2, 58, 179, 3, 2, 2, 2, 60, 181, 
	3, 2, 2, 2, 62, 183, 3, 2, 2, 2, 64, 185, 3, 2, 2, 2, 66, 187, 3, 2, 2, 
	2, 68, 189, 3, 2, 2, 2, 70, 71, 5, 4, 3, 2, 71, 3, 3, 2, 2, 2, 72, 74, 
	5, 6, 4, 2, 73, 75, 7, 32, 2, 2, 74, 73, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 
	75, 77, 3, 2, 2, 2, 76, 72, 3, 2, 2, 2, 77, 80, 3, 2, 2, 2, 78, 76, 3, 
	2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 5, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 81, 
	85, 5, 8, 5, 2, 82, 85, 5, 10, 6, 2, 83, 85, 5, 12, 7, 2, 84, 81, 3, 2, 
	2, 2, 84, 82, 3, 2, 2, 2, 84, 83, 3, 2, 2, 2, 85, 7, 3, 2, 2, 2, 86, 87, 
	7, 3, 2, 2, 87, 88, 5, 4, 3, 2, 88, 89, 7, 4, 2, 2, 89, 9, 3, 2, 2, 2, 
	90, 91, 7, 12, 2, 2, 91, 93, 5, 12, 7, 2, 92, 94, 5, 10, 6, 2, 93, 92, 
	3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 11, 3, 2, 2, 2, 95, 97, 5, 14, 8, 2, 
	96, 95, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 5, 
	18, 10, 2, 99, 100, 7, 5, 2, 2, 100, 102, 5, 16, 9, 2, 101, 103, 5, 8, 
	5, 2, 102, 101, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 13, 3, 2, 2, 2, 
	104, 109, 7, 35, 2, 2, 105, 109, 5, 24, 13, 2, 106, 107, 7, 35, 2, 2, 107, 
	109, 5, 24, 13, 2, 108, 104, 3, 2, 2, 2, 108, 105, 3, 2, 2, 2, 108, 106, 
	3, 2, 2, 2, 109, 15, 3, 2, 2, 2, 110, 111, 7, 6, 2, 2, 111, 112, 7, 29, 
	2, 2, 112, 113, 7, 7, 2, 2, 113, 114, 7, 29, 2, 2, 114, 115, 7, 8, 2, 2, 
	115, 116, 5, 28, 15, 2, 116, 17, 3, 2, 2, 2, 117, 118, 5, 20, 11, 2, 118, 
	119, 5, 26, 14, 2, 119, 120, 5, 22, 12, 2, 120, 19, 3, 2, 2, 2, 121, 124, 
	5, 60, 31, 2, 122, 124, 5, 62, 32, 2, 123, 121, 3, 2, 2, 2, 123, 122, 3, 
	2, 2, 2, 124, 21, 3, 2, 2, 2, 125, 132, 5, 48, 25, 2, 126, 132, 5, 50, 
	26, 2, 127, 132, 5, 52, 27, 2, 128, 132, 5, 54, 28, 2, 129, 132, 5, 56, 
	29, 2, 130, 132, 5, 58, 30, 2, 131, 125, 3, 2, 2, 2, 131, 126, 3, 2, 2, 
	2, 131, 127, 3, 2, 2, 2, 131, 128, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 131, 
	130, 3, 2, 2, 2, 132, 23, 3, 2, 2, 2, 133, 137, 5, 64, 33, 2, 134, 137, 
	5, 66, 34, 2, 135, 137, 5, 68, 35, 2, 136, 133, 3, 2, 2, 2, 136, 134, 3, 
	2, 2, 2, 136, 135, 3, 2, 2, 2, 137, 25, 3, 2, 2, 2, 138, 145, 5, 36, 19, 
	2, 139, 145, 5, 38, 20, 2, 140, 145, 5, 40, 21, 2, 141, 145, 5, 42, 22, 
	2, 142, 145, 5, 44, 23, 2, 143, 145, 5, 46, 24, 2, 144, 138, 3, 2, 2, 2, 
	144, 139, 3, 2, 2, 2, 144, 140, 3, 2, 2, 2, 144, 141, 3, 2, 2, 2, 144, 
	142, 3, 2, 2, 2, 144, 143, 3, 2, 2, 2, 145, 27, 3, 2, 2, 2, 146, 150, 5, 
	30, 16, 2, 147, 150, 5, 32, 17, 2, 148, 150, 5, 34, 18, 2, 149, 146, 3, 
	2, 2, 2, 149, 147, 3, 2, 2, 2, 149, 148, 3, 2, 2, 2, 150, 29, 3, 2, 2, 
	2, 151, 152, 7, 9, 2, 2, 152, 31, 3, 2, 2, 2, 153, 154, 7, 10, 2, 2, 154, 
	33, 3, 2, 2, 2, 155, 156, 7, 11, 2, 2, 156, 35, 3, 2, 2, 2, 157, 158, 7, 
	13, 2, 2, 158, 37, 3, 2, 2, 2, 159, 160, 7, 14, 2, 2, 160, 39, 3, 2, 2, 
	2, 161, 162, 7, 17, 2, 2, 162, 41, 3, 2, 2, 2, 163, 164, 7, 18, 2, 2, 164, 
	43, 3, 2, 2, 2, 165, 166, 7, 19, 2, 2, 166, 45, 3, 2, 2, 2, 167, 168, 7, 
	20, 2, 2, 168, 47, 3, 2, 2, 2, 169, 170, 7, 26, 2, 2, 170, 49, 3, 2, 2, 
	2, 171, 172, 7, 25, 2, 2, 172, 51, 3, 2, 2, 2, 173, 174, 7, 27, 2, 2, 174, 
	53, 3, 2, 2, 2, 175, 176, 7, 28, 2, 2, 176, 55, 3, 2, 2, 2, 177, 178, 7, 
	29, 2, 2, 178, 57, 3, 2, 2, 2, 179, 180, 7, 37, 2, 2, 180, 59, 3, 2, 2, 
	2, 181, 182, 7, 36, 2, 2, 182, 61, 3, 2, 2, 2, 183, 184, 7, 35, 2, 2, 184, 
	63, 3, 2, 2, 2, 185, 186, 7, 21, 2, 2, 186, 65, 3, 2, 2, 2, 187, 188, 7, 
	23, 2, 2, 188, 67, 3, 2, 2, 2, 189, 190, 7, 24, 2, 2, 190, 69, 3, 2, 2, 
	2, 14, 74, 78, 84, 93, 96, 102, 108, 123, 131, 136, 144, 149,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "'=>'", "'['", "','", "']'", "'!!!'", "'!!'", "'!'", 
	"'~'", "", "", "", "", "", "", "", "", "'@len'", "'@float'", "'@dateTimeLong'", 
	"'@afterCurrentTime'", "", "", "", "", "", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "FATAL", "SEVERE", "WARNING", "LINKED", "EQ", 
	"NE", "AND", "OR", "GT", "LT", "GTE", "LTE", "LEN_FUNC", "FLOAT_FUNC", 
	"DATETIME_LONG", "AFTER_CURR_TIME", "FLOAT", "INT", "BOOL", "NULL", "STRING", 
	"TRUE", "FALSE", "COMMENT", "WS", "TERMINATOR", "DECL_ID", "ATTR_ID", "CTX_ID",
}

var ruleNames = []string{
	"mcrule", "statementList", "statement", "block", "linkedStmt", "simpleStmt", 
	"inlineStmt", "errorStmt", "exprStmt", "identifier", "literal", "function", 
	"operator", "errorType", "fatalError", "severeError", "warningError", "equalOperator", 
	"notEqualOperator", "greaterThanOperator", "lessThanOperator", "greaterThanOrEqualOperator", 
	"lessThanOrEqualOperator", "intLiteral", "floatLiteral", "boolLiteral", 
	"nullLiteral", "stringLiteral", "ctxLiteral", "attributeIdentifier", "declarationIdentifier", 
	"lengthFunction", "dateTimeLongFunction", "afterCurrentTimeFunction",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type MooncakeParser struct {
	*antlr.BaseParser
}

func NewMooncakeParser(input antlr.TokenStream) *MooncakeParser {
	this := new(MooncakeParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Mooncake.g4"

	return this
}

// MooncakeParser tokens.
const (
	MooncakeParserEOF = antlr.TokenEOF
	MooncakeParserT__0 = 1
	MooncakeParserT__1 = 2
	MooncakeParserT__2 = 3
	MooncakeParserT__3 = 4
	MooncakeParserT__4 = 5
	MooncakeParserT__5 = 6
	MooncakeParserFATAL = 7
	MooncakeParserSEVERE = 8
	MooncakeParserWARNING = 9
	MooncakeParserLINKED = 10
	MooncakeParserEQ = 11
	MooncakeParserNE = 12
	MooncakeParserAND = 13
	MooncakeParserOR = 14
	MooncakeParserGT = 15
	MooncakeParserLT = 16
	MooncakeParserGTE = 17
	MooncakeParserLTE = 18
	MooncakeParserLEN_FUNC = 19
	MooncakeParserFLOAT_FUNC = 20
	MooncakeParserDATETIME_LONG = 21
	MooncakeParserAFTER_CURR_TIME = 22
	MooncakeParserFLOAT = 23
	MooncakeParserINT = 24
	MooncakeParserBOOL = 25
	MooncakeParserNULL = 26
	MooncakeParserSTRING = 27
	MooncakeParserTRUE = 28
	MooncakeParserFALSE = 29
	MooncakeParserCOMMENT = 30
	MooncakeParserWS = 31
	MooncakeParserTERMINATOR = 32
	MooncakeParserDECL_ID = 33
	MooncakeParserATTR_ID = 34
	MooncakeParserCTX_ID = 35
)

// MooncakeParser rules.
const (
	MooncakeParserRULE_mcrule = 0
	MooncakeParserRULE_statementList = 1
	MooncakeParserRULE_statement = 2
	MooncakeParserRULE_block = 3
	MooncakeParserRULE_linkedStmt = 4
	MooncakeParserRULE_simpleStmt = 5
	MooncakeParserRULE_inlineStmt = 6
	MooncakeParserRULE_errorStmt = 7
	MooncakeParserRULE_exprStmt = 8
	MooncakeParserRULE_identifier = 9
	MooncakeParserRULE_literal = 10
	MooncakeParserRULE_function = 11
	MooncakeParserRULE_operator = 12
	MooncakeParserRULE_errorType = 13
	MooncakeParserRULE_fatalError = 14
	MooncakeParserRULE_severeError = 15
	MooncakeParserRULE_warningError = 16
	MooncakeParserRULE_equalOperator = 17
	MooncakeParserRULE_notEqualOperator = 18
	MooncakeParserRULE_greaterThanOperator = 19
	MooncakeParserRULE_lessThanOperator = 20
	MooncakeParserRULE_greaterThanOrEqualOperator = 21
	MooncakeParserRULE_lessThanOrEqualOperator = 22
	MooncakeParserRULE_intLiteral = 23
	MooncakeParserRULE_floatLiteral = 24
	MooncakeParserRULE_boolLiteral = 25
	MooncakeParserRULE_nullLiteral = 26
	MooncakeParserRULE_stringLiteral = 27
	MooncakeParserRULE_ctxLiteral = 28
	MooncakeParserRULE_attributeIdentifier = 29
	MooncakeParserRULE_declarationIdentifier = 30
	MooncakeParserRULE_lengthFunction = 31
	MooncakeParserRULE_dateTimeLongFunction = 32
	MooncakeParserRULE_afterCurrentTimeFunction = 33
)

// IMcruleContext is an interface to support dynamic dispatch.
type IMcruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMcruleContext differentiates from other interfaces.
	IsMcruleContext()
}

type McruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMcruleContext() *McruleContext {
	var p = new(McruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_mcrule
	return p
}

func (*McruleContext) IsMcruleContext() {}

func NewMcruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *McruleContext {
	var p = new(McruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_mcrule

	return p
}

func (s *McruleContext) GetParser() antlr.Parser { return s.parser }

func (s *McruleContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *McruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *McruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *McruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterMcrule(s)
	}
}

func (s *McruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitMcrule(s)
	}
}

func (s *McruleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitMcrule(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) Mcrule() (localctx IMcruleContext) {
	localctx = NewMcruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MooncakeParserRULE_mcrule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.StatementList()
	}



	return localctx
}


// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(MooncakeParserCOMMENT)
}

func (s *StatementListContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(MooncakeParserCOMMENT, i)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MooncakeParserRULE_statementList)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << MooncakeParserT__0) | (1 << MooncakeParserLINKED) | (1 << MooncakeParserLEN_FUNC) | (1 << MooncakeParserDATETIME_LONG) | (1 << MooncakeParserAFTER_CURR_TIME))) != 0) || _la == MooncakeParserDECL_ID || _la == MooncakeParserATTR_ID {
		{
			p.SetState(70)
			p.Statement()
		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MooncakeParserCOMMENT {
			{
				p.SetState(71)
				p.Match(MooncakeParserCOMMENT)
			}

		}


		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) LinkedStmt() ILinkedStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILinkedStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILinkedStmtContext)
}

func (s *StatementContext) SimpleStmt() ISimpleStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MooncakeParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(82)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MooncakeParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.Block()
		}


	case MooncakeParserLINKED:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.LinkedStmt()
		}


	case MooncakeParserLEN_FUNC, MooncakeParserDATETIME_LONG, MooncakeParserAFTER_CURR_TIME, MooncakeParserDECL_ID, MooncakeParserATTR_ID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(81)
			p.SimpleStmt()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MooncakeParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(MooncakeParserT__0)
	}
	{
		p.SetState(85)
		p.StatementList()
	}
	{
		p.SetState(86)
		p.Match(MooncakeParserT__1)
	}



	return localctx
}


// ILinkedStmtContext is an interface to support dynamic dispatch.
type ILinkedStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLinkedStmtContext differentiates from other interfaces.
	IsLinkedStmtContext()
}

type LinkedStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLinkedStmtContext() *LinkedStmtContext {
	var p = new(LinkedStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_linkedStmt
	return p
}

func (*LinkedStmtContext) IsLinkedStmtContext() {}

func NewLinkedStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinkedStmtContext {
	var p = new(LinkedStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_linkedStmt

	return p
}

func (s *LinkedStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *LinkedStmtContext) LINKED() antlr.TerminalNode {
	return s.GetToken(MooncakeParserLINKED, 0)
}

func (s *LinkedStmtContext) SimpleStmt() ISimpleStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *LinkedStmtContext) LinkedStmt() ILinkedStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILinkedStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILinkedStmtContext)
}

func (s *LinkedStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinkedStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LinkedStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterLinkedStmt(s)
	}
}

func (s *LinkedStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitLinkedStmt(s)
	}
}

func (s *LinkedStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitLinkedStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) LinkedStmt() (localctx ILinkedStmtContext) {
	localctx = NewLinkedStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MooncakeParserRULE_linkedStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(MooncakeParserLINKED)
	}
	{
		p.SetState(89)
		p.SimpleStmt()
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(90)
			p.LinkedStmt()
		}


	}



	return localctx
}


// ISimpleStmtContext is an interface to support dynamic dispatch.
type ISimpleStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleStmtContext differentiates from other interfaces.
	IsSimpleStmtContext()
}

type SimpleStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStmtContext() *SimpleStmtContext {
	var p = new(SimpleStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_simpleStmt
	return p
}

func (*SimpleStmtContext) IsSimpleStmtContext() {}

func NewSimpleStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStmtContext {
	var p = new(SimpleStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_simpleStmt

	return p
}

func (s *SimpleStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStmtContext) ExprStmt() IExprStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprStmtContext)
}

func (s *SimpleStmtContext) ErrorStmt() IErrorStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IErrorStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IErrorStmtContext)
}

func (s *SimpleStmtContext) InlineStmt() IInlineStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineStmtContext)
}

func (s *SimpleStmtContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SimpleStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SimpleStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterSimpleStmt(s)
	}
}

func (s *SimpleStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitSimpleStmt(s)
	}
}

func (s *SimpleStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitSimpleStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) SimpleStmt() (localctx ISimpleStmtContext) {
	localctx = NewSimpleStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MooncakeParserRULE_simpleStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(94)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(93)
			p.InlineStmt()
		}


	}
	{
		p.SetState(96)
		p.ExprStmt()
	}
	{
		p.SetState(97)
		p.Match(MooncakeParserT__2)
	}
	{
		p.SetState(98)
		p.ErrorStmt()
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(99)
			p.Block()
		}


	}



	return localctx
}


// IInlineStmtContext is an interface to support dynamic dispatch.
type IInlineStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token 


	// SetId sets the id token.
	SetId(antlr.Token) 


	// GetFn returns the fn rule contexts.
	GetFn() IFunctionContext


	// SetFn sets the fn rule contexts.
	SetFn(IFunctionContext)


	// IsInlineStmtContext differentiates from other interfaces.
	IsInlineStmtContext()
}

type InlineStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id antlr.Token
	fn IFunctionContext 
}

func NewEmptyInlineStmtContext() *InlineStmtContext {
	var p = new(InlineStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_inlineStmt
	return p
}

func (*InlineStmtContext) IsInlineStmtContext() {}

func NewInlineStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineStmtContext {
	var p = new(InlineStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_inlineStmt

	return p
}

func (s *InlineStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineStmtContext) GetId() antlr.Token { return s.id }


func (s *InlineStmtContext) SetId(v antlr.Token) { s.id = v }


func (s *InlineStmtContext) GetFn() IFunctionContext { return s.fn }


func (s *InlineStmtContext) SetFn(v IFunctionContext) { s.fn = v }


func (s *InlineStmtContext) DECL_ID() antlr.TerminalNode {
	return s.GetToken(MooncakeParserDECL_ID, 0)
}

func (s *InlineStmtContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *InlineStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InlineStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterInlineStmt(s)
	}
}

func (s *InlineStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitInlineStmt(s)
	}
}

func (s *InlineStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitInlineStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) InlineStmt() (localctx IInlineStmtContext) {
	localctx = NewInlineStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MooncakeParserRULE_inlineStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)

			var _m = p.Match(MooncakeParserDECL_ID)

			localctx.(*InlineStmtContext).id = _m
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)

			var _x = p.Function()


			localctx.(*InlineStmtContext).fn = _x
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)

			var _m = p.Match(MooncakeParserDECL_ID)

			localctx.(*InlineStmtContext).id = _m
		}
		{
			p.SetState(105)

			var _x = p.Function()


			localctx.(*InlineStmtContext).fn = _x
		}

	}


	return localctx
}


// IErrorStmtContext is an interface to support dynamic dispatch.
type IErrorStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCode returns the code token.
	GetCode() antlr.Token 

	// GetInfo returns the info token.
	GetInfo() antlr.Token 


	// SetCode sets the code token.
	SetCode(antlr.Token) 

	// SetInfo sets the info token.
	SetInfo(antlr.Token) 


	// GetErrType returns the errType rule contexts.
	GetErrType() IErrorTypeContext


	// SetErrType sets the errType rule contexts.
	SetErrType(IErrorTypeContext)


	// IsErrorStmtContext differentiates from other interfaces.
	IsErrorStmtContext()
}

type ErrorStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	code antlr.Token
	info antlr.Token
	errType IErrorTypeContext 
}

func NewEmptyErrorStmtContext() *ErrorStmtContext {
	var p = new(ErrorStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_errorStmt
	return p
}

func (*ErrorStmtContext) IsErrorStmtContext() {}

func NewErrorStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ErrorStmtContext {
	var p = new(ErrorStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_errorStmt

	return p
}

func (s *ErrorStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ErrorStmtContext) GetCode() antlr.Token { return s.code }

func (s *ErrorStmtContext) GetInfo() antlr.Token { return s.info }


func (s *ErrorStmtContext) SetCode(v antlr.Token) { s.code = v }

func (s *ErrorStmtContext) SetInfo(v antlr.Token) { s.info = v }


func (s *ErrorStmtContext) GetErrType() IErrorTypeContext { return s.errType }


func (s *ErrorStmtContext) SetErrType(v IErrorTypeContext) { s.errType = v }


func (s *ErrorStmtContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(MooncakeParserSTRING)
}

func (s *ErrorStmtContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(MooncakeParserSTRING, i)
}

func (s *ErrorStmtContext) ErrorType() IErrorTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IErrorTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IErrorTypeContext)
}

func (s *ErrorStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrorStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ErrorStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterErrorStmt(s)
	}
}

func (s *ErrorStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitErrorStmt(s)
	}
}

func (s *ErrorStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitErrorStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) ErrorStmt() (localctx IErrorStmtContext) {
	localctx = NewErrorStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MooncakeParserRULE_errorStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(MooncakeParserT__3)
	}
	{
		p.SetState(109)

		var _m = p.Match(MooncakeParserSTRING)

		localctx.(*ErrorStmtContext).code = _m
	}
	{
		p.SetState(110)
		p.Match(MooncakeParserT__4)
	}
	{
		p.SetState(111)

		var _m = p.Match(MooncakeParserSTRING)

		localctx.(*ErrorStmtContext).info = _m
	}
	{
		p.SetState(112)
		p.Match(MooncakeParserT__5)
	}
	{
		p.SetState(113)

		var _x = p.ErrorType()


		localctx.(*ErrorStmtContext).errType = _x
	}



	return localctx
}


// IExprStmtContext is an interface to support dynamic dispatch.
type IExprStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id rule contexts.
	GetId() IIdentifierContext

	// GetOp returns the op rule contexts.
	GetOp() IOperatorContext

	// GetVal returns the val rule contexts.
	GetVal() ILiteralContext


	// SetId sets the id rule contexts.
	SetId(IIdentifierContext)

	// SetOp sets the op rule contexts.
	SetOp(IOperatorContext)

	// SetVal sets the val rule contexts.
	SetVal(ILiteralContext)


	// IsExprStmtContext differentiates from other interfaces.
	IsExprStmtContext()
}

type ExprStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id IIdentifierContext 
	op IOperatorContext 
	val ILiteralContext 
}

func NewEmptyExprStmtContext() *ExprStmtContext {
	var p = new(ExprStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_exprStmt
	return p
}

func (*ExprStmtContext) IsExprStmtContext() {}

func NewExprStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprStmtContext {
	var p = new(ExprStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_exprStmt

	return p
}

func (s *ExprStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprStmtContext) GetId() IIdentifierContext { return s.id }

func (s *ExprStmtContext) GetOp() IOperatorContext { return s.op }

func (s *ExprStmtContext) GetVal() ILiteralContext { return s.val }


func (s *ExprStmtContext) SetId(v IIdentifierContext) { s.id = v }

func (s *ExprStmtContext) SetOp(v IOperatorContext) { s.op = v }

func (s *ExprStmtContext) SetVal(v ILiteralContext) { s.val = v }


func (s *ExprStmtContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ExprStmtContext) Operator() IOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *ExprStmtContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExprStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterExprStmt(s)
	}
}

func (s *ExprStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitExprStmt(s)
	}
}

func (s *ExprStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitExprStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) ExprStmt() (localctx IExprStmtContext) {
	localctx = NewExprStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MooncakeParserRULE_exprStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)

		var _x = p.Identifier()


		localctx.(*ExprStmtContext).id = _x
	}
	{
		p.SetState(116)

		var _x = p.Operator()


		localctx.(*ExprStmtContext).op = _x
	}
	{
		p.SetState(117)

		var _x = p.Literal()


		localctx.(*ExprStmtContext).val = _x
	}



	return localctx
}


// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) AttributeIdentifier() IAttributeIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributeIdentifierContext)
}

func (s *IdentifierContext) DeclarationIdentifier() IDeclarationIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationIdentifierContext)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MooncakeParserRULE_identifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MooncakeParserATTR_ID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.AttributeIdentifier()
		}


	case MooncakeParserDECL_ID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.DeclarationIdentifier()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) IntLiteral() IIntLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntLiteralContext)
}

func (s *LiteralContext) FloatLiteral() IFloatLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatLiteralContext)
}

func (s *LiteralContext) BoolLiteral() IBoolLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolLiteralContext)
}

func (s *LiteralContext) NullLiteral() INullLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INullLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INullLiteralContext)
}

func (s *LiteralContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *LiteralContext) CtxLiteral() ICtxLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICtxLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICtxLiteralContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MooncakeParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(129)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MooncakeParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.IntLiteral()
		}


	case MooncakeParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.FloatLiteral()
		}


	case MooncakeParserBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.BoolLiteral()
		}


	case MooncakeParserNULL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(126)
			p.NullLiteral()
		}


	case MooncakeParserSTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(127)
			p.StringLiteral()
		}


	case MooncakeParserCTX_ID:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(128)
			p.CtxLiteral()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) LengthFunction() ILengthFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILengthFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILengthFunctionContext)
}

func (s *FunctionContext) DateTimeLongFunction() IDateTimeLongFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateTimeLongFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateTimeLongFunctionContext)
}

func (s *FunctionContext) AfterCurrentTimeFunction() IAfterCurrentTimeFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAfterCurrentTimeFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAfterCurrentTimeFunctionContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MooncakeParserRULE_function)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MooncakeParserLEN_FUNC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.LengthFunction()
		}


	case MooncakeParserDATETIME_LONG:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.DateTimeLongFunction()
		}


	case MooncakeParserAFTER_CURR_TIME:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(133)
			p.AfterCurrentTimeFunction()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_operator
	return p
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) EqualOperator() IEqualOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualOperatorContext)
}

func (s *OperatorContext) NotEqualOperator() INotEqualOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotEqualOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotEqualOperatorContext)
}

func (s *OperatorContext) GreaterThanOperator() IGreaterThanOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGreaterThanOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGreaterThanOperatorContext)
}

func (s *OperatorContext) LessThanOperator() ILessThanOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILessThanOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILessThanOperatorContext)
}

func (s *OperatorContext) GreaterThanOrEqualOperator() IGreaterThanOrEqualOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGreaterThanOrEqualOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGreaterThanOrEqualOperatorContext)
}

func (s *OperatorContext) LessThanOrEqualOperator() ILessThanOrEqualOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILessThanOrEqualOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILessThanOrEqualOperatorContext)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (s *OperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MooncakeParserRULE_operator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(142)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MooncakeParserEQ:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.EqualOperator()
		}


	case MooncakeParserNE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.NotEqualOperator()
		}


	case MooncakeParserGT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(138)
			p.GreaterThanOperator()
		}


	case MooncakeParserLT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(139)
			p.LessThanOperator()
		}


	case MooncakeParserGTE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(140)
			p.GreaterThanOrEqualOperator()
		}


	case MooncakeParserLTE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(141)
			p.LessThanOrEqualOperator()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IErrorTypeContext is an interface to support dynamic dispatch.
type IErrorTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsErrorTypeContext differentiates from other interfaces.
	IsErrorTypeContext()
}

type ErrorTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyErrorTypeContext() *ErrorTypeContext {
	var p = new(ErrorTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_errorType
	return p
}

func (*ErrorTypeContext) IsErrorTypeContext() {}

func NewErrorTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ErrorTypeContext {
	var p = new(ErrorTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_errorType

	return p
}

func (s *ErrorTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ErrorTypeContext) FatalError() IFatalErrorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFatalErrorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFatalErrorContext)
}

func (s *ErrorTypeContext) SevereError() ISevereErrorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISevereErrorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISevereErrorContext)
}

func (s *ErrorTypeContext) WarningError() IWarningErrorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWarningErrorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWarningErrorContext)
}

func (s *ErrorTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrorTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ErrorTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterErrorType(s)
	}
}

func (s *ErrorTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitErrorType(s)
	}
}

func (s *ErrorTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitErrorType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) ErrorType() (localctx IErrorTypeContext) {
	localctx = NewErrorTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MooncakeParserRULE_errorType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MooncakeParserFATAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.FatalError()
		}


	case MooncakeParserSEVERE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.SevereError()
		}


	case MooncakeParserWARNING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(146)
			p.WarningError()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IFatalErrorContext is an interface to support dynamic dispatch.
type IFatalErrorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFatalErrorContext differentiates from other interfaces.
	IsFatalErrorContext()
}

type FatalErrorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFatalErrorContext() *FatalErrorContext {
	var p = new(FatalErrorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_fatalError
	return p
}

func (*FatalErrorContext) IsFatalErrorContext() {}

func NewFatalErrorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FatalErrorContext {
	var p = new(FatalErrorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_fatalError

	return p
}

func (s *FatalErrorContext) GetParser() antlr.Parser { return s.parser }

func (s *FatalErrorContext) FATAL() antlr.TerminalNode {
	return s.GetToken(MooncakeParserFATAL, 0)
}

func (s *FatalErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FatalErrorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FatalErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterFatalError(s)
	}
}

func (s *FatalErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitFatalError(s)
	}
}

func (s *FatalErrorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitFatalError(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) FatalError() (localctx IFatalErrorContext) {
	localctx = NewFatalErrorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MooncakeParserRULE_fatalError)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(MooncakeParserFATAL)
	}



	return localctx
}


// ISevereErrorContext is an interface to support dynamic dispatch.
type ISevereErrorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSevereErrorContext differentiates from other interfaces.
	IsSevereErrorContext()
}

type SevereErrorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySevereErrorContext() *SevereErrorContext {
	var p = new(SevereErrorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_severeError
	return p
}

func (*SevereErrorContext) IsSevereErrorContext() {}

func NewSevereErrorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SevereErrorContext {
	var p = new(SevereErrorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_severeError

	return p
}

func (s *SevereErrorContext) GetParser() antlr.Parser { return s.parser }

func (s *SevereErrorContext) SEVERE() antlr.TerminalNode {
	return s.GetToken(MooncakeParserSEVERE, 0)
}

func (s *SevereErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SevereErrorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SevereErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterSevereError(s)
	}
}

func (s *SevereErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitSevereError(s)
	}
}

func (s *SevereErrorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitSevereError(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) SevereError() (localctx ISevereErrorContext) {
	localctx = NewSevereErrorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MooncakeParserRULE_severeError)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(MooncakeParserSEVERE)
	}



	return localctx
}


// IWarningErrorContext is an interface to support dynamic dispatch.
type IWarningErrorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWarningErrorContext differentiates from other interfaces.
	IsWarningErrorContext()
}

type WarningErrorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWarningErrorContext() *WarningErrorContext {
	var p = new(WarningErrorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_warningError
	return p
}

func (*WarningErrorContext) IsWarningErrorContext() {}

func NewWarningErrorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WarningErrorContext {
	var p = new(WarningErrorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_warningError

	return p
}

func (s *WarningErrorContext) GetParser() antlr.Parser { return s.parser }

func (s *WarningErrorContext) WARNING() antlr.TerminalNode {
	return s.GetToken(MooncakeParserWARNING, 0)
}

func (s *WarningErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WarningErrorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *WarningErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterWarningError(s)
	}
}

func (s *WarningErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitWarningError(s)
	}
}

func (s *WarningErrorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitWarningError(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) WarningError() (localctx IWarningErrorContext) {
	localctx = NewWarningErrorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MooncakeParserRULE_warningError)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(MooncakeParserWARNING)
	}



	return localctx
}


// IEqualOperatorContext is an interface to support dynamic dispatch.
type IEqualOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualOperatorContext differentiates from other interfaces.
	IsEqualOperatorContext()
}

type EqualOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualOperatorContext() *EqualOperatorContext {
	var p = new(EqualOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_equalOperator
	return p
}

func (*EqualOperatorContext) IsEqualOperatorContext() {}

func NewEqualOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualOperatorContext {
	var p = new(EqualOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_equalOperator

	return p
}

func (s *EqualOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualOperatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(MooncakeParserEQ, 0)
}

func (s *EqualOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *EqualOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterEqualOperator(s)
	}
}

func (s *EqualOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitEqualOperator(s)
	}
}

func (s *EqualOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitEqualOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) EqualOperator() (localctx IEqualOperatorContext) {
	localctx = NewEqualOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MooncakeParserRULE_equalOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(MooncakeParserEQ)
	}



	return localctx
}


// INotEqualOperatorContext is an interface to support dynamic dispatch.
type INotEqualOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotEqualOperatorContext differentiates from other interfaces.
	IsNotEqualOperatorContext()
}

type NotEqualOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotEqualOperatorContext() *NotEqualOperatorContext {
	var p = new(NotEqualOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_notEqualOperator
	return p
}

func (*NotEqualOperatorContext) IsNotEqualOperatorContext() {}

func NewNotEqualOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotEqualOperatorContext {
	var p = new(NotEqualOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_notEqualOperator

	return p
}

func (s *NotEqualOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *NotEqualOperatorContext) NE() antlr.TerminalNode {
	return s.GetToken(MooncakeParserNE, 0)
}

func (s *NotEqualOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotEqualOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NotEqualOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterNotEqualOperator(s)
	}
}

func (s *NotEqualOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitNotEqualOperator(s)
	}
}

func (s *NotEqualOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitNotEqualOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) NotEqualOperator() (localctx INotEqualOperatorContext) {
	localctx = NewNotEqualOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MooncakeParserRULE_notEqualOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(MooncakeParserNE)
	}



	return localctx
}


// IGreaterThanOperatorContext is an interface to support dynamic dispatch.
type IGreaterThanOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGreaterThanOperatorContext differentiates from other interfaces.
	IsGreaterThanOperatorContext()
}

type GreaterThanOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGreaterThanOperatorContext() *GreaterThanOperatorContext {
	var p = new(GreaterThanOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_greaterThanOperator
	return p
}

func (*GreaterThanOperatorContext) IsGreaterThanOperatorContext() {}

func NewGreaterThanOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GreaterThanOperatorContext {
	var p = new(GreaterThanOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_greaterThanOperator

	return p
}

func (s *GreaterThanOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *GreaterThanOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(MooncakeParserGT, 0)
}

func (s *GreaterThanOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *GreaterThanOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterGreaterThanOperator(s)
	}
}

func (s *GreaterThanOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitGreaterThanOperator(s)
	}
}

func (s *GreaterThanOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitGreaterThanOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) GreaterThanOperator() (localctx IGreaterThanOperatorContext) {
	localctx = NewGreaterThanOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MooncakeParserRULE_greaterThanOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(MooncakeParserGT)
	}



	return localctx
}


// ILessThanOperatorContext is an interface to support dynamic dispatch.
type ILessThanOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLessThanOperatorContext differentiates from other interfaces.
	IsLessThanOperatorContext()
}

type LessThanOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLessThanOperatorContext() *LessThanOperatorContext {
	var p = new(LessThanOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_lessThanOperator
	return p
}

func (*LessThanOperatorContext) IsLessThanOperatorContext() {}

func NewLessThanOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LessThanOperatorContext {
	var p = new(LessThanOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_lessThanOperator

	return p
}

func (s *LessThanOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LessThanOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(MooncakeParserLT, 0)
}

func (s *LessThanOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LessThanOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterLessThanOperator(s)
	}
}

func (s *LessThanOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitLessThanOperator(s)
	}
}

func (s *LessThanOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitLessThanOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) LessThanOperator() (localctx ILessThanOperatorContext) {
	localctx = NewLessThanOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MooncakeParserRULE_lessThanOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(MooncakeParserLT)
	}



	return localctx
}


// IGreaterThanOrEqualOperatorContext is an interface to support dynamic dispatch.
type IGreaterThanOrEqualOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGreaterThanOrEqualOperatorContext differentiates from other interfaces.
	IsGreaterThanOrEqualOperatorContext()
}

type GreaterThanOrEqualOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGreaterThanOrEqualOperatorContext() *GreaterThanOrEqualOperatorContext {
	var p = new(GreaterThanOrEqualOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_greaterThanOrEqualOperator
	return p
}

func (*GreaterThanOrEqualOperatorContext) IsGreaterThanOrEqualOperatorContext() {}

func NewGreaterThanOrEqualOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GreaterThanOrEqualOperatorContext {
	var p = new(GreaterThanOrEqualOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_greaterThanOrEqualOperator

	return p
}

func (s *GreaterThanOrEqualOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *GreaterThanOrEqualOperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(MooncakeParserGTE, 0)
}

func (s *GreaterThanOrEqualOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanOrEqualOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *GreaterThanOrEqualOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterGreaterThanOrEqualOperator(s)
	}
}

func (s *GreaterThanOrEqualOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitGreaterThanOrEqualOperator(s)
	}
}

func (s *GreaterThanOrEqualOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitGreaterThanOrEqualOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) GreaterThanOrEqualOperator() (localctx IGreaterThanOrEqualOperatorContext) {
	localctx = NewGreaterThanOrEqualOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MooncakeParserRULE_greaterThanOrEqualOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(MooncakeParserGTE)
	}



	return localctx
}


// ILessThanOrEqualOperatorContext is an interface to support dynamic dispatch.
type ILessThanOrEqualOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLessThanOrEqualOperatorContext differentiates from other interfaces.
	IsLessThanOrEqualOperatorContext()
}

type LessThanOrEqualOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLessThanOrEqualOperatorContext() *LessThanOrEqualOperatorContext {
	var p = new(LessThanOrEqualOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_lessThanOrEqualOperator
	return p
}

func (*LessThanOrEqualOperatorContext) IsLessThanOrEqualOperatorContext() {}

func NewLessThanOrEqualOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LessThanOrEqualOperatorContext {
	var p = new(LessThanOrEqualOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_lessThanOrEqualOperator

	return p
}

func (s *LessThanOrEqualOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LessThanOrEqualOperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(MooncakeParserLTE, 0)
}

func (s *LessThanOrEqualOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanOrEqualOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LessThanOrEqualOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterLessThanOrEqualOperator(s)
	}
}

func (s *LessThanOrEqualOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitLessThanOrEqualOperator(s)
	}
}

func (s *LessThanOrEqualOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitLessThanOrEqualOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) LessThanOrEqualOperator() (localctx ILessThanOrEqualOperatorContext) {
	localctx = NewLessThanOrEqualOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MooncakeParserRULE_lessThanOrEqualOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(MooncakeParserLTE)
	}



	return localctx
}


// IIntLiteralContext is an interface to support dynamic dispatch.
type IIntLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntLiteralContext differentiates from other interfaces.
	IsIntLiteralContext()
}

type IntLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntLiteralContext() *IntLiteralContext {
	var p = new(IntLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_intLiteral
	return p
}

func (*IntLiteralContext) IsIntLiteralContext() {}

func NewIntLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntLiteralContext {
	var p = new(IntLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_intLiteral

	return p
}

func (s *IntLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntLiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(MooncakeParserINT, 0)
}

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IntLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterIntLiteral(s)
	}
}

func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitIntLiteral(s)
	}
}

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitIntLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) IntLiteral() (localctx IIntLiteralContext) {
	localctx = NewIntLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MooncakeParserRULE_intLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(MooncakeParserINT)
	}



	return localctx
}


// IFloatLiteralContext is an interface to support dynamic dispatch.
type IFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatLiteralContext differentiates from other interfaces.
	IsFloatLiteralContext()
}

type FloatLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLiteralContext() *FloatLiteralContext {
	var p = new(FloatLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_floatLiteral
	return p
}

func (*FloatLiteralContext) IsFloatLiteralContext() {}

func NewFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_floatLiteral

	return p
}

func (s *FloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLiteralContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(MooncakeParserFLOAT, 0)
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) FloatLiteral() (localctx IFloatLiteralContext) {
	localctx = NewFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MooncakeParserRULE_floatLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(MooncakeParserFLOAT)
	}



	return localctx
}


// IBoolLiteralContext is an interface to support dynamic dispatch.
type IBoolLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolLiteralContext differentiates from other interfaces.
	IsBoolLiteralContext()
}

type BoolLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolLiteralContext() *BoolLiteralContext {
	var p = new(BoolLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_boolLiteral
	return p
}

func (*BoolLiteralContext) IsBoolLiteralContext() {}

func NewBoolLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolLiteralContext {
	var p = new(BoolLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_boolLiteral

	return p
}

func (s *BoolLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolLiteralContext) BOOL() antlr.TerminalNode {
	return s.GetToken(MooncakeParserBOOL, 0)
}

func (s *BoolLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BoolLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitBoolLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) BoolLiteral() (localctx IBoolLiteralContext) {
	localctx = NewBoolLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MooncakeParserRULE_boolLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(MooncakeParserBOOL)
	}



	return localctx
}


// INullLiteralContext is an interface to support dynamic dispatch.
type INullLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNullLiteralContext differentiates from other interfaces.
	IsNullLiteralContext()
}

type NullLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullLiteralContext() *NullLiteralContext {
	var p = new(NullLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_nullLiteral
	return p
}

func (*NullLiteralContext) IsNullLiteralContext() {}

func NewNullLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullLiteralContext {
	var p = new(NullLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_nullLiteral

	return p
}

func (s *NullLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NullLiteralContext) NULL() antlr.TerminalNode {
	return s.GetToken(MooncakeParserNULL, 0)
}

func (s *NullLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NullLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterNullLiteral(s)
	}
}

func (s *NullLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitNullLiteral(s)
	}
}

func (s *NullLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitNullLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) NullLiteral() (localctx INullLiteralContext) {
	localctx = NewNullLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MooncakeParserRULE_nullLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(MooncakeParserNULL)
	}



	return localctx
}


// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(MooncakeParserSTRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MooncakeParserRULE_stringLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(MooncakeParserSTRING)
	}



	return localctx
}


// ICtxLiteralContext is an interface to support dynamic dispatch.
type ICtxLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCtxLiteralContext differentiates from other interfaces.
	IsCtxLiteralContext()
}

type CtxLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCtxLiteralContext() *CtxLiteralContext {
	var p = new(CtxLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_ctxLiteral
	return p
}

func (*CtxLiteralContext) IsCtxLiteralContext() {}

func NewCtxLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CtxLiteralContext {
	var p = new(CtxLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_ctxLiteral

	return p
}

func (s *CtxLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *CtxLiteralContext) CTX_ID() antlr.TerminalNode {
	return s.GetToken(MooncakeParserCTX_ID, 0)
}

func (s *CtxLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CtxLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CtxLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterCtxLiteral(s)
	}
}

func (s *CtxLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitCtxLiteral(s)
	}
}

func (s *CtxLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitCtxLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) CtxLiteral() (localctx ICtxLiteralContext) {
	localctx = NewCtxLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, MooncakeParserRULE_ctxLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(MooncakeParserCTX_ID)
	}



	return localctx
}


// IAttributeIdentifierContext is an interface to support dynamic dispatch.
type IAttributeIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeIdentifierContext differentiates from other interfaces.
	IsAttributeIdentifierContext()
}

type AttributeIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeIdentifierContext() *AttributeIdentifierContext {
	var p = new(AttributeIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_attributeIdentifier
	return p
}

func (*AttributeIdentifierContext) IsAttributeIdentifierContext() {}

func NewAttributeIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeIdentifierContext {
	var p = new(AttributeIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_attributeIdentifier

	return p
}

func (s *AttributeIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeIdentifierContext) ATTR_ID() antlr.TerminalNode {
	return s.GetToken(MooncakeParserATTR_ID, 0)
}

func (s *AttributeIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AttributeIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterAttributeIdentifier(s)
	}
}

func (s *AttributeIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitAttributeIdentifier(s)
	}
}

func (s *AttributeIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitAttributeIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) AttributeIdentifier() (localctx IAttributeIdentifierContext) {
	localctx = NewAttributeIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, MooncakeParserRULE_attributeIdentifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(MooncakeParserATTR_ID)
	}



	return localctx
}


// IDeclarationIdentifierContext is an interface to support dynamic dispatch.
type IDeclarationIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationIdentifierContext differentiates from other interfaces.
	IsDeclarationIdentifierContext()
}

type DeclarationIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationIdentifierContext() *DeclarationIdentifierContext {
	var p = new(DeclarationIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_declarationIdentifier
	return p
}

func (*DeclarationIdentifierContext) IsDeclarationIdentifierContext() {}

func NewDeclarationIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationIdentifierContext {
	var p = new(DeclarationIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_declarationIdentifier

	return p
}

func (s *DeclarationIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationIdentifierContext) DECL_ID() antlr.TerminalNode {
	return s.GetToken(MooncakeParserDECL_ID, 0)
}

func (s *DeclarationIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DeclarationIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterDeclarationIdentifier(s)
	}
}

func (s *DeclarationIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitDeclarationIdentifier(s)
	}
}

func (s *DeclarationIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitDeclarationIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) DeclarationIdentifier() (localctx IDeclarationIdentifierContext) {
	localctx = NewDeclarationIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, MooncakeParserRULE_declarationIdentifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(MooncakeParserDECL_ID)
	}



	return localctx
}


// ILengthFunctionContext is an interface to support dynamic dispatch.
type ILengthFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLengthFunctionContext differentiates from other interfaces.
	IsLengthFunctionContext()
}

type LengthFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthFunctionContext() *LengthFunctionContext {
	var p = new(LengthFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_lengthFunction
	return p
}

func (*LengthFunctionContext) IsLengthFunctionContext() {}

func NewLengthFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthFunctionContext {
	var p = new(LengthFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_lengthFunction

	return p
}

func (s *LengthFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthFunctionContext) LEN_FUNC() antlr.TerminalNode {
	return s.GetToken(MooncakeParserLEN_FUNC, 0)
}

func (s *LengthFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LengthFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterLengthFunction(s)
	}
}

func (s *LengthFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitLengthFunction(s)
	}
}

func (s *LengthFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitLengthFunction(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) LengthFunction() (localctx ILengthFunctionContext) {
	localctx = NewLengthFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, MooncakeParserRULE_lengthFunction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(MooncakeParserLEN_FUNC)
	}



	return localctx
}


// IDateTimeLongFunctionContext is an interface to support dynamic dispatch.
type IDateTimeLongFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateTimeLongFunctionContext differentiates from other interfaces.
	IsDateTimeLongFunctionContext()
}

type DateTimeLongFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateTimeLongFunctionContext() *DateTimeLongFunctionContext {
	var p = new(DateTimeLongFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_dateTimeLongFunction
	return p
}

func (*DateTimeLongFunctionContext) IsDateTimeLongFunctionContext() {}

func NewDateTimeLongFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateTimeLongFunctionContext {
	var p = new(DateTimeLongFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_dateTimeLongFunction

	return p
}

func (s *DateTimeLongFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *DateTimeLongFunctionContext) DATETIME_LONG() antlr.TerminalNode {
	return s.GetToken(MooncakeParserDATETIME_LONG, 0)
}

func (s *DateTimeLongFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateTimeLongFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DateTimeLongFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterDateTimeLongFunction(s)
	}
}

func (s *DateTimeLongFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitDateTimeLongFunction(s)
	}
}

func (s *DateTimeLongFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitDateTimeLongFunction(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) DateTimeLongFunction() (localctx IDateTimeLongFunctionContext) {
	localctx = NewDateTimeLongFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, MooncakeParserRULE_dateTimeLongFunction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(MooncakeParserDATETIME_LONG)
	}



	return localctx
}


// IAfterCurrentTimeFunctionContext is an interface to support dynamic dispatch.
type IAfterCurrentTimeFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAfterCurrentTimeFunctionContext differentiates from other interfaces.
	IsAfterCurrentTimeFunctionContext()
}

type AfterCurrentTimeFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAfterCurrentTimeFunctionContext() *AfterCurrentTimeFunctionContext {
	var p = new(AfterCurrentTimeFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_afterCurrentTimeFunction
	return p
}

func (*AfterCurrentTimeFunctionContext) IsAfterCurrentTimeFunctionContext() {}

func NewAfterCurrentTimeFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AfterCurrentTimeFunctionContext {
	var p = new(AfterCurrentTimeFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_afterCurrentTimeFunction

	return p
}

func (s *AfterCurrentTimeFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *AfterCurrentTimeFunctionContext) AFTER_CURR_TIME() antlr.TerminalNode {
	return s.GetToken(MooncakeParserAFTER_CURR_TIME, 0)
}

func (s *AfterCurrentTimeFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AfterCurrentTimeFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AfterCurrentTimeFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterAfterCurrentTimeFunction(s)
	}
}

func (s *AfterCurrentTimeFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitAfterCurrentTimeFunction(s)
	}
}

func (s *AfterCurrentTimeFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitAfterCurrentTimeFunction(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) AfterCurrentTimeFunction() (localctx IAfterCurrentTimeFunctionContext) {
	localctx = NewAfterCurrentTimeFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, MooncakeParserRULE_afterCurrentTimeFunction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(MooncakeParserAFTER_CURR_TIME)
	}



	return localctx
}


