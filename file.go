package zdpgo_file

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/zhangdapeng520/zdpgo_log"
)

// File 文件对象
type File struct {
	Path    string         // 文件路径
	Name    string         // 文件名称
	Suffix  string         // 文件后缀
	Size    int64          // 文件大小
	ModTime time.Time      // 修改时间
	IsDir   bool           // 是否为文件夹
	log     *zdpgo_log.Log // 日志
	config  *FileConfig    // 配置信息
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
	if config.Path == "" {
		f.log.Warning("文件路径为空")
	} else {
		f.Path = config.Path

		// 判断文件是否存在
		isExist := err == nil || os.IsExist(err)
		if !isExist {
			f.log.Panic("文件不存在：", config.Path)
		}

		// 获取文件相关数据
		f.Size = fileInfo.Size()       // 大小
		f.Name = fileInfo.Name()       // 名称
		f.ModTime = fileInfo.ModTime() // 修改时间
		f.IsDir = fileInfo.IsDir()     // 是否为文件夹

		// 获取后缀
		arr := strings.Split(config.Path, ".")
		if len(arr) <= 1 {
			if f.IsDir {
				f.log.Info("该路径为一个文件夹")
			} else {
				f.log.Warning("该文件不包含后缀")
			}
		} else {
			f.Suffix = fmt.Sprintf(".%s", arr[len(arr)-1])
		}
	}

	return &f
}
