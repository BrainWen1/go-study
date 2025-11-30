package main

import (
	"fmt"
	"sort"
)

func main() {
	// 标准形式
	var arr1 [3]string = [3]string{"hello", "world", "go"}
	fmt.Println(arr1)

	// 可省略类型
	var arr2 = [3]string{"Golang", "is", "awesome"}
	fmt.Println(arr2)

	// 根据初始值数量自动推断长度
	var arr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	// 索引：与 C/C++ 无异
	fmt.Println(arr3[0]) // 访问第一个元素
	arr3[1] = 20         // 修改第二个元素
	fmt.Println(arr3)

	// len 函数获取数组长度
	fmt.Println(arr3[len(arr3)-1]) // 访问最后一个元素

	// 切片
	var vec1 []int = []int{10, 20, 30, 40, 50}
	fmt.Println(vec1)

	// append 函数添加元素
	vec1 = append(vec1, 666)
	fmt.Print(vec1)
	fmt.Println("", len(vec1))

	// make 函数创建切片
	vec2 := make([]string, 3)
	vec2[0] = "Go"
	vec2[1] = "is"
	vec2[2] = "fun"
	fmt.Println(vec2, len(vec2), cap(vec2))
	fmt.Println(vec2 == nil) // 判断是否为 nil 切片 --> false

	vec3 := make([]int, 0)
	fmt.Println(vec3 == nil) // 空切片不是 nil 切片 --> false

	var vec4 []int           // nil 切片
	fmt.Println(vec4 == nil) // true

	vec5 := make([]int, len(vec1), cap(vec1))
	fmt.Println(vec5, len(vec5), cap(vec5)) // [0 0 0 0 0 0] 6 10

	// 切片截取
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := nums[2:5] // 包含索引2到4的元素
	fmt.Println(slice1)

	slice2 := nums[:4] // 包含索引0到3的元素
	fmt.Println(slice2)

	slice3 := nums[5:] // 包含索引5到最后的元素
	fmt.Println(slice3)

	slice4 := nums[:] // 包含所有元素
	fmt.Println(slice4)

	// 切片排序
	fmt.Println("--------------------------")
	nums = []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println(nums)

	sort.Ints(nums)   // 使用 sort 包对切片进行排序
	fmt.Println(nums) // 升序

	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // 降序
	fmt.Println(nums)

	// map
	fmt.Println("--------------------------")
	var dict1 map[string]int = map[string]int{ // string 类型的键，int 类型的值
		"one": 1,
		"two": 2,
		"ten": 10,
	}
	fmt.Println(dict1)

	var dict2 map[string]string
	fmt.Println(dict2 == nil) // 判断是否为 nil map --> true

	dict2 = make(map[string]string) // 使用 make 函数创建 map
	fmt.Println(dict2 == nil)       // false

	// 添加元素
	dict2["name"] = "Golang"
	dict2["type"] = "programming language"
	fmt.Println(dict2)

	// 访问和修改元素
	fmt.Println(dict1["two"]) // 访问键为 "two" 的值
	dict1["ten"] = 42         // 修改键为 "ten" 的值
	fmt.Println(dict1)

	// 删除元素
	delete(dict1, "one") // 删除键为 "one" 的键值对
	fmt.Println(dict1)

	// 检查键是否存在：不存在时，第一个参数为该类型的零值
	// 第二个参数为布尔值，表示键是否存在
	value, exists := dict1["three"]
	fmt.Println(value, exists) // 0 false

	value, exists = dict1["two"]
	fmt.Println(value, exists) // 2 true
}
