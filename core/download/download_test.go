package download

import (
	"fmt"
	"testing"
)

func getDownload() *Download {
	return New()
}

func TestFile_DownloadNetImage(t *testing.T) {
	f := getDownload()
	savePath := "./"
	url := "https://avatar.csdnimg.cn/2/9/0/1_togolife.jpg"
	err := f.Download(savePath, url)
	if err != nil {
		fmt.Println("下载失败：", err)
	} else {
		fmt.Println("下载成功：")
	}
}

func TestFile_DownloadNetImageMany(t *testing.T) {
	urls := []string{
		"https://alifei04.cfp.cn/creative/vcg/nowarter800/new/VCG41N695593548.jpg",
		"https://tenfei02.cfp.cn/creative/vcg/nowarter800/new/VCG41N1014325904.jpg",
		"https://tenfei05.cfp.cn/creative/vcg/nowater800/new/VCG41545444880.jpg",
	}
	f := getDownload()
	savePath := "./"
	f.Downloads(savePath, urls)
}
