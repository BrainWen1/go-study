package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建 TCP 地址
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	// 连接到服务器
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 发送数据给服务器
	message := "Hello, Server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		panic(err)
	}

	// 读取服务器的响应
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Received %d bytes from server:\n%s\n", n, string(buffer[:n]))
}
