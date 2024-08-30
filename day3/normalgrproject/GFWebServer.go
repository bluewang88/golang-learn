package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	// This is a web server that listens on port 8199
	// and returns a simple message to the client.
	// You can access the server by visiting http://localhost:8199
	// in your browser.
	// This is a simple example of how to create a web server using GoFrame.
	// You can create more complex web servers by adding more routes and handlers.
	// For more information on how to create web servers using GoFrame,
	// please refer to the documentation at https://goframe.org.
	// You can also find more examples of web servers in the GoFrame repository
	// at
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("Hello, GoFrame!")
	})
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Write("Hello, world!")
	})

	s.SetPort(8199)
	s.Run()
}
