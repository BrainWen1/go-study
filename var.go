//go:build var

package main

import "fmt"

// 全局变量声明
var globalVar string = "I am a global variable" // 可以只定义不使用

// 多个全局变量声明
var (
	globalA int    = 10
	globalB string = "globalB"
)

func main() {
	// 常规变量声明
	var x int = 42
	fmt.Println("The value of x is:", x)

	// 自动类型推导
	var foo = "haha"
	fmt.Println("Foo is:", foo)

	// 简短变量声明
	name := "Brian"
	fmt.Println("Hello,", name)

	// 多变量声明
	var a, b, c int = 1, 2, 3
	fmt.Println("Values are:", a, b, c)

	// 常量声明
	const pi = 3.14
	fmt.Println("The value of pi is:", pi)

	fmt.Println(globalVar)
	fmt.Println("Global A:", globalA)
	fmt.Println("Global B:", globalB)
}
