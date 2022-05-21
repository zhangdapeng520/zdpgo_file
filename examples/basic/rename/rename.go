package rename

import (
	"github.com/zhangdapeng520/zdpgo_file"
)

/*
@Time : 2022/5/21 9:40
@Author : 张大鹏
@File : rename.go
@Software: Goland2021.3.1
@Description: rename 重命名相关
*/

func RenameDirFilesName(f *zdpgo_file.File) {
	b := f.ReplaceDirFilesName("./test", "[test]", "")
	f.Log.Debug("重命名文件夹中所有文件", b)
}
