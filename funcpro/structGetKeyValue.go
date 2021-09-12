package funcpro

import (
	"log"
	"reflect"
)

//StructGetKeyValue 获取结构体的字段名和值类型
func StructGetKeyValue(data interface{}) []FieldInfo {
	return getKeyValue(data)
}

type FieldInfo struct {
	Key       string
	ValueType reflect.Kind
	Value     reflect.Value
}

func getKeyValue(data interface{}) []FieldInfo {
	result := []FieldInfo{}
	valRef := reflect.ValueOf(data)
	t := valRef.Kind().String()
	if t != "struct" {
		log.Println("ERROR: keys仅处理struct类型。当前类型为" + t)
		return result
	}

	typeRef := reflect.TypeOf(data)
	fieldCount := typeRef.NumField()
	for i := 0; i < fieldCount; i++ {
		field := typeRef.Field(i)
		value := valRef.Field(i)
		result = append(result, FieldInfo{
			Key:       field.Name,
			Value:     value,
			ValueType: value.Kind(),
		})
	}

	return result
}
