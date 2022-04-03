package download

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Download struct {
}

func New() *Download {
	d := Download{}

	return &d
}

// Download 下载
func (f *Download) Download(saveDir string, imageUrl string) (err error) {
	// 拼接图片路径
	pic := saveDir
	if !strings.HasSuffix(pic, "/") && !strings.HasSuffix(pic, "\\") {
		pic = pic + "/"
	}
	idx := strings.LastIndex(imageUrl, "/")
	if idx < 0 {
		pic += "/" + imageUrl
	} else {
		pic += imageUrl[idx+1:]
	}

	// 获取文件流
	v, err := http.Get(imageUrl)
	if err != nil {
		return err
	}
	defer v.Body.Close()
	content, err := ioutil.ReadAll(v.Body)
	if err != nil {
		return err
	}

	// 写入文件
	err = ioutil.WriteFile(pic, content, 0766)
	if err != nil {
		return err
	}

	// 正常返回
	return nil
}

// Downloads 批量下载
func (f *Download) Downloads(saveDir string, imageUrls []string) (err error) {
	for _, url := range imageUrls {
		err = f.Download(saveDir, url)
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}
