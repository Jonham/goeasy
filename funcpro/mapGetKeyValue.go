package funcpro

import (
	"github.com/sirupsen/logrus"
	"reflect"
)

//MapGetKeyValue Map数据获取key-value
func MapGetKeyValue(data interface{}) []FieldInfo {
	return MapValueGetKeyValue(reflect.ValueOf(data))
}

func MapValueGetKeyValue(valRef reflect.Value) []FieldInfo {
	result := []FieldInfo{}
	if valRef.Kind() == reflect.Invalid {
		//nil不处理
		return result
	}
	if valRef.Kind() != reflect.Map {
		logrus.Errorln("ERROR: keys仅处理map类型。当前类型为" + valRef.Kind().String())
		return result
	}

	for r := valRef.MapRange(); r.Next(); {
		//log.Println(r.Key().String(), r.Value().Elem().Kind())
		result = append(result, FieldInfo{
			Key:       r.Key().String(),
			Value:     r.Value().Elem(),
			ValueType: r.Value().Elem().Kind(),
		})
	}

	return result
}

//MapGetter 快速获取Map的字段
type MapGetter struct {
	Raw       reflect.Value
	FieldList []string
}

//String 获取string
func (m *MapGetter) String(fieldName string) string {
	if m.noField(fieldName) {
		return ""
	}
	return m.Raw.MapIndex(reflect.ValueOf(fieldName)).Elem().String()
}

//GetStringList 批量获取field
func (m *MapGetter) GetStringList(fieldNames []string) []string {
	if len(fieldNames) == 0 {
		return []string{}
	}
	result := []string{}
	for _, name := range fieldNames {
		result = append(result, m.String(name))
	}
	return result
}

//Float 获取string
func (m *MapGetter) Float(fieldName string) float64 {
	if m.noField(fieldName) {
		return 0
	}
	return m.Raw.MapIndex(reflect.ValueOf(fieldName)).Elem().Float()
}

//Int 获取string
func (m *MapGetter) Int(fieldName string) int64 {
	if m.noField(fieldName) {
		return 0
	}
	return m.Raw.MapIndex(reflect.ValueOf(fieldName)).Elem().Int()
}

//noField 字段不存在
func (m *MapGetter) noField(fieldName string) bool {
	for _, s := range m.FieldList {
		if s == fieldName {
			return false
		}
	}
	return true
}

//MapGetField Map的value获取字段的值
//支持string和float和int
func MapGetField(raw interface{}) *MapGetter {
	val := reflect.ValueOf(raw)
	return ValueMapGetField(val)
}

//ValueMapGetField Map的value获取字段的值
//支持string和float和int
func ValueMapGetField(val reflect.Value) *MapGetter {
	if val.Kind() != reflect.Map {
		return nil
	}
	fields := []string{}
	for _, value := range val.MapKeys() {
		fields = append(fields, value.String())
	}
	return &MapGetter{
		Raw:       val,
		FieldList: fields,
	}
}

//AnyMap 字段
type AnyMap = map[string]AnyField

type AnyField struct {
	RawValue reflect.Value
	Type     reflect.Kind
}

func (f AnyField) String() string {
	if f.Type != reflect.String {
		return ""
	}
	return f.RawValue.String()
}
func (f AnyField) Float() float64 {
	switch f.Type {
	case reflect.Float32, reflect.Float64:
		return f.RawValue.Float()
	}
	return 0
}
func (f AnyField) Int() int64 {
	switch f.Type {
	case reflect.Int, reflect.Int32, reflect.Int8, reflect.Int16, reflect.Int64:
		return f.RawValue.Int()
	}
	return 0
}

//ValueMapToAnyMap Map的value获取字段的值
//支持string和float和int
func ValueMapToAnyMap(val reflect.Value) AnyMap {
	if val.Kind() != reflect.Map {
		return nil
	}
	result := AnyMap{}
	for _, key := range val.MapKeys() {
		value := val.MapIndex(key)
		result[key.String()] = AnyField{
			RawValue: value,
			Type:     value.Kind(),
		}
	}
	return result
}
