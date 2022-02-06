package zdpgo_file

import (
	"github.com/xuri/excelize/v2"
	"github.com/zhangdapeng520/zdpgo_zap"
)

// File 文件对象
type File struct {
	log         *zdpgo_zap.Zap   // 日志
	config      *FileConfig      // 配置信息
	excelConfig *FileExcelConfig // Excel配置信息
	excel       *excelize.File   // Excel文件对象
}

// FileConfig 文件配置
type FileConfig struct {
	Debug       bool   // 是否为debug模式
	LogFilePath string // 日志存放路径
}

// NewFile 新建文件对象
// @param config 配置对象
func NewFile(config FileConfig) *File {

	f := File{}

	// 初始化日志
	if config.LogFilePath == "" {
		config.LogFilePath = "zdpgo_file.log"
	}
	l := zdpgo_zap.New(zdpgo_zap.ZapConfig{
		Debug:        config.Debug,
		OpenGlobal:   false,
		OpenFileName: false,
		LogFilePath:  config.LogFilePath,
	})
	f.log = l

	// 初始化配置
	f.config = &config

	return &f
}
