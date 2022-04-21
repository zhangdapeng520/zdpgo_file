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

func TestFile_RemoveDirFilesSuffix(t *testing.T) {
	f := getFile()
	err := f.RemoveDirFilesSuffix("./test")
	if err != nil {
		t.Error(err)
	}
}
