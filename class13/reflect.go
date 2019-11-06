package main

import "reflect"

func Walk(in interface{}, fn func(input string)) {
	val := reflect.ValueOf(in)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}

	}
}
