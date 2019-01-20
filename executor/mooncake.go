package executor

import (
	"../domain"
	"../parser"
	"encoding/json"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/yalp/jsonpath"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type inline struct {
	id string
	fn antlr.Token
}

type expr struct {
	identifier interface{}
	literal    interface{}
	operator   antlr.Token
}

type err struct {
	humanCode   string
	numericCode int
	errorType   antlr.Token
}

type rule struct {
	inline
	expr
	err
}

type mooncakeVisitor struct {
	*parser.BaseMooncakeVisitor

	data        interface{}
	ctx         interface{}
	result      domain.ValidationResult
	symbolTable map[string]interface{}
}

func (mv *mooncakeVisitor) VisitErule(ctx *parser.McruleContext) interface{} {
	log.Printf("VisitErule - %v", ctx.GetText())
	return mv.Visit(ctx)
}

func (mv *mooncakeVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	log.Printf("VisitBlock - %v", ctx.GetText())
	return mv.Visit(ctx)
}

func (mv *mooncakeVisitor) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	log.Printf("VisitStatementList - %v", ctx.GetText())
	return mv.Visit(ctx)
}

func (mv *mooncakeVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	log.Printf("VisitStatement - %v", ctx.GetText())
	return mv.Visit(ctx)
}

func (mv *mooncakeVisitor) VisitInlineStmt(ctx *parser.InlineStmtContext) interface{} {
	log.Printf("VisitInlineStmt - %v", ctx.GetText())
	inline := inline{}

	if id := ctx.GetId(); id != nil {
		inline.id = id.GetText()
	}

	if fn := ctx.GetFn(); fn != nil {
		inline.fn = fn.Accept(mv).(antlr.Token)
	}

	return inline
}

func (mv *mooncakeVisitor) VisitErrorStmt(ctx *parser.ErrorStmtContext) interface{} {
	log.Printf("VisitErrorStmt - %v", ctx.GetText())
	humanCode := ctx.GetHumanCode().GetText()
	errCode, _ := strconv.Atoi(ctx.GetErrCode().GetText())
	errType := ctx.ErrorType().Accept(mv).(antlr.Token)

	err := err{
		humanCode,
		errCode,
		errType,
	}
	return err
}

func (mv *mooncakeVisitor) VisitLinkedStmt(ctx *parser.LinkedStmtContext) interface{} {
	log.Printf("VisitLinkedStmt - %v", ctx.GetText())
	isValid := ctx.SimpleStmt().Accept(mv).(bool)

	// check linked statements
	if isValid {
		mv.Visit(ctx.LinkedStmt())
	}

	return isValid
}

func (mv *mooncakeVisitor) VisitSimpleStmt(ctx *parser.SimpleStmtContext) interface{} {
	log.Printf("VisitSimpleStmt - %v", ctx.GetText())
	rule := rule{}

	rule.expr = ctx.Expression().Accept(mv).(expr)
	rule.err = ctx.ErrorStmt().Accept(mv).(err)

	identifier := rule.expr.identifier

	if inStmt := ctx.InlineStmt(); inStmt != nil {
		rule.inline = inStmt.Accept(mv).(inline)
		log.Printf("Evaluating InlineStmt - %v", rule.inline)

		if fn := rule.inline.fn; fn != nil {
			switch fn.GetTokenType() {
			case parser.MooncakeParserLEN_FUNC:
				identifier = len(identifier.([]interface{}))
			case parser.MooncakeParserDATETIME_LONG:
				// to datetime long
			case parser.MooncakeParserFLOAT_FUNC:
				// to float
			}
		}

		if id := rule.inline.id; id != "" {
			mv.symbolTable[id] = identifier
		}
	}

	log.Printf("Evaluating Expr - {%v, %v, %v} Err - %v:%v",
		rule.expr.identifier, rule.expr.operator.GetText(),
		rule.expr.literal, rule.err.humanCode,
		rule.err.numericCode)

	isValid := true
	switch rule.expr.operator.GetTokenType() {
	case parser.MooncakeParserEQ:
		if identifier == rule.expr.literal {
			isValid = false
		}
	case parser.MooncakeParserNE:
		if rule.expr.identifier != rule.expr.literal {
			isValid = false
		}
	}

	log.Printf("Evaluated Expr - %v", mv.result)

	if !isValid {
		mv.setError(rule.err)
	} else if reflect.ValueOf(ctx.Block()).IsValid() {
		ctx.Block().Accept(mv)
	}

	return isValid
}

func (mv *mooncakeVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	log.Printf("VisitExpression - %v", ctx.GetText())
	expr := expr{
		ctx.Identifier().Accept(mv),
		ctx.Literal().Accept(mv),
		ctx.Operator().Accept(mv).(antlr.Token),
	}
	return expr
}

func (mv *mooncakeVisitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	log.Printf("VisitIdentifier - %v", ctx.GetText())
	var value interface{}

	switch ctx.GetStart().GetTokenType() {
	case parser.MooncakeParserIDENTIFIER:
		value, _ = jsonpath.Read(mv.data, "$."+ctx.GetText())
		log.Printf("Evaluated IDENTIFIER - %v  %v", ctx.GetText(), value)
	case parser.MooncakeParserINLINE_ID:
		value = mv.symbolTable[ctx.GetText()]
		log.Printf("Evaluated INLINE_ID - %v  %v", ctx.GetText(), value)
	}

	return value
}

func (mv *mooncakeVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	log.Printf("VisitLiteral - %v", ctx.GetText())
	var value interface{}

	switch ctx.GetStart().GetTokenType() {
	case parser.MooncakeParserCTX_ID:
		r := strings.NewReplacer("${", "", "}", "")
		id := r.Replace(ctx.GetText())
		value = reflect.ValueOf(mv.ctx).FieldByName(id)
		log.Printf("Evaluated CTX_ID - %v  %v", ctx.GetText(), value)
	case parser.MooncakeParserINT:
		value, _ = strconv.Atoi(ctx.GetText())
		log.Printf("Evaluated INT - %v  %v", ctx.GetText(), value)
	case parser.MooncakeParserBOOL:
		value, _ = strconv.ParseBool(ctx.GetText())
		log.Printf("Evaluated BOOL - %v  %v", ctx.GetText(), value)
	case parser.MooncakeParserNULL:
		value = nil
		log.Printf("Evaluated NULL - %v  %v", ctx.GetText(), value)
	}

	return value
}

func (mv *mooncakeVisitor) VisitFunction(ctx *parser.FunctionContext) interface{} {
	log.Printf("VisitFunction - %v", ctx.GetText())
	return ctx.GetStart()
}

func (mv *mooncakeVisitor) VisitOperator(ctx *parser.OperatorContext) interface{} {
	log.Printf("VisitOperator - %v", ctx.GetText())
	return ctx.GetStart()
}

func (mv *mooncakeVisitor) VisitErrorType(ctx *parser.ErrorTypeContext) interface{} {
	log.Printf("VisitErrorType - %v", ctx.GetText())
	return ctx.GetStart()
}

func (mv *mooncakeVisitor) Visit(tree antlr.ParseTree) interface{} {
	for _, child := range tree.GetChildren() {
		log.Printf("Visiting child - %v", child)

		switch child := child.(type) {
		case antlr.TerminalNode:
			mv.VisitTerminal(child)
		case antlr.ErrorNode:
			mv.VisitErrorNode(child)
		case antlr.RuleNode:
			child.Accept(mv)
		default:
			// can this happen??
		}
	}

	return nil // not aggregating results for now
}

func (v *mooncakeVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return node.GetSymbol()
}

func (v *mooncakeVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return node.GetSymbol()
}

func (v *mooncakeVisitor) setError(err err) {
	switch err.errorType.GetTokenType() {
	case parser.MooncakeParserWARNING:
		v.result.WarningCodes = append(v.result.WarningCodes, domain.Error{err.numericCode, err.humanCode})
	case parser.MooncakeParserSEVERE:
		v.result.SevereCodes = append(v.result.SevereCodes, domain.Error{err.numericCode, err.humanCode})
	case parser.MooncakeParserFATAL:
		v.result.FatalCodes = append(v.result.FatalCodes, domain.Error{err.numericCode, err.humanCode})
	}
}

func Execute(rules string, filePath string, ctx interface{}) domain.ValidationResult {
	is := antlr.NewInputStream(rules)

	lexer := parser.NewMooncakeLexer(is)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewMooncakeParser(stream)

	// Init Visitor
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(file)

	var fileData interface{}
	if err := json.Unmarshal(byteValue, &fileData); err != nil {
		panic(err)
	}

	mooncakeVisitor := &mooncakeVisitor{
		data:        fileData,
		ctx:         ctx,
		symbolTable: map[string]interface{}{},
	}

	antlr.ParseTreeVisitor.Visit(mooncakeVisitor, p.Mcrule())

	return mooncakeVisitor.result
}
