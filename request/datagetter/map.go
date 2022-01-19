package datagetter

import (
	"fmt"
	"log"
	"reflect"

	"github.com/Jonham/goeasy/checkerror"
	"github.com/Jonham/goeasy/gotype"
)

//interface快捷mapper
//缓存机制

//ProxyGetter 代理获取数据
type ProxyGetter struct {
	Raw  interface{}
	Data map[string]interface{}
}

//InitProxyGetter 初始化
func InitProxyGetter(data interface{}) ProxyGetter {
	return ProxyGetter{
		Raw: data,
	}
}

//String 直接转化为string
func (p ProxyGetter) String() string {
	return fmt.Sprintf("%s", p.Raw)
}

//GetChild 直接转化为string
func (p ProxyGetter) GetChild(field string) ProxyGetter {
	m, err := gotype.ParseToMap(p.Data)
	checkerror.CheckLog(err)
	value := reflect.ValueOf(p.Data)
	city := value.MapIndex(reflect.ValueOf("city"))
	log.Println(city)

	for _, v := range value.MapKeys() {
		field := value.MapIndex(v)
		m[v.String()] = field.Interface()
	}
	return InitProxyGetter(m[field])
}
