package model

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	Pool *redis.Pool
}

// 工厂模式
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		Pool: pool,
	}
	return
}

func (userDao UserDao) GetUserById(conn redis.Conn, userId int) (user *User, err error) {
	res, err := conn.Do("HGET", "users", userId)
	if res == nil || err != nil {
		err = ERROR_USER_NOTEXISTS
		return
	}
	err = json.Unmarshal(res.([]byte), &user)
	if err != nil {
		return
	}
	return
}

func (userDao UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := userDao.Pool.Get()
	defer conn.Close()
	user, err = userDao.GetUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
