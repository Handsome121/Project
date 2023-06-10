package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip        string
	Port      int
	OnlineMap map[string]*User
	mapLock   sync.Mutex
	//消息广播的channel
	Message chan string
}

//创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
		//存储当前上线用户
		OnlineMap: make(map[string]*User),
		//用于广播的消息
		Message: make(chan string),
	}
	return server
}

//监听Message广播消息channel的goroutine，一旦有消息，就发送给全部用户的channel
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		//将msg发送给全部的在线User
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}

}
func (this *Server) BroadCast(user *User, s string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + s
	this.Message <- sendMsg
}
func (this *Server) Handler(conn net.Conn) {
	//Fmt.Println("连接建立成功...")
	//用户上线，将用户加入到onlineMap中
	user := NewUser(conn, this)
	user.Online()
	//监听用户是否活跃的channel
	islive := make(chan bool)
	//接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
			}
			if err != nil && err != io.EOF {
				fmt.Printf("Conn Read err:", err)
				return
			}
			//提取用户的消息
			msg := string(buf[:n-1])
			//将得到的消息进行广播
			user.DoMessage(msg)
			//用户的任意消息，代表当前用户是一个活跃的
			islive <- true
		}
	}()
	//当前handler阻塞
	for {
		select {
		case <-islive:
		//当前用户是活跃的，应该重置定时器
		case <-time.After(time.Second * 10):
			//已经超时
			//将当前用户强制踢出
			user.SendMsg("你被踢出了")
			//销毁用的资源
			close(user.C)
			//关闭连接
			conn.Close()
			//退出当前的Handler
			return
		}
	}

}

//启动服务器的接口
func (this *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}
	defer listener.Close()
	//启动监听Message的goroutine
	go this.ListenMessage()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err", err)
			continue
		}
		go this.Handler(conn)
	}
}
