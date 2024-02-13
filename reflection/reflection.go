package reflection

import (
	"reflect"
)

func walk(x any, f func(input string)) {
	v := reflect.VisibleFields(reflect.TypeOf(x))
	for _, vf := range v {
		if vf.Type.Kind() == reflect.String {
			f(reflect.ValueOf(x).FieldByName(vf.Name).String())
		}
		if vf.Type.Kind() == reflect.Struct {
			walk(reflect.ValueOf(x).FieldByName(vf.Name).Interface(), f)
		}
	}
}
