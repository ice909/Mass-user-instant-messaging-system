package process

import (
	"encoding/json"
	"fmt"
	"net"
	"server/model"

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
	user, err := model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMsg.Code = 500
			loginResMsg.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMsg.Code = 403
			loginResMsg.Error = err.Error()
		} else {
			loginResMsg.Code = 505
			loginResMsg.Error = "服务器内部错误..."
		}
	} else {
		loginResMsg.Code = 200
		fmt.Println(user, "登录成功")
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
