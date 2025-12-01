package main

import (
	"fmt"
)

func main() {
	// if 语句
	a := 10
	if a > 5 {
		fmt.Println("a is greater than 5")
	}

	// if-else 语句
	a = 7
	if a%2 == 0 {
		fmt.Println("a is even")
	} else {
		fmt.Println("a is odd")
	}

	// if-else if-else 语句
	a = 0
	if a < 0 {
		fmt.Println("a is negative")
	} else if a == 0 {
		fmt.Println("a is zero")
	} else {
		fmt.Println("a is positive")
	}

	// 嵌套 if 语句
	a = 15
	if a > 10 {
		if a < 20 {
			fmt.Println("a is between 10 and 20")
		} else {
			fmt.Println("a is 20 or greater")
		}
	}

	// 多条件判断
	a = 25
	if a > 20 && a < 30 {
		fmt.Println("a is between 20 and 30")
	} else {
		fmt.Println("a is not between 20 and 30")
	}

	// 带初始化语句的 if
	if b := a * 2; b > 50 {
		fmt.Println("b is greater than 50")
	} else {
		fmt.Println("b is 50 or less")
	}

	// switch 语句
	var day int
	fmt.Print("Enter a day number (1-7): ")
	fmt.Scan(&day)

	switch day { // Go 语言中的 switch 默认不需要 break
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	switch day {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Println("It's a weekday")
	case 6:
		fallthrough
	case 7:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("Invalid day")
	}

	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	switch {
	case age >= 0 && age <= 18:
		fmt.Println("You are a minor")
	case age > 18 && age <= 65:
		fmt.Println("You are an adult")
	case age > 65:
		fmt.Println("You are a senior citizen")
	default:
		fmt.Println("Invalid age")
	}
}
