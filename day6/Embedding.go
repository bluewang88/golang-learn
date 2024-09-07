package main

import (
	"fmt"
)

type A struct {
	Name string
	age int
}

func (a *A) Greet() {
	fmt.Println("Hello, I am ", a.Name)
}

func (a *A) GetAge() int {
	return a.age
}
func (a *A) SetAge(age int) {
	a.age = age
}

func (a *A) SetName(name string) {
	a.Name = name
}
func (a *A) GetName() string {
	return a.Name
}

func main() {
	fmt.Println("Embedding in Go")
}
