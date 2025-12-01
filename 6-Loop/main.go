package main

import (
	"fmt"
)

func main() {
	// for 循环
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("\nSum of first 100 natural numbers: %d\n", sum)

	// 无限循环
	count := 0
	for {
		if count >= 5 {
			break
		}
		count++
		fmt.Println("Count:", count)
	}

	// while 风格的 for 循环：只有条件部分
	n := 1
	for n < 10 {
		n *= 2
	}
	fmt.Printf("Final value of n: %d\n", n)

	// do-while 风格的 for 循环：先执行一次循环体，然后检查条件
	m := 1
	for {
		m *= 3
		if m >= 10 {
			break
		}
	}
	fmt.Printf("Final value of m: %d\n", m)

	// for-range 循环
	numbers := []int{1, 2, 3, 4, 5}
	for index, num := range numbers {
		fmt.Println(index, num)
	}

	// 拿到索引
	for index := range numbers {
		fmt.Print(index, " ")
	}
	fmt.Println()
	// 拿到值
	for _, num := range numbers {
		fmt.Print(num, " ")
	}
	fmt.Println()

	// for-range 遍历 map
	dict := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for key, value := range dict {
		fmt.Println(key, ":", value)
	}

	// break continue
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			continue // 跳过偶数
		}
		if i == 15 {
			break // 遇到15时退出循环
		}
		fmt.Print(i, " ")
	}
}
