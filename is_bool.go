package zdpgo_file

import "os"

/*
@Time : 2022/5/12 14:53
@Author : 张大鹏
@File : is_bool.go
@Software: Goland2021.3.1
@Description: is类型的判断方法
*/

// IsExist 判断所给路径文件/文件夹是否存在(返回true是存在)
func (f *File) IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
