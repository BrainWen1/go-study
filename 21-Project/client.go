package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	// 要连接的服务器IP地址和端口
	ServerIp   string
	ServerPort int

	// 连接到服务器的TCP连接
	conn net.Conn

	Name   string
	option int // 用户选择的菜单选项
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

func (c *Client) menu() bool {
	fmt.Println("1. Public Chat")
	fmt.Println("2. Private Chat")
	fmt.Println("3. Update User Name")
	fmt.Println("0. Exit")

	fmt.Scanln(&c.option)

	if c.option >= 0 && c.option <= 3 {
		return true
	} else {
		fmt.Println("Invalid option, please try again.")
		return false
	}
}

func (c *Client) UpdateName() bool {
	fmt.Println("Enter new user name:")
	fmt.Scanln(&c.Name)

	// 发送更新用户名的消息给服务器
	msg := fmt.Sprintf("rename:%s\n", c.Name)
	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("c.conn.Write err:", err)
		return false
	}
	return true
}

func (c *Client) Run() {
	for {
		if !c.menu() {
			continue
		}

		// 根据用户选择的菜单选项执行相应的操作
		switch c.option {
		case 1: // 公聊
			fmt.Println("Public Chat selected")
		case 2: // 私聊
			fmt.Println("Private Chat selected")
		case 3: // 更新用户名
			fmt.Println("Update User Name selected")
			c.UpdateName()
		case 0: // 退出
			fmt.Println("Exiting...")
			return
		}
	}
}

func (c *Client) ListenMsg() {
	// 把连接重定向到标准输出
	io.Copy(os.Stdout, c.conn)
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
	go client.ListenMsg() // 启动监听服务器消息的goroutine
	client.Run()
}
