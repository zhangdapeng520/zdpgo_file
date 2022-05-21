package create

import (
	"github.com/zhangdapeng520/zdpgo_file"
)

/*
@Time : 2022/5/21 10:02
@Author : 张大鹏
@File : create.go
@Software: Goland2021.3.1
@Description: create 创建文件相关
*/

func CreateDir(f *zdpgo_file.File) {
	dirs := []string{
		"test",
		"test/a",
		"test/b",
		"test/c",
	}

	for _, d := range dirs {
		b := f.CreateDir(d)
		f.Log.Debug("创建文件夹", "dir", d, "result", b)
	}
}

// 测试在指定文件夹中创建指定文件
func CreateDirFile(f *zdpgo_file.File) {
	data := []struct {
		DirPath  string
		FileName string
		Content  string
	}{
		{"test", "a[test].txt", "aaa"},
		{"test", "b[test].txt", "bbb"},
		{"test", "c[test].txt", "bbb"},
		{"test/a", "a.txt", "ccc"},
		{"test/b", "b.txt", "ccc"},
		{"test/c", "c.txt", "ccc"},
		{"test/d", "d.txt", "ccc"},
	}

	for _, d := range data {
		b := f.CreateDirFile(d.DirPath, d.FileName, d.Content)
		f.Log.Debug("指定目录创建文件", "dir", d.DirPath, "file", d.FileName, "result", b)
	}
}

// 测试创建文件
func CreateFile(f *zdpgo_file.File) {
	data := []struct {
		FilePath string
		Content  string
	}{
		{"test/a[test].txt", "aaa"},
		{"test/b[test].txt", "aaa"},
		{"test/c[test].txt", "aaa"},
	}

	for _, d := range data {
		b := f.CreateFile(d.FilePath, d.Content)
		f.Log.Debug("创建文件", "file", d.FilePath, "result", b)
	}
}
