package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"github.com/ice909/go-common/message"
	"github.com/ice909/go-common/utils"
)

func login(userId int, userPwd string) (err error) {
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
	return nil
}
