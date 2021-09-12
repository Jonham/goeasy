package checkerror

import "log"

func CheckLog(err error) {
	if err != nil {
		log.Println("ERR: " + err.Error())
	}
}

//Check 如果有错误直接退出
func Check(err error) {
	if err != nil {
		log.Println("ERR: " + err.Error())
		panic(err)
	}
}
