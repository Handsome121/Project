package models

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"sync"
)

type Message struct {
	gorm.Model
	FromId   string // 发送者
	TargetId string // 接收者
	Type     string // 消息类型 群聊、私聊、广播
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
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

func Chat(writer http.ResponseWriter, request *http.Request) {
	// 校验token
	query := request.URL.Query()
	userId := query.Get("userId")
	userIds, _ := strconv.ParseInt(userId, 10, 64)
	msgType := query.Get("type")
	targetId := query.Get("targetId")
	context := query.Get("context")
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

func broadMsg(data []byte) {

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
