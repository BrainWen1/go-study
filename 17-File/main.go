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

	// 一次性写
	err = os.WriteFile("write.txt", []byte("Writing some text..."), 0644)
	if err != nil {
		panic(err)
	}

	// 流式写
	file, err = os.OpenFile("write.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(" Adding more text...\n")
	if err != nil {
		panic(err)
	}
	writer.Flush()

	// 文件信息
	info, err := os.Stat("write.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File Name: %s\nSize: %d bytes\nPermissions: %s\n",
		info.Name(), info.Size(), info.Mode())

	// 重命名、删除文件
	file, err = os.Create("temp.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("temp.txt created !")
	}
	file.Close()

	err = os.Rename("temp.txt", "renamed.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("temp.txt renamed to renamed.txt !")
	}
	err = os.Remove("renamed.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("renamed.txt removed !")
	}

	// 文件复制
	src, _ := os.Open("hello.txt")
	defer src.Close()
	dst, _ := os.Create("copy.txt")
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}

	// 目录操作
	dir, _ := os.ReadDir("..") // 读取上级目录"go-study"
	for _, entry := range dir {
		info, _ := entry.Info()
		fmt.Println(entry.Name(), info.Size(), entry.IsDir())
	}
}
