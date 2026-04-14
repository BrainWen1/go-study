package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 一次性读取
	data, err := os.ReadFile("hello.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// 分片读
	file, err := os.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("------")
	buf := make([]byte, 5) // 每次读取 5 个字节
	for {
		n, err := file.Read(buf)
		if err != nil || err == io.EOF {
			break
		}
		fmt.Print(string(buf[:n]))
	}

	// 带缓冲读
	fmt.Println("\n------")
	file, err = os.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf1 := bufio.NewReader(file)
	for {
		line, _, err := buf1.ReadLine()
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	// 指定分隔符
	fmt.Println("------")
	file, err = os.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf2 := bufio.NewScanner(file)
	buf2.Split(bufio.ScanWords) // 按单词分割
	for buf2.Scan() {           // 读取每个单词并且判断是否读取成功
		fmt.Println(buf2.Text()) // 输出当前单词
	}
}
