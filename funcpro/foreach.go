package funcpro

import (
	"reflect"
)

type MapIterFun = func(key string, value interface{})

func MapForEach(d interface{}, onEach MapIterFun) {
	keyList := reflect.ValueOf(d).MapKeys()
	for _, item := range keyList {
		key := item.String()
		value := item.MapIndex(item).Elem()
		//log.Println(key, value)
		onEach(key, value)
	}
}
