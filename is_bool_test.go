package zdpgo_file

import (
	"fmt"
	"testing"
)

/*
@Time : 2022/5/12 14:54
@Author : 张大鹏
@File : is_bool_test.go
@Software: Goland2021.3.1
@Description: is类型的测试
*/

func TestFile_IsExist(t *testing.T) {
	f := getFile()
	fmt.Println(f.IsExist("D:\\BaiduNetdiskWorkspace\\文档"))
}
