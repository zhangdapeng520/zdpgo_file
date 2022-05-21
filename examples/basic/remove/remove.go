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
	b := f.RemoveDirFilesSuffix("test")
	f.Log.Debug("移除文件夹中所有文件的后缀", "dir", "test", "result", b)
}
