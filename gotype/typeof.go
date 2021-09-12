package gotype

import (
	"fmt"
	"reflect"
)

func TypeOf(n interface{}) reflect.Kind {
	typeRef := reflect.TypeOf(n)
	return typeRef.Kind()
}

func KindOf(n interface{}) reflect.Kind {
	typeRef := reflect.TypeOf(n)
	return typeRef.Kind()
}

func TypeOfString(n interface{}) string {
	typeRef := reflect.TypeOf(n)
	return typeRef.Kind().String()
}

func ValueTypeOfString(n interface{}) string {
	typeRef := reflect.ValueOf(n)
	return typeRef.Kind().String()
}

func InterfaceToString(v interface{}) string {
	valueRef := reflect.ValueOf(v)
	switch valueRef.Kind() {
	case reflect.String:
		return valueRef.String()
	case reflect.Float64:
		return fmt.Sprintf("%f", valueRef.Float())
	}
	return fmt.Sprintf("%s", v)
}
