package process

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/ice909/go-common/message"
	"github.com/ice909/go-common/utils"
)

type UserProcess struct {
	Conn net.Conn
}

// 编写一个函数serverProcessLogin函数
// 功能：专门处理登录请求
func (userProcess UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
	err = utils.WritePkg(userProcess.Conn, data)
	return
}
