package osutil

import "os"

// EnsureFolder 保证文件夹存在
func EnsureFolder(p string) (err error) {
	existed, err := FolderExisted(p)
	if existed {
		return
	}
	// 新建文件夹
	os.MkdirAll(p, os.ModePerm)
	return
}
