package zdpgo_file

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)
var (
	rw = new(sync.RWMutex) // 读写互斥锁
)

// 将数据保存为csv文件
func SaveCsv(filePath string, data [][]string) bool{
	fileHandler, err := os.Create(filePath)
	if err != nil{
		fmt.Println("创建csv文件失败：", err)
		return false
	}
	defer fileHandler.Close()

	// 写入UTF-8 BOM
	fileHandler.WriteString("\xEF\xBB\xBF")

	// 创建一个新的写入文件流
	w := csv.NewWriter(fileHandler)

	// 使其同步，保证线程安全
	rw.Lock()
	w.WriteAll(data)
	w.Flush()
	rw.Unlock()
	
	return true
}


// 将数据保存为csv文件
func SaveCsvAsync(filePath string, data [][]string, ch chan bool){

	fileHandler, err := os.Create(filePath)
	if err != nil{
		fmt.Println("创建csv文件失败：", err)
		ch <- false
		return
	}
	defer fileHandler.Close()

	// 写入UTF-8 BOM
	fileHandler.WriteString("\xEF\xBB\xBF")

	// 创建一个新的写入文件流
	w := csv.NewWriter(fileHandler)

	// 使其同步，保证线程安全
	rw.Lock()
	w.WriteAll(data)
	w.Flush()
	rw.Unlock()

	ch <- true
}