package main

import "fmt"

// 定义一个函数，接受函数作为参数
func compute(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// 定义一个函数，返回函数作为结果
func getOperation(op string) func(int, int) int {
	if op == "add" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "subtract" {
		return func(a, b int) int {
			return a - b
		}
	}
	return nil
}

func main() {
	// 使用函数作为参数
	result := compute(5, 3, func(a, b int) int {
		return a * b
	})
	fmt.Println("Result of multiplication:", result)

	// 使用函数作为返回值
	addFunc := getOperation("add")
	result = addFunc(10, 5)
	fmt.Println("Result of addition:", result)
}
