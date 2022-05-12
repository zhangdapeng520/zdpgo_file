package zdpgo_file

import (
	"fmt"
	"testing"
)

/*
@Time : 2022/5/12 14:58
@Author : 张大鹏
@File : read_test.go
@Software: Goland2021.3.1
@Description: 读取类型的相关测试
*/
func TestFile_ReadCsv(t *testing.T) {
	f := getFile()
	data, err := f.ReadCsv("test.csv")
	fmt.Println(data, err)
}
