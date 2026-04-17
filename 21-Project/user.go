package main

import (
	"fmt"
	"net"
)

type User struct {
	Name string      `json:"name"`
	Addr string      `json:"address"`
	Ch   chan string `json:"channel"`
	Conn net.Conn    `json:"connection"`
	serv *Server     `json:"server"`
}

func NewUser(conn net.Conn, server *Server) *User {
	user := User{
		Name: conn.RemoteAddr().String(),
		Addr: conn.RemoteAddr().String(),
		Ch:   make(chan string),
		Conn: conn,
		serv: server,
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

func (u *User) Offline() {
	// 更新在线用户列表
	u.serv.mapLock.Lock()
	delete(u.serv.UserMap, u.Name)
	u.serv.mapLock.Unlock()
	// 广播用户下线消息
	u.serv.Broadcast(fmt.Sprintf("[%s]%s is offline", u.Addr, u.Name))
}

func (u *User) Online() {
	// 更新在线用户列表
	u.serv.mapLock.Lock()
	u.serv.UserMap[u.Name] = u
	u.serv.mapLock.Unlock()
	// 广播用户上线消息
	u.serv.Broadcast(fmt.Sprintf("[%s]%s is online", u.Addr, u.Name))
}

func (u *User) DoMessage(msg string) {
	// 处理查询在线用户操作
	if msg == "who" {
		u.serv.mapLock.Lock()

		onlineMsg := "Oneline Users: [ "
		for _, user := range u.serv.UserMap {
			onlineMsg = fmt.Sprintf("%s%s ", onlineMsg, user.Name)
		}
		u.SendMsg(fmt.Sprintf("%s]", onlineMsg))

		u.serv.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename:" {
		newName := msg[7:] // 提取新用户名

		// 检查新用户名是否已存在
		_, exists := u.serv.UserMap[newName]
		if exists {
			u.SendMsg("Username already exists")
			return
		}

		u.serv.mapLock.Lock()
		// 从在线用户列表中删除旧用户名，并添加新用户名
		delete(u.serv.UserMap, u.Name)
		u.serv.UserMap[newName] = u

		u.serv.mapLock.Unlock()

		u.Name = newName // 更新用户对象的用户名

		// 向该用户回显新用户名修改成功的消息
		u.SendMsg(fmt.Sprintf("Username changed to %s", newName))
	} else {
		// 广播用户消息
		u.serv.Broadcast(fmt.Sprintf("[%s]: %s", u.Name, msg))
	}
}

func (u *User) SendMsg(msg string) {
	u.Conn.Write([]byte(msg + "\r\n"))
}
