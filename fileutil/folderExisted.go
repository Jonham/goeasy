package fileutil

import "os"

// FolderExisted 文件夹是否存在
func FolderExisted(p string) (bool, error) {
	info, err := os.Stat(p)

	if err != nil {
		return false, err
	}
	if !info.IsDir() {
		return false, err
	}
	return true, nil
}
