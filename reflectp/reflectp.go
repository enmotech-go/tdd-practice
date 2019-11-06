package reflectp

import "reflect"

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func walk(x interface{}, fn func(input string))  {
	val := getValue(x)
	numOfValues := 0
	var getField func(int) reflect.Value
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice:
		numOfValues = val.Len()
		getField = val.Index
	}

	for i:=0;i<numOfValues;i++ {
		walk(getField(i).Interface(), fn)
	}
}
