package lang

type Params struct {
	// TODO: encapsulate generic params interface with this struct
}

type Function interface {
	apply(params []interface{}) interface{}
}

type LengthFunction struct {
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

type DateTimeLongFunction struct {
	Function
}

func (f *DateTimeLongFunction) apply(params []interface{}) interface{} {
	panic("DateTimeLongFunction not implemented!")
}

type AfterCurrentTimeFunction struct {
	Function
}

func (f *AfterCurrentTimeFunction) apply(params []interface{}) interface{} {
	panic("AfterCurrentTimeFunction not implemented!")
}
