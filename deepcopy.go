package deepcopy

import (
	"reflect"
	"time"
)

func Clone(dst interface{}, original interface{}) {
	origin := reflect.ValueOf(original).Elem()
	clone := reflect.ValueOf(dst).Elem()
	copy(origin, clone)

	return
}

func copy(_origin, _clone reflect.Value) {
	oq, cq := Values{}, Values{}
	oq.Push(_origin)
	cq.Push(_clone)

	for oq.Len() > 0 {
		o, c := oq.Shift(), cq.Shift()

		switch o.Kind() {
		case reflect.Ptr:
			originVal := o.Elem()
			if originVal.IsValid() == false {
				continue
			}

			c.Set(reflect.New(originVal.Type()))

			oq.Push(originVal)
			cq.Push(c.Elem())

		case reflect.Interface:
			if o.IsNil() {
				continue
			}

			originVal := o.Elem()
			cloneVal := reflect.New(originVal.Type()).Elem()
			c.Set(cloneVal)
			oq.Push(originVal)
			cq.Push(cloneVal)

		case reflect.Struct:
			t, ok := o.Interface().(time.Time)
			if ok {
				c.Set(reflect.ValueOf(t))
				continue
			}

			for i := 0; i < o.NumField(); i++ {
				if o.Type().Field(i).PkgPath != "" {
					continue
				}

				oq.Push(o.Field(i))
				cq.Push(c.Field(i))
			}

		case reflect.Slice:
			if o.IsNil() {
				continue
			}

			c.Set(reflect.MakeSlice(o.Type(), o.Len(), o.Cap()))
			for i := 0; i < o.Len(); i++ {
				oq.Push(o.Index(i))
				cq.Push(c.Index(i))
			}

		case reflect.Map:
			if o.IsNil() {
				continue
			}
			c.Set(reflect.MakeMap(o.Type()))
			for _, key := range o.MapKeys() {
				originVal := o.MapIndex(key)
				cloneVal := reflect.New(originVal.Type()).Elem()
				var copyKey interface{}
				Clone(copyKey, key.Interface())
				c.SetMapIndex(reflect.ValueOf(copyKey), cloneVal)
			}

		default:
			c.Set(o)
		}
	}
}
