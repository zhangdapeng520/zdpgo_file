package zdpgo_file

import (
	"fmt"
	"testing"
)

func getFile() *File {
	return NewWithConfig(Config{
		Debug: true,
	})
}

func TestFile_IsExists(t *testing.T) {
	f := getFile()
	fmt.Println(f.IsExists("./file.go"))
	fmt.Println(f.IsExists("./file111.go"))
}

func TestFile_Size(t *testing.T) {
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

func TestFile_RenameDirFilesName(t *testing.T) {
	f := getFile()
	err := f.ReplaceDirFilesName("./test", "[test]", "")
	if err != nil {
		t.Error(err)
	}
}
