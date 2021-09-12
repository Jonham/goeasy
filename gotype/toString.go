package gotype

import (
	"fmt"
)

//ToString 任意类型转string。直接转化
func ToString(d interface{}) string {
	return fmt.Sprintf("%s", d)
}
