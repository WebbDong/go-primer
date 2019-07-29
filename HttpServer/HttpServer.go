package main

import (
	"fmt"
	"net/http"
)

// http服务器
func main() {
	// 注册回调函数，该函数在服务器被访问时调用
	http.HandleFunc("/test", httpHandler)
	// 绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func httpHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World"))

	fmt.Println("request.URL:", request.URL)
	fmt.Println("request.Header:", request.Header)
	fmt.Println("request.RemoteAddr:", request.RemoteAddr)
	fmt.Println("request.Method:", request.Method)
	fmt.Println("request.Body:", request.Body)
	fmt.Println("request.Close:", request.Close)
	fmt.Println("request.Host:", request.Host)
}
