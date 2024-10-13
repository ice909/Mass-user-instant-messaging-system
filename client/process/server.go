package process

import (
	"fmt"
	"net"
	"os"

	"github.com/ice909/go-common/utils"
)

// 显示登录成功后的界面
func ShowMenu() {
	fmt.Println("---------------恭喜xxx登录成功---------------")
	fmt.Println("\t\t1. 显示在线用户列表")
	fmt.Println("\t\t2. 发送消息")
	fmt.Println("\t\t3. 信息列表")
	fmt.Println("\t\t4. 退出系统")
	fmt.Println("--------------------------------------------")
	fmt.Println("请选择(1-4):")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入有误，请重新输入")
	}
}

func ProcessServerMes(conn net.Conn) {
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := utils.ReadPkg(conn)
		if err != nil {
			fmt.Println("utils.ReadPkg(conn) err=", err)
			return
		}
		// 如果读取到消息，又是下一步处理逻辑
		fmt.Println("mes=", mes)
	}
}
