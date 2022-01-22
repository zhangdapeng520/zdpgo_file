package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_file"
)

func main() {
	fileConfig := zdpgo_file.FileConfig{
		Path:  "E:\\static\\images\\4k\\4k动漫壁纸\\[大]3D美女 休闲运动装 盘腿坐姿 4k壁纸.jpg",
		Debug: true, // 控制台输出日志
	}
	f := zdpgo_file.New(fileConfig)
	fmt.Println("客户端：", f)

	// 发送文件到服务器
	f.SendFile("127.0.0.1", 8888)
}
