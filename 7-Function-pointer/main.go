package main

import (
	"fmt"
)

// 多个相同类型参数可以简写
func add(x, y int) int {
	return x + y
}

// 多个参数
func addpro(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// 多返回值
func fun(x, y int) (int, int) {
	return y, x
}

// 命名返回值
func namedFun(x, y int) (a int, b int) {
	a = y
	b = x
	return
}

// 闭包
func Counter() func() int {
	count := 0          // 闭包捕获的变量
	return func() int { // 返回的函数就是闭包
		count++
		return count
	}
}

func main() {
	a, b := 10, 20

	fmt.Println("Sum of a and b:", add(a, b))

	fmt.Println("Sum of 1 to 5:", addpro(1, 2, 3, 4, 5))
	fmt.Println("Sum of 10 to 12:", addpro(10, 11, 12))

	fmt.Println("a =", a, " b =", b)
	a, b = fun(a, b)
	fmt.Println("a =", a, " b =", b)

	a, b = namedFun(a, b)
	fmt.Println("a =", a, " b =", b)

	// 匿名函数
	result := func(x, y int) int {
		return x * y
	}(3, 4) // 立即调用
	fmt.Println("Result of anonymous function:", result)

	lambda := func(x, y int) int { // 匿名函数赋值给变量，相当于 lambda 表达式
		return x - y
	}
	fmt.Println("Result of lambda function:", lambda(10, 5)) // 调用

	// 以函数作为参数传递
	operate := func(x, y int, op func(int, int) int) int {
		return op(x, y)
	}
	fmt.Println("Result of higher-order function:", operate(10, 5, lambda))

	// 函数作为返回值
	makeAdder := func(x int) func(int) int {
		return func(y int) int {
			return x + y
		}
	}
	addFive := makeAdder(5)                                         // 返回一个将参数加5的函数
	fmt.Println("Result of function as return value:", addFive(10)) // 10 + 5 = 15

	// 递归函数
	var factorial func(n int) int // 递归函数要求首先声明
	factorial = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * factorial(n-1)
	}
	fmt.Println("Factorial of 5:", factorial(5))

	// 高阶函数
	var option int
	fmt.Println(`Choose an option:
1. Add
2. Subtract
3. Multiply
4. Divide`)
	fmt.Scan(&option)

	options := map[int]func(int, int) int{
		1: func(x, y int) int { return x + y },
		2: func(x, y int) int { return x - y },
		3: func(x, y int) int { return x * y },
		4: func(x, y int) int { return x / y },
	}

	fmt.Println("Result of 20 and 10:", options[option](20, 10)) // 根据用户选择调用不同的函数

	// 把函数提出来
	Add := func(x, y int) int {
		return x + y
	}
	Subtract := func(x, y int) int {
		return x - y
	}
	Multiply := func(x, y int) int {
		return x * y
	}
	Divide := func(x, y int) int {
		return x / y
	}

	funcMap := map[int]func(int, int) int{
		1: Add,
		2: Subtract,
		3: Multiply,
		4: Divide,
	}
	fmt.Println("Result of 30 and 15:", funcMap[option](30, 15))

	// 闭包
	c1 := Counter()
	fmt.Print(c1(), "")
	fmt.Print(c1(), "")
	fmt.Println(c1(), "")

	c2 := c1 // 复制闭包，c2 和 c1 共享同一个捕获的变量
	fmt.Print(c2(), "")
	fmt.Print(c2(), "")
	fmt.Println(c1(), "")

	c3 := Counter() // 新的闭包，捕获新的变量
	fmt.Print(c3(), "")
	fmt.Println(c3(), "")

	// 值传递和引用传递
	modify := func(num int) {
		fmt.Println(&num)
		num += 10
	}

	num := 10
	fmt.Println(num, &num)
	modify(num)
	fmt.Println(num, &num)

	modifypro := func(numPtr *int) {
		fmt.Println(numPtr)
		*numPtr += 100
	}

	fmt.Println(num, &num)
	modifypro(&num)
	fmt.Println(num, &num)
}
