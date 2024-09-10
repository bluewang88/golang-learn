package main

import (
	"fmt"
)

// 显示家庭收支软件主界面
func main() {
	// fmt.Println("Welcome to my account")
	// 显示主菜单
	// 1. 收支明细
	// 2. 登记收入
	// 3. 登记支出
	// 4. 退出软件
	// 请选择(1-4):
	// 获取用户输入
	// 根据用户输入执行相应的函数
	// 声明一个变量，接收用户输入
	var key int
	// 控制是否退出for循环
	loop := true
	// 循环显示主菜单
	for {
		fmt.Println("----------家庭收支记账软件----------")
		fmt.Println("          1 收支明细")
		fmt.Println("          2 登记收入")
		fmt.Println("          3 登记支出")
		fmt.Println("          4 退出软件")
		fmt.Print("请选择(1-4):")

		// 获取用户输入
		fmt.Scanln(&key)
		// 根据用户输入执行相应的函数
		switch key {
		case 1:
			fmt.Println("收支明细")
		case 2:
			fmt.Println("登记收入")
		case 3:
			fmt.Println("登记支出")
		case 4:
			fmt.Println("退出软件")
			loop = false
			// break
		default:
			fmt.Println("请输入正确的选项")
		}

		if !loop {
			break
		}
	}

	fmt.Println("你退出了家庭收支记账软件的使用...")

}
