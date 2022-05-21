package get

import (
	"github.com/zhangdapeng520/zdpgo_file"
)

/*
@Time : 2022/5/21 9:38
@Author : 张大鹏
@File : get.go
@Software: Goland2021.3.1
@Description: get 类型的获取方法
*/

func GetFileSize(f *zdpgo_file.File) {
	data := []string{"test/a[test].txt", "test/b[test].txt", "test/c[test].txt"}
	for _, d := range data {
		f.Log.Debug("文件大小", "file", d, "result", f.GetFileSize(d))
	}
}

func GetDirectorySize(f *zdpgo_file.File) {
	f.Log.Debug("文件夹大小", "test", f.GetDirSize("test"))
	f.Log.Debug("文件夹大小", "test111", f.GetDirSize("test111"))
}
