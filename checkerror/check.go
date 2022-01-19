package checkerror

import (
	"log"
	"strings"
)

func CheckLog(err error, tags ...string) {
	if err != nil {
		prefix := ""
		if len(tags) > 0 {
			prefix = strings.Join(tags, "\n") + "\n"
		}
		log.Println(prefix + "ERR: " + err.Error())
	}
}

//Check 如果有错误直接退出
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
