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
	// 增加一个字段，表示该Conn是哪个用户
	UserId int
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
		// 这里，因为用户登录成功，我们就把该登录成功的用户放入到userMgr中
		userProcess.UserId = loginMsg.UserId
		userMgr.AddOnlineUser(&userProcess)
		userProcess.NotifyOthersOnlineUser(loginMsg.UserId)
		// 将当前在线用户的id，放入到loginResMsg.UserIds
		for id := range userMgr.onlineUsers {
			loginResMsg.UserIds = append(loginResMsg.UserIds, id)
		}
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

func (userProcess UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 从mes中取出mes.Data,并直接反序列化成RegisterMsg
	var registerMsg message.RegisterMsg
	err = json.Unmarshal([]byte(mes.Data), &registerMsg)
	if err != nil {
		fmt.Println("serverProcessRegister() json.Unmarshal fail, err=", err)
		return
	}

	// 返回的消息
	var resMes message.Message
	resMes.Type = message.RegisterResMsgType

	// 再声明一个 RegisterResMsg
	var registerResMsg message.RegisterResMsg
	err = model.MyUserDao.Register(registerMsg.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMsg.Code = 505
			registerResMsg.Error = err.Error()
		} else {
			registerResMsg.Code = 500
			registerResMsg.Error = "服务器内部错误..."
		}
	} else {
		registerResMsg.Code = 200
		registerResMsg.Error = "注册成功"
	}

	// 对registerResMsg序列化
	data, err := json.Marshal(registerResMsg)
	if err != nil {
		fmt.Println("registerResMsg json.Marshal fail, err=", err)
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

// 通知所有用户
func (userProcess UserProcess) NotifyOthersOnlineUser(userId int) {
	// 遍历onlineUsers，然后一个一个发送NotifyUserStatusMsg
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}

func (userProcess UserProcess) NotifyMeOnline(userId int) {
	// 组装NotifyUserStatusMsg
	var mes message.Message
	mes.Type = message.NotifyUserStatusMsgType
	var notifyUserStatusMsg message.NotifyUserStatusMsg
	notifyUserStatusMsg.UserId = userId
	notifyUserStatusMsg.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMsg)
	if err != nil {
		fmt.Println("NotifyMeOnline json.Marshal fail, err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("NotifyMeOnline json.Marshal fail, err=", err)
		return
	}
	err = utils.WritePkg(userProcess.Conn, data)
	if err != nil {
		fmt.Println("NotifyMeOnline WritePkg fail, err=", err)
		return
	}
}
