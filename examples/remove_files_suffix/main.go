package main

import (
	"github.com/zhangdapeng520/zdpgo_file"
)

func main() {
	f := zdpgo_file.New()
	err := f.File.RemoveDirFilesSuffix("test")
	if err != nil {
		panic(err)
	}
}
