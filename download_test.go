package zdpgo_file

import (
	"fmt"
	"testing"
)

func TestFile_DownloadNetImage(t *testing.T) {
	f := prepareFile()
	savePath := "D:\\data\\image\\spider"
	url := "https://avatar.csdnimg.cn/2/9/0/1_togolife.jpg"
	err := f.DownloadNetImage(savePath, url)
	if err != nil {
		fmt.Println("Download pic file failed!", err)
	} else {
		fmt.Println("Download file success.")
	}
}

func TestFile_DownloadNetImageMany(t *testing.T) {
	urls := []string{
		"https://alifei04.cfp.cn/creative/vcg/nowarter800/new/VCG41N695593548.jpg",
		"https://tenfei02.cfp.cn/creative/vcg/nowarter800/new/VCG41N1014325904.jpg",
		"https://tenfei05.cfp.cn/creative/vcg/nowater800/new/VCG41545444880.jpg",
	}
	f := prepareFile()
	savePath := "D:\\data\\image\\spider"
	f.DownloadNetImageMany(savePath, urls)
}
