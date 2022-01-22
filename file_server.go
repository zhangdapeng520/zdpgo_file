package zdpgo_file

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	lck sync.Mutex // 同步锁
)

// 获取文件内容，并写入到文件中
func (f File) RecvFile(conn net.Conn, fileName string) {
	// 按照文件名创建新文件
	f1, err := os.Create(fileName)
	if err != nil {
		f.log.Error("创建文件失败：", err)
		return
	}
	defer f1.Close()

	// 从网络中读数据，写入本地文件
	buf := make([]byte, 1024*33)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			f.log.Info("接收文件完成。")
			return
		}
		// 写入本地文件，读多少，写多少。
		lck.Lock()
		f1.Write(buf[:n])
		lck.Unlock()
	}
}

// CreateServer 创建文件服务器
// @param host 文件服务器主机地址，可以收0.0.0.0
// @param port 文件服务器端口号
func (f File) CreateServer(host string, port uint16) {
	// 创建监听退出chan
	c := make(chan os.Signal)

	// 监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 优雅退出
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				f.log.Info("退出", s)
				f.Exit()
			default:
				f.log.Info("其他信号：", s)
			}
		}
	}()

	// 创建用于监听的socket
	address := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		f.log.Error("创建文件服务监听器失败：", err)
		return
	}
	defer listener.Close()

	// 不断的接收来自客户端的文件并保存
	for {
		// 阻塞监听
		conn, err := listener.Accept()
		if err != nil {
			f.log.Error("listener.Accept()监听数据失败：", err)
			return
		}

		// 开启协程，接收数据并写入文件
		go func(conn net.Conn) {
			defer conn.Close()
			// 获取文件名，保存
			buf := make([]byte, 1024*1024*1) // 1M
			n, err := conn.Read(buf)
			if err != nil {
				f.log.Error("从远程读取连接数据失败：", err)
				return
			}
			fileName := string(buf[:n])

			// 回写 ok 给发送端
			conn.Write([]byte("ok"))

			// 获取文件内容
			f.RecvFile(conn, fileName)
		}(conn)
	}
}
