package main

import (
	"basic/remove"
	"github.com/zhangdapeng520/zdpgo_file"
)

func main() {
	f := zdpgo_file.NewWithConfig(zdpgo_file.Config{
		Debug: true,
	})
	remove.RemoveSuffix(f) // 移除文件后缀
}
