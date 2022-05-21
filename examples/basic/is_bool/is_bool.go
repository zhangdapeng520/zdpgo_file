package is_bool

import (
	"github.com/zhangdapeng520/zdpgo_file"
)

/*
@Time : 2022/5/21 9:35
@Author : 张大鹏
@File : is_bool.go
@Software: Goland2021.3.1
@Description: is类型的判断方法
*/

func IsExists(f *zdpgo_file.File) {
	data := []string{"test/a[test].txt", "test/b[test].txt", "test/c[test].txt"}
	for _, d := range data {
		b := f.IsExists(d)
		f.Log.Debug("文件是否存在", "file", d, "result", b)
	}
}

// 判断文件夹中是否包含某个文件
func IsDirContainsFile(f *zdpgo_file.File) {
	data := []struct {
		DirPath  string
		FileName string
	}{
		{"test", "a[test].txt"},
		{"test/a", "a"},
		{"test/a", "a.txt"},
		{"test/b", "b"},
		{"test/b", "b.txt"},
		{"test/c", "c"},
		{"test/c", "c.txt"},
		{"test", "a.txt"},
		{"test", "b.txt"},
	}

	for _, d := range data {
		b := f.IsDirContainsFile(d.DirPath, d.FileName)
		f.Log.Debug("文件夹是否包含文件", "dir", d.DirPath, "file", d.FileName, "result", b)
	}
}
