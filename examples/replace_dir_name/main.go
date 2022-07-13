package main

import "github.com/zhangdapeng520/zdpgo_file"

func main() {
	// 将111改为222
	f := zdpgo_file.New()
	f.ReplaceDirName("examples/replace_dir_name/test", "222", "100")
}
