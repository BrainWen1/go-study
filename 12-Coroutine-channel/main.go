package main

import (
	"fmt"
	"sync"
	"time"
)

// var wait int // 方法一

var moneyChan = make(chan int) // 声明并初始化一个长度为0的信道

func Shopping(name string, wait *sync.WaitGroup, money int) {
	fmt.Println("购买商品：", name)
	time.Sleep(1 * time.Second)
	fmt.Println("商品购买完成：", name)

	moneyChan <- money // 将购买商品花费的钱放入信道

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

	go Shopping("ZhangSan", &wait, 100)
	go Shopping("LiSi", &wait, 200)
	go Shopping("WangWu", &wait, 300)

	go func() {
		defer close(moneyChan)
		// 在协程函数里面等待上面三个协程函数结束
		wait.Wait()
	}()

	for {
		money, ok := <-moneyChan
		fmt.Println(money, ok)
		if !ok {
			break
		}
	}

	// for { // 方法一
	// 	if wait <= 0 {
	// 		break
	// 	}
	// }

	fmt.Println("-------- 购物完成 --------", time.Since(start))

	// channel
	var ch chan int = make(chan int, 1)
	ch <- 10          // 放入数据
	fmt.Println(<-ch) // 取出数据
	// fmt.Println(<-ch) // 再次取出数据，发生阻塞：fatal error: all goroutines are asleep - deadlock!

	ch <- 20
	fmt.Println(<-ch) // 取出数据

	// num, ok := <-ch // 从 channel 取数据，并判断 channel 是否关闭
	// if ok {
	// 	fmt.Println("取到数据：", num)
	// } else {
	// 	fmt.Println("channel 已关闭")
	// }

	// close(ch) // 关闭 channel

	// num, ok = <-ch // 再次从 channel 取数据
	// if ok {
	// 	fmt.Println("取到数据：", num)
	// } else {
	// 	fmt.Println("channel 已关闭")
	// }
}
