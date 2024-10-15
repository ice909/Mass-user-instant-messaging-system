package main

import (
	"fmt"
	"net"

	"server/process"

	"github.com/ice909/go-common/message"
	"github.com/ice909/go-common/utils"
)

type Processor struct {
	Conn net.Conn
}

// 编写一个serverProcessMes函数
// 功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (processor Processor) ServerProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMsgType: // 处理登录
		up := &process.UserProcess{Conn: processor.Conn}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMsgType: // 处理注册
		up := &process.UserProcess{Conn: processor.Conn}
		err = up.ServerProcessRegister(mes)
	case message.SmsMsgType: // 处理群发消息
		sp := &process.SmsProcess{}
		err = sp.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (processor Processor) process2() {
	// 循环读取客户端发送的数据
	for {
		mes, err := utils.ReadPkg(processor.Conn)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("readPkg err=", err)
			}
			return
		}
		processor := &Processor{Conn: processor.Conn}
		err = processor.ServerProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes err=", err)
			return
		}
	}
}
