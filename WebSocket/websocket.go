package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", "127.0.0.1:10000", "http service address")
var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var websocketConnMap = make(map[string]*websocket.Conn)
var handleChanel = make(chan string, 5)

// https://github.com/gorilla/websocket
// go get github.com/gorilla/websocket
func main() {
	flag.Parse()
	log.SetFlags(0)
	go handle()
	go ticker()
	http.HandleFunc("/websocket", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	clientId := string(message)
	websocketConnMap[clientId] = conn
}

func handle() {
	for {
		if clientId, ok := <-handleChanel; ok {
			conn := websocketConnMap[clientId]
			if conn == nil {
				continue
			}
			err := conn.WriteMessage(websocket.TextMessage, []byte("1"))
			if err != nil {
				log.Println("write:", err)
			}
		}
	}
}

func ticker() {
	quit := make(chan bool)
	// 每隔2秒执行一次
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		i := 0
		for {
			<-ticker.C
			if i == 0 {
				handleChanel <- "f6156870-305f-47b4-8556-b9763c16bb83"
			} else {
				handleChanel <- "f6156870-305f-47b4-8556-b9763c16bb83abc"
			}
			if i == 2 {
				// 停止ticker
				ticker.Stop()
				// 结束主Go程，子Go程也一起结束
				quit <- true
			}
			i++
		}
	}()
	<-quit
}
