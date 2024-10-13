package main

import (
	"fmt"
	"net"
)

// main

func main() {
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()

	// 监听成功,等待客户端连接服务器
	for {
		fmt.Println("等待客户端来连接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
			return
		}
		// 一旦连接成功，则启动一个协程和客户端保持通讯
		processor := &Processor{Conn: conn}
		go processor.process2()
	}
}
