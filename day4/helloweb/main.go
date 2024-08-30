package main

import (
	"fmt"
	"net/http"
)

//创建处理器函数

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Web", r.URL.Path)
}

func main() {
	//注册处理器函数
	http.HandleFunc("/", handler)
	//启动web服务,创建路由
	http.ListenAndServe(":8080", nil)
}
