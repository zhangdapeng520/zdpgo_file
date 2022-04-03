package directory

import (
	"fmt"
	"testing"
)

func getDirectory() *Directory {
	return New()
}

func TestDirectorySize(t *testing.T) {
	f := getDirectory()
	fmt.Println(f.GetDirectorySize("../download"))
}

func TestFile_IsExist(t *testing.T) {
	f := getDirectory()
	fmt.Println(f.IsExist("D:\\BaiduNetdiskWorkspace\\文档"))
}

func TestFile_CreateMultiDir(t *testing.T) {
	f := getDirectory()
	fmt.Println(f.CreateMultiDir("logs/zdpgo"))
}
