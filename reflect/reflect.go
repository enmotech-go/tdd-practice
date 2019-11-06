package reflect

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	var numOfFields int
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		numOfFields = val.Len()
		getField = val.Index
	case reflect.Struct:
		numOfFields = val.NumField()
		getField = val.Field
	case reflect.String:
		fn(val.String())
	}

	for i := 0; i < numOfFields; i++ {
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
