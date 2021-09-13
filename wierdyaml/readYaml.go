package wierdyaml

import (
	"log"
	"os"

	"open-git.jonham.cn/Jonham/goeasy/checkerror"
)

//ReadYAML 读取YAML
func ReadYAML(filename string) {
	b, err := os.ReadFile(filename)
	checkerror.CheckLog(err)

	if err == nil {
		str := string(b)
		log.Println(str)
	}
}
