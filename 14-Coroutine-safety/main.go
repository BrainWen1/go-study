package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

// 同步锁
var global int
var mutex sync.Mutex

func add() {
	for i := 0; i < 10000; i++ {
		mutex.Lock()
		global++
		mutex.Unlock()
	}
	wait.Done()
}

func sub() {
	for i := 0; i < 10000; i++ {
		mutex.Lock()
		global--
		mutex.Unlock()
	}
	wait.Done()
}

// 线程安全的map
var mp = sync.Map{}

func reader() {
	for true {
		fmt.Println(mp.Load("key"))
	}
	wait.Done()
}

func writer() {
	for i := 0; true; i++ {
		mp.Store("key", i)
	}
	wait.Done()
}

func main() {
	for i := 0; i < 100; i++ {
		wait.Add(2)

		go add()
		go sub()

		wait.Wait()
		fmt.Print(global, " ")
	}

	wait.Add(2)
	go reader()
	go writer()
	wait.Wait()
}
