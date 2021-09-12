package gotype

import (
	"log"
	"reflect"
	"strings"
)

var verboseLog = false

//InterfaceGetField 根据字段，获取指定路径的数据
func InterfaceGetField(d interface{}, pathName string) reflect.Value {
	pathList := strings.Split(pathName, ".")

	//TODO 需要剔除，不合法的基础类型。只有map才能获取
	valueRef := reflect.ValueOf(d)
	if valueRef.Kind().String() != "map" {
		panic("只支持map获取字段" + valueRef.Kind().String())
	}

	if verboseLog {
		log.Println("pathList: ", pathList)
	}

	if len(pathList) == 0 {
		return valueRef
	}
	for index, p := range pathList {
		if verboseLog {
			log.Println("==finding==>", p)
		}
		mapKeys := valueRef.MapKeys()

		for _, value := range mapKeys {
			if verboseLog {
				log.Println("->", value)
			}
			if value.String() == p {
				valueRef = valueRef.MapIndex(value).Elem()
				t := valueRef.Kind().String()

				if index == len(pathList)-1 {
					//找到结果了
					return valueRef
				}
				if t != "map" {
					panic("获取Field失败: 类型： " + t)
				}
				break
			}
		}
	}

	return reflect.ValueOf("")
}

//GetMapOnField 指定字段是否是Map类型
func GetMapOnField(mapValue reflect.Value, field string) (interface{}, bool) {
	val := InterfaceGetField(mapValue.Interface(), field)
	if val.IsValid() && val.Kind() == reflect.Map {
		return val.Interface(), true
	}
	return nil, false
}
