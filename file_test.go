package zdpgo_file

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

func TestFile_RenameDirFilesName(t *testing.T) {
	f := getFile()
	err := f.ReplaceDirFilesName("./test", "[test]", "")
	if err != nil {
		t.Error(err)
	}
}

func TestFile_Copy(t *testing.T) {
	f := getFile()
	testData := []struct {
		Src  string
		Dest string
	}{
		{"test/test1.txt", "test/test11.txt"},
		{"test/test2.txt", "test/test22.txt"},
		{"test/test3.txt", "test/test33.txt"},
	}

	for _, data := range testData {
		err := f.Copy(data.Src, data.Dest)
		if err != nil {
			panic(err)
		}
	}
}
