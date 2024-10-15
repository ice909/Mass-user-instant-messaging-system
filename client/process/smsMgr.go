package process

import (
	"encoding/json"
	"fmt"

	"github.com/ice909/go-common/message"
)

// 输出群聊消息
func outputGroupMes(mes *message.Message) {
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(mes.Data), &smsMsg)
	if err != nil {
		fmt.Println("OutputGroupMes json.Unmarshal fail, err=", err)
		return
	}
	fmt.Printf("用户id:\t%d 说:\t%s\n", smsMsg.UserId, smsMsg.Content)
}
