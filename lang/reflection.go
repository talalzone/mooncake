package lang

import "reflect"

func GetValue(identifier string, ctx interface{}) reflect.Value {
	val := reflect.ValueOf(ctx).FieldByName(identifier)
	return val
}