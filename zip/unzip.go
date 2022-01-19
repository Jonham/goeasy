package zip

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var ErrInvalidFilePath = errors.New("invalid file path")

// Unzip 解压文件到指定目录
func Unzip(zipFilePath, outputFolder string, verbose ...string) error {
	verboseOutput := len(verbose) > 0
	dst := outputFolder

	archive, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)
		if verboseOutput {
			fmt.Println("unzipping file ", filePath)
		}

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			return ErrInvalidFilePath
		}
		if f.FileInfo().IsDir() {
			if verboseOutput {
				fmt.Println("creating directory..." + filePath)
			}
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		// permission ?
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}

		dstFile.Close()
		fileInArchive.Close()
	}
	return nil
}
