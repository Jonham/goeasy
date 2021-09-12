package defaultparams

type MapAny = map[string]Any

//MapAny 返回第一个数据
func Map(other []MapAny, defaultStr MapAny) MapAny {
	if len(other) > 0 {
		return other[0]
	}
	return defaultStr
}
