package zdpgo_file

import (
	"github.com/zhangdapeng520/zdpgo_file/core/csv"
	"github.com/zhangdapeng520/zdpgo_file/core/directory"
	"github.com/zhangdapeng520/zdpgo_file/core/download"
	"github.com/zhangdapeng520/zdpgo_file/core/file"
)

// File 文件对象
type File struct {
	Csv       *csv.Csv             // CSV处理对象
	File      *file.File           // File文件处理对象
	Directory *directory.Directory // 文件夹处理对象
	Download  *download.Download   // 下载对象
}

// New 新建文件对象
func New() *File {

	f := File{}

	// 初始化操作对象
	f.Csv = csv.New()
	f.Directory = directory.New()
	f.File = file.New()
	f.Download = download.New()

	return &f
}
