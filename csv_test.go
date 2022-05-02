package zdpgo_file

import (
	"fmt"
	"testing"
)

func TestFile_CsvSave(t *testing.T) {
	f := getFile()
	data := [][]string{
		{"a[test]", "b", "c[test]"},
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
	f.WriteCsv("test.csv", data)
}

func TestFile_CsvRead(t *testing.T) {
	f := getFile()
	data, err := f.ReadCsv("test.csv")
	fmt.Println(data, err)
}
