package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_file"
)

func main() {
	f := zdpgo_file.New()
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
