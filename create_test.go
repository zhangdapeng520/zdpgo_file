package zdpgo_file

import (
	"testing"
)

/*
@Time : 2022/5/12 14:55
@Author : 张大鹏
@File : create_test.go
@Software: Goland2021.3.1
@Description: create创建类型的相关测试
*/

func TestFile_CreateMultiDir(t *testing.T) {
	f := getFile()
	data := []struct {
		DirPath string
		Status  bool
	}{
		{"test", true},
		{"test1", true},
	}

	for _, d := range data {
		b := f.CreateMultiDir(d.DirPath)
		if b != d.Status {
			panic("测试失败：不是预期的结果")
		}
	}
}

// 测试在指定文件夹中创建指定文件
func TestFile_CreateDirFile(t *testing.T) {
	f := getFile()
	data := []struct {
		DirPath  string
		FileName string
		Content  string
		Status   bool
	}{
		{"test", "a.txt", "aaa", true},
		{"test", "b.txt", "bbb", true},
		{"test", "c", "ccc", true},
		{"test111", "a.txt", "aaa", false},
		{"test111", "b.txt", "bbb", false},
		{"test111", "c", "ccc", false},
	}

	for _, d := range data {
		b := f.CreateDirFile(d.DirPath, d.FileName, d.Content)
		if b != d.Status {
			panic("测试失败：不是预期的结果")
		}
	}
}

// 测试创建文件
func TestFile_CreateFile(t *testing.T) {
	f := getFile()
	data := []struct {
		FilePath string
		Content  string
		Status   bool
	}{
		{"test/aaa.txt", "aaa", true},
		{"test/bbb.txt", "aaa", true},
		{"test/ccc.txt", "aaa", true},
		{"test1111/ccc.txt", "aaa", false},
		{"test1111/acc.txt", "aaa", false},
		{"test1111/bcc.txt", "aaa", false},
	}

	for _, d := range data {
		b := f.CreateFile(d.FilePath, d.Content)
		if b != d.Status {
			panic("测试失败：不是预期的结果")
		}
	}
}
