package main

import (
	"fmt"
	"time"
)

func Task1() {
	for {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "正在处理Task1")
		time.Sleep(3 * time.Second)
	}

}

func Task2() {
	for {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "正在处理Task2")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go Task1()
	go Task2()
	// for {
	// 	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "main goroutine")
	// 	time.Sleep(2 * time.Second)
	// }
}
