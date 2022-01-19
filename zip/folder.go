package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/Jonham/goeasy/checkerror"
)

//ZipFilesInFolder 压缩所有文件夹里的文件
func ZipFilesInFolder(srcFile string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer func(zipfile *os.File) {
		err := zipfile.Close()
		if err != nil {
			checkerror.CheckLog(err)
		}
	}(zipfile)

	archive := zip.NewWriter(zipfile)
	defer func(archive *zip.Writer) {
		err := archive.Close()
		if err != nil {
			checkerror.CheckLog(err)
		}
	}(archive)

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(srcFile)+"/")
		// header.Name = path
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					checkerror.CheckLog(err)
				}
			}(file)
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}
