package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_file"
)

func main() {
	f := zdpgo_file.New()
	fmt.Println(f.File.Exists("./file.go"))
	fmt.Println(f.File.Exists("./file111.go"))

	fmt.Println(f.File.Size("./file.go"))
	fmt.Println(f.File.Size("./file1.go"))
}
