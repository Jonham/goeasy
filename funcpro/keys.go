package funcpro

import (
	"log"
	"reflect"
)

//StructKeys 获取结构体的字段名
func StructKeys (data interface{})[]string {
	return getKeyList(data)
}

//TODO 处理非struct类型
func getKeyList(data interface{}) []string {
	result := []string{}
	valRef := reflect.ValueOf(data)
	t := valRef.Kind().String()
	if t != "struct" {
		log.Println("ERROR: keys仅处理struct类型。当前类型为" + t)
		return result
	}

	typeRef := reflect.TypeOf(data)
	fieldCount := typeRef.NumField()
	//log.Println(fieldCount)
	for i := 0; i < fieldCount; i++ {
		field := typeRef.Field(i)
		result = append(result, field.Name)
	}

	return result
}
