package list

type Value interface{}

type Element struct {
	value Value
	prev  *Element
	next  *Element
}

func NewElement(v Value) *Element {
	e := &Element{value: v}
	return e
}

func (e *Element) GetValue() Value {
	return e.value
}

func (e *Element) GetNext() *Element {
	return e.next
}
