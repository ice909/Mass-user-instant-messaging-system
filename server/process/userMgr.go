package process

import "errors"

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUsers的添加
func (userMgr *UserMgr) AddOnlineUser(up *UserProcess) {
	userMgr.onlineUsers[up.UserId] = up
}

// 完成对onlineUsers的删除
func (userMgr *UserMgr) DelOnlineUser(userId int) {
	delete(userMgr.onlineUsers, userId)
}

// 获取所有在线的用户
func (userMgr *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return userMgr.onlineUsers
}

// 获取单个在线用户
func (userMgr *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := userMgr.onlineUsers[userId]
	if !ok {
		err = errors.New("用户不在线")
		return
	}
	return
}
