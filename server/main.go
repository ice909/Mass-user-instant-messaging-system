package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"server/message"
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
		go process(conn)
	}
}

func process(conn net.Conn) {
	// 这里需要延时关闭conn
	defer conn.Close()
	// 循环读取客户端发送的数据
	for {
		fmt.Println("服务器在等待客户端发送消息...")
		mes, err := readPkg(conn)
		if err != nil {
			fmt.Println("readPkg err=", err)
			return
		}
		fmt.Println("mes=", mes)
	}
}

func readPkg(conn net.Conn) (mes message.LoginMsg, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("read pkg header fail, err=", err)
		return
	}

	// 根据buf[:4]转成一个uint32类型
	var pkgLen uint32 = binary.BigEndian.Uint32(buf[0:4])

	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("read pkg dat fail, err = ", err)
		return
	}
	// 把buf[:pkgLen]反序列化成message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail, err=", err)
		return
	}

	return
}
