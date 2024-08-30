package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func main() {
	s := g.Server()
	g.Cfg()

	// 测试日志
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("Hello, GoFrame!")
		glog.Info(r.GetCtx(), "你来了！")
		glog.Error(r.GetCtx(), "你异常了！")
	})

	s.BindHandler("/panic", func(r *ghttp.Request) {
		glog.Panic(r.GetCtx(), "你崩溃了！")
	})

	//post请求
	s.BindHandler("POST:/post", func(r *ghttp.Request) {
		r.Response.Write("Hello, POST GoFrame!")
	})

	// s.SetPort(8199) 配置文件编写
	s.Run()
}
