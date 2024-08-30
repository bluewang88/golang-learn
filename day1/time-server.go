package main

import (
	"fmt"
	"net"
	"time"
)

func handler(conn net.Conn) {
	fmt.Fprintf(conn, "Hello, I am a time server\n %s", time.Now().String())
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil { // 如果监听失败
		fmt.Println("Error listening", err.Error())
		return // 终止程序
	}
	fmt.Println("Listening on localhost:8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		go handler(conn) // 开启协程处理连接
	}
}
