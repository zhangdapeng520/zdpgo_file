package directory

import (
	"os"
	"path/filepath"
)

type Directory struct {
}

func New() *Directory {
	d := Directory{}

	return &d
}

// GetDirectorySize 文件夹占用磁盘大小
func (f *Directory) GetDirectorySize(path string) (int64, error) {
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
func (f *Directory) CreateMultiDir(filePath string) error {
	if !f.IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

// MakeDirs 创建多级目录
func (f *Directory) MakeDirs(filePath string) error {
	return f.CreateMultiDir(filePath)
}

// IsExist 判断所给路径文件/文件夹是否存在(返回true是存在)
func (f *Directory) IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
