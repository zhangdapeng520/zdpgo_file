package zdpgo_file

// 二维字符串切片结构体
type DoubleArrayString struct{
	Data [][]string
}

// 实现保存为CSV的方法
func (das *DoubleArrayString) SaveCsv(filepath string){
	ch := make(chan bool)
	go SaveCsvAsync(filepath, das.Data, ch)
	<- ch
}