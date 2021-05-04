package deepcopy

import "reflect"

type values struct {
	queue []reflect.Value
}

func (v *values) Init(c int) {
	v.queue = make([]reflect.Value, 0, c)
}

func (v *values) Len() int {
	return len(v.queue)
}

func (v *values) Push(n reflect.Value) {
	v.queue = append(v.queue, n)
}

func (v *values) Shift() reflect.Value {
	ret := v.queue[0]
	v.queue = v.queue[1:]

	return ret
}
