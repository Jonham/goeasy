package osutil

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
)

// FileMD5sum 获取指定文件的md5
func FileMD5sum(fp string) (result string, err error) {
	if !FileExisted(fp) {
		err = errors.New("fileNotExist")
		return
	}

	f, err := os.Open(fp)
	if err != nil {
		return
	}
	defer f.Close()

	h := md5.New()
	if _, err = io.Copy(h, f); err != nil {
		return
	}
	result = fmt.Sprintf("%x", h.Sum(nil))
	return
}
