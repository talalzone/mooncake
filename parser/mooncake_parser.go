// Code generated from /Users/talal/Development/antlr-lang/mooncake/grammar/Mooncake.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 39, 90, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4, 
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9, 
	13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 
	4, 3, 4, 5, 4, 39, 10, 4, 7, 4, 41, 10, 4, 12, 4, 14, 4, 44, 11, 4, 3, 
	5, 3, 5, 3, 5, 5, 5, 49, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 55, 10, 6, 
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 5, 8, 65, 10, 8, 3, 9, 
	5, 9, 68, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 74, 10, 9, 3, 10, 3, 10, 
	3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 
	15, 3, 15, 3, 15, 2, 2, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 
	26, 28, 2, 7, 4, 2, 37, 37, 39, 39, 5, 2, 27, 29, 32, 32, 38, 38, 3, 2, 
	21, 24, 4, 2, 13, 14, 17, 20, 3, 2, 7, 9, 2, 84, 2, 30, 3, 2, 2, 2, 4, 
	32, 3, 2, 2, 2, 6, 42, 3, 2, 2, 2, 8, 48, 3, 2, 2, 2, 10, 54, 3, 2, 2, 
	2, 12, 56, 3, 2, 2, 2, 14, 61, 3, 2, 2, 2, 16, 67, 3, 2, 2, 2, 18, 75, 
	3, 2, 2, 2, 20, 79, 3, 2, 2, 2, 22, 81, 3, 2, 2, 2, 24, 83, 3, 2, 2, 2, 
	26, 85, 3, 2, 2, 2, 28, 87, 3, 2, 2, 2, 30, 31, 5, 6, 4, 2, 31, 3, 3, 2, 
	2, 2, 32, 33, 7, 3, 2, 2, 33, 34, 5, 6, 4, 2, 34, 35, 7, 4, 2, 2, 35, 5, 
	3, 2, 2, 2, 36, 38, 5, 8, 5, 2, 37, 39, 7, 34, 2, 2, 38, 37, 3, 2, 2, 2, 
	38, 39, 3, 2, 2, 2, 39, 41, 3, 2, 2, 2, 40, 36, 3, 2, 2, 2, 41, 44, 3, 
	2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 7, 3, 2, 2, 2, 44, 
	42, 3, 2, 2, 2, 45, 49, 5, 4, 3, 2, 46, 49, 5, 14, 8, 2, 47, 49, 5, 16, 
	9, 2, 48, 45, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 47, 3, 2, 2, 2, 49, 9, 
	3, 2, 2, 2, 50, 55, 7, 39, 2, 2, 51, 55, 5, 24, 13, 2, 52, 53, 7, 39, 2, 
	2, 53, 55, 5, 24, 13, 2, 54, 50, 3, 2, 2, 2, 54, 51, 3, 2, 2, 2, 54, 52, 
	3, 2, 2, 2, 55, 11, 3, 2, 2, 2, 56, 57, 7, 33, 2, 2, 57, 58, 7, 5, 2, 2, 
	58, 59, 7, 28, 2, 2, 59, 60, 5, 28, 15, 2, 60, 13, 3, 2, 2, 2, 61, 62, 
	7, 10, 2, 2, 62, 64, 5, 16, 9, 2, 63, 65, 5, 14, 8, 2, 64, 63, 3, 2, 2, 
	2, 64, 65, 3, 2, 2, 2, 65, 15, 3, 2, 2, 2, 66, 68, 5, 10, 6, 2, 67, 66, 
	3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 70, 5, 18, 10, 
	2, 70, 71, 7, 6, 2, 2, 71, 73, 5, 12, 7, 2, 72, 74, 5, 4, 3, 2, 73, 72, 
	3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 17, 3, 2, 2, 2, 75, 76, 5, 20, 11, 
	2, 76, 77, 5, 26, 14, 2, 77, 78, 5, 22, 12, 2, 78, 19, 3, 2, 2, 2, 79, 
	80, 9, 2, 2, 2, 80, 21, 3, 2, 2, 2, 81, 82, 9, 3, 2, 2, 82, 23, 3, 2, 2, 
	2, 83, 84, 9, 4, 2, 2, 84, 25, 3, 2, 2, 2, 85, 86, 9, 5, 2, 2, 86, 27, 
	3, 2, 2, 2, 87, 88, 9, 6, 2, 2, 88, 29, 3, 2, 2, 2, 9, 38, 42, 48, 54, 
	64, 67, 73,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "':'", "'=>'", "'!!!'", "'!!'", "'!'", "'~'", "'if'", 
	"'elif'", "", "", "", "", "", "", "", "", "'@len'", "'@float'", "'@dateTimeLong'", 
	"'@afterCurrentTime'", "'exists'", "'empty'", "", "", "", "'true'", "'false'", 
	"'nil'",
}
var symbolicNames = []string{
	"", "", "", "", "", "FATAL", "SEVERE", "WARNING", "LINKED", "IF", "ELIF", 
	"EQ", "NE", "AND", "OR", "GT", "LT", "GTE", "LTE", "LEN_FUNC", "FLOAT_FUNC", 
	"DATETIME_LONG", "AFTER_CURR_TIME", "EXISTS", "EMPTY", "DOUBLE", "INT", 
	"BOOL", "TRUE", "FALSE", "NULL", "ERROR_CODE", "COMMENT", "WS", "TERMINATOR", 
	"IDENTIFIER", "CTX_ID", "INLINE_ID",
}

var ruleNames = []string{
	"mcrule", "block", "statementList", "statement", "inlineStmt", "errorStmt", 
	"linkedStmt", "simpleStmt", "expression", "identifier", "literal", "function", 
	"operator", "errorType",
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
	MooncakeParserFATAL = 5
	MooncakeParserSEVERE = 6
	MooncakeParserWARNING = 7
	MooncakeParserLINKED = 8
	MooncakeParserIF = 9
	MooncakeParserELIF = 10
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
	MooncakeParserEXISTS = 23
	MooncakeParserEMPTY = 24
	MooncakeParserDOUBLE = 25
	MooncakeParserINT = 26
	MooncakeParserBOOL = 27
	MooncakeParserTRUE = 28
	MooncakeParserFALSE = 29
	MooncakeParserNULL = 30
	MooncakeParserERROR_CODE = 31
	MooncakeParserCOMMENT = 32
	MooncakeParserWS = 33
	MooncakeParserTERMINATOR = 34
	MooncakeParserIDENTIFIER = 35
	MooncakeParserCTX_ID = 36
	MooncakeParserINLINE_ID = 37
)

// MooncakeParser rules.
const (
	MooncakeParserRULE_mcrule = 0
	MooncakeParserRULE_block = 1
	MooncakeParserRULE_statementList = 2
	MooncakeParserRULE_statement = 3
	MooncakeParserRULE_inlineStmt = 4
	MooncakeParserRULE_errorStmt = 5
	MooncakeParserRULE_linkedStmt = 6
	MooncakeParserRULE_simpleStmt = 7
	MooncakeParserRULE_expression = 8
	MooncakeParserRULE_identifier = 9
	MooncakeParserRULE_literal = 10
	MooncakeParserRULE_function = 11
	MooncakeParserRULE_operator = 12
	MooncakeParserRULE_errorType = 13
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
		p.SetState(28)
		p.StatementList()
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
	p.EnterRule(localctx, 2, MooncakeParserRULE_block)

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
		p.SetState(30)
		p.Match(MooncakeParserT__0)
	}
	{
		p.SetState(31)
		p.StatementList()
	}
	{
		p.SetState(32)
		p.Match(MooncakeParserT__1)
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
	p.EnterRule(localctx, 4, MooncakeParserRULE_statementList)
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
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << MooncakeParserT__0) | (1 << MooncakeParserLINKED) | (1 << MooncakeParserLEN_FUNC) | (1 << MooncakeParserFLOAT_FUNC) | (1 << MooncakeParserDATETIME_LONG) | (1 << MooncakeParserAFTER_CURR_TIME))) != 0) || _la == MooncakeParserIDENTIFIER || _la == MooncakeParserINLINE_ID {
		{
			p.SetState(34)
			p.Statement()
		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MooncakeParserCOMMENT {
			{
				p.SetState(35)
				p.Match(MooncakeParserCOMMENT)
			}

		}


		p.SetState(42)
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
	p.EnterRule(localctx, 6, MooncakeParserRULE_statement)

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

	p.SetState(46)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MooncakeParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.Block()
		}


	case MooncakeParserLINKED:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.LinkedStmt()
		}


	case MooncakeParserLEN_FUNC, MooncakeParserFLOAT_FUNC, MooncakeParserDATETIME_LONG, MooncakeParserAFTER_CURR_TIME, MooncakeParserIDENTIFIER, MooncakeParserINLINE_ID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(45)
			p.SimpleStmt()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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


func (s *InlineStmtContext) INLINE_ID() antlr.TerminalNode {
	return s.GetToken(MooncakeParserINLINE_ID, 0)
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
	p.EnterRule(localctx, 8, MooncakeParserRULE_inlineStmt)

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

	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)

			var _m = p.Match(MooncakeParserINLINE_ID)

			localctx.(*InlineStmtContext).id = _m
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)

			var _x = p.Function()


			localctx.(*InlineStmtContext).fn = _x
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)

			var _m = p.Match(MooncakeParserINLINE_ID)

			localctx.(*InlineStmtContext).id = _m
		}
		{
			p.SetState(51)

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

	// GetHumanCode returns the humanCode token.
	GetHumanCode() antlr.Token 

	// GetErrCode returns the errCode token.
	GetErrCode() antlr.Token 


	// SetHumanCode sets the humanCode token.
	SetHumanCode(antlr.Token) 

	// SetErrCode sets the errCode token.
	SetErrCode(antlr.Token) 


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
	humanCode antlr.Token
	errCode antlr.Token
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

func (s *ErrorStmtContext) GetHumanCode() antlr.Token { return s.humanCode }

func (s *ErrorStmtContext) GetErrCode() antlr.Token { return s.errCode }


func (s *ErrorStmtContext) SetHumanCode(v antlr.Token) { s.humanCode = v }

func (s *ErrorStmtContext) SetErrCode(v antlr.Token) { s.errCode = v }


func (s *ErrorStmtContext) GetErrType() IErrorTypeContext { return s.errType }


func (s *ErrorStmtContext) SetErrType(v IErrorTypeContext) { s.errType = v }


func (s *ErrorStmtContext) ERROR_CODE() antlr.TerminalNode {
	return s.GetToken(MooncakeParserERROR_CODE, 0)
}

func (s *ErrorStmtContext) INT() antlr.TerminalNode {
	return s.GetToken(MooncakeParserINT, 0)
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
	p.EnterRule(localctx, 10, MooncakeParserRULE_errorStmt)

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
		p.SetState(54)

		var _m = p.Match(MooncakeParserERROR_CODE)

		localctx.(*ErrorStmtContext).humanCode = _m
	}
	{
		p.SetState(55)
		p.Match(MooncakeParserT__2)
	}
	{
		p.SetState(56)

		var _m = p.Match(MooncakeParserINT)

		localctx.(*ErrorStmtContext).errCode = _m
	}
	{
		p.SetState(57)

		var _x = p.ErrorType()


		localctx.(*ErrorStmtContext).errType = _x
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
	p.EnterRule(localctx, 12, MooncakeParserRULE_linkedStmt)

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
		p.SetState(59)
		p.Match(MooncakeParserLINKED)
	}
	{
		p.SetState(60)
		p.SimpleStmt()
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(61)
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

func (s *SimpleStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
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
	p.EnterRule(localctx, 14, MooncakeParserRULE_simpleStmt)

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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(64)
			p.InlineStmt()
		}


	}
	{
		p.SetState(67)
		p.Expression()
	}
	{
		p.SetState(68)
		p.Match(MooncakeParserT__3)
	}
	{
		p.SetState(69)
		p.ErrorStmt()
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(70)
			p.Block()
		}


	}



	return localctx
}


// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
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


	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id IIdentifierContext 
	op IOperatorContext 
	val ILiteralContext 
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MooncakeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MooncakeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetId() IIdentifierContext { return s.id }

func (s *ExpressionContext) GetOp() IOperatorContext { return s.op }

func (s *ExpressionContext) GetVal() ILiteralContext { return s.val }


func (s *ExpressionContext) SetId(v IIdentifierContext) { s.id = v }

func (s *ExpressionContext) SetOp(v IOperatorContext) { s.op = v }

func (s *ExpressionContext) SetVal(v ILiteralContext) { s.val = v }


func (s *ExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ExpressionContext) Operator() IOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *ExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MooncakeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MooncakeVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MooncakeParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MooncakeParserRULE_expression)

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
		p.SetState(73)

		var _x = p.Identifier()


		localctx.(*ExpressionContext).id = _x
	}
	{
		p.SetState(74)

		var _x = p.Operator()


		localctx.(*ExpressionContext).op = _x
	}
	{
		p.SetState(75)

		var _x = p.Literal()


		localctx.(*ExpressionContext).val = _x
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

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MooncakeParserIDENTIFIER, 0)
}

func (s *IdentifierContext) INLINE_ID() antlr.TerminalNode {
	return s.GetToken(MooncakeParserINLINE_ID, 0)
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
	{
		p.SetState(77)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MooncakeParserIDENTIFIER || _la == MooncakeParserINLINE_ID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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

func (s *LiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(MooncakeParserINT, 0)
}

func (s *LiteralContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(MooncakeParserDOUBLE, 0)
}

func (s *LiteralContext) BOOL() antlr.TerminalNode {
	return s.GetToken(MooncakeParserBOOL, 0)
}

func (s *LiteralContext) NULL() antlr.TerminalNode {
	return s.GetToken(MooncakeParserNULL, 0)
}

func (s *LiteralContext) CTX_ID() antlr.TerminalNode {
	return s.GetToken(MooncakeParserCTX_ID, 0)
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
	{
		p.SetState(79)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 25)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 25))) & ((1 << (MooncakeParserDOUBLE - 25)) | (1 << (MooncakeParserINT - 25)) | (1 << (MooncakeParserBOOL - 25)) | (1 << (MooncakeParserNULL - 25)) | (1 << (MooncakeParserCTX_ID - 25)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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

func (s *FunctionContext) LEN_FUNC() antlr.TerminalNode {
	return s.GetToken(MooncakeParserLEN_FUNC, 0)
}

func (s *FunctionContext) FLOAT_FUNC() antlr.TerminalNode {
	return s.GetToken(MooncakeParserFLOAT_FUNC, 0)
}

func (s *FunctionContext) DATETIME_LONG() antlr.TerminalNode {
	return s.GetToken(MooncakeParserDATETIME_LONG, 0)
}

func (s *FunctionContext) AFTER_CURR_TIME() antlr.TerminalNode {
	return s.GetToken(MooncakeParserAFTER_CURR_TIME, 0)
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
	{
		p.SetState(81)
		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << MooncakeParserLEN_FUNC) | (1 << MooncakeParserFLOAT_FUNC) | (1 << MooncakeParserDATETIME_LONG) | (1 << MooncakeParserAFTER_CURR_TIME))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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

func (s *OperatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(MooncakeParserEQ, 0)
}

func (s *OperatorContext) NE() antlr.TerminalNode {
	return s.GetToken(MooncakeParserNE, 0)
}

func (s *OperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(MooncakeParserGT, 0)
}

func (s *OperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(MooncakeParserLT, 0)
}

func (s *OperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(MooncakeParserGTE, 0)
}

func (s *OperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(MooncakeParserLTE, 0)
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
	{
		p.SetState(83)
		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << MooncakeParserEQ) | (1 << MooncakeParserNE) | (1 << MooncakeParserGT) | (1 << MooncakeParserLT) | (1 << MooncakeParserGTE) | (1 << MooncakeParserLTE))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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

func (s *ErrorTypeContext) FATAL() antlr.TerminalNode {
	return s.GetToken(MooncakeParserFATAL, 0)
}

func (s *ErrorTypeContext) SEVERE() antlr.TerminalNode {
	return s.GetToken(MooncakeParserSEVERE, 0)
}

func (s *ErrorTypeContext) WARNING() antlr.TerminalNode {
	return s.GetToken(MooncakeParserWARNING, 0)
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
	{
		p.SetState(85)
		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << MooncakeParserFATAL) | (1 << MooncakeParserSEVERE) | (1 << MooncakeParserWARNING))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


