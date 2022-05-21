package zdpgo_file

import (
	"encoding/csv"
	"os"
)

/*
@Time : 2022/5/12 14:56
@Author : 张大鹏
@File : write.go
@Software: Goland2021.3.1
@Description: write类型的写入相关方法
*/

// WriteCsv 将数据保存为csv文件
func (f *File) WriteCsv(filePath string, data [][]string) bool {
	fileHandler, err := os.Create(filePath)
	defer fileHandler.Close()
	if err != nil {
		f.Log.Error("创建CSV文件失败", "error", err)
		return false
	}

	// 写入UTF-8 BOM
	_, err = fileHandler.WriteString("\xEF\xBB\xBF")
	if err != nil {
		f.Log.Error("写入字符集失败", "error", err)
		return false
	}

	// 创建一个新的写入文件流
	w := csv.NewWriter(fileHandler)

	// 使其同步，保证线程安全
	f.Lock()
	err = w.WriteAll(data)
	if err != nil {
		f.Log.Error("写入CSV数据失败", "error", err)
		return false
	}
	w.Flush()
	f.Unlock()

	return true
}
