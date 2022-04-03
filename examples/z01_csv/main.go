package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_file"
)

func main() {
	f := zdpgo_file.New()
	data := [][]string{
		{"a", "b", "c"},
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
	f.Csv.Write("test.csv", data)

	// 读取
	dataNew, err := f.Csv.Read("test.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dataNew)
}
