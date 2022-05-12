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
	logConfig := zdpgo_log.Config{
		Debug:       config.Debug,
		OpenJsonLog: true,
		LogFilePath: config.LogFilePath,
	}
	if config.Debug {
		logConfig.IsShowConsole = true
	}
	f.Log = zdpgo_log.NewWithConfig(logConfig)

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
