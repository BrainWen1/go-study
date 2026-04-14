package main

import "fmt"

// 1. 定义接口：只规定 '说话' 的行为
type AnimalInterface interface {
	speak()
}

// 2. 具体类型实现接口（非侵入式，无需显式声明implements）
type chicken struct{ name string }

func (c chicken) speak() { fmt.Println("咯咯咯") }

type cat struct{ name string }

func (c cat) speak() { fmt.Println("喵喵喵") }

// 3. 通用函数：只依赖接口，不关心具体类型
func speak(obj AnimalInterface) {
	obj.speak()
}

// 类型断言
func sing(obj AnimalInterface) {
	switch obj.(type) {
	case chicken:
		fmt.Println("鸡")
	case cat:
		fmt.Println("猫")
	}
	obj.speak()
}

func say(obj AnimalInterface) {
	Type, judge := obj.(chicken) // 断言类型 布尔类型
	if judge {
		fmt.Println("是只鸡，名字叫：", Type.name)
	} else {
		fmt.Println("不是鸡")
	}

	// c := obj.(cat) // 断言类型, 不正确会报错
	// c.speak()
	c, ok := obj.(cat) // 断言类型 布尔类型
	fmt.Println(c, ok)
	c1, ok1 := obj.(chicken)
	fmt.Println(c1, ok1)
}

// 空接口
// type EmptyInterface interface{}
type EmptyInterface any

func print(obj EmptyInterface) {
	fmt.Print(obj, " ")
}

func println(obj any) {
	fmt.Println(obj)
}

func main() {
	// 接口是一种类型，是方法的集合
	// 接口类型的变量可以保存任何实现了该接口的值
	ch := chicken{name: "kunkun"}
	ca := cat{name: "miaomiao"}

	var animal AnimalInterface
	animal = ch
	animal.speak()
	animal = ca
	animal.speak()

	speak(ch)
	speak(ca)

	sing(ch)
	sing(ca)

	say(ch)
	say(ca)

	// 空接口
	print("hello world")
	print(123)
	print(3.14159)
	fmt.Println()

	println("hello go")
	println(456)
	println(2.71828)
}
