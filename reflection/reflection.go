package reflection

import (
	// "fmt"
	"reflect"
)

/**
class_12
demand: 编写函数walk(x interface{}, fn func(string))，对x中的所有字符串字段递归地调用fn函数。
反射提供检查自身结构体类型的能力，是元编程的一种形式。使用Interface{}会失去对类型安全的检查，
*/

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn) // 回调
	}

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
		for _, k := range val.MapKeys() {
			walkValue(val.MapIndex(k))
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
