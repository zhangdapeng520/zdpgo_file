package zdpgo_file

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
)

// File 文件对象
type File struct {
	sync.Mutex                // 同步锁
	Log        *zdpgo_log.Log // 日志
	Config     *Config        // 配置对象
}

// New 新建文件对象
func New() *File {
	return NewWithConfig(Config{})
}

// NewWithConfig 使用配置创建File对象
func NewWithConfig(config Config) *File {
	f := File{}

	// 日志
	if config.LogFilePath == "" {
		config.LogFilePath = "logs/zdpgo/zdpgo_file.log"
	}
	f.Log = zdpgo_log.NewWithDebug(config.Debug, config.LogFilePath)

	// 配置
	f.Config = &config

	// 返回对象
	return &f
}

// Size 获取文件大小
func (f *File) Size(path string) int64 {
	if !f.IsExists(path) {
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
		f.Log.Error("读取文件夹失败", "error", err, "dirPath", dirPath)
		return err
	}

	// 遍历修改文件
	for _, file := range files {
		// 获取文件名
		fileName := file.Name()
		newName := f.GetFileName(fileName)

		// 拼接文件夹
		fileName = path.Join(dirPath, fileName)

		// 如果是文件夹，递归移除
		if f.IsDir(fileName) {
			err = f.RemoveDirFilesSuffix(fileName)
			if err != nil {
				f.Log.Error("递归移除文件夹后缀失败", "error", err, "fileName", fileName)
				return err
			}
			continue
		}

		if f.GetFileSuffix(fileName) == "" {
			f.Log.Debug("该文件不包含后缀，跳过", "fileName", fileName)
			continue
		}

		// 不是文件夹，则移除后缀
		newName = path.Join(dirPath, newName)

		// 重命名
		err = os.Rename(fileName, newName)
		if err != nil {
			f.Log.Error("重命名文件失败", "error", err, "fileName", fileName, "newName", newName)
			return err
		}

		f.Log.Debug("移除文件后缀成功", "fileName", fileName, "newName", newName)
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
