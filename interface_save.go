package zdpgo_file

// 保存CSV文件
type ICsv interface{
	// 保存csv的方法
	SaveCsv(fileName string)
}