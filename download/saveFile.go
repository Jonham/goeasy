package download

import (
	"io"
	"net/http"
	"os"

	"github.com/Jonham/goeasy/checkerror"
)

// SaveHTTPResouceAsFile 下载文件，并保存到指定位置
func SaveHTTPResouceAsFile(fileURL, filePath string) {
	res, err := http.Get(fileURL)
	checkerror.Check(err)

	f, _ := os.Create(filePath)
	io.Copy(f, res.Body)
}
