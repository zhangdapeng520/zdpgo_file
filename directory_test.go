package zdpgo_file

import (
	"fmt"
	"testing"
)

func TestFile_IsExist(t *testing.T) {
	f := getFile()
	fmt.Println(f.IsExist("D:\\BaiduNetdiskWorkspace\\文档"))
}

func TestFile_CreateMultiDir(t *testing.T) {
	f := getFile()
	fmt.Println(f.CreateMultiDir("logs/zdpgo"))
}
