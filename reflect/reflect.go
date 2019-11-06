package reflect

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	var numOfFields int
	var getField func(int) reflect.Value
	walkValue := func(val reflect.Value) {
		walk(val.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		numOfFields = val.Len()
		getField = val.Index
	case reflect.Struct:
		numOfFields = val.NumField()
		getField = val.Field
	case reflect.String:
		fn(val.String())
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	}

	for i := 0; i < numOfFields; i++ {
		walkValue(getField(i))
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
