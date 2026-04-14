package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int)
var nameChan = make(chan string)
var doneChan = make(chan struct{}) // 用于通知主协程所有购物完成

func shopping(name string, money int, wait *sync.WaitGroup) {
	fmt.Println("购买商品：", name)

	time.Sleep(1 * time.Second)
	moneyChan <- money
	nameChan <- name

	fmt.Println("商品购买完成：", name)
	wait.Done()
}

func main() {
	var wait sync.WaitGroup
	wait.Add(3)

	startTime := time.Now()
	// 三个协程执行
	go shopping("苹果", 10, &wait)
	go shopping("香蕉", 20, &wait)
	go shopping("橙子", 30, &wait)
	// 该协程等待上面三个协程执行完成后，关闭信道
	go func() {
		defer close(moneyChan)
		defer close(nameChan)
		defer close(doneChan)
		wait.Wait()
	}()

	moneyList := []int{}
	nameList := []string{}

	event := func() {
		for {
			select { // select 监听多个信道的事件
			case money := <-moneyChan:
				moneyList = append(moneyList, money)
			case name := <-nameChan:
				nameList = append(nameList, name)
			case <-doneChan:
				return
			case <-time.After(2 * time.Second):
				fmt.Println("等待超时，退出程序")
				return
			}
		}
	}
	event()

	fmt.Println("购物完成", time.Since(startTime))
	fmt.Println("购买的商品：", nameList)
	fmt.Println("花费的金额：", moneyList)
}
