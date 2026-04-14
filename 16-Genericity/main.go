package main

import "fmt"

//
func add[T int | float64](a ...T) T {
	var sum T
	for _, val := range a {
		sum += val
	}
	return sum
}

func print[T any](a T) T {
	fmt.Println(a)
	return a
}

//
type Number interface {
	int | int64 | float32 | float64
}

func sum[T Number](a ...T) T {
	var total T
	for _, val := range a {
		total += val
	}
	return total
}

// 泛型结构体
type Pair[T any] struct {
	first  T
	second T
}

// 泛型切片
type List[T any] []T

// 泛型map
type Map[K int | string, V any] map[K]V

func main() {
	fmt.Println(add(1, 2, 3))       // 输出: 6
	fmt.Println(add(1.5, 2.5, 3.1)) // 输出: 7.1
	fmt.Println(sum(1, 2, 3))       // 输出: 6
	fmt.Println(sum(1.5, 2.5, 3.1)) // 输出: 7.1

	print("Hello, World!")
	print(666)
	print(3.14)
	print([]string{"Go", "Generics", "Example"})

	pair1 := Pair[int]{first: 1, second: 2}
	pair2 := Pair[string]{first: "Hello", second: "Go"}
	fmt.Printf("Pair1: %v, Pair2: %v\n", pair1, pair2)

	list1 := List[int]{1, 2, 3}
	list2 := List[string]{"Hello", "Go"}
	fmt.Printf("List1: %v, List2: %v\n", list1, list2)

	map1 := Map[int, string]{1: "Hello", 2: "Go"}
	map2 := Map[string, float64]{"Hello": 1.2, "Go": 3.14}
	fmt.Printf("Map1: %v, Map2: %v\n", map1, map2)
}
