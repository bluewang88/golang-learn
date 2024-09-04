package main

import (
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Score []float64
}

func main() {
	//直接声明struct变量
	var p1 Person
	fmt.Println("declare struct variable directly p1:", p1)

	//声明struct变量并初始化
	var p2 Person = Person{}
	fmt.Println("declare struct variable and init p2:", p2)

	// if p1 == p2 {
	// 	fmt.Println("p1 == p2")
	// } else {
	// 	fmt.Println("p1 != p2")
	// }

	//声明struct变量并初始化
	p3 := Person{}
	fmt.Println("declare struct variable and init p3:", p3)
	p3.Name = "Tom"
	p3.Age = 18
	p3.Score = []float64{98.5, 88.5, 78.5}
	fmt.Println("after assigining values to p3:", p3)

	//声明struct变量并同时初始化
	p4 := Person{Name: "Jerry", Age: 20, Score: []float64{98.5, 88.5, 78.5}}
	fmt.Println("declare struct variable and init p4:", p4)

	//声明struct变量并同时初始化
	p5 := Person{"Jack", 22, []float64{98.5, 88.5, 78.5}}
	fmt.Println("declare struct variable and init p5:", p5)

	//使用new函数创建struct变量,其返回的是指针
	var p6 *Person = new(Person)
	//因为p6是指针，所以访问成员变量时需要使用(*p6).Name
	(*p6).Name = "Mike"
	(*p6).Age = 25
	(*p6).Score = []float64{98.5, 88.5, 78.5}
	fmt.Println("declare struct variable and init *p6:", *p6)

	//(*p6).Name = "Mike" 可以简写为 p6.Name
	//原因是go编译器会自动根据p6的类型来判断是指针还是变量
	//如果p6是指针，底层则自动转换为(*p6).Name

	p6.Name = "Smith"
	p6.Age = 26
	p6.Score = []float64{97.5, 98.5, 88.5}
	fmt.Println("after assigining values to p6:", p6)

	//使用&符号创建struct变量,其返回的是指针
	p7 := &Person{}
	//因为p7是指针，所以访问成员变量时需要使用(*p7).Name
	(*p7).Name = "John"
	(*p7).Age = 28
	(*p7).Score = []float64{98.5, 88.5, 78.5}
	fmt.Println("declare struct variable and init *p7:", *p7)

	//(*p7).Name = "John" 可以简写为 p7.Name
	//原因是go编译器会自动根据p7的类型来判断是指针还是变量
	//如果p7是指针，底层则自动转换为(*p7).Name

	p7.Name = "Lily"
	p7.Age = 29
	p7.Score = []float64{97.5, 98.5, 88.5}
	fmt.Println("after assigining values to p7:", p7)

	//指针变量保存的是变量的地址

	var p8 *Person = &p2

	fmt.Println("p8:", p8)
	fmt.Printf("p2在内存中的地址%p\n", &p2)
	fmt.Printf("p8保存的值%p\n", p8)
	fmt.Printf("p8在内存中的地址%p\n", &p8)

}
