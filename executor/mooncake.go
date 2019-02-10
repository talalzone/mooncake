package executor

import (
	"encoding/json"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/yalp/jsonpath"
	"io/ioutil"
	"log"
	"mooncake/lang"
	"mooncake/parser"
	"os"
)

type mooncakeVisitor struct {
	*parser.BaseMooncakeVisitor

	data        interface{}
	ctx         interface{}
	result      lang.ValidationResult
	symbolTable lang.SymbolTable
}

func (mv *mooncakeVisitor) VisitMcrule(ctx *parser.McruleContext) interface{} {
	log.Printf("VisitMcrule - %v", ctx.GetText())
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

func (mv *mooncakeVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	log.Printf("VisitBlock - %v", ctx.GetText())
	return mv.Visit(ctx)
}

func (mv *mooncakeVisitor) VisitLinkedStmt(ctx *parser.LinkedStmtContext) interface{} {
	log.Printf("VisitLinkedStmt - %v", ctx.GetText())
	valid := ctx.SimpleStmt().Accept(mv).(bool)

	// check linked statements
	if !valid && ctx.LinkedStmt() != nil {
		ctx.LinkedStmt().Accept(mv)
	}

	return valid
}

func (mv *mooncakeVisitor) VisitSimpleStmt(ctx *parser.SimpleStmtContext) interface{} {
	log.Printf("VisitSimpleStmt - %v", ctx.GetText())
	stmt := lang.RuleStatement{}
	stmt.SymbolTable = &mv.symbolTable
	stmt.ExpressionStatement = ctx.ExprStmt().Accept(mv).(lang.ExpressionStatement)
	stmt.ErrorStatement = ctx.ErrorStmt().Accept(mv).(lang.ErrorStatement)

	if inlineStmt := ctx.InlineStmt(); inlineStmt != nil {
		stmt.InlineStatement = inlineStmt.Accept(mv).(lang.InlineStatement)
	}

	valid, err := stmt.Evaluate()

	if valid {
		mv.setError(err)
	} else if ctx.Block() != nil {
		ctx.Block().Accept(mv)
	}

	return valid
}

func (mv *mooncakeVisitor) VisitInlineStmt(ctx *parser.InlineStmtContext) interface{} {
	log.Printf("VisitInlineStmt - %v", ctx.GetText())
	stmt := lang.InlineStatement{}

	if id := ctx.GetId(); id != nil {
		stmt.DeclIdentifier = &lang.DeclIdentifier{Name: id.GetText(), Val: nil}
	}
	if fn := ctx.GetFn(); fn != nil {
		stmt.Function = fn.Accept(mv).(lang.Function)
	}
	return stmt
}

func (mv *mooncakeVisitor) VisitExprStmt(ctx *parser.ExprStmtContext) interface{} {
	log.Printf("VisitExpression - %v", ctx.GetText())
	identifier := ctx.Identifier().Accept(mv).(lang.Identifier)
	operator := ctx.Operator().Accept(mv).(lang.Operator)
	literal := ctx.Literal().Accept(mv).(lang.Literal)

	expr := lang.Expression{Identifier: identifier, Operator: operator, Literal: literal}
	return lang.ExpressionStatement{Expression: expr}
}

func (mv *mooncakeVisitor) VisitErrorStmt(ctx *parser.ErrorStmtContext) interface{} {
	log.Printf("VisitErrorStmt - %v", ctx.GetText())
	code := lang.ToUnescapedStr(ctx.GetCode().GetText())
	info := lang.ToUnescapedStr(ctx.GetInfo().GetText())
	severity := ctx.ErrorType().Accept(mv).(int)

	return lang.ErrorStatement{Code: code, Info: info, Severity: severity}
}

func (mv *mooncakeVisitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	log.Printf("VisitIdentifier - %v", ctx.GetText())
	return mv.Visit(ctx)
}

func (mv *mooncakeVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	log.Printf("VisitLiteral - %v", ctx.GetText())
	return mv.Visit(ctx)
}

func (mv *mooncakeVisitor) VisitFunction(ctx *parser.FunctionContext) interface{} {
	log.Printf("VisitFunction - %v", ctx.GetText())
	return mv.Visit(ctx)
}

func (mv *mooncakeVisitor) VisitOperator(ctx *parser.OperatorContext) interface{} {
	log.Printf("VisitOperator - %v", ctx.GetText())
	return mv.Visit(ctx)
}

func (mv *mooncakeVisitor) VisitErrorType(ctx *parser.ErrorTypeContext) interface{} {
	log.Printf("VisitErrorType - %v", ctx.GetText())
	return mv.Visit(ctx)
}

func (mv *mooncakeVisitor) VisitFatalError(ctx *parser.FatalErrorContext) interface{} {
	log.Printf("VisitFatalError - %v", ctx.GetText())
	return ctx.GetStart().GetTokenType()
}

func (mv *mooncakeVisitor) VisitSevereError(ctx *parser.SevereErrorContext) interface{} {
	log.Printf("VisitSevereError - %v", ctx.GetText())
	return ctx.GetStart().GetTokenType()
}

func (mv *mooncakeVisitor) VisitWarningError(ctx *parser.WarningErrorContext) interface{} {
	log.Printf("VisitWarningError - %v", ctx.GetText())
	return ctx.GetStart().GetTokenType()
}

func (mv *mooncakeVisitor) VisitEqualOperator(ctx *parser.EqualOperatorContext) interface{} {
	log.Printf("VisitEqualOperator - %v", ctx.GetText())
	return &lang.EqualOperator{}
}

func (mv *mooncakeVisitor) VisitNotEqualOperator(ctx *parser.NotEqualOperatorContext) interface{} {
	log.Printf("VisitNotEqualOperator - %v", ctx.GetText())
	return &lang.NotEqualOperator{}
}

func (mv *mooncakeVisitor) VisitGreaterThanOperator(ctx *parser.GreaterThanOperatorContext) interface{} {
	log.Printf("VisitGreaterThanOperator - %v", ctx.GetText())
	return &lang.GreaterThanOperator{}
}

func (mv *mooncakeVisitor) VisitLessThanOperator(ctx *parser.LessThanOperatorContext) interface{} {
	log.Printf("VisitLessThanOperator - %v", ctx.GetText())
	return &lang.LessThanOperator{}
}

func (mv *mooncakeVisitor) VisitGreaterThanOrEqualOperator(ctx *parser.GreaterThanOrEqualOperatorContext) interface{} {
	log.Printf("VisitGreaterThanOrEqualOperator - %v", ctx.GetText())
	return &lang.GreaterThanOrEqualOperator{}
}

func (mv *mooncakeVisitor) VisitLessThanOrEqualOperator(ctx *parser.LessThanOrEqualOperatorContext) interface{} {
	log.Printf("VisitLessThanOrEqualOperator - %v", ctx.GetText())
	return &lang.LessThanOrEqualOperator{}
}

func (mv *mooncakeVisitor) VisitIntLiteral(ctx *parser.IntLiteralContext) interface{} {
	log.Printf("VisitIntLiteral - %v", ctx.GetText())
	val := lang.StrToInt(ctx.GetText())
	return &lang.IntLiteral{Val: val}
}

func (mv *mooncakeVisitor) VisitFloatLiteral(ctx *parser.FloatLiteralContext) interface{} {
	log.Printf("VisitFloatLiteral - %v", ctx.GetText())
	val := lang.StrToFloat(ctx.GetText())
	return &lang.FloatLiteral{Val: val}
}

func (mv *mooncakeVisitor) VisitBoolLiteral(ctx *parser.BoolLiteralContext) interface{} {
	log.Printf("VisitBoolLiteral - %v", ctx.GetText())
	val := lang.StrToBool(ctx.GetText())
	return &lang.BoolLiteral{Val: val}
}

func (mv *mooncakeVisitor) VisitNullLiteral(ctx *parser.NullLiteralContext) interface{} {
	log.Printf("VisitNullLiteral - %v", ctx.GetText())
	return &lang.NullLiteral{}
}

func (mv *mooncakeVisitor) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	log.Printf("VisitStringLiteral - %v", ctx.GetText())
	val := lang.ToUnescapedStr(ctx.GetText())
	return &lang.StringLiteral{Val: val}
}

func (mv *mooncakeVisitor) VisitCtxLiteral(ctx *parser.CtxLiteralContext) interface{} {
	log.Printf("VisitCtxLiteral - %v", ctx.GetText())
	id := lang.StrToCtx(ctx.GetText())
	val := lang.GetValue(id, mv.ctx)
	return &lang.CtxLiteral{Val: val}
}

func (mv *mooncakeVisitor) VisitAttributeIdentifier(ctx *parser.AttributeIdentifierContext) interface{} {
	log.Printf("VisitAttributeIdentifier - %v", ctx.GetText())
	val, _ := jsonpath.Read(mv.data, "$."+ctx.GetText())
	return &lang.AttrIdentifier{Name: ctx.GetText(), Val: val}
}

func (mv *mooncakeVisitor) VisitDeclarationIdentifier(ctx *parser.DeclarationIdentifierContext) interface{} {
	log.Printf("VisitDeclarationIdentifier - %v", ctx.GetText())
	val := mv.symbolTable.Get(ctx.GetText())
	return &lang.DeclIdentifier{Name: ctx.GetText(), Val: val}
}

func (v *mooncakeVisitor) VisitLengthFunction(ctx *parser.LengthFunctionContext) interface{} {
	return &lang.LengthFunction{}
}

func (v *mooncakeVisitor) VisitDateTimeLongFunction(ctx *parser.DateTimeLongFunctionContext) interface{} {
	return &lang.DateTimeLongFunction{}
}

func (v *mooncakeVisitor) VisitAfterCurrentTimeFunction(ctx *parser.AfterCurrentTimeFunctionContext) interface{} {
	return &lang.AfterCurrentTimeFunction{}
}

func (mv *mooncakeVisitor) Visit(tree antlr.ParseTree) interface{} {
	var value interface{}

	for _, child := range tree.GetChildren() {
		log.Printf("Visiting child - %v", child)

		switch child := child.(type) {
		case antlr.TerminalNode:
			mv.VisitTerminal(child)
		case antlr.ErrorNode:
			mv.VisitErrorNode(child)
		case antlr.RuleNode:
			value = child.Accept(mv)
		default:
			// can this happen??
		}
	}
	return value // not aggregating results
}

func (v *mooncakeVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return node.GetSymbol()
}

func (v *mooncakeVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return node.GetSymbol()
}

func (v *mooncakeVisitor) setError(err lang.ErrorStatement) {
	e := lang.Error{Code: err.Code, Info: err.Info}

	switch err.Severity {
	case parser.MooncakeParserWARNING:
		v.result.Warning = append(v.result.Warning, e)
	case parser.MooncakeParserSEVERE:
		v.result.Severe = append(v.result.Severe, e)
	case parser.MooncakeParserFATAL:
		v.result.Fatal = append(v.result.Fatal, e)
	}
}

func Execute(rules string, filePath string, ctx interface{}) lang.ValidationResult {
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
		result:      lang.NewValidationResult(),
		symbolTable: lang.NewSymbolTable(),
	}

	antlr.ParseTreeVisitor.Visit(mooncakeVisitor, p.Mcrule())

	return mooncakeVisitor.result
}
