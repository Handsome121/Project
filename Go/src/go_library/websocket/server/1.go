package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var conns []*websocket.Conn

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := UP.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	conns = append(conns, conn)
	for {
		message, i, err := conn.ReadMessage()
		if err != nil {
			break
		}
		for i := range conns {
			err = conns[i].WriteMessage(websocket.TextMessage, []byte("你说的是:"+string(i)+"吗？"))
			if err != nil {
				break
			}
		}

		fmt.Println(message, string(i))
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	fmt.Println("服务关闭")
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		return
	}
}
