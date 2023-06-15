package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string      //用户名称
	Addr   string      //用户地址
	C      chan string //储存用户需要向客户端发送的消息的channel
	conn   net.Conn    //客户端连接，向客户端发送消息
	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	//启动当前user channel消息的goroutine
	go user.ListemMessage()
	return user
}

//用户的上线业务
func (this *User) Online() {
	//用户上线，将用户加入到onlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	//广播当前用户上线消息
	this.server.BroadCast(this, "已上线")
}
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询当前在线用户
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := strings.Split(msg, "|")[1]
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("当前用户名被使用\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()
			this.Name = newName
			this.SendMsg("您已经更新用户名" + this.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		//获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("消息格式不正确")
			return
		}
		//根据用户名得到对方的user对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("该用户名不存在")
		}
		//获取消息内容
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("无消息内容")
			return
		}
		remoteUser.SendMsg(this.Name + "对您说" + content)
		this.server.BroadCast(this, msg)
	}

}

//用户的下线业务
func (this *User) Offline() {
	//用户下线，将用户从onlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
}
func (this *User) ListemMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + '\n'))
	}
}
