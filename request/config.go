package zhuhe

import (
	"os"
	"path"
)

var BaseURL string = "https://test.rlb-cn.com/request"

//导出目录
var ExportFileFolder = "temp/"

func init() {
	//从环境变量，获取BaseURL
	devBaseURL := os.Getenv("BaseURL")
	BaseURL = devBaseURL

	folder := os.Getenv("ExportFileFolder")
	if folder != "" {
		ExportFileFolder = folder
	}
}

//ParseDownloadFolder 合并为下载目录
func ParseDownloadFolder(pathName string) string {
	return path.Join(ExportFileFolder, pathName)
}
