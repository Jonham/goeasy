package download

import (
	"io"
	"net/http"
	"os"
)

// SaveHTTPResouceAsFile 下载文件，并保存到指定位置
func SaveHTTPResouceAsFile(fileURL, filePath string) error {
	res, err := http.Get(fileURL)
	if err != nil {
		return err
	}

	f, _ := os.Create(filePath)
	_, err = io.Copy(f, res.Body)
	return err
}
