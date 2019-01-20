// Code generated from /Users/talal/Development/antlr-lang/mooncake/grammar/Mooncake.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Mooncake

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMooncakeListener is a complete listener for a parse tree produced by MooncakeParser.
type BaseMooncakeListener struct{}

var _ MooncakeListener = &BaseMooncakeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMooncakeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMooncakeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMooncakeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMooncakeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMcrule is called when production mcrule is entered.
func (s *BaseMooncakeListener) EnterMcrule(ctx *McruleContext) {}

// ExitMcrule is called when production mcrule is exited.
func (s *BaseMooncakeListener) ExitMcrule(ctx *McruleContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseMooncakeListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseMooncakeListener) ExitBlock(ctx *BlockContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseMooncakeListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseMooncakeListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseMooncakeListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseMooncakeListener) ExitStatement(ctx *StatementContext) {}

// EnterInlineStmt is called when production inlineStmt is entered.
func (s *BaseMooncakeListener) EnterInlineStmt(ctx *InlineStmtContext) {}

// ExitInlineStmt is called when production inlineStmt is exited.
func (s *BaseMooncakeListener) ExitInlineStmt(ctx *InlineStmtContext) {}

// EnterErrorStmt is called when production errorStmt is entered.
func (s *BaseMooncakeListener) EnterErrorStmt(ctx *ErrorStmtContext) {}

// ExitErrorStmt is called when production errorStmt is exited.
func (s *BaseMooncakeListener) ExitErrorStmt(ctx *ErrorStmtContext) {}

// EnterLinkedStmt is called when production linkedStmt is entered.
func (s *BaseMooncakeListener) EnterLinkedStmt(ctx *LinkedStmtContext) {}

// ExitLinkedStmt is called when production linkedStmt is exited.
func (s *BaseMooncakeListener) ExitLinkedStmt(ctx *LinkedStmtContext) {}

// EnterSimpleStmt is called when production simpleStmt is entered.
func (s *BaseMooncakeListener) EnterSimpleStmt(ctx *SimpleStmtContext) {}

// ExitSimpleStmt is called when production simpleStmt is exited.
func (s *BaseMooncakeListener) ExitSimpleStmt(ctx *SimpleStmtContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMooncakeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMooncakeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseMooncakeListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseMooncakeListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseMooncakeListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseMooncakeListener) ExitLiteral(ctx *LiteralContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseMooncakeListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseMooncakeListener) ExitFunction(ctx *FunctionContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseMooncakeListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseMooncakeListener) ExitOperator(ctx *OperatorContext) {}

// EnterErrorType is called when production errorType is entered.
func (s *BaseMooncakeListener) EnterErrorType(ctx *ErrorTypeContext) {}

// ExitErrorType is called when production errorType is exited.
func (s *BaseMooncakeListener) ExitErrorType(ctx *ErrorTypeContext) {}
