package zdpgo_file

import (
	"fmt"
	"testing"
)

func TestFileExists(t *testing.T) {
	fmt.Println(FileExists("./file.go"))
	fmt.Println(FileExists("./file111.go"))
}

func TestFileSize(t *testing.T) {
	fmt.Println(FileSize("./file.go"))
	fmt.Println(FileSize("./file1.go"))
}
