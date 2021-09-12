package gotype

import (
	"fmt"
	"reflect"
)

func display(path string, v reflect.Value) {
	switch v.Kind() {

	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)

	case reflect.Slice, reflect.Array:
		fmt.Printf("%s: [\n", path)
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("\t[%d]", i), v.Index(i))
		}
		fmt.Printf("]\n")

	case reflect.Struct:
		fmt.Printf("\t-Type: %s \n\t-Kind: %s \n", v.Type(), v.Kind())
		fmt.Printf("%s: {\n", path)
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("\t%s", v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
		fmt.Printf("}\n")

		fmt.Println(v.NumMethod())
		for i := 0; i < v.NumMethod(); i++ {
			fieldPath := fmt.Sprintf("\t%s", v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}

	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}

	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s: nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}

	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s: nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}

	default: // basic types, channels, funcs
		fmt.Printf("%s: %s\n", path, formatAtom(v))
	}
}
