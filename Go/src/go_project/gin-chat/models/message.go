package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
	"sync"
)

type Message struct {
	gorm.Model
	FromId   int64  // 发送者
	TargetId int64  // 接收者
	Type     int64  // 消息类型 群聊、私聊、广播
	Media    string // 消息类型 文字、图片、音频
	Content  string // 消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int // 其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 10)

// 读写锁
var rwLocker sync.RWMutex

func Chat(writer http.ResponseWriter, request *http.Request) {
	// 校验token
	query := request.URL.Query()
	userId := query.Get("userId")
	userIds, _ := strconv.ParseInt(userId, 10, 64)
	//msgType := query.Get("type")
	//targetId := query.Get("targetId")
	//context := query.Get("context")
	isValid := true

	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isValid
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取连接
	node := &Node{
		conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	// 用户关系
	// userId跟node绑定并加锁
	rwLocker.Lock()
	clientMap[userIds] = node
	rwLocker.Unlock()
	// 完成发送逻辑
	go sendProc(node)
	// 完成接受逻辑
	go recProc(node)
	sendMsg(userIds, []byte("欢迎进入聊天室......"))

}

func recProc(node *Node) {
	for {
		_, data, err := node.conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws] <<<<<", data)
	}
}

var udpSendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpSendChan <- data
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func init() {
	go udpSendProc()
	go udpRecvProc()
}

func udpRecvProc() {
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer func(udpConn *net.UDPConn) {
		err := udpConn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(udpConn)

	for {
		var buf [512]byte
		n, err := udpConn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
		}
		dispatch(buf[0:n])
	}
}

func dispatch(data []byte) {
	var msg Message

	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
	}

	switch msg.Type {
	case 1: // 私信
		sendMsg(msg.FromId, data)
	case 2: // 群发
		sendGroupMsg()
	case 3: // 广播
		sendAllMsg()
	}

}

func sendAllMsg() {

}

func sendGroupMsg() {

}

func sendMsg(userId int64, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.Unlock()
	if ok {
		node.DataQueue <- msg
	}
}

func udpSendProc() {
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})
	defer func(udpConn *net.UDPConn) {
		err := udpConn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(udpConn)
	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case data := <-udpSendChan:
			_, err := udpConn.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
