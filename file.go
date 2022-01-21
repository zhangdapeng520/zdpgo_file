package zdpgo_file

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"
	"os"
	"strings"
)

// File 文件对象
type File struct {
	Path   string         // 文件路径
	Name   string         // 文件名称
	Suffix string         // 文件后缀
	Size   int64          // 文件大小
	log    *zdpgo_log.Log // 日志
	config *FileConfig    // 配置信息
}

// FileConfig 文件配置
type FileConfig struct {
	Path        string // 文件路径
	LogFilePath string // 日志存放路径
	Debug       bool   // 是否为debug模式
}

// New 新建文件对象
// @param config 配置对象
func New(config FileConfig) *File {

	f := File{}

	// 初始化日志
	if config.LogFilePath == "" {
		config.LogFilePath = "zdpgo_file.log"
	}
	logConfig := zdpgo_log.LogConfig{
		Debug: config.Debug,
		Path:  config.LogFilePath,
	}
	l := zdpgo_log.New(logConfig)
	f.log = l

	fileInfo, err := os.Stat(config.Path)

	// 边界条件
	if config.Path == ""{
		f.log.Panic("文件路径不能为空")
	}else{
		f.Path = config.Path
	}

	// 判断文件是否存在
	isExist := err == nil || os.IsExist(err)
	if !isExist {
		f.log.Panic("文件不存在：", config.Path)
	}

	// 获取文件名和大小
	f.Size = fileInfo.Size()
	f.Name = fileInfo.Name()

	// 获取后缀
	arr := strings.Split(config.Path, ".")
	if len(arr) <= 1 {
		f.log.Warning("该文件不包含后缀")
	} else {
		f.Suffix = fmt.Sprintf(".%s", arr[len(arr)-1])
	}
	return &f
}
