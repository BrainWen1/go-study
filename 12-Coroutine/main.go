package main

import (
	"fmt"
	"sync"
	"time"
)

// var wait int // 方法一

func Shopping(name string, wait *sync.WaitGroup) {
	fmt.Println("购买商品：", name)
	time.Sleep(1 * time.Second)
	fmt.Println("商品购买完成：", name)
	// wait-- // 方法一
	wait.Done() // 方法二：通知 WaitGroup 当前协程完成
}

func main() {
	start := time.Now()

	// Shopping("ZhangSan")
	// Shopping("LiSi")
	// Shopping("WangWu")

	// wait = 3 // 方法一
	var wait sync.WaitGroup // 方法二：Go 语言官方推荐使用 sync 包中的 WaitGroup 来实现协程同步
	wait.Add(3)             // 添加三个协程

	go Shopping("ZhangSan", &wait)
	go Shopping("LiSi", &wait)
	go Shopping("WangWu", &wait)

	wait.Wait() // 方法二：等待所有的协程完成

	// for { // 方法一
	// 	if wait <= 0 {
	// 		break
	// 	}
	// }

	fmt.Println("-------- 购物完成 --------", time.Since(start))
}
