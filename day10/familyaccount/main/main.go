package main

import (
	"fmt"
	"utils"
)

func main() {
	fmt.Println("Hello, My Account Software")
	// 创建一个FamilyAccount实例
	familyAccount := utils.NewFamilyAccount()
	// 显示主菜单
	familyAccount.MainMenu()
}
