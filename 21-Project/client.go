package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	// 要连接的服务器IP地址和端口
	ServerIp   string
	ServerPort int

	// 连接到服务器的TCP连接
	conn net.Conn

	Name string
}

func NewClient(serverIp string, serverPort int) *Client {
	// 连接服务器
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return nil
	}

	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		conn:       conn,
		Name:       conn.LocalAddr().String(),
	}

	return client
}

// 命令行参数绑定
var (
	serverIp   string
	serverPort int
)

func init() {
	// 命令行参数绑定
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "Remote server ip address")
	flag.IntVar(&serverPort, "port", 8080, "Remote server port")
	// 解析命令行参数
	flag.Parse()
}

// 客户端入口
func main() {
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("NewClient err")
		return
	}

	fmt.Printf("Success to connect server, ServerAddr: [%s, %d]\n", client.ServerIp, client.ServerPort)

	// 阻塞处理业务
	for {
	}
}
