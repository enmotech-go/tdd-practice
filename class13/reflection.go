package class13

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	}

}

func getValue(x interface{}) reflect.Value {
	var val = reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
