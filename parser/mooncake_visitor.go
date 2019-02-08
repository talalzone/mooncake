// Code generated from /Users/talal/go/src/mooncake/grammar/Mooncake.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Mooncake

import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by MooncakeParser.
type MooncakeVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MooncakeParser#mcrule.
	VisitMcrule(ctx *McruleContext) interface{}

	// Visit a parse tree produced by MooncakeParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by MooncakeParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by MooncakeParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by MooncakeParser#linkedStmt.
	VisitLinkedStmt(ctx *LinkedStmtContext) interface{}

	// Visit a parse tree produced by MooncakeParser#simpleStmt.
	VisitSimpleStmt(ctx *SimpleStmtContext) interface{}

	// Visit a parse tree produced by MooncakeParser#inlineStmt.
	VisitInlineStmt(ctx *InlineStmtContext) interface{}

	// Visit a parse tree produced by MooncakeParser#exprStmt.
	VisitExprStmt(ctx *ExprStmtContext) interface{}

	// Visit a parse tree produced by MooncakeParser#errorStmt.
	VisitErrorStmt(ctx *ErrorStmtContext) interface{}

	// Visit a parse tree produced by MooncakeParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by MooncakeParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by MooncakeParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by MooncakeParser#operator.
	VisitOperator(ctx *OperatorContext) interface{}

	// Visit a parse tree produced by MooncakeParser#errorType.
	VisitErrorType(ctx *ErrorTypeContext) interface{}

	// Visit a parse tree produced by MooncakeParser#fatalError.
	VisitFatalError(ctx *FatalErrorContext) interface{}

	// Visit a parse tree produced by MooncakeParser#severeError.
	VisitSevereError(ctx *SevereErrorContext) interface{}

	// Visit a parse tree produced by MooncakeParser#warningError.
	VisitWarningError(ctx *WarningErrorContext) interface{}

	// Visit a parse tree produced by MooncakeParser#equalOperator.
	VisitEqualOperator(ctx *EqualOperatorContext) interface{}

	// Visit a parse tree produced by MooncakeParser#notEqualOperator.
	VisitNotEqualOperator(ctx *NotEqualOperatorContext) interface{}

	// Visit a parse tree produced by MooncakeParser#greaterThanOperator.
	VisitGreaterThanOperator(ctx *GreaterThanOperatorContext) interface{}

	// Visit a parse tree produced by MooncakeParser#lessThanOperator.
	VisitLessThanOperator(ctx *LessThanOperatorContext) interface{}

	// Visit a parse tree produced by MooncakeParser#greaterThanOrEqualOperator.
	VisitGreaterThanOrEqualOperator(ctx *GreaterThanOrEqualOperatorContext) interface{}

	// Visit a parse tree produced by MooncakeParser#lessThanOrEqualOperator.
	VisitLessThanOrEqualOperator(ctx *LessThanOrEqualOperatorContext) interface{}

	// Visit a parse tree produced by MooncakeParser#intLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}

	// Visit a parse tree produced by MooncakeParser#floatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by MooncakeParser#boolLiteral.
	VisitBoolLiteral(ctx *BoolLiteralContext) interface{}

	// Visit a parse tree produced by MooncakeParser#nullLiteral.
	VisitNullLiteral(ctx *NullLiteralContext) interface{}

	// Visit a parse tree produced by MooncakeParser#ctxLiteral.
	VisitCtxLiteral(ctx *CtxLiteralContext) interface{}

	// Visit a parse tree produced by MooncakeParser#attributeIdentifier.
	VisitAttributeIdentifier(ctx *AttributeIdentifierContext) interface{}

	// Visit a parse tree produced by MooncakeParser#declarationIdentifier.
	VisitDeclarationIdentifier(ctx *DeclarationIdentifierContext) interface{}

	// Visit a parse tree produced by MooncakeParser#lengthFunction.
	VisitLengthFunction(ctx *LengthFunctionContext) interface{}

	// Visit a parse tree produced by MooncakeParser#dateTimeLongFunction.
	VisitDateTimeLongFunction(ctx *DateTimeLongFunctionContext) interface{}

	// Visit a parse tree produced by MooncakeParser#afterCurrentTimeFunction.
	VisitAfterCurrentTimeFunction(ctx *AfterCurrentTimeFunctionContext) interface{}

}