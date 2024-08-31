package main

import "fmt"

type Person struct {
	//首字母大写表示public，可以被其他包访问
	Name  string
	Age   int
	Score [3]int
}

func main() {

	//标准实例化
	var p Person
	fmt.Printf("标准结构体声明后结构体变量p的地址：%p\n", &p) //{ 0}
	fmt.Println("标准结构体声明后结构体变量p的值：", p)      //{ 0}
	p.Name = "John"
	p.Age = 25
	fmt.Println("标准实例化p：", p) //{John 25}

	//使用new关键字实例化
	p1 := new(Person)
	p1.Name = "John"
	p1.Age = 25
	println("使用new关键字实例化p1：", p1) //&{John 25}

	//使用&符号实例化
	p2 := &Person{}
	p2.Name = "John"
	p2.Age = 25
	println("使用&符号实例化p2：", p2) //&{John 25}

	//使用&符号实例化并初始化
	p3 := &Person{Name: "John", Age: 25}
	println("使用&符号实例化并初始化p3：", p3) //&{John 25}

}
