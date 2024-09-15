package main

import (
	"fmt"
	"model"
	"service"
)

type customerView struct {
	// 定义必要字段
	key     string // 接收用户输入的字符
	loop    bool   // 表示是否循环的显示主菜单
	service *service.CustomerService
}

// 显示所有的客户信息
func (cv *customerView) list() {
	// 获取到当前所有的客户信息，放在切片中
	customers := cv.service.List()
	// 显示
	fmt.Println("----------客户列表----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("----------客户列表完成----------")
}

// 得到用户输入信息，构建新的客户，并完成添加
func (cv *customerView) add() {
	fmt.Println("----------添加客户----------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)
	// 构建一个customer对象
	customer := model.NewCustomerWithoutId(name, gender, age, phone, email)
	// 调用service完成添加
	if cv.service.Add(customer) {
		fmt.Println("-----------添加成功----------")
	} else {
		fmt.Println("-----------添加失败----------")
	}
}

// 删除客户
func (cv *customerView) delete() {
	fmt.Println("----------删除客户----------")
	fmt.Println("请选择待删除客户编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(Y/N):")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if cv.service.Delete(id) {
			fmt.Println("-----------删除成功----------")
		} else {
			fmt.Println("-----------删除失败，输入的id号不存在----------")
		}
	}
}

// 退出程序
func (cv *customerView) exit() {
	fmt.Println("确认是否退出(Y/N):")
	for {
		fmt.Scanln(&cv.key)
		if cv.key == "y" || cv.key == "Y" || cv.key == "n" || cv.key == "N" {
			break
		}
		fmt.Println("你的输入有误，请重新输入")
	}
	if cv.key == "y" || cv.key == "Y" {
		cv.loop = false
	}
}

// 显示主菜单
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("----------客户信息管理软件----------")
		fmt.Println("            1 添加客户")
		fmt.Println("            2 修改客户")
		fmt.Println("            3 删除客户")
		fmt.Println("            4 客户列表")
		fmt.Println("            5 退    出")
		fmt.Print("请选择(1-5):")

		//获取用户输入
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			fmt.Println("添加客户")
			cv.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
			cv.delete()
		case "4":
			// fmt.Println("客户列表")
			cv.list()
		case "5":
			fmt.Println("退出")
			cv.exit()
		default:
			fmt.Println("输入有误，请重新输入")

		}

		if !cv.loop {
			break
		}

	}

	fmt.Println("你退出了客户关系管理系统")

}

func main() {
	// 创建一个CustomerService实例
	// customerService := service.NewCustomerService()
	// 显示主菜单
	// 创建一个customerView实例，并运行主菜单
	customerView := customerView{
		key:     "",
		loop:    true,
		service: service.NewCustomerService(),
	}
	customerView.mainMenu()
}
