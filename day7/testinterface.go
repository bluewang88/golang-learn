package main

import "fmt"

type usb interface {
	Start()
	Stop()
}

type phone struct {
	Name  string
	Brand string
}

func (p phone) Start() {
	fmt.Println(p.Brand, " ", p.Name, "Phone is starting...")
}

func (p phone) Stop() {
	fmt.Println(p.Brand, " ", p.Name, "Phone is stopping...")
}

type Camera struct {
	Name  string
	Brand string
}

func (c Camera) Start() {
	fmt.Println(c.Brand, " ", c.Name, "Camera is starting...")
}

func (c Camera) Stop() {
	fmt.Println(c.Brand, " ", c.Name, "Camera is stopping...")
}

type Computer struct {
}

// Working method of Computer struct, which accepts usb interface
func (c Computer) Working(usb usb) {
	// Calling Start method of usb interface
	usb.Start()
	usb.Stop()
}

func main() {
	// Creating Phone struct
	p := phone{
		Name:  "OnePlus 7",
		Brand: "OnePlus",
	}

	// Creating Camera struct
	c := Camera{
		Name:  "Nikon D850",
		Brand: "Nikon",
	}

	// Creating Computer struct
	computer := Computer{}

	// Passing Phone struct to Working method of Computer struct
	computer.Working(p)

	// Passing Camera struct to Working method of Computer struct
	computer.Working(c)
}
