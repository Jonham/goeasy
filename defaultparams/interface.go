package defaultparams

type Any = interface{}

//Interface 返回第一个数据
func Interface(other []Any, defaultStr Any) Any {
	if len(other) > 0 {
		return other[0]
	}
	return defaultStr
}
