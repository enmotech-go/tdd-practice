package main

import "reflect"

func Walk(in interface{}, fn func(input string)) {
	val := reflect.ValueOf(in)
	fieldOne := val.Field(0)
	fn(fieldOne.String())
}
