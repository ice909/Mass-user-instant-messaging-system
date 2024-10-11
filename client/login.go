package main

import "fmt"

func login(userId int, userPwd string) (err error) {
	fmt.Println("userId:", userId)
	fmt.Println("userPwd:", userPwd)
	return nil
}
