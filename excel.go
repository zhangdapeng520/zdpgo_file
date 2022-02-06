package zdpgo_file

import (
	"github.com/xuri/excelize/v2"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// 参考文档：https://github.com/qax-os/excelize

// ExcelNew 创建Excel文件
func (f *File) ExcelNew(config FileExcelConfig) {
	f.excel = excelize.NewFile()
	f.excelConfig = &config
}

// ExcelNewSheet 创建工作簿
func (f *File) ExcelNewSheet(name string) int {
	index := f.excel.NewSheet(name)
	return index
}

// ExcelSetCellValue 设置单元格的值
func (f *File) ExcelSetCellValue(sheet string, cell string, value interface{}) error {
	err := f.excel.SetCellValue(sheet, cell, value)
	if err != nil {
		return err
	}
	return nil
}

// ExcelSetActiveSheet 激活单元格
func (f *File) ExcelSetActiveSheet(sheetIndex int) {
	f.excel.SetActiveSheet(sheetIndex)
}

// ExcelSave 保存Excel
func (f *File) ExcelSave() error {
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

// ExcelAddChart 添加图表到Excel
func (f *File) ExcelAddChart(sheetName string, cellName string, jsonData string) (err error) {
	if f.excelConfig.FileName == "" {
		f.log.Warning("Excel文件名为空，将使用默认的文件名：zdpgo_file.xlsx")
		f.excelConfig.FileName = "zdpgo_file.xlsx"
	}

	// 打开Excel
	if f.excel == nil {
		f.excel, err = excelize.OpenFile(f.excelConfig.FileName)
		if err != nil {
			f.log.Error("打开Excel失败", "error", err.Error())
			return err
		}
	}

	// 添加图
	err = f.excel.AddChart(sheetName, cellName, jsonData)
	if err != nil {
		f.log.Error("添加图表失败", "error", err.Error())
	}
	return err
}

func (f *File) ExcelAddPicture(sheetName string, cellName string, img string, jsonConfig string) (err error) {
	if f.excelConfig.FileName == "" {
		f.log.Warning("Excel文件名为空，将使用默认的文件名：zdpgo_file.xlsx")
		f.excelConfig.FileName = "zdpgo_file.xlsx"
	}

	// 打开Excel
	if f.excel == nil {
		f.excel, err = excelize.OpenFile(f.excelConfig.FileName)
		if err != nil {
			f.log.Error("打开Excel失败", "error", err.Error())
			return err
		}
	}

	// 添加图
	err = f.excel.AddPicture(sheetName, cellName, img, jsonConfig)
	f.log.Info("添加图片的配置", "jsonConfig", jsonConfig)
	if err != nil {
		f.log.Error("添加图片失败", "error", err.Error())
	}
	return err
}

func (f *File) ExcelClose() (err error) {
	// excel 没有打开
	if f.excel == nil {
		return nil
	}

	// 关闭
	err = f.excel.Close()
	if err != nil {
		f.log.Error("关闭Excel失败", "error", err.Error())
	}
	return err
}
