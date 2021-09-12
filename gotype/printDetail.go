package gotype

import (
	"fmt"
	"strings"
)

//PrintDetail 打印详细结构
func PrintDetail (m interface{}) {
	fmt.Printf("%#v\n", m)
}

//PrintDetailBeautified 打印详细结构
func PrintDetailBeautified (m interface{}) {
	content := fmt.Sprintf("%#v", m)
	content = strings.ReplaceAll(content, ",", ",\n\t")
	content = strings.Replace(content, "{", "{\n\t", 1)
	//fmt.Printf("%s%s", t, content)
	fmt.Println(content)
}
