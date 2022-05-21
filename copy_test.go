package zdpgo_file

import "testing"

/*
@Time : 2022/5/12 15:38
@Author : 张大鹏
@File : copy_test.go
@Software: Goland2021.3.1
@Description: 复制相关的测试
*/
func TestFile_CopyFile(t *testing.T) {
	f := getFile()
	testData := []struct {
		Src  string
		Dest string
	}{
		{"test/a.txt", "test/test11.txt"},
		{"test/a.txt", "test/test22.txt"},
		{"test/a.txt", "test/test33.txt"},
	}

	for _, data := range testData {
		err := f.CopyFile(data.Src, data.Dest)
		if err != nil {
			panic(err)
		}
	}
}
