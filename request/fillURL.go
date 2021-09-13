package zhuhe

import (
	"fmt"
	"open-git.jonham.cn/Jonham/goeasy/funcpro"
	"reflect"
	"strconv"
	"strings"
)

//FillURL 填充URL中的模板值
func (api API) FillURL(urlTemplate string, params interface{}) string {
	return api.FillURLMap(urlTemplate, params)
}

//FillURLMap 填充URL中的模板值
func (api API) FillURLMap(urlTemplate string, params interface{}) string {
	keyValueList := funcpro.MapGetKeyValue(params)
	//keys组合成模板字符串
	for _, fieldInfo := range keyValueList {
		oldValueTemplate := fmt.Sprintf("${%s}", fieldInfo.Key)

		switch fieldInfo.ValueType {
		case reflect.String:
			urlTemplate = strings.Replace(urlTemplate, oldValueTemplate, fieldInfo.Value.String(), -1)
		case reflect.Float64:
			fallthrough
		case reflect.Float32:
			floatStr := strconv.FormatFloat(fieldInfo.Value.Float(), 'g', 64, 64)
			urlTemplate = strings.Replace(urlTemplate, oldValueTemplate, floatStr, -1)
		case reflect.Int:
			fallthrough
		case reflect.Int64:
			fallthrough
		case reflect.Int32:
			urlTemplate = strings.Replace(urlTemplate, oldValueTemplate, strconv.Itoa(int(fieldInfo.Value.Int())), -1)
		}
	}
	return urlTemplate
}

//FillURLStruct 填充URL中的模板值，使用结构体作为参数
func (api API) FillURLStruct(urlTemplate string, params interface{}) string {
	//strings.Replace(s, old, new, -1)
	keyValueList := funcpro.StructGetKeyValue(params)
	//keys组合成模板字符串
	for _, fieldInfo := range keyValueList {
		oldValueTemplate := fmt.Sprintf("${%s}", fieldInfo.Key)

		switch fieldInfo.ValueType {
		case reflect.String:
			urlTemplate = strings.Replace(urlTemplate, oldValueTemplate, fieldInfo.Value.String(), -1)
		case reflect.Int:
			fallthrough
		case reflect.Int64:
			fallthrough
		case reflect.Int32:
			urlTemplate = strings.Replace(urlTemplate, oldValueTemplate, strconv.Itoa(int(fieldInfo.Value.Int())), -1)
		}
	}
	return urlTemplate
}
