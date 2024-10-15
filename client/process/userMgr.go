package process

import (
	"fmt"

	"github.com/ice909/go-common/message"
)

// 客户端要维护一个在线用户的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 0)

// 在客户端显示在线的用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表:")
	for id, user := range onlineUsers {
		fmt.Println("用户id:\t", id, "状态:\t", user.UserStatus)
	}
}

// 处理返回的NotifyUserStatusMsg消息
func updateUserStatus(notifyUserStatusMsg *message.NotifyUserStatusMsg) {
	user, ok := onlineUsers[notifyUserStatusMsg.UserId]
	if !ok {
		user = &message.User{
			UserId:     notifyUserStatusMsg.UserId,
			UserStatus: notifyUserStatusMsg.Status,
		}
		onlineUsers[notifyUserStatusMsg.UserId] = user
	} else {
		user.UserStatus = notifyUserStatusMsg.Status
	}

	outputOnlineUser()
}
