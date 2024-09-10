package utils

import "fmt"

type FamilyAccount struct {
	// 声明必要的字段
	// 保存接收用户输入的选项
	key string
	// 控制是否退出for循环
	loop bool
	// 定义账户的余额
	balance float64
	// 每次收支的金额
	money float64
	// 每次收支的说明
	note string
	// 收支的详情使用字符串来记录
	// 当有收支时，只需要对details进行拼接处理即可
	details string
	// 是否有收支的标记
	flag bool
	// 退出前提示信息
	exitMsg string
	// 退出提示信息
	exitMsgFlag bool
}

//给结构体添加方法

// 显示主菜单
func (this *FamilyAccount) MainMenu() {
	// 显示主菜单
	// 1. 收支明细
	// 2. 登记收入
	// 3. 登记支出
	// 4. 退出软件
	// 请选择(1-4):
	// 获取用户输入
	// 根据用户输入执行相应的函数
	// 声明一个变量，接收用户输入
	// 控制是否退出for循环
	// 循环显示主菜单
	for {
		//判断是否退出软件
		if this.exitMsgFlag {
			break
		}
		if this.loop {
			this.showMenu()
		}
	}
}
func (this *FamilyAccount) showMenu() {
	// 循环显示主菜单
	fmt.Println("----------家庭收支记账软件----------")
	fmt.Println("          1 收支明细")
	fmt.Println("          2 登记收入")
	fmt.Println("          3 登记支出")
	fmt.Println("          4 退出软件")
	fmt.Print("请选择(1-4):")

	// 获取用户输入
	fmt.Scanln(&this.key)
	// 根据用户输入执行相应的函数
	switch this.key {
	case "1":
		this.showDetails()
	case "2":
		this.income()
	case "3":
		this.pay()
	case "4":
		this.exit()
	default:
		fmt.Println("请输入正确的选项")
	}
}

// 显示明细
func (this *FamilyAccount) showDetails() {
	fmt.Println("----------当前收支明细记录----------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细...来一笔吧!")
	}
}

// 登记收入
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	// 将本次收入情况，拼接到details变量
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 登记支出
func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
		return
	}
	this.balance -= this.money
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	// 将本次支出情况，拼接到details变量
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 退出
func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗?y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入y/n")
	}
	if choice == "y" {
		this.exitMsgFlag = true
	}
}

// 工厂模式
// 创建一个FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:         "",
		loop:        true,
		balance:     10000.0,
		money:       0.0,
		note:        "",
		details:     "收支\t账户金额\t收支金额\t说    明",
		flag:        false,
		exitMsg:     "你退出了家庭收支记账软件的使用...",
		exitMsgFlag: false,
	}
}
