package main

import "fmt"

// 定义一个结构体
// 该结构体的属性含有指针、slice、map等引用类型
type Person struct {
	Name  string
	Age   int
	Score [5]float64 // 数组,保存5门成绩
	// 指针
	Ptr *int
	// 切片
	Slice []int
	// map
	Map1 map[string]string
}

func main() {
	// 定义一个结构体变量
	var p1 Person
	fmt.Println("Print var struct p1 without init :", p1)
	if p1.Ptr == nil {
		fmt.Println(p1.Ptr, " :p1.Ptr is nil")
	}
	if p1.Slice == nil {
		fmt.Println(p1.Slice, " :p1.Slice is nil")
	}
	if p1.Map1 == nil {
		fmt.Println(p1.Map1, " :p1.Map1 is nil")

	}
	fmt.Println("Print var struct p1.Ptr before make :", p1.Ptr)
	fmt.Println("Print var struct p1.Slice before make :", p1.Slice)
	fmt.Println("Print var struct p1.Map1 before make :", p1.Map1)
	p1.Name = "Tom"
	p1.Age = 18
	// 为指针分配空间
	var num = 10
	p1.Ptr = &num
	// 为切片分配空间
	p1.Slice = make([]int, 10)
	// 为map分配空间
	p1.Map1 = make(map[string]string, 10)
	p1.Map1["key1"] = "value1"
	p1.Map1["key2"] = "value2"
	p1.Map1["key3"] = "value3"
	p1.Map1["key4"] = "value4"
	p1.Map1["key5"] = "value5"
	p1.Map1["key6"] = "value6"
	p1.Map1["key7"] = "value7"
	p1.Map1["key8"] = "value8"
	p1.Map1["key9"] = "value9"
	p1.Map1["key10"] = "value10"

	fmt.Println("Print var struct p1 after init :", p1)

	// 格式化输出结构体
	fmt.Printf("Print var struct p1 after init (%%v): %v\n", p1)
	fmt.Printf("Print var struct p1 after init (%%+v): %+v\n", p1)
	fmt.Printf("Print var struct p1 after init (%%#v): %#v\n", p1)

	// 分行输出结构体
	fmt.Printf("分行输出结构体:\nName: %s\nAge: %d\nPtr: %v\nSlice: %v\nMap1: %v\n",
		p1.Name, p1.Age, p1.Ptr, p1.Slice, p1.Map1)
}
