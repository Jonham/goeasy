package gotype

import (
	"fmt"
	"reflect"
	"testing"
)

func displayValue(path string, value interface{}) {
	display(path, reflect.ValueOf(value))
}

type FuncString = func(v interface{}) string

func TestDisplay(t *testing.T) {
	strList := []string{"JOnham", "JCC"}
	fmt.Println("strList")
	display("strList", reflect.ValueOf(strList))

	fmt.Println("")
	var noValue ValueAndTypes
	fmt.Println("noValue ValueAndTypes")
	displayValue("ValueAndTypes", noValue)

	fmt.Println("")
	valueType := ValueAndTypes{}
	fmt.Println("valueType ValueAndTypes")
	displayValue("ValueAndTypes", valueType)

	fmt.Println("")
	var noFunc FuncString
	fmt.Println("noFunc FuncString")
	displayValue("FuncString", noFunc)
}
