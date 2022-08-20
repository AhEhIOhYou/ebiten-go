package objectpool

import "github.com/AhEhIOhYou/project2/prj2/internal/list"

type Iterator struct {
	current *list.Element
}

func (ite *Iterator) HasNext() bool {
	return ite.current != nil
}

func (ite *Iterator) Next() *Object {
	e := ite.current
	ite.current = e.GetNext()
	return (*Object)(e.GetValue())
}
