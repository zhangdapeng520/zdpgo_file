package main

import (
	"basic/copy"
	"basic/create"
	"basic/csv"
	"basic/delete"
	"basic/get"
	"basic/is_bool"
	"basic/remove"
	"basic/rename"
	"github.com/zhangdapeng520/zdpgo_file"
)

func main() {
	f := zdpgo_file.NewWithConfig(zdpgo_file.Config{
		Debug: true,
	})
	create.CreateDirFile(f) // 在指定文件夹中创建文件
	create.CreateFile(f)    // 创建文件
	create.CreateDir(f)     // 创建文件夹

	csv.Write(f) // 写入csv文件

	copy.CopyFile(f) // 复制文件

	is_bool.IsExists(f)          // 判断文件是否存在
	is_bool.IsDirContainsFile(f) // 判断文件夹中是否包含某文件

	get.GetFileSize(f)      // 获取文件大小
	get.GetDirectorySize(f) // 获取文件夹大小

	remove.RemoveSuffix(f) // 移除文件后缀

	rename.RenameDirFilesName(f) // 重命名文件夹中的文件

	delete.DeleteDirFile(f) // 删除文件
	delete.DeleteDir(f)     // 删除文件夹
}
