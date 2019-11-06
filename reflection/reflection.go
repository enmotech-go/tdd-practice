package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		filed := val.Field(i)
		if filed.Kind() == reflect.String {
			fn(filed.String())
		}

		if filed.Kind() == reflect.Struct {
			walk(filed.Interface(), fn)
		}
	}
}
