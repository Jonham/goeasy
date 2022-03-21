package osutil

import (
	"io"
	"io/fs"
	"io/ioutil"
)

// PipeTo 将reader内容导向本地一个文件
func PipeTo(reader io.Reader, filename string, perm fs.FileMode) {
	b, _ := ioutil.ReadAll(reader)
	ioutil.WriteFile(filename, b, perm)
}
