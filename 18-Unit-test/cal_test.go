package main

import (
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"1+2", 1, 2, 3},
		{"-1+1", -1, 1, 0},
		{"0+0", 0, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Add(c.a, c.b)
			if got != c.want {
				t.Errorf("add(%v, %v) = %v; want %v", c.a, c.b, got, c.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"1-2", 1, 2, -1},
		{"-1-1", -1, 1, -2},
		{"0-0", 0, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Sub(c.a, c.b)
			if got != c.want {
				t.Errorf("sub(%v, %v) = %v; want %v", c.a, c.b, got, c.want)
			}
		})
	}
}

// TestMain 函数可以用于在测试运行前后执行一些 setup 和 teardown 的操作
func TestMain(m *testing.M) {
	// setup 代码
	println("Setup before tests")

	// 运行测试
	exitCode := m.Run()

	// teardown 代码
	println("Teardown after tests")
	// 退出
	os.Exit(exitCode)
}
