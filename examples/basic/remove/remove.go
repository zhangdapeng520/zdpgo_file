package remove

import "github.com/zhangdapeng520/zdpgo_file"

/*
@Time : 2022/5/21 9:21
@Author : 张大鹏
@File : remove.go
@Software: Goland2021.3.1
@Description: remove移除相关
*/

func RemoveSuffix(f *zdpgo_file.File) {
	err := f.RemoveDirFilesSuffix("test")
	if err != nil {
		panic(err)
	}
}
