package zip

import "regexp"

//ParseToSafeFileName 处理非法文件名的字符
func ParseToSafeFileName(rawFileName string) string {
	rg := regexp.MustCompile(`[\/、]`)
	return rg.ReplaceAllString(rawFileName, "_")
}
