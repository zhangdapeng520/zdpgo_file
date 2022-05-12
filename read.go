package zdpgo_file

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

/*
@Time : 2022/5/12 14:57
@Author : 张大鹏
@File : read.go
@Software: Goland2021.3.1
@Description: read类型的读取相关方法
*/

// ReadCsv 读取csv
func (f *File) ReadCsv(fileName string) (data [][]string, err error) {
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
