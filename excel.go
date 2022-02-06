package zdpgo_file

import (
	"github.com/xuri/excelize/v2"
)

// 参考文档：https://github.com/qax-os/excelize

// NewExcel 创建Excel文件
func (f *File) NewExcel(config FileExcelConfig) {
	f.excel = excelize.NewFile()
	f.excelConfig = &config
}

// NewSheet 创建工作簿
func (f *File) NewSheet(name string) int {
	index := f.excel.NewSheet(name)
	return index
}

// SetCellValue 设置单元格的值
func (f *File) SetCellValue(sheet string, cell string, value interface{}) error {
	err := f.excel.SetCellValue(sheet, cell, value)
	if err != nil {
		return err
	}
	return nil
}

// SetActiveSheet 激活单元格
func (f *File) SetActiveSheet(sheetIndex int) error {
	err := f.SetActiveSheet(sheetIndex)
	if err != nil {
		return err
	}
	return nil
}

// SaveExcel 保存Excel
func (f *File) SaveExcel() error {
	if f.excelConfig.FileName == "" {
		f.log.Warning("Excel文件名为空，将使用默认的文件名：zdpgo_file.xlsx")
		f.excelConfig.FileName = "zdpgo_file.xlsx"
	}
	err := f.excel.SaveAs(f.excelConfig.FileName)
	if err != nil {
		f.log.Error("保存Excel失败", "error", err.Error())
		return err
	}
	return nil
}

// GetCellValue 读取单元格数据
func (f *File) GetCellValue(sheetName string, cellName string) (value string, err error) {
	if f.excelConfig.FileName == "" {
		f.log.Warning("Excel文件名为空，将使用默认的文件名：zdpgo_file.xlsx")
		f.excelConfig.FileName = "zdpgo_file.xlsx"
	}

	// 打开Excel
	f.excel, err = excelize.OpenFile(f.excelConfig.FileName)
	if err != nil {
		f.log.Error("打开Excel失败", "error", err.Error())
		return "", err
	}

	// 读取单元格
	value, err = f.excel.GetCellValue(sheetName, cellName)
	if err != nil {
		f.log.Error("读取Excel单元格失败", "error", err.Error())
		return "", err
	}

	// 正常情况
	return value, nil
}

// GetSheetData 读取整张工作簿
func (f *File) GetSheetData(sheetName string) (data [][]string, err error) {
	if f.excelConfig.FileName == "" {
		f.log.Warning("Excel文件名为空，将使用默认的文件名：zdpgo_file.xlsx")
		f.excelConfig.FileName = "zdpgo_file.xlsx"
	}

	// 打开Excel
	f.excel, err = excelize.OpenFile(f.excelConfig.FileName)
	if err != nil {
		f.log.Error("打开Excel失败", "error", err.Error())
		return data, err
	}

	// 读取单元格
	data, err = f.excel.GetRows(sheetName)
	if err != nil {
		f.log.Error("读取Excel单元格失败", "error", err.Error())
	}
	return data, err
}
