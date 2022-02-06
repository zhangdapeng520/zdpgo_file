package zdpgo_file

import (
	"fmt"
	"io"
	"net"
	"os"
)

// SendFile 发送文件到服务器
func (f File) sendFile(conn net.Conn, filePath string) {
	// 只读打开文件
	f1, err := os.Open(filePath)
	if err != nil {
		f.log.Error("打开文件失败：", err)
		return
	}
	defer f1.Close()

	// 从本文件中，读数据，写给网络接收端。 读多少，写多少。原封不动。
	buf := make([]byte, 1024*1024*1)
	for {
		n, err := f1.Read(buf)
		if err != nil {
			if err == io.EOF {
				f.log.Info("发送文件完成。")
			} else {
				f.log.Error("打开文件失败：", err)
			}
			return
		}

		// 写到网络socket中
		_, err = conn.Write(buf[:n])
		if err != nil {
			f.log.Error("写数据到conn连接中失败：", err)
			return
		}
		f.log.Info("向服务器传输文件成功!")
	}
}

// SendFile 向文件服务器发送文件
// @param host 文件服务器主机地址
// @param port 文件服务器端口号
func (f File) SendFile(host string, port uint16, path string, name string) {
	// 主动发起连接请求
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		f.log.Error("建立远程连接失败：", err)
		return
	}
	defer conn.Close()

	// 发送文件名给 接收端
	_, err = conn.Write([]byte(name))
	if err != nil {
		f.log.Error("写入文件到conn失败：", err)
		return
	}

	// 读取服务器回发的 OK
	buf := make([]byte, 16)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("从conn读取数据失败：", err)
		return
	}

	if string(buf[:n]) == "ok" {
		// 写文件内容给服务器——借助conn
		f.sendFile(conn, path)
	}

}
