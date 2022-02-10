package zdpgo_file

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetDirectorySize 文件夹占用磁盘大小
func (f *File) GetDirectorySize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// CreateMultiDir 调用os.MkdirAll递归创建文件夹
func (f *File) CreateMultiDir(filePath string) error {
	if !f.IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败,error info:", err)
			return err
		}
		return err
	}
	return nil
}

// IsExist 判断所给路径文件/文件夹是否存在(返回true是存在)
func (f *File) IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
