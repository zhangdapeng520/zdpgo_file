package csv

import (
	"fmt"
	"testing"
)

func getCsv() *Csv {
	return New()
}
func TestFile_CsvSave(t *testing.T) {
	f := getCsv()
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
	f.Write("test.csv", data)
}

func TestFile_CsvRead(t *testing.T) {
	f := getCsv()
	data, err := f.Read("test.csv")
	fmt.Println(data, err)
}
