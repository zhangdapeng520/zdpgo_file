package zdpgo_file

import "os"

/*
@Time : 2022/5/12 14:49
@Author : 张大鹏
@File : create.go
@Software: Goland2021.3.1
@Description: 创建相关的方法
*/

// CreateMultiDir 调用os.MkdirAll递归创建文件夹
func (f *File) CreateMultiDir(filePath string) error {
	if !f.IsExists(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}
