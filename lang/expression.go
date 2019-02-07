package lang

type Expression struct {
	Identifier Identifier
	Operator   Operator
	Literal    Literal
}

func (e *Expression) Evaluate() bool {
	literal := e.Literal

	// CtxLiteral needs to be coerced to specific literal type
	ctxLiteral, isCtx := e.Literal.(*CtxLiteral)
	if isCtx {
		literal = ToLiteral(ctxLiteral.Val)
	}

	return e.Operator.Apply(e.Identifier.GetValue(), literal)
}
