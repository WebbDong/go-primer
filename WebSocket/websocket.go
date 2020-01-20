package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var addr = flag.String("addr", "127.0.0.1:10000", "http service address")
var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// https://github.com/gorilla/websocket
// go get github.com/gorilla/websocket
func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/websocket", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		/*		mt, message, err := c.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					break
				}
				log.Printf("recv: %s", message)*/
		err = c.WriteMessage(websocket.TextMessage, []byte("Hello World!!!"))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
