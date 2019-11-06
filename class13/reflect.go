package reflect

import "reflect"

func Walk(in interface{}, fn func(input string)) {
	val := getValue(in)

	walkValue := func(val reflect.Value) {
		Walk(val.Interface(), fn)
	}

	switch val.Kind() {
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

	case reflect.String:
		fn(val.String())
	}

}

func getValue(i interface{}) reflect.Value {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}
