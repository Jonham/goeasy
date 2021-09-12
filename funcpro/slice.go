package funcpro

import "reflect"

//ToValueSlice 转化为Value的slice
func ToValueSlice (val reflect.Value) []reflect.Value {
	result := []reflect.Value{}
	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)
		result = append(result, elem)
	}
	return result
}