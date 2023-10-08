package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
)

func main() {
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial("ws://127.0.0.1:8888", nil)
	if err != nil {
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte("wo"))
	if err != nil {
		return
	}

	go send(conn)

	for {
		m, i, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(m, string(i))
	}
}

func send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		err := conn.WriteMessage(websocket.TextMessage, []byte(l))
		if err != nil {
			break
		}
	}
}
