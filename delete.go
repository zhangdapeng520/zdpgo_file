package zdpgo_file

import (
	"os"
	"path/filepath"
)

/*
@Time : 2022/5/12 16:00
@Author : 张大鹏
@File : delete.go
@Software: Goland2021.3.1
@Description: delete删除相关的方法
*/

// DeleteDirFile 删除指定目录中的指定文件
func (f *File) DeleteDirFile(dirPath, fileName string) bool {
	// 拼接文件路径
	filePath := filepath.Join(dirPath, fileName)

	// 删除文件
	err := os.Remove(filePath) //删除文件test.txt
	if err != nil {
		f.Log.Error("删除指定目录中的指定文件失败", "error", err, "dirPath", dirPath, "fileName", fileName)
		return false
	}
	return true
}

// DeleteDir 删除文件夹
func (f *File) DeleteDir(dirPath string) bool {
	err := os.RemoveAll(dirPath)
	if err != nil {
		f.Log.Error("删除文件夹失败", "error", err, "dirPath", dirPath)
		return false
	}
	return true
}
