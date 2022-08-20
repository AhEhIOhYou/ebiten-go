package objectpool

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/list"
	"unsafe"
)

type Object struct {
	data     unsafe.Pointer
	isActive bool
	elem     *list.Element
}

func (o *Object) GetData() unsafe.Pointer {
	return o.data
}

func (o *Object) SetInactive() {
	o.isActive = false
}
