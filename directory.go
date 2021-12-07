package zdpgo_file

import (
	"os"
	"path/filepath"
)


// 文件夹占用磁盘大小
func DirectorySize(path string) (int64, error) {
    var size int64
    err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            size += info.Size()
        }
        return err
    })
    return size, err
}
