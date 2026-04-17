package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
)

type Server struct {
	// 服务器IP地址和端口
	Ip   string
	Port int

	// 在线用户列表
	UserMap map[string]*User
	mapLock sync.RWMutex // 保护UserMap的读写锁

	// 消息广播的channel
	ChBroad chan string
}

func NewServer(ip string, port int) *Server {
	ser := &Server{
		Ip:      ip,
		Port:    port,
		UserMap: make(map[string]*User),
		ChBroad: make(chan string, 100),
	}
	return ser
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen error: ", err)
		return
	}
	defer listener.Close()

	// 启动监听广播消息的goroutine
	go s.ListenBroad()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error: ", err)
			continue
		}

		// 在服务端播报连接信息
		fmt.Println("New client connected: ", conn.RemoteAddr().String())

		// 为每个新连接启动一个goroutine处理客户端请求
		go s.Handler(conn)
	}
}

func (s *Server) Handler(conn net.Conn) {
	// 创建一个用户对象
	user := NewUser(conn)

	// 更新在线用户列表
	s.mapLock.Lock()
	s.UserMap[user.Name] = user
	s.mapLock.Unlock()

	// 广播用户上线消息
	s.Broadcast(fmt.Sprintf("[%s]%s is online", user.Addr, user.Name))

	// 处理用户消息
	go s.ReadUserMsg(user)
}

func (s *Server) ReadUserMsg(user *User) {
	reader := bufio.NewReader(user.Conn)
	for {
		msg, err := reader.ReadString('\n')

		if err != nil {
			if err != io.EOF {
				fmt.Println("conn.Read error: ", err)
				return
			}

			// 用户下线，更新在线用户列表
			s.mapLock.Lock()
			delete(s.UserMap, user.Name)
			s.mapLock.Unlock()

			// 广播用户下线消息
			s.Broadcast(fmt.Sprintf("[%s]%s is offline", user.Addr, user.Name))

			fmt.Println("Client disconnected: ", user.Conn.RemoteAddr().String())
			return
		}

		if len(msg) > 0 {
			msg = strings.TrimRight(msg, "\r\n")
			s.Broadcast(fmt.Sprintf("[%s]: %s", user.Name, msg))
		}
	}
}

func (s *Server) Broadcast(msg string) {
	s.ChBroad <- msg
}

func (s *Server) ListenBroad() {
	// 阻塞等待广播消息，一旦有消息就发送给所有在线用户
	for {
		msg := <-s.ChBroad

		s.mapLock.Lock()
		for _, user := range s.UserMap {
			user.Ch <- msg
		}
		s.mapLock.Unlock()
	}
}
