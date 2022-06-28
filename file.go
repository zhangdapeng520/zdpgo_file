package zdpgo_file

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
)

// File 文件对象
type File struct {
	sync.Mutex         // 同步锁
	Config     *Config // 配置对象
}

// New 新建文件对象
func New() *File {
	return NewWithConfig(Config{})
}

// NewWithConfig 使用配置创建File对象
func NewWithConfig(config Config) *File {
	f := File{}

	// 配置
	f.Config = &config

	// 返回对象
	return &f
}

// GetFileSize 获取文件大小
func (f *File) GetFileSize(path string) int64 {
	if !f.IsExists(path) {
		return 0
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

// RemoveDirFilesSuffix 去掉指定目录下所有文件的后缀
func (f *File) RemoveDirFilesSuffix(dirPath string) bool {
	// 读取文件夹
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return false
	}

	// 遍历修改文件
	for _, file := range files {
		// 获取文件名
		fileName := file.Name()
		newName := f.GetFileNameNoSuffix(fileName)

		// 拼接文件夹
		fileName = path.Join(dirPath, fileName)

		// 如果是文件夹，递归移除
		if f.IsDir(fileName) {
			f.RemoveDirFilesSuffix(fileName)
			continue
		}

		if f.GetFileSuffix(fileName) == "" {
			continue
		}

		// 不是文件夹，则移除后缀
		newName = path.Join(dirPath, newName)

		// 重命名
		err = os.Rename(fileName, newName)
		if err != nil {
			return false
		}
	}
	return true
}

// ReplaceDirFilesName 替换指定目录下的文件名
func (f *File) ReplaceDirFilesName(dirPath string, oldStr, newStr string) bool {
	// 读取文件夹
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return false
	}

	// 遍历修改文件
	for _, file := range files {
		// 获取文件名
		fileName := file.Name()
		newName := strings.Replace(fileName, oldStr, newStr, 1)

		// 拼接文件夹
		fileName = path.Join(dirPath, fileName)

		// 如果是文件夹，递归的重命名
		if f.IsDir(fileName) {
			f.ReplaceDirFilesName(fileName, oldStr, newStr)
			continue
		}

		// 不包含要替换的字符串，跳过
		if !strings.Contains(fileName, oldStr) {
			continue
		}

		// 重命名
		newName = path.Join(dirPath, newName)
		err = os.Rename(fileName, newName)
		if err != nil {
			return false
		}
	}
	return true
}
