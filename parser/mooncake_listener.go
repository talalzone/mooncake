// Code generated from /Users/talal/Development/antlr-lang/mooncake/grammar/Mooncake.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Mooncake

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MooncakeListener is a complete listener for a parse tree produced by MooncakeParser.
type MooncakeListener interface {
	antlr.ParseTreeListener

	// EnterMcrule is called when entering the mcrule production.
	EnterMcrule(c *McruleContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterInlineStmt is called when entering the inlineStmt production.
	EnterInlineStmt(c *InlineStmtContext)

	// EnterErrorStmt is called when entering the errorStmt production.
	EnterErrorStmt(c *ErrorStmtContext)

	// EnterLinkedStmt is called when entering the linkedStmt production.
	EnterLinkedStmt(c *LinkedStmtContext)

	// EnterSimpleStmt is called when entering the simpleStmt production.
	EnterSimpleStmt(c *SimpleStmtContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterErrorType is called when entering the errorType production.
	EnterErrorType(c *ErrorTypeContext)

	// ExitMcrule is called when exiting the mcrule production.
	ExitMcrule(c *McruleContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitInlineStmt is called when exiting the inlineStmt production.
	ExitInlineStmt(c *InlineStmtContext)

	// ExitErrorStmt is called when exiting the errorStmt production.
	ExitErrorStmt(c *ErrorStmtContext)

	// ExitLinkedStmt is called when exiting the linkedStmt production.
	ExitLinkedStmt(c *LinkedStmtContext)

	// ExitSimpleStmt is called when exiting the simpleStmt production.
	ExitSimpleStmt(c *SimpleStmtContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitErrorType is called when exiting the errorType production.
	ExitErrorType(c *ErrorTypeContext)
}
