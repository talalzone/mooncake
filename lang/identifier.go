package lang

type Identifier interface {
	GetName() string
	GetValue() interface{}
	SetValue(interface{})
}

type AttrIdentifier struct {
	Identifier

	Name string
	Val  interface{}
}

func (i *AttrIdentifier) GetName() string {
	return i.Name
}

func (i *AttrIdentifier) GetValue() interface{} {
	return i.Val
}

func (i *AttrIdentifier) SetValue(val interface{}) {
	i.Val = val
}

type DeclIdentifier struct {
	Identifier

	Name string
	Val  interface{}
}

func (i *DeclIdentifier) GetName() string {
	return i.Name
}

func (i *DeclIdentifier) GetValue() interface{} {
	return i.Val
}

func (i *DeclIdentifier) SetValue(val interface{}) {
	i.Val = val
}
