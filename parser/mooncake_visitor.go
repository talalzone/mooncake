// Code generated from /Users/talal/Development/antlr-lang/mooncake/grammar/Mooncake.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Mooncake

import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by MooncakeParser.
type MooncakeVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MooncakeParser#mcrule.
	VisitMcrule(ctx *McruleContext) interface{}

	// Visit a parse tree produced by MooncakeParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by MooncakeParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by MooncakeParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by MooncakeParser#inlineStmt.
	VisitInlineStmt(ctx *InlineStmtContext) interface{}

	// Visit a parse tree produced by MooncakeParser#errorStmt.
	VisitErrorStmt(ctx *ErrorStmtContext) interface{}

	// Visit a parse tree produced by MooncakeParser#linkedStmt.
	VisitLinkedStmt(ctx *LinkedStmtContext) interface{}

	// Visit a parse tree produced by MooncakeParser#simpleStmt.
	VisitSimpleStmt(ctx *SimpleStmtContext) interface{}

	// Visit a parse tree produced by MooncakeParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

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

}