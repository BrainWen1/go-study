package main

import "net/http"

// handler
func handler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("ok, your request is " + req.Method))
}

func main() {
	// 回调函数
	http.HandleFunc("/index", handler)
	// 绑定服务
	http.ListenAndServe(":8080", nil)
}
