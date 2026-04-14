package main

import "fmt"

// 主动返回error并处理
func divide(a, b int) (ret int, err error) {
	if b == 0 {
		ret = 0
		err = fmt.Errorf("除数不能为 0")
		return
	}
	ret = a / b
	err = nil
	return
}

func test() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("发生错误: %v\n", r)
		}
	}()
	panic("主动抛出一个错误")
}

func main() {
	ret, err := divide(10, 0)
	if err != nil {
		fmt.Printf("发生错误: %v\n", err)
	} else {
		fmt.Printf("结果: %v\n", ret)
	}

	test()
	fmt.Println("程序继续执行...")
}
