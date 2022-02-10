package zdpgo_file

import (
	"fmt"
	"testing"
)

func TestDirectorySize(t *testing.T) {
	f := prepareFile()
	fmt.Println(f.GetDirectorySize("D:\\BaiduNetdiskWorkspace\\文档"))
}

func TestFile_IsExist(t *testing.T) {
	f := prepareFile()
	fmt.Println(f.IsExist("D:\\BaiduNetdiskWorkspace\\文档"))
}

func TestFile_CreateMultiDir(t *testing.T) {
	f := prepareFile()
	fmt.Println(f.CreateMultiDir("logs/zdpgo"))
}
