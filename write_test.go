package zdpgo_file

import "testing"

/*
@Time : 2022/5/12 15:34
@Author : 张大鹏
@File : write_test.go
@Software: Goland2021.3.1
@Description: write写入相关的测试
*/

func TestFile_WriteCsv(t *testing.T) {
	f := getFile()
	data := [][]string{
		{"a[test].txt", "b[test].txt", "c"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
	}
	f.WriteCsv("test.csv", data)
}
