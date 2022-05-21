package delete

import (
	"github.com/zhangdapeng520/zdpgo_file"
)

/*
@Time : 2022/5/21 10:20
@Author : 张大鹏
@File : delete.go
@Software: Goland2021.3.1
@Description: delete 删除相关
*/

func DeleteDirFile(f *zdpgo_file.File) {
	data := []struct {
		DirPath  string
		FileName string
	}{
		{"test", "test11.txt"},
		{"test", "test22.txt"},
		{"test", "test33.txt"},
	}

	for _, d := range data {
		b := f.DeleteDirFile(d.DirPath, d.FileName)
		f.Log.Debug("删除文件夹中的文件", "dir", d.DirPath, "file", d.FileName, "result", b)
	}
}

// 测试删除文件夹
func DeleteDir(f *zdpgo_file.File) {
	paths := []string{"test", "test111"}

	for _, d := range paths {
		b := f.DeleteDir(d)
		f.Log.Debug("删除文件夹", "dir", d, "result", b)
	}
}
