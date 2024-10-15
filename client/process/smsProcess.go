package process

import (
	"encoding/json"
	"fmt"

	"github.com/ice909/go-common/message"
	"github.com/ice909/go-common/utils"
)

type SmsProcess struct{}

// 发送群聊消息
func (sp *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMsgType
	var smsMsg message.SmsMsg
	smsMsg.Content = content
	smsMsg.UserId = curUser.UserId
	smsMsg.UserStatus = curUser.UserStatus

	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail, err=", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail, err=", err)
		return
	}

	// 将mes发送给服务器
	err = utils.WritePkg(curUser.Conn, data)
	if err != nil {
		fmt.Println("SendGroupMes utils.WritePkg fail, err=", err)
		return
	}
	return
}
