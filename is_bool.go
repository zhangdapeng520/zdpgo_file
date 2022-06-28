package zdpgo_file

import (
	"io/ioutil"
	"os"
)

/*
@Time : 2022/5/12 14:53
@Author : 张大鹏
@File : is_bool.go
@Software: Goland2021.3.1
@Description: is类型的判断方法
*/

// IsExists 判断所给路径文件/文件夹是否存在(返回true是存在)
func (f *File) IsExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDirContainsFile 判断文件夹中是否包含指定文件
func (f *File) IsDirContainsFile(dirPath, fileName string) bool {
	dirList, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return false
	}

	// 遍历文件夹，判断是否包含指定文件
	for _, dirFile := range dirList {
		if dirFile.Name() == fileName {
			return true
		}
	}

	return false
}

// IsDir 判断所给路径是否为文件夹
func (f *File) IsDir(dirPath string) bool {
	s, err := os.Stat(dirPath)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func (f *File) IsFile(path string) bool {
	return !f.IsDir(path)
}
