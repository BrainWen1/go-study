package main

import (
	"fmt"
)

var global int = 0

// init 函数在 main 函数之前执行，可以有多个 init 函数，按顺序执行
// 一般用于初始化包级变量或执行启动时需要的设置
func init() {
	fmt.Println(global)
	fmt.Println("init1")
}

func init() {
	global = 1024
	fmt.Println(global)
	fmt.Println("init2")
}

func main() {
	fmt.Println("main function")

	fmt.Println(global)

	// defer 函数用于注册延迟调用，这些调用直到 return 前才被执行，后进先出
	// 常用于资源释放、文件关闭等操作
	defer fmt.Println("World")
	defer fmt.Println("!")

	fmt.Println("Hello")
}
