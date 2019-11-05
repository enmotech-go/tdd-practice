package class13

import "reflect"

func walk(x interface{}, fn func(input string)) {
	var val = reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())

}
