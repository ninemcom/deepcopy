package deepcopy

import "reflect"

type Values struct {
	queue []reflect.Value
}

func (v *Values) Len() int {
	return len(v.queue)
}

func (v *Values) Push(n reflect.Value) {
	v.queue = append(v.queue, n)
}

func (v *Values) Shift() reflect.Value {
	ret := v.queue[0]
	v.queue = v.queue[1:]

	return ret
}
