package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/ice909/go-common/message"
	"github.com/ice909/go-common/utils"
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
		mes, err := utils.ReadPkg(conn)
		if err != nil {
			fmt.Println("readPkg err=", err)
			return
		}
		err = serverProcessMes(conn, &mes)
		if err != nil {
			fmt.Println("serverProcessMes err=", err)
			return
		}
	}
}

// 编写一个函数serverProcessLogin函数
// 功能：专门处理登录请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	// 从mes中取出mes.Data,并直接反序列化成LoginMsg
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(mes.Data), &loginMsg)
	if err != nil {
		fmt.Println("serverProcessLogin() json.Unmarshal fail, err=", err)
		return
	}

	// 返回的消息
	var resMes message.Message
	resMes.Type = message.LoginResMsgType

	// 再声明一个 LoginResMsg
	var loginResMsg message.LoginResMsg

	// 如果用户id=100，密码=123456，认为合法，否则不合法
	if loginMsg.UserId == 100 && loginMsg.UserPwd == "123456" {
		// 合法
		loginResMsg.Code = 200
	} else if loginMsg.UserId != 100 {
		// 不合法
		loginResMsg.Code = 500 // 500表示该用户不存在
		loginResMsg.Error = "该用户不存在，请注册再使用..."
	} else {
		// 不合法
		loginResMsg.Code = 403 // 403表示密码不正确
		loginResMsg.Error = "密码不正确，请重新输入..."
	}

	// 对loginResMsg序列化
	data, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("loginResMsg json.Marshal fail, err=", err)
		return
	}
	resMes.Data = string(data)
	// 对resMes序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("resMes json.Marshal fail, err=", err)
		return
	}
	err = utils.WritePkg(conn, data)
	return
}

// 编写一个serverProcessMes函数
// 功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMsgType:
		// 处理登录
		err = serverProcessLogin(conn, mes)
		// case message.RegisterMsgType:
		// 处理注册
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}
