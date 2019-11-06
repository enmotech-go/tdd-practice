package reflect

import "reflect"

func Walk(in interface{}, fn func(input string)) {
	val := getValue(in)

	var getFiledFun func(i int) reflect.Value
	var numberOfValue int

	switch val.Kind() {
	case reflect.Struct:
		numberOfValue = val.NumField()
		getFiledFun = val.Field
	case reflect.Slice:
		numberOfValue = val.Len()
		getFiledFun = val.Index
	case reflect.String:
		fn(val.String())
	}
	for i := 0; i < numberOfValue; i++ {
		Walk(getFiledFun(i).Interface(), fn)
	}

}

func getValue(i interface{}) reflect.Value {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}
