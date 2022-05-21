package zdpgo_file

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

/*
@Time : 2022/5/12 14:54
@Author : 张大鹏
@File : get.go
@Software: Goland2021.3.1
@Description: get类型的获取方法
*/

// GetDirSize 文件夹占用磁盘大小
func (f *File) GetDirSize(path string) (size int64) {
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	if err != nil {
		f.Log.Error("获取文件夹大小失败", "error", err, "dirPath", path)
	}
	return
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

// GetFileNameNoSuffix 获取文件名称，不包含后缀
func (f *File) GetFileNameNoSuffix(absolutePath string) (fileName string) {
	fileName = filepath.Base(absolutePath)
	suffix := f.GetFileSuffix(fileName)
	fileName = strings.Replace(fileName, suffix, "", 1)
	return
}

// GetFileSuffix 获取文件后缀
func (f *File) GetFileSuffix(absolutePath string) (suffix string) {
	suffix = path.Ext(absolutePath)
	return
}
