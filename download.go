package zdpgo_file

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// DownloadNetImage 下载网络图片
func (f *File) DownloadNetImage(saveDir string, imageUrl string) error {
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
	v, err := http.Get(imageUrl)
	if err != nil {
		f.log.Error("获取图片内容失败", "error", err, "imageUrl", imageUrl)
		return err
	}
	defer v.Body.Close()
	content, err := ioutil.ReadAll(v.Body)
	if err != nil {
		f.log.Error("读取图片内容失败", "error", err)
		return err
	}
	f.log.Info("保存图片", "pic", pic)
	err = ioutil.WriteFile(pic, content, 0766)
	if err != nil {
		f.log.Error("保存图片失败", "error", err)
		return err
	}
	return nil
}

// DownloadNetImageMany 批量下载图片
func (f *File) DownloadNetImageMany(saveDir string, imageUrls []string) {
	for _, url := range imageUrls {
		err := f.DownloadNetImage(saveDir, url)
		if err != nil {
			f.log.Error("下载图片失败", "error", err.Error(), "url", url)
		}
		f.log.Info("下载图片成功", "url", url)
		time.Sleep(time.Second)
	}
}
