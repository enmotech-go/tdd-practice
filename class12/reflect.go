package class12

import "reflect"

func Walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}
