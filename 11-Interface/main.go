package main

import "fmt"

type AnimalInterface interface {
	speak()
}

type chicken struct {
	name string
}

type cat struct {
	name string
}

func (c chicken) speak() {
	fmt.Println("咯咯咯")
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

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
