package file

import (
	"os"
)

// File 文件对象
type File struct {
}

func New() *File {
	f := File{}
	return &f
}

// Exists 判断文件是否存在
func (f *File) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// Size 获取文件大小
func (f *File) Size(path string) int64 {
	if !f.Exists(path) {
		return 0
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}
