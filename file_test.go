package zdpgo_file

import (
	"fmt"
	"testing"
)

func getFile() *File {
	return New()
}

func TestFile_csv(t *testing.T) {
	f := getFile()
	data := [][]string{
		{"a", "b", "c"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
	}

	// 写入
	f.Csv.Write("test.csv", data)

	// 读取
	dataNew, err := f.Csv.Read("test.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dataNew)

}

func TestFile_directory(t *testing.T) {
	f := getFile()
	fmt.Println(f.Directory.GetDirectorySize("../download"))
	fmt.Println(f.Directory.IsExist("D:\\BaiduNetdiskWorkspace\\文档"))
	fmt.Println(f.Directory.CreateMultiDir("logs/zdpgo"))
}

func TestFile_download(t *testing.T) {

	f := getFile()
	savePath := "./"

	// 单个下载
	url := "https://avatar.csdnimg.cn/2/9/0/1_togolife.jpg"
	err := f.Download.Download(savePath, url)
	if err != nil {
		fmt.Println("下载失败：", err)
	} else {
		fmt.Println("下载成功")
	}

	// 批量下载
	urls := []string{
		"https://alifei04.cfp.cn/creative/vcg/nowarter800/new/VCG41N695593548.jpg",
		"https://tenfei02.cfp.cn/creative/vcg/nowarter800/new/VCG41N1014325904.jpg",
		"https://tenfei05.cfp.cn/creative/vcg/nowater800/new/VCG41545444880.jpg",
	}
	f.Download.Downloads(savePath, urls)
}

func TestFile_file(t *testing.T) {
	f := getFile()
	fmt.Println(f.File.Exists("./file.go"))
	fmt.Println(f.File.Exists("./file111.go"))

	fmt.Println(f.File.Size("./file.go"))
	fmt.Println(f.File.Size("./file1.go"))
}
