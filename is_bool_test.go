package zdpgo_file

import (
	"fmt"
	"testing"
)

/*
@Time : 2022/5/12 14:54
@Author : 张大鹏
@File : is_bool_test.go
@Software: Goland2021.3.1
@Description: is类型的测试
*/

func TestFile_IsExist(t *testing.T) {
	f := getFile()
	fmt.Println(f.IsExists("D:\\BaiduNetdiskWorkspace\\文档"))
}

// 判断文件夹中是否包含某个文件
func TestFile_IsDirContainsFile(t *testing.T) {
	f := getFile()
	data := []struct {
		DirPath  string
		FileName string
		Status   bool
	}{
		{"test", "a", true},
		{"test", "aa", false},
		{"test", "test11.txt", true},
		{"test", "test111.txt", false},
		{"test", "test333.txt", false},
		{"test", "test33.txt", true},
	}

	for _, d := range data {
		b := f.IsDirContainsFile(d.DirPath, d.FileName)
		if b != d.Status {
			panic("测试失败：不是预期的结果")
		}
	}
}
