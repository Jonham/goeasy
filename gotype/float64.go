package gotype

import "reflect"

//GetFloat64 转化为float64
// value float64, parseResult bool
func GetFloat64(v interface{}) (value float64, parseResult bool) {
	typeRef := reflect.ValueOf(v)
	if typeRef.Kind() == reflect.Float64 {
		return typeRef.Float(), true
	}
	return 0, false
}

//GetInt64 转化为Int64
// value float64, parseResult bool
func GetInt64(v interface{}) (value int64, parseResult bool) {
	typeRef := reflect.ValueOf(v)
	switch typeRef.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return typeRef.Int(), true
	}
	return 0, false
}
