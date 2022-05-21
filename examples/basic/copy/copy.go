package copy

import (
	"github.com/zhangdapeng520/zdpgo_file"
)

/*
@Time : 2022/5/21 9:59
@Author : 张大鹏
@File : copy.go
@Software: Goland2021.3.1
@Description: copy复制文件相关
*/

func CopyFile(f *zdpgo_file.File) {
	testData := []struct {
		Src  string
		Dest string
	}{
		{"test/a[test].txt", "test/test11.txt"},
		{"test/a[test].txt", "test/test22.txt"},
		{"test/a[test].txt", "test/test33.txt"},
		{"test/a[test].txt", "test111/test11.txt"},
		{"test/a[test].txt", "test111/test22.txt"},
		{"test/a[test].txt", "test111/test33.txt"},
	}

	for _, data := range testData {
		b := f.CopyFile(data.Src, data.Dest)
		f.Log.Debug("复制文件", "src", data.Src, "dest", data.Dest, "result", b)
	}
}
