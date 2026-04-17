package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 实例化一个http客户端
	client := new(http.Client)
	// 构造请求对象
	req, err := http.NewRequest("GET", "http://localhost:8080/index", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	// 发请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()
	// 获取响应
	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println(string(b))
}
