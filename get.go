package zdpgo_file

import (
	"os"
	"path"
	"path/filepath"
)

/*
@Time : 2022/5/12 14:54
@Author : 张大鹏
@File : get.go
@Software: Goland2021.3.1
@Description: get类型的获取方法
*/

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

// GetFileDirAndFileName 获取文件路径和文件名
func (f *File) GetFileDirAndFileName(absolutePath string) (fileDir string, fileName string) {
	fileDir, fileName = filepath.Split(absolutePath)
	return
}

// GetFileDir 获取文件路径
func (f *File) GetFileDir(absolutePath string) (fileDir string) {
	fileDir, _ = filepath.Split(absolutePath)
	return
}

// GetFileName 获取文件名称
func (f *File) GetFileName(absolutePath string) (fileName string) {
	fileName = filepath.Base(absolutePath)
	return
}

// GetFileSuffix 获取文件后缀
func (f *File) GetFileSuffix(absolutePath string) (suffix string) {
	suffix = path.Ext(absolutePath)
	return
}
