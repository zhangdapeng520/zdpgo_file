package zdpgo_file

import (
	"os"
)

// Exit 退出程序
func (f File) Exit() {
	f.log.Info("开始退出...")
	f.log.Info("执行清理...")
	f.log.Info("结束退出...")
	os.Exit(0)
}
