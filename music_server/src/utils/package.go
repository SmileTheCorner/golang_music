package utils

import (
	"archive/zip"
	"io"
	"strings"
)

// 将文件打成zip包
func ZipFiles(name string, data map[string]string, buf io.Writer) error {
	//创建压缩包
	zipWriter := zip.NewWriter(buf)

	//遍历数据
	for k, v := range data {
		writer, err := zipWriter.Create(k + "/" + k + ".go")
		if err != nil {
			return err
		}
		if _, err := io.Copy(writer, strings.NewReader(v)); err != nil {
			return err
		}
	}

	//关闭zipWriter
	defer zipWriter.Close()
	return nil
}
