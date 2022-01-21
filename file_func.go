package zdpgo_file

import (
	"os"
)

// FileExists 判断文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// FileSize 获取文件大小
func FileSize(path string) int64 {
	if !FileExists(path) {
		return 0
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

