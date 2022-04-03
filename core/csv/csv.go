package csv

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

type Csv struct {
}

func New() *Csv {
	c := Csv{}
	return &c
}

// Write 将数据保存为csv文件
func (f *Csv) Write(filePath string, data [][]string) (err error) {
	fileHandler, err := os.Create(filePath)
	defer fileHandler.Close()
	if err != nil {
		return err
	}

	// 写入UTF-8 BOM
	_, err = fileHandler.WriteString("\xEF\xBB\xBF")
	if err != nil {
		return err
	}

	// 创建一个新的写入文件流
	w := csv.NewWriter(fileHandler)

	// 使其同步，保证线程安全
	rw.Lock()
	err = w.WriteAll(data)
	if err != nil {
		return err
	}
	w.Flush()
	rw.Unlock()

	return nil
}

// Read 读取csv
func (f *Csv) Read(fileName string) (data [][]string, err error) {
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
			return nil, err
		}

		// 追加数据
		data = append(data, line)
	}

	// 返回数据
	return data, nil
}
