package zdpgo_file

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"sync"
)

var (
	rw = new(sync.RWMutex) // 读写互斥锁
)

// CsvSave 将数据保存为csv文件
func (f *File) CsvSave(filePath string, data [][]string) bool {
	fileHandler, err := os.Create(filePath)
	if err != nil {
		f.log.Error("创建csv文件失败", "error", err)
		return false
	}
	defer func(fileHandler *os.File) {
		err := fileHandler.Close()
		if err != nil {
			f.log.Error("关闭文件失败", "error", err.Error())
		}
	}(fileHandler)

	// 写入UTF-8 BOM
	_, err = fileHandler.WriteString("\xEF\xBB\xBF")
	if err != nil {
		f.log.Error("写入UTF-8失败", "error", err.Error())
		return false
	}

	// 创建一个新的写入文件流
	w := csv.NewWriter(fileHandler)

	// 使其同步，保证线程安全
	rw.Lock()
	err = w.WriteAll(data)
	if err != nil {
		f.log.Error("写入所有数据失败", "error", err.Error())
		return false
	}
	w.Flush()
	rw.Unlock()

	return true
}

// CsvRead 读取csv
func (f *File) CsvRead(fileName string) (data [][]string, err error) {
	// 打开文件
	csvFile, _ := os.Open(fileName)

	// 读取文件
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		// 读取一行
		var line []string
		line, err = reader.Read()

		// 处理错误
		if err == io.EOF {
			break
		} else if err != nil {
			f.log.Error("读取csv失败", "error", err.Error())
			return nil, err
		}

		// 追加数据
		data = append(data, line)
	}

	// 返回数据
	return data, nil
}
