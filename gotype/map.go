package gotype

import (
	"reflect"
)

//SliceValueGetIndex 从slice的value中获取指定index的值
func SliceValueGetIndex(mapValue reflect.Value, index int) reflect.Value {
	item := mapValue.Index(index)
	return item.Elem()
}

type StringMapInterface = map[string]interface{}
type ParseToMapError struct {
	Msg string
}

func (t ParseToMapError) Error() string {
	return t.Msg
}

//ParseToMap 转显性Map
func ParseToMap(raw interface{}) (StringMapInterface, error) {
	m := StringMapInterface{}
	value := reflect.ValueOf(raw)
	if value.Kind() != reflect.Map {
		return m, ParseToMapError{"数据类型不是Map"}
	}

	for _, v := range value.MapKeys() {
		field := value.MapIndex(v)
		m[v.String()] = field.Interface()
	}
	return m, nil
}
