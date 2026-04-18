package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

type Client struct {
	// 要连接的服务器IP地址和端口
	ServerIp   string
	ServerPort int

	// 连接到服务器的TCP连接
	conn net.Conn

	Name   string
	option int // 用户选择的菜单选项

	input *bufio.Reader // 用于读取用户输入的缓冲读取器
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
		input:      bufio.NewReader(os.Stdin),
	}

	return client
}

func (c *Client) menu() bool {
	fmt.Println("1. Public Chat")
	fmt.Println("2. Private Chat")
	fmt.Println("3. Update User Name")
	fmt.Println("0. Exit")

	// 使用有缓冲读取整行输入，直到用户按下回车键
	line, err := c.input.ReadString('\n')
	if err != nil {
		fmt.Println("read menu option err:", err)
		return false
	}

	// 去除输入开头和结尾的空白字符，并将输入转换为整数
	option, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		fmt.Println("Invalid option, please enter number 0-3.")
		return false
	}

	c.option = option

	if c.option >= 0 && c.option <= 3 {
		return true
	} else {
		fmt.Println("Invalid option, please try again.")
		return false
	}
}

func (c *Client) UpdateName() bool {
	fmt.Println("Enter new user name:")

	name, err := c.input.ReadString('\n')
	if err != nil {
		fmt.Println("read user name err:", err)
		return false
	}

	c.Name = strings.TrimSpace(name)
	if c.Name == "" {
		fmt.Println("User name cannot be empty.")
		return false
	}

	// 发送更新用户名的消息给服务器
	msg := fmt.Sprintf("rename:%s\n", c.Name)
	_, err = c.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("c.conn.Write err:", err)
		return false
	}
	return true
}

func (c *Client) PublicChat() {
	for {
		fmt.Println("Enter message to send (type 'exit' to return to menu):")

		// 使用有缓冲读取整行输入，直到用户按下回车键
		msg, err := c.input.ReadString('\n')
		if err != nil {
			fmt.Println("read message err:", err)
			return
		}

		// 去除输入消息开头和结尾的空白字符
		msg = strings.TrimSpace(msg)

		if msg == "exit" {
			break
		} else if msg == "" {
			continue
		}

		// 发送消息给服务器
		_, err = c.conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("c.conn.Write err:", err)
			return
		}
	}
}

func (c *Client) PrivateChat() {
	// 显示在线用户列表
	c.ShowOnlineUsers()

	for {
		// 选择私聊对象
		var target string
		fmt.Println("user name to chat with (type 'exit' to return to menu):")

		target, err := c.input.ReadString('\n')
		if err != nil {
			fmt.Println("read target user name err:", err)
			return
		}

		target = strings.TrimSpace(target)

		if target == "exit" {
			break
		} else if target == "" {
			fmt.Println("Target user name cannot be empty.")
			continue
		}

		// 输入私聊消息并处理后发送回服务器
		for {
			fmt.Printf("message to send to '%s' (type 'exit' to choose another user):\n", target)

			// 使用有缓冲读取整行输入，直到用户按下回车键
			msg, err := c.input.ReadString('\n')
			if err != nil {
				fmt.Println("read message err:", err)
				return
			}

			// 去除输入消息开头和结尾的空白字符
			msg = strings.TrimSpace(msg)

			if msg == "exit" {
				break
			} else if msg == "" {
				continue
			}

			// 发送私聊消息给服务器，格式为 "to|target|message"
			_, err = c.conn.Write([]byte(fmt.Sprintf("to:%s:%s\n", target, msg)))
			if err != nil {
				fmt.Println("c.conn.Write err:", err)
				return
			}
		}
	}
}

func (c *Client) ShowOnlineUsers() {
	// 发送查询在线用户的消息给服务器
	_, err := c.conn.Write([]byte("who\n"))
	if err != nil {
		fmt.Println("c.conn.Write err:", err)
		return
	}
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
			c.PublicChat()
		case 2: // 私聊
			fmt.Println("Private Chat selected")
			c.PrivateChat()
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
