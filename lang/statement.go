package lang

type Statement interface {
	Evaluate() (bool, Error)
}

type InlineStatement struct {
	Statement

	Function       Function
	DeclIdentifier Identifier
}

type ExpressionStatement struct {
	Statement

	Expression
}

type ErrorStatement struct {
	Statement

	Code     string
	Info     string
	Severity int
}

type RuleStatement struct {
	Statement

	*SymbolTable

	InlineStatement
	ExpressionStatement
	ErrorStatement
}

func (rs *RuleStatement) Evaluate() (bool, ErrorStatement) {
	val := rs.Identifier.GetValue()

	if fn := rs.Function; fn != nil {
		params := []interface{}{val}
		val = rs.Function.apply(params)
		rs.Identifier.SetValue(val)
	}

	if decl := rs.DeclIdentifier; decl != nil {
		decl.SetValue(val)
		rs.SymbolTable.Put(decl.GetName(), val)
	}

	valid := rs.Expression.Evaluate()

	return valid, rs.ErrorStatement
}
