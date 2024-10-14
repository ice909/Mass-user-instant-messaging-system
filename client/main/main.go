package main

import (
	"client/process"
	"fmt"
	"os"
)

// main

var userId int
var userPwd string
var userName string

func main() {
	// 接收用户选择
	var key int
	// 判断用户选择
	for {
		fmt.Println("---------------欢迎登录多人聊天系统---------------")
		fmt.Println("\t\t1. 登录聊天室")
		fmt.Println("\t\t2. 注册用户")
		fmt.Println("\t\t3. 退出系统")
		fmt.Println("------------------------------------------------")
		fmt.Println("请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码:")
			fmt.Scanf("%s\n", &userPwd)
			// 先把登录函数写到login.go
			userProcess := &process.UserProcess{}
			err := userProcess.Login(userId, userPwd)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("登录成功")
			}
		case 2:
			fmt.Println("请输入用户的id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户的名字:")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Printf("\n您的输入有误，请重新输入!\n")
		}
	}
}
