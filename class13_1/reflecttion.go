package class13_1

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	numberOfValues := 0
	var getField func(int2 int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}

	}
	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
