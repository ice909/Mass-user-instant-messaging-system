package process

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"github.com/ice909/go-common/message"
	"github.com/ice909/go-common/utils"
)

type UserProcess struct{}

func (userProcess UserProcess) Login(userId int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	// 准备通过conn发送消息给服务器
	var mes message.Message

	mes.Type = message.LoginMsgType

	// 创建一个LoginMsg结构体
	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd

	// 将loginMsg序列化
	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	mes.Data = string(data)

	// 将mes序列化
	data, err = json.Marshal(mes)

	// 先把data长度发送给服务器
	var dataLen uint32 = uint32(len(data))

	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], dataLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("client send msg fail", err)
		return
	}

	// 这里还需要处理服务器端返回的消息
	mes, err = utils.ReadPkg(conn)
	if err != nil {
		fmt.Println("utils.ReadPkg(conn) err=", err)
		return
	}
	// 将mes的Data部分反序列化成LoginResMsg
	var loginResMsg message.LoginResMsg
	err = json.Unmarshal([]byte(mes.Data), &loginResMsg)
	if loginResMsg.Code != 200 {
		return errors.New(loginResMsg.Error)
	}
	go ProcessServerMes(conn)
	// 显示登录成功菜单
	for {
		ShowMenu()
	}
}

func (userProcess UserProcess) Register(userId int, userPwd, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegisterMsgType
	var registerMes message.RegisterMsg
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		return
	}

	var dataLen uint32 = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], dataLen)

	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("client send msg fail", err)
		return
	}

	// 这里还需要处理服务器端返回的消息
	mes, err = utils.ReadPkg(conn)
	if err != nil {
		fmt.Println("utils.ReadPkg(conn) err=", err)
		return
	}
	// 将mes的Data部分反序列化成LoginResMsg
	var registerResMsg message.RegisterResMsg
	err = json.Unmarshal([]byte(mes.Data), &registerResMsg)
	if registerResMsg.Code != 200 {
		return errors.New(registerResMsg.Error)
	}
	// 注册成功
	fmt.Println("注册成功,请重新登录")
	return
}
