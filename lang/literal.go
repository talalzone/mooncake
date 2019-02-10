package lang

import "reflect"

type Literal interface {
	Operation

	GetValue() interface{}
}

type NullLiteral struct {
	Literal
}

func (l *NullLiteral) GetValue() interface{} {
	return nil
}

func (l *NullLiteral) eq(value interface{}) bool {
	return value == nil
}

func (l *NullLiteral) ne(value interface{}) bool {
	return value != nil
}

func (l *NullLiteral) lt(value interface{}) bool {
	panic("Illegal operator usage!")
}

func (l *NullLiteral) gt(value interface{}) bool {
	panic("Illegal operator usage!")
}

func (l *NullLiteral) lte(value interface{}) bool {
	panic("Illegal operator usage!")
}

func (l *NullLiteral) gte(value interface{}) bool {
	panic("Illegal operator usage!")
}

type BoolLiteral struct {
	Literal

	Val bool
}

func (l *BoolLiteral) GetValue() interface{} {
	return l.Val
}

func (l *BoolLiteral) eq(value interface{}) bool {
	return value.(bool) == l.Val
}

func (l *BoolLiteral) ne(value interface{}) bool {
	return value.(bool) != l.Val
}

func (l *BoolLiteral) lt(value interface{}) bool {
	panic("Illegal operator usage!")
}

func (l *BoolLiteral) gt(value interface{}) bool {
	panic("Illegal operator usage!")
}

func (l *BoolLiteral) lte(value interface{}) bool {
	panic("Illegal operator usage!")
}

func (l *BoolLiteral) gte(value interface{}) bool {
	panic("Illegal operator usage!")
}

type StringLiteral struct {
	Literal

	Val string
}

func (l *StringLiteral) GetValue() interface{} {
	return l.Val
}

func (l *StringLiteral) eq(value interface{}) bool {
	return value == l.Val
}

func (l *StringLiteral) ne(value interface{}) bool {
	return value != l.Val
}

func (l *StringLiteral) lt(value interface{}) bool {
	panic("Illegal operator usage!")
}

func (l *StringLiteral) gt(value interface{}) bool {
	panic("Illegal operator usage!")
}

func (l *StringLiteral) lte(value interface{}) bool {
	panic("Illegal operator usage!")
}

func (l *StringLiteral) gte(value interface{}) bool {
	panic("Illegal operator usage!")
}

type IntLiteral struct {
	Literal

	Val int
}

func (l *IntLiteral) GetValue() interface{} {
	return l.Val
}

func (l *IntLiteral) eq(value interface{}) bool {
	return value == l.Val
}

func (l *IntLiteral) ne(value interface{}) bool {
	return value != l.Val
}

func (l *IntLiteral) gt(value interface{}) bool {
	return ToInt(value) > l.Val
}

func (l *IntLiteral) lt(value interface{}) bool {
	return ToInt(value) < l.Val
}

func (l *IntLiteral) gte(value interface{}) bool {
	return ToInt(value) >= l.Val
}

func (l *IntLiteral) lte(value interface{}) bool {
	return ToInt(value) <= l.Val
}

type FloatLiteral struct {
	Literal

	Val float64
}

func (l *FloatLiteral) GetValue() interface{} {
	return l.Val
}

func (l *FloatLiteral) eq(value interface{}) bool {
	return value == l.Val
}

func (l *FloatLiteral) ne(value interface{}) bool {
	return value != l.Val
}

func (l *FloatLiteral) gt(value interface{}) bool {
	return ToFloat(value) > l.Val
}

func (l *FloatLiteral) lt(value interface{}) bool {
	return ToFloat(value) < l.Val
}

func (l *FloatLiteral) gte(value interface{}) bool {
	return ToFloat(value) >= l.Val
}

func (l *FloatLiteral) lte(value interface{}) bool {
	return ToFloat(value) <= l.Val
}

type CtxLiteral struct {
	Literal

	Val reflect.Value
}

func (l *CtxLiteral) GetValue() interface{} {
	return l.Val
}

func (l *CtxLiteral) eq(value interface{}) bool {
	panic("Illegal operator usage!")
}

func (l *CtxLiteral) ne(value interface{}) bool {
	panic("Illegal operator usage!")
}

func (l *CtxLiteral) gt(value interface{}) bool {
	panic("Illegal operator usage!")
}

func (l *CtxLiteral) lt(value interface{}) bool {
	panic("Illegal operator usage!")
}

func (l *CtxLiteral) gte(value interface{}) bool {
	panic("Illegal operator usage!")
}

func (l *CtxLiteral) lte(value interface{}) bool {
	panic("Illegal operator usage!")
}
