package main

import "fmt"

// 自定义类型
type Code int

const (
	Success     Code = 200 // 请求成功
	NotFound    Code = 404 // 未找到
	ServerError Code = 500 // 服务器错误
)

// 为自定义类型添加方法
func (c Code) GetMsg() string {
	switch c {
	case Success:
		return "请求成功"
	case NotFound:
		return "未找到"
	case ServerError:
		return "服务器错误"
	}
	return "未知状态"
}

func WebServer(status int) (code Code, msg string) {
	switch status {
	case 1:
		code = Success
	case 2:
		code = NotFound
	case 3:
		code = ServerError
	default:
		code = Code(status)
	}

	msg = code.GetMsg()
	return
}

// 类型别名
// 1. 类型别名不能绑定方法
// 2. %T 显示的是别名的原始类型
// 3. 可以直接和原始类型互相赋值和比较
type MyInt = int // alias for int

func main() {
	statusCodes := []int{1, 2, 3, 4}
	for _, status := range statusCodes {
		code, msg := WebServer(status)
		fmt.Println(int(code), msg)
	}

	var code Code = 200
	var num MyInt = 100
	fmt.Printf("%v, %T\n", code, code)
	fmt.Printf("%v, %T\n", num, num)

	val := 200
	// fmt.Println(code == val)    // 自定义类型不可以直接和原始类型比较
	fmt.Println(code == Code(val)) // 需要转换类型后才能比较
	fmt.Println(num == val)        // 类型别名可以直接和原始类型比较
}
