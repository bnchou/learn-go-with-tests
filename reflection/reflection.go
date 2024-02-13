package reflection

import (
	"reflect"
)

func walk(x any, f func(input string)) {
	val := getValue(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			f(field.String())
		case reflect.Struct:
			walk(field.Interface(), f)
		}
	}
}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
