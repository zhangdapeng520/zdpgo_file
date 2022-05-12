package zdpgo_file

/*
@Time : 2022/5/12 14:42
@Author : 张大鹏
@File : config.go
@Software: Goland2021.3.1
@Description: 文件组件相关的配置
*/

type Config struct {
	Debug       bool   `yaml:"debug" json:"debug"`                 // 是否为debug模式
	LogFilePath string `yaml:"log_file_path" json:"log_file_path"` // 日志路径
}
