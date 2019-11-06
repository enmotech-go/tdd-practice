package reflection

import (
	// "fmt"
	"reflect"
)

/**
func main() {
	fmt.Println("开始进行reflect用法学习")
	mock := struct {
		// Age int
		Name string
	}{
		"小明",
	}
	fn := func(s string) {
		fmt.Printf("val is %s \n", s)
	}
	walk(mock, fn)
	fmt.Println("reflect用法学习结束啦！")
}
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
