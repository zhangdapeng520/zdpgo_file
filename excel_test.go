package zdpgo_file

import (
	"fmt"
	"testing"
)

func TestFile_NewExcel(t *testing.T) {
	f := prepareFile()
	f.NewExcel(FileExcelConfig{FileName: "test.xlsx"})
	f.NewSheet("测试")
	f.SetCellValue("测试", "A1", "姓名")
	f.SetCellValue("测试", "B1", "性别")
	f.SetCellValue("测试", "C1", "年龄")
	f.SetCellValue("测试", "A2", "张三")
	f.SetCellValue("测试", "B2", "男")
	f.SetCellValue("测试", "C2", "22")
	f.SaveExcel()
}

func TestFile_GetCellValue(t *testing.T) {
	f := prepareFile()
	f.NewExcel(FileExcelConfig{FileName: "test.xlsx"})
	value, err := f.GetCellValue("测试", "A1")
	fmt.Println(value, err)
}

func TestFile_GetSheetData(t *testing.T) {
	f := prepareFile()
	f.NewExcel(FileExcelConfig{FileName: "test.xlsx"})
	data, err := f.GetSheetData("测试")
	fmt.Println(data, err)

	for _, row := range data { // 读取一行
		for _, col := range row { // 读取一行中的每列
			fmt.Print(col, "\t")
		}
		fmt.Println()
	}
}
