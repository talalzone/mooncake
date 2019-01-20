// Code generated from /Users/talal/Development/antlr-lang/mooncake/grammar/Mooncake.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Mooncake

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseMooncakeVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMooncakeVisitor) VisitMcrule(ctx *McruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitInlineStmt(ctx *InlineStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitErrorStmt(ctx *ErrorStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitLinkedStmt(ctx *LinkedStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitSimpleStmt(ctx *SimpleStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMooncakeVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
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
