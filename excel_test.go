package zdpgo_file

import (
	"fmt"
	"testing"
)

func TestFile_NewExcel(t *testing.T) {
	f := prepareFile()
	f.ExcelNew(FileExcelConfig{FileName: "test.xlsx"})
	f.ExcelNewSheet("测试")
	f.ExcelSetCellValue("测试", "A1", "姓名")
	f.ExcelSetCellValue("测试", "B1", "性别")
	f.ExcelSetCellValue("测试", "C1", "年龄")
	f.ExcelSetCellValue("测试", "A2", "张三")
	f.ExcelSetCellValue("测试", "B2", "男")
	f.ExcelSetCellValue("测试", "C2", "22")
	f.ExcelSave()
}

func TestFile_GetCellValue(t *testing.T) {
	f := prepareFile()
	f.ExcelNew(FileExcelConfig{FileName: "test.xlsx"})
	value, err := f.GetCellValue("测试", "A1")
	fmt.Println(value, err)
}

func TestFile_GetSheetData(t *testing.T) {
	f := prepareFile()
	f.ExcelNew(FileExcelConfig{FileName: "test.xlsx"})
	data, err := f.GetSheetData("测试")
	fmt.Println(data, err)

	for _, row := range data { // 读取一行
		for _, col := range row { // 读取一行中的每列
			fmt.Print(col, "\t")
		}
		fmt.Println()
	}
}

func TestFile_ExcelAddChart(t *testing.T) {
	f := prepareFile()
	f.ExcelNew(FileExcelConfig{FileName: "test.xlsx"})

	// 分类
	categories := map[string]string{
		"A2": "Small", "A3": "Normal", "A4": "Large",
		"B1": "Apple", "C1": "Orange", "D1": "Pear"}

	// 值
	values := map[string]int{
		"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}

	// 添加分类
	for k, v := range categories {
		f.ExcelSetCellValue("Sheet1", k, v)
	}

	// 添加值
	for k, v := range values {
		f.ExcelSetCellValue("Sheet1", k, v)
	}

	// 添加图表
	if err := f.ExcelAddChart("Sheet1", "E1", `{
        "type": "col3DClustered",
        "series": [
        {
            "name": "Sheet1!$A$2",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$2:$D$2"
        },
        {
            "name": "Sheet1!$A$3",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$3:$D$3"
        },
        {
            "name": "Sheet1!$A$4",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$4:$D$4"
        }],
        "title":
        {
            "name": "Fruit 3D Clustered Column Chart"
        }
    }`); err != nil {
		fmt.Println(err)
		return
	}
	f.ExcelSave()
}

func TestFile_ExcelAddPicture(t *testing.T) {
	f := prepareFile()
	f.ExcelNew(FileExcelConfig{FileName: "test.xlsx"})

	defer func() {
		if err := f.ExcelClose(); err != nil {
			fmt.Println(err)
		}
	}()

	img := "D:\\data\\image\\wzry\\1.jpg"

	// 插入一张图片
	if err := f.ExcelAddPicture("Sheet1", "A2", img, ""); err != nil {
		fmt.Println(err)
	}

	// 插入一张图片并缩放
	if err := f.ExcelAddPicture("Sheet1", "D2", img,
		`{"x_scale": 0.5, "y_scale": 0.5}`); err != nil {
		fmt.Println(err)
	}

	// 插入一张图片并支持打印
	if err := f.ExcelAddPicture("Sheet1", "H2", img, `{
        "x_offset": 15,
        "y_offset": 10,
        "print_obj": true,
        "lock_aspect_ratio": false,
        "locked": false
    }`); err != nil {
		fmt.Println(err)
	}

	f.ExcelSave()
}
