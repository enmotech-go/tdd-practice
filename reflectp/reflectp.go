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

func walk(x interface{}, fn func(input string))  {
	val := reflect.ValueOf(x)
	for i:=0;i<val.NumField();i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}
		if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
	}
}
