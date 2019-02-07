package lang

type Params struct {
	// TODO:
}

type Function interface {
	apply(params []interface{}) interface{}
}

type LengthFunction struct {
	Function
}

type DateTimeLongFunction struct {
	Function
}

type AfterCurrentTimeFunction struct {
	Function
}

func (f *LengthFunction) apply(params []interface{}) interface{} {
	switch t := params[0].(type) {
	case []interface{}:
		return len(t)
	default:
		panic("Unknown parameter. Required list.")
	}
}

func (f *DateTimeLongFunction) apply(params []interface{}) interface{} {
	panic("DateTimeLongFunction not implemented!")
}

func (f *AfterCurrentTimeFunction) apply(params []interface{}) interface{} {
	panic("AfterCurrentTimeFunction not implemented!")
}
