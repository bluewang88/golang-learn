package main

import (
	"fmt"
)

// 定义接口
type BookInfo interface {
	GetTitle() string
	GetAuthor() string
}

// 定义结构体
type Book struct {
	Title  string
	Author string
	ID     int
}

// 实现接口方法
func (b Book) GetTitle() string {
	return b.Title
}

func (b Book) GetAuthor() string {
	return b.Author
}

func main() {
	// 创建 Book 实例
	book := Book{Title: "The Go Programming Language", Author: "Alan A. A. Donovan", ID: 1}

	// 定义接口类型变量
	var bookInfo BookInfo

	// 将 Book 实例的值赋值给接口类型变量
	bookInfo = book
	fmt.Printf("Using value: Title: %s, Author: %s\n", bookInfo.GetTitle(), bookInfo.GetAuthor())

	// 将 Book 实例的指针赋值给接口类型变量
	bookInfo = &book
	fmt.Printf("Using pointer: Title: %s, Author: %s\n", bookInfo.GetTitle(), bookInfo.GetAuthor())
}
