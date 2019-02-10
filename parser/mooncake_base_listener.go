// Code generated from /Users/talal/go/src/mooncake/grammar/Mooncake.g4 by ANTLR 4.7.2. DO NOT EDIT.

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

// EnterStatementList is called when production statementList is entered.
func (s *BaseMooncakeListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseMooncakeListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseMooncakeListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseMooncakeListener) ExitStatement(ctx *StatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseMooncakeListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseMooncakeListener) ExitBlock(ctx *BlockContext) {}

// EnterLinkedStmt is called when production linkedStmt is entered.
func (s *BaseMooncakeListener) EnterLinkedStmt(ctx *LinkedStmtContext) {}

// ExitLinkedStmt is called when production linkedStmt is exited.
func (s *BaseMooncakeListener) ExitLinkedStmt(ctx *LinkedStmtContext) {}

// EnterSimpleStmt is called when production simpleStmt is entered.
func (s *BaseMooncakeListener) EnterSimpleStmt(ctx *SimpleStmtContext) {}

// ExitSimpleStmt is called when production simpleStmt is exited.
func (s *BaseMooncakeListener) ExitSimpleStmt(ctx *SimpleStmtContext) {}

// EnterInlineStmt is called when production inlineStmt is entered.
func (s *BaseMooncakeListener) EnterInlineStmt(ctx *InlineStmtContext) {}

// ExitInlineStmt is called when production inlineStmt is exited.
func (s *BaseMooncakeListener) ExitInlineStmt(ctx *InlineStmtContext) {}

// EnterErrorStmt is called when production errorStmt is entered.
func (s *BaseMooncakeListener) EnterErrorStmt(ctx *ErrorStmtContext) {}

// ExitErrorStmt is called when production errorStmt is exited.
func (s *BaseMooncakeListener) ExitErrorStmt(ctx *ErrorStmtContext) {}

// EnterExprStmt is called when production exprStmt is entered.
func (s *BaseMooncakeListener) EnterExprStmt(ctx *ExprStmtContext) {}

// ExitExprStmt is called when production exprStmt is exited.
func (s *BaseMooncakeListener) ExitExprStmt(ctx *ExprStmtContext) {}

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

// EnterFatalError is called when production fatalError is entered.
func (s *BaseMooncakeListener) EnterFatalError(ctx *FatalErrorContext) {}

// ExitFatalError is called when production fatalError is exited.
func (s *BaseMooncakeListener) ExitFatalError(ctx *FatalErrorContext) {}

// EnterSevereError is called when production severeError is entered.
func (s *BaseMooncakeListener) EnterSevereError(ctx *SevereErrorContext) {}

// ExitSevereError is called when production severeError is exited.
func (s *BaseMooncakeListener) ExitSevereError(ctx *SevereErrorContext) {}

// EnterWarningError is called when production warningError is entered.
func (s *BaseMooncakeListener) EnterWarningError(ctx *WarningErrorContext) {}

// ExitWarningError is called when production warningError is exited.
func (s *BaseMooncakeListener) ExitWarningError(ctx *WarningErrorContext) {}

// EnterEqualOperator is called when production equalOperator is entered.
func (s *BaseMooncakeListener) EnterEqualOperator(ctx *EqualOperatorContext) {}

// ExitEqualOperator is called when production equalOperator is exited.
func (s *BaseMooncakeListener) ExitEqualOperator(ctx *EqualOperatorContext) {}

// EnterNotEqualOperator is called when production notEqualOperator is entered.
func (s *BaseMooncakeListener) EnterNotEqualOperator(ctx *NotEqualOperatorContext) {}

// ExitNotEqualOperator is called when production notEqualOperator is exited.
func (s *BaseMooncakeListener) ExitNotEqualOperator(ctx *NotEqualOperatorContext) {}

// EnterGreaterThanOperator is called when production greaterThanOperator is entered.
func (s *BaseMooncakeListener) EnterGreaterThanOperator(ctx *GreaterThanOperatorContext) {}

// ExitGreaterThanOperator is called when production greaterThanOperator is exited.
func (s *BaseMooncakeListener) ExitGreaterThanOperator(ctx *GreaterThanOperatorContext) {}

// EnterLessThanOperator is called when production lessThanOperator is entered.
func (s *BaseMooncakeListener) EnterLessThanOperator(ctx *LessThanOperatorContext) {}

// ExitLessThanOperator is called when production lessThanOperator is exited.
func (s *BaseMooncakeListener) ExitLessThanOperator(ctx *LessThanOperatorContext) {}

// EnterGreaterThanOrEqualOperator is called when production greaterThanOrEqualOperator is entered.
func (s *BaseMooncakeListener) EnterGreaterThanOrEqualOperator(ctx *GreaterThanOrEqualOperatorContext) {}

// ExitGreaterThanOrEqualOperator is called when production greaterThanOrEqualOperator is exited.
func (s *BaseMooncakeListener) ExitGreaterThanOrEqualOperator(ctx *GreaterThanOrEqualOperatorContext) {}

// EnterLessThanOrEqualOperator is called when production lessThanOrEqualOperator is entered.
func (s *BaseMooncakeListener) EnterLessThanOrEqualOperator(ctx *LessThanOrEqualOperatorContext) {}

// ExitLessThanOrEqualOperator is called when production lessThanOrEqualOperator is exited.
func (s *BaseMooncakeListener) ExitLessThanOrEqualOperator(ctx *LessThanOrEqualOperatorContext) {}

// EnterIntLiteral is called when production intLiteral is entered.
func (s *BaseMooncakeListener) EnterIntLiteral(ctx *IntLiteralContext) {}

// ExitIntLiteral is called when production intLiteral is exited.
func (s *BaseMooncakeListener) ExitIntLiteral(ctx *IntLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseMooncakeListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseMooncakeListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterBoolLiteral is called when production boolLiteral is entered.
func (s *BaseMooncakeListener) EnterBoolLiteral(ctx *BoolLiteralContext) {}

// ExitBoolLiteral is called when production boolLiteral is exited.
func (s *BaseMooncakeListener) ExitBoolLiteral(ctx *BoolLiteralContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseMooncakeListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseMooncakeListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseMooncakeListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseMooncakeListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterCtxLiteral is called when production ctxLiteral is entered.
func (s *BaseMooncakeListener) EnterCtxLiteral(ctx *CtxLiteralContext) {}

// ExitCtxLiteral is called when production ctxLiteral is exited.
func (s *BaseMooncakeListener) ExitCtxLiteral(ctx *CtxLiteralContext) {}

// EnterAttributeIdentifier is called when production attributeIdentifier is entered.
func (s *BaseMooncakeListener) EnterAttributeIdentifier(ctx *AttributeIdentifierContext) {}

// ExitAttributeIdentifier is called when production attributeIdentifier is exited.
func (s *BaseMooncakeListener) ExitAttributeIdentifier(ctx *AttributeIdentifierContext) {}

// EnterDeclarationIdentifier is called when production declarationIdentifier is entered.
func (s *BaseMooncakeListener) EnterDeclarationIdentifier(ctx *DeclarationIdentifierContext) {}

// ExitDeclarationIdentifier is called when production declarationIdentifier is exited.
func (s *BaseMooncakeListener) ExitDeclarationIdentifier(ctx *DeclarationIdentifierContext) {}

// EnterLengthFunction is called when production lengthFunction is entered.
func (s *BaseMooncakeListener) EnterLengthFunction(ctx *LengthFunctionContext) {}

// ExitLengthFunction is called when production lengthFunction is exited.
func (s *BaseMooncakeListener) ExitLengthFunction(ctx *LengthFunctionContext) {}

// EnterDateTimeLongFunction is called when production dateTimeLongFunction is entered.
func (s *BaseMooncakeListener) EnterDateTimeLongFunction(ctx *DateTimeLongFunctionContext) {}

// ExitDateTimeLongFunction is called when production dateTimeLongFunction is exited.
func (s *BaseMooncakeListener) ExitDateTimeLongFunction(ctx *DateTimeLongFunctionContext) {}

// EnterAfterCurrentTimeFunction is called when production afterCurrentTimeFunction is entered.
func (s *BaseMooncakeListener) EnterAfterCurrentTimeFunction(ctx *AfterCurrentTimeFunctionContext) {}

// ExitAfterCurrentTimeFunction is called when production afterCurrentTimeFunction is exited.
func (s *BaseMooncakeListener) ExitAfterCurrentTimeFunction(ctx *AfterCurrentTimeFunctionContext) {}
