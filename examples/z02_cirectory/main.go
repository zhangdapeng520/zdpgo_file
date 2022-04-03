package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_file"
)

func main() {
	f := zdpgo_file.New()
	fmt.Println(f.Directory.GetDirectorySize("../download"))
	fmt.Println(f.Directory.IsExist("D:\\BaiduNetdiskWorkspace\\文档"))
	fmt.Println(f.Directory.CreateMultiDir("logs/zdpgo"))
}
