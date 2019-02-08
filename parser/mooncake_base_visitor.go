// Code generated from /Users/talal/go/src/mooncake/grammar/Mooncake.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Mooncake

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseMooncakeVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMooncakeVisitor) VisitMcrule(ctx *McruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitLinkedStmt(ctx *LinkedStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitSimpleStmt(ctx *SimpleStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitInlineStmt(ctx *InlineStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitExprStmt(ctx *ExprStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitErrorStmt(ctx *ErrorStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitOperator(ctx *OperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitErrorType(ctx *ErrorTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitFatalError(ctx *FatalErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitSevereError(ctx *SevereErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitWarningError(ctx *WarningErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitEqualOperator(ctx *EqualOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitNotEqualOperator(ctx *NotEqualOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitGreaterThanOperator(ctx *GreaterThanOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitLessThanOperator(ctx *LessThanOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitGreaterThanOrEqualOperator(ctx *GreaterThanOrEqualOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitLessThanOrEqualOperator(ctx *LessThanOrEqualOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitBoolLiteral(ctx *BoolLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitCtxLiteral(ctx *CtxLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitAttributeIdentifier(ctx *AttributeIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitDeclarationIdentifier(ctx *DeclarationIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitLengthFunction(ctx *LengthFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitDateTimeLongFunction(ctx *DateTimeLongFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitAfterCurrentTimeFunction(ctx *AfterCurrentTimeFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}
