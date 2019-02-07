// Code generated from /Users/talal/go/src/mooncake/grammar/Mooncake.g4 by ANTLR 4.7.2. DO NOT EDIT.

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

	// EnterExprStmt is called when entering the exprStmt production.
	EnterExprStmt(c *ExprStmtContext)

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

	// EnterFatalError is called when entering the fatalError production.
	EnterFatalError(c *FatalErrorContext)

	// EnterSevereError is called when entering the severeError production.
	EnterSevereError(c *SevereErrorContext)

	// EnterWarningError is called when entering the warningError production.
	EnterWarningError(c *WarningErrorContext)

	// EnterEqualOperator is called when entering the equalOperator production.
	EnterEqualOperator(c *EqualOperatorContext)

	// EnterNotEqualOperator is called when entering the notEqualOperator production.
	EnterNotEqualOperator(c *NotEqualOperatorContext)

	// EnterGreaterThanOperator is called when entering the greaterThanOperator production.
	EnterGreaterThanOperator(c *GreaterThanOperatorContext)

	// EnterLessThanOperator is called when entering the lessThanOperator production.
	EnterLessThanOperator(c *LessThanOperatorContext)

	// EnterGreaterThanOrEqualOperator is called when entering the greaterThanOrEqualOperator production.
	EnterGreaterThanOrEqualOperator(c *GreaterThanOrEqualOperatorContext)

	// EnterLessThanOrEqualOperator is called when entering the lessThanOrEqualOperator production.
	EnterLessThanOrEqualOperator(c *LessThanOrEqualOperatorContext)

	// EnterIntLiteral is called when entering the intLiteral production.
	EnterIntLiteral(c *IntLiteralContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterBoolLiteral is called when entering the boolLiteral production.
	EnterBoolLiteral(c *BoolLiteralContext)

	// EnterNullLiteral is called when entering the nullLiteral production.
	EnterNullLiteral(c *NullLiteralContext)

	// EnterCtxLiteral is called when entering the ctxLiteral production.
	EnterCtxLiteral(c *CtxLiteralContext)

	// EnterAttributeIdentifier is called when entering the attributeIdentifier production.
	EnterAttributeIdentifier(c *AttributeIdentifierContext)

	// EnterDeclarationIdentifier is called when entering the declarationIdentifier production.
	EnterDeclarationIdentifier(c *DeclarationIdentifierContext)

	// EnterLengthFunction is called when entering the lengthFunction production.
	EnterLengthFunction(c *LengthFunctionContext)

	// EnterDateTimeLongFunction is called when entering the dateTimeLongFunction production.
	EnterDateTimeLongFunction(c *DateTimeLongFunctionContext)

	// EnterAfterCurrentTimeFunction is called when entering the afterCurrentTimeFunction production.
	EnterAfterCurrentTimeFunction(c *AfterCurrentTimeFunctionContext)

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

	// ExitExprStmt is called when exiting the exprStmt production.
	ExitExprStmt(c *ExprStmtContext)

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

	// ExitFatalError is called when exiting the fatalError production.
	ExitFatalError(c *FatalErrorContext)

	// ExitSevereError is called when exiting the severeError production.
	ExitSevereError(c *SevereErrorContext)

	// ExitWarningError is called when exiting the warningError production.
	ExitWarningError(c *WarningErrorContext)

	// ExitEqualOperator is called when exiting the equalOperator production.
	ExitEqualOperator(c *EqualOperatorContext)

	// ExitNotEqualOperator is called when exiting the notEqualOperator production.
	ExitNotEqualOperator(c *NotEqualOperatorContext)

	// ExitGreaterThanOperator is called when exiting the greaterThanOperator production.
	ExitGreaterThanOperator(c *GreaterThanOperatorContext)

	// ExitLessThanOperator is called when exiting the lessThanOperator production.
	ExitLessThanOperator(c *LessThanOperatorContext)

	// ExitGreaterThanOrEqualOperator is called when exiting the greaterThanOrEqualOperator production.
	ExitGreaterThanOrEqualOperator(c *GreaterThanOrEqualOperatorContext)

	// ExitLessThanOrEqualOperator is called when exiting the lessThanOrEqualOperator production.
	ExitLessThanOrEqualOperator(c *LessThanOrEqualOperatorContext)

	// ExitIntLiteral is called when exiting the intLiteral production.
	ExitIntLiteral(c *IntLiteralContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitBoolLiteral is called when exiting the boolLiteral production.
	ExitBoolLiteral(c *BoolLiteralContext)

	// ExitNullLiteral is called when exiting the nullLiteral production.
	ExitNullLiteral(c *NullLiteralContext)

	// ExitCtxLiteral is called when exiting the ctxLiteral production.
	ExitCtxLiteral(c *CtxLiteralContext)

	// ExitAttributeIdentifier is called when exiting the attributeIdentifier production.
	ExitAttributeIdentifier(c *AttributeIdentifierContext)

	// ExitDeclarationIdentifier is called when exiting the declarationIdentifier production.
	ExitDeclarationIdentifier(c *DeclarationIdentifierContext)

	// ExitLengthFunction is called when exiting the lengthFunction production.
	ExitLengthFunction(c *LengthFunctionContext)

	// ExitDateTimeLongFunction is called when exiting the dateTimeLongFunction production.
	ExitDateTimeLongFunction(c *DateTimeLongFunctionContext)

	// ExitAfterCurrentTimeFunction is called when exiting the afterCurrentTimeFunction production.
	ExitAfterCurrentTimeFunction(c *AfterCurrentTimeFunctionContext)
}
