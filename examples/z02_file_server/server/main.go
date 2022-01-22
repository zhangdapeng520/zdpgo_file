package main

import (
	"github.com/zhangdapeng520/zdpgo_file"
)

func main() {
	fileConfig := zdpgo_file.FileConfig{
		Debug: true,
	}
	f := zdpgo_file.New(fileConfig)
	f.CreateServer("0.0.0.0", 8888)
}
