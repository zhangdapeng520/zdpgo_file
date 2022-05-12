package zdpgo_file

import "testing"

/*
@Time : 2022/5/12 16:09
@Author : 张大鹏
@File : delete_test.go
@Software: Goland2021.3.1
@Description: delete删除相关的测试
*/

func TestFile_DeleteDirFile(t *testing.T) {
	f := getFile()
	data := []struct {
		DirPath  string
		FileName string
		Status   bool
	}{
		{"test", "a", true},
		{"test", "aa", false},
		{"test", "b", true},
		{"test", "test11.txt", true},
		{"test", "test22.txt", true},
		{"test", "test33.txt", true},
	}

	for _, d := range data {
		b := f.DeleteDirFile(d.DirPath, d.FileName)
		if b != d.Status {
			panic("测试失败：不是预期的结果")
		}
	}
}
