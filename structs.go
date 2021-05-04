package deepcopy

import (
	"math"
	"reflect"
)

type values struct {
	queue []reflect.Value
}

func (v *values) Init(c int) {
	v.queue = make([]reflect.Value, 0, c)
}

func (v *values) Len() int {
	return len(v.queue)
}

func (v *values) Cap() int {
	return cap(v.queue)
}

func (v *values) Push(n reflect.Value) {
	v.queue = append(v.queue, n)
}

func (v *values) Shift() reflect.Value {
	ret := v.queue[0]
	v.queue = v.queue[1:]

	return ret
}

var sizeTable = []int{8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}

type sizeCache map[string]int

func (c sizeCache) Register(t string, s int) {
	i := 0
	if s > 0 {
		i = int(math.Log2(float64(s))) - 4
		if i < 0 {
			i = 0
		} else if i >= len(sizeTable) {
			i = len(sizeTable) - 1
		}
	}
	c[t] = sizeTable[i]
}

func (c sizeCache) ObtainSize(t string) int {
	if s, ok := c[t]; ok {
		return s
	}

	return sizeTable[0]
}
