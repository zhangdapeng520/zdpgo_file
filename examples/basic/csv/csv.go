package csv

import (
	"github.com/zhangdapeng520/zdpgo_file"
)

/*
@Time : 2022/5/21 10:27
@Author : 张大鹏
@File : csv.go
@Software: Goland2021.3.1
@Description: csv文件相关
*/

func Write(f *zdpgo_file.File) {
	data := [][]string{
		{"a[test].txt[test]", "b[test].txt", "c[test]"},
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

	// 写入
	f.WriteCsv("test/test.csv", data)
	if !f.IsExists("test/test.csv") {
		f.Log.Panic("写入CSV数据失败")
	}

	// 读取
	dataNew, err := f.ReadCsv("test/test.csv")
	if err != nil {
		f.Log.Error("读取CSV文件失败", "error", err, "file", "test.csv")
	}
	f.Log.Debug("读取CSV文件成功", "data", dataNew)
}
