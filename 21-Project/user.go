package main

import "net"

type User struct {
	Name string
	Addr string
	Ch   chan string
	Conn net.Conn
}

func NewUser(conn net.Conn) *User {
	user := User{
		Name: conn.RemoteAddr().String(),
		Addr: conn.RemoteAddr().String(),
		Ch:   make(chan string),
		Conn: conn,
	}

	// 启动监听用户消息的goroutine
	go user.ListenMsg()

	return &user
}

func (u *User) ListenMsg() {
	// 阻塞等待来自Ch的消息，一旦有消息就发送给客户端
	for {
		msg := <-u.Ch
		u.Conn.Write([]byte(msg + "\r\n"))
	}
}
