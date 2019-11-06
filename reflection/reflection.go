package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		filed := val.Field(i)

		switch filed.Kind() {
		case reflect.String:
			fn(filed.String())
		case reflect.Struct:
			walk(filed.Interface(), fn)
		}
	}
}
