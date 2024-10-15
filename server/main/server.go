package main

import (
	"fmt"
	"net"
	"server/model"
	"time"
)

// main

func main() {
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()

	// 监听成功,等待客户端连接服务器
	for {
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

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
