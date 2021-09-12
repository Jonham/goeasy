package gotype

import (
	"fmt"
	"reflect"
	"testing"
)

type ValueAndTypes struct {
	Value interface{}
	Type  reflect.Kind
	Fn    FuncString
}

func (v ValueAndTypes) hi(i string) string {
	return v.Type.String() + i
}

func (v ValueAndTypes) Public(i string) string {
	return v.Type.String() + i
}

func (v *ValueAndTypes) Good(i string) *ValueAndTypes {
	return v
}

func print(i interface{}) {
	fmt.Println(Any(i))
}
func TestFormat(t *testing.T) {
	i := 1
	print(i)

	var typedValueI int64 = 12
	print(typedValueI)

	var noValue ValueAndTypes
	print(noValue)

	valueTypes := ValueAndTypes{Value: nil, Type: reflect.Uint}
	print(valueTypes)
}
