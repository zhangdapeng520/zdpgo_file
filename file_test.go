package zdpgo_file

import (
	"fmt"
	"strings"
	"testing"
)

func TestFile_New(t *testing.T) {
	fileConfig := FileConfig{
		Path: "file.go",
	}
	f := New(fileConfig)
	fmt.Println(f)
	fmt.Println(f.Name)
	fmt.Println(f.Size)
	fmt.Println(f.Suffix)
}

func TestDemo(t *testing.T) {
	f := "./file.go"
	arr := strings.Split(f, ".")
	fmt.Println(arr)
	fmt.Println(arr[len(arr)-1])
}
