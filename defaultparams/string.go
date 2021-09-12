package defaultparams

//String 返回第一个字符串或者默认字符串
func String(other []string, defaultStr string) string {
	if len(other) > 0 {
		return other[0]
	}
	return defaultStr
}
//StringNotEmpty 返回第一个字符串或者默认字符串
func StringNotEmpty(other []string, defaultStr string) string {
	if len(other) > 0 && other[0] != "" {
		return other[0]
	}
	return defaultStr
}
