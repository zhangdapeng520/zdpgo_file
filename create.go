package zdpgo_file

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

/*
@Time : 2022/5/12 14:49
@Author : 张大鹏
@File : create.go
@Software: Goland2021.3.1
@Description: 创建相关的方法
*/

// CreateDir 调用os.MkdirAll递归创建文件夹
func (f *File) CreateDir(dirtPath string) bool {
	if !f.IsExists(dirtPath) {
		err := os.MkdirAll(dirtPath, os.ModePerm)
		if err != nil {
			f.Log.Error("创建文件夹失败", "error", err, "dirPath", dirtPath)
			return false
		}
	}
	return true
}

// CreateDirFile 在指定文件夹中创建文件
func (f *File) CreateDirFile(dirtPath, fileName, content string) bool {
	// 文件夹不存在
	if !f.IsExists(dirtPath) {
		f.CreateDir(dirtPath)
	}

	// 拼接文件路径
	filePath := filepath.Join(dirtPath, fileName)

	// 写入数据
	err := ioutil.WriteFile(filePath, []byte(content), os.ModePerm)
	if err != nil {
		f.Log.Error("写入数据到文件失败", "error", err, "filePath", filePath)
		return false
	}

	return true
}

// CreateFile 创建文件
func (f *File) CreateFile(filePath, content string) bool {
	// 创建文件夹
	dirPath := f.GetFileDir(filePath)
	if !f.IsExists(dirPath) {
		f.CreateDir(dirPath)
	}

	// 创建文件
	err := ioutil.WriteFile(filePath, []byte(content), os.ModePerm)
	if err != nil {
		f.Log.Error("写入数据到文件失败", "error", err, "filePath", filePath)
		return false
	}
	return true
}
