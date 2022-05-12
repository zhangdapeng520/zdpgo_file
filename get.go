package zdpgo_file

import (
	"os"
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
