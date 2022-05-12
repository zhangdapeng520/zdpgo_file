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
func (f *File) CopyFile(src, dest string) error {
	// 读取源文件
	srcFile, err := ioutil.ReadFile(src)
	if err != nil {
		f.Log.Error("读取源文件失败", "error", err, "src", src)
		return err
	}

	// 写入目标文件
	err = ioutil.WriteFile(dest, srcFile, 0644)
	if err != nil {
		f.Log.Error("写入目标文件失败", "error", err, "dest", dest)
		return err
	}

	return nil
}