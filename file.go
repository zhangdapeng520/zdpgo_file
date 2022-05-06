package zdpgo_file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
)

// File 文件对象
type File struct {
	lock sync.Mutex
}

// New 新建文件对象
func New() *File {

	f := File{}
	return &f
}

// Exists 判断文件是否存在
func (f *File) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// Size 获取文件大小
func (f *File) Size(path string) int64 {
	if !f.Exists(path) {
		return 0
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

func (f *File) Rename() {
	dir := "static"
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		fileName := file.Name()
		newName := strings.Split(fileName, ".")[0]

		fileName = path.Join(dir, fileName)
		newName = path.Join(dir, newName)

		fmt.Println("正在修改文件：", fileName, newName)
		err := os.Rename(fileName, newName)
		if err != nil {
			panic(err)
		}
	}
}

// RemoveDirFilesSuffix 去掉指定目录下所有文件的后缀
func (f *File) RemoveDirFilesSuffix(dirPath string) error {
	// 读取文件夹
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	// 遍历修改文件
	for _, file := range files {
		// 获取文件名
		fileName := file.Name()
		newName := strings.Split(fileName, ".")[0]

		// 拼接文件夹
		fileName = path.Join(dirPath, fileName)
		newName = path.Join(dirPath, newName)

		// 重命名
		err := os.Rename(fileName, newName)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReplaceDirFilesName 替换指定目录下的文件名
func (f *File) ReplaceDirFilesName(dirPath string, oldStr, newStr string) error {
	// 读取文件夹
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	// 遍历修改文件
	for _, file := range files {
		// 获取文件名
		fileName := file.Name()
		newName := strings.Replace(fileName, oldStr, newStr, 1)

		// 拼接文件夹
		fileName = path.Join(dirPath, fileName)
		newName = path.Join(dirPath, newName)

		// 重命名
		err = os.Rename(fileName, newName)
		if err != nil {
			return err
		}
	}
	return nil
}

// Copy 复制文件
// @param src 源文件地址
// @param dest 目标文件地址
func (f *File) Copy(src, dest string) error {
	// 读取文件夹
	srcFile, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dest, srcFile, 0644)
	if err != nil {
		return err
	}

	return nil
}
