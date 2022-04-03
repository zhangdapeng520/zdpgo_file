package file

import (
	"fmt"
	"testing"
)

func getFile() *File {
	return New()
}

func TestFileExists(t *testing.T) {
	f := getFile()
	fmt.Println(f.Exists("./file.go"))
	fmt.Println(f.Exists("./file111.go"))
}

func TestFileSize(t *testing.T) {
	f := getFile()
	fmt.Println(f.Size("./file.go"))
	fmt.Println(f.Size("./file1.go"))
}
