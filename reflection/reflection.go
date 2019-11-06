package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	filed := val.Field(0)
	fn(filed.String())
}
