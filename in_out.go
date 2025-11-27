//go:build in_out

package main

import "fmt"

func main() {
	// 自动换行输出函数
	fmt.Println("Hello, Println!")
	fmt.Println(true)
	fmt.Println(123)

	// 格式化输出函数
	fmt.Printf("Hello, %s!", "Print")
	fmt.Printf("Pi is approximately %.2f\n", 3.14159)
	fmt.Printf("%v %v %v\n", 1, "two", 3.0) // 使用%v通用占位符
	fmt.Printf("%#v\n", "")                 // 显示值的Go语法表示a

	name := fmt.Sprintf("Hello, %s!", "Sprintf")
	fmt.Println(name)

	fmt.Println("----------------")

	// 输入函数
	var input string
	fmt.Print("Enter something: ")
	fmt.Scan(&input)
	fmt.Println("You entered:", input)
}
