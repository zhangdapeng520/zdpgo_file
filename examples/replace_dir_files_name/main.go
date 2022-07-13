package main

import "github.com/zhangdapeng520/zdpgo_file"

func main() {
	// 将111改为222
	f := zdpgo_file.New()
	f.ReplaceDirFilesName("examples/replace_dir_files_name/test", "111", "222")
}
