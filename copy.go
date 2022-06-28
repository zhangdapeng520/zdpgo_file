package zdpgo_file

import "io/ioutil"

/*
@Time : 2022/5/12 15:37
@Author : 张大鹏
@File : copy.go
@Software: Goland2021.3.1
@Description: copy复制相关的方法
*/

// CopyFile 复制文件
// @param src 源文件地址
// @param dest 目标文件地址
func (f *File) CopyFile(src, dest string) bool {
	// 读取源文件
	srcFile, err := ioutil.ReadFile(src)
	if err != nil {
		return false
	}

	// 创建目标文件夹
	destDir := f.GetFileDir(dest)
	if !f.IsExists(destDir) {
		f.CreateDir(destDir)
	}

	// 写入目标文件
	err = ioutil.WriteFile(dest, srcFile, 0644)
	if err != nil {
		return false
	}

	return true
}
