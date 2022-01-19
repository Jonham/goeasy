package checkerror

import (
	"log"
	"strings"
)

// CheckLog 如果错误存在，打印错误，并返回true
// 	示例：
// 	err := errors.New("someError")
// 	if checkerror.CheckError(err, "error happens") return err
// ---
// 	print logs as follow:
// 	error happens
// 	someError
func CheckLog(err error, tags ...string) (hasError bool) {
	if err != nil {
		prefix := ""
		if len(tags) > 0 {
			prefix = strings.Join(tags, "\n") + "\n"
		}
		log.Println(prefix + "ERR: " + err.Error())
		return true
	}
	return
}

// Check 如果有错误直接退出
// 	will panic
func Check(err error, tags ...string) {
	if err != nil {
		prefix := ""
		if len(tags) > 0 {
			prefix = strings.Join(tags, "\n") + "\n"
		}
		log.Println(prefix + "ERR: " + err.Error())
		panic(err)
	}
}
