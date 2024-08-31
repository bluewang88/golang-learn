package main

import (
	"fmt"
	"mode"
)

func main() {
	//创建一个Stuent的实例
	student := mode.Student{
		Name:  "Tom",
		Age:   18,
		Score: 98.5,
	}
	fmt.Println(student)

	// // 通过工厂模式创建一个Student实例
	// student := factory.NewStudent("Tom", 18, 98.5)
	// fmt.Println(student)
}
