package zdpgo_file

import (
	"fmt"
	"testing"
)

/*
@Time : 2022/5/12 14:55
@Author : 张大鹏
@File : create_test.go
@Software: Goland2021.3.1
@Description: create创建类型的相关测试
*/

func TestFile_CreateMultiDir(t *testing.T) {
	f := getFile()
	fmt.Println(f.CreateMultiDir("logs/zdpgo"))
}
