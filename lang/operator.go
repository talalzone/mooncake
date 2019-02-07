package lang

type Operation interface {
	eq(interface{}) bool
	ne(interface{}) bool
	gt(interface{}) bool
	lt(interface{}) bool
	gte(interface{}) bool
	lte(interface{}) bool
}

type Operator interface {
	Apply(interface{}, Literal) bool
}

type EqualOperator struct {
	Operator
}

type NotEqualOperator struct {
	Operator
}

type GreaterThanOperator struct {
	Operator
}

type LessThanOperator struct {
	Operator
}

type GreaterThanOrEqualOperator struct {
	Operator
}

type LessThanOrEqualOperator struct {
	Operator
}

func (o *EqualOperator) Apply(value interface{}, literal Literal) bool {
	return literal.eq(value)
}

func (o *NotEqualOperator) Apply(value interface{}, literal Literal) bool {
	return literal.ne(value)
}

func (o *GreaterThanOperator) Apply(value interface{}, literal Literal) bool {
	return literal.gt(value)
}

func (o *LessThanOperator) Apply(value interface{}, literal Literal) bool {
	return literal.lt(value)
}

func (o *GreaterThanOrEqualOperator) Apply(value interface{}, literal Literal) bool {
	return literal.gte(value)
}
func (o *LessThanOrEqualOperator) Apply(value interface{}, literal Literal) bool {
	return literal.lte(value)
}
