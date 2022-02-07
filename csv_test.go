package zdpgo_file

import (
	"fmt"
	"testing"
)

func TestFile_CsvSave(t *testing.T) {
	f := prepareFile()
	data := [][]string{
		{"a", "b", "c"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
	}
	f.CsvSave("test.csv", data)
}

func TestFile_CsvRead(t *testing.T) {
	f := prepareFile()
	data, err := f.CsvRead("test.csv")
	fmt.Println(data, err)
}
