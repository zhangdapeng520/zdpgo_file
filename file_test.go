package zdpgo_file

import (
	"fmt"
	"strings"
	"testing"
)

func prepareFile() *File {
	f := NewFile(FileConfig{
		Debug: true,
	})
	return f
}

func TestFile_New(t *testing.T) {
	f := prepareFile()
	fmt.Println(f)
}

func TestDemo(t *testing.T) {
	f := "./file.go"
	arr := strings.Split(f, ".")
	fmt.Println(arr)
	fmt.Println(arr[len(arr)-1])
}
