package lang

import (
	"reflect"
	"strconv"
	"strings"
)

var replacer = strings.NewReplacer("${", "", "}", "")

func ToInt(val interface{}) int {
	f, ok := val.(float64)
	if ok {
		return int(f)
	}

	i, ok := val.(int)
	if ok {
		return i
	}

	panic("Can not convert value to Int.")
}

func ToFloat(val interface{}) float64 {
	i, ok := val.(int)
	if ok {
		return float64(i)
	}

	f, ok := val.(float64)
	if ok {
		return f
	}

	panic("Can not convert value to Float.")
}

func ToLiteral(val reflect.Value) Literal {
	if !val.IsValid() {
		return &NullLiteral{}
	}

	switch val.Kind() {
	case reflect.Int:
		return &IntLiteral{Val: int(val.Int())}
	case reflect.Float64, reflect.Float32:
		return &FloatLiteral{Val: val.Float()}
	case reflect.String:
		return &StringLiteral{Val: val.String()}
	case reflect.Bool:
		return &BoolLiteral{Val: val.Bool()}
	default:
		panic("Unable to coerce type!")
	}
}

func ToUnescapedStr(val string) string {
	unescapedStr := val[1 : len(val)-1] // for now removing superfluous quotes
	return unescapedStr
}

func StrToInt(val string) int {
	intVal, _ := strconv.Atoi(val)
	return intVal
}

func StrToFloat(val string) float64 {
	floatVal, _ := strconv.ParseFloat(val, 64)
	return floatVal
}

func StrToBool(val string) bool {
	boolVal, _ := strconv.ParseBool(val)
	return boolVal
}

func StrToCtx(val string) string {
	id := replacer.Replace(val)
	return id
}
