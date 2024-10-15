package process

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/ice909/go-common/message"
	"github.com/ice909/go-common/utils"
)

type SmsProcess struct{}

// 转发群聊消息
func (sp *SmsProcess) SendGroupMes(mes *message.Message) (err error) {
	var smsMsg message.SmsMsg
	err = json.Unmarshal([]byte(mes.Data), &smsMsg)
	if err != nil {
		fmt.Println("SendGroupMes json.Unmarshal fail, err=", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail, err=", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		if id == smsMsg.UserId {
			continue
		}
		sp.SendMesToEachOnlineUser(up.Conn, data)
	}
	return
}

func (sp *SmsProcess) SendMesToEachOnlineUser(conn net.Conn, data []byte) {
	err := utils.WritePkg(conn, data)
	if err != nil {
		fmt.Println("SendMesToEachOnlineUser utils.WritePkg fail, err=", err)
	}
}
