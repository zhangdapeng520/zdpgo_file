package zdpgo_file

import "testing"

func TestDoubleArrayString_SaveCsv(t *testing.T) {
	var das = DoubleArrayString{
		Data: [][]string{
			{"姓名", "年龄", "性别"},
			{"张大鹏", "22", "男"},
			{"张大鹏", "22", "男"},
			{"张大鹏", "22", "男"},
			{"张大鹏", "22", "男"},
			{"张大鹏", "22", "男"},
			{"张大鹏", "22", "男"},
			{"张大鹏", "22", "男"},
			{"张大鹏", "22", "男"},
			{"张大鹏", "22", "男"},
		},
	}
	das.SaveCsv("test10.csv")
	das.SaveCsv("test11.csv")
	das.SaveCsv("test12.csv")
	das.SaveCsv("test13.csv")
}