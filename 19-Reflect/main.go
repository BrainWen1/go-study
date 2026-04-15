package main

import (
	"fmt"
	"reflect"
)

// 获取类型信息、值
func reflectType(val any) {
	t := reflect.TypeOf(val)

	switch t.Kind() {
	case reflect.Int:
		fmt.Println("Int")
	case reflect.Float64:
		fmt.Println("Float64")
	case reflect.String:
		fmt.Println("String")
	case reflect.Array:
		fmt.Println("Array")
	case reflect.Slice:
		fmt.Println("Slice")
	case reflect.Struct:
		fmt.Println("Struct")
	default:
		fmt.Println("Unknown type")
	}
}

func reflectValue(val any) {
	v := reflect.ValueOf(val)

	switch v.Kind() {
	case reflect.Int:
		fmt.Printf("Value: %d\n", v.Int())
	case reflect.Float64:
		fmt.Printf("Value: %f\n", v.Float())
	case reflect.String:
		fmt.Printf("Value: %s\n", v.String())
	default:
		fmt.Println("Unsupported type")
	}
}

// 修改值
func modifyValue(val any) {
	v := reflect.ValueOf(val)
	elem := v.Elem()

	switch elem.Kind() {
	case reflect.Int:
		elem.SetInt(666)
	case reflect.Float64:
		elem.SetFloat(6.66)
	case reflect.String:
		elem.SetString("666666666qpqpqp")
	default:
		fmt.Println("Unsupported type for modification")
	}
}

func main() {
	reflectType(100)
	reflectType(3.14)
	reflectType("Hello, World!")
	reflectType([]int{1, 2, 3})
	reflectType(struct {
		Name string
		Age  int
	}{
		Name: "Alice",
		Age:  30,
	})

	reflectValue(100)
	reflectValue(3.14)
	reflectValue("Hello, World!")
	reflectValue([]int{1, 2, 3})
	reflectValue(struct {
		Name string
		Age  int
	}{
		Name: "Andrew",
		Age:  25,
	})

	var num int = 100
	var str string = "Hello, World!"
	var f float64 = 3.14
	modifyValue(&num)
	modifyValue(&str)
	modifyValue(&f)
	fmt.Println(num)
	fmt.Println(str)
	fmt.Println(f)
}
