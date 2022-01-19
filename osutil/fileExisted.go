package osutil

import "os"

// FileExisted 文件是否存在
func FileExisted(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}
