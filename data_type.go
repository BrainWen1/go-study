//gp:build data_type
package main

import "fmt"

func main() {
	// 整型
	var intVar int = 100
	fmt.Println(intVar)

	// 浮点型
	var floatVar float64 = 3.1415
	fmt.Println(floatVar)

	// 布尔型
	var boolVar bool = true
	fmt.Println(boolVar)

	// 字符型
	var charVar rune = 'A'
	fmt.Println(charVar)

	// 字符串型
	var strVar string = "Hello, Go!"
	fmt.Println(strVar)
}
